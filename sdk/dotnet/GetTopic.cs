// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Kafka
{
    public static class GetTopic
    {
        public static Task<GetTopicResult> InvokeAsync(GetTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicResult>("kafka:index:getTopic", args ?? new GetTopicArgs(), options.WithVersion());

        public static Output<GetTopicResult> Invoke(GetTopicInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTopicResult>("kafka:index:getTopic", args ?? new GetTopicInvokeArgs(), options.WithVersion());
    }


    public sealed class GetTopicArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTopicArgs()
        {
        }
    }

    public sealed class GetTopicInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTopicInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicResult
    {
        public readonly ImmutableDictionary<string, object> Config;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly int Partitions;
        public readonly int ReplicationFactor;

        [OutputConstructor]
        private GetTopicResult(
            ImmutableDictionary<string, object> config,

            string id,

            string name,

            int partitions,

            int replicationFactor)
        {
            Config = config;
            Id = id;
            Name = name;
            Partitions = partitions;
            ReplicationFactor = replicationFactor;
        }
    }
}