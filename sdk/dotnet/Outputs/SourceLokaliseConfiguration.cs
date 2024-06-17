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
    public sealed class SourceLokaliseConfiguration
    {
        /// <summary>
        /// Lokalise API Key with read-access. Available at Profile settings &gt; API tokens. See &lt;a href="https://docs.lokalise.com/en/articles/1929556-api-tokens"&gt;here&lt;/a&gt;.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// Lokalise project ID. Available at Project Settings &gt; General.
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private SourceLokaliseConfiguration(
            string apiKey,

            string projectId)
        {
            ApiKey = apiKey;
            ProjectId = projectId;
        }
    }
}