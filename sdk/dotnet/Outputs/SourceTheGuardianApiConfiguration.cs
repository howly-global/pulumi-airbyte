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
    public sealed class SourceTheGuardianApiConfiguration
    {
        /// <summary>
        /// Your API Key. See &lt;a href="https://open-platform.theguardian.com/access/"&gt;here&lt;/a&gt;. The key is case sensitive.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// (Optional) Use this to set the maximum date (YYYY-MM-DD) of the results. Results newer than the end_date will not be shown. Default is set to the current date (today) for incremental syncs.
        /// </summary>
        public readonly string? EndDate;
        /// <summary>
        /// (Optional) The query (q) parameter filters the results to only those that include that search term. The q parameter supports AND, OR and NOT operators.
        /// </summary>
        public readonly string? Query;
        /// <summary>
        /// (Optional) Use this to filter the results by a particular section. See &lt;a href="https://content.guardianapis.com/sections?api-key=test"&gt;here&lt;/a&gt; for a list of all sections, and &lt;a href="https://open-platform.theguardian.com/documentation/section"&gt;here&lt;/a&gt; for the sections endpoint documentation.
        /// </summary>
        public readonly string? Section;
        /// <summary>
        /// Use this to set the minimum date (YYYY-MM-DD) of the results. Results older than the start_date will not be shown.
        /// </summary>
        public readonly string StartDate;
        /// <summary>
        /// (Optional) A tag is a piece of data that is used by The Guardian to categorise content. Use this parameter to filter results by showing only the ones matching the entered tag. See &lt;a href="https://content.guardianapis.com/tags?api-key=test"&gt;here&lt;/a&gt; for a list of all tags, and &lt;a href="https://open-platform.theguardian.com/documentation/tag"&gt;here&lt;/a&gt; for the tags endpoint documentation.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private SourceTheGuardianApiConfiguration(
            string apiKey,

            string? endDate,

            string? query,

            string? section,

            string startDate,

            string? tag)
        {
            ApiKey = apiKey;
            EndDate = endDate;
            Query = query;
            Section = section;
            StartDate = startDate;
            Tag = tag;
        }
    }
}