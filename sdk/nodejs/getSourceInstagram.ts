// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSourceInstagram(args: GetSourceInstagramArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceInstagramResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceInstagram:getSourceInstagram", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceInstagram.
 */
export interface GetSourceInstagramArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceInstagram.
 */
export interface GetSourceInstagramResult {
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
export function getSourceInstagramOutput(args: GetSourceInstagramOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceInstagramResult> {
    return pulumi.output(args).apply((a: any) => getSourceInstagram(a, opts))
}

/**
 * A collection of arguments for invoking getSourceInstagram.
 */
export interface GetSourceInstagramOutputArgs {
    sourceId: pulumi.Input<string>;
}