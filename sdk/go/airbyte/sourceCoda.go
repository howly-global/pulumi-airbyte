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

type SourceCoda struct {
	pulumi.CustomResourceState

	Configuration SourceCodaConfigurationOutput `pulumi:"configuration"`
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

// NewSourceCoda registers a new resource with the given unique name, arguments, and options.
func NewSourceCoda(ctx *pulumi.Context,
	name string, args *SourceCodaArgs, opts ...pulumi.ResourceOption) (*SourceCoda, error) {
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
	var resource SourceCoda
	err := ctx.RegisterResource("airbyte:index/sourceCoda:SourceCoda", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCoda gets an existing SourceCoda resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCoda(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCodaState, opts ...pulumi.ResourceOption) (*SourceCoda, error) {
	var resource SourceCoda
	err := ctx.ReadResource("airbyte:index/sourceCoda:SourceCoda", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCoda resources.
type sourceCodaState struct {
	Configuration *SourceCodaConfiguration `pulumi:"configuration"`
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

type SourceCodaState struct {
	Configuration SourceCodaConfigurationPtrInput
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

func (SourceCodaState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCodaState)(nil)).Elem()
}

type sourceCodaArgs struct {
	Configuration SourceCodaConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceCoda resource.
type SourceCodaArgs struct {
	Configuration SourceCodaConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceCodaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCodaArgs)(nil)).Elem()
}

type SourceCodaInput interface {
	pulumi.Input

	ToSourceCodaOutput() SourceCodaOutput
	ToSourceCodaOutputWithContext(ctx context.Context) SourceCodaOutput
}

func (*SourceCoda) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCoda)(nil)).Elem()
}

func (i *SourceCoda) ToSourceCodaOutput() SourceCodaOutput {
	return i.ToSourceCodaOutputWithContext(context.Background())
}

func (i *SourceCoda) ToSourceCodaOutputWithContext(ctx context.Context) SourceCodaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodaOutput)
}

// SourceCodaArrayInput is an input type that accepts SourceCodaArray and SourceCodaArrayOutput values.
// You can construct a concrete instance of `SourceCodaArrayInput` via:
//
//	SourceCodaArray{ SourceCodaArgs{...} }
type SourceCodaArrayInput interface {
	pulumi.Input

	ToSourceCodaArrayOutput() SourceCodaArrayOutput
	ToSourceCodaArrayOutputWithContext(context.Context) SourceCodaArrayOutput
}

type SourceCodaArray []SourceCodaInput

func (SourceCodaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCoda)(nil)).Elem()
}

func (i SourceCodaArray) ToSourceCodaArrayOutput() SourceCodaArrayOutput {
	return i.ToSourceCodaArrayOutputWithContext(context.Background())
}

func (i SourceCodaArray) ToSourceCodaArrayOutputWithContext(ctx context.Context) SourceCodaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodaArrayOutput)
}

// SourceCodaMapInput is an input type that accepts SourceCodaMap and SourceCodaMapOutput values.
// You can construct a concrete instance of `SourceCodaMapInput` via:
//
//	SourceCodaMap{ "key": SourceCodaArgs{...} }
type SourceCodaMapInput interface {
	pulumi.Input

	ToSourceCodaMapOutput() SourceCodaMapOutput
	ToSourceCodaMapOutputWithContext(context.Context) SourceCodaMapOutput
}

type SourceCodaMap map[string]SourceCodaInput

func (SourceCodaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCoda)(nil)).Elem()
}

func (i SourceCodaMap) ToSourceCodaMapOutput() SourceCodaMapOutput {
	return i.ToSourceCodaMapOutputWithContext(context.Background())
}

func (i SourceCodaMap) ToSourceCodaMapOutputWithContext(ctx context.Context) SourceCodaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCodaMapOutput)
}

type SourceCodaOutput struct{ *pulumi.OutputState }

func (SourceCodaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCoda)(nil)).Elem()
}

func (o SourceCodaOutput) ToSourceCodaOutput() SourceCodaOutput {
	return o
}

func (o SourceCodaOutput) ToSourceCodaOutputWithContext(ctx context.Context) SourceCodaOutput {
	return o
}

func (o SourceCodaOutput) Configuration() SourceCodaConfigurationOutput {
	return o.ApplyT(func(v *SourceCoda) SourceCodaConfigurationOutput { return v.Configuration }).(SourceCodaConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceCodaOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCoda) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceCodaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoda) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceCodaOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCoda) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceCodaOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoda) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceCodaOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoda) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceCodaOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoda) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceCodaArrayOutput struct{ *pulumi.OutputState }

func (SourceCodaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCoda)(nil)).Elem()
}

func (o SourceCodaArrayOutput) ToSourceCodaArrayOutput() SourceCodaArrayOutput {
	return o
}

func (o SourceCodaArrayOutput) ToSourceCodaArrayOutputWithContext(ctx context.Context) SourceCodaArrayOutput {
	return o
}

func (o SourceCodaArrayOutput) Index(i pulumi.IntInput) SourceCodaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceCoda {
		return vs[0].([]*SourceCoda)[vs[1].(int)]
	}).(SourceCodaOutput)
}

type SourceCodaMapOutput struct{ *pulumi.OutputState }

func (SourceCodaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCoda)(nil)).Elem()
}

func (o SourceCodaMapOutput) ToSourceCodaMapOutput() SourceCodaMapOutput {
	return o
}

func (o SourceCodaMapOutput) ToSourceCodaMapOutputWithContext(ctx context.Context) SourceCodaMapOutput {
	return o
}

func (o SourceCodaMapOutput) MapIndex(k pulumi.StringInput) SourceCodaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceCoda {
		return vs[0].(map[string]*SourceCoda)[vs[1].(string)]
	}).(SourceCodaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCodaInput)(nil)).Elem(), &SourceCoda{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCodaArrayInput)(nil)).Elem(), SourceCodaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCodaMapInput)(nil)).Elem(), SourceCodaMap{})
	pulumi.RegisterOutputType(SourceCodaOutput{})
	pulumi.RegisterOutputType(SourceCodaArrayOutput{})
	pulumi.RegisterOutputType(SourceCodaMapOutput{})
}