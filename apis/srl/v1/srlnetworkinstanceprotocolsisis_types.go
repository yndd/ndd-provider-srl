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
	// NetworkinstanceProtocolsIsisFinalizer is the name of the finalizer added to
	// NetworkinstanceProtocolsIsis to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceProtocolsIsisFinalizer string = "isis.srl.ndd.yndd.io"
)

// NetworkinstanceProtocolsIsis struct
type NetworkinstanceProtocolsIsis struct {
	Instance []*NetworkinstanceProtocolsIsisInstance `json:"instance,omitempty"`
}

// NetworkinstanceProtocolsIsisInstance struct
type NetworkinstanceProtocolsIsisInstance struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState                    *string                                                            `json:"admin-state,omitempty"`
	AttachedBit                   *NetworkinstanceProtocolsIsisInstanceAttachedBit                   `json:"attached-bit,omitempty"`
	Authentication                *NetworkinstanceProtocolsIsisInstanceAuthentication                `json:"authentication,omitempty"`
	AutoCost                      *NetworkinstanceProtocolsIsisInstanceAutoCost                      `json:"auto-cost,omitempty"`
	ExportPolicy                  *string                                                            `json:"export-policy,omitempty"`
	GracefulRestart               *NetworkinstanceProtocolsIsisInstanceGracefulRestart               `json:"graceful-restart,omitempty"`
	InterLevelPropagationPolicies *NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPolicies `json:"inter-level-propagation-policies,omitempty"`
	Interface                     []*NetworkinstanceProtocolsIsisInstanceInterface                   `json:"interface,omitempty"`
	Ipv4Unicast                   *NetworkinstanceProtocolsIsisInstanceIpv4Unicast                   `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast                   *NetworkinstanceProtocolsIsisInstanceIpv6Unicast                   `json:"ipv6-unicast,omitempty"`
	LdpSynchronization            *NetworkinstanceProtocolsIsisInstanceLdpSynchronization            `json:"ldp-synchronization,omitempty"`
	Level                         []*NetworkinstanceProtocolsIsisInstanceLevel                       `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`L1`;`L1L2`;`L2`
	LevelCapability *string `json:"level-capability,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MaxEcmpPaths *uint8 `json:"max-ecmp-paths,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[a-fA-F0-9]{2}(\.[a-fA-F0-9]{4}){3,9}\.[0]{2}`
	Net                *string                                                 `json:"net,omitempty"`
	Overload           *NetworkinstanceProtocolsIsisInstanceOverload           `json:"overload,omitempty"`
	PoiTlv             *bool                                                   `json:"poi-tlv,omitempty"`
	TeDatabaseInstall  *NetworkinstanceProtocolsIsisInstanceTeDatabaseInstall  `json:"te-database-install,omitempty"`
	Timers             *NetworkinstanceProtocolsIsisInstanceTimers             `json:"timers,omitempty"`
	TraceOptions       *NetworkinstanceProtocolsIsisInstanceTraceOptions       `json:"trace-options,omitempty"`
	TrafficEngineering *NetworkinstanceProtocolsIsisInstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	Transport          *NetworkinstanceProtocolsIsisInstanceTransport          `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAttachedBit struct
type NetworkinstanceProtocolsIsisInstanceAttachedBit struct {
	Ignore   *bool `json:"ignore,omitempty"`
	Suppress *bool `json:"suppress,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAuthentication struct
type NetworkinstanceProtocolsIsisInstanceAuthentication struct {
	CsnpAuthentication  *bool   `json:"csnp-authentication,omitempty"`
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
	PsnpAuthentication  *bool   `json:"psnp-authentication,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAutoCost struct
type NetworkinstanceProtocolsIsisInstanceAutoCost struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=8000000000
	ReferenceBandwidth *uint64 `json:"reference-bandwidth,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceGracefulRestart struct
type NetworkinstanceProtocolsIsisInstanceGracefulRestart struct {
	HelperMode *bool `json:"helper-mode,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPolicies struct
type NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPolicies struct {
	Level1ToLevel2 *NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 `json:"level1-to-level2,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 struct
type NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 struct {
	SummaryAddress []*NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress `json:"summary-address,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress struct
type NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	RouteTag *uint32 `json:"route-tag,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterface struct
type NetworkinstanceProtocolsIsisInstanceInterface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState     *string                                                      `json:"admin-state,omitempty"`
	Authentication *NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	CircuitType *string `json:"circuit-type,omitempty"`
	// +kubebuilder:validation:Enum=`adaptive`;`disable`;`loose`;`strict`
	HelloPadding       *string                                                          `json:"hello-padding,omitempty"`
	InterfaceName      *string                                                          `json:"interface-name"`
	Ipv4Unicast        *NetworkinstanceProtocolsIsisInstanceInterfaceIpv4Unicast        `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast        *NetworkinstanceProtocolsIsisInstanceInterfaceIpv6Unicast        `json:"ipv6-unicast,omitempty"`
	LdpSynchronization *NetworkinstanceProtocolsIsisInstanceInterfaceLdpSynchronization `json:"ldp-synchronization,omitempty"`
	Level              []*NetworkinstanceProtocolsIsisInstanceInterfaceLevel            `json:"level,omitempty"`
	Passive            *bool                                                            `json:"passive,omitempty"`
	Timers             *NetworkinstanceProtocolsIsisInstanceInterfaceTimers             `json:"timers,omitempty"`
	TraceOptions       *NetworkinstanceProtocolsIsisInstanceInterfaceTraceOptions       `json:"trace-options,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct
type NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct {
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceIpv4Unicast struct
type NetworkinstanceProtocolsIsisInstanceInterfaceIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState    *string `json:"admin-state,omitempty"`
	EnableBfd     *bool   `json:"enable-bfd,omitempty"`
	IncludeBfdTlv *bool   `json:"include-bfd-tlv,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceIpv6Unicast struct
type NetworkinstanceProtocolsIsisInstanceInterfaceIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState    *string `json:"admin-state,omitempty"`
	EnableBfd     *bool   `json:"enable-bfd,omitempty"`
	IncludeBfdTlv *bool   `json:"include-bfd-tlv,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLdpSynchronization struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLdpSynchronization struct {
	Disable  *string `json:"disable,omitempty"`
	EndOfLib *bool   `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	HoldDownTimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevel struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevel struct {
	Authentication *NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication `json:"authentication,omitempty"`
	Disable        *bool                                                             `json:"disable,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Ipv6UnicastMetric *uint32 `json:"ipv6-unicast-metric,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	LevelNumber *uint8 `json:"level-number"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16777215
	Metric *uint32 `json:"metric,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=127
	Priority *uint8                                                    `json:"priority,omitempty"`
	Timers   *NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers `json:"timers,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=20000
	HelloInterval *uint32 `json:"hello-interval,omitempty"`
	// kubebuilder:validation:Minimum=2
	// kubebuilder:validation:Maximum=100
	HelloMultiplier *uint8 `json:"hello-multiplier,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceTimers struct
type NetworkinstanceProtocolsIsisInstanceInterfaceTimers struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	CsnpInterval *uint16 `json:"csnp-interval,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000
	LspPacingInterval *uint64 `json:"lsp-pacing-interval,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceTraceOptions struct
type NetworkinstanceProtocolsIsisInstanceInterfaceTraceOptions struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`
	Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceIpv4Unicast struct
type NetworkinstanceProtocolsIsisInstanceIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceIpv6Unicast struct
type NetworkinstanceProtocolsIsisInstanceIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState    *string `json:"admin-state,omitempty"`
	MultiTopology *bool   `json:"multi-topology,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLdpSynchronization struct
type NetworkinstanceProtocolsIsisInstanceLdpSynchronization struct {
	EndOfLib *bool `json:"end-of-lib,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=1800
	HoldDownTimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevel struct
type NetworkinstanceProtocolsIsisInstanceLevel struct {
	Authentication *NetworkinstanceProtocolsIsisInstanceLevelAuthentication `json:"authentication,omitempty"`
	BgpLsExclude   *bool                                                    `json:"bgp-ls-exclude,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	LevelNumber *uint8 `json:"level-number"`
	// +kubebuilder:validation:Enum=`narrow`;`wide`
	MetricStyle     *string                                                   `json:"metric-style,omitempty"`
	RoutePreference *NetworkinstanceProtocolsIsisInstanceLevelRoutePreference `json:"route-preference,omitempty"`
	TraceOptions    *NetworkinstanceProtocolsIsisInstanceLevelTraceOptions    `json:"trace-options,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelAuthentication struct
type NetworkinstanceProtocolsIsisInstanceLevelAuthentication struct {
	CsnpAuthentication  *bool   `json:"csnp-authentication,omitempty"`
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
	PsnpAuthentication  *bool   `json:"psnp-authentication,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelRoutePreference struct
type NetworkinstanceProtocolsIsisInstanceLevelRoutePreference struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	External *uint8 `json:"external,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Internal *uint8 `json:"internal,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelTraceOptions struct
type NetworkinstanceProtocolsIsisInstanceLevelTraceOptions struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`lsdb`;`routes`;`spf`
	Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverload struct
type NetworkinstanceProtocolsIsisInstanceOverload struct {
	AdvertiseExternal   *bool                                                  `json:"advertise-external,omitempty"`
	AdvertiseInterlevel *bool                                                  `json:"advertise-interlevel,omitempty"`
	Immediate           *NetworkinstanceProtocolsIsisInstanceOverloadImmediate `json:"immediate,omitempty"`
	OnBoot              *NetworkinstanceProtocolsIsisInstanceOverloadOnBoot    `json:"on-boot,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverloadImmediate struct
type NetworkinstanceProtocolsIsisInstanceOverloadImmediate struct {
	MaxMetric *bool `json:"max-metric,omitempty"`
	SetBit    *bool `json:"set-bit,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverloadOnBoot struct
type NetworkinstanceProtocolsIsisInstanceOverloadOnBoot struct {
	MaxMetric *bool `json:"max-metric,omitempty"`
	SetBit    *bool `json:"set-bit,omitempty"`
	// kubebuilder:validation:Minimum=60
	// kubebuilder:validation:Maximum=1800
	Timeout *uint16 `json:"timeout,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTeDatabaseInstall struct
type NetworkinstanceProtocolsIsisInstanceTeDatabaseInstall struct {
	BgpLs *NetworkinstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs struct
type NetworkinstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	BgpLsIdentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=-1
	IgpIdentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimers struct
type NetworkinstanceProtocolsIsisInstanceTimers struct {
	LspGeneration *NetworkinstanceProtocolsIsisInstanceTimersLspGeneration `json:"lsp-generation,omitempty"`
	// kubebuilder:validation:Minimum=350
	// kubebuilder:validation:Maximum=65535
	LspLifetime *uint16                                               `json:"lsp-lifetime,omitempty"`
	LspRefresh  *NetworkinstanceProtocolsIsisInstanceTimersLspRefresh `json:"lsp-refresh,omitempty"`
	Spf         *NetworkinstanceProtocolsIsisInstanceTimersSpf        `json:"spf,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersLspGeneration struct
type NetworkinstanceProtocolsIsisInstanceTimersLspGeneration struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	InitialWait *uint64 `json:"initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	MaxWait *uint64 `json:"max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	SecondWait *uint64 `json:"second-wait,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersLspRefresh struct
type NetworkinstanceProtocolsIsisInstanceTimersLspRefresh struct {
	HalfLifetime *bool `json:"half-lifetime,omitempty"`
	// kubebuilder:validation:Minimum=150
	// kubebuilder:validation:Maximum=65535
	Interval *uint16 `json:"interval,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersSpf struct
type NetworkinstanceProtocolsIsisInstanceTimersSpf struct {
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	InitialWait *uint64 `json:"initial-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=120000
	MaxWait *uint64 `json:"max-wait,omitempty"`
	// kubebuilder:validation:Minimum=10
	// kubebuilder:validation:Maximum=100000
	SecondWait *uint64 `json:"second-wait,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTraceOptions struct
type NetworkinstanceProtocolsIsisInstanceTraceOptions struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`graceful-restart`;`interfaces`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`;`routes`;`summary-addresses`
	Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTrafficEngineering struct
type NetworkinstanceProtocolsIsisInstanceTrafficEngineering struct {
	Advertisement                    *bool `json:"advertisement,omitempty"`
	LegacyLinkAttributeAdvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTransport struct
type NetworkinstanceProtocolsIsisInstanceTransport struct {
	// kubebuilder:validation:Minimum=490
	// kubebuilder:validation:Maximum=9490
	LspMtuSize *uint16 `json:"lsp-mtu-size,omitempty"`
}

// NetworkinstanceProtocolsIsisParameters are the parameter fields of a NetworkinstanceProtocolsIsis.
type NetworkinstanceProtocolsIsisParameters struct {
	NetworkInstanceName             *string                       `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsIsis *NetworkinstanceProtocolsIsis `json:"isis,omitempty"`
}

// NetworkinstanceProtocolsIsisObservation are the observable fields of a NetworkinstanceProtocolsIsis.
type NetworkinstanceProtocolsIsisObservation struct {
}

// A NetworkinstanceProtocolsIsisSpec defines the desired state of a NetworkinstanceProtocolsIsis.
type NetworkinstanceProtocolsIsisSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceProtocolsIsisParameters `json:"forNetworkNode"`
}

// A NetworkinstanceProtocolsIsisStatus represents the observed state of a NetworkinstanceProtocolsIsis.
type NetworkinstanceProtocolsIsisStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceProtocolsIsisObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsIsis is the Schema for the NetworkinstanceProtocolsIsis API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceProtocolsIsis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsIsisSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsIsisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsIsisList contains a list of NetworkinstanceProtocolsIsiss
type SrlNetworkinstanceProtocolsIsisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsIsis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsIsis{}, &SrlNetworkinstanceProtocolsIsisList{})
}

// NetworkinstanceProtocolsIsis type metadata.
var (
	NetworkinstanceProtocolsIsisKind             = reflect.TypeOf(SrlNetworkinstanceProtocolsIsis{}).Name()
	NetworkinstanceProtocolsIsisGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsIsisKind}.String()
	NetworkinstanceProtocolsIsisKindAPIVersion   = NetworkinstanceProtocolsIsisKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsIsisGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsIsisKind)
)
