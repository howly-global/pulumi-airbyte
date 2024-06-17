// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceMicrosoftTeams
    {
        public static Task<GetSourceMicrosoftTeamsResult> InvokeAsync(GetSourceMicrosoftTeamsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceMicrosoftTeamsResult>("airbyte:index/getSourceMicrosoftTeams:getSourceMicrosoftTeams", args ?? new GetSourceMicrosoftTeamsArgs(), options.WithDefaults());

        public static Output<GetSourceMicrosoftTeamsResult> Invoke(GetSourceMicrosoftTeamsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceMicrosoftTeamsResult>("airbyte:index/getSourceMicrosoftTeams:getSourceMicrosoftTeams", args ?? new GetSourceMicrosoftTeamsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceMicrosoftTeamsArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceMicrosoftTeamsArgs()
        {
        }
        public static new GetSourceMicrosoftTeamsArgs Empty => new GetSourceMicrosoftTeamsArgs();
    }

    public sealed class GetSourceMicrosoftTeamsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceMicrosoftTeamsInvokeArgs()
        {
        }
        public static new GetSourceMicrosoftTeamsInvokeArgs Empty => new GetSourceMicrosoftTeamsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceMicrosoftTeamsResult
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
        private GetSourceMicrosoftTeamsResult(
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