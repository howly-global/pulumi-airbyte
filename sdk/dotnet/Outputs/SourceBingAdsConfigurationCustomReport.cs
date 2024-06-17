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
    public sealed class SourceBingAdsConfigurationCustomReport
    {
        /// <summary>
        /// The name of the custom report, this name would be used as stream name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of available aggregations. Default: "[Hourly]"
        /// </summary>
        public readonly string? ReportAggregation;
        /// <summary>
        /// A list of available report object columns. You can find it in description of reporting object that you want to add to custom report.
        /// </summary>
        public readonly ImmutableArray<string> ReportColumns;
        /// <summary>
        /// The name of the the object derives from the ReportRequest object. You can find it in Bing Ads Api docs - Reporting API - Reporting Data Objects. must be one of ["AccountPerformanceReportRequest", "AdDynamicTextPerformanceReportRequest", "AdExtensionByAdReportRequest", "AdExtensionByKeywordReportRequest", "AdExtensionDetailReportRequest", "AdGroupPerformanceReportRequest", "AdPerformanceReportRequest", "AgeGenderAudienceReportRequest", "AudiencePerformanceReportRequest", "CallDetailReportRequest", "CampaignPerformanceReportRequest", "ConversionPerformanceReportRequest", "DestinationUrlPerformanceReportRequest", "DSAAutoTargetPerformanceReportRequest", "DSACategoryPerformanceReportRequest", "DSASearchQueryPerformanceReportRequest", "GeographicPerformanceReportRequest", "GoalsAndFunnelsReportRequest", "HotelDimensionPerformanceReportRequest", "HotelGroupPerformanceReportRequest", "KeywordPerformanceReportRequest", "NegativeKeywordConflictReportRequest", "ProductDimensionPerformanceReportRequest", "ProductMatchCountReportRequest", "ProductNegativeKeywordConflictReportRequest", "ProductPartitionPerformanceReportRequest", "ProductPartitionUnitPerformanceReportRequest", "ProductSearchQueryPerformanceReportRequest", "ProfessionalDemographicsAudienceReportRequest", "PublisherUsagePerformanceReportRequest", "SearchCampaignChangeHistoryReportRequest", "SearchQueryPerformanceReportRequest", "ShareOfVoiceReportRequest", "UserLocationPerformanceReportRequest"]
        /// </summary>
        public readonly string ReportingObject;

        [OutputConstructor]
        private SourceBingAdsConfigurationCustomReport(
            string name,

            string? reportAggregation,

            ImmutableArray<string> reportColumns,

            string reportingObject)
        {
            Name = name;
            ReportAggregation = reportAggregation;
            ReportColumns = reportColumns;
            ReportingObject = reportingObject;
        }
    }
}