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
    public sealed class DestinationMssqlConfigurationSslMethod
    {
        /// <summary>
        /// Use the certificate provided by the server without verification. (For testing purposes only!)
        /// </summary>
        public readonly Outputs.DestinationMssqlConfigurationSslMethodEncryptedTrustServerCertificate? EncryptedTrustServerCertificate;
        /// <summary>
        /// Verify and use the certificate provided by the server.
        /// </summary>
        public readonly Outputs.DestinationMssqlConfigurationSslMethodEncryptedVerifyCertificate? EncryptedVerifyCertificate;

        [OutputConstructor]
        private DestinationMssqlConfigurationSslMethod(
            Outputs.DestinationMssqlConfigurationSslMethodEncryptedTrustServerCertificate? encryptedTrustServerCertificate,

            Outputs.DestinationMssqlConfigurationSslMethodEncryptedVerifyCertificate? encryptedVerifyCertificate)
        {
            EncryptedTrustServerCertificate = encryptedTrustServerCertificate;
            EncryptedVerifyCertificate = encryptedVerifyCertificate;
        }
    }
}