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
    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterNumericFilterValue
    {
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterNumericFilterValueDoubleValue? DoubleValue;
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterNumericFilterValueInt64Value? Int64Value;

        [OutputConstructor]
        private SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterNumericFilterValue(
            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterNumericFilterValueDoubleValue? doubleValue,

            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterNumericFilterValueInt64Value? int64Value)
        {
            DoubleValue = doubleValue;
            Int64Value = int64Value;
        }
    }
}