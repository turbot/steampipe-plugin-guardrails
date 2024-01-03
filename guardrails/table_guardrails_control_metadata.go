package turbot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsControlMetadata(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_control_metadata",
		Description: "Controls show the current state of checks in the Turbot Guardrails workspace.",
		List: &plugin.ListConfig{
			Hydrate: listControlMetadata,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the control."},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: plugin.HydrateFunc(getTurbotGuardrailsWorkspace).WithCache(), Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	queryControlMetadataList = `
	query controlMetadataList($filter: [String!]) {
	controls(filter: $filter) {
		metadata {
      stats {
        total
        control {
          alarm
          error
          invalid
          ok
          skipped
          tbd
          total
        }
      }
    }
  }
}
`
)

func listControlMetadata(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_control_metadata.listControlMetadata", "connection_error", err)
		return nil, err
	}

	variables := map[string]interface{}{}

	//controlColumnIncludes(&variables, d.QueryContext.Columns)

	result := &ControlMetadataResponse{}
	err = conn.DoRequest(queryControlMetadataList, variables, result)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_control_metadata.listControlMetadata", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}
