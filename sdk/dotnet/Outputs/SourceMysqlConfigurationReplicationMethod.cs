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
    public sealed class SourceMysqlConfigurationReplicationMethod
    {
        /// <summary>
        /// &lt;i&gt;Recommended&lt;/i&gt; - Incrementally reads new inserts, updates, and deletes using the MySQL &lt;a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc"&gt;binary log&lt;/a&gt;. This must be enabled on your database.
        /// </summary>
        public readonly Outputs.SourceMysqlConfigurationReplicationMethodReadChangesUsingBinaryLogCdc? ReadChangesUsingBinaryLogCdc;
        /// <summary>
        /// Incrementally detects new inserts and updates using the &lt;a href="https://docs.airbyte.com/understanding-airbyte/connections/incremental-append/#user-defined-cursor"&gt;cursor column&lt;/a&gt; chosen when configuring a connection (e.g. created_at, updated_at).
        /// </summary>
        public readonly Outputs.SourceMysqlConfigurationReplicationMethodScanChangesWithUserDefinedCursor? ScanChangesWithUserDefinedCursor;

        [OutputConstructor]
        private SourceMysqlConfigurationReplicationMethod(
            Outputs.SourceMysqlConfigurationReplicationMethodReadChangesUsingBinaryLogCdc? readChangesUsingBinaryLogCdc,

            Outputs.SourceMysqlConfigurationReplicationMethodScanChangesWithUserDefinedCursor? scanChangesWithUserDefinedCursor)
        {
            ReadChangesUsingBinaryLogCdc = readChangesUsingBinaryLogCdc;
            ScanChangesWithUserDefinedCursor = scanChangesWithUserDefinedCursor;
        }
    }
}