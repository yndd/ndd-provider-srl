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
	// TunnelinterfaceVxlaninterfaceFinalizer is the name of the finalizer added to
	// TunnelinterfaceVxlaninterface to block delete operations until the physical node can be
	// deprovisioned.
	TunnelinterfaceVxlaninterfaceFinalizer string = "vxlanInterface.srl.ndd.yndd.io"
)

// TunnelinterfaceVxlaninterface struct
type TunnelinterfaceVxlaninterface struct {
	BridgeTable *TunnelinterfaceVxlaninterfaceBridgeTable `json:"bridge-table,omitempty"`
	Egress      *TunnelinterfaceVxlaninterfaceEgress      `json:"egress,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=99999999
	Index   *uint32                               `json:"index"`
	Ingress *TunnelinterfaceVxlaninterfaceIngress `json:"ingress,omitempty"`
	Type    *string                               `json:"type"`
}

// TunnelinterfaceVxlaninterfaceBridgeTable struct
type TunnelinterfaceVxlaninterfaceBridgeTable struct {
}

// TunnelinterfaceVxlaninterfaceEgress struct
type TunnelinterfaceVxlaninterfaceEgress struct {
	DestinationGroups   *TunnelinterfaceVxlaninterfaceEgressDestinationGroups   `json:"destination-groups,omitempty"`
	InnerEthernetHeader *TunnelinterfaceVxlaninterfaceEgressInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
	SourceIp            *string                                                 `json:"source-ip,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroups struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroups struct {
	Group []*TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroup `json:"group,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroup struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                                                 `json:"admin-state,omitempty"`
	Destination []*TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestination `json:"destination,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi *string `json:"esi,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestination struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestination struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Index               *uint16                                                                                  `json:"index"`
	InnerEthernetHeader *TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	DestinationMac *string `json:"destination-mac,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressInnerEthernetHeader struct
type TunnelinterfaceVxlaninterfaceEgressInnerEthernetHeader struct {
	SourceMac *string `json:"source-mac,omitempty"`
}

// TunnelinterfaceVxlaninterfaceIngress struct
type TunnelinterfaceVxlaninterfaceIngress struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni"`
}

// TunnelinterfaceVxlaninterfaceParameters are the parameter fields of a TunnelinterfaceVxlaninterface.
type TunnelinterfaceVxlaninterfaceParameters struct {
	TunnelInterfaceName              *string                        `json:"tunnel-interface-name"`
	SrlTunnelinterfaceVxlaninterface *TunnelinterfaceVxlaninterface `json:"vxlan-interface,omitempty"`
}

// TunnelinterfaceVxlaninterfaceObservation are the observable fields of a TunnelinterfaceVxlaninterface.
type TunnelinterfaceVxlaninterfaceObservation struct {
}

// A TunnelinterfaceVxlaninterfaceSpec defines the desired state of a TunnelinterfaceVxlaninterface.
type TunnelinterfaceVxlaninterfaceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     TunnelinterfaceVxlaninterfaceParameters `json:"forNetworkNode"`
}

// A TunnelinterfaceVxlaninterfaceStatus represents the observed state of a TunnelinterfaceVxlaninterface.
type TunnelinterfaceVxlaninterfaceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        TunnelinterfaceVxlaninterfaceObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceVxlaninterface is the Schema for the TunnelinterfaceVxlaninterface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlTunnelinterfaceVxlaninterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TunnelinterfaceVxlaninterfaceSpec   `json:"spec,omitempty"`
	Status TunnelinterfaceVxlaninterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceVxlaninterfaceList contains a list of TunnelinterfaceVxlaninterfaces
type SrlTunnelinterfaceVxlaninterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlTunnelinterfaceVxlaninterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlTunnelinterfaceVxlaninterface{}, &SrlTunnelinterfaceVxlaninterfaceList{})
}

// TunnelinterfaceVxlaninterface type metadata.
var (
	TunnelinterfaceVxlaninterfaceKind             = reflect.TypeOf(SrlTunnelinterfaceVxlaninterface{}).Name()
	TunnelinterfaceVxlaninterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: TunnelinterfaceVxlaninterfaceKind}.String()
	TunnelinterfaceVxlaninterfaceKindAPIVersion   = TunnelinterfaceVxlaninterfaceKind + "." + GroupVersion.String()
	TunnelinterfaceVxlaninterfaceGroupVersionKind = GroupVersion.WithKind(TunnelinterfaceVxlaninterfaceKind)
)
