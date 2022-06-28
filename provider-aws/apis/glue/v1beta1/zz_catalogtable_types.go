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

type CatalogTableObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PartitionIndex []PartitionIndexObservation `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`
}

type CatalogTableParameters struct {

	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionIndex []PartitionIndexParameters `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionKeys []PartitionKeysParameters `json:"partitionKeys,omitempty" tf:"partition_keys,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// +kubebuilder:validation:Optional
	StorageDescriptor []StorageDescriptorParameters `json:"storageDescriptor,omitempty" tf:"storage_descriptor,omitempty"`

	// +kubebuilder:validation:Optional
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// +kubebuilder:validation:Optional
	TargetTable []TargetTableParameters `json:"targetTable,omitempty" tf:"target_table,omitempty"`

	// +kubebuilder:validation:Optional
	ViewExpandedText *string `json:"viewExpandedText,omitempty" tf:"view_expanded_text,omitempty"`

	// +kubebuilder:validation:Optional
	ViewOriginalText *string `json:"viewOriginalText,omitempty" tf:"view_original_text,omitempty"`
}

type ColumnsObservation struct {
}

type ColumnsParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PartitionIndexObservation struct {
	IndexStatus *string `json:"indexStatus,omitempty" tf:"index_status,omitempty"`
}

type PartitionIndexParameters struct {

	// +kubebuilder:validation:Required
	IndexName *string `json:"indexName" tf:"index_name,omitempty"`

	// +kubebuilder:validation:Required
	Keys []*string `json:"keys" tf:"keys,omitempty"`
}

type PartitionKeysObservation struct {
}

type PartitionKeysParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchemaIDObservation struct {
}

type SchemaIDParameters struct {

	// +kubebuilder:validation:Optional
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaArn *string `json:"schemaArn,omitempty" tf:"schema_arn,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`
}

type SchemaReferenceObservation struct {
}

type SchemaReferenceParameters struct {

	// +kubebuilder:validation:Optional
	SchemaID []SchemaIDParameters `json:"schemaId,omitempty" tf:"schema_id,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaVersionID *string `json:"schemaVersionId,omitempty" tf:"schema_version_id,omitempty"`

	// +kubebuilder:validation:Required
	SchemaVersionNumber *float64 `json:"schemaVersionNumber" tf:"schema_version_number,omitempty"`
}

type SerDeInfoObservation struct {
}

type SerDeInfoParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	SerializationLibrary *string `json:"serializationLibrary,omitempty" tf:"serialization_library,omitempty"`
}

type SkewedInfoObservation struct {
}

type SkewedInfoParameters struct {

	// +kubebuilder:validation:Optional
	SkewedColumnNames []*string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names,omitempty"`

	// +kubebuilder:validation:Optional
	SkewedColumnValueLocationMaps map[string]*string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps,omitempty"`

	// +kubebuilder:validation:Optional
	SkewedColumnValues []*string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values,omitempty"`
}

type SortColumnsObservation struct {
}

type SortColumnsParameters struct {

	// +kubebuilder:validation:Required
	Column *string `json:"column" tf:"column,omitempty"`

	// +kubebuilder:validation:Required
	SortOrder *float64 `json:"sortOrder" tf:"sort_order,omitempty"`
}

type StorageDescriptorObservation struct {
}

type StorageDescriptorParameters struct {

	// +kubebuilder:validation:Optional
	BucketColumns []*string `json:"bucketColumns,omitempty" tf:"bucket_columns,omitempty"`

	// +kubebuilder:validation:Optional
	Columns []ColumnsParameters `json:"columns,omitempty" tf:"columns,omitempty"`

	// +kubebuilder:validation:Optional
	Compressed *bool `json:"compressed,omitempty" tf:"compressed,omitempty"`

	// +kubebuilder:validation:Optional
	InputFormat *string `json:"inputFormat,omitempty" tf:"input_format,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	NumberOfBuckets *float64 `json:"numberOfBuckets,omitempty" tf:"number_of_buckets,omitempty"`

	// +kubebuilder:validation:Optional
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaReference []SchemaReferenceParameters `json:"schemaReference,omitempty" tf:"schema_reference,omitempty"`

	// +kubebuilder:validation:Optional
	SerDeInfo []SerDeInfoParameters `json:"serDeInfo,omitempty" tf:"ser_de_info,omitempty"`

	// +kubebuilder:validation:Optional
	SkewedInfo []SkewedInfoParameters `json:"skewedInfo,omitempty" tf:"skewed_info,omitempty"`

	// +kubebuilder:validation:Optional
	SortColumns []SortColumnsParameters `json:"sortColumns,omitempty" tf:"sort_columns,omitempty"`

	// +kubebuilder:validation:Optional
	StoredAsSubDirectories *bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories,omitempty"`
}

type TargetTableObservation struct {
}

type TargetTableParameters struct {

	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// CatalogTableSpec defines the desired state of CatalogTable
type CatalogTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogTableParameters `json:"forProvider"`
}

// CatalogTableStatus defines the observed state of CatalogTable.
type CatalogTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogTable is the Schema for the CatalogTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CatalogTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogTableSpec   `json:"spec"`
	Status            CatalogTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogTableList contains a list of CatalogTables
type CatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogTable `json:"items"`
}

// Repository type metadata.
var (
	CatalogTable_Kind             = "CatalogTable"
	CatalogTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogTable_Kind}.String()
	CatalogTable_KindAPIVersion   = CatalogTable_Kind + "." + CRDGroupVersion.String()
	CatalogTable_GroupVersionKind = CRDGroupVersion.WithKind(CatalogTable_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogTable{}, &CatalogTableList{})
}
