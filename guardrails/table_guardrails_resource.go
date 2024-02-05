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

func tableGuardrailsResource(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_resource",
		Description: "Resources from the Turbot Guardrails CMDB.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Optional},
				{Name: "resource_type_id", Require: plugin.Optional},
				{Name: "resource_type_uri", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listResource,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the resource.", Hydrate: resourceHydrateId},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the resource.", Hydrate: resourceHydrateTitle},
			{Name: "trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title with full path of the resource.", Hydrate: resourceHydrateTrunkTitle},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "Tags for the resource.", Hydrate: resourceHydrateTags},
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "AKA (also known as) identifiers for the resource.", Hydrate: resourceHydrateAkas},
			// Other columns
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the resource was first discovered by Turbot. (It may have been created earlier.)", Hydrate: resourceHydrateCreateTimestamp},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "Resource data.", Transform: transform.FromValue(), Hydrate: resourceHydrateData},
			{Name: "object", Type: proto.ColumnType_JSON, Description: "Extended Resource data.", Transform: transform.FromValue(), Hydrate: resourceHydrateObject},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used for this resource list."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Resource custom metadata.", Transform: transform.FromValue(), Hydrate: resourceHydrateMetadata},
			{Name: "parent_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID for the parent of this resource. For the Turbot root resource this is null.", Hydrate: resourceHydrateParentId},
			{Name: "path", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(pathToArray), Description: "Hierarchy path with all identifiers of ancestors of the resource.", Hydrate: resourceHydratePath},
			{Name: "resource_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource type for this resource.", Hydrate: resourceHydrateResourceTypeId},
			{Name: "resource_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the resource type for this resource.", Hydrate: resourceHydrateResourceTypeUri},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the resource was last modified (created, updated or deleted).", Hydrate: resourceHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the resource was last updated in Turbot.", Hydrate: resourceHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the resource.", Hydrate: resourceHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: plugin.HydrateFunc(getTurbotGuardrailsWorkspace).WithCache(), Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryResourceList = `
query resourceList($filter: [String!], $next_token: String, $includeResourceObject: Boolean!, $includeResourceData: Boolean!, $includeResourceMetadata: Boolean!, $includeResourceTrunkTitle: Boolean!, $includeResourceId: Boolean!, $includeResourceTitle: Boolean!, $includeResourceTags: Boolean!, $includeResourceAkas: Boolean!, $includeResourceTimestamp: Boolean!, $includeResourceCreateTimestamp: Boolean!, $includeResourceUpdateTimestamp: Boolean!, $includeResourceVersionId: Boolean!, $includeResourceParentId: Boolean!, $includeResourcePath: Boolean!, $includeResourceTypeId: Boolean!, $includeResourceTypeUri: Boolean!) {
  resources(filter: $filter, paging: $next_token) {
    items {
      data @include(if: $includeResourceData)
      object @include(if: $includeResourceObject)
      metadata @include(if: $includeResourceMetadata)
      trunk {
        title @include(if: $includeResourceTrunkTitle)
      }
      turbot {
        id @include(if: $includeResourceId)
        title @include(if: $includeResourceTitle)
        tags @include(if: $includeResourceTags)
        akas @include(if: $includeResourceAkas)
        timestamp @include(if: $includeResourceTimestamp)
        createTimestamp @include(if: $includeResourceCreateTimestamp)
        updateTimestamp @include(if: $includeResourceUpdateTimestamp)
        versionId @include(if: $includeResourceVersionId)
        parentId @include(if: $includeResourceParentId)
        path @include(if: $includeResourcePath)
        resourceTypeId @include(if: $includeResourceTypeId)
      }
      type {
        uri @include(if: $includeResourceTypeUri)
      }
    }
    paging {
      next
    }
  }
}
`
)

func listResource(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_resource.listResource", "connection_error", err)
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
		filters = append(filters, fmt.Sprintf("resourceId:%s level:self", getQualListValues(ctx, quals, "id", "int64")))
	}
	if quals["resource_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_type_id", "int64")))
	}
	if quals["resource_type_uri"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_type_uri", "string")))
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

	plugin.Logger(ctx).Debug("guardrails_resource.listResource", "quals", quals)
	plugin.Logger(ctx).Debug("guardrails_resource.listResource", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	appendResourceColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &ResourcesResponse{}
		err = conn.DoRequest(queryResourceList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_resource.listResource", "query_error", err)
			return nil, err
		}
		for _, r := range result.Resources.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !pageResults || result.Resources.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.Resources.Paging.Next
	}

	return nil, nil
}
