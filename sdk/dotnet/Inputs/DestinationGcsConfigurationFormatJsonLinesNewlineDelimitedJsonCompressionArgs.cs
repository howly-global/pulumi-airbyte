// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGcsConfigurationFormatJsonLinesNewlineDelimitedJsonCompressionArgs : global::Pulumi.ResourceArgs
    {
        [Input("gzip")]
        public Input<Inputs.DestinationGcsConfigurationFormatJsonLinesNewlineDelimitedJsonCompressionGzipArgs>? Gzip { get; set; }

        [Input("noCompression")]
        public Input<Inputs.DestinationGcsConfigurationFormatJsonLinesNewlineDelimitedJsonCompressionNoCompressionArgs>? NoCompression { get; set; }

        public DestinationGcsConfigurationFormatJsonLinesNewlineDelimitedJsonCompressionArgs()
        {
        }
        public static new DestinationGcsConfigurationFormatJsonLinesNewlineDelimitedJsonCompressionArgs Empty => new DestinationGcsConfigurationFormatJsonLinesNewlineDelimitedJsonCompressionArgs();
    }
}