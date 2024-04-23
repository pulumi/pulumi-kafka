// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kafka
{
    /// <summary>
    /// A resource for managing Kafka quotas.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Kafka = Pulumi.Kafka;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var quota = new Kafka.Quota("quota", new()
    ///     {
    ///         EntityName = "app_consumer",
    ///         EntityType = "client-id",
    ///         Config = 
    ///         {
    ///             { "consumer_byte_rate", "5000000" },
    ///             { "producer_byte_rate", "2500000" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [KafkaResourceType("kafka:index/quota:Quota")]
    public partial class Quota : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A map of string k/v attributes.
        /// </summary>
        [Output("config")]
        public Output<ImmutableDictionary<string, object>?> Config { get; private set; } = null!;

        /// <summary>
        /// The name of the entity to target.
        /// </summary>
        [Output("entityName")]
        public Output<string> EntityName { get; private set; } = null!;

        /// <summary>
        /// The type of entity. Valid values are `client-id`, `user`, `ip`.
        /// </summary>
        [Output("entityType")]
        public Output<string> EntityType { get; private set; } = null!;


        /// <summary>
        /// Create a Quota resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Quota(string name, QuotaArgs args, CustomResourceOptions? options = null)
            : base("kafka:index/quota:Quota", name, args ?? new QuotaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Quota(string name, Input<string> id, QuotaState? state = null, CustomResourceOptions? options = null)
            : base("kafka:index/quota:Quota", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Quota resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Quota Get(string name, Input<string> id, QuotaState? state = null, CustomResourceOptions? options = null)
        {
            return new Quota(name, id, state, options);
        }
    }

    public sealed class QuotaArgs : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<object>? _config;

        /// <summary>
        /// A map of string k/v attributes.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// The name of the entity to target.
        /// </summary>
        [Input("entityName", required: true)]
        public Input<string> EntityName { get; set; } = null!;

        /// <summary>
        /// The type of entity. Valid values are `client-id`, `user`, `ip`.
        /// </summary>
        [Input("entityType", required: true)]
        public Input<string> EntityType { get; set; } = null!;

        public QuotaArgs()
        {
        }
        public static new QuotaArgs Empty => new QuotaArgs();
    }

    public sealed class QuotaState : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<object>? _config;

        /// <summary>
        /// A map of string k/v attributes.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// The name of the entity to target.
        /// </summary>
        [Input("entityName")]
        public Input<string>? EntityName { get; set; }

        /// <summary>
        /// The type of entity. Valid values are `client-id`, `user`, `ip`.
        /// </summary>
        [Input("entityType")]
        public Input<string>? EntityType { get; set; }

        public QuotaState()
        {
        }
        public static new QuotaState Empty => new QuotaState();
    }
}
