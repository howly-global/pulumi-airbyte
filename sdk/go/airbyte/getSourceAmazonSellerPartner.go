// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/howly-global/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceAmazonSellerPartner(ctx *pulumi.Context, args *LookupSourceAmazonSellerPartnerArgs, opts ...pulumi.InvokeOption) (*LookupSourceAmazonSellerPartnerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceAmazonSellerPartnerResult
	err := ctx.Invoke("airbyte:index/getSourceAmazonSellerPartner:getSourceAmazonSellerPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAmazonSellerPartner.
type LookupSourceAmazonSellerPartnerArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAmazonSellerPartner.
type LookupSourceAmazonSellerPartnerResult struct {
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceAmazonSellerPartnerOutput(ctx *pulumi.Context, args LookupSourceAmazonSellerPartnerOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAmazonSellerPartnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAmazonSellerPartnerResult, error) {
			args := v.(LookupSourceAmazonSellerPartnerArgs)
			r, err := LookupSourceAmazonSellerPartner(ctx, &args, opts...)
			var s LookupSourceAmazonSellerPartnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAmazonSellerPartnerResultOutput)
}

// A collection of arguments for invoking getSourceAmazonSellerPartner.
type LookupSourceAmazonSellerPartnerOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceAmazonSellerPartnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmazonSellerPartnerArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAmazonSellerPartner.
type LookupSourceAmazonSellerPartnerResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAmazonSellerPartnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmazonSellerPartnerResult)(nil)).Elem()
}

func (o LookupSourceAmazonSellerPartnerResultOutput) ToLookupSourceAmazonSellerPartnerResultOutput() LookupSourceAmazonSellerPartnerResultOutput {
	return o
}

func (o LookupSourceAmazonSellerPartnerResultOutput) ToLookupSourceAmazonSellerPartnerResultOutputWithContext(ctx context.Context) LookupSourceAmazonSellerPartnerResultOutput {
	return o
}

func (o LookupSourceAmazonSellerPartnerResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSellerPartnerResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAmazonSellerPartnerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSellerPartnerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonSellerPartnerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSellerPartnerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonSellerPartnerResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSellerPartnerResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonSellerPartnerResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSellerPartnerResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonSellerPartnerResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSellerPartnerResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAmazonSellerPartnerResultOutput{})
}