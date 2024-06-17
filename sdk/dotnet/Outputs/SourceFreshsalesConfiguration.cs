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
    public sealed class SourceFreshsalesConfiguration
    {
        /// <summary>
        /// Freshsales API Key. See &lt;a href="https://crmsupport.freshworks.com/support/solutions/articles/50000002503-how-to-find-my-api-key-"&gt;here&lt;/a&gt;. The key is case sensitive.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// The Name of your Freshsales domain
        /// </summary>
        public readonly string DomainName;

        [OutputConstructor]
        private SourceFreshsalesConfiguration(
            string apiKey,

            string domainName)
        {
            ApiKey = apiKey;
            DomainName = domainName;
        }
    }
}