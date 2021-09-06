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
	// NetworkinstanceProtocolsLinuxFinalizer is the name of the finalizer added to
	// NetworkinstanceProtocolsLinux to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceProtocolsLinuxFinalizer string = "linux.srl.ndd.yndd.io"
)

// NetworkinstanceProtocolsLinux struct
type NetworkinstanceProtocolsLinux struct {
	// +kubebuilder:default:=true
	ExportNeighbors *bool `json:"export-neighbors,omitempty"`
	// +kubebuilder:default:=false
	ExportRoutes *bool `json:"export-routes,omitempty"`
	// +kubebuilder:default:=false
	ImportRoutes *bool `json:"import-routes,omitempty"`
}

// NetworkinstanceProtocolsLinuxParameters struct defines the resource Parameters
type NetworkinstanceProtocolsLinuxParameters struct {
	NetworkInstanceName              *string                        `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsLinux *NetworkinstanceProtocolsLinux `json:"linux,omitempty"`
}

// NetworkinstanceProtocolsLinuxObservation struct defines the resource Observation
type NetworkinstanceProtocolsLinuxObservation struct {
}

// A NetworkinstanceProtocolsLinuxSpec defines the desired state of a NetworkinstanceProtocolsLinux.
type NetworkinstanceProtocolsLinuxSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceProtocolsLinuxParameters `json:"forNetworkNode"`
}

// A NetworkinstanceProtocolsLinuxStatus represents the observed state of a NetworkinstanceProtocolsLinux.
type NetworkinstanceProtocolsLinuxStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceProtocolsLinuxObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsLinux is the Schema for the NetworkinstanceProtocolsLinux API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceProtocolsLinux struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsLinuxSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsLinuxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsLinuxList contains a list of NetworkinstanceProtocolsLinuxs
type SrlNetworkinstanceProtocolsLinuxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsLinux `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsLinux{}, &SrlNetworkinstanceProtocolsLinuxList{})
}

// NetworkinstanceProtocolsLinux type metadata.
var (
	NetworkinstanceProtocolsLinuxKind             = reflect.TypeOf(SrlNetworkinstanceProtocolsLinux{}).Name()
	NetworkinstanceProtocolsLinuxGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsLinuxKind}.String()
	NetworkinstanceProtocolsLinuxKindAPIVersion   = NetworkinstanceProtocolsLinuxKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsLinuxGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsLinuxKind)
)
