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
    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValue
    {
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueDoubleValue? DoubleValue;
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueInt64Value? Int64Value;

        [OutputConstructor]
        private SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValue(
            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueDoubleValue? doubleValue,

            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueInt64Value? int64Value)
        {
            DoubleValue = doubleValue;
            Int64Value = int64Value;
        }
    }
}