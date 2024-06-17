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
    public sealed class SourceTempoConfiguration
    {
        /// <summary>
        /// Tempo API Token. Go to Tempo&gt;Settings, scroll down to Data Access and select API integration.
        /// </summary>
        public readonly string ApiToken;

        [OutputConstructor]
        private SourceTempoConfiguration(string apiToken)
        {
            ApiToken = apiToken;
        }
    }
}