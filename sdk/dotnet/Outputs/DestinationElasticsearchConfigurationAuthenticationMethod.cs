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
    public sealed class DestinationElasticsearchConfigurationAuthenticationMethod
    {
        /// <summary>
        /// Use a api key and secret combination to authenticate
        /// </summary>
        public readonly Outputs.DestinationElasticsearchConfigurationAuthenticationMethodApiKeySecret? ApiKeySecret;
        /// <summary>
        /// Basic auth header with a username and password
        /// </summary>
        public readonly Outputs.DestinationElasticsearchConfigurationAuthenticationMethodUsernamePassword? UsernamePassword;

        [OutputConstructor]
        private DestinationElasticsearchConfigurationAuthenticationMethod(
            Outputs.DestinationElasticsearchConfigurationAuthenticationMethodApiKeySecret? apiKeySecret,

            Outputs.DestinationElasticsearchConfigurationAuthenticationMethodUsernamePassword? usernamePassword)
        {
            ApiKeySecret = apiKeySecret;
            UsernamePassword = usernamePassword;
        }
    }
}