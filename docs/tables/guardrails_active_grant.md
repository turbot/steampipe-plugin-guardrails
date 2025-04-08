---
title: "Steampipe Table: guardrails_active_grant - Query Guardrails Active Grants using SQL"
description: "Allows users to query Active Grants in Guardrails, providing insights into the current permissions and access levels granted within the system."
folder: "Active Grant"
---

# Table: guardrails_active_grant - Query Guardrails Active Grants using SQL

Guardrails is a security management service that enables organizations to implement and enforce consistent security policies across their environment. It provides a unified view of security posture, with the ability to monitor and manage security configurations, compliance status, and incident response. Active Grants within Guardrails represent the permissions and access levels currently granted to users or entities.

## Table Usage Guide

The `guardrails_active_grant` table provides insights into the current permissions and access levels within Guardrails. As a security administrator, you can explore grant-specific details through this table, including the grantee, the grantor, the permission level, and the time the grant was made. Use it to monitor and manage access levels, ensuring that only the appropriate permissions are granted and maintained.

**Important Notes**
- The `guardrails_active_grant` table will only return active grants. Use the `guardrails_grant` table to get a list of all grants.

## Examples

### Basic info
Explore which active grants are in place by examining the status, associated email and profile ID. This allows you to assess the various levels of access and resources granted, providing a comprehensive overview of permissions within your organization.

```sql+postgres
select
  grant_id,
  identity_status,
  identity_email,
  identity_profile_id,
  identity_trunk_title,
  level_title,
  resource_trunk_title
from
  guardrails_active_grant;
```

```sql+sqlite
select
  grant_id,
  identity_status,
  identity_email,
  identity_profile_id,
  identity_trunk_title,
  level_title,
  resource_trunk_title
from
  guardrails_active_grant;
```

### List active grants for an identity
Determine the active grants associated with a specific user's email. This is useful for understanding and managing the access rights and privileges of individual users within a system.

```sql+postgres
select
  grant_id,
  identity_status,
  identity_email,
  identity_trunk_title,
  level_title,
  resource_trunk_title
from
  guardrails_active_grant
where
  identity_email = 'abc@gmail.com';
```

```sql+sqlite
select
  grant_id,
  identity_status,
  identity_email,
  identity_trunk_title,
  level_title,
  resource_trunk_title
from
  guardrails_active_grant
where
  identity_email = 'abc@gmail.com';
```

### List active grants for inactive identities
Explore which active grants are associated with inactive identities to assess potential security risks and manage access control. This would be particularly useful in maintaining organizational security by ensuring that inactive identities do not have unnecessary access privileges.

```sql+postgres
select
  grant_id,
  identity_status,
  identity_email,
  level_title,
  resource_trunk_title
from
  guardrails_active_grant
where
  identity_status = 'Inactive';
```

```sql+sqlite
select
  grant_id,
  identity_status,
  identity_email,
  level_title,
  resource_trunk_title
from
  guardrails_active_grant
where
  identity_status = 'Inactive';
```

### List inactive grants
Identify the grants that are currently inactive. This could be useful for auditing purposes or to clean up unused or unnecessary access permissions.

```sql+postgres
select 
  grant_id, 
  identity_email,
  level_title,
  level_trunk_title,
  level_uri,
  resource_trunk_title,
  resource_type_trunk_title
from 
  guardrails_grant 
where grant_id not in (select grant_id from guardrails_active_grant);
```

```sql+sqlite
select 
  grant_id, 
  identity_email,
  level_title,
  level_trunk_title,
  level_uri,
  resource_trunk_title,
  resource_type_trunk_title
from
  guardrails_grant
where grant_id not in (select grant_id from guardrails_active_grant);
```