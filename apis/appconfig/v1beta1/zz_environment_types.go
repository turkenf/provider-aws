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

type EnvironmentInitParameters struct {

	// Description of the environment. Can be at most 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitor []MonitorInitParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// Name for the environment. Must be between 1 and 64 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentObservation struct {

	// AppConfig application ID. Must be between 4 and 7 characters in length.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// ARN of the AppConfig Environment.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the environment. Can be at most 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// AppConfig environment ID.
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// (Deprecated) AppConfig environment ID and application ID separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitor []MonitorObservation `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// Name for the environment. Must be between 1 and 64 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// State of the environment. Possible values are READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK
	// or ROLLED_BACK.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EnvironmentParameters struct {

	// AppConfig application ID. Must be between 4 and 7 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Description of the environment. Can be at most 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	// +kubebuilder:validation:Optional
	Monitor []MonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// Name for the environment. Must be between 1 and 64 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorInitParameters struct {
}

type MonitorObservation struct {

	// ARN of the Amazon CloudWatch alarm.
	AlarmArn *string `json:"alarmArn,omitempty" tf:"alarm_arn,omitempty"`

	// ARN of an IAM role for AWS AppConfig to monitor alarm_arn.
	AlarmRoleArn *string `json:"alarmRoleArn,omitempty" tf:"alarm_role_arn,omitempty"`
}

type MonitorParameters struct {

	// ARN of the Amazon CloudWatch alarm.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatch/v1beta1.MetricAlarm
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	AlarmArn *string `json:"alarmArn,omitempty" tf:"alarm_arn,omitempty"`

	// Reference to a MetricAlarm in cloudwatch to populate alarmArn.
	// +kubebuilder:validation:Optional
	AlarmArnRef *v1.Reference `json:"alarmArnRef,omitempty" tf:"-"`

	// Selector for a MetricAlarm in cloudwatch to populate alarmArn.
	// +kubebuilder:validation:Optional
	AlarmArnSelector *v1.Selector `json:"alarmArnSelector,omitempty" tf:"-"`

	// ARN of an IAM role for AWS AppConfig to monitor alarm_arn.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	AlarmRoleArn *string `json:"alarmRoleArn,omitempty" tf:"alarm_role_arn,omitempty"`

	// Reference to a Role in iam to populate alarmRoleArn.
	// +kubebuilder:validation:Optional
	AlarmRoleArnRef *v1.Reference `json:"alarmRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate alarmRoleArn.
	// +kubebuilder:validation:Optional
	AlarmRoleArnSelector *v1.Selector `json:"alarmRoleArnSelector,omitempty" tf:"-"`
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider EnvironmentInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Environment is the Schema for the Environments API. Provides an AppConfig Environment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   EnvironmentSpec   `json:"spec"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	Environment_Kind             = "Environment"
	Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Environment_Kind}.String()
	Environment_KindAPIVersion   = Environment_Kind + "." + CRDGroupVersion.String()
	Environment_GroupVersionKind = CRDGroupVersion.WithKind(Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
