// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kafka
{
    [KafkaResourceType("kafka:index/userScramCredential:UserScramCredential")]
    public partial class UserScramCredential : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The password of the credential
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The number of SCRAM iterations used when generating the credential
        /// </summary>
        [Output("scramIterations")]
        public Output<int?> ScramIterations { get; private set; } = null!;

        /// <summary>
        /// The SCRAM mechanism used to generate the credential (SCRAM-SHA-256, SCRAM-SHA-512)
        /// </summary>
        [Output("scramMechanism")]
        public Output<string> ScramMechanism { get; private set; } = null!;

        /// <summary>
        /// The name of the credential
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a UserScramCredential resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserScramCredential(string name, UserScramCredentialArgs args, CustomResourceOptions? options = null)
            : base("kafka:index/userScramCredential:UserScramCredential", name, args ?? new UserScramCredentialArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserScramCredential(string name, Input<string> id, UserScramCredentialState? state = null, CustomResourceOptions? options = null)
            : base("kafka:index/userScramCredential:UserScramCredential", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserScramCredential resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserScramCredential Get(string name, Input<string> id, UserScramCredentialState? state = null, CustomResourceOptions? options = null)
        {
            return new UserScramCredential(name, id, state, options);
        }
    }

    public sealed class UserScramCredentialArgs : global::Pulumi.ResourceArgs
    {
        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The password of the credential
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The number of SCRAM iterations used when generating the credential
        /// </summary>
        [Input("scramIterations")]
        public Input<int>? ScramIterations { get; set; }

        /// <summary>
        /// The SCRAM mechanism used to generate the credential (SCRAM-SHA-256, SCRAM-SHA-512)
        /// </summary>
        [Input("scramMechanism", required: true)]
        public Input<string> ScramMechanism { get; set; } = null!;

        /// <summary>
        /// The name of the credential
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public UserScramCredentialArgs()
        {
        }
        public static new UserScramCredentialArgs Empty => new UserScramCredentialArgs();
    }

    public sealed class UserScramCredentialState : global::Pulumi.ResourceArgs
    {
        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password of the credential
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The number of SCRAM iterations used when generating the credential
        /// </summary>
        [Input("scramIterations")]
        public Input<int>? ScramIterations { get; set; }

        /// <summary>
        /// The SCRAM mechanism used to generate the credential (SCRAM-SHA-256, SCRAM-SHA-512)
        /// </summary>
        [Input("scramMechanism")]
        public Input<string>? ScramMechanism { get; set; }

        /// <summary>
        /// The name of the credential
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public UserScramCredentialState()
        {
        }
        public static new UserScramCredentialState Empty => new UserScramCredentialState();
    }
}
