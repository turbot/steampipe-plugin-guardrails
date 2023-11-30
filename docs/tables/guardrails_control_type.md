---
title: "Steampipe Table: guardrails_control_type - Query Guardrails Control Types using SQL"
description: "Allows users to query Guardrails Control Types, specifically the control type, severity, and category, providing insights into the control type classification and its criticality."
---

# Table: guardrails_control_type - Query Guardrails Control Types using SQL

Guardrails is a policy as code service that allows you to define and enforce policies across your cloud infrastructure. It provides a centralized way to set up and manage controls for various cloud resources, including virtual machines, databases, web applications, and more. Guardrails helps you stay informed about the compliance status of your cloud resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `guardrails_control_type` table provides insights into control types within Guardrails. As a Security or Compliance officer, explore control type-specific details through this table, including control type, severity, and category. Utilize it to uncover information about control types, such as their classification and criticality, facilitating better risk management and compliance assurance.

## Examples

### List all control types
Discover the segments that are sorted by the title within the guardrails control system, allowing you to analyze and organize the control types effectively. This can be useful for gaining insights into the structure and organization of your control system.

```sql
select
  id,
  uri,
  trunk_title
from
  guardrails_control_type
order by
  trunk_title;
```

### List all control types for AWS S3
Identify all control types related to AWS S3 to gain insights into the various security and configuration measures available. This could be useful for assessing the elements within your S3 setup and optimizing for best practices.

```sql
select
  id,
  uri,
  trunk_title
from
  guardrails_control_type
where
  mod_uri like 'tmod:@turbot/aws-s3%'
order by
  trunk_title;
```

### Count control types by cloud provider
Explore the distribution of control types across various cloud providers to understand their usage patterns and make informed decisions about resource allocation and risk management. This can help in identifying the most utilized cloud provider and strategizing resource management accordingly.

```sql
select
  sum(case when mod_uri like 'tmod:@turbot/aws-%' then 1 else 0 end) as aws,
  sum(case when mod_uri like 'tmod:@turbot/azure-%' then 1 else 0 end) as azure,
  sum(case when mod_uri like 'tmod:@turbot/gcp-%' then 1 else 0 end) as gcp,
  count(*) as total
from
  guardrails_control_type;
```

### Control types that target AWS > S3 > Bucket
Explore the control types that specifically target AWS S3 Buckets to better manage and secure your cloud resources. This is particularly useful for ensuring that your AWS S3 Buckets adhere to best practices and regulatory compliance.

```sql
select
  trunk_title,
  uri,
  targets
from
  guardrails_control_type
where
  targets ? 'tmod:@turbot/aws-s3#/resource/types/bucket';
```