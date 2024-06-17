// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceIp2whois(ctx *pulumi.Context, args *LookupSourceIp2whoisArgs, opts ...pulumi.InvokeOption) (*LookupSourceIp2whoisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceIp2whoisResult
	err := ctx.Invoke("airbyte:index/getSourceIp2whois:getSourceIp2whois", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceIp2whois.
type LookupSourceIp2whoisArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceIp2whois.
type LookupSourceIp2whoisResult struct {
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceIp2whoisOutput(ctx *pulumi.Context, args LookupSourceIp2whoisOutputArgs, opts ...pulumi.InvokeOption) LookupSourceIp2whoisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceIp2whoisResult, error) {
			args := v.(LookupSourceIp2whoisArgs)
			r, err := LookupSourceIp2whois(ctx, &args, opts...)
			var s LookupSourceIp2whoisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceIp2whoisResultOutput)
}

// A collection of arguments for invoking getSourceIp2whois.
type LookupSourceIp2whoisOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceIp2whoisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceIp2whoisArgs)(nil)).Elem()
}

// A collection of values returned by getSourceIp2whois.
type LookupSourceIp2whoisResultOutput struct{ *pulumi.OutputState }

func (LookupSourceIp2whoisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceIp2whoisResult)(nil)).Elem()
}

func (o LookupSourceIp2whoisResultOutput) ToLookupSourceIp2whoisResultOutput() LookupSourceIp2whoisResultOutput {
	return o
}

func (o LookupSourceIp2whoisResultOutput) ToLookupSourceIp2whoisResultOutputWithContext(ctx context.Context) LookupSourceIp2whoisResultOutput {
	return o
}

func (o LookupSourceIp2whoisResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIp2whoisResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceIp2whoisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIp2whoisResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceIp2whoisResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIp2whoisResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceIp2whoisResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIp2whoisResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceIp2whoisResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIp2whoisResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceIp2whoisResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIp2whoisResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceIp2whoisResultOutput{})
}