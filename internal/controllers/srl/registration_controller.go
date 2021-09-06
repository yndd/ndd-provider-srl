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

package srl

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/yndd/ndd-provider-srl/internal/collector"
	"github.com/yndd/ndd-yang/pkg/parser"

	"github.com/karimra/gnmic/target"
	"github.com/karimra/gnmic/types"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/reconciler/managed"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-runtime/pkg/utils"
	corev1 "k8s.io/api/core/v1"

	srlv1 "github.com/yndd/ndd-provider-srl/apis/srl/v1"
)

const (
	// RegistrationFinalizer defines the finalizer for Registration
	RegistrationFinalizer = "Registration.srl.ndd.yndd.io"

	// Errors
	errUnexpectedRegistration       = "the managed resource is not a Registration resource"
	errKubeUpdateRegistrationFailed = "cannot update Registration"
	errRegistrationGet              = "cannot get Registration"
	errRegistrationCreate           = "cannot create Registration"
	errRegistrationUpdate           = "cannot update Registration"
	errRegistrationDelete           = "cannot delete Registration"
)

// SetupRegistration adds a controller that reconciles Registrations.
func SetupRegistration(mgr ctrl.Manager, o controller.Options, l logging.Logger, poll time.Duration, namespace string, subChan chan collector.TargetUpdate) error {

	name := managed.ControllerName(srlv1.RegistrationGroupKind)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(srlv1.RegistrationGroupVersionKind),
		managed.WithExternalConnecter(&connectorRegistration{
			log:     l,
			subChan: subChan,
			kube:    mgr.GetClient(),
			usage:   resource.NewNetworkNodeUsageTracker(mgr.GetClient(), &ndrv1.NetworkNodeUsage{}),
			//newClientFn: regclient.NewClient},
			newClientFn: target.NewTarget},
		),
		managed.WithValidator(&validatorRegistration{log: l}),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1.Registration{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		//Watches(
		//	&source.Kind{Type: &ndrv1.NetworkNode{}},
		//	handler.EnqueueRequestsFromMapFunc(r.NetworkNodeMapFunc),
		//).
		Complete(r)
}

type validatorRegistration struct {
	log logging.Logger
}

func (v *validatorRegistration) ValidateLocalleafRef(ctx context.Context, mg resource.Managed) (managed.ValidateLocalleafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateLocalleafRef success")
	return managed.ValidateLocalleafRefObservation{Success: true}, nil
}

func (v *validatorRegistration) ValidateExternalleafRef(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateExternalleafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateExternalleafRef success")
	return managed.ValidateExternalleafRefObservation{Success: true}, nil
}

func (v *validatorRegistration) ValidateParentDependency(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateParentDependencyObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateParentDependency success")
	return managed.ValidateParentDependencyObservation{Success: true}, nil
}

func (v *validatorRegistration) ValidateResourceIndexes(ctx context.Context, mg resource.Managed) (managed.ValidateResourceIndexesObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateResourceIndexes success")
	return managed.ValidateResourceIndexesObservation{ResourceDeletes: make([]*gnmi.Path, 0)}, nil
}

// A connectorRegistration is expected to produce an ExternalClient when its Connect method
// is called.
type connectorRegistration struct {
	log         logging.Logger
	subChan     chan collector.TargetUpdate
	kube        client.Client
	usage       resource.Tracker
	newClientFn func(c *types.TargetConfig) *target.Target
}

// Connect produces an ExternalClient by:
// 1. Tracking that the managed resource is using a NetworkNode Reference.
// 2. Getting the managed resource's NetworkNode with connection details
// For registration we did a trick to use aall network nodes in the system since
// we want to register to all nodes, this is an exception for registration
func (c *connectorRegistration) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	log := c.log.WithValues("resource", mg.GetName())
	log.Debug("Connect")
	o, ok := mg.(*srlv1.Registration)
	if !ok {
		return nil, errors.New(errUnexpectedRegistration)
	}
	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackTCUsage)
	}

	selectors := []client.ListOption{}
	nnl := &ndrv1.NetworkNodeList{}
	if err := c.kube.List(ctx, nnl, selectors...); err != nil {
		return nil, errors.Wrap(err, errGetNetworkNode)
	}

	// find all targets that have are in configured status
	var ts []*nddv1.Target
	for _, nn := range nnl.Items {
		log.Debug("Network Node", "Name", nn.GetName(), "Status", nn.GetCondition(ndrv1.ConditionKindDeviceDriverConfigured).Status)
		if nn.GetCondition(ndrv1.ConditionKindDeviceDriverConfigured).Status == corev1.ConditionTrue {
			t := &nddv1.Target{
				Name: nn.GetName(),
				Config: &types.TargetConfig{
					Name:       nn.GetName(),
					Address:    ndrv1.PrefixService + "-" + nn.Name + "." + ndrv1.NamespaceLocalK8sDNS + strconv.Itoa(*nn.Spec.GrpcServerPort),
					Username:   utils.StringPtr("admin"),
					Password:   utils.StringPtr("admin"),
					Timeout:    10 * time.Second,
					SkipVerify: utils.BoolPtr(true),
					Insecure:   utils.BoolPtr(true),
					TLSCA:      utils.StringPtr(""), //TODO TLS
					TLSCert:    utils.StringPtr(""), //TODO TLS
					TLSKey:     utils.StringPtr(""),
					Gzip:       utils.BoolPtr(false),
				},
			}
			ts = append(ts, t)
		}
	}
	log.Debug("Active targets", "targets", ts)

	// Validate if targets got added or deleted, based on this information the subscription Server
	// should be informed over the channel
	// check for deletes
	deletedTargets := make([]collector.TargetUpdate, 0)
	for _, origTarget := range o.Status.Target {
		found := false
		for _, newTarget := range ts {
			if origTarget == newTarget.Name {
				found = true
			}
		}
		if !found {
			deletedTargets = append(deletedTargets, collector.TargetUpdate{
				Name:   origTarget,
				Action: collector.TargetDelete,
			})
		}
	}
	// udate all targets
	allTargets := make([]collector.TargetUpdate, 0)
	for _, allTarget := range ts {
		allTargets = append(allTargets, collector.TargetUpdate{
			Name:   allTarget.Name,
			Action: collector.TargetAdd,
			TargetConfig: &types.TargetConfig{
				Name:       allTarget.Name,
				Address:    allTarget.Config.Address,
				Username:   utils.StringPtr("admin"),
				Password:   utils.StringPtr("admin"),
				Timeout:    10 * time.Second,
				SkipVerify: utils.BoolPtr(true),
				Insecure:   utils.BoolPtr(true),
				TLSCA:      utils.StringPtr(""), //TODO TLS
				TLSCert:    utils.StringPtr(""), //TODO TLS
				TLSKey:     utils.StringPtr(""),
				Gzip:       utils.BoolPtr(false),
			},
		})
	}

	for _, sub := range deletedTargets {
		log.Debug("Stop Subscription", "target", sub.Name)
		c.subChan <- sub
	}
	for _, sub := range allTargets {
		log.Debug("Start Subscription", "target", sub.Name)
		c.subChan <- sub
	}

	// when no targets are found we return a not found error
	// this unifies the reconcile code when a dedicate network node is looked up
	if len(ts) == 0 {
		return nil, errors.New(errNoTargetFound)
	}

	//get clients for each target
	cls := make([]*target.Target, 0)
	tns := make([]string, 0)
	for _, t := range ts {
		cl := target.NewTarget(t.Config)
		if err := cl.CreateGNMIClient(ctx); err != nil {
			return nil, errors.Wrap(err, errNewClient)
		}
		cls = append(cls, cl)
		tns = append(tns, t.Name)
		/*
			cl, err := c.newClientFn(t.Cfg)
			if err != nil {
				return nil, errors.Wrap(err, errNewClient)
			}
			cls = append(cls, cl)
			tns = append(tns, t.Name)
		*/
	}

	//get clients for each target -> config target
	/*
		cls := make([]register.RegistrationClient, 0)
		tns := make([]string, 0)
		for _, t := range ts {
			cl, err := c.newClientFn(ctx, t.Cfg)
			if err != nil {
				return nil, errors.Wrap(err, errNewClient)
			}
			cls = append(cls, cl)
			tns = append(tns, t.Name)
		}
	*/

	log.Debug("Connect info", "clients", cls, "targets", tns)

	return &externalRegistration{clients: cls, targets: tns, log: log, parser: *parser.NewParser(parser.WithLogger(log))}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
/*
type externalRegistration struct {
	clients []register.RegistrationClient
	targets []string
	log     logging.Logger
}
*/
type externalRegistration struct {
	clients []*target.Target
	targets []string
	parser  parser.Parser
	log     logging.Logger
}

func (e *externalRegistration) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	o, ok := mg.(*srlv1.Registration)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedRegistration)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Observing ...")

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{Name: nddv1.RegisterPathElemName, Key: map[string]string{nddv1.RegisterPathElemKey: string(srlv1.DeviceTypeSRL)}},
			},
		},
	}
	//d, err := json.Marshal(o.Spec.ForNetworkNode)
	//if err != nil {
	//	return managed.ExternalObservation{}, errors.Wrap(err, errJSONMarshal)
	//}
	req := &gnmi.GetRequest{
		Path:     path,
		Encoding: gnmi.Encoding_JSON,
	}

	for _, cl := range e.clients {
		rsp, err := cl.Get(ctx, req)
		if err != nil {
			// if a single network device driver reports an error this is applicable to all
			// network devices
			return managed.ExternalObservation{}, errors.Wrap(err, errRegistrationGet)
		}
		// if a network device driver reports a different device type we trigger
		// a recreation of the configuration on all devices by returning
		// Exists = false and
		log.Debug("Observing response", "Response", rsp)
		if deviceType, ok := rsp.GetNotification()[0].GetUpdate()[0].GetPath().GetElem()[0].GetKey()[nddv1.RegisterPathElemKey]; ok {
			log.Debug("Observing response", "Data", deviceType)
			if nddv1.DeviceType(deviceType) != srlv1.DeviceTypeSRL {
				return managed.ExternalObservation{
					ResourceExists:   false,
					ResourceUpToDate: false,
					ResourceHasData:  false,
					Ready:            true,
				}, nil
			}
		}

	}
	/*
		for _, cl := range e.clients {
			r, err := cl.Get(ctx, &register.DeviceType{
				DeviceType: string(srlv1.DeviceTypeSRL),
			})
			if err != nil {
				// if a single network device driver reports an error this is applicable to all
				// network devices
				return managed.ExternalObservation{}, errors.New(errRegistrationGet)
			}
			// if a network device driver reports a different device type we trigger
			// a recreation of the configuration on all devices by returning
			// Exists = false and
			if r.DeviceType != string(srlv1.DeviceTypeSRL) {
				return managed.ExternalObservation{
					ResourceExists:   false,
					ResourceUpToDate: false,
					ResourceHasData:  false,
				}, nil
			}
		}
	*/

	// when all network device driver reports the proper device type
	// we return exists and up to date
	return managed.ExternalObservation{
		Ready:            true,
		ResourceExists:   true,
		ResourceUpToDate: true,
		ResourceHasData:  true, // we fake that we have data since it is not relevant
	}, nil

}

func (e *externalRegistration) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	o, ok := mg.(*srlv1.Registration)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedRegistration)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Creating ...")

	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: nddv1.RegisterPathElemName, Key: map[string]string{nddv1.RegisterPathElemKey: string(srlv1.DeviceTypeSRL)}},
		},
	}
	d, err := json.Marshal(o.Spec.ForNetworkNode)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errJSONMarshal)
	}
	req := &gnmi.SetRequest{
		Replace: []*gnmi.Update{
			{
				Path: path,
				Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: d}},
			},
		},
	}
	for _, cl := range e.clients {
		_, err := cl.Set(ctx, req)
		if err != nil {
			return managed.ExternalCreation{}, errors.New(errRegistrationCreate)
		}
	}

	/*
		for _, cl := range e.clients {
			_, err := cl.Create(ctx, &register.Request{
				DeviceType:             string(srlv1.DeviceTypeSRL),
				MatchString:            srlv1.DeviceMatch,
				Subscriptions:          o.GetSubscriptions(),
				ExceptionPaths:         o.GetExceptionPaths(),
				ExplicitExceptionPaths: o.GetExplicitExceptionPaths(),
			})
			if err != nil {
				return managed.ExternalCreation{}, errors.New(errRegistrationCreate)
			}
		}
	*/

	return managed.ExternalCreation{}, nil
}

func (e *externalRegistration) Update(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) (managed.ExternalUpdate, error) {
	o, ok := mg.(*srlv1.Registration)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedRegistration)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Updating ...")

	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: nddv1.RegisterPathElemName, Key: map[string]string{nddv1.RegisterPathElemKey: string(srlv1.DeviceTypeSRL)}},
		},
	}
	d, err := json.Marshal(o.Spec.ForNetworkNode)
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errJSONMarshal)
	}
	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path,
				Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_JsonVal{JsonVal: d}},
			},
		},
	}
	for _, cl := range e.clients {
		_, err := cl.Set(ctx, req)
		if err != nil {
			return managed.ExternalUpdate{}, errors.New(errRegistrationUpdate)
		}
	}

	/*
		for _, cl := range e.clients {
			_, err := cl.Update(ctx, &register.Request{
				DeviceType:             string(srlv1.DeviceTypeSRL),
				MatchString:            srlv1.DeviceMatch,
				Subscriptions:          o.GetSubscriptions(),
				ExceptionPaths:         o.GetExceptionPaths(),
				ExplicitExceptionPaths: o.GetExplicitExceptionPaths(),
			})
			if err != nil {
				return managed.ExternalUpdate{}, errors.New(errRegistrationUpdate)
			}
		}
	*/
	return managed.ExternalUpdate{}, nil
}

func (e *externalRegistration) Delete(ctx context.Context, mg resource.Managed) error {
	o, ok := mg.(*srlv1.Registration)
	if !ok {
		return errors.New(errUnexpectedRegistration)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Deleting ...")

	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: nddv1.RegisterPathElemName, Key: map[string]string{nddv1.RegisterPathElemKey: string(srlv1.DeviceTypeSRL)}},
		},
	}
	paths := make([]*gnmi.Path, 0)
	paths = append(paths, path)
	req := &gnmi.SetRequest{
		Delete: paths,
	}
	for _, cl := range e.clients {
		_, err := cl.Set(ctx, req)
		if err != nil {
			return errors.New(errRegistrationDelete)
		}
	}

	/*
		for _, cl := range e.clients {
			_, err := cl.Delete(ctx, &register.DeviceType{
				DeviceType: string(srlv1.DeviceTypeSRL),
			})
			if err != nil {
				return errors.New(errRegistrationDelete)
			}
		}
	*/
	return nil
}

func (e *externalRegistration) GetTarget() []string {
	return e.targets
}

func (e *externalRegistration) GetConfig(ctx context.Context) ([]byte, error) {
	return make([]byte, 0), nil
}

func (e *externalRegistration) GetResourceName(ctx context.Context, path []*gnmi.Path) (string, error) {
	return "", nil
}
