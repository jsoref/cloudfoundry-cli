// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeStagingEnvironmentVariableGroupActor struct {
	GetEnvironmentVariableGroupStub        func(constant.EnvironmentVariableGroupName) (v7action.EnvironmentVariableGroup, v7action.Warnings, error)
	getEnvironmentVariableGroupMutex       sync.RWMutex
	getEnvironmentVariableGroupArgsForCall []struct {
		arg1 constant.EnvironmentVariableGroupName
	}
	getEnvironmentVariableGroupReturns struct {
		result1 v7action.EnvironmentVariableGroup
		result2 v7action.Warnings
		result3 error
	}
	getEnvironmentVariableGroupReturnsOnCall map[int]struct {
		result1 v7action.EnvironmentVariableGroup
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStagingEnvironmentVariableGroupActor) GetEnvironmentVariableGroup(arg1 constant.EnvironmentVariableGroupName) (v7action.EnvironmentVariableGroup, v7action.Warnings, error) {
	fake.getEnvironmentVariableGroupMutex.Lock()
	ret, specificReturn := fake.getEnvironmentVariableGroupReturnsOnCall[len(fake.getEnvironmentVariableGroupArgsForCall)]
	fake.getEnvironmentVariableGroupArgsForCall = append(fake.getEnvironmentVariableGroupArgsForCall, struct {
		arg1 constant.EnvironmentVariableGroupName
	}{arg1})
	fake.recordInvocation("GetEnvironmentVariableGroup", []interface{}{arg1})
	fake.getEnvironmentVariableGroupMutex.Unlock()
	if fake.GetEnvironmentVariableGroupStub != nil {
		return fake.GetEnvironmentVariableGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getEnvironmentVariableGroupReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStagingEnvironmentVariableGroupActor) GetEnvironmentVariableGroupCallCount() int {
	fake.getEnvironmentVariableGroupMutex.RLock()
	defer fake.getEnvironmentVariableGroupMutex.RUnlock()
	return len(fake.getEnvironmentVariableGroupArgsForCall)
}

func (fake *FakeStagingEnvironmentVariableGroupActor) GetEnvironmentVariableGroupCalls(stub func(constant.EnvironmentVariableGroupName) (v7action.EnvironmentVariableGroup, v7action.Warnings, error)) {
	fake.getEnvironmentVariableGroupMutex.Lock()
	defer fake.getEnvironmentVariableGroupMutex.Unlock()
	fake.GetEnvironmentVariableGroupStub = stub
}

func (fake *FakeStagingEnvironmentVariableGroupActor) GetEnvironmentVariableGroupArgsForCall(i int) constant.EnvironmentVariableGroupName {
	fake.getEnvironmentVariableGroupMutex.RLock()
	defer fake.getEnvironmentVariableGroupMutex.RUnlock()
	argsForCall := fake.getEnvironmentVariableGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStagingEnvironmentVariableGroupActor) GetEnvironmentVariableGroupReturns(result1 v7action.EnvironmentVariableGroup, result2 v7action.Warnings, result3 error) {
	fake.getEnvironmentVariableGroupMutex.Lock()
	defer fake.getEnvironmentVariableGroupMutex.Unlock()
	fake.GetEnvironmentVariableGroupStub = nil
	fake.getEnvironmentVariableGroupReturns = struct {
		result1 v7action.EnvironmentVariableGroup
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStagingEnvironmentVariableGroupActor) GetEnvironmentVariableGroupReturnsOnCall(i int, result1 v7action.EnvironmentVariableGroup, result2 v7action.Warnings, result3 error) {
	fake.getEnvironmentVariableGroupMutex.Lock()
	defer fake.getEnvironmentVariableGroupMutex.Unlock()
	fake.GetEnvironmentVariableGroupStub = nil
	if fake.getEnvironmentVariableGroupReturnsOnCall == nil {
		fake.getEnvironmentVariableGroupReturnsOnCall = make(map[int]struct {
			result1 v7action.EnvironmentVariableGroup
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getEnvironmentVariableGroupReturnsOnCall[i] = struct {
		result1 v7action.EnvironmentVariableGroup
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStagingEnvironmentVariableGroupActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEnvironmentVariableGroupMutex.RLock()
	defer fake.getEnvironmentVariableGroupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStagingEnvironmentVariableGroupActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.StagingEnvironmentVariableGroupActor = new(FakeStagingEnvironmentVariableGroupActor)