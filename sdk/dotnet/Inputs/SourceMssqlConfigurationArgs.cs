// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMssqlConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The hostname of the database.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&amp;'. (example: key1=value1&amp;key2=value2&amp;key3=value3).
        /// </summary>
        [Input("jdbcUrlParams")]
        public Input<string>? JdbcUrlParams { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The password associated with the username.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The port of the database.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Configures how data is extracted from the database.
        /// </summary>
        [Input("replicationMethod")]
        public Input<Inputs.SourceMssqlConfigurationReplicationMethodArgs>? ReplicationMethod { get; set; }

        [Input("schemas")]
        private InputList<string>? _schemas;

        /// <summary>
        /// The list of schemas to sync from. Defaults to user. Case sensitive.
        /// </summary>
        public InputList<string> Schemas
        {
            get => _schemas ?? (_schemas = new InputList<string>());
            set => _schemas = value;
        }

        /// <summary>
        /// The encryption method which is used when communicating with the database.
        /// </summary>
        [Input("sslMethod")]
        public Input<Inputs.SourceMssqlConfigurationSslMethodArgs>? SslMethod { get; set; }

        /// <summary>
        /// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
        /// </summary>
        [Input("tunnelMethod")]
        public Input<Inputs.SourceMssqlConfigurationTunnelMethodArgs>? TunnelMethod { get; set; }

        /// <summary>
        /// The username which is used to access the database.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SourceMssqlConfigurationArgs()
        {
        }
        public static new SourceMssqlConfigurationArgs Empty => new SourceMssqlConfigurationArgs();
    }
}