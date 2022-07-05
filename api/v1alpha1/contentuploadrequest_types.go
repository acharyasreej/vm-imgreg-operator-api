// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UploadPhase indicates the phase of a ContentLibraryItem's content.
type UploadPhase string

const (
	// UploadingContent phase indicates that the content library item content is being uploaded.
	UploadingContent UploadPhase = "Uploading"

	// UploadedContent phase indicates that the content library item content has been uploaded.
	UploadedContent UploadPhase = "Uploaded"
)

// LibraryItem defines the desried state of a content library item.
type LibraryItem struct {
	// Name specifies the name of the content library item.
	// +required
	Name string `json:"name"`

	// Description is a human-readable description for this library item.
	// +optional
	Description string `json:"description"`

	// Overwrite indicates if the content of an existing content library item
	// with the specified name should be overwritten.
	// +optional
	// +default=false
	Overwrite bool `json:"overwrite"`
}

// UploadSpec defines the spec for the content of the library item.
type UploadSpec struct {
	// SourceType defines how the file content is retrieved.
	SourceType string `json:"sourceType"`

	// SourceEndpoint is the source endpoint from which the file will be retrieved.
	SourceEndpoint string `json:"sourceEndpoint"`
}

// ContentUploadRequestSpec defines the desired state of a ContentUploadRequest.
type ContentUploadRequestSpec struct {
	// ContentLibraryName is the name of the library that contains the library item.
	// +required
	LibraryName string `json:"libraryName"`

	// ContentLibraryItem defines the desired state of a content library item.
	// +required
	LibraryItem LibraryItem `json:"libraryItem"`

	// UploadSpec defines the spec for the content of the library item.
	// +required
	UploadSpec UploadSpec `json:"uploadSpec"`
}

// ContentUploadRequestStatus defines the observed state of a ContentUploadRequest.
type ContentUploadRequestStatus struct {
	// ItemUUID is the identifier which uniquely identifies the library item in vCenter.
	ItemUUID string `json:"itemUUID,omitempty"`

	// ItemVersion indicates the version of the library item metadata.
	ItemVersion string `json:"itemVersion,omitempty"`

	// ContentVersion indicates the version of the library item content.
	ContentVersion string `json:"contentVersion,omitempty"`

	// ItemType string indicates the type of the library item.
	ItemType string `json:"itemType,omitempty"`

	// Phase indicates the phase of a ContentLibraryItem's lifecycle.
	Phase UploadPhase `json:"phase,omitempty"`

	// Ready denotes that the library item is ready to be used.
	Ready bool `json:"ready"`

	// Conditions describes the current condition information of the ContentLibraryItem.
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
}

func (contentUploadRequest *ContentUploadRequest) GetConditions() Conditions {
	return contentUploadRequest.Status.Conditions
}

func (contentUploadRequest *ContentUploadRequest) SetConditions(conditions Conditions) {
	contentUploadRequest.Status.Conditions = conditions
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=uploadreq
// +kubebuilder:printcolumn:name="LibraryName",type="string",JSONPath=".spec.libraryName"
// +kubebuilder:printcolumn:name="ItemName",type="string",JSONPath=".spec.libraryItem.name"
// +kubebuilder:printcolumn:name="ItemUUID",type="string",JSONPath=".status.itemUUID"
// +kubebuilder:printcolumn:name="ItemType",type="string",JSONPath=".status.itemType"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready"

// ContentUploadRequest is the schema for the content library item content upload request API.
type ContentUploadRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentUploadRequestSpec   `json:"spec,omitempty"`
	Status ContentUploadRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentUploadRequestList contains a list of ContentUploadRequest.
type ContentUploadRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentUploadRequest `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ContentUploadRequest{}, &ContentUploadRequestList{})
}
