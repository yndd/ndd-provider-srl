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
	// NetworkinstanceStaticroutesFinalizer is the name of the finalizer added to
	// NetworkinstanceStaticroutes to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceStaticroutesFinalizer string = "staticRoutes.srl.ndd.yndd.io"
)

// NetworkinstanceStaticroutes struct
type NetworkinstanceStaticroutes struct {
	Route []*NetworkinstanceStaticroutesRoute `json:"route,omitempty"`
}

// NetworkinstanceStaticroutesRoute struct
type NetworkinstanceStaticroutesRoute struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=1
	Metric       *uint32 `json:"metric,omitempty"`
	NextHopGroup *string `json:"next-hop-group,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Preference *uint8 `json:"preference,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix,omitempty"`
}

// NetworkinstanceStaticroutesSpec struct
type NetworkinstanceStaticroutesParameters struct {
	NetworkInstanceName            *string                      `json:"network-instance-name"`
	SrlNetworkinstanceStaticroutes *NetworkinstanceStaticroutes `json:"static-routes,omitempty"`
}

// NetworkinstanceStaticroutesStatus struct
type NetworkinstanceStaticroutesObservation struct {
}

// A NetworkinstanceStaticroutesSpec defines the desired state of a NetworkinstanceStaticroutes.
type NetworkinstanceStaticroutesSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceStaticroutesParameters `json:"forNetworkNode"`
}

// A NetworkinstanceStaticroutesStatus represents the observed state of a NetworkinstanceStaticroutes.
type NetworkinstanceStaticroutesStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceStaticroutesObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceStaticroutes is the Schema for the NetworkinstanceStaticroutes API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceStaticroutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceStaticroutesSpec   `json:"spec,omitempty"`
	Status NetworkinstanceStaticroutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceStaticroutesList contains a list of NetworkinstanceStaticroutess
type SrlNetworkinstanceStaticroutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceStaticroutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceStaticroutes{}, &SrlNetworkinstanceStaticroutesList{})
}

// NetworkinstanceStaticroutes type metadata.
var (
	NetworkinstanceStaticroutesKind             = reflect.TypeOf(SrlNetworkinstanceStaticroutes{}).Name()
	NetworkinstanceStaticroutesGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceStaticroutesKind}.String()
	NetworkinstanceStaticroutesKindAPIVersion   = NetworkinstanceStaticroutesKind + "." + GroupVersion.String()
	NetworkinstanceStaticroutesGroupVersionKind = GroupVersion.WithKind(NetworkinstanceStaticroutesKind)
)
