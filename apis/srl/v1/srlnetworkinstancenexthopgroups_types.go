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
	// NetworkinstanceNexthopgroupsFinalizer is the name of the finalizer added to
	// NetworkinstanceNexthopgroups to block delete operations until the physical node can be
	// deprovisioned.
	NetworkinstanceNexthopgroupsFinalizer string = "nextHopGroups.srl.ndd.yndd.io"
)

// NetworkinstanceNexthopgroups struct
type NetworkinstanceNexthopgroups struct {
	Group []*NetworkinstanceNexthopgroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceNexthopgroupsGroup struct
type NetworkinstanceNexthopgroupsGroup struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                     `json:"admin-state,omitempty"`
	Blackhole  *NetworkinstanceNexthopgroupsGroupBlackhole `json:"blackhole,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name    *string                                     `json:"name,omitempty"`
	Nexthop []*NetworkinstanceNexthopgroupsGroupNexthop `json:"nexthop,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupBlackhole struct
type NetworkinstanceNexthopgroupsGroupBlackhole struct {
	// +kubebuilder:default:=false
	GenerateIcmp *bool `json:"generate-icmp,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthop struct
type NetworkinstanceNexthopgroupsGroupNexthop struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState       *string                                                   `json:"admin-state,omitempty"`
	FailureDetection *NetworkinstanceNexthopgroupsGroupNexthopFailureDetection `json:"failure-detection,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	IpAddress            *string `json:"ip-address,omitempty"`
	PushedMplsLabelStack *string `json:"pushed-mpls-label-stack,omitempty"`
	// +kubebuilder:default:=true
	Resolve *bool `json:"resolve,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthopFailureDetection struct
type NetworkinstanceNexthopgroupsGroupNexthopFailureDetection struct {
	EnableBfd *NetworkinstanceNexthopgroupsGroupNexthopFailureDetectionEnableBfd `json:"enable-bfd,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthopFailureDetectionEnableBfd struct
type NetworkinstanceNexthopgroupsGroupNexthopFailureDetectionEnableBfd struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16384
	LocalDiscriminator *uint32 `json:"local-discriminator,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=16384
	RemoteDiscriminator *uint32 `json:"remote-discriminator,omitempty"`
}

// NetworkinstanceNexthopgroupsSpec struct
type NetworkinstanceNexthopgroupsParameters struct {
	NetworkInstanceName             *string                       `json:"network-instance-name"`
	SrlNetworkinstanceNexthopgroups *NetworkinstanceNexthopgroups `json:"next-hop-groups,omitempty"`
}

// NetworkinstanceNexthopgroupsStatus struct
type NetworkinstanceNexthopgroupsObservation struct {
}

// A NetworkinstanceNexthopgroupsSpec defines the desired state of a NetworkinstanceNexthopgroups.
type NetworkinstanceNexthopgroupsSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     NetworkinstanceNexthopgroupsParameters `json:"forNetworkNode"`
}

// A NetworkinstanceNexthopgroupsStatus represents the observed state of a NetworkinstanceNexthopgroups.
type NetworkinstanceNexthopgroupsStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        NetworkinstanceNexthopgroupsObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceNexthopgroups is the Schema for the NetworkinstanceNexthopgroups API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlNetworkinstanceNexthopgroups struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceNexthopgroupsSpec   `json:"spec,omitempty"`
	Status NetworkinstanceNexthopgroupsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceNexthopgroupsList contains a list of NetworkinstanceNexthopgroupss
type SrlNetworkinstanceNexthopgroupsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceNexthopgroups `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceNexthopgroups{}, &SrlNetworkinstanceNexthopgroupsList{})
}

// NetworkinstanceNexthopgroups type metadata.
var (
	NetworkinstanceNexthopgroupsKind             = reflect.TypeOf(SrlNetworkinstanceNexthopgroups{}).Name()
	NetworkinstanceNexthopgroupsGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceNexthopgroupsKind}.String()
	NetworkinstanceNexthopgroupsKindAPIVersion   = NetworkinstanceNexthopgroupsKind + "." + GroupVersion.String()
	NetworkinstanceNexthopgroupsGroupVersionKind = GroupVersion.WithKind(NetworkinstanceNexthopgroupsKind)
)
