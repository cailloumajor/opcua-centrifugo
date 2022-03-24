// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package proxy

import (
	"context"
	"github.com/cailloumajor/opcua-proxy/internal/opcua"
	"github.com/centrifugal/gocent/v3"
	gopcua "github.com/gopcua/opcua"
	"sync"
)

// Ensure, that MonitorProviderMock does implement MonitorProvider.
// If this is not the case, regenerate this file with moq.
var _ MonitorProvider = &MonitorProviderMock{}

// MonitorProviderMock is a mock implementation of MonitorProvider.
//
// 	func TestSomethingThatUsesMonitorProvider(t *testing.T) {
//
// 		// make and configure a mocked MonitorProvider
// 		mockedMonitorProvider := &MonitorProviderMock{
// 			StateFunc: func() gopcua.ConnState {
// 				panic("mock out the State method")
// 			},
// 			SubscribeFunc: func(ctx context.Context, nsURI string, ch opcua.ChannelProvider, nodes []opcua.NodeIDProvider) error {
// 				panic("mock out the Subscribe method")
// 			},
// 		}
//
// 		// use mockedMonitorProvider in code that requires MonitorProvider
// 		// and then make assertions.
//
// 	}
type MonitorProviderMock struct {
	// StateFunc mocks the State method.
	StateFunc func() gopcua.ConnState

	// SubscribeFunc mocks the Subscribe method.
	SubscribeFunc func(ctx context.Context, nsURI string, ch opcua.ChannelProvider, nodes []opcua.NodeIDProvider) error

	// calls tracks calls to the methods.
	calls struct {
		// State holds details about calls to the State method.
		State []struct {
		}
		// Subscribe holds details about calls to the Subscribe method.
		Subscribe []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// NsURI is the nsURI argument value.
			NsURI string
			// Ch is the ch argument value.
			Ch opcua.ChannelProvider
			// Nodes is the nodes argument value.
			Nodes []opcua.NodeIDProvider
		}
	}
	lockState     sync.RWMutex
	lockSubscribe sync.RWMutex
}

// State calls StateFunc.
func (mock *MonitorProviderMock) State() gopcua.ConnState {
	if mock.StateFunc == nil {
		panic("MonitorProviderMock.StateFunc: method is nil but MonitorProvider.State was just called")
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
//     len(mockedMonitorProvider.StateCalls())
func (mock *MonitorProviderMock) StateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockState.RLock()
	calls = mock.calls.State
	mock.lockState.RUnlock()
	return calls
}

// Subscribe calls SubscribeFunc.
func (mock *MonitorProviderMock) Subscribe(ctx context.Context, nsURI string, ch opcua.ChannelProvider, nodes []opcua.NodeIDProvider) error {
	if mock.SubscribeFunc == nil {
		panic("MonitorProviderMock.SubscribeFunc: method is nil but MonitorProvider.Subscribe was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		NsURI string
		Ch    opcua.ChannelProvider
		Nodes []opcua.NodeIDProvider
	}{
		Ctx:   ctx,
		NsURI: nsURI,
		Ch:    ch,
		Nodes: nodes,
	}
	mock.lockSubscribe.Lock()
	mock.calls.Subscribe = append(mock.calls.Subscribe, callInfo)
	mock.lockSubscribe.Unlock()
	return mock.SubscribeFunc(ctx, nsURI, ch, nodes)
}

// SubscribeCalls gets all the calls that were made to Subscribe.
// Check the length with:
//     len(mockedMonitorProvider.SubscribeCalls())
func (mock *MonitorProviderMock) SubscribeCalls() []struct {
	Ctx   context.Context
	NsURI string
	Ch    opcua.ChannelProvider
	Nodes []opcua.NodeIDProvider
} {
	var calls []struct {
		Ctx   context.Context
		NsURI string
		Ch    opcua.ChannelProvider
		Nodes []opcua.NodeIDProvider
	}
	mock.lockSubscribe.RLock()
	calls = mock.calls.Subscribe
	mock.lockSubscribe.RUnlock()
	return calls
}

// Ensure, that CentrifugoChannelParserMock does implement CentrifugoChannelParser.
// If this is not the case, regenerate this file with moq.
var _ CentrifugoChannelParser = &CentrifugoChannelParserMock{}

// CentrifugoChannelParserMock is a mock implementation of CentrifugoChannelParser.
//
// 	func TestSomethingThatUsesCentrifugoChannelParser(t *testing.T) {
//
// 		// make and configure a mocked CentrifugoChannelParser
// 		mockedCentrifugoChannelParser := &CentrifugoChannelParserMock{
// 			ParseChannelFunc: func(s string, namespace string) (opcua.ChannelProvider, error) {
// 				panic("mock out the ParseChannel method")
// 			},
// 		}
//
// 		// use mockedCentrifugoChannelParser in code that requires CentrifugoChannelParser
// 		// and then make assertions.
//
// 	}
type CentrifugoChannelParserMock struct {
	// ParseChannelFunc mocks the ParseChannel method.
	ParseChannelFunc func(s string, namespace string) (opcua.ChannelProvider, error)

	// calls tracks calls to the methods.
	calls struct {
		// ParseChannel holds details about calls to the ParseChannel method.
		ParseChannel []struct {
			// S is the s argument value.
			S string
			// Namespace is the namespace argument value.
			Namespace string
		}
	}
	lockParseChannel sync.RWMutex
}

// ParseChannel calls ParseChannelFunc.
func (mock *CentrifugoChannelParserMock) ParseChannel(s string, namespace string) (opcua.ChannelProvider, error) {
	if mock.ParseChannelFunc == nil {
		panic("CentrifugoChannelParserMock.ParseChannelFunc: method is nil but CentrifugoChannelParser.ParseChannel was just called")
	}
	callInfo := struct {
		S         string
		Namespace string
	}{
		S:         s,
		Namespace: namespace,
	}
	mock.lockParseChannel.Lock()
	mock.calls.ParseChannel = append(mock.calls.ParseChannel, callInfo)
	mock.lockParseChannel.Unlock()
	return mock.ParseChannelFunc(s, namespace)
}

// ParseChannelCalls gets all the calls that were made to ParseChannel.
// Check the length with:
//     len(mockedCentrifugoChannelParser.ParseChannelCalls())
func (mock *CentrifugoChannelParserMock) ParseChannelCalls() []struct {
	S         string
	Namespace string
} {
	var calls []struct {
		S         string
		Namespace string
	}
	mock.lockParseChannel.RLock()
	calls = mock.calls.ParseChannel
	mock.lockParseChannel.RUnlock()
	return calls
}

// Ensure, that CentrifugoInfoProviderMock does implement CentrifugoInfoProvider.
// If this is not the case, regenerate this file with moq.
var _ CentrifugoInfoProvider = &CentrifugoInfoProviderMock{}

// CentrifugoInfoProviderMock is a mock implementation of CentrifugoInfoProvider.
//
// 	func TestSomethingThatUsesCentrifugoInfoProvider(t *testing.T) {
//
// 		// make and configure a mocked CentrifugoInfoProvider
// 		mockedCentrifugoInfoProvider := &CentrifugoInfoProviderMock{
// 			InfoFunc: func(ctx context.Context) (gocent.InfoResult, error) {
// 				panic("mock out the Info method")
// 			},
// 		}
//
// 		// use mockedCentrifugoInfoProvider in code that requires CentrifugoInfoProvider
// 		// and then make assertions.
//
// 	}
type CentrifugoInfoProviderMock struct {
	// InfoFunc mocks the Info method.
	InfoFunc func(ctx context.Context) (gocent.InfoResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// Info holds details about calls to the Info method.
		Info []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockInfo sync.RWMutex
}

// Info calls InfoFunc.
func (mock *CentrifugoInfoProviderMock) Info(ctx context.Context) (gocent.InfoResult, error) {
	if mock.InfoFunc == nil {
		panic("CentrifugoInfoProviderMock.InfoFunc: method is nil but CentrifugoInfoProvider.Info was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockInfo.Lock()
	mock.calls.Info = append(mock.calls.Info, callInfo)
	mock.lockInfo.Unlock()
	return mock.InfoFunc(ctx)
}

// InfoCalls gets all the calls that were made to Info.
// Check the length with:
//     len(mockedCentrifugoInfoProvider.InfoCalls())
func (mock *CentrifugoInfoProviderMock) InfoCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockInfo.RLock()
	calls = mock.calls.Info
	mock.lockInfo.RUnlock()
	return calls
}
