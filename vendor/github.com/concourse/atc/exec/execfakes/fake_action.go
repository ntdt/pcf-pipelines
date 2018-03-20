// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/exec"
	"github.com/concourse/atc/worker"
)

type FakeAction struct {
	RunStub        func(context.Context, lager.Logger, *worker.ArtifactRepository) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *worker.ArtifactRepository
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	ExitStatusStub        func() exec.ExitStatus
	exitStatusMutex       sync.RWMutex
	exitStatusArgsForCall []struct{}
	exitStatusReturns     struct {
		result1 exec.ExitStatus
	}
	exitStatusReturnsOnCall map[int]struct {
		result1 exec.ExitStatus
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAction) Run(arg1 context.Context, arg2 lager.Logger, arg3 *worker.ArtifactRepository) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *worker.ArtifactRepository
	}{arg1, arg2, arg3})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *FakeAction) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeAction) RunArgsForCall(i int) (context.Context, lager.Logger, *worker.ArtifactRepository) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2, fake.runArgsForCall[i].arg3
}

func (fake *FakeAction) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAction) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAction) ExitStatus() exec.ExitStatus {
	fake.exitStatusMutex.Lock()
	ret, specificReturn := fake.exitStatusReturnsOnCall[len(fake.exitStatusArgsForCall)]
	fake.exitStatusArgsForCall = append(fake.exitStatusArgsForCall, struct{}{})
	fake.recordInvocation("ExitStatus", []interface{}{})
	fake.exitStatusMutex.Unlock()
	if fake.ExitStatusStub != nil {
		return fake.ExitStatusStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.exitStatusReturns.result1
}

func (fake *FakeAction) ExitStatusCallCount() int {
	fake.exitStatusMutex.RLock()
	defer fake.exitStatusMutex.RUnlock()
	return len(fake.exitStatusArgsForCall)
}

func (fake *FakeAction) ExitStatusReturns(result1 exec.ExitStatus) {
	fake.ExitStatusStub = nil
	fake.exitStatusReturns = struct {
		result1 exec.ExitStatus
	}{result1}
}

func (fake *FakeAction) ExitStatusReturnsOnCall(i int, result1 exec.ExitStatus) {
	fake.ExitStatusStub = nil
	if fake.exitStatusReturnsOnCall == nil {
		fake.exitStatusReturnsOnCall = make(map[int]struct {
			result1 exec.ExitStatus
		})
	}
	fake.exitStatusReturnsOnCall[i] = struct {
		result1 exec.ExitStatus
	}{result1}
}

func (fake *FakeAction) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.exitStatusMutex.RLock()
	defer fake.exitStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAction) recordInvocation(key string, args []interface{}) {
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

var _ exec.Action = new(FakeAction)