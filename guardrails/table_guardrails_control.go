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

func tableGuardrailsControl(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_control",
		Description: "Controls show the current state of checks in the Turbot Guardrails workspace.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Optional},
				{Name: "control_type_id", Require: plugin.Optional},
				{Name: "control_type_uri", Require: plugin.Optional},
				{Name: "resource_type_id", Require: plugin.Optional},
				{Name: "resource_type_uri", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listControl,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the control.", Hydrate: controlHydrateId},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the control.", Transform: transform.FromValue(), Hydrate: controlHydrateState},
			{Name: "reason", Type: proto.ColumnType_STRING, Description: "Reason for this control state.", Transform: transform.FromValue(), Hydrate: controlHydrateReason},
			{Name: "details", Type: proto.ColumnType_JSON, Description: "Details associated with this control state.", Transform: transform.FromValue(), Hydrate: controlHydrateDetails},
			{Name: "resource_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource this control is associated with.", Hydrate: controlHydrateResourceId},
			{Name: "resource_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the resource.", Hydrate: controlHydrateResourceTrunkTitle},
			{Name: "control_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the control type.", Hydrate: controlHydrateControlTypeTrunkTitle},

			// Other columns
			{Name: "control_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the control type for this control.", Hydrate: controlHydrateControlTypeId},
			{Name: "control_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the control type for this control.", Hydrate: controlHydrateControlTypeUri},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the control was first discovered by Turbot. (It may have been created earlier.)", Hydrate: controlHydrateCreateTimestamp},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used for this control list."},
			{Name: "resource_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the resource type for this control.", Hydrate: controlHydrateResourceTypeId},
			{Name: "resource_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the resource type for this control.", Hydrate: controlHydrateResourceTypeUri},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Timestamp when the control was last modified (created, updated or deleted).", Hydrate: controlHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the control was last updated in Turbot.", Hydrate: controlHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier for this version of the control.", Hydrate: controlHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: plugin.HydrateFunc(getTurbotGuardrailsWorkspace).WithCache(), Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Hydrate: controlHydrateMetadata, Transform: transform.FromValue(), Description: "The control metadata."},
		},
	}
}

const (
	queryControlList = `
	query controlList($filter: [String!], $next_token: String, $includeControlState: Boolean!, $includeControlReason: Boolean!, $includeControlDetails: Boolean!, $includeControlResourceTypeUri: Boolean!, $includeControlResourceTrunkTitle: Boolean!, $includeControlTypeUri: Boolean!, $includeControlTypeTrunkTitle: Boolean!, $includeControlId: Boolean!, $includeControlTimestamp: Boolean!, $includeControlCreateTimestamp: Boolean!, $includeControlUpdateTimestamp: Boolean!, $includeControlVersionId: Boolean!, $includeControlTypeId: Boolean!, $includeControlResourceId: Boolean!, $includeControlResourceTypeId: Boolean!, $includeControlMetadata: Boolean!, $includeControlItems: Boolean!) {
	controls(filter: $filter, paging: $next_token) {
		metadata @include(if: $includeControlMetadata){
      stats {
        total
      }
    }
		items @include(if: $includeControlItems){
			state @include(if: $includeControlState)
			reason @include(if: $includeControlReason)
			details @include(if: $includeControlDetails)
			resource {
				type {
					uri @include(if: $includeControlResourceTypeUri)
				}
				trunk {
					title @include(if: $includeControlResourceTrunkTitle)
				}
			}
			type {
				uri @include(if: $includeControlTypeUri)
				trunk {
					title @include(if: $includeControlTypeTrunkTitle)
				}
			}
			turbot {
				id @include(if: $includeControlId)
				timestamp @include(if: $includeControlTimestamp)
				createTimestamp @include(if: $includeControlCreateTimestamp)
				updateTimestamp @include(if: $includeControlUpdateTimestamp)
				versionId @include(if: $includeControlVersionId)
				controlTypeId @include(if: $includeControlTypeId)
				resourceId @include(if: $includeControlResourceId)
				resourceTypeId @include(if: $includeControlResourceTypeId)
			}
		}
		paging {
			next
		}
	}
}
`
)

func listControl(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_control.listControl", "connection_error", err)
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
	if quals["control_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("controlTypeId:%s controlTypeLevel:self", getQualListValues(ctx, quals, "control_type_id", "int64")))
	}
	if quals["control_type_uri"] != nil {
		filters = append(filters, fmt.Sprintf("controlTypeId:%s controlTypeLevel:self", getQualListValues(ctx, quals, "control_type_uri", "string")))
	}
	if quals["resource_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_type_id", "int64")))
	}
	if quals["resource_type_uri"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_type_uri", "string")))
	}
	if quals["state"] != nil {
		filters = append(filters, fmt.Sprintf("state:%s", getQualListValues(ctx, quals, "state", "string")))
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

	plugin.Logger(ctx).Debug("guardrails_control.listControl", "quals", quals)
	plugin.Logger(ctx).Debug("guardrails_control.listControl", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	controlColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &ControlsResponse{}
		err = conn.DoRequest(queryControlList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_control.listControl", "query_error", err)
			return nil, err
		}
		if len(result.Controls.Items) == 0 {
			d.StreamListItem(ctx, ControlItem{result.Controls.Metadata, Control{}})
			break
		} else {
			for _, r := range result.Controls.Items {
				d.StreamListItem(ctx, ControlItem{result.Controls.Metadata, r})

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}
		if !pageResults || result.Controls.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.Controls.Paging.Next
	}

	return nil, nil
}

type ControlItem struct {
	Metadata interface{}
	Item     Control
}
