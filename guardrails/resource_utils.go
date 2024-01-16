package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendResourceTypeColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeResourceTypeId"] = slices.Contains(cols, "id")
    (*m)["includeResourceTypeUri"] = slices.Contains(cols, "uri")
    (*m)["includeResourceTypeTitle"] = slices.Contains(cols, "title")
    (*m)["includeResourceTypeTrunkTitle"] = slices.Contains(cols, "trunk_title")
    (*m)["includeResourceTypeDescription"] = slices.Contains(cols, "description")
    (*m)["includeResourceTypeAkas"] = slices.Contains(cols, "akas")
    (*m)["includeResourceTypeCategoryId"] = slices.Contains(cols, "category_id")
    (*m)["includeResourceTypeCategoryUri"] = slices.Contains(cols, "category_uri")
    (*m)["includeResourceTypeCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeResourceTypeIcon"] = slices.Contains(cols, "icon")
    (*m)["includeResourceTypeModUri"] = slices.Contains(cols, "mod_uri")
    (*m)["includeResourceTypeParentId"] = slices.Contains(cols, "parent_id")
    (*m)["includeResourceTypePath"] = slices.Contains(cols, "path")
    (*m)["includeResourceTypeUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeResourceTypeVersionId"] = slices.Contains(cols, "version_id")
}

func extractResourceTypeFromHydrateItem(h *plugin.HydrateData) (ResourceType, error) {
    if resourceType, ok := h.Item.(ResourceType); ok {
        return resourceType, nil
    } else {
        return ResourceType{}, fmt.Errorf("unable to parse hydrate item %v as a ResourceType", h.Item)
    }
}

func resourceTypeHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Turbot.ID, nil
}

func resourceTypeHydrateUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.URI, nil
}

func resourceTypeHydrateTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Title, nil
}

func resourceTypeHydrateTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Trunk.Title, nil
}

func resourceTypeHydrateDescription(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Description, nil
}

func resourceTypeHydrateAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Turbot.Akas, nil
}

func resourceTypeHydrateCategoryId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Category.Turbot.ID, nil
}

func resourceTypeHydrateCategoryUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.CategoryURI, nil
}

func resourceTypeHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Turbot.CreateTimestamp, nil
}

func resourceTypeHydrateIcon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Icon, nil
}

func resourceTypeHydrateModUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.ModURI, nil
}

func resourceTypeHydrateParentId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Turbot.ParentID, nil
}

func resourceTypeHydratePath(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Turbot.Path, nil
}

func resourceTypeHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Turbot.UpdateTimestamp, nil
}

func resourceTypeHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resourceType, err := extractResourceTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resourceType.Turbot.VersionID, nil
}

func appendResourceColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeResourceId"] = slices.Contains(cols, "id")
    (*m)["includeResourceTitle"] = slices.Contains(cols, "title")
    (*m)["includeResourceTrunkTitle"] = slices.Contains(cols, "trunk_title")
    (*m)["includeResourceTags"] = slices.Contains(cols, "tags")
    (*m)["includeResourceAkas"] = slices.Contains(cols, "akas")
    (*m)["includeResourceCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeResourceData"] = slices.Contains(cols, "data")
    (*m)["includeResourceObject"] = slices.Contains(cols, "object")
    (*m)["includeResourceFilter"] = slices.Contains(cols, "filter")
    (*m)["includeResourceMetadata"] = slices.Contains(cols, "metadata")
    (*m)["includeResourceParentId"] = slices.Contains(cols, "parent_id")
    (*m)["includeResourcePath"] = slices.Contains(cols, "path")
    (*m)["includeResourceTypeId"] = slices.Contains(cols, "resource_type_id")
    (*m)["includeResourceTypeUri"] = slices.Contains(cols, "resource_type_uri")
    (*m)["includeResourceTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includeResourceUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeResourceVersionId"] = slices.Contains(cols, "version_id")
}

func extractResourceFromHydrateItem(h *plugin.HydrateData) (Resource, error) {
    if resource, ok := h.Item.(Resource); ok {
        return resource, nil
    } else {
        return Resource{}, fmt.Errorf("unable to parse hydrate item %v as a Resource", h.Item)
    }
}

func resourceHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.ID, nil
}

func resourceHydrateTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Title, nil
}

func resourceHydrateTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Trunk.Title, nil
}

func resourceHydrateTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Tags, nil
}

func resourceHydrateAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Akas, nil
}

func resourceHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.CreateTimestamp, nil
}

func resourceHydrateData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Data, nil
}

func resourceHydrateObject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Object, nil
}

func resourceHydrateMetadata(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Metadata, nil
}

func resourceHydrateParentId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.ParentID, nil
}

func resourceHydratePath(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Path, nil
}

func resourceHydrateResourceTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.ResourceTypeID, nil
}

func resourceHydrateResourceTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Type.URI, nil
}

func resourceHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.Timestamp, nil
}

func resourceHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.UpdateTimestamp, nil
}

func resourceHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    resource, err := extractResourceFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return resource.Turbot.VersionID, nil
}
