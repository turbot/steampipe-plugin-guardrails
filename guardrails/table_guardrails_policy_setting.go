package turbot

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsPolicySetting(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_policy_setting",
		Description: "Policy settings defined in the Turbot Guardrails workspace.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Optional},
				{Name: "resource_id", Require: plugin.Optional},
				{Name: "policy_type_id", Require: plugin.Optional},
				{Name: "policy_type_uri", Require: plugin.Optional},
				{Name: "orphan", Require: plugin.Optional},
				{Name: "exception", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listPolicySetting,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the policy setting.", Hydrate: policySettingHydrateId},
			{Name: "precedence", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Precedence of the setting: REQUIRED or RECOMMENDED.", Hydrate: policySettingHydratePrecedence},
			{Name: "resource_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource this policy setting is associated with.", Hydrate: policySettingHydrateResourceId},
			{Name: "resource_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the resource.", Hydrate: policySettingHydrateResourceTrunkTitle},
			{Name: "policy_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the policy type for this policy setting.", Hydrate: policySettingHydratePolicyTypeUri},
			{Name: "policy_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the policy type.", Hydrate: policySettingHydratePolicyTypeTrunkTitle},
			{Name: "value", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Value of the policy setting (for non-calculated policy settings).", Hydrate: policySettingHydrateValue},
			{Name: "is_calculated", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "True if this is a policy setting will be calculated for each value.", Hydrate: policySettingHydrateIsCalculated},
			{Name: "exception", Type: proto.ColumnType_BOOL, Transform: transform.FromValue().Transform(intToBool), Description: "True if this setting is an exception to a higher level setting.", Hydrate: policySettingHydrateException},
			{Name: "orphan", Type: proto.ColumnType_BOOL, Transform: transform.FromValue().Transform(intToBool), Description: "True if this setting is orphaned by a higher level setting.", Hydrate: policySettingHydrateOrphan},
			{Name: "note", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Optional note or comment for the setting.", Hydrate: policySettingHydrateNote},
			// Other columns
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the policy setting was first discovered by Turbot. (It may have been created earlier.)", Hydrate: policySettingHydrateCreateTimestamp},
			{Name: "default", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "True if this policy setting is the default.", Hydrate: policySettingHydrateDefault},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used for this policy setting list."},
			{Name: "input", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "For calculated policy settings, this is the input GraphQL query.", Hydrate: policySettingHydrateInput},
			{Name: "policy_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the policy type for this policy setting.", Hydrate: policySettingHydratePolicyTypeId},

			{Name: "template", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "For a calculated policy setting, this is the nunjucks template string defining a YAML string which is parsed to get the value.", Hydrate: policySettingHydrateTemplate},
			{Name: "template_input", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "For calculated policy settings, this GraphQL query is run and used as input to the template.", Hydrate: policySettingHydrateTemplateInput},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the policy setting was last modified (created, updated or deleted).", Hydrate: policySettingHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the policy setting was last updated in Turbot.", Hydrate: policySettingHydrateUpdateTimestamp},
			{Name: "valid_from_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the policy setting becomes valid.", Hydrate: policySettingHydrateValidFromTimestamp},
			{Name: "valid_to_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the policy setting expires.", Hydrate: policySettingHydrateValidToTimestamp},
			{Name: "value_source", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The raw value in YAML format. If the setting was made via YAML template including comments, these will be included here.", Hydrate: policySettingHydrateValueSource},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the policy setting.", Hydrate: policySettingHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: getTurbotGuardrailsWorkspace, Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryPolicySettingList = `
query policySettingList($filter: [String!], $next_token: String, $includePolicySettingDefault: Boolean!, $includePolicySettingException: Boolean!, $includePolicySettingInput: Boolean!, $includePolicySettingIsCalculated: Boolean!, $includePolicySettingNote: Boolean!, $includePolicySettingOrphan: Boolean!, $includePolicySettingPrecedence: Boolean!, $includePolicySettingResourceTrunkTitle: Boolean!, $includePolicySettingTemplate: Boolean!, $includePolicySettingTemplateInput: Boolean!, $includePolicySettingTypeUri: Boolean!, $includePolicySettingTypeTrunkTitle: Boolean!, $includePolicySettingTurbotId: Boolean!, $includePolicySettingTurbotTimestamp: Boolean!, $includePolicySettingTurbotCreateTimestamp: Boolean!, $includePolicySettingTurbotUpdateTimestamp: Boolean!, $includePolicySettingTurbotVersionId: Boolean!, $includePolicySettingTurbotPolicyTypeId: Boolean!, $includePolicySettingTurbotResourceId: Boolean!, $includePolicySettingValidFromTimestamp: Boolean!, $includePolicySettingValidToTimestamp: Boolean!, $includePolicySettingValue: Boolean!, $includePolicySettingValueSource: Boolean!) {
  policySettings(filter: $filter, paging: $next_token) {
    items {
      default @include(if: $includePolicySettingDefault)
      exception @include(if: $includePolicySettingException)
      input @include(if: $includePolicySettingInput)
      isCalculated @include(if: $includePolicySettingIsCalculated)
      note @include(if: $includePolicySettingNote)
      orphan @include(if: $includePolicySettingOrphan)
      precedence @include(if: $includePolicySettingPrecedence)
      resource {
        trunk {
          title @include(if: $includePolicySettingResourceTrunkTitle)
        }
      }
      template @include(if: $includePolicySettingTemplate)
      templateInput @include(if: $includePolicySettingTemplateInput)
      type {
        uri @include(if: $includePolicySettingTypeUri)
        trunk {
          title @include(if: $includePolicySettingTypeTrunkTitle)
        }
      }
      turbot {
        id @include(if: $includePolicySettingTurbotId)
        timestamp @include(if: $includePolicySettingTurbotTimestamp)
        createTimestamp @include(if: $includePolicySettingTurbotCreateTimestamp)
        updateTimestamp @include(if: $includePolicySettingTurbotUpdateTimestamp)
        versionId @include(if: $includePolicySettingTurbotVersionId)
        policyTypeId @include(if: $includePolicySettingTurbotPolicyTypeId)
        resourceId @include(if: $includePolicySettingTurbotResourceId)
      }
      validFromTimestamp @include(if: $includePolicySettingValidFromTimestamp)
      validToTimestamp @include(if: $includePolicySettingValidToTimestamp)
      value @include(if: $includePolicySettingValue)
      valueSource @include(if: $includePolicySettingValueSource)
    }
    paging {
      next
    }
  }
}
`
)

func listPolicySetting(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_policy_setting.listPolicySetting", "connection_error", err)
		return nil, err
	}

	filters := []string{}
	quals := d.EqualsQuals

	filter := ""
	if quals["filter"] != nil {
		filter = quals["filter"].GetStringValue()
		filters = append(filters, filter)
	}

	// Additional filters
	if quals["id"] != nil {
		filters = append(filters, fmt.Sprintf("id:%s", getQualListValues(ctx, quals, "id", "int64")))
	}

	if quals["policy_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("policyTypeId:%s policyTypeLevel:self", getQualListValues(ctx, quals, "policy_type_id", "int64")))
	}

	if quals["policy_type_uri"] != nil {
		filters = append(filters, fmt.Sprintf("policyTypeId:%s policyTypeLevel:self", getQualListValues(ctx, quals, "policy_type_uri", "string")))
	}

	if quals["resource_id"] != nil {
		filters = append(filters, fmt.Sprintf("resourceId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_id", "int64")))
	}

	if quals["orphan"] != nil {
		orphan := quals["orphan"].GetBoolValue()
		if orphan {
			filters = append(filters, "is:orphan")
		} else {
			filters = append(filters, "-is:orphan")
		}
	}

	if quals["exception"] != nil {
		exception := quals["exception"].GetBoolValue()
		if exception {
			filters = append(filters, "is:exception")
		} else {
			filters = append(filters, "-is:exception")
		}
	}

	// Default to a very large page size. Page sizes earlier in the filter string
	// win, so this is only used as a fallback.
	pageResults := false
	// Add a limit if they haven't given one in the filter field
	re := regexp.MustCompile(`(^|\s)limit:[0-9]+($|\s)`)
	if !re.MatchString(filter) {
		// The caller did not specify a limit, so set a high limit and page all
		// results.
		pageResults = true
		var pageLimit int64 = 5000

		// Adjust page limit, if less than default value
		limit := d.QueryContext.Limit
		if d.QueryContext.Limit != nil {
			if *limit < pageLimit {
				pageLimit = *limit
			}
		}
		filters = append(filters, fmt.Sprintf("limit:%s", strconv.Itoa(int(pageLimit))))
	}

	plugin.Logger(ctx).Debug("guardrails_policy_setting.listPolicySetting", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}
	appendPolicySettingColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &PolicySettingsResponse{}
		err = conn.DoRequest(queryPolicySettingList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_policy_setting.listPolicySetting", "query_error", err)
			return nil, err
		}
		for _, r := range result.PolicySettings.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !pageResults || result.PolicySettings.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.PolicySettings.Paging.Next
	}

	return nil, nil
}
