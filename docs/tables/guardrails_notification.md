---
title: "Steampipe Table: guardrails_notification - Query Guardrails Notifications using SQL"
description: "Allows users to query Guardrails Notifications, particularly the notification details including the type, status, and related metadata, providing insights into the system's security and compliance status."
folder: "Notification"
---

# Table: guardrails_notification - Query Guardrails Notifications using SQL

Guardrails is a service that provides continuous compliance and security for cloud infrastructure. It offers real-time detection of potential threats and violations, and automatically sends notifications when such events occur. These notifications contain crucial information about the identified issues, including their type, status, and related metadata.

## Table Usage Guide

The `guardrails_notification` table provides insights into the notifications generated by Guardrails. As a security analyst, you can explore notification-specific details through this table, including the type of violation, its status, and related metadata. Utilize it to understand the security posture of your cloud infrastructure, identify potential threats, and take necessary remedial actions.

**Important Notes**
- When querying this table, we recommend using at least one of these columns (usually in the `where` clause):
  - `id`
  - `resource_id`
  - `notification_type`
  - `control_id`
  - `control_type_id`
  - `control_type_uri`
  - `resource_type_id`
  - `resource_type_uri`
  - `policy_setting_type_id`
  - `policy_setting_type_uri`
  - `actor_identity_id`
  - `create_timestamp`
  - `filter`

For more information on how to construct a `filter`, please see [Notifications examples](https://turbot.com/guardrails/docs/reference/filter/notifications#examples).

## Examples

### Find all Turbot grants activations in last 1 week using `filter`
Explore recent Turbot grant activations from the past week. This is useful for keeping track of security permissions and understanding who has been given access to what within your system.

```sql+postgres
select
  active_grant_id,
  notification_type,
  active_grant_type_title,
  active_grant_level_title,
  create_timestamp,
  actor_identity_trunk_title,
  active_grant_identity_trunk_title,
  active_grant_valid_to_timestamp,
  active_grant_identity_profile_id,
  resource_title
from
  guardrails_notification
where
  filter = 'notificationType:activeGrant createTimestamp:>T-1w'
  and active_grant_type_title = 'Turbot'
order by
  create_timestamp desc,
  notification_type,
  actor_identity_trunk_title,
  resource_title;
```

```sql+sqlite
select
  active_grant_id,
  notification_type,
  active_grant_type_title,
  active_grant_level_title,
  create_timestamp,
  actor_identity_trunk_title,
  active_grant_identity_trunk_title,
  active_grant_valid_to_timestamp,
  active_grant_identity_profile_id,
  resource_title
from
  guardrails_notification
where
  filter = 'notificationType:activeGrant createTimestamp:>T-1w'
  and active_grant_type_title = 'Turbot'
order by
  create_timestamp desc,
  notification_type,
  actor_identity_trunk_title,
  resource_title;
```

### Find all AWS grants activations in last 7 days
Discover the segments that have been granted AWS access in the past week. This can be useful for auditing purposes and to ensure that all access grants align with your organization's security policy.

```sql+postgres
select
  active_grant_id,
  notification_type,
  active_grant_type_title,
  active_grant_level_title,
  create_timestamp,
  actor_identity_trunk_title,
  active_grant_identity_trunk_title,
  active_grant_valid_to_timestamp,
  active_grant_identity_profile_id,
  resource_title
from
  guardrails_notification
where
  notification_type = 'active_grants_created'
  and create_timestamp >= (current_date - interval '7' day)
  and active_grant_type_title = 'AWS'
order by
  create_timestamp desc,
  notification_type,
  actor_identity_trunk_title,
  resource_title;
```

```sql+sqlite
select
  active_grant_id,
  notification_type,
  active_grant_type_title,
  active_grant_level_title,
  create_timestamp,
  actor_identity_trunk_title,
  active_grant_identity_trunk_title,
  active_grant_valid_to_timestamp,
  active_grant_identity_profile_id,
  resource_title
from
  guardrails_notification
where
  notification_type = 'active_grants_created'
  and create_timestamp >= date('now','-7 day')
  and active_grant_type_title = 'AWS'
order by
  create_timestamp desc,
  notification_type,
  actor_identity_trunk_title,
  resource_title;
```

### Find all AWS S3 buckets created notifications in last 7 days
Discover the recent activities related to AWS S3 bucket creation. This query is useful for gaining insights into new resources, helping you monitor and manage your AWS S3 bucket inventory effectively.

```sql+postgres
select
  create_timestamp,
  resource_id,
  resource_title,
  resource_trunk_title,
  actor_identity_trunk_title
from
  guardrails_notification
where
  notification_type = 'resource_created'
  and create_timestamp >= (current_date - interval '120' day)
  and resource_type_uri = 'tmod:@turbot/aws-s3#/resource/types/bucket'
order by
  create_timestamp desc;
```

```sql+sqlite
select
  create_timestamp,
  resource_id,
  resource_title,
  resource_trunk_title,
  actor_identity_trunk_title
from
  guardrails_notification
where
  notification_type = 'resource_created'
  and create_timestamp >= date('now','-120 day')
  and resource_type_uri = 'tmod:@turbot/aws-s3#/resource/types/bucket'
order by
  create_timestamp desc;
```

### All policy settings notifications on a given resource or below in last 90 days
This example helps you to monitor all policy settings notifications related to a specific resource or its descendants over the past 90 days. It is useful for tracking changes and updates in policy settings, aiding in resource management and security compliance.

```sql+postgres
select
  notification_type,
  create_timestamp,
  policy_setting_id,
  policy_setting_type_trunk_title,
  policy_setting_type_uri,
  resource_trunk_title,
  resource_type_trunk_title,
  policy_setting_type_read_only,
  policy_setting_type_secret,
  policy_setting_value
from
  guardrails_notification
where
  resource_id = 191382256916538
  and create_timestamp >= (current_date - interval '90' day)
  and filter = 'notificationType:policySetting level:self,descendant'
order by
  create_timestamp desc;
```

```sql+sqlite
select
  notification_type,
  create_timestamp,
  policy_setting_id,
  policy_setting_type_trunk_title,
  policy_setting_type_uri,
  resource_trunk_title,
  resource_type_trunk_title,
  policy_setting_type_read_only,
  policy_setting_type_secret,
  policy_setting_value
from
  guardrails_notification
where
  resource_id = 191382256916538
  and create_timestamp >= date('now','-90 day')
  and filter = 'notificationType:policySetting level:self,descendant'
order by
  create_timestamp desc;
```

### All policy settings notifications for AWS > Account > Regions policy
Explore the notifications related to policy settings for your AWS account's regional policies. This is particularly useful for keeping track of policy changes and ensuring compliance with your organization's standards.

```sql+postgres
select
  notification_type,
  create_timestamp,
  policy_setting_id,
  resource_id,
  resource_trunk_title,
  jsonb_pretty(policy_setting_value::jsonb) as policy_setting_value
from
  guardrails_notification
where
  policy_setting_type_uri = 'tmod:@turbot/aws#/policy/types/regionsDefault'
  and filter = 'notificationType:policySetting level:self'
order by
  create_timestamp desc;
```

```sql+sqlite
select
  notification_type,
  create_timestamp,
  policy_setting_id,
  resource_id,
  resource_trunk_title,
  policy_setting_value
from
  guardrails_notification
where
  policy_setting_type_uri = 'tmod:@turbot/aws#/policy/types/regionsDefault'
  and filter = 'notificationType:policySetting level:self'
order by
  create_timestamp desc;
```

### All notifications for AWS > Account > Budget > Budget control
Determine the areas in which budget controls have been notified for your AWS account. This allows you to assess the state and reason for each control, providing insights for better financial management.

```sql+postgres
select
  notification_type,
  create_timestamp,
  control_id,
  resource_trunk_title,
  control_state,
  control_reason
from
  guardrails_notification
where
  control_type_uri = 'tmod:@turbot/aws#/control/types/budget'
  and filter = 'notificationType:control level:self'
order by
  resource_id,
  create_timestamp desc;
```

```sql+sqlite
select
  notification_type,
  create_timestamp,
  control_id,
  resource_trunk_title,
  control_state,
  control_reason
from
  guardrails_notification
where
  control_type_uri = 'tmod:@turbot/aws#/control/types/budget'
  and filter = 'notificationType:control level:self'
order by
  resource_id,
  create_timestamp desc;
```