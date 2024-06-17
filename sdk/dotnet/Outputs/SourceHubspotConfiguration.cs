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
    public sealed class SourceHubspotConfiguration
    {
        /// <summary>
        /// Choose how to authenticate to HubSpot.
        /// </summary>
        public readonly Outputs.SourceHubspotConfigurationCredentials Credentials;
        /// <summary>
        /// If enabled then experimental streams become available for sync. Default: false
        /// </summary>
        public readonly bool? EnableExperimentalStreams;
        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. If not set, "2006-06-01T00:00:00Z" (Hubspot creation date) will be used as start date. It's recommended to provide relevant to your data start date value to optimize synchronization.
        /// </summary>
        public readonly string? StartDate;

        [OutputConstructor]
        private SourceHubspotConfiguration(
            Outputs.SourceHubspotConfigurationCredentials credentials,

            bool? enableExperimentalStreams,

            string? startDate)
        {
            Credentials = credentials;
            EnableExperimentalStreams = enableExperimentalStreams;
            StartDate = startDate;
        }
    }
}