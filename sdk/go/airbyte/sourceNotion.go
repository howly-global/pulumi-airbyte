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

type SourceNotion struct {
	pulumi.CustomResourceState

	Configuration SourceNotionConfigurationOutput `pulumi:"configuration"`
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

// NewSourceNotion registers a new resource with the given unique name, arguments, and options.
func NewSourceNotion(ctx *pulumi.Context,
	name string, args *SourceNotionArgs, opts ...pulumi.ResourceOption) (*SourceNotion, error) {
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
	var resource SourceNotion
	err := ctx.RegisterResource("airbyte:index/sourceNotion:SourceNotion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceNotion gets an existing SourceNotion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceNotion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceNotionState, opts ...pulumi.ResourceOption) (*SourceNotion, error) {
	var resource SourceNotion
	err := ctx.ReadResource("airbyte:index/sourceNotion:SourceNotion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceNotion resources.
type sourceNotionState struct {
	Configuration *SourceNotionConfiguration `pulumi:"configuration"`
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

type SourceNotionState struct {
	Configuration SourceNotionConfigurationPtrInput
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

func (SourceNotionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNotionState)(nil)).Elem()
}

type sourceNotionArgs struct {
	Configuration SourceNotionConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceNotion resource.
type SourceNotionArgs struct {
	Configuration SourceNotionConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceNotionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNotionArgs)(nil)).Elem()
}

type SourceNotionInput interface {
	pulumi.Input

	ToSourceNotionOutput() SourceNotionOutput
	ToSourceNotionOutputWithContext(ctx context.Context) SourceNotionOutput
}

func (*SourceNotion) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNotion)(nil)).Elem()
}

func (i *SourceNotion) ToSourceNotionOutput() SourceNotionOutput {
	return i.ToSourceNotionOutputWithContext(context.Background())
}

func (i *SourceNotion) ToSourceNotionOutputWithContext(ctx context.Context) SourceNotionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNotionOutput)
}

// SourceNotionArrayInput is an input type that accepts SourceNotionArray and SourceNotionArrayOutput values.
// You can construct a concrete instance of `SourceNotionArrayInput` via:
//
//	SourceNotionArray{ SourceNotionArgs{...} }
type SourceNotionArrayInput interface {
	pulumi.Input

	ToSourceNotionArrayOutput() SourceNotionArrayOutput
	ToSourceNotionArrayOutputWithContext(context.Context) SourceNotionArrayOutput
}

type SourceNotionArray []SourceNotionInput

func (SourceNotionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceNotion)(nil)).Elem()
}

func (i SourceNotionArray) ToSourceNotionArrayOutput() SourceNotionArrayOutput {
	return i.ToSourceNotionArrayOutputWithContext(context.Background())
}

func (i SourceNotionArray) ToSourceNotionArrayOutputWithContext(ctx context.Context) SourceNotionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNotionArrayOutput)
}

// SourceNotionMapInput is an input type that accepts SourceNotionMap and SourceNotionMapOutput values.
// You can construct a concrete instance of `SourceNotionMapInput` via:
//
//	SourceNotionMap{ "key": SourceNotionArgs{...} }
type SourceNotionMapInput interface {
	pulumi.Input

	ToSourceNotionMapOutput() SourceNotionMapOutput
	ToSourceNotionMapOutputWithContext(context.Context) SourceNotionMapOutput
}

type SourceNotionMap map[string]SourceNotionInput

func (SourceNotionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceNotion)(nil)).Elem()
}

func (i SourceNotionMap) ToSourceNotionMapOutput() SourceNotionMapOutput {
	return i.ToSourceNotionMapOutputWithContext(context.Background())
}

func (i SourceNotionMap) ToSourceNotionMapOutputWithContext(ctx context.Context) SourceNotionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNotionMapOutput)
}

type SourceNotionOutput struct{ *pulumi.OutputState }

func (SourceNotionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNotion)(nil)).Elem()
}

func (o SourceNotionOutput) ToSourceNotionOutput() SourceNotionOutput {
	return o
}

func (o SourceNotionOutput) ToSourceNotionOutputWithContext(ctx context.Context) SourceNotionOutput {
	return o
}

func (o SourceNotionOutput) Configuration() SourceNotionConfigurationOutput {
	return o.ApplyT(func(v *SourceNotion) SourceNotionConfigurationOutput { return v.Configuration }).(SourceNotionConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceNotionOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceNotion) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceNotionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNotion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceNotionOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceNotion) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceNotionOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNotion) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceNotionOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNotion) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceNotionOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNotion) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceNotionArrayOutput struct{ *pulumi.OutputState }

func (SourceNotionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceNotion)(nil)).Elem()
}

func (o SourceNotionArrayOutput) ToSourceNotionArrayOutput() SourceNotionArrayOutput {
	return o
}

func (o SourceNotionArrayOutput) ToSourceNotionArrayOutputWithContext(ctx context.Context) SourceNotionArrayOutput {
	return o
}

func (o SourceNotionArrayOutput) Index(i pulumi.IntInput) SourceNotionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceNotion {
		return vs[0].([]*SourceNotion)[vs[1].(int)]
	}).(SourceNotionOutput)
}

type SourceNotionMapOutput struct{ *pulumi.OutputState }

func (SourceNotionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceNotion)(nil)).Elem()
}

func (o SourceNotionMapOutput) ToSourceNotionMapOutput() SourceNotionMapOutput {
	return o
}

func (o SourceNotionMapOutput) ToSourceNotionMapOutputWithContext(ctx context.Context) SourceNotionMapOutput {
	return o
}

func (o SourceNotionMapOutput) MapIndex(k pulumi.StringInput) SourceNotionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceNotion {
		return vs[0].(map[string]*SourceNotion)[vs[1].(string)]
	}).(SourceNotionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNotionInput)(nil)).Elem(), &SourceNotion{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNotionArrayInput)(nil)).Elem(), SourceNotionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNotionMapInput)(nil)).Elem(), SourceNotionMap{})
	pulumi.RegisterOutputType(SourceNotionOutput{})
	pulumi.RegisterOutputType(SourceNotionArrayOutput{})
	pulumi.RegisterOutputType(SourceNotionMapOutput{})
}