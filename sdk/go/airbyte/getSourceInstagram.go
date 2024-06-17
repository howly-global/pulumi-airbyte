// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceInstagram(ctx *pulumi.Context, args *LookupSourceInstagramArgs, opts ...pulumi.InvokeOption) (*LookupSourceInstagramResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceInstagramResult
	err := ctx.Invoke("airbyte:index/getSourceInstagram:getSourceInstagram", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceInstagram.
type LookupSourceInstagramArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceInstagram.
type LookupSourceInstagramResult struct {
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceInstagramOutput(ctx *pulumi.Context, args LookupSourceInstagramOutputArgs, opts ...pulumi.InvokeOption) LookupSourceInstagramResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceInstagramResult, error) {
			args := v.(LookupSourceInstagramArgs)
			r, err := LookupSourceInstagram(ctx, &args, opts...)
			var s LookupSourceInstagramResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceInstagramResultOutput)
}

// A collection of arguments for invoking getSourceInstagram.
type LookupSourceInstagramOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceInstagramOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceInstagramArgs)(nil)).Elem()
}

// A collection of values returned by getSourceInstagram.
type LookupSourceInstagramResultOutput struct{ *pulumi.OutputState }

func (LookupSourceInstagramResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceInstagramResult)(nil)).Elem()
}

func (o LookupSourceInstagramResultOutput) ToLookupSourceInstagramResultOutput() LookupSourceInstagramResultOutput {
	return o
}

func (o LookupSourceInstagramResultOutput) ToLookupSourceInstagramResultOutputWithContext(ctx context.Context) LookupSourceInstagramResultOutput {
	return o
}

func (o LookupSourceInstagramResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstagramResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceInstagramResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstagramResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceInstagramResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstagramResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceInstagramResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstagramResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceInstagramResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstagramResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceInstagramResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstagramResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceInstagramResultOutput{})
}