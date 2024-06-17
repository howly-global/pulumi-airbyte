// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSurveymonkeyConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// Access Token for making authenticated requests. See the &lt;a href="https://docs.airbyte.io/integrations/sources/surveymonkey"&gt;docs&lt;/a&gt; for information on how to generate this key.
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
        /// The Client ID of the SurveyMonkey developer application.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The Client Secret of the SurveyMonkey developer application.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        public SourceSurveymonkeyConfigurationCredentialsArgs()
        {
        }
        public static new SourceSurveymonkeyConfigurationCredentialsArgs Empty => new SourceSurveymonkeyConfigurationCredentialsArgs();
    }
}