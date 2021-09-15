// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_s3_storage_lens", storageLensDataSourceType)
}

// storageLensDataSourceType returns the Terraform awscc_s3_storage_lens data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::S3::StorageLens resource type.
func storageLensDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"storage_lens_configuration": {
			// Property: StorageLensConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies the details of Amazon S3 Storage Lens configuration.",
			//   "properties": {
			//     "AccountLevel": {
			//       "additionalProperties": false,
			//       "description": "Account-level metrics configurations.",
			//       "properties": {
			//         "ActivityMetrics": {
			//           "additionalProperties": false,
			//           "description": "Enables activity metrics.",
			//           "properties": {
			//             "IsEnabled": {
			//               "description": "Specifies whether activity metrics are enabled or disabled.",
			//               "type": "boolean"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "BucketLevel": {
			//           "additionalProperties": false,
			//           "description": "Bucket-level metrics configurations.",
			//           "properties": {
			//             "ActivityMetrics": {
			//               "additionalProperties": false,
			//               "description": "Enables activity metrics.",
			//               "properties": {
			//                 "IsEnabled": {
			//                   "description": "Specifies whether activity metrics are enabled or disabled.",
			//                   "type": "boolean"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "PrefixLevel": {
			//               "additionalProperties": false,
			//               "description": "Prefix-level metrics configurations.",
			//               "properties": {
			//                 "StorageMetrics": {
			//                   "additionalProperties": false,
			//                   "properties": {
			//                     "IsEnabled": {
			//                       "description": "Specifies whether prefix-level storage metrics are enabled or disabled.",
			//                       "type": "boolean"
			//                     },
			//                     "SelectionCriteria": {
			//                       "additionalProperties": false,
			//                       "description": "Selection criteria for prefix-level metrics.",
			//                       "properties": {
			//                         "Delimiter": {
			//                           "description": "Delimiter to divide S3 key into hierarchy of prefixes.",
			//                           "type": "string"
			//                         },
			//                         "MaxDepth": {
			//                           "description": "Max depth of prefixes of S3 key that Amazon S3 Storage Lens will analyze.",
			//                           "type": "integer"
			//                         },
			//                         "MinStorageBytesPercentage": {
			//                           "description": "The minimum storage bytes threshold for the prefixes to be included in the analysis.",
			//                           "type": "number"
			//                         }
			//                       },
			//                       "type": "object"
			//                     }
			//                   },
			//                   "type": "object"
			//                 }
			//               },
			//               "required": [
			//                 "StorageMetrics"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "BucketLevel"
			//       ],
			//       "type": "object"
			//     },
			//     "AwsOrg": {
			//       "additionalProperties": false,
			//       "description": "The AWS Organizations ARN to use in the Amazon S3 Storage Lens configuration.",
			//       "properties": {
			//         "Arn": {
			//           "description": "The Amazon Resource Name (ARN) of the specified resource.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Arn"
			//       ],
			//       "type": "object"
			//     },
			//     "DataExport": {
			//       "additionalProperties": false,
			//       "description": "Specifies how Amazon S3 Storage Lens metrics should be exported.",
			//       "properties": {
			//         "S3BucketDestination": {
			//           "additionalProperties": false,
			//           "description": "S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.",
			//           "properties": {
			//             "AccountId": {
			//               "description": "The AWS account ID that owns the destination S3 bucket.",
			//               "type": "string"
			//             },
			//             "Arn": {
			//               "description": "The ARN of the bucket to which Amazon S3 Storage Lens exports will be placed.",
			//               "type": "string"
			//             },
			//             "Encryption": {
			//               "description": "Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).",
			//               "oneOf": [
			//                 {
			//                   "required": [
			//                     "SSES3"
			//                   ]
			//                 },
			//                 {
			//                   "required": [
			//                     "SSEKMS"
			//                   ]
			//                 }
			//               ],
			//               "type": "object"
			//             },
			//             "Format": {
			//               "description": "Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.",
			//               "enum": [
			//                 "CSV",
			//                 "Parquet"
			//               ],
			//               "type": "string"
			//             },
			//             "OutputSchemaVersion": {
			//               "description": "The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.",
			//               "enum": [
			//                 "V_1"
			//               ],
			//               "type": "string"
			//             },
			//             "Prefix": {
			//               "description": "The prefix to use for Amazon S3 Storage Lens export.",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "OutputSchemaVersion",
			//             "Format",
			//             "AccountId",
			//             "Arn"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "S3BucketDestination"
			//       ],
			//       "type": "object"
			//     },
			//     "Exclude": {
			//       "additionalProperties": false,
			//       "description": "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
			//       "properties": {
			//         "Buckets": {
			//           "insertionOrder": false,
			//           "items": {
			//             "description": "The Amazon Resource Name (ARN) of the specified resource.",
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         },
			//         "Regions": {
			//           "insertionOrder": false,
			//           "items": {
			//             "description": "An AWS Region.",
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Id": {
			//       "description": "The ID that identifies the Amazon S3 Storage Lens configuration.",
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "Include": {
			//       "additionalProperties": false,
			//       "description": "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
			//       "properties": {
			//         "Buckets": {
			//           "insertionOrder": false,
			//           "items": {
			//             "description": "The Amazon Resource Name (ARN) of the specified resource.",
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         },
			//         "Regions": {
			//           "insertionOrder": false,
			//           "items": {
			//             "description": "An AWS Region.",
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "IsEnabled": {
			//       "description": "Specifies whether the Amazon S3 Storage Lens configuration is enabled or disabled.",
			//       "type": "boolean"
			//     },
			//     "StorageLensArn": {
			//       "description": "The ARN for the Amazon S3 Storage Lens configuration.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Id",
			//     "AccountLevel",
			//     "IsEnabled"
			//   ],
			//   "type": "object"
			// }
			Description: "Specifies the details of Amazon S3 Storage Lens configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"account_level": {
						// Property: AccountLevel
						Description: "Account-level metrics configurations.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"activity_metrics": {
									// Property: ActivityMetrics
									Description: "Enables activity metrics.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"is_enabled": {
												// Property: IsEnabled
												Description: "Specifies whether activity metrics are enabled or disabled.",
												Type:        types.BoolType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"bucket_level": {
									// Property: BucketLevel
									Description: "Bucket-level metrics configurations.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"activity_metrics": {
												// Property: ActivityMetrics
												Description: "Enables activity metrics.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"is_enabled": {
															// Property: IsEnabled
															Description: "Specifies whether activity metrics are enabled or disabled.",
															Type:        types.BoolType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
											"prefix_level": {
												// Property: PrefixLevel
												Description: "Prefix-level metrics configurations.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"storage_metrics": {
															// Property: StorageMetrics
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"is_enabled": {
																		// Property: IsEnabled
																		Description: "Specifies whether prefix-level storage metrics are enabled or disabled.",
																		Type:        types.BoolType,
																		Computed:    true,
																	},
																	"selection_criteria": {
																		// Property: SelectionCriteria
																		Description: "Selection criteria for prefix-level metrics.",
																		Attributes: tfsdk.SingleNestedAttributes(
																			map[string]tfsdk.Attribute{
																				"delimiter": {
																					// Property: Delimiter
																					Description: "Delimiter to divide S3 key into hierarchy of prefixes.",
																					Type:        types.StringType,
																					Computed:    true,
																				},
																				"max_depth": {
																					// Property: MaxDepth
																					Description: "Max depth of prefixes of S3 key that Amazon S3 Storage Lens will analyze.",
																					Type:        types.NumberType,
																					Computed:    true,
																				},
																				"min_storage_bytes_percentage": {
																					// Property: MinStorageBytesPercentage
																					Description: "The minimum storage bytes threshold for the prefixes to be included in the analysis.",
																					Type:        types.NumberType,
																					Computed:    true,
																				},
																			},
																		),
																		Computed: true,
																	},
																},
															),
															Computed: true,
														},
													},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"aws_org": {
						// Property: AwsOrg
						Description: "The AWS Organizations ARN to use in the Amazon S3 Storage Lens configuration.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Description: "The Amazon Resource Name (ARN) of the specified resource.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"data_export": {
						// Property: DataExport
						Description: "Specifies how Amazon S3 Storage Lens metrics should be exported.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"s3_bucket_destination": {
									// Property: S3BucketDestination
									Description: "S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"account_id": {
												// Property: AccountId
												Description: "The AWS account ID that owns the destination S3 bucket.",
												Type:        types.StringType,
												Computed:    true,
											},
											"arn": {
												// Property: Arn
												Description: "The ARN of the bucket to which Amazon S3 Storage Lens exports will be placed.",
												Type:        types.StringType,
												Computed:    true,
											},
											"encryption": {
												// Property: Encryption
												Description: "Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).",
												Type:        types.MapType{ElemType: types.StringType},
												Computed:    true,
											},
											"format": {
												// Property: Format
												Description: "Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.",
												Type:        types.StringType,
												Computed:    true,
											},
											"output_schema_version": {
												// Property: OutputSchemaVersion
												Description: "The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.",
												Type:        types.StringType,
												Computed:    true,
											},
											"prefix": {
												// Property: Prefix
												Description: "The prefix to use for Amazon S3 Storage Lens export.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"exclude": {
						// Property: Exclude
						Description: "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"buckets": {
									// Property: Buckets
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
								"regions": {
									// Property: Regions
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"id": {
						// Property: Id
						Description: "The ID that identifies the Amazon S3 Storage Lens configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"include": {
						// Property: Include
						Description: "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"buckets": {
									// Property: Buckets
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
								"regions": {
									// Property: Regions
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"is_enabled": {
						// Property: IsEnabled
						Description: "Specifies whether the Amazon S3 Storage Lens configuration is enabled or disabled.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"storage_lens_arn": {
						// Property: StorageLensArn
						Description: "The ARN for the Amazon S3 Storage Lens configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
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
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.",
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::S3::StorageLens",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::StorageLens").WithTerraformTypeName("awscc_s3_storage_lens")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                   "AccountId",
		"account_level":                "AccountLevel",
		"activity_metrics":             "ActivityMetrics",
		"arn":                          "Arn",
		"aws_org":                      "AwsOrg",
		"bucket_level":                 "BucketLevel",
		"buckets":                      "Buckets",
		"data_export":                  "DataExport",
		"delimiter":                    "Delimiter",
		"encryption":                   "Encryption",
		"exclude":                      "Exclude",
		"format":                       "Format",
		"id":                           "Id",
		"include":                      "Include",
		"is_enabled":                   "IsEnabled",
		"key":                          "Key",
		"max_depth":                    "MaxDepth",
		"min_storage_bytes_percentage": "MinStorageBytesPercentage",
		"output_schema_version":        "OutputSchemaVersion",
		"prefix":                       "Prefix",
		"prefix_level":                 "PrefixLevel",
		"regions":                      "Regions",
		"s3_bucket_destination":        "S3BucketDestination",
		"selection_criteria":           "SelectionCriteria",
		"storage_lens_arn":             "StorageLensArn",
		"storage_lens_configuration":   "StorageLensConfiguration",
		"storage_metrics":              "StorageMetrics",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
