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
    public sealed class SourceGoogleSearchConsoleConfigurationAuthorizationOAuth
    {
        /// <summary>
        /// Access token for making authenticated requests. Read more &lt;a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing"&gt;here&lt;/a&gt;.
        /// </summary>
        public readonly string? AccessToken;
        /// <summary>
        /// The client ID of your Google Search Console developer application. Read more &lt;a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing"&gt;here&lt;/a&gt;.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The client secret of your Google Search Console developer application. Read more &lt;a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing"&gt;here&lt;/a&gt;.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The token for obtaining a new access token. Read more &lt;a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing"&gt;here&lt;/a&gt;.
        /// </summary>
        public readonly string RefreshToken;

        [OutputConstructor]
        private SourceGoogleSearchConsoleConfigurationAuthorizationOAuth(
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