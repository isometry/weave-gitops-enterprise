// Code generated by counterfeiter. DO NOT EDIT.
package storefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/internal/models"
	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/store"
)

type FakeStore struct {
	CountObjectsStub        func(context.Context, string) (int64, error)
	countObjectsMutex       sync.RWMutex
	countObjectsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	countObjectsReturns struct {
		result1 int64
		result2 error
	}
	countObjectsReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	DeleteObjectStub        func(context.Context, models.Object) error
	deleteObjectMutex       sync.RWMutex
	deleteObjectArgsForCall []struct {
		arg1 context.Context
		arg2 models.Object
	}
	deleteObjectReturns struct {
		result1 error
	}
	deleteObjectReturnsOnCall map[int]struct {
		result1 error
	}
	GetAccessRulesStub        func() ([]models.AccessRule, error)
	getAccessRulesMutex       sync.RWMutex
	getAccessRulesArgsForCall []struct {
	}
	getAccessRulesReturns struct {
		result1 []models.AccessRule
		result2 error
	}
	getAccessRulesReturnsOnCall map[int]struct {
		result1 []models.AccessRule
		result2 error
	}
	GetObjectsStub        func() ([]models.Object, error)
	getObjectsMutex       sync.RWMutex
	getObjectsArgsForCall []struct {
	}
	getObjectsReturns struct {
		result1 []models.Object
		result2 error
	}
	getObjectsReturnsOnCall map[int]struct {
		result1 []models.Object
		result2 error
	}
	StoreAccessRulesStub        func(context.Context, []models.AccessRule) error
	storeAccessRulesMutex       sync.RWMutex
	storeAccessRulesArgsForCall []struct {
		arg1 context.Context
		arg2 []models.AccessRule
	}
	storeAccessRulesReturns struct {
		result1 error
	}
	storeAccessRulesReturnsOnCall map[int]struct {
		result1 error
	}
	StoreObjectsStub        func(context.Context, []models.Object) error
	storeObjectsMutex       sync.RWMutex
	storeObjectsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Object
	}
	storeObjectsReturns struct {
		result1 error
	}
	storeObjectsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) CountObjects(arg1 context.Context, arg2 string) (int64, error) {
	fake.countObjectsMutex.Lock()
	ret, specificReturn := fake.countObjectsReturnsOnCall[len(fake.countObjectsArgsForCall)]
	fake.countObjectsArgsForCall = append(fake.countObjectsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.CountObjectsStub
	fakeReturns := fake.countObjectsReturns
	fake.recordInvocation("CountObjects", []interface{}{arg1, arg2})
	fake.countObjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) CountObjectsCallCount() int {
	fake.countObjectsMutex.RLock()
	defer fake.countObjectsMutex.RUnlock()
	return len(fake.countObjectsArgsForCall)
}

func (fake *FakeStore) CountObjectsCalls(stub func(context.Context, string) (int64, error)) {
	fake.countObjectsMutex.Lock()
	defer fake.countObjectsMutex.Unlock()
	fake.CountObjectsStub = stub
}

func (fake *FakeStore) CountObjectsArgsForCall(i int) (context.Context, string) {
	fake.countObjectsMutex.RLock()
	defer fake.countObjectsMutex.RUnlock()
	argsForCall := fake.countObjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) CountObjectsReturns(result1 int64, result2 error) {
	fake.countObjectsMutex.Lock()
	defer fake.countObjectsMutex.Unlock()
	fake.CountObjectsStub = nil
	fake.countObjectsReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) CountObjectsReturnsOnCall(i int, result1 int64, result2 error) {
	fake.countObjectsMutex.Lock()
	defer fake.countObjectsMutex.Unlock()
	fake.CountObjectsStub = nil
	if fake.countObjectsReturnsOnCall == nil {
		fake.countObjectsReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.countObjectsReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) DeleteObject(arg1 context.Context, arg2 models.Object) error {
	fake.deleteObjectMutex.Lock()
	ret, specificReturn := fake.deleteObjectReturnsOnCall[len(fake.deleteObjectArgsForCall)]
	fake.deleteObjectArgsForCall = append(fake.deleteObjectArgsForCall, struct {
		arg1 context.Context
		arg2 models.Object
	}{arg1, arg2})
	stub := fake.DeleteObjectStub
	fakeReturns := fake.deleteObjectReturns
	fake.recordInvocation("DeleteObject", []interface{}{arg1, arg2})
	fake.deleteObjectMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) DeleteObjectCallCount() int {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	return len(fake.deleteObjectArgsForCall)
}

func (fake *FakeStore) DeleteObjectCalls(stub func(context.Context, models.Object) error) {
	fake.deleteObjectMutex.Lock()
	defer fake.deleteObjectMutex.Unlock()
	fake.DeleteObjectStub = stub
}

func (fake *FakeStore) DeleteObjectArgsForCall(i int) (context.Context, models.Object) {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	argsForCall := fake.deleteObjectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) DeleteObjectReturns(result1 error) {
	fake.deleteObjectMutex.Lock()
	defer fake.deleteObjectMutex.Unlock()
	fake.DeleteObjectStub = nil
	fake.deleteObjectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) DeleteObjectReturnsOnCall(i int, result1 error) {
	fake.deleteObjectMutex.Lock()
	defer fake.deleteObjectMutex.Unlock()
	fake.DeleteObjectStub = nil
	if fake.deleteObjectReturnsOnCall == nil {
		fake.deleteObjectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteObjectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetAccessRules() ([]models.AccessRule, error) {
	fake.getAccessRulesMutex.Lock()
	ret, specificReturn := fake.getAccessRulesReturnsOnCall[len(fake.getAccessRulesArgsForCall)]
	fake.getAccessRulesArgsForCall = append(fake.getAccessRulesArgsForCall, struct {
	}{})
	stub := fake.GetAccessRulesStub
	fakeReturns := fake.getAccessRulesReturns
	fake.recordInvocation("GetAccessRules", []interface{}{})
	fake.getAccessRulesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetAccessRulesCallCount() int {
	fake.getAccessRulesMutex.RLock()
	defer fake.getAccessRulesMutex.RUnlock()
	return len(fake.getAccessRulesArgsForCall)
}

func (fake *FakeStore) GetAccessRulesCalls(stub func() ([]models.AccessRule, error)) {
	fake.getAccessRulesMutex.Lock()
	defer fake.getAccessRulesMutex.Unlock()
	fake.GetAccessRulesStub = stub
}

func (fake *FakeStore) GetAccessRulesReturns(result1 []models.AccessRule, result2 error) {
	fake.getAccessRulesMutex.Lock()
	defer fake.getAccessRulesMutex.Unlock()
	fake.GetAccessRulesStub = nil
	fake.getAccessRulesReturns = struct {
		result1 []models.AccessRule
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetAccessRulesReturnsOnCall(i int, result1 []models.AccessRule, result2 error) {
	fake.getAccessRulesMutex.Lock()
	defer fake.getAccessRulesMutex.Unlock()
	fake.GetAccessRulesStub = nil
	if fake.getAccessRulesReturnsOnCall == nil {
		fake.getAccessRulesReturnsOnCall = make(map[int]struct {
			result1 []models.AccessRule
			result2 error
		})
	}
	fake.getAccessRulesReturnsOnCall[i] = struct {
		result1 []models.AccessRule
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetObjects() ([]models.Object, error) {
	fake.getObjectsMutex.Lock()
	ret, specificReturn := fake.getObjectsReturnsOnCall[len(fake.getObjectsArgsForCall)]
	fake.getObjectsArgsForCall = append(fake.getObjectsArgsForCall, struct {
	}{})
	stub := fake.GetObjectsStub
	fakeReturns := fake.getObjectsReturns
	fake.recordInvocation("GetObjects", []interface{}{})
	fake.getObjectsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetObjectsCallCount() int {
	fake.getObjectsMutex.RLock()
	defer fake.getObjectsMutex.RUnlock()
	return len(fake.getObjectsArgsForCall)
}

func (fake *FakeStore) GetObjectsCalls(stub func() ([]models.Object, error)) {
	fake.getObjectsMutex.Lock()
	defer fake.getObjectsMutex.Unlock()
	fake.GetObjectsStub = stub
}

func (fake *FakeStore) GetObjectsReturns(result1 []models.Object, result2 error) {
	fake.getObjectsMutex.Lock()
	defer fake.getObjectsMutex.Unlock()
	fake.GetObjectsStub = nil
	fake.getObjectsReturns = struct {
		result1 []models.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetObjectsReturnsOnCall(i int, result1 []models.Object, result2 error) {
	fake.getObjectsMutex.Lock()
	defer fake.getObjectsMutex.Unlock()
	fake.GetObjectsStub = nil
	if fake.getObjectsReturnsOnCall == nil {
		fake.getObjectsReturnsOnCall = make(map[int]struct {
			result1 []models.Object
			result2 error
		})
	}
	fake.getObjectsReturnsOnCall[i] = struct {
		result1 []models.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) StoreAccessRules(arg1 context.Context, arg2 []models.AccessRule) error {
	var arg2Copy []models.AccessRule
	if arg2 != nil {
		arg2Copy = make([]models.AccessRule, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeAccessRulesMutex.Lock()
	ret, specificReturn := fake.storeAccessRulesReturnsOnCall[len(fake.storeAccessRulesArgsForCall)]
	fake.storeAccessRulesArgsForCall = append(fake.storeAccessRulesArgsForCall, struct {
		arg1 context.Context
		arg2 []models.AccessRule
	}{arg1, arg2Copy})
	stub := fake.StoreAccessRulesStub
	fakeReturns := fake.storeAccessRulesReturns
	fake.recordInvocation("StoreAccessRules", []interface{}{arg1, arg2Copy})
	fake.storeAccessRulesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) StoreAccessRulesCallCount() int {
	fake.storeAccessRulesMutex.RLock()
	defer fake.storeAccessRulesMutex.RUnlock()
	return len(fake.storeAccessRulesArgsForCall)
}

func (fake *FakeStore) StoreAccessRulesCalls(stub func(context.Context, []models.AccessRule) error) {
	fake.storeAccessRulesMutex.Lock()
	defer fake.storeAccessRulesMutex.Unlock()
	fake.StoreAccessRulesStub = stub
}

func (fake *FakeStore) StoreAccessRulesArgsForCall(i int) (context.Context, []models.AccessRule) {
	fake.storeAccessRulesMutex.RLock()
	defer fake.storeAccessRulesMutex.RUnlock()
	argsForCall := fake.storeAccessRulesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) StoreAccessRulesReturns(result1 error) {
	fake.storeAccessRulesMutex.Lock()
	defer fake.storeAccessRulesMutex.Unlock()
	fake.StoreAccessRulesStub = nil
	fake.storeAccessRulesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreAccessRulesReturnsOnCall(i int, result1 error) {
	fake.storeAccessRulesMutex.Lock()
	defer fake.storeAccessRulesMutex.Unlock()
	fake.StoreAccessRulesStub = nil
	if fake.storeAccessRulesReturnsOnCall == nil {
		fake.storeAccessRulesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeAccessRulesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreObjects(arg1 context.Context, arg2 []models.Object) error {
	var arg2Copy []models.Object
	if arg2 != nil {
		arg2Copy = make([]models.Object, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeObjectsMutex.Lock()
	ret, specificReturn := fake.storeObjectsReturnsOnCall[len(fake.storeObjectsArgsForCall)]
	fake.storeObjectsArgsForCall = append(fake.storeObjectsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Object
	}{arg1, arg2Copy})
	stub := fake.StoreObjectsStub
	fakeReturns := fake.storeObjectsReturns
	fake.recordInvocation("StoreObjects", []interface{}{arg1, arg2Copy})
	fake.storeObjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) StoreObjectsCallCount() int {
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	return len(fake.storeObjectsArgsForCall)
}

func (fake *FakeStore) StoreObjectsCalls(stub func(context.Context, []models.Object) error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = stub
}

func (fake *FakeStore) StoreObjectsArgsForCall(i int) (context.Context, []models.Object) {
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	argsForCall := fake.storeObjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) StoreObjectsReturns(result1 error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = nil
	fake.storeObjectsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreObjectsReturnsOnCall(i int, result1 error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = nil
	if fake.storeObjectsReturnsOnCall == nil {
		fake.storeObjectsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeObjectsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.countObjectsMutex.RLock()
	defer fake.countObjectsMutex.RUnlock()
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	fake.getAccessRulesMutex.RLock()
	defer fake.getAccessRulesMutex.RUnlock()
	fake.getObjectsMutex.RLock()
	defer fake.getObjectsMutex.RUnlock()
	fake.storeAccessRulesMutex.RLock()
	defer fake.storeAccessRulesMutex.RUnlock()
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
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

var _ store.Store = new(FakeStore)
