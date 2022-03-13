// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package opcua

import (
	"context"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
	"sync"
	"time"
)

// Ensure, that ClientProviderMock does implement ClientProvider.
// If this is not the case, regenerate this file with moq.
var _ ClientProvider = &ClientProviderMock{}

// ClientProviderMock is a mock implementation of ClientProvider.
//
// 	func TestSomethingThatUsesClientProvider(t *testing.T) {
//
// 		// make and configure a mocked ClientProvider
// 		mockedClientProvider := &ClientProviderMock{
// 			CloseWithContextFunc: func(ctx context.Context) error {
// 				panic("mock out the CloseWithContext method")
// 			},
// 			NamespaceIndexFunc: func(ctx context.Context, nsURI string) (uint16, error) {
// 				panic("mock out the NamespaceIndex method")
// 			},
// 			StateFunc: func() opcua.ConnState {
// 				panic("mock out the State method")
// 			},
// 			SubscribeFunc: func(ctx context.Context, params *opcua.SubscriptionParameters, notifyCh chan<- *opcua.PublishNotificationData) (SubscriptionProvider, error) {
// 				panic("mock out the Subscribe method")
// 			},
// 		}
//
// 		// use mockedClientProvider in code that requires ClientProvider
// 		// and then make assertions.
//
// 	}
type ClientProviderMock struct {
	// CloseWithContextFunc mocks the CloseWithContext method.
	CloseWithContextFunc func(ctx context.Context) error

	// NamespaceIndexFunc mocks the NamespaceIndex method.
	NamespaceIndexFunc func(ctx context.Context, nsURI string) (uint16, error)

	// StateFunc mocks the State method.
	StateFunc func() opcua.ConnState

	// SubscribeFunc mocks the Subscribe method.
	SubscribeFunc func(ctx context.Context, params *opcua.SubscriptionParameters, notifyCh chan<- *opcua.PublishNotificationData) (SubscriptionProvider, error)

	// calls tracks calls to the methods.
	calls struct {
		// CloseWithContext holds details about calls to the CloseWithContext method.
		CloseWithContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// NamespaceIndex holds details about calls to the NamespaceIndex method.
		NamespaceIndex []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// NsURI is the nsURI argument value.
			NsURI string
		}
		// State holds details about calls to the State method.
		State []struct {
		}
		// Subscribe holds details about calls to the Subscribe method.
		Subscribe []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *opcua.SubscriptionParameters
			// NotifyCh is the notifyCh argument value.
			NotifyCh chan<- *opcua.PublishNotificationData
		}
	}
	lockCloseWithContext sync.RWMutex
	lockNamespaceIndex   sync.RWMutex
	lockState            sync.RWMutex
	lockSubscribe        sync.RWMutex
}

// CloseWithContext calls CloseWithContextFunc.
func (mock *ClientProviderMock) CloseWithContext(ctx context.Context) error {
	if mock.CloseWithContextFunc == nil {
		panic("ClientProviderMock.CloseWithContextFunc: method is nil but ClientProvider.CloseWithContext was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCloseWithContext.Lock()
	mock.calls.CloseWithContext = append(mock.calls.CloseWithContext, callInfo)
	mock.lockCloseWithContext.Unlock()
	return mock.CloseWithContextFunc(ctx)
}

// CloseWithContextCalls gets all the calls that were made to CloseWithContext.
// Check the length with:
//     len(mockedClientProvider.CloseWithContextCalls())
func (mock *ClientProviderMock) CloseWithContextCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCloseWithContext.RLock()
	calls = mock.calls.CloseWithContext
	mock.lockCloseWithContext.RUnlock()
	return calls
}

// NamespaceIndex calls NamespaceIndexFunc.
func (mock *ClientProviderMock) NamespaceIndex(ctx context.Context, nsURI string) (uint16, error) {
	if mock.NamespaceIndexFunc == nil {
		panic("ClientProviderMock.NamespaceIndexFunc: method is nil but ClientProvider.NamespaceIndex was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		NsURI string
	}{
		Ctx:   ctx,
		NsURI: nsURI,
	}
	mock.lockNamespaceIndex.Lock()
	mock.calls.NamespaceIndex = append(mock.calls.NamespaceIndex, callInfo)
	mock.lockNamespaceIndex.Unlock()
	return mock.NamespaceIndexFunc(ctx, nsURI)
}

// NamespaceIndexCalls gets all the calls that were made to NamespaceIndex.
// Check the length with:
//     len(mockedClientProvider.NamespaceIndexCalls())
func (mock *ClientProviderMock) NamespaceIndexCalls() []struct {
	Ctx   context.Context
	NsURI string
} {
	var calls []struct {
		Ctx   context.Context
		NsURI string
	}
	mock.lockNamespaceIndex.RLock()
	calls = mock.calls.NamespaceIndex
	mock.lockNamespaceIndex.RUnlock()
	return calls
}

// State calls StateFunc.
func (mock *ClientProviderMock) State() opcua.ConnState {
	if mock.StateFunc == nil {
		panic("ClientProviderMock.StateFunc: method is nil but ClientProvider.State was just called")
	}
	callInfo := struct {
	}{}
	mock.lockState.Lock()
	mock.calls.State = append(mock.calls.State, callInfo)
	mock.lockState.Unlock()
	return mock.StateFunc()
}

// StateCalls gets all the calls that were made to State.
// Check the length with:
//     len(mockedClientProvider.StateCalls())
func (mock *ClientProviderMock) StateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockState.RLock()
	calls = mock.calls.State
	mock.lockState.RUnlock()
	return calls
}

// Subscribe calls SubscribeFunc.
func (mock *ClientProviderMock) Subscribe(ctx context.Context, params *opcua.SubscriptionParameters, notifyCh chan<- *opcua.PublishNotificationData) (SubscriptionProvider, error) {
	if mock.SubscribeFunc == nil {
		panic("ClientProviderMock.SubscribeFunc: method is nil but ClientProvider.Subscribe was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Params   *opcua.SubscriptionParameters
		NotifyCh chan<- *opcua.PublishNotificationData
	}{
		Ctx:      ctx,
		Params:   params,
		NotifyCh: notifyCh,
	}
	mock.lockSubscribe.Lock()
	mock.calls.Subscribe = append(mock.calls.Subscribe, callInfo)
	mock.lockSubscribe.Unlock()
	return mock.SubscribeFunc(ctx, params, notifyCh)
}

// SubscribeCalls gets all the calls that were made to Subscribe.
// Check the length with:
//     len(mockedClientProvider.SubscribeCalls())
func (mock *ClientProviderMock) SubscribeCalls() []struct {
	Ctx      context.Context
	Params   *opcua.SubscriptionParameters
	NotifyCh chan<- *opcua.PublishNotificationData
} {
	var calls []struct {
		Ctx      context.Context
		Params   *opcua.SubscriptionParameters
		NotifyCh chan<- *opcua.PublishNotificationData
	}
	mock.lockSubscribe.RLock()
	calls = mock.calls.Subscribe
	mock.lockSubscribe.RUnlock()
	return calls
}

// Ensure, that SubscriptionProviderMock does implement SubscriptionProvider.
// If this is not the case, regenerate this file with moq.
var _ SubscriptionProvider = &SubscriptionProviderMock{}

// SubscriptionProviderMock is a mock implementation of SubscriptionProvider.
//
// 	func TestSomethingThatUsesSubscriptionProvider(t *testing.T) {
//
// 		// make and configure a mocked SubscriptionProvider
// 		mockedSubscriptionProvider := &SubscriptionProviderMock{
// 			CancelFunc: func(ctx context.Context) error {
// 				panic("mock out the Cancel method")
// 			},
// 			IDFunc: func() uint32 {
// 				panic("mock out the ID method")
// 			},
// 			MonitorWithContextFunc: func(ctx context.Context, ts ua.TimestampsToReturn, items ...*ua.MonitoredItemCreateRequest) (*ua.CreateMonitoredItemsResponse, error) {
// 				panic("mock out the MonitorWithContext method")
// 			},
// 		}
//
// 		// use mockedSubscriptionProvider in code that requires SubscriptionProvider
// 		// and then make assertions.
//
// 	}
type SubscriptionProviderMock struct {
	// CancelFunc mocks the Cancel method.
	CancelFunc func(ctx context.Context) error

	// IDFunc mocks the ID method.
	IDFunc func() uint32

	// MonitorWithContextFunc mocks the MonitorWithContext method.
	MonitorWithContextFunc func(ctx context.Context, ts ua.TimestampsToReturn, items ...*ua.MonitoredItemCreateRequest) (*ua.CreateMonitoredItemsResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// Cancel holds details about calls to the Cancel method.
		Cancel []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// MonitorWithContext holds details about calls to the MonitorWithContext method.
		MonitorWithContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Ts is the ts argument value.
			Ts ua.TimestampsToReturn
			// Items is the items argument value.
			Items []*ua.MonitoredItemCreateRequest
		}
	}
	lockCancel             sync.RWMutex
	lockID                 sync.RWMutex
	lockMonitorWithContext sync.RWMutex
}

// Cancel calls CancelFunc.
func (mock *SubscriptionProviderMock) Cancel(ctx context.Context) error {
	if mock.CancelFunc == nil {
		panic("SubscriptionProviderMock.CancelFunc: method is nil but SubscriptionProvider.Cancel was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCancel.Lock()
	mock.calls.Cancel = append(mock.calls.Cancel, callInfo)
	mock.lockCancel.Unlock()
	return mock.CancelFunc(ctx)
}

// CancelCalls gets all the calls that were made to Cancel.
// Check the length with:
//     len(mockedSubscriptionProvider.CancelCalls())
func (mock *SubscriptionProviderMock) CancelCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCancel.RLock()
	calls = mock.calls.Cancel
	mock.lockCancel.RUnlock()
	return calls
}

// ID calls IDFunc.
func (mock *SubscriptionProviderMock) ID() uint32 {
	if mock.IDFunc == nil {
		panic("SubscriptionProviderMock.IDFunc: method is nil but SubscriptionProvider.ID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	mock.lockID.Unlock()
	return mock.IDFunc()
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//     len(mockedSubscriptionProvider.IDCalls())
func (mock *SubscriptionProviderMock) IDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockID.RLock()
	calls = mock.calls.ID
	mock.lockID.RUnlock()
	return calls
}

// MonitorWithContext calls MonitorWithContextFunc.
func (mock *SubscriptionProviderMock) MonitorWithContext(ctx context.Context, ts ua.TimestampsToReturn, items ...*ua.MonitoredItemCreateRequest) (*ua.CreateMonitoredItemsResponse, error) {
	if mock.MonitorWithContextFunc == nil {
		panic("SubscriptionProviderMock.MonitorWithContextFunc: method is nil but SubscriptionProvider.MonitorWithContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Ts    ua.TimestampsToReturn
		Items []*ua.MonitoredItemCreateRequest
	}{
		Ctx:   ctx,
		Ts:    ts,
		Items: items,
	}
	mock.lockMonitorWithContext.Lock()
	mock.calls.MonitorWithContext = append(mock.calls.MonitorWithContext, callInfo)
	mock.lockMonitorWithContext.Unlock()
	return mock.MonitorWithContextFunc(ctx, ts, items...)
}

// MonitorWithContextCalls gets all the calls that were made to MonitorWithContext.
// Check the length with:
//     len(mockedSubscriptionProvider.MonitorWithContextCalls())
func (mock *SubscriptionProviderMock) MonitorWithContextCalls() []struct {
	Ctx   context.Context
	Ts    ua.TimestampsToReturn
	Items []*ua.MonitoredItemCreateRequest
} {
	var calls []struct {
		Ctx   context.Context
		Ts    ua.TimestampsToReturn
		Items []*ua.MonitoredItemCreateRequest
	}
	mock.lockMonitorWithContext.RLock()
	calls = mock.calls.MonitorWithContext
	mock.lockMonitorWithContext.RUnlock()
	return calls
}

// Ensure, that ChannelProviderMock does implement ChannelProvider.
// If this is not the case, regenerate this file with moq.
var _ ChannelProvider = &ChannelProviderMock{}

// ChannelProviderMock is a mock implementation of ChannelProvider.
//
// 	func TestSomethingThatUsesChannelProvider(t *testing.T) {
//
// 		// make and configure a mocked ChannelProvider
// 		mockedChannelProvider := &ChannelProviderMock{
// 			IntervalFunc: func() time.Duration {
// 				panic("mock out the Interval method")
// 			},
// 			StringFunc: func() string {
// 				panic("mock out the String method")
// 			},
// 		}
//
// 		// use mockedChannelProvider in code that requires ChannelProvider
// 		// and then make assertions.
//
// 	}
type ChannelProviderMock struct {
	// IntervalFunc mocks the Interval method.
	IntervalFunc func() time.Duration

	// StringFunc mocks the String method.
	StringFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Interval holds details about calls to the Interval method.
		Interval []struct {
		}
		// String holds details about calls to the String method.
		String []struct {
		}
	}
	lockInterval sync.RWMutex
	lockString   sync.RWMutex
}

// Interval calls IntervalFunc.
func (mock *ChannelProviderMock) Interval() time.Duration {
	if mock.IntervalFunc == nil {
		panic("ChannelProviderMock.IntervalFunc: method is nil but ChannelProvider.Interval was just called")
	}
	callInfo := struct {
	}{}
	mock.lockInterval.Lock()
	mock.calls.Interval = append(mock.calls.Interval, callInfo)
	mock.lockInterval.Unlock()
	return mock.IntervalFunc()
}

// IntervalCalls gets all the calls that were made to Interval.
// Check the length with:
//     len(mockedChannelProvider.IntervalCalls())
func (mock *ChannelProviderMock) IntervalCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockInterval.RLock()
	calls = mock.calls.Interval
	mock.lockInterval.RUnlock()
	return calls
}

// String calls StringFunc.
func (mock *ChannelProviderMock) String() string {
	if mock.StringFunc == nil {
		panic("ChannelProviderMock.StringFunc: method is nil but ChannelProvider.String was just called")
	}
	callInfo := struct {
	}{}
	mock.lockString.Lock()
	mock.calls.String = append(mock.calls.String, callInfo)
	mock.lockString.Unlock()
	return mock.StringFunc()
}

// StringCalls gets all the calls that were made to String.
// Check the length with:
//     len(mockedChannelProvider.StringCalls())
func (mock *ChannelProviderMock) StringCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockString.RLock()
	calls = mock.calls.String
	mock.lockString.RUnlock()
	return calls
}
