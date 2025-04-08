---
title: "Steampipe Table: guardrails_mod_version - Query Guardrails Mod Versions using SQL"
description: "Allows users to query Guardrails Mod Versions, providing insights into the different mod versions and their associated metadata."
folder: "Mod"
---

# Table: guardrails_mod_version - Query Guardrails Mod Versions using SQL

Guardrails is a service that helps manage and enforce security, compliance, and operational policies across your cloud environments. It provides a unified view of policy enforcement and compliance status across multiple cloud providers, accounts, and regions. The Mod Versions within Guardrails represent different versions of these policies, including their details and associated metadata.

## Table Usage Guide

The `guardrails_mod_version` table provides insights into the different versions of policies within Guardrails. As a security or compliance officer, explore policy version-specific details through this table, including version numbers, descriptions, and associated metadata. Utilize it to track changes over time, understand the evolution of your policies, and ensure compliance with your organizational standards.

## Examples

### Version details for aws mod
Analyze the settings to understand the status and version of the AWS module in use within your workspace. This is useful to ensure you're working with the most current version and to troubleshoot any version-related issues.

```sql+postgres
select
  name,
  version,
  status,
  workspace
from 
  guardrails_mod_version where name = 'aws';
```

```sql+sqlite
select
  name,
  version,
  status,
  workspace
from 
  guardrails_mod_version where name = 'aws';
```

### Get recommended mod version for aws-acm
Explore the recommended version status for a specific module in the Guardrails, allowing you to ensure you're using the most suitable version for optimal performance.

```sql+postgres
select
  name,
  version,
  status
from
  guardrails_mod_version where name = 'aws-acm' and status = 'RECOMMENDED';
```

```sql+sqlite
select
  name,
  version,
  status
from
  guardrails_mod_version where name = 'aws-acm' and status = 'RECOMMENDED';
```

### List available mod versions for aws-acm
Explore the available versions of a specific module in your AWS Certificate Manager to keep track of updates or changes. This can be useful in maintaining the security and functionality of your environment.

```sql+postgres
select
  name,
  version,
  status
from
  guardrails_mod_version where name = 'aws-acm' and status = 'AVAILABLE';
```

```sql+sqlite
select
  name,
  version,
  status
from
  guardrails_mod_version where name = 'aws-acm' and status = 'AVAILABLE';
```

### List mod versions using the filter syntax
Explore which mod versions are currently in use and their status, specifically those associated with the 'aws-x' filter. This can be useful for understanding the distribution and application of different mod versions in your environment.

```sql+postgres
select
  name,
  version,
  status
from
  guardrails_mod_version where filter = 'aws-x';
```

```sql+sqlite
select
  name,
  version,
  status
from
  guardrails_mod_version where filter = 'aws-x';
```