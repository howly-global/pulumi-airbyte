// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationSnowflake
    {
        public static Task<GetDestinationSnowflakeResult> InvokeAsync(GetDestinationSnowflakeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationSnowflakeResult>("airbyte:index/getDestinationSnowflake:getDestinationSnowflake", args ?? new GetDestinationSnowflakeArgs(), options.WithDefaults());

        public static Output<GetDestinationSnowflakeResult> Invoke(GetDestinationSnowflakeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationSnowflakeResult>("airbyte:index/getDestinationSnowflake:getDestinationSnowflake", args ?? new GetDestinationSnowflakeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationSnowflakeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationSnowflakeArgs()
        {
        }
        public static new GetDestinationSnowflakeArgs Empty => new GetDestinationSnowflakeArgs();
    }

    public sealed class GetDestinationSnowflakeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationSnowflakeInvokeArgs()
        {
        }
        public static new GetDestinationSnowflakeInvokeArgs Empty => new GetDestinationSnowflakeInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationSnowflakeResult
    {
        public readonly string Configuration;
        public readonly string DestinationId;
        public readonly string DestinationType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetDestinationSnowflakeResult(
            string configuration,

            string destinationId,

            string destinationType,

            string id,

            string name,

            string workspaceId)
        {
            Configuration = configuration;
            DestinationId = destinationId;
            DestinationType = destinationType;
            Id = id;
            Name = name;
            WorkspaceId = workspaceId;
        }
    }
}