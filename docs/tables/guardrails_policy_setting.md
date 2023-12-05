---
title: "Steampipe Table: guardrails_policy_setting - Query Guardrails Policy Settings using SQL"
description: "Allows users to query Guardrails Policy Settings, specifically providing insights into policy configurations and their associated metadata."
---

# Table: guardrails_policy_setting - Query Guardrails Policy Settings using SQL

Guardrails is a policy-as-code service that helps you manage and enforce your cloud security, compliance, and operational policies. It allows you to define your own policies, or use pre-built ones, to continuously monitor and remediate security and compliance issues. Guardrails Policy Settings are a key component of this service, defining the specific configurations for each policy.

## Table Usage Guide

The `guardrails_policy_setting` table provides insights into policy settings within Guardrails. As a Security Engineer, explore policy-specific details through this table, including policy configurations, associated metadata, and the status of each policy. Utilize it to uncover information about policies, such as those which are currently active, the specific configurations of each policy, and the potential impact of these policies on your cloud resources.

**Important Notes**
- When querying this table, we recommend using at least one of these columns (usually in the `where` clause):
  - `id`
  - `resource_id`
  - `exception`
  - `orphan`
  - `policy_type_id`
  - `policy_type_uri`
  - `filter`

## Examples

### Find all policy settings that are exceptions to another policy
Discover the segments that constitute exceptions to other policies, enabling you to assess the elements within your system that deviate from the standard protocol. This is useful in identifying instances where modifications may be necessary to ensure consistency and compliance.

```sql
select
  policy_type_uri,
  resource_id,
  is_calculated,
  exception,
  value
from
  guardrails_policy_setting
where
  exception;
```

### List policy settings with full resource and policy type information
Explore the configuration of policy settings, including their associated resources and policy types. This can help in identifying any exceptions and understanding the calculated values, aiding in effective policy management.

```sql
select
  r.trunk_title as resource,
  pt.trunk_title as policy_type,
  ps.value,
  ps.is_calculated,
  ps.exception
from
  guardrails_policy_setting as ps
  left join guardrails_policy_type as pt on pt.id = ps.policy_type_id
  left join guardrails_resource as r on r.id = ps.resource_id;
```

### All policy settings set on a given resource
Explore which policy settings are applied to a specific resource to understand the current configuration and its calculated status. This can aid in identifying potential security gaps or compliance issues.

```sql
select
  r.trunk_title as resource,
  ps.resource_id,
  pt.trunk_title as policy_type,
  ps.value,
  ps.is_calculated
from
  guardrails_policy_setting as ps
  left join guardrails_policy_type as pt on pt.id = ps.policy_type_id
  left join guardrails_resource as r on r.id = ps.resource_id
where
  ps.resource_id = 173434983560398;
```

### All policy settings set on a given resource or below
This query is used to identify all the policy settings applied to a specific resource or its sublevels. This is useful in managing and understanding the security measures in place for that resource.

```sql
select
  r.trunk_title as resource,
  ps.resource_id,
  pt.trunk_title as policy_type,
  ps.value,
  ps.is_calculated
from
  guardrails_policy_setting as ps
  left join guardrails_policy_type as pt on pt.id = ps.policy_type_id
  left join guardrails_resource as r on r.id = ps.resource_id
where
  ps.filter = 'resourceId:173434983560398 level:self,descendant';
```

### All policy settings related to AWS > S3 > Bucket
Discover the segments that have specific policy settings related to AWS S3 Buckets. This allows for a comprehensive overview of policy types and their respective values, useful for security audits and compliance checks.

```sql
select
  r.trunk_title as resource,
  ps.resource_id,
  pt.trunk_title as policy_type,
  ps.value,
  ps.is_calculated
from
  guardrails_policy_setting as ps
  left join guardrails_policy_type as pt on pt.id = ps.policy_type_id
  left join guardrails_resource as r on r.id = ps.resource_id
where
  ps.filter = 'resourceTypeId:"tmod:@turbot/aws-s3#/resource/types/bucket"';
```