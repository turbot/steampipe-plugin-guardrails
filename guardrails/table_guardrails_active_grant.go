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

func tableGuardrailsActiveGrant(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_active_grant",
		Description: "All active grants of resources by Turbot Guardrails.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "grant_id", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listActiveGrants,
		},
		Columns: []*plugin.Column{
			{Name: "grant_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the grant.", Hydrate: activeGrantHydrateGrantId},
			{Name: "resource_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the resource.", Hydrate: activeGrantHydrateResourceId},
			{Name: "identity_status", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Status of the identity.", Hydrate: activeGrantHydrateIdentityStatus},
			{Name: "identity_display_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Display name of the identity.", Hydrate: activeGrantHydrateIdentityDisplayName},
			{Name: "identity_email", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Email identity for the identity.", Hydrate: activeGrantHydrateIdentityEmail},
			{Name: "identity_family_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Family name of the identity.", Hydrate: activeGrantHydrateIdentityFamilyName},
			{Name: "identity_given_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Given name of the identity.", Hydrate: activeGrantHydrateIdentityGivenName},
			{Name: "identity_last_login_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Last login timestamp.", Hydrate: activeGrantHydrateIdentityLastLoginTimestamp},
			{Name: "identity_profile_id", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Profile id of the identity.", Hydrate: activeGrantHydrateIdentityProfileId},
			{Name: "identity_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the grant identity.", Hydrate: activeGrantHydrateIdentityTrunkTitle},
			{Name: "level_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The title of the level.", Hydrate: activeGrantHydrateLevelTitle},
			{Name: "level_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the level.", Hydrate: activeGrantHydrateLevelTrunkTitle},
			{Name: "level_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The URI of the level.", Hydrate: activeGrantHydrateLevelUri},
			{Name: "resource_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the grant type.", Hydrate: activeGrantHydrateResourceTypeTrunkTitle},
			{Name: "resource_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Full title (including ancestor trunk) of the resource.", Hydrate: activeGrantHydrateResourceTrunkTitle},
			{Name: "resource_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the resource type.", Hydrate: activeGrantHydrateResourceTypeUri},
			{Name: "identity_akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "AKA (also known as) identifiers for the identity", Hydrate: activeGrantHydrateIdentityAkas},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue().NullIfEqual(""), Description: "The create time of the grant.", Hydrate: activeGrantHydrateCreateTimestamp},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used for this grant list."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue().NullIfEqual(""), Description: "Timestamp when the grant was last modified (created, updated or deleted).", Hydrate: activeGrantHydrateTimestamp},
			{Name: "update_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the grant was last updated in Turbot.", Hydrate: activeGrantHydrateUpdateTimestamp},
			{Name: "version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().NullIfEqual(""), Description: "Unique identifier for this version of the identity.", Hydrate: activeGrantHydrateVersionId},
			{Name: "workspace", Type: proto.ColumnType_STRING, Hydrate: getTurbotGuardrailsWorkspace, Transform: transform.FromValue(), Description: "Specifies the workspace URL."},
		},
	}
}

const (
	activeGrants = `
  query MyQuery($filter: [String!], $paging: String, $includeActiveGrantResourceTrunkTitle: Boolean!, $includeActiveGrantResourceAkas: Boolean! $includeActiveGrantResourceTitle: Boolean!, $includeActiveGrantResourceCreateTimestamp: Boolean!, $includeActiveGrantResourceUpdateTimestamp: Boolean!, $includeActiveGrantResourceDeleteTimestamp: Boolean!, $includeActiveGrantResourceTimestamp: Boolean!, $includeActiveGrantResourceVersionId: Boolean!, $includeActiveGrantResourceTypeURI: Boolean!, $includeActiveGrantResourceTypeTrunkTitle: Boolean!, $includeActiveGrantResourceId: Boolean!, $includeActiveGrantIdentityAkas: Boolean!, $includeActiveGrantIdentityEmail: Boolean!, $includeActiveGrantIdentityStatus: Boolean!, $includeActiveGrantIdentityGivenName: Boolean!, $includeActiveGrantIdentityProfileId: Boolean!, $includeActiveGrantIdentityFamilyName: Boolean!, $includeActiveGrantIdentityDisplayName: Boolean!, $includeActiveGrantIdentityLastLoginTimestamp: Boolean!, $includeActiveGrantIdentityTrunkTitle: Boolean!, $includeActiveGrantLevelTitle: Boolean!, $includeActiveGrantDeleteTimestamp: Boolean!, $includeActiveGrantTitle: Boolean!, $includeActiveGrantLevelURI: Boolean!, $includeActiveGrantLevelTrunkTitle: Boolean!, $includeActiveGrantId: Boolean!, $includeActiveGrantCreateTimestamp: Boolean!, $includeActiveGrantUpdateTimestamp: Boolean!, $includeActiveGrantTimestamp: Boolean!, $includeActiveGrantVersionId: Boolean!) {
  activeGrants(filter: $filter, paging: $paging) {
    items {
      resource {
        akas @include(if: $includeActiveGrantResourceAkas)
        title @include(if: $includeActiveGrantResourceTitle)
        trunk {
          title @include(if: $includeActiveGrantResourceTrunkTitle)
        }
        type {
          uri @include(if: $includeActiveGrantResourceTypeURI)
          trunk {
            title @include(if: $includeActiveGrantResourceTypeTrunkTitle)
          }
        }
        turbot {
          id @include(if: $includeActiveGrantResourceId)
          createTimestamp @include(if: $includeActiveGrantResourceCreateTimestamp)
          deleteTimestamp @include(if: $includeActiveGrantResourceDeleteTimestamp)
          timestamp @include(if: $includeActiveGrantResourceTimestamp)
          versionId @include(if: $includeActiveGrantResourceVersionId)
          updateTimestamp @include(if: $includeActiveGrantResourceUpdateTimestamp)
        }
      }
      grant {
        identity {
          akas @include(if: $includeActiveGrantIdentityAkas)
          email: get(path: "email") @include(if: $includeActiveGrantIdentityEmail)
          status: get(path: "status") @include(if: $includeActiveGrantIdentityStatus)
          givenName: get(path: "givenName") @include(if: $includeActiveGrantIdentityGivenName)
          profileId: get(path: "profileId") @include(if: $includeActiveGrantIdentityProfileId)
          familyName: get(path: "familyName") @include(if: $includeActiveGrantIdentityFamilyName)
          displayName: get(path: "displayName") @include(if: $includeActiveGrantIdentityDisplayName)
          lastLoginTimestamp: get(path: "lastLoginTimestamp") @include(if: $includeActiveGrantIdentityLastLoginTimestamp)
          trunk {
            title @include(if: $includeActiveGrantIdentityTrunkTitle)
          }
        }
        level {
          title @include(if: $includeActiveGrantLevelTitle)
          uri @include(if: $includeActiveGrantLevelURI)
          trunk {
            title @include(if: $includeActiveGrantLevelTrunkTitle)
          }
        }
        turbot {
          id @include(if: $includeActiveGrantId)
        }
      }
      turbot {
        createTimestamp @include(if: $includeActiveGrantCreateTimestamp)
        deleteTimestamp @include(if: $includeActiveGrantDeleteTimestamp)
        updateTimestamp @include(if: $includeActiveGrantUpdateTimestamp)
        title @include(if: $includeActiveGrantTitle)
        timestamp @include(if: $includeActiveGrantTimestamp)
        versionId @include(if: $includeActiveGrantVersionId)
      }
    }
    paging {
      next
    }
  }
  }
 `
)

func listActiveGrants(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_active_grants.listActiveGrants", "connection_error", err)
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
	if quals["grant_id"] != nil {
		filters = append(filters, fmt.Sprintf("id:%s", getQualListValues(ctx, quals, "grant_id", "int64")))
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

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	appendActiveGrantColumnIncludes(&variables, d.QueryContext.Columns)

	for {
		result := &ActiveGrantInfo{}
		err = conn.DoRequest(activeGrants, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_active_grants.listActiveGrants", "query_error", err)
		}
		for _, ActiveGrantDetails := range result.ActiveGrants.Items {

			d.StreamListItem(ctx, ActiveGrantDetails)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !pageResults || result.ActiveGrants.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.ActiveGrants.Paging.Next
	}

	return nil, err
}
