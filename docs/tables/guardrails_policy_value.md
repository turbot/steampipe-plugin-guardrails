---
title: "Steampipe Table: guardrails_policy_value - Query Guardrails Policy Values using SQL"
description: "Allows users to query Guardrails Policy Values, providing insights into policy configurations and associated metadata."
folder: "Policy"
---

# Table: guardrails_policy_value - Query Guardrails Policy Values using SQL

Guardrails is a policy as code service that enables users to manage and enforce policies across their infrastructure. It allows users to define policy values that dictate the desired state of system configuration. These policy values can be used to ensure compliance, enforce security measures, and manage resource configurations.

## Table Usage Guide

The `guardrails_policy_value` table provides insights into Guardrails policy values. As a system administrator or a compliance manager, explore policy-specific details through this table, including policy value, associated metadata, and the desired state of system configuration. Utilize it to uncover information about policy values, such as those associated with specific compliance requirements, and to verify the desired state of system configuration.

**Important Notes**
- When querying this table, we recommend using at least one of these columns (usually in the `where` clause):
  - `state`
  - `policy_type_id`
  - `resource_type_id`
  - `resource_type_uri`
  - `filter`

## Examples

### List policy values by policy type ID
Explore specific policy values based on their type ID to understand their status, defaults, and calculations. This can help in analyzing and managing guardrail policies effectively.

```sql+postgres
select
  id,
  state,
  is_default,
  is_calculated,
  policy_type_id,
  type_mod_uri
from
  guardrails_policy_value
where
  policy_type_id = 221505068398189;
```

```sql+sqlite
select
  id,
  state,
  is_default,
  is_calculated,
  policy_type_id,
  type_mod_uri
from
  guardrails_policy_value
where
  policy_type_id = 221505068398189;
```

### List policy values by resource ID
Identify the status and types of policy values associated with a specific resource. This can aid in understanding the configuration and management of that resource.

```sql+postgres
select
  id,
  state,
  is_default,
  is_calculated,
  resource_id,
  type_mod_uri
from
  guardrails_policy_value
where
  resource_id = 161587219904115;
```

```sql+sqlite
select
  id,
  state,
  is_default,
  is_calculated,
  resource_id,
  type_mod_uri
from
  guardrails_policy_value
where
  resource_id = 161587219904115;
```

### List non-default calculated policy values
Analyze the settings to understand the non-standard calculated policy values. This is beneficial in identifying any deviations from the default settings, which could potentially impact resource management and security.

```sql+postgres
select
  id,
  state,
  is_default,
  is_calculated,
  resource_type_id,
  type_mod_uri
from
  guardrails_policy_value
where
  is_calculated and not is_default;
```

```sql+sqlite
select
  id,
  state,
  is_default,
  is_calculated,
  resource_type_id,
  type_mod_uri
from
  guardrails_policy_value
where
  is_calculated = 1 and not is_default = 1;
```

### Filter policy values using Turbot filter syntax
Analyze the settings to understand the status of different policy values, specifically those that are currently in an 'ok' state. This allows for efficient monitoring and management of system policies.

```sql+postgres
select
  id,
  state,
  is_default,
  is_calculated,
  policy_type_id,
  resource_id,
  resource_type_id
from
  guardrails_policy_value
where
  filter = 'state:ok';
```

```sql+sqlite
select
  id,
  state,
  is_default,
  is_calculated,
  policy_type_id,
  resource_id,
  resource_type_id
from
  guardrails_policy_value
where
  filter = 'state:ok';
```