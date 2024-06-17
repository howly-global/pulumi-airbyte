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
    public sealed class SourceFreshdeskConfiguration
    {
        /// <summary>
        /// Freshdesk API Key. See the &lt;a href="https://docs.airbyte.com/integrations/sources/freshdesk"&gt;docs&lt;/a&gt; for more information on how to obtain this key.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// Freshdesk domain
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// Number of days for lookback window for the stream Satisfaction Ratings. Default: 14
        /// </summary>
        public readonly int? LookbackWindowInDays;
        /// <summary>
        /// The number of requests per minute that this source allowed to use. There is a rate limit of 50 requests per minute per app per account.
        /// </summary>
        public readonly int? RequestsPerMinute;
        /// <summary>
        /// UTC date and time. Any data created after this date will be replicated. If this parameter is not set, all data will be replicated.
        /// </summary>
        public readonly string? StartDate;

        [OutputConstructor]
        private SourceFreshdeskConfiguration(
            string apiKey,

            string domain,

            int? lookbackWindowInDays,

            int? requestsPerMinute,

            string? startDate)
        {
            ApiKey = apiKey;
            Domain = domain;
            LookbackWindowInDays = lookbackWindowInDays;
            RequestsPerMinute = requestsPerMinute;
            StartDate = startDate;
        }
    }
}