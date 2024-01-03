package turbot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsControlType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_control_type",
		Description: "Control types define the types of controls known to Turbot Guardrails.",
		List: &plugin.ListConfig{
			Hydrate: listControlType,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "category_uri", Require: plugin.Optional},
				{Name: "uri", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getControlType,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the control type.", Hydrate: controlTypeHydrateId},
			{Name: "uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the control type.", Hydrate: controlTypeHydrateUri},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the control type.", Hydrate: controlTypeHydrateTitle},
			{Name: "trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title with full path of the control type.", Hydrate: controlTypeHydrateTrunkTitle},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Description of the control type.", Hydrate: controlTypeHydrateDescription},
			{Name: "targets", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "URIs of the resource types targeted by this control type.", Hydrate: controlTypeHydrateTargets},
			// Other columns
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "AKA (also known as) identifiers for the control type.", Hydrate: controlTypeHydrateAkas},
			{Name: "category_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the control category for the control type.", Hydrate: controlTypeHydrateCategoryId},
			{Name: "category_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the control category for the control type.", Hydrate: controlTypeHydrateCategoryUri},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the control type was first discovered by Turbot. (It may have been created earlier.)", Hydrate: controlTypeHydrateCreateTimestamp},
			{Name: "icon", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Icon of the control type.", Hydrate: controlTypeHydrateIcon},
			{Name: "mod_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the mod that contains the control type.", Hydrate: controlTypeHydrateModUri},
			{Name: "parent_id", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "ID for the parent of this control type.", Hydrate: controlTypeHydrateParentId},
			{Name: "path", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(pathToArray), Description: "Hierarchy path with all identifiers of ancestors of the control type.", Hydrate: controlTypeHydratePath},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the control type was last updated in Turbot.", Hydrate: controlTypeHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the control type.", Hydrate: controlTypeHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: plugin.HydrateFunc(getTurbotGuardrailsWorkspace).WithCache(), Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryControlTypeList = `
query controlTypeList($filter: [String!], $next_token: String, $includeControlTypeCategoryId: Boolean!, $includeControlTypeCategoryUri: Boolean!, $includeControlTypeDescription: Boolean!, $includeControlTypeIcon: Boolean!, $includeControlTypeModUri: Boolean!, $includeControlTypeTargets: Boolean!, $includeControlTypeTitle: Boolean!, $includeControlTypeTrunkTitle: Boolean!, $includeControlTypeTurbotAkas: Boolean!, $includeControlTypeTurbotCreateTimestamp: Boolean!, $includeControlTypeTurbotParentId: Boolean!, $includeControlTypeTurbotPath: Boolean!, $includeControlTypeTurbotUpdateTimestamp: Boolean!, $includeControlTypeTurbotVersionId: Boolean!, $includeControlTypeUri: Boolean!) {
  controlTypes(filter: $filter, paging: $next_token) {
    items {
      category {
        turbot {
          id @include(if: $includeControlTypeCategoryId)
        }
        uri @include(if: $includeControlTypeCategoryUri)
      }
      description @include(if: $includeControlTypeDescription)
      icon @include(if: $includeControlTypeIcon)
      modUri @include(if: $includeControlTypeModUri)
      targets @include(if: $includeControlTypeTargets)
      title @include(if: $includeControlTypeTitle)
      trunk {
        title @include(if: $includeControlTypeTrunkTitle)
      }
      turbot {
        akas @include(if: $includeControlTypeTurbotAkas)
        createTimestamp @include(if: $includeControlTypeTurbotCreateTimestamp)
        parentId @include(if: $includeControlTypeTurbotParentId)
        path @include(if: $includeControlTypeTurbotPath)
        updateTimestamp @include(if: $includeControlTypeTurbotUpdateTimestamp)
        versionId @include(if: $includeControlTypeTurbotVersionId)
      }
      uri @include(if: $includeControlTypeUri)
    }
    paging {
      next
    }
  }
}
`

	queryControlTypeGet = `
query controlTypeGet($id: ID!, $includeControlTypeCategoryId: Boolean!, $includeControlTypeCategoryUri: Boolean!, $includeControlTypeDescription: Boolean!, $includeControlTypeIcon: Boolean!, $includeControlTypeModUri: Boolean!, $includeControlTypeTargets: Boolean!, $includeControlTypeTitle: Boolean!, $includeControlTypeTrunkTitle: Boolean!, $includeControlTypeTurbotAkas: Boolean!, $includeControlTypeTurbotCreateTimestamp: Boolean!, $includeControlTypeTurbotParentId: Boolean!, $includeControlTypeTurbotPath: Boolean!, $includeControlTypeTurbotUpdateTimestamp: Boolean!, $includeControlTypeTurbotVersionId: Boolean!, $includeControlTypeUri: Boolean!) {
  controlType(id: $id) {
    category {
      turbot {
        id @include(if: $includeControlTypeCategoryId)
      }
      uri @include(if: $includeControlTypeCategoryUri)
    }
    description @include(if: $includeControlTypeDescription)
    icon @include(if: $includeControlTypeIcon)
    modUri @include(if: $includeControlTypeModUri)
    targets @include(if: $includeControlTypeTargets)
    title @include(if: $includeControlTypeTitle)
    trunk {
      title @include(if: $includeControlTypeTrunkTitle)
    }
    turbot {
      akas @include(if: $includeControlTypeTurbotAkas)
      createTimestamp @include(if: $includeControlTypeTurbotCreateTimestamp)
      parentId @include(if: $includeControlTypeTurbotParentId)
      path @include(if: $includeControlTypeTurbotPath)
      updateTimestamp @include(if: $includeControlTypeTurbotUpdateTimestamp)
      versionId @include(if: $includeControlTypeTurbotVersionId)
    }
    uri @include(if: $includeControlTypeUri)
  }
}
`
)

func listControlType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_control_type.listControlType", "connection_error", err)
		return nil, err
	}

	filters := []string{}
	quals := d.EqualsQuals

	// Additional filters
	if quals["uri"] != nil {
		filters = append(filters, fmt.Sprintf("controlTypeId:%s controlTypeLevel:self", getQualListValues(ctx, quals, "uri", "string")))
	}

	if quals["category_uri"] != nil {
		filters = append(filters, fmt.Sprintf("controlCategory:%s", getQualListValues(ctx, quals, "category_uri", "string")))
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

	plugin.Logger(ctx).Debug("guardrails_control_type.listControlType", "quals", quals)
	plugin.Logger(ctx).Debug("guardrails_control_type.listControlType", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	appendControlTypeColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &ControlTypesResponse{}
		err = conn.DoRequest(queryControlTypeList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_control_type.listControlType", "query_error", err)
			return nil, err
		}
		for _, r := range result.ControlTypes.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if result.ControlTypes.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.ControlTypes.Paging.Next
	}

	return nil, nil
}

func getControlType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_control_type.getControlType", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	variables := map[string]interface{}{
		"id": id,
	}

	appendControlTypeColumnIncludes(&variables, d.QueryContext.Columns)

	result := &ControlTypeResponse{}
	err = conn.DoRequest(queryControlTypeGet, variables, result)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_control_type.getControlType", "query_error", err)
		return nil, err
	}
	return result.ControlType, nil
}
