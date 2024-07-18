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

func tableGuardrailsTag(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_tag",
		Description: "All tags discovered on cloud resources by Turbot Guardrails.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Optional},
				{Name: "key", Require: plugin.Optional},
				{Name: "value", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listTag,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the tag.", Hydrate: tagHydrateId},
			{Name: "key", Type: proto.ColumnType_STRING, Description: "Tag key.", Transform: transform.FromValue(), Hydrate: tagHydrateKey},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Tag value.", Transform: transform.FromValue(), Hydrate: tagHydrateValue},
			{Name: "resource_ids", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(tagResourcesToIdArray), Description: "Turbot IDs of resources with this tag.", Hydrate: tagHydrateResources},
			// Other columns
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the tag was first discovered by Turbot. (It may have been created earlier.)", Hydrate: tagHydrateCreateTimestamp},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used for this tag list."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the tag was last modified (created, updated or deleted).", Hydrate: tagHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the tag was last updated in Turbot.", Hydrate: tagHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the tag.", Hydrate: tagHydrateVersionID},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: getTurbotGuardrailsWorkspace, Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryTagList = `
query tagList($filter: [String!], $paging: String, $includeTagKey: Boolean!, $includeTagValue: Boolean!, $includeTagTurbotId: Boolean!, $includeTagTurbotTimestamp: Boolean!, $includeTagTurbotCreateTimestamp: Boolean!, $includeTagTurbotUpdateTimestamp: Boolean!, $includeTagTurbotVersionId: Boolean!, $includeTagResources: Boolean!) {
  tags(filter: $filter, paging: $paging) {
    items {
      key @include(if: $includeTagKey)
      value @include(if: $includeTagValue)
      turbot {
        id @include(if: $includeTagTurbotId)
        timestamp @include(if: $includeTagTurbotTimestamp)
        createTimestamp @include(if: $includeTagTurbotCreateTimestamp)
        updateTimestamp @include(if: $includeTagTurbotUpdateTimestamp)
        versionId @include(if: $includeTagTurbotVersionId)
      }
      resources @include(if: $includeTagResources) {
        items {
          turbot {
            id
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

func listTag(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_tag.listTag", "connection_error", err)
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
	if quals["key"] != nil {
		filters = append(filters, fmt.Sprintf("key:%s", getQualListValues(ctx, quals, "key", "string")))
	}
	if quals["value"] != nil {
		filters = append(filters, fmt.Sprintf("value:%s", getQualListValues(ctx, quals, "value", "string")))
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

	plugin.Logger(ctx).Debug("guardrails_tag.listTag", "quals", quals)
	plugin.Logger(ctx).Debug("guardrails_tag.listTag", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	appendTagColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &TagsResponse{}
		err = conn.DoRequest(queryTagList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_tag.listTag", "query_error", err)
			// TODO - this is a bit risk and should not be necessary, but there is a
			// bug in Turbot where sometimes resource requests within the tags table fail
			//return nil, err
		}
		for _, r := range result.Tags.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !pageResults || result.Tags.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.Tags.Paging.Next
	}

	return nil, nil
}

func tagResourcesToIdArray(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	resources := d.Value.(TagResources)
	ids := []int64{}
	for _, r := range resources.Items {
		id, err := strconv.ParseInt(r.Turbot.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
