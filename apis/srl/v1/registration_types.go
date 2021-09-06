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
	// DeviceMatch is the matching string for device registration
	DeviceMatch                    = "srl_nokia"
	// DeviceTypeSRL defines the device type the provider supports
	DeviceTypeSRL nddv1.DeviceType = "nokia-srl"
)

// RegistrationParameters are the parameter fields of a Registration.
type RegistrationParameters struct {
	nddv1.Register `json:",inline"`
}

// RegistrationObservation are the observable fields of a Registration.
type RegistrationObservation struct {
}

// A RegistrationSpec defines the desired state of a Registration.
type RegistrationSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     RegistrationParameters `json:"forNetworkNode"`
}

// A RegistrationStatus represents the observed state of a Registration.
type RegistrationStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        RegistrationObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// Registration is the Schema for the Registration API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl},shortName=srlreg
type Registration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegistrationSpec   `json:"spec"`
	Status RegistrationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RegistrationList contains a list of Registration
type RegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Registration{}, &RegistrationList{})
}

// GetSubscriptions defines a method to get subscriptions
func (o *Registration) GetSubscriptions() []string {
	return o.Spec.ForNetworkNode.Subscriptions
}

// SetSubscriptions defines a method to set subscriptions
func (o *Registration) SetSubscriptions(sub []string) {
	o.Spec.ForNetworkNode.Subscriptions = sub
}

// Registration type metadata.
var (
	RegistrationKind             = reflect.TypeOf(Registration{}).Name()
	RegistrationGroupKind        = schema.GroupKind{Group: Group, Kind: RegistrationKind}.String()
	RegistrationKindAPIVersion   = RegistrationKind + "." + GroupVersion.String()
	RegistrationGroupVersionKind = GroupVersion.WithKind(RegistrationKind)
)
