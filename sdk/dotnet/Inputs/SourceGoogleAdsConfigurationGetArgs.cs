// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAdsConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A conversion window is the number of days after an ad interaction (such as an ad click or video view) during which a conversion, such as a purchase, is recorded in Google Ads. For more information, see &lt;a href="https://support.google.com/google-ads/answer/3123169?hl=en"&gt;Google's documentation&lt;/a&gt;. Default: 14
        /// </summary>
        [Input("conversionWindowDays")]
        public Input<int>? ConversionWindowDays { get; set; }

        [Input("credentials", required: true)]
        public Input<Inputs.SourceGoogleAdsConfigurationCredentialsGetArgs> Credentials { get; set; } = null!;

        [Input("customQueriesArrays")]
        private InputList<Inputs.SourceGoogleAdsConfigurationCustomQueriesArrayGetArgs>? _customQueriesArrays;
        public InputList<Inputs.SourceGoogleAdsConfigurationCustomQueriesArrayGetArgs> CustomQueriesArrays
        {
            get => _customQueriesArrays ?? (_customQueriesArrays = new InputList<Inputs.SourceGoogleAdsConfigurationCustomQueriesArrayGetArgs>());
            set => _customQueriesArrays = value;
        }

        /// <summary>
        /// Comma-separated list of (client) customer IDs. Each customer ID must be specified as a 10-digit number without dashes. For detailed instructions on finding this value, refer to our &lt;a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide"&gt;documentation&lt;/a&gt;.
        /// </summary>
        [Input("customerId")]
        public Input<string>? CustomerId { get; set; }

        [Input("customerStatusFilters")]
        private InputList<string>? _customerStatusFilters;

        /// <summary>
        /// A list of customer statuses to filter on. For detailed info about what each status mean refer to Google Ads &lt;a href="https://developers.google.com/google-ads/api/reference/rpc/v15/CustomerStatusEnum.CustomerStatus"&gt;documentation&lt;/a&gt;.
        /// </summary>
        public InputList<string> CustomerStatusFilters
        {
            get => _customerStatusFilters ?? (_customerStatusFilters = new InputList<string>());
            set => _customerStatusFilters = value;
        }

        /// <summary>
        /// UTC date in the format YYYY-MM-DD. Any data after this date will not be replicated. (Default value of today is used if not set)
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated. (Default value of two years ago is used if not set)
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceGoogleAdsConfigurationGetArgs()
        {
        }
        public static new SourceGoogleAdsConfigurationGetArgs Empty => new SourceGoogleAdsConfigurationGetArgs();
    }
}
