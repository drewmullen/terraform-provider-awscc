// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iotwireless_wireless_gateway", wirelessGatewayDataSourceType)
}

// wirelessGatewayDataSourceType returns the Terraform awscc_iotwireless_wireless_gateway data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTWireless::WirelessGateway resource type.
func wirelessGatewayDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn for Wireless Gateway. Returned upon successful create.",
			//   "type": "string"
			// }
			Description: "Arn for Wireless Gateway. Returned upon successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of Wireless Gateway.",
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Description: "Description of Wireless Gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id for Wireless Gateway. Returned upon successful create.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Id for Wireless Gateway. Returned upon successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_uplink_received_at": {
			// Property: LastUplinkReceivedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the most recent uplink was received.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date and time when the most recent uplink was received.",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "GatewayEui": {
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "RfRegion": {
			//       "maxLength": 64,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "GatewayEui",
			//     "RfRegion"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"gateway_eui": {
						// Property: GatewayEui
						Type:     types.StringType,
						Computed: true,
					},
					"rf_region": {
						// Property: RfRegion
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Wireless Gateway.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Name of Wireless Gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the gateway.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the gateway.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"thing_arn": {
			// Property: ThingArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
			//   "type": "string"
			// }
			Description: "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"thing_name": {
			// Property: ThingName
			// CloudFormation resource type schema:
			// {
			//   "description": "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
			//   "type": "string"
			// }
			Description: "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTWireless::WirelessGateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::WirelessGateway").WithTerraformTypeName("awscc_iotwireless_wireless_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"description":             "Description",
		"gateway_eui":             "GatewayEui",
		"id":                      "Id",
		"key":                     "Key",
		"last_uplink_received_at": "LastUplinkReceivedAt",
		"lo_ra_wan":               "LoRaWAN",
		"name":                    "Name",
		"rf_region":               "RfRegion",
		"tags":                    "Tags",
		"thing_arn":               "ThingArn",
		"thing_name":              "ThingName",
		"value":                   "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
