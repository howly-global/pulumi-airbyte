// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3ConfigurationFormatCsvCommaSeparatedValuesCompressionArgs : global::Pulumi.ResourceArgs
    {
        [Input("gzip")]
        public Input<Inputs.DestinationS3ConfigurationFormatCsvCommaSeparatedValuesCompressionGzipArgs>? Gzip { get; set; }

        [Input("noCompression")]
        public Input<Inputs.DestinationS3ConfigurationFormatCsvCommaSeparatedValuesCompressionNoCompressionArgs>? NoCompression { get; set; }

        public DestinationS3ConfigurationFormatCsvCommaSeparatedValuesCompressionArgs()
        {
        }
        public static new DestinationS3ConfigurationFormatCsvCommaSeparatedValuesCompressionArgs Empty => new DestinationS3ConfigurationFormatCsvCommaSeparatedValuesCompressionArgs();
    }
}