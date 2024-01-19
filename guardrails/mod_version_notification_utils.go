package turbot

import (
    "context"
    "fmt"
    "slices"

    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func appendNotificationColumnIncludes(m *map[string]interface{}, cols []string) {
    (*m)["includeNotificationTurbotId"] = slices.Contains(cols, "id")
    (*m)["includeNotificationTurbotResourceNewVersionId"] = slices.Contains(cols, "resource_new_version_id")
    (*m)["includeNotificationTurbotResourceOldVersionId"] = slices.Contains(cols, "resource_old_version_id")
    (*m)["includeNotificationResourceTypeTurbotId"] = slices.Contains(cols, "resource_type_id")
    (*m)["includeNotificationResourceTypeUri"] = slices.Contains(cols, "resource_type_uri")
    (*m)["includeNotificationResourceTypeTrunkTitle"] = slices.Contains(cols, "resource_type_trunk_title")
    (*m)["includeNotificationTurbotProcessId"] = slices.Contains(cols, "process_id")
    (*m)["includeNotificationTurbotCreateTimestamp"] = slices.Contains(cols, "create_timestamp")
    (*m)["includeNotificationTurbotResourceId"] = slices.Contains(cols, "resource_id")
    (*m)["includeNotificationIcon"] = slices.Contains(cols, "icon")
    (*m)["includeNotificationMessage"] = slices.Contains(cols, "message")
    (*m)["includeNotificationType"] = slices.Contains(cols, "notification_type")
    (*m)["includeNotificationActorIdentityTrunkTitle"] = slices.Contains(cols, "actor_identity_trunk_title")
    (*m)["includeNotificationActorIdentityId"] = slices.Contains(cols, "actor_identity_id")
    (*m)["includeNotificationResourceData"] = slices.Contains(cols, "resource_data")
    (*m)["includeNotificationResourceObject"] = slices.Contains(cols, "resource_object")
    (*m)["includeNotificationResourceTrunkTitle"] = slices.Contains(cols, "resource_trunk_title")
    (*m)["includeNotificationResourceTurbotTitle"] = slices.Contains(cols, "resource_title")
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
    (*m)["includeNotificationTurbotControlTypeId"] = slices.Contains(cols, "control_type_id")
    (*m)["includeNotificationTurbotPolicySettingId"] = slices.Contains(cols, "policy_setting_id")
    (*m)["includeNotificationTurbotPolicySettingNewVersionId"] = slices.Contains(cols, "policy_setting_new_version_id")
    (*m)["includeNotificationTurbotPolicySettingOldVersionId"] = slices.Contains(cols, "policy_setting_old_version_id")
    (*m)["includeNotificationTurbotPolicySettingDefaultTemplate"] = slices.Contains(cols, "policy_setting_default_template")
    (*m)["includeNotificationTurbotPolicySettingDefaultTemplateInput"] = slices.Contains(cols, "policy_setting_default_template_input")
    (*m)["includeNotificationTurbotPolicySettingTypeId"] = slices.Contains(cols, "policy_setting_type_id")
}

func extractNotificationFromHydrateItem(h *plugin.HydrateData) (Notification, error) {
    if notification, ok := h.Item.(Notification); ok {
        return notification, nil
    } else {
        return Notification{}, fmt.Errorf("unable to parse hydrate item %v as a Notification", h.Item)
    }
}

func notificationHydrateId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ID, nil
}

func notificationHydrateResourceNewVersionID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ResourceNewVersionID, nil
}

func notificationHydrateResourceOldVersionID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ResourceOldVersionID, nil
}

func notificationHydrateProcessId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ProcessID, nil
}

func notificationHydrateCreateTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.CreateTimestamp, nil
}

func notificationHydrateResourceID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ResourceID, nil
}

func notificationHydrateResourceTypeURI(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Type.URI, nil
}

func notificationHydrateResourceTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Type.Trunk.Title, nil
}

func notificationHydrateResourceTypeID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Type.Turbot.ID, nil
}

func notificationHydrateIcon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Icon, nil
}

func notificationHydrateMessage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Message, nil
}

func notificationHydrateNotificationType(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.NotificationType, nil
}

func notificationHydrateActorIdentityTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Actor.Identity.Trunk.Title, nil
}

func notificationHydrateActorIdentityId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Actor.Identity.Turbot.ID, nil
}

func notificationHydrateResourceData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Data, nil
}

func notificationHydrateResourceTurbotTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Turbot.Title, nil
}

func notificationHydrateResourceObject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Object, nil
}

func notificationHydrateResourceTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Trunk.Title, nil
}

func notificationHydrateResourceTurbotAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Turbot.Akas, nil
}

func notificationHydrateResourceParentId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Turbot.ParentID, nil
}

func notificationHydrateResourcePath(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Turbot.Path, nil
}

func notificationHydrateResourceTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Resource.Turbot.Tags, nil
}

func notificationHydratePolicySettingIsCalculated(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.isCalculated, nil
}

func notificationHydratePolicySettingTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Type.URI, nil
}

func notificationHydratePolicySettingTypeReadOnly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Type.ReadOnly, nil
}

func notificationHydratePolicySettingTypeSecret(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Type.Secret, nil
}

func notificationHydratePolicySettingTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Type.Trunk.Title, nil
}

func notificationHydratePolicySettingValue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Value, nil
}

func notificationHydrateControlState(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Control.State, nil
}

func notificationHydrateControlReason(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Control.Reason, nil
}

func notificationHydrateControlDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Control.Details, nil
}

func notificationHydrateControlTypeUri(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Control.Type.URI, nil
}

func notificationHydrateControlTypeTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Control.Type.Trunk.Title, nil
}

func notificationHydrateActiveGrantId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ActiveGrantsID, nil
}

func notificationHydrateActiveGrantNewVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ActiveGrantsNewVersionID, nil
}

func notificationHydrateActiveGrantOldVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ActiveGrantsOldVersionID, nil
}

func notificationHydrateActiveGrantValidToTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.ValidToTimestamp, nil
}

func notificationHydrateActiveGrantIdentityProfileId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.Identity.ProfileID, nil
}

func notificationHydrateActiveGrantIdentityTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.Identity.Trunk.Title, nil
}

func notificationHydrateActiveGrantLevelTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.Level.Title, nil
}

func notificationHydrateActiveGrantPermissionLevelId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.PermissionLevelId, nil
}

func notificationHydrateActiveGrantPermissionTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.PermissionTypeID, nil
}

func notificationHydrateActiveGrantRoleName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.RoleName, nil
}

func notificationHydrateActiveGrantTypeTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.ActiveGrant.Grant.Type.Title, nil
}

func notificationHydrateTurbotGrantId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.GrantID, nil
}

func notificationHydrateTurbotGrantNewVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.GrantNewVersionID, nil
}

func notificationHydrateTurbotGrantOldVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.GrantOldVersionID, nil
}

func notificationHydrateTurbotGrantValidToTimestamp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.ValidToTimestamp, nil
}

func notificationHydrateTurbotGrantIdentityProfileId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.Identity.ProfileID, nil
}

func notificationHydrateTurbotGrantIdentityTrunkTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.Identity.Trunk.Title, nil
}

func notificationHydrateTurbotGrantLevelTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.Level.Title, nil
}

func notificationHydrateTurbotGrantPermissionLevelId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.PermissionLevelId, nil
}

func notificationHydrateTurbotGrantPermissionTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.PermissionTypeID, nil
}

func notificationHydrateTurbotGrantRoleName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.RoleName, nil
}

func notificationHydrateTurbotGrantTypeTitle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Grant.Type.Title, nil
}

func notificationHydrateTurbotControlId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ControlID, nil
}

func notificationHydrateTurbotControlNewVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ControlNewVersionID, nil
}

func notificationHydrateTurbotControlOldVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.ControlOldVersionID, nil
}

func notificationHydrateTurbotControlTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Control.Type.Turbot.ID, nil
}

func notificationHydrateTurbotPolicySettingId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.PolicySettingID, nil
}

func notificationHydrateTurbotPolicySettingNewVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.PolicySettingNewVersionID, nil
}

func notificationHydrateTurbotPolicySettingOldVersionId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.Turbot.PolicySettingOldVersionID, nil
}

func notificationHydrateTurbotPolicySettingDefaultTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Type.DefaultTemplate, nil
}

func notificationHydrateTurbotPolicySettingDefaultTemplateInput(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Type.DefaultTemplateInput, nil
}

func notificationHydrateTurbotPolicySettingTypeId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    notification, err := extractNotificationFromHydrateItem(h)
    if err != nil {
        return nil, err
    }
    return notification.PolicySetting.Type.Turbot.ID, nil
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
