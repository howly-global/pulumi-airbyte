// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGoogleSheetsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Google API Credentials for connecting to Google Sheets and Google Drive APIs
        /// </summary>
        [Input("credentials", required: true)]
        public Input<Inputs.DestinationGoogleSheetsConfigurationCredentialsArgs> Credentials { get; set; } = null!;

        /// <summary>
        /// The link to your spreadsheet. See &lt;a href='https://docs.airbyte.com/integrations/destinations/google-sheets#sheetlink'&gt;this guide&lt;/a&gt; for more details.
        /// </summary>
        [Input("spreadsheetId", required: true)]
        public Input<string> SpreadsheetId { get; set; } = null!;

        public DestinationGoogleSheetsConfigurationArgs()
        {
        }
        public static new DestinationGoogleSheetsConfigurationArgs Empty => new DestinationGoogleSheetsConfigurationArgs();
    }
}