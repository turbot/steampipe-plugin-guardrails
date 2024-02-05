package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendControlTypeColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeControlTypeId"] = slices.Contains(cols, "id")
    (*m)["includeControlTypeUri"] = slices.Contains(cols, "uri")
    (*m)["includeControlTypeTitle"] = slices.Contains(cols, "title")
    (*m)["includeControlTypeTrunkTitle"] = slices.Contains(cols, "trunk_title")
    (*m)["includeControlTypeDescription"] = slices.Contains(cols, "description")
    (*m)["includeControlTypeTargets"] = slices.Contains(cols, "targets")
    (*m)["includeControlTypeTurbotAkas"] = slices.Contains(cols, "akas")
    (*m)["includeControlTypeCategoryId"] = slices.Contains(cols, "category_id")
    (*m)["includeControlTypeCategoryUri"] = slices.Contains(cols, "category_uri")
    (*m)["includeControlTypeTurbotCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeControlTypeIcon"] = slices.Contains(cols, "icon")
    (*m)["includeControlTypeModUri"] = slices.Contains(cols, "mod_uri")
    (*m)["includeControlTypeTurbotParentId"] = slices.Contains(cols, "parent_id")
    (*m)["includeControlTypeTurbotPath"] = slices.Contains(cols, "path")
    (*m)["includeControlTypeTurbotUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includeControlTypeTurbotVersionId"] = slices.Contains(cols, "version_id")
}

func extractControlTypeFromHydrateItem(h *plugin.HydrateData) (ControlType, error) {
    if controlType, ok := h.Item.(ControlType); ok {
        return controlType, nil
    } else {
        return ControlType{}, fmt.Errorf("unable to parse hydrate item %v as a Control Type", h.Item)
    }
}

func controlTypeHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Turbot.ID, nil
}

func controlTypeHydrateUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.URI, nil
}

func controlTypeHydrateTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Title, nil
}

func controlTypeHydrateTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Trunk.Title, nil
}

func controlTypeHydrateDescription(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Description, nil
}

func controlTypeHydrateTargets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Targets, nil
}

func controlTypeHydrateAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Turbot.Akas, nil
}

func controlTypeHydrateCategoryId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Category.Turbot.ID, nil
}

func controlTypeHydrateCategoryUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Category.URI, nil
}

func controlTypeHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Turbot.CreateTimestamp, nil
}

func controlTypeHydrateIcon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Icon, nil
}

func controlTypeHydrateModUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.ModURI, nil
}

func controlTypeHydrateParentId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Turbot.ParentID, nil
}

func controlTypeHydratePath(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Turbot.Path, nil
}

func controlTypeHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Turbot.UpdateTimestamp, nil
}

func controlTypeHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    controlType, err := extractControlTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return controlType.Turbot.VersionID, nil
}

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
}

func extractControlFromHydrateItem(h *plugin.HydrateData) (Control, error) {
    if control, ok := h.Item.(Control); ok {
        return control, nil
    } else {
        return Control{}, fmt.Errorf("unable to parse hydrate item %v as a Control", h.Item)
    }
}

func controlHydrateId(_ context.Context, _ *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.ID, nil
}

func controlHydrateState(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.State, nil
}

func controlHydrateReason(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Reason, nil
}

func controlHydrateDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Details, nil
}

func controlHydrateResourceId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.ResourceID, nil
}

func controlHydrateResourceTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Resource.Trunk.Title, nil
}

func controlHydrateControlTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Type.Trunk.Title, nil
}

func controlHydrateControlTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.ControlTypeID, nil
}

func controlHydrateControlTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Type.URI, nil
}

func controlHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.Timestamp, nil
}

func controlHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.CreateTimestamp, nil
}

func controlHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.UpdateTimestamp, nil
}

func controlHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.VersionID, nil
}

func controlHydrateResourceTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Turbot.ResourceTypeID, nil
}

func controlHydrateResourceTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    control, err := extractControlFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return control.Resource.Type.URI, nil
}
