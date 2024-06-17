// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceSonarCloud
    {
        public static Task<GetSourceSonarCloudResult> InvokeAsync(GetSourceSonarCloudArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceSonarCloudResult>("airbyte:index/getSourceSonarCloud:getSourceSonarCloud", args ?? new GetSourceSonarCloudArgs(), options.WithDefaults());

        public static Output<GetSourceSonarCloudResult> Invoke(GetSourceSonarCloudInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceSonarCloudResult>("airbyte:index/getSourceSonarCloud:getSourceSonarCloud", args ?? new GetSourceSonarCloudInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceSonarCloudArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceSonarCloudArgs()
        {
        }
        public static new GetSourceSonarCloudArgs Empty => new GetSourceSonarCloudArgs();
    }

    public sealed class GetSourceSonarCloudInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceSonarCloudInvokeArgs()
        {
        }
        public static new GetSourceSonarCloudInvokeArgs Empty => new GetSourceSonarCloudInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceSonarCloudResult
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
        private GetSourceSonarCloudResult(
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