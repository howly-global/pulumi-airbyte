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
    public sealed class SourceGitlabConfiguration
    {
        /// <summary>
        /// Please enter your basic URL from GitLab instance. Default: "gitlab.com"
        /// </summary>
        public readonly string? ApiUrl;
        public readonly Outputs.SourceGitlabConfigurationCredentials Credentials;
        /// <summary>
        /// [DEPRECATED] Space-delimited list of groups. e.g. airbyte.io.
        /// </summary>
        public readonly string? Groups;
        /// <summary>
        /// List of groups. e.g. airbyte.io.
        /// </summary>
        public readonly ImmutableArray<string> GroupsLists;
        /// <summary>
        /// [DEPRECATED] Space-delimited list of projects. e.g. airbyte.io/documentation meltano/tap-gitlab.
        /// </summary>
        public readonly string? Projects;
        /// <summary>
        /// Space-delimited list of projects. e.g. airbyte.io/documentation meltano/tap-gitlab.
        /// </summary>
        public readonly ImmutableArray<string> ProjectsLists;
        /// <summary>
        /// The date from which you'd like to replicate data for GitLab API, in the format YYYY-MM-DDT00:00:00Z. Optional. If not set, all data will be replicated. All data generated after this date will be replicated.
        /// </summary>
        public readonly string? StartDate;

        [OutputConstructor]
        private SourceGitlabConfiguration(
            string? apiUrl,

            Outputs.SourceGitlabConfigurationCredentials credentials,

            string? groups,

            ImmutableArray<string> groupsLists,

            string? projects,

            ImmutableArray<string> projectsLists,

            string? startDate)
        {
            ApiUrl = apiUrl;
            Credentials = credentials;
            Groups = groups;
            GroupsLists = groupsLists;
            Projects = projects;
            ProjectsLists = projectsLists;
            StartDate = startDate;
        }
    }
}