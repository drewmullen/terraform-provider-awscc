---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_route_table Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::RouteTable
---

# awscc_ec2_route_table (Resource)

Resource Type definition for AWS::EC2::RouteTable



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vpc_id` (String) The ID of the VPC.

### Optional

- `tags` (Attributes List) Any tags assigned to the route table. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `route_table_id` (String) The route table ID.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_route_table.example <resource ID>
```
