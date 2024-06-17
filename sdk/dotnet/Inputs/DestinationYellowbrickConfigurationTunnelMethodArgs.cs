// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationYellowbrickConfigurationTunnelMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("noTunnel")]
        public Input<Inputs.DestinationYellowbrickConfigurationTunnelMethodNoTunnelArgs>? NoTunnel { get; set; }

        [Input("passwordAuthentication")]
        public Input<Inputs.DestinationYellowbrickConfigurationTunnelMethodPasswordAuthenticationArgs>? PasswordAuthentication { get; set; }

        [Input("sshKeyAuthentication")]
        public Input<Inputs.DestinationYellowbrickConfigurationTunnelMethodSshKeyAuthenticationArgs>? SshKeyAuthentication { get; set; }

        public DestinationYellowbrickConfigurationTunnelMethodArgs()
        {
        }
        public static new DestinationYellowbrickConfigurationTunnelMethodArgs Empty => new DestinationYellowbrickConfigurationTunnelMethodArgs();
    }
}