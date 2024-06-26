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
    public sealed class SourceOrbitConfiguration
    {
        /// <summary>
        /// Authorizes you to work with Orbit workspaces associated with the token.
        /// </summary>
        public readonly string ApiToken;
        /// <summary>
        /// Date in the format 2022-06-26. Only load members whose last activities are after this date.
        /// </summary>
        public readonly string? StartDate;
        /// <summary>
        /// The unique name of the workspace that your API token is associated with.
        /// </summary>
        public readonly string Workspace;

        [OutputConstructor]
        private SourceOrbitConfiguration(
            string apiToken,

            string? startDate,

            string workspace)
        {
            ApiToken = apiToken;
            StartDate = startDate;
            Workspace = workspace;
        }
    }
}
