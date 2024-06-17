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

type SourceAmazonAds struct {
	pulumi.CustomResourceState

	Configuration SourceAmazonAdsConfigurationOutput `pulumi:"configuration"`
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

// NewSourceAmazonAds registers a new resource with the given unique name, arguments, and options.
func NewSourceAmazonAds(ctx *pulumi.Context,
	name string, args *SourceAmazonAdsArgs, opts ...pulumi.ResourceOption) (*SourceAmazonAds, error) {
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
	var resource SourceAmazonAds
	err := ctx.RegisterResource("airbyte:index/sourceAmazonAds:SourceAmazonAds", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAmazonAds gets an existing SourceAmazonAds resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAmazonAds(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAmazonAdsState, opts ...pulumi.ResourceOption) (*SourceAmazonAds, error) {
	var resource SourceAmazonAds
	err := ctx.ReadResource("airbyte:index/sourceAmazonAds:SourceAmazonAds", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAmazonAds resources.
type sourceAmazonAdsState struct {
	Configuration *SourceAmazonAdsConfiguration `pulumi:"configuration"`
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

type SourceAmazonAdsState struct {
	Configuration SourceAmazonAdsConfigurationPtrInput
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

func (SourceAmazonAdsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmazonAdsState)(nil)).Elem()
}

type sourceAmazonAdsArgs struct {
	Configuration SourceAmazonAdsConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAmazonAds resource.
type SourceAmazonAdsArgs struct {
	Configuration SourceAmazonAdsConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
	// replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAmazonAdsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmazonAdsArgs)(nil)).Elem()
}

type SourceAmazonAdsInput interface {
	pulumi.Input

	ToSourceAmazonAdsOutput() SourceAmazonAdsOutput
	ToSourceAmazonAdsOutputWithContext(ctx context.Context) SourceAmazonAdsOutput
}

func (*SourceAmazonAds) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmazonAds)(nil)).Elem()
}

func (i *SourceAmazonAds) ToSourceAmazonAdsOutput() SourceAmazonAdsOutput {
	return i.ToSourceAmazonAdsOutputWithContext(context.Background())
}

func (i *SourceAmazonAds) ToSourceAmazonAdsOutputWithContext(ctx context.Context) SourceAmazonAdsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmazonAdsOutput)
}

// SourceAmazonAdsArrayInput is an input type that accepts SourceAmazonAdsArray and SourceAmazonAdsArrayOutput values.
// You can construct a concrete instance of `SourceAmazonAdsArrayInput` via:
//
//	SourceAmazonAdsArray{ SourceAmazonAdsArgs{...} }
type SourceAmazonAdsArrayInput interface {
	pulumi.Input

	ToSourceAmazonAdsArrayOutput() SourceAmazonAdsArrayOutput
	ToSourceAmazonAdsArrayOutputWithContext(context.Context) SourceAmazonAdsArrayOutput
}

type SourceAmazonAdsArray []SourceAmazonAdsInput

func (SourceAmazonAdsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAmazonAds)(nil)).Elem()
}

func (i SourceAmazonAdsArray) ToSourceAmazonAdsArrayOutput() SourceAmazonAdsArrayOutput {
	return i.ToSourceAmazonAdsArrayOutputWithContext(context.Background())
}

func (i SourceAmazonAdsArray) ToSourceAmazonAdsArrayOutputWithContext(ctx context.Context) SourceAmazonAdsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmazonAdsArrayOutput)
}

// SourceAmazonAdsMapInput is an input type that accepts SourceAmazonAdsMap and SourceAmazonAdsMapOutput values.
// You can construct a concrete instance of `SourceAmazonAdsMapInput` via:
//
//	SourceAmazonAdsMap{ "key": SourceAmazonAdsArgs{...} }
type SourceAmazonAdsMapInput interface {
	pulumi.Input

	ToSourceAmazonAdsMapOutput() SourceAmazonAdsMapOutput
	ToSourceAmazonAdsMapOutputWithContext(context.Context) SourceAmazonAdsMapOutput
}

type SourceAmazonAdsMap map[string]SourceAmazonAdsInput

func (SourceAmazonAdsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAmazonAds)(nil)).Elem()
}

func (i SourceAmazonAdsMap) ToSourceAmazonAdsMapOutput() SourceAmazonAdsMapOutput {
	return i.ToSourceAmazonAdsMapOutputWithContext(context.Background())
}

func (i SourceAmazonAdsMap) ToSourceAmazonAdsMapOutputWithContext(ctx context.Context) SourceAmazonAdsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmazonAdsMapOutput)
}

type SourceAmazonAdsOutput struct{ *pulumi.OutputState }

func (SourceAmazonAdsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmazonAds)(nil)).Elem()
}

func (o SourceAmazonAdsOutput) ToSourceAmazonAdsOutput() SourceAmazonAdsOutput {
	return o
}

func (o SourceAmazonAdsOutput) ToSourceAmazonAdsOutputWithContext(ctx context.Context) SourceAmazonAdsOutput {
	return o
}

func (o SourceAmazonAdsOutput) Configuration() SourceAmazonAdsConfigurationOutput {
	return o.ApplyT(func(v *SourceAmazonAds) SourceAmazonAdsConfigurationOutput { return v.Configuration }).(SourceAmazonAdsConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
// replacement if changed.
func (o SourceAmazonAdsOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAmazonAds) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceAmazonAdsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonAds) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceAmazonAdsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAmazonAds) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAmazonAdsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonAds) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAmazonAdsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonAds) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAmazonAdsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonAds) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceAmazonAdsArrayOutput struct{ *pulumi.OutputState }

func (SourceAmazonAdsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAmazonAds)(nil)).Elem()
}

func (o SourceAmazonAdsArrayOutput) ToSourceAmazonAdsArrayOutput() SourceAmazonAdsArrayOutput {
	return o
}

func (o SourceAmazonAdsArrayOutput) ToSourceAmazonAdsArrayOutputWithContext(ctx context.Context) SourceAmazonAdsArrayOutput {
	return o
}

func (o SourceAmazonAdsArrayOutput) Index(i pulumi.IntInput) SourceAmazonAdsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceAmazonAds {
		return vs[0].([]*SourceAmazonAds)[vs[1].(int)]
	}).(SourceAmazonAdsOutput)
}

type SourceAmazonAdsMapOutput struct{ *pulumi.OutputState }

func (SourceAmazonAdsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAmazonAds)(nil)).Elem()
}

func (o SourceAmazonAdsMapOutput) ToSourceAmazonAdsMapOutput() SourceAmazonAdsMapOutput {
	return o
}

func (o SourceAmazonAdsMapOutput) ToSourceAmazonAdsMapOutputWithContext(ctx context.Context) SourceAmazonAdsMapOutput {
	return o
}

func (o SourceAmazonAdsMapOutput) MapIndex(k pulumi.StringInput) SourceAmazonAdsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceAmazonAds {
		return vs[0].(map[string]*SourceAmazonAds)[vs[1].(string)]
	}).(SourceAmazonAdsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmazonAdsInput)(nil)).Elem(), &SourceAmazonAds{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmazonAdsArrayInput)(nil)).Elem(), SourceAmazonAdsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmazonAdsMapInput)(nil)).Elem(), SourceAmazonAdsMap{})
	pulumi.RegisterOutputType(SourceAmazonAdsOutput{})
	pulumi.RegisterOutputType(SourceAmazonAdsArrayOutput{})
	pulumi.RegisterOutputType(SourceAmazonAdsMapOutput{})
}