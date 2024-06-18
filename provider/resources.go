// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package airbyte

import (
	"path"
	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	// Replace this provider with the provider you are bridging.
	airbyte "github.com/howly-global/terraform-provider-airbyte/provider"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"

	"github.com/howly-global/pulumi-airbyte/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "airbyte"
	// modules:
	mainMod = "index" // the airbyte module
)

//go:embed cmd/pulumi-resource-airbyte/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	computeSourceID := tfbridge.DelegateIDField("sourceId", "pulumi-airbyte", "https://github.com/howly-global/terraform-provider-airbyte")
	computeDestinationID := tfbridge.DelegateIDField("destinationId", "pulumi-airbyte", "https://github.com/howly-global/terraform-provider-airbyte")
	computeConnectionID := tfbridge.DelegateIDField("connectionId", "pulumi-airbyte", "https://github.com/howly-global/terraform-provider-airbyte")
	computeWorkspaceID := tfbridge.DelegateIDField("workspaceId", "pulumi-airbyte", "https://github.com/howly-global/terraform-provider-airbyte")
	computePermissionID := tfbridge.DelegateIDField("permissionId", "pulumi-airbyte", "https://github.com/howly-global/terraform-provider-airbyte")

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		//
		// The [pulumi-terraform-bridge](https://github.com/pulumi/pulumi-terraform-bridge) supports 3
		// types of Terraform providers:
		//
		// 1. Providers written with the terraform-plugin-sdk/v1:
		//
		//    If the provider you are bridging is written with the terraform-plugin-sdk/v1, then you
		//    will need to adapt the boilerplate:
		//
		//    - Change the import "shimv2" to "shimv1" and change the associated import to
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1".
		//
		//    You can then proceed as normal.
		//
		// 2. Providers written with terraform-plugin-sdk/v2:
		//
		//    This boilerplate is already geared towards providers written with the
		//    terraform-plugin-sdk/v2, since it is the most common provider framework used. No
		//    adaptions are needed.
		//
		// 3. Providers written with terraform-plugin-framework:
		//
		//    If the provider you are bridging is written with the terraform-plugin-framework, then
		//    you will need to adapt the boilerplate:
		//
		//    - Remove the `shimv2` import and add:
		//
		//      	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
		//
		//    - Replace `shimv2.NewProvider` with `pfbridge.ShimProvider`.
		//
		//    - In provider/cmd/pulumi-tfgen-airbyte/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/pf/tfgen". Remove the `version.Version`
		//      argument to `tfgen.Main`.
		//
		//    - In provider/cmd/pulumi-resource-airbyte/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge". Replace the arguments to the
		//      `tfbridge.Main` so it looks like this:
		//
		//      	tfbridge.Main(context.Background(), "airbyte", airbyte.Provider(),
		//			tfbridge.ProviderMetadata{PulumiSchema: pulumiSchema})
		//
		//   Detailed instructions can be found at
		//   https://github.com/pulumi/pulumi-terraform-bridge/blob/master/pf/README.md#how-to-upgrade-a-bridged-provider-to-plugin-framework.
		//   After that, you can proceed as normal.
		//
		// This is where you give the bridge a handle to the upstream terraform provider. SDKv2
		// convention is to have a function at "github.com/iwahbe/terraform-provider-airbyte/provider".New
		// which takes a version and produces a factory function. The provider you are bridging may
		// not do that. You will need to find the function (generally called in upstream's main.go)
		// that produces a:
		//
		// - *"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema".Provider (for SDKv2)
		// - *"github.com/hashicorp/terraform-plugin-sdk/v1/helper/schema".Provider (for SDKv1)
		// - "github.com/hashicorp/terraform-plugin-framework/provider".Provider (for plugin-framework)
		//
		//nolint:lll
		P: pfbridge.ShimProvider(airbyte.New(version.Version)()),

		Name:    "airbyte",
		Version: version.Version,
		// DisplayName is a way to be able to change the casing of the provider name when being
		// displayed on the Pulumi registry
		DisplayName: "",
		// Change this to your personal name (or a company name) that you would like to be shown in
		// the Pulumi Registry if this package is published there.
		Publisher: "howly-global",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/howly-global",
		Description:       "A Pulumi package for creating and managing airbyte cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"howly-global", "airbyte", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/howly-global/pulumi-airbyte",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this should
		// match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "howly-global",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config:       map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"airbyte_source_google_search_console":                    {ComputeID: computeSourceID},
			"airbyte_source_smartsheets":                              {ComputeID: computeSourceID},
			"airbyte_source_braze":                                    {ComputeID: computeSourceID},
			"airbyte_source_mongodb_v2":                               {ComputeID: computeSourceID},
			"airbyte_source_oracle":                                   {ComputeID: computeSourceID},
			"airbyte_source_secoda":                                   {ComputeID: computeSourceID},
			"airbyte_source_zenloop":                                  {ComputeID: computeSourceID},
			"airbyte_destination_qdrant":                              {ComputeID: computeDestinationID},
			"airbyte_source_iterable":                                 {ComputeID: computeSourceID},
			"airbyte_source_outbrain_amplify":                         {ComputeID: computeSourceID},
			"airbyte_source_wikipedia_pageviews":                      {ComputeID: computeSourceID},
			"airbyte_destination_elasticsearch":                       {ComputeID: computeDestinationID},
			"airbyte_source_getlago":                                  {ComputeID: computeSourceID},
			"airbyte_source_aws_cloudtrail":                           {ComputeID: computeSourceID},
			"airbyte_source_glassfrog":                                {ComputeID: computeSourceID},
			"airbyte_source_lemlist":                                  {ComputeID: computeSourceID},
			"airbyte_source_lever_hiring":                             {ComputeID: computeSourceID},
			"airbyte_source_salesloft":                                {ComputeID: computeSourceID},
			"airbyte_source_typeform":                                 {ComputeID: computeSourceID},
			"airbyte_destination_databricks":                          {ComputeID: computeDestinationID},
			"airbyte_source_microsoft_teams":                          {ComputeID: computeSourceID},
			"airbyte_source_google_webfonts":                          {ComputeID: computeSourceID},
			"airbyte_source_redshift":                                 {ComputeID: computeSourceID},
			"airbyte_source_sentry":                                   {ComputeID: computeSourceID},
			"airbyte_source_smartengage":                              {ComputeID: computeSourceID},
			"airbyte_destination_s3_glue":                             {ComputeID: computeDestinationID},
			"airbyte_source_google_analytics_data_api":                {ComputeID: computeSourceID},
			"airbyte_source_chartmogul":                               {ComputeID: computeSourceID},
			"airbyte_source_punk_api":                                 {ComputeID: computeSourceID},
			"airbyte_source_zoom":                                     {ComputeID: computeSourceID},
			"airbyte_source_bigquery":                                 {ComputeID: computeSourceID},
			"airbyte_source_github":                                   {ComputeID: computeSourceID},
			"airbyte_source_monday":                                   {ComputeID: computeSourceID},
			"airbyte_source_amazon_seller_partner":                    {ComputeID: computeSourceID},
			"airbyte_source_recreation":                               {ComputeID: computeSourceID},
			"airbyte_source_vantage":                                  {ComputeID: computeSourceID},
			"airbyte_source_recharge":                                 {ComputeID: computeSourceID},
			"airbyte_source_faker":                                    {ComputeID: computeSourceID},
			"airbyte_source_gnews":                                    {ComputeID: computeSourceID},
			"airbyte_source_apify_dataset":                            {ComputeID: computeSourceID},
			"airbyte_source_insightly":                                {ComputeID: computeSourceID},
			"airbyte_source_emailoctopus":                             {ComputeID: computeSourceID},
			"airbyte_source_hubplanner":                               {ComputeID: computeSourceID},
			"airbyte_source_snapchat_marketing":                       {ComputeID: computeSourceID},
			"airbyte_destination_dynamodb":                            {ComputeID: computeDestinationID},
			"airbyte_destination_mssql":                               {ComputeID: computeDestinationID},
			"airbyte_source_coin_api":                                 {ComputeID: computeSourceID},
			"airbyte_source_google_drive":                             {ComputeID: computeSourceID},
			"airbyte_destination_milvus":                              {ComputeID: computeDestinationID},
			"airbyte_source_google_ads":                               {ComputeID: computeSourceID},
			"airbyte_source_freshsales":                               {ComputeID: computeSourceID},
			"airbyte_source_gitlab":                                   {ComputeID: computeSourceID},
			"airbyte_source_launchdarkly":                             {ComputeID: computeSourceID},
			"airbyte_source_nytimes":                                  {ComputeID: computeSourceID},
			"airbyte_source_polygon_stock_api":                        {ComputeID: computeSourceID},
			"airbyte_source_tvmaze_schedule":                          {ComputeID: computeSourceID},
			"airbyte_destination_pubsub":                              {ComputeID: computeDestinationID},
			"airbyte_source_coinmarketcap":                            {ComputeID: computeSourceID},
			"airbyte_source_confluence":                               {ComputeID: computeSourceID},
			"airbyte_source_dremio":                                   {ComputeID: computeSourceID},
			"airbyte_source_lokalise":                                 {ComputeID: computeSourceID},
			"airbyte_source_spacex_api":                               {ComputeID: computeSourceID},
			"airbyte_source_zendesk_sell":                             {ComputeID: computeSourceID},
			"airbyte_destination_dev_null":                            {ComputeID: computeDestinationID},
			"airbyte_source_gainsight_px":                             {ComputeID: computeSourceID},
			"airbyte_source_amazon_sqs":                               {ComputeID: computeSourceID},
			"airbyte_source_appfollow":                                {ComputeID: computeSourceID},
			"airbyte_source_dixa":                                     {ComputeID: computeSourceID},
			"airbyte_source_linnworks":                                {ComputeID: computeSourceID},
			"airbyte_source_outreach":                                 {ComputeID: computeSourceID},
			"airbyte_source_paypal_transaction":                       {ComputeID: computeSourceID},
			"airbyte_source_pypi":                                     {ComputeID: computeSourceID},
			"airbyte_destination_teradata":                            {ComputeID: computeDestinationID},
			"airbyte_source_mssql":                                    {ComputeID: computeSourceID},
			"airbyte_source_jira":                                     {ComputeID: computeSourceID},
			"airbyte_source_kyve":                                     {ComputeID: computeSourceID},
			"airbyte_source_delighted":                                {ComputeID: computeSourceID},
			"airbyte_source_custom":                                   {ComputeID: computeSourceID},
			"airbyte_source_asana":                                    {ComputeID: computeSourceID},
			"airbyte_source_my_hours":                                 {ComputeID: computeSourceID},
			"airbyte_source_postgres":                                 {ComputeID: computeSourceID},
			"airbyte_source_tempo":                                    {ComputeID: computeSourceID},
			"airbyte_source_woocommerce":                              {ComputeID: computeSourceID},
			"airbyte_source_azure_blob_storage":                       {ComputeID: computeSourceID},
			"airbyte_destination_firestore":                           {ComputeID: computeDestinationID},
			"airbyte_source_facebook_marketing":                       {ComputeID: computeSourceID},
			"airbyte_source_qualaroo":                                 {ComputeID: computeSourceID},
			"airbyte_source_slack":                                    {ComputeID: computeSourceID},
			"airbyte_source_surveymonkey":                             {ComputeID: computeSourceID},
			"airbyte_destination_aws_datalake":                        {ComputeID: computeDestinationID},
			"airbyte_source_sendinblue":                               {ComputeID: computeSourceID},
			"airbyte_source_pocket":                                   {ComputeID: computeSourceID},
			"airbyte_source_salesforce":                               {ComputeID: computeSourceID},
			"airbyte_source_marketo":                                  {ComputeID: computeSourceID},
			"airbyte_source_postmarkapp":                              {ComputeID: computeSourceID},
			"airbyte_source_zendesk_support":                          {ComputeID: computeSourceID},
			"airbyte_source_dockerhub":                                {ComputeID: computeSourceID},
			"airbyte_source_fauna":                                    {ComputeID: computeSourceID},
			"airbyte_workspace":                                       {ComputeID: computeWorkspaceID},
			"airbyte_destination_convex":                              {ComputeID: computeDestinationID},
			"airbyte_source_amplitude":                                {ComputeID: computeSourceID},
			"airbyte_source_braintree":                                {ComputeID: computeSourceID},
			"airbyte_source_sftp":                                     {ComputeID: computeSourceID},
			"airbyte_source_square":                                   {ComputeID: computeSourceID},
			"airbyte_destination_langchain":                           {ComputeID: computeDestinationID},
			"airbyte_source_mailjet_sms":                              {ComputeID: computeSourceID},
			"airbyte_source_trello":                                   {ComputeID: computeSourceID},
			"airbyte_destination_yellowbrick":                         {ComputeID: computeDestinationID},
			"airbyte_source_microsoft_sharepoint":                     {ComputeID: computeSourceID},
			"airbyte_source_hubspot":                                  {ComputeID: computeSourceID},
			"airbyte_source_orb":                                      {ComputeID: computeSourceID},
			"airbyte_source_sendgrid":                                 {ComputeID: computeSourceID},
			"airbyte_destination_snowflake":                           {ComputeID: computeDestinationID},
			"airbyte_destination_sftp_json":                           {ComputeID: computeDestinationID},
			"airbyte_source_aha":                                      {ComputeID: computeSourceID},
			"airbyte_source_azure_table":                              {ComputeID: computeSourceID},
			"airbyte_source_mixpanel":                                 {ComputeID: computeSourceID},
			"airbyte_source_us_census":                                {ComputeID: computeSourceID},
			"airbyte_destination_duckdb":                              {ComputeID: computeDestinationID},
			"airbyte_destination_typesense":                           {ComputeID: computeDestinationID},
			"airbyte_destination_postgres":                            {ComputeID: computeDestinationID},
			"airbyte_destination_mysql":                               {ComputeID: computeDestinationID},
			"airbyte_source_file":                                     {ComputeID: computeSourceID},
			"airbyte_source_firebolt":                                 {ComputeID: computeSourceID},
			"airbyte_source_paystack":                                 {ComputeID: computeSourceID},
			"airbyte_source_senseforce":                               {ComputeID: computeSourceID},
			"airbyte_destination_custom":                              {ComputeID: computeDestinationID},
			"airbyte_source_clickhouse":                               {ComputeID: computeSourceID},
			"airbyte_source_dynamodb":                                 {ComputeID: computeSourceID},
			"airbyte_source_ip2whois":                                 {ComputeID: computeSourceID},
			"airbyte_source_shortio":                                  {ComputeID: computeSourceID},
			"airbyte_source_strava":                                   {ComputeID: computeSourceID},
			"airbyte_source_twilio_taskrouter":                        {ComputeID: computeSourceID},
			"airbyte_destination_redshift":                            {ComputeID: computeDestinationID},
			"airbyte_source_freshdesk":                                {ComputeID: computeSourceID},
			"airbyte_source_mailchimp":                                {ComputeID: computeSourceID},
			"airbyte_source_cart":                                     {ComputeID: computeSourceID},
			"airbyte_source_mailgun":                                  {ComputeID: computeSourceID},
			"airbyte_source_orbit":                                    {ComputeID: computeSourceID},
			"airbyte_source_recruitee":                                {ComputeID: computeSourceID},
			"airbyte_source_sap_fieldglass":                           {ComputeID: computeSourceID},
			"airbyte_permission":                                      {ComputeID: computePermissionID},
			"airbyte_source_onesignal":                                {ComputeID: computeSourceID},
			"airbyte_source_google_analytics_v4_service_account_only": {ComputeID: computeSourceID},
			"airbyte_source_harvest":                                  {ComputeID: computeSourceID},
			"airbyte_source_pinterest":                                {ComputeID: computeSourceID},
			"airbyte_destination_gcs":                                 {ComputeID: computeDestinationID},
			"airbyte_source_amazon_ads":                               {ComputeID: computeSourceID},
			"airbyte_source_klarna":                                   {ComputeID: computeSourceID},
			"airbyte_source_okta":                                     {ComputeID: computeSourceID},
			"airbyte_source_omnisend":                                 {ComputeID: computeSourceID},
			"airbyte_source_s3":                                       {ComputeID: computeSourceID},
			"airbyte_source_stripe":                                   {ComputeID: computeSourceID},
			"airbyte_source_youtube_analytics":                        {ComputeID: computeSourceID},
			"airbyte_destination_vectara":                             {ComputeID: computeDestinationID},
			"airbyte_source_zendesk_sunshine":                         {ComputeID: computeSourceID},
			"airbyte_source_twilio":                                   {ComputeID: computeSourceID},
			"airbyte_source_metabase":                                 {ComputeID: computeSourceID},
			"airbyte_source_gridly":                                   {ComputeID: computeSourceID},
			"airbyte_source_bamboo_hr":                                {ComputeID: computeSourceID},
			"airbyte_source_greenhouse":                               {ComputeID: computeSourceID},
			"airbyte_source_yotpo":                                    {ComputeID: computeSourceID},
			"airbyte_source_coda":                                     {ComputeID: computeSourceID},
			"airbyte_source_trustpilot":                               {ComputeID: computeSourceID},
			"airbyte_source_twitter":                                  {ComputeID: computeSourceID},
			"airbyte_source_netsuite":                                 {ComputeID: computeSourceID},
			"airbyte_destination_snowflake_cortex":                    {ComputeID: computeDestinationID},
			"airbyte_source_auth0":                                    {ComputeID: computeSourceID},
			"airbyte_source_gcs":                                      {ComputeID: computeSourceID},
			"airbyte_source_k6_cloud":                                 {ComputeID: computeSourceID},
			"airbyte_source_microsoft_onedrive":                       {ComputeID: computeSourceID},
			"airbyte_source_zoho_crm":                                 {ComputeID: computeSourceID},
			"airbyte_destination_clickhouse":                          {ComputeID: computeDestinationID},
			"airbyte_source_tiktok_marketing":                         {ComputeID: computeSourceID},
			"airbyte_source_pipedrive":                                {ComputeID: computeSourceID},
			"airbyte_source_exchange_rates":                           {ComputeID: computeSourceID},
			"airbyte_source_instatus":                                 {ComputeID: computeSourceID},
			"airbyte_destination_oracle":                              {ComputeID: computeDestinationID},
			"airbyte_source_prestashop":                               {ComputeID: computeSourceID},
			"airbyte_source_pexels_api":                               {ComputeID: computeSourceID},
			"airbyte_destination_pinecone":                            {ComputeID: computeDestinationID},
			"airbyte_source_airtable":                                 {ComputeID: computeSourceID},
			"airbyte_source_linkedin_pages":                           {ComputeID: computeSourceID},
			"airbyte_destination_google_sheets":                       {ComputeID: computeDestinationID},
			"airbyte_source_freshcaller":                              {ComputeID: computeSourceID},
			"airbyte_source_pendo":                                    {ComputeID: computeSourceID},
			"airbyte_source_persistiq":                                {ComputeID: computeSourceID},
			"airbyte_source_configcat":                                {ComputeID: computeSourceID},
			"airbyte_source_instagram":                                {ComputeID: computeSourceID},
			"airbyte_source_sonar_cloud":                              {ComputeID: computeSourceID},
			"airbyte_source_zendesk_chat":                             {ComputeID: computeSourceID},
			"airbyte_source_convex":                                   {ComputeID: computeSourceID},
			"airbyte_source_mysql":                                    {ComputeID: computeSourceID},
			"airbyte_source_notion":                                   {ComputeID: computeSourceID},
			"airbyte_source_survey_sparrow":                           {ComputeID: computeSourceID},
			"airbyte_source_zendesk_talk":                             {ComputeID: computeSourceID},
			"airbyte_source_mongodb_internal_poc":                     {ComputeID: computeSourceID},
			"airbyte_source_recurly":                                  {ComputeID: computeSourceID},
			"airbyte_source_close_com":                                {ComputeID: computeSourceID},
			"airbyte_source_webflow":                                  {ComputeID: computeSourceID},
			"airbyte_source_yandex_metrica":                           {ComputeID: computeSourceID},
			"airbyte_source_aircall":                                  {ComputeID: computeSourceID},
			"airbyte_source_rki_covid":                                {ComputeID: computeSourceID},
			"airbyte_source_railz":                                    {ComputeID: computeSourceID},
			"airbyte_destination_bigquery":                            {ComputeID: computeDestinationID},
			"airbyte_source_datascope":                                {ComputeID: computeSourceID},
			"airbyte_source_google_directory":                         {ComputeID: computeSourceID},
			"airbyte_source_intercom":                                 {ComputeID: computeSourceID},
			"airbyte_source_pokeapi":                                  {ComputeID: computeSourceID},
			"airbyte_destination_redis":                               {ComputeID: computeDestinationID},
			"airbyte_destination_azure_blob_storage":                  {ComputeID: computeDestinationID},
			"airbyte_source_clockify":                                 {ComputeID: computeSourceID},
			"airbyte_source_klaviyo":                                  {ComputeID: computeSourceID},
			"airbyte_source_posthog":                                  {ComputeID: computeSourceID},
			"airbyte_source_retently":                                 {ComputeID: computeSourceID},
			"airbyte_source_xkcd":                                     {ComputeID: computeSourceID},
			"airbyte_connection":                                      {ComputeID: computeConnectionID},
			"airbyte_source_google_pagespeed_insights":                {ComputeID: computeSourceID},
			"airbyte_source_linkedin_ads":                             {ComputeID: computeSourceID},
			"airbyte_source_snowflake":                                {ComputeID: computeSourceID},
			"airbyte_source_whisky_hunter":                            {ComputeID: computeSourceID},
			"airbyte_destination_mongodb":                             {ComputeID: computeDestinationID},
			"airbyte_source_sftp_bulk":                                {ComputeID: computeSourceID},
			"airbyte_source_shopify":                                  {ComputeID: computeSourceID},
			"airbyte_source_the_guardian_api":                         {ComputeID: computeSourceID},
			"airbyte_source_rss":                                      {ComputeID: computeSourceID},
			"airbyte_source_clickup_api":                              {ComputeID: computeSourceID},
			"airbyte_destination_s3":                                  {ComputeID: computeDestinationID},
			"airbyte_destination_weaviate":                            {ComputeID: computeDestinationID},
			"airbyte_source_bing_ads":                                 {ComputeID: computeSourceID},
			"airbyte_source_chargebee":                                {ComputeID: computeSourceID},
			"airbyte_source_google_sheets":                            {ComputeID: computeSourceID},
			"airbyte_source_smaily":                                   {ComputeID: computeSourceID},
			"airbyte_destination_astra":                               {ComputeID: computeDestinationID},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/howly-global/pulumi-airbyte/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("airbyte_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
