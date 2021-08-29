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
	// RoutingpolicyPolicyFinalizer is the name of the finalizer added to
	// RoutingpolicyPolicy to block delete operations until the physical node can be
	// deprovisioned.
	RoutingpolicyPolicyFinalizer string = "policy.srl.ndd.yndd.io"
)

// RoutingpolicyPolicy struct
type RoutingpolicyPolicy struct {
	DefaultAction *RoutingpolicyPolicyDefaultAction `json:"default-action,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name      *string                         `json:"name,omitempty"`
	Statement []*RoutingpolicyPolicyStatement `json:"statement,omitempty"`
}

// RoutingpolicyPolicyDefaultAction struct
type RoutingpolicyPolicyDefaultAction struct {
	Accept     *RoutingpolicyPolicyDefaultActionAccept     `json:"accept,omitempty"`
	NextEntry  *RoutingpolicyPolicyDefaultActionNextEntry  `json:"next-entry,omitempty"`
	NextPolicy *RoutingpolicyPolicyDefaultActionNextPolicy `json:"next-policy,omitempty"`
	Reject     *RoutingpolicyPolicyDefaultActionReject     `json:"reject,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAccept struct
type RoutingpolicyPolicyDefaultActionAccept struct {
	Bgp *RoutingpolicyPolicyDefaultActionAcceptBgp `json:"bgp,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgp struct
type RoutingpolicyPolicyDefaultActionAcceptBgp struct {
	AsPath          *RoutingpolicyPolicyDefaultActionAcceptBgpAsPath          `json:"as-path,omitempty"`
	Communities     *RoutingpolicyPolicyDefaultActionAcceptBgpCommunities     `json:"communities,omitempty"`
	LocalPreference *RoutingpolicyPolicyDefaultActionAcceptBgpLocalPreference `json:"local-preference,omitempty"`
	Origin          *RoutingpolicyPolicyDefaultActionAcceptBgpOrigin          `json:"origin,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpAsPath struct
type RoutingpolicyPolicyDefaultActionAcceptBgpAsPath struct {
	Prepend *RoutingpolicyPolicyDefaultActionAcceptBgpAsPathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                   `json:"remove,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpAsPathPrepend struct
type RoutingpolicyPolicyDefaultActionAcceptBgpAsPathPrepend struct {
	AsNumber *string `json:"as-number,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	RepeatN *uint8 `json:"repeat-n,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpCommunities struct
type RoutingpolicyPolicyDefaultActionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpLocalPreference struct
type RoutingpolicyPolicyDefaultActionAcceptBgpLocalPreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpOrigin struct
type RoutingpolicyPolicyDefaultActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set *string `json:"set,omitempty"`
}

// RoutingpolicyPolicyDefaultActionNextEntry struct
type RoutingpolicyPolicyDefaultActionNextEntry struct {
}

// RoutingpolicyPolicyDefaultActionNextPolicy struct
type RoutingpolicyPolicyDefaultActionNextPolicy struct {
}

// RoutingpolicyPolicyDefaultActionReject struct
type RoutingpolicyPolicyDefaultActionReject struct {
}

// RoutingpolicyPolicyStatement struct
type RoutingpolicyPolicyStatement struct {
	Action *RoutingpolicyPolicyStatementAction `json:"action,omitempty"`
	Match  *RoutingpolicyPolicyStatementMatch  `json:"match,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	SequenceId *uint32 `json:"sequence-id,omitempty"`
}

// RoutingpolicyPolicyStatementAction struct
type RoutingpolicyPolicyStatementAction struct {
	Accept     *RoutingpolicyPolicyStatementActionAccept     `json:"accept,omitempty"`
	NextEntry  *RoutingpolicyPolicyStatementActionNextEntry  `json:"next-entry,omitempty"`
	NextPolicy *RoutingpolicyPolicyStatementActionNextPolicy `json:"next-policy,omitempty"`
	Reject     *RoutingpolicyPolicyStatementActionReject     `json:"reject,omitempty"`
}

// RoutingpolicyPolicyStatementActionAccept struct
type RoutingpolicyPolicyStatementActionAccept struct {
	Bgp *RoutingpolicyPolicyStatementActionAcceptBgp `json:"bgp,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgp struct
type RoutingpolicyPolicyStatementActionAcceptBgp struct {
	AsPath          *RoutingpolicyPolicyStatementActionAcceptBgpAsPath          `json:"as-path,omitempty"`
	Communities     *RoutingpolicyPolicyStatementActionAcceptBgpCommunities     `json:"communities,omitempty"`
	LocalPreference *RoutingpolicyPolicyStatementActionAcceptBgpLocalPreference `json:"local-preference,omitempty"`
	Origin          *RoutingpolicyPolicyStatementActionAcceptBgpOrigin          `json:"origin,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpAsPath struct
type RoutingpolicyPolicyStatementActionAcceptBgpAsPath struct {
	Prepend *RoutingpolicyPolicyStatementActionAcceptBgpAsPathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                     `json:"remove,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpAsPathPrepend struct
type RoutingpolicyPolicyStatementActionAcceptBgpAsPathPrepend struct {
	AsNumber *string `json:"as-number,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	RepeatN *uint8 `json:"repeat-n,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpCommunities struct
type RoutingpolicyPolicyStatementActionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpLocalPreference struct
type RoutingpolicyPolicyStatementActionAcceptBgpLocalPreference struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpOrigin struct
type RoutingpolicyPolicyStatementActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set *string `json:"set,omitempty"`
}

// RoutingpolicyPolicyStatementActionNextEntry struct
type RoutingpolicyPolicyStatementActionNextEntry struct {
}

// RoutingpolicyPolicyStatementActionNextPolicy struct
type RoutingpolicyPolicyStatementActionNextPolicy struct {
}

// RoutingpolicyPolicyStatementActionReject struct
type RoutingpolicyPolicyStatementActionReject struct {
}

// RoutingpolicyPolicyStatementMatch struct
type RoutingpolicyPolicyStatementMatch struct {
	Bgp       *RoutingpolicyPolicyStatementMatchBgp  `json:"bgp,omitempty"`
	Family    *string                                `json:"family,omitempty"`
	Isis      *RoutingpolicyPolicyStatementMatchIsis `json:"isis,omitempty"`
	Ospf      *RoutingpolicyPolicyStatementMatchOspf `json:"ospf,omitempty"`
	PrefixSet *string                                `json:"prefix-set,omitempty"`
	Protocol  *string                                `json:"protocol,omitempty"`
}

// RoutingpolicyPolicyStatementMatchBgp struct
type RoutingpolicyPolicyStatementMatchBgp struct {
	AsPathLength *RoutingpolicyPolicyStatementMatchBgpAsPathLength `json:"as-path-length,omitempty"`
	AsPathSet    *string                                           `json:"as-path-set,omitempty"`
	CommunitySet *string                                           `json:"community-set,omitempty"`
	Evpn         *RoutingpolicyPolicyStatementMatchBgpEvpn         `json:"evpn,omitempty"`
}

// RoutingpolicyPolicyStatementMatchBgpAsPathLength struct
type RoutingpolicyPolicyStatementMatchBgpAsPathLength struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	// +kubebuilder:default:=eq
	Operator *string `json:"operator,omitempty"`
	// +kubebuilder:default:=false
	Unique *bool `json:"unique,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	Value *uint8 `json:"value"`
}

// RoutingpolicyPolicyStatementMatchBgpEvpn struct
type RoutingpolicyPolicyStatementMatchBgpEvpn struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=5
	RouteType *uint8 `json:"route-type,omitempty"`
}

// RoutingpolicyPolicyStatementMatchIsis struct
type RoutingpolicyPolicyStatementMatchIsis struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=2
	Level *uint8 `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`external`;`internal`
	RouteType *string `json:"route-type,omitempty"`
}

// RoutingpolicyPolicyStatementMatchOspf struct
type RoutingpolicyPolicyStatementMatchOspf struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	AreaId *string `json:"area-id,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=255
	InstanceId *uint32 `json:"instance-id,omitempty"`
	RouteType  *string `json:"route-type,omitempty"`
}

// RoutingpolicyPolicySpec struct
type RoutingpolicyPolicyParameters struct {
	SrlRoutingpolicyPolicy *RoutingpolicyPolicy `json:"policy,omitempty"`
}

// RoutingpolicyPolicyStatus struct
type RoutingpolicyPolicyObservation struct {
}

// A RoutingpolicyPolicySpec defines the desired state of a RoutingpolicyPolicy.
type RoutingpolicyPolicySpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     RoutingpolicyPolicyParameters `json:"forNetworkNode"`
}

// A RoutingpolicyPolicyStatus represents the observed state of a RoutingpolicyPolicy.
type RoutingpolicyPolicyStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        RoutingpolicyPolicyObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPolicy is the Schema for the RoutingpolicyPolicy API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlRoutingpolicyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoutingpolicyPolicySpec   `json:"spec,omitempty"`
	Status RoutingpolicyPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPolicyList contains a list of RoutingpolicyPolicys
type SrlRoutingpolicyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyPolicy{}, &SrlRoutingpolicyPolicyList{})
}

// RoutingpolicyPolicy type metadata.
var (
	RoutingpolicyPolicyKind             = reflect.TypeOf(SrlRoutingpolicyPolicy{}).Name()
	RoutingpolicyPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: RoutingpolicyPolicyKind}.String()
	RoutingpolicyPolicyKindAPIVersion   = RoutingpolicyPolicyKind + "." + GroupVersion.String()
	RoutingpolicyPolicyGroupVersionKind = GroupVersion.WithKind(RoutingpolicyPolicyKind)
)
