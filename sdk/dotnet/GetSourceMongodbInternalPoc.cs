// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceMongodbInternalPoc
    {
        public static Task<GetSourceMongodbInternalPocResult> InvokeAsync(GetSourceMongodbInternalPocArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceMongodbInternalPocResult>("airbyte:index/getSourceMongodbInternalPoc:getSourceMongodbInternalPoc", args ?? new GetSourceMongodbInternalPocArgs(), options.WithDefaults());

        public static Output<GetSourceMongodbInternalPocResult> Invoke(GetSourceMongodbInternalPocInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceMongodbInternalPocResult>("airbyte:index/getSourceMongodbInternalPoc:getSourceMongodbInternalPoc", args ?? new GetSourceMongodbInternalPocInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceMongodbInternalPocArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceMongodbInternalPocArgs()
        {
        }
        public static new GetSourceMongodbInternalPocArgs Empty => new GetSourceMongodbInternalPocArgs();
    }

    public sealed class GetSourceMongodbInternalPocInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceMongodbInternalPocInvokeArgs()
        {
        }
        public static new GetSourceMongodbInternalPocInvokeArgs Empty => new GetSourceMongodbInternalPocInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceMongodbInternalPocResult
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
        private GetSourceMongodbInternalPocResult(
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
