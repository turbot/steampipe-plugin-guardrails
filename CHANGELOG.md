## v1.2.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#73](https://github.com/turbot/steampipe-plugin-guardrails/pull/73))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#73](https://github.com/turbot/steampipe-plugin-guardrails/pull/73))

## v1.1.0 [2025-04-15]

_Enhancements_

- Added `folder` metadata to the documentation of all the Guardrails tables for improved organization on the Steampipe Hub. ([#74](https://github.com/turbot/steampipe-plugin-guardrails/pull/74))

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#73](https://github.com/turbot/steampipe-plugin-guardrails/pull/73))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/releases/tag/v5.11.5) that addresses critical and high vulnerabilities in dependent packages. ([#73](https://github.com/turbot/steampipe-plugin-guardrails/pull/73))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#54](https://github.com/turbot/steampipe-plugin-guardrails/pull/54))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#54](https://github.com/turbot/steampipe-plugin-guardrails/pull/54))

## v0.17.1 [2024-07-18]

_Bug fixes_

- Fixed plugin loading issues by eliminating the need for manual caching, ensuring smoother and more reliable plugin installations. ([#50](https://github.com/turbot/steampipe-plugin-guardrails/pull/50))

## v0.17.0 [2024-07-18]

_What's new?_

- Added the `insecure_skip_verify` connection config argument to support bypassing the `SSL/TLS` certificate verification while querying the tables. ([#48](https://github.com/turbot/steampipe-plugin-guardrails/pull/48))

_Enhancements_

- The Plugin and the Steampipe Anywhere binaries are now built with the `netgo` package.

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.10.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v5101-2024-05-09) that adds support for connection key columns. ([#49](https://github.com/turbot/steampipe-plugin-guardrails/pull/49))

## v0.16.0 [2024-02-05]

_Enhancements_

- Updated all the tables to fetch the column data using hydrate functions to optimize the API calls and increase query speed when querying specific columns. ([#30](https://github.com/turbot/steampipe-plugin-guardrails/pull/30))

## v0.15.0 [2024-01-30]

_What's new?_

- New tables added
  - [guardrails_query](https://hub.steampipe.io/plugins/turbot/guardrails/tables/guardrails_query) ([#31](https://github.com/turbot/steampipe-plugin-guardrails/pull/31))

## v0.14.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#26](https://github.com/turbot/steampipe-plugin-guardrails/pull/26))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#26](https://github.com/turbot/steampipe-plugin-guardrails/pull/26))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-guardrails/blob/main/docs/LICENSE). ([#26](https://github.com/turbot/steampipe-plugin-guardrails/pull/26))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#25](https://github.com/turbot/steampipe-plugin-guardrails/pull/25))

## v0.13.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#16](https://github.com/turbot/steampipe-plugin-guardrails/pull/16))

## v0.13.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#11](https://github.com/turbot/steampipe-plugin-guardrails/pull/11))
- Recompiled plugin with Go version `1.21`. ([#11](https://github.com/turbot/steampipe-plugin-guardrails/pull/11))

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

