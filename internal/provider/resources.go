// Code generated by generators/schema/main.go; DO NOT EDIT.
//go:generate go run generators/resource/main.go -resource aws_logs_log_group -cfschema /Users/ewbankkit/src/github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation/schemas/us-west-2/aws-logs-loggroup.json -package logs -- ../aws/logs/log_group_gen.go

package provider

import (
	_ "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/aws/logs"
)
