// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.API":                  schema_pkg_apis_wso2_v1alpha1_API(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.APISpec":              schema_pkg_apis_wso2_v1alpha1_APISpec(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.APIStatus":            schema_pkg_apis_wso2_v1alpha1_APIStatus(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.RateLimiting":         schema_pkg_apis_wso2_v1alpha1_RateLimiting(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.RateLimitingSpec":     schema_pkg_apis_wso2_v1alpha1_RateLimitingSpec(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.RateLimitingStatus":   schema_pkg_apis_wso2_v1alpha1_RateLimitingStatus(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Security":             schema_pkg_apis_wso2_v1alpha1_Security(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecuritySpec":         schema_pkg_apis_wso2_v1alpha1_SecuritySpec(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecurityStatus":       schema_pkg_apis_wso2_v1alpha1_SecurityStatus(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.TargetEndpoint":       schema_pkg_apis_wso2_v1alpha1_TargetEndpoint(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.TargetEndpointSpec":   schema_pkg_apis_wso2_v1alpha1_TargetEndpointSpec(ref),
		"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.TargetEndpointStatus": schema_pkg_apis_wso2_v1alpha1_TargetEndpointStatus(ref),
	}
}

func schema_pkg_apis_wso2_v1alpha1_API(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "API is the Schema for the apis API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.APISpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.APIStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.APISpec", "github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.APIStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_APISpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APISpec defines the desired state of API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"updateTimeStamp": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"definition": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Definition"),
						},
					},
					"override": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"replicas", "definition"},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Definition"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_APIStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIStatus defines the observed state of API",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_wso2_v1alpha1_RateLimiting(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RateLimiting is the Schema for the ratelimitings API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.RateLimitingSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.RateLimitingSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_RateLimitingSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RateLimitingSpec defines the desired state of RateLimiting",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"timeUnit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"unitTime": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"requestCount": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.RequestCount"),
						},
					},
					"stopOnQuotaReach": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bandwidth": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Bandwidth"),
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Conditions"),
						},
					},
				},
				Required: []string{"type", "timeUnit", "unitTime", "requestCount", "stopOnQuotaReach", "description", "bandwidth", "conditions"},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Bandwidth", "github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Conditions", "github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.RequestCount"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_RateLimitingStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RateLimitingStatus defines the observed state of RateLimiting",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_wso2_v1alpha1_Security(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Security is the Schema for the securities API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecuritySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecurityStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecuritySpec", "github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecurityStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_SecuritySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecuritySpec defines the desired state of Security",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"securityConfig": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecurityConfig"),
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "securityConfig"},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.SecurityConfig"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_SecurityStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecurityStatus defines the observed state of Security",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_wso2_v1alpha1_TargetEndpoint(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TargetEndpoint is the Schema for the targetendpoints API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.TargetEndpointSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.TargetEndpointStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.TargetEndpointSpec", "github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.TargetEndpointStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_TargetEndpointSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TargetEndpointSpec defines the desired state of TargetEndpoint",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"protocol": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"targetPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ports": {
						SchemaProps: spec.SchemaProps{
							Description: "List of optional ports of the target endpoint.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Ports"),
									},
								},
							},
						},
					},
					"deploy": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Deploy"),
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "Mode of the Target Endpoint. Applicable values: (privateJet|sidecar|serverless). Default value: privateJet",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"protocol", "port", "targetPort", "deploy"},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Deploy", "github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1.Ports"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_TargetEndpointStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TargetEndpointStatus defines the observed state of TargetEndpoint",
				Type:        []string{"object"},
			},
		},
	}
}
