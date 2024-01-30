package turbot

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGuardrailsNotification(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "guardrails_notification",
		Description: "Notifications from the Turbot Guardrails CMDB.",
		List: &plugin.ListConfig{
			Hydrate: listNotification,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "id", Require: plugin.Optional},
				{Name: "notification_type", Require: plugin.Optional},
				{Name: "control_id", Require: plugin.Optional},
				{Name: "control_type_id", Require: plugin.Optional},
				{Name: "control_type_uri", Require: plugin.Optional},
				{Name: "resource_id", Require: plugin.Optional},
				{Name: "resource_type_id", Require: plugin.Optional},
				{Name: "resource_type_uri", Require: plugin.Optional},
				{Name: "policy_setting_type_id", Require: plugin.Optional},
				{Name: "policy_setting_type_uri", Require: plugin.Optional},
				{Name: "actor_identity_id", Require: plugin.Optional},
				{Name: "create_timestamp", Require: plugin.Optional, Operators: []string{">", ">=", "=", "<", "<="}},
				{Name: "filter", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getNotification,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Unique identifier of the notification.", Hydrate: notificationHydrateId},
			{Name: "process_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the process that created this notification.", Hydrate: notificationHydrateProcessId},
			{Name: "icon", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Icon for this notification type.", Hydrate: notificationHydrateIcon},
			{Name: "message", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Message for the notification.", Hydrate: notificationHydrateMessage},
			{Name: "notification_type", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Type of the notification: resource, action, policySetting, control, grant, activeGrant.", Hydrate: notificationHydrateNotificationType},
			{Name: "create_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "When the resource was first discovered by Turbot. (It may have been created earlier.)", Hydrate: notificationHydrateCreateTimestamp},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Filter used to search for notifications."},

			// Actor info for the notification
			{Name: "actor_identity_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue().NullIfZero(), Description: "Title hierarchy of the actor from the root down to the actor of this event.", Hydrate: notificationHydrateActorIdentityTrunkTitle},
			{Name: "actor_identity_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().NullIfZero(), Description: "Identity ID of the actor that performed this event.", Hydrate: notificationHydrateActorIdentityId},

			// Resource info for notification
			{Name: "resource_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().NullIfZero(), Description: "ID of the resource for this notification.", Hydrate: notificationHydrateResourceID},
			{Name: "resource_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the resource hierarchy from the root down to this resource.", Hydrate: notificationHydrateResourceTrunkTitle},
			{Name: "resource_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the resource.", Hydrate: notificationHydrateResourceTurbotTitle},
			{Name: "resource_new_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the resource after the event.", Hydrate: notificationHydrateResourceNewVersionID},
			{Name: "resource_old_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the resource before the event.", Hydrate: notificationHydrateResourceOldVersionID},
			{Name: "resource_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().NullIfZero(), Description: "ID of the resource type for this notification.", Hydrate: notificationHydrateResourceTypeID},
			{Name: "resource_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the resource type for this notification.", Hydrate: notificationHydrateResourceTypeURI},
			{Name: "resource_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Title of the resource type hierarchy from the root down to this resource.", Hydrate: notificationHydrateResourceTypeTrunkTitle},
			{Name: "resource_data", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "The data for this resource", Hydrate: notificationHydrateResourceData},
			{Name: "resource_object", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "More detailed and extensive resource data", Hydrate: notificationHydrateResourceObject},
			{Name: "resource_akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "The globally-unique akas for this resource.", Hydrate: notificationHydrateResourceTurbotAkas},
			{Name: "resource_parent_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().NullIfZero(), Description: "The id of the parent resource of this resource.", Hydrate: notificationHydrateResourceParentId},
			{Name: "resource_path", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The string of resource ids separated by \".\" from root down to this resource.", Hydrate: notificationHydrateResourcePath},
			{Name: "resource_tags", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "Tags attached to this resource.", Hydrate: notificationHydrateResourceTags},

			// Policy settings notification details
			{Name: "policy_setting_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the policy setting for this notification.", Hydrate: notificationHydrateTurbotPolicySettingId},
			{Name: "policy_setting_new_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the policy setting after the event.", Hydrate: notificationHydrateTurbotPolicySettingNewVersionId},
			{Name: "policy_setting_old_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the policy setting before the event.", Hydrate: notificationHydrateTurbotPolicySettingOldVersionId},
			{Name: "policy_setting_default_template", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The Nunjucks template if this setting is for a calculated value.", Hydrate: notificationHydrateTurbotPolicySettingDefaultTemplate},
			{Name: "policy_setting_default_template_input", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(formatPolicyFieldsValue), Description: "The GraphQL Input query if this setting is for a calculated value.", Hydrate: notificationHydrateTurbotPolicySettingDefaultTemplateInput},
			{Name: "policy_setting_is_calculated", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "If true this setting contains calculated inputs e.g. templateInput and template.", Hydrate: notificationHydratePolicySettingIsCalculated},
			{Name: "policy_setting_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue().NullIfZero(), Description: "ID of the policy setting type for this notification.", Hydrate: notificationHydrateTurbotPolicySettingTypeId},
			{Name: "policy_setting_type_read_only", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "If true user-defined policy settings are blocked from being created.", Hydrate: notificationHydratePolicySettingTypeReadOnly},
			{Name: "policy_setting_type_secret", Type: proto.ColumnType_BOOL, Transform: transform.FromValue(), Description: "If true policy value will be encrypted.", Hydrate: notificationHydratePolicySettingTypeSecret},
			{Name: "policy_setting_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "This is the title of hierarchy from the root down to this policy type.", Hydrate: notificationHydratePolicySettingTypeTrunkTitle},
			{Name: "policy_setting_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the policy setting type for this notification.", Hydrate: notificationHydratePolicySettingTypeUri},
			{Name: "policy_setting_value", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(formatPolicyFieldsValue), Description: "The value of the policy setting after this event.", Hydrate: notificationHydratePolicySettingValue},

			// Controls notification details
			{Name: "control_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the control for this notification.", Hydrate: notificationHydrateTurbotControlId},
			{Name: "control_new_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the control after the event.", Hydrate: notificationHydrateTurbotControlNewVersionId},
			{Name: "control_old_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the control before the event.", Hydrate: notificationHydrateTurbotControlOldVersionId},
			{Name: "control_details", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "Optional details provided at the last state update of this control.", Hydrate: notificationHydrateControlDetails},
			{Name: "control_reason", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Optional reason provided at the last state update of this control.", Hydrate: notificationHydrateControlReason},
			{Name: "control_state", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The current state of the control.", Hydrate: notificationHydrateControlState},
			{Name: "control_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the control type for this control.", Hydrate: notificationHydrateTurbotControlTypeId},
			{Name: "control_type_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "This is the title of hierarchy from the root down to this control type.", Hydrate: notificationHydrateControlTypeTrunkTitle},
			{Name: "control_type_uri", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "URI of the control type for this control.", Hydrate: notificationHydrateControlTypeUri},

			// ActiveGrants notification details
			{Name: "active_grant_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Active grant ID for this notification.", Hydrate: notificationHydrateActiveGrantId},
			{Name: "active_grant_new_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Active grant version ID of the grant after the notification.", Hydrate: notificationHydrateActiveGrantNewVersionId},
			{Name: "active_grant_old_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the active grant before the event.", Hydrate: notificationHydrateActiveGrantOldVersionId},
			{Name: "active_grant_valid_to_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Optional end date for the active grant to expire.", Hydrate: notificationHydrateActiveGrantValidToTimestamp},
			{Name: "active_grant_identity_profile_id", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The identity of profile id for this active grant.", Hydrate: notificationHydrateActiveGrantIdentityProfileId},
			{Name: "active_grant_identity_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "This is the title of hierarchy from the root down to this identity (i.e. Identity whoes access got revoked/permiited) for this active grant.", Hydrate: notificationHydrateActiveGrantIdentityTrunkTitle},
			{Name: "active_grant_level_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The name of the active grant level.", Hydrate: notificationHydrateActiveGrantLevelTitle},
			{Name: "active_grant_permission_level_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "The unique identifier for the active grant permission level.", Hydrate: notificationHydrateActiveGrantPermissionLevelId},
			{Name: "active_grant_permission_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "The unique identifier for the active grant permission type.", Hydrate: notificationHydrateActiveGrantPermissionTypeId},
			{Name: "active_grant_role_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Optional custom roleName for this active grant, when using existing roles rather than Turbot-managed ones.", Hydrate: notificationHydrateActiveGrantRoleName},
			{Name: "active_grant_type_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The name of the active grant type.", Hydrate: notificationHydrateActiveGrantTypeTitle},

			// Grants notification details
			{Name: "grant_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "ID of the grant for this notification.", Hydrate: notificationHydrateTurbotGrantId},
			{Name: "grant_new_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the grant after the event.", Hydrate: notificationHydrateTurbotGrantNewVersionId},
			{Name: "grant_old_version_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "Version ID of the grant before the event.", Hydrate: notificationHydrateTurbotGrantOldVersionId},
			{Name: "grant_valid_to_timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue(), Description: "Optional end date for the grant.", Hydrate: notificationHydrateTurbotGrantValidToTimestamp},
			{Name: "grant_identity_profile_id", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The identity profile id for this grant.", Hydrate: notificationHydrateTurbotGrantIdentityProfileId},
			{Name: "grant_identity_trunk_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "This is the title of hierarchy from the root down to this identity (i.e. Identity whoes access got revoked/permiited) for this grant.", Hydrate: notificationHydrateTurbotGrantIdentityTrunkTitle},
			{Name: "grant_level_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The name of the permission level.", Hydrate: notificationHydrateTurbotGrantLevelTitle},
			{Name: "grant_permission_level_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "The unique identifier for the permission level.", Hydrate: notificationHydrateTurbotGrantPermissionLevelId},
			{Name: "grant_permission_type_id", Type: proto.ColumnType_INT, Transform: transform.FromValue(), Description: "The unique identifier for the permission type.", Hydrate: notificationHydrateTurbotGrantPermissionTypeId},
			{Name: "grant_role_name", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "Optional custom roleName for this grant, when using existing roles rather than Turbot-managed ones.", Hydrate: notificationHydrateTurbotGrantRoleName},
			{Name: "grant_type_title", Type: proto.ColumnType_STRING, Transform: transform.FromValue(), Description: "The name of the permission type.", Hydrate: notificationHydrateTurbotGrantTypeTitle},
		},
	}
}

const (
	queryNotificationList = `
	   query notificationList($filter: [String!], $next_token: String, $includeNotificationIcon: Boolean!, $includeNotificationMessage: Boolean!, $includeNotificationType: Boolean!, $includeNotificationActorIdentityTrunkTitle: Boolean!, $includeNotificationActorIdentityTurbotId: Boolean!, $includeNotificationControlState: Boolean!, $includeNotificationControlReason: Boolean!, $includeNotificationControlDetails: Boolean!, $includeNotificationControlTypeUri: Boolean!, $includeNotificationControlTypeTrunkTitle: Boolean!, $includeNotificationControlTypeTurbotId: Boolean!, $includeNotificationResourceData: Boolean!, $includeNotificationResourceObject: Boolean!, $includeNotificationResourceTrunkTitle: Boolean!, $includeNotificationResourceTurbotAkas: Boolean!, $includeNotificationResourceTurbotParentId: Boolean!, $includeNotificationResourceTurbotPath: Boolean!, $includeNotificationResourceTurbotTags: Boolean!, $includeNotificationResourceTurbotTitle: Boolean!, $includeNotificationResourceTypeUri: Boolean!, $includeNotificationResourceTypeTrunkTitle: Boolean!, $includeNotificationResourceTypeTurbotId: Boolean!, $includeNotificationPolicySettingIsCalculated: Boolean!, $includeNotificationPolicySettingTypeUri: Boolean!, $includeNotificationPolicySettingTypeReadOnly: Boolean!, $includeNotificationPolicySettingTypeDefaultTemplate: Boolean!, $includeNotificationPolicySettingTypeDefaultTemplateInput: Boolean!, $includeNotificationPolicySettingTypeSecret: Boolean!, $includeNotificationPolicySettingTypeTrunkTitle: Boolean!, $includeNotificationPolicySettingTypeTurbotId: Boolean!, $includeNotificationPolicySettingValue: Boolean!, $includeNotificationGrantRoleName: Boolean!, $includeNotificationGrantPermissionTypeId: Boolean!, $includeNotificationGrantPermissionLevelId: Boolean!, $includeNotificationGrantValidToTimestamp: Boolean!, $includeNotificationGrantLevelTitle: Boolean!, $includeNotificationGrantTypeTitle: Boolean!, $includeNotificationGrantIdentityTrunkTitle: Boolean!, $includeNotificationGrantIdentityProfileId: Boolean!, $includeNotificationActiveGrantRoleName: Boolean!, $includeNotificationActiveGrantPermissionTypeId: Boolean!, $includeNotificationActiveGrantPermissionLevelId: Boolean!, $includeNotificationActiveGrantValidToTimestamp: Boolean!, $includeNotificationActiveGrantLevelTitle: Boolean!, $includeNotificationActiveGrantTypeTitle: Boolean!, $includeNotificationActiveGrantIdentityTrunkTitle: Boolean!, $includeNotificationActiveGrantIdentityProfileId: Boolean!, $includeNotificationTurbotControlId: Boolean!, $includeNotificationTurbotControlNewVersionId: Boolean!, $includeNotificationTurbotControlOldVersionId: Boolean!, $includeNotificationTurbotCreateTimestamp: Boolean!, $includeNotificationTurbotGrantId: Boolean!, $includeNotificationTurbotGrantNewVersionId: Boolean!, $includeNotificationTurbotGrantOldVersionId: Boolean!, $includeNotificationTurbotId: Boolean!, $includeNotificationTurbotPolicySettingId: Boolean!, $includeNotificationTurbotPolicySettingNewVersionId: Boolean!, $includeNotificationTurbotPolicySettingOldVersionId: Boolean!, $includeNotificationTurbotProcessId: Boolean!, $includeNotificationTurbotResourceId: Boolean!, $includeNotificationTurbotResourceNewVersionId: Boolean!, $includeNotificationTurbotResourceOldVersionId: Boolean!, $includeNotificationTurbotActiveGrantsId: Boolean!, $includeNotificationTurbotActiveGrantsNewVersionId: Boolean!, $includeNotificationTurbotActiveGrantsOldVersionId: Boolean!) {
            notifications(filter: $filter, paging: $next_token) {
                items {
                    icon @include(if: $includeNotificationIcon)
                    message @include(if: $includeNotificationMessage)
                    notificationType @include(if: $includeNotificationType)
                    actor {
                        identity {
                            trunk { title @include(if: $includeNotificationActorIdentityTrunkTitle) }
                            turbot {
                                id @include(if: $includeNotificationActorIdentityTurbotId)
                            }
                        }
                    }
                    control {
                        state @include(if: $includeNotificationControlState)
                        reason @include(if: $includeNotificationControlReason)
                        details @include(if: $includeNotificationControlDetails)
                        type {
                            uri @include(if: $includeNotificationControlTypeUri)
                            trunk {
                                title @include(if: $includeNotificationControlTypeTrunkTitle)
                            }
                            turbot {
                                id @include(if: $includeNotificationControlTypeTurbotId)
                            }
                        }
                    }
                    resource {
                        data @include(if: $includeNotificationResourceData)
                        object @include(if: $includeNotificationResourceObject)
                        trunk {
                            title @include(if: $includeNotificationResourceTrunkTitle)
                        }
                        turbot {
                            akas @include(if: $includeNotificationResourceTurbotAkas)
                            parentId @include(if: $includeNotificationResourceTurbotParentId)
                            path @include(if: $includeNotificationResourceTurbotPath)
                            tags @include(if: $includeNotificationResourceTurbotTags)
                            title @include(if: $includeNotificationResourceTurbotTitle)
                        }
                        type {
                            uri @include(if: $includeNotificationResourceTypeUri)
                            trunk {
                                title @include(if: $includeNotificationResourceTypeTrunkTitle)
                            }
                            turbot {
                                id @include(if: $includeNotificationResourceTypeTurbotId)
                            }
                        }
                    }
                    policySetting {
                        isCalculated @include(if: $includeNotificationPolicySettingIsCalculated)
                        type {
                            uri @include(if: $includeNotificationPolicySettingTypeUri)
                            readOnly @include(if: $includeNotificationPolicySettingTypeReadOnly)
                            defaultTemplate @include(if: $includeNotificationPolicySettingTypeDefaultTemplate)
                            defaultTemplateInput @include(if: $includeNotificationPolicySettingTypeDefaultTemplateInput)
                            secret @include(if: $includeNotificationPolicySettingTypeSecret)
                            trunk {
                                title @include(if: $includeNotificationPolicySettingTypeTrunkTitle)
                            }
                            turbot {
                                id @include(if: $includeNotificationPolicySettingTypeTurbotId)
                            }
                        }
                        value @include(if: $includeNotificationPolicySettingValue)
                    }
                    grant {
                        roleName @include(if: $includeNotificationGrantRoleName)
                        permissionTypeId @include(if: $includeNotificationGrantPermissionTypeId)
                        permissionLevelId @include(if: $includeNotificationGrantPermissionLevelId)
                        validToTimestamp @include(if: $includeNotificationGrantValidToTimestamp)
                        level {
                            title @include(if: $includeNotificationGrantLevelTitle)
                        }
                        type {
                            title @include(if: $includeNotificationGrantTypeTitle)
                        }
                        identity {
                            trunk { title @include(if: $includeNotificationGrantIdentityTrunkTitle) }
                            profileId: get(path: "profileId") @include(if: $includeNotificationGrantIdentityProfileId)
                        }
                    }
                    activeGrant {
                        grant {
                            roleName @include(if: $includeNotificationActiveGrantRoleName)
                            permissionTypeId @include(if: $includeNotificationActiveGrantPermissionTypeId)
                            permissionLevelId @include(if: $includeNotificationActiveGrantPermissionLevelId)
                            validToTimestamp @include(if: $includeNotificationActiveGrantValidToTimestamp)
                            level {
                                title @include(if: $includeNotificationActiveGrantLevelTitle)
                            }
                            type {
                                title @include(if: $includeNotificationActiveGrantTypeTitle)
                            }
                            identity {
                                trunk { title @include(if: $includeNotificationActiveGrantIdentityTrunkTitle) }
                                profileId: get(path: "profileId") @include(if: $includeNotificationActiveGrantIdentityProfileId)
                            }
                        }
                    }
                    turbot {
                        controlId @include(if: $includeNotificationTurbotControlId)
                        controlNewVersionId @include(if: $includeNotificationTurbotControlNewVersionId)
                        controlOldVersionId @include(if: $includeNotificationTurbotControlOldVersionId)
                        createTimestamp @include(if: $includeNotificationTurbotCreateTimestamp)
                        grantId @include(if: $includeNotificationTurbotGrantId)
                        grantNewVersionId @include(if: $includeNotificationTurbotGrantNewVersionId)
                        grantOldVersionId @include(if: $includeNotificationTurbotGrantOldVersionId)
                        id @include(if: $includeNotificationTurbotId)
                        policySettingId @include(if: $includeNotificationTurbotPolicySettingId)
                        policySettingNewVersionId @include(if: $includeNotificationTurbotPolicySettingNewVersionId)
                        policySettingOldVersionId @include(if: $includeNotificationTurbotPolicySettingOldVersionId)
                        processId @include(if: $includeNotificationTurbotProcessId)
                        resourceId @include(if: $includeNotificationTurbotResourceId)
                        resourceNewVersionId @include(if: $includeNotificationTurbotResourceNewVersionId)
                        resourceOldVersionId @include(if: $includeNotificationTurbotResourceOldVersionId)
                        activeGrantsId @include(if: $includeNotificationTurbotActiveGrantsId)
                        activeGrantsNewVersionId @include(if: $includeNotificationTurbotActiveGrantsNewVersionId)
                        activeGrantsOldVersionId @include(if: $includeNotificationTurbotActiveGrantsOldVersionId)
                    }
                }
                paging {
                    next
                }
            }
        }
	`

	queryNotificationGet = `
		query notificationGet($id: ID!, $includeNotificationIcon: Boolean!, $includeNotificationMessage: Boolean!, $includeNotificationType: Boolean!, $includeNotificationActorIdentityTrunkTitle: Boolean!, $includeNotificationActorIdentityTurbotId: Boolean!, $includeNotificationControlState: Boolean!, $includeNotificationControlReason: Boolean!, $includeNotificationControlDetails: Boolean!, $includeNotificationControlTypeUri: Boolean!, $includeNotificationControlTypeTrunkTitle: Boolean!, $includeNotificationControlTypeTurbotId: Boolean!, $includeNotificationResourceData: Boolean!, $includeNotificationResourceObject: Boolean!, $includeNotificationResourceTrunkTitle: Boolean!, $includeNotificationResourceTurbotAkas: Boolean!, $includeNotificationResourceTurbotParentId: Boolean!, $includeNotificationResourceTurbotPath: Boolean!, $includeNotificationResourceTurbotTags: Boolean!, $includeNotificationResourceTurbotTitle: Boolean!, $includeNotificationResourceTypeUri: Boolean!, $includeNotificationResourceTypeTrunkTitle: Boolean!, $includeNotificationResourceTypeTurbotId: Boolean!, $includeNotificationPolicySettingIsCalculated: Boolean!, $includeNotificationPolicySettingTypeUri: Boolean!, $includeNotificationPolicySettingTypeReadOnly: Boolean!, $includeNotificationPolicySettingTypeDefaultTemplate: Boolean!, $includeNotificationPolicySettingTypeDefaultTemplateInput: Boolean!, $includeNotificationPolicySettingTypeSecret: Boolean!, $includeNotificationPolicySettingTypeTrunkTitle: Boolean!, $includeNotificationPolicySettingTypeTurbotId: Boolean!, $includeNotificationPolicySettingValue: Boolean!, $includeNotificationGrantRoleName: Boolean!, $includeNotificationGrantPermissionTypeId: Boolean!, $includeNotificationGrantPermissionLevelId: Boolean!, $includeNotificationGrantValidToTimestamp: Boolean!, $includeNotificationGrantLevelTitle: Boolean!, $includeNotificationGrantTypeTitle: Boolean!, $includeNotificationGrantIdentityTrunkTitle: Boolean!, $includeNotificationGrantIdentityProfileId: Boolean!, $includeNotificationActiveGrantRoleName: Boolean!, $includeNotificationActiveGrantPermissionTypeId: Boolean!, $includeNotificationActiveGrantPermissionLevelId: Boolean!, $includeNotificationActiveGrantValidToTimestamp: Boolean!, $includeNotificationActiveGrantLevelTitle: Boolean!, $includeNotificationActiveGrantTypeTitle: Boolean!, $includeNotificationActiveGrantIdentityTrunkTitle: Boolean!, $includeNotificationActiveGrantIdentityProfileId: Boolean!, $includeNotificationTurbotControlId: Boolean!, $includeNotificationTurbotControlNewVersionId: Boolean!, $includeNotificationTurbotControlOldVersionId: Boolean!, $includeNotificationTurbotCreateTimestamp: Boolean!, $includeNotificationTurbotGrantId: Boolean!, $includeNotificationTurbotGrantNewVersionId: Boolean!, $includeNotificationTurbotGrantOldVersionId: Boolean!, $includeNotificationTurbotId: Boolean!, $includeNotificationTurbotPolicySettingId: Boolean!, $includeNotificationTurbotPolicySettingNewVersionId: Boolean!, $includeNotificationTurbotPolicySettingOldVersionId: Boolean!, $includeNotificationTurbotProcessId: Boolean!, $includeNotificationTurbotResourceId: Boolean!, $includeNotificationTurbotResourceNewVersionId: Boolean!, $includeNotificationTurbotResourceOldVersionId: Boolean!, $includeNotificationTurbotActiveGrantsId: Boolean!, $includeNotificationTurbotActiveGrantsNewVersionId: Boolean!, $includeNotificationTurbotActiveGrantsOldVersionId: Boolean!) {
  notification(id: $id) {
    icon @include(if: $includeNotificationIcon)
    message @include(if: $includeNotificationMessage)
    notificationType @include(if: $includeNotificationType)
    actor {
      identity {
        trunk {
          title @include(if: $includeNotificationActorIdentityTrunkTitle)
        }
        turbot {
          id @include(if: $includeNotificationActorIdentityTurbotId)
        }
      }
    }
    control {
      state @include(if: $includeNotificationControlState)
      reason @include(if: $includeNotificationControlReason)
      details @include(if: $includeNotificationControlDetails)
      type {
        uri @include(if: $includeNotificationControlTypeUri)
        trunk {
          title @include(if: $includeNotificationControlTypeTrunkTitle)
        }
        turbot {
          id @include(if: $includeNotificationControlTypeTurbotId)
        }
      }
    }
    resource {
      data @include(if: $includeNotificationResourceData)
      object @include(if: $includeNotificationResourceObject)
      trunk {
        title @include(if: $includeNotificationResourceTrunkTitle)
      }
      turbot {
        akas @include(if: $includeNotificationResourceTurbotAkas)
        parentId @include(if: $includeNotificationResourceTurbotParentId)
        path @include(if: $includeNotificationResourceTurbotPath)
        tags @include(if: $includeNotificationResourceTurbotTags)
        title @include(if: $includeNotificationResourceTurbotTitle)
      }
      type {
        uri @include(if: $includeNotificationResourceTypeUri)
        trunk {
          title @include(if: $includeNotificationResourceTypeTrunkTitle)
        }
        turbot {
          id @include(if: $includeNotificationResourceTypeTurbotId)
        }
      }
    }
    policySetting {
      isCalculated @include(if: $includeNotificationPolicySettingIsCalculated)
      type {
        uri @include(if: $includeNotificationPolicySettingTypeUri)
        readOnly @include(if: $includeNotificationPolicySettingTypeReadOnly)
        defaultTemplate @include(if: $includeNotificationPolicySettingTypeDefaultTemplate)
        defaultTemplateInput @include(if: $includeNotificationPolicySettingTypeDefaultTemplateInput)
        secret @include(if: $includeNotificationPolicySettingTypeSecret)
        trunk {
          title @include(if: $includeNotificationPolicySettingTypeTrunkTitle)
        }
        turbot {
          id @include(if: $includeNotificationPolicySettingTypeTurbotId)
        }
      }
      value @include(if: $includeNotificationPolicySettingValue)
    }
    grant {
      roleName @include(if: $includeNotificationGrantRoleName)
      permissionTypeId @include(if: $includeNotificationGrantPermissionTypeId)
      permissionLevelId @include(if: $includeNotificationGrantPermissionLevelId)
      validToTimestamp @include(if: $includeNotificationGrantValidToTimestamp)
      level {
        title @include(if: $includeNotificationGrantLevelTitle)
      }
      type {
        title @include(if: $includeNotificationGrantTypeTitle)
      }
      identity {
        trunk {
          title @include(if: $includeNotificationGrantIdentityTrunkTitle)
        }
        profileId: get(path: "profileId") @include(if: $includeNotificationGrantIdentityProfileId)
      }
    }
    activeGrant {
      grant {
        roleName @include(if: $includeNotificationActiveGrantRoleName)
        permissionTypeId @include(if: $includeNotificationActiveGrantPermissionTypeId)
        permissionLevelId @include(if: $includeNotificationActiveGrantPermissionLevelId)
        validToTimestamp @include(if: $includeNotificationActiveGrantValidToTimestamp)
        level {
          title @include(if: $includeNotificationActiveGrantLevelTitle)
        }
        type {
          title @include(if: $includeNotificationActiveGrantTypeTitle)
        }
        identity {
          trunk {
            title @include(if: $includeNotificationActiveGrantIdentityTrunkTitle)
          }
          profileId: get(path: "profileId") @include(if: $includeNotificationActiveGrantIdentityProfileId)
        }
      }
    }
    turbot {
      controlId @include(if: $includeNotificationTurbotControlId)
      controlNewVersionId @include(if: $includeNotificationTurbotControlNewVersionId)
      controlOldVersionId @include(if: $includeNotificationTurbotControlOldVersionId)
      createTimestamp @include(if: $includeNotificationTurbotCreateTimestamp)
      grantId @include(if: $includeNotificationTurbotGrantId)
      grantNewVersionId @include(if: $includeNotificationTurbotGrantNewVersionId)
      grantOldVersionId @include(if: $includeNotificationTurbotGrantOldVersionId)
      id @include(if: $includeNotificationTurbotId)
      policySettingId @include(if: $includeNotificationTurbotPolicySettingId)
      policySettingNewVersionId @include(if: $includeNotificationTurbotPolicySettingNewVersionId)
      policySettingOldVersionId @include(if: $includeNotificationTurbotPolicySettingOldVersionId)
      processId @include(if: $includeNotificationTurbotProcessId)
      resourceId @include(if: $includeNotificationTurbotResourceId)
      resourceNewVersionId @include(if: $includeNotificationTurbotResourceNewVersionId)
      resourceOldVersionId @include(if: $includeNotificationTurbotResourceOldVersionId)
      activeGrantsId @include(if: $includeNotificationTurbotActiveGrantsId)
      activeGrantsNewVersionId @include(if: $includeNotificationTurbotActiveGrantsNewVersionId)
      activeGrantsOldVersionId @include(if: $includeNotificationTurbotActiveGrantsOldVersionId)
    }
  }
}
`
)

func listNotification(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_notification.listNotification", "connection_error", err)
		return nil, err
	}

	//build the Quals/Filters for the query
	filters := []string{}
	quals := d.EqualsQuals
	allQuals := d.Quals
	filter := ""
	if quals["filter"] != nil {
		filter = quals["filter"].GetStringValue()
		filters = append(filters, filter)
	}
	if quals["id"] != nil {
		filters = append(filters, fmt.Sprintf("id:%s", getQualListValues(ctx, quals, "id", "int64")))
	}

	if quals["notification_type"] != nil {
		filters = append(filters, fmt.Sprintf("notificationType:%s", getQualListValues(ctx, quals, "notification_type", "string")))
	}

	if quals["actor_identity_id"] != nil {
		filters = append(filters, fmt.Sprintf("actorIdentityId:%s", getQualListValues(ctx, quals, "actor_identity_id", "int64")))
	}

	if quals["resource_id"] != nil {
		filters = append(filters, fmt.Sprintf("resourceId:%s", getQualListValues(ctx, quals, "resource_id", "int64")))
	}

	if quals["resource_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_type_id", "int64")))
	}

	if quals["resource_type_uri"] != nil {
		filters = append(filters, fmt.Sprintf("resourceTypeId:%s resourceTypeLevel:self", getQualListValues(ctx, quals, "resource_type_uri", "string")))
	}

	if quals["control_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("controlTypeId:%s controlTypeLevel:self", getQualListValues(ctx, quals, "control_type_id", "int64")))
	}

	if quals["control_type_uri"] != nil {
		filters = append(filters, fmt.Sprintf("controlTypeId:%s controlTypeLevel:self", getQualListValues(ctx, quals, "control_type_uri", "string")))
	}

	if quals["policy_type_id"] != nil {
		filters = append(filters, fmt.Sprintf("policyTypeId:%s policyTypeLevel:self", getQualListValues(ctx, quals, "policy_type_id", "int64")))
	}

	if quals["policy_type_uri"] != nil {
		filters = append(filters, fmt.Sprintf("policyTypeId:%s policyTypeLevel:self", getQualListValues(ctx, quals, "policy_type_uri", "string")))
	}

	if allQuals["create_timestamp"] != nil {
		for _, q := range allQuals["create_timestamp"].Quals {
			// Subtracted 1 minute to FilterFrom time and Added 1 minute to FilterTo time to miss any results due to time conersions in steampipe
			switch q.Operator {
			case "=":
				filters = append(filters, fmt.Sprintf("createTimestamp:'%s'", q.Value.GetTimestampValue().AsTime().Format(filterTimeFormat)))
			case ">=", ">":
				filters = append(filters, fmt.Sprintf("createTimestamp:>='%s'", q.Value.GetTimestampValue().AsTime().Add(-1*time.Minute).Format(filterTimeFormat)))
			case "<", "<=":
				filters = append(filters, fmt.Sprintf("createTimestamp:<='%s'", q.Value.GetTimestampValue().AsTime().Add(1*time.Minute).Format(filterTimeFormat)))
			}
		}
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

	plugin.Logger(ctx).Warn("guardrails_notification.listNotification", "filters", filters)

	variables := map[string]interface{}{
		"filter":     filters,
		"next_token": "",
	}

	appendNotificationColumnIncludes(&variables, d.QueryContext.Columns)
	for {
		result := &NotificationsResponse{}
		err = conn.DoRequest(queryNotificationList, variables, result)
		if err != nil {
			plugin.Logger(ctx).Error("guardrails_notification.listNotification", "query_error", err)
			// Not returning for function in case of errors because of resources/policies/controls referred might be deleted and
			// graphql queries may fail to retrieve few properties for such items
			// return nil, err
		}
		for _, r := range result.Notifications.Items {
			d.StreamListItem(ctx, r)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !pageResults || result.Notifications.Paging.Next == "" {
			break
		}
		variables["next_token"] = result.Notifications.Paging.Next
	}

	return nil, nil
}

func getNotification(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_notification.getNotification", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()
	variables := map[string]interface{}{
		"id": id,
	}

	appendNotificationColumnIncludes(&variables, d.QueryContext.Columns)

	result := &NotificationsGetResponse{}
	err = conn.DoRequest(queryNotificationGet, variables, result)
	if err != nil {
		plugin.Logger(ctx).Error("guardrails_notification.getNotification", "query_error", err)
		return nil, err
	}
	return result.Notification, nil
}

//// TRANSFORM FUNCTION

// formatPolicyValue:: Policy value can be a string, hcl or a json.
// It will transform the raw value from api into a string if a hcl or json
func formatPolicyFieldsValue(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var item = d.HydrateItem.(Notification)
	columnName := d.ColumnName
	var value interface{}

	if item.PolicySetting != nil {
		if columnName == "policy_template_input" {
			value = item.PolicySetting.Type.DefaultTemplateInput
		} else {
			value = item.PolicySetting.Value
		}
	}

	if value != nil {
		switch val := value.(type) {
		case string:
			return val, nil
		case []string, map[string]interface{}, interface{}:
			data, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			return string(data), nil
		}
	}

	return nil, nil
}

// fromField:: generates a value by retrieving a field or a set of fields from the source item
func fromField(fieldNames ...string) *transform.ColumnTransforms {
	var fieldNameArray []string
	fieldNameArray = append(fieldNameArray, fieldNames...)
	return &transform.ColumnTransforms{Transforms: []*transform.TransformCall{{Transform: fieldValue, Param: fieldNameArray}}}
}

// fieldValue function is intended for the start of a transform chain.
// This returns a field value of either the hydrate call result (if present)  or the root item if not
// the field name is in the 'Param'
func fieldValue(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	var item = d.HydrateItem
	var fieldNames []string

	switch p := d.Param.(type) {
	case []string:
		fieldNames = p
	case string:
		fieldNames = []string{p}
	default:
		return nil, fmt.Errorf("'FieldValue' requires one or more string parameters containing property path but received %v", d.Param)
	}

	for _, propertyPath := range fieldNames {
		fieldValue, ok := helpers.GetNestedFieldValueFromInterface(item, propertyPath)
		if ok && !helpers.IsNil(fieldValue) {
			return fieldValue, nil

		}

	}
	return nil, nil
}
