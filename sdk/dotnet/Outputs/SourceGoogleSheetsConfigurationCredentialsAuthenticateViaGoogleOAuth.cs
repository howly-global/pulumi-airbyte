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
    public sealed class SourceGoogleSheetsConfigurationCredentialsAuthenticateViaGoogleOAuth
    {
        /// <summary>
        /// Enter your Google application's Client ID. See &lt;a href='https://developers.google.com/identity/protocols/oauth2'&gt;Google's documentation&lt;/a&gt; for more information.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Enter your Google application's Client Secret. See &lt;a href='https://developers.google.com/identity/protocols/oauth2'&gt;Google's documentation&lt;/a&gt; for more information.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// Enter your Google application's refresh token. See &lt;a href='https://developers.google.com/identity/protocols/oauth2'&gt;Google's documentation&lt;/a&gt; for more information.
        /// </summary>
        public readonly string RefreshToken;

        [OutputConstructor]
        private SourceGoogleSheetsConfigurationCredentialsAuthenticateViaGoogleOAuth(
            string clientId,

            string clientSecret,

            string refreshToken)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            RefreshToken = refreshToken;
        }
    }
}