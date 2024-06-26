// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskTalkConfigurationCredentialsOAuth20Args : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// The value of the API token generated. See the &lt;a href="https://docs.airbyte.com/integrations/sources/zendesk-talk"&gt;docs&lt;/a&gt; for more information.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Parsed as JSON.
        /// </summary>
        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        /// <summary>
        /// Client ID
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client Secret
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        public SourceZendeskTalkConfigurationCredentialsOAuth20Args()
        {
        }
        public static new SourceZendeskTalkConfigurationCredentialsOAuth20Args Empty => new SourceZendeskTalkConfigurationCredentialsOAuth20Args();
    }
}
