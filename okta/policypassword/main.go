// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policypassword

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyPassword.PolicyPassword",
		reflect.TypeOf((*PolicyPassword)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authProvider", GoGetter: "AuthProvider"},
			_jsii_.MemberProperty{JsiiProperty: "authProviderInput", GoGetter: "AuthProviderInput"},
			_jsii_.MemberProperty{JsiiProperty: "callRecovery", GoGetter: "CallRecovery"},
			_jsii_.MemberProperty{JsiiProperty: "callRecoveryInput", GoGetter: "CallRecoveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailRecovery", GoGetter: "EmailRecovery"},
			_jsii_.MemberProperty{JsiiProperty: "emailRecoveryInput", GoGetter: "EmailRecoveryInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupsIncluded", GoGetter: "GroupsIncluded"},
			_jsii_.MemberProperty{JsiiProperty: "groupsIncludedInput", GoGetter: "GroupsIncludedInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordAutoUnlockMinutes", GoGetter: "PasswordAutoUnlockMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "passwordAutoUnlockMinutesInput", GoGetter: "PasswordAutoUnlockMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordDictionaryLookup", GoGetter: "PasswordDictionaryLookup"},
			_jsii_.MemberProperty{JsiiProperty: "passwordDictionaryLookupInput", GoGetter: "PasswordDictionaryLookupInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExcludeFirstName", GoGetter: "PasswordExcludeFirstName"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExcludeFirstNameInput", GoGetter: "PasswordExcludeFirstNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExcludeLastName", GoGetter: "PasswordExcludeLastName"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExcludeLastNameInput", GoGetter: "PasswordExcludeLastNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExcludeUsername", GoGetter: "PasswordExcludeUsername"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExcludeUsernameInput", GoGetter: "PasswordExcludeUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExpireWarnDays", GoGetter: "PasswordExpireWarnDays"},
			_jsii_.MemberProperty{JsiiProperty: "passwordExpireWarnDaysInput", GoGetter: "PasswordExpireWarnDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordHistoryCount", GoGetter: "PasswordHistoryCount"},
			_jsii_.MemberProperty{JsiiProperty: "passwordHistoryCountInput", GoGetter: "PasswordHistoryCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordLockoutNotificationChannels", GoGetter: "PasswordLockoutNotificationChannels"},
			_jsii_.MemberProperty{JsiiProperty: "passwordLockoutNotificationChannelsInput", GoGetter: "PasswordLockoutNotificationChannelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMaxAgeDays", GoGetter: "PasswordMaxAgeDays"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMaxAgeDaysInput", GoGetter: "PasswordMaxAgeDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMaxLockoutAttempts", GoGetter: "PasswordMaxLockoutAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMaxLockoutAttemptsInput", GoGetter: "PasswordMaxLockoutAttemptsInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinAgeMinutes", GoGetter: "PasswordMinAgeMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinAgeMinutesInput", GoGetter: "PasswordMinAgeMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinLength", GoGetter: "PasswordMinLength"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinLengthInput", GoGetter: "PasswordMinLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinLowercase", GoGetter: "PasswordMinLowercase"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinLowercaseInput", GoGetter: "PasswordMinLowercaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinNumber", GoGetter: "PasswordMinNumber"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinNumberInput", GoGetter: "PasswordMinNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinSymbol", GoGetter: "PasswordMinSymbol"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinSymbolInput", GoGetter: "PasswordMinSymbolInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinUppercase", GoGetter: "PasswordMinUppercase"},
			_jsii_.MemberProperty{JsiiProperty: "passwordMinUppercaseInput", GoGetter: "PasswordMinUppercaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordShowLockoutFailures", GoGetter: "PasswordShowLockoutFailures"},
			_jsii_.MemberProperty{JsiiProperty: "passwordShowLockoutFailuresInput", GoGetter: "PasswordShowLockoutFailuresInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "questionMinLength", GoGetter: "QuestionMinLength"},
			_jsii_.MemberProperty{JsiiProperty: "questionMinLengthInput", GoGetter: "QuestionMinLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "questionRecovery", GoGetter: "QuestionRecovery"},
			_jsii_.MemberProperty{JsiiProperty: "questionRecoveryInput", GoGetter: "QuestionRecoveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryEmailToken", GoGetter: "RecoveryEmailToken"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryEmailTokenInput", GoGetter: "RecoveryEmailTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthProvider", GoMethod: "ResetAuthProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resetCallRecovery", GoMethod: "ResetCallRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailRecovery", GoMethod: "ResetEmailRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupsIncluded", GoMethod: "ResetGroupsIncluded"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordAutoUnlockMinutes", GoMethod: "ResetPasswordAutoUnlockMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordDictionaryLookup", GoMethod: "ResetPasswordDictionaryLookup"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordExcludeFirstName", GoMethod: "ResetPasswordExcludeFirstName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordExcludeLastName", GoMethod: "ResetPasswordExcludeLastName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordExcludeUsername", GoMethod: "ResetPasswordExcludeUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordExpireWarnDays", GoMethod: "ResetPasswordExpireWarnDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordHistoryCount", GoMethod: "ResetPasswordHistoryCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordLockoutNotificationChannels", GoMethod: "ResetPasswordLockoutNotificationChannels"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMaxAgeDays", GoMethod: "ResetPasswordMaxAgeDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMaxLockoutAttempts", GoMethod: "ResetPasswordMaxLockoutAttempts"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMinAgeMinutes", GoMethod: "ResetPasswordMinAgeMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMinLength", GoMethod: "ResetPasswordMinLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMinLowercase", GoMethod: "ResetPasswordMinLowercase"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMinNumber", GoMethod: "ResetPasswordMinNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMinSymbol", GoMethod: "ResetPasswordMinSymbol"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordMinUppercase", GoMethod: "ResetPasswordMinUppercase"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordShowLockoutFailures", GoMethod: "ResetPasswordShowLockoutFailures"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetQuestionMinLength", GoMethod: "ResetQuestionMinLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetQuestionRecovery", GoMethod: "ResetQuestionRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecoveryEmailToken", GoMethod: "ResetRecoveryEmailToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipUnlock", GoMethod: "ResetSkipUnlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmsRecovery", GoMethod: "ResetSmsRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberProperty{JsiiProperty: "skipUnlock", GoGetter: "SkipUnlock"},
			_jsii_.MemberProperty{JsiiProperty: "skipUnlockInput", GoGetter: "SkipUnlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "smsRecovery", GoGetter: "SmsRecovery"},
			_jsii_.MemberProperty{JsiiProperty: "smsRecoveryInput", GoGetter: "SmsRecoveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyPassword{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.policyPassword.PolicyPasswordConfig",
		reflect.TypeOf((*PolicyPasswordConfig)(nil)).Elem(),
	)
}
