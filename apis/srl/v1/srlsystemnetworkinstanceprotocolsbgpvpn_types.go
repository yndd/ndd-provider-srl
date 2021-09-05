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
	// SystemNetworkinstanceProtocolsBgpvpnFinalizer is the name of the finalizer added to
	// SystemNetworkinstanceProtocolsBgpvpn to block delete operations until the physical node can be
	// deprovisioned.
	SystemNetworkinstanceProtocolsBgpvpnFinalizer string = "bgpVpn.srl.ndd.yndd.io"
)

// SystemNetworkinstanceProtocolsBgpvpn struct
type SystemNetworkinstanceProtocolsBgpvpn struct {
	BgpInstance []*SystemNetworkinstanceProtocolsBgpvpnBgpInstance `json:"bgp-instance,omitempty"`
}

// SystemNetworkinstanceProtocolsBgpvpnBgpInstance struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstance struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Id                 *uint8                                                             `json:"id,omitempty"`
	RouteDistinguisher *SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher `json:"route-distinguisher,omitempty"`
	RouteTarget        *SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget        `json:"route-target,omitempty"`
}

// SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher struct {
}

// SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget struct {
}

// SystemNetworkinstanceProtocolsBgpvpnSpec struct
type SystemNetworkinstanceProtocolsBgpvpnParameters struct {
	SrlSystemNetworkinstanceProtocolsBgpvpn *SystemNetworkinstanceProtocolsBgpvpn `json:"bgp-vpn,omitempty"`
}

// SystemNetworkinstanceProtocolsBgpvpnStatus struct
type SystemNetworkinstanceProtocolsBgpvpnObservation struct {
}

// A SystemNetworkinstanceProtocolsBgpvpnSpec defines the desired state of a SystemNetworkinstanceProtocolsBgpvpn.
type SystemNetworkinstanceProtocolsBgpvpnSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     SystemNetworkinstanceProtocolsBgpvpnParameters `json:"forNetworkNode"`
}

// A SystemNetworkinstanceProtocolsBgpvpnStatus represents the observed state of a SystemNetworkinstanceProtocolsBgpvpn.
type SystemNetworkinstanceProtocolsBgpvpnStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        SystemNetworkinstanceProtocolsBgpvpnObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsBgpvpn is the Schema for the SystemNetworkinstanceProtocolsBgpvpn API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlSystemNetworkinstanceProtocolsBgpvpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsBgpvpnSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsBgpvpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsBgpvpnList contains a list of SystemNetworkinstanceProtocolsBgpvpns
type SrlSystemNetworkinstanceProtocolsBgpvpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsBgpvpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsBgpvpn{}, &SrlSystemNetworkinstanceProtocolsBgpvpnList{})
}

// SystemNetworkinstanceProtocolsBgpvpn type metadata.
var (
	SystemNetworkinstanceProtocolsBgpvpnKind             = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsBgpvpn{}).Name()
	SystemNetworkinstanceProtocolsBgpvpnGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsBgpvpnKind}.String()
	SystemNetworkinstanceProtocolsBgpvpnKindAPIVersion   = SystemNetworkinstanceProtocolsBgpvpnKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsBgpvpnGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsBgpvpnKind)
)
