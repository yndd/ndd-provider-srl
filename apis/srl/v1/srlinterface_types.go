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
	// InterfaceFinalizer is the name of the finalizer added to
	// Interface to block delete operations until the physical node can be
	// deprovisioned.
	InterfaceFinalizer string = "interface.srl.ndd.yndd.io"
)

// Interface struct
type Interface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description  *string            `json:"description,omitempty"`
	Ethernet     *InterfaceEthernet `json:"ethernet,omitempty"`
	Lag          *InterfaceLag      `json:"lag,omitempty"`
	LoopbackMode *bool              `json:"loopback-mode,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	Mtu *uint16 `json:"mtu,omitempty"`
	// kubebuilder:validation:MinLength=3
	// kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0|mgmt0-standby|system0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))`
	Name        *string               `json:"name"`
	Qos         *InterfaceQos         `json:"qos,omitempty"`
	Sflow       *InterfaceSflow       `json:"sflow,omitempty"`
	Transceiver *InterfaceTransceiver `json:"transceiver,omitempty"`
	VlanTagging *bool                 `json:"vlan-tagging,omitempty"`
}

// InterfaceEthernet struct
type InterfaceEthernet struct {
	AggregateId   *string `json:"aggregate-id,omitempty"`
	AutoNegotiate *bool   `json:"auto-negotiate,omitempty"`
	// +kubebuilder:validation:Enum=`full`;`half`
	DuplexMode  *string                       `json:"duplex-mode,omitempty"`
	FlowControl *InterfaceEthernetFlowControl `json:"flow-control,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	LacpPortPriority *uint16 `json:"lacp-port-priority,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`1T`;`200G`;`25G`;`400G`;`40G`;`50G`
	PortSpeed *string `json:"port-speed,omitempty"`
}

// InterfaceEthernetFlowControl struct
type InterfaceEthernetFlowControl struct {
	Receive  *bool `json:"receive,omitempty"`
	Transmit *bool `json:"transmit,omitempty"`
}

// InterfaceLag struct
type InterfaceLag struct {
	Lacp *InterfaceLagLacp `json:"lacp,omitempty"`
	// +kubebuilder:validation:Enum=`static`
	LacpFallbackMode *string `json:"lacp-fallback-mode,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=3600
	LacpFallbackTimeout *uint16 `json:"lacp-fallback-timeout,omitempty"`
	// +kubebuilder:validation:Enum=`lacp`;`static`
	LagType *string `json:"lag-type,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`25G`;`400G`;`40G`
	MemberSpeed *string `json:"member-speed,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	MinLinks *uint16 `json:"min-links,omitempty"`
}

// InterfaceLagLacp struct
type InterfaceLagLacp struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	AdminKey *uint16 `json:"admin-key,omitempty"`
	// +kubebuilder:validation:Enum=`FAST`;`SLOW`
	Interval *string `json:"interval,omitempty"`
	// +kubebuilder:validation:Enum=`ACTIVE`;`PASSIVE`
	LacpMode *string `json:"lacp-mode,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	SystemIdMac *string `json:"system-id-mac,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	SystemPriority *uint16 `json:"system-priority,omitempty"`
}

// InterfaceQos struct
type InterfaceQos struct {
	Output *InterfaceQosOutput `json:"output,omitempty"`
}

// InterfaceQosOutput struct
type InterfaceQosOutput struct {
	MulticastQueue []*InterfaceQosOutputMulticastQueue `json:"multicast-queue,omitempty"`
	Scheduler      *InterfaceQosOutputScheduler        `json:"scheduler,omitempty"`
	UnicastQueue   []*InterfaceQosOutputUnicastQueue   `json:"unicast-queue,omitempty"`
}

// InterfaceQosOutputMulticastQueue struct
type InterfaceQosOutputMulticastQueue struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	QueueId    *uint8                                      `json:"queue-id"`
	Scheduling *InterfaceQosOutputMulticastQueueScheduling `json:"scheduling,omitempty"`
	Template   *string                                     `json:"template,omitempty"`
}

// InterfaceQosOutputMulticastQueueScheduling struct
type InterfaceQosOutputMulticastQueueScheduling struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	PeakRatePercent *uint8 `json:"peak-rate-percent,omitempty"`
}

// InterfaceQosOutputScheduler struct
type InterfaceQosOutputScheduler struct {
	Tier []*InterfaceQosOutputSchedulerTier `json:"tier,omitempty"`
}

// InterfaceQosOutputSchedulerTier struct
type InterfaceQosOutputSchedulerTier struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4
	Level *uint8                                 `json:"level"`
	Node  []*InterfaceQosOutputSchedulerTierNode `json:"node,omitempty"`
}

// InterfaceQosOutputSchedulerTierNode struct
type InterfaceQosOutputSchedulerTierNode struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=11
	NodeNumber     *uint8 `json:"node-number"`
	StrictPriority *bool  `json:"strict-priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=127
	Weight *uint8 `json:"weight,omitempty"`
}

// InterfaceQosOutputUnicastQueue struct
type InterfaceQosOutputUnicastQueue struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	QueueId     *uint8                                    `json:"queue-id"`
	Scheduling  *InterfaceQosOutputUnicastQueueScheduling `json:"scheduling,omitempty"`
	Template    *string                                   `json:"template,omitempty"`
	VoqTemplate *string                                   `json:"voq-template,omitempty"`
}

// InterfaceQosOutputUnicastQueueScheduling struct
type InterfaceQosOutputUnicastQueueScheduling struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	PeakRatePercent *uint8 `json:"peak-rate-percent,omitempty"`
	StrictPriority  *bool  `json:"strict-priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	Weight *uint8 `json:"weight,omitempty"`
}

// InterfaceSflow struct
type InterfaceSflow struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
}

// InterfaceTransceiver struct
type InterfaceTransceiver struct {
	DdmEvents *bool `json:"ddm-events,omitempty"`
	// +kubebuilder:validation:Enum=`base-r`;`disabled`;`rs-108`;`rs-528`;`rs-544`
	ForwardErrorCorrection *string `json:"forward-error-correction,omitempty"`
	TxLaser                *bool   `json:"tx-laser,omitempty"`
}

// InterfaceParameters are the parameter fields of a Interface.
type InterfaceParameters struct {
	SrlInterface *Interface `json:"interface,omitempty"`
}

// InterfaceObservation are the observable fields of a Interface.
type InterfaceObservation struct {
}

// A InterfaceSpec defines the desired state of a Interface.
type InterfaceSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     InterfaceParameters `json:"forNetworkNode"`
}

// A InterfaceStatus represents the observed state of a Interface.
type InterfaceStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        InterfaceObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlInterface is the Schema for the Interface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlInterfaceList contains a list of Interfaces
type SrlInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlInterface{}, &SrlInterfaceList{})
}

// Interface type metadata.
var (
	InterfaceKind             = reflect.TypeOf(SrlInterface{}).Name()
	InterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceKind}.String()
	InterfaceKindAPIVersion   = InterfaceKind + "." + GroupVersion.String()
	InterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceKind)
)
