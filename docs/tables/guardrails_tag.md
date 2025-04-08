---
title: "Steampipe Table: guardrails_tag - Query Guardrails Tags using SQL"
description: "Allows users to query Guardrails Tags, providing critical insights into the tag data associated with Guardrails resources."
folder: "Tag"
---

# Table: guardrails_tag - Query Guardrails Tags using SQL

Guardrails is a tool that helps in managing the security and compliance of cloud resources. It allows users to define and enforce policies across different cloud services, ensuring that resources are secure and compliant with industry standards. Guardrails tags are metadata that can be assigned to Guardrails resources to help organize and manage them.

## Table Usage Guide

The `guardrails_tag` table offers a comprehensive view into Guardrails Tags within the Guardrails service. As a Security Engineer, you can leverage this table to analyze and manage tag data associated with Guardrails resources, including their names and values. This can facilitate effective organization, identification, and management of resources based on custom-defined metadata.

**Important Notes**
- When querying this table, we recommend using at least one of these columns (usually in the `where` clause):
  - `id`
  - `key`
  - `value`
  - `filter`

## Examples

### List all tags
Explore all tags to understand their key-value pairings, which can help in organizing and locating specific resources within the Guardrails system.

```sql+postgres
select
  *
from
  guardrails_tag
order by
  key,
  value;
```

```sql+sqlite
select
  *
from
  guardrails_tag
order by
  key,
  value;
```

### Find all resources for the Sales department
Explore which resources are specifically allocated for the Sales department, assisting in resource management and departmental budgeting.

```sql+postgres
select
  key,
  value,
  resource_ids
from
  guardrails_tag
where
  key = 'Department'
  and value = 'Sales';
```

```sql+sqlite
select
  key,
  value,
  resource_ids
from
  guardrails_tag
where
  key = 'Department'
  and value = 'Sales';
```

### Find departments with the most tagged resources
Analyze the settings to understand which departments have the most resources tagged to them. This can help to identify areas that may require more oversight or resource allocation.

```sql+postgres
select
  key,
  value,
  jsonb_array_length(resource_ids) as count
from
  guardrails_tag
where
  key = 'Department'
order by
  count desc;
```

```sql+sqlite
select
  key,
  value,
  json_array_length(resource_ids) as count
from
  guardrails_tag
where
  key = 'Department'
order by
  count desc;
```

### List tags without values
Discover the segments that contain tags without assigned values. This can be useful in identifying potential gaps or inconsistencies in your data tagging practices.

```sql+postgres
select
  *
from
  guardrails_tag
where
  value is null or trim(value) = '';
```

```sql+sqlite
select
  *
from
  guardrails_tag
where
  value is null or trim(value) = '';
```