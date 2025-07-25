// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    [KafkaResourceType("pulumi:providers:kafka")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// CA certificate file to validate the server's certificate.
        /// </summary>
        [Output("caCert")]
        public Output<string?> CaCert { get; private set; } = null!;

        /// <summary>
        /// Path to a CA certificate file to validate the server's certificate.
        /// </summary>
        [Output("caCertFile")]
        public Output<string?> CaCertFile { get; private set; } = null!;

        /// <summary>
        /// The client certificate.
        /// </summary>
        [Output("clientCert")]
        public Output<string?> ClientCert { get; private set; } = null!;

        /// <summary>
        /// Path to a file containing the client certificate.
        /// </summary>
        [Output("clientCertFile")]
        public Output<string?> ClientCertFile { get; private set; } = null!;

        /// <summary>
        /// The private key that the certificate was issued for.
        /// </summary>
        [Output("clientKey")]
        public Output<string?> ClientKey { get; private set; } = null!;

        /// <summary>
        /// Path to a file containing the private key that the certificate was issued for.
        /// </summary>
        [Output("clientKeyFile")]
        public Output<string?> ClientKeyFile { get; private set; } = null!;

        /// <summary>
        /// The passphrase for the private key that the certificate was issued for.
        /// </summary>
        [Output("clientKeyPassphrase")]
        public Output<string?> ClientKeyPassphrase { get; private set; } = null!;

        /// <summary>
        /// The version of Kafka protocol to use in `$MAJOR.$MINOR.$PATCH` format. Some features may not be available on older
        /// versions. Default is 2.7.0.
        /// </summary>
        [Output("kafkaVersion")]
        public Output<string?> KafkaVersion { get; private set; } = null!;

        /// <summary>
        /// The AWS access key.
        /// </summary>
        [Output("saslAwsAccessKey")]
        public Output<string?> SaslAwsAccessKey { get; private set; } = null!;

        /// <summary>
        /// Path to a file containing the AWS pod identity authorization token
        /// </summary>
        [Output("saslAwsContainerAuthorizationTokenFile")]
        public Output<string?> SaslAwsContainerAuthorizationTokenFile { get; private set; } = null!;

        /// <summary>
        /// URI to retrieve AWS credentials from
        /// </summary>
        [Output("saslAwsContainerCredentialsFullUri")]
        public Output<string?> SaslAwsContainerCredentialsFullUri { get; private set; } = null!;

        /// <summary>
        /// External ID of the AWS IAM role to assume
        /// </summary>
        [Output("saslAwsExternalId")]
        public Output<string?> SaslAwsExternalId { get; private set; } = null!;

        /// <summary>
        /// AWS profile name to use
        /// </summary>
        [Output("saslAwsProfile")]
        public Output<string?> SaslAwsProfile { get; private set; } = null!;

        /// <summary>
        /// AWS region where MSK is deployed.
        /// </summary>
        [Output("saslAwsRegion")]
        public Output<string?> SaslAwsRegion { get; private set; } = null!;

        /// <summary>
        /// Arn of an AWS IAM role to assume
        /// </summary>
        [Output("saslAwsRoleArn")]
        public Output<string?> SaslAwsRoleArn { get; private set; } = null!;

        /// <summary>
        /// The AWS secret key.
        /// </summary>
        [Output("saslAwsSecretKey")]
        public Output<string?> SaslAwsSecretKey { get; private set; } = null!;

        /// <summary>
        /// The AWS session token. Only required if you are using temporary security credentials.
        /// </summary>
        [Output("saslAwsToken")]
        public Output<string?> SaslAwsToken { get; private set; } = null!;

        /// <summary>
        /// SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
        /// </summary>
        [Output("saslMechanism")]
        public Output<string?> SaslMechanism { get; private set; } = null!;

        /// <summary>
        /// Password for SASL authentication.
        /// </summary>
        [Output("saslPassword")]
        public Output<string?> SaslPassword { get; private set; } = null!;

        /// <summary>
        /// The url to retrieve oauth2 tokens from, when using sasl mechanism oauthbearer
        /// </summary>
        [Output("saslTokenUrl")]
        public Output<string?> SaslTokenUrl { get; private set; } = null!;

        /// <summary>
        /// Username for SASL authentication.
        /// </summary>
        [Output("saslUsername")]
        public Output<string?> SaslUsername { get; private set; } = null!;


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

        /// <summary>
        /// This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
        /// </summary>
        public global::Pulumi.Output<ProviderTerraformConfigResult> TerraformConfig()
            => global::Pulumi.Deployment.Instance.Call<ProviderTerraformConfigResult>("pulumi:providers:kafka/terraformConfig", CallArgs.Empty, this);
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
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
        /// The version of Kafka protocol to use in `$MAJOR.$MINOR.$PATCH` format. Some features may not be available on older
        /// versions. Default is 2.7.0.
        /// </summary>
        [Input("kafkaVersion")]
        public Input<string>? KafkaVersion { get; set; }

        /// <summary>
        /// The AWS access key.
        /// </summary>
        [Input("saslAwsAccessKey")]
        public Input<string>? SaslAwsAccessKey { get; set; }

        /// <summary>
        /// Path to a file containing the AWS pod identity authorization token
        /// </summary>
        [Input("saslAwsContainerAuthorizationTokenFile")]
        public Input<string>? SaslAwsContainerAuthorizationTokenFile { get; set; }

        /// <summary>
        /// URI to retrieve AWS credentials from
        /// </summary>
        [Input("saslAwsContainerCredentialsFullUri")]
        public Input<string>? SaslAwsContainerCredentialsFullUri { get; set; }

        /// <summary>
        /// Set this to true to turn AWS credentials debug.
        /// </summary>
        [Input("saslAwsCredsDebug", json: true)]
        public Input<bool>? SaslAwsCredsDebug { get; set; }

        /// <summary>
        /// External ID of the AWS IAM role to assume
        /// </summary>
        [Input("saslAwsExternalId")]
        public Input<string>? SaslAwsExternalId { get; set; }

        /// <summary>
        /// AWS profile name to use
        /// </summary>
        [Input("saslAwsProfile")]
        public Input<string>? SaslAwsProfile { get; set; }

        /// <summary>
        /// AWS region where MSK is deployed.
        /// </summary>
        [Input("saslAwsRegion")]
        public Input<string>? SaslAwsRegion { get; set; }

        /// <summary>
        /// Arn of an AWS IAM role to assume
        /// </summary>
        [Input("saslAwsRoleArn")]
        public Input<string>? SaslAwsRoleArn { get; set; }

        /// <summary>
        /// The AWS secret key.
        /// </summary>
        [Input("saslAwsSecretKey")]
        public Input<string>? SaslAwsSecretKey { get; set; }

        [Input("saslAwsSharedConfigFiles", json: true)]
        private InputList<string>? _saslAwsSharedConfigFiles;

        /// <summary>
        /// List of paths to AWS shared config files.
        /// </summary>
        public InputList<string> SaslAwsSharedConfigFiles
        {
            get => _saslAwsSharedConfigFiles ?? (_saslAwsSharedConfigFiles = new InputList<string>());
            set => _saslAwsSharedConfigFiles = value;
        }

        /// <summary>
        /// The AWS session token. Only required if you are using temporary security credentials.
        /// </summary>
        [Input("saslAwsToken")]
        public Input<string>? SaslAwsToken { get; set; }

        /// <summary>
        /// SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
        /// </summary>
        [Input("saslMechanism")]
        public Input<string>? SaslMechanism { get; set; }

        [Input("saslOauthScopes", json: true)]
        private InputList<string>? _saslOauthScopes;

        /// <summary>
        /// OAuth scopes to request when using the oauthbearer mechanism
        /// </summary>
        public InputList<string> SaslOauthScopes
        {
            get => _saslOauthScopes ?? (_saslOauthScopes = new InputList<string>());
            set => _saslOauthScopes = value;
        }

        /// <summary>
        /// Password for SASL authentication.
        /// </summary>
        [Input("saslPassword")]
        public Input<string>? SaslPassword { get; set; }

        /// <summary>
        /// The url to retrieve oauth2 tokens from, when using sasl mechanism oauthbearer
        /// </summary>
        [Input("saslTokenUrl")]
        public Input<string>? SaslTokenUrl { get; set; }

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
            SaslMechanism = Utilities.GetEnv("KAFKA_SASL_MECHANISM") ?? "plain";
            SkipTlsVerify = Utilities.GetEnvBoolean("KAFKA_SKIP_VERIFY") ?? false;
            TlsEnabled = Utilities.GetEnvBoolean("KAFKA_ENABLE_TLS") ?? true;
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }

    /// <summary>
    /// The results of the <see cref="Provider.TerraformConfig"/> method.
    /// </summary>
    [OutputType]
    public sealed class ProviderTerraformConfigResult
    {
        public readonly ImmutableDictionary<string, object> Result;

        [OutputConstructor]
        private ProviderTerraformConfigResult(ImmutableDictionary<string, object> result)
        {
            Result = result;
        }
    }
}
