package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendPolicyValueColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includePolicyValueTurbotId"] = slices.Contains(cols, "id")
    (*m)["includePolicyValueTypeTitle"] = slices.Contains(cols, "policy_type_title")
    (*m)["includePolicyValueTypeTrunkTitle"] = slices.Contains(cols, "policy_type_trunk_title")
    (*m)["includePolicyValueDefault"] = slices.Contains(cols, "is_default")
    (*m)["includePolicyValueIsCalculated"] = slices.Contains(cols, "is_calculated")
    (*m)["includePolicyValuePrecedence"] = slices.Contains(cols, "precedence")
    (*m)["includePolicyValueTurbotResourceId"] = slices.Contains(cols, "resource_id")
    (*m)["includePolicyValueResourceTrunkTitle"] = slices.Contains(cols, "resource_trunk_title")
    (*m)["includePolicyValueTurbotResourceTypeId"] = slices.Contains(cols, "resource_type_id")
    (*m)["includePolicyValueState"] = slices.Contains(cols, "state")
    (*m)["includePolicyValueSecretValue"] = slices.Contains(cols, "secret_value")
    (*m)["includePolicyValue"] = slices.Contains(cols, "value")
    (*m)["includePolicyValueTypeModUri"] = slices.Contains(cols, "type_mod_uri")
    (*m)["includePolicyValueTurbotPolicyTypeId"] = slices.Contains(cols, "policy_type_id")
    (*m)["includePolicyValueTypeDefaultTemplate"] = slices.Contains(cols, "policy_type_default_template")
    (*m)["includePolicyValueTurbotSettingId"] = slices.Contains(cols, "setting_id")
    (*m)["includePolicyValueDependentControls"] = slices.Contains(cols, "dependent_controls")
    (*m)["includePolicyValueDependentPolicyValues"] = slices.Contains(cols, "dependent_policy_values")
    (*m)["includePolicyValueTurbotCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includePolicyValueTurbotTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includePolicyValueTurbotUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includePolicyValueTurbotVersionId"] = slices.Contains(cols, "version_id")
}

func extractPolicyValueFromHydrateItem(h *plugin.HydrateData) (PolicyValue, error) {
    if policyValue, ok := h.Item.(PolicyValue); ok {
        return policyValue, nil
    } else {
        return PolicyValue{}, fmt.Errorf("unable to parse hydrate item %v as a PolicyValue", h.Item)
    }
}

func policyValueHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.ID, nil
}

func policyValueHydratePolicyTypeTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Type.Title, nil
}

func policyValueHydratePolicyTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Type.Trunk.Title, nil
}

func policyValueHydrateIsDefault(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Default, nil
}

func policyValueHydrateIsCalculated(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.IsCalculated, nil
}

func policyValueHydratePrecedence(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Precedence, nil
}

func policyValueHydrateResourceId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.ResourceId, nil
}

func policyValueHydrateResourceTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Resource.Trunk.Title, nil
}

func policyValueHydrateResourceTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.ResourceTypeID, nil
}

func policyValueHydrateState(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.State, nil
}

func policyValueHydrateSecretValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.SecretValue, nil
}

func policyValueHydrateValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Value, nil
}

func policyValueHydrateTypeModUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Type.ModURI, nil
}

func policyValueHydratePolicyTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.PolicyTypeId, nil
}

func policyValueHydratePolicyTypeDefaultTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Type.DefaultTemplate, nil
}

func policyValueHydrateSettingId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.SettingId, nil
}

func policyValueHydrateDependentControls(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.DependentControls, nil
}

func policyValueHydrateDependentPolicyValues(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.DependentPolicyValues, nil
}

func policyValueHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.CreateTimestamp, nil
}

func policyValueHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.Timestamp, nil
}

func policyValueHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.UpdateTimestamp, nil
}

func policyValueHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyValue, err := extractPolicyValueFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyValue.Turbot.VersionID, nil
}

func appendPolicyTypeColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includePolicyTypeTurbotId"] = slices.Contains(cols, "id")
    (*m)["includePolicyTypeUri"] = slices.Contains(cols, "uri")
    (*m)["includePolicyTypeTitle"] = slices.Contains(cols, "title")
    (*m)["includePolicyTypeTrunkTitle"] = slices.Contains(cols, "trunk_title")
    (*m)["includePolicyTypeDescription"] = slices.Contains(cols, "description")
    (*m)["includePolicyTypeTargets"] = slices.Contains(cols, "targets")
    (*m)["includePolicyTypeTurbotAkas"] = slices.Contains(cols, "akas")
    (*m)["includePolicyTypeCategoryTurbotId"] = slices.Contains(cols, "category_id")
    (*m)["includePolicyTypeCategoryUri"] = slices.Contains(cols, "category_uri")
    (*m)["includePolicyTypeTurbotCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includePolicyTypeDefaultTemplate"] = slices.Contains(cols, "default_template")
    (*m)["includePolicyTypeIcon"] = slices.Contains(cols, "icon")
    (*m)["includePolicyTypeModUri"] = slices.Contains(cols, "mod_uri")
    (*m)["includePolicyTypeTurbotParentId"] = slices.Contains(cols, "parent_id")
    (*m)["includePolicyTypeTurbotPath"] = slices.Contains(cols, "path")
    (*m)["includePolicyTypeReadOnly"] = slices.Contains(cols, "read_only")
    (*m)["includePolicyTypeSchema"] = slices.Contains(cols, "schema")
    (*m)["includePolicyTypeSecret"] = slices.Contains(cols, "secret")
    (*m)["includePolicyTypeSecretLevel"] = slices.Contains(cols, "secret_level")
    (*m)["includePolicyTypeTurbotUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includePolicyTypeTurbotVersionId"] = slices.Contains(cols, "version_id")
}

func extractPolicyTypeFromHydrateItem(h *plugin.HydrateData) (PolicyType, error) {
    if policyType, ok := h.Item.(PolicyType); ok {
        return policyType, nil
    } else {
        return PolicyType{}, fmt.Errorf("unable to parse hydrate item %v as a PolicyType", h.Item)
    }
}

func policyTypeHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Turbot.ID, nil
}

func policyTypeHydrateUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.URI, nil
}

func policyTypeHydrateTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Title, nil
}

func policyTypeHydrateTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Trunk.Title, nil
}

func policyTypeHydrateDescription(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Description, nil
}

func policyTypeHydrateTargets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Targets, nil
}

func policyTypeHydrateAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Turbot.Akas, nil
}

func policyTypeHydrateCategoryId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Category.Turbot.ID, nil
}

func policyTypeHydrateCategoryUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Category.URI, nil
}

func policyTypeHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Turbot.CreateTimestamp, nil
}

func policyTypeHydrateDefaultTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.DefaultTemplate, nil
}

func policyTypeHydrateIcon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Icon, nil
}

func policyTypeHydrateModUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.ModURI, nil
}

func policyTypeHydrateParentId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Turbot.ParentID, nil
}

func policyTypeHydratePath(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Turbot.Path, nil
}

func policyTypeHydrateReadOnly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.ReadOnly, nil
}

func policyTypeHydrateSchema(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Schema, nil
}

func policyTypeHydrateSecret(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Secret, nil
}

func policyTypeHydrateSecretLevel(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.SecretLevel, nil
}

func policyTypeHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Turbot.UpdateTimestamp, nil
}

func policyTypeHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policyType, err := extractPolicyTypeFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policyType.Turbot.VersionID, nil
}

func appendPolicySettingColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includePolicySettingId"] = slices.Contains(cols, "id")
    (*m)["includePolicySettingPrecedence"] = slices.Contains(cols, "precedence")
    (*m)["includePolicySettingResourceId"] = slices.Contains(cols, "resource_id")
    (*m)["includePolicySettingResourceTrunkTitle"] = slices.Contains(cols, "resource_trunk_title")
    (*m)["includePolicySettingPolicyTypeUri"] = slices.Contains(cols, "policy_type_uri")
    (*m)["includePolicySettingPolicyTypeTrunkTitle"] = slices.Contains(cols, "policy_type_trunk_title")
    (*m)["includePolicySettingValue"] = slices.Contains(cols, "value")
    (*m)["includePolicySettingIsCalculated"] = slices.Contains(cols, "is_calculated")
    (*m)["includePolicySettingException"] = slices.Contains(cols, "exception")
    (*m)["includePolicySettingOrphan"] = slices.Contains(cols, "orphan")
    (*m)["includePolicySettingNote"] = slices.Contains(cols, "note")
    (*m)["includePolicySettingCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includePolicySettingDefault"] = slices.Contains(cols, "default")
    (*m)["includePolicySettingInput"] = slices.Contains(cols, "input")
    (*m)["includePolicySettingPolicyTypeId"] = slices.Contains(cols, "policy_type_id")
    (*m)["includePolicySettingTemplate"] = slices.Contains(cols, "template")
    (*m)["includePolicySettingTemplateInput"] = slices.Contains(cols, "template_input")
    (*m)["includePolicySettingTimestamp"] = slices.Contains(cols, "timestamp")
    (*m)["includePolicySettingUpdateTimestamp"] = slices.Contains(cols, "update_timestamp")
    (*m)["includePolicySettingValidFromTimestamp"] = slices.Contains(cols, "valid_from_timestamp")
    (*m)["includePolicySettingValidToTimestamp"] = slices.Contains(cols, "valid_to_timestamp")
    (*m)["includePolicySettingValueSource"] = slices.Contains(cols, "value_source")
    (*m)["includePolicySettingVersionId"] = slices.Contains(cols, "version_id")
}

func extractPolicySettingFromHydrateItem(h *plugin.HydrateData) (PolicySetting, error) {
    if policySetting, ok := h.Item.(PolicySetting); ok {
        return policySetting, nil
    } else {
        return PolicySetting{}, fmt.Errorf("unable to parse hydrate item %v as a PolicySetting", h.Item)
    }
}

func policySettingHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Turbot.ID, nil
}

func policySettingHydratePrecedence(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Precedence, nil
}

func policySettingHydrateResourceId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Turbot.ResourceID, nil
}

func policySettingHydrateResourceTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Resource.Trunk.Title, nil
}

func policySettingHydratePolicyTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Type.URI, nil
}

func policySettingHydratePolicyTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Type.Trunk.Title, nil
}

func policySettingHydrateValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Value, nil
}

func policySettingHydrateIsCalculated(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.IsCalculated, nil
}

func policySettingHydrateException(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Exception, nil
}

func policySettingHydrateOrphan(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Orphan, nil
}

func policySettingHydrateNote(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Note, nil
}

func policySettingHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Turbot.CreateTimestamp, nil
}

func policySettingHydrateDefault(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Default, nil
}

func policySettingHydrateInput(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Input, nil
}

func policySettingHydratePolicyTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Turbot.PolicyTypeID, nil
}

func policySettingHydrateTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Template, nil
}

func policySettingHydrateTemplateInput(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.TemplateInput, nil
}

func policySettingHydrateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Turbot.Timestamp, nil
}

func policySettingHydrateUpdateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Turbot.UpdateTimestamp, nil
}

func policySettingHydrateValidFromTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.ValidFromTimestamp, nil
}

func policySettingHydrateValidToTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.ValidToTimestamp, nil
}

func policySettingHydrateValueSource(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.ValueSource, nil
}

func policySettingHydrateVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    policySetting, err := extractPolicySettingFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return policySetting.Turbot.VersionID, nil
}
