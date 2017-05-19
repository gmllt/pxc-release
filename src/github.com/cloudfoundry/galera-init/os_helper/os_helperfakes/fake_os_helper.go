// This file was generated by counterfeiter
package os_helperfakes

import (
	"os/exec"
	"sync"
	"time"

	"github.com/cloudfoundry/mariadb_ctrl/os_helper"
)

type FakeOsHelper struct {
	RunCommandStub        func(executable string, args ...string) (string, error)
	runCommandMutex       sync.RWMutex
	runCommandArgsForCall []struct {
		executable string
		args       []string
	}
	runCommandReturns struct {
		result1 string
		result2 error
	}
	runCommandReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RunCommandWithTimeoutStub        func(timeout int, logFileName string, executable string, args ...string) error
	runCommandWithTimeoutMutex       sync.RWMutex
	runCommandWithTimeoutArgsForCall []struct {
		timeout     int
		logFileName string
		executable  string
		args        []string
	}
	runCommandWithTimeoutReturns struct {
		result1 error
	}
	runCommandWithTimeoutReturnsOnCall map[int]struct {
		result1 error
	}
	StartCommandStub        func(logFileName string, executable string, args ...string) (*exec.Cmd, error)
	startCommandMutex       sync.RWMutex
	startCommandArgsForCall []struct {
		logFileName string
		executable  string
		args        []string
	}
	startCommandReturns struct {
		result1 *exec.Cmd
		result2 error
	}
	startCommandReturnsOnCall map[int]struct {
		result1 *exec.Cmd
		result2 error
	}
	WaitForCommandStub        func(cmd *exec.Cmd) chan error
	waitForCommandMutex       sync.RWMutex
	waitForCommandArgsForCall []struct {
		cmd *exec.Cmd
	}
	waitForCommandReturns struct {
		result1 chan error
	}
	waitForCommandReturnsOnCall map[int]struct {
		result1 chan error
	}
	FileExistsStub        func(filename string) bool
	fileExistsMutex       sync.RWMutex
	fileExistsArgsForCall []struct {
		filename string
	}
	fileExistsReturns struct {
		result1 bool
	}
	fileExistsReturnsOnCall map[int]struct {
		result1 bool
	}
	ReadFileStub        func(filename string) (string, error)
	readFileMutex       sync.RWMutex
	readFileArgsForCall []struct {
		filename string
	}
	readFileReturns struct {
		result1 string
		result2 error
	}
	readFileReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	WriteStringToFileStub        func(filename string, contents string) error
	writeStringToFileMutex       sync.RWMutex
	writeStringToFileArgsForCall []struct {
		filename string
		contents string
	}
	writeStringToFileReturns struct {
		result1 error
	}
	writeStringToFileReturnsOnCall map[int]struct {
		result1 error
	}
	SleepStub        func(duration time.Duration)
	sleepMutex       sync.RWMutex
	sleepArgsForCall []struct {
		duration time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOsHelper) RunCommand(executable string, args ...string) (string, error) {
	fake.runCommandMutex.Lock()
	ret, specificReturn := fake.runCommandReturnsOnCall[len(fake.runCommandArgsForCall)]
	fake.runCommandArgsForCall = append(fake.runCommandArgsForCall, struct {
		executable string
		args       []string
	}{executable, args})
	fake.recordInvocation("RunCommand", []interface{}{executable, args})
	fake.runCommandMutex.Unlock()
	if fake.RunCommandStub != nil {
		return fake.RunCommandStub(executable, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runCommandReturns.result1, fake.runCommandReturns.result2
}

func (fake *FakeOsHelper) RunCommandCallCount() int {
	fake.runCommandMutex.RLock()
	defer fake.runCommandMutex.RUnlock()
	return len(fake.runCommandArgsForCall)
}

func (fake *FakeOsHelper) RunCommandArgsForCall(i int) (string, []string) {
	fake.runCommandMutex.RLock()
	defer fake.runCommandMutex.RUnlock()
	return fake.runCommandArgsForCall[i].executable, fake.runCommandArgsForCall[i].args
}

func (fake *FakeOsHelper) RunCommandReturns(result1 string, result2 error) {
	fake.RunCommandStub = nil
	fake.runCommandReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) RunCommandReturnsOnCall(i int, result1 string, result2 error) {
	fake.RunCommandStub = nil
	if fake.runCommandReturnsOnCall == nil {
		fake.runCommandReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.runCommandReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) RunCommandWithTimeout(timeout int, logFileName string, executable string, args ...string) error {
	fake.runCommandWithTimeoutMutex.Lock()
	ret, specificReturn := fake.runCommandWithTimeoutReturnsOnCall[len(fake.runCommandWithTimeoutArgsForCall)]
	fake.runCommandWithTimeoutArgsForCall = append(fake.runCommandWithTimeoutArgsForCall, struct {
		timeout     int
		logFileName string
		executable  string
		args        []string
	}{timeout, logFileName, executable, args})
	fake.recordInvocation("RunCommandWithTimeout", []interface{}{timeout, logFileName, executable, args})
	fake.runCommandWithTimeoutMutex.Unlock()
	if fake.RunCommandWithTimeoutStub != nil {
		return fake.RunCommandWithTimeoutStub(timeout, logFileName, executable, args...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runCommandWithTimeoutReturns.result1
}

func (fake *FakeOsHelper) RunCommandWithTimeoutCallCount() int {
	fake.runCommandWithTimeoutMutex.RLock()
	defer fake.runCommandWithTimeoutMutex.RUnlock()
	return len(fake.runCommandWithTimeoutArgsForCall)
}

func (fake *FakeOsHelper) RunCommandWithTimeoutArgsForCall(i int) (int, string, string, []string) {
	fake.runCommandWithTimeoutMutex.RLock()
	defer fake.runCommandWithTimeoutMutex.RUnlock()
	return fake.runCommandWithTimeoutArgsForCall[i].timeout, fake.runCommandWithTimeoutArgsForCall[i].logFileName, fake.runCommandWithTimeoutArgsForCall[i].executable, fake.runCommandWithTimeoutArgsForCall[i].args
}

func (fake *FakeOsHelper) RunCommandWithTimeoutReturns(result1 error) {
	fake.RunCommandWithTimeoutStub = nil
	fake.runCommandWithTimeoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsHelper) RunCommandWithTimeoutReturnsOnCall(i int, result1 error) {
	fake.RunCommandWithTimeoutStub = nil
	if fake.runCommandWithTimeoutReturnsOnCall == nil {
		fake.runCommandWithTimeoutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runCommandWithTimeoutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsHelper) StartCommand(logFileName string, executable string, args ...string) (*exec.Cmd, error) {
	fake.startCommandMutex.Lock()
	ret, specificReturn := fake.startCommandReturnsOnCall[len(fake.startCommandArgsForCall)]
	fake.startCommandArgsForCall = append(fake.startCommandArgsForCall, struct {
		logFileName string
		executable  string
		args        []string
	}{logFileName, executable, args})
	fake.recordInvocation("StartCommand", []interface{}{logFileName, executable, args})
	fake.startCommandMutex.Unlock()
	if fake.StartCommandStub != nil {
		return fake.StartCommandStub(logFileName, executable, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startCommandReturns.result1, fake.startCommandReturns.result2
}

func (fake *FakeOsHelper) StartCommandCallCount() int {
	fake.startCommandMutex.RLock()
	defer fake.startCommandMutex.RUnlock()
	return len(fake.startCommandArgsForCall)
}

func (fake *FakeOsHelper) StartCommandArgsForCall(i int) (string, string, []string) {
	fake.startCommandMutex.RLock()
	defer fake.startCommandMutex.RUnlock()
	return fake.startCommandArgsForCall[i].logFileName, fake.startCommandArgsForCall[i].executable, fake.startCommandArgsForCall[i].args
}

func (fake *FakeOsHelper) StartCommandReturns(result1 *exec.Cmd, result2 error) {
	fake.StartCommandStub = nil
	fake.startCommandReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) StartCommandReturnsOnCall(i int, result1 *exec.Cmd, result2 error) {
	fake.StartCommandStub = nil
	if fake.startCommandReturnsOnCall == nil {
		fake.startCommandReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
			result2 error
		})
	}
	fake.startCommandReturnsOnCall[i] = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) WaitForCommand(cmd *exec.Cmd) chan error {
	fake.waitForCommandMutex.Lock()
	ret, specificReturn := fake.waitForCommandReturnsOnCall[len(fake.waitForCommandArgsForCall)]
	fake.waitForCommandArgsForCall = append(fake.waitForCommandArgsForCall, struct {
		cmd *exec.Cmd
	}{cmd})
	fake.recordInvocation("WaitForCommand", []interface{}{cmd})
	fake.waitForCommandMutex.Unlock()
	if fake.WaitForCommandStub != nil {
		return fake.WaitForCommandStub(cmd)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.waitForCommandReturns.result1
}

func (fake *FakeOsHelper) WaitForCommandCallCount() int {
	fake.waitForCommandMutex.RLock()
	defer fake.waitForCommandMutex.RUnlock()
	return len(fake.waitForCommandArgsForCall)
}

func (fake *FakeOsHelper) WaitForCommandArgsForCall(i int) *exec.Cmd {
	fake.waitForCommandMutex.RLock()
	defer fake.waitForCommandMutex.RUnlock()
	return fake.waitForCommandArgsForCall[i].cmd
}

func (fake *FakeOsHelper) WaitForCommandReturns(result1 chan error) {
	fake.WaitForCommandStub = nil
	fake.waitForCommandReturns = struct {
		result1 chan error
	}{result1}
}

func (fake *FakeOsHelper) WaitForCommandReturnsOnCall(i int, result1 chan error) {
	fake.WaitForCommandStub = nil
	if fake.waitForCommandReturnsOnCall == nil {
		fake.waitForCommandReturnsOnCall = make(map[int]struct {
			result1 chan error
		})
	}
	fake.waitForCommandReturnsOnCall[i] = struct {
		result1 chan error
	}{result1}
}

func (fake *FakeOsHelper) FileExists(filename string) bool {
	fake.fileExistsMutex.Lock()
	ret, specificReturn := fake.fileExistsReturnsOnCall[len(fake.fileExistsArgsForCall)]
	fake.fileExistsArgsForCall = append(fake.fileExistsArgsForCall, struct {
		filename string
	}{filename})
	fake.recordInvocation("FileExists", []interface{}{filename})
	fake.fileExistsMutex.Unlock()
	if fake.FileExistsStub != nil {
		return fake.FileExistsStub(filename)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.fileExistsReturns.result1
}

func (fake *FakeOsHelper) FileExistsCallCount() int {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	return len(fake.fileExistsArgsForCall)
}

func (fake *FakeOsHelper) FileExistsArgsForCall(i int) string {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	return fake.fileExistsArgsForCall[i].filename
}

func (fake *FakeOsHelper) FileExistsReturns(result1 bool) {
	fake.FileExistsStub = nil
	fake.fileExistsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOsHelper) FileExistsReturnsOnCall(i int, result1 bool) {
	fake.FileExistsStub = nil
	if fake.fileExistsReturnsOnCall == nil {
		fake.fileExistsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.fileExistsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOsHelper) ReadFile(filename string) (string, error) {
	fake.readFileMutex.Lock()
	ret, specificReturn := fake.readFileReturnsOnCall[len(fake.readFileArgsForCall)]
	fake.readFileArgsForCall = append(fake.readFileArgsForCall, struct {
		filename string
	}{filename})
	fake.recordInvocation("ReadFile", []interface{}{filename})
	fake.readFileMutex.Unlock()
	if fake.ReadFileStub != nil {
		return fake.ReadFileStub(filename)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readFileReturns.result1, fake.readFileReturns.result2
}

func (fake *FakeOsHelper) ReadFileCallCount() int {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return len(fake.readFileArgsForCall)
}

func (fake *FakeOsHelper) ReadFileArgsForCall(i int) string {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return fake.readFileArgsForCall[i].filename
}

func (fake *FakeOsHelper) ReadFileReturns(result1 string, result2 error) {
	fake.ReadFileStub = nil
	fake.readFileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) ReadFileReturnsOnCall(i int, result1 string, result2 error) {
	fake.ReadFileStub = nil
	if fake.readFileReturnsOnCall == nil {
		fake.readFileReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.readFileReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) WriteStringToFile(filename string, contents string) error {
	fake.writeStringToFileMutex.Lock()
	ret, specificReturn := fake.writeStringToFileReturnsOnCall[len(fake.writeStringToFileArgsForCall)]
	fake.writeStringToFileArgsForCall = append(fake.writeStringToFileArgsForCall, struct {
		filename string
		contents string
	}{filename, contents})
	fake.recordInvocation("WriteStringToFile", []interface{}{filename, contents})
	fake.writeStringToFileMutex.Unlock()
	if fake.WriteStringToFileStub != nil {
		return fake.WriteStringToFileStub(filename, contents)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writeStringToFileReturns.result1
}

func (fake *FakeOsHelper) WriteStringToFileCallCount() int {
	fake.writeStringToFileMutex.RLock()
	defer fake.writeStringToFileMutex.RUnlock()
	return len(fake.writeStringToFileArgsForCall)
}

func (fake *FakeOsHelper) WriteStringToFileArgsForCall(i int) (string, string) {
	fake.writeStringToFileMutex.RLock()
	defer fake.writeStringToFileMutex.RUnlock()
	return fake.writeStringToFileArgsForCall[i].filename, fake.writeStringToFileArgsForCall[i].contents
}

func (fake *FakeOsHelper) WriteStringToFileReturns(result1 error) {
	fake.WriteStringToFileStub = nil
	fake.writeStringToFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsHelper) WriteStringToFileReturnsOnCall(i int, result1 error) {
	fake.WriteStringToFileStub = nil
	if fake.writeStringToFileReturnsOnCall == nil {
		fake.writeStringToFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeStringToFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsHelper) Sleep(duration time.Duration) {
	fake.sleepMutex.Lock()
	fake.sleepArgsForCall = append(fake.sleepArgsForCall, struct {
		duration time.Duration
	}{duration})
	fake.recordInvocation("Sleep", []interface{}{duration})
	fake.sleepMutex.Unlock()
	if fake.SleepStub != nil {
		fake.SleepStub(duration)
	}
}

func (fake *FakeOsHelper) SleepCallCount() int {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return len(fake.sleepArgsForCall)
}

func (fake *FakeOsHelper) SleepArgsForCall(i int) time.Duration {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return fake.sleepArgsForCall[i].duration
}

func (fake *FakeOsHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runCommandMutex.RLock()
	defer fake.runCommandMutex.RUnlock()
	fake.runCommandWithTimeoutMutex.RLock()
	defer fake.runCommandWithTimeoutMutex.RUnlock()
	fake.startCommandMutex.RLock()
	defer fake.startCommandMutex.RUnlock()
	fake.waitForCommandMutex.RLock()
	defer fake.waitForCommandMutex.RUnlock()
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	fake.writeStringToFileMutex.RLock()
	defer fake.writeStringToFileMutex.RUnlock()
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOsHelper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ os_helper.OsHelper = new(FakeOsHelper)
