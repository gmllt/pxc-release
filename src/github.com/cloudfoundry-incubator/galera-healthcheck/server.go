package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudfoundry-incubator/galera-healthcheck/config"
	"github.com/cloudfoundry-incubator/galera-healthcheck/sequence_number"

	"github.com/cloudfoundry-incubator/galera-healthcheck/healthcheck"
	"github.com/pivotal-golang/lager"

	_ "github.com/go-sql-driver/mysql"
)

var healthchecker *healthcheck.Healthchecker
var sequence_number_checker *sequence_number.SequenceNumberchecker

func handler(w http.ResponseWriter, r *http.Request, logger lager.Logger) {
	result, msg := healthchecker.Check()
	if result {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	body := fmt.Sprintf("Galera Cluster Node Status: %s", msg)
	fmt.Fprint(w, body)

	logger.Debug(fmt.Sprintf("Healhcheck Response Body: %s", body))
}

func main() {

	rootConfig, err := config.NewConfig(os.Args)

	logger := rootConfig.Logger()

	err = rootConfig.Validate()
	if err != nil {
		logger.Fatal("Failed to validate config", err)
	}

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/",
			rootConfig.DB.User,
			rootConfig.DB.Password,
			rootConfig.DB.Host,
			rootConfig.DB.Port))
	if err != nil {
		// sql.Open may not actually check that the DB is reachable
		err = db.Ping()
	}
	if err != nil {
		logger.Fatal("Failed to open DB connection", err, lager.Data{
			"dbHost": rootConfig.DB.Host,
			"dbPort": rootConfig.DB.Port,
			"dbUser": rootConfig.DB.User,
		})
	}

	logger.Info("Opened DB connection", lager.Data{
		"dbHost": rootConfig.DB.Host,
		"dbPort": rootConfig.DB.Port,
		"dbUser": rootConfig.DB.User,
	})

	healthchecker = healthcheck.New(db, config, logger)
	StatusEndpointHandler := fmt.Sprintf("/%s", *statusEndpoint)
	http.HandleFunc(StatusEndpointHandler, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, logger)
	})

	sequence_number_checker = sequence_number.New(db, config, logger)
	http.HandleFunc("sequence_number", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, logger)
	})

	address := fmt.Sprintf("%s:%d", rootConfig.Host, rootConfig.Port)
	url := fmt.Sprintf("http://%s/", address)
	logger.Info("Serving healthcheck endpoint", lager.Data{
		"url": url,
	})

	go func() {
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		logger.Info("Attempting to GET endpoint...", lager.Data{
			"url": url,
		})

		var resp *http.Response
		retryAttemptsRemaining := 3
		for ; retryAttemptsRemaining > 0; retryAttemptsRemaining-- {
			resp, err = client.Get(url)
			if err != nil {
				logger.Info("GET endpoint failed, retrying...", lager.Data{
					"url": url,
					"err": err,
				})
				time.Sleep(time.Second * 10)
			} else {
				break
			}
		}
		if retryAttemptsRemaining == 0 {
			logger.Fatal(
				"Initialization failed: Coudn't GET endpoint",
				err,
				lager.Data{
					"url":     url,
					"retries": retryAttemptsRemaining,
				})
		}
		logger.Info("GET endpoint succeeded, now accepting connections", lager.Data{
			"url":        url,
			"statusCode": resp.StatusCode,
		})

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Fatal("Initialization failed: reading response body", err, lager.Data{
				"url":         url,
				"status-code": resp.StatusCode,
			})
		}
		logger.Info(fmt.Sprintf("Initial Response: %s", body))

		// existence of pid file means the server is running
		pid := os.Getpid()
		err = ioutil.WriteFile(rootConfig.PidFile, []byte(strconv.Itoa(os.Getpid())), 0644)
		if err != nil {
			logger.Fatal("Failed to write pid file", err, lager.Data{
				"pid":     pid,
				"pidFile": rootConfig.PidFile,
			})
		}

		// Used by tests to deterministically know that the healthcheck is accepting incoming connections
		logger.Info("Healthcheck Started")
	}()

	server := &http.Server{Addr: address}
	server.ListenAndServe()
}
