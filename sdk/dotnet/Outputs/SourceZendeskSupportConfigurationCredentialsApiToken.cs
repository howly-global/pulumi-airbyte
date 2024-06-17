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
    public sealed class SourceZendeskSupportConfigurationCredentialsApiToken
    {
        /// <summary>
        /// Parsed as JSON.
        /// </summary>
        public readonly string? AdditionalProperties;
        /// <summary>
        /// The value of the API token generated. See our &lt;a href="https://docs.airbyte.com/integrations/sources/zendesk-support#setup-guide"&gt;full documentation&lt;/a&gt; for more information on generating this token.
        /// </summary>
        public readonly string ApiToken;
        /// <summary>
        /// The user email for your Zendesk account.
        /// </summary>
        public readonly string Email;

        [OutputConstructor]
        private SourceZendeskSupportConfigurationCredentialsApiToken(
            string? additionalProperties,

            string apiToken,

            string email)
        {
            AdditionalProperties = additionalProperties;
            ApiToken = apiToken;
            Email = email;
        }
    }
}