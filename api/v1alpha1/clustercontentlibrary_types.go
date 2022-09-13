// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterContentLibrarySpec defines the desired state of a ClusterContentLibrary.
type ClusterContentLibrarySpec struct {
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

// ClusterContentLibraryStatus defines the observed state of ClusterContentLibrary.
type ClusterContentLibraryStatus struct {
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

func (contentLibrary *ClusterContentLibrary) GetConditions() Conditions {
	return contentLibrary.Status.Conditions
}

func (contentLibrary *ClusterContentLibrary) SetConditions(conditions Conditions) {
	contentLibrary.Status.Conditions = conditions
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=clustercl
// +kubebuilder:printcolumn:name="LibraryName",type="string",JSONPath=".spec.libraryName"
// +kubebuilder:printcolumn:name="UUID",type="string",JSONPath=".status.libraryUUID"
// +kubebuilder:printcolumn:name="LibraryType",type="string",JSONPath=".status.libraryType"
// +kubebuilder:printcolumn:name="StorageType",type="string",JSONPath=".spec.storageBacking.storageType"

// ClusterContentLibrary is the schema for the cluster scoped content library API.
// Currently, ClusterContentLibrary is immutable to end users.
type ClusterContentLibrary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterContentLibrarySpec   `json:"spec,omitempty"`
	Status ClusterContentLibraryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterContentLibraryList contains a list of ClusterContentLibrary.
type ClusterContentLibraryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterContentLibrary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterContentLibrary{}, &ClusterContentLibraryList{})
}
