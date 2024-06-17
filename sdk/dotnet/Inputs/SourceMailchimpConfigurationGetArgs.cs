// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMailchimpConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        public Input<Inputs.SourceMailchimpConfigurationCredentialsGetArgs>? Credentials { get; set; }

        /// <summary>
        /// Technical fields used to identify datacenter to send request to
        /// </summary>
        [Input("dataCenter")]
        public Input<string>? DataCenter { get; set; }

        /// <summary>
        /// The date from which you want to start syncing data for Incremental streams. Only records that have been created or modified since this date will be synced. If left blank, all data will by synced.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceMailchimpConfigurationGetArgs()
        {
        }
        public static new SourceMailchimpConfigurationGetArgs Empty => new SourceMailchimpConfigurationGetArgs();
    }
}