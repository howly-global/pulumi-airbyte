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
    public sealed class SourceSlackConfigurationCredentialsApiToken
    {
        /// <summary>
        /// A Slack bot token. See the &lt;a href="https://docs.airbyte.com/integrations/sources/slack"&gt;docs&lt;/a&gt; for instructions on how to generate it.
        /// </summary>
        public readonly string ApiToken;

        [OutputConstructor]
        private SourceSlackConfigurationCredentialsApiToken(string apiToken)
        {
            ApiToken = apiToken;
        }
    }
}