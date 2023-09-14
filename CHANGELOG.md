## v0.12.0 [2023-09-14]

_Enhancements_

- Added the `resource_object` and `object` columns to `guardrails_notification` and `guardrails_resource` tables respectively. ([#7](https://github.com/turbot/steampipe-plugin-guardrails/pull/7))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v551-2023-07-26). ([#3](https://github.com/turbot/steampipe-plugin-guardrails/pull/3))
- Recompiled plugin with `github.com/turbot/go-kit v0.7.0`. ([#5](https://github.com/turbot/steampipe-plugin-guardrails/pull/5))
- Recompiled plugin with `github.com/stretchr/testify v1.8.4`. ([#4](https://github.com/turbot/steampipe-plugin-guardrails/pull/4))

## v0.11.1 [2023-07-28]

_Bug fixes_

- Fix typo in the `docs/index.md` file.

## v0.11.0 [2023-07-28]

_What's new?_

- The [Turbot](https://hub.steampipe.io/plugins/turbot/turbot/tables) tables have now been rebranded to use `Guardrails` instead:
  - [guardrails_active_grant](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_active_grant)
  - [guardrails_control](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_control)
  - [guardrails_control_type](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_control_type)
  - [guardrails_grant](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_grant)
  - [guardrails_mod_version](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_mod_version)
  - [guardrails_notification](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_notification)
  - [guardrails_policy_setting](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_policy_setting)
  - [guardrails_policy_type](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_policy_type)
  - [guardrails_policy_value](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_policy_value)
  - [guardrails_resource](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_resource)
  - [guardrails_resource_type](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_resource_type)
  - [guardrails_smart_folder](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_smart_folder)
  - [guardrails_tag](https://hub.steampipe.io/plugins/turbot/turbot/tables/guardrails_tag)

