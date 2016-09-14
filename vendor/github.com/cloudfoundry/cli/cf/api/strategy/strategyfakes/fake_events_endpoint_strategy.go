// This file was generated by counterfeiter
package strategyfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/resources"
	"github.com/cloudfoundry/cli/cf/api/strategy"
)

type FakeEventsEndpointStrategy struct {
	EventsURLStub        func(appGUID string, limit int64) string
	eventsURLMutex       sync.RWMutex
	eventsURLArgsForCall []struct {
		appGUID string
		limit   int64
	}
	eventsURLReturns struct {
		result1 string
	}
	EventsResourceStub        func() resources.EventResource
	eventsResourceMutex       sync.RWMutex
	eventsResourceArgsForCall []struct{}
	eventsResourceReturns     struct {
		result1 resources.EventResource
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventsEndpointStrategy) EventsURL(appGUID string, limit int64) string {
	fake.eventsURLMutex.Lock()
	fake.eventsURLArgsForCall = append(fake.eventsURLArgsForCall, struct {
		appGUID string
		limit   int64
	}{appGUID, limit})
	fake.recordInvocation("EventsURL", []interface{}{appGUID, limit})
	fake.eventsURLMutex.Unlock()
	if fake.EventsURLStub != nil {
		return fake.EventsURLStub(appGUID, limit)
	} else {
		return fake.eventsURLReturns.result1
	}
}

func (fake *FakeEventsEndpointStrategy) EventsURLCallCount() int {
	fake.eventsURLMutex.RLock()
	defer fake.eventsURLMutex.RUnlock()
	return len(fake.eventsURLArgsForCall)
}

func (fake *FakeEventsEndpointStrategy) EventsURLArgsForCall(i int) (string, int64) {
	fake.eventsURLMutex.RLock()
	defer fake.eventsURLMutex.RUnlock()
	return fake.eventsURLArgsForCall[i].appGUID, fake.eventsURLArgsForCall[i].limit
}

func (fake *FakeEventsEndpointStrategy) EventsURLReturns(result1 string) {
	fake.EventsURLStub = nil
	fake.eventsURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventsEndpointStrategy) EventsResource() resources.EventResource {
	fake.eventsResourceMutex.Lock()
	fake.eventsResourceArgsForCall = append(fake.eventsResourceArgsForCall, struct{}{})
	fake.recordInvocation("EventsResource", []interface{}{})
	fake.eventsResourceMutex.Unlock()
	if fake.EventsResourceStub != nil {
		return fake.EventsResourceStub()
	} else {
		return fake.eventsResourceReturns.result1
	}
}

func (fake *FakeEventsEndpointStrategy) EventsResourceCallCount() int {
	fake.eventsResourceMutex.RLock()
	defer fake.eventsResourceMutex.RUnlock()
	return len(fake.eventsResourceArgsForCall)
}

func (fake *FakeEventsEndpointStrategy) EventsResourceReturns(result1 resources.EventResource) {
	fake.EventsResourceStub = nil
	fake.eventsResourceReturns = struct {
		result1 resources.EventResource
	}{result1}
}

func (fake *FakeEventsEndpointStrategy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.eventsURLMutex.RLock()
	defer fake.eventsURLMutex.RUnlock()
	fake.eventsResourceMutex.RLock()
	defer fake.eventsResourceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEventsEndpointStrategy) recordInvocation(key string, args []interface{}) {
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

var _ strategy.EventsEndpointStrategy = new(FakeEventsEndpointStrategy)