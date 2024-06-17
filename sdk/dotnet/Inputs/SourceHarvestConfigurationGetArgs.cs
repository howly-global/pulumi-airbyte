// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceHarvestConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Harvest account ID. Required for all Harvest requests in pair with Personal Access Token
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// Choose how to authenticate to Harvest.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.SourceHarvestConfigurationCredentialsGetArgs>? Credentials { get; set; }

        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data after this date will not be replicated.
        /// </summary>
        [Input("replicationEndDate")]
        public Input<string>? ReplicationEndDate { get; set; }

        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
        /// </summary>
        [Input("replicationStartDate", required: true)]
        public Input<string> ReplicationStartDate { get; set; } = null!;

        public SourceHarvestConfigurationGetArgs()
        {
        }
        public static new SourceHarvestConfigurationGetArgs Empty => new SourceHarvestConfigurationGetArgs();
    }
}