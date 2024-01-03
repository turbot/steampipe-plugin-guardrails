package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendSmartFolderColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeAttachedResourcesId"] = slices.Contains(cols, "attached_resource_ids")
    (*m)["includeData"] = slices.Contains(cols, "description")
    (*m)["includeMetadata"] = slices.Contains(cols, "metadata")
    (*m)["includeTrunkTitle"] = slices.Contains(cols, "trunk_title")
    (*m)["includeTurbotId"] = slices.Contains(cols, "id")
    (*m)["includeTurbotTitle"] = slices.Contains(cols, "title")
    (*m)["includeTurbotTags"] = slices.Contains(cols, "tags")
    (*m)["includeTurbotAkas"] = slices.Contains(cols, "akas")
    (*m)["includeTurbotTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includeTurbotCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeTurbotUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeTurbotVersionId"] = slices.Contains(cols, "version_id")
    (*m)["includeTurbotParentId"] = slices.Contains(cols, "parent_id")
    (*m)["includeTurbotPath"] = slices.Contains(cols, "path")
    (*m)["includeTurbotResourceTypeId"] = slices.Contains(cols, "resource_type_id")
    (*m)["includeTypeUri"] = slices.Contains(cols, "resource_type_uri")
}

func extractSmartFolderFromHydrateItem(h *plugin.HydrateData) (Resource, error) {
    if resource, ok := h.Item.(Resource); ok {
        return resource, nil
    } else {
        return Resource{}, fmt.Errorf("unable to parse hydrate item %v as a Resource", h.Item)
    }
}

func smartFolderHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.ID, nil
}

func smartFolderHydrateTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Title, nil
}

func smartFolderHydrateTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Tags, nil
}

func smartFolderHydrateAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Akas, nil
}

func smartFolderHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Timestamp, nil
}

func smartFolderHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.CreateTimestamp, nil
}

func smartFolderHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.UpdateTimestamp, nil
}

func smartFolderHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.VersionID, nil
}

func smartFolderHydrateParentId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.ParentID, nil
}

func smartFolderHydratePath(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Path, nil
}

func smartFolderHydrateResourceTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.ResourceTypeID, nil
}

func smartFolderHydrateTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Type.URI, nil
}

func smartFolderHydrateData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Data, nil
}

func smartFolderHydrateMetadata(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Metadata, nil
}

func smartFolderHydrateTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Trunk.Title, nil
}

func smartFolderHydrateAttachedResources(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractSmartFolderFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.AttachedResources.Items, nil
}

func appendTagColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeTagKey"] = slices.Contains(cols, "key")
    (*m)["includeTagValue"] = slices.Contains(cols, "value")
    (*m)["includeTagTurbotId"] = slices.Contains(cols, "id")
    (*m)["includeTagTurbotTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includeTagTurbotCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeTagTurbotUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeTagTurbotVersionId"] = slices.Contains(cols, "version_id")
    (*m)["includeTagResources"] = slices.Contains(cols, "resource_ids")
}

func extractTagFromHydrateItem(h *plugin.HydrateData) (Tag, error) {
    if tag, ok := h.Item.(Tag); ok {
        return tag, nil
    } else {
        return Tag{}, fmt.Errorf("unable to parse hydrate item %v as a Tag", h.Item)
    }
}

func tagHydrateKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Key, nil
}

func tagHydrateValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Value, nil
}

func tagHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Turbot.ID, nil
}

func tagHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Turbot.Timestamp, nil
}

func tagHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Turbot.CreateTimestamp, nil
}

func tagHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Turbot.UpdateTimestamp, nil
}

func tagHydrateVersionID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Turbot.VersionID, nil
}

func tagHydrateResources(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    tag, err := extractTagFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return tag.Resources, nil
}
