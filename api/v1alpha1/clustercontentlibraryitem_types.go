// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterContentLibraryItemSpec defines the desired state of a ClusterContentLibraryItem.
type ClusterContentLibraryItemSpec struct {
	// ContentLibraryRefName is the name of the ContentLibrary custom resource.
        // +required
        ContentLibraryRefName string `json:"contentLibraryRefName"`

	// ItemDescription is a human-readable description for this library item
	// +optional
	ItemDescription string `json:"itemDescription,omitempty"`
}

// ClusterContentLibraryItemStatus defines the observed state of ContentLibraryItem.
type ClusterContentLibraryItemStatus struct {
	// ItemUUID is the identifier which uniquely identifies the library item in vCenter
	ItemUUID string `json:"itemUUID,omitempty"`

	// ItemName specifies the name of the content library item in vCenter
        ItemName string `json:"itemName"`

	// ItemVersion indicates the version of the library item metadata
	ItemVersion string `json:"itemVersion,omitempty"`

	// ContentVersion indicates the version of the library item content
	ContentVersion string `json:"contentVersion,omitempty"`

	// ItemType string indicates the type of the library item in vCenter
	ItemType string `json:"itemType,omitempty"`

	// Cached indicates if the files are on disk in vCenter
	// +optional
	Cached bool `json:"cached,omitempty"`

	// Ready denotes that the library item is ready to be used
	Ready bool `json:"ready"`

	// Conditions describes the current condition information of the ContentLibraryItem
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
}

func (contentLibraryItem *ClusterContentLibraryItem) GetConditions() Conditions {
	return contentLibraryItem.Status.Conditions
}

func (contentLibraryItem *ClusterContentLibraryItem) SetConditions(conditions Conditions) {
	contentLibraryItem.Status.Conditions = conditions
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=clusterclitem
// +kubebuilder:printcolumn:name="ContentLibraryRefName",type="string",JSONPath=".spec.contentLibraryRefName"
// +kubebuilder:printcolumn:name="ItemName",type="string",JSONPath=".status.itemName"
// +kubebuilder:printcolumn:name="ItemUUID",type="string",JSONPath=".status.itemUUID"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ClusterContentLibraryItem is the schema for the content library item API at the cluster scope.
// Currently, ClusterContentLibraryItem are immutable to end users.
type ClusterContentLibraryItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterContentLibraryItemSpec   `json:"spec,omitempty"`
	Status ClusterContentLibraryItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterContentLibraryItemList contains a list of ClusterContentLibraryItem.
type ClusterContentLibraryItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterContentLibraryItem `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ClusterContentLibraryItem{}, &ClusterContentLibraryItemList{})
}
