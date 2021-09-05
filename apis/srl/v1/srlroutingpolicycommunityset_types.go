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
	// RoutingpolicyCommunitysetFinalizer is the name of the finalizer added to
	// RoutingpolicyCommunityset to block delete operations until the physical node can be
	// deprovisioned.
	RoutingpolicyCommunitysetFinalizer string = "communitySet.srl.ndd.yndd.io"
)

// RoutingpolicyCommunityset struct
type RoutingpolicyCommunityset struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
	Member *string `json:"member,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name,omitempty"`
}

// RoutingpolicyCommunitysetSpec struct
type RoutingpolicyCommunitysetParameters struct {
	SrlRoutingpolicyCommunityset *RoutingpolicyCommunityset `json:"community-set,omitempty"`
}

// RoutingpolicyCommunitysetStatus struct
type RoutingpolicyCommunitysetObservation struct {
}

// A RoutingpolicyCommunitysetSpec defines the desired state of a RoutingpolicyCommunityset.
type RoutingpolicyCommunitysetSpec struct {
	nddv1.ResourceSpec `json:",inline"`
	ForNetworkNode     RoutingpolicyCommunitysetParameters `json:"forNetworkNode"`
}

// A RoutingpolicyCommunitysetStatus represents the observed state of a RoutingpolicyCommunityset.
type RoutingpolicyCommunitysetStatus struct {
	nddv1.ResourceStatus `json:",inline"`
	AtNetworkNode        RoutingpolicyCommunitysetObservation `json:"atNetworkNode,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyCommunityset is the Schema for the RoutingpolicyCommunityset API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="TARGET",type="string",JSONPath=".status.conditions[?(@.kind=='TargetFound')].status"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCALLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='InternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="EXTLEAFREF",type="string",JSONPath=".status.conditions[?(@.kind=='ExternalLeafrefValidationSuccess')].status"
// +kubebuilder:printcolumn:name="PARENTDEP",type="string",JSONPath=".status.conditions[?(@.kind=='ParentValidationSuccess')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={ndd,srl}
type SrlRoutingpolicyCommunityset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoutingpolicyCommunitysetSpec   `json:"spec,omitempty"`
	Status RoutingpolicyCommunitysetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyCommunitysetList contains a list of RoutingpolicyCommunitysets
type SrlRoutingpolicyCommunitysetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyCommunityset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyCommunityset{}, &SrlRoutingpolicyCommunitysetList{})
}

// RoutingpolicyCommunityset type metadata.
var (
	RoutingpolicyCommunitysetKind             = reflect.TypeOf(SrlRoutingpolicyCommunityset{}).Name()
	RoutingpolicyCommunitysetGroupKind        = schema.GroupKind{Group: Group, Kind: RoutingpolicyCommunitysetKind}.String()
	RoutingpolicyCommunitysetKindAPIVersion   = RoutingpolicyCommunitysetKind + "." + GroupVersion.String()
	RoutingpolicyCommunitysetGroupVersionKind = GroupVersion.WithKind(RoutingpolicyCommunitysetKind)
)
