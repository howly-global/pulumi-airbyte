// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceDynamodb
    {
        public static Task<GetSourceDynamodbResult> InvokeAsync(GetSourceDynamodbArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceDynamodbResult>("airbyte:index/getSourceDynamodb:getSourceDynamodb", args ?? new GetSourceDynamodbArgs(), options.WithDefaults());

        public static Output<GetSourceDynamodbResult> Invoke(GetSourceDynamodbInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceDynamodbResult>("airbyte:index/getSourceDynamodb:getSourceDynamodb", args ?? new GetSourceDynamodbInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceDynamodbArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceDynamodbArgs()
        {
        }
        public static new GetSourceDynamodbArgs Empty => new GetSourceDynamodbArgs();
    }

    public sealed class GetSourceDynamodbInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceDynamodbInvokeArgs()
        {
        }
        public static new GetSourceDynamodbInvokeArgs Empty => new GetSourceDynamodbInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceDynamodbResult
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
        private GetSourceDynamodbResult(
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