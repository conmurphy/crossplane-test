/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type L3ExternalSubnetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type L3ExternalSubnetParameters struct {

	// +kubebuilder:validation:Optional
	Aggregate *string `json:"aggregate,omitempty" tf:"aggregate,omitempty"`

	// +kubebuilder:validation:Optional
	Annotation *string `json:"annotation,omitempty" tf:"annotation,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=L3ExternalNetworkInstanceProfile
	// +kubebuilder:validation:Optional
	ExternalNetworkInstanceProfileDn *string `json:"externalNetworkInstanceProfileDn,omitempty" tf:"external_network_instance_profile_dn,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalNetworkInstanceProfileDnRef *v1.Reference `json:"externalNetworkInstanceProfileDnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ExternalNetworkInstanceProfileDnSelector *v1.Selector `json:"externalNetworkInstanceProfileDnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	IP *string `json:"ip" tf:"ip,omitempty"`

	// +kubebuilder:validation:Optional
	NameAlias *string `json:"nameAlias,omitempty" tf:"name_alias,omitempty"`

	// +kubebuilder:validation:Optional
	RelationL3ExtRsSubnetToProfile []RelationL3ExtRsSubnetToProfileParameters `json:"relationL3ExtRsSubnetToProfile,omitempty" tf:"relation_l3ext_rs_subnet_to_profile,omitempty"`

	// +kubebuilder:validation:Optional
	RelationL3ExtRsSubnetToRtSumm *string `json:"relationL3ExtRsSubnetToRtSumm,omitempty" tf:"relation_l3ext_rs_subnet_to_rt_summ,omitempty"`

	// +kubebuilder:validation:Optional
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type RelationL3ExtRsSubnetToProfileObservation struct {
}

type RelationL3ExtRsSubnetToProfileParameters struct {

	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// +kubebuilder:validation:Optional
	TnRtctrlProfileDn *string `json:"tnRtctrlProfileDn,omitempty" tf:"tn_rtctrl_profile_dn,omitempty"`

	// +kubebuilder:validation:Optional
	TnRtctrlProfileName *string `json:"tnRtctrlProfileName,omitempty" tf:"tn_rtctrl_profile_name,omitempty"`
}

// L3ExternalSubnetSpec defines the desired state of L3ExternalSubnet
type L3ExternalSubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     L3ExternalSubnetParameters `json:"forProvider"`
}

// L3ExternalSubnetStatus defines the observed state of L3ExternalSubnet.
type L3ExternalSubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        L3ExternalSubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// L3ExternalSubnet is the Schema for the L3ExternalSubnets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,acijet}
type L3ExternalSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              L3ExternalSubnetSpec   `json:"spec"`
	Status            L3ExternalSubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// L3ExternalSubnetList contains a list of L3ExternalSubnets
type L3ExternalSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L3ExternalSubnet `json:"items"`
}

// Repository type metadata.
var (
	L3ExternalSubnet_Kind             = "L3ExternalSubnet"
	L3ExternalSubnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: L3ExternalSubnet_Kind}.String()
	L3ExternalSubnet_KindAPIVersion   = L3ExternalSubnet_Kind + "." + CRDGroupVersion.String()
	L3ExternalSubnet_GroupVersionKind = CRDGroupVersion.WithKind(L3ExternalSubnet_Kind)
)

func init() {
	SchemeBuilder.Register(&L3ExternalSubnet{}, &L3ExternalSubnetList{})
}
