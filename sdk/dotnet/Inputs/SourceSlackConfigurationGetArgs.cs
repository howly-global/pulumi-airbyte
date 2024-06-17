// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSlackConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("channelFilters")]
        private InputList<string>? _channelFilters;

        /// <summary>
        /// A channel name list (without leading '#' char) which limit the channels from which you'd like to sync. Empty list means no filter.
        /// </summary>
        public InputList<string> ChannelFilters
        {
            get => _channelFilters ?? (_channelFilters = new InputList<string>());
            set => _channelFilters = value;
        }

        /// <summary>
        /// Choose how to authenticate into Slack
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.SourceSlackConfigurationCredentialsGetArgs>? Credentials { get; set; }

        /// <summary>
        /// Whether to read information from private channels that the bot is already in.  If false, only public channels will be read.  If true, the bot must be manually added to private channels. . Default: false
        /// </summary>
        [Input("includePrivateChannels")]
        public Input<bool>? IncludePrivateChannels { get; set; }

        /// <summary>
        /// Whether to join all channels or to sync data only from channels the bot is already in.  If false, you'll need to manually add the bot to all the channels from which you'd like to sync messages. . Default: true
        /// </summary>
        [Input("joinChannels")]
        public Input<bool>? JoinChannels { get; set; }

        /// <summary>
        /// How far into the past to look for messages in threads, default is 0 days. Default: 0
        /// </summary>
        [Input("lookbackWindow")]
        public Input<int>? LookbackWindow { get; set; }

        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceSlackConfigurationGetArgs()
        {
        }
        public static new SourceSlackConfigurationGetArgs Empty => new SourceSlackConfigurationGetArgs();
    }
}