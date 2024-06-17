// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceS3ConfigurationFormat
    {
        /// <summary>
        /// This connector utilises &lt;a href="https://fastavro.readthedocs.io/en/latest/" target="_blank"&gt;fastavro&lt;/a&gt; for Avro parsing.
        /// </summary>
        public readonly Outputs.SourceS3ConfigurationFormatAvro? Avro;
        /// <summary>
        /// This connector utilises &lt;a href="https: // arrow.apache.org/docs/python/generated/pyarrow.csv.open_csv.html" target="_blank"&gt;PyArrow (Apache Arrow)&lt;/a&gt; for CSV parsing.
        /// </summary>
        public readonly Outputs.SourceS3ConfigurationFormatCsv? Csv;
        /// <summary>
        /// This connector uses &lt;a href="https://arrow.apache.org/docs/python/json.html" target="_blank"&gt;PyArrow&lt;/a&gt; for JSON Lines (jsonl) file parsing.
        /// </summary>
        public readonly Outputs.SourceS3ConfigurationFormatJsonl? Jsonl;
        /// <summary>
        /// This connector utilises &lt;a href="https://arrow.apache.org/docs/python/generated/pyarrow.parquet.ParquetFile.html" target="_blank"&gt;PyArrow (Apache Arrow)&lt;/a&gt; for Parquet parsing.
        /// </summary>
        public readonly Outputs.SourceS3ConfigurationFormatParquet? Parquet;

        [OutputConstructor]
        private SourceS3ConfigurationFormat(
            Outputs.SourceS3ConfigurationFormatAvro? avro,

            Outputs.SourceS3ConfigurationFormatCsv? csv,

            Outputs.SourceS3ConfigurationFormatJsonl? jsonl,

            Outputs.SourceS3ConfigurationFormatParquet? parquet)
        {
            Avro = avro;
            Csv = csv;
            Jsonl = jsonl;
            Parquet = parquet;
        }
    }
}