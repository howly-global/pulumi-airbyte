// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSftpBulkConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingViaApiParameterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the unstructured API parameter to use
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value of the parameter
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public SourceSftpBulkConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingViaApiParameterGetArgs()
        {
        }
        public static new SourceSftpBulkConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingViaApiParameterGetArgs Empty => new SourceSftpBulkConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingViaApiParameterGetArgs();
    }
}