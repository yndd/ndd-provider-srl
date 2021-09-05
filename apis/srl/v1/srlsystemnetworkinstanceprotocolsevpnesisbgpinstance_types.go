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
	// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceFinalizer is the name of the finalizer added to
	// SystemNetworkinstanceProtocolsEvpnEsisBgpinstance to block delete operations until the physical node can be
	// deprovisioned.
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceFinalizer string = "bgpInstance.srl.ndd.yndd.io"
)

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct {
	Id *string `json:"id,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceParameters struct {
	SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance *SystemNetworkinstanceProtocolsEvpnEsisBgpinstance `json:"bgp-instance,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceObservation struct {
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec defines the desired state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstance.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceParameters `json:"forNetworkNode"`
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus represents the observed state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstance.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance is the Schema for the SystemNetworkinstanceProtocolsEvpnEsisBgpinstance API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList contains a list of SystemNetworkinstanceProtocolsEvpnEsisBgpinstances
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}, &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList{})
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstance type metadata.
var (
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKind             = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}).Name()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKind}.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKindAPIVersion   = SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceKind)
)
