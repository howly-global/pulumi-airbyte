// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskChatConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        public Input<Inputs.SourceZendeskChatConfigurationCredentialsAccessTokenArgs>? AccessToken { get; set; }

        [Input("oAuth20")]
        public Input<Inputs.SourceZendeskChatConfigurationCredentialsOAuth20Args>? OAuth20 { get; set; }

        public SourceZendeskChatConfigurationCredentialsArgs()
        {
        }
        public static new SourceZendeskChatConfigurationCredentialsArgs Empty => new SourceZendeskChatConfigurationCredentialsArgs();
    }
}