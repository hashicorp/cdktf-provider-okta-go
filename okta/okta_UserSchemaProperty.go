// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-okta-go/okta/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-okta-go/okta/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/user_schema_property okta_user_schema_property}.
type UserSchemaProperty interface {
	cdktf.TerraformResource
	ArrayEnum() *[]*string
	SetArrayEnum(val *[]*string)
	ArrayEnumInput() *[]*string
	ArrayOneOf() UserSchemaPropertyArrayOneOfList
	ArrayOneOfInput() interface{}
	ArrayType() *string
	SetArrayType(val *string)
	ArrayTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enum() *[]*string
	SetEnum(val *[]*string)
	EnumInput() *[]*string
	ExternalName() *string
	SetExternalName(val *string)
	ExternalNameInput() *string
	ExternalNamespace() *string
	SetExternalNamespace(val *string)
	ExternalNamespaceInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Master() *string
	SetMaster(val *string)
	MasterInput() *string
	MasterOverridePriority() UserSchemaPropertyMasterOverridePriorityList
	MasterOverridePriorityInput() interface{}
	MaxLength() *float64
	SetMaxLength(val *float64)
	MaxLengthInput() *float64
	MinLength() *float64
	SetMinLength(val *float64)
	MinLengthInput() *float64
	// The tree node.
	Node() constructs.Node
	OneOf() UserSchemaPropertyOneOfList
	OneOfInput() interface{}
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
	Permissions() *string
	SetPermissions(val *string)
	PermissionsInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Required() interface{}
	SetRequired(val interface{})
	RequiredInput() interface{}
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Unique() *string
	SetUnique(val *string)
	UniqueInput() *string
	UserType() *string
	SetUserType(val *string)
	UserTypeInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutArrayOneOf(value interface{})
	PutMasterOverridePriority(value interface{})
	PutOneOf(value interface{})
	ResetArrayEnum()
	ResetArrayOneOf()
	ResetArrayType()
	ResetDescription()
	ResetEnum()
	ResetExternalName()
	ResetExternalNamespace()
	ResetId()
	ResetMaster()
	ResetMasterOverridePriority()
	ResetMaxLength()
	ResetMinLength()
	ResetOneOf()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPattern()
	ResetPermissions()
	ResetRequired()
	ResetScope()
	ResetUnique()
	ResetUserType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for UserSchemaProperty
type jsiiProxy_UserSchemaProperty struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_UserSchemaProperty) ArrayEnum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ArrayEnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ArrayOneOf() UserSchemaPropertyArrayOneOfList {
	var returns UserSchemaPropertyArrayOneOfList
	_jsii_.Get(
		j,
		"arrayOneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ArrayOneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arrayOneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ArrayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ArrayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) EnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ExternalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ExternalNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ExternalNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Master() *string {
	var returns *string
	_jsii_.Get(
		j,
		"master",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) MasterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) MasterOverridePriority() UserSchemaPropertyMasterOverridePriorityList {
	var returns UserSchemaPropertyMasterOverridePriorityList
	_jsii_.Get(
		j,
		"masterOverridePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) MasterOverridePriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterOverridePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) MaxLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) MinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) OneOf() UserSchemaPropertyOneOfList {
	var returns UserSchemaPropertyOneOfList
	_jsii_.Get(
		j,
		"oneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) OneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Permissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) PermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) Unique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) UniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchemaProperty) UserTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/user_schema_property okta_user_schema_property} Resource.
func NewUserSchemaProperty(scope constructs.Construct, id *string, config *UserSchemaPropertyConfig) UserSchemaProperty {
	_init_.Initialize()

	j := jsiiProxy_UserSchemaProperty{}

	_jsii_.Create(
		"@cdktf/provider-okta.UserSchemaProperty",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/user_schema_property okta_user_schema_property} Resource.
func NewUserSchemaProperty_Override(u UserSchemaProperty, scope constructs.Construct, id *string, config *UserSchemaPropertyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.UserSchemaProperty",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetArrayEnum(val *[]*string) {
	_jsii_.Set(
		j,
		"arrayEnum",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetArrayType(val *string) {
	_jsii_.Set(
		j,
		"arrayType",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetEnum(val *[]*string) {
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetExternalName(val *string) {
	_jsii_.Set(
		j,
		"externalName",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetExternalNamespace(val *string) {
	_jsii_.Set(
		j,
		"externalNamespace",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetIndex(val *string) {
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetMaster(val *string) {
	_jsii_.Set(
		j,
		"master",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetMaxLength(val *float64) {
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetMinLength(val *float64) {
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetPattern(val *string) {
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetPermissions(val *string) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetRequired(val interface{}) {
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetTitle(val *string) {
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetUnique(val *string) {
	_jsii_.Set(
		j,
		"unique",
		val,
	)
}

func (j *jsiiProxy_UserSchemaProperty) SetUserType(val *string) {
	_jsii_.Set(
		j,
		"userType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func UserSchemaProperty_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.UserSchemaProperty",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func UserSchemaProperty_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.UserSchemaProperty",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UserSchemaProperty) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_UserSchemaProperty) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_UserSchemaProperty) PutArrayOneOf(value interface{}) {
	_jsii_.InvokeVoid(
		u,
		"putArrayOneOf",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UserSchemaProperty) PutMasterOverridePriority(value interface{}) {
	_jsii_.InvokeVoid(
		u,
		"putMasterOverridePriority",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UserSchemaProperty) PutOneOf(value interface{}) {
	_jsii_.InvokeVoid(
		u,
		"putOneOf",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetArrayEnum() {
	_jsii_.InvokeVoid(
		u,
		"resetArrayEnum",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetArrayOneOf() {
	_jsii_.InvokeVoid(
		u,
		"resetArrayOneOf",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetArrayType() {
	_jsii_.InvokeVoid(
		u,
		"resetArrayType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetDescription() {
	_jsii_.InvokeVoid(
		u,
		"resetDescription",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetEnum() {
	_jsii_.InvokeVoid(
		u,
		"resetEnum",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetExternalName() {
	_jsii_.InvokeVoid(
		u,
		"resetExternalName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetExternalNamespace() {
	_jsii_.InvokeVoid(
		u,
		"resetExternalNamespace",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetMaster() {
	_jsii_.InvokeVoid(
		u,
		"resetMaster",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetMasterOverridePriority() {
	_jsii_.InvokeVoid(
		u,
		"resetMasterOverridePriority",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetMaxLength() {
	_jsii_.InvokeVoid(
		u,
		"resetMaxLength",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetMinLength() {
	_jsii_.InvokeVoid(
		u,
		"resetMinLength",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetOneOf() {
	_jsii_.InvokeVoid(
		u,
		"resetOneOf",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetPattern() {
	_jsii_.InvokeVoid(
		u,
		"resetPattern",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetPermissions() {
	_jsii_.InvokeVoid(
		u,
		"resetPermissions",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetRequired() {
	_jsii_.InvokeVoid(
		u,
		"resetRequired",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetScope() {
	_jsii_.InvokeVoid(
		u,
		"resetScope",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetUnique() {
	_jsii_.InvokeVoid(
		u,
		"resetUnique",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) ResetUserType() {
	_jsii_.InvokeVoid(
		u,
		"resetUserType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchemaProperty) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchemaProperty) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

