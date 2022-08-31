//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppBookmarkUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppBookmarkUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppBookmarkUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppBookmarkUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppBookmarkUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppBookmarkUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppBookmarkUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

