// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationBigqueryConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Google BigQuery client's chunk (buffer) size (MIN=1, MAX = 15) for each table. The size that will be written by a single RPC. Written data will be buffered and only flushed upon reaching this size or closing the channel. The default 15MB value is used if not set explicitly. Read more &lt;a href="https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.client.Client.html"&gt;here&lt;/a&gt;. Default: 15
        /// </summary>
        [Input("bigQueryClientBufferSizeMb")]
        public Input<int>? BigQueryClientBufferSizeMb { get; set; }

        /// <summary>
        /// The contents of the JSON service account key. Check out the &lt;a href="https://docs.airbyte.com/integrations/destinations/bigquery#service-account-key"&gt;docs&lt;/a&gt; if you need help generating this key. Default credentials will be used if this field is left empty.
        /// </summary>
        [Input("credentialsJson")]
        public Input<string>? CredentialsJson { get; set; }

        /// <summary>
        /// The default BigQuery Dataset ID that tables are replicated to if the source does not specify a namespace. Read more &lt;a href="https://cloud.google.com/bigquery/docs/datasets#create-dataset"&gt;here&lt;/a&gt;.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// The location of the dataset. Warning: Changes made after creation will not be applied. Read more &lt;a href="https://cloud.google.com/bigquery/docs/locations"&gt;here&lt;/a&gt;. must be one of ["US", "EU", "asia-east1", "asia-east2", "asia-northeast1", "asia-northeast2", "asia-northeast3", "asia-south1", "asia-south2", "asia-southeast1", "asia-southeast2", "australia-southeast1", "australia-southeast2", "europe-central1", "europe-central2", "europe-north1", "europe-southwest1", "europe-west1", "europe-west2", "europe-west3", "europe-west4", "europe-west6", "europe-west7", "europe-west8", "europe-west9", "europe-west12", "me-central1", "me-central2", "me-west1", "northamerica-northeast1", "northamerica-northeast2", "southamerica-east1", "southamerica-west1", "us-central1", "us-east1", "us-east2", "us-east3", "us-east4", "us-east5", "us-south1", "us-west1", "us-west2", "us-west3", "us-west4"]
        /// </summary>
        [Input("datasetLocation", required: true)]
        public Input<string> DatasetLocation { get; set; } = null!;

        /// <summary>
        /// Disable Writing Final Tables. WARNING! The data format in _airbyte_data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions. Default: false
        /// </summary>
        [Input("disableTypeDedupe")]
        public Input<bool>? DisableTypeDedupe { get; set; }

        /// <summary>
        /// The way data will be uploaded to BigQuery.
        /// </summary>
        [Input("loadingMethod")]
        public Input<Inputs.DestinationBigqueryConfigurationLoadingMethodGetArgs>? LoadingMethod { get; set; }

        /// <summary>
        /// The GCP project ID for the project containing the target BigQuery dataset. Read more &lt;a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects"&gt;here&lt;/a&gt;.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The dataset to write raw tables into (default: airbyte_internal)
        /// </summary>
        [Input("rawDataDataset")]
        public Input<string>? RawDataDataset { get; set; }

        /// <summary>
        /// Interactive run type means that the query is executed as soon as possible, and these queries count towards concurrent rate limit and daily limit. Read more about interactive run type &lt;a href="https://cloud.google.com/bigquery/docs/running-queries#queries"&gt;here&lt;/a&gt;. Batch queries are queued and started as soon as idle resources are available in the BigQuery shared resource pool, which usually occurs within a few minutes. Batch queries don’t count towards your concurrent rate limit. Read more about batch queries &lt;a href="https://cloud.google.com/bigquery/docs/running-queries#batch"&gt;here&lt;/a&gt;. The default "interactive" value is used if not set explicitly. must be one of ["interactive", "batch"]; Default: "interactive"
        /// </summary>
        [Input("transformationPriority")]
        public Input<string>? TransformationPriority { get; set; }

        public DestinationBigqueryConfigurationGetArgs()
        {
        }
        public static new DestinationBigqueryConfigurationGetArgs Empty => new DestinationBigqueryConfigurationGetArgs();
    }
}
