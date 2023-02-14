// Copyright (c) 2022-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// ContentLibraryItemType is a constant for the type of a content library item in vCenter.
type ContentLibraryItemType string

const (
	// ContentLibraryItemTypeOvf indicates an OVF content library item in vCenter.
	ContentLibraryItemTypeOvf ContentLibraryItemType = "OVF"

	// ContentLibraryItemTypeIso indicates an ISO content library item in vCenter.
	ContentLibraryItemTypeIso ContentLibraryItemType = "ISO"
)

// ContentLibraryItemSpec defines the desired state of a ContentLibraryItem.
type ContentLibraryItemSpec struct {
	// UUID is the identifier which uniquely identifies the library item in vCenter. This field is immutable.
	// +required
	UUID types.UID `json:"uuid"`
}

// ContentLibraryItemStatus defines the observed state of ContentLibraryItem.
type ContentLibraryItemStatus struct {
	// Name specifies the name of the content library item in vCenter specified by the user.
	// +optional
	Name string `json:"name,omitempty"`

	// ContentLibraryRef refers to the ContentLibrary custom resource that this item belongs to.
	// +optional
	ContentLibraryRef *NameAndKindRef `json:"contentLibraryRef,omitempty"`

	// Description is a human-readable description for this library item.
	// +optional
	Description string `json:"description,omitempty"`

	// MetadataVersion indicates the version of the library item metadata in vCenter.
	// This integer value is incremented when the library item properties such as name or description are changed in vCenter.
	// +optional
	MetadataVersion string `json:"metadataVersion,omitempty"`

	// ContentVersion indicates the version of the library item content in vCenter.
	// This integer value is incremented when the files comprising the content library item are changed in vCenter.
	// +optional
	ContentVersion string `json:"contentVersion,omitempty"`

	// Type string indicates the type of the library item in vCenter.
	// +kubebuilder:validation:Enum=OVF;ISO
	// +optional
	Type ContentLibraryItemType `json:"type,omitempty"`

	// Size indicates the library item size in bytes in vCenter.
	// +optional
	Size resource.Quantity `json:"size,omitempty"`

	// Cached indicates if the library item files are on disk in vCenter.
	// +optional
	Cached bool `json:"cached,omitempty"`

	// Ready denotes that the library item is ready to be used.
	// This flag is set only after all the files associated with the content library item
	// have been uploaded to vCenter.
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Files represent zero, one or more files belonging to the content library item in vCenter.
	// +optional
	// Files []string `json:"files,omitempty"`

	// CreationTime indicates the date and time when this library item was created in vCenter.
	// +optional
	CreationTime metav1.Time `json:"creationTime,omitempty"`

	// LastModifiedTime indicates the date and time when this library item was last updated in vCenter.
	// This field is updated when the library item properties are changed or the file content is changed.
	// +optional
	LastModifiedTime metav1.Time `json:"lastModifiedTime,omitempty"`

	// LastSyncTime indicates the date and time when this library item was last synchronized in vCenter.
	// This field applies only to the library items belonging to the library of Type=Subscribed.
	// +optional
	LastSyncTime metav1.Time `json:"lastSyncTime,omitempty"`

	// Conditions describes the current condition information of the ContentLibraryItem.
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
}

func (contentLibraryItem *ContentLibraryItem) GetConditions() Conditions {
	return contentLibraryItem.Status.Conditions
}

func (contentLibraryItem *ContentLibraryItem) SetConditions(conditions Conditions) {
	contentLibraryItem.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=clitem
// +kubebuilder:printcolumn:name="ContentLibraryRef",type="string",JSONPath=".status.contentLibraryRef.name"
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".status.type"
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready"
// +kubebuilder:printcolumn:name="LastSyncTime",type="string",JSONPath=".status.lastSyncTime"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ContentLibraryItem is the schema for the content library item API.
type ContentLibraryItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibraryItemSpec   `json:"spec,omitempty"`
	Status ContentLibraryItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentLibraryItemList contains a list of ContentLibraryItem.
type ContentLibraryItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentLibraryItem `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=cclitem
// +kubebuilder:printcolumn:name="ClusterContentLibraryRef",type="string",JSONPath=".status.clusterContentLibraryRef"
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".status.type"
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready"
// +kubebuilder:printcolumn:name="LastSyncTime",type="string",JSONPath=".status.lastSyncTime"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ClusterContentLibraryItem is the schema for the content library item API at the cluster scope.
// Currently, ClusterContentLibraryItem are immutable to end users.
type ClusterContentLibraryItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibraryItemSpec   `json:"spec,omitempty"`
	Status ContentLibraryItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterContentLibraryItemList contains a list of ClusterContentLibraryItem.
type ClusterContentLibraryItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterContentLibraryItem `json:"items"`
}

func init() {
	RegisterTypeWithScheme(
		&ContentLibraryItem{},
		&ContentLibraryItemList{},
		&ClusterContentLibraryItem{},
		&ClusterContentLibraryItemList{})
}
