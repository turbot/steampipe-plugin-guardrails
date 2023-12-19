package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendGrantColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeGrantTurbotId"] = slices.Contains(cols, "id")
    (*m)["includeGrantResourceId"] = slices.Contains(cols, "resource_id")
    (*m)["includeGrantIdentityStatus"] = slices.Contains(cols, "identity_status")
    (*m)["includeGrantIdentityDisplayName"] = slices.Contains(cols, "identity_display_name")
    (*m)["includeGrantIdentityEmail"] = slices.Contains(cols, "identity_email")
    (*m)["includeGrantIdentityFamilyName"] = slices.Contains(cols, "identity_family_name")
    (*m)["includeGrantIdentityGivenName"] = slices.Contains(cols, "identity_given_name")
    (*m)["includeGrantIdentityLastLoginTimestamp"] = slices.Contains(cols, "identity_last_login_timestamp")
    (*m)["includeGrantIdentityProfileId"] = slices.Contains(cols, "identity_profile_id")
    (*m)["includeGrantIdentityTrunkTitle"] = slices.Contains(cols, "identity_trunk_title")
    (*m)["includeGrantLevelTitle"] = slices.Contains(cols, "level_title")
    (*m)["includeGrantLevelTrunkTitle"] = slices.Contains(cols, "level_trunk_title")
    (*m)["includeGrantLevelURI"] = slices.Contains(cols, "level_uri")
    (*m)["includeGrantResourceTypeTrunkTitle"] = slices.Contains(cols, "resource_type_trunk_title")
    (*m)["includeGrantResourceTrunkTitle"] = slices.Contains(cols, "resource_trunk_title")
    (*m)["includeGrantResourceTypeURI"] = slices.Contains(cols, "resource_type_uri")
    (*m)["includeGrantIdentityAkas"] = slices.Contains(cols, "identity_akas")
    (*m)["includeGrantTurbotCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeGrantTurbotTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includeGrantTurbotUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeGrantTurbotVersionId"] = slices.Contains(cols, "version_id")

    // columns which are not part of the table
    (*m)["includeGrantResourceAkas"] = slices.Contains(cols, "resource_akas")
    (*m)["includeGrantResourceTitle"] = slices.Contains(cols, "resource_title")
    (*m)["includeGrantResourceCreateTimestamp"] = slices.Contains(cols, "resource_create_timestamp")
    (*m)["includeGrantResourceUpdateTimestamp"] = slices.Contains(cols, "resource_update_timestamp")
    (*m)["includeGrantResourceDeleteTimestamp"] = slices.Contains(cols, "resource_delete_timestamp")
    (*m)["includeGrantResourceTimestamp"] = slices.Contains(cols, "resource_timestamp")
    (*m)["includeGrantResourceVersionId"] = slices.Contains(cols, "resource_version_id")
    (*m)["includeGrantTurbotDeleteTimestamp"] = slices.Contains(cols, "turbot_delete_timestamp")
}

func extractGrantFromHydrateItem(h *plugin.HydrateData) (Grant, error) {
    if grant, ok := h.Item.(Grant); ok {
        return grant, nil
    } else {
        return Grant{}, fmt.Errorf("unable to parse hydrate item %v as a Grant", h.Item)
    }
}

func grantHydrateGrantId(_ context.Context, _ *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Turbot.ID, nil
}

func grantHydrateResourceId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Resource.Turbot.ID, nil
}

func grantHydrateIdentityStatus(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.Status, nil
}

func grantHydrateIdentityDisplayName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.DisplayName, nil
}

func grantHydrateIdentityEmail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.Email, nil
}

func grantHydrateIdentityFamilyName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.FamilyName, nil
}

func grantHydrateIdentityGivenName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.GivenName, nil
}

func grantHydrateIdentityLastLoginTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.LastLoginTimestamp, nil
}

func grantHydrateIdentityProfileId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.ProfileID, nil
}

func grantHydrateIdentityTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.Trunk.Title, nil
}

func grantHydrateLevelTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Level.Title, nil
}

func grantHydrateLevelTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Level.Trunk.Title, nil
}

func grantHydrateLevelUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Level.URI, nil
}

func grantHydrateResourceTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Resource.Type.Trunk.Title, nil
}

func grantHydrateResourceTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Resource.Trunk.Title, nil
}

func grantHydrateResourceTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Resource.Type.URI, nil
}

func grantHydrateIdentityAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Identity.Akas, nil
}

func grantHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Turbot.CreateTimestamp, nil
}

func grantHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Turbot.Timestamp, nil
}

func grantHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Turbot.UpdateTimestamp, nil
}

func grantHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    grant, err := extractGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return grant.Turbot.VersionID, nil
}

func appendActiveGrantColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeActiveGrantId"] = slices.Contains(cols, "grant_id")
    (*m)["includeActiveGrantResourceId"] = slices.Contains(cols, "resource_id")
    (*m)["includeActiveGrantIdentityStatus"] = slices.Contains(cols, "identity_status")
    (*m)["includeActiveGrantIdentityDisplayName"] = slices.Contains(cols, "identity_display_name")
    (*m)["includeActiveGrantIdentityEmail"] = slices.Contains(cols, "identity_email")
    (*m)["includeActiveGrantIdentityFamilyName"] = slices.Contains(cols, "identity_family_name")
    (*m)["includeActiveGrantIdentityGivenName"] = slices.Contains(cols, "identity_given_name")
    (*m)["includeActiveGrantIdentityLastLoginTimestamp"] = slices.Contains(cols, "identity_last_login_timestamp")
    (*m)["includeActiveGrantIdentityProfileId"] = slices.Contains(cols, "identity_profile_id")
    (*m)["includeActiveGrantIdentityTrunkTitle"] = slices.Contains(cols, "identity_trunk_title")
    (*m)["includeActiveGrantLevelTitle"] = slices.Contains(cols, "level_title")
    (*m)["includeActiveGrantLevelTrunkTitle"] = slices.Contains(cols, "level_trunk_title")
    (*m)["includeActiveGrantLevelURI"] = slices.Contains(cols, "level_uri")
    (*m)["includeActiveGrantResourceTypeTrunkTitle"] = slices.Contains(cols, "resource_type_trunk_title")
    (*m)["includeActiveGrantResourceTrunkTitle"] = slices.Contains(cols, "resource_trunk_title")
    (*m)["includeActiveGrantResourceTypeURI"] = slices.Contains(cols, "resource_type_uri")
    (*m)["includeActiveGrantIdentityAkas"] = slices.Contains(cols, "identity_akas")
    (*m)["includeActiveGrantCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeActiveGrantTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includeActiveGrantUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeActiveGrantVersionId"] = slices.Contains(cols, "version_id")

    // columns which are not part of the table
    (*m)["includeActiveGrantResourceAkas"] = slices.Contains(cols, "resource_akas")
    (*m)["includeActiveGrantResourceTitle"] = slices.Contains(cols, "resource_title")
    (*m)["includeActiveGrantResourceCreateTimestamp"] = slices.Contains(cols, "resource_create_timestamp")
    (*m)["includeActiveGrantResourceUpdateTimestamp"] = slices.Contains(cols, "resource_update_timestamp")
    (*m)["includeActiveGrantResourceDeleteTimestamp"] = slices.Contains(cols, "resource_delete_timestamp")
    (*m)["includeActiveGrantResourceTimestamp"] = slices.Contains(cols, "resource_timestamp")
    (*m)["includeActiveGrantResourceVersionId"] = slices.Contains(cols, "resource_version_id")
    (*m)["includeActiveGrantDeleteTimestamp"] = slices.Contains(cols, "turbot_delete_timestamp")
    (*m)["includeActiveGrantTitle"] = slices.Contains(cols, "turbot_title")
}

func extractActiveGrantFromHydrateItem(h *plugin.HydrateData) (ActiveGrant, error) {
    if activeGrant, ok := h.Item.(ActiveGrant); ok {
        return activeGrant, nil
    } else {
        return ActiveGrant{}, fmt.Errorf("unable to parse hydrate item %v as a Active Grant", h.Item)
    }
}

func activeGrantHydrateGrantId(_ context.Context, _ *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Turbot.ID, nil
}

func activeGrantHydrateResourceId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Resource.Turbot.ID, nil
}

func activeGrantHydrateIdentityStatus(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.Status, nil
}

func activeGrantHydrateIdentityDisplayName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.DisplayName, nil
}

func activeGrantHydrateIdentityEmail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.Email, nil
}

func activeGrantHydrateIdentityFamilyName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.FamilyName, nil
}

func activeGrantHydrateIdentityGivenName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.GivenName, nil
}

func activeGrantHydrateIdentityLastLoginTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.LastLoginTimestamp, nil
}

func activeGrantHydrateIdentityProfileId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.ProfileID, nil
}

func activeGrantHydrateIdentityTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.Trunk.Title, nil
}

func activeGrantHydrateLevelTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Level.Title, nil
}

func activeGrantHydrateLevelTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Level.Trunk.Title, nil
}

func activeGrantHydrateLevelUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Level.URI, nil
}

func activeGrantHydrateResourceTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Resource.Type.Trunk.Title, nil
}

func activeGrantHydrateResourceTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Resource.Trunk.Title, nil
}

func activeGrantHydrateResourceTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Resource.Type.URI, nil
}

func activeGrantHydrateIdentityAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Grant.Identity.Akas, nil
}

func activeGrantHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Turbot.CreateTimestamp, nil
}

func activeGrantHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Turbot.Timestamp, nil
}

func activeGrantHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Turbot.UpdateTimestamp, nil
}

func activeGrantHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    activeGrant, err := extractActiveGrantFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return activeGrant.Turbot.VersionID, nil
}
