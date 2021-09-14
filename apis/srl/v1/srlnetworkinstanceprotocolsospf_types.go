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
	// NetworkinstanceProtocolsOspfFinalizer is the name of the finalizer added to
	// NetworkinstanceProtocolsOspf to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceProtocolsOspfFinalizer string = "ospf.srl.ndd.yndd.io"
)

// NetworkinstanceProtocolsOspf struct
type NetworkinstanceProtocolsOspf struct {
	Instance []*NetworkinstanceProtocolsOspfInstance `json:"instance,omitempty"`
}

// NetworkinstanceProtocolsOspfInstance struct
type NetworkinstanceProtocolsOspfInstance struct {
	AddressFamily *string `json:"address-family,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Enum=`area`;`as`;`false`;`link`
	AdvertiseRouterCapability *string                                                 `json:"advertise-router-capability,omitempty"`
	Area                      []*NetworkinstanceProtocolsOspfInstanceArea             `json:"area,omitempty"`
	Asbr                      *NetworkinstanceProtocolsOspfInstanceAsbr               `json:"asbr,omitempty"`
	ExportLimit               *NetworkinstanceProtocolsOspfInstanceExportLimit        `json:"export-limit,omitempty"`
	ExportPolicy              *string                                                 `json:"export-policy,omitempty"`
	ExternalDbOverflow        *NetworkinstanceProtocolsOspfInstanceExternalDbOverflow `json:"external-db-overflow,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	ExternalPreference *uint8                                               `json:"external-preference,omitempty"`
	GracefulRestart    *NetworkinstanceProtocolsOspfInstanceGracefulRestart `json:"graceful-restart,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	InstanceId *uint32 `json:"instance-id,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxEcmpPaths *uint8 `json:"max-ecmp-paths,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name     *string                                       `json:"name"`
	Overload *NetworkinstanceProtocolsOspfInstanceOverload `json:"overload,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Preference *uint8 `json:"preference,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8000000000
	ReferenceBandwidth *uint64 `json:"reference-bandwidth,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId           *string                                                 `json:"router-id,omitempty"`
	TeDatabaseInstall  *NetworkinstanceProtocolsOspfInstanceTeDatabaseInstall  `json:"te-database-install,omitempty"`
	Timers             *NetworkinstanceProtocolsOspfInstanceTimers             `json:"timers,omitempty"`
	TraceOptions       *NetworkinstanceProtocolsOspfInstanceTraceOptions       `json:"trace-options,omitempty"`
	TrafficEngineering *NetworkinstanceProtocolsOspfInstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	Version            *string                                                 `json:"version"`
}

// NetworkinstanceProtocolsOspfInstanceArea struct
type NetworkinstanceProtocolsOspfInstanceArea struct {
	AdvertiseRouterCapability *bool `json:"advertise-router-capability,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	AreaId             *string                                              `json:"area-id"`
	AreaRange          []*NetworkinstanceProtocolsOspfInstanceAreaAreaRange `json:"area-range,omitempty"`
	BgpLsExclude       *bool                                                `json:"bgp-ls-exclude,omitempty"`
	BlackholeAggregate *bool                                                `json:"blackhole-aggregate,omitempty"`
	ExportPolicy       *string                                              `json:"export-policy,omitempty"`
	Interface          []*NetworkinstanceProtocolsOspfInstanceAreaInterface `json:"interface,omitempty"`
	Nssa               *NetworkinstanceProtocolsOspfInstanceAreaNssa        `json:"nssa,omitempty"`
	Stub               *NetworkinstanceProtocolsOspfInstanceAreaStub        `json:"stub,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaAreaRange struct
type NetworkinstanceProtocolsOspfInstanceAreaAreaRange struct {
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefixMask *string `json:"ip-prefix-mask"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterface struct
type NetworkinstanceProtocolsOspfInstanceAreaInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState                *string                                                          `json:"admin-state,omitempty"`
	AdvertiseRouterCapability *bool                                                            `json:"advertise-router-capability,omitempty"`
	AdvertiseSubnet           *bool                                                            `json:"advertise-subnet,omitempty"`
	Authentication            *NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication `json:"authentication,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=65535
	DeadInterval     *uint32                                                            `json:"dead-interval,omitempty"`
	FailureDetection *NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailureDetection `json:"failure-detection,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	HelloInterval *uint32 `json:"hello-interval,omitempty"`
	InterfaceName *string `json:"interface-name"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	InterfaceType *string `json:"interface-type,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`except-own-rtrlsa`;`except-own-rtrlsa-and-defaults`;`none`
	LsaFilterOut *string `json:"lsa-filter-out,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Metric *uint16 `json:"metric,omitempty"`
	// kubebuilder:validation:Minimum=512
	// kubebuilder:validation:Maximum=9486
	Mtu     *uint32 `json:"mtu,omitempty"`
	Passive *bool   `json:"passive,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Priority *uint16 `json:"priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	RetransmitInterval *uint32                                                        `json:"retransmit-interval,omitempty"`
	TraceOptions       *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptions `json:"trace-options,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	TransitDelay *uint32 `json:"transit-delay,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailureDetection struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailureDetection struct {
	EnableBfd *bool `json:"enable-bfd,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptions struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptions struct {
	Trace *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace struct {
	Adjacencies *string                                                                   `json:"adjacencies,omitempty"`
	Interfaces  *string                                                                   `json:"interfaces,omitempty"`
	Packet      *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket `json:"packet,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssa struct
type NetworkinstanceProtocolsOspfInstanceAreaNssa struct {
	AreaRange             []*NetworkinstanceProtocolsOspfInstanceAreaNssaAreaRange           `json:"area-range,omitempty"`
	OriginateDefaultRoute *NetworkinstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute `json:"originate-default-route,omitempty"`
	RedistributeExternal  *bool                                                              `json:"redistribute-external,omitempty"`
	Summaries             *bool                                                              `json:"summaries,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssaAreaRange struct
type NetworkinstanceProtocolsOspfInstanceAreaNssaAreaRange struct {
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefixMask *string `json:"ip-prefix-mask"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute struct
type NetworkinstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute struct {
	AdjacencyCheck *bool `json:"adjacency-check,omitempty"`
	TypeNssa       *bool `json:"type-nssa,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaStub struct
type NetworkinstanceProtocolsOspfInstanceAreaStub struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	DefaultMetric *uint16 `json:"default-metric,omitempty"`
	Summaries     *bool   `json:"summaries,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAsbr struct
type NetworkinstanceProtocolsOspfInstanceAsbr struct {
	TracePath *string `json:"trace-path,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceExportLimit struct
type NetworkinstanceProtocolsOspfInstanceExportLimit struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	LogPercent *uint32 `json:"log-percent,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Number *uint32 `json:"number"`
}

// NetworkinstanceProtocolsOspfInstanceExternalDbOverflow struct
type NetworkinstanceProtocolsOspfInstanceExternalDbOverflow struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=2147483647
	Interval *uint32 `json:"interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=2147483647
	Limit *uint32 `json:"limit,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceGracefulRestart struct
type NetworkinstanceProtocolsOspfInstanceGracefulRestart struct {
	HelperMode        *bool `json:"helper-mode,omitempty"`
	StrictLsaChecking *bool `json:"strict-lsa-checking,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverload struct
type NetworkinstanceProtocolsOspfInstanceOverload struct {
	Active                 *bool                                                       `json:"active,omitempty"`
	OverloadIncludeExt1    *bool                                                       `json:"overload-include-ext-1,omitempty"`
	OverloadIncludeExt2    *bool                                                       `json:"overload-include-ext-2,omitempty"`
	OverloadIncludeExtStub *bool                                                       `json:"overload-include-ext-stub,omitempty"`
	OverloadOnBoot         *NetworkinstanceProtocolsOspfInstanceOverloadOverloadOnBoot `json:"overload-on-boot,omitempty"`
	RtrAdvLsaLimit         *NetworkinstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit `json:"rtr-adv-lsa-limit,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverloadOverloadOnBoot struct
type NetworkinstanceProtocolsOspfInstanceOverloadOverloadOnBoot struct {
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=1800
	Timeout *uint32 `json:"timeout,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit struct
type NetworkinstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit struct {
	LogOnly *bool `json:"log-only,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	MaxLsaCount *uint32 `json:"max-lsa-count,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	OverloadTimeout *uint16 `json:"overload-timeout,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	WarningThreshold *uint8 `json:"warning-threshold,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTeDatabaseInstall struct
type NetworkinstanceProtocolsOspfInstanceTeDatabaseInstall struct {
	BgpLs *NetworkinstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs struct
type NetworkinstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	BgpLsIdentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	IgpIdentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimers struct
type NetworkinstanceProtocolsOspfInstanceTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	IncrementalSpfWait *uint32 `json:"incremental-spf-wait,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	LsaAccumulate *uint32 `json:"lsa-accumulate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=600000
	LsaArrival  *uint32                                                `json:"lsa-arrival,omitempty"`
	LsaGenerate *NetworkinstanceProtocolsOspfInstanceTimersLsaGenerate `json:"lsa-generate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=1000
	RedistributeDelay *uint32                                            `json:"redistribute-delay,omitempty"`
	SpfWait           *NetworkinstanceProtocolsOspfInstanceTimersSpfWait `json:"spf-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimersLsaGenerate struct
type NetworkinstanceProtocolsOspfInstanceTimersLsaGenerate struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	LsaInitialWait *uint32 `json:"lsa-initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	LsaSecondWait *uint32 `json:"lsa-second-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=600000
	MaxLsaWait *uint32 `json:"max-lsa-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimersSpfWait struct
type NetworkinstanceProtocolsOspfInstanceTimersSpfWait struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	SpfInitialWait *uint32 `json:"spf-initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	SpfMaxWait *uint32 `json:"spf-max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	SpfSecondWait *uint32 `json:"spf-second-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptions struct
type NetworkinstanceProtocolsOspfInstanceTraceOptions struct {
	Trace *NetworkinstanceProtocolsOspfInstanceTraceOptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTrace struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTrace struct {
	Adjacencies     *string                                                      `json:"adjacencies,omitempty"`
	GracefulRestart *string                                                      `json:"graceful-restart,omitempty"`
	Interfaces      *string                                                      `json:"interfaces,omitempty"`
	Lsdb            *NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceLsdb   `json:"lsdb,omitempty"`
	Misc            *string                                                      `json:"misc,omitempty"`
	Packet          *NetworkinstanceProtocolsOspfInstanceTraceOptionsTracePacket `json:"packet,omitempty"`
	Routes          *NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceRoutes `json:"routes,omitempty"`
	Spf             *NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceSpf    `json:"spf,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceLsdb struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceLsdb struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	LinkStateId *string `json:"link-state-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId *string `json:"router-id,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`external`;`inter-area-prefix`;`inter-area-router`;`intra-area-prefix`;`network`;`nssa`;`opaque`;`router`;`summary`
	Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTracePacket struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceRoutes struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceRoutes struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	DestAddress *string `json:"dest-address,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceSpf struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceSpf struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	DestAddress *string `json:"dest-address,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTrafficEngineering struct
type NetworkinstanceProtocolsOspfInstanceTrafficEngineering struct {
	Advertisement                    *bool `json:"advertisement,omitempty"`
	LegacyLinkAttributeAdvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// NetworkinstanceProtocolsOspfParameters are the parameter fields of a NetworkinstanceProtocolsOspf.
type NetworkinstanceProtocolsOspfParameters struct {
	NetworkInstanceName             *string                       `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsOspf *NetworkinstanceProtocolsOspf `json:"ospf,omitempty"`
}

// NetworkinstanceProtocolsOspfObservation are the observable fields of a NetworkinstanceProtocolsOspf.
type NetworkinstanceProtocolsOspfObservation struct {
}

// A NetworkinstanceProtocolsOspfSpec defines the desired state of a NetworkinstanceProtocolsOspf.
type NetworkinstanceProtocolsOspfSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceProtocolsOspfParameters `json:"forNetworkNode"`
}

// A NetworkinstanceProtocolsOspfStatus represents the observed state of a NetworkinstanceProtocolsOspf.
type NetworkinstanceProtocolsOspfStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceProtocolsOspfObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsOspf is the Schema for the NetworkinstanceProtocolsOspf API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceProtocolsOspf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsOspfSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsOspfStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsOspfList contains a list of NetworkinstanceProtocolsOspfs
type SrlNetworkinstanceProtocolsOspfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsOspf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsOspf{}, &SrlNetworkinstanceProtocolsOspfList{})
}

// NetworkinstanceProtocolsOspf type metadata.
var (
	NetworkinstanceProtocolsOspfKind             = reflect.TypeOf(SrlNetworkinstanceProtocolsOspf{}).Name()
	NetworkinstanceProtocolsOspfGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsOspfKind}.String()
	NetworkinstanceProtocolsOspfKindAPIVersion   = NetworkinstanceProtocolsOspfKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsOspfGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsOspfKind)
)
