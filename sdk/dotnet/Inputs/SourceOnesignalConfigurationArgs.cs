// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceOnesignalConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("applications", required: true)]
        private InputList<Inputs.SourceOnesignalConfigurationApplicationArgs>? _applications;

        /// <summary>
        /// Applications keys, see the &lt;a href="https://documentation.onesignal.com/docs/accounts-and-keys"&gt;docs&lt;/a&gt; for more information on how to obtain this data
        /// </summary>
        public InputList<Inputs.SourceOnesignalConfigurationApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.SourceOnesignalConfigurationApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Comma-separated list of names and the value (sum/count) for the returned outcome data. See the &lt;a href="https://documentation.onesignal.com/reference/view-outcomes"&gt;docs&lt;/a&gt; for more details
        /// </summary>
        [Input("outcomeNames", required: true)]
        public Input<string> OutcomeNames { get; set; } = null!;

        /// <summary>
        /// The date from which you'd like to replicate data for OneSignal API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        [Input("userAuthKey", required: true)]
        private Input<string>? _userAuthKey;

        /// <summary>
        /// OneSignal User Auth Key, see the &lt;a href="https://documentation.onesignal.com/docs/accounts-and-keys#user-auth-key"&gt;docs&lt;/a&gt; for more information on how to obtain this key.
        /// </summary>
        public Input<string>? UserAuthKey
        {
            get => _userAuthKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _userAuthKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceOnesignalConfigurationArgs()
        {
        }
        public static new SourceOnesignalConfigurationArgs Empty => new SourceOnesignalConfigurationArgs();
    }
}