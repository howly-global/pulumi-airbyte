// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourcePostgres
    {
        public static Task<GetSourcePostgresResult> InvokeAsync(GetSourcePostgresArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourcePostgresResult>("airbyte:index/getSourcePostgres:getSourcePostgres", args ?? new GetSourcePostgresArgs(), options.WithDefaults());

        public static Output<GetSourcePostgresResult> Invoke(GetSourcePostgresInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourcePostgresResult>("airbyte:index/getSourcePostgres:getSourcePostgres", args ?? new GetSourcePostgresInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourcePostgresArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourcePostgresArgs()
        {
        }
        public static new GetSourcePostgresArgs Empty => new GetSourcePostgresArgs();
    }

    public sealed class GetSourcePostgresInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourcePostgresInvokeArgs()
        {
        }
        public static new GetSourcePostgresInvokeArgs Empty => new GetSourcePostgresInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourcePostgresResult
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
        private GetSourcePostgresResult(
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