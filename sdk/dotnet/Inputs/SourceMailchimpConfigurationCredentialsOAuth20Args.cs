// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMailchimpConfigurationCredentialsOAuth20Args : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// An access token generated using the above client ID and secret.
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
        /// The Client ID of your OAuth application.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The Client Secret of your OAuth application.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        public SourceMailchimpConfigurationCredentialsOAuth20Args()
        {
        }
        public static new SourceMailchimpConfigurationCredentialsOAuth20Args Empty => new SourceMailchimpConfigurationCredentialsOAuth20Args();
    }
}