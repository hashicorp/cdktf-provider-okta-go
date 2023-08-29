// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policydeviceassurancechromeos

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
		reflect.TypeOf((*PolicyDeviceAssuranceChromeos)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdBy", GoGetter: "CreatedBy"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdate", GoGetter: "LastUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedBy", GoGetter: "LastUpdatedBy"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspAllowScreenLock", GoMethod: "ResetTpspAllowScreenLock"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspBrowserVersion", GoMethod: "ResetTpspBrowserVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspBuiltinDnsClientEnabled", GoMethod: "ResetTpspBuiltinDnsClientEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspChromeRemoteDesktopAppBlocked", GoMethod: "ResetTpspChromeRemoteDesktopAppBlocked"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspDeviceEnrollmentDomain", GoMethod: "ResetTpspDeviceEnrollmentDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspDiskEncrypted", GoMethod: "ResetTpspDiskEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspKeyTrustLevel", GoMethod: "ResetTpspKeyTrustLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspOsFirewall", GoMethod: "ResetTpspOsFirewall"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspOsVersion", GoMethod: "ResetTpspOsVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspPasswordProctectionWarningTrigger", GoMethod: "ResetTpspPasswordProctectionWarningTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspRealtimeUrlCheckMode", GoMethod: "ResetTpspRealtimeUrlCheckMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspSafeBrowsingProtectionLevel", GoMethod: "ResetTpspSafeBrowsingProtectionLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspScreenLockSecured", GoMethod: "ResetTpspScreenLockSecured"},
			_jsii_.MemberMethod{JsiiMethod: "resetTpspSiteIsolationEnabled", GoMethod: "ResetTpspSiteIsolationEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tpspAllowScreenLock", GoGetter: "TpspAllowScreenLock"},
			_jsii_.MemberProperty{JsiiProperty: "tpspAllowScreenLockInput", GoGetter: "TpspAllowScreenLockInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspBrowserVersion", GoGetter: "TpspBrowserVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tpspBrowserVersionInput", GoGetter: "TpspBrowserVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspBuiltinDnsClientEnabled", GoGetter: "TpspBuiltinDnsClientEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "tpspBuiltinDnsClientEnabledInput", GoGetter: "TpspBuiltinDnsClientEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspChromeRemoteDesktopAppBlocked", GoGetter: "TpspChromeRemoteDesktopAppBlocked"},
			_jsii_.MemberProperty{JsiiProperty: "tpspChromeRemoteDesktopAppBlockedInput", GoGetter: "TpspChromeRemoteDesktopAppBlockedInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspDeviceEnrollmentDomain", GoGetter: "TpspDeviceEnrollmentDomain"},
			_jsii_.MemberProperty{JsiiProperty: "tpspDeviceEnrollmentDomainInput", GoGetter: "TpspDeviceEnrollmentDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspDiskEncrypted", GoGetter: "TpspDiskEncrypted"},
			_jsii_.MemberProperty{JsiiProperty: "tpspDiskEncryptedInput", GoGetter: "TpspDiskEncryptedInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspKeyTrustLevel", GoGetter: "TpspKeyTrustLevel"},
			_jsii_.MemberProperty{JsiiProperty: "tpspKeyTrustLevelInput", GoGetter: "TpspKeyTrustLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspOsFirewall", GoGetter: "TpspOsFirewall"},
			_jsii_.MemberProperty{JsiiProperty: "tpspOsFirewallInput", GoGetter: "TpspOsFirewallInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspOsVersion", GoGetter: "TpspOsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tpspOsVersionInput", GoGetter: "TpspOsVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspPasswordProctectionWarningTrigger", GoGetter: "TpspPasswordProctectionWarningTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "tpspPasswordProctectionWarningTriggerInput", GoGetter: "TpspPasswordProctectionWarningTriggerInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspRealtimeUrlCheckMode", GoGetter: "TpspRealtimeUrlCheckMode"},
			_jsii_.MemberProperty{JsiiProperty: "tpspRealtimeUrlCheckModeInput", GoGetter: "TpspRealtimeUrlCheckModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspSafeBrowsingProtectionLevel", GoGetter: "TpspSafeBrowsingProtectionLevel"},
			_jsii_.MemberProperty{JsiiProperty: "tpspSafeBrowsingProtectionLevelInput", GoGetter: "TpspSafeBrowsingProtectionLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspScreenLockSecured", GoGetter: "TpspScreenLockSecured"},
			_jsii_.MemberProperty{JsiiProperty: "tpspScreenLockSecuredInput", GoGetter: "TpspScreenLockSecuredInput"},
			_jsii_.MemberProperty{JsiiProperty: "tpspSiteIsolationEnabled", GoGetter: "TpspSiteIsolationEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "tpspSiteIsolationEnabledInput", GoGetter: "TpspSiteIsolationEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyDeviceAssuranceChromeos{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeosConfig",
		reflect.TypeOf((*PolicyDeviceAssuranceChromeosConfig)(nil)).Elem(),
	)
}
