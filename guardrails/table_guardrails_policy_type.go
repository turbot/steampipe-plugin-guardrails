package turbot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsPolicyType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_policy_type",
		Description: "Policy types define the types of controls known to Turbot Guardrails.",
		List: &plugin.ListConfig{
			Hydrate: listPolicyType,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "uri", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPolicyType,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the policy type.", Hydrate: policyTypeHydrateId},
			{Name: "uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the policy type.", Hydrate: policyTypeHydrateUri},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the policy type.", Hydrate: policyTypeHydrateTitle},
			{Name: "trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title with full path of the policy type.", Hydrate: policyTypeHydrateTrunkTitle},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Description of the policy type.", Hydrate: policyTypeHydrateDescription},
			{Name: "targets", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(emptyListIfNil), Description: "URIs of the resource types targeted by this policy type.", Hydrate: policyTypeHydrateTargets},
			// Other columns
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(emptyListIfNil), Description: "AKA (also known as) identifiers for the policy type.", Hydrate: policyTypeHydrateAkas},
			{Name: "category_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the control category for the policy type.", Hydrate: policyTypeHydrateCategoryId},
			{Name: "category_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the control category for the policy type.", Hydrate: policyTypeHydrateCategoryUri},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the policy type was first discovered by Turbot. (It may have been created earlier.)", Hydrate: policyTypeHydrateCreateTimestamp},
			{Name: "default_template", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Default template used to calculate template-based policy values. Should be a Jinja based YAML string.", Hydrate: policyTypeHydrateDefaultTemplate},
			{Name: "icon", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Icon of the policy type.", Hydrate: policyTypeHydrateIcon},
			{Name: "mod_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the mod that contains the policy type.", Hydrate: policyTypeHydrateModUri},
			{Name: "parent_id", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "ID for the parent of this policy type.", Hydrate: policyTypeHydrateParentId},
			{Name: "path", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(pathToArray), Description: "Hierarchy path with all identifiers of ancestors of the policy type.", Hydrate: policyTypeHydratePath},
			{Name: "read_only", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "If true user-defined policy settings are blocked from being created.", Hydrate: policyTypeHydrateReadOnly},
			{Name: "schema", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "JSON schema defining the allowed schema for policy values for any targeted resources.", Hydrate: policyTypeHydrateSchema},
			{Name: "secret", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "JSON schema defining valid values for the policy type.", Hydrate: policyTypeHydrateSecret},
			{Name: "secret_level", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Secret Level: SECRET, CONFIDENTIAL or NONE.", Hydrate: policyTypeHydrateSecretLevel},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the policy type was last updated in Turbot.", Hydrate: policyTypeHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the policy type.", Hydrate: policyTypeHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: plugin.HydrateFunc(getTurbotGuardrailsWorkspace).WithCache(), Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryPolicyTypeList = `
query policyTypeList($filter: [String!], $next_token: String, $includePolicyTypeCategoryTurbotId: Boolean!, $includePolicyTypeCategoryUri: Boolean!, $includePolicyTypeDescription: Boolean!, $includePolicyTypeDefaultTemplate: Boolean!, $includePolicyTypeIcon: Boolean!, $includePolicyTypeModUri: Boolean!, $includePolicyTypeReadOnly: Boolean!, $includePolicyTypeSchema: Boolean!, $includePolicyTypeSecret: Boolean!, $includePolicyTypeSecretLevel: Boolean!, $includePolicyTypeTargets: Boolean!, $includePolicyTypeTitle: Boolean!, $includePolicyTypeTrunkTitle: Boolean!, $includePolicyTypeTurbotAkas: Boolean!, $includePolicyTypeTurbotCreateTimestamp: Boolean!, $includePolicyTypeTurbotId: Boolean!, $includePolicyTypeTurbotParentId: Boolean!, $includePolicyTypeTurbotPath: Boolean!, $includePolicyTypeTurbotUpdateTimestamp: Boolean!, $includePolicyTypeTurbotVersionId: Boolean!, $includePolicyTypeUri: Boolean!) {
  policyTypes(filter: $filter, paging: $next_token) {
    items {
      category {
        turbot {
          id @include(if: $includePolicyTypeCategoryTurbotId)
        }
        uri @include(if: $includePolicyTypeCategoryUri)
      }
      description @include(if: $includePolicyTypeDescription)
      defaultTemplate @include(if: $includePolicyTypeDefaultTemplate)
      icon @include(if: $includePolicyTypeIcon)
      modUri @include(if: $includePolicyTypeModUri)
      readOnly @include(if: $includePolicyTypeReadOnly)
      schema @include(if: $includePolicyTypeSchema)
      secret @include(if: $includePolicyTypeSecret)
      secretLevel @include(if: $includePolicyTypeSecretLevel)
      targets @include(if: $includePolicyTypeTargets)
      title @include(if: $includePolicyTypeTitle)
      trunk {
        title @include(if: $includePolicyTypeTrunkTitle)
      }
      turbot {
        akas @include(if: $includePolicyTypeTurbotAkas)
        createTimestamp @include(if: $includePolicyTypeTurbotCreateTimestamp)
        id @include(if: $includePolicyTypeTurbotId)
        parentId @include(if: $includePolicyTypeTurbotParentId)
        path @include(if: $includePolicyTypeTurbotPath)
        updateTimestamp @include(if: $includePolicyTypeTurbotUpdateTimestamp)
        versionId @include(if: $includePolicyTypeTurbotVersionId)
      }
      uri @include(if: $includePolicyTypeUri)
    }
    paging {
      next
    }
  }
}
`

	queryPolicyTypeGet = `
query policyTypeGet($id: ID!, $includePolicyTypeCategoryTurbotId: Boolean!, $includePolicyTypeCategoryUri: Boolean!, $includePolicyTypeDescription: Boolean!, $includePolicyTypeDefaultTemplate: Boolean!, $includePolicyTypeIcon: Boolean!, $includePolicyTypeModUri: Boolean!, $includePolicyTypeReadOnly: Boolean!, $includePolicyTypeSchema: Boolean!, $includePolicyTypeSecret: Boolean!, $includePolicyTypeSecretLevel: Boolean!, $includePolicyTypeTargets: Boolean!, $includePolicyTypeTitle: Boolean!, $includePolicyTypeTrunkTitle: Boolean!, $includePolicyTypeTurbotAkas: Boolean!, $includePolicyTypeTurbotCategoryId: Boolean!, $includePolicyTypeTurbotCreateTimestamp: Boolean!, $includePolicyTypeTurbotId: Boolean!, $includePolicyTypeTurbotParentId: Boolean!, $includePolicyTypeTurbotPath: Boolean!, $includePolicyTypeTurbotUpdateTimestamp: Boolean!, $includePolicyTypeTurbotVersionId: Boolean!, $includePolicyTypeUri: Boolean!) {
  policyType(id: $id) {
    category {
      turbot {
        id @include(if: $includePolicyTypeCategoryTurbotId)
      }
      uri @include(if: $includePolicyTypeCategoryUri)
    }
    description @include(if: $includePolicyTypeDescription)
    defaultTemplate @include(if: $includePolicyTypeDefaultTemplate)
    icon @include(if: $includePolicyTypeIcon)
    modUri @include(if: $includePolicyTypeModUri)
    readOnly @include(if: $includePolicyTypeReadOnly)
    schema @include(if: $includePolicyTypeSchema)
    secret @include(if: $includePolicyTypeSecret)
    secretLevel @include(if: $includePolicyTypeSecretLevel)
    targets @include(if: $includePolicyTypeTargets)
    title @include(if: $includePolicyTypeTitle)
    trunk {
      title @include(if: $includePolicyTypeTrunkTitle)
    }
    turbot {
      akas @include(if: $includePolicyTypeTurbotAkas)
      createTimestamp @include(if: $includePolicyTypeTurbotCreateTimestamp)
      id @include(if: $includePolicyTypeTurbotId)
      parentId @include(if: $includePolicyTypeTurbotParentId)
      path @include(if: $includePolicyTypeTurbotPath)
      updateTimestamp @include(if: $includePolicyTypeTurbotUpdateTimestamp)
      versionId @include(if: $includePolicyTypeTurbotVersionId)
    }
    uri @include(if: $includePolicyTypeUri)
  }
}
`
)

func listPolicyType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_policy_type.listPolicyType", "connection_error", err)
		return nil, err
	}

	filters := []string{}
	quals := d.EqualsQuals

	// Additional filters
	if quals["uri"] != nil {
		filters = append(filters, fmt.Sprintf("policyTypeId:%s policyTypeLevel:self", getQualListValues(ctx, quals, "uri", "string")))
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

	plugin.Logger(ctx).Debug("guardrails_policy_type.listPolicyType", "quals", quals)
	plugin.Logger(ctx).Debug("guardrails_policy_type.listPolicyType", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}
	appendPolicyTypeColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &PolicyTypesResponse{}
		err = conn.DoRequest(queryPolicyTypeList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_policy_type.listPolicyType", "query_error", err)
			return nil, err
		}
		for _, r := range result.PolicyTypes.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if result.PolicyTypes.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.PolicyTypes.Paging.Next
	}
	return nil, nil
}

func getPolicyType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_policy_type.getPolicyType", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	variables := map[string]interface{}{
		"id": id,
	}
	appendPolicyTypeColumnIncludes(&variables, d.QueryContext.Columns)

	result := &PolicyTypeResponse{}
	err = conn.DoRequest(queryPolicyTypeGet, variables, result)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_policy_type.getPolicyType", "query_error", err)
		return nil, err
	}
	return result.PolicyType, nil
}
