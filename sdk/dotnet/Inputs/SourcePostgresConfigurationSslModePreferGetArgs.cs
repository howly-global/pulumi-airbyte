// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePostgresConfigurationSslModePreferGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parsed as JSON.
        /// </summary>
        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        public SourcePostgresConfigurationSslModePreferGetArgs()
        {
        }
        public static new SourcePostgresConfigurationSslModePreferGetArgs Empty => new SourcePostgresConfigurationSslModePreferGetArgs();
    }
}
