package subscription

import (
	"github.com/karimra/gnmic/collector"
	"github.com/netw-device-driver/ndd-runtime/pkg/logging"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

// A SubscriptionAction represents an action on a GNMI subscription
type SubscriptionAction string

// Condition Kinds.
const (
	// start
	SubscriptionActionStart SubscriptionAction = "start"
	// stop
	SubscriptionActionStop SubscriptionAction = "stop"
)

type Subscription struct {
	Name   string
	Action SubscriptionAction
	Target *collector.Collector
}

// DeviationServer contains the device driver information
type SubscriptionServer struct {
	eventChs map[string]chan event.GenericEvent
	subCh    chan Subscription
	log      logging.Logger
	stopCh   chan struct{}
	//ctx      context.Context
}

// Option is a function to initialize the options
type Option func(d *SubscriptionServer)

// WithDeviationServer initializes the deviation server in the srl operator
func WithEventChannels(e map[string]chan event.GenericEvent) Option {
	return func(s *SubscriptionServer) {
		s.eventChs = e
	}
}

func WithSubscriptionChannel(sub chan Subscription) Option {
	return func(d *SubscriptionServer) {
		d.subCh = sub
	}
}

func WithLogging(l logging.Logger) Option {
	return func(d *SubscriptionServer) {
		d.log = l
	}
}

func WithStopChannel(stopCh chan struct{}) Option {
	return func(d *SubscriptionServer) {
		d.stopCh = stopCh
	}
}

// NewSubscriptionServer function defines a new SubscriptionServer
func NewSubscriptionServer(opts ...Option) *SubscriptionServer {
	s := &SubscriptionServer{}

	for _, o := range opts {
		o(s)
	}

	return s
}

// StartDeviationGRPCServer function starts the deviation server
func (s *SubscriptionServer) StartSubscriptionHandler() {
	s.log.Debug("Starting subscription gnmi server...")

	for {
		select {
		case subscription := <-s.subCh:
			s.log.Debug("subscription server", "Action", subscription.Action, "Target", subscription.Name)
			// TODO Add target subscription
		case <-s.stopCh:
			s.log.Debug("stopping subscription handler")

		}
	}
}
