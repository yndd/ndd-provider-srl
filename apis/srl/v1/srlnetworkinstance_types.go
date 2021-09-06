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
	// NetworkinstanceFinalizer is the name of the finalizer added to
	// Networkinstance to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceFinalizer string = "networkInstance.srl.ndd.yndd.io"
)

// Networkinstance struct
type Networkinstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState  *string                     `json:"admin-state,omitempty"`
	BridgeTable *NetworkinstanceBridgeTable `json:"bridge-table,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description     *string                         `json:"description,omitempty"`
	Interface       []*NetworkinstanceInterface     `json:"interface,omitempty"`
	IpForwarding    *NetworkinstanceIpForwarding    `json:"ip-forwarding,omitempty"`
	IpLoadBalancing *NetworkinstanceIpLoadBalancing `json:"ip-load-balancing,omitempty"`
	Mpls            *NetworkinstanceMpls            `json:"mpls,omitempty"`
	Mtu             *NetworkinstanceMtu             `json:"mtu,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name      *string                   `json:"name,omitempty"`
	Protocols *NetworkinstanceProtocols `json:"protocols,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId           *string                            `json:"router-id,omitempty"`
	TrafficEngineering *NetworkinstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	// +kubebuilder:default:=default
	Type           *string                          `json:"type,omitempty"`
	VxlanInterface []*NetworkinstanceVxlanInterface `json:"vxlan-interface,omitempty"`
}

// NetworkinstanceBridgeTable struct
type NetworkinstanceBridgeTable struct {
	// +kubebuilder:default:=false
	DiscardUnknownDestMac *bool                                     `json:"discard-unknown-dest-mac,omitempty"`
	MacDuplication        *NetworkinstanceBridgeTableMacDuplication `json:"mac-duplication,omitempty"`
	MacLearning           *NetworkinstanceBridgeTableMacLearning    `json:"mac-learning,omitempty"`
	MacLimit              *NetworkinstanceBridgeTableMacLimit       `json:"mac-limit,omitempty"`
	// +kubebuilder:default:=false
	ProtectAnycastGwMac *bool                                `json:"protect-anycast-gw-mac,omitempty"`
	StaticMac           *NetworkinstanceBridgeTableStaticMac `json:"static-mac,omitempty"`
}

// NetworkinstanceBridgeTableMacDuplication struct
type NetworkinstanceBridgeTableMacDuplication struct {
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`
	// +kubebuilder:default:=stop-learning
	Action *string `json:"action,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=9
	HoldDownTime *uint32 `json:"hold-down-time,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=3
	MonitoringWindow *uint32 `json:"monitoring-window,omitempty"`
	// kubebuilder:validation:Minimum=3
	// kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=5
	NumMoves *uint32 `json:"num-moves,omitempty"`
}

// NetworkinstanceBridgeTableMacLearning struct
type NetworkinstanceBridgeTableMacLearning struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                     `json:"admin-state,omitempty"`
	Aging      *NetworkinstanceBridgeTableMacLearningAging `json:"aging,omitempty"`
}

// NetworkinstanceBridgeTableMacLearningAging struct
type NetworkinstanceBridgeTableMacLearningAging struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=86400
	// +kubebuilder:default:=300
	AgeTime *int32 `json:"age-time,omitempty"`
}

// NetworkinstanceBridgeTableMacLimit struct
type NetworkinstanceBridgeTableMacLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	MaximumEntries *int32 `json:"maximum-entries,omitempty"`
	// kubebuilder:validation:Minimum=6
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	WarningThresholdPct *int32 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceBridgeTableStaticMac struct
type NetworkinstanceBridgeTableStaticMac struct {
	Mac []*NetworkinstanceBridgeTableStaticMacMac `json:"mac,omitempty"`
}

// NetworkinstanceBridgeTableStaticMacMac struct
type NetworkinstanceBridgeTableStaticMacMac struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Address     *string `json:"address,omitempty"`
	Destination *string `json:"destination"`
}

// NetworkinstanceInterface struct
type NetworkinstanceInterface struct {
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name,omitempty"`
}

// NetworkinstanceIpForwarding struct
type NetworkinstanceIpForwarding struct {
	ReceiveIpv4Check *bool `json:"receive-ipv4-check,omitempty"`
	ReceiveIpv6Check *bool `json:"receive-ipv6-check,omitempty"`
}

// NetworkinstanceIpLoadBalancing struct
type NetworkinstanceIpLoadBalancing struct {
	ResilientHashPrefix []*NetworkinstanceIpLoadBalancingResilientHashPrefix `json:"resilient-hash-prefix,omitempty"`
}

// NetworkinstanceIpLoadBalancingResilientHashPrefix struct
type NetworkinstanceIpLoadBalancingResilientHashPrefix struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=32
	// +kubebuilder:default:=1
	HashBucketsPerPath *uint8 `json:"hash-buckets-per-path,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPaths *uint8 `json:"max-paths,omitempty"`
}

// NetworkinstanceMpls struct
type NetworkinstanceMpls struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState      *string                               `json:"admin-state,omitempty"`
	StaticMplsEntry []*NetworkinstanceMplsStaticMplsEntry `json:"static-mpls-entry,omitempty"`
	// +kubebuilder:default:=false
	TtlPropagation *bool `json:"ttl-propagation,omitempty"`
}

// NetworkinstanceMplsStaticMplsEntry struct
type NetworkinstanceMplsStaticMplsEntry struct {
	// +kubebuilder:default:=false
	CollectStats *bool   `json:"collect-stats,omitempty"`
	NextHopGroup *string `json:"next-hop-group,omitempty"`
	// +kubebuilder:validation:Enum=`pop`;`swap`
	// +kubebuilder:default:=swap
	Operation *string `json:"operation,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Preference *uint8  `json:"preference,omitempty"`
	TopLabel   *string `json:"top-label,omitempty"`
}

// NetworkinstanceMtu struct
type NetworkinstanceMtu struct {
	PathMtuDiscovery *bool `json:"path-mtu-discovery,omitempty"`
}

// NetworkinstanceProtocols struct
type NetworkinstanceProtocols struct {
	DirectlyConnected *NetworkinstanceProtocolsDirectlyConnected `json:"directly-connected,omitempty"`
	Ldp               *NetworkinstanceProtocolsLdp               `json:"ldp,omitempty"`
}

// NetworkinstanceProtocolsDirectlyConnected struct
type NetworkinstanceProtocolsDirectlyConnected struct {
	TeDatabaseInstall *NetworkinstanceProtocolsDirectlyConnectedTeDatabaseInstall `json:"te-database-install,omitempty"`
}

// NetworkinstanceProtocolsDirectlyConnectedTeDatabaseInstall struct
type NetworkinstanceProtocolsDirectlyConnectedTeDatabaseInstall struct {
	BgpLs *NetworkinstanceProtocolsDirectlyConnectedTeDatabaseInstallBgpLs `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsDirectlyConnectedTeDatabaseInstallBgpLs struct
type NetworkinstanceProtocolsDirectlyConnectedTeDatabaseInstallBgpLs struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	BgpLsIdentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	IgpIdentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsLdp struct
type NetworkinstanceProtocolsLdp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState        *string                                     `json:"admin-state,omitempty"`
	Discovery         *NetworkinstanceProtocolsLdpDiscovery       `json:"discovery,omitempty"`
	DynamicLabelBlock *string                                     `json:"dynamic-label-block"`
	GracefulRestart   *NetworkinstanceProtocolsLdpGracefulRestart `json:"graceful-restart,omitempty"`
	Ipv4              *NetworkinstanceProtocolsLdpIpv4            `json:"ipv4,omitempty"`
	Multipath         *NetworkinstanceProtocolsLdpMultipath       `json:"multipath,omitempty"`
	Peers             *NetworkinstanceProtocolsLdpPeers           `json:"peers,omitempty"`
	TraceOptions      *NetworkinstanceProtocolsLdpTraceOptions    `json:"trace-options,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscovery struct
type NetworkinstanceProtocolsLdpDiscovery struct {
	Interfaces *NetworkinstanceProtocolsLdpDiscoveryInterfaces `json:"interfaces,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfaces struct
type NetworkinstanceProtocolsLdpDiscoveryInterfaces struct {
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=15
	HelloHoldtime *uint16 `json:"hello-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=5
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=5
	HelloInterval *uint16                                                    `json:"hello-interval,omitempty"`
	Interface     []*NetworkinstanceProtocolsLdpDiscoveryInterfacesInterface `json:"interface,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfacesInterface struct
type NetworkinstanceProtocolsLdpDiscoveryInterfacesInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=15
	HelloHoldtime *uint16 `json:"hello-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=5
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=5
	HelloInterval *uint16                                                      `json:"hello-interval,omitempty"`
	Ipv4          *NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 `json:"ipv4,omitempty"`
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name,omitempty"`
}

// NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 struct
type NetworkinstanceProtocolsLdpDiscoveryInterfacesInterfaceIpv4 struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
}

// NetworkinstanceProtocolsLdpGracefulRestart struct
type NetworkinstanceProtocolsLdpGracefulRestart struct {
	// +kubebuilder:default:=false
	HelperEnable *bool `json:"helper-enable,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=120
	MaxReconnectTime *uint16 `json:"max-reconnect-time,omitempty"`
	// kubebuilder:validation:Minimum=30
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=120
	MaxRecoveryTime *uint16 `json:"max-recovery-time,omitempty"`
}

// NetworkinstanceProtocolsLdpIpv4 struct
type NetworkinstanceProtocolsLdpIpv4 struct {
	FecResolution *NetworkinstanceProtocolsLdpIpv4FecResolution `json:"fec-resolution,omitempty"`
}

// NetworkinstanceProtocolsLdpIpv4FecResolution struct
type NetworkinstanceProtocolsLdpIpv4FecResolution struct {
	// +kubebuilder:default:=false
	LongestPrefix *bool `json:"longest-prefix,omitempty"`
}

// NetworkinstanceProtocolsLdpMultipath struct
type NetworkinstanceProtocolsLdpMultipath struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxPaths *uint8 `json:"max-paths,omitempty"`
}

// NetworkinstanceProtocolsLdpPeers struct
type NetworkinstanceProtocolsLdpPeers struct {
	Peer []*NetworkinstanceProtocolsLdpPeersPeer `json:"peer,omitempty"`
	// kubebuilder:validation:Minimum=45
	// kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=180
	SessionKeepaliveHoldtime *uint16 `json:"session-keepalive-holdtime,omitempty"`
	// kubebuilder:validation:Minimum=15
	// kubebuilder:validation:Maximum=1200
	// +kubebuilder:default:=60
	SessionKeepaliveInterval *uint16                                       `json:"session-keepalive-interval,omitempty"`
	TcpTransport             *NetworkinstanceProtocolsLdpPeersTcpTransport `json:"tcp-transport,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeer struct
type NetworkinstanceProtocolsLdpPeersPeer struct {
	Ipv4 *NetworkinstanceProtocolsLdpPeersPeerIpv4 `json:"ipv4,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	LabelSpaceId *uint16 `json:"label-space-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	LsrId        *string                                           `json:"lsr-id,omitempty"`
	TcpTransport *NetworkinstanceProtocolsLdpPeersPeerTcpTransport `json:"tcp-transport,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerIpv4 struct
type NetworkinstanceProtocolsLdpPeersPeerIpv4 struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	FecLimit *uint32 `json:"fec-limit,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerTcpTransport struct
type NetworkinstanceProtocolsLdpPeersPeerTcpTransport struct {
	Authentication *NetworkinstanceProtocolsLdpPeersPeerTcpTransportAuthentication `json:"authentication,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersPeerTcpTransportAuthentication struct
type NetworkinstanceProtocolsLdpPeersPeerTcpTransportAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersTcpTransport struct
type NetworkinstanceProtocolsLdpPeersTcpTransport struct {
	Authentication *NetworkinstanceProtocolsLdpPeersTcpTransportAuthentication `json:"authentication,omitempty"`
}

// NetworkinstanceProtocolsLdpPeersTcpTransportAuthentication struct
type NetworkinstanceProtocolsLdpPeersTcpTransportAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsLdpTraceOptions struct
type NetworkinstanceProtocolsLdpTraceOptions struct {
	Interface []*NetworkinstanceProtocolsLdpTraceOptionsInterface `json:"interface,omitempty"`
	Peer      []*NetworkinstanceProtocolsLdpTraceOptionsPeer      `json:"peer,omitempty"`
}

// NetworkinstanceProtocolsLdpTraceOptionsInterface struct
type NetworkinstanceProtocolsLdpTraceOptionsInterface struct {
	// kubebuilder:validation:MinLength=5
	// kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name,omitempty"`
}

// NetworkinstanceProtocolsLdpTraceOptionsPeer struct
type NetworkinstanceProtocolsLdpTraceOptionsPeer struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	LabelSpaceId *uint16 `json:"label-space-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	LsrId *string `json:"lsr-id,omitempty"`
}

// NetworkinstanceTrafficEngineering struct
type NetworkinstanceTrafficEngineering struct {
	AdminGroups *NetworkinstanceTrafficEngineeringAdminGroups `json:"admin-groups,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	AutonomousSystem *uint32                                       `json:"autonomous-system,omitempty"`
	Interface        []*NetworkinstanceTrafficEngineeringInterface `json:"interface,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4TeRouterId *string `json:"ipv4-te-router-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6TeRouterId       *string                                                `json:"ipv6-te-router-id,omitempty"`
	SharedRiskLinkGroups *NetworkinstanceTrafficEngineeringSharedRiskLinkGroups `json:"shared-risk-link-groups,omitempty"`
}

// NetworkinstanceTrafficEngineeringAdminGroups struct
type NetworkinstanceTrafficEngineeringAdminGroups struct {
	Group []*NetworkinstanceTrafficEngineeringAdminGroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceTrafficEngineeringAdminGroupsGroup struct
type NetworkinstanceTrafficEngineeringAdminGroupsGroup struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=31
	BitPosition *uint32 `json:"bit-position,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name,omitempty"`
}

// NetworkinstanceTrafficEngineeringInterface struct
type NetworkinstanceTrafficEngineeringInterface struct {
	AdminGroup     *string                                          `json:"admin-group,omitempty"`
	Delay          *NetworkinstanceTrafficEngineeringInterfaceDelay `json:"delay,omitempty"`
	InterfaceName  *string                                          `json:"interface-name,omitempty"`
	SrlgMembership *string                                          `json:"srlg-membership,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	TeMetric *uint32 `json:"te-metric,omitempty"`
}

// NetworkinstanceTrafficEngineeringInterfaceDelay struct
type NetworkinstanceTrafficEngineeringInterfaceDelay struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Static *uint32 `json:"static,omitempty"`
}

// NetworkinstanceTrafficEngineeringSharedRiskLinkGroups struct
type NetworkinstanceTrafficEngineeringSharedRiskLinkGroups struct {
	Group []*NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroup struct
type NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroup struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Cost *uint32 `json:"cost,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name         *string                                                                   `json:"name,omitempty"`
	StaticMember []*NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember `json:"static-member,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Value *uint32 `json:"value,omitempty"`
}

// NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember struct
type NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	FromAddress *string `json:"from-address,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	ToAddress *string `json:"to-address,omitempty"`
}

// NetworkinstanceVxlanInterface struct
type NetworkinstanceVxlanInterface struct {
	// kubebuilder:validation:MinLength=8
	// kubebuilder:validation:MaxLength=17
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,8}))`
	Name *string `json:"name,omitempty"`
}

// NetworkinstanceParameters struct defines the resource Parameters
type NetworkinstanceParameters struct {
	SrlNetworkinstance *Networkinstance `json:"network-instance,omitempty"`
}

// NetworkinstanceObservation struct defines the resource Observation
type NetworkinstanceObservation struct {
}

// A NetworkinstanceSpec defines the desired state of a Networkinstance.
type NetworkinstanceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceParameters `json:"forNetworkNode"`
}

// A NetworkinstanceStatus represents the observed state of a Networkinstance.
type NetworkinstanceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstance is the Schema for the Networkinstance API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceSpec   `json:"spec,omitempty"`
	Status NetworkinstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceList contains a list of Networkinstances
type SrlNetworkinstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstance{}, &SrlNetworkinstanceList{})
}

// Networkinstance type metadata.
var (
	NetworkinstanceKind             = reflect.TypeOf(SrlNetworkinstance{}).Name()
	NetworkinstanceGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceKind}.String()
	NetworkinstanceKindAPIVersion   = NetworkinstanceKind + "." + GroupVersion.String()
	NetworkinstanceGroupVersionKind = GroupVersion.WithKind(NetworkinstanceKind)
)
