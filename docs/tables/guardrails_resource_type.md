---
title: "Steampipe Table: guardrails_resource_type - Query Guardrails Resource Types using SQL"
description: "Allows users to query Guardrails Resource Types, specifically providing details on the resource type, description, and associated metadata. It offers insights into the resource categorization and description used within the Guardrails service."
---

# Table: guardrails_resource_type - Query Guardrails Resource Types using SQL

Guardrails is a service that helps manage and enforce security, compliance, and operational policies in real-time across your cloud infrastructure. It provides a way to categorize resources and apply specific policies based on the resource type, enabling more granular control and governance. Guardrails Resource Types are the different categories of resources that can be managed and monitored using the Guardrails service.

## Table Usage Guide

The `guardrails_resource_type` table provides insights into the resource types within the Guardrails service. As a security engineer or cloud administrator, explore resource-specific details through this table, including type description and associated metadata. Utilize it to understand the categorization of resources within your cloud infrastructure, helping you to apply and manage policies more effectively.

## Examples

### List all resource types
Explore various resource types to understand the different categories available, which can assist in organizing and managing your resources more effectively.

```sql
select
  id,
  uri,
  trunk_title
from
  guardrails_resource_type
order by
  trunk_title;
```

### List all resource types for AWS S3
Explore the variety of resource types associated with AWS S3 to better manage and understand your cloud storage resources. This is useful in comprehending the structure and organization of your S3 resources, aiding in efficient resource utilization and management.

```sql
select
  id,
  uri,
  trunk_title
from
  guardrails_resource_type
where
  mod_uri like 'tmod:@turbot/aws-s3%'
order by
  trunk_title;
```

### Count resource types by cloud provider
The query helps to analyze the distribution of resource types across different cloud providers, such as AWS, Azure, and GCP. This can be useful in understanding the spread of resources and making decisions about resource management or optimization.

```sql
select
  sum(case when mod_uri like 'tmod:@turbot/aws-%' then 1 else 0 end) as aws,
  sum(case when mod_uri like 'tmod:@turbot/azure-%' then 1 else 0 end) as azure,
  sum(case when mod_uri like 'tmod:@turbot/gcp-%' then 1 else 0 end) as gcp,
  count(*) as total
from
  guardrails_resource_type;
```