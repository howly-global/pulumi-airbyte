// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterNotExpressionExpressionFilterBetweenFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("fromValue", required: true)]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterNotExpressionExpressionFilterBetweenFilterFromValueArgs> FromValue { get; set; } = null!;

        [Input("toValue", required: true)]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterNotExpressionExpressionFilterBetweenFilterToValueArgs> ToValue { get; set; } = null!;

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterNotExpressionExpressionFilterBetweenFilterArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterNotExpressionExpressionFilterBetweenFilterArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterNotExpressionExpressionFilterBetweenFilterArgs();
    }
}