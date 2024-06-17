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
    public sealed class SourceSftpBulkConfigurationStreamFormat
    {
        public readonly Outputs.SourceSftpBulkConfigurationStreamFormatAvroFormat? AvroFormat;
        public readonly Outputs.SourceSftpBulkConfigurationStreamFormatCsvFormat? CsvFormat;
        /// <summary>
        /// Extract text from document formats (.pdf, .docx, .md, .pptx) and emit as one record per file.
        /// </summary>
        public readonly Outputs.SourceSftpBulkConfigurationStreamFormatDocumentFileTypeFormatExperimental? DocumentFileTypeFormatExperimental;
        public readonly Outputs.SourceSftpBulkConfigurationStreamFormatJsonlFormat? JsonlFormat;
        public readonly Outputs.SourceSftpBulkConfigurationStreamFormatParquetFormat? ParquetFormat;

        [OutputConstructor]
        private SourceSftpBulkConfigurationStreamFormat(
            Outputs.SourceSftpBulkConfigurationStreamFormatAvroFormat? avroFormat,

            Outputs.SourceSftpBulkConfigurationStreamFormatCsvFormat? csvFormat,

            Outputs.SourceSftpBulkConfigurationStreamFormatDocumentFileTypeFormatExperimental? documentFileTypeFormatExperimental,

            Outputs.SourceSftpBulkConfigurationStreamFormatJsonlFormat? jsonlFormat,

            Outputs.SourceSftpBulkConfigurationStreamFormatParquetFormat? parquetFormat)
        {
            AvroFormat = avroFormat;
            CsvFormat = csvFormat;
            DocumentFileTypeFormatExperimental = documentFileTypeFormatExperimental;
            JsonlFormat = jsonlFormat;
            ParquetFormat = parquetFormat;
        }
    }
}