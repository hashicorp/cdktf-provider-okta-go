package appoauth

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.appOauth.AppOauth",
		reflect.TypeOf((*AppOauth)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "appSettingsJson", GoGetter: "AppSettingsJson"},
			_jsii_.MemberProperty{JsiiProperty: "appSettingsJsonInput", GoGetter: "AppSettingsJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationPolicy", GoGetter: "AuthenticationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationPolicyInput", GoGetter: "AuthenticationPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoKeyRotation", GoGetter: "AutoKeyRotation"},
			_jsii_.MemberProperty{JsiiProperty: "autoKeyRotationInput", GoGetter: "AutoKeyRotationInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbar", GoGetter: "AutoSubmitToolbar"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbarInput", GoGetter: "AutoSubmitToolbarInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientBasicSecret", GoGetter: "ClientBasicSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientBasicSecretInput", GoGetter: "ClientBasicSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientUri", GoGetter: "ClientUri"},
			_jsii_.MemberProperty{JsiiProperty: "clientUriInput", GoGetter: "ClientUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "consentMethod", GoGetter: "ConsentMethod"},
			_jsii_.MemberProperty{JsiiProperty: "consentMethodInput", GoGetter: "ConsentMethodInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "grantTypes", GoGetter: "GrantTypes"},
			_jsii_.MemberProperty{JsiiProperty: "grantTypesInput", GoGetter: "GrantTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupsClaim", GoGetter: "GroupsClaim"},
			_jsii_.MemberProperty{JsiiProperty: "groupsClaimInput", GoGetter: "GroupsClaimInput"},
			_jsii_.MemberProperty{JsiiProperty: "hideIos", GoGetter: "HideIos"},
			_jsii_.MemberProperty{JsiiProperty: "hideIosInput", GoGetter: "HideIosInput"},
			_jsii_.MemberProperty{JsiiProperty: "hideWeb", GoGetter: "HideWeb"},
			_jsii_.MemberProperty{JsiiProperty: "hideWebInput", GoGetter: "HideWebInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "implicitAssignment", GoGetter: "ImplicitAssignment"},
			_jsii_.MemberProperty{JsiiProperty: "implicitAssignmentInput", GoGetter: "ImplicitAssignmentInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuerMode", GoGetter: "IssuerMode"},
			_jsii_.MemberProperty{JsiiProperty: "issuerModeInput", GoGetter: "IssuerModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwks", GoGetter: "Jwks"},
			_jsii_.MemberProperty{JsiiProperty: "jwksInput", GoGetter: "JwksInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwksUri", GoGetter: "JwksUri"},
			_jsii_.MemberProperty{JsiiProperty: "jwksUriInput", GoGetter: "JwksUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loginMode", GoGetter: "LoginMode"},
			_jsii_.MemberProperty{JsiiProperty: "loginModeInput", GoGetter: "LoginModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "loginScopes", GoGetter: "LoginScopes"},
			_jsii_.MemberProperty{JsiiProperty: "loginScopesInput", GoGetter: "LoginScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "loginUri", GoGetter: "LoginUri"},
			_jsii_.MemberProperty{JsiiProperty: "loginUriInput", GoGetter: "LoginUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "logo", GoGetter: "Logo"},
			_jsii_.MemberProperty{JsiiProperty: "logoInput", GoGetter: "LogoInput"},
			_jsii_.MemberProperty{JsiiProperty: "logoUri", GoGetter: "LogoUri"},
			_jsii_.MemberProperty{JsiiProperty: "logoUriInput", GoGetter: "LogoUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "logoUrl", GoGetter: "LogoUrl"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "omitSecret", GoGetter: "OmitSecret"},
			_jsii_.MemberProperty{JsiiProperty: "omitSecretInput", GoGetter: "OmitSecretInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pkceRequired", GoGetter: "PkceRequired"},
			_jsii_.MemberProperty{JsiiProperty: "pkceRequiredInput", GoGetter: "PkceRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyUri", GoGetter: "PolicyUri"},
			_jsii_.MemberProperty{JsiiProperty: "policyUriInput", GoGetter: "PolicyUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "postLogoutRedirectUris", GoGetter: "PostLogoutRedirectUris"},
			_jsii_.MemberProperty{JsiiProperty: "postLogoutRedirectUrisInput", GoGetter: "PostLogoutRedirectUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "profile", GoGetter: "Profile"},
			_jsii_.MemberProperty{JsiiProperty: "profileInput", GoGetter: "ProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putGroupsClaim", GoMethod: "PutGroupsClaim"},
			_jsii_.MemberMethod{JsiiMethod: "putJwks", GoMethod: "PutJwks"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUris", GoGetter: "RedirectUris"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUrisInput", GoGetter: "RedirectUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenLeeway", GoGetter: "RefreshTokenLeeway"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenLeewayInput", GoGetter: "RefreshTokenLeewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenRotation", GoGetter: "RefreshTokenRotation"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenRotationInput", GoGetter: "RefreshTokenRotationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilityErrorRedirectUrl", GoMethod: "ResetAccessibilityErrorRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilityLoginRedirectUrl", GoMethod: "ResetAccessibilityLoginRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilitySelfService", GoMethod: "ResetAccessibilitySelfService"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminNote", GoMethod: "ResetAdminNote"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLinksJson", GoMethod: "ResetAppLinksJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppSettingsJson", GoMethod: "ResetAppSettingsJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthenticationPolicy", GoMethod: "ResetAuthenticationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoKeyRotation", GoMethod: "ResetAutoKeyRotation"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoSubmitToolbar", GoMethod: "ResetAutoSubmitToolbar"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientBasicSecret", GoMethod: "ResetClientBasicSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientUri", GoMethod: "ResetClientUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsentMethod", GoMethod: "ResetConsentMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnduserNote", GoMethod: "ResetEnduserNote"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrantTypes", GoMethod: "ResetGrantTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupsClaim", GoMethod: "ResetGroupsClaim"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideIos", GoMethod: "ResetHideIos"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideWeb", GoMethod: "ResetHideWeb"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetImplicitAssignment", GoMethod: "ResetImplicitAssignment"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuerMode", GoMethod: "ResetIssuerMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwks", GoMethod: "ResetJwks"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwksUri", GoMethod: "ResetJwksUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginMode", GoMethod: "ResetLoginMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginScopes", GoMethod: "ResetLoginScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginUri", GoMethod: "ResetLoginUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogo", GoMethod: "ResetLogo"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogoUri", GoMethod: "ResetLogoUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetOmitSecret", GoMethod: "ResetOmitSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPkceRequired", GoMethod: "ResetPkceRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyUri", GoMethod: "ResetPolicyUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostLogoutRedirectUris", GoMethod: "ResetPostLogoutRedirectUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetProfile", GoMethod: "ResetProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectUris", GoMethod: "ResetRedirectUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenLeeway", GoMethod: "ResetRefreshTokenLeeway"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenRotation", GoMethod: "ResetRefreshTokenRotation"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseTypes", GoMethod: "ResetResponseTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenEndpointAuthMethod", GoMethod: "ResetTokenEndpointAuthMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetTosUri", GoMethod: "ResetTosUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplate", GoMethod: "ResetUserNameTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplatePushStatus", GoMethod: "ResetUserNameTemplatePushStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplateSuffix", GoMethod: "ResetUserNameTemplateSuffix"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplateType", GoMethod: "ResetUserNameTemplateType"},
			_jsii_.MemberMethod{JsiiMethod: "resetWildcardRedirect", GoMethod: "ResetWildcardRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "responseTypes", GoGetter: "ResponseTypes"},
			_jsii_.MemberProperty{JsiiProperty: "responseTypesInput", GoGetter: "ResponseTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "signOnMode", GoGetter: "SignOnMode"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointAuthMethod", GoGetter: "TokenEndpointAuthMethod"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointAuthMethodInput", GoGetter: "TokenEndpointAuthMethodInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "tosUri", GoGetter: "TosUri"},
			_jsii_.MemberProperty{JsiiProperty: "tosUriInput", GoGetter: "TosUriInput"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplate", GoGetter: "UserNameTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateInput", GoGetter: "UserNameTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplatePushStatus", GoGetter: "UserNameTemplatePushStatus"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplatePushStatusInput", GoGetter: "UserNameTemplatePushStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateSuffix", GoGetter: "UserNameTemplateSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateSuffixInput", GoGetter: "UserNameTemplateSuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateType", GoGetter: "UserNameTemplateType"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateTypeInput", GoGetter: "UserNameTemplateTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "wildcardRedirect", GoGetter: "WildcardRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "wildcardRedirectInput", GoGetter: "WildcardRedirectInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppOauth{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.appOauth.AppOauthConfig",
		reflect.TypeOf((*AppOauthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.appOauth.AppOauthGroupsClaim",
		reflect.TypeOf((*AppOauthGroupsClaim)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.appOauth.AppOauthGroupsClaimOutputReference",
		reflect.TypeOf((*AppOauthGroupsClaimOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "filterType", GoGetter: "FilterType"},
			_jsii_.MemberProperty{JsiiProperty: "filterTypeInput", GoGetter: "FilterTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "issuerMode", GoGetter: "IssuerMode"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilterType", GoMethod: "ResetFilterType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppOauthGroupsClaimOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.appOauth.AppOauthJwks",
		reflect.TypeOf((*AppOauthJwks)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.appOauth.AppOauthJwksList",
		reflect.TypeOf((*AppOauthJwksList)(nil)).Elem(),
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
			j := jsiiProxy_AppOauthJwksList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.appOauth.AppOauthJwksOutputReference",
		reflect.TypeOf((*AppOauthJwksOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "e", GoGetter: "E"},
			_jsii_.MemberProperty{JsiiProperty: "eInput", GoGetter: "EInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kid", GoGetter: "Kid"},
			_jsii_.MemberProperty{JsiiProperty: "kidInput", GoGetter: "KidInput"},
			_jsii_.MemberProperty{JsiiProperty: "kty", GoGetter: "Kty"},
			_jsii_.MemberProperty{JsiiProperty: "ktyInput", GoGetter: "KtyInput"},
			_jsii_.MemberProperty{JsiiProperty: "n", GoGetter: "N"},
			_jsii_.MemberProperty{JsiiProperty: "nInput", GoGetter: "NInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetE", GoMethod: "ResetE"},
			_jsii_.MemberMethod{JsiiMethod: "resetN", GoMethod: "ResetN"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppOauthJwksOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.appOauth.AppOauthTimeouts",
		reflect.TypeOf((*AppOauthTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.appOauth.AppOauthTimeoutsOutputReference",
		reflect.TypeOf((*AppOauthTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_AppOauthTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
