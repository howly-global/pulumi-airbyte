// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSalesloftConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticateViaApiKey")]
        public Input<Inputs.SourceSalesloftConfigurationCredentialsAuthenticateViaApiKeyGetArgs>? AuthenticateViaApiKey { get; set; }

        [Input("authenticateViaOAuth")]
        public Input<Inputs.SourceSalesloftConfigurationCredentialsAuthenticateViaOAuthGetArgs>? AuthenticateViaOAuth { get; set; }

        public SourceSalesloftConfigurationCredentialsGetArgs()
        {
        }
        public static new SourceSalesloftConfigurationCredentialsGetArgs Empty => new SourceSalesloftConfigurationCredentialsGetArgs();
    }
}