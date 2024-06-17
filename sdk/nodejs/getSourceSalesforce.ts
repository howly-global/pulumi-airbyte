// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSourceSalesforce(args: GetSourceSalesforceArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceSalesforceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceSalesforce:getSourceSalesforce", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceSalesforce.
 */
export interface GetSourceSalesforceArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceSalesforce.
 */
export interface GetSourceSalesforceResult {
    readonly configuration: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly sourceId: string;
    readonly sourceType: string;
    readonly workspaceId: string;
}
export function getSourceSalesforceOutput(args: GetSourceSalesforceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceSalesforceResult> {
    return pulumi.output(args).apply((a: any) => getSourceSalesforce(a, opts))
}

/**
 * A collection of arguments for invoking getSourceSalesforce.
 */
export interface GetSourceSalesforceOutputArgs {
    sourceId: pulumi.Input<string>;
}