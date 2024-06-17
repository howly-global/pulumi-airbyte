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
    public sealed class SourceSmartsheetsConfigurationCredentialsApiAccessToken
    {
        /// <summary>
        /// The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account &gt; Apps &amp; Integrations &gt; API Access. See the &lt;a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide"&gt;setup guide&lt;/a&gt; for information on how to obtain this token.
        /// </summary>
        public readonly string AccessToken;

        [OutputConstructor]
        private SourceSmartsheetsConfigurationCredentialsApiAccessToken(string accessToken)
        {
            AccessToken = accessToken;
        }
    }
}