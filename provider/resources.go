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
	"context"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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

func computeID(ctx context.Context, state resource.PropertyMap) (resource.ID, error) {
	if id, has := state["id"]; has && id.IsString() {
		return resource.ID(id.StringValue()), nil
	}
	return "", fmt.Errorf("missing id property")
}

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
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
		PluginDownloadURL: "",
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
			"airbyte_source_google_search_console":                    {ComputeID: computeID},
			"airbyte_source_smartsheets":                              {ComputeID: computeID},
			"airbyte_source_braze":                                    {ComputeID: computeID},
			"airbyte_source_mongodb_v2":                               {ComputeID: computeID},
			"airbyte_source_oracle":                                   {ComputeID: computeID},
			"airbyte_source_secoda":                                   {ComputeID: computeID},
			"airbyte_source_zenloop":                                  {ComputeID: computeID},
			"airbyte_destination_qdrant":                              {ComputeID: computeID},
			"airbyte_source_iterable":                                 {ComputeID: computeID},
			"airbyte_source_outbrain_amplify":                         {ComputeID: computeID},
			"airbyte_source_wikipedia_pageviews":                      {ComputeID: computeID},
			"airbyte_destination_elasticsearch":                       {ComputeID: computeID},
			"airbyte_source_getlago":                                  {ComputeID: computeID},
			"airbyte_source_aws_cloudtrail":                           {ComputeID: computeID},
			"airbyte_source_glassfrog":                                {ComputeID: computeID},
			"airbyte_source_lemlist":                                  {ComputeID: computeID},
			"airbyte_source_lever_hiring":                             {ComputeID: computeID},
			"airbyte_source_salesloft":                                {ComputeID: computeID},
			"airbyte_source_typeform":                                 {ComputeID: computeID},
			"airbyte_destination_databricks":                          {ComputeID: computeID},
			"airbyte_source_microsoft_teams":                          {ComputeID: computeID},
			"airbyte_source_google_webfonts":                          {ComputeID: computeID},
			"airbyte_source_redshift":                                 {ComputeID: computeID},
			"airbyte_source_sentry":                                   {ComputeID: computeID},
			"airbyte_source_smartengage":                              {ComputeID: computeID},
			"airbyte_destination_s3_glue":                             {ComputeID: computeID},
			"airbyte_source_google_analytics_data_api":                {ComputeID: computeID},
			"airbyte_source_chartmogul":                               {ComputeID: computeID},
			"airbyte_source_punk_api":                                 {ComputeID: computeID},
			"airbyte_source_zoom":                                     {ComputeID: computeID},
			"airbyte_source_bigquery":                                 {ComputeID: computeID},
			"airbyte_source_github":                                   {ComputeID: computeID},
			"airbyte_source_monday":                                   {ComputeID: computeID},
			"airbyte_source_amazon_seller_partner":                    {ComputeID: computeID},
			"airbyte_source_recreation":                               {ComputeID: computeID},
			"airbyte_source_vantage":                                  {ComputeID: computeID},
			"airbyte_source_recharge":                                 {ComputeID: computeID},
			"airbyte_source_faker":                                    {ComputeID: computeID},
			"airbyte_source_gnews":                                    {ComputeID: computeID},
			"airbyte_source_apify_dataset":                            {ComputeID: computeID},
			"airbyte_source_insightly":                                {ComputeID: computeID},
			"airbyte_source_emailoctopus":                             {ComputeID: computeID},
			"airbyte_source_hubplanner":                               {ComputeID: computeID},
			"airbyte_source_snapchat_marketing":                       {ComputeID: computeID},
			"airbyte_destination_dynamodb":                            {ComputeID: computeID},
			"airbyte_destination_mssql":                               {ComputeID: computeID},
			"airbyte_source_coin_api":                                 {ComputeID: computeID},
			"airbyte_source_google_drive":                             {ComputeID: computeID},
			"airbyte_destination_milvus":                              {ComputeID: computeID},
			"airbyte_source_google_ads":                               {ComputeID: computeID},
			"airbyte_source_freshsales":                               {ComputeID: computeID},
			"airbyte_source_gitlab":                                   {ComputeID: computeID},
			"airbyte_source_launchdarkly":                             {ComputeID: computeID},
			"airbyte_source_nytimes":                                  {ComputeID: computeID},
			"airbyte_source_polygon_stock_api":                        {ComputeID: computeID},
			"airbyte_source_tvmaze_schedule":                          {ComputeID: computeID},
			"airbyte_destination_pubsub":                              {ComputeID: computeID},
			"airbyte_source_coinmarketcap":                            {ComputeID: computeID},
			"airbyte_source_confluence":                               {ComputeID: computeID},
			"airbyte_source_dremio":                                   {ComputeID: computeID},
			"airbyte_source_lokalise":                                 {ComputeID: computeID},
			"airbyte_source_spacex_api":                               {ComputeID: computeID},
			"airbyte_source_zendesk_sell":                             {ComputeID: computeID},
			"airbyte_destination_dev_null":                            {ComputeID: computeID},
			"airbyte_source_gainsight_px":                             {ComputeID: computeID},
			"airbyte_source_amazon_sqs":                               {ComputeID: computeID},
			"airbyte_source_appfollow":                                {ComputeID: computeID},
			"airbyte_source_dixa":                                     {ComputeID: computeID},
			"airbyte_source_linnworks":                                {ComputeID: computeID},
			"airbyte_source_outreach":                                 {ComputeID: computeID},
			"airbyte_source_paypal_transaction":                       {ComputeID: computeID},
			"airbyte_source_pypi":                                     {ComputeID: computeID},
			"airbyte_destination_teradata":                            {ComputeID: computeID},
			"airbyte_source_mssql":                                    {ComputeID: computeID},
			"airbyte_source_jira":                                     {ComputeID: computeID},
			"airbyte_source_kyve":                                     {ComputeID: computeID},
			"airbyte_source_delighted":                                {ComputeID: computeID},
			"airbyte_source_custom":                                   {ComputeID: computeID},
			"airbyte_source_asana":                                    {ComputeID: computeID},
			"airbyte_source_my_hours":                                 {ComputeID: computeID},
			"airbyte_source_postgres":                                 {ComputeID: computeID},
			"airbyte_source_tempo":                                    {ComputeID: computeID},
			"airbyte_source_woocommerce":                              {ComputeID: computeID},
			"airbyte_source_azure_blob_storage":                       {ComputeID: computeID},
			"airbyte_destination_firestore":                           {ComputeID: computeID},
			"airbyte_source_facebook_marketing":                       {ComputeID: computeID},
			"airbyte_source_qualaroo":                                 {ComputeID: computeID},
			"airbyte_source_slack":                                    {ComputeID: computeID},
			"airbyte_source_surveymonkey":                             {ComputeID: computeID},
			"airbyte_destination_aws_datalake":                        {ComputeID: computeID},
			"airbyte_source_sendinblue":                               {ComputeID: computeID},
			"airbyte_source_pocket":                                   {ComputeID: computeID},
			"airbyte_source_salesforce":                               {ComputeID: computeID},
			"airbyte_source_marketo":                                  {ComputeID: computeID},
			"airbyte_source_postmarkapp":                              {ComputeID: computeID},
			"airbyte_source_zendesk_support":                          {ComputeID: computeID},
			"airbyte_source_dockerhub":                                {ComputeID: computeID},
			"airbyte_source_fauna":                                    {ComputeID: computeID},
			"airbyte_workspace":                                       {ComputeID: computeID},
			"airbyte_destination_convex":                              {ComputeID: computeID},
			"airbyte_source_amplitude":                                {ComputeID: computeID},
			"airbyte_source_braintree":                                {ComputeID: computeID},
			"airbyte_source_sftp":                                     {ComputeID: computeID},
			"airbyte_source_square":                                   {ComputeID: computeID},
			"airbyte_destination_langchain":                           {ComputeID: computeID},
			"airbyte_source_mailjet_sms":                              {ComputeID: computeID},
			"airbyte_source_trello":                                   {ComputeID: computeID},
			"airbyte_destination_yellowbrick":                         {ComputeID: computeID},
			"airbyte_source_microsoft_sharepoint":                     {ComputeID: computeID},
			"airbyte_source_hubspot":                                  {ComputeID: computeID},
			"airbyte_source_orb":                                      {ComputeID: computeID},
			"airbyte_source_sendgrid":                                 {ComputeID: computeID},
			"airbyte_destination_snowflake":                           {ComputeID: computeID},
			"airbyte_destination_sftp_json":                           {ComputeID: computeID},
			"airbyte_source_aha":                                      {ComputeID: computeID},
			"airbyte_source_azure_table":                              {ComputeID: computeID},
			"airbyte_source_mixpanel":                                 {ComputeID: computeID},
			"airbyte_source_us_census":                                {ComputeID: computeID},
			"airbyte_destination_duckdb":                              {ComputeID: computeID},
			"airbyte_destination_typesense":                           {ComputeID: computeID},
			"airbyte_destination_postgres":                            {ComputeID: computeID},
			"airbyte_destination_mysql":                               {ComputeID: computeID},
			"airbyte_source_file":                                     {ComputeID: computeID},
			"airbyte_source_firebolt":                                 {ComputeID: computeID},
			"airbyte_source_paystack":                                 {ComputeID: computeID},
			"airbyte_source_senseforce":                               {ComputeID: computeID},
			"airbyte_destination_custom":                              {ComputeID: computeID},
			"airbyte_source_clickhouse":                               {ComputeID: computeID},
			"airbyte_source_dynamodb":                                 {ComputeID: computeID},
			"airbyte_source_ip2whois":                                 {ComputeID: computeID},
			"airbyte_source_shortio":                                  {ComputeID: computeID},
			"airbyte_source_strava":                                   {ComputeID: computeID},
			"airbyte_source_twilio_taskrouter":                        {ComputeID: computeID},
			"airbyte_destination_redshift":                            {ComputeID: computeID},
			"airbyte_source_freshdesk":                                {ComputeID: computeID},
			"airbyte_source_mailchimp":                                {ComputeID: computeID},
			"airbyte_source_cart":                                     {ComputeID: computeID},
			"airbyte_source_mailgun":                                  {ComputeID: computeID},
			"airbyte_source_orbit":                                    {ComputeID: computeID},
			"airbyte_source_recruitee":                                {ComputeID: computeID},
			"airbyte_source_sap_fieldglass":                           {ComputeID: computeID},
			"airbyte_permission":                                      {ComputeID: computeID},
			"airbyte_source_onesignal":                                {ComputeID: computeID},
			"airbyte_source_google_analytics_v4_service_account_only": {ComputeID: computeID},
			"airbyte_source_harvest":                                  {ComputeID: computeID},
			"airbyte_source_pinterest":                                {ComputeID: computeID},
			"airbyte_destination_gcs":                                 {ComputeID: computeID},
			"airbyte_source_amazon_ads":                               {ComputeID: computeID},
			"airbyte_source_klarna":                                   {ComputeID: computeID},
			"airbyte_source_okta":                                     {ComputeID: computeID},
			"airbyte_source_omnisend":                                 {ComputeID: computeID},
			"airbyte_source_s3":                                       {ComputeID: computeID},
			"airbyte_source_stripe":                                   {ComputeID: computeID},
			"airbyte_source_youtube_analytics":                        {ComputeID: computeID},
			"airbyte_destination_vectara":                             {ComputeID: computeID},
			"airbyte_source_zendesk_sunshine":                         {ComputeID: computeID},
			"airbyte_source_twilio":                                   {ComputeID: computeID},
			"airbyte_source_metabase":                                 {ComputeID: computeID},
			"airbyte_source_gridly":                                   {ComputeID: computeID},
			"airbyte_source_bamboo_hr":                                {ComputeID: computeID},
			"airbyte_source_greenhouse":                               {ComputeID: computeID},
			"airbyte_source_yotpo":                                    {ComputeID: computeID},
			"airbyte_source_coda":                                     {ComputeID: computeID},
			"airbyte_source_trustpilot":                               {ComputeID: computeID},
			"airbyte_source_twitter":                                  {ComputeID: computeID},
			"airbyte_source_netsuite":                                 {ComputeID: computeID},
			"airbyte_destination_snowflake_cortex":                    {ComputeID: computeID},
			"airbyte_source_auth0":                                    {ComputeID: computeID},
			"airbyte_source_gcs":                                      {ComputeID: computeID},
			"airbyte_source_k6_cloud":                                 {ComputeID: computeID},
			"airbyte_source_microsoft_onedrive":                       {ComputeID: computeID},
			"airbyte_source_zoho_crm":                                 {ComputeID: computeID},
			"airbyte_destination_clickhouse":                          {ComputeID: computeID},
			"airbyte_source_tiktok_marketing":                         {ComputeID: computeID},
			"airbyte_source_pipedrive":                                {ComputeID: computeID},
			"airbyte_source_exchange_rates":                           {ComputeID: computeID},
			"airbyte_source_instatus":                                 {ComputeID: computeID},
			"airbyte_destination_oracle":                              {ComputeID: computeID},
			"airbyte_source_prestashop":                               {ComputeID: computeID},
			"airbyte_source_pexels_api":                               {ComputeID: computeID},
			"airbyte_destination_pinecone":                            {ComputeID: computeID},
			"airbyte_source_airtable":                                 {ComputeID: computeID},
			"airbyte_source_linkedin_pages":                           {ComputeID: computeID},
			"airbyte_destination_google_sheets":                       {ComputeID: computeID},
			"airbyte_source_freshcaller":                              {ComputeID: computeID},
			"airbyte_source_pendo":                                    {ComputeID: computeID},
			"airbyte_source_persistiq":                                {ComputeID: computeID},
			"airbyte_source_configcat":                                {ComputeID: computeID},
			"airbyte_source_instagram":                                {ComputeID: computeID},
			"airbyte_source_sonar_cloud":                              {ComputeID: computeID},
			"airbyte_source_zendesk_chat":                             {ComputeID: computeID},
			"airbyte_source_convex":                                   {ComputeID: computeID},
			"airbyte_source_mysql":                                    {ComputeID: computeID},
			"airbyte_source_notion":                                   {ComputeID: computeID},
			"airbyte_source_survey_sparrow":                           {ComputeID: computeID},
			"airbyte_source_zendesk_talk":                             {ComputeID: computeID},
			"airbyte_source_mongodb_internal_poc":                     {ComputeID: computeID},
			"airbyte_source_recurly":                                  {ComputeID: computeID},
			"airbyte_source_close_com":                                {ComputeID: computeID},
			"airbyte_source_webflow":                                  {ComputeID: computeID},
			"airbyte_source_yandex_metrica":                           {ComputeID: computeID},
			"airbyte_source_aircall":                                  {ComputeID: computeID},
			"airbyte_source_rki_covid":                                {ComputeID: computeID},
			"airbyte_source_railz":                                    {ComputeID: computeID},
			"airbyte_destination_bigquery":                            {ComputeID: computeID},
			"airbyte_source_datascope":                                {ComputeID: computeID},
			"airbyte_source_google_directory":                         {ComputeID: computeID},
			"airbyte_source_intercom":                                 {ComputeID: computeID},
			"airbyte_source_pokeapi":                                  {ComputeID: computeID},
			"airbyte_destination_redis":                               {ComputeID: computeID},
			"airbyte_destination_azure_blob_storage":                  {ComputeID: computeID},
			"airbyte_source_clockify":                                 {ComputeID: computeID},
			"airbyte_source_klaviyo":                                  {ComputeID: computeID},
			"airbyte_source_posthog":                                  {ComputeID: computeID},
			"airbyte_source_retently":                                 {ComputeID: computeID},
			"airbyte_source_xkcd":                                     {ComputeID: computeID},
			"airbyte_connection":                                      {ComputeID: computeID},
			"airbyte_source_google_pagespeed_insights":                {ComputeID: computeID},
			"airbyte_source_linkedin_ads":                             {ComputeID: computeID},
			"airbyte_source_snowflake":                                {ComputeID: computeID},
			"airbyte_source_whisky_hunter":                            {ComputeID: computeID},
			"airbyte_destination_mongodb":                             {ComputeID: computeID},
			"airbyte_source_sftp_bulk":                                {ComputeID: computeID},
			"airbyte_source_shopify":                                  {ComputeID: computeID},
			"airbyte_source_the_guardian_api":                         {ComputeID: computeID},
			"airbyte_source_rss":                                      {ComputeID: computeID},
			"airbyte_source_clickup_api":                              {ComputeID: computeID},
			"airbyte_destination_s3":                                  {ComputeID: computeID},
			"airbyte_destination_weaviate":                            {ComputeID: computeID},
			"airbyte_source_bing_ads":                                 {ComputeID: computeID},
			"airbyte_source_chargebee":                                {ComputeID: computeID},
			"airbyte_source_google_sheets":                            {ComputeID: computeID},
			"airbyte_source_smaily":                                   {ComputeID: computeID},
			"airbyte_destination_astra":                               {ComputeID: computeID},
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
