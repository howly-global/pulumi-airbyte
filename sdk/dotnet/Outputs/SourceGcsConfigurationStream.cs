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
    public sealed class SourceGcsConfigurationStream
    {
        /// <summary>
        /// When the state history of the file store is full, syncs will only read files that were last modified in the provided day range. Default: 3
        /// </summary>
        public readonly int? DaysToSyncIfHistoryIsFull;
        /// <summary>
        /// The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.
        /// </summary>
        public readonly Outputs.SourceGcsConfigurationStreamFormat Format;
        /// <summary>
        /// The pattern used to specify which files should be selected from the file system. For more information on glob pattern matching look &lt;a href="https://en.wikipedia.org/wiki/Glob_(programming)"&gt;here&lt;/a&gt;.
        /// </summary>
        public readonly ImmutableArray<string> Globs;
        /// <summary>
        /// The schema that will be used to validate records extracted from the file. This will override the stream schema that is auto-detected from incoming files.
        /// </summary>
        public readonly string? InputSchema;
        /// <summary>
        /// The path prefix configured in previous versions of the GCS connector. This option is deprecated in favor of a single glob.
        /// </summary>
        public readonly string? LegacyPrefix;
        /// <summary>
        /// The name of the stream.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The column or columns (for a composite key) that serves as the unique identifier of a record. If empty, the primary key will default to the parser's default primary key.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// When enabled, syncs will not validate or structure records against the stream's schema. Default: false
        /// </summary>
        public readonly bool? Schemaless;
        /// <summary>
        /// The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema. must be one of ["Emit Record", "Skip Record", "Wait for Discover"]; Default: "Emit Record"
        /// </summary>
        public readonly string? ValidationPolicy;

        [OutputConstructor]
        private SourceGcsConfigurationStream(
            int? daysToSyncIfHistoryIsFull,

            Outputs.SourceGcsConfigurationStreamFormat format,

            ImmutableArray<string> globs,

            string? inputSchema,

            string? legacyPrefix,

            string name,

            string? primaryKey,

            bool? schemaless,

            string? validationPolicy)
        {
            DaysToSyncIfHistoryIsFull = daysToSyncIfHistoryIsFull;
            Format = format;
            Globs = globs;
            InputSchema = inputSchema;
            LegacyPrefix = legacyPrefix;
            Name = name;
            PrimaryKey = primaryKey;
            Schemaless = schemaless;
            ValidationPolicy = validationPolicy;
        }
    }
}
