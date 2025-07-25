// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kafka.Outputs
{

    [OutputType]
    public sealed class GetTopicsListResult
    {
        /// <summary>
        /// A map of string k/v attributes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Config;
        /// <summary>
        /// Number of partitions.
        /// </summary>
        public readonly int Partitions;
        /// <summary>
        /// Number of replicas.
        /// </summary>
        public readonly int ReplicationFactor;
        /// <summary>
        /// The name of the topic.
        /// </summary>
        public readonly string TopicName;

        [OutputConstructor]
        private GetTopicsListResult(
            ImmutableDictionary<string, string> config,

            int partitions,

            int replicationFactor,

            string topicName)
        {
            Config = config;
            Partitions = partitions;
            ReplicationFactor = replicationFactor;
            TopicName = topicName;
        }
    }
}
