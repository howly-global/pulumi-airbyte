// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceLinkedinAdsConfigurationCredentialsOAuth20GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ID of your developer application. Refer to our &lt;a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'&gt;documentation&lt;/a&gt; for more information.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The client secret of your developer application. Refer to our &lt;a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'&gt;documentation&lt;/a&gt; for more information.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        [Input("refreshToken", required: true)]
        private Input<string>? _refreshToken;

        /// <summary>
        /// The key to refresh the expired access token. Refer to our &lt;a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'&gt;documentation&lt;/a&gt; for more information.
        /// </summary>
        public Input<string>? RefreshToken
        {
            get => _refreshToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _refreshToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceLinkedinAdsConfigurationCredentialsOAuth20GetArgs()
        {
        }
        public static new SourceLinkedinAdsConfigurationCredentialsOAuth20GetArgs Empty => new SourceLinkedinAdsConfigurationCredentialsOAuth20GetArgs();
    }
}
