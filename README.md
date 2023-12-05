![image](https://hub.steampipe.io/images/plugins/turbot/guardrails-social-graphic.png)

# Turbot Guardrails Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, identities, policy settings, grants and more from Turbot Guardrails.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/guardrails)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/turbot/guardrails)
- Community: [Join #steampipe on Slack →](https:/turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-guardrails/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install guardrails
```

Run a query:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-guardrails.git
cd steampipe-plugin-guardrails
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```shell
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/guardrails.spc
```

Try it!

```shell
steampipe query
> .inspect guardrails
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-guardrails/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-guardrails/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Turbot Plugin](https://github.com/turbot/steampipe-plugin-guardrails/labels/help%20wanted)
