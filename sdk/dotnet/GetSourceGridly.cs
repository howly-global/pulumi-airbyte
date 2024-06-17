// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceGridly
    {
        public static Task<GetSourceGridlyResult> InvokeAsync(GetSourceGridlyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceGridlyResult>("airbyte:index/getSourceGridly:getSourceGridly", args ?? new GetSourceGridlyArgs(), options.WithDefaults());

        public static Output<GetSourceGridlyResult> Invoke(GetSourceGridlyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceGridlyResult>("airbyte:index/getSourceGridly:getSourceGridly", args ?? new GetSourceGridlyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceGridlyArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceGridlyArgs()
        {
        }
        public static new GetSourceGridlyArgs Empty => new GetSourceGridlyArgs();
    }

    public sealed class GetSourceGridlyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceGridlyInvokeArgs()
        {
        }
        public static new GetSourceGridlyInvokeArgs Empty => new GetSourceGridlyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceGridlyResult
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
        private GetSourceGridlyResult(
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