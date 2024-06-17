// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskSupportConfigurationCredentialsApiTokenGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parsed as JSON.
        /// </summary>
        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        [Input("apiToken", required: true)]
        private Input<string>? _apiToken;

        /// <summary>
        /// The value of the API token generated. See our &lt;a href="https://docs.airbyte.com/integrations/sources/zendesk-support#setup-guide"&gt;full documentation&lt;/a&gt; for more information on generating this token.
        /// </summary>
        public Input<string>? ApiToken
        {
            get => _apiToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The user email for your Zendesk account.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public SourceZendeskSupportConfigurationCredentialsApiTokenGetArgs()
        {
        }
        public static new SourceZendeskSupportConfigurationCredentialsApiTokenGetArgs Empty => new SourceZendeskSupportConfigurationCredentialsApiTokenGetArgs();
    }
}