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
    public sealed class SourceWoocommerceConfiguration
    {
        /// <summary>
        /// Customer Key for API in WooCommerce shop
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// Customer Secret for API in WooCommerce shop
        /// </summary>
        public readonly string ApiSecret;
        /// <summary>
        /// The name of the store. For https://EXAMPLE.com, the shop name is 'EXAMPLE.com'.
        /// </summary>
        public readonly string Shop;
        /// <summary>
        /// The date you would like to replicate data from. Format: YYYY-MM-DD
        /// </summary>
        public readonly string StartDate;

        [OutputConstructor]
        private SourceWoocommerceConfiguration(
            string apiKey,

            string apiSecret,

            string shop,

            string startDate)
        {
            ApiKey = apiKey;
            ApiSecret = apiSecret;
            Shop = shop;
            StartDate = startDate;
        }
    }
}