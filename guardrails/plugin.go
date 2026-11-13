package turbot

import (
	"context"

	"github.com/turbot/steampipe-plugin-guardrails/errors"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-guardrails",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: func(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
					return errors.NotFoundError(err)
				},
			},
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"guardrails_active_grant":   tableGuardrailsActiveGrant(ctx),
			"guardrails_control":        tableGuardrailsControl(ctx),
			"guardrails_control_type":   tableGuardrailsControlType(ctx),
			"guardrails_grant":          tableGuardrailsGrant(ctx),
			"guardrails_mod_version":    tableGuardrailsModVersion(ctx),
			"guardrails_notification":   tableGuardrailsNotification(ctx),
			"guardrails_policy_setting": tableGuardrailsPolicySetting(ctx),
			"guardrails_policy_type":    tableGuardrailsPolicyType(ctx),
			"guardrails_policy_value":   tableGuardrailsPolicyValue(ctx),
			"guardrails_query":          tableGuardrailsQuery(ctx),
			"guardrails_resource":       tableGuardrailsResource(ctx),
			"guardrails_resource_type":  tableGuardrailsResourceType(ctx),
			"guardrails_smart_folder":   tableGuardrailsSmartFolder(ctx),
			"guardrails_tag":            tableGuardrailsTag(ctx),
		},
	}
	return p
}
