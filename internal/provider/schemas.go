//go:generate go run generators/schema/main.go -config all_schemas.hcl -generated-code-root .. -import-path-root github.com/hashicorp/terraform-provider-aws-cloudapi/internal -- resources.go

package provider
