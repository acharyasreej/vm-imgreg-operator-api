// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CLItemPhase indicates the phase of a ContentLibraryItem's lifecycle.
type CLItemPhase string

const (
	// CreatingItem phase indicates that the content library item is being created by the backing infrastructure provider.
	CreatingItem CLItemPhase = "Creating"

	// CreatedItem phase indicates that the content library item has been already been created by the backing infrastructure provider.
	CreatedItem CLItemPhase = "Created"

	// UpdatingItem phase indicates that the content library item is being updated by the backing infrastructure provider.
	UpdatingItem CLItemPhase = "Updating"

	// UpdatedItem phase indicates that the content library item has been already been updated by the backing infrastructure provider.
	UpdatedItem CLItemPhase = "Updated"

	// DeletingItem phase indicates that the content library item is being deleted by the backing infrastructure provider.
	DeletingItem CLItemPhase = "Deleting"

	// DeletedItem phase indicates that the content library item has been deleted by the backing infrastructure provider.
	DeletedItem CLItemPhase = "Deleted"

	// UnknownItemStatus phase indicates that the content library item status cannot be determined from the backing infrastructure provider.
	UnknownItemStatus CLItemPhase = "Unknown"
)

// ContentLibraryItemSpec defines the desired state of a ContentLibraryItem.
type ContentLibraryItemSpec struct {
	// ContentLibraryName is the name of the library that contains the library item.
	// +required
	ContentLibraryName string `json:"contentLibraryName"`

	// ItemName specifies the name of the content library item.
	// +required
	ItemName string `json:"itemName"`

	// ItemDescription is a human-readable description for this library item.
	// +optional
	ItemDescription string `json:"itemDescription,omitempty"`
}

// ContentLibraryItemStatus defines the observed state of ContentLibraryItem.
type ContentLibraryItemStatus struct {
	// ItemUUID is the identifier which uniquely identifies the library item in vCenter.
	ItemUUID string `json:"itemUUID,omitempty"`

	// ItemVersion indicates the version of the library item metadata.
	ItemVersion string `json:"itemVersion,omitempty"`

	// ContentVersion indicates the version of the library item content.
	ContentVersion string `json:"contentVersion,omitempty"`

	// ItemType string indicates the type of the library item.
	ItemType string `json:"itemType,omitempty"`

	// Cached indicates if the files are on disk in vCenter.
	// +optional
	Cached bool `json:"cached,omitempty"`

	// Phase indicates the phase of a ContentLibraryItem's lifecycle.
	Phase CLItemPhase `json:"phase,omitempty"`

	// Ready denotes that the library item is ready to be used.
	Ready bool `json:"ready"`

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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=clitem
// +kubebuilder:printcolumn:name="LibraryName",type="string",JSONPath=".spec.contentLibraryRef.name"
// +kubebuilder:printcolumn:name="ItemName",type="string",JSONPath=".spec.itemName"
// +kubebuilder:printcolumn:name="ItemUUID",type="string",JSONPath=".status.itemUUID"
// +kubebuilder:printcolumn:name="ItemType",type="string",JSONPath=".status.itemType"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContentLibraryItem is the schema for the content library item API.
type ContentLibraryItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibraryItemSpec   `json:"spec,omitempty"`
	Status ContentLibraryItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContentLibraryItemList contains a list of ContentLibraryItem.
type ContentLibraryItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentLibraryItem `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContentLibraryItem{}, &ContentLibraryItemList{})
}
