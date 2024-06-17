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

type SourceClickupApi struct {
	pulumi.CustomResourceState

	Configuration SourceClickupApiConfigurationOutput `pulumi:"configuration"`
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

// NewSourceClickupApi registers a new resource with the given unique name, arguments, and options.
func NewSourceClickupApi(ctx *pulumi.Context,
	name string, args *SourceClickupApiArgs, opts ...pulumi.ResourceOption) (*SourceClickupApi, error) {
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
	var resource SourceClickupApi
	err := ctx.RegisterResource("airbyte:index/sourceClickupApi:SourceClickupApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceClickupApi gets an existing SourceClickupApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceClickupApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceClickupApiState, opts ...pulumi.ResourceOption) (*SourceClickupApi, error) {
	var resource SourceClickupApi
	err := ctx.ReadResource("airbyte:index/sourceClickupApi:SourceClickupApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceClickupApi resources.
type sourceClickupApiState struct {
	Configuration *SourceClickupApiConfiguration `pulumi:"configuration"`
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

type SourceClickupApiState struct {
	Configuration SourceClickupApiConfigurationPtrInput
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

func (SourceClickupApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceClickupApiState)(nil)).Elem()
}

type sourceClickupApiArgs struct {
	Configuration SourceClickupApiConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceClickupApi resource.
type SourceClickupApiArgs struct {
	Configuration SourceClickupApiConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceClickupApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceClickupApiArgs)(nil)).Elem()
}

type SourceClickupApiInput interface {
	pulumi.Input

	ToSourceClickupApiOutput() SourceClickupApiOutput
	ToSourceClickupApiOutputWithContext(ctx context.Context) SourceClickupApiOutput
}

func (*SourceClickupApi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceClickupApi)(nil)).Elem()
}

func (i *SourceClickupApi) ToSourceClickupApiOutput() SourceClickupApiOutput {
	return i.ToSourceClickupApiOutputWithContext(context.Background())
}

func (i *SourceClickupApi) ToSourceClickupApiOutputWithContext(ctx context.Context) SourceClickupApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceClickupApiOutput)
}

// SourceClickupApiArrayInput is an input type that accepts SourceClickupApiArray and SourceClickupApiArrayOutput values.
// You can construct a concrete instance of `SourceClickupApiArrayInput` via:
//
//	SourceClickupApiArray{ SourceClickupApiArgs{...} }
type SourceClickupApiArrayInput interface {
	pulumi.Input

	ToSourceClickupApiArrayOutput() SourceClickupApiArrayOutput
	ToSourceClickupApiArrayOutputWithContext(context.Context) SourceClickupApiArrayOutput
}

type SourceClickupApiArray []SourceClickupApiInput

func (SourceClickupApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceClickupApi)(nil)).Elem()
}

func (i SourceClickupApiArray) ToSourceClickupApiArrayOutput() SourceClickupApiArrayOutput {
	return i.ToSourceClickupApiArrayOutputWithContext(context.Background())
}

func (i SourceClickupApiArray) ToSourceClickupApiArrayOutputWithContext(ctx context.Context) SourceClickupApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceClickupApiArrayOutput)
}

// SourceClickupApiMapInput is an input type that accepts SourceClickupApiMap and SourceClickupApiMapOutput values.
// You can construct a concrete instance of `SourceClickupApiMapInput` via:
//
//	SourceClickupApiMap{ "key": SourceClickupApiArgs{...} }
type SourceClickupApiMapInput interface {
	pulumi.Input

	ToSourceClickupApiMapOutput() SourceClickupApiMapOutput
	ToSourceClickupApiMapOutputWithContext(context.Context) SourceClickupApiMapOutput
}

type SourceClickupApiMap map[string]SourceClickupApiInput

func (SourceClickupApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceClickupApi)(nil)).Elem()
}

func (i SourceClickupApiMap) ToSourceClickupApiMapOutput() SourceClickupApiMapOutput {
	return i.ToSourceClickupApiMapOutputWithContext(context.Background())
}

func (i SourceClickupApiMap) ToSourceClickupApiMapOutputWithContext(ctx context.Context) SourceClickupApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceClickupApiMapOutput)
}

type SourceClickupApiOutput struct{ *pulumi.OutputState }

func (SourceClickupApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceClickupApi)(nil)).Elem()
}

func (o SourceClickupApiOutput) ToSourceClickupApiOutput() SourceClickupApiOutput {
	return o
}

func (o SourceClickupApiOutput) ToSourceClickupApiOutputWithContext(ctx context.Context) SourceClickupApiOutput {
	return o
}

func (o SourceClickupApiOutput) Configuration() SourceClickupApiConfigurationOutput {
	return o.ApplyT(func(v *SourceClickupApi) SourceClickupApiConfigurationOutput { return v.Configuration }).(SourceClickupApiConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceClickupApiOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceClickupApi) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceClickupApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickupApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceClickupApiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceClickupApi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceClickupApiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickupApi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceClickupApiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickupApi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceClickupApiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickupApi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceClickupApiArrayOutput struct{ *pulumi.OutputState }

func (SourceClickupApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceClickupApi)(nil)).Elem()
}

func (o SourceClickupApiArrayOutput) ToSourceClickupApiArrayOutput() SourceClickupApiArrayOutput {
	return o
}

func (o SourceClickupApiArrayOutput) ToSourceClickupApiArrayOutputWithContext(ctx context.Context) SourceClickupApiArrayOutput {
	return o
}

func (o SourceClickupApiArrayOutput) Index(i pulumi.IntInput) SourceClickupApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceClickupApi {
		return vs[0].([]*SourceClickupApi)[vs[1].(int)]
	}).(SourceClickupApiOutput)
}

type SourceClickupApiMapOutput struct{ *pulumi.OutputState }

func (SourceClickupApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceClickupApi)(nil)).Elem()
}

func (o SourceClickupApiMapOutput) ToSourceClickupApiMapOutput() SourceClickupApiMapOutput {
	return o
}

func (o SourceClickupApiMapOutput) ToSourceClickupApiMapOutputWithContext(ctx context.Context) SourceClickupApiMapOutput {
	return o
}

func (o SourceClickupApiMapOutput) MapIndex(k pulumi.StringInput) SourceClickupApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceClickupApi {
		return vs[0].(map[string]*SourceClickupApi)[vs[1].(string)]
	}).(SourceClickupApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceClickupApiInput)(nil)).Elem(), &SourceClickupApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceClickupApiArrayInput)(nil)).Elem(), SourceClickupApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceClickupApiMapInput)(nil)).Elem(), SourceClickupApiMap{})
	pulumi.RegisterOutputType(SourceClickupApiOutput{})
	pulumi.RegisterOutputType(SourceClickupApiArrayOutput{})
	pulumi.RegisterOutputType(SourceClickupApiMapOutput{})
}