// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/tigera/operator/pkg/apis/operator/v1beta1.AmazonCloudIntegration":       schema_pkg_apis_operator_v1beta1_AmazonCloudIntegration(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1beta1.AmazonCloudIntegrationSpec":   schema_pkg_apis_operator_v1beta1_AmazonCloudIntegrationSpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1beta1.AmazonCloudIntegrationStatus": schema_pkg_apis_operator_v1beta1_AmazonCloudIntegrationStatus(ref),
	}
}

func schema_pkg_apis_operator_v1beta1_AmazonCloudIntegration(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AmazonCloudIntegration is the Schema for the amazoncloudintegrations API",
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
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1beta1.AmazonCloudIntegrationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1beta1.AmazonCloudIntegrationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1beta1.AmazonCloudIntegrationSpec", "github.com/tigera/operator/pkg/apis/operator/v1beta1.AmazonCloudIntegrationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1beta1_AmazonCloudIntegrationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AmazonCloudIntegrationSpec defines the desired state of AmazonCloudIntegration",
				Properties: map[string]spec.Schema{
					"defaultPodMetadataAccess": {
						SchemaProps: spec.SchemaProps{
							Description: "DefaultPodMetadataAccess defines what the default behavior will be for accessing the AWS metadata service from a pod. Default: Denied",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodeSecurityGroupIDs": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSecurityGroupIDs is a list of Security Group IDs that all nodes and masters will be in.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"podSecurityGroupID": {
						SchemaProps: spec.SchemaProps{
							Description: "PodSecurityGroupID is the ID of the Security Group which all pods should be placed in by default.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"vpcs": {
						SchemaProps: spec.SchemaProps{
							Description: "Vpcs is a list of VPC IDs to monitor for ENIs and Security Groups, only one is supported.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"sqsURL": {
						SchemaProps: spec.SchemaProps{
							Description: "SqsURL is the SQS URL needed to access the Simple Queue Service.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"awsRegion": {
						SchemaProps: spec.SchemaProps{
							Description: "AwsRegion is the region in which your cluster is located.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"enforcedSecurityGroupID": {
						SchemaProps: spec.SchemaProps{
							Description: "EnforcedSecurityGroupID is the ID of the Security Group which will be applied to all ENIs that are on a host that is also part of the Kubernetes cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"trustEnforcedSecurityGroupID": {
						SchemaProps: spec.SchemaProps{
							Description: "TrustEnforcedSecurityGroupID is the ID of the Security Group which will be applied to all ENIs in the VPC.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"trustEnforcedSecurityGroupID"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1beta1_AmazonCloudIntegrationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AmazonCloudIntegrationStatus defines the observed state of AmazonCloudIntegration",
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "State provides user-readable status.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}