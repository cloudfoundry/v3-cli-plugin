// This file was generated by counterfeiter
package requirementsfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeApplicationRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns     struct {
		result1 error
	}
	GetApplicationStub        func() models.Application
	getApplicationMutex       sync.RWMutex
	getApplicationArgsForCall []struct{}
	getApplicationReturns     struct {
		result1 models.Application
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApplicationRequirement) Execute() error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.recordInvocation("Execute", []interface{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeApplicationRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeApplicationRequirement) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApplicationRequirement) GetApplication() models.Application {
	fake.getApplicationMutex.Lock()
	fake.getApplicationArgsForCall = append(fake.getApplicationArgsForCall, struct{}{})
	fake.recordInvocation("GetApplication", []interface{}{})
	fake.getApplicationMutex.Unlock()
	if fake.GetApplicationStub != nil {
		return fake.GetApplicationStub()
	} else {
		return fake.getApplicationReturns.result1
	}
}

func (fake *FakeApplicationRequirement) GetApplicationCallCount() int {
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	return len(fake.getApplicationArgsForCall)
}

func (fake *FakeApplicationRequirement) GetApplicationReturns(result1 models.Application) {
	fake.GetApplicationStub = nil
	fake.getApplicationReturns = struct {
		result1 models.Application
	}{result1}
}

func (fake *FakeApplicationRequirement) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeApplicationRequirement) recordInvocation(key string, args []interface{}) {
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

var _ requirements.ApplicationRequirement = new(FakeApplicationRequirement)