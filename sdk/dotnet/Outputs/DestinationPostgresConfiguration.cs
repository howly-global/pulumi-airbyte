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
    public sealed class DestinationPostgresConfiguration
    {
        /// <summary>
        /// Name of the database.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// Disable Writing Final Tables. WARNING! The data format in _airbyte_data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions. Default: false
        /// </summary>
        public readonly bool? DisableTypeDedupe;
        /// <summary>
        /// Drop tables with CASCADE. WARNING! This will delete all data in all dependent objects (views, etc.). Use with caution. This option is intended for usecases which can easily rebuild the dependent objects. Default: false
        /// </summary>
        public readonly bool? DropCascade;
        /// <summary>
        /// Hostname of the database.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&amp;'. (example: key1=value1&amp;key2=value2&amp;key3=value3).
        /// </summary>
        public readonly string? JdbcUrlParams;
        /// <summary>
        /// Password associated with the username.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Port of the database. Default: 5432
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The schema to write raw tables into
        /// </summary>
        public readonly string? RawDataSchema;
        /// <summary>
        /// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public". Default: "public"
        /// </summary>
        public readonly string? Schema;
        /// <summary>
        /// SSL connection modes. 
        ///  &lt;b&gt;disable&lt;/b&gt; - Chose this mode to disable encryption of communication between Airbyte and destination database
        ///  &lt;b&gt;allow&lt;/b&gt; - Chose this mode to enable encryption only when required by the source database
        ///  &lt;b&gt;prefer&lt;/b&gt; - Chose this mode to allow unencrypted connection only if the source database does not support encryption
        ///  &lt;b&gt;require&lt;/b&gt; - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
        ///   &lt;b&gt;verify-ca&lt;/b&gt; - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
        ///   &lt;b&gt;verify-full&lt;/b&gt; - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
        ///  See more information - &lt;a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"&gt; in the docs&lt;/a&gt;.
        /// </summary>
        public readonly Outputs.DestinationPostgresConfigurationSslMode? SslMode;
        /// <summary>
        /// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
        /// </summary>
        public readonly Outputs.DestinationPostgresConfigurationTunnelMethod? TunnelMethod;
        /// <summary>
        /// Username to use to access the database.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private DestinationPostgresConfiguration(
            string database,

            bool? disableTypeDedupe,

            bool? dropCascade,

            string host,

            string? jdbcUrlParams,

            string? password,

            int? port,

            string? rawDataSchema,

            string? schema,

            Outputs.DestinationPostgresConfigurationSslMode? sslMode,

            Outputs.DestinationPostgresConfigurationTunnelMethod? tunnelMethod,

            string username)
        {
            Database = database;
            DisableTypeDedupe = disableTypeDedupe;
            DropCascade = dropCascade;
            Host = host;
            JdbcUrlParams = jdbcUrlParams;
            Password = password;
            Port = port;
            RawDataSchema = rawDataSchema;
            Schema = schema;
            SslMode = sslMode;
            TunnelMethod = tunnelMethod;
            Username = username;
        }
    }
}