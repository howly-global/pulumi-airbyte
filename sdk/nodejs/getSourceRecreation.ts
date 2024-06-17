// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSourceRecreation(args: GetSourceRecreationArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceRecreationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceRecreation:getSourceRecreation", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceRecreation.
 */
export interface GetSourceRecreationArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceRecreation.
 */
export interface GetSourceRecreationResult {
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
export function getSourceRecreationOutput(args: GetSourceRecreationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceRecreationResult> {
    return pulumi.output(args).apply((a: any) => getSourceRecreation(a, opts))
}

/**
 * A collection of arguments for invoking getSourceRecreation.
 */
export interface GetSourceRecreationOutputArgs {
    sourceId: pulumi.Input<string>;
}