// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceLeverHiringConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticateViaLeverApiKey")]
        public Input<Inputs.SourceLeverHiringConfigurationCredentialsAuthenticateViaLeverApiKeyGetArgs>? AuthenticateViaLeverApiKey { get; set; }

        [Input("authenticateViaLeverOAuth")]
        public Input<Inputs.SourceLeverHiringConfigurationCredentialsAuthenticateViaLeverOAuthGetArgs>? AuthenticateViaLeverOAuth { get; set; }

        public SourceLeverHiringConfigurationCredentialsGetArgs()
        {
        }
        public static new SourceLeverHiringConfigurationCredentialsGetArgs Empty => new SourceLeverHiringConfigurationCredentialsGetArgs();
    }
}