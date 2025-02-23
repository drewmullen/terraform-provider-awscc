---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_topic_rule Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IoT::TopicRule
---

# awscc_iot_topic_rule (Resource)

Resource Type definition for AWS::IoT::TopicRule



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `topic_rule_payload` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload))

### Optional

- `rule_name` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--topic_rule_payload"></a>
### Nested Schema for `topic_rule_payload`

Required:

- `actions` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--actions))
- `aws_iot_sql_version` (String)
- `description` (String)
- `error_action` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action))
- `rule_disabled` (Boolean)
- `sql` (String)

<a id="nestedatt--topic_rule_payload--actions"></a>
### Nested Schema for `topic_rule_payload.actions`

Required:

- `cloudwatch_alarm` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--cloudwatch_alarm))
- `cloudwatch_logs` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--cloudwatch_logs))
- `cloudwatch_metric` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--cloudwatch_metric))
- `dynamo_d_bv_2` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--dynamo_d_bv_2))
- `dynamo_db` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--dynamo_db))
- `elasticsearch` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--elasticsearch))
- `firehose` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--firehose))
- `http` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--http))
- `iot_analytics` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--iot_analytics))
- `iot_events` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--iot_events))
- `iot_site_wise` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--iot_site_wise))
- `kafka` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--kafka))
- `kinesis` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--kinesis))
- `lambda` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--lambda))
- `open_search` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--open_search))
- `republish` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--republish))
- `s3` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--s3))
- `sns` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--sns))
- `sqs` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--sqs))
- `step_functions` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--step_functions))
- `timestream` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--timestream))

<a id="nestedatt--topic_rule_payload--actions--cloudwatch_alarm"></a>
### Nested Schema for `topic_rule_payload.actions.cloudwatch_alarm`

Required:

- `alarm_name` (String)
- `role_arn` (String)
- `state_reason` (String)
- `state_value` (String)


<a id="nestedatt--topic_rule_payload--actions--cloudwatch_logs"></a>
### Nested Schema for `topic_rule_payload.actions.cloudwatch_logs`

Required:

- `log_group_name` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--actions--cloudwatch_metric"></a>
### Nested Schema for `topic_rule_payload.actions.cloudwatch_metric`

Required:

- `metric_name` (String)
- `metric_namespace` (String)
- `metric_timestamp` (String)
- `metric_unit` (String)
- `metric_value` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--actions--dynamo_d_bv_2"></a>
### Nested Schema for `topic_rule_payload.actions.dynamo_d_bv_2`

Required:

- `put_item` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--dynamo_d_bv_2--put_item))
- `role_arn` (String)

<a id="nestedatt--topic_rule_payload--actions--dynamo_d_bv_2--put_item"></a>
### Nested Schema for `topic_rule_payload.actions.dynamo_d_bv_2.role_arn`

Required:

- `table_name` (String)



<a id="nestedatt--topic_rule_payload--actions--dynamo_db"></a>
### Nested Schema for `topic_rule_payload.actions.dynamo_db`

Required:

- `hash_key_field` (String)
- `hash_key_type` (String)
- `hash_key_value` (String)
- `payload_field` (String)
- `range_key_field` (String)
- `range_key_type` (String)
- `range_key_value` (String)
- `role_arn` (String)
- `table_name` (String)


<a id="nestedatt--topic_rule_payload--actions--elasticsearch"></a>
### Nested Schema for `topic_rule_payload.actions.elasticsearch`

Required:

- `endpoint` (String)
- `id` (String) The ID of this resource.
- `index` (String)
- `role_arn` (String)
- `type` (String)


<a id="nestedatt--topic_rule_payload--actions--firehose"></a>
### Nested Schema for `topic_rule_payload.actions.firehose`

Required:

- `batch_mode` (Boolean)
- `delivery_stream_name` (String)
- `role_arn` (String)
- `separator` (String)


<a id="nestedatt--topic_rule_payload--actions--http"></a>
### Nested Schema for `topic_rule_payload.actions.http`

Required:

- `auth` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--http--auth))
- `confirmation_url` (String)
- `headers` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--http--headers))
- `url` (String)

<a id="nestedatt--topic_rule_payload--actions--http--auth"></a>
### Nested Schema for `topic_rule_payload.actions.http.url`

Required:

- `sigv_4` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--http--url--sigv_4))

<a id="nestedatt--topic_rule_payload--actions--http--url--sigv_4"></a>
### Nested Schema for `topic_rule_payload.actions.http.url.sigv_4`

Required:

- `role_arn` (String)
- `service_name` (String)
- `signing_region` (String)



<a id="nestedatt--topic_rule_payload--actions--http--headers"></a>
### Nested Schema for `topic_rule_payload.actions.http.url`

Required:

- `key` (String)
- `value` (String)



<a id="nestedatt--topic_rule_payload--actions--iot_analytics"></a>
### Nested Schema for `topic_rule_payload.actions.iot_analytics`

Required:

- `batch_mode` (Boolean)
- `channel_name` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--actions--iot_events"></a>
### Nested Schema for `topic_rule_payload.actions.iot_events`

Required:

- `batch_mode` (Boolean)
- `input_name` (String)
- `message_id` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--actions--iot_site_wise"></a>
### Nested Schema for `topic_rule_payload.actions.iot_site_wise`

Required:

- `put_asset_property_value_entries` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--iot_site_wise--put_asset_property_value_entries))
- `role_arn` (String)

<a id="nestedatt--topic_rule_payload--actions--iot_site_wise--put_asset_property_value_entries"></a>
### Nested Schema for `topic_rule_payload.actions.iot_site_wise.role_arn`

Required:

- `asset_id` (String)
- `entry_id` (String)
- `property_alias` (String)
- `property_id` (String)
- `property_values` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--iot_site_wise--role_arn--property_values))

<a id="nestedatt--topic_rule_payload--actions--iot_site_wise--role_arn--property_values"></a>
### Nested Schema for `topic_rule_payload.actions.iot_site_wise.role_arn.property_values`

Required:

- `quality` (String)
- `timestamp` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--iot_site_wise--role_arn--property_values--timestamp))
- `value` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--iot_site_wise--role_arn--property_values--value))

<a id="nestedatt--topic_rule_payload--actions--iot_site_wise--role_arn--property_values--timestamp"></a>
### Nested Schema for `topic_rule_payload.actions.iot_site_wise.role_arn.property_values.value`

Required:

- `offset_in_nanos` (String)
- `time_in_seconds` (String)


<a id="nestedatt--topic_rule_payload--actions--iot_site_wise--role_arn--property_values--value"></a>
### Nested Schema for `topic_rule_payload.actions.iot_site_wise.role_arn.property_values.value`

Required:

- `boolean_value` (String)
- `double_value` (String)
- `integer_value` (String)
- `string_value` (String)





<a id="nestedatt--topic_rule_payload--actions--kafka"></a>
### Nested Schema for `topic_rule_payload.actions.kafka`

Required:

- `client_properties` (Map of String)
- `destination_arn` (String)
- `key` (String)
- `partition` (String)
- `topic` (String)


<a id="nestedatt--topic_rule_payload--actions--kinesis"></a>
### Nested Schema for `topic_rule_payload.actions.kinesis`

Required:

- `partition_key` (String)
- `role_arn` (String)
- `stream_name` (String)


<a id="nestedatt--topic_rule_payload--actions--lambda"></a>
### Nested Schema for `topic_rule_payload.actions.lambda`

Required:

- `function_arn` (String)


<a id="nestedatt--topic_rule_payload--actions--open_search"></a>
### Nested Schema for `topic_rule_payload.actions.open_search`

Required:

- `endpoint` (String)
- `id` (String) The ID of this resource.
- `index` (String)
- `role_arn` (String)
- `type` (String)


<a id="nestedatt--topic_rule_payload--actions--republish"></a>
### Nested Schema for `topic_rule_payload.actions.republish`

Required:

- `qos` (Number)
- `role_arn` (String)
- `topic` (String)


<a id="nestedatt--topic_rule_payload--actions--s3"></a>
### Nested Schema for `topic_rule_payload.actions.s3`

Required:

- `bucket_name` (String)
- `canned_acl` (String)
- `key` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--actions--sns"></a>
### Nested Schema for `topic_rule_payload.actions.sns`

Required:

- `message_format` (String)
- `role_arn` (String)
- `target_arn` (String)


<a id="nestedatt--topic_rule_payload--actions--sqs"></a>
### Nested Schema for `topic_rule_payload.actions.sqs`

Required:

- `queue_url` (String)
- `role_arn` (String)
- `use_base_64` (Boolean)


<a id="nestedatt--topic_rule_payload--actions--step_functions"></a>
### Nested Schema for `topic_rule_payload.actions.step_functions`

Required:

- `execution_name_prefix` (String)
- `role_arn` (String)
- `state_machine_name` (String)


<a id="nestedatt--topic_rule_payload--actions--timestream"></a>
### Nested Schema for `topic_rule_payload.actions.timestream`

Required:

- `batch_mode` (Boolean)
- `database_name` (String)
- `dimensions` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--timestream--dimensions))
- `role_arn` (String)
- `table_name` (String)
- `timestamp` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--actions--timestream--timestamp))

<a id="nestedatt--topic_rule_payload--actions--timestream--dimensions"></a>
### Nested Schema for `topic_rule_payload.actions.timestream.timestamp`

Required:

- `name` (String)
- `value` (String)


<a id="nestedatt--topic_rule_payload--actions--timestream--timestamp"></a>
### Nested Schema for `topic_rule_payload.actions.timestream.timestamp`

Required:

- `unit` (String)
- `value` (String)




<a id="nestedatt--topic_rule_payload--error_action"></a>
### Nested Schema for `topic_rule_payload.error_action`

Required:

- `cloudwatch_alarm` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--cloudwatch_alarm))
- `cloudwatch_logs` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--cloudwatch_logs))
- `cloudwatch_metric` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--cloudwatch_metric))
- `dynamo_d_bv_2` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--dynamo_d_bv_2))
- `dynamo_db` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--dynamo_db))
- `elasticsearch` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--elasticsearch))
- `firehose` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--firehose))
- `http` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--http))
- `iot_analytics` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--iot_analytics))
- `iot_events` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--iot_events))
- `iot_site_wise` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--iot_site_wise))
- `kafka` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--kafka))
- `kinesis` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--kinesis))
- `lambda` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--lambda))
- `open_search` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--open_search))
- `republish` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--republish))
- `s3` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--s3))
- `sns` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--sns))
- `sqs` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--sqs))
- `step_functions` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--step_functions))
- `timestream` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--timestream))

<a id="nestedatt--topic_rule_payload--error_action--cloudwatch_alarm"></a>
### Nested Schema for `topic_rule_payload.error_action.cloudwatch_alarm`

Required:

- `alarm_name` (String)
- `role_arn` (String)
- `state_reason` (String)
- `state_value` (String)


<a id="nestedatt--topic_rule_payload--error_action--cloudwatch_logs"></a>
### Nested Schema for `topic_rule_payload.error_action.cloudwatch_logs`

Required:

- `log_group_name` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--error_action--cloudwatch_metric"></a>
### Nested Schema for `topic_rule_payload.error_action.cloudwatch_metric`

Required:

- `metric_name` (String)
- `metric_namespace` (String)
- `metric_timestamp` (String)
- `metric_unit` (String)
- `metric_value` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--error_action--dynamo_d_bv_2"></a>
### Nested Schema for `topic_rule_payload.error_action.dynamo_d_bv_2`

Required:

- `put_item` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--dynamo_d_bv_2--put_item))
- `role_arn` (String)

<a id="nestedatt--topic_rule_payload--error_action--dynamo_d_bv_2--put_item"></a>
### Nested Schema for `topic_rule_payload.error_action.dynamo_d_bv_2.role_arn`

Required:

- `table_name` (String)



<a id="nestedatt--topic_rule_payload--error_action--dynamo_db"></a>
### Nested Schema for `topic_rule_payload.error_action.dynamo_db`

Required:

- `hash_key_field` (String)
- `hash_key_type` (String)
- `hash_key_value` (String)
- `payload_field` (String)
- `range_key_field` (String)
- `range_key_type` (String)
- `range_key_value` (String)
- `role_arn` (String)
- `table_name` (String)


<a id="nestedatt--topic_rule_payload--error_action--elasticsearch"></a>
### Nested Schema for `topic_rule_payload.error_action.elasticsearch`

Required:

- `endpoint` (String)
- `id` (String) The ID of this resource.
- `index` (String)
- `role_arn` (String)
- `type` (String)


<a id="nestedatt--topic_rule_payload--error_action--firehose"></a>
### Nested Schema for `topic_rule_payload.error_action.firehose`

Required:

- `batch_mode` (Boolean)
- `delivery_stream_name` (String)
- `role_arn` (String)
- `separator` (String)


<a id="nestedatt--topic_rule_payload--error_action--http"></a>
### Nested Schema for `topic_rule_payload.error_action.http`

Required:

- `auth` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--http--auth))
- `confirmation_url` (String)
- `headers` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--http--headers))
- `url` (String)

<a id="nestedatt--topic_rule_payload--error_action--http--auth"></a>
### Nested Schema for `topic_rule_payload.error_action.http.url`

Required:

- `sigv_4` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--http--url--sigv_4))

<a id="nestedatt--topic_rule_payload--error_action--http--url--sigv_4"></a>
### Nested Schema for `topic_rule_payload.error_action.http.url.sigv_4`

Required:

- `role_arn` (String)
- `service_name` (String)
- `signing_region` (String)



<a id="nestedatt--topic_rule_payload--error_action--http--headers"></a>
### Nested Schema for `topic_rule_payload.error_action.http.url`

Required:

- `key` (String)
- `value` (String)



<a id="nestedatt--topic_rule_payload--error_action--iot_analytics"></a>
### Nested Schema for `topic_rule_payload.error_action.iot_analytics`

Required:

- `batch_mode` (Boolean)
- `channel_name` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--error_action--iot_events"></a>
### Nested Schema for `topic_rule_payload.error_action.iot_events`

Required:

- `batch_mode` (Boolean)
- `input_name` (String)
- `message_id` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--error_action--iot_site_wise"></a>
### Nested Schema for `topic_rule_payload.error_action.iot_site_wise`

Required:

- `put_asset_property_value_entries` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--iot_site_wise--put_asset_property_value_entries))
- `role_arn` (String)

<a id="nestedatt--topic_rule_payload--error_action--iot_site_wise--put_asset_property_value_entries"></a>
### Nested Schema for `topic_rule_payload.error_action.iot_site_wise.role_arn`

Required:

- `asset_id` (String)
- `entry_id` (String)
- `property_alias` (String)
- `property_id` (String)
- `property_values` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--iot_site_wise--role_arn--property_values))

<a id="nestedatt--topic_rule_payload--error_action--iot_site_wise--role_arn--property_values"></a>
### Nested Schema for `topic_rule_payload.error_action.iot_site_wise.role_arn.property_values`

Required:

- `quality` (String)
- `timestamp` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--iot_site_wise--role_arn--property_values--timestamp))
- `value` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--iot_site_wise--role_arn--property_values--value))

<a id="nestedatt--topic_rule_payload--error_action--iot_site_wise--role_arn--property_values--timestamp"></a>
### Nested Schema for `topic_rule_payload.error_action.iot_site_wise.role_arn.property_values.value`

Required:

- `offset_in_nanos` (String)
- `time_in_seconds` (String)


<a id="nestedatt--topic_rule_payload--error_action--iot_site_wise--role_arn--property_values--value"></a>
### Nested Schema for `topic_rule_payload.error_action.iot_site_wise.role_arn.property_values.value`

Required:

- `boolean_value` (String)
- `double_value` (String)
- `integer_value` (String)
- `string_value` (String)





<a id="nestedatt--topic_rule_payload--error_action--kafka"></a>
### Nested Schema for `topic_rule_payload.error_action.kafka`

Required:

- `client_properties` (Map of String)
- `destination_arn` (String)
- `key` (String)
- `partition` (String)
- `topic` (String)


<a id="nestedatt--topic_rule_payload--error_action--kinesis"></a>
### Nested Schema for `topic_rule_payload.error_action.kinesis`

Required:

- `partition_key` (String)
- `role_arn` (String)
- `stream_name` (String)


<a id="nestedatt--topic_rule_payload--error_action--lambda"></a>
### Nested Schema for `topic_rule_payload.error_action.lambda`

Required:

- `function_arn` (String)


<a id="nestedatt--topic_rule_payload--error_action--open_search"></a>
### Nested Schema for `topic_rule_payload.error_action.open_search`

Required:

- `endpoint` (String)
- `id` (String) The ID of this resource.
- `index` (String)
- `role_arn` (String)
- `type` (String)


<a id="nestedatt--topic_rule_payload--error_action--republish"></a>
### Nested Schema for `topic_rule_payload.error_action.republish`

Required:

- `qos` (Number)
- `role_arn` (String)
- `topic` (String)


<a id="nestedatt--topic_rule_payload--error_action--s3"></a>
### Nested Schema for `topic_rule_payload.error_action.s3`

Required:

- `bucket_name` (String)
- `canned_acl` (String)
- `key` (String)
- `role_arn` (String)


<a id="nestedatt--topic_rule_payload--error_action--sns"></a>
### Nested Schema for `topic_rule_payload.error_action.sns`

Required:

- `message_format` (String)
- `role_arn` (String)
- `target_arn` (String)


<a id="nestedatt--topic_rule_payload--error_action--sqs"></a>
### Nested Schema for `topic_rule_payload.error_action.sqs`

Required:

- `queue_url` (String)
- `role_arn` (String)
- `use_base_64` (Boolean)


<a id="nestedatt--topic_rule_payload--error_action--step_functions"></a>
### Nested Schema for `topic_rule_payload.error_action.step_functions`

Required:

- `execution_name_prefix` (String)
- `role_arn` (String)
- `state_machine_name` (String)


<a id="nestedatt--topic_rule_payload--error_action--timestream"></a>
### Nested Schema for `topic_rule_payload.error_action.timestream`

Required:

- `batch_mode` (Boolean)
- `database_name` (String)
- `dimensions` (Attributes List) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--timestream--dimensions))
- `role_arn` (String)
- `table_name` (String)
- `timestamp` (Attributes) (see [below for nested schema](#nestedatt--topic_rule_payload--error_action--timestream--timestamp))

<a id="nestedatt--topic_rule_payload--error_action--timestream--dimensions"></a>
### Nested Schema for `topic_rule_payload.error_action.timestream.timestamp`

Required:

- `name` (String)
- `value` (String)


<a id="nestedatt--topic_rule_payload--error_action--timestream--timestamp"></a>
### Nested Schema for `topic_rule_payload.error_action.timestream.timestamp`

Required:

- `unit` (String)
- `value` (String)





<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iot_topic_rule.example <resource ID>
```
