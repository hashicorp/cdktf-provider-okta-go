//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package samlapp

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlAppKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SamlAppKeysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SamlAppKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SamlAppKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SamlAppKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSamlAppKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

