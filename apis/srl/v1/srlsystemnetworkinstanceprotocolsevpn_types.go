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
	// SystemNetworkinstanceProtocolsEvpnFinalizer is the name of the finalizer added to
	// SystemNetworkinstanceProtocolsEvpn to block delete operations until the physical node can be
	// deprovisioned.
	SystemNetworkinstanceProtocolsEvpnFinalizer string = "evpn.srl.ndd.yndd.io"
)

// SystemNetworkinstanceProtocolsEvpn struct
type SystemNetworkinstanceProtocolsEvpn struct {
	EthernetSegments *SystemNetworkinstanceProtocolsEvpnEthernetSegments `json:"ethernet-segments,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEthernetSegments struct
type SystemNetworkinstanceProtocolsEvpnEthernetSegments struct {
	Timers *SystemNetworkinstanceProtocolsEvpnEthernetSegmentsTimers `json:"timers,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEthernetSegmentsTimers struct
type SystemNetworkinstanceProtocolsEvpnEthernetSegmentsTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	ActivationTimer *uint32 `json:"activation-timer,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=6000
	// +kubebuilder:default:=10
	BootTimer *uint32 `json:"boot-timer,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnParameters struct defines the resource Parameters
type SystemNetworkinstanceProtocolsEvpnParameters struct {
	SrlSystemNetworkinstanceProtocolsEvpn *SystemNetworkinstanceProtocolsEvpn `json:"evpn,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnObservation struct defines the resource Observation
type SystemNetworkinstanceProtocolsEvpnObservation struct {
}

// A SystemNetworkinstanceProtocolsEvpnSpec defines the desired state of a SystemNetworkinstanceProtocolsEvpn.
type SystemNetworkinstanceProtocolsEvpnSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     SystemNetworkinstanceProtocolsEvpnParameters `json:"forNetworkNode"`
}

// A SystemNetworkinstanceProtocolsEvpnStatus represents the observed state of a SystemNetworkinstanceProtocolsEvpn.
type SystemNetworkinstanceProtocolsEvpnStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        SystemNetworkinstanceProtocolsEvpnObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpn is the Schema for the SystemNetworkinstanceProtocolsEvpn API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlSystemNetworkinstanceProtocolsEvpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsEvpnSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsEvpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnList contains a list of SystemNetworkinstanceProtocolsEvpns
type SrlSystemNetworkinstanceProtocolsEvpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpn{}, &SrlSystemNetworkinstanceProtocolsEvpnList{})
}

// SystemNetworkinstanceProtocolsEvpn type metadata.
var (
	SystemNetworkinstanceProtocolsEvpnKind             = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsEvpn{}).Name()
	SystemNetworkinstanceProtocolsEvpnGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsEvpnKind}.String()
	SystemNetworkinstanceProtocolsEvpnKindAPIVersion   = SystemNetworkinstanceProtocolsEvpnKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsEvpnGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsEvpnKind)
)
