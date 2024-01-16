package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendNotificationColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeNotificationIcon"] = slices.Contains(cols, "icon")
    (*m)["includeNotificationMessage"] = slices.Contains(cols, "message")
    (*m)["includeNotificationType"] = slices.Contains(cols, "notification_type")
    (*m)["includeNotificationData"] = slices.Contains(cols, "data")
    (*m)["includeNotificationActorIdentityTrunkTitle"] = slices.Contains(cols, "actor_identity_trunk_title")
    (*m)["includeNotificationActorIdentityId"] = slices.Contains(cols, "actor_identity_id")
    (*m)["includeNotificationResourceData"] = slices.Contains(cols, "resource_data")
    (*m)["includeNotificationResourceObject"] = slices.Contains(cols, "resource_object")
    (*m)["includeNotificationResourceMetadata"] = slices.Contains(cols, "resource_metadata")
    (*m)["includeNotificationResourceTrunkTitle"] = slices.Contains(cols, "resource_trunk_title")
    (*m)["includeNotificationResourceTurbotAkas"] = slices.Contains(cols, "resource_akas")
    (*m)["includeNotificationResourceParentId"] = slices.Contains(cols, "resource_parent_id")
    (*m)["includeNotificationResourcePath"] = slices.Contains(cols, "resource_path")
    (*m)["includeNotificationResourceTags"] = slices.Contains(cols, "resource_tags")
    (*m)["includeNotificationPolicySettingIsCalculated"] = slices.Contains(cols, "policy_setting_is_calculated")
    (*m)["includeNotificationPolicySettingTypeUri"] = slices.Contains(cols, "policy_setting_type_uri")
    (*m)["includeNotificationPolicySettingTypeReadOnly"] = slices.Contains(cols, "policy_setting_type_read_only")
    (*m)["includeNotificationPolicySettingTypeSecret"] = slices.Contains(cols, "policy_setting_type_secret")
    (*m)["includeNotificationPolicySettingTypeTrunkTitle"] = slices.Contains(cols, "policy_setting_type_trunk_title")
    (*m)["includeNotificationPolicySettingValue"] = slices.Contains(cols, "policy_setting_value")
    (*m)["includeNotificationControlState"] = slices.Contains(cols, "control_state")
    (*m)["includeNotificationControlReason"] = slices.Contains(cols, "control_reason")
    (*m)["includeNotificationControlDetails"] = slices.Contains(cols, "control_details")
    (*m)["includeNotificationControlTypeUri"] = slices.Contains(cols, "control_type_uri")
    (*m)["includeNotificationControlTypeTrunkTitle"] = slices.Contains(cols, "control_type_trunk_title")
    (*m)["includeNotificationTurbotActiveGrantsId"] = slices.Contains(cols, "active_grant_id")
    (*m)["includeNotificationTurbotActiveGrantsNewVersionId"] = slices.Contains(cols, "active_grant_new_version_id")
    (*m)["includeNotificationTurbotActiveGrantsOldVersionId"] = slices.Contains(cols, "active_grant_old_version_id")
    (*m)["includeNotificationTurbotActiveGrantsValidToTimestamp"] = slices.Contains(cols, "active_grant_valid_to_timestamp")
    (*m)["includeNotificationTurbotActiveGrantsIdentityProfileId"] = slices.Contains(cols, "active_grant_identity_profile_id")
    (*m)["includeNotificationTurbotActiveGrantsIdentityTrunkTitle"] = slices.Contains(cols, "active_grant_identity_trunk_title")
    (*m)["includeNotificationTurbotActiveGrantsLevelTitle"] = slices.Contains(cols, "active_grant_level_title")
    (*m)["includeNotificationTurbotActiveGrantsPermissionLevelId"] = slices.Contains(cols, "active_grant_permission_level_id")
    (*m)["includeNotificationTurbotActiveGrantsPermissionTypeId"] = slices.Contains(cols, "active_grant_permission_type_id")
    (*m)["includeNotificationTurbotActiveGrantsRoleName"] = slices.Contains(cols, "active_grant_role_name")
    (*m)["includeNotificationTurbotActiveGrantsTypeTitle"] = slices.Contains(cols, "active_grant_type_title")
    (*m)["includeNotificationTurbotGrantId"] = slices.Contains(cols, "grant_id")
    (*m)["includeNotificationTurbotGrantNewVersionId"] = slices.Contains(cols, "grant_new_version_id")
    (*m)["includeNotificationTurbotGrantOldVersionId"] = slices.Contains(cols, "grant_old_version_id")
    (*m)["includeNotificationTurbotGrantValidToTimestamp"] = slices.Contains(cols, "grant_valid_to_timestamp")
    (*m)["includeNotificationTurbotGrantIdentityProfileId"] = slices.Contains(cols, "grant_identity_profile_id")
    (*m)["includeNotificationTurbotGrantIdentityTrunkTitle"] = slices.Contains(cols, "grant_identity_trunk_title")
    (*m)["includeNotificationTurbotGrantLevelTitle"] = slices.Contains(cols, "grant_level_title")
    (*m)["includeNotificationTurbotGrantPermissionLevelId"] = slices.Contains(cols, "grant_permission_level_id")
    (*m)["includeNotificationTurbotGrantPermissionTypeId"] = slices.Contains(cols, "grant_permission_type_id")
    (*m)["includeNotificationTurbotGrantRoleName"] = slices.Contains(cols, "grant_role_name")
    (*m)["includeNotificationTurbotGrantTypeTitle"] = slices.Contains(cols, "grant_type_title")
    (*m)["includeNotificationTurbotControlId"] = slices.Contains(cols, "control_id")
    (*m)["includeNotificationTurbotControlNewVersionId"] = slices.Contains(cols, "control_new_version_id")
    (*m)["includeNotificationTurbotControlOldVersionId"] = slices.Contains(cols, "control_old_version_id")
    (*m)["includeNotificationTurbotControlDetails"] = slices.Contains(cols, "control_details")
    (*m)["includeNotificationTurbotControlReason"] = slices.Contains(cols, "control_reason")
    (*m)["includeNotificationTurbotControlState"] = slices.Contains(cols, "control_state")
    (*m)["includeNotificationTurbotControlTypeId"] = slices.Contains(cols, "control_type_id")
    (*m)["includeNotificationTurbotControlTypeTrunkTitle"] = slices.Contains(cols, "control_type_trunk_title")
    (*m)["includeNotificationTurbotControlTypeUri"] = slices.Contains(cols, "control_type_uri")
    (*m)["includeNotificationTurbotPolicySettingId"] = slices.Contains(cols, "policy_setting_id")
    (*m)["includeNotificationTurbotPolicySettingNewVersionId"] = slices.Contains(cols, "policy_setting_new_version_id")
    (*m)["includeNotificationTurbotPolicySettingOldVersionId"] = slices.Contains(cols, "policy_setting_old_version_id")
    (*m)["includeNotificationTurbotPolicySettingDefaultTemplate"] = slices.Contains(cols, "policy_setting_default_template")
    (*m)["includeNotificationTurbotPolicySettingDefaultTemplateInput"] = slices.Contains(cols, "policy_setting_default_template_input")
    (*m)["includeNotificationTurbotPolicySettingIsCalculated"] = slices.Contains(cols, "policy_setting_is_calculated")
    (*m)["includeNotificationTurbotPolicySettingTypeId"] = slices.Contains(cols, "policy_setting_type_id")
    (*m)["includeNotificationTurbotPolicySettingTypeReadOnly"] = slices.Contains(cols, "policy_setting_type_read_only")
    (*m)["includeNotificationTurbotPolicySettingTypeSecret"] = slices.Contains(cols, "policy_setting_type_secret")
    (*m)["includeNotificationTurbotPolicySettingTypeTrunkTitle"] = slices.Contains(cols, "policy_setting_type_trunk_title")
    (*m)["includeNotificationTurbotPolicySettingTypeUri"] = slices.Contains(cols, "policy_setting_type_uri")
    (*m)["includeNotificationTurbotPolicySettingValue"] = slices.Contains(cols, "policy_setting_value")
}

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
