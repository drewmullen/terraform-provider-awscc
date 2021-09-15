// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_networkfirewall_firewall_policy", firewallPolicyResourceType)
}

// firewallPolicyResourceType returns the Terraform awscc_networkfirewall_firewall_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkFirewall::FirewallPolicy resource type.
func firewallPolicyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 512),
			},
		},
		"firewall_policy": {
			// Property: FirewallPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "StatefulRuleGroupReferences": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ResourceArn": {
			//             "description": "A resource ARN.",
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ResourceArn"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "StatelessCustomActions": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ActionDefinition": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "PublishMetricAction": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Dimensions": {
			//                     "insertionOrder": false,
			//                     "items": {
			//                       "additionalProperties": false,
			//                       "properties": {
			//                         "Value": {
			//                           "maxLength": 128,
			//                           "minLength": 1,
			//                           "pattern": "",
			//                           "type": "string"
			//                         }
			//                       },
			//                       "required": [
			//                         "Value"
			//                       ],
			//                       "type": "object"
			//                     },
			//                     "type": "array",
			//                     "uniqueItems": true
			//                   }
			//                 },
			//                 "required": [
			//                   "Dimensions"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ActionName": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ActionName",
			//           "ActionDefinition"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "StatelessDefaultActions": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "StatelessFragmentDefaultActions": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "StatelessRuleGroupReferences": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Priority": {
			//             "maximum": 65535,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "ResourceArn": {
			//             "description": "A resource ARN.",
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ResourceArn",
			//           "Priority"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "required": [
			//     "StatelessDefaultActions",
			//     "StatelessFragmentDefaultActions"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"stateful_rule_group_references": {
						// Property: StatefulRuleGroupReferences
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"resource_arn": {
									// Property: ResourceArn
									Description: "A resource ARN.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 256),
									},
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Optional: true,
					},
					"stateless_custom_actions": {
						// Property: StatelessCustomActions
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"action_definition": {
									// Property: ActionDefinition
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"publish_metric_action": {
												// Property: PublishMetricAction
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"dimensions": {
															// Property: Dimensions
															Attributes: tfsdk.SetNestedAttributes(
																map[string]tfsdk.Attribute{
																	"value": {
																		// Property: Value
																		Type:     types.StringType,
																		Required: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringLenBetween(1, 128),
																		},
																	},
																},
																tfsdk.SetNestedAttributesOptions{},
															),
															Required: true,
														},
													},
												),
												Optional: true,
											},
										},
									),
									Required: true,
								},
								"action_name": {
									// Property: ActionName
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 128),
									},
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Optional: true,
					},
					"stateless_default_actions": {
						// Property: StatelessDefaultActions
						Type:     types.SetType{ElemType: types.StringType},
						Required: true,
					},
					"stateless_fragment_default_actions": {
						// Property: StatelessFragmentDefaultActions
						Type:     types.SetType{ElemType: types.StringType},
						Required: true,
					},
					"stateless_rule_group_references": {
						// Property: StatelessRuleGroupReferences
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"priority": {
									// Property: Priority
									Type:     types.NumberType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(1, 65535),
									},
								},
								"resource_arn": {
									// Property: ResourceArn
									Description: "A resource ARN.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 256),
									},
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"firewall_policy_arn": {
			// Property: FirewallPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"firewall_policy_id": {
			// Property: FirewallPolicyId
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
		"firewall_policy_name": {
			// Property: FirewallPolicyName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 0,
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
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
							validate.StringLenBetween(0, 255),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource type definition for AWS::NetworkFirewall::FirewallPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::FirewallPolicy").WithTerraformTypeName("awscc_networkfirewall_firewall_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action_definition":                  "ActionDefinition",
		"action_name":                        "ActionName",
		"description":                        "Description",
		"dimensions":                         "Dimensions",
		"firewall_policy":                    "FirewallPolicy",
		"firewall_policy_arn":                "FirewallPolicyArn",
		"firewall_policy_id":                 "FirewallPolicyId",
		"firewall_policy_name":               "FirewallPolicyName",
		"key":                                "Key",
		"priority":                           "Priority",
		"publish_metric_action":              "PublishMetricAction",
		"resource_arn":                       "ResourceArn",
		"stateful_rule_group_references":     "StatefulRuleGroupReferences",
		"stateless_custom_actions":           "StatelessCustomActions",
		"stateless_default_actions":          "StatelessDefaultActions",
		"stateless_fragment_default_actions": "StatelessFragmentDefaultActions",
		"stateless_rule_group_references":    "StatelessRuleGroupReferences",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
