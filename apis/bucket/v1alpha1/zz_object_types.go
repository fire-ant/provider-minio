/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ObjectObservation struct {
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ObjectName *string `json:"objectName,omitempty" tf:"object_name,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type ObjectParameters struct {

	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectName *string `json:"objectName,omitempty" tf:"object_name,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

// ObjectSpec defines the desired state of Object
type ObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectParameters `json:"forProvider"`
}

// ObjectStatus defines the observed state of Object.
type ObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Object is the Schema for the Objects API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,minio}
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.objectName)",message="objectName is a required parameter"
	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Objects
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}

// Repository type metadata.
var (
	Object_Kind             = "Object"
	Object_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Object_Kind}.String()
	Object_KindAPIVersion   = Object_Kind + "." + CRDGroupVersion.String()
	Object_GroupVersionKind = CRDGroupVersion.WithKind(Object_Kind)
)

func init() {
	SchemeBuilder.Register(&Object{}, &ObjectList{})
}
