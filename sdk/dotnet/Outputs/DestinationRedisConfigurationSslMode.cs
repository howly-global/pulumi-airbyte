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
    public sealed class DestinationRedisConfigurationSslMode
    {
        /// <summary>
        /// Disable SSL.
        /// </summary>
        public readonly Outputs.DestinationRedisConfigurationSslModeDisable? Disable;
        /// <summary>
        /// Verify-full SSL mode.
        /// </summary>
        public readonly Outputs.DestinationRedisConfigurationSslModeVerifyFull? VerifyFull;

        [OutputConstructor]
        private DestinationRedisConfigurationSslMode(
            Outputs.DestinationRedisConfigurationSslModeDisable? disable,

            Outputs.DestinationRedisConfigurationSslModeVerifyFull? verifyFull)
        {
            Disable = disable;
            VerifyFull = verifyFull;
        }
    }
}
