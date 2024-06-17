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

type SourceSmaily struct {
	pulumi.CustomResourceState

	Configuration SourceSmailyConfigurationOutput `pulumi:"configuration"`
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

// NewSourceSmaily registers a new resource with the given unique name, arguments, and options.
func NewSourceSmaily(ctx *pulumi.Context,
	name string, args *SourceSmailyArgs, opts ...pulumi.ResourceOption) (*SourceSmaily, error) {
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
	var resource SourceSmaily
	err := ctx.RegisterResource("airbyte:index/sourceSmaily:SourceSmaily", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSmaily gets an existing SourceSmaily resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSmaily(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSmailyState, opts ...pulumi.ResourceOption) (*SourceSmaily, error) {
	var resource SourceSmaily
	err := ctx.ReadResource("airbyte:index/sourceSmaily:SourceSmaily", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSmaily resources.
type sourceSmailyState struct {
	Configuration *SourceSmailyConfiguration `pulumi:"configuration"`
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

type SourceSmailyState struct {
	Configuration SourceSmailyConfigurationPtrInput
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

func (SourceSmailyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSmailyState)(nil)).Elem()
}

type sourceSmailyArgs struct {
	Configuration SourceSmailyConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSmaily resource.
type SourceSmailyArgs struct {
	Configuration SourceSmailyConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSmailyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSmailyArgs)(nil)).Elem()
}

type SourceSmailyInput interface {
	pulumi.Input

	ToSourceSmailyOutput() SourceSmailyOutput
	ToSourceSmailyOutputWithContext(ctx context.Context) SourceSmailyOutput
}

func (*SourceSmaily) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSmaily)(nil)).Elem()
}

func (i *SourceSmaily) ToSourceSmailyOutput() SourceSmailyOutput {
	return i.ToSourceSmailyOutputWithContext(context.Background())
}

func (i *SourceSmaily) ToSourceSmailyOutputWithContext(ctx context.Context) SourceSmailyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSmailyOutput)
}

// SourceSmailyArrayInput is an input type that accepts SourceSmailyArray and SourceSmailyArrayOutput values.
// You can construct a concrete instance of `SourceSmailyArrayInput` via:
//
//	SourceSmailyArray{ SourceSmailyArgs{...} }
type SourceSmailyArrayInput interface {
	pulumi.Input

	ToSourceSmailyArrayOutput() SourceSmailyArrayOutput
	ToSourceSmailyArrayOutputWithContext(context.Context) SourceSmailyArrayOutput
}

type SourceSmailyArray []SourceSmailyInput

func (SourceSmailyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSmaily)(nil)).Elem()
}

func (i SourceSmailyArray) ToSourceSmailyArrayOutput() SourceSmailyArrayOutput {
	return i.ToSourceSmailyArrayOutputWithContext(context.Background())
}

func (i SourceSmailyArray) ToSourceSmailyArrayOutputWithContext(ctx context.Context) SourceSmailyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSmailyArrayOutput)
}

// SourceSmailyMapInput is an input type that accepts SourceSmailyMap and SourceSmailyMapOutput values.
// You can construct a concrete instance of `SourceSmailyMapInput` via:
//
//	SourceSmailyMap{ "key": SourceSmailyArgs{...} }
type SourceSmailyMapInput interface {
	pulumi.Input

	ToSourceSmailyMapOutput() SourceSmailyMapOutput
	ToSourceSmailyMapOutputWithContext(context.Context) SourceSmailyMapOutput
}

type SourceSmailyMap map[string]SourceSmailyInput

func (SourceSmailyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSmaily)(nil)).Elem()
}

func (i SourceSmailyMap) ToSourceSmailyMapOutput() SourceSmailyMapOutput {
	return i.ToSourceSmailyMapOutputWithContext(context.Background())
}

func (i SourceSmailyMap) ToSourceSmailyMapOutputWithContext(ctx context.Context) SourceSmailyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSmailyMapOutput)
}

type SourceSmailyOutput struct{ *pulumi.OutputState }

func (SourceSmailyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSmaily)(nil)).Elem()
}

func (o SourceSmailyOutput) ToSourceSmailyOutput() SourceSmailyOutput {
	return o
}

func (o SourceSmailyOutput) ToSourceSmailyOutputWithContext(ctx context.Context) SourceSmailyOutput {
	return o
}

func (o SourceSmailyOutput) Configuration() SourceSmailyConfigurationOutput {
	return o.ApplyT(func(v *SourceSmaily) SourceSmailyConfigurationOutput { return v.Configuration }).(SourceSmailyConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceSmailyOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceSmailyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceSmailyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSmailyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSmailyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSmailyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceSmailyArrayOutput struct{ *pulumi.OutputState }

func (SourceSmailyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSmaily)(nil)).Elem()
}

func (o SourceSmailyArrayOutput) ToSourceSmailyArrayOutput() SourceSmailyArrayOutput {
	return o
}

func (o SourceSmailyArrayOutput) ToSourceSmailyArrayOutputWithContext(ctx context.Context) SourceSmailyArrayOutput {
	return o
}

func (o SourceSmailyArrayOutput) Index(i pulumi.IntInput) SourceSmailyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceSmaily {
		return vs[0].([]*SourceSmaily)[vs[1].(int)]
	}).(SourceSmailyOutput)
}

type SourceSmailyMapOutput struct{ *pulumi.OutputState }

func (SourceSmailyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSmaily)(nil)).Elem()
}

func (o SourceSmailyMapOutput) ToSourceSmailyMapOutput() SourceSmailyMapOutput {
	return o
}

func (o SourceSmailyMapOutput) ToSourceSmailyMapOutputWithContext(ctx context.Context) SourceSmailyMapOutput {
	return o
}

func (o SourceSmailyMapOutput) MapIndex(k pulumi.StringInput) SourceSmailyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceSmaily {
		return vs[0].(map[string]*SourceSmaily)[vs[1].(string)]
	}).(SourceSmailyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSmailyInput)(nil)).Elem(), &SourceSmaily{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSmailyArrayInput)(nil)).Elem(), SourceSmailyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSmailyMapInput)(nil)).Elem(), SourceSmailyMap{})
	pulumi.RegisterOutputType(SourceSmailyOutput{})
	pulumi.RegisterOutputType(SourceSmailyArrayOutput{})
	pulumi.RegisterOutputType(SourceSmailyMapOutput{})
}