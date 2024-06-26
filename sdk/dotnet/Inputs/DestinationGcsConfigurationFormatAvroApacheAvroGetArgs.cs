// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGcsConfigurationFormatAvroApacheAvroGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compression algorithm used to compress data. Default to no compression.
        /// </summary>
        [Input("compressionCodec", required: true)]
        public Input<Inputs.DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecGetArgs> CompressionCodec { get; set; } = null!;

        /// <summary>
        /// must be one of ["Avro"]; Default: "Avro"
        /// </summary>
        [Input("formatType")]
        public Input<string>? FormatType { get; set; }

        public DestinationGcsConfigurationFormatAvroApacheAvroGetArgs()
        {
        }
        public static new DestinationGcsConfigurationFormatAvroApacheAvroGetArgs Empty => new DestinationGcsConfigurationFormatAvroApacheAvroGetArgs();
    }
}
