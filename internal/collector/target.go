/*
Copyright 2021 Wim Henderickx.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package collector

import (
	"context"
	"time"

	"github.com/karimra/gnmic/target"
	"github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

const (
	//
	configSubscription = "ConfigChangesubscription"

	// errors
	errCreateGnmiClient = "cannot create gnmi client"

	// timers
	defaultTimeout = 5 * time.Second
)

// A TargetAction represents an action on a target
type TargetAction string

// Condition Kinds.
const (
	// start
	TargetAdd TargetAction = "target added"
	// stop
	TargetDelete TargetAction = "target deleted"
)

type TargetUpdate struct {
	Name         string
	Action       TargetAction
	TargetConfig *types.TargetConfig
}

// DeviationServer contains the device driver information
type DeviationServer struct {
	eventChs map[string]chan event.GenericEvent
	tuCh     chan TargetUpdate
	log      logging.Logger
	stopCh   chan struct{}
	Targets  map[string]*Target
	ctx      context.Context
}

type Target struct {
	Config    *types.TargetConfig
	Target    *target.Target
	StopCh    chan struct{}
	log       logging.Logger
	Collector *GNMICollector
}

// Option is a function to initialize the options
type Option func(d *DeviationServer)

// WithDeviationServer initializes the deviation server in the srl operator
func WithEventChannels(e map[string]chan event.GenericEvent) Option {
	return func(s *DeviationServer) {
		s.eventChs = e
	}
}

func WithTargetUpdateChannel(tu chan TargetUpdate) Option {
	return func(d *DeviationServer) {
		d.tuCh = tu
	}
}

func WithLogging(l logging.Logger) Option {
	return func(d *DeviationServer) {
		d.log = l
	}
}

func WithStopChannel(stopCh chan struct{}) Option {
	return func(d *DeviationServer) {
		d.stopCh = stopCh
	}
}

// NewSubscriptionServer function defines a new SubscriptionServer
func NewDeviationServer(opts ...Option) *DeviationServer {
	s := &DeviationServer{
		Targets: make(map[string]*Target),
		ctx:     context.Background(),
	}

	for _, o := range opts {
		o(s)
	}

	return s
}

// StartTargetChangeHandler changes to targets, targets can be deleted or created
// this function handles the changes to the targets
func (d *DeviationServer) StartTargetChangeHandler() {
	d.log.Debug("Starting subscription gnmi server...")

	for {
		select {
		case tu := <-d.tuCh:
			d.log.Debug("subscription server", "Action", tu.Action, "Target", tu.Name)

			if err := d.HandleTargetUpdate(d.ctx, tu); err != nil {
				d.log.Debug("HandleSubscription", "Error", err)
			}
		case <-d.stopCh:
			d.log.Debug("stopping subscription handler")

		}
	}
}

func (d *DeviationServer) HandleTargetUpdate(ctx context.Context, tu TargetUpdate) error {
	switch tu.Action {
	case TargetAdd:
		// it is possible that during a restart the subscription got removed
		if _, ok := d.Targets[tu.Name]; !ok {
			t := target.NewTarget(tu.TargetConfig)
			d.log.Debug("Target", "Config", tu.TargetConfig, "Target", t)

			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()

			if err := t.CreateGNMIClient(ctx); err != nil {
				d.log.Debug("Error Creating client", "Error", err)
				return errors.Wrap(err, errCreateGnmiClient)
			}
			d.Targets[tu.Name] = &Target{
				log:       d.log,
				Config:    tu.TargetConfig,
				Target:    t,
				StopCh:    make(chan struct{}),
				Collector: NewGNMICollector(t, WithDeviceCollectorLogger(d.log)),
			}

			// start gnmi subscription handler
			go func() {
				d.log.Debug("Target", "TargetName", tu.Name, "Target Info", d.Targets[tu.Name].Target)

				d.Targets[tu.Name].StartGnmiSubscriptionHandler(d.ctx)
				// we should delete the target since an error occured and we will get
				// most likely a device driver got restarted
				delete(d.Targets, tu.Name)

			}()
		}

	case TargetDelete:
		// delete the
		if _, ok := d.Targets[tu.Name]; ok {
			if err := d.Targets[tu.Name].Collector.StopSubscription(ctx, configSubscription); err != nil {
				return err
			}
			d.Targets[tu.Name].StopCh <- struct{}{}
		}

		delete(d.Targets, tu.Name)

	}
	return nil
}

func (t *Target) StartGnmiSubscriptionHandler(ctx context.Context) {
	t.log.Debug("Starting GNMI subscription...", "Target", t.Target.Config.Name)

	t.Collector.Lock()
	go t.Collector.StartSubscription(ctx, t.Config.Name, configSubscription,
		[]*gnmi.Path{{
			Elem: []*gnmi.PathElem{
				{Name: "provider-resource-update"},
			},
		}})
	t.Collector.Unlock()

	chanSubResp, chanSubErr := t.Target.ReadSubscriptions()

	for {
		select {
		case resp := <-chanSubResp:
			//log.Infof("SubRsp Response %v", resp)
			// TODO error handling
			t.ReconcileOnChange(resp.Response)
		case tErr := <-chanSubErr:
			t.log.Debug("subscribe", "error", tErr)
			return
		case <-t.StopCh:
			t.log.Debug("Stopping subscription process...")
			return
		}
	}
}

func (t *Target) ReconcileOnChange(resp *gnmi.SubscribeResponse) error {
	switch resp.GetResponse().(type) {
	case *gnmi.SubscribeResponse_Update:
		// handle deletes
		du := resp.GetUpdate().Delete
		for _, del := range du {
			t.log.Debug("ReconcileOnChange", "Delete", del)
		}

		// handle updates
		u := resp.GetUpdate().Update
		// subscription UPDATE per xpath
		for _, upd := range u {
			t.log.Debug("ReconcileOnChange", "Update", upd)
		}

	case *gnmi.SubscribeResponse_SyncResponse:
		t.log.Debug("SyncResponse")
	}

	return nil
}
