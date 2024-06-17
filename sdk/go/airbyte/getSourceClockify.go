// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceClockify(ctx *pulumi.Context, args *LookupSourceClockifyArgs, opts ...pulumi.InvokeOption) (*LookupSourceClockifyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceClockifyResult
	err := ctx.Invoke("airbyte:index/getSourceClockify:getSourceClockify", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceClockify.
type LookupSourceClockifyArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceClockify.
type LookupSourceClockifyResult struct {
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceClockifyOutput(ctx *pulumi.Context, args LookupSourceClockifyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceClockifyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceClockifyResult, error) {
			args := v.(LookupSourceClockifyArgs)
			r, err := LookupSourceClockify(ctx, &args, opts...)
			var s LookupSourceClockifyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceClockifyResultOutput)
}

// A collection of arguments for invoking getSourceClockify.
type LookupSourceClockifyOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceClockifyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceClockifyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceClockify.
type LookupSourceClockifyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceClockifyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceClockifyResult)(nil)).Elem()
}

func (o LookupSourceClockifyResultOutput) ToLookupSourceClockifyResultOutput() LookupSourceClockifyResultOutput {
	return o
}

func (o LookupSourceClockifyResultOutput) ToLookupSourceClockifyResultOutputWithContext(ctx context.Context) LookupSourceClockifyResultOutput {
	return o
}

func (o LookupSourceClockifyResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClockifyResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceClockifyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClockifyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceClockifyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClockifyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceClockifyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClockifyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceClockifyResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClockifyResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceClockifyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClockifyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceClockifyResultOutput{})
}