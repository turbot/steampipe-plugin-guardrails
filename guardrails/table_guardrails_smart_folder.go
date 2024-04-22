package turbot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsSmartFolder(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_smart_folder",
		Description: "Smart folders allow policy settings to be attached as groups to resources.",
		List: &plugin.ListConfig{
			Hydrate: listSmartFolder,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSmartFolder,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the smart folder.", Hydrate: smartFolderHydrateId},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the smart folder.", Hydrate: smartFolderHydrateTitle},
			{Name: "trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title with full path of the smart folder.", Hydrate: smartFolderHydrateTrunkTitle},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromValue().TransformP(getMapValue, "description"), Description: "Description of the smart folder.", Hydrate: smartFolderHydrateData},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(emptyMapIfNil), Description: "Tags for the smart folder.", Hydrate: smartFolderHydrateTags},
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(emptyListIfNil), Description: "AKA (also known as) identifiers for the smart folder.", Hydrate: smartFolderHydrateAkas},
			{Name: "attached_resource_ids", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(attachedResourceIDs), Description: "", Hydrate: smartFolderHydrateAttachedResources},
			// Other columns
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the smart folder was first discovered by Turbot. (It may have been created earlier.)", Hydrate: smartFolderHydrateCreateTimestamp},
			{Name: "color", Type: proto.ColumnType_STRING, Transform: transform.FromValue().TransformP(getMapValue, "color"), Description: "Color of the smart folder in the UI.", Hydrate: smartFolderHydrateData},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "Resource data.", Hydrate: smartFolderHydrateData, Transform: transform.FromValue()},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Resource custom metadata.", Hydrate: smartFolderHydrateMetadata, Transform: transform.FromValue()},
			{Name: "parent_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID for the parent of this smart folder.", Hydrate: smartFolderHydrateParentId},
			{Name: "path", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(pathToArray), Description: "Hierarchy path with all identifiers of ancestors of the smart folder.", Hydrate: smartFolderHydratePath},
			{Name: "resource_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource type for this smart folder.", Hydrate: smartFolderHydrateResourceTypeId},
			{Name: "resource_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the resource type for this smart folder.", Hydrate: smartFolderHydrateTypeUri},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the smart folder was last modified (created, updated or deleted).", Hydrate: smartFolderHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the smart folder was last updated in Turbot.", Hydrate: smartFolderHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the smart folder.", Hydrate: smartFolderHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: getTurbotGuardrailsWorkspace, Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	querySmartFolderList = `
query smartFolderList($filter: [String!], $next_token: String, $includeAttachedResourcesId: Boolean!, $includeData: Boolean!, $includeMetadata: Boolean!, $includeTrunkTitle: Boolean!, $includeTurbotId: Boolean!, $includeTurbotTitle: Boolean!, $includeTurbotTags: Boolean!, $includeTurbotAkas: Boolean!, $includeTurbotTimestamp: Boolean!, $includeTurbotCreateTimestamp: Boolean!, $includeTurbotUpdateTimestamp: Boolean!, $includeTurbotVersionId: Boolean!, $includeTurbotParentId: Boolean!, $includeTurbotPath: Boolean!, $includeTurbotResourceTypeId: Boolean!, $includeTypeUri: Boolean!) {
  resources(filter: $filter, paging: $next_token) {
    items {
      attachedResources @include(if: $includeAttachedResourcesId) {
        items {
          turbot {
            id
          }
        }
      }
      data @include(if: $includeData)
      metadata @include(if: $includeMetadata)
      trunk {
        title @include(if: $includeTrunkTitle)
      }
      turbot {
        id @include(if: $includeTurbotId)
        title @include(if: $includeTurbotTitle)
        tags @include(if: $includeTurbotTags)
        akas @include(if: $includeTurbotAkas)
        timestamp @include(if: $includeTurbotTimestamp)
        createTimestamp @include(if: $includeTurbotCreateTimestamp)
        updateTimestamp @include(if: $includeTurbotUpdateTimestamp)
        versionId @include(if: $includeTurbotVersionId)
        parentId @include(if: $includeTurbotParentId)
        path @include(if: $includeTurbotPath)
        resourceTypeId @include(if: $includeTurbotResourceTypeId)
      }
      type {
        uri @include(if: $includeTypeUri)
      }
    }
    paging {
      next
    }
  }
}
`

	querySmartFolderGet = `
query smartFolderGet($id: ID!, $includeAttachedResourcesId: Boolean!, $includeData: Boolean!, $includeMetadata: Boolean!, $includeTrunkTitle: Boolean!, $includeTurbotId: Boolean!, $includeTurbotTitle: Boolean!, $includeTurbotTags: Boolean!, $includeTurbotAkas: Boolean!, $includeTurbotTimestamp: Boolean!, $includeTurbotCreateTimestamp: Boolean!, $includeTurbotUpdateTimestamp: Boolean!, $includeTurbotVersionId: Boolean!, $includeTurbotParentId: Boolean!, $includeTurbotPath: Boolean!, $includeTurbotResourceTypeId: Boolean!, $includeTypeUri: Boolean!) {
  resource(id: $id) {
    attachedResources @include(if: $includeAttachedResourcesId) {
      items {
        turbot {
          id
        }
      }
    }
    data @include(if: $includeData)
    metadata @include(if: $includeMetadata)
    trunk {
      title @include(if: $includeTrunkTitle)
    }
    turbot {
      id @include(if: $includeTurbotId)
      title @include(if: $includeTurbotTitle)
      tags @include(if: $includeTurbotTags)
      akas @include(if: $includeTurbotAkas)
      timestamp @include(if: $includeTurbotTimestamp)
      createTimestamp @include(if: $includeTurbotCreateTimestamp)
      updateTimestamp @include(if: $includeTurbotUpdateTimestamp)
      versionId @include(if: $includeTurbotVersionId)
      parentId @include(if: $includeTurbotParentId)
      path @include(if: $includeTurbotPath)
      resourceTypeId @include(if: $includeTurbotResourceTypeId)
    }
    type {
      uri @include(if: $includeTypeUri)
    }
  }
}
`
)

func listSmartFolder(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_smart_folder.listSmartFolder", "connection_error", err)
		return nil, err
	}

	var pageLimit int64 = 5000

	// Adjust page limit, if less than default value
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < pageLimit {
			pageLimit = *limit
		}
	}
	filter := fmt.Sprintf("resourceTypeId:'tmod:@turbot/turbot#/resource/types/smartFolder' resourceTypeLevel:self limit:%s", strconv.Itoa(int(pageLimit)))

	variables := map[string]interface{}{
		"filter":     filter,
		"next_token": "",
	}

	appendSmartFolderColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &ResourcesResponse{}
		err = conn.DoRequest(querySmartFolderList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_smart_folder.listSmartFolder", "query_error", err)
			return nil, err
		}
		for _, r := range result.Resources.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if result.Resources.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.Resources.Paging.Next
	}

	return nil, nil
}

func getSmartFolder(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_smart_folder.getSmartFolder", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	variables := map[string]interface{}{
		"id": id,
	}

	appendSmartFolderColumnIncludes(&variables, d.QueryContext.Columns)

	result := &ResourceResponse{}
	err = conn.DoRequest(querySmartFolderGet, variables, result)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_smart_folder.getSmartFolder", "query_error", err)
		return nil, err
	}
	return result.Resource, nil
}
