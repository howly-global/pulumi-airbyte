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

type SourceNetsuite struct {
	pulumi.CustomResourceState

	Configuration SourceNetsuiteConfigurationOutput `pulumi:"configuration"`
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

// NewSourceNetsuite registers a new resource with the given unique name, arguments, and options.
func NewSourceNetsuite(ctx *pulumi.Context,
	name string, args *SourceNetsuiteArgs, opts ...pulumi.ResourceOption) (*SourceNetsuite, error) {
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
	var resource SourceNetsuite
	err := ctx.RegisterResource("airbyte:index/sourceNetsuite:SourceNetsuite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceNetsuite gets an existing SourceNetsuite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceNetsuite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceNetsuiteState, opts ...pulumi.ResourceOption) (*SourceNetsuite, error) {
	var resource SourceNetsuite
	err := ctx.ReadResource("airbyte:index/sourceNetsuite:SourceNetsuite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceNetsuite resources.
type sourceNetsuiteState struct {
	Configuration *SourceNetsuiteConfiguration `pulumi:"configuration"`
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

type SourceNetsuiteState struct {
	Configuration SourceNetsuiteConfigurationPtrInput
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

func (SourceNetsuiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNetsuiteState)(nil)).Elem()
}

type sourceNetsuiteArgs struct {
	Configuration SourceNetsuiteConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceNetsuite resource.
type SourceNetsuiteArgs struct {
	Configuration SourceNetsuiteConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceNetsuiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNetsuiteArgs)(nil)).Elem()
}

type SourceNetsuiteInput interface {
	pulumi.Input

	ToSourceNetsuiteOutput() SourceNetsuiteOutput
	ToSourceNetsuiteOutputWithContext(ctx context.Context) SourceNetsuiteOutput
}

func (*SourceNetsuite) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNetsuite)(nil)).Elem()
}

func (i *SourceNetsuite) ToSourceNetsuiteOutput() SourceNetsuiteOutput {
	return i.ToSourceNetsuiteOutputWithContext(context.Background())
}

func (i *SourceNetsuite) ToSourceNetsuiteOutputWithContext(ctx context.Context) SourceNetsuiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNetsuiteOutput)
}

// SourceNetsuiteArrayInput is an input type that accepts SourceNetsuiteArray and SourceNetsuiteArrayOutput values.
// You can construct a concrete instance of `SourceNetsuiteArrayInput` via:
//
//	SourceNetsuiteArray{ SourceNetsuiteArgs{...} }
type SourceNetsuiteArrayInput interface {
	pulumi.Input

	ToSourceNetsuiteArrayOutput() SourceNetsuiteArrayOutput
	ToSourceNetsuiteArrayOutputWithContext(context.Context) SourceNetsuiteArrayOutput
}

type SourceNetsuiteArray []SourceNetsuiteInput

func (SourceNetsuiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceNetsuite)(nil)).Elem()
}

func (i SourceNetsuiteArray) ToSourceNetsuiteArrayOutput() SourceNetsuiteArrayOutput {
	return i.ToSourceNetsuiteArrayOutputWithContext(context.Background())
}

func (i SourceNetsuiteArray) ToSourceNetsuiteArrayOutputWithContext(ctx context.Context) SourceNetsuiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNetsuiteArrayOutput)
}

// SourceNetsuiteMapInput is an input type that accepts SourceNetsuiteMap and SourceNetsuiteMapOutput values.
// You can construct a concrete instance of `SourceNetsuiteMapInput` via:
//
//	SourceNetsuiteMap{ "key": SourceNetsuiteArgs{...} }
type SourceNetsuiteMapInput interface {
	pulumi.Input

	ToSourceNetsuiteMapOutput() SourceNetsuiteMapOutput
	ToSourceNetsuiteMapOutputWithContext(context.Context) SourceNetsuiteMapOutput
}

type SourceNetsuiteMap map[string]SourceNetsuiteInput

func (SourceNetsuiteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceNetsuite)(nil)).Elem()
}

func (i SourceNetsuiteMap) ToSourceNetsuiteMapOutput() SourceNetsuiteMapOutput {
	return i.ToSourceNetsuiteMapOutputWithContext(context.Background())
}

func (i SourceNetsuiteMap) ToSourceNetsuiteMapOutputWithContext(ctx context.Context) SourceNetsuiteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNetsuiteMapOutput)
}

type SourceNetsuiteOutput struct{ *pulumi.OutputState }

func (SourceNetsuiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNetsuite)(nil)).Elem()
}

func (o SourceNetsuiteOutput) ToSourceNetsuiteOutput() SourceNetsuiteOutput {
	return o
}

func (o SourceNetsuiteOutput) ToSourceNetsuiteOutputWithContext(ctx context.Context) SourceNetsuiteOutput {
	return o
}

func (o SourceNetsuiteOutput) Configuration() SourceNetsuiteConfigurationOutput {
	return o.ApplyT(func(v *SourceNetsuite) SourceNetsuiteConfigurationOutput { return v.Configuration }).(SourceNetsuiteConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceNetsuiteOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceNetsuiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceNetsuiteOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceNetsuiteOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceNetsuiteOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceNetsuiteOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceNetsuiteArrayOutput struct{ *pulumi.OutputState }

func (SourceNetsuiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceNetsuite)(nil)).Elem()
}

func (o SourceNetsuiteArrayOutput) ToSourceNetsuiteArrayOutput() SourceNetsuiteArrayOutput {
	return o
}

func (o SourceNetsuiteArrayOutput) ToSourceNetsuiteArrayOutputWithContext(ctx context.Context) SourceNetsuiteArrayOutput {
	return o
}

func (o SourceNetsuiteArrayOutput) Index(i pulumi.IntInput) SourceNetsuiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceNetsuite {
		return vs[0].([]*SourceNetsuite)[vs[1].(int)]
	}).(SourceNetsuiteOutput)
}

type SourceNetsuiteMapOutput struct{ *pulumi.OutputState }

func (SourceNetsuiteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceNetsuite)(nil)).Elem()
}

func (o SourceNetsuiteMapOutput) ToSourceNetsuiteMapOutput() SourceNetsuiteMapOutput {
	return o
}

func (o SourceNetsuiteMapOutput) ToSourceNetsuiteMapOutputWithContext(ctx context.Context) SourceNetsuiteMapOutput {
	return o
}

func (o SourceNetsuiteMapOutput) MapIndex(k pulumi.StringInput) SourceNetsuiteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceNetsuite {
		return vs[0].(map[string]*SourceNetsuite)[vs[1].(string)]
	}).(SourceNetsuiteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNetsuiteInput)(nil)).Elem(), &SourceNetsuite{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNetsuiteArrayInput)(nil)).Elem(), SourceNetsuiteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNetsuiteMapInput)(nil)).Elem(), SourceNetsuiteMap{})
	pulumi.RegisterOutputType(SourceNetsuiteOutput{})
	pulumi.RegisterOutputType(SourceNetsuiteArrayOutput{})
	pulumi.RegisterOutputType(SourceNetsuiteMapOutput{})
}