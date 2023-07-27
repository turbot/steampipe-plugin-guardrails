package turbot

import (
	"context"

	"github.com/turbot/steampipe-plugin-guardrails/errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-guardrails",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: errors.NotFoundError,
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"guardrails_active_grant":   tableTurbotActiveGrant(ctx),
			"guardrails_control":        tableTurbotControl(ctx),
			"guardrails_control_type":   tableTurbotControlType(ctx),
			"guardrails_grant":          tableTurbotGrant(ctx),
			"guardrails_mod_version":    tableTurbotModVersion(ctx),
			"guardrails_notification":   tableTurbotNotification(ctx),
			"guardrails_policy_setting": tableTurbotPolicySetting(ctx),
			"guardrails_policy_type":    tableTurbotPolicyType(ctx),
			"guardrails_policy_value":   tableTurbotPolicyValue(ctx),
			"guardrails_resource":       tableTurbotResource(ctx),
			"guardrails_resource_type":  tableTurbotResourceType(ctx),
			"guardrails_smart_folder":   tableTurbotSmartFolder(ctx),
			"guardrails_tag":            tableTurbotTag(ctx),
		},
	}
	return p
}
