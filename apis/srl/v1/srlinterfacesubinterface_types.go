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
	// InterfaceSubinterfaceFinalizer is the name of the finalizer added to
	// InterfaceSubinterface to block delete operations until the physical node can be
	// deprovisioned.
	InterfaceSubinterfaceFinalizer string = "subinterface.srl.ndd.yndd.io"
)

// InterfaceSubinterface struct
type InterfaceSubinterface struct {
	Acl *InterfaceSubinterfaceAcl `json:"acl,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState  *string                           `json:"admin-state,omitempty"`
	AnycastGw   *InterfaceSubinterfaceAnycastGw   `json:"anycast-gw,omitempty"`
	BridgeTable *InterfaceSubinterfaceBridgeTable `json:"bridge-table,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9999
	Index *uint32 `json:"index,omitempty"`
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	IpMtu *uint16                    `json:"ip-mtu,omitempty"`
	Ipv4  *InterfaceSubinterfaceIpv4 `json:"ipv4,omitempty"`
	Ipv6  *InterfaceSubinterfaceIpv6 `json:"ipv6,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	L2Mtu                  *uint16                                      `json:"l2-mtu,omitempty"`
	LocalMirrorDestination *InterfaceSubinterfaceLocalMirrorDestination `json:"local-mirror-destination,omitempty"`
	Qos                    *InterfaceSubinterfaceQos                    `json:"qos,omitempty"`
	Type                   *string                                      `json:"type,omitempty"`
	Vlan                   *InterfaceSubinterfaceVlan                   `json:"vlan,omitempty"`
}

// InterfaceSubinterfaceAcl struct
type InterfaceSubinterfaceAcl struct {
	Input  *InterfaceSubinterfaceAclInput  `json:"input,omitempty"`
	Output *InterfaceSubinterfaceAclOutput `json:"output,omitempty"`
}

// InterfaceSubinterfaceAclInput struct
type InterfaceSubinterfaceAclInput struct {
	Ipv4Filter *string `json:"ipv4-filter,omitempty"`
	Ipv6Filter *string `json:"ipv6-filter,omitempty"`
}

// InterfaceSubinterfaceAclOutput struct
type InterfaceSubinterfaceAclOutput struct {
	Ipv4Filter *string `json:"ipv4-filter,omitempty"`
	Ipv6Filter *string `json:"ipv6-filter,omitempty"`
}

// InterfaceSubinterfaceAnycastGw struct
type InterfaceSubinterfaceAnycastGw struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	AnycastGwMac *string `json:"anycast-gw-mac,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	VirtualRouterId *uint8 `json:"virtual-router-id,omitempty"`
}

// InterfaceSubinterfaceBridgeTable struct
type InterfaceSubinterfaceBridgeTable struct {
	// +kubebuilder:default:=false
	DiscardUnknownSrcMac *bool                                           `json:"discard-unknown-src-mac,omitempty"`
	MacDuplication       *InterfaceSubinterfaceBridgeTableMacDuplication `json:"mac-duplication,omitempty"`
	MacLearning          *InterfaceSubinterfaceBridgeTableMacLearning    `json:"mac-learning,omitempty"`
	MacLimit             *InterfaceSubinterfaceBridgeTableMacLimit       `json:"mac-limit,omitempty"`
}

// InterfaceSubinterfaceBridgeTableMacDuplication struct
type InterfaceSubinterfaceBridgeTableMacDuplication struct {
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`;`use-net-instance-action`
	// +kubebuilder:default:=use-net-instance-action
	Action *string `json:"action,omitempty"`
}

// InterfaceSubinterfaceBridgeTableMacLearning struct
type InterfaceSubinterfaceBridgeTableMacLearning struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                           `json:"admin-state,omitempty"`
	Aging      *InterfaceSubinterfaceBridgeTableMacLearningAging `json:"aging,omitempty"`
}

// InterfaceSubinterfaceBridgeTableMacLearningAging struct
type InterfaceSubinterfaceBridgeTableMacLearningAging struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
}

// InterfaceSubinterfaceBridgeTableMacLimit struct
type InterfaceSubinterfaceBridgeTableMacLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	MaximumEntries *int32 `json:"maximum-entries,omitempty"`
	// kubebuilder:validation:Minimum=6
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	WarningThresholdPct *int32 `json:"warning-threshold-pct,omitempty"`
}

// InterfaceSubinterfaceIpv4 struct
type InterfaceSubinterfaceIpv4 struct {
	Address []*InterfaceSubinterfaceIpv4Address `json:"address,omitempty"`
	// +kubebuilder:default:=false
	AllowDirectedBroadcast *bool                                `json:"allow-directed-broadcast,omitempty"`
	Arp                    *InterfaceSubinterfaceIpv4Arp        `json:"arp,omitempty"`
	DhcpClient             *InterfaceSubinterfaceIpv4DhcpClient `json:"dhcp-client,omitempty"`
	DhcpRelay              *InterfaceSubinterfaceIpv4DhcpRelay  `json:"dhcp-relay,omitempty"`
	Vrrp                   *InterfaceSubinterfaceIpv4Vrrp       `json:"vrrp,omitempty"`
}

// InterfaceSubinterfaceIpv4Address struct
type InterfaceSubinterfaceIpv4Address struct {
	AnycastGw *bool `json:"anycast-gw,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	Primary  *string `json:"primary,omitempty"`
}

// InterfaceSubinterfaceIpv4Arp struct
type InterfaceSubinterfaceIpv4Arp struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug *string `json:"debug,omitempty"`
	// +kubebuilder:default:=true
	DuplicateAddressDetection *bool                                  `json:"duplicate-address-detection,omitempty"`
	Evpn                      *InterfaceSubinterfaceIpv4ArpEvpn      `json:"evpn,omitempty"`
	HostRoute                 *InterfaceSubinterfaceIpv4ArpHostRoute `json:"host-route,omitempty"`
	// +kubebuilder:default:=false
	LearnUnsolicited *bool                                   `json:"learn-unsolicited,omitempty"`
	Neighbor         []*InterfaceSubinterfaceIpv4ArpNeighbor `json:"neighbor,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	Timeout *uint16 `json:"timeout,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpEvpn struct
type InterfaceSubinterfaceIpv4ArpEvpn struct {
	Advertise []*InterfaceSubinterfaceIpv4ArpEvpnAdvertise `json:"advertise,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpEvpnAdvertise struct
type InterfaceSubinterfaceIpv4ArpEvpnAdvertise struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	AdminTag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	RouteType *string `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpHostRoute struct
type InterfaceSubinterfaceIpv4ArpHostRoute struct {
	Populate []*InterfaceSubinterfaceIpv4ArpHostRoutePopulate `json:"populate,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpHostRoutePopulate struct
type InterfaceSubinterfaceIpv4ArpHostRoutePopulate struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	AdminTag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	RouteType *string `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv4ArpNeighbor struct
type InterfaceSubinterfaceIpv4ArpNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4Address *string `json:"ipv4-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	LinkLayerAddress *string `json:"link-layer-address"`
}

// InterfaceSubinterfaceIpv4DhcpClient struct
type InterfaceSubinterfaceIpv4DhcpClient struct {
	TraceOptions *InterfaceSubinterfaceIpv4DhcpClientTraceOptions `json:"trace-options,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcpClientTraceOptions struct
type InterfaceSubinterfaceIpv4DhcpClientTraceOptions struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcpRelay struct
type InterfaceSubinterfaceIpv4DhcpRelay struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	GiAddress *string `json:"gi-address,omitempty"`
	// +kubebuilder:validation:Enum=`circuit-id`;`remote-id`
	Option *string `json:"option,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server       *string                                         `json:"server,omitempty"`
	TraceOptions *InterfaceSubinterfaceIpv4DhcpRelayTraceOptions `json:"trace-options,omitempty"`
	// +kubebuilder:default:=false
	UseGiAddrAsSrcIpAddr *bool `json:"use-gi-addr-as-src-ip-addr,omitempty"`
}

// InterfaceSubinterfaceIpv4DhcpRelayTraceOptions struct
type InterfaceSubinterfaceIpv4DhcpRelayTraceOptions struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv4Vrrp struct
type InterfaceSubinterfaceIpv4Vrrp struct {
	VrrpGroup []*InterfaceSubinterfaceIpv4VrrpVrrpGroup `json:"vrrp-group,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpGroup struct
type InterfaceSubinterfaceIpv4VrrpVrrpGroup struct {
	AcceptMode *bool `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	AdvertiseInterval *uint16                                               `json:"advertise-interval,omitempty"`
	Authentication    *InterfaceSubinterfaceIpv4VrrpVrrpGroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	InitDelay         *uint16                                                  `json:"init-delay,omitempty"`
	InterfaceTracking *InterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:default:=false
	MasterInheritInterval *bool `json:"master-inherit-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	OperInterval *uint16 `json:"oper-interval,omitempty"`
	Preempt      *bool   `json:"preempt,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	PreemptDelay *uint16 `json:"preempt-delay,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority   *uint8                                            `json:"priority,omitempty"`
	Statistics *InterfaceSubinterfaceIpv4VrrpVrrpGroupStatistics `json:"statistics,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=2
	Version *uint8 `json:"version,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	VirtualAddress *string `json:"virtual-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	VirtualRouterId *uint8 `json:"virtual-router-id,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpGroupAuthentication struct
type InterfaceSubinterfaceIpv4VrrpVrrpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTracking struct
type InterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTracking struct {
	TrackInterface []*InterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTrackingTrackInterface `json:"track-interface,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTrackingTrackInterface struct
type InterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTrackingTrackInterface struct {
	Interface *string `json:"interface,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	PriorityDecrement *uint8 `json:"priority-decrement,omitempty"`
}

// InterfaceSubinterfaceIpv4VrrpVrrpGroupStatistics struct
type InterfaceSubinterfaceIpv4VrrpVrrpGroupStatistics struct {
}

// InterfaceSubinterfaceIpv6 struct
type InterfaceSubinterfaceIpv6 struct {
	Address             []*InterfaceSubinterfaceIpv6Address           `json:"address,omitempty"`
	DhcpClient          *InterfaceSubinterfaceIpv6DhcpClient          `json:"dhcp-client,omitempty"`
	DhcpRelay           *InterfaceSubinterfaceIpv6DhcpRelay           `json:"dhcp-relay,omitempty"`
	NeighborDiscovery   *InterfaceSubinterfaceIpv6NeighborDiscovery   `json:"neighbor-discovery,omitempty"`
	RouterAdvertisement *InterfaceSubinterfaceIpv6RouterAdvertisement `json:"router-advertisement,omitempty"`
	Vrrp                *InterfaceSubinterfaceIpv6Vrrp                `json:"vrrp,omitempty"`
}

// InterfaceSubinterfaceIpv6Address struct
type InterfaceSubinterfaceIpv6Address struct {
	AnycastGw *bool `json:"anycast-gw,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	Primary  *string `json:"primary,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcpClient struct
type InterfaceSubinterfaceIpv6DhcpClient struct {
	TraceOptions *InterfaceSubinterfaceIpv6DhcpClientTraceOptions `json:"trace-options,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcpClientTraceOptions struct
type InterfaceSubinterfaceIpv6DhcpClientTraceOptions struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcpRelay struct
type InterfaceSubinterfaceIpv6DhcpRelay struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Enum=`interface-id`;`remote-id`
	Option *string `json:"option,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))|((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server *string `json:"server,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	SourceAddress *string                                         `json:"source-address,omitempty"`
	TraceOptions  *InterfaceSubinterfaceIpv6DhcpRelayTraceOptions `json:"trace-options,omitempty"`
}

// InterfaceSubinterfaceIpv6DhcpRelayTraceOptions struct
type InterfaceSubinterfaceIpv6DhcpRelayTraceOptions struct {
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighborDiscovery struct
type InterfaceSubinterfaceIpv6NeighborDiscovery struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug *string `json:"debug,omitempty"`
	// +kubebuilder:default:=true
	DuplicateAddressDetection *bool                                                `json:"duplicate-address-detection,omitempty"`
	Evpn                      *InterfaceSubinterfaceIpv6NeighborDiscoveryEvpn      `json:"evpn,omitempty"`
	HostRoute                 *InterfaceSubinterfaceIpv6NeighborDiscoveryHostRoute `json:"host-route,omitempty"`
	// +kubebuilder:validation:Enum=`both`;`global`;`link-local`;`none`
	// +kubebuilder:default:=none
	LearnUnsolicited *string                                               `json:"learn-unsolicited,omitempty"`
	Neighbor         []*InterfaceSubinterfaceIpv6NeighborDiscoveryNeighbor `json:"neighbor,omitempty"`
	// kubebuilder:validation:Minimum=30
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=30
	ReachableTime *uint32 `json:"reachable-time,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	StaleTime *uint32 `json:"stale-time,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighborDiscoveryEvpn struct
type InterfaceSubinterfaceIpv6NeighborDiscoveryEvpn struct {
	Advertise []*InterfaceSubinterfaceIpv6NeighborDiscoveryEvpnAdvertise `json:"advertise,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighborDiscoveryEvpnAdvertise struct
type InterfaceSubinterfaceIpv6NeighborDiscoveryEvpnAdvertise struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	AdminTag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	RouteType *string `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighborDiscoveryHostRoute struct
type InterfaceSubinterfaceIpv6NeighborDiscoveryHostRoute struct {
	Populate []*InterfaceSubinterfaceIpv6NeighborDiscoveryHostRoutePopulate `json:"populate,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighborDiscoveryHostRoutePopulate struct
type InterfaceSubinterfaceIpv6NeighborDiscoveryHostRoutePopulate struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	AdminTag *uint32 `json:"admin-tag,omitempty"`
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	RouteType *string `json:"route-type,omitempty"`
}

// InterfaceSubinterfaceIpv6NeighborDiscoveryNeighbor struct
type InterfaceSubinterfaceIpv6NeighborDiscoveryNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6Address *string `json:"ipv6-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	LinkLayerAddress *string `json:"link-layer-address"`
}

// InterfaceSubinterfaceIpv6RouterAdvertisement struct
type InterfaceSubinterfaceIpv6RouterAdvertisement struct {
	// +kubebuilder:validation:Enum=`messages`
	Debug      *string                                                 `json:"debug,omitempty"`
	RouterRole *InterfaceSubinterfaceIpv6RouterAdvertisementRouterRole `json:"router-role,omitempty"`
}

// InterfaceSubinterfaceIpv6RouterAdvertisementRouterRole struct
type InterfaceSubinterfaceIpv6RouterAdvertisementRouterRole struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=64
	CurrentHopLimit *uint8 `json:"current-hop-limit,omitempty"`
	// kubebuilder:validation:Minimum=1280
	// kubebuilder:validation:Maximum=9486
	IpMtu *uint16 `json:"ip-mtu,omitempty"`
	// +kubebuilder:default:=false
	ManagedConfigurationFlag *bool `json:"managed-configuration-flag,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=600
	MaxAdvertisementInterval *uint16 `json:"max-advertisement-interval,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=1350
	// +kubebuilder:default:=200
	MinAdvertisementInterval *uint16 `json:"min-advertisement-interval,omitempty"`
	// +kubebuilder:default:=false
	OtherConfigurationFlag *bool                                                           `json:"other-configuration-flag,omitempty"`
	Prefix                 []*InterfaceSubinterfaceIpv6RouterAdvertisementRouterRolePrefix `json:"prefix,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600000
	// +kubebuilder:default:=0
	ReachableTime *uint32 `json:"reachable-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1800000
	// +kubebuilder:default:=0
	RetransmitTime *uint32 `json:"retransmit-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9000
	// +kubebuilder:default:=1800
	RouterLifetime *uint16 `json:"router-lifetime,omitempty"`
}

// InterfaceSubinterfaceIpv6RouterAdvertisementRouterRolePrefix struct
type InterfaceSubinterfaceIpv6RouterAdvertisementRouterRolePrefix struct {
	// +kubebuilder:default:=false
	AutonomousFlag *bool `json:"autonomous-flag,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipv6Prefix *string `json:"ipv6-prefix,omitempty"`
	// +kubebuilder:default:=false
	OnLinkFlag *bool `json:"on-link-flag,omitempty"`
	// +kubebuilder:default:=604800
	PreferredLifetime *uint32 `json:"preferred-lifetime,omitempty"`
	// +kubebuilder:default:=2592000
	ValidLifetime *uint32 `json:"valid-lifetime,omitempty"`
}

// InterfaceSubinterfaceIpv6Vrrp struct
type InterfaceSubinterfaceIpv6Vrrp struct {
	VrrpGroup []*InterfaceSubinterfaceIpv6VrrpVrrpGroup `json:"vrrp-group,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpGroup struct
type InterfaceSubinterfaceIpv6VrrpVrrpGroup struct {
	AcceptMode *bool `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	AdvertiseInterval *uint16                                               `json:"advertise-interval,omitempty"`
	Authentication    *InterfaceSubinterfaceIpv6VrrpVrrpGroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	InitDelay         *uint16                                                  `json:"init-delay,omitempty"`
	InterfaceTracking *InterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:default:=false
	MasterInheritInterval *bool `json:"master-inherit-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	OperInterval *uint16 `json:"oper-interval,omitempty"`
	Preempt      *bool   `json:"preempt,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	PreemptDelay *uint16 `json:"preempt-delay,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority   *uint8                                            `json:"priority,omitempty"`
	Statistics *InterfaceSubinterfaceIpv6VrrpVrrpGroupStatistics `json:"statistics,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=3
	Version *uint8 `json:"version,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	VirtualAddress *string `json:"virtual-address,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	VirtualRouterId *uint8 `json:"virtual-router-id,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpGroupAuthentication struct
type InterfaceSubinterfaceIpv6VrrpVrrpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTracking struct
type InterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTracking struct {
	TrackInterface []*InterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTrackingTrackInterface `json:"track-interface,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTrackingTrackInterface struct
type InterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTrackingTrackInterface struct {
	Interface *string `json:"interface,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	PriorityDecrement *uint8 `json:"priority-decrement,omitempty"`
}

// InterfaceSubinterfaceIpv6VrrpVrrpGroupStatistics struct
type InterfaceSubinterfaceIpv6VrrpVrrpGroupStatistics struct {
}

// InterfaceSubinterfaceLocalMirrorDestination struct
type InterfaceSubinterfaceLocalMirrorDestination struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
}

// InterfaceSubinterfaceQos struct
type InterfaceSubinterfaceQos struct {
	Input  *InterfaceSubinterfaceQosInput  `json:"input,omitempty"`
	Output *InterfaceSubinterfaceQosOutput `json:"output,omitempty"`
}

// InterfaceSubinterfaceQosInput struct
type InterfaceSubinterfaceQosInput struct {
	Classifiers *InterfaceSubinterfaceQosInputClassifiers `json:"classifiers,omitempty"`
}

// InterfaceSubinterfaceQosInputClassifiers struct
type InterfaceSubinterfaceQosInputClassifiers struct {
	Dscp             *string `json:"dscp,omitempty"`
	Ipv4Dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6Dscp         *string `json:"ipv6-dscp,omitempty"`
	MplsTrafficClass *string `json:"mpls-traffic-class,omitempty"`
}

// InterfaceSubinterfaceQosOutput struct
type InterfaceSubinterfaceQosOutput struct {
	RewriteRules *InterfaceSubinterfaceQosOutputRewriteRules `json:"rewrite-rules,omitempty"`
}

// InterfaceSubinterfaceQosOutputRewriteRules struct
type InterfaceSubinterfaceQosOutputRewriteRules struct {
	Dscp             *string `json:"dscp,omitempty"`
	Ipv4Dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6Dscp         *string `json:"ipv6-dscp,omitempty"`
	MplsTrafficClass *string `json:"mpls-traffic-class,omitempty"`
}

// InterfaceSubinterfaceVlan struct
type InterfaceSubinterfaceVlan struct {
	Encap *InterfaceSubinterfaceVlanEncap `json:"encap,omitempty"`
}

// InterfaceSubinterfaceVlanEncap struct
type InterfaceSubinterfaceVlanEncap struct {
	SingleTagged *InterfaceSubinterfaceVlanEncapSingleTagged `json:"single-tagged,omitempty"`
	Untagged     *InterfaceSubinterfaceVlanEncapUntagged     `json:"untagged,omitempty"`
}

// InterfaceSubinterfaceVlanEncapSingleTagged struct
type InterfaceSubinterfaceVlanEncapSingleTagged struct {
	VlanId *string `json:"vlan-id,omitempty"`
}

// InterfaceSubinterfaceVlanEncapUntagged struct
type InterfaceSubinterfaceVlanEncapUntagged struct {
}

// InterfaceSubinterfaceSpec struct
type InterfaceSubinterfaceParameters struct {
	InterfaceName            *string                `json:"interface-name"`
	SrlInterfaceSubinterface *InterfaceSubinterface `json:"subinterface,omitempty"`
}

// InterfaceSubinterfaceStatus struct
type InterfaceSubinterfaceObservation struct {
}

// A InterfaceSubinterfaceSpec defines the desired state of a InterfaceSubinterface.
type InterfaceSubinterfaceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     InterfaceSubinterfaceParameters `json:"forNetworkNode"`
}

// A InterfaceSubinterfaceStatus represents the observed state of a InterfaceSubinterface.
type InterfaceSubinterfaceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        InterfaceSubinterfaceObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlInterfaceSubinterface is the Schema for the InterfaceSubinterface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlInterfaceSubinterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSubinterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceSubinterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlInterfaceSubinterfaceList contains a list of InterfaceSubinterfaces
type SrlInterfaceSubinterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlInterfaceSubinterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlInterfaceSubinterface{}, &SrlInterfaceSubinterfaceList{})
}

// InterfaceSubinterface type metadata.
var (
	InterfaceSubinterfaceKind             = reflect.TypeOf(SrlInterfaceSubinterface{}).Name()
	InterfaceSubinterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceSubinterfaceKind}.String()
	InterfaceSubinterfaceKindAPIVersion   = InterfaceSubinterfaceKind + "." + GroupVersion.String()
	InterfaceSubinterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceSubinterfaceKind)
)
