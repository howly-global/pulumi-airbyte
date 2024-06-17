// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceNotionConfigurationCredentialsOAuth20Args : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// The Access Token received by completing the OAuth flow for your Notion integration. See our &lt;a href='https://docs.airbyte.com/integrations/sources/notion#step-2-set-permissions-and-acquire-authorization-credentials'&gt;docs&lt;/a&gt; for more information.
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
        /// The Client ID of your Notion integration. See our &lt;a href='https://docs.airbyte.com/integrations/sources/notion#step-2-set-permissions-and-acquire-authorization-credentials'&gt;docs&lt;/a&gt; for more information.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The Client Secret of your Notion integration. See our &lt;a href='https://docs.airbyte.com/integrations/sources/notion#step-2-set-permissions-and-acquire-authorization-credentials'&gt;docs&lt;/a&gt; for more information.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        public SourceNotionConfigurationCredentialsOAuth20Args()
        {
        }
        public static new SourceNotionConfigurationCredentialsOAuth20Args Empty => new SourceNotionConfigurationCredentialsOAuth20Args();
    }
}