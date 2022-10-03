package swaapp

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.swaApp.SwaApp",
		reflect.TypeOf((*SwaApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessibilityErrorRedirectUrl", GoGetter: "AccessibilityErrorRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilityErrorRedirectUrlInput", GoGetter: "AccessibilityErrorRedirectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilityLoginRedirectUrl", GoGetter: "AccessibilityLoginRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilityLoginRedirectUrlInput", GoGetter: "AccessibilityLoginRedirectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilitySelfService", GoGetter: "AccessibilitySelfService"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilitySelfServiceInput", GoGetter: "AccessibilitySelfServiceInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminNote", GoGetter: "AdminNote"},
			_jsii_.MemberProperty{JsiiProperty: "adminNoteInput", GoGetter: "AdminNoteInput"},
			_jsii_.MemberProperty{JsiiProperty: "appLinksJson", GoGetter: "AppLinksJson"},
			_jsii_.MemberProperty{JsiiProperty: "appLinksJsonInput", GoGetter: "AppLinksJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbar", GoGetter: "AutoSubmitToolbar"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbarInput", GoGetter: "AutoSubmitToolbarInput"},
			_jsii_.MemberProperty{JsiiProperty: "buttonField", GoGetter: "ButtonField"},
			_jsii_.MemberProperty{JsiiProperty: "buttonFieldInput", GoGetter: "ButtonFieldInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "checkbox", GoGetter: "Checkbox"},
			_jsii_.MemberProperty{JsiiProperty: "checkboxInput", GoGetter: "CheckboxInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enduserNote", GoGetter: "EnduserNote"},
			_jsii_.MemberProperty{JsiiProperty: "enduserNoteInput", GoGetter: "EnduserNoteInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberProperty{JsiiProperty: "groupsInput", GoGetter: "GroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hideIos", GoGetter: "HideIos"},
			_jsii_.MemberProperty{JsiiProperty: "hideIosInput", GoGetter: "HideIosInput"},
			_jsii_.MemberProperty{JsiiProperty: "hideWeb", GoGetter: "HideWeb"},
			_jsii_.MemberProperty{JsiiProperty: "hideWebInput", GoGetter: "HideWebInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logo", GoGetter: "Logo"},
			_jsii_.MemberProperty{JsiiProperty: "logoInput", GoGetter: "LogoInput"},
			_jsii_.MemberProperty{JsiiProperty: "logoUrl", GoGetter: "LogoUrl"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordField", GoGetter: "PasswordField"},
			_jsii_.MemberProperty{JsiiProperty: "passwordFieldInput", GoGetter: "PasswordFieldInput"},
			_jsii_.MemberProperty{JsiiProperty: "preconfiguredApp", GoGetter: "PreconfiguredApp"},
			_jsii_.MemberProperty{JsiiProperty: "preconfiguredAppInput", GoGetter: "PreconfiguredAppInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putUsers", GoMethod: "PutUsers"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUrl", GoGetter: "RedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUrlInput", GoGetter: "RedirectUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilityErrorRedirectUrl", GoMethod: "ResetAccessibilityErrorRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilityLoginRedirectUrl", GoMethod: "ResetAccessibilityLoginRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilitySelfService", GoMethod: "ResetAccessibilitySelfService"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminNote", GoMethod: "ResetAdminNote"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLinksJson", GoMethod: "ResetAppLinksJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoSubmitToolbar", GoMethod: "ResetAutoSubmitToolbar"},
			_jsii_.MemberMethod{JsiiMethod: "resetButtonField", GoMethod: "ResetButtonField"},
			_jsii_.MemberMethod{JsiiMethod: "resetCheckbox", GoMethod: "ResetCheckbox"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnduserNote", GoMethod: "ResetEnduserNote"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroups", GoMethod: "ResetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideIos", GoMethod: "ResetHideIos"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideWeb", GoMethod: "ResetHideWeb"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogo", GoMethod: "ResetLogo"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordField", GoMethod: "ResetPasswordField"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreconfiguredApp", GoMethod: "ResetPreconfiguredApp"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectUrl", GoMethod: "ResetRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipGroups", GoMethod: "ResetSkipGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipUsers", GoMethod: "ResetSkipUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrl", GoMethod: "ResetUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRegex", GoMethod: "ResetUrlRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsernameField", GoMethod: "ResetUsernameField"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplate", GoMethod: "ResetUserNameTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplatePushStatus", GoMethod: "ResetUserNameTemplatePushStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplateSuffix", GoMethod: "ResetUserNameTemplateSuffix"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplateType", GoMethod: "ResetUserNameTemplateType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsers", GoMethod: "ResetUsers"},
			_jsii_.MemberProperty{JsiiProperty: "signOnMode", GoGetter: "SignOnMode"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroups", GoGetter: "SkipGroups"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroupsInput", GoGetter: "SkipGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipUsers", GoGetter: "SkipUsers"},
			_jsii_.MemberProperty{JsiiProperty: "skipUsersInput", GoGetter: "SkipUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "urlRegex", GoGetter: "UrlRegex"},
			_jsii_.MemberProperty{JsiiProperty: "urlRegexInput", GoGetter: "UrlRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "usernameField", GoGetter: "UsernameField"},
			_jsii_.MemberProperty{JsiiProperty: "usernameFieldInput", GoGetter: "UsernameFieldInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplate", GoGetter: "UserNameTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateInput", GoGetter: "UserNameTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplatePushStatus", GoGetter: "UserNameTemplatePushStatus"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplatePushStatusInput", GoGetter: "UserNameTemplatePushStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateSuffix", GoGetter: "UserNameTemplateSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateSuffixInput", GoGetter: "UserNameTemplateSuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateType", GoGetter: "UserNameTemplateType"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateTypeInput", GoGetter: "UserNameTemplateTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
			_jsii_.MemberProperty{JsiiProperty: "usersInput", GoGetter: "UsersInput"},
		},
		func() interface{} {
			j := jsiiProxy_SwaApp{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.swaApp.SwaAppConfig",
		reflect.TypeOf((*SwaAppConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.swaApp.SwaAppTimeouts",
		reflect.TypeOf((*SwaAppTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.swaApp.SwaAppTimeoutsOutputReference",
		reflect.TypeOf((*SwaAppTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_SwaAppTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.swaApp.SwaAppUsers",
		reflect.TypeOf((*SwaAppUsers)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.swaApp.SwaAppUsersList",
		reflect.TypeOf((*SwaAppUsersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SwaAppUsersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.swaApp.SwaAppUsersOutputReference",
		reflect.TypeOf((*SwaAppUsersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_SwaAppUsersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
