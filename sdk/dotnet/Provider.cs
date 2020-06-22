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
    /// The provider type for the kafka package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("kafka", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        [Input("bootstrapServers", required: true, json: true)]
        private InputList<string>? _bootstrapServers;

        /// <summary>
        /// A list of kafka brokers
        /// </summary>
        public InputList<string> BootstrapServers
        {
            get => _bootstrapServers ?? (_bootstrapServers = new InputList<string>());
            set => _bootstrapServers = value;
        }

        /// <summary>
        /// CA certificate file to validate the server's certificate.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        /// <summary>
        /// Path to a CA certificate file to validate the server's certificate.
        /// </summary>
        [Input("caCertFile")]
        public Input<string>? CaCertFile { get; set; }

        /// <summary>
        /// The client certificate.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// Path to a file containing the client certificate.
        /// </summary>
        [Input("clientCertFile")]
        public Input<string>? ClientCertFile { get; set; }

        /// <summary>
        /// The private key that the certificate was issued for.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// Path to a file containing the private key that the certificate was issued for.
        /// </summary>
        [Input("clientKeyFile")]
        public Input<string>? ClientKeyFile { get; set; }

        /// <summary>
        /// The passphrase for the private key that the certificate was issued for.
        /// </summary>
        [Input("clientKeyPassphrase")]
        public Input<string>? ClientKeyPassphrase { get; set; }

        /// <summary>
        /// SASL mechanism, can be plain, scram-sha512, scram-sha256
        /// </summary>
        [Input("saslMechanism")]
        public Input<string>? SaslMechanism { get; set; }

        /// <summary>
        /// Password for SASL authentication.
        /// </summary>
        [Input("saslPassword")]
        public Input<string>? SaslPassword { get; set; }

        /// <summary>
        /// Username for SASL authentication.
        /// </summary>
        [Input("saslUsername")]
        public Input<string>? SaslUsername { get; set; }

        /// <summary>
        /// Set this to true only if the target Kafka server is an insecure development instance.
        /// </summary>
        [Input("skipTlsVerify", json: true)]
        public Input<bool>? SkipTlsVerify { get; set; }

        /// <summary>
        /// Timeout in seconds
        /// </summary>
        [Input("timeout", json: true)]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Enable communication with the Kafka Cluster over TLS.
        /// </summary>
        [Input("tlsEnabled", json: true)]
        public Input<bool>? TlsEnabled { get; set; }

        public ProviderArgs()
        {
            CaCert = Utilities.GetEnv("KAFKA_CA_CERT");
            ClientCert = Utilities.GetEnv("KAFKA_CLIENT_CERT");
            ClientKey = Utilities.GetEnv("KAFKA_CLIENT_KEY");
            SaslMechanism = Utilities.GetEnv("KAFKA_SASL_MECHANISM") ?? "plain";
            SaslPassword = Utilities.GetEnv("KAFKA_SASL_PASSWORD");
            SaslUsername = Utilities.GetEnv("KAFKA_SASL_USERNAME");
            SkipTlsVerify = Utilities.GetEnvBoolean("KAFKA_SKIP_VERIFY") ?? false;
            TlsEnabled = Utilities.GetEnvBoolean("KAFKA_ENABLE_TLS") ?? true;
        }
    }
}
