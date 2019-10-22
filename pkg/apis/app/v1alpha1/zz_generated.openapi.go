// +build !ignore_autogenerated

// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/app/v1alpha1.HelmChartSubscription": schema_pkg_apis_app_v1alpha1_HelmChartSubscription(ref),
		"./pkg/apis/app/v1alpha1.HelmRelease":           schema_pkg_apis_app_v1alpha1_HelmRelease(ref),
		"./pkg/apis/app/v1alpha1.HelmReleaseSpec":       schema_pkg_apis_app_v1alpha1_HelmReleaseSpec(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_HelmChartSubscription(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HelmChartSubscription is the Schema for the subscriptions API",
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
							Ref: ref("./pkg/apis/app/v1alpha1.HelmChartSubscriptionSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/app/v1alpha1.HelmChartSubscriptionStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/app/v1alpha1.HelmChartSubscriptionSpec", "./pkg/apis/app/v1alpha1.HelmChartSubscriptionStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_HelmRelease(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HelmRelease is the Schema for the subscriptionreleases API",
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
							Ref: ref("./pkg/apis/app/v1alpha1.HelmReleaseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/app/v1alpha1.HelmReleaseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/app/v1alpha1.HelmReleaseSpec", "./pkg/apis/app/v1alpha1.HelmReleaseStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_HelmReleaseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HelmReleaseSpec defines the desired state of HelmRelease",
				Properties: map[string]spec.Schema{
					"source": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html Source holds the url toward the helm-chart",
							Ref:         ref("./pkg/apis/app/v1alpha1.Source"),
						},
					},
					"chartName": {
						SchemaProps: spec.SchemaProps{
							Description: "ChartName is the name of the chart within the repo",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"releaseName": {
						SchemaProps: spec.SchemaProps{
							Description: "ReleaseName is the Name of the release given to Tiller. Defaults to namespace-name. Must not be changed after initial object creation.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version is the chart version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"values": {
						SchemaProps: spec.SchemaProps{
							Description: "Values is a string containing (unparsed) YAML values",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret to use to access the helm-repo defined in the CatalogSource.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"configMapRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration parameters to access the helm-repo defined in the CatalogSource",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/app/v1alpha1.Source", "k8s.io/api/core/v1.ObjectReference"},
	}
}
