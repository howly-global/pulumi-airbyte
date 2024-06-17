// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceRss
    {
        public static Task<GetSourceRssResult> InvokeAsync(GetSourceRssArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceRssResult>("airbyte:index/getSourceRss:getSourceRss", args ?? new GetSourceRssArgs(), options.WithDefaults());

        public static Output<GetSourceRssResult> Invoke(GetSourceRssInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceRssResult>("airbyte:index/getSourceRss:getSourceRss", args ?? new GetSourceRssInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceRssArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceRssArgs()
        {
        }
        public static new GetSourceRssArgs Empty => new GetSourceRssArgs();
    }

    public sealed class GetSourceRssInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceRssInvokeArgs()
        {
        }
        public static new GetSourceRssInvokeArgs Empty => new GetSourceRssInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceRssResult
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
        private GetSourceRssResult(
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