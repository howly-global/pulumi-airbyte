// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("autogenerated")]
        public Input<Inputs.SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionAutogeneratedArgs>? Autogenerated { get; set; }

        [Input("fromCsv")]
        public Input<Inputs.SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionFromCsvArgs>? FromCsv { get; set; }

        [Input("userProvided")]
        public Input<Inputs.SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionUserProvidedArgs>? UserProvided { get; set; }

        public SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionArgs()
        {
        }
        public static new SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionArgs Empty => new SourceSftpBulkConfigurationStreamFormatCsvFormatHeaderDefinitionArgs();
    }
}