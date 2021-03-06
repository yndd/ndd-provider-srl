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

package v1

import (
	"reflect"

	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// NetworkinstanceAggregateroutesFinalizer is the name of the finalizer added to
	// NetworkinstanceAggregateroutes to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceAggregateroutesFinalizer string = "aggregateRoutes.srl.ndd.yndd.io"
)

// NetworkinstanceAggregateroutes struct
type NetworkinstanceAggregateroutes struct {
	Route []*NetworkinstanceAggregateroutesRoute `json:"route,omitempty"`
}

// NetworkinstanceAggregateroutesRoute struct
type NetworkinstanceAggregateroutesRoute struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState   *string                                         `json:"admin-state,omitempty"`
	Aggregator   *NetworkinstanceAggregateroutesRouteAggregator  `json:"aggregator,omitempty"`
	Communities  *NetworkinstanceAggregateroutesRouteCommunities `json:"communities,omitempty"`
	GenerateIcmp *bool                                           `json:"generate-icmp,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix      *string `json:"prefix"`
	SummaryOnly *bool   `json:"summary-only,omitempty"`
}

// NetworkinstanceAggregateroutesRouteAggregator struct
type NetworkinstanceAggregateroutesRouteAggregator struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	AsNumber *uint32 `json:"as-number,omitempty"`
}

// NetworkinstanceAggregateroutesRouteCommunities struct
type NetworkinstanceAggregateroutesRouteCommunities struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
	Add *string `json:"add,omitempty"`
}

// NetworkinstanceAggregateroutesParameters are the parameter fields of a NetworkinstanceAggregateroutes.
type NetworkinstanceAggregateroutesParameters struct {
	NetworkInstanceName               *string                         `json:"network-instance-name"`
	SrlNetworkinstanceAggregateroutes *NetworkinstanceAggregateroutes `json:"aggregate-routes,omitempty"`
}

// NetworkinstanceAggregateroutesObservation are the observable fields of a NetworkinstanceAggregateroutes.
type NetworkinstanceAggregateroutesObservation struct {
}

// A NetworkinstanceAggregateroutesSpec defines the desired state of a NetworkinstanceAggregateroutes.
type NetworkinstanceAggregateroutesSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceAggregateroutesParameters `json:"forNetworkNode"`
}

// A NetworkinstanceAggregateroutesStatus represents the observed state of a NetworkinstanceAggregateroutes.
type NetworkinstanceAggregateroutesStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceAggregateroutesObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceAggregateroutes is the Schema for the NetworkinstanceAggregateroutes API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceAggregateroutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceAggregateroutesSpec   `json:"spec,omitempty"`
	Status NetworkinstanceAggregateroutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceAggregateroutesList contains a list of NetworkinstanceAggregateroutess
type SrlNetworkinstanceAggregateroutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceAggregateroutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceAggregateroutes{}, &SrlNetworkinstanceAggregateroutesList{})
}

// NetworkinstanceAggregateroutes type metadata.
var (
	NetworkinstanceAggregateroutesKind             = reflect.TypeOf(SrlNetworkinstanceAggregateroutes{}).Name()
	NetworkinstanceAggregateroutesGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceAggregateroutesKind}.String()
	NetworkinstanceAggregateroutesKindAPIVersion   = NetworkinstanceAggregateroutesKind + "." + GroupVersion.String()
	NetworkinstanceAggregateroutesGroupVersionKind = GroupVersion.WithKind(NetworkinstanceAggregateroutesKind)
)
