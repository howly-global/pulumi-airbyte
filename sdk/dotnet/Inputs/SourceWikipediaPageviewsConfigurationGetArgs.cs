// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceWikipediaPageviewsConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If you want to filter by access method, use one of desktop, mobile-app or mobile-web. If you are interested in pageviews regardless of access method, use all-access.
        /// </summary>
        [Input("access", required: true)]
        public Input<string> Access { get; set; } = null!;

        /// <summary>
        /// If you want to filter by agent type, use one of user, automated or spider. If you are interested in pageviews regardless of agent type, use all-agents.
        /// </summary>
        [Input("agent", required: true)]
        public Input<string> Agent { get; set; } = null!;

        /// <summary>
        /// The title of any article in the specified project. Any spaces should be replaced with underscores. It also should be URI-encoded, so that non-URI-safe characters like %, / or ? are accepted.
        /// </summary>
        [Input("article", required: true)]
        public Input<string> Article { get; set; } = null!;

        /// <summary>
        /// The ISO 3166-1 alpha-2 code of a country for which to retrieve top articles.
        /// </summary>
        [Input("country", required: true)]
        public Input<string> Country { get; set; } = null!;

        /// <summary>
        /// The date of the last day to include, in YYYYMMDD or YYYYMMDDHH format.
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// If you want to filter by project, use the domain of any Wikimedia project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The date of the first day to include, in YYYYMMDD or YYYYMMDDHH format.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public SourceWikipediaPageviewsConfigurationGetArgs()
        {
        }
        public static new SourceWikipediaPageviewsConfigurationGetArgs Empty => new SourceWikipediaPageviewsConfigurationGetArgs();
    }
}