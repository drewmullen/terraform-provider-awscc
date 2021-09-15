// Code generated by generators/resource/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_amplify_domain", domainResourceType)
}

// domainResourceType returns the Terraform awscc_amplify_domain resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Amplify::Domain resource type.
func domainResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_id": {
			// Property: AppId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 20),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"auto_sub_domain_creation_patterns": {
			// Property: AutoSubDomainCreationPatterns
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 2048,
			//     "minLength": 1,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"auto_sub_domain_iam_role": {
			// Property: AutoSubDomainIAMRole
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1000),
			},
		},
		"certificate_record": {
			// Property: CertificateRecord
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"domain_status": {
			// Property: DomainStatus
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"enable_auto_sub_domain": {
			// Property: EnableAutoSubDomain
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"status_reason": {
			// Property: StatusReason
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"sub_domain_settings": {
			// Property: SubDomainSettings
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "BranchName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Prefix": {
			//         "maxLength": 255,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Prefix",
			//       "BranchName"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 255,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"branch_name": {
						// Property: BranchName
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"prefix": {
						// Property: Prefix
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 255),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 255,
				},
			),
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Domain").WithTerraformTypeName("awscc_amplify_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                            "AppId",
		"arn":                               "Arn",
		"auto_sub_domain_creation_patterns": "AutoSubDomainCreationPatterns",
		"auto_sub_domain_iam_role":          "AutoSubDomainIAMRole",
		"branch_name":                       "BranchName",
		"certificate_record":                "CertificateRecord",
		"domain_name":                       "DomainName",
		"domain_status":                     "DomainStatus",
		"enable_auto_sub_domain":            "EnableAutoSubDomain",
		"prefix":                            "Prefix",
		"status_reason":                     "StatusReason",
		"sub_domain_settings":               "SubDomainSettings",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
