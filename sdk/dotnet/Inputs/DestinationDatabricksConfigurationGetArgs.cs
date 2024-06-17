// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationDatabricksConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// You must agree to the Databricks JDBC Driver &lt;a href="https://databricks.com/jdbc-odbc-driver-license"&gt;Terms &amp; Conditions&lt;/a&gt; to use this connector. Default: false
        /// </summary>
        [Input("acceptTerms")]
        public Input<bool>? AcceptTerms { get; set; }

        /// <summary>
        /// Storage on which the delta lake is built.
        /// </summary>
        [Input("dataSource", required: true)]
        public Input<Inputs.DestinationDatabricksConfigurationDataSourceGetArgs> DataSource { get; set; } = null!;

        /// <summary>
        /// The name of the catalog. If not specified otherwise, the "hive_metastore" will be used.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Databricks Cluster HTTP Path.
        /// </summary>
        [Input("databricksHttpPath", required: true)]
        public Input<string> DatabricksHttpPath { get; set; } = null!;

        [Input("databricksPersonalAccessToken", required: true)]
        private Input<string>? _databricksPersonalAccessToken;

        /// <summary>
        /// Databricks Personal Access Token for making authenticated requests.
        /// </summary>
        public Input<string>? DatabricksPersonalAccessToken
        {
            get => _databricksPersonalAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _databricksPersonalAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Databricks Cluster Port. Default: "443"
        /// </summary>
        [Input("databricksPort")]
        public Input<string>? DatabricksPort { get; set; }

        /// <summary>
        /// Databricks Cluster Server Hostname.
        /// </summary>
        [Input("databricksServerHostname", required: true)]
        public Input<string> DatabricksServerHostname { get; set; } = null!;

        /// <summary>
        /// Support schema evolution for all streams. If "false", the connector might fail when a stream's schema changes. Default: false
        /// </summary>
        [Input("enableSchemaEvolution")]
        public Input<bool>? EnableSchemaEvolution { get; set; }

        /// <summary>
        /// Default to 'true'. Switch it to 'false' for debugging purpose. Default: true
        /// </summary>
        [Input("purgeStagingData")]
        public Input<bool>? PurgeStagingData { get; set; }

        /// <summary>
        /// The default schema tables are written. If not specified otherwise, the "default" will be used. Default: "default"
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public DestinationDatabricksConfigurationGetArgs()
        {
        }
        public static new DestinationDatabricksConfigurationGetArgs Empty => new DestinationDatabricksConfigurationGetArgs();
    }
}