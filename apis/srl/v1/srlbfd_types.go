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

	nddv1 "github.com/netw-device-driver/ndd-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// BfdFinalizer is the name of the finalizer added to
	// Bfd to block delete operations until the physical node can be
	// deprovisioned.
	BfdFinalizer string = "bfd.srl.ndd.yndd.io"
)

// Bfd struct
type Bfd struct {
	MicroBfdSessions *BfdMicroBfdSessions `json:"micro-bfd-sessions,omitempty"`
	Subinterface     []*BfdSubinterface   `json:"subinterface,omitempty"`
}

// BfdMicroBfdSessions struct
type BfdMicroBfdSessions struct {
	LagInterface []*BfdMicroBfdSessionsLagInterface `json:"lag-interface,omitempty"`
}

// BfdMicroBfdSessionsLagInterface struct
type BfdMicroBfdSessionsLagInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	DesiredMinimumTransmitInterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	DetectionMultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
	Name         *string `json:"name,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	RemoteAddress *string `json:"remote-address,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	RequiredMinimumReceive *uint32 `json:"required-minimum-receive,omitempty"`
}

// BfdSubinterface struct
type BfdSubinterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	DesiredMinimumTransmitInterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	DetectionMultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Id *string `json:"id,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	// +kubebuilder:default:=0
	MinimumEchoReceiveInterval *uint32 `json:"minimum-echo-receive-interval,omitempty"`
	// kubebuilder:validation:Minimum=10000
	// kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	RequiredMinimumReceive *uint32 `json:"required-minimum-receive,omitempty"`
}

// BfdSpec struct
type BfdParameters struct {
	SrlBfd *Bfd `json:"bfd,omitempty"`
}

// BfdStatus struct
type BfdObservation struct {
}

// A BfdSpec defines the desired state of a Bfd.
type BfdSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     BfdParameters `json:"forNetworkNode"`
}

// A BfdStatus represents the observed state of a Bfd.
type BfdStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        BfdObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlBfd is the Schema for the Bfd API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlBfd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BfdSpec   `json:"spec,omitempty"`
	Status BfdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlBfdList contains a list of Bfds
type SrlBfdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlBfd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlBfd{}, &SrlBfdList{})
}

// Bfd type metadata.
var (
	BfdKind             = reflect.TypeOf(SrlBfd{}).Name()
	BfdGroupKind        = schema.GroupKind{Group: Group, Kind: BfdKind}.String()
	BfdKindAPIVersion   = BfdKind + "." + GroupVersion.String()
	BfdGroupVersionKind = GroupVersion.WithKind(BfdKind)
)
