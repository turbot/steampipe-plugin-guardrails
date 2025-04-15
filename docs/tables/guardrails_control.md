---
title: "Steampipe Table: guardrails_control - Query Guardrails Controls using SQL"
description: "Allows users to query Guardrails Controls, specifically the control status, control category, and control standard, providing insights into the guardrails compliance and security posture."
folder: "Control"
---

# Table: guardrails_control - Query Guardrails Controls using SQL

Guardrails is a governance, risk management, and compliance service that provides a set of predefined controls to help organizations manage their security and compliance posture. These controls are designed to be easily integrated into existing security and compliance processes and can be customized to meet specific organizational needs. Guardrails controls are categorized into various standards and categories to provide a comprehensive and structured approach to security and compliance management.

## Table Usage Guide

The `guardrails_control` table provides insights into the controls within Guardrails. As a security analyst, explore control-specific details through this table, including control status, control category, and control standard. Utilize it to uncover information about controls, such as their compliance status, the categories they fall under, and the standards they adhere to.

**Important Notes**
- When querying this table, we recommend using at least one of these columns (usually in the `where` clause):
  - `id`
  - `control_type_id`
  - `control_type_uri`
  - `resource_type_id`
  - `resource_type_uri`
  - `state`
  - `filter`

## Examples

### Control summary for AWS > IAM > Role > Approved
Explore the status of approved roles in AWS IAM by counting their occurrence. This allows you to identify potential issues and understand the overall health of your IAM roles.
Simple table:


```sql+postgres
select
  state,
  count(*)
from
  guardrails_control
where
  control_type_uri = 'tmod:@turbot/aws-iam#/control/types/roleApproved'
group by
  state
order by
  count desc;
```

```sql+sqlite
select
  state,
  count(*)
from
  guardrails_control
where
  control_type_uri = 'tmod:@turbot/aws-iam#/control/types/roleApproved'
group by
  state
order by
  count(*) desc;
```

Or, if you prefer a full view of all states:

```sql+postgres
select
  control_type_uri,
  sum(case when state = 'ok' then 1 else 0 end) as ok,
  sum(case when state = 'tbd' then 1 else 0 end) as tbd,
  sum(case when state = 'invalid' then 1 else 0 end) as invalid,
  sum(case when state = 'alarm' then 1 else 0 end) as alarm,
  sum(case when state = 'skipped' then 1 else 0 end) as skipped,
  sum(case when state = 'error' then 1 else 0 end) as error,
  sum(case when state in ('alarm', 'error', 'invalid') then 1 else 0 end) as alert,
  count(*) as total
from
  guardrails_control as c
where
  control_type_uri = 'tmod:@turbot/aws-iam#/control/types/roleApproved'
group by
  control_type_uri
order by
  total desc;
```

```sql+sqlite
select
  control_type_uri,
  sum(case when state = 'ok' then 1 else 0 end) as ok,
  sum(case when state = 'tbd' then 1 else 0 end) as tbd,
  sum(case when state = 'invalid' then 1 else 0 end) as invalid,
  sum(case when state = 'alarm' then 1 else 0 end) as alarm,
  sum(case when state = 'skipped' then 1 else 0 end) as skipped,
  sum(case when state = 'error' then 1 else 0 end) as error,
  sum(case when state in ('alarm', 'error', 'invalid') then 1 else 0 end) as alert,
  count(*) as total
from
  guardrails_control as c
where
  control_type_uri = 'tmod:@turbot/aws-iam#/control/types/roleApproved'
group by
  control_type_uri
order by
  total desc;
```

### Control summary for all AWS > IAM controls
This query helps in assessing the security posture of your AWS Identity and Access Management (IAM) controls. It aids in identifying the number of controls in different states, allowing you to quickly pinpoint areas that might need attention or remediation, thereby enhancing your overall security management.

```sql+postgres
select
  state,
  count(*)
from
  guardrails_control
where
  filter = 'controlTypeId:"tmod:@turbot/aws-iam#/resource/types/iam"'
group by
  state
order by
  count desc;
```

```sql+sqlite
select
  state,
  count(*)
from
  guardrails_control
where
  filter = 'controlTypeId:"tmod:@turbot/aws-iam#/resource/types/iam"'
group by
  state
order by
  count(*) desc;
```

Or, if you prefer a full view of all states:

```sql+postgres
select
  control_type_uri,
  sum(case when state = 'ok' then 1 else 0 end) as ok,
  sum(case when state = 'tbd' then 1 else 0 end) as tbd,
  sum(case when state = 'invalid' then 1 else 0 end) as invalid,
  sum(case when state = 'alarm' then 1 else 0 end) as alarm,
  sum(case when state = 'skipped' then 1 else 0 end) as skipped,
  sum(case when state = 'error' then 1 else 0 end) as error,
  sum(case when state in ('alarm', 'error', 'invalid') then 1 else 0 end) as alert,
  count(*) as total
from
  guardrails_control as c
where
  filter = 'controlTypeId:"tmod:@turbot/aws-iam#/resource/types/iam"'
group by
  control_type_uri
order by
  total desc;
```

```sql+sqlite
select
  control_type_uri,
  sum(case when state = 'ok' then 1 else 0 end) as ok,
  sum(case when state = 'tbd' then 1 else 0 end) as tbd,
  sum(case when state = 'invalid' then 1 else 0 end) as invalid,
  sum(case when state = 'alarm' then 1 else 0 end) as alarm,
  sum(case when state = 'skipped' then 1 else 0 end) as skipped,
  sum(case when state = 'error' then 1 else 0 end) as error,
  sum(case when state in ('alarm', 'error', 'invalid') then 1 else 0 end) as alert,
  count(*) as total
from
  guardrails_control as c
where
  filter = 'controlTypeId:"tmod:@turbot/aws-iam#/resource/types/iam"'
group by
  control_type_uri
order by
  total desc;
```

### List controls for AWS > IAM > Role > Approved
Explore the history of changes related to approved roles in AWS IAM. This can help in understanding the compliance status and identifying any unauthorized or accidental modifications.

```sql+postgres
select
  timestamp,
  state,
  reason,
  resource_id,
  control_type_uri
from
  guardrails_control
where
  filter = 'controlTypeId:"tmod:@turbot/aws-iam#/control/types/roleApproved" controlTypeLevel:self'
order by
  timestamp desc;
```

```sql+sqlite
select
  timestamp,
  state,
  reason,
  resource_id,
  control_type_uri
from
  guardrails_control
where
  filter = 'controlTypeId:"tmod:@turbot/aws-iam#/control/types/roleApproved" controlTypeLevel:self'
order by
  timestamp desc;
```

### Query the most recent 10 controls
Explore the latest 10 guardrail controls to understand their current state and the reasons behind it, which is crucial for maintaining system integrity and security.
Note: It's more efficient to have Turbot Guardrails limit the results to the last 10
(`filter = 'limit:10'`), rather than using `limit 10` which will pull all rows
from Turbot Guardrails and will then filter them afterwards on the Steampipe side.


```sql+postgres
select
  timestamp,
  state,
  reason,
  resource_id,
  control_type_uri
from
  guardrails_control
where
  filter = 'limit:10'
order by
  timestamp desc;
```

```sql+sqlite
select
  timestamp,
  state,
  reason,
  resource_id,
  control_type_uri
from
  guardrails_control
where
  filter = 'limit:10'
order by
  timestamp desc;
```

### Control & Resource data for for AWS > IAM > Role > Approved
This query is used to gain insights into the status and reasons for approval of IAM roles in AWS. It helps in managing access controls by identifying roles that are approved, providing a better understanding of the security posture.

```sql+postgres
select
  r.trunk_title,
  r.data ->> 'Arn' as arn,
  r.metadata -> 'aws' ->> 'accountId' as account_id,
  c.state,
  c.reason
from
  guardrails_control as c,
  guardrails_resource as r
where
  -- Filter to the control type
  c.control_type_uri = 'tmod:@turbot/aws-iam#/control/types/roleApproved'
  -- Filter to the resource type as well, reducing the size of the join
  and r.resource_type_uri = 'tmod:@turbot/aws-iam#/resource/types/role'
  and r.id = c.resource_id
order by
  r.trunk_title;
```

```sql+sqlite
select
  r.trunk_title,
  json_extract(r.data, '$.Arn') as arn,
  json_extract(r.metadata, '$.aws.accountId') as account_id,
  c.state,
  c.reason
from
  guardrails_control as c,
  guardrails_resource as r
where
  c.control_type_uri = 'tmod:@turbot/aws-iam#/control/types/roleApproved'
  and r.resource_type_uri = 'tmod:@turbot/aws-iam#/resource/types/role'
  and r.id = c.resource_id
order by
  r.trunk_title;
```

### Extract all controls from Turbot Guardrails
Discover the segments that fall under Turbot Guardrails' controls. This can provide a comprehensive overview, aiding in the efficient management and review of security measures.
WARNING - This is a large query and may take minutes to run. It is not recommended and may timeout.
It's included here as a reference for those who need to extract all data.


```sql+postgres
select
  *
from
  guardrails_control;
```

```sql+sqlite
select
  *
from
  guardrails_control;
```