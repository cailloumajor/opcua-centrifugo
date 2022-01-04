package opcua

import (
	"context"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
)

//go:generate moq -out mocks_test.go -pkg opcua_test . EndpointsGetter EndpointSelector Client ClientCreator NodeMonitor NodeMonitorCreator

// EndpointsGetter models an OPC-UA server endpoints getter
type EndpointsGetter interface {
	GetEndpoints(ctx context.Context, endpoint string, opts ...opcua.Option) ([]*ua.EndpointDescription, error)
}

// EndpointSelector models an OPC-UA server endpoint selector
type EndpointSelector interface {
	SelectEndpoint(endpoints []*ua.EndpointDescription, policy string, mode ua.MessageSecurityMode) *ua.EndpointDescription
}

// Client models an OPC-UA client
type Client interface {
	Connect(context.Context) (err error)
}

// ClientCreator models an OPC-UA client creator
type ClientCreator interface {
	NewClient(endpoint string, opts ...opcua.Option) Client
}

// NodeMonitor models an OPC-UA node monitor
type NodeMonitor interface {
	ChanSubscribe(context.Context, *opcua.SubscriptionParameters, chan<- *monitor.DataChangeMessage, ...string) (*monitor.Subscription, error)
}

// NodeMonitorCreator models an OPC-UA node monitor creator
type NodeMonitorCreator interface {
	NewNodeMonitor(client Client) (NodeMonitor, error)
}

// Config holds the OPC-UA part of the configuration
type Config struct {
	ServerURL string
	User      string `envconfig:"optional"`
	Password  string `envconfig:"optional"`
	CertFile  string `envconfig:"optional"`
	KeyFile   string `envconfig:"optional"`
}

// NewMonitor creates an OPC-UA node monitor
func NewMonitor(ctx context.Context, config *Config, eg EndpointsGetter, es EndpointSelector, cc ClientCreator, nmc NodeMonitorCreator) (NodeMonitor, error) {
	ep, _ := eg.GetEndpoints(ctx, config.ServerURL)

	es.SelectEndpoint(ep, "None", ua.MessageSecurityModeNone)

	cc.NewClient("blabla")

	return nil, nil
}
