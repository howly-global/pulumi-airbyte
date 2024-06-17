// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSourceSnowflake(args: GetSourceSnowflakeArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceSnowflakeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceSnowflake:getSourceSnowflake", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceSnowflake.
 */
export interface GetSourceSnowflakeArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceSnowflake.
 */
export interface GetSourceSnowflakeResult {
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
export function getSourceSnowflakeOutput(args: GetSourceSnowflakeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceSnowflakeResult> {
    return pulumi.output(args).apply((a: any) => getSourceSnowflake(a, opts))
}

/**
 * A collection of arguments for invoking getSourceSnowflake.
 */
export interface GetSourceSnowflakeOutputArgs {
    sourceId: pulumi.Input<string>;
}