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
	// TunnelinterfaceFinalizer is the name of the finalizer added to
	// Tunnelinterface to block delete operations until the physical node can be
	// deprovisioned.
	TunnelinterfaceFinalizer string = "tunnelInterface.srl.ndd.yndd.io"
)

// Tunnelinterface struct
type Tunnelinterface struct {
	// kubebuilder:validation:MinLength=6
	// kubebuilder:validation:MaxLength=8
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9]))`
	Name *string `json:"name,omitempty"`
}

// TunnelinterfaceSpec struct
type TunnelinterfaceParameters struct {
	SrlTunnelinterface *Tunnelinterface `json:"tunnel-interface,omitempty"`
}

// TunnelinterfaceStatus struct
type TunnelinterfaceObservation struct {
}

// A TunnelinterfaceSpec defines the desired state of a Tunnelinterface.
type TunnelinterfaceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     TunnelinterfaceParameters `json:"forNetworkNode"`
}

// A TunnelinterfaceStatus represents the observed state of a Tunnelinterface.
type TunnelinterfaceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        TunnelinterfaceObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterface is the Schema for the Tunnelinterface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlTunnelinterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TunnelinterfaceSpec   `json:"spec,omitempty"`
	Status TunnelinterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceList contains a list of Tunnelinterfaces
type SrlTunnelinterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlTunnelinterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlTunnelinterface{}, &SrlTunnelinterfaceList{})
}

// Tunnelinterface type metadata.
var (
	TunnelinterfaceKind             = reflect.TypeOf(SrlTunnelinterface{}).Name()
	TunnelinterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: TunnelinterfaceKind}.String()
	TunnelinterfaceKindAPIVersion   = TunnelinterfaceKind + "." + GroupVersion.String()
	TunnelinterfaceGroupVersionKind = GroupVersion.WithKind(TunnelinterfaceKind)
)
