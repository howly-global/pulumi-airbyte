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
    public sealed class SourceMongodbV2Configuration
    {
        /// <summary>
        /// Configures the MongoDB cluster type.
        /// </summary>
        public readonly Outputs.SourceMongodbV2ConfigurationDatabaseConfig DatabaseConfig;
        /// <summary>
        /// The maximum number of documents to sample when attempting to discover the unique fields for a collection. Default: 10000
        /// </summary>
        public readonly int? DiscoverSampleSize;
        /// <summary>
        /// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Default: 300
        /// </summary>
        public readonly int? InitialWaitingSeconds;
        /// <summary>
        /// Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss. must be one of ["Fail sync", "Re-sync data"]; Default: "Fail sync"
        /// </summary>
        public readonly string? InvalidCdcCursorPositionBehavior;
        /// <summary>
        /// The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful. Default: 10000
        /// </summary>
        public readonly int? QueueSize;
        /// <summary>
        /// Determines how Airbyte looks up the value of an updated document. If 'Lookup' is chosen, the current value of the document will be read. If 'Post Image' is chosen, then the version of the document immediately after an update will be read. WARNING : Severe data loss will occur if this option is chosen and the appropriate settings are not set on your Mongo instance : https://www.mongodb.com/docs/manual/changeStreams/#change-streams-with-document-pre-and-post-images. must be one of ["Lookup", "Post Image"]; Default: "Lookup"
        /// </summary>
        public readonly string? UpdateCaptureMode;

        [OutputConstructor]
        private SourceMongodbV2Configuration(
            Outputs.SourceMongodbV2ConfigurationDatabaseConfig databaseConfig,

            int? discoverSampleSize,

            int? initialWaitingSeconds,

            string? invalidCdcCursorPositionBehavior,

            int? queueSize,

            string? updateCaptureMode)
        {
            DatabaseConfig = databaseConfig;
            DiscoverSampleSize = discoverSampleSize;
            InitialWaitingSeconds = initialWaitingSeconds;
            InvalidCdcCursorPositionBehavior = invalidCdcCursorPositionBehavior;
            QueueSize = queueSize;
            UpdateCaptureMode = updateCaptureMode;
        }
    }
}