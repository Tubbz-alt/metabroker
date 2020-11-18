/*
Copyright 2020 SUSE

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OfferingSpec defines the Service Offering spec.
type OfferingSpec struct {
	// The name of the Provider this Offering belongs to.
	Provider string `json:"provider,omitempty"`

	// A unique ID for the Offering.
	ID string `json:"id,omitempty"`
	// A description for the Offering.
	Description string `json:"description,omitempty"`
	// Whether the Plans linked to this Offering are bindable or not.
	Bindable bool `json:"bindable,omitempty"`
}

// OfferingStatus defines the observed state of Offering.
type OfferingStatus struct {
	// TODO: implement.
}

// +kubebuilder:object:root=true

// Offering is the Schema for the offerings API.
type Offering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OfferingSpec   `json:"spec,omitempty"`
	Status OfferingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OfferingList contains a list of Offering.
type OfferingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Offering `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Offering{}, &OfferingList{})
}
