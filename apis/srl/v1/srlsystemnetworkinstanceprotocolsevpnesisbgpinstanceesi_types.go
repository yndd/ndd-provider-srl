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
	// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiFinalizer is the name of the finalizer added to
	// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi to block delete operations until the physical node can be
	// deprovisioned.
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiFinalizer string = "ethernetSegment.srl.ndd.yndd.io"
)

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string                                                         `json:"admin-state,omitempty"`
	DfElection *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElection `json:"df-election,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi       *string `json:"esi,omitempty"`
	Interface *string `json:"interface,omitempty"`
	// +kubebuilder:validation:Enum=`all-active`;`single-active`
	// +kubebuilder:default:=all-active
	MultiHomingMode *string `json:"multi-homing-mode,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name   *string                                                     `json:"name,omitempty"`
	Routes *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes `json:"routes,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElection struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElection struct {
	Algorithm                        *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithm `json:"algorithm,omitempty"`
	InterfaceStandbySignalingOnNonDf *bool                                                                    `json:"interface-standby-signaling-on-non-df,omitempty"`
	Timers                           *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionTimers    `json:"timers,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithm struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithm struct {
	DefaultAlg    *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlg    `json:"default-alg,omitempty"`
	PreferenceAlg *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlg `json:"preference-alg,omitempty"`
	// +kubebuilder:validation:Enum=`default`;`preference`
	// +kubebuilder:default:=default
	Type *string `json:"type,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlg struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlg struct {
	Capabilities *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlgCapabilities `json:"capabilities,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlgCapabilities struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlgCapabilities struct {
	// +kubebuilder:default:=true
	AcDf *bool `json:"ac-df,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlg struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlg struct {
	Capabilities *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlgCapabilities `json:"capabilities,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=32767
	PreferenceValue *uint32 `json:"preference-value,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlgCapabilities struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlgCapabilities struct {
	// +kubebuilder:default:=true
	AcDf *bool `json:"ac-df,omitempty"`
	// +kubebuilder:default:=false
	NonRevertive *bool `json:"non-revertive,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionTimers struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	ActivationTimer *uint32 `json:"activation-timer,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes struct {
	Esi *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetSegment `json:"esi,omitempty"`
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:=use-system-ipv4-address
	NextHop *string `json:"next-hop,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetSegment struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetSegment struct {
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:=use-system-ipv4-address
	OriginatingIp *string `json:"originating-ip,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiParameters struct {
	BgpInstanceId                                           *string                                               `json:"bgp-instance-id"`
	SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi `json:"ethernet-segment,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiObservation struct {
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec defines the desired state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiParameters `json:"forNetworkNode"`
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus represents the observed state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi is the Schema for the SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList contains a list of SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}, &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList{})
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi type metadata.
var (
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKind             = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}).Name()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKind}.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindAPIVersion   = SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKind)
)
