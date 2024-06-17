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
    public sealed class SourceGoogleAnalyticsDataApiConfigurationCredentialsAuthenticateViaGoogleOauth
    {
        /// <summary>
        /// Access Token for making authenticated requests.
        /// </summary>
        public readonly string? AccessToken;
        /// <summary>
        /// The Client ID of your Google Analytics developer application.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The Client Secret of your Google Analytics developer application.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The token for obtaining a new access token.
        /// </summary>
        public readonly string RefreshToken;

        [OutputConstructor]
        private SourceGoogleAnalyticsDataApiConfigurationCredentialsAuthenticateViaGoogleOauth(
            string? accessToken,

            string clientId,

            string clientSecret,

            string refreshToken)
        {
            AccessToken = accessToken;
            ClientId = clientId;
            ClientSecret = clientSecret;
            RefreshToken = refreshToken;
        }
    }
}