// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageType is used to indicate the type of the storage backing where the content would be stored.
type StorageType string
type LibraryType string

const (
	Datastore  StorageType = "Datastore"
	Other      StorageType = "Other"
	Local      LibraryType = "Local"
	Subscribed LibraryType = "Subscribed"
)

// StorageBacking describes the default storage backing which is available for the library.
type StorageBacking struct {
	// StorageType indicates the type of storage where the content would be stored.
	// +required
	StorageType StorageType `json:"storageType"`

	// DatastoreID indicates the identifier of the datastore used to store the content in the library for the DATASTORE storageType.
	// +optional
	DatastoreID string `json:"datastoreID,omitempty"`
}

// ContentLibrarySpec defines the desired state of a ContentLibrary.
type ContentLibrarySpec struct {
	// LibraryName specifies the name of the content library.
	// +required
	LibraryName string `json:"libraryName"`

	// LibraryDescription is a human-readable description for this library.
	// +optional
	LibraryDescription string `json:"libraryDescription,omitempty"`

	// StorageBackings indicates the default storage backings which are available for this library.
	// +required
	StorageBacking StorageBacking `json:"storageBacking"`
}

// ContentLibraryStatus defines the observed state of ContentLibrary.
type ContentLibraryStatus struct {
	// LibraryUUID is the identifier which uniquely identifies the library in vCenter.
	LibraryUUID string `json:"libraryUUID,omitempty"`

	// Type indicates the type of a library in vCenter.
	// Possible types are Local and Subscribed.
	LibraryType LibraryType `json:"libraryType,omitempty"`

	// Version is the version number that can identify metadata changes.
	Version string `json:"version,omitempty"`

	// Conditions describes the current condition information of the ContentLibrary.
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
}

func (contentLibrary *ContentLibrary) GetConditions() Conditions {
	return contentLibrary.Status.Conditions
}

func (contentLibrary *ContentLibrary) SetConditions(conditions Conditions) {
	contentLibrary.Status.Conditions = conditions
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=cl
// +kubebuilder:printcolumn:name="LibraryName",type="string",JSONPath=".spec.libraryName"
// +kubebuilder:printcolumn:name="UUID",type="string",JSONPath=".status.libraryUUID"
// +kubebuilder:printcolumn:name="LibraryType",type="string",JSONPath=".status.libraryType"
// +kubebuilder:printcolumn:name="StorageType",type="string",JSONPath=".spec.storageBacking.storageType"

// ContentLibrary is the schema for the content library API.
// Currently, ContentLibrary is immutable to end users.
type ContentLibrary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibrarySpec   `json:"spec,omitempty"`
	Status ContentLibraryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentLibraryList contains a list of ContentLibrary.
type ContentLibraryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentLibrary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContentLibrary{}, &ContentLibraryList{})
}
