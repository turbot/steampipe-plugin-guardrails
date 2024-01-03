package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendModVersionColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeModIdentityName"] = slices.Contains(cols, "identity_name")
    (*m)["includeModName"] = slices.Contains(cols, "name")
    (*m)["includeModVersion"] = slices.Contains(cols, "version")
    (*m)["includeModStatus"] = slices.Contains(cols, "status")
    (*m)["includeModHead"] = slices.Contains(cols, "mod_peer_dependency")
}

func extractModVersionFromHydrateItem(h *plugin.HydrateData) (ModVersionInfo, error) {
    if modVersionInfo, ok := h.Item.(ModVersionInfo); ok {
        return modVersionInfo, nil
    } else {
        return ModVersionInfo{}, fmt.Errorf("unable to parse hydrate item %v as a Mod Version", h.Item)
    }
}

func modVersionHydrateIdentityName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    modVersionInfo, err := extractModVersionFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return modVersionInfo.IdentityName, nil
}

func modVersionHydrateName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    modVersionInfo, err := extractModVersionFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return modVersionInfo.Name, nil
}

func modVersionHydrateVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    modVersionInfo, err := extractModVersionFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return modVersionInfo.Version, nil
}

func modVersionHydrateStatus(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    modVersionInfo, err := extractModVersionFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return modVersionInfo.Status, nil
}

func modVersionHydratePeerDependencies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    modVersionInfo, err := extractModVersionFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return modVersionInfo.Head.PeerDependencies, nil
}
