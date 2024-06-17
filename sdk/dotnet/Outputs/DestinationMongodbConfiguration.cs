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
    public sealed class DestinationMongodbConfiguration
    {
        /// <summary>
        /// Authorization type.
        /// </summary>
        public readonly Outputs.DestinationMongodbConfigurationAuthType AuthType;
        /// <summary>
        /// Name of the database.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
        /// </summary>
        public readonly Outputs.DestinationMongodbConfigurationInstanceType? InstanceType;
        /// <summary>
        /// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
        /// </summary>
        public readonly Outputs.DestinationMongodbConfigurationTunnelMethod? TunnelMethod;

        [OutputConstructor]
        private DestinationMongodbConfiguration(
            Outputs.DestinationMongodbConfigurationAuthType authType,

            string database,

            Outputs.DestinationMongodbConfigurationInstanceType? instanceType,

            Outputs.DestinationMongodbConfigurationTunnelMethod? tunnelMethod)
        {
            AuthType = authType;
            Database = database;
            InstanceType = instanceType;
            TunnelMethod = tunnelMethod;
        }
    }
}