// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueInt64ValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueInt64ValueArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueInt64ValueArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueInt64ValueArgs();
    }
}