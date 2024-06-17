// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceRss(ctx *pulumi.Context, args *LookupSourceRssArgs, opts ...pulumi.InvokeOption) (*LookupSourceRssResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceRssResult
	err := ctx.Invoke("airbyte:index/getSourceRss:getSourceRss", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceRss.
type LookupSourceRssArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceRss.
type LookupSourceRssResult struct {
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceRssOutput(ctx *pulumi.Context, args LookupSourceRssOutputArgs, opts ...pulumi.InvokeOption) LookupSourceRssResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceRssResult, error) {
			args := v.(LookupSourceRssArgs)
			r, err := LookupSourceRss(ctx, &args, opts...)
			var s LookupSourceRssResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceRssResultOutput)
}

// A collection of arguments for invoking getSourceRss.
type LookupSourceRssOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceRssOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRssArgs)(nil)).Elem()
}

// A collection of values returned by getSourceRss.
type LookupSourceRssResultOutput struct{ *pulumi.OutputState }

func (LookupSourceRssResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRssResult)(nil)).Elem()
}

func (o LookupSourceRssResultOutput) ToLookupSourceRssResultOutput() LookupSourceRssResultOutput {
	return o
}

func (o LookupSourceRssResultOutput) ToLookupSourceRssResultOutputWithContext(ctx context.Context) LookupSourceRssResultOutput {
	return o
}

func (o LookupSourceRssResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceRssResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceRssResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceRssResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceRssResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceRssResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceRssResultOutput{})
}