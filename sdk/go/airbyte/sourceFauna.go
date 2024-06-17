// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceFauna struct {
	pulumi.CustomResourceState

	Configuration SourceFaunaConfigurationOutput `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrOutput `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceFauna registers a new resource with the given unique name, arguments, and options.
func NewSourceFauna(ctx *pulumi.Context,
	name string, args *SourceFaunaArgs, opts ...pulumi.ResourceOption) (*SourceFauna, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SourceFauna
	err := ctx.RegisterResource("airbyte:index/sourceFauna:SourceFauna", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceFauna gets an existing SourceFauna resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceFauna(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceFaunaState, opts ...pulumi.ResourceOption) (*SourceFauna, error) {
	var resource SourceFauna
	err := ctx.ReadResource("airbyte:index/sourceFauna:SourceFauna", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceFauna resources.
type sourceFaunaState struct {
	Configuration *SourceFaunaConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceFaunaState struct {
	Configuration SourceFaunaConfigurationPtrInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceFaunaState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFaunaState)(nil)).Elem()
}

type sourceFaunaArgs struct {
	Configuration SourceFaunaConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceFauna resource.
type SourceFaunaArgs struct {
	Configuration SourceFaunaConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceFaunaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFaunaArgs)(nil)).Elem()
}

type SourceFaunaInput interface {
	pulumi.Input

	ToSourceFaunaOutput() SourceFaunaOutput
	ToSourceFaunaOutputWithContext(ctx context.Context) SourceFaunaOutput
}

func (*SourceFauna) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFauna)(nil)).Elem()
}

func (i *SourceFauna) ToSourceFaunaOutput() SourceFaunaOutput {
	return i.ToSourceFaunaOutputWithContext(context.Background())
}

func (i *SourceFauna) ToSourceFaunaOutputWithContext(ctx context.Context) SourceFaunaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFaunaOutput)
}

// SourceFaunaArrayInput is an input type that accepts SourceFaunaArray and SourceFaunaArrayOutput values.
// You can construct a concrete instance of `SourceFaunaArrayInput` via:
//
//	SourceFaunaArray{ SourceFaunaArgs{...} }
type SourceFaunaArrayInput interface {
	pulumi.Input

	ToSourceFaunaArrayOutput() SourceFaunaArrayOutput
	ToSourceFaunaArrayOutputWithContext(context.Context) SourceFaunaArrayOutput
}

type SourceFaunaArray []SourceFaunaInput

func (SourceFaunaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceFauna)(nil)).Elem()
}

func (i SourceFaunaArray) ToSourceFaunaArrayOutput() SourceFaunaArrayOutput {
	return i.ToSourceFaunaArrayOutputWithContext(context.Background())
}

func (i SourceFaunaArray) ToSourceFaunaArrayOutputWithContext(ctx context.Context) SourceFaunaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFaunaArrayOutput)
}

// SourceFaunaMapInput is an input type that accepts SourceFaunaMap and SourceFaunaMapOutput values.
// You can construct a concrete instance of `SourceFaunaMapInput` via:
//
//	SourceFaunaMap{ "key": SourceFaunaArgs{...} }
type SourceFaunaMapInput interface {
	pulumi.Input

	ToSourceFaunaMapOutput() SourceFaunaMapOutput
	ToSourceFaunaMapOutputWithContext(context.Context) SourceFaunaMapOutput
}

type SourceFaunaMap map[string]SourceFaunaInput

func (SourceFaunaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceFauna)(nil)).Elem()
}

func (i SourceFaunaMap) ToSourceFaunaMapOutput() SourceFaunaMapOutput {
	return i.ToSourceFaunaMapOutputWithContext(context.Background())
}

func (i SourceFaunaMap) ToSourceFaunaMapOutputWithContext(ctx context.Context) SourceFaunaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFaunaMapOutput)
}

type SourceFaunaOutput struct{ *pulumi.OutputState }

func (SourceFaunaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFauna)(nil)).Elem()
}

func (o SourceFaunaOutput) ToSourceFaunaOutput() SourceFaunaOutput {
	return o
}

func (o SourceFaunaOutput) ToSourceFaunaOutputWithContext(ctx context.Context) SourceFaunaOutput {
	return o
}

func (o SourceFaunaOutput) Configuration() SourceFaunaConfigurationOutput {
	return o.ApplyT(func(v *SourceFauna) SourceFaunaConfigurationOutput { return v.Configuration }).(SourceFaunaConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceFaunaOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFauna) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceFaunaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFauna) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceFaunaOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFauna) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceFaunaOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFauna) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceFaunaOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFauna) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceFaunaOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFauna) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceFaunaArrayOutput struct{ *pulumi.OutputState }

func (SourceFaunaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceFauna)(nil)).Elem()
}

func (o SourceFaunaArrayOutput) ToSourceFaunaArrayOutput() SourceFaunaArrayOutput {
	return o
}

func (o SourceFaunaArrayOutput) ToSourceFaunaArrayOutputWithContext(ctx context.Context) SourceFaunaArrayOutput {
	return o
}

func (o SourceFaunaArrayOutput) Index(i pulumi.IntInput) SourceFaunaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceFauna {
		return vs[0].([]*SourceFauna)[vs[1].(int)]
	}).(SourceFaunaOutput)
}

type SourceFaunaMapOutput struct{ *pulumi.OutputState }

func (SourceFaunaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceFauna)(nil)).Elem()
}

func (o SourceFaunaMapOutput) ToSourceFaunaMapOutput() SourceFaunaMapOutput {
	return o
}

func (o SourceFaunaMapOutput) ToSourceFaunaMapOutputWithContext(ctx context.Context) SourceFaunaMapOutput {
	return o
}

func (o SourceFaunaMapOutput) MapIndex(k pulumi.StringInput) SourceFaunaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceFauna {
		return vs[0].(map[string]*SourceFauna)[vs[1].(string)]
	}).(SourceFaunaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFaunaInput)(nil)).Elem(), &SourceFauna{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFaunaArrayInput)(nil)).Elem(), SourceFaunaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFaunaMapInput)(nil)).Elem(), SourceFaunaMap{})
	pulumi.RegisterOutputType(SourceFaunaOutput{})
	pulumi.RegisterOutputType(SourceFaunaArrayOutput{})
	pulumi.RegisterOutputType(SourceFaunaMapOutput{})
}