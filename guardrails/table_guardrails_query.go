package turbot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsQuery(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_query",
		Description: "Query Turbot Guardrails.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", Require: plugin.Required},
			},
			Hydrate: getQueryOutput,
		},
		Columns: []*plugin.Column{
			{Name: "output", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "The output of the query."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "The graphql query."},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: getTurbotGuardrailsWorkspace, Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

func getQueryOutput(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_query.getQueryOutput", "connection_error", err)
		return nil, err
	}

	query := d.EqualsQualString("query")

	var result interface{}
	err = conn.DoRequest(query, nil, &result)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_query.getQueryOutput", "query_error", err)
	}

	d.StreamListItem(ctx, result)

	return nil, nil
}
