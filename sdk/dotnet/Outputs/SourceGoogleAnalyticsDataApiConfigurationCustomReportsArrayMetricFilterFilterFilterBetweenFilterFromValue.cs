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
    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValue
    {
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueDoubleValue? DoubleValue;
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueInt64Value? Int64Value;

        [OutputConstructor]
        private SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValue(
            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueDoubleValue? doubleValue,

            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterFilterFilterBetweenFilterFromValueInt64Value? int64Value)
        {
            DoubleValue = doubleValue;
            Int64Value = int64Value;
        }
    }
}