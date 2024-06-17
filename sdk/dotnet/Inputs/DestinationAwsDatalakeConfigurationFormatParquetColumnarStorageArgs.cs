// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationAwsDatalakeConfigurationFormatParquetColumnarStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compression algorithm used to compress data. must be one of ["UNCOMPRESSED", "SNAPPY", "GZIP", "ZSTD"]; Default: "SNAPPY"
        /// </summary>
        [Input("compressionCodec")]
        public Input<string>? CompressionCodec { get; set; }

        /// <summary>
        /// must be one of ["Parquet"]; Default: "Parquet"
        /// </summary>
        [Input("formatType")]
        public Input<string>? FormatType { get; set; }

        public DestinationAwsDatalakeConfigurationFormatParquetColumnarStorageArgs()
        {
        }
        public static new DestinationAwsDatalakeConfigurationFormatParquetColumnarStorageArgs Empty => new DestinationAwsDatalakeConfigurationFormatParquetColumnarStorageArgs();
    }
}