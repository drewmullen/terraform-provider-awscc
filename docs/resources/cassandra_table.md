---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cassandra_table Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Cassandra::Table
---

# awscc_cassandra_table (Resource)

Resource schema for AWS::Cassandra::Table



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `keyspace_name` (String) Name for Cassandra keyspace
- `partition_key_columns` (Attributes List) Partition key columns of the table (see [below for nested schema](#nestedatt--partition_key_columns))

### Optional

- `billing_mode` (Attributes) (see [below for nested schema](#nestedatt--billing_mode))
- `clustering_key_columns` (Attributes List) Clustering key columns of the table (see [below for nested schema](#nestedatt--clustering_key_columns))
- `default_time_to_live` (Number) Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
- `encryption_specification` (Attributes) Represents the settings used to enable server-side encryption (see [below for nested schema](#nestedatt--encryption_specification))
- `point_in_time_recovery_enabled` (Boolean) Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
- `regular_columns` (Attributes Set) Non-key columns of the table (see [below for nested schema](#nestedatt--regular_columns))
- `table_name` (String) Name for Cassandra table
- `tags` (Attributes List) An array of key-value pairs to apply to this resource (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--partition_key_columns"></a>
### Nested Schema for `partition_key_columns`

Required:

- `column_name` (String)
- `column_type` (String)


<a id="nestedatt--billing_mode"></a>
### Nested Schema for `billing_mode`

Optional:

- `mode` (String) Capacity mode for the specified table
- `provisioned_throughput` (Attributes) Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits (see [below for nested schema](#nestedatt--billing_mode--provisioned_throughput))

<a id="nestedatt--billing_mode--provisioned_throughput"></a>
### Nested Schema for `billing_mode.provisioned_throughput`

Optional:

- `read_capacity_units` (Number)
- `write_capacity_units` (Number)



<a id="nestedatt--clustering_key_columns"></a>
### Nested Schema for `clustering_key_columns`

Optional:

- `column` (Attributes) (see [below for nested schema](#nestedatt--clustering_key_columns--column))
- `order_by` (String)

<a id="nestedatt--clustering_key_columns--column"></a>
### Nested Schema for `clustering_key_columns.column`

Optional:

- `column_name` (String)
- `column_type` (String)



<a id="nestedatt--encryption_specification"></a>
### Nested Schema for `encryption_specification`

Optional:

- `encryption_type` (String) Server-side encryption type
- `kms_key_identifier` (String) The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN.


<a id="nestedatt--regular_columns"></a>
### Nested Schema for `regular_columns`

Optional:

- `column_name` (String)
- `column_type` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cassandra_table.example <resource ID>
```
