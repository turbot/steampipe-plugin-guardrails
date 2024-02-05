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

func tableGuardrailsGrant(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_grant",
		Description: "All grants of resources by Turbot Guardrails.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listGrants,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the grant.", Hydrate: grantHydrateGrantId},
			{Name: "resource_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the resource.", Hydrate: grantHydrateResourceId},
			{Name: "identity_status", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Status of the identity.", Hydrate: grantHydrateIdentityStatus},
			{Name: "identity_display_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Display name of the identity.", Hydrate: grantHydrateIdentityDisplayName},
			{Name: "identity_email", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Email identity for the identity.", Hydrate: grantHydrateIdentityEmail},
			{Name: "identity_family_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Family name of the identity.", Hydrate: grantHydrateIdentityFamilyName},
			{Name: "identity_given_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Given name of the identity.", Hydrate: grantHydrateIdentityGivenName},
			{Name: "identity_last_login_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Last login timestamp.", Hydrate: grantHydrateIdentityLastLoginTimestamp},
			{Name: "identity_profile_id", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Profile id of the identity.", Hydrate: grantHydrateIdentityProfileId},
			{Name: "identity_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the grant identity.", Hydrate: grantHydrateIdentityTrunkTitle},
			{Name: "level_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The title of the level.", Hydrate: grantHydrateLevelTitle},
			{Name: "level_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the level.", Hydrate: grantHydrateLevelTrunkTitle},
			{Name: "level_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The URI of the level.", Hydrate: grantHydrateLevelUri},
			{Name: "resource_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the grant type.", Hydrate: grantHydrateResourceTypeTrunkTitle},
			{Name: "resource_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the resource.", Hydrate: grantHydrateResourceTrunkTitle},
			{Name: "resource_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the resource type.", Hydrate: grantHydrateResourceTypeUri},
			{Name: "identity_akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "AKA (also known as) identifiers for the identity", Hydrate: grantHydrateIdentityAkas},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue().NullIfEqual(""), Description: "The create time of the grant.", Hydrate: grantHydrateCreateTimestamp},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used for this grant list."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue().NullIfEqual(""), Description: "Timestamp when the grant was last modified (created, updated or deleted).", Hydrate: grantHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the grant was last updated in Turbot.", Hydrate: grantHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().NullIfEqual(""), Description: "Unique identifier for this version of the identity.", Hydrate: grantHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: plugin.HydrateFunc(getTurbotGuardrailsWorkspace).WithCache(), Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	grants = `
	query MyQuery($filter: [String!], $paging: String, $includeGrantResourceAkas: Boolean!, $includeGrantResourceTitle: Boolean!, $includeGrantResourceTrunkTitle: Boolean!, $includeGrantResourceTypeURI: Boolean!, $includeGrantResourceTypeTrunkTitle: Boolean!, $includeGrantResourceId: Boolean!, $includeGrantResourceCreateTimestamp: Boolean!, $includeGrantResourceDeleteTimestamp: Boolean!, $includeGrantResourceTimestamp: Boolean!, $includeGrantResourceVersionId: Boolean!, $includeGrantResourceUpdateTimestamp: Boolean!, $includeGrantIdentityAkas: Boolean!, $includeGrantIdentityEmail: Boolean!, $includeGrantIdentityStatus: Boolean!, $includeGrantIdentityGivenName: Boolean!, $includeGrantIdentityProfileId: Boolean!, $includeGrantIdentityFamilyName: Boolean!, $includeGrantIdentityDisplayName: Boolean!, $includeGrantIdentityLastLoginTimestamp: Boolean!, $includeGrantIdentityTrunkTitle: Boolean!, $includeGrantLevelTitle: Boolean!, $includeGrantLevelURI: Boolean!, $includeGrantLevelTrunkTitle: Boolean!, $includeGrantTurbotId: Boolean!, $includeGrantTurbotCreateTimestamp: Boolean!, $includeGrantTurbotDeleteTimestamp: Boolean!, $includeGrantTurbotTimestamp: Boolean!, $includeGrantTurbotVersionId: Boolean!, $includeGrantTurbotUpdateTimestamp: Boolean!) {
		grants(filter: $filter, paging: $paging) {
		  items {
			resource {
			  akas @include(if: $includeGrantResourceAkas)
			  title @include(if: $includeGrantResourceTitle)
			  trunk {
				title @include(if: $includeGrantResourceTrunkTitle)
			  }
			  type {
				uri @include(if: $includeGrantResourceTypeURI)
				trunk {
				  title @include(if: $includeGrantResourceTypeTrunkTitle)
				}
			  }
			  turbot {
				id @include(if: $includeGrantResourceId)
				createTimestamp @include(if: $includeGrantResourceCreateTimestamp)
				deleteTimestamp @include(if: $includeGrantResourceDeleteTimestamp)
				timestamp @include(if: $includeGrantResourceTimestamp)
				versionId @include(if: $includeGrantResourceVersionId)
				updateTimestamp @include(if: $includeGrantResourceUpdateTimestamp)
			  }
			}
		identity {
			akas @include(if: $includeGrantIdentityAkas)
			email: get(path: "email") @include(if: $includeGrantIdentityEmail)
			status: get(path: "status") @include(if: $includeGrantIdentityStatus)
			givenName: get(path: "givenName") @include(if: $includeGrantIdentityGivenName)
			profileId: get(path: "profileId") @include(if: $includeGrantIdentityProfileId)
			familyName: get(path: "familyName") @include(if: $includeGrantIdentityFamilyName)
			displayName: get(path: "displayName") @include(if: $includeGrantIdentityDisplayName)
			lastLoginTimestamp: get(path: "lastLoginTimestamp") @include(if: $includeGrantIdentityLastLoginTimestamp)
			trunk {
			title @include(if: $includeGrantIdentityTrunkTitle)
			}
		}
		level {
			title @include(if: $includeGrantLevelTitle)
			uri @include(if: $includeGrantLevelURI)
			trunk {
			title @include(if: $includeGrantLevelTrunkTitle)
			}
		}
		turbot {
			id @include(if: $includeGrantTurbotId)
			createTimestamp @include(if: $includeGrantTurbotCreateTimestamp)
			deleteTimestamp @include(if: $includeGrantTurbotDeleteTimestamp)
			timestamp @include(if: $includeGrantTurbotTimestamp)
			versionId @include(if: $includeGrantTurbotVersionId)
			updateTimestamp @include(if: $includeGrantTurbotUpdateTimestamp)
		}
		}
		paging {
		next
	}
	}
	}
`
)

func listGrants(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_grants.listGrants", "connection_error", err)
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

	plugin.Logger(ctx).Debug("guardrails_grants.listGrants", "quals", quals)
	plugin.Logger(ctx).Debug("guardrails_grants.listGrants", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	appendGrantColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &GrantInfo{}
		err = conn.DoRequest(grants, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_grants.listGrants", "query_error", err)
		}
		for _, grantDetails := range result.Grants.Items {

			d.StreamListItem(ctx, grantDetails)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !pageResults || result.Grants.Paging.Next == "" {
			break
		}

		variables["next_token"] = result.Grants.Paging.Next
	}

	return nil, err
}
