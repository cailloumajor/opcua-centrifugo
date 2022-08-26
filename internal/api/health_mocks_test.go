// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package api

import (
	"context"
	"sync"
)

// Ensure, that HealtherMock does implement Healther.
// If this is not the case, regenerate this file with moq.
var _ Healther = &HealtherMock{}

// HealtherMock is a mock implementation of Healther.
//
//	func TestSomethingThatUsesHealther(t *testing.T) {
//
//		// make and configure a mocked Healther
//		mockedHealther := &HealtherMock{
//			HealthFunc: func(ctx context.Context) (bool, string) {
//				panic("mock out the Health method")
//			},
//		}
//
//		// use mockedHealther in code that requires Healther
//		// and then make assertions.
//
//	}
type HealtherMock struct {
	// HealthFunc mocks the Health method.
	HealthFunc func(ctx context.Context) (bool, string)

	// calls tracks calls to the methods.
	calls struct {
		// Health holds details about calls to the Health method.
		Health []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockHealth sync.RWMutex
}

// Health calls HealthFunc.
func (mock *HealtherMock) Health(ctx context.Context) (bool, string) {
	if mock.HealthFunc == nil {
		panic("HealtherMock.HealthFunc: method is nil but Healther.Health was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockHealth.Lock()
	mock.calls.Health = append(mock.calls.Health, callInfo)
	mock.lockHealth.Unlock()
	return mock.HealthFunc(ctx)
}

// HealthCalls gets all the calls that were made to Health.
// Check the length with:
//
//	len(mockedHealther.HealthCalls())
func (mock *HealtherMock) HealthCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockHealth.RLock()
	calls = mock.calls.Health
	mock.lockHealth.RUnlock()
	return calls
}
