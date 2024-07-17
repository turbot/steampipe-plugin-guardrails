package turbot

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type guardrailsConfig struct {
	Profile            *string `hcl:"profile"`
	AccessKey          *string `hcl:"access_key"`
	SecretKey          *string `hcl:"secret_key"`
	Workspace          *string `hcl:"workspace"`
	InsecureSkipVerify *bool   `hcl:"insecure_skip_verify,optional"`
}

func ConfigInstance() interface{} {
	return &guardrailsConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) guardrailsConfig {
	if connection == nil || connection.Config == nil {
		return guardrailsConfig{}
	}
	config, _ := connection.Config.(guardrailsConfig)
	return config
}
