//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataOktaAppSamlAttributeStatementsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataOktaAppSamlAttributeStatementsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataOktaAppSamlAttributeStatementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataOktaAppSamlAttributeStatementsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataOktaAppSamlAttributeStatementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataOktaAppSamlAttributeStatementsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

