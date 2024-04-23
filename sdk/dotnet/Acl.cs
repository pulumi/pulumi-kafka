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
    /// A resource for managing Kafka ACLs.
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
    ///     var test = new Kafka.Acl("test", new()
    ///     {
    ///         AclResourceName = "syslog",
    ///         AclResourceType = "Topic",
    ///         AclPrincipal = "User:Alice",
    ///         AclHost = "*",
    ///         AclOperation = "Write",
    ///         AclPermissionType = "Deny",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ACLs can be imported using the following pattern
    /// 
    /// ```sh
    /// $ pulumi import kafka:index/acl:Acl test "acl_principal|acl_host|acl_operation|acl_permission_type|resource_type|resource_name|resource_pattern_type_filter"
    /// ```
    /// e.g.
    /// 
    /// ```sh
    /// $ pulumi import kafka:index/acl:Acl test "User:Alice|*|Write|Deny|Topic|syslog|Prefixed"
    /// ```
    /// </summary>
    [KafkaResourceType("kafka:index/acl:Acl")]
    public partial class Acl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Host from which principal listed in `acl_principal`
        /// will have access.
        /// </summary>
        [Output("aclHost")]
        public Output<string> AclHost { get; private set; } = null!;

        /// <summary>
        /// Operation that is being allowed or denied. Valid
        /// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
        /// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
        /// </summary>
        [Output("aclOperation")]
        public Output<string> AclOperation { get; private set; } = null!;

        /// <summary>
        /// Type of permission. Valid values are `Unknown`,
        /// `Any`, `Allow`, `Deny`.
        /// </summary>
        [Output("aclPermissionType")]
        public Output<string> AclPermissionType { get; private set; } = null!;

        /// <summary>
        /// Principal that is being allowed or denied.
        /// </summary>
        [Output("aclPrincipal")]
        public Output<string> AclPrincipal { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("aclResourceName")]
        public Output<string> AclResourceName { get; private set; } = null!;

        /// <summary>
        /// The type of resource. Valid values are `Unknown`,
        /// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
        /// </summary>
        [Output("aclResourceType")]
        public Output<string> AclResourceType { get; private set; } = null!;

        /// <summary>
        /// The pattern filter. Valid values
        /// are `Prefixed`, `Any`, `Match`, `Literal`. Default `Literal`.
        /// </summary>
        [Output("resourcePatternTypeFilter")]
        public Output<string?> ResourcePatternTypeFilter { get; private set; } = null!;


        /// <summary>
        /// Create a Acl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Acl(string name, AclArgs args, CustomResourceOptions? options = null)
            : base("kafka:index/acl:Acl", name, args ?? new AclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Acl(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
            : base("kafka:index/acl:Acl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Acl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Acl Get(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
        {
            return new Acl(name, id, state, options);
        }
    }

    public sealed class AclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host from which principal listed in `acl_principal`
        /// will have access.
        /// </summary>
        [Input("aclHost", required: true)]
        public Input<string> AclHost { get; set; } = null!;

        /// <summary>
        /// Operation that is being allowed or denied. Valid
        /// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
        /// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
        /// </summary>
        [Input("aclOperation", required: true)]
        public Input<string> AclOperation { get; set; } = null!;

        /// <summary>
        /// Type of permission. Valid values are `Unknown`,
        /// `Any`, `Allow`, `Deny`.
        /// </summary>
        [Input("aclPermissionType", required: true)]
        public Input<string> AclPermissionType { get; set; } = null!;

        /// <summary>
        /// Principal that is being allowed or denied.
        /// </summary>
        [Input("aclPrincipal", required: true)]
        public Input<string> AclPrincipal { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("aclResourceName", required: true)]
        public Input<string> AclResourceName { get; set; } = null!;

        /// <summary>
        /// The type of resource. Valid values are `Unknown`,
        /// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
        /// </summary>
        [Input("aclResourceType", required: true)]
        public Input<string> AclResourceType { get; set; } = null!;

        /// <summary>
        /// The pattern filter. Valid values
        /// are `Prefixed`, `Any`, `Match`, `Literal`. Default `Literal`.
        /// </summary>
        [Input("resourcePatternTypeFilter")]
        public Input<string>? ResourcePatternTypeFilter { get; set; }

        public AclArgs()
        {
        }
        public static new AclArgs Empty => new AclArgs();
    }

    public sealed class AclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host from which principal listed in `acl_principal`
        /// will have access.
        /// </summary>
        [Input("aclHost")]
        public Input<string>? AclHost { get; set; }

        /// <summary>
        /// Operation that is being allowed or denied. Valid
        /// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
        /// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
        /// </summary>
        [Input("aclOperation")]
        public Input<string>? AclOperation { get; set; }

        /// <summary>
        /// Type of permission. Valid values are `Unknown`,
        /// `Any`, `Allow`, `Deny`.
        /// </summary>
        [Input("aclPermissionType")]
        public Input<string>? AclPermissionType { get; set; }

        /// <summary>
        /// Principal that is being allowed or denied.
        /// </summary>
        [Input("aclPrincipal")]
        public Input<string>? AclPrincipal { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("aclResourceName")]
        public Input<string>? AclResourceName { get; set; }

        /// <summary>
        /// The type of resource. Valid values are `Unknown`,
        /// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
        /// </summary>
        [Input("aclResourceType")]
        public Input<string>? AclResourceType { get; set; }

        /// <summary>
        /// The pattern filter. Valid values
        /// are `Prefixed`, `Any`, `Match`, `Literal`. Default `Literal`.
        /// </summary>
        [Input("resourcePatternTypeFilter")]
        public Input<string>? ResourcePatternTypeFilter { get; set; }

        public AclState()
        {
        }
        public static new AclState Empty => new AclState();
    }
}
