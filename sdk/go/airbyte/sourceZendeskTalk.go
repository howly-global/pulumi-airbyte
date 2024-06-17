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

type SourceZendeskTalk struct {
	pulumi.CustomResourceState

	Configuration SourceZendeskTalkConfigurationOutput `pulumi:"configuration"`
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

// NewSourceZendeskTalk registers a new resource with the given unique name, arguments, and options.
func NewSourceZendeskTalk(ctx *pulumi.Context,
	name string, args *SourceZendeskTalkArgs, opts ...pulumi.ResourceOption) (*SourceZendeskTalk, error) {
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
	var resource SourceZendeskTalk
	err := ctx.RegisterResource("airbyte:index/sourceZendeskTalk:SourceZendeskTalk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceZendeskTalk gets an existing SourceZendeskTalk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceZendeskTalk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceZendeskTalkState, opts ...pulumi.ResourceOption) (*SourceZendeskTalk, error) {
	var resource SourceZendeskTalk
	err := ctx.ReadResource("airbyte:index/sourceZendeskTalk:SourceZendeskTalk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceZendeskTalk resources.
type sourceZendeskTalkState struct {
	Configuration *SourceZendeskTalkConfiguration `pulumi:"configuration"`
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

type SourceZendeskTalkState struct {
	Configuration SourceZendeskTalkConfigurationPtrInput
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

func (SourceZendeskTalkState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceZendeskTalkState)(nil)).Elem()
}

type sourceZendeskTalkArgs struct {
	Configuration SourceZendeskTalkConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceZendeskTalk resource.
type SourceZendeskTalkArgs struct {
	Configuration SourceZendeskTalkConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceZendeskTalkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceZendeskTalkArgs)(nil)).Elem()
}

type SourceZendeskTalkInput interface {
	pulumi.Input

	ToSourceZendeskTalkOutput() SourceZendeskTalkOutput
	ToSourceZendeskTalkOutputWithContext(ctx context.Context) SourceZendeskTalkOutput
}

func (*SourceZendeskTalk) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceZendeskTalk)(nil)).Elem()
}

func (i *SourceZendeskTalk) ToSourceZendeskTalkOutput() SourceZendeskTalkOutput {
	return i.ToSourceZendeskTalkOutputWithContext(context.Background())
}

func (i *SourceZendeskTalk) ToSourceZendeskTalkOutputWithContext(ctx context.Context) SourceZendeskTalkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZendeskTalkOutput)
}

// SourceZendeskTalkArrayInput is an input type that accepts SourceZendeskTalkArray and SourceZendeskTalkArrayOutput values.
// You can construct a concrete instance of `SourceZendeskTalkArrayInput` via:
//
//	SourceZendeskTalkArray{ SourceZendeskTalkArgs{...} }
type SourceZendeskTalkArrayInput interface {
	pulumi.Input

	ToSourceZendeskTalkArrayOutput() SourceZendeskTalkArrayOutput
	ToSourceZendeskTalkArrayOutputWithContext(context.Context) SourceZendeskTalkArrayOutput
}

type SourceZendeskTalkArray []SourceZendeskTalkInput

func (SourceZendeskTalkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceZendeskTalk)(nil)).Elem()
}

func (i SourceZendeskTalkArray) ToSourceZendeskTalkArrayOutput() SourceZendeskTalkArrayOutput {
	return i.ToSourceZendeskTalkArrayOutputWithContext(context.Background())
}

func (i SourceZendeskTalkArray) ToSourceZendeskTalkArrayOutputWithContext(ctx context.Context) SourceZendeskTalkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZendeskTalkArrayOutput)
}

// SourceZendeskTalkMapInput is an input type that accepts SourceZendeskTalkMap and SourceZendeskTalkMapOutput values.
// You can construct a concrete instance of `SourceZendeskTalkMapInput` via:
//
//	SourceZendeskTalkMap{ "key": SourceZendeskTalkArgs{...} }
type SourceZendeskTalkMapInput interface {
	pulumi.Input

	ToSourceZendeskTalkMapOutput() SourceZendeskTalkMapOutput
	ToSourceZendeskTalkMapOutputWithContext(context.Context) SourceZendeskTalkMapOutput
}

type SourceZendeskTalkMap map[string]SourceZendeskTalkInput

func (SourceZendeskTalkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceZendeskTalk)(nil)).Elem()
}

func (i SourceZendeskTalkMap) ToSourceZendeskTalkMapOutput() SourceZendeskTalkMapOutput {
	return i.ToSourceZendeskTalkMapOutputWithContext(context.Background())
}

func (i SourceZendeskTalkMap) ToSourceZendeskTalkMapOutputWithContext(ctx context.Context) SourceZendeskTalkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZendeskTalkMapOutput)
}

type SourceZendeskTalkOutput struct{ *pulumi.OutputState }

func (SourceZendeskTalkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceZendeskTalk)(nil)).Elem()
}

func (o SourceZendeskTalkOutput) ToSourceZendeskTalkOutput() SourceZendeskTalkOutput {
	return o
}

func (o SourceZendeskTalkOutput) ToSourceZendeskTalkOutputWithContext(ctx context.Context) SourceZendeskTalkOutput {
	return o
}

func (o SourceZendeskTalkOutput) Configuration() SourceZendeskTalkConfigurationOutput {
	return o.ApplyT(func(v *SourceZendeskTalk) SourceZendeskTalkConfigurationOutput { return v.Configuration }).(SourceZendeskTalkConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceZendeskTalkOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceZendeskTalk) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceZendeskTalkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskTalk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceZendeskTalkOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceZendeskTalk) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceZendeskTalkOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskTalk) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceZendeskTalkOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskTalk) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceZendeskTalkOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskTalk) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceZendeskTalkArrayOutput struct{ *pulumi.OutputState }

func (SourceZendeskTalkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceZendeskTalk)(nil)).Elem()
}

func (o SourceZendeskTalkArrayOutput) ToSourceZendeskTalkArrayOutput() SourceZendeskTalkArrayOutput {
	return o
}

func (o SourceZendeskTalkArrayOutput) ToSourceZendeskTalkArrayOutputWithContext(ctx context.Context) SourceZendeskTalkArrayOutput {
	return o
}

func (o SourceZendeskTalkArrayOutput) Index(i pulumi.IntInput) SourceZendeskTalkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceZendeskTalk {
		return vs[0].([]*SourceZendeskTalk)[vs[1].(int)]
	}).(SourceZendeskTalkOutput)
}

type SourceZendeskTalkMapOutput struct{ *pulumi.OutputState }

func (SourceZendeskTalkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceZendeskTalk)(nil)).Elem()
}

func (o SourceZendeskTalkMapOutput) ToSourceZendeskTalkMapOutput() SourceZendeskTalkMapOutput {
	return o
}

func (o SourceZendeskTalkMapOutput) ToSourceZendeskTalkMapOutputWithContext(ctx context.Context) SourceZendeskTalkMapOutput {
	return o
}

func (o SourceZendeskTalkMapOutput) MapIndex(k pulumi.StringInput) SourceZendeskTalkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceZendeskTalk {
		return vs[0].(map[string]*SourceZendeskTalk)[vs[1].(string)]
	}).(SourceZendeskTalkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZendeskTalkInput)(nil)).Elem(), &SourceZendeskTalk{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZendeskTalkArrayInput)(nil)).Elem(), SourceZendeskTalkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZendeskTalkMapInput)(nil)).Elem(), SourceZendeskTalkMap{})
	pulumi.RegisterOutputType(SourceZendeskTalkOutput{})
	pulumi.RegisterOutputType(SourceZendeskTalkArrayOutput{})
	pulumi.RegisterOutputType(SourceZendeskTalkMapOutput{})
}