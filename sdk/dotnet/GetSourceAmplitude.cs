// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceAmplitude
    {
        public static Task<GetSourceAmplitudeResult> InvokeAsync(GetSourceAmplitudeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceAmplitudeResult>("airbyte:index/getSourceAmplitude:getSourceAmplitude", args ?? new GetSourceAmplitudeArgs(), options.WithDefaults());

        public static Output<GetSourceAmplitudeResult> Invoke(GetSourceAmplitudeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceAmplitudeResult>("airbyte:index/getSourceAmplitude:getSourceAmplitude", args ?? new GetSourceAmplitudeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceAmplitudeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceAmplitudeArgs()
        {
        }
        public static new GetSourceAmplitudeArgs Empty => new GetSourceAmplitudeArgs();
    }

    public sealed class GetSourceAmplitudeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceAmplitudeInvokeArgs()
        {
        }
        public static new GetSourceAmplitudeInvokeArgs Empty => new GetSourceAmplitudeInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceAmplitudeResult
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
        private GetSourceAmplitudeResult(
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