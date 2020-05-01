// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kafka
{
    public partial class Acl : Pulumi.CustomResource
    {
        [Output("aclHost")]
        public Output<string> AclHost { get; private set; } = null!;

        [Output("aclOperation")]
        public Output<string> AclOperation { get; private set; } = null!;

        [Output("aclPermissionType")]
        public Output<string> AclPermissionType { get; private set; } = null!;

        [Output("aclPrincipal")]
        public Output<string> AclPrincipal { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("aclResourceName")]
        public Output<string> AclResourceName { get; private set; } = null!;

        [Output("aclResourceType")]
        public Output<string> AclResourceType { get; private set; } = null!;

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

    public sealed class AclArgs : Pulumi.ResourceArgs
    {
        [Input("aclHost", required: true)]
        public Input<string> AclHost { get; set; } = null!;

        [Input("aclOperation", required: true)]
        public Input<string> AclOperation { get; set; } = null!;

        [Input("aclPermissionType", required: true)]
        public Input<string> AclPermissionType { get; set; } = null!;

        [Input("aclPrincipal", required: true)]
        public Input<string> AclPrincipal { get; set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Input("aclResourceName", required: true)]
        public Input<string> AclResourceName { get; set; } = null!;

        [Input("aclResourceType", required: true)]
        public Input<string> AclResourceType { get; set; } = null!;

        [Input("resourcePatternTypeFilter")]
        public Input<string>? ResourcePatternTypeFilter { get; set; }

        public AclArgs()
        {
        }
    }

    public sealed class AclState : Pulumi.ResourceArgs
    {
        [Input("aclHost")]
        public Input<string>? AclHost { get; set; }

        [Input("aclOperation")]
        public Input<string>? AclOperation { get; set; }

        [Input("aclPermissionType")]
        public Input<string>? AclPermissionType { get; set; }

        [Input("aclPrincipal")]
        public Input<string>? AclPrincipal { get; set; }

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Input("aclResourceName")]
        public Input<string>? AclResourceName { get; set; }

        [Input("aclResourceType")]
        public Input<string>? AclResourceType { get; set; }

        [Input("resourcePatternTypeFilter")]
        public Input<string>? ResourcePatternTypeFilter { get; set; }

        public AclState()
        {
        }
    }
}
