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

type DestinationSftpJson struct {
	pulumi.CustomResourceState

	Configuration DestinationSftpJsonConfigurationOutput `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId    pulumi.StringPtrOutput `pulumi:"definitionId"`
	DestinationId   pulumi.StringOutput    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput    `pulumi:"destinationType"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringOutput `pulumi:"name"`
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewDestinationSftpJson registers a new resource with the given unique name, arguments, and options.
func NewDestinationSftpJson(ctx *pulumi.Context,
	name string, args *DestinationSftpJsonArgs, opts ...pulumi.ResourceOption) (*DestinationSftpJson, error) {
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
	var resource DestinationSftpJson
	err := ctx.RegisterResource("airbyte:index/destinationSftpJson:DestinationSftpJson", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationSftpJson gets an existing DestinationSftpJson resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationSftpJson(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationSftpJsonState, opts ...pulumi.ResourceOption) (*DestinationSftpJson, error) {
	var resource DestinationSftpJson
	err := ctx.ReadResource("airbyte:index/destinationSftpJson:DestinationSftpJson", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationSftpJson resources.
type destinationSftpJsonState struct {
	Configuration *DestinationSftpJsonConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId    *string `pulumi:"definitionId"`
	DestinationId   *string `pulumi:"destinationId"`
	DestinationType *string `pulumi:"destinationType"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        *string `pulumi:"name"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type DestinationSftpJsonState struct {
	Configuration DestinationSftpJsonConfigurationPtrInput
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId    pulumi.StringPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (DestinationSftpJsonState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationSftpJsonState)(nil)).Elem()
}

type destinationSftpJsonArgs struct {
	Configuration DestinationSftpJsonConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        *string `pulumi:"name"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationSftpJson resource.
type DestinationSftpJsonArgs struct {
	Configuration DestinationSftpJsonConfigurationInput
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (DestinationSftpJsonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationSftpJsonArgs)(nil)).Elem()
}

type DestinationSftpJsonInput interface {
	pulumi.Input

	ToDestinationSftpJsonOutput() DestinationSftpJsonOutput
	ToDestinationSftpJsonOutputWithContext(ctx context.Context) DestinationSftpJsonOutput
}

func (*DestinationSftpJson) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationSftpJson)(nil)).Elem()
}

func (i *DestinationSftpJson) ToDestinationSftpJsonOutput() DestinationSftpJsonOutput {
	return i.ToDestinationSftpJsonOutputWithContext(context.Background())
}

func (i *DestinationSftpJson) ToDestinationSftpJsonOutputWithContext(ctx context.Context) DestinationSftpJsonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationSftpJsonOutput)
}

// DestinationSftpJsonArrayInput is an input type that accepts DestinationSftpJsonArray and DestinationSftpJsonArrayOutput values.
// You can construct a concrete instance of `DestinationSftpJsonArrayInput` via:
//
//	DestinationSftpJsonArray{ DestinationSftpJsonArgs{...} }
type DestinationSftpJsonArrayInput interface {
	pulumi.Input

	ToDestinationSftpJsonArrayOutput() DestinationSftpJsonArrayOutput
	ToDestinationSftpJsonArrayOutputWithContext(context.Context) DestinationSftpJsonArrayOutput
}

type DestinationSftpJsonArray []DestinationSftpJsonInput

func (DestinationSftpJsonArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationSftpJson)(nil)).Elem()
}

func (i DestinationSftpJsonArray) ToDestinationSftpJsonArrayOutput() DestinationSftpJsonArrayOutput {
	return i.ToDestinationSftpJsonArrayOutputWithContext(context.Background())
}

func (i DestinationSftpJsonArray) ToDestinationSftpJsonArrayOutputWithContext(ctx context.Context) DestinationSftpJsonArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationSftpJsonArrayOutput)
}

// DestinationSftpJsonMapInput is an input type that accepts DestinationSftpJsonMap and DestinationSftpJsonMapOutput values.
// You can construct a concrete instance of `DestinationSftpJsonMapInput` via:
//
//	DestinationSftpJsonMap{ "key": DestinationSftpJsonArgs{...} }
type DestinationSftpJsonMapInput interface {
	pulumi.Input

	ToDestinationSftpJsonMapOutput() DestinationSftpJsonMapOutput
	ToDestinationSftpJsonMapOutputWithContext(context.Context) DestinationSftpJsonMapOutput
}

type DestinationSftpJsonMap map[string]DestinationSftpJsonInput

func (DestinationSftpJsonMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationSftpJson)(nil)).Elem()
}

func (i DestinationSftpJsonMap) ToDestinationSftpJsonMapOutput() DestinationSftpJsonMapOutput {
	return i.ToDestinationSftpJsonMapOutputWithContext(context.Background())
}

func (i DestinationSftpJsonMap) ToDestinationSftpJsonMapOutputWithContext(ctx context.Context) DestinationSftpJsonMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationSftpJsonMapOutput)
}

type DestinationSftpJsonOutput struct{ *pulumi.OutputState }

func (DestinationSftpJsonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationSftpJson)(nil)).Elem()
}

func (o DestinationSftpJsonOutput) ToDestinationSftpJsonOutput() DestinationSftpJsonOutput {
	return o
}

func (o DestinationSftpJsonOutput) ToDestinationSftpJsonOutputWithContext(ctx context.Context) DestinationSftpJsonOutput {
	return o
}

func (o DestinationSftpJsonOutput) Configuration() DestinationSftpJsonConfigurationOutput {
	return o.ApplyT(func(v *DestinationSftpJson) DestinationSftpJsonConfigurationOutput { return v.Configuration }).(DestinationSftpJsonConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
// replacement if changed.
func (o DestinationSftpJsonOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationSftpJson) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

func (o DestinationSftpJsonOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSftpJson) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationSftpJsonOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSftpJson) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

// Name of the destination e.g. dev-mysql-instance.
func (o DestinationSftpJsonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSftpJson) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationSftpJsonOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSftpJson) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationSftpJsonArrayOutput struct{ *pulumi.OutputState }

func (DestinationSftpJsonArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationSftpJson)(nil)).Elem()
}

func (o DestinationSftpJsonArrayOutput) ToDestinationSftpJsonArrayOutput() DestinationSftpJsonArrayOutput {
	return o
}

func (o DestinationSftpJsonArrayOutput) ToDestinationSftpJsonArrayOutputWithContext(ctx context.Context) DestinationSftpJsonArrayOutput {
	return o
}

func (o DestinationSftpJsonArrayOutput) Index(i pulumi.IntInput) DestinationSftpJsonOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationSftpJson {
		return vs[0].([]*DestinationSftpJson)[vs[1].(int)]
	}).(DestinationSftpJsonOutput)
}

type DestinationSftpJsonMapOutput struct{ *pulumi.OutputState }

func (DestinationSftpJsonMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationSftpJson)(nil)).Elem()
}

func (o DestinationSftpJsonMapOutput) ToDestinationSftpJsonMapOutput() DestinationSftpJsonMapOutput {
	return o
}

func (o DestinationSftpJsonMapOutput) ToDestinationSftpJsonMapOutputWithContext(ctx context.Context) DestinationSftpJsonMapOutput {
	return o
}

func (o DestinationSftpJsonMapOutput) MapIndex(k pulumi.StringInput) DestinationSftpJsonOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationSftpJson {
		return vs[0].(map[string]*DestinationSftpJson)[vs[1].(string)]
	}).(DestinationSftpJsonOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationSftpJsonInput)(nil)).Elem(), &DestinationSftpJson{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationSftpJsonArrayInput)(nil)).Elem(), DestinationSftpJsonArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationSftpJsonMapInput)(nil)).Elem(), DestinationSftpJsonMap{})
	pulumi.RegisterOutputType(DestinationSftpJsonOutput{})
	pulumi.RegisterOutputType(DestinationSftpJsonArrayOutput{})
	pulumi.RegisterOutputType(DestinationSftpJsonMapOutput{})
}