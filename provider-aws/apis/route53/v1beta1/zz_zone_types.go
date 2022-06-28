/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPCObservation struct {
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	VPCRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`
}

type VPCParameters struct {
}

type ZoneObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NameServers []*string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VPC []VPCObservation `json:"vpc,omitempty" tf:"vpc,omitempty"`

	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ZoneParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +crossplane:generate:reference:type=DelegationSet
	// +kubebuilder:validation:Optional
	DelegationSetID *string `json:"delegationSetId,omitempty" tf:"delegation_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	DelegationSetIDRef *v1.Reference `json:"delegationSetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DelegationSetIDSelector *v1.Selector `json:"delegationSetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Zone is the Schema for the Zones API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneSpec   `json:"spec"`
	Status            ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
