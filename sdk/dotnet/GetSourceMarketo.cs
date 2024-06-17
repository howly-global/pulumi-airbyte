// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceMarketo
    {
        public static Task<GetSourceMarketoResult> InvokeAsync(GetSourceMarketoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceMarketoResult>("airbyte:index/getSourceMarketo:getSourceMarketo", args ?? new GetSourceMarketoArgs(), options.WithDefaults());

        public static Output<GetSourceMarketoResult> Invoke(GetSourceMarketoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceMarketoResult>("airbyte:index/getSourceMarketo:getSourceMarketo", args ?? new GetSourceMarketoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceMarketoArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceMarketoArgs()
        {
        }
        public static new GetSourceMarketoArgs Empty => new GetSourceMarketoArgs();
    }

    public sealed class GetSourceMarketoInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceMarketoInvokeArgs()
        {
        }
        public static new GetSourceMarketoInvokeArgs Empty => new GetSourceMarketoInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceMarketoResult
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
        private GetSourceMarketoResult(
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