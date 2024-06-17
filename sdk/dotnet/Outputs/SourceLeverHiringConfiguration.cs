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
    public sealed class SourceLeverHiringConfiguration
    {
        /// <summary>
        /// Choose how to authenticate to Lever Hiring.
        /// </summary>
        public readonly Outputs.SourceLeverHiringConfigurationCredentials? Credentials;
        /// <summary>
        /// The environment in which you'd like to replicate data for Lever. This is used to determine which Lever API endpoint to use. must be one of ["Production", "Sandbox"]; Default: "Sandbox"
        /// </summary>
        public readonly string? Environment;
        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. Note that it will be used only in the following incremental streams: comments, commits, and issues.
        /// </summary>
        public readonly string StartDate;

        [OutputConstructor]
        private SourceLeverHiringConfiguration(
            Outputs.SourceLeverHiringConfigurationCredentials? credentials,

            string? environment,

            string startDate)
        {
            Credentials = credentials;
            Environment = environment;
            StartDate = startDate;
        }
    }
}