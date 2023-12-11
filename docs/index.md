---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/guardrails.svg"
brand_color: "#FCC119"
display_name: Turbot Guardrails
short_name: guardrails
description: Steampipe plugin to query resources, controls, policies and more from Turbot Guardrails.
og_description: Query Turbot Guardrails with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/guardrails-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Turbot Guardrails + Steampipe

[Turbot Guardrails](https://turbot.com/guardrails) is the leading platform for policy-based control and automatic remediation of enterprise clouds.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  trunk_title,
  uri
from
  guardrails_resource_type;
```

```
+---------------------------------+---------------------------------------------------------+
| trunk_title                     | uri                                                     |
+---------------------------------+---------------------------------------------------------+
| Turbot > IAM > Access Key       | tmod:@turbot/turbot-iam#/resource/types/accessKey       |
| GCP > Monitoring > Alert Policy | tmod:@turbot/gcp-monitoring#/resource/types/alertPolicy |
| AWS > IAM > Access Key          | tmod:@turbot/aws-iam#/resource/types/accessKey          |
| AWS > EC2 > AMI                 | tmod:@turbot/aws-ec2#/resource/types/ami                |
| AWS > SSM > Association         | tmod:@turbot/aws-ssm#/resource/types/association        |
| GCP > Network > Address         | tmod:@turbot/gcp-network#/resource/types/address        |
+---------------------------------+---------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/guardrails/tables)**

## Get started

### Install

Download and install the latest Turbot Guardrails plugin:

```bash
steampipe plugin install guardrails
```

### Credentials

Installing the latest guardrails plugin will create a config file (`~/.steampipe/config/guardrails.spc`) with a single connection named `guardrails`. By default, Steampipe will use your [Turbot Guardrails profiles and credentials](https://turbot.com/guardrails/docs/reference/cli/installation#set-up-your-turbot-guardrails-credentials) exactly the same as the Turbot Guardrails CLI and Turbot Guardrails Terraform provider. In many cases, no extra configuration is required to use Steampipe.

```hcl
connection "guardrails" {
  plugin = "guardrails"
}
```

## Advanced configuration options

If you have a `default` profile setup using the Turbot Guardrails CLI, Steampipe just works with that connection.

For users with multiple workspaces and more complex authentication use cases, here are some examples of advanced configuration options:

### Credentials via key pair

The Turbot Guardrails plugin allows you to set static credentials with the `access_key`, `secret_key`, and `workspace` arguments in any connection profile.

```hcl
connection "guardrails" {
  plugin = "guardrails"
  access_key = "c8e2c2ed-1ca8-429b-b369-010e3cf75aac"
  secret_key = "a3d8385d-47f7-40c5-a90c-bfdf5b43c8dd"
  workspace  = "https://turbot-acme.cloud.turbot.com/"
}
```

### Credentials via Turbot Guardrails config profiles

You can use an existing Turbot Guardrails named profile configured in `/Users/jsmyth/.config/turbot/credentials.yml`. A connect per workspace is a common configuration:

```hcl
connection "guardrails_acme" {
  plugin = "guardrails"
  profile = "turbot-acme"
}

connection "guardrails_dmi" {
  plugin = "guardrails"
  profile = "turbot-dmi"
}

```

### Credentials from environment variables

Environment variables provide another way to specify default Turbot Guardrails CLI credentials:

```sh
export TURBOT_SECRET_KEY=3d397816-575f-4b2a-a470-a96abe29b81a
export TURBOT_ACCESS_KEY=86835f29-1c88-46d9-b6ce-cbe5016842d3
export TURBOT_WORKSPACE=https://turbot-acme.cloud.turbot.com
```

You can also change the default profile to a named profile with the TURBOT_PROFILE  environment variable:

```sh
export TURBOT_PROFILE=turbot-acme
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-guardrails
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)

