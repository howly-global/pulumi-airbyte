// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSlackConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken")]
        public Input<Inputs.SourceSlackConfigurationCredentialsApiTokenGetArgs>? ApiToken { get; set; }

        [Input("signInViaSlackOAuth")]
        public Input<Inputs.SourceSlackConfigurationCredentialsSignInViaSlackOAuthGetArgs>? SignInViaSlackOAuth { get; set; }

        public SourceSlackConfigurationCredentialsGetArgs()
        {
        }
        public static new SourceSlackConfigurationCredentialsGetArgs Empty => new SourceSlackConfigurationCredentialsGetArgs();
    }
}