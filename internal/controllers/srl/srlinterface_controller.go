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

	"github.com/karimra/gnmic/target"
	gnmitypes "github.com/karimra/gnmic/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/pkg/errors"
	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/event"
	"github.com/yndd/ndd-runtime/pkg/gext"
	"github.com/yndd/ndd-runtime/pkg/gvk"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/ndd-runtime/pkg/reconciler/managed"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"github.com/yndd/ndd-runtime/pkg/utils"
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
	errUnexpectedInterface       = "the managed resource is not a Interface resource"
	errKubeUpdateFailedInterface = "cannot update Interface"
	errReadInterface             = "cannot read Interface"
	errCreateInterface           = "cannot create Interface"
	erreUpdateInterface          = "cannot update Interface"
	errDeleteInterface           = "cannot delete Interface"

	// resource information
	levelInterface = 1
	// resourcePrefixInterface = "srl.ndd.yndd.io.v1.Interface"
)

var ResourceRefPathsInterface = []*gnmi.Path{
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "ethernet"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "ethernet"},
			{Name: "flow-control"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "lag"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "lag"},
			{Name: "lacp"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
			{Name: "multicast-queue", Key: map[string]string{"queue-id": ""}},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
			{Name: "multicast-queue", Key: map[string]string{"queue-id": ""}},
			{Name: "scheduling"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
			{Name: "scheduler"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
			{Name: "scheduler"},
			{Name: "tier", Key: map[string]string{"level": ""}},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
			{Name: "scheduler"},
			{Name: "tier", Key: map[string]string{"level": ""}},
			{Name: "node", Key: map[string]string{"node-number": ""}},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
			{Name: "unicast-queue", Key: map[string]string{"queue-id": ""}},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "qos"},
			{Name: "output"},
			{Name: "unicast-queue", Key: map[string]string{"queue-id": ""}},
			{Name: "scheduling"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "sflow"},
		},
	},
	{
		Elem: []*gnmi.PathElem{
			{Name: "interface", Key: map[string]string{"name": ""}},
			{Name: "transceiver"},
		},
	},
}
var DependencyInterface = []*parser.LeafRefGnmi{}
var LocalleafRefInterface = []*parser.LeafRefGnmi{}
var ExternalleafRefInterface = []*parser.LeafRefGnmi{
	{
		LocalPath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "interface"},
				{Name: "ethernet"},
				{Name: "aggregate-id"},
			},
		},
		RemotePath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "interface", Key: map[string]string{"name": ""}},
			},
		},
	},
	{
		LocalPath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "interface"},
				{Name: "qos"},
				{Name: "output"},
				{Name: "multicast-queue", Key: map[string]string{"queue-id": ""}},
				{Name: "template"},
			},
		},
		RemotePath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "qos"},
				{Name: "queue-templates"},
				{Name: "queue-template", Key: map[string]string{"name": ""}},
			},
		},
	},
	{
		LocalPath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "interface"},
				{Name: "qos"},
				{Name: "output"},
				{Name: "unicast-queue", Key: map[string]string{"queue-id": ""}},
				{Name: "template"},
			},
		},
		RemotePath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "qos"},
				{Name: "queue-templates"},
				{Name: "queue-template", Key: map[string]string{"name": ""}},
			},
		},
	},
	{
		LocalPath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "interface"},
				{Name: "qos"},
				{Name: "output"},
				{Name: "unicast-queue", Key: map[string]string{"queue-id": ""}},
				{Name: "voq-template"},
			},
		},
		RemotePath: &gnmi.Path{
			Elem: []*gnmi.PathElem{
				{Name: "qos"},
				{Name: "queue-templates"},
				{Name: "queue-template", Key: map[string]string{"name": ""}},
			},
		},
	},
}

// SetupInterface adds a controller that reconciles Interfaces.
func SetupInterface(mgr ctrl.Manager, o controller.Options, l logging.Logger, poll time.Duration, namespace string) (string, chan cevent.GenericEvent, error) {

	name := managed.ControllerName(srlv1.InterfaceGroupKind)

	events := make(chan cevent.GenericEvent)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(srlv1.InterfaceGroupVersionKind),
		managed.WithExternalConnecter(&connectorInterface{
			log:         l,
			kube:        mgr.GetClient(),
			usage:       resource.NewNetworkNodeUsageTracker(mgr.GetClient(), &ndrv1.NetworkNodeUsage{}),
			newClientFn: target.NewTarget},
		),
		managed.WithParser(l),
		managed.WithValidator(&validatorInterface{log: l, parser: *parser.NewParser(parser.WithLogger(l))}),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return srlv1.InterfaceGroupKind, events, ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o).
		For(&srlv1.SrlInterface{}).
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

type validatorInterface struct {
	log    logging.Logger
	parser parser.Parser
}

func (v *validatorInterface) ValidateLocalleafRef(ctx context.Context, mg resource.Managed) (managed.ValidateLocalleafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateLocalleafRef...")

	// json unmarshal the resource
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return managed.ValidateLocalleafRefObservation{}, errors.New(errUnexpectedInterface)
	}
	d, err := json.Marshal(&o.Spec.ForNetworkNode)
	if err != nil {
		return managed.ValidateLocalleafRefObservation{}, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	// For local leafref validation we dont need to supply the external data so we use nil
	success, resultleafRefValidation, err := v.parser.ValidateLeafRefGnmi(
		parser.LeafRefValidationLocal, x1, nil, LocalleafRefInterface, log)
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

func (v *validatorInterface) ValidateExternalleafRef(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateExternalleafRefObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateExternalleafRef...")

	// json unmarshal the resource
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return managed.ValidateExternalleafRefObservation{}, errors.New(errUnexpectedInterface)
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
	success, resultleafRefValidation, err := v.parser.ValidateLeafRefGnmi(
		parser.LeafRefValidationExternal, x1, x2, ExternalleafRefInterface, log)
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

func (v *validatorInterface) ValidateParentDependency(ctx context.Context, mg resource.Managed, cfg []byte) (managed.ValidateParentDependencyObservation, error) {
	log := v.log.WithValues("resource", mg.GetName())
	log.Debug("ValidateParentDependency...")

	// we initialize a global list for finer information on the resolution
	resultleafRefValidation := make([]*parser.ResolvedLeafRefGnmi, 0)
	log.Debug("ValidateParentDependency success", "resultParentValidation", resultleafRefValidation)
	return managed.ValidateParentDependencyObservation{
		Success:          true,
		ResolvedLeafRefs: resultleafRefValidation}, nil
}

// ValidateResourceIndexes validates if the indexes of a resource got changed
// if so we need to delete the original resource, because it will be dangling if we dont delete it
func (v *validatorInterface) ValidateResourceIndexes(ctx context.Context, mg resource.Managed) (managed.ValidateResourceIndexesObservation, error) {
	log := v.log.WithValues("resosurce", mg.GetName())

	// json unmarshal the resource
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return managed.ValidateResourceIndexesObservation{}, errors.New(errUnexpectedInterface)
	}
	log.Debug("ValidateResourceIndexes", "Spec", o.Spec)

	rootPath := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{Name: "interface", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlInterface.Name}},
			},
		},
	}

	origResourceIndex := mg.GetResourceIndexes()
	// we call the CompareConfigPathsWithResourceKeys irrespective is the get resource index returns nil
	changed, deletPaths, newResourceIndex := v.parser.CompareGnmiPathsWithResourceKeys(rootPath[0], origResourceIndex)
	if changed {
		log.Debug("ValidateResourceIndexes changed", "deletPaths", deletPaths[0])
		return managed.ValidateResourceIndexesObservation{Changed: true, ResourceDeletes: deletPaths, ResourceIndexes: newResourceIndex}, nil
	}

	log.Debug("ValidateResourceIndexes success")
	return managed.ValidateResourceIndexesObservation{Changed: false, ResourceIndexes: newResourceIndex}, nil
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connectorInterface struct {
	log         logging.Logger
	kube        client.Client
	usage       resource.Tracker
	newClientFn func(c *gnmitypes.TargetConfig) *target.Target
	//newClientFn func(ctx context.Context, cfg ndd.Config) (config.ConfigurationClient, error)
}

// Connect produces an ExternalClient by:
// 1. Tracking that the managed resource is using a NetworkNode.
// 2. Getting the managed resource's NetworkNode with connection details
// A resource is mapped to a single target
func (c *connectorInterface) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	log := c.log.WithValues("resource", mg.GetName())
	log.Debug("Connect")
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return nil, errors.New(errUnexpectedInterface)
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
	cfg := &gnmitypes.TargetConfig{
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
	}

	cl := target.NewTarget(cfg)
	if err := cl.CreateGNMIClient(ctx); err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	// we make a string here since we use a trick in registration to go to multiple targets
	// while here the object is mapped to a single target/network node
	tns := make([]string, 0)
	tns = append(tns, nn.GetName())

	return &externalInterface{client: cl, targets: tns, log: log, parser: *parser.NewParser(parser.WithLogger(log))}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type externalInterface struct {
	//client  config.ConfigurationClient
	client  *target.Target
	targets []string
	log     logging.Logger
	parser  parser.Parser
}

func (e *externalInterface) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedInterface)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Observing ...")

	// rootpath of the resource
	rootPath := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{Name: "interface", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlInterface.Name}},
			},
		},
	}

	// gvk: group, version, kind, name, namespace of the resource
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

	// gext: gni extension information for the resource: action, gvk name and level
	gextInfo := &gext.GEXT{
		Action: gext.GEXTActionGet,
		Name:   gvkstring,
		Level:  levelInterface,
	}
	gextInfoString, err := gextInfo.String()
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errGetGextInfo)
	}

	// gnmi get request
	req := &gnmi.GetRequest{
		Path:     rootPath,
		Encoding: gnmi.Encoding_JSON,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gextInfoString)}}},
		},
	}

	// gnmi get response
	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errReadInterface)
	}

	// validate if the extension matches or not
	if resp.GetExtension()[0].GetRegisteredExt().GetId() != gnmi_ext.ExtensionID_EID_EXPERIMENTAL {
		log.Debug("Observe response GNMI Extension mismatch", "Extension Info", resp.GetExtension()[0])
		return managed.ExternalObservation{}, errors.New(errGnmiExtensionMismatch)
	}

	// get gnmi extension metadata
	meta := resp.GetExtension()[0].GetRegisteredExt().GetMsg()
	respMeta := &gext.GEXT{}
	if err := json.Unmarshal(meta, &respMeta); err != nil {
		log.Debug("Observe response gext unmarshal issue", "Extension Info", meta)
		return managed.ExternalObservation{}, errors.Wrap(err, errJSONMarshal)
	}

	// prepare the input data to compare against the response data
	d, err := json.Marshal(&o.Spec.ForNetworkNode)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errJSONMarshal)
	}
	var x1 interface{}
	if err := json.Unmarshal(d, &x1); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errJSONUnMarshal)
	}

	// remove the hierarchical elements for data processing, comparison, etc
	// they are used in the provider for parent dependency resolution
	// but are not relevant in the data, they are referenced in the rootPath
	// when interacting with the device driver
	hids := make([]string, 0)
	x1 = e.parser.RemoveLeafsFromJSONData(x1, hids)

	// validate gnmi resp information
	var x2 interface{}
	if len(resp.GetNotification()) != 0 {
		if len(resp.GetNotification()[0].GetUpdate()) != 0 {
			// get value from gnmi get response
			x2, err = e.parser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
			if err != nil {
				log.Debug("Observe response get value issue")
				return managed.ExternalObservation{}, errors.Wrap(err, errJSONMarshal)
			}
		}
	}

	// logging information that will be used to provide the response
	log.Debug("Observer Response", "Meta", string(meta))
	log.Debug("Spec Data", "X1", x1)
	log.Debug("Resp Data", "X2", x2)

	// if the cache is not ready we back off and return
	if !respMeta.CacheReady {
		log.Debug("Cache Not Ready ...")
		return managed.ExternalObservation{
			Ready:            false,
			ResourceExists:   false,
			ResourceHasData:  true,
			ResourceUpToDate: false,
		}, nil
	}

	if !respMeta.Exists {
		// Resource Does not Exists
		if respMeta.HasData {
			// this is an umnaged resource which has data and will be moved to a managed resource
			// for lists with keys we need to create a list before calulating the paths since this is what
			// the object eventually happens to be based upon. We avoid having multiple entries in a list object
			// and hence we have to add this step
			x1, err = e.parser.AddJSONDataToList(x1)
			if err != nil {
				return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
			}

			updatesx1 := e.parser.GetUpdatesFromJSONDataGnmi(rootPath[0], e.parser.XpathToGnmiPath("/", 0), x1, ResourceRefPathsInterface)
			for _, update := range updatesx1 {
				log.Debug("Observe Fine Grane Updates X1", "Path", e.parser.GnmiPathToXPath(update.Path, true), "Value", update.GetVal())
			}
			// for lists with keys we need to create a list before calulating the paths since this is what
			// the object eventually happens to be based upon. We avoid having multiple entries in a list object
			// and hence we have to add this step
			x2, err = e.parser.AddJSONDataToList(x2)
			if err != nil {
				return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
			}
			updatesx2 := e.parser.GetUpdatesFromJSONDataGnmi(rootPath[0], e.parser.XpathToGnmiPath("/", 0), x2, ResourceRefPathsInterface)
			for _, update := range updatesx2 {
				log.Debug("Observe Fine Grane Updates X2", "Path", e.parser.GnmiPathToXPath(update.Path, true), "Value", update.GetVal())
			}

			deletes, updates, err := e.parser.FindResourceDeltaGnmi(updatesx1, updatesx2, log)
			if err != nil {
				return managed.ExternalObservation{}, err
			}
			if len(deletes) != 0 || len(updates) != 0 {
				// UMR -> MR with data, which is NOT up to date
				log.Debug("Observing Respone: resource NOT up to date", "Exists", false, "HasData", true, "UpToDate", false, "Response", resp, "Updates", updates, "Deletes", deletes)
				for _, del := range deletes {
					log.Debug("Observing Respone: resource NOT up to date, deletes", "path", e.parser.GnmiPathToXPath(del, true))
				}
				for _, upd := range updates {
					val, _ := e.parser.GetValue(upd.GetVal())
					log.Debug("Observing Respone: resource NOT up to date, updates", "path", e.parser.GnmiPathToXPath(upd.GetPath(), true), "data", val)
				}
				return managed.ExternalObservation{
					Ready:            true,
					ResourceExists:   false,
					ResourceHasData:  true,
					ResourceUpToDate: false,
					ResourceDeletes:  deletes,
					ResourceUpdates:  updates,
				}, nil
			}
			// UMR -> MR with data, which is up to date
			log.Debug("Observing Respone: resource up to date", "Exists", false, "HasData", true, "UpToDate", true, "Response", resp)
			return managed.ExternalObservation{
				Ready:            true,
				ResourceExists:   false,
				ResourceHasData:  true,
				ResourceUpToDate: true,
			}, nil
		} else {
			// UMR -> MR without data
			log.Debug("Observing Respone:", "Exists", false, "HasData", false, "UpToDate", false, "Response", resp)
			return managed.ExternalObservation{
				Ready:            true,
				ResourceExists:   false,
				ResourceHasData:  false,
				ResourceUpToDate: false,
			}, nil
		}
	} else {
		// Resource Exists
		switch respMeta.Status {
		case gext.ResourceStatusSuccess:
			if respMeta.HasData {
				// data is present
				// for lists with keys we need to create a list before calulating the paths since this is what
				// the object eventually happens to be based upon. We avoid having multiple entries in a list object
				// and hence we have to add this step
				x1, err = e.parser.AddJSONDataToList(x1)
				if err != nil {
					return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
				}

				updatesx1 := e.parser.GetUpdatesFromJSONDataGnmi(rootPath[0], e.parser.XpathToGnmiPath("/", 0), x1, ResourceRefPathsInterface)
				for _, update := range updatesx1 {
					log.Debug("Observe Fine Grane Updates X1", "Path", e.parser.GnmiPathToXPath(update.Path, true), "Value", update.GetVal())
				}
				// for lists with keys we need to create a list before calulating the paths since this is what
				// the object eventually happens to be based upon. We avoid having multiple entries in a list object
				// and hence we have to add this step
				x2, err = e.parser.AddJSONDataToList(x2)
				if err != nil {
					return managed.ExternalObservation{}, errors.Wrap(err, errWrongInputdata)
				}
				updatesx2 := e.parser.GetUpdatesFromJSONDataGnmi(rootPath[0], e.parser.XpathToGnmiPath("/", 0), x2, ResourceRefPathsInterface)
				for _, update := range updatesx2 {
					log.Debug("Observe Fine Grane Updates X2", "Path", e.parser.GnmiPathToXPath(update.Path, true), "Value", update.GetVal())
				}

				deletes, updates, err := e.parser.FindResourceDeltaGnmi(updatesx1, updatesx2, log)
				if err != nil {
					return managed.ExternalObservation{}, err
				}
				// MR -> MR, resource is NOT up to date
				if len(deletes) != 0 || len(updates) != 0 {
					// resource is NOT up to date
					log.Debug("Observing Respone: resource NOT up to date", "Exists", true, "HasData", true, "UpToDate", false, "Response", resp, "Updates", updates, "Deletes", deletes)
					for _, del := range deletes {
						log.Debug("Observing Respone: resource NOT up to date, deletes", "path", e.parser.GnmiPathToXPath(del, true))
					}
					for _, upd := range updates {
						val, _ := e.parser.GetValue(upd.GetVal())
						log.Debug("Observing Respone: resource NOT up to date, updates", "path", e.parser.GnmiPathToXPath(upd.GetPath(), true), "data", val)
					}
					return managed.ExternalObservation{
						Ready:            true,
						ResourceExists:   true,
						ResourceHasData:  true,
						ResourceUpToDate: false,
						ResourceDeletes:  deletes,
						ResourceUpdates:  updates,
					}, nil
				}
				// MR -> MR, resource is up to date
				log.Debug("Observing Respone: resource up to date", "Exists", true, "HasData", true, "UpToDate", true, "Response", resp)
				return managed.ExternalObservation{
					Ready:            true,
					ResourceExists:   true,
					ResourceHasData:  true,
					ResourceUpToDate: true,
				}, nil
			} else {
				// MR -> MR, resource has no data, strange, someone could have deleted the resource
				log.Debug("Observing Respone", "Exists", true, "HasData", false, "UpToDate", false, "Status", respMeta.Status)
				return managed.ExternalObservation{
					Ready:            true,
					ResourceExists:   true,
					ResourceHasData:  false,
					ResourceUpToDate: false,
				}, nil
			}
		default:
			// MR -> MR, resource is not in a success state, so the object might still be in creation phase
			log.Debug("Observing Respone", "Exists", true, "HasData", false, "UpToDate", false, "Status", respMeta.Status)
			return managed.ExternalObservation{
				Ready:            true,
				ResourceExists:   true,
				ResourceHasData:  false,
				ResourceUpToDate: false,
			}, nil
		}
	}
}

func (e *externalInterface) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedInterface)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Creating ...")

	rootPath := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{Name: "interface", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlInterface.Name}},
			},
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

	updates := e.parser.GetUpdatesFromJSONDataGnmi(rootPath[0], e.parser.XpathToGnmiPath("/", 0), x1, ResourceRefPathsInterface)
	for _, update := range updates {
		log.Debug("Create Fine Grane Updates", "Path", update.Path, "Value", update.GetVal())
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

	gextInfo := &gext.GEXT{
		Action:   gext.GEXTActionCreate,
		Name:     gvkstring,
		Level:    levelInterface,
		RootPath: rootPath[0],
	}
	gextInfoString, err := gextInfo.String()
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errGetGextInfo)
	}

	if len(updates) == 0 {
		log.Debug("cannot create object since there are no updates present")
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateObject)
	}

	req := &gnmi.SetRequest{
		Replace: updates,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gextInfoString)}}},
		},
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errReadInterface)
	}

	return managed.ExternalCreation{}, nil
}

func (e *externalInterface) Update(ctx context.Context, mg resource.Managed, obs managed.ExternalObservation) (managed.ExternalUpdate, error) {
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedInterface)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Updating ...")

	for _, u := range obs.ResourceUpdates {
		log.Debug("Update -> Update", "Path", u.Path, "Value", u.GetVal())
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

	gextInfo := &gext.GEXT{
		Action: gext.GEXTActionUpdate,
		Name:   gvkstring,
		Level:  levelInterface,
	}
	gextInfoString, err := gextInfo.String()
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errGetGextInfo)
	}

	req := &gnmi.SetRequest{
		Update: obs.ResourceUpdates,
		Delete: obs.ResourceDeletes,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gextInfoString)}}},
		},
	}

	_, err = e.client.Set(ctx, req)
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errReadInterface)
	}

	return managed.ExternalUpdate{}, nil
}

func (e *externalInterface) Delete(ctx context.Context, mg resource.Managed) error {
	o, ok := mg.(*srlv1.SrlInterface)
	if !ok {
		return errors.New(errUnexpectedInterface)
	}
	log := e.log.WithValues("Resource", o.GetName())
	log.Debug("Deleting ...")

	rootPath := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{Name: "interface", Key: map[string]string{"name": *o.Spec.ForNetworkNode.SrlInterface.Name}},
			},
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

	gextInfo := &gext.GEXT{
		Action: gext.GEXTActionDelete,
		Name:   gvkstring,
		Level:  levelInterface,
	}
	gextInfoString, err := gextInfo.String()
	if err != nil {
		return errors.Wrap(err, errGetGextInfo)
	}

	req := gnmi.SetRequest{
		Delete: rootPath,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gextInfoString)}}},
		},
	}

	_, err = e.client.Set(ctx, &req)
	if err != nil {
		return errors.Wrap(err, errDeleteInterface)
	}

	return nil
}

func (e *externalInterface) GetTarget() []string {
	return e.targets
}

func (e *externalInterface) GetConfig(ctx context.Context) ([]byte, error) {
	e.log.Debug("Get Config ...")
	req := &gnmi.GetRequest{
		Path:     []*gnmi.Path{},
		Encoding: gnmi.Encoding_JSON,
	}

	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return make([]byte, 0), errors.Wrap(err, errGetConfig)
	}

	if len(resp.GetNotification()) != 0 {
		if len(resp.GetNotification()[0].GetUpdate()) != 0 {
			x2, err := e.parser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
			if err != nil {
				return make([]byte, 0), errors.Wrap(err, errGetConfig)
			}

			data, err := json.Marshal(x2)
			if err != nil {
				return make([]byte, 0), errors.Wrap(err, errJSONMarshal)
			}
			return data, nil
		}
	}
	e.log.Debug("Get Config Empty response")
	return nil, nil
}

func (e *externalInterface) GetResourceName(ctx context.Context, path []*gnmi.Path) (string, error) {
	e.log.Debug("Get ResourceName ...")

	gextInfo := &gext.GEXT{
		Action: gext.GEXTActionGetResourceName,
	}
	gextInfoString, err := gextInfo.String()
	if err != nil {
		return "", errors.Wrap(err, errGetGextInfo)
	}

	req := &gnmi.GetRequest{
		Path:     path,
		Encoding: gnmi.Encoding_JSON,
		Extension: []*gnmi_ext.Extension{
			{Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{Id: gnmi_ext.ExtensionID_EID_EXPERIMENTAL, Msg: []byte(gextInfoString)}}},
		},
	}

	resp, err := e.client.Get(ctx, req)
	if err != nil {
		return "", errors.Wrap(err, errGetResourceName)
	}

	x2, err := e.parser.GetValue(resp.GetNotification()[0].GetUpdate()[0].Val)
	if err != nil {
		return "", errors.Wrap(err, errJSONMarshal)
	}

	d, err := json.Marshal(x2)
	if err != nil {
		return "", errors.Wrap(err, errJSONMarshal)
	}

	var resourceName nddv1.ResourceName
	if err := json.Unmarshal(d, &resourceName); err != nil {
		return "", errors.Wrap(err, errJSONUnMarshal)
	}

	e.log.Debug("Get ResourceName Response", "ResourceName", resourceName)

	return resourceName.Name, nil
}
