// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterFromValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("doubleValue")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterFromValueDoubleValueArgs>? DoubleValue { get; set; }

        [Input("int64Value")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterFromValueInt64ValueArgs>? Int64Value { get; set; }

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterFromValueArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterFromValueArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterFromValueArgs();
    }
}