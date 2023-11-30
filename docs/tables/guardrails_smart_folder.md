---
title: "Steampipe Table: guardrails_smart_folder - Query Guardrails Smart Folders using SQL"
description: "Allows users to query Guardrails Smart Folders, specifically providing insights into the organization and grouping of guardrails based on their attributes."
---

# Table: guardrails_smart_folder - Query Guardrails Smart Folders using SQL

A Guardrails Smart Folder is a feature within Guardrails that allows users to organize and group guardrails based on their attributes. It provides a centralized way to manage and categorize guardrails, enhancing the efficiency of security and compliance management. Guardrails Smart Folder helps users to streamline their guardrails management process and improve the visibility of their security posture.

## Table Usage Guide

The `guardrails_smart_folder` table provides insights into the organization and grouping of guardrails within Guardrails. As a security engineer, you can explore smart folder-specific details through this table, including the guardrails grouped under each smart folder, their attributes, and associated metadata. Utilize it to uncover information about smart folders, such as their grouping logic, the number of guardrails under each folder, and the overall organization of your guardrails.

## Examples

### List all smart folders
Discover the segments that list all your smart folders. This can help you manage and organize your data more efficiently, particularly when dealing with large volumes of data.

```sql
select
  id,
  title
from
  guardrails_smart_folder;
```

### List smart folders with their policy settings
Explore which smart folders have specific policy settings assigned to them. This can help you understand and manage the security measures applied to different folders in your system.

```sql
select
  sf.trunk_title as smart_folder,
  pt.trunk_title as policy,
  ps.id,
  ps.precedence,
  ps.is_calculated,
  ps.value
from
  guardrails_smart_folder as sf
  left join guardrails_policy_setting as ps on ps.resource_id = sf.id
  left join guardrails_policy_type as pt on pt.id = ps.policy_type_id
order by
  smart_folder;
```

### List smart folders with their attached resources
Discover the segments that are linked to specific smart folders. This information can be useful in understanding how resources are grouped and managed, providing insights into your resource organization strategy.
Get each smart folder with an array of the resources attached to it:


```sql
select
  title,
  attached_resource_ids
from
  guardrails_smart_folder
order by
  title;
```

Create a row per smart folder and resource:

```sql
select
  sf.title as smart_folder,
  sf_resource_id
from
  guardrails_smart_folder as sf,
  jsonb_array_elements(sf.attached_resource_ids) as sf_resource_id
order by
  smart_folder,
  sf_resource_id;
```

Unfortunately, this query to join the smart folder with its resources does not
work yet due to issues with qualifier handling in the Steampipe Postgres FDW:

```sql
select
  sf.title as smart_folder,
  r.trunk_title as resource,
  r.id
from
  guardrails_smart_folder as sf
  cross join jsonb_array_elements(sf.attached_resource_ids) as sf_resource_id
  left join guardrails_resource as r on r.id = sf_resource_id::bigint;
```