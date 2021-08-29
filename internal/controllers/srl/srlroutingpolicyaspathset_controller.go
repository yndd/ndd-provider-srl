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

	ndrv1 "github.com/netw-device-driver/ndd-core/apis/dvr/v1"
	cfgclient "github.com/netw-device-driver/ndd-grpc/config/client"
	config "github.com/netw-device-driver/ndd-grpc/config/configpb"
	"github.com/netw-device-driver/ndd-grpc/ndd"
	"github.com/netw-device-driver/ndd-runtime/pkg/event"
	"github.com/netw-device-driver/ndd-runtime/pkg/gvk"
	"github.com/netw-device-driver/ndd-runtime/pkg/logging"
	"github.com/netw-device-driver/ndd-runtime/pkg/reconciler/managed"
	"github.com/netw-device-driver/ndd-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-yang/pkg/parser"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	cevent "sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	srlv1 "github.com/yndd/ndd-provider-srl/apis/srl/v1"
)

const (
	// Errors
	errUnexpectedRoutingpolicyAspathset       = "the managed resource is not a RoutingpolicyAspathset resource"
	errKubeUpdateFailedRoutingpolicyAspathset = "cannot update RoutingpolicyAspathset"
	errReadRoutingpolicyAspathset             = "cannot read RoutingpolicyAspathset"
	errCreateRoutingpolicyAspathset           = "cannot create RoutingpolicyAspathset"
	erreUpdateRoutingpolicyAspathset          = "cannot update RoutingpolicyAspathset"
	errDeleteRoutingpolicyAspathset           = "cannot delete RoutingpolicyAspathset"

	// resource information
	levelRoutingpolicyAspathset = 2
	// resourcePrefixRoutingpolicyAspathset = "srl.ndd.yndd.io.v1.RoutingpolicyAspathset"
)

var ResourceRefPathsRoutingpolicyAspathset = []*config.Path{
	{
		Elem: []*config.PathElem{
			{Name: "as-path-set", Key: map[string]string{"name": ""}},
		},
	},
}
var DependencyRoutingpolicyAspathset = []*parser.LeafRef{}
var LocalleafRefRoutingpolicyAspathset = []*parser.LeafRef{}
var ExternalleafRefRoutingpolicyAspathset = []*parser.LeafRef{}

// SetupRoutingpolicyAspathset adds a controller that reconciles RoutingpolicyAspathsets.
func SetupRoutingpolicyAspathset(mgr ctrl.Manager, o controller.Options, l logging.Logger, poll time.Duration, namespace string) (string, chan cevent.GenericEvent, error) {

	name := managed.ControllerName(srlv1.RoutingpolicyAspathsetGroupKind)

	events := make(chan cevent.GenericEvent)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(srlv1.RoutingpolicyAspathsetGroupVersionKind),
		managed.WithExternalConnecter(&connectorRoutingpolicyAspathset{
			log:         l,
			kube:        mgr.GetClient(),
			usage:       resource.NewNetworkNodeUsageTracker(mgr.GetClient(), &ndrv1.NetworkNodeUsage{}),
			newClientFn: cfgclient.NewClient},
		),
		managed.WithParser(l),
		managed.WithValidator(&validatorRoutingpolicyAspathset{log: l, parser: *parser.NewParser(parser.WithLogger(l))}),
		//managed.WithResolver(&resolverRoutingpolicyAspathset{log: l}),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return srlv1.RoutingpolicyAspathsetGroupKind, events, ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1.SrlRoutingpolicyAspathset{}).
		WithEventFilter(resource.IgnoreUpdateWithoutGenerationChangePredicate()).
		Watches(
			&source.Channel{Source: events},
			&handler.EnqueueRequestForObject{},
		).
		//Watches(
		//	&source.Kind{Type: &ndrv1.NetworkNode{}},
		//	handler.EnqueueRequestsFromMapFunc(r.NetworkNodeMapFunc),
		//).
		Complete(r)
}

/*
type resolverRoutingpolicyAspathset struct {
	log logging.Logger
}

func (r *resolverRoutingpolicyAspathset) GetManagedResource(ctx context.Context, resourceName string) (resource.Managed, error) {
	return getManagedResource(resourceName)
}
*/

type validatorRoutingpolicyAspathset struct {
	log    logging.Logger
	parser parser.Parser
}

func (v *validatorRoutingpolicyAspathset) ValidateLocalleafRef(ctx context.Context, mg resource.Managed) (managed.ValidateLocalleafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateLocalleafRef...")

	// json unmarshal the resource
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return managed.ValidateLocalleafRefObservation{}, errors.New(errUnexpectedRoutingpolicyAspathset)
	}
	d, err := json.Marshal(&o.Spec.ForNetworkNode)
	if err != nil {
		return managed.ValidateLocalleafRefObservation{}, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	// For local leafref validation we dont need to supply the external data so we use nil
	success, resultleafRefValidation, err := v.parser.ValidateLeafRef(
		parser.LeafRefValidationLocal, x1, nil, LocalleafRefRoutingpolicyAspathset, log)
	if err != nil {
		return managed.ValidateLocalleafRefObservation{
			Success: false,
		}, nil
	}
	if !success {
		log.Debug("ValidateLocalleafRef failed", "resultleafRefValidation", resultleafRefValidation)
		return managed.ValidateLocalleafRefObservation{
			Success:          false,
			ResolvedLeafRefs: resultleafRefValidation}, nil
	}
	log.Debug("ValidateLocalleafRef success", "resultleafRefValidation", resultleafRefValidation)
	return managed.ValidateLocalleafRefObservation{
		Success:          true,
		ResolvedLeafRefs: resultleafRefValidation}, nil
}

func (v *validatorRoutingpolicyAspathset) ValidateExternalleafRef(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateExternalleafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateExternalleafRef...")

	// json unmarshal the resource
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return managed.ValidateExternalleafRefObservation{}, errors.New(errUnexpectedRoutingpolicyAspathset)
	}
	d, err := json.Marshal(&o.Spec.ForNetworkNode)
	if err != nil {
		return managed.ValidateExternalleafRefObservation{}, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	// json unmarshal the external data
	var x2 interface{}
	json.Unmarshal(cfg, &x2)

	// For local external leafref validation we need to supply the external
	// data to validate the remote leafref, we use x2 for this
	success, resultleafRefValidation, err := v.parser.ValidateLeafRef(
		parser.LeafRefValidationExternal, x1, x2, ExternalleafRefRoutingpolicyAspathset, log)
	if err != nil {
		return managed.ValidateExternalleafRefObservation{
			Success: false,
		}, nil
	}
	if !success {
		log.Debug("ValidateExternalLeafRef failed", "resultleafRefValidation", resultleafRefValidation)
		return managed.ValidateExternalleafRefObservation{
			Success:          false,
			ResolvedLeafRefs: resultleafRefValidation}, nil
	}
	log.Debug("ValidateExternalLeafRef success", "resultleafRefValidation", resultleafRefValidation)
	return managed.ValidateExternalleafRefObservation{
		Success:          true,
		ResolvedLeafRefs: resultleafRefValidation}, nil
}

func (v *validatorRoutingpolicyAspathset) ValidateParentDependency(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateParentDependencyObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateParentDependency...")

	// we initialize a global list for finer information on the resolution
	resultleafRefValidation := make([]*parser.ResolvedLeafRef, 0)
	log.Debug("ValidateParentDependency success", "resultParentValidation", resultleafRefValidation)
	return managed.ValidateParentDependencyObservation{
		Success:          true,
		ResolvedLeafRefs: resultleafRefValidation}, nil
}

// ValidateResourceIndexes validates if the indexes of a resource got changed
// if so we need to delete the original resource, because it will be dangling if we dont delete it
func (v *validatorRoutingpolicyAspathset) ValidateResourceIndexes(ctx context.Context, mg resource.Managed) (managed.ValidateResourceIndexesObservation, error) {
	log := v.log.WithValues("resosurce", mg.GetName())

	// json unmarshal the resource
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return managed.ValidateResourceIndexesObservation{}, errors.New(errUnexpectedInterface)
	}
	log.Debug("ValidateResourceIndexes", "Spec", o.Spec)

	rootPath := &config.Path{
		Elem: []*config.PathElem{
			{Name: "routing-policy"},
			{Name: "as-path-set", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlRoutingpolicyAspathset.Name}},
		},
	}

	origResourceIndex := mg.GetResourceIndexes()
	// we call the CompareConfigPathsWithResourceKeys irrespective is the get resource index returns nil
	changed, deletPaths, newResourceIndex := v.parser.CompareConfigPathsWithResourceKeys(rootPath, origResourceIndex)
	if changed {
		log.Debug("ValidateResourceIndexes changed", "deletPaths", deletPaths[0])
		return managed.ValidateResourceIndexesObservation{Changed: true, ResourceDeletes: deletPaths, ResourceIndexes: newResourceIndex}, nil
	}

	log.Debug("ValidateResourceIndexes success")
	return managed.ValidateResourceIndexesObservation{Changed: false, ResourceIndexes: newResourceIndex}, nil
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connectorRoutingpolicyAspathset struct {
	log         logging.Logger
	kube        client.Client
	usage       resource.Tracker
	newClientFn func(ctx context.Context, cfg ndd.Config) (config.ConfigurationClient, error)
}

// Connect produces an ExternalClient by:
// 1. Tracking that the managed resource is using a NetworkNode.
// 2. Getting the managed resource's NetworkNode with connection details
// A resource is mapped to a single target
func (c *connectorRoutingpolicyAspathset) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	log := c.log.WithValues("resource", mg.GetName())
	log.Debug("Connect")
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return nil, errors.New(errUnexpectedRoutingpolicyAspathset)
	}
	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackTCUsage)
	}

	// find network node that is configured status
	nn := &ndrv1.NetworkNode{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: o.GetNetworkNodeReference().Name}, nn); err != nil {
		return nil, errors.Wrap(err, errGetNetworkNode)
	}

	if nn.GetCondition(ndrv1.ConditionKindDeviceDriverConfigured).Status != corev1.ConditionTrue {
		return nil, errors.New(targetNotConfigured)
	}

	cfg := ndd.Config{
		SkipVerify: true,
		Insecure:   true,
		Target:     ndrv1.PrefixService + "-" + nn.Name + "." + ndrv1.NamespaceLocalK8sDNS + strconv.Itoa(*nn.Spec.GrpcServerPort),
	}
	log.Debug("Client config", "config", cfg)

	cl, err := c.newClientFn(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	// we make a string here since we use a trick in registration to go to multiple targets
	// while here the object is mapped to a single target/network node
	tns := make([]string, 0)
	tns = append(tns, nn.GetName())

	log.Debug("Client info", "client", cl)

	return &externalRoutingpolicyAspathset{client: cl, targets: tns, log: log, parser: *parser.NewParser(parser.WithLogger(log))}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type externalRoutingpolicyAspathset struct {
	client  config.ConfigurationClient
	targets []string
	log     logging.Logger
	parser  parser.Parser
}

func (e *externalRoutingpolicyAspathset) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedRoutingpolicyAspathset)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Observing ...")

	rootPath := &config.Path{
		Elem: []*config.PathElem{
			{Name: "routing-policy"},
			{Name: "as-path-set", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlRoutingpolicyAspathset.Name}},
		},
	}

	gvk := &gvk.GVK{
		Group:     mg.GetObjectKind().GroupVersionKind().Group,
		Version:   mg.GetObjectKind().GroupVersionKind().Version,
		Kind:      mg.GetObjectKind().GroupVersionKind().Kind,
		Name:      mg.GetName(),
		NameSpace: mg.GetNamespace(),
	}
	gvkstring, err := gvk.String()
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	resp, err := e.client.Get(ctx, &config.ResourceKey{
		Name: gvkstring,
		//Name: resourcePrefixRoutingpolicyAspathset + "." + o.GetName(),
		Level: levelRoutingpolicyAspathset,
		Path:  rootPath,
	})
	if err != nil {
		return managed.ExternalObservation{}, errors.New(errReadRoutingpolicyAspathset)
	}

	if !resp.Exists {
		// Resource Does not Exists
		if resp.Data != nil {
			// this is an umnaged resource which has data and will be moved to an unmanaged resource

			d, err := json.Marshal(&o.Spec.ForNetworkNode)
			if err != nil {
				return managed.ExternalObservation{}, errors.Wrap(err, errJSONMarshal)
			}

			var x1 interface{}
			if err := json.Unmarshal(d, &x1); err != nil {
				return managed.ExternalObservation{}, errors.Wrap(err, errJSONUnMarshal)
			}
			log.Debug("Spec Data Before", "X1", x1)

			// remove the hierarchical elements for data processing, comparison, etc
			// they are used in the provider for parent dependency resolution
			// but are not relevant in the data, they are referenced in the rootPath
			// when interacting with the device driver
			hids := make([]string, 0)
			x1 = e.parser.RemoveLeafsFromJSONData(x1, hids)

			var x2 interface{}
			if err := json.Unmarshal(resp.Data, &x2); err != nil {
				return managed.ExternalObservation{}, errors.Wrap(err, errJSONUnMarshal)
			}

			log.Debug("Spec Data", "X1", x1)
			log.Debug("Resp Data", "X2", x2)
			// for lists with keys we need to create a list before calulating the paths since this is what
			// the object eventually happens to be based upon. We avoid having multiple entries in a list object
			// and hence we have to add this step
			x1, err = e.parser.AddJSONDataToList(x1)
			if err != nil {
				return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
			}

			updatesx1 := e.parser.GetUpdatesFromJSONData(rootPath, e.parser.XpathToConfigGnmiPath("/", 0), x1, ResourceRefPathsRoutingpolicyAspathset)
			for _, update := range updatesx1 {
				log.Debug("Observe Fine Grane Updates X1", "Path", update.Path, "Value", string(update.Value))
			}
			// for lists with keys we need to create a list before calulating the paths since this is what
			// the object eventually happens to be based upon. We avoid having multiple entries in a list object
			// and hence we have to add this step
			x2, err = e.parser.AddJSONDataToList(x2)
			if err != nil {
				return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
			}
			updatesx2 := e.parser.GetUpdatesFromJSONData(rootPath, e.parser.XpathToConfigGnmiPath("/", 0), x2, ResourceRefPathsRoutingpolicyAspathset)
			for _, update := range updatesx2 {
				log.Debug("Observe Fine Grane Updates X2", "Path", update.Path, "Value", string(update.Value))
			}

			deletes, updates, err := e.parser.FindResourceDelta(updatesx1, updatesx2, log)
			if err != nil {
				return managed.ExternalObservation{}, err
			}
			if len(deletes) != 0 || len(updates) != 0 {
				// resource is NOT up to date
				log.Debug("Observing resource not up to date", "Updates", updates, "Deletes", deletes)
				log.Debug("Observing  Respone", "Exists", false, "HasData", true, "UpToDate", false, "Response", resp)
				return managed.ExternalObservation{
					ResourceExists:   false,
					ResourceHasData:  true,
					ResourceUpToDate: false,
					ResourceDeletes:  deletes,
					ResourceUpdates:  updates,
				}, nil
			}
			// resource is up to date
			log.Debug("Observing  Respone", "Exists", false, "HasData", true, "UpToDate", true, "Response", resp)
			return managed.ExternalObservation{
				ResourceExists:   false,
				ResourceHasData:  true,
				ResourceUpToDate: true,
			}, nil
		} else {
			log.Debug("Observing  Respone", "Exists", false, "HasData", false, "UpToDate", false, "Response", resp)
			return managed.ExternalObservation{
				ResourceExists:   false,
				ResourceHasData:  false,
				ResourceUpToDate: false,
			}, nil
		}
	} else {
		// Resource Exists
		switch resp.Status {
		case config.Status_Success:
			if resp.Data != nil {
				// data is present
				d, err := json.Marshal(&o.Spec.ForNetworkNode)
				if err != nil {
					return managed.ExternalObservation{}, errors.Wrap(err, errJSONMarshal)
				}

				var x1 interface{}
				if err := json.Unmarshal(d, &x1); err != nil {
					return managed.ExternalObservation{}, errors.Wrap(err, errJSONUnMarshal)
				}
				log.Debug("Spec Data Before", "X1", x1)

				// remove the hierarchical elements for data processing, comparison, etc
				// they are used in the provider for parent dependency resolution
				// but are not relevant in the data, they are referenced in the rootPath
				// when interacting with the device driver
				hids := make([]string, 0)
				x1 = e.parser.RemoveLeafsFromJSONData(x1, hids)

				var x2 interface{}
				if err := json.Unmarshal(resp.Data, &x2); err != nil {
					return managed.ExternalObservation{}, errors.Wrap(err, errJSONUnMarshal)
				}

				log.Debug("Spec Data", "X1", x1)
				log.Debug("Resp Data", "X2", x2)
				// for lists with keys we need to create a list before calulating the paths since this is what
				// the object eventually happens to be based upon. We avoid having multiple entries in a list object
				// and hence we have to add this step
				x1, err = e.parser.AddJSONDataToList(x1)
				if err != nil {
					return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
				}

				updatesx1 := e.parser.GetUpdatesFromJSONData(rootPath, e.parser.XpathToConfigGnmiPath("/", 0), x1, ResourceRefPathsRoutingpolicyAspathset)
				for _, update := range updatesx1 {
					log.Debug("Observe Fine Grane Updates X1", "Path", update.Path, "Value", string(update.Value))
				}
				// for lists with keys we need to create a list before calulating the paths since this is what
				// the object eventually happens to be based upon. We avoid having multiple entries in a list object
				// and hence we have to add this step
				x2, err = e.parser.AddJSONDataToList(x2)
				if err != nil {
					return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
				}
				updatesx2 := e.parser.GetUpdatesFromJSONData(rootPath, e.parser.XpathToConfigGnmiPath("/", 0), x2, ResourceRefPathsRoutingpolicyAspathset)
				for _, update := range updatesx2 {
					log.Debug("Observe Fine Grane Updates X2", "Path", update.Path, "Value", string(update.Value))
				}

				deletes, updates, err := e.parser.FindResourceDelta(updatesx1, updatesx2, log)
				if err != nil {
					return managed.ExternalObservation{}, err
				}
				if len(deletes) != 0 || len(updates) != 0 {
					// resource is NOT up to date
					log.Debug("Observind resource not up to date", "Updates", updates, "Deletes", deletes)
					log.Debug("Observing  Respone", "Exists", true, "HasData", true, "UpToDate", false, "Response", resp)
					return managed.ExternalObservation{
						ResourceExists:   true,
						ResourceHasData:  true,
						ResourceUpToDate: false,
						ResourceDeletes:  deletes,
						ResourceUpdates:  updates,
					}, nil
				}
				// resource is up to date
				log.Debug("Observing  Respone", "Exists", true, "HasData", true, "UpToDate", true, "Response", resp)
				return managed.ExternalObservation{
					ResourceExists:   true,
					ResourceHasData:  true,
					ResourceUpToDate: true,
				}, nil
			} else {
				log.Debug("Observing  Respone", "Exists", true, "HasData", false, "UpToDate", false, "Status", resp.Status)
				return managed.ExternalObservation{
					ResourceExists:   true,
					ResourceHasData:  false,
					ResourceUpToDate: false,
				}, nil
			}
		default:
			log.Debug("Observing  Respone", "Exists", true, "HasData", false, "UpToDate", false, "Status", resp.Status)
			return managed.ExternalObservation{
				ResourceExists:   true,
				ResourceHasData:  false,
				ResourceUpToDate: false,
			}, nil
		}
	}
}

func (e *externalRoutingpolicyAspathset) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedRoutingpolicyAspathset)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Creating ...")

	rootPath := &config.Path{
		Elem: []*config.PathElem{
			{Name: "routing-policy"},
			{Name: "as-path-set", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlRoutingpolicyAspathset.Name}},
		},
	}

	d, err := json.Marshal(&o.Spec.ForNetworkNode)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errJSONMarshal)
	}

	var x1 interface{}
	if err := json.Unmarshal(d, &x1); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errJSONUnMarshal)
	}

	// remove the hierarchical elements for data processing, comparison, etc
	// they are used in the provider for parent dependency resolution
	// but are not relevant in the data, they are referenced in the rootPath
	// when interacting with the device driver
	hids := make([]string, 0)
	x1 = e.parser.RemoveLeafsFromJSONData(x1, hids)
	// for lists with keys we need to create a list before calulating the paths since this is what
	// the object eventually happens to be based upon. We avoid having multiple entries in a list object
	// and hence we have to add this step
	x1, err = e.parser.AddJSONDataToList(x1)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errWrongInputdata)
	}

	updates := e.parser.GetUpdatesFromJSONData(rootPath, e.parser.XpathToConfigGnmiPath("/", 0), x1, ResourceRefPathsRoutingpolicyAspathset)
	for _, update := range updates {
		log.Debug("Create Fine Grane Updates", "Path", update.Path, "Value", update.Value)
	}

	gvk := &gvk.GVK{
		Group:     mg.GetObjectKind().GroupVersionKind().Group,
		Version:   mg.GetObjectKind().GroupVersionKind().Version,
		Kind:      mg.GetObjectKind().GroupVersionKind().Kind,
		Name:      mg.GetName(),
		NameSpace: mg.GetNamespace(),
	}
	gvkstring, err := gvk.String()
	if err != nil {
		return managed.ExternalCreation{}, err
	}

	_, err = e.client.Create(ctx, &config.Request{
		Name: gvkstring,
		//Name:  resourcePrefixRoutingpolicyAspathset + "." + o.GetName(),
		Level:  levelRoutingpolicyAspathset,
		Path:   rootPath,
		Data:   d, // depreciated and can be removed later
		Update: updates,
	})
	if err != nil {
		return managed.ExternalCreation{}, errors.New(errReadRoutingpolicyAspathset)
	}

	return managed.ExternalCreation{}, nil
}

func (e *externalRoutingpolicyAspathset) Update(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) (managed.ExternalUpdate, error) {
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedRoutingpolicyAspathset)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Updating ...")

	for _, u := range obs.ResourceUpdates {
		log.Debug("Update -> Update", "Path", u.Path, "Value", string(u.Value))
	}
	for _, d := range obs.ResourceDeletes {
		log.Debug("Update -> Delete", "Path", d)
	}

	gvk := &gvk.GVK{
		Group:     mg.GetObjectKind().GroupVersionKind().Group,
		Version:   mg.GetObjectKind().GroupVersionKind().Version,
		Kind:      mg.GetObjectKind().GroupVersionKind().Kind,
		Name:      mg.GetName(),
		NameSpace: mg.GetNamespace(),
	}
	gvkstring, err := gvk.String()
	if err != nil {
		return managed.ExternalUpdate{}, err
	}

	_, err = e.client.Update(ctx, &config.Notification{
		Name: gvkstring,
		//Name:  resourcePrefixInterface + "." + o.GetName(),
		Level:  levelInterface,
		Delete: obs.ResourceDeletes,
		Update: obs.ResourceUpdates,
	})
	if err != nil {
		return managed.ExternalUpdate{}, errors.New(errReadInterface)
	}

	return managed.ExternalUpdate{}, nil
}

func (e *externalRoutingpolicyAspathset) Delete(ctx context.Context, mg resource.Managed) error {
	o, ok := mg.(*srlv1.SrlRoutingpolicyAspathset)
	if !ok {
		return errors.New(errUnexpectedRoutingpolicyAspathset)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Deleting ...")

	rootPath := &config.Path{
		Elem: []*config.PathElem{
			{Name: "routing-policy"},
			{Name: "as-path-set", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlRoutingpolicyAspathset.Name}},
		},
	}

	gvk := &gvk.GVK{
		Group:     mg.GetObjectKind().GroupVersionKind().Group,
		Version:   mg.GetObjectKind().GroupVersionKind().Version,
		Kind:      mg.GetObjectKind().GroupVersionKind().Kind,
		Name:      mg.GetName(),
		NameSpace: mg.GetNamespace(),
	}
	gvkstring, err := gvk.String()
	if err != nil {
		return err
	}

	_, err = e.client.Delete(ctx, &config.ResourceKey{
		Name: gvkstring,
		//Name: resourcePrefixRoutingpolicyAspathset + "." + o.GetName(),
		Level: levelRoutingpolicyAspathset,
		Path:  rootPath,
	})
	if err != nil {
		return errors.New(errDeleteRoutingpolicyAspathset)
	}

	return nil
}

func (e *externalRoutingpolicyAspathset) GetTarget() []string {
	return e.targets
}

func (e *externalRoutingpolicyAspathset) GetConfig(ctx context.Context) ([]byte, error) {
	resp, err := e.client.GetConfig(ctx, &config.ConfigRequest{})
	if err != nil {
		return make([]byte, 0), errors.Wrap(err, "err get config")
	}
	return resp.Data, nil
}

func (e *externalRoutingpolicyAspathset) GetResourceName(ctx context.Context, path *config.Path) (string, error) {
	resp, err := e.client.GetResourceName(ctx, &config.ResourceRequest{Path: path})
	if err != nil {
		return "", errors.Wrap(err, "err get resourceName")
	}
	return resp.GetName(), nil
}
