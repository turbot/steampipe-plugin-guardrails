package turbot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsPolicyValue(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_policy_value",
		Description: "Policy value define the value of policy known to Turbot Guardrails.",
		List: &plugin.ListConfig{
			Hydrate: listPolicyValue,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "state", Require: plugin.Optional},
				{Name: "policy_type_id", Require: plugin.Optional},
				{Name: "resource_id", Require: plugin.Optional},
				{Name: "resource_type_id", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the policy value.", Hydrate: policyValueHydrateId},
			{Name: "policy_type_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the policy type.", Hydrate: policyValueHydratePolicyTypeTitle},
			{Name: "policy_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title with full path of the policy type.", Hydrate: policyValueHydratePolicyTypeTrunkTitle},
			{Name: "is_default", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "If true this value is derived from the default value of the type.", Hydrate: policyValueHydrateIsDefault},
			{Name: "is_calculated", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "If true this value is derived from calculated setting inputs e.g. templateInput and template.", Hydrate: policyValueHydrateIsCalculated},
			{Name: "precedence", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Precedence of the setting: REQUIRED or RECOMMENDED.", Hydrate: policyValueHydratePrecedence},
			{Name: "resource_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource for the policy value.", Hydrate: policyValueHydrateResourceId},
			{Name: "resource_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the resource.", Hydrate: policyValueHydrateResourceTrunkTitle},
			{Name: "resource_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource type for this policy setting.", Hydrate: policyValueHydrateResourceTypeId},
			{Name: "state", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "State of the policy value.", Hydrate: policyValueHydrateState},
			{Name: "secret_value", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(convToString), Description: "Secrect value of the policy value.", Hydrate: policyValueHydrateSecretValue},
			{Name: "value", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(convToString), Description: "Value of the policy value.", Hydrate: policyValueHydrateValue},
			{Name: "type_mod_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the mod that contains the policy value.", Hydrate: policyValueHydrateTypeModUri},

			// Other columns
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used for this policy value list."},
			{Name: "policy_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the policy type for this policy value.", Hydrate: policyValueHydratePolicyTypeId},
			{Name: "policy_type_default_template", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Default template used to calculate template-based policy values. Should be a Jinja based YAML string.", Hydrate: policyValueHydratePolicyTypeDefaultTemplate},
			{Name: "setting_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().Transform(transform.NullIfZeroValue), Description: "Policy setting Id for the policy value.", Hydrate: policyValueHydrateSettingId},
			{Name: "dependent_controls", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "The controls that depends on this policy value.", Hydrate: policyValueHydrateDependentControls},
			{Name: "dependent_policy_values", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "The policy values that depends on this policy value.", Hydrate: policyValueHydrateDependentPolicyValues},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the policy value was first set by Turbot. (It may have been created earlier.)", Hydrate: policyValueHydrateCreateTimestamp},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the policy value was last modified (created, updated or deleted).", Hydrate: policyValueHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the policy value was last updated in Turbot.", Hydrate: policyValueHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the policy value.", Hydrate: policyValueHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: getTurbotGuardrailsWorkspace, Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryPolicyValueList = `
query MyQuery($filter: [String!], $next_token: String, $includePolicyValueDefault: Boolean!, $includePolicyValue: Boolean!, $includePolicyValueState: Boolean!, $includePolicyValueSecretValue: Boolean!, $includePolicyValueIsCalculated: Boolean!, $includePolicyValuePrecedence: Boolean!, $includePolicyValueTypeModUri: Boolean!, $includePolicyValueTypeDefaultTemplate: Boolean!, $includePolicyValueTypeTitle: Boolean!, $includePolicyValueTypeTrunkTitle: Boolean!, $includePolicyValueResourceTrunkTitle: Boolean!, $includePolicyValueTurbotId: Boolean!, $includePolicyValueTurbotPolicyTypeId: Boolean!, $includePolicyValueTurbotResourceId: Boolean!, $includePolicyValueTurbotResourceTypeId: Boolean!, $includePolicyValueTurbotSettingId: Boolean!, $includePolicyValueTurbotCreateTimestamp: Boolean!, $includePolicyValueTurbotTimestamp: Boolean!, $includePolicyValueTurbotUpdateTimestamp: Boolean!, $includePolicyValueTurbotVersionId: Boolean!, $includePolicyValueDependentControls: Boolean!, $includePolicyValueDependentPolicyValues: Boolean!) {
  policyValues(filter: $filter, paging: $next_token) {
    items {
      default @include(if: $includePolicyValueDefault)
      value @include(if: $includePolicyValue)
      state @include(if: $includePolicyValueState)
      secretValue @include(if: $includePolicyValueSecretValue)
      isCalculated @include(if: $includePolicyValueIsCalculated)
      precedence @include(if: $includePolicyValuePrecedence)
      type {
        modUri @include(if: $includePolicyValueTypeModUri)
        defaultTemplate @include(if: $includePolicyValueTypeDefaultTemplate)
        title @include(if: $includePolicyValueTypeTitle)
        trunk {
          title @include(if: $includePolicyValueTypeTrunkTitle)
        }
      }
      resource {
        trunk {
          title @include(if: $includePolicyValueResourceTrunkTitle)
        }
      }
      turbot {
        id @include(if: $includePolicyValueTurbotId)
        policyTypeId @include(if: $includePolicyValueTurbotPolicyTypeId)
        resourceId @include(if: $includePolicyValueTurbotResourceId)
        resourceTypeId @include(if: $includePolicyValueTurbotResourceTypeId)
        settingId @include(if: $includePolicyValueTurbotSettingId)
        createTimestamp @include(if: $includePolicyValueTurbotCreateTimestamp)
        timestamp @include(if: $includePolicyValueTurbotTimestamp)
        updateTimestamp @include(if: $includePolicyValueTurbotUpdateTimestamp)
        versionId @include(if: $includePolicyValueTurbotVersionId)
      }
      dependentControls @include(if: $includePolicyValueDependentControls) {
        items {
          turbot {
            controlTypeId
            controlTypePath
            controlCategoryId
            controlCategoryPath
            id
            resourceId
            resourceTypeId
          }
          type {
            modUri
            title
            trunk {
              title
            }
          }
        }
      }
      dependentPolicyValues @include(if: $includePolicyValueDependentPolicyValues) {
        items {
          type {
            modUri
            uri
            title
            trunk {
              title
            }
            turbot {
              id
              title
            }
          }
        }
      }
    }
    paging {
      next
    }
  }
}
`
)

func listPolicyValue(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_policy_type.listPolicyType", "connection_error", err)
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
	if quals["state"] != nil {
		filters = append(filters, fmt.Sprintf("state:%s ", getQualListValues(ctx, quals, "state", "string")))
	}

	if quals["policy_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("policyTypeId:%s policyTypeLevel:self", getQualListValues(ctx, quals, "policy_type_id", "int64")))
	}

	if quals["resource_id"] != nil {
		filters = append(filters, fmt.Sprintf("resourceId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_id", "int64")))
	}

	if quals["resource_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_type_id", "int64")))
	}

	// Setting a high limit and page all results
	var pageLimit int64 = 5000

	// Adjust page limit, if less than default value
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < pageLimit {
			pageLimit = *limit
		}
	}

	// Setting page limit
	filters = append(filters, fmt.Sprintf("limit:%s", strconv.Itoa(int(pageLimit))))

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}
	appendPolicyValueColumnIncludes(&variables, d.QueryContext.Columns)
	for {
		result := &PolicyValuesResponse{}
		err = conn.DoRequest(queryPolicyValueList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_policy_value.listPolicyValue", "query_error", err)
			return nil, err
		}
		for _, r := range result.PolicyValues.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if result.PolicyValues.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.PolicyValues.Paging.Next
	}
	return nil, nil
}
