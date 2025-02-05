/*
Copyright 2022.

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

type State string

// Valid IstioCR States.
const (
	Ready      State = "Ready"
	Processing State = "Processing"
	Error      State = "Error"
	Deleting   State = "Deleting"
)

// Defines the desired specification for installing or updating Istio.
type IstioSpec struct {
	// +kubebuilder:validation:Optional
	Config Config `json:"config,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:JSONPath=".status.state",name="State",type="string"

// Contains Istio CR specification and current status.
type Istio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IstioSpec   `json:"spec,omitempty"`
	Status IstioStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of Istio's.
type IstioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Istio `json:"items"`
}

// IstioStatus defines the observed state of IstioCR.
type IstioStatus struct {
	// State signifies current state of CustomObject. Value
	// can be one of ("Ready", "Processing", "Error", "Deleting").
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=Processing;Deleting;Ready;Error
	State State `json:"state"`
	//  Conditions associated with IstioStatus.
	Conditions *[]metav1.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Istio{}, &IstioList{})
}

func (i *Istio) GetStatus() IstioStatus {
	return i.Status
}

func (i *Istio) SetStatus(status IstioStatus) {
	i.Status = status
}

func (i *Istio) ComponentName() string {
	return "istio"
}
