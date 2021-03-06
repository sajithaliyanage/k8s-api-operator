// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RateLimitingSpec defines the desired state of RateLimiting
// +k8s:openapi-gen=true
type RateLimitingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	Type             string       `json:"type"`
	TimeUnit         string       `json:"timeUnit"`
	UnitTime         int          `json:"unitTime"`
	RequestCount     RequestCount `json:"requestCount"`
	StopOnQuotaReach bool         `json:"stopOnQuotaReach"`
	Description      string       `json:"description"`
	Bandwidth        Bandwidth    `json:"bandwidth"`
	Conditions       Conditions   `json:"conditions"`
}

//RequestCount is exported type in Ratelimiting Spec
type RequestCount struct {
	Limit int `json:"limit"`
}

//Bandwidth is exported type in Ratelimiting Spec
type Bandwidth struct {
	DataAmount string `json:"dataAmount"`
	DataUnit   string `json:"dataUnit"`
}

//Conditions is exported type in Ratelimiting Spec
type Conditions struct {
	HeaderCondition HeaderCondition `json:"headerCondition"`
	IPCondition     IPCondition     `json:"ipCondition"`
}

//HeaderCondition is exported type in Ratelimiting Spec
type HeaderCondition struct {
	HeaderName  string `json:"headerName"`
	HeaderValue string `json:"headerValue"`
}

//IPCondition is exported type in Ratelimiting Spec
type IPCondition struct {
	Type       string `json:"type"`
	SpecificIP string `json:"specificIp"`
	Negation   bool   `json:"negation"`
	StartIP    string `json:"startIp"`
	EndIP      string `json:"endIp"`
}

// RateLimitingStatus defines the observed state of RateLimiting
// +k8s:openapi-gen=true
type RateLimitingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RateLimiting is the Schema for the ratelimitings API
// +k8s:openapi-gen=true
type RateLimiting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RateLimitingSpec `json:"spec,omitempty"`
	//Status RateLimitingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RateLimitingList contains a list of RateLimiting
type RateLimitingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RateLimiting `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RateLimiting{}, &RateLimitingList{})
}
