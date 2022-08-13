// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta


type AppOauthUsers struct {
	// User ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#id AppOauth#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Password for user application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#password AppOauth#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username for user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#username AppOauth#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

