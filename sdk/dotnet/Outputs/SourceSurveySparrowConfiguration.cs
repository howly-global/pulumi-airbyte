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
    public sealed class SourceSurveySparrowConfiguration
    {
        /// <summary>
        /// Your access token. See &lt;a href="https://developers.surveysparrow.com/rest-apis#authentication"&gt;here&lt;/a&gt;. The key is case sensitive.
        /// </summary>
        public readonly string AccessToken;
        /// <summary>
        /// Is your account location is EU based? If yes, the base url to retrieve data will be different.
        /// </summary>
        public readonly Outputs.SourceSurveySparrowConfigurationRegion? Region;
        /// <summary>
        /// A List of your survey ids for survey-specific stream
        /// </summary>
        public readonly ImmutableArray<string> SurveyIds;

        [OutputConstructor]
        private SourceSurveySparrowConfiguration(
            string accessToken,

            Outputs.SourceSurveySparrowConfigurationRegion? region,

            ImmutableArray<string> surveyIds)
        {
            AccessToken = accessToken;
            Region = region;
            SurveyIds = surveyIds;
        }
    }
}