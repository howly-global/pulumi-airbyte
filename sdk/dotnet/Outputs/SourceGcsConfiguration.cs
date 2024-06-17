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
    public sealed class SourceGcsConfiguration
    {
        /// <summary>
        /// Name of the GCS bucket where the file(s) exist.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Enter your Google Cloud &lt;a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys"&gt;service account key&lt;/a&gt; in JSON format
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00.000000Z. Any file modified before this date will not be replicated.
        /// </summary>
        public readonly string? StartDate;
        /// <summary>
        /// Each instance of this configuration defines a &lt;a href=https://docs.airbyte.com/cloud/core-concepts#stream&gt;stream&lt;/a&gt;. Use this to define which files belong in the stream, their format, and how they should be parsed and validated. When sending data to warehouse destination such as Snowflake or BigQuery, each stream is a separate table.
        /// </summary>
        public readonly ImmutableArray<Outputs.SourceGcsConfigurationStream> Streams;

        [OutputConstructor]
        private SourceGcsConfiguration(
            string bucket,

            string serviceAccount,

            string? startDate,

            ImmutableArray<Outputs.SourceGcsConfigurationStream> streams)
        {
            Bucket = bucket;
            ServiceAccount = serviceAccount;
            StartDate = startDate;
            Streams = streams;
        }
    }
}