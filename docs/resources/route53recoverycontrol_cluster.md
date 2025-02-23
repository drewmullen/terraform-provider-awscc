---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53recoverycontrol_cluster Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  AWS Route53 Recovery Control Cluster resource schema
---

# awscc_route53recoverycontrol_cluster (Resource)

AWS Route53 Recovery Control Cluster resource schema



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name` (String) Name of a Cluster. You can use any non-white space character in the name
- `tags` (Attributes List) A collection of tags associated with a resource (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `cluster_arn` (String) The Amazon Resource Name (ARN) of the cluster.
- `cluster_endpoints` (Attributes List) Endpoints for the cluster. (see [below for nested schema](#nestedatt--cluster_endpoints))
- `id` (String) Uniquely identifies the resource.
- `status` (String) Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--cluster_endpoints"></a>
### Nested Schema for `cluster_endpoints`

Read-Only:

- `endpoint` (String)
- `region` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_route53recoverycontrol_cluster.example <resource ID>
```
