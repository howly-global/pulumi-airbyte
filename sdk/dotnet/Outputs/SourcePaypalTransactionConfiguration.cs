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
    public sealed class SourcePaypalTransactionConfiguration
    {
        /// <summary>
        /// The Client ID of your Paypal developer application.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The Client Secret of your Paypal developer application.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// Start Date parameter for the list dispute endpoint in &lt;a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\"&gt;ISO format&lt;/a&gt;. This Start Date must be in range within 180 days before present time, and requires ONLY 3 miliseconds(mandatory). If you don't use this option, it defaults to a start date set 180 days in the past.
        /// </summary>
        public readonly string? DisputeStartDate;
        /// <summary>
        /// End Date for data extraction in &lt;a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\"&gt;ISO format&lt;/a&gt;. This can be help you select specific range of time, mainly for test purposes  or data integrity tests. When this is not used, now_utc() is used by the streams. This does not apply to Disputes and Product streams.
        /// </summary>
        public readonly string? EndDate;
        /// <summary>
        /// Determines whether to use the sandbox or production environment. Default: false
        /// </summary>
        public readonly bool? IsSandbox;
        /// <summary>
        /// The key to refresh the expired access token.
        /// </summary>
        public readonly string? RefreshToken;
        /// <summary>
        /// Start Date for data extraction in &lt;a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\"&gt;ISO format&lt;/a&gt;. Date must be in range from 3 years till 12 hrs before present time.
        /// </summary>
        public readonly string StartDate;
        /// <summary>
        /// The number of days per request. Must be a number between 1 and 31. Default: 7
        /// </summary>
        public readonly int? TimeWindow;

        [OutputConstructor]
        private SourcePaypalTransactionConfiguration(
            string clientId,

            string clientSecret,

            string? disputeStartDate,

            string? endDate,

            bool? isSandbox,

            string? refreshToken,

            string startDate,

            int? timeWindow)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            DisputeStartDate = disputeStartDate;
            EndDate = endDate;
            IsSandbox = isSandbox;
            RefreshToken = refreshToken;
            StartDate = startDate;
            TimeWindow = timeWindow;
        }
    }
}