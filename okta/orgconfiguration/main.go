// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orgconfiguration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.orgConfiguration.OrgConfiguration",
		reflect.TypeOf((*OrgConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "address1", GoGetter: "Address1"},
			_jsii_.MemberProperty{JsiiProperty: "address1Input", GoGetter: "Address1Input"},
			_jsii_.MemberProperty{JsiiProperty: "address2", GoGetter: "Address2"},
			_jsii_.MemberProperty{JsiiProperty: "address2Input", GoGetter: "Address2Input"},
			_jsii_.MemberProperty{JsiiProperty: "billingContactUser", GoGetter: "BillingContactUser"},
			_jsii_.MemberProperty{JsiiProperty: "billingContactUserInput", GoGetter: "BillingContactUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "city", GoGetter: "City"},
			_jsii_.MemberProperty{JsiiProperty: "cityInput", GoGetter: "CityInput"},
			_jsii_.MemberProperty{JsiiProperty: "companyName", GoGetter: "CompanyName"},
			_jsii_.MemberProperty{JsiiProperty: "companyNameInput", GoGetter: "CompanyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "country", GoGetter: "Country"},
			_jsii_.MemberProperty{JsiiProperty: "countryInput", GoGetter: "CountryInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "endUserSupportHelpUrl", GoGetter: "EndUserSupportHelpUrl"},
			_jsii_.MemberProperty{JsiiProperty: "endUserSupportHelpUrlInput", GoGetter: "EndUserSupportHelpUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "expiresAt", GoGetter: "ExpiresAt"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logo", GoGetter: "Logo"},
			_jsii_.MemberProperty{JsiiProperty: "logoInput", GoGetter: "LogoInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "optOutCommunicationEmails", GoGetter: "OptOutCommunicationEmails"},
			_jsii_.MemberProperty{JsiiProperty: "optOutCommunicationEmailsInput", GoGetter: "OptOutCommunicationEmailsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "phoneNumber", GoGetter: "PhoneNumber"},
			_jsii_.MemberProperty{JsiiProperty: "phoneNumberInput", GoGetter: "PhoneNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "postalCode", GoGetter: "PostalCode"},
			_jsii_.MemberProperty{JsiiProperty: "postalCodeInput", GoGetter: "PostalCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddress1", GoMethod: "ResetAddress1"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddress2", GoMethod: "ResetAddress2"},
			_jsii_.MemberMethod{JsiiMethod: "resetBillingContactUser", GoMethod: "ResetBillingContactUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetCity", GoMethod: "ResetCity"},
			_jsii_.MemberMethod{JsiiMethod: "resetCountry", GoMethod: "ResetCountry"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndUserSupportHelpUrl", GoMethod: "ResetEndUserSupportHelpUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogo", GoMethod: "ResetLogo"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptOutCommunicationEmails", GoMethod: "ResetOptOutCommunicationEmails"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhoneNumber", GoMethod: "ResetPhoneNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostalCode", GoMethod: "ResetPostalCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetState", GoMethod: "ResetState"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportPhoneNumber", GoMethod: "ResetSupportPhoneNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resetTechnicalContactUser", GoMethod: "ResetTechnicalContactUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebsite", GoMethod: "ResetWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "stateInput", GoGetter: "StateInput"},
			_jsii_.MemberProperty{JsiiProperty: "subdomain", GoGetter: "Subdomain"},
			_jsii_.MemberProperty{JsiiProperty: "supportPhoneNumber", GoGetter: "SupportPhoneNumber"},
			_jsii_.MemberProperty{JsiiProperty: "supportPhoneNumberInput", GoGetter: "SupportPhoneNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "technicalContactUser", GoGetter: "TechnicalContactUser"},
			_jsii_.MemberProperty{JsiiProperty: "technicalContactUserInput", GoGetter: "TechnicalContactUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "website", GoGetter: "Website"},
			_jsii_.MemberProperty{JsiiProperty: "websiteInput", GoGetter: "WebsiteInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrgConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.orgConfiguration.OrgConfigurationConfig",
		reflect.TypeOf((*OrgConfigurationConfig)(nil)).Elem(),
	)
}
