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

type DestinationS3Glue struct {
	pulumi.CustomResourceState

	Configuration DestinationS3GlueConfigurationOutput `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId    pulumi.StringPtrOutput `pulumi:"definitionId"`
	DestinationId   pulumi.StringOutput    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput    `pulumi:"destinationType"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringOutput `pulumi:"name"`
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewDestinationS3Glue registers a new resource with the given unique name, arguments, and options.
func NewDestinationS3Glue(ctx *pulumi.Context,
	name string, args *DestinationS3GlueArgs, opts ...pulumi.ResourceOption) (*DestinationS3Glue, error) {
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
	var resource DestinationS3Glue
	err := ctx.RegisterResource("airbyte:index/destinationS3Glue:DestinationS3Glue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationS3Glue gets an existing DestinationS3Glue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationS3Glue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationS3GlueState, opts ...pulumi.ResourceOption) (*DestinationS3Glue, error) {
	var resource DestinationS3Glue
	err := ctx.ReadResource("airbyte:index/destinationS3Glue:DestinationS3Glue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationS3Glue resources.
type destinationS3GlueState struct {
	Configuration *DestinationS3GlueConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId    *string `pulumi:"definitionId"`
	DestinationId   *string `pulumi:"destinationId"`
	DestinationType *string `pulumi:"destinationType"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        *string `pulumi:"name"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type DestinationS3GlueState struct {
	Configuration DestinationS3GlueConfigurationPtrInput
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId    pulumi.StringPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (DestinationS3GlueState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationS3GlueState)(nil)).Elem()
}

type destinationS3GlueArgs struct {
	Configuration DestinationS3GlueConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        *string `pulumi:"name"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationS3Glue resource.
type DestinationS3GlueArgs struct {
	Configuration DestinationS3GlueConfigurationInput
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (DestinationS3GlueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationS3GlueArgs)(nil)).Elem()
}

type DestinationS3GlueInput interface {
	pulumi.Input

	ToDestinationS3GlueOutput() DestinationS3GlueOutput
	ToDestinationS3GlueOutputWithContext(ctx context.Context) DestinationS3GlueOutput
}

func (*DestinationS3Glue) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationS3Glue)(nil)).Elem()
}

func (i *DestinationS3Glue) ToDestinationS3GlueOutput() DestinationS3GlueOutput {
	return i.ToDestinationS3GlueOutputWithContext(context.Background())
}

func (i *DestinationS3Glue) ToDestinationS3GlueOutputWithContext(ctx context.Context) DestinationS3GlueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationS3GlueOutput)
}

// DestinationS3GlueArrayInput is an input type that accepts DestinationS3GlueArray and DestinationS3GlueArrayOutput values.
// You can construct a concrete instance of `DestinationS3GlueArrayInput` via:
//
//	DestinationS3GlueArray{ DestinationS3GlueArgs{...} }
type DestinationS3GlueArrayInput interface {
	pulumi.Input

	ToDestinationS3GlueArrayOutput() DestinationS3GlueArrayOutput
	ToDestinationS3GlueArrayOutputWithContext(context.Context) DestinationS3GlueArrayOutput
}

type DestinationS3GlueArray []DestinationS3GlueInput

func (DestinationS3GlueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationS3Glue)(nil)).Elem()
}

func (i DestinationS3GlueArray) ToDestinationS3GlueArrayOutput() DestinationS3GlueArrayOutput {
	return i.ToDestinationS3GlueArrayOutputWithContext(context.Background())
}

func (i DestinationS3GlueArray) ToDestinationS3GlueArrayOutputWithContext(ctx context.Context) DestinationS3GlueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationS3GlueArrayOutput)
}

// DestinationS3GlueMapInput is an input type that accepts DestinationS3GlueMap and DestinationS3GlueMapOutput values.
// You can construct a concrete instance of `DestinationS3GlueMapInput` via:
//
//	DestinationS3GlueMap{ "key": DestinationS3GlueArgs{...} }
type DestinationS3GlueMapInput interface {
	pulumi.Input

	ToDestinationS3GlueMapOutput() DestinationS3GlueMapOutput
	ToDestinationS3GlueMapOutputWithContext(context.Context) DestinationS3GlueMapOutput
}

type DestinationS3GlueMap map[string]DestinationS3GlueInput

func (DestinationS3GlueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationS3Glue)(nil)).Elem()
}

func (i DestinationS3GlueMap) ToDestinationS3GlueMapOutput() DestinationS3GlueMapOutput {
	return i.ToDestinationS3GlueMapOutputWithContext(context.Background())
}

func (i DestinationS3GlueMap) ToDestinationS3GlueMapOutputWithContext(ctx context.Context) DestinationS3GlueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationS3GlueMapOutput)
}

type DestinationS3GlueOutput struct{ *pulumi.OutputState }

func (DestinationS3GlueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationS3Glue)(nil)).Elem()
}

func (o DestinationS3GlueOutput) ToDestinationS3GlueOutput() DestinationS3GlueOutput {
	return o
}

func (o DestinationS3GlueOutput) ToDestinationS3GlueOutputWithContext(ctx context.Context) DestinationS3GlueOutput {
	return o
}

func (o DestinationS3GlueOutput) Configuration() DestinationS3GlueConfigurationOutput {
	return o.ApplyT(func(v *DestinationS3Glue) DestinationS3GlueConfigurationOutput { return v.Configuration }).(DestinationS3GlueConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
// replacement if changed.
func (o DestinationS3GlueOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationS3Glue) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

func (o DestinationS3GlueOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationS3Glue) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationS3GlueOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationS3Glue) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

// Name of the destination e.g. dev-mysql-instance.
func (o DestinationS3GlueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationS3Glue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationS3GlueOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationS3Glue) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationS3GlueArrayOutput struct{ *pulumi.OutputState }

func (DestinationS3GlueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationS3Glue)(nil)).Elem()
}

func (o DestinationS3GlueArrayOutput) ToDestinationS3GlueArrayOutput() DestinationS3GlueArrayOutput {
	return o
}

func (o DestinationS3GlueArrayOutput) ToDestinationS3GlueArrayOutputWithContext(ctx context.Context) DestinationS3GlueArrayOutput {
	return o
}

func (o DestinationS3GlueArrayOutput) Index(i pulumi.IntInput) DestinationS3GlueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationS3Glue {
		return vs[0].([]*DestinationS3Glue)[vs[1].(int)]
	}).(DestinationS3GlueOutput)
}

type DestinationS3GlueMapOutput struct{ *pulumi.OutputState }

func (DestinationS3GlueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationS3Glue)(nil)).Elem()
}

func (o DestinationS3GlueMapOutput) ToDestinationS3GlueMapOutput() DestinationS3GlueMapOutput {
	return o
}

func (o DestinationS3GlueMapOutput) ToDestinationS3GlueMapOutputWithContext(ctx context.Context) DestinationS3GlueMapOutput {
	return o
}

func (o DestinationS3GlueMapOutput) MapIndex(k pulumi.StringInput) DestinationS3GlueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationS3Glue {
		return vs[0].(map[string]*DestinationS3Glue)[vs[1].(string)]
	}).(DestinationS3GlueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationS3GlueInput)(nil)).Elem(), &DestinationS3Glue{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationS3GlueArrayInput)(nil)).Elem(), DestinationS3GlueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationS3GlueMapInput)(nil)).Elem(), DestinationS3GlueMap{})
	pulumi.RegisterOutputType(DestinationS3GlueOutput{})
	pulumi.RegisterOutputType(DestinationS3GlueArrayOutput{})
	pulumi.RegisterOutputType(DestinationS3GlueMapOutput{})
}