// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSalesforceConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enter your Salesforce developer application's &lt;a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG"&gt;Client ID&lt;/a&gt;
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Enter your Salesforce developer application's &lt;a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG"&gt;Client secret&lt;/a&gt;
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// Toggle to use Bulk API (this might cause empty fields for some streams). Default: false
        /// </summary>
        [Input("forceUseBulkApi")]
        public Input<bool>? ForceUseBulkApi { get; set; }

        /// <summary>
        /// Toggle if you're using a &lt;a href="https://help.salesforce.com/s/articleView?id=sf.deploy_sandboxes_parent.htm&amp;type=5"&gt;Salesforce Sandbox&lt;/a&gt;. Default: false
        /// </summary>
        [Input("isSandbox")]
        public Input<bool>? IsSandbox { get; set; }

        [Input("refreshToken", required: true)]
        private Input<string>? _refreshToken;

        /// <summary>
        /// Enter your application's &lt;a href="https://developer.salesforce.com/docs/atlas.en-us.mobile_sdk.meta/mobile_sdk/oauth_refresh_token_flow.htm"&gt;Salesforce Refresh Token&lt;/a&gt; used for Airbyte to access your Salesforce account.
        /// </summary>
        public Input<string>? RefreshToken
        {
            get => _refreshToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _refreshToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enter the date (or date-time) in the YYYY-MM-DD or YYYY-MM-DDTHH:mm:ssZ format. Airbyte will replicate the data updated on and after this date. If this field is blank, Airbyte will replicate the data for last two years.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// The size of the time window (ISO8601 duration) to slice requests. Default: "P30D"
        /// </summary>
        [Input("streamSliceStep")]
        public Input<string>? StreamSliceStep { get; set; }

        [Input("streamsCriterias")]
        private InputList<Inputs.SourceSalesforceConfigurationStreamsCriteriaGetArgs>? _streamsCriterias;

        /// <summary>
        /// Add filters to select only required stream based on `SObject` name. Use this field to filter which tables are displayed by this connector. This is useful if your Salesforce account has a large number of tables (&gt;1000), in which case you may find it easier to navigate the UI and speed up the connector's performance if you restrict the tables displayed by this connector.
        /// </summary>
        public InputList<Inputs.SourceSalesforceConfigurationStreamsCriteriaGetArgs> StreamsCriterias
        {
            get => _streamsCriterias ?? (_streamsCriterias = new InputList<Inputs.SourceSalesforceConfigurationStreamsCriteriaGetArgs>());
            set => _streamsCriterias = value;
        }

        public SourceSalesforceConfigurationGetArgs()
        {
        }
        public static new SourceSalesforceConfigurationGetArgs Empty => new SourceSalesforceConfigurationGetArgs();
    }
}