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
    public sealed class DestinationMssqlConfigurationSslMethodEncryptedVerifyCertificate
    {
        /// <summary>
        /// Specifies the host name of the server. The value of this property must match the subject property of the certificate.
        /// </summary>
        public readonly string? HostNameInCertificate;

        [OutputConstructor]
        private DestinationMssqlConfigurationSslMethodEncryptedVerifyCertificate(string? hostNameInCertificate)
        {
            HostNameInCertificate = hostNameInCertificate;
        }
    }
}