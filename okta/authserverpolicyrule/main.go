// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authserverpolicyrule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
		reflect.TypeOf((*AuthServerPolicyRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessTokenLifetimeMinutes", GoGetter: "AccessTokenLifetimeMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "accessTokenLifetimeMinutesInput", GoGetter: "AccessTokenLifetimeMinutesInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authServerId", GoGetter: "AuthServerId"},
			_jsii_.MemberProperty{JsiiProperty: "authServerIdInput", GoGetter: "AuthServerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantTypeWhitelist", GoGetter: "GrantTypeWhitelist"},
			_jsii_.MemberProperty{JsiiProperty: "grantTypeWhitelistInput", GoGetter: "GrantTypeWhitelistInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupBlacklist", GoGetter: "GroupBlacklist"},
			_jsii_.MemberProperty{JsiiProperty: "groupBlacklistInput", GoGetter: "GroupBlacklistInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupWhitelist", GoGetter: "GroupWhitelist"},
			_jsii_.MemberProperty{JsiiProperty: "groupWhitelistInput", GoGetter: "GroupWhitelistInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "inlineHookId", GoGetter: "InlineHookId"},
			_jsii_.MemberProperty{JsiiProperty: "inlineHookIdInput", GoGetter: "InlineHookIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdInput", GoGetter: "PolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenLifetimeMinutes", GoGetter: "RefreshTokenLifetimeMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenLifetimeMinutesInput", GoGetter: "RefreshTokenLifetimeMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenWindowMinutes", GoGetter: "RefreshTokenWindowMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenWindowMinutesInput", GoGetter: "RefreshTokenWindowMinutesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessTokenLifetimeMinutes", GoMethod: "ResetAccessTokenLifetimeMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupBlacklist", GoMethod: "ResetGroupBlacklist"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupWhitelist", GoMethod: "ResetGroupWhitelist"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInlineHookId", GoMethod: "ResetInlineHookId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenLifetimeMinutes", GoMethod: "ResetRefreshTokenLifetimeMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenWindowMinutes", GoMethod: "ResetRefreshTokenWindowMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopeWhitelist", GoMethod: "ResetScopeWhitelist"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserBlacklist", GoMethod: "ResetUserBlacklist"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserWhitelist", GoMethod: "ResetUserWhitelist"},
			_jsii_.MemberProperty{JsiiProperty: "scopeWhitelist", GoGetter: "ScopeWhitelist"},
			_jsii_.MemberProperty{JsiiProperty: "scopeWhitelistInput", GoGetter: "ScopeWhitelistInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "userBlacklist", GoGetter: "UserBlacklist"},
			_jsii_.MemberProperty{JsiiProperty: "userBlacklistInput", GoGetter: "UserBlacklistInput"},
			_jsii_.MemberProperty{JsiiProperty: "userWhitelist", GoGetter: "UserWhitelist"},
			_jsii_.MemberProperty{JsiiProperty: "userWhitelistInput", GoGetter: "UserWhitelistInput"},
		},
		func() interface{} {
			j := jsiiProxy_AuthServerPolicyRule{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRuleConfig",
		reflect.TypeOf((*AuthServerPolicyRuleConfig)(nil)).Elem(),
	)
}
