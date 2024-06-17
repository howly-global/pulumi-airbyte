// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceWebflow
    {
        public static Task<GetSourceWebflowResult> InvokeAsync(GetSourceWebflowArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceWebflowResult>("airbyte:index/getSourceWebflow:getSourceWebflow", args ?? new GetSourceWebflowArgs(), options.WithDefaults());

        public static Output<GetSourceWebflowResult> Invoke(GetSourceWebflowInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceWebflowResult>("airbyte:index/getSourceWebflow:getSourceWebflow", args ?? new GetSourceWebflowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceWebflowArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceWebflowArgs()
        {
        }
        public static new GetSourceWebflowArgs Empty => new GetSourceWebflowArgs();
    }

    public sealed class GetSourceWebflowInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceWebflowInvokeArgs()
        {
        }
        public static new GetSourceWebflowInvokeArgs Empty => new GetSourceWebflowInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceWebflowResult
    {
        public readonly string Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string SourceId;
        public readonly string SourceType;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceWebflowResult(
            string configuration,

            string id,

            string name,

            string sourceId,

            string sourceType,

            string workspaceId)
        {
            Configuration = configuration;
            Id = id;
            Name = name;
            SourceId = sourceId;
            SourceType = sourceType;
            WorkspaceId = workspaceId;
        }
    }
}