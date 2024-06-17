// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationClickhouseConfigurationTunnelMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("noTunnel")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodNoTunnelArgs>? NoTunnel { get; set; }

        [Input("passwordAuthentication")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodPasswordAuthenticationArgs>? PasswordAuthentication { get; set; }

        [Input("sshKeyAuthentication")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodSshKeyAuthenticationArgs>? SshKeyAuthentication { get; set; }

        public DestinationClickhouseConfigurationTunnelMethodArgs()
        {
        }
        public static new DestinationClickhouseConfigurationTunnelMethodArgs Empty => new DestinationClickhouseConfigurationTunnelMethodArgs();
    }
}