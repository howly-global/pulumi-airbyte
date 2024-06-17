// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceLokaliseConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        private Input<string>? _apiKey;

        /// <summary>
        /// Lokalise API Key with read-access. Available at Profile settings &gt; API tokens. See &lt;a href="https://docs.lokalise.com/en/articles/1929556-api-tokens"&gt;here&lt;/a&gt;.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Lokalise project ID. Available at Project Settings &gt; General.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public SourceLokaliseConfigurationArgs()
        {
        }
        public static new SourceLokaliseConfigurationArgs Empty => new SourceLokaliseConfigurationArgs();
    }
}