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
    public sealed class SourceLinkedinAdsConfigurationCredentialsAccessToken
    {
        /// <summary>
        /// The access token generated for your developer application. Refer to our &lt;a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'&gt;documentation&lt;/a&gt; for more information.
        /// </summary>
        public readonly string AccessToken;

        [OutputConstructor]
        private SourceLinkedinAdsConfigurationCredentialsAccessToken(string accessToken)
        {
            AccessToken = accessToken;
        }
    }
}