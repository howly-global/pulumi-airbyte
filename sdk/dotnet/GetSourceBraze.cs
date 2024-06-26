// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceBraze
    {
        public static Task<GetSourceBrazeResult> InvokeAsync(GetSourceBrazeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceBrazeResult>("airbyte:index/getSourceBraze:getSourceBraze", args ?? new GetSourceBrazeArgs(), options.WithDefaults());

        public static Output<GetSourceBrazeResult> Invoke(GetSourceBrazeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceBrazeResult>("airbyte:index/getSourceBraze:getSourceBraze", args ?? new GetSourceBrazeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceBrazeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceBrazeArgs()
        {
        }
        public static new GetSourceBrazeArgs Empty => new GetSourceBrazeArgs();
    }

    public sealed class GetSourceBrazeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceBrazeInvokeArgs()
        {
        }
        public static new GetSourceBrazeInvokeArgs Empty => new GetSourceBrazeInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceBrazeResult
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
        private GetSourceBrazeResult(
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
