// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSourcePinterest(args: GetSourcePinterestArgs, opts?: pulumi.InvokeOptions): Promise<GetSourcePinterestResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourcePinterest:getSourcePinterest", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourcePinterest.
 */
export interface GetSourcePinterestArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourcePinterest.
 */
export interface GetSourcePinterestResult {
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
export function getSourcePinterestOutput(args: GetSourcePinterestOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourcePinterestResult> {
    return pulumi.output(args).apply((a: any) => getSourcePinterest(a, opts))
}

/**
 * A collection of arguments for invoking getSourcePinterest.
 */
export interface GetSourcePinterestOutputArgs {
    sourceId: pulumi.Input<string>;
}