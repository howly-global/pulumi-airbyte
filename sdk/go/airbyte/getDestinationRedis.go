// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDestinationRedis(ctx *pulumi.Context, args *LookupDestinationRedisArgs, opts ...pulumi.InvokeOption) (*LookupDestinationRedisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationRedisResult
	err := ctx.Invoke("airbyte:index/getDestinationRedis:getDestinationRedis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationRedis.
type LookupDestinationRedisArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationRedis.
type LookupDestinationRedisResult struct {
	Configuration   string `pulumi:"configuration"`
	DestinationId   string `pulumi:"destinationId"`
	DestinationType string `pulumi:"destinationType"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationRedisOutput(ctx *pulumi.Context, args LookupDestinationRedisOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationRedisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationRedisResult, error) {
			args := v.(LookupDestinationRedisArgs)
			r, err := LookupDestinationRedis(ctx, &args, opts...)
			var s LookupDestinationRedisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationRedisResultOutput)
}

// A collection of arguments for invoking getDestinationRedis.
type LookupDestinationRedisOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationRedisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationRedisArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationRedis.
type LookupDestinationRedisResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationRedisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationRedisResult)(nil)).Elem()
}

func (o LookupDestinationRedisResultOutput) ToLookupDestinationRedisResultOutput() LookupDestinationRedisResultOutput {
	return o
}

func (o LookupDestinationRedisResultOutput) ToLookupDestinationRedisResultOutputWithContext(ctx context.Context) LookupDestinationRedisResultOutput {
	return o
}

func (o LookupDestinationRedisResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedisResult) string { return v.Configuration }).(pulumi.StringOutput)
}

func (o LookupDestinationRedisResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedisResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

func (o LookupDestinationRedisResultOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedisResult) string { return v.DestinationType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationRedisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedisResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationRedisResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedisResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationRedisResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedisResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationRedisResultOutput{})
}