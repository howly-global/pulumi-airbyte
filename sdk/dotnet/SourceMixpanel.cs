// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    [AirbyteResourceType("airbyte:index/sourceMixpanel:SourceMixpanel")]
    public partial class SourceMixpanel : global::Pulumi.CustomResource
    {
        [Output("configuration")]
        public Output<Outputs.SourceMixpanelConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
        /// replacement if changed.
        /// </summary>
        [Output("definitionId")]
        public Output<string?> DefinitionId { get; private set; } = null!;

        /// <summary>
        /// Name of the source e.g. dev-mysql-instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
        /// </summary>
        [Output("secretId")]
        public Output<string?> SecretId { get; private set; } = null!;

        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;

        [Output("sourceType")]
        public Output<string> SourceType { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a SourceMixpanel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceMixpanel(string name, SourceMixpanelArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceMixpanel:SourceMixpanel", name, args ?? new SourceMixpanelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceMixpanel(string name, Input<string> id, SourceMixpanelState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceMixpanel:SourceMixpanel", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SourceMixpanel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceMixpanel Get(string name, Input<string> id, SourceMixpanelState? state = null, CustomResourceOptions? options = null)
        {
            return new SourceMixpanel(name, id, state, options);
        }
    }

    public sealed class SourceMixpanelArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration", required: true)]
        public Input<Inputs.SourceMixpanelConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
        /// replacement if changed.
        /// </summary>
        [Input("definitionId")]
        public Input<string>? DefinitionId { get; set; }

        /// <summary>
        /// Name of the source e.g. dev-mysql-instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public SourceMixpanelArgs()
        {
        }
        public static new SourceMixpanelArgs Empty => new SourceMixpanelArgs();
    }

    public sealed class SourceMixpanelState : global::Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<Inputs.SourceMixpanelConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires
        /// replacement if changed.
        /// </summary>
        [Input("definitionId")]
        public Input<string>? DefinitionId { get; set; }

        /// <summary>
        /// Name of the source e.g. dev-mysql-instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public SourceMixpanelState()
        {
        }
        public static new SourceMixpanelState Empty => new SourceMixpanelState();
    }
}