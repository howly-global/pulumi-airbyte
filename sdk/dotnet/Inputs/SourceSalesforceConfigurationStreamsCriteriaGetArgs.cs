// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSalesforceConfigurationStreamsCriteriaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// must be one of ["starts with", "ends with", "contains", "exacts", "starts not with", "ends not with", "not contains", "not exacts"]; Default: "contains"
        /// </summary>
        [Input("criteria")]
        public Input<string>? Criteria { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public SourceSalesforceConfigurationStreamsCriteriaGetArgs()
        {
        }
        public static new SourceSalesforceConfigurationStreamsCriteriaGetArgs Empty => new SourceSalesforceConfigurationStreamsCriteriaGetArgs();
    }
}