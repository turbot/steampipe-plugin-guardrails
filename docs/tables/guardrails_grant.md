---
title: "Steampipe Table: guardrails_grant - Query Guardrails Grants using SQL"
description: "Allows users to query Guardrails Grants, specifically the grantee, grantor, and guardrail details, providing insights into the permissions and rules set within a system."
folder: "Grant"
---

# Table: guardrails_grant - Query Guardrails Grants using SQL

Guardrails Grant is a feature in Guardrails that allows users to manage and control the permissions and rules within a system. It provides a way to set up and manage grants for various resources, including users, roles, and groups. Guardrails Grant helps you stay informed about the permissions and rules of your resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `guardrails_grant` table provides insights into the grants within Guardrails. As a System Administrator, explore grant-specific details through this table, including grantee, grantor, and guardrail details. Utilize it to uncover information about grants, such as those with specific permissions, the relationships between grants, and the verification of guardrail rules.

## Examples

### Basic info
Explore the status and details of various identities and their associated profiles and levels to better understand the configuration and organization of your guardrails grant. This can help in managing access controls and permissions effectively.

```sql+postgres
select
  id,
  identity_status,
  identity_email,
  identity_profile_id,
  identity_trunk_title,
  level_title,
  resource_trunk_title
from
  guardrails_grant;
```

```sql+sqlite
select
  id,
  identity_status,
  identity_email,
  identity_profile_id,
  identity_trunk_title,
  level_title,
  resource_trunk_title
from
  guardrails_grant;
```

### List grants for an identity
Explore the level of access granted to a specific user. This is useful for auditing purposes, ensuring that each user has the appropriate level of system permissions.

```sql+postgres
select
  id,
  identity_email,
  identity_family_name,
  level_title,
  level_trunk_title,
from
  guardrails_grant
where
  identity_email = 'xyz@gmail.com';
```

```sql+sqlite
select
  id,
  identity_email,
  identity_family_name,
  level_title,
  level_trunk_title
from
  guardrails_grant
where
  identity_email = 'xyz@gmail.com';
```

### List SuperUser grants
Discover the segments that have been granted SuperUser access. This helps in maintaining security by identifying who has high-level permissions and where these permissions are applied.

```sql+postgres
select
  id,
  identity_email,
  identity_family_name,
  level_title,
  resource_trunk_title
from
  guardrails_grant
where
  level_uri  = 'tmod:@turbot/turbot-iam#/permission/levels/superuser';
```

```sql+sqlite
select
  id,
  identity_email,
  identity_family_name,
  level_title,
  resource_trunk_title
from
  guardrails_grant
where
  level_uri  = 'tmod:@turbot/turbot-iam#/permission/levels/superuser';
```

### List grants for inactive identities
Discover the segments that have been granted access to inactive identities. This is useful to ensure that no unnecessary permissions are given to inactive users, thereby enhancing security measures.

```sql+postgres
select
  id,
  identity_email,
  identity_status,
  resource_trunk_title
from
  guardrails_grant
where
  identity_status = 'Inactive';
```

```sql+sqlite
select
  id,
  identity_email,
  identity_status,
  resource_trunk_title
from
  guardrails_grant
where
  identity_status = 'Inactive';
```