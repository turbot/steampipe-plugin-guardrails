package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func controlColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeControlId"] = slices.Contains(cols, "id")
    (*m)["includeControlState"] = slices.Contains(cols, "state")
    (*m)["includeControlReason"] = slices.Contains(cols, "reason")
    (*m)["includeControlDetails"] = slices.Contains(cols, "details")
    (*m)["includeControlResourceId"] = slices.Contains(cols, "resource_id")
    (*m)["includeControlTypeTrunkTitle"] = slices.Contains(cols, "control_type_trunk_title")
    (*m)["includeControlResourceTrunkTitle"] = slices.Contains(cols, "resource_trunk_title")
    (*m)["includeControlTypeUri"] = slices.Contains(cols, "control_type_uri")
    (*m)["includeControlTypeId"] = slices.Contains(cols, "control_type_id")
    (*m)["includeControlTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includeControlCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeControlUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeControlVersionId"] = slices.Contains(cols, "version_id")
    (*m)["includeControlResourceTypeId"] = slices.Contains(cols, "resource_type_id")
    (*m)["includeControlResourceTypeUri"] = slices.Contains(cols, "resource_type_uri")
    (*m)["includeControlMetadata"] = slices.Contains(cols, "metadata")

}

func extractControlFromHydrateItem(h *plugin.HydrateData) (ControlItem, error) {
    if controlItem, ok := h.Item.(ControlItem); ok {
        return controlItem, nil
    } else {
        return ControlItem{}, fmt.Errorf("unable to parse hydrate item %v as a Control", h.Item)
    }
}

func controlHydrateMetadata(_ context.Context, _ *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Metadata, nil
}

func controlHydrateId(_ context.Context, _ *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.ID, nil
}

func controlHydrateState(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.State, nil
}

func controlHydrateReason(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Reason, nil
}

func controlHydrateDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Details, nil
}

func controlHydrateResourceId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.ResourceID, nil
}

func controlHydrateResourceTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Resource.Trunk.Title, nil
}

func controlHydrateControlTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Type.Trunk.Title, nil
}

func controlHydrateControlTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.ControlTypeID, nil
}

func controlHydrateControlTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Type.URI, nil
}

func controlHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.Timestamp, nil
}

func controlHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.CreateTimestamp, nil
}

func controlHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.UpdateTimestamp, nil
}

func controlHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.VersionID, nil
}

func controlHydrateResourceTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Turbot.ResourceTypeID, nil
}

func controlHydrateResourceTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlItem, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlItem.Item.Resource.Type.URI, nil
}
