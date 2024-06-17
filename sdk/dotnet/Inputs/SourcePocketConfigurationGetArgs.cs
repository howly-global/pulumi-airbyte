// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePocketConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// The user's Pocket access token.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("consumerKey", required: true)]
        private Input<string>? _consumerKey;

        /// <summary>
        /// Your application's Consumer Key.
        /// </summary>
        public Input<string>? ConsumerKey
        {
            get => _consumerKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _consumerKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Select the content type of the items to retrieve. must be one of ["article", "video", "image"]
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// Select the granularity of the information about each item. must be one of ["simple", "complete"]
        /// </summary>
        [Input("detailType")]
        public Input<string>? DetailType { get; set; }

        /// <summary>
        /// Only return items from a particular `domain`.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Retrieve only favorited items. Default: false
        /// </summary>
        [Input("favorite")]
        public Input<bool>? Favorite { get; set; }

        /// <summary>
        /// Only return items whose title or url contain the `search` string.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// Only return items modified since the given timestamp.
        /// </summary>
        [Input("since")]
        public Input<string>? Since { get; set; }

        /// <summary>
        /// Sort retrieved items by the given criteria. must be one of ["newest", "oldest", "title", "site"]
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        /// <summary>
        /// Select the state of the items to retrieve. must be one of ["unread", "archive", "all"]
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Return only items tagged with this tag name. Use _untagged_ for retrieving only untagged items.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public SourcePocketConfigurationGetArgs()
        {
        }
        public static new SourcePocketConfigurationGetArgs Empty => new SourcePocketConfigurationGetArgs();
    }
}