// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceSnowflake(ctx *pulumi.Context, args *LookupSourceSnowflakeArgs, opts ...pulumi.InvokeOption) (*LookupSourceSnowflakeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSnowflakeResult
	err := ctx.Invoke("airbyte:index/getSourceSnowflake:getSourceSnowflake", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSnowflake.
type LookupSourceSnowflakeArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSnowflake.
type LookupSourceSnowflakeResult struct {
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceSnowflakeOutput(ctx *pulumi.Context, args LookupSourceSnowflakeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSnowflakeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSnowflakeResult, error) {
			args := v.(LookupSourceSnowflakeArgs)
			r, err := LookupSourceSnowflake(ctx, &args, opts...)
			var s LookupSourceSnowflakeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSnowflakeResultOutput)
}

// A collection of arguments for invoking getSourceSnowflake.
type LookupSourceSnowflakeOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceSnowflakeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSnowflakeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSnowflake.
type LookupSourceSnowflakeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSnowflakeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSnowflakeResult)(nil)).Elem()
}

func (o LookupSourceSnowflakeResultOutput) ToLookupSourceSnowflakeResultOutput() LookupSourceSnowflakeResultOutput {
	return o
}

func (o LookupSourceSnowflakeResultOutput) ToLookupSourceSnowflakeResultOutputWithContext(ctx context.Context) LookupSourceSnowflakeResultOutput {
	return o
}

func (o LookupSourceSnowflakeResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnowflakeResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSnowflakeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnowflakeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSnowflakeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnowflakeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSnowflakeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnowflakeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSnowflakeResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnowflakeResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceSnowflakeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnowflakeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSnowflakeResultOutput{})
}