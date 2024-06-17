// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cohort reports creates a time series of user retention for the cohort.
        /// </summary>
        [Input("cohortSpec")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecGetArgs>? CohortSpec { get; set; }

        /// <summary>
        /// Dimensions filter
        /// </summary>
        [Input("dimensionFilter")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterGetArgs>? DimensionFilter { get; set; }

        [Input("dimensions", required: true)]
        private InputList<string>? _dimensions;

        /// <summary>
        /// A list of dimensions.
        /// </summary>
        public InputList<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Metrics filter
        /// </summary>
        [Input("metricFilter")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterGetArgs>? MetricFilter { get; set; }

        [Input("metrics", required: true)]
        private InputList<string>? _metrics;

        /// <summary>
        /// A list of metrics.
        /// </summary>
        public InputList<string> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<string>());
            set => _metrics = value;
        }

        /// <summary>
        /// The name of the custom report, this name would be used as stream name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayGetArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayGetArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayGetArgs();
    }
}