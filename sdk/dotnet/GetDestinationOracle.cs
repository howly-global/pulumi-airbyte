// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationOracle
    {
        public static Task<GetDestinationOracleResult> InvokeAsync(GetDestinationOracleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationOracleResult>("airbyte:index/getDestinationOracle:getDestinationOracle", args ?? new GetDestinationOracleArgs(), options.WithDefaults());

        public static Output<GetDestinationOracleResult> Invoke(GetDestinationOracleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationOracleResult>("airbyte:index/getDestinationOracle:getDestinationOracle", args ?? new GetDestinationOracleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationOracleArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationOracleArgs()
        {
        }
        public static new GetDestinationOracleArgs Empty => new GetDestinationOracleArgs();
    }

    public sealed class GetDestinationOracleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationOracleInvokeArgs()
        {
        }
        public static new GetDestinationOracleInvokeArgs Empty => new GetDestinationOracleInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationOracleResult
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
        private GetDestinationOracleResult(
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