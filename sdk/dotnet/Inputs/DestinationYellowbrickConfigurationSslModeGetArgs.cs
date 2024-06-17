// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationYellowbrickConfigurationSslModeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow SSL mode.
        /// </summary>
        [Input("allow")]
        public Input<Inputs.DestinationYellowbrickConfigurationSslModeAllowGetArgs>? Allow { get; set; }

        /// <summary>
        /// Disable SSL.
        /// </summary>
        [Input("disable")]
        public Input<Inputs.DestinationYellowbrickConfigurationSslModeDisableGetArgs>? Disable { get; set; }

        /// <summary>
        /// Prefer SSL mode.
        /// </summary>
        [Input("prefer")]
        public Input<Inputs.DestinationYellowbrickConfigurationSslModePreferGetArgs>? Prefer { get; set; }

        /// <summary>
        /// Require SSL mode.
        /// </summary>
        [Input("require")]
        public Input<Inputs.DestinationYellowbrickConfigurationSslModeRequireGetArgs>? Require { get; set; }

        /// <summary>
        /// Verify-ca SSL mode.
        /// </summary>
        [Input("verifyCa")]
        public Input<Inputs.DestinationYellowbrickConfigurationSslModeVerifyCaGetArgs>? VerifyCa { get; set; }

        /// <summary>
        /// Verify-full SSL mode.
        /// </summary>
        [Input("verifyFull")]
        public Input<Inputs.DestinationYellowbrickConfigurationSslModeVerifyFullGetArgs>? VerifyFull { get; set; }

        public DestinationYellowbrickConfigurationSslModeGetArgs()
        {
        }
        public static new DestinationYellowbrickConfigurationSslModeGetArgs Empty => new DestinationYellowbrickConfigurationSslModeGetArgs();
    }
}