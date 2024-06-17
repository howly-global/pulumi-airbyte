// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMicrosoftSharepointConfigurationStreamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When the state history of the file store is full, syncs will only read files that were last modified in the provided day range. Default: 3
        /// </summary>
        [Input("daysToSyncIfHistoryIsFull")]
        public Input<int>? DaysToSyncIfHistoryIsFull { get; set; }

        /// <summary>
        /// The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.
        /// </summary>
        [Input("format", required: true)]
        public Input<Inputs.SourceMicrosoftSharepointConfigurationStreamFormatArgs> Format { get; set; } = null!;

        [Input("globs")]
        private InputList<string>? _globs;

        /// <summary>
        /// The pattern used to specify which files should be selected from the file system. For more information on glob pattern matching look &lt;a href="https://en.wikipedia.org/wiki/Glob_(programming)"&gt;here&lt;/a&gt;.
        /// </summary>
        public InputList<string> Globs
        {
            get => _globs ?? (_globs = new InputList<string>());
            set => _globs = value;
        }

        /// <summary>
        /// The schema that will be used to validate records extracted from the file. This will override the stream schema that is auto-detected from incoming files.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// The name of the stream.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("primaryKey")]
        private Input<string>? _primaryKey;

        /// <summary>
        /// The column or columns (for a composite key) that serves as the unique identifier of a record. If empty, the primary key will default to the parser's default primary key.
        /// </summary>
        public Input<string>? PrimaryKey
        {
            get => _primaryKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _primaryKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// When enabled, syncs will not validate or structure records against the stream's schema. Default: false
        /// </summary>
        [Input("schemaless")]
        public Input<bool>? Schemaless { get; set; }

        /// <summary>
        /// The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema. must be one of ["Emit Record", "Skip Record", "Wait for Discover"]; Default: "Emit Record"
        /// </summary>
        [Input("validationPolicy")]
        public Input<string>? ValidationPolicy { get; set; }

        public SourceMicrosoftSharepointConfigurationStreamArgs()
        {
        }
        public static new SourceMicrosoftSharepointConfigurationStreamArgs Empty => new SourceMicrosoftSharepointConfigurationStreamArgs();
    }
}