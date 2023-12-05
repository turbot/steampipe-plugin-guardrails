---
title: "Steampipe Table: guardrails_resource - Query Guardrails Resources using SQL"
description: "Allows users to query Guardrails Resources, providing insights into the security, compliance, and operational risk of their cloud infrastructure."
---

# Table: guardrails_resource - Query Guardrails Resources using SQL

Guardrails is a service that provides a unified view of the security, compliance, and operational health of your cloud infrastructure. It enables you to continuously monitor and manage your cloud resources across multiple cloud platforms, including AWS, Azure, and Google Cloud. Guardrails helps you identify potential risks and vulnerabilities, and take appropriate actions to mitigate them.

## Table Usage Guide

The `guardrails_resource` table provides insights into the resources monitored by Guardrails. As a cloud security engineer, you can use this table to explore resource-specific details, including their current compliance status, potential risks, and associated metadata. This table is beneficial for identifying non-compliant resources, understanding the security posture of your cloud infrastructure, and taking necessary actions to ensure compliance and mitigate risks.

**Important Notes**
- When querying this table, we recommend using at least one of these columns (usually in the `where` clause):
  - `id`
  - `resource_type_id`
  - `resource_type_uri`
  - `filter`

## Examples

### List all AWS IAM Roles
Explore the different roles within your AWS IAM setup to understand their configurations and creation times. This can aid in managing access controls and ensuring optimal security practices.

```sql
select
  id,
  title,
  create_timestamp,
  metadata,
  data
from
  guardrails_resource
where
  resource_type_uri = 'tmod:@turbot/aws-iam#/resource/types/role';
```

### List all S3 buckets with a given Owner tag
Explore which S3 buckets are associated with a specific owner. This is useful to manage and track resources based on ownership within an AWS environment.

```sql
select
  id,
  title,
  tags
from
  guardrails_resource
where
  resource_type_uri = 'tmod:@turbot/aws-s3#/resource/types/bucket'
  and tags ->> 'Owner' = 'Jane';
```

### Get a specific resource by ID
Determine the details of a specific resource using its unique identifier. This can be particularly useful when you need to quickly access and review the specifics of a resource in your system.

```sql
select
  id,
  title,
  create_timestamp,
  metadata,
  data
from
  guardrails_resource
where
  id = 216005088871602;
```

### Filter for resources using Turbot filter syntax
Analyze the distribution of resources based on their type to understand the prevalence of specific resource types in your AWS IAM configuration. This query is particularly useful in identifying and managing resource types that are heavily utilized.

```sql
select
  resource_type_uri,
  count(*)
from
  guardrails_resource
where
  filter = 'resourceTypeId:"tmod:@turbot/aws-iam#/resource/types/iam"'
group by
  resource_type_uri
order by
  count desc;
```

### Search for AWS IAM Roles by name (Turbot side)
Explore which AWS IAM roles have 'admin' access to better understand and manage permissions within your AWS environment. This can help in enhancing security by identifying potential areas of risk.
This query will ask Turbot to filter the resources down to the given `filter`,
limiting the results by name.


```sql
select
  id,
  title,
  create_timestamp,
  metadata,
  data
from
  guardrails_resource
where
  resource_type_uri = 'tmod:@turbot/aws-iam#/resource/types/role'
  and filter = 'admin';
```

### Search for AWS IAM Roles by name (Steampipe side)
Discover the segments that include AWS IAM roles with a specific name. This is valuable for identifying roles that may have overly broad or administrative permissions, helping to enhance security and compliance management.
This query gathers all the AWS IAM roles from Turbot and then uses Postgres
level filters to limit the results.


```sql
select
  id,
  title,
  create_timestamp,
  metadata,
  data
from
  guardrails_resource
where
  resource_type_uri = 'tmod:@turbot/aws-iam#/resource/types/role'
  and title ilike '%admin%';
```

### Search for console logins within 7 days
Determine the areas in which console logins have occurred within the past week. This can help in monitoring user activity and identifying any unusual login patterns for enhanced security.

```sql
select
  id,
  title,
  data ->> 'email' as email,
  array_to_string(regexp_matches(trunk_title, '^Turbot > (.*) >'), ' ' ) as "directory",
  trunk_title,
  to_char((data ->> 'lastLoginTimestamp') :: timestamp, 'YYYY-MM-DD HH24:MI') as "last_login"
from
  guardrails_resource
where
  filter = 'resourceTypeId:"tmod:@turbot/turbot-iam#/resource/types/profile" $.lastLoginTimestamp:>=T-7d';
```

### Search for resources created within 7 days, join with count of controls in alarm state
Explore resources created within the past week and assess the number of controls in an alarm state. This is useful for monitoring new resources and their potential risks.

```sql
select
  r.id,
  r.title,
  r.trunk_title,
  r.resource_type_uri,
  to_char(r.create_timestamp, 'YYYY-MM-DD HH24:MI') as create_timestamp,
  count(c.*) as alarm_count
from
  guardrails_resource as r
  left join
    guardrails_control as c
    on r.id = c.resource_id
    and c.state = 'alarm'
where
  r.filter = 'notificationType:resource timestamp:>=T-7d'
group by
  r.id,
  r.title,
  r.trunk_title,
  r.resource_type_uri,
  r.create_timestamp
order by
  r.create_timestamp desc;
```

### Extract all resources from Turbot Guardrails
Explore the full range of resources available in Turbot Guardrails to gain a comprehensive understanding of the scope of your resources and their management. This is beneficial for assessing overall resource utilization and identifying areas for improvement.
WARNING - This is a large query and may take minutes to run. It is not recommended and may timeout.
It's included here as a reference for those who need to extract all data.


```sql
select
  *
from
  guardrails_resource;
```