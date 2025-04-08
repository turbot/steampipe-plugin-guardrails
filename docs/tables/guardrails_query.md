---
title: "Steampipe Table: guardrails_query - Query Guardrails using SQL"
description: "Allows users to query Guardrails, providing critical insights about the resource types, permissions, and any settings."
folder: "Query"
---

# Table: guardrails_query - Query Guardrails using SQL

Guardrails is a tool that helps in managing the security and compliance of cloud resources. It allows users to define and enforce policies across different cloud services, ensuring that resources are secure and compliant with industry standards.

## Table Usage Guide

The `guardrails_query` table provides a detailed overview of Guardrails, making it an invaluable tool for Security Engineers to examine and manage data related to Guardrails resources and settings, encompassing their names and values.

**Important Notes**
- When querying this table, we must have to pass the `query` in `where` clause.

## Examples

### List recomended mod version
Analyze the settings to understand the RECOMMENDED version of all the module within your workspace. This is useful to ensure you're working with the most current version and to troubleshoot any version-related issues.

```sql+postgres
select
  query,
  m ->> 'name' as mod_name,
  m -> 'versions' as mod_version
from
  guardrails_query,
  jsonb_array_elements(output -> 'modVersionSearches' -> 'items') as m
where
  query = 'query modVersionSearchByName($status: [ModVersionStatus!] = RECOMMENDED) {
  modVersionSearches(status: $status) {
      items {
        name
        versions {
          version
          status
        }
      }
    }
  }';
```

```sql+sqlite
select
  query,
  json_extract(m.value, '$.name') as mod_name,
  json_extract(m.value, '$.versions') as mod_version
from
  guardrails_query,
  json_each(json_extract(output, '$.modVersionSearches.items')) as m
where
  query = 'query modVersionSearchByName($status: [ModVersionStatus!] = RECOMMENDED) {
  modVersionSearches(status: $status) {
      items {
        name
        versions {
          version
          status
        }
      }
    }
  }';

```

### Get control metadata status
Count the number of controls in a workspace.

```sql+postgres
select
  query,
  output -> 'controls' -> 'metadata' -> 'stats' ->> 'total' as control_count
from
  guardrails_query
where
  query = 'query controlMetadataList($filter: [String!]) {
  controls(filter: $filter) {
    metadata {
      stats {
        total
      }
    }
  }
}';
```

```sql+sqlite
select
  query,
  json_extract(output, '$.controls.metadata.stats.total')) as control_count
from
  guardrails_query
where
  query = 'query controlMetadataList($filter: [String!]) {
  controls(filter: $filter) {
    metadata {
      stats {
        total
      }
    }
  }
}';
```