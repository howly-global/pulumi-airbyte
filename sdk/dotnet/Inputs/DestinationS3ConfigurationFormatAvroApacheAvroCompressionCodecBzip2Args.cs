// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3ConfigurationFormatAvroApacheAvroCompressionCodecBzip2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// must be one of ["bzip2"]; Default: "bzip2"
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        public DestinationS3ConfigurationFormatAvroApacheAvroCompressionCodecBzip2Args()
        {
        }
        public static new DestinationS3ConfigurationFormatAvroApacheAvroCompressionCodecBzip2Args Empty => new DestinationS3ConfigurationFormatAvroApacheAvroCompressionCodecBzip2Args();
    }
}
