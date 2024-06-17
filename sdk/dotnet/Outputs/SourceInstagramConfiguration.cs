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
    public sealed class SourceInstagramConfiguration
    {
        /// <summary>
        /// The value of the access token generated with &lt;b&gt;instagram_basic, instagram_manage_insights, pages_show_list, pages_read_engagement, Instagram Public Content Access&lt;/b&gt; permissions. See the &lt;a href="https://docs.airbyte.com/integrations/sources/instagram/#step-1-set-up-instagram"&gt;docs&lt;/a&gt; for more information
        /// </summary>
        public readonly string AccessToken;
        /// <summary>
        /// The Client ID for your Oauth application
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// The Client Secret for your Oauth application
        /// </summary>
        public readonly string? ClientSecret;
        /// <summary>
        /// The date from which you'd like to replicate data for User Insights, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated. If left blank, the start date will be set to 2 years before the present date.
        /// </summary>
        public readonly string? StartDate;

        [OutputConstructor]
        private SourceInstagramConfiguration(
            string accessToken,

            string? clientId,

            string? clientSecret,

            string? startDate)
        {
            AccessToken = accessToken;
            ClientId = clientId;
            ClientSecret = clientSecret;
            StartDate = startDate;
        }
    }
}