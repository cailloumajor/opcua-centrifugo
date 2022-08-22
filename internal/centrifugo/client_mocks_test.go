// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package centrifugo

import (
	"context"
	"github.com/centrifugal/gocent/v3"
	"sync"
)

// Ensure, that ClientProviderMock does implement ClientProvider.
// If this is not the case, regenerate this file with moq.
var _ ClientProvider = &ClientProviderMock{}

// ClientProviderMock is a mock implementation of ClientProvider.
//
//	func TestSomethingThatUsesClientProvider(t *testing.T) {
//
//		// make and configure a mocked ClientProvider
//		mockedClientProvider := &ClientProviderMock{
//			ChannelsFunc: func(ctx context.Context, opts ...gocent.ChannelsOption) (gocent.ChannelsResult, error) {
//				panic("mock out the Channels method")
//			},
//			InfoFunc: func(ctx context.Context) (gocent.InfoResult, error) {
//				panic("mock out the Info method")
//			},
//			PublishFunc: func(ctx context.Context, channel string, data []byte, opts ...gocent.PublishOption) (gocent.PublishResult, error) {
//				panic("mock out the Publish method")
//			},
//		}
//
//		// use mockedClientProvider in code that requires ClientProvider
//		// and then make assertions.
//
//	}
type ClientProviderMock struct {
	// ChannelsFunc mocks the Channels method.
	ChannelsFunc func(ctx context.Context, opts ...gocent.ChannelsOption) (gocent.ChannelsResult, error)

	// InfoFunc mocks the Info method.
	InfoFunc func(ctx context.Context) (gocent.InfoResult, error)

	// PublishFunc mocks the Publish method.
	PublishFunc func(ctx context.Context, channel string, data []byte, opts ...gocent.PublishOption) (gocent.PublishResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// Channels holds details about calls to the Channels method.
		Channels []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts []gocent.ChannelsOption
		}
		// Info holds details about calls to the Info method.
		Info []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Publish holds details about calls to the Publish method.
		Publish []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Channel is the channel argument value.
			Channel string
			// Data is the data argument value.
			Data []byte
			// Opts is the opts argument value.
			Opts []gocent.PublishOption
		}
	}
	lockChannels sync.RWMutex
	lockInfo     sync.RWMutex
	lockPublish  sync.RWMutex
}

// Channels calls ChannelsFunc.
func (mock *ClientProviderMock) Channels(ctx context.Context, opts ...gocent.ChannelsOption) (gocent.ChannelsResult, error) {
	if mock.ChannelsFunc == nil {
		panic("ClientProviderMock.ChannelsFunc: method is nil but ClientProvider.Channels was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts []gocent.ChannelsOption
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockChannels.Lock()
	mock.calls.Channels = append(mock.calls.Channels, callInfo)
	mock.lockChannels.Unlock()
	return mock.ChannelsFunc(ctx, opts...)
}

// ChannelsCalls gets all the calls that were made to Channels.
// Check the length with:
//
//	len(mockedClientProvider.ChannelsCalls())
func (mock *ClientProviderMock) ChannelsCalls() []struct {
	Ctx  context.Context
	Opts []gocent.ChannelsOption
} {
	var calls []struct {
		Ctx  context.Context
		Opts []gocent.ChannelsOption
	}
	mock.lockChannels.RLock()
	calls = mock.calls.Channels
	mock.lockChannels.RUnlock()
	return calls
}

// Info calls InfoFunc.
func (mock *ClientProviderMock) Info(ctx context.Context) (gocent.InfoResult, error) {
	if mock.InfoFunc == nil {
		panic("ClientProviderMock.InfoFunc: method is nil but ClientProvider.Info was just called")
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
//
//	len(mockedClientProvider.InfoCalls())
func (mock *ClientProviderMock) InfoCalls() []struct {
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

// Publish calls PublishFunc.
func (mock *ClientProviderMock) Publish(ctx context.Context, channel string, data []byte, opts ...gocent.PublishOption) (gocent.PublishResult, error) {
	if mock.PublishFunc == nil {
		panic("ClientProviderMock.PublishFunc: method is nil but ClientProvider.Publish was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Channel string
		Data    []byte
		Opts    []gocent.PublishOption
	}{
		Ctx:     ctx,
		Channel: channel,
		Data:    data,
		Opts:    opts,
	}
	mock.lockPublish.Lock()
	mock.calls.Publish = append(mock.calls.Publish, callInfo)
	mock.lockPublish.Unlock()
	return mock.PublishFunc(ctx, channel, data, opts...)
}

// PublishCalls gets all the calls that were made to Publish.
// Check the length with:
//
//	len(mockedClientProvider.PublishCalls())
func (mock *ClientProviderMock) PublishCalls() []struct {
	Ctx     context.Context
	Channel string
	Data    []byte
	Opts    []gocent.PublishOption
} {
	var calls []struct {
		Ctx     context.Context
		Channel string
		Data    []byte
		Opts    []gocent.PublishOption
	}
	mock.lockPublish.RLock()
	calls = mock.calls.Publish
	mock.lockPublish.RUnlock()
	return calls
}