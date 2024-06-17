// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DestinationS3Glue extends pulumi.CustomResource {
    /**
     * Get an existing DestinationS3Glue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DestinationS3GlueState, opts?: pulumi.CustomResourceOptions): DestinationS3Glue {
        return new DestinationS3Glue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'airbyte:index/destinationS3Glue:DestinationS3Glue';

    /**
     * Returns true if the given object is an instance of DestinationS3Glue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DestinationS3Glue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DestinationS3Glue.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.DestinationS3GlueConfiguration>;
    /**
     * The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
     * replacement if changed.
     */
    public readonly definitionId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly destinationId!: pulumi.Output<string>;
    public /*out*/ readonly destinationType!: pulumi.Output<string>;
    /**
     * Name of the destination e.g. dev-mysql-instance.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a DestinationS3Glue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DestinationS3GlueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DestinationS3GlueArgs | DestinationS3GlueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DestinationS3GlueState | undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["definitionId"] = state ? state.definitionId : undefined;
            resourceInputs["destinationId"] = state ? state.destinationId : undefined;
            resourceInputs["destinationType"] = state ? state.destinationType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as DestinationS3GlueArgs | undefined;
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.workspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceId'");
            }
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["definitionId"] = args ? args.definitionId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["workspaceId"] = args ? args.workspaceId : undefined;
            resourceInputs["destinationId"] = undefined /*out*/;
            resourceInputs["destinationType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DestinationS3Glue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DestinationS3Glue resources.
 */
export interface DestinationS3GlueState {
    configuration?: pulumi.Input<inputs.DestinationS3GlueConfiguration>;
    /**
     * The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
     * replacement if changed.
     */
    definitionId?: pulumi.Input<string>;
    destinationId?: pulumi.Input<string>;
    destinationType?: pulumi.Input<string>;
    /**
     * Name of the destination e.g. dev-mysql-instance.
     */
    name?: pulumi.Input<string>;
    workspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DestinationS3Glue resource.
 */
export interface DestinationS3GlueArgs {
    configuration: pulumi.Input<inputs.DestinationS3GlueConfiguration>;
    /**
     * The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
     * replacement if changed.
     */
    definitionId?: pulumi.Input<string>;
    /**
     * Name of the destination e.g. dev-mysql-instance.
     */
    name?: pulumi.Input<string>;
    workspaceId: pulumi.Input<string>;
}