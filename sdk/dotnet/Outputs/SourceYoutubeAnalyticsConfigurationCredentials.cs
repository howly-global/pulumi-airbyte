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
    public sealed class SourceYoutubeAnalyticsConfigurationCredentials
    {
        /// <summary>
        /// Parsed as JSON.
        /// </summary>
        public readonly string? AdditionalProperties;
        /// <summary>
        /// The Client ID of your developer application
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The client secret of your developer application
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// A refresh token generated using the above client ID and secret
        /// </summary>
        public readonly string RefreshToken;

        [OutputConstructor]
        private SourceYoutubeAnalyticsConfigurationCredentials(
            string? additionalProperties,

            string clientId,

            string clientSecret,

            string refreshToken)
        {
            AdditionalProperties = additionalProperties;
            ClientId = clientId;
            ClientSecret = clientSecret;
            RefreshToken = refreshToken;
        }
    }
}
