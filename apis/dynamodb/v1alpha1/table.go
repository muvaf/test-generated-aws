// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TableParameters defines the desired state of Table
type TableParameters struct {
	AttributeDefinitions []*AttributeDefinition `json:"attributeDefinitions,omitempty"`
	BillingMode *string `json:"billingMode,omitempty"`
	GlobalSecondaryIndexes []*GlobalSecondaryIndex `json:"globalSecondaryIndexes,omitempty"`
	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	LocalSecondaryIndexes []*LocalSecondaryIndex `json:"localSecondaryIndexes,omitempty"`
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
	SSESpecification *SSESpecification `json:"sseSpecification,omitempty"`
	StreamSpecification *StreamSpecification `json:"streamSpecification,omitempty"`
	TableName *string `json:"tableName,omitempty"`
	Tags []*Tag `json:"tags,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider TableParameters `json:"forProvider"`
}

// TableObservation defines the observed state of Table
type TableObservation struct {
	ArchivalSummary *ArchivalSummary `json:"archivalSummary,omitempty"`
	BillingModeSummary *BillingModeSummary `json:"billingModeSummary,omitempty"`
	CreationDateTime *metav1.Time `json:"creationDateTime,omitempty"`
	GlobalTableVersion *string `json:"globalTableVersion,omitempty"`
	ItemCount *int64 `json:"itemCount,omitempty"`
	LatestStreamARN *string `json:"latestStreamARN,omitempty"`
	LatestStreamLabel *string `json:"latestStreamLabel,omitempty"`
	Replicas []*ReplicaDescription `json:"replicas,omitempty"`
	RestoreSummary *RestoreSummary `json:"restoreSummary,omitempty"`
	SSEDescription *SSEDescription `json:"sseDescription,omitempty"`
	TableID *string `json:"tableID,omitempty"`
	TableSizeBytes *int64 `json:"tableSizeBytes,omitempty"`
	TableStatus *string `json:"tableStatus,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider TableObservation `json:"atProvider"`
}


// +kubebuilder:object:root=true

// Table is the Schema for the Tables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   TableSpec   `json:"spec,omitempty"`
	Status TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []Table `json:"items"`
}

// Repository type metadata.
var (
	TableKind             = "Table"
	TableGroupKind        = schema.GroupKind{Group: Group, Kind: TableKind}.String()
	TableKindAPIVersion   = TableKind + "." + GroupVersion.String()
	TableGroupVersionKind = GroupVersion.WithKind(TableKind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}

