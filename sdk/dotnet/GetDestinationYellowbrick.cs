// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationYellowbrick
    {
        public static Task<GetDestinationYellowbrickResult> InvokeAsync(GetDestinationYellowbrickArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationYellowbrickResult>("airbyte:index/getDestinationYellowbrick:getDestinationYellowbrick", args ?? new GetDestinationYellowbrickArgs(), options.WithDefaults());

        public static Output<GetDestinationYellowbrickResult> Invoke(GetDestinationYellowbrickInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationYellowbrickResult>("airbyte:index/getDestinationYellowbrick:getDestinationYellowbrick", args ?? new GetDestinationYellowbrickInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationYellowbrickArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationYellowbrickArgs()
        {
        }
        public static new GetDestinationYellowbrickArgs Empty => new GetDestinationYellowbrickArgs();
    }

    public sealed class GetDestinationYellowbrickInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationYellowbrickInvokeArgs()
        {
        }
        public static new GetDestinationYellowbrickInvokeArgs Empty => new GetDestinationYellowbrickInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationYellowbrickResult
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
        private GetDestinationYellowbrickResult(
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