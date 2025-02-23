---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ecs_task_definition Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Schema describing various properties for ECS TaskDefinition
---

# awscc_ecs_task_definition (Resource)

Resource Schema describing various properties for ECS TaskDefinition



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `container_definitions` (Attributes Set) (see [below for nested schema](#nestedatt--container_definitions))
- `cpu` (String)
- `ephemeral_storage` (Attributes) (see [below for nested schema](#nestedatt--ephemeral_storage))
- `execution_role_arn` (String)
- `family` (String)
- `inference_accelerators` (Attributes Set) (see [below for nested schema](#nestedatt--inference_accelerators))
- `ipc_mode` (String)
- `memory` (String)
- `network_mode` (String)
- `pid_mode` (String)
- `placement_constraints` (Attributes Set) (see [below for nested schema](#nestedatt--placement_constraints))
- `proxy_configuration` (Attributes) (see [below for nested schema](#nestedatt--proxy_configuration))
- `requires_compatibilities` (Set of String)
- `runtime_platform` (Attributes) (see [below for nested schema](#nestedatt--runtime_platform))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `task_role_arn` (String)
- `volumes` (Attributes Set) (see [below for nested schema](#nestedatt--volumes))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `task_definition_arn` (String) The Amazon Resource Name (ARN) of the Amazon ECS task definition

<a id="nestedatt--container_definitions"></a>
### Nested Schema for `container_definitions`

Optional:

- `command` (List of String)
- `cpu` (Number)
- `depends_on` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--depends_on))
- `disable_networking` (Boolean)
- `dns_search_domains` (List of String)
- `dns_servers` (List of String)
- `docker_labels` (Map of String)
- `docker_security_options` (List of String)
- `entry_point` (List of String)
- `environment` (Attributes List) The environment variables to pass to a container (see [below for nested schema](#nestedatt--container_definitions--environment))
- `environment_files` (Attributes List) The list of one or more files that contain the environment variables to pass to a container (see [below for nested schema](#nestedatt--container_definitions--environment_files))
- `essential` (Boolean)
- `extra_hosts` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--extra_hosts))
- `firelens_configuration` (Attributes) (see [below for nested schema](#nestedatt--container_definitions--firelens_configuration))
- `health_check` (Attributes) The health check command and associated configuration parameters for the container. (see [below for nested schema](#nestedatt--container_definitions--health_check))
- `hostname` (String)
- `image` (String) The image used to start a container. This string is passed directly to the Docker daemon.
- `interactive` (Boolean)
- `links` (Set of String)
- `linux_parameters` (Attributes) (see [below for nested schema](#nestedatt--container_definitions--linux_parameters))
- `log_configuration` (Attributes) (see [below for nested schema](#nestedatt--container_definitions--log_configuration))
- `memory` (Number) The amount (in MiB) of memory to present to the container. If your container attempts to exceed the memory specified here, the container is killed.
- `memory_reservation` (Number)
- `mount_points` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--mount_points))
- `name` (String) The name of a container. Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed
- `port_mappings` (Attributes Set) Port mappings allow containers to access ports on the host container instance to send or receive traffic. (see [below for nested schema](#nestedatt--container_definitions--port_mappings))
- `privileged` (Boolean)
- `pseudo_terminal` (Boolean)
- `readonly_root_filesystem` (Boolean)
- `repository_credentials` (Attributes) (see [below for nested schema](#nestedatt--container_definitions--repository_credentials))
- `resource_requirements` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--resource_requirements))
- `secrets` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--secrets))
- `start_timeout` (Number)
- `stop_timeout` (Number)
- `system_controls` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--system_controls))
- `ulimits` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--ulimits))
- `user` (String)
- `volumes_from` (Attributes Set) (see [below for nested schema](#nestedatt--container_definitions--volumes_from))
- `working_directory` (String)

<a id="nestedatt--container_definitions--depends_on"></a>
### Nested Schema for `container_definitions.depends_on`

Optional:

- `condition` (String)
- `container_name` (String)


<a id="nestedatt--container_definitions--environment"></a>
### Nested Schema for `container_definitions.environment`

Optional:

- `name` (String)
- `value` (String)


<a id="nestedatt--container_definitions--environment_files"></a>
### Nested Schema for `container_definitions.environment_files`

Optional:

- `type` (String)
- `value` (String)


<a id="nestedatt--container_definitions--extra_hosts"></a>
### Nested Schema for `container_definitions.extra_hosts`

Optional:

- `hostname` (String)
- `ip_address` (String)


<a id="nestedatt--container_definitions--firelens_configuration"></a>
### Nested Schema for `container_definitions.firelens_configuration`

Optional:

- `options` (Map of String)
- `type` (String)


<a id="nestedatt--container_definitions--health_check"></a>
### Nested Schema for `container_definitions.health_check`

Optional:

- `command` (List of String) A string array representing the command that the container runs to determine if it is healthy.
- `interval` (Number) The time period in seconds between each health check execution. You may specify between 5 and 300 seconds. The default value is 30 seconds.
- `retries` (Number) The number of times to retry a failed health check before the container is considered unhealthy. You may specify between 1 and 10 retries. The default value is three retries.
- `start_period` (Number) The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries. You may specify between 0 and 300 seconds. The startPeriod is disabled by default.
- `timeout` (Number) The time period in seconds to wait for a health check to succeed before it is considered a failure. You may specify between 2 and 60 seconds. The default value is 5 seconds.


<a id="nestedatt--container_definitions--linux_parameters"></a>
### Nested Schema for `container_definitions.linux_parameters`

Optional:

- `capabilities` (Attributes) (see [below for nested schema](#nestedatt--container_definitions--linux_parameters--capabilities))
- `devices` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--linux_parameters--devices))
- `init_process_enabled` (Boolean)
- `max_swap` (Number)
- `shared_memory_size` (Number)
- `swappiness` (Number)
- `tmpfs` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--linux_parameters--tmpfs))

<a id="nestedatt--container_definitions--linux_parameters--capabilities"></a>
### Nested Schema for `container_definitions.linux_parameters.capabilities`

Optional:

- `add` (List of String)
- `drop` (List of String)


<a id="nestedatt--container_definitions--linux_parameters--devices"></a>
### Nested Schema for `container_definitions.linux_parameters.devices`

Optional:

- `container_path` (String)
- `host_path` (String)
- `permissions` (Set of String)


<a id="nestedatt--container_definitions--linux_parameters--tmpfs"></a>
### Nested Schema for `container_definitions.linux_parameters.tmpfs`

Optional:

- `container_path` (String)
- `mount_options` (List of String)
- `size` (Number)



<a id="nestedatt--container_definitions--log_configuration"></a>
### Nested Schema for `container_definitions.log_configuration`

Optional:

- `log_driver` (String)
- `options` (Map of String)
- `secret_options` (Attributes List) (see [below for nested schema](#nestedatt--container_definitions--log_configuration--secret_options))

<a id="nestedatt--container_definitions--log_configuration--secret_options"></a>
### Nested Schema for `container_definitions.log_configuration.secret_options`

Optional:

- `name` (String)
- `value_from` (String)



<a id="nestedatt--container_definitions--mount_points"></a>
### Nested Schema for `container_definitions.mount_points`

Optional:

- `container_path` (String)
- `read_only` (Boolean)
- `source_volume` (String)


<a id="nestedatt--container_definitions--port_mappings"></a>
### Nested Schema for `container_definitions.port_mappings`

Optional:

- `container_port` (Number)
- `host_port` (Number)
- `protocol` (String)


<a id="nestedatt--container_definitions--repository_credentials"></a>
### Nested Schema for `container_definitions.repository_credentials`

Optional:

- `credentials_parameter` (String)


<a id="nestedatt--container_definitions--resource_requirements"></a>
### Nested Schema for `container_definitions.resource_requirements`

Optional:

- `type` (String)
- `value` (String)


<a id="nestedatt--container_definitions--secrets"></a>
### Nested Schema for `container_definitions.secrets`

Optional:

- `name` (String)
- `value_from` (String)


<a id="nestedatt--container_definitions--system_controls"></a>
### Nested Schema for `container_definitions.system_controls`

Optional:

- `namespace` (String)
- `value` (String)


<a id="nestedatt--container_definitions--ulimits"></a>
### Nested Schema for `container_definitions.ulimits`

Optional:

- `hard_limit` (Number)
- `name` (String)
- `soft_limit` (Number)


<a id="nestedatt--container_definitions--volumes_from"></a>
### Nested Schema for `container_definitions.volumes_from`

Optional:

- `read_only` (Boolean)
- `source_container` (String)



<a id="nestedatt--ephemeral_storage"></a>
### Nested Schema for `ephemeral_storage`

Optional:

- `size_in_gi_b` (Number)


<a id="nestedatt--inference_accelerators"></a>
### Nested Schema for `inference_accelerators`

Optional:

- `device_name` (String)
- `device_type` (String)


<a id="nestedatt--placement_constraints"></a>
### Nested Schema for `placement_constraints`

Optional:

- `expression` (String)
- `type` (String)


<a id="nestedatt--proxy_configuration"></a>
### Nested Schema for `proxy_configuration`

Optional:

- `container_name` (String)
- `proxy_configuration_properties` (Attributes Set) (see [below for nested schema](#nestedatt--proxy_configuration--proxy_configuration_properties))
- `type` (String)

<a id="nestedatt--proxy_configuration--proxy_configuration_properties"></a>
### Nested Schema for `proxy_configuration.proxy_configuration_properties`

Optional:

- `name` (String)
- `value` (String)



<a id="nestedatt--runtime_platform"></a>
### Nested Schema for `runtime_platform`

Optional:

- `cpu_architecture` (String)
- `operating_system_family` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--volumes"></a>
### Nested Schema for `volumes`

Optional:

- `docker_volume_configuration` (Attributes) (see [below for nested schema](#nestedatt--volumes--docker_volume_configuration))
- `efs_volume_configuration` (Attributes) (see [below for nested schema](#nestedatt--volumes--efs_volume_configuration))
- `host` (Attributes) (see [below for nested schema](#nestedatt--volumes--host))
- `name` (String)

<a id="nestedatt--volumes--docker_volume_configuration"></a>
### Nested Schema for `volumes.docker_volume_configuration`

Optional:

- `autoprovision` (Boolean)
- `driver` (String)
- `driver_opts` (Map of String)
- `labels` (Map of String)
- `scope` (String)


<a id="nestedatt--volumes--efs_volume_configuration"></a>
### Nested Schema for `volumes.efs_volume_configuration`

Optional:

- `authorization_config` (Attributes) (see [below for nested schema](#nestedatt--volumes--efs_volume_configuration--authorization_config))
- `filesystem_id` (String)
- `root_directory` (String)
- `transit_encryption` (String)
- `transit_encryption_port` (Number)

<a id="nestedatt--volumes--efs_volume_configuration--authorization_config"></a>
### Nested Schema for `volumes.efs_volume_configuration.authorization_config`

Optional:

- `access_point_id` (String)
- `iam` (String)



<a id="nestedatt--volumes--host"></a>
### Nested Schema for `volumes.host`

Optional:

- `source_path` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ecs_task_definition.example <resource ID>
```
