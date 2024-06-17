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
    public sealed class SourceS3ConfigurationFormatJsonl
    {
        /// <summary>
        /// The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors. Default: 0
        /// </summary>
        public readonly int? BlockSize;
        /// <summary>
        /// Whether newline characters are allowed in JSON values. Turning this on may affect performance. Leave blank to default to False. Default: false
        /// </summary>
        public readonly bool? NewlinesInValues;
        /// <summary>
        /// How JSON fields outside of explicit_schema (if given) are treated. Check &lt;a href="https://arrow.apache.org/docs/python/generated/pyarrow.json.ParseOptions.html" target="_blank"&gt;PyArrow documentation&lt;/a&gt; for details. must be one of ["ignore", "infer", "error"]; Default: "infer"
        /// </summary>
        public readonly string? UnexpectedFieldBehavior;

        [OutputConstructor]
        private SourceS3ConfigurationFormatJsonl(
            int? blockSize,

            bool? newlinesInValues,

            string? unexpectedFieldBehavior)
        {
            BlockSize = blockSize;
            NewlinesInValues = newlinesInValues;
            UnexpectedFieldBehavior = unexpectedFieldBehavior;
        }
    }
}