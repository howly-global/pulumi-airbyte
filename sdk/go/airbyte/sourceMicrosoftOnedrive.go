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

type SourceMicrosoftOnedrive struct {
	pulumi.CustomResourceState

	// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification. This class combines the authentication
	// details with additional configuration for the OneDrive API.
	Configuration SourceMicrosoftOnedriveConfigurationOutput `pulumi:"configuration"`
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

// NewSourceMicrosoftOnedrive registers a new resource with the given unique name, arguments, and options.
func NewSourceMicrosoftOnedrive(ctx *pulumi.Context,
	name string, args *SourceMicrosoftOnedriveArgs, opts ...pulumi.ResourceOption) (*SourceMicrosoftOnedrive, error) {
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
	var resource SourceMicrosoftOnedrive
	err := ctx.RegisterResource("airbyte:index/sourceMicrosoftOnedrive:SourceMicrosoftOnedrive", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceMicrosoftOnedrive gets an existing SourceMicrosoftOnedrive resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceMicrosoftOnedrive(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceMicrosoftOnedriveState, opts ...pulumi.ResourceOption) (*SourceMicrosoftOnedrive, error) {
	var resource SourceMicrosoftOnedrive
	err := ctx.ReadResource("airbyte:index/sourceMicrosoftOnedrive:SourceMicrosoftOnedrive", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceMicrosoftOnedrive resources.
type sourceMicrosoftOnedriveState struct {
	// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification. This class combines the authentication
	// details with additional configuration for the OneDrive API.
	Configuration *SourceMicrosoftOnedriveConfiguration `pulumi:"configuration"`
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

type SourceMicrosoftOnedriveState struct {
	// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification. This class combines the authentication
	// details with additional configuration for the OneDrive API.
	Configuration SourceMicrosoftOnedriveConfigurationPtrInput
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

func (SourceMicrosoftOnedriveState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMicrosoftOnedriveState)(nil)).Elem()
}

type sourceMicrosoftOnedriveArgs struct {
	// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification. This class combines the authentication
	// details with additional configuration for the OneDrive API.
	Configuration SourceMicrosoftOnedriveConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceMicrosoftOnedrive resource.
type SourceMicrosoftOnedriveArgs struct {
	// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification. This class combines the authentication
	// details with additional configuration for the OneDrive API.
	Configuration SourceMicrosoftOnedriveConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceMicrosoftOnedriveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMicrosoftOnedriveArgs)(nil)).Elem()
}

type SourceMicrosoftOnedriveInput interface {
	pulumi.Input

	ToSourceMicrosoftOnedriveOutput() SourceMicrosoftOnedriveOutput
	ToSourceMicrosoftOnedriveOutputWithContext(ctx context.Context) SourceMicrosoftOnedriveOutput
}

func (*SourceMicrosoftOnedrive) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMicrosoftOnedrive)(nil)).Elem()
}

func (i *SourceMicrosoftOnedrive) ToSourceMicrosoftOnedriveOutput() SourceMicrosoftOnedriveOutput {
	return i.ToSourceMicrosoftOnedriveOutputWithContext(context.Background())
}

func (i *SourceMicrosoftOnedrive) ToSourceMicrosoftOnedriveOutputWithContext(ctx context.Context) SourceMicrosoftOnedriveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMicrosoftOnedriveOutput)
}

// SourceMicrosoftOnedriveArrayInput is an input type that accepts SourceMicrosoftOnedriveArray and SourceMicrosoftOnedriveArrayOutput values.
// You can construct a concrete instance of `SourceMicrosoftOnedriveArrayInput` via:
//
//	SourceMicrosoftOnedriveArray{ SourceMicrosoftOnedriveArgs{...} }
type SourceMicrosoftOnedriveArrayInput interface {
	pulumi.Input

	ToSourceMicrosoftOnedriveArrayOutput() SourceMicrosoftOnedriveArrayOutput
	ToSourceMicrosoftOnedriveArrayOutputWithContext(context.Context) SourceMicrosoftOnedriveArrayOutput
}

type SourceMicrosoftOnedriveArray []SourceMicrosoftOnedriveInput

func (SourceMicrosoftOnedriveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMicrosoftOnedrive)(nil)).Elem()
}

func (i SourceMicrosoftOnedriveArray) ToSourceMicrosoftOnedriveArrayOutput() SourceMicrosoftOnedriveArrayOutput {
	return i.ToSourceMicrosoftOnedriveArrayOutputWithContext(context.Background())
}

func (i SourceMicrosoftOnedriveArray) ToSourceMicrosoftOnedriveArrayOutputWithContext(ctx context.Context) SourceMicrosoftOnedriveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMicrosoftOnedriveArrayOutput)
}

// SourceMicrosoftOnedriveMapInput is an input type that accepts SourceMicrosoftOnedriveMap and SourceMicrosoftOnedriveMapOutput values.
// You can construct a concrete instance of `SourceMicrosoftOnedriveMapInput` via:
//
//	SourceMicrosoftOnedriveMap{ "key": SourceMicrosoftOnedriveArgs{...} }
type SourceMicrosoftOnedriveMapInput interface {
	pulumi.Input

	ToSourceMicrosoftOnedriveMapOutput() SourceMicrosoftOnedriveMapOutput
	ToSourceMicrosoftOnedriveMapOutputWithContext(context.Context) SourceMicrosoftOnedriveMapOutput
}

type SourceMicrosoftOnedriveMap map[string]SourceMicrosoftOnedriveInput

func (SourceMicrosoftOnedriveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMicrosoftOnedrive)(nil)).Elem()
}

func (i SourceMicrosoftOnedriveMap) ToSourceMicrosoftOnedriveMapOutput() SourceMicrosoftOnedriveMapOutput {
	return i.ToSourceMicrosoftOnedriveMapOutputWithContext(context.Background())
}

func (i SourceMicrosoftOnedriveMap) ToSourceMicrosoftOnedriveMapOutputWithContext(ctx context.Context) SourceMicrosoftOnedriveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMicrosoftOnedriveMapOutput)
}

type SourceMicrosoftOnedriveOutput struct{ *pulumi.OutputState }

func (SourceMicrosoftOnedriveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMicrosoftOnedrive)(nil)).Elem()
}

func (o SourceMicrosoftOnedriveOutput) ToSourceMicrosoftOnedriveOutput() SourceMicrosoftOnedriveOutput {
	return o
}

func (o SourceMicrosoftOnedriveOutput) ToSourceMicrosoftOnedriveOutputWithContext(ctx context.Context) SourceMicrosoftOnedriveOutput {
	return o
}

// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification. This class combines the authentication
// details with additional configuration for the OneDrive API.
func (o SourceMicrosoftOnedriveOutput) Configuration() SourceMicrosoftOnedriveConfigurationOutput {
	return o.ApplyT(func(v *SourceMicrosoftOnedrive) SourceMicrosoftOnedriveConfigurationOutput { return v.Configuration }).(SourceMicrosoftOnedriveConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceMicrosoftOnedriveOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMicrosoftOnedrive) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceMicrosoftOnedriveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftOnedrive) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceMicrosoftOnedriveOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMicrosoftOnedrive) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceMicrosoftOnedriveOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftOnedrive) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceMicrosoftOnedriveOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftOnedrive) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceMicrosoftOnedriveOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftOnedrive) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceMicrosoftOnedriveArrayOutput struct{ *pulumi.OutputState }

func (SourceMicrosoftOnedriveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMicrosoftOnedrive)(nil)).Elem()
}

func (o SourceMicrosoftOnedriveArrayOutput) ToSourceMicrosoftOnedriveArrayOutput() SourceMicrosoftOnedriveArrayOutput {
	return o
}

func (o SourceMicrosoftOnedriveArrayOutput) ToSourceMicrosoftOnedriveArrayOutputWithContext(ctx context.Context) SourceMicrosoftOnedriveArrayOutput {
	return o
}

func (o SourceMicrosoftOnedriveArrayOutput) Index(i pulumi.IntInput) SourceMicrosoftOnedriveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceMicrosoftOnedrive {
		return vs[0].([]*SourceMicrosoftOnedrive)[vs[1].(int)]
	}).(SourceMicrosoftOnedriveOutput)
}

type SourceMicrosoftOnedriveMapOutput struct{ *pulumi.OutputState }

func (SourceMicrosoftOnedriveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMicrosoftOnedrive)(nil)).Elem()
}

func (o SourceMicrosoftOnedriveMapOutput) ToSourceMicrosoftOnedriveMapOutput() SourceMicrosoftOnedriveMapOutput {
	return o
}

func (o SourceMicrosoftOnedriveMapOutput) ToSourceMicrosoftOnedriveMapOutputWithContext(ctx context.Context) SourceMicrosoftOnedriveMapOutput {
	return o
}

func (o SourceMicrosoftOnedriveMapOutput) MapIndex(k pulumi.StringInput) SourceMicrosoftOnedriveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceMicrosoftOnedrive {
		return vs[0].(map[string]*SourceMicrosoftOnedrive)[vs[1].(string)]
	}).(SourceMicrosoftOnedriveOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMicrosoftOnedriveInput)(nil)).Elem(), &SourceMicrosoftOnedrive{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMicrosoftOnedriveArrayInput)(nil)).Elem(), SourceMicrosoftOnedriveArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMicrosoftOnedriveMapInput)(nil)).Elem(), SourceMicrosoftOnedriveMap{})
	pulumi.RegisterOutputType(SourceMicrosoftOnedriveOutput{})
	pulumi.RegisterOutputType(SourceMicrosoftOnedriveArrayOutput{})
	pulumi.RegisterOutputType(SourceMicrosoftOnedriveMapOutput{})
}