// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSnapchatMarketingConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the principle for conversion reporting. must be one of ["conversion", "impression"]; Default: "conversion"
        /// </summary>
        [Input("actionReportTime")]
        public Input<string>? ActionReportTime { get; set; }

        /// <summary>
        /// The Client ID of your Snapchat developer application.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The Client Secret of your Snapchat developer application.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// Date in the format 2017-01-25. Any data after this date will not be replicated.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        [Input("refreshToken", required: true)]
        private Input<string>? _refreshToken;

        /// <summary>
        /// Refresh Token to renew the expired Access Token.
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

        /// <summary>
        /// Date in the format 2022-01-01. Any data before this date will not be replicated. Default: "2022-01-01"
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// Attribution window for swipe ups. must be one of ["1_DAY", "7_DAY", "28_DAY"]; Default: "28_DAY"
        /// </summary>
        [Input("swipeUpAttributionWindow")]
        public Input<string>? SwipeUpAttributionWindow { get; set; }

        /// <summary>
        /// Attribution window for views. must be one of ["1_HOUR", "3_HOUR", "6_HOUR", "1_DAY", "7_DAY"]; Default: "1_DAY"
        /// </summary>
        [Input("viewAttributionWindow")]
        public Input<string>? ViewAttributionWindow { get; set; }

        public SourceSnapchatMarketingConfigurationGetArgs()
        {
        }
        public static new SourceSnapchatMarketingConfigurationGetArgs Empty => new SourceSnapchatMarketingConfigurationGetArgs();
    }
}
