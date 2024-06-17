// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceZohoCrm(ctx *pulumi.Context, args *LookupSourceZohoCrmArgs, opts ...pulumi.InvokeOption) (*LookupSourceZohoCrmResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceZohoCrmResult
	err := ctx.Invoke("airbyte:index/getSourceZohoCrm:getSourceZohoCrm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceZohoCrm.
type LookupSourceZohoCrmArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceZohoCrm.
type LookupSourceZohoCrmResult struct {
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceZohoCrmOutput(ctx *pulumi.Context, args LookupSourceZohoCrmOutputArgs, opts ...pulumi.InvokeOption) LookupSourceZohoCrmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceZohoCrmResult, error) {
			args := v.(LookupSourceZohoCrmArgs)
			r, err := LookupSourceZohoCrm(ctx, &args, opts...)
			var s LookupSourceZohoCrmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceZohoCrmResultOutput)
}

// A collection of arguments for invoking getSourceZohoCrm.
type LookupSourceZohoCrmOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceZohoCrmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZohoCrmArgs)(nil)).Elem()
}

// A collection of values returned by getSourceZohoCrm.
type LookupSourceZohoCrmResultOutput struct{ *pulumi.OutputState }

func (LookupSourceZohoCrmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZohoCrmResult)(nil)).Elem()
}

func (o LookupSourceZohoCrmResultOutput) ToLookupSourceZohoCrmResultOutput() LookupSourceZohoCrmResultOutput {
	return o
}

func (o LookupSourceZohoCrmResultOutput) ToLookupSourceZohoCrmResultOutputWithContext(ctx context.Context) LookupSourceZohoCrmResultOutput {
	return o
}

func (o LookupSourceZohoCrmResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceZohoCrmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceZohoCrmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceZohoCrmResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceZohoCrmResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceZohoCrmResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceZohoCrmResultOutput{})
}