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
	"sync"
	"time"

	"github.com/karimra/gnmic/target"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/logging"
)

const (
	defaultTargetReceivebuffer = 1000
	defaultLockRetry           = 5 * time.Second
	defaultRetryTimer          = 10 * time.Second

	// errors
	errCreateSubscriptionRequest = "cannot create subscription request"
)

// Collector defines the interfaces for the collector
type Collector interface {
	Lock()
	Unlock()
	GetSubscription(subName string) bool
	StopSubscription(ctx context.Context, sub string) error
	StartSubscription(ctx context.Context, subName string, sub []string) error
}

// DeviceCollectorOption can be used to manipulate Options.
type DeviceCollectorOption func(*GNMICollector)

// WithTargetLogger specifies how the object should log messages.
func WithDeviceCollectorLogger(log logging.Logger) DeviceCollectorOption {
	return func(o *GNMICollector) {
		o.log = log
	}
}

// GNMICollector defines the parameters for the collector
type GNMICollector struct {
	TargetReceiveBuffer uint
	RetryTimer          time.Duration
	Target              *target.Target
	//targetSubRespChan   chan *collector.SubscribeResponse
	//targetSubErrChan    chan *collector.TargetError
	Subscriptions map[string]*Subscription
	Mutex         sync.RWMutex
	log           logging.Logger
}

// Subscription defines the parameters for the subscription
type Subscription struct {
	StopCh   chan bool
	CancelFn context.CancelFunc
}

// NewGNMICollector creates a new GNMI collector
func NewGNMICollector(t *target.Target, opts ...DeviceCollectorOption) *GNMICollector {
	c := &GNMICollector{
		Target:              t,
		Subscriptions:       make(map[string]*Subscription),
		Mutex:               sync.RWMutex{},
		TargetReceiveBuffer: defaultTargetReceivebuffer,
		RetryTimer:          defaultRetryTimer,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Lock locks a gnmi collector
func (c *GNMICollector) Lock() {
	c.Mutex.RLock()
}

// Unlock unlocks a gnmi collector
func (c *GNMICollector) Unlock() {
	c.Mutex.RUnlock()
}

// GetSubscription returns a bool based on a subscription name
func (c *GNMICollector) GetSubscription(subName string) bool {
	if _, ok := c.Subscriptions[subName]; !ok {
		return true
	}
	return false
}

// StopSubscription stops a subscription
func (c *GNMICollector) StopSubscription(ctx context.Context, sub string) error {
	c.log.WithValues("subscription", sub)
	c.log.Debug("subscription stop...")
	c.Subscriptions[sub].StopCh <- true // trigger quit

	c.log.Debug("subscription stopped")
	return nil
}

// StartSubscription starts a subscription
func (c *GNMICollector) StartSubscription(dctx context.Context, target, subName string, paths []*gnmi.Path) error {
	log := c.log.WithValues("subscription", subName, "Paths", paths)
	log.Debug("subscription start...")
	// initialize new subscription
	ctx, cancel := context.WithCancel(dctx)

	c.Subscriptions[subName] = &Subscription{
		StopCh:   make(chan bool),
		CancelFn: cancel,
	}

	req, err := CreateSubscriptionRequest(target, subName, paths)
	if err != nil {
		c.log.Debug(errCreateSubscriptionRequest, "error", err)
		return errors.Wrap(err, errCreateSubscriptionRequest)
	}

	go func() {
		c.Target.Subscribe(ctx, req, subName)
	}()
	log.Debug("subscription started ...")

	for {
		select {
		case <-c.Subscriptions[subName].StopCh: // execute quit
			c.Subscriptions[subName].CancelFn()
			c.Mutex.Lock()
			delete(c.Subscriptions, subName)
			c.Mutex.Unlock()
			c.log.Debug("subscription cancelled")
			return nil
		}
	}
}
