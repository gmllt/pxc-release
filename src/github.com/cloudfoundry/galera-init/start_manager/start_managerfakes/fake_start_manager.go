// This file was generated by counterfeiter
package start_managerfakes

import (
	"os/exec"
	"sync"

	"github.com/cloudfoundry/mariadb_ctrl/start_manager"
)

type FakeStartManager struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns     struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	GetMysqlCmdStub        func() (*exec.Cmd, error)
	getMysqlCmdMutex       sync.RWMutex
	getMysqlCmdArgsForCall []struct{}
	getMysqlCmdReturns     struct {
		result1 *exec.Cmd
		result2 error
	}
	getMysqlCmdReturnsOnCall map[int]struct {
		result1 *exec.Cmd
		result2 error
	}
	ShutdownStub        func() error
	shutdownMutex       sync.RWMutex
	shutdownArgsForCall []struct{}
	shutdownReturns     struct {
		result1 error
	}
	shutdownReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStartManager) Execute() error {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.recordInvocation("Execute", []interface{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.executeReturns.result1
}

func (fake *FakeStartManager) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeStartManager) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStartManager) ExecuteReturnsOnCall(i int, result1 error) {
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStartManager) GetMysqlCmd() (*exec.Cmd, error) {
	fake.getMysqlCmdMutex.Lock()
	ret, specificReturn := fake.getMysqlCmdReturnsOnCall[len(fake.getMysqlCmdArgsForCall)]
	fake.getMysqlCmdArgsForCall = append(fake.getMysqlCmdArgsForCall, struct{}{})
	fake.recordInvocation("GetMysqlCmd", []interface{}{})
	fake.getMysqlCmdMutex.Unlock()
	if fake.GetMysqlCmdStub != nil {
		return fake.GetMysqlCmdStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getMysqlCmdReturns.result1, fake.getMysqlCmdReturns.result2
}

func (fake *FakeStartManager) GetMysqlCmdCallCount() int {
	fake.getMysqlCmdMutex.RLock()
	defer fake.getMysqlCmdMutex.RUnlock()
	return len(fake.getMysqlCmdArgsForCall)
}

func (fake *FakeStartManager) GetMysqlCmdReturns(result1 *exec.Cmd, result2 error) {
	fake.GetMysqlCmdStub = nil
	fake.getMysqlCmdReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeStartManager) GetMysqlCmdReturnsOnCall(i int, result1 *exec.Cmd, result2 error) {
	fake.GetMysqlCmdStub = nil
	if fake.getMysqlCmdReturnsOnCall == nil {
		fake.getMysqlCmdReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
			result2 error
		})
	}
	fake.getMysqlCmdReturnsOnCall[i] = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeStartManager) Shutdown() error {
	fake.shutdownMutex.Lock()
	ret, specificReturn := fake.shutdownReturnsOnCall[len(fake.shutdownArgsForCall)]
	fake.shutdownArgsForCall = append(fake.shutdownArgsForCall, struct{}{})
	fake.recordInvocation("Shutdown", []interface{}{})
	fake.shutdownMutex.Unlock()
	if fake.ShutdownStub != nil {
		return fake.ShutdownStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.shutdownReturns.result1
}

func (fake *FakeStartManager) ShutdownCallCount() int {
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	return len(fake.shutdownArgsForCall)
}

func (fake *FakeStartManager) ShutdownReturns(result1 error) {
	fake.ShutdownStub = nil
	fake.shutdownReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStartManager) ShutdownReturnsOnCall(i int, result1 error) {
	fake.ShutdownStub = nil
	if fake.shutdownReturnsOnCall == nil {
		fake.shutdownReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.shutdownReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStartManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.getMysqlCmdMutex.RLock()
	defer fake.getMysqlCmdMutex.RUnlock()
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStartManager) recordInvocation(key string, args []interface{}) {
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

var _ start_manager.StartManager = new(FakeStartManager)
