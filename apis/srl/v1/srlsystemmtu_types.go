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
	// SystemMtuFinalizer is the name of the finalizer added to
	// SystemMtu to block delete operations until the physical node can be
	// deprovisioned.
	SystemMtuFinalizer string = "mtu.srl.ndd.yndd.io"
)

// SystemMtu struct
type SystemMtu struct {
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	// +kubebuilder:default:=1500
	DefaultIpMtu *uint16 `json:"default-ip-mtu,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	DefaultL2Mtu *uint16 `json:"default-l2-mtu,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	DefaultPortMtu *uint16 `json:"default-port-mtu,omitempty"`
	// kubebuilder:validation:Minimum=552
	// kubebuilder:validation:Maximum=9232
	// +kubebuilder:default:=552
	MinPathMtu *uint16 `json:"min-path-mtu,omitempty"`
}

// SystemMtuSpec struct
type SystemMtuParameters struct {
	SrlSystemMtu *SystemMtu `json:"system-mtu"`
}

// SystemMtuStatus struct
type SystemMtuObservation struct {
}

// A SystemMtuSpec defines the desired state of a SystemMtu.
type SystemMtuSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     SystemMtuParameters `json:"forNetworkNode"`
}

// A SystemMtuStatus represents the observed state of a SystemMtu.
type SystemMtuStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        SystemMtuObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemMtu is the Schema for the SystemMtu API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl},shortName=srlint
type SrlSystemMtu struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemMtuSpec   `json:"spec,omitempty"`
	Status SystemMtuStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemMtuList contains a list of SystemMtus
type SrlSystemMtuList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemMtu `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemMtu{}, &SrlSystemMtuList{})
}

// SystemMtu type metadata.
var (
	SystemMtuKind             = reflect.TypeOf(SrlSystemMtu{}).Name()
	SystemMtuGroupKind        = schema.GroupKind{Group: Group, Kind: SystemMtuKind}.String()
	SystemMtuKindAPIVersion   = SystemMtuKind + "." + GroupVersion.String()
	SystemMtuGroupVersionKind = GroupVersion.WithKind(SystemMtuKind)
)
