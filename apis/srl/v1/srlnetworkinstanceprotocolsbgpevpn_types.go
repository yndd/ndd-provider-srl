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
	// NetworkinstanceProtocolsBgpevpnFinalizer is the name of the finalizer added to
	// NetworkinstanceProtocolsBgpevpn to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceProtocolsBgpevpnFinalizer string = "bgpEvpn.srl.ndd.yndd.io"
)

// NetworkinstanceProtocolsBgpevpn struct
type NetworkinstanceProtocolsBgpevpn struct {
	BgpInstance []*NetworkinstanceProtocolsBgpevpnBgpInstance `json:"bgp-instance,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstance struct
type NetworkinstanceProtocolsBgpevpnBgpInstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	DefaultAdminTag *uint32 `json:"default-admin-tag,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8
	// +kubebuilder:default:=1
	Ecmp *uint8 `json:"ecmp,omitempty"`
	// +kubebuilder:validation:Enum=`vxlan`
	// +kubebuilder:default:=vxlan
	EncapsulationType *string `json:"encapsulation-type,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Evi            *uint32                                           `json:"evi"`
	Id             *string                                           `json:"id,omitempty"`
	Routes         *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutes `json:"routes,omitempty"`
	VxlanInterface *string                                           `json:"vxlan-interface,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutes struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutes struct {
	BridgeTable *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable `json:"bridge-table,omitempty"`
	RouteTable  *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable  `json:"route-table,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable struct {
	InclusiveMcast *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast `json:"inclusive-mcast,omitempty"`
	MacIp          *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp          `json:"mac-ip,omitempty"`
	// +kubebuilder:default:=use-system-ipv4-address
	NextHop *string `json:"next-hop,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	OriginatingIp *string `json:"originating-ip,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable struct {
	MacIp *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp `json:"mac-ip,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp struct {
	// +kubebuilder:default:=false
	AdvertiseGatewayMac *bool `json:"advertise-gateway-mac,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnParameters struct defines the resource Parameters
type NetworkinstanceProtocolsBgpevpnParameters struct {
	NetworkInstanceName                *string                          `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsBgpevpn *NetworkinstanceProtocolsBgpevpn `json:"bgp-evpn,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnObservation struct defines the resource Observation
type NetworkinstanceProtocolsBgpevpnObservation struct {
}

// A NetworkinstanceProtocolsBgpevpnSpec defines the desired state of a NetworkinstanceProtocolsBgpevpn.
type NetworkinstanceProtocolsBgpevpnSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceProtocolsBgpevpnParameters `json:"forNetworkNode"`
}

// A NetworkinstanceProtocolsBgpevpnStatus represents the observed state of a NetworkinstanceProtocolsBgpevpn.
type NetworkinstanceProtocolsBgpevpnStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceProtocolsBgpevpnObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpevpn is the Schema for the NetworkinstanceProtocolsBgpevpn API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceProtocolsBgpevpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsBgpevpnSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsBgpevpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpevpnList contains a list of NetworkinstanceProtocolsBgpevpns
type SrlNetworkinstanceProtocolsBgpevpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsBgpevpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsBgpevpn{}, &SrlNetworkinstanceProtocolsBgpevpnList{})
}

// NetworkinstanceProtocolsBgpevpn type metadata.
var (
	NetworkinstanceProtocolsBgpevpnKind             = reflect.TypeOf(SrlNetworkinstanceProtocolsBgpevpn{}).Name()
	NetworkinstanceProtocolsBgpevpnGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsBgpevpnKind}.String()
	NetworkinstanceProtocolsBgpevpnKindAPIVersion   = NetworkinstanceProtocolsBgpevpnKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsBgpevpnGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsBgpevpnKind)
)
