// Code generated by generators/resource/main.go; DO NOT EDIT.

package fms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_fms_policy", policyResourceType)
}

// policyResourceType returns the Terraform awscc_fms_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FMS::Policy resource type.
func policyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"delete_all_policy_resources": {
			// Property: DeleteAllPolicyResources
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
			// DeleteAllPolicyResources is a write-only property.
		},
		"exclude_map": {
			// Property: ExcludeMap
			// CloudFormation resource type schema:
			// {
			//   "description": "An FMS includeMap or excludeMap.",
			//   "properties": {
			//     "ACCOUNT": {
			//       "items": {
			//         "description": "An AWS account ID.",
			//         "maxLength": 12,
			//         "minLength": 12,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "ORGUNIT": {
			//       "items": {
			//         "description": "An Organizational Unit ID.",
			//         "maxLength": 68,
			//         "minLength": 16,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An FMS includeMap or excludeMap.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"account": {
						// Property: ACCOUNT
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"orgunit": {
						// Property: ORGUNIT
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"exclude_resource_tags": {
			// Property: ExcludeResourceTags
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Required: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"include_map": {
			// Property: IncludeMap
			// CloudFormation resource type schema:
			// {
			//   "description": "An FMS includeMap or excludeMap.",
			//   "properties": {
			//     "ACCOUNT": {
			//       "items": {
			//         "description": "An AWS account ID.",
			//         "maxLength": 12,
			//         "minLength": 12,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "ORGUNIT": {
			//       "items": {
			//         "description": "An Organizational Unit ID.",
			//         "maxLength": 68,
			//         "minLength": 16,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An FMS includeMap or excludeMap.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"account": {
						// Property: ACCOUNT
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"orgunit": {
						// Property: ORGUNIT
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"policy_name": {
			// Property: PolicyName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1024),
			},
		},
		"remediation_enabled": {
			// Property: RemediationEnabled
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Required: true,
		},
		"resource_tags": {
			// Property: ResourceTags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "description": "A resource tag.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxLength": 8,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "An AWS resource type",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "An AWS resource type",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
		},
		"resource_type_list": {
			// Property: ResourceTypeList
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "description": "An AWS resource type",
			//     "maxLength": 128,
			//     "minLength": 1,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"security_service_policy_data": {
			// Property: SecurityServicePolicyData
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "ManagedServiceData": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Type": {
			//       "enum": [
			//         "WAF",
			//         "WAFV2",
			//         "SHIELD_ADVANCED",
			//         "SECURITY_GROUPS_COMMON",
			//         "SECURITY_GROUPS_CONTENT_AUDIT",
			//         "SECURITY_GROUPS_USAGE_AUDIT",
			//         "NETWORK_FIREWALL",
			//         "DNS_FIREWALL"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"managed_service_data": {
						// Property: ManagedServiceData
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 4096),
						},
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"WAF",
								"WAFV2",
								"SHIELD_ADVANCED",
								"SECURITY_GROUPS_COMMON",
								"SECURITY_GROUPS_CONTENT_AUDIT",
								"SECURITY_GROUPS_USAGE_AUDIT",
								"NETWORK_FIREWALL",
								"DNS_FIREWALL",
							}),
						},
					},
				},
			),
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "description": "A policy tag.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Creates an AWS Firewall Manager policy.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FMS::Policy").WithTerraformTypeName("awscc_fms_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account":                      "ACCOUNT",
		"arn":                          "Arn",
		"delete_all_policy_resources":  "DeleteAllPolicyResources",
		"exclude_map":                  "ExcludeMap",
		"exclude_resource_tags":        "ExcludeResourceTags",
		"id":                           "Id",
		"include_map":                  "IncludeMap",
		"key":                          "Key",
		"managed_service_data":         "ManagedServiceData",
		"orgunit":                      "ORGUNIT",
		"policy_name":                  "PolicyName",
		"remediation_enabled":          "RemediationEnabled",
		"resource_tags":                "ResourceTags",
		"resource_type":                "ResourceType",
		"resource_type_list":           "ResourceTypeList",
		"security_service_policy_data": "SecurityServicePolicyData",
		"tags":                         "Tags",
		"type":                         "Type",
		"value":                        "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DeleteAllPolicyResources",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
