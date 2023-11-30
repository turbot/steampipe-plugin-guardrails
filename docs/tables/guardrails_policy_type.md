---
title: "Steampipe Table: guardrails_policy_type - Query Guardrails Policy Types using SQL"
description: "Allows users to query Policy Types in Guardrails, specifically providing details about each policy type, including its name, description, and associated metadata."
---

# Table: guardrails_policy_type - Query Guardrails Policy Types using SQL

Guardrails is a service that provides policy management and enforcement capabilities. It allows you to define and manage policies for various resources, ensuring compliance with your organization's security and governance requirements. Policy Types in Guardrails are the different categories of policies that can be created and managed.

## Table Usage Guide

The `guardrails_policy_type` table provides insights into Policy Types within Guardrails. As a Security Engineer, explore policy type-specific details through this table, including their names, descriptions, and associated metadata. Utilize it to uncover information about different policy types, such as those related to resource management, access control, and security compliance.

## Examples

### List all policy types
Explore the various policy types available in your system in an organized manner to better understand your security infrastructure and manage your policies effectively. This query is useful for identifying and managing the range of policy types in your system.

```sql
select
  id,
  uri,
  trunk_title
from
  guardrails_policy_type
order by
  trunk_title;
```

### List all policy types with additional detail
Explore policy settings by gaining insights into additional details such as descriptions and related links. This allows for a better understanding of each policy type and its specific configurations.

```sql
select
  trunk_title as "policy_name",
  description,
  schema ->> 'enum' as "policy_settings",
  uri as "policy_uri"
from
  guardrails_policy_type
order by
  trunk_title;
```

### List all policy types for AWS S3
Discover the variety of policy types available for AWS S3 to better manage and secure your cloud storage resources. This can help you understand the different levels of control you can exert over your S3 resources.

```sql
select
  id,
  uri,
  trunk_title
from
  guardrails_policy_type
where
  mod_uri like 'tmod:@turbot/aws-s3%'
order by
  trunk_title;
```

### Count policy types by cloud provider
Explore the distribution of policy types across different cloud providers such as AWS, Azure, and GCP to understand their usage and prevalence. This information can be beneficial for assessing your organization's cloud utilization and security posture.

```sql
select
  sum(case when mod_uri like 'tmod:@turbot/aws-%' then 1 else 0 end) as aws,
  sum(case when mod_uri like 'tmod:@turbot/azure-%' then 1 else 0 end) as azure,
  sum(case when mod_uri like 'tmod:@turbot/gcp-%' then 1 else 0 end) as gcp,
  count(*) as total
from
  guardrails_policy_type;
```

### Policy types that target AWS > S3 > Bucket
Explore which policy types are specifically targeting your AWS S3 buckets. This is useful to assess and manage the security and compliance of your S3 resources.

```sql
select
  trunk_title,
  uri,
  targets
from
  guardrails_policy_type
where
  targets ? 'tmod:@turbot/aws-s3#/resource/types/bucket';
```