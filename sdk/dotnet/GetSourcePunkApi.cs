// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourcePunkApi
    {
        public static Task<GetSourcePunkApiResult> InvokeAsync(GetSourcePunkApiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourcePunkApiResult>("airbyte:index/getSourcePunkApi:getSourcePunkApi", args ?? new GetSourcePunkApiArgs(), options.WithDefaults());

        public static Output<GetSourcePunkApiResult> Invoke(GetSourcePunkApiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourcePunkApiResult>("airbyte:index/getSourcePunkApi:getSourcePunkApi", args ?? new GetSourcePunkApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourcePunkApiArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourcePunkApiArgs()
        {
        }
        public static new GetSourcePunkApiArgs Empty => new GetSourcePunkApiArgs();
    }

    public sealed class GetSourcePunkApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourcePunkApiInvokeArgs()
        {
        }
        public static new GetSourcePunkApiInvokeArgs Empty => new GetSourcePunkApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourcePunkApiResult
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
        private GetSourcePunkApiResult(
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