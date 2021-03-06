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
	// SystemNameFinalizer is the name of the finalizer added to
	// SystemName to block delete operations until the physical node can be
	// deprovisioned.
	SystemNameFinalizer string = "name.srl.ndd.yndd.io"
)

// SystemName struct
type SystemName struct {
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	DomainName *string `json:"domain-name,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])`
	HostName *string `json:"host-name,omitempty"`
}

// SystemNameParameters are the parameter fields of a SystemName.
type SystemNameParameters struct {
	SrlSystemName *SystemName `json:"name,omitempty"`
}

// SystemNameObservation are the observable fields of a SystemName.
type SystemNameObservation struct {
}

// A SystemNameSpec defines the desired state of a SystemName.
type SystemNameSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     SystemNameParameters `json:"forNetworkNode"`
}

// A SystemNameStatus represents the observed state of a SystemName.
type SystemNameStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        SystemNameObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemName is the Schema for the SystemName API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlSystemName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNameSpec   `json:"spec,omitempty"`
	Status SystemNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNameList contains a list of SystemNames
type SrlSystemNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemName `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemName{}, &SrlSystemNameList{})
}

// SystemName type metadata.
var (
	SystemNameKind             = reflect.TypeOf(SrlSystemName{}).Name()
	SystemNameGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNameKind}.String()
	SystemNameKindAPIVersion   = SystemNameKind + "." + GroupVersion.String()
	SystemNameGroupVersionKind = GroupVersion.WithKind(SystemNameKind)
)
