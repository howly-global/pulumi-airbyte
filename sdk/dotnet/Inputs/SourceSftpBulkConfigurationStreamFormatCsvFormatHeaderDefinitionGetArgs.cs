// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("autogenerated")]
        public Input<Inputs.SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionAutogeneratedGetArgs>? Autogenerated { get; set; }

        [Input("fromCsv")]
        public Input<Inputs.SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionFromCsvGetArgs>? FromCsv { get; set; }

        [Input("userProvided")]
        public Input<Inputs.SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionUserProvidedGetArgs>? UserProvided { get; set; }

        public SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionGetArgs()
        {
        }
        public static new SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionGetArgs Empty => new SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionGetArgs();
    }
}