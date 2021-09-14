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
	// NetworkinstanceProtocolsBgpFinalizer is the name of the finalizer added to
	// NetworkinstanceProtocolsBgp to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceProtocolsBgpFinalizer string = "bgp.srl.ndd.yndd.io"
)

// NetworkinstanceProtocolsBgp struct
type NetworkinstanceProtocolsBgp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState     *string                                    `json:"admin-state,omitempty"`
	AsPathOptions  *NetworkinstanceProtocolsBgpAsPathOptions  `json:"as-path-options,omitempty"`
	Authentication *NetworkinstanceProtocolsBgpAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	AutonomousSystem  *uint32                                       `json:"autonomous-system"`
	Convergence       *NetworkinstanceProtocolsBgpConvergence       `json:"convergence,omitempty"`
	DynamicNeighbors  *NetworkinstanceProtocolsBgpDynamicNeighbors  `json:"dynamic-neighbors,omitempty"`
	EbgpDefaultPolicy *NetworkinstanceProtocolsBgpEbgpDefaultPolicy `json:"ebgp-default-policy,omitempty"`
	Evpn              *NetworkinstanceProtocolsBgpEvpn              `json:"evpn,omitempty"`
	ExportPolicy      *string                                       `json:"export-policy,omitempty"`
	FailureDetection  *NetworkinstanceProtocolsBgpFailureDetection  `json:"failure-detection,omitempty"`
	GracefulRestart   *NetworkinstanceProtocolsBgpGracefulRestart   `json:"graceful-restart,omitempty"`
	Group             []*NetworkinstanceProtocolsBgpGroup           `json:"group,omitempty"`
	ImportPolicy      *string                                       `json:"import-policy,omitempty"`
	Ipv4Unicast       *NetworkinstanceProtocolsBgpIpv4Unicast       `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast       *NetworkinstanceProtocolsBgpIpv6Unicast       `json:"ipv6-unicast,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	LocalPreference    *uint32                                        `json:"local-preference,omitempty"`
	Neighbor           []*NetworkinstanceProtocolsBgpNeighbor         `json:"neighbor,omitempty"`
	Preference         *NetworkinstanceProtocolsBgpPreference         `json:"preference,omitempty"`
	RouteAdvertisement *NetworkinstanceProtocolsBgpRouteAdvertisement `json:"route-advertisement,omitempty"`
	RouteReflector     *NetworkinstanceProtocolsBgpRouteReflector     `json:"route-reflector,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	RouterId      *string                                   `json:"router-id"`
	SendCommunity *NetworkinstanceProtocolsBgpSendCommunity `json:"send-community,omitempty"`
	TraceOptions  *NetworkinstanceProtocolsBgpTraceOptions  `json:"trace-options,omitempty"`
	Transport     *NetworkinstanceProtocolsBgpTransport     `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsBgpAsPathOptions struct
type NetworkinstanceProtocolsBgpAsPathOptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	AllowOwnAs      *uint8                                                   `json:"allow-own-as,omitempty"`
	RemovePrivateAs *NetworkinstanceProtocolsBgpAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
}

// NetworkinstanceProtocolsBgpAsPathOptionsRemovePrivateAs struct
type NetworkinstanceProtocolsBgpAsPathOptionsRemovePrivateAs struct {
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	LeadingOnly  *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode *string `json:"mode,omitempty"`
}

// NetworkinstanceProtocolsBgpAuthentication struct
type NetworkinstanceProtocolsBgpAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpConvergence struct
type NetworkinstanceProtocolsBgpConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	MinWaitToAdvertise *uint16 `json:"min-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicNeighbors struct
type NetworkinstanceProtocolsBgpDynamicNeighbors struct {
	Accept *NetworkinstanceProtocolsBgpDynamicNeighborsAccept `json:"accept,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicNeighborsAccept struct
type NetworkinstanceProtocolsBgpDynamicNeighborsAccept struct {
	Match []*NetworkinstanceProtocolsBgpDynamicNeighborsAcceptMatch `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	MaxSessions *uint16 `json:"max-sessions,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicNeighborsAcceptMatch struct
type NetworkinstanceProtocolsBgpDynamicNeighborsAcceptMatch struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([1-9][0-9]*)|([1-9][0-9]*)\.\.([1-9][0-9]*)`
	AllowedPeerAs *string `json:"allowed-peer-as,omitempty"`
	PeerGroup     *string `json:"peer-group"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
}

// NetworkinstanceProtocolsBgpEbgpDefaultPolicy struct
type NetworkinstanceProtocolsBgpEbgpDefaultPolicy struct {
	ExportRejectAll *bool `json:"export-reject-all,omitempty"`
	ImportRejectAll *bool `json:"import-reject-all,omitempty"`
}

// NetworkinstanceProtocolsBgpEvpn struct
type NetworkinstanceProtocolsBgpEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                   `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                     `json:"advertise-ipv6-next-hops,omitempty"`
	KeepAllRoutes         *bool                                     `json:"keep-all-routes,omitempty"`
	Multipath             *NetworkinstanceProtocolsBgpEvpnMultipath `json:"multipath,omitempty"`
	RapidUpdate           *bool                                     `json:"rapid-update,omitempty"`
}

// NetworkinstanceProtocolsBgpEvpnMultipath struct
type NetworkinstanceProtocolsBgpEvpnMultipath struct {
	AllowMultipleAs *bool `json:"allow-multiple-as,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxPathsLevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxPathsLevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// NetworkinstanceProtocolsBgpFailureDetection struct
type NetworkinstanceProtocolsBgpFailureDetection struct {
	EnableBfd    *bool `json:"enable-bfd,omitempty"`
	FastFailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpGracefulRestart struct
type NetworkinstanceProtocolsBgpGracefulRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	StaleRoutesTime *uint16 `json:"stale-routes-time,omitempty"`
}

// NetworkinstanceProtocolsBgpGroup struct
type NetworkinstanceProtocolsBgpGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState     *string                                         `json:"admin-state,omitempty"`
	AsPathOptions  *NetworkinstanceProtocolsBgpGroupAsPathOptions  `json:"as-path-options,omitempty"`
	Authentication *NetworkinstanceProtocolsBgpGroupAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description      *string                                           `json:"description,omitempty"`
	Evpn             *NetworkinstanceProtocolsBgpGroupEvpn             `json:"evpn,omitempty"`
	ExportPolicy     *string                                           `json:"export-policy,omitempty"`
	FailureDetection *NetworkinstanceProtocolsBgpGroupFailureDetection `json:"failure-detection,omitempty"`
	GracefulRestart  *NetworkinstanceProtocolsBgpGroupGracefulRestart  `json:"graceful-restart,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	GroupName    *string                                      `json:"group-name"`
	ImportPolicy *string                                      `json:"import-policy,omitempty"`
	Ipv4Unicast  *NetworkinstanceProtocolsBgpGroupIpv4Unicast `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast  *NetworkinstanceProtocolsBgpGroupIpv6Unicast `json:"ipv6-unicast,omitempty"`
	LocalAs      []*NetworkinstanceProtocolsBgpGroupLocalAs   `json:"local-as,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	LocalPreference *uint32 `json:"local-preference,omitempty"`
	NextHopSelf     *bool   `json:"next-hop-self,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	PeerAs           *uint32                                           `json:"peer-as,omitempty"`
	RouteReflector   *NetworkinstanceProtocolsBgpGroupRouteReflector   `json:"route-reflector,omitempty"`
	SendCommunity    *NetworkinstanceProtocolsBgpGroupSendCommunity    `json:"send-community,omitempty"`
	SendDefaultRoute *NetworkinstanceProtocolsBgpGroupSendDefaultRoute `json:"send-default-route,omitempty"`
	Timers           *NetworkinstanceProtocolsBgpGroupTimers           `json:"timers,omitempty"`
	TraceOptions     *NetworkinstanceProtocolsBgpGroupTraceOptions     `json:"trace-options,omitempty"`
	Transport        *NetworkinstanceProtocolsBgpGroupTransport        `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupAsPathOptions struct
type NetworkinstanceProtocolsBgpGroupAsPathOptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	AllowOwnAs      *uint8                                                        `json:"allow-own-as,omitempty"`
	RemovePrivateAs *NetworkinstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
	ReplacePeerAs   *bool                                                         `json:"replace-peer-as,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs struct
type NetworkinstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs struct {
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	LeadingOnly  *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode *string `json:"mode"`
}

// NetworkinstanceProtocolsBgpGroupAuthentication struct
type NetworkinstanceProtocolsBgpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupEvpn struct
type NetworkinstanceProtocolsBgpGroupEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                          `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                            `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpGroupEvpnPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupEvpnPrefixLimit struct
type NetworkinstanceProtocolsBgpGroupEvpnPrefixLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupFailureDetection struct
type NetworkinstanceProtocolsBgpGroupFailureDetection struct {
	EnableBfd    *bool `json:"enable-bfd,omitempty"`
	FastFailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupGracefulRestart struct
type NetworkinstanceProtocolsBgpGroupGracefulRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	StaleRoutesTime *uint16 `json:"stale-routes-time,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv4Unicast struct
type NetworkinstanceProtocolsBgpGroupIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                 `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                   `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpGroupIpv4UnicastPrefixLimit `json:"prefix-limit,omitempty"`
	ReceiveIpv6NextHops   *bool                                                   `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv4UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpGroupIpv4UnicastPrefixLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv6Unicast struct
type NetworkinstanceProtocolsBgpGroupIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                                 `json:"admin-state,omitempty"`
	PrefixLimit *NetworkinstanceProtocolsBgpGroupIpv6UnicastPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv6UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpGroupIpv6UnicastPrefixLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupLocalAs struct
type NetworkinstanceProtocolsBgpGroupLocalAs struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	AsNumber        *uint32 `json:"as-number"`
	PrependGlobalAs *bool   `json:"prepend-global-as,omitempty"`
	PrependLocalAs  *bool   `json:"prepend-local-as,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupRouteReflector struct
type NetworkinstanceProtocolsBgpGroupRouteReflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupSendCommunity struct
type NetworkinstanceProtocolsBgpGroupSendCommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupSendDefaultRoute struct
type NetworkinstanceProtocolsBgpGroupSendDefaultRoute struct {
	ExportPolicy *string `json:"export-policy,omitempty"`
	Ipv4Unicast  *bool   `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast  *bool   `json:"ipv6-unicast,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTimers struct
type NetworkinstanceProtocolsBgpGroupTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	ConnectRetry *uint16 `json:"connect-retry,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	HoldTime *uint16 `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=21845
	KeepaliveInterval *uint16 `json:"keepalive-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	MinimumAdvertisementInterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTraceOptions struct
type NetworkinstanceProtocolsBgpGroupTraceOptions struct {
	Flag []*NetworkinstanceProtocolsBgpGroupTraceOptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTraceOptionsFlag struct
type NetworkinstanceProtocolsBgpGroupTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
}

// NetworkinstanceProtocolsBgpGroupTransport struct
type NetworkinstanceProtocolsBgpGroupTransport struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
	PassiveMode  *bool   `json:"passive-mode,omitempty"`
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4Unicast struct
type NetworkinstanceProtocolsBgpIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                            `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                              `json:"advertise-ipv6-next-hops,omitempty"`
	Convergence           *NetworkinstanceProtocolsBgpIpv4UnicastConvergence `json:"convergence,omitempty"`
	Multipath             *NetworkinstanceProtocolsBgpIpv4UnicastMultipath   `json:"multipath,omitempty"`
	ReceiveIpv6NextHops   *bool                                              `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4UnicastConvergence struct
type NetworkinstanceProtocolsBgpIpv4UnicastConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	MaxWaitToAdvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4UnicastMultipath struct
type NetworkinstanceProtocolsBgpIpv4UnicastMultipath struct {
	AllowMultipleAs *bool `json:"allow-multiple-as,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxPathsLevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxPathsLevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6Unicast struct
type NetworkinstanceProtocolsBgpIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                            `json:"admin-state,omitempty"`
	Convergence *NetworkinstanceProtocolsBgpIpv6UnicastConvergence `json:"convergence,omitempty"`
	Multipath   *NetworkinstanceProtocolsBgpIpv6UnicastMultipath   `json:"multipath,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6UnicastConvergence struct
type NetworkinstanceProtocolsBgpIpv6UnicastConvergence struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=3600
	MaxWaitToAdvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6UnicastMultipath struct
type NetworkinstanceProtocolsBgpIpv6UnicastMultipath struct {
	AllowMultipleAs *bool `json:"allow-multiple-as,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxPathsLevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxPathsLevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighbor struct
type NetworkinstanceProtocolsBgpNeighbor struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState     *string                                            `json:"admin-state,omitempty"`
	AsPathOptions  *NetworkinstanceProtocolsBgpNeighborAsPathOptions  `json:"as-path-options,omitempty"`
	Authentication *NetworkinstanceProtocolsBgpNeighborAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description      *string                                              `json:"description,omitempty"`
	Evpn             *NetworkinstanceProtocolsBgpNeighborEvpn             `json:"evpn,omitempty"`
	ExportPolicy     *string                                              `json:"export-policy,omitempty"`
	FailureDetection *NetworkinstanceProtocolsBgpNeighborFailureDetection `json:"failure-detection,omitempty"`
	GracefulRestart  *NetworkinstanceProtocolsBgpNeighborGracefulRestart  `json:"graceful-restart,omitempty"`
	ImportPolicy     *string                                              `json:"import-policy,omitempty"`
	Ipv4Unicast      *NetworkinstanceProtocolsBgpNeighborIpv4Unicast      `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast      *NetworkinstanceProtocolsBgpNeighborIpv6Unicast      `json:"ipv6-unicast,omitempty"`
	LocalAs          []*NetworkinstanceProtocolsBgpNeighborLocalAs        `json:"local-as,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	LocalPreference *uint32 `json:"local-preference,omitempty"`
	NextHopSelf     *bool   `json:"next-hop-self,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	PeerAddress *string `json:"peer-address"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	PeerAs           *uint32                                              `json:"peer-as,omitempty"`
	PeerGroup        *string                                              `json:"peer-group"`
	RouteReflector   *NetworkinstanceProtocolsBgpNeighborRouteReflector   `json:"route-reflector,omitempty"`
	SendCommunity    *NetworkinstanceProtocolsBgpNeighborSendCommunity    `json:"send-community,omitempty"`
	SendDefaultRoute *NetworkinstanceProtocolsBgpNeighborSendDefaultRoute `json:"send-default-route,omitempty"`
	Timers           *NetworkinstanceProtocolsBgpNeighborTimers           `json:"timers,omitempty"`
	TraceOptions     *NetworkinstanceProtocolsBgpNeighborTraceOptions     `json:"trace-options,omitempty"`
	Transport        *NetworkinstanceProtocolsBgpNeighborTransport        `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborAsPathOptions struct
type NetworkinstanceProtocolsBgpNeighborAsPathOptions struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	AllowOwnAs      *uint8                                                           `json:"allow-own-as,omitempty"`
	RemovePrivateAs *NetworkinstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
	ReplacePeerAs   *bool                                                            `json:"replace-peer-as,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs struct
type NetworkinstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs struct {
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	LeadingOnly  *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode *string `json:"mode"`
}

// NetworkinstanceProtocolsBgpNeighborAuthentication struct
type NetworkinstanceProtocolsBgpNeighborAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborEvpn struct
type NetworkinstanceProtocolsBgpNeighborEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                             `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                               `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpNeighborEvpnPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborEvpnPrefixLimit struct
type NetworkinstanceProtocolsBgpNeighborEvpnPrefixLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborFailureDetection struct
type NetworkinstanceProtocolsBgpNeighborFailureDetection struct {
	EnableBfd    *bool `json:"enable-bfd,omitempty"`
	FastFailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborGracefulRestart struct
type NetworkinstanceProtocolsBgpNeighborGracefulRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=3600
	StaleRoutesTime *uint16                                                        `json:"stale-routes-time,omitempty"`
	WarmRestart     *NetworkinstanceProtocolsBgpNeighborGracefulRestartWarmRestart `json:"warm-restart,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborGracefulRestartWarmRestart struct
type NetworkinstanceProtocolsBgpNeighborGracefulRestartWarmRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv4Unicast struct
type NetworkinstanceProtocolsBgpNeighborIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                    `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                      `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit `json:"prefix-limit,omitempty"`
	ReceiveIpv6NextHops   *bool                                                      `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv6Unicast struct
type NetworkinstanceProtocolsBgpNeighborIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                                    `json:"admin-state,omitempty"`
	PrefixLimit *NetworkinstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborLocalAs struct
type NetworkinstanceProtocolsBgpNeighborLocalAs struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	AsNumber        *uint32 `json:"as-number"`
	PrependGlobalAs *bool   `json:"prepend-global-as,omitempty"`
	PrependLocalAs  *bool   `json:"prepend-local-as,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborRouteReflector struct
type NetworkinstanceProtocolsBgpNeighborRouteReflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborSendCommunity struct
type NetworkinstanceProtocolsBgpNeighborSendCommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborSendDefaultRoute struct
type NetworkinstanceProtocolsBgpNeighborSendDefaultRoute struct {
	ExportPolicy *string `json:"export-policy,omitempty"`
	Ipv4Unicast  *bool   `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast  *bool   `json:"ipv6-unicast,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTimers struct
type NetworkinstanceProtocolsBgpNeighborTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	ConnectRetry *uint16 `json:"connect-retry,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=0
	HoldTime *uint16 `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=21845
	KeepaliveInterval *uint16 `json:"keepalive-interval,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	MinimumAdvertisementInterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTraceOptions struct
type NetworkinstanceProtocolsBgpNeighborTraceOptions struct {
	Flag []*NetworkinstanceProtocolsBgpNeighborTraceOptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTraceOptionsFlag struct
type NetworkinstanceProtocolsBgpNeighborTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
}

// NetworkinstanceProtocolsBgpNeighborTransport struct
type NetworkinstanceProtocolsBgpNeighborTransport struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
	PassiveMode  *bool   `json:"passive-mode,omitempty"`
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgpPreference struct
type NetworkinstanceProtocolsBgpPreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Ebgp *uint8 `json:"ebgp,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Ibgp *uint8 `json:"ibgp,omitempty"`
}

// NetworkinstanceProtocolsBgpRouteAdvertisement struct
type NetworkinstanceProtocolsBgpRouteAdvertisement struct {
	RapidWithdrawal   *bool `json:"rapid-withdrawal,omitempty"`
	WaitForFibInstall *bool `json:"wait-for-fib-install,omitempty"`
}

// NetworkinstanceProtocolsBgpRouteReflector struct
type NetworkinstanceProtocolsBgpRouteReflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpSendCommunity struct
type NetworkinstanceProtocolsBgpSendCommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpTraceOptions struct
type NetworkinstanceProtocolsBgpTraceOptions struct {
	Flag []*NetworkinstanceProtocolsBgpTraceOptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpTraceOptionsFlag struct
type NetworkinstanceProtocolsBgpTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
}

// NetworkinstanceProtocolsBgpTransport struct
type NetworkinstanceProtocolsBgpTransport struct {
	// kubebuilder:validation:Minimum=536
	// kubebuilder:validation:Maximum=9446
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgpParameters are the parameter fields of a NetworkinstanceProtocolsBgp.
type NetworkinstanceProtocolsBgpParameters struct {
	NetworkInstanceName            *string                      `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsBgp *NetworkinstanceProtocolsBgp `json:"bgp,omitempty"`
}

// NetworkinstanceProtocolsBgpObservation are the observable fields of a NetworkinstanceProtocolsBgp.
type NetworkinstanceProtocolsBgpObservation struct {
}

// A NetworkinstanceProtocolsBgpSpec defines the desired state of a NetworkinstanceProtocolsBgp.
type NetworkinstanceProtocolsBgpSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceProtocolsBgpParameters `json:"forNetworkNode"`
}

// A NetworkinstanceProtocolsBgpStatus represents the observed state of a NetworkinstanceProtocolsBgp.
type NetworkinstanceProtocolsBgpStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceProtocolsBgpObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgp is the Schema for the NetworkinstanceProtocolsBgp API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceProtocolsBgp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsBgpSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsBgpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpList contains a list of NetworkinstanceProtocolsBgps
type SrlNetworkinstanceProtocolsBgpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsBgp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsBgp{}, &SrlNetworkinstanceProtocolsBgpList{})
}

// NetworkinstanceProtocolsBgp type metadata.
var (
	NetworkinstanceProtocolsBgpKind             = reflect.TypeOf(SrlNetworkinstanceProtocolsBgp{}).Name()
	NetworkinstanceProtocolsBgpGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsBgpKind}.String()
	NetworkinstanceProtocolsBgpKindAPIVersion   = NetworkinstanceProtocolsBgpKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsBgpGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsBgpKind)
)
