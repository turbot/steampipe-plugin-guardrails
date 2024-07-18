package turbot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsResourceType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_resource_type",
		Description: "Resource types define the types of resources known to Turbot Guardrails.",
		List: &plugin.ListConfig{
			Hydrate: listResourceType,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "category_uri", Require: plugin.Optional},
				{Name: "uri", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getResourceType,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the resource type.", Hydrate: resourceTypeHydrateId},
			{Name: "uri", Type: proto.ColumnType_STRING, Description: "URI of the resource type.", Transform: transform.FromValue(), Hydrate: resourceTypeHydrateUri},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the resource type.", Transform: transform.FromValue(), Hydrate: resourceTypeHydrateTitle},
			{Name: "trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title with full path of the resource type.", Hydrate: resourceTypeHydrateTrunkTitle},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the resource type.", Transform: transform.FromValue(), Hydrate: resourceTypeHydrateDescription},
			// Other columns
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "AKA (also known as) identifiers for the resource type.", Hydrate: resourceTypeHydrateAkas},
			{Name: "category_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource category for the resource type.", Hydrate: resourceTypeHydrateCategoryId},
			{Name: "category_uri", Type: proto.ColumnType_STRING, Description: "URI of the resource category for the resource type.", Transform: transform.FromValue(), Hydrate: resourceTypeHydrateCategoryUri},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the resource type was first discovered by Turbot. (It may have been created earlier.)", Hydrate: resourceTypeHydrateCreateTimestamp},
			{Name: "icon", Type: proto.ColumnType_STRING, Description: "Icon of the resource type.", Transform: transform.FromValue(), Hydrate: resourceTypeHydrateIcon},
			{Name: "mod_uri", Type: proto.ColumnType_STRING, Description: "URI of the mod that contains the resource type.", Transform: transform.FromValue(), Hydrate: resourceTypeHydrateModUri},
			{Name: "parent_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID for the parent of this resource type.", Hydrate: resourceTypeHydrateParentId},
			{Name: "path", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(pathToArray), Description: "Hierarchy path with all identifiers of ancestors of the resource type.", Hydrate: resourceTypeHydratePath},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the resource type was last updated in Turbot.", Hydrate: resourceTypeHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the resource type.", Hydrate: resourceTypeHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: getTurbotGuardrailsWorkspace, Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryResourceTypeList = `
query resourceTypeList($filter: [String!], $next_token: String, $includeResourceTypeCategoryId: Boolean!, $includeResourceTypeCategoryUri: Boolean!, $includeResourceTypeDescription: Boolean!, $includeResourceTypeIcon: Boolean!, $includeResourceTypeModUri: Boolean!, $includeResourceTypeTitle: Boolean!, $includeResourceTypeTrunkTitle: Boolean!, $includeResourceTypeAkas: Boolean!, $includeResourceTypeCreateTimestamp: Boolean!, $includeResourceTypeId: Boolean!, $includeResourceTypeParentId: Boolean!, $includeResourceTypePath: Boolean!, $includeResourceTypeUpdateTimestamp: Boolean!, $includeResourceTypeVersionId: Boolean!, $includeResourceTypeUri: Boolean!) {
  resourceTypes(filter: $filter, paging: $next_token) {
    items {
      category @include(if: $includeResourceTypeCategoryId) {
        turbot {
          id
        }
      }
      categoryUri @include(if: $includeResourceTypeCategoryUri)
      description @include(if: $includeResourceTypeDescription)
      icon @include(if: $includeResourceTypeIcon)
      modUri @include(if: $includeResourceTypeModUri)
      title @include(if: $includeResourceTypeTitle)
      trunk @include(if: $includeResourceTypeTrunkTitle) {
        title
      }
      turbot {
        akas @include(if: $includeResourceTypeAkas)
        createTimestamp @include(if: $includeResourceTypeCreateTimestamp)
        id @include(if: $includeResourceTypeId)
        parentId @include(if: $includeResourceTypeParentId)
        path @include(if: $includeResourceTypePath)
        updateTimestamp @include(if: $includeResourceTypeUpdateTimestamp)
        versionId @include(if: $includeResourceTypeVersionId)
      }
      uri @include(if: $includeResourceTypeUri)
    }
    paging {
      next
    }
  }
}
`

	queryResourceTypeGet = `
query resourceGet($id: ID!, $includeResourceTypeCategoryId: Boolean!, $includeResourceTypeCategoryUri: Boolean!, $includeResourceTypeDescription: Boolean!, $includeResourceTypeIcon: Boolean!, $includeResourceTypeModUri: Boolean!, $includeResourceTypeTitle: Boolean!, $includeResourceTypeTrunkTitle: Boolean!, $includeResourceTypeAkas: Boolean!, $includeResourceTypeCreateTimestamp: Boolean!, $includeResourceTypeId: Boolean!, $includeResourceTypeParentId: Boolean!, $includeResourceTypePath: Boolean!, $includeResourceTypeUpdateTimestamp: Boolean!, $includeResourceTypeVersionId: Boolean!, $includeResourceTypeUri: Boolean!) {
  resourceType(id: $id) {
    category @include(if: $includeResourceTypeCategoryId) {
      turbot {
        id
      }
    }
    categoryUri @include(if: $includeResourceTypeCategoryUri)
    description @include(if: $includeResourceTypeDescription)
    icon @include(if: $includeResourceTypeIcon)
    modUri @include(if: $includeResourceTypeModUri)
    title @include(if: $includeResourceTypeTitle)
    trunk @include(if: $includeResourceTypeTrunkTitle) {
      title
    }
    turbot {
      akas @include(if: $includeResourceTypeAkas)
      createTimestamp @include(if: $includeResourceTypeCreateTimestamp)
      id @include(if: $includeResourceTypeId)
      parentId @include(if: $includeResourceTypeParentId)
      path @include(if: $includeResourceTypePath)
      updateTimestamp @include(if: $includeResourceTypeUpdateTimestamp)
      versionId @include(if: $includeResourceTypeVersionId)
    }
    uri @include(if: $includeResourceTypeUri)
  }
}
`
)

func listResourceType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_resource_type.listResourceType", "connection_error", err)
		return nil, err
	}

	filters := []string{}
	quals := d.EqualsQuals

	// Additional filters
	if quals["uri"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "uri", "string")))
	}

	if quals["category_uri"] != nil {
		filters = append(filters, fmt.Sprintf("resourceCategory:%s", getQualListValues(ctx, quals, "category_uri", "string")))
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

	plugin.Logger(ctx).Debug("guardrails_resource_type.listResourceType", "quals", quals)
	plugin.Logger(ctx).Debug("guardrails_resource_type.listResourceType", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	appendResourceTypeColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &ResourceTypesResponse{}
		err = conn.DoRequest(queryResourceTypeList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_resource_type.listResourceType", "query_error", err)
			return nil, err
		}
		for _, r := range result.ResourceTypes.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if result.ResourceTypes.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.ResourceTypes.Paging.Next
	}

	return nil, nil
}

func getResourceType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_resource_type.getResourceType", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	variables := map[string]interface{}{
		"id": id,
	}

	appendResourceTypeColumnIncludes(&variables, d.QueryContext.Columns)

	result := &ResourceTypeResponse{}
	err = conn.DoRequest(queryResourceTypeGet, variables, result)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_resource_type.getResourceType", "query_error", err)
		return nil, err
	}
	return result.ResourceType, nil
}
