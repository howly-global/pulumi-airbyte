// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSourceMetabase(args: GetSourceMetabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceMetabaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceMetabase:getSourceMetabase", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceMetabase.
 */
export interface GetSourceMetabaseArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceMetabase.
 */
export interface GetSourceMetabaseResult {
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
export function getSourceMetabaseOutput(args: GetSourceMetabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceMetabaseResult> {
    return pulumi.output(args).apply((a: any) => getSourceMetabase(a, opts))
}

/**
 * A collection of arguments for invoking getSourceMetabase.
 */
export interface GetSourceMetabaseOutputArgs {
    sourceId: pulumi.Input<string>;
}