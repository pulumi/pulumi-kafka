// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Kafka
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("kafka");

        private static readonly __Value<ImmutableArray<string>> _bootstrapServers = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("bootstrapServers"));
        /// <summary>
        /// A list of kafka brokers
        /// </summary>
        public static ImmutableArray<string> BootstrapServers
        {
            get => _bootstrapServers.Get();
            set => _bootstrapServers.Set(value);
        }

        private static readonly __Value<string?> _caCert = new __Value<string?>(() => __config.Get("caCert"));
        /// <summary>
        /// CA certificate file to validate the server's certificate.
        /// </summary>
        public static string? CaCert
        {
            get => _caCert.Get();
            set => _caCert.Set(value);
        }

        private static readonly __Value<string?> _caCertFile = new __Value<string?>(() => __config.Get("caCertFile"));
        /// <summary>
        /// Path to a CA certificate file to validate the server's certificate.
        /// </summary>
        public static string? CaCertFile
        {
            get => _caCertFile.Get();
            set => _caCertFile.Set(value);
        }

        private static readonly __Value<string?> _clientCert = new __Value<string?>(() => __config.Get("clientCert"));
        /// <summary>
        /// The client certificate.
        /// </summary>
        public static string? ClientCert
        {
            get => _clientCert.Get();
            set => _clientCert.Set(value);
        }

        private static readonly __Value<string?> _clientCertFile = new __Value<string?>(() => __config.Get("clientCertFile"));
        /// <summary>
        /// Path to a file containing the client certificate.
        /// </summary>
        public static string? ClientCertFile
        {
            get => _clientCertFile.Get();
            set => _clientCertFile.Set(value);
        }

        private static readonly __Value<string?> _clientKey = new __Value<string?>(() => __config.Get("clientKey"));
        /// <summary>
        /// The private key that the certificate was issued for.
        /// </summary>
        public static string? ClientKey
        {
            get => _clientKey.Get();
            set => _clientKey.Set(value);
        }

        private static readonly __Value<string?> _clientKeyFile = new __Value<string?>(() => __config.Get("clientKeyFile"));
        /// <summary>
        /// Path to a file containing the private key that the certificate was issued for.
        /// </summary>
        public static string? ClientKeyFile
        {
            get => _clientKeyFile.Get();
            set => _clientKeyFile.Set(value);
        }

        private static readonly __Value<string?> _clientKeyPassphrase = new __Value<string?>(() => __config.Get("clientKeyPassphrase"));
        /// <summary>
        /// The passphrase for the private key that the certificate was issued for.
        /// </summary>
        public static string? ClientKeyPassphrase
        {
            get => _clientKeyPassphrase.Get();
            set => _clientKeyPassphrase.Set(value);
        }

        private static readonly __Value<string?> _kafkaVersion = new __Value<string?>(() => __config.Get("kafkaVersion"));
        /// <summary>
        /// The version of Kafka protocol to use in `$MAJOR.$MINOR.$PATCH` format. Some features may not be available on older
        /// versions. Default is 2.7.0.
        /// </summary>
        public static string? KafkaVersion
        {
            get => _kafkaVersion.Get();
            set => _kafkaVersion.Set(value);
        }

        private static readonly __Value<bool?> _saslAwsCredsDebug = new __Value<bool?>(() => __config.GetBoolean("saslAwsCredsDebug"));
        /// <summary>
        /// Set this to true to turn AWS credentials debug.
        /// </summary>
        public static bool? SaslAwsCredsDebug
        {
            get => _saslAwsCredsDebug.Get();
            set => _saslAwsCredsDebug.Set(value);
        }

        private static readonly __Value<string?> _saslAwsProfile = new __Value<string?>(() => __config.Get("saslAwsProfile"));
        /// <summary>
        /// AWS profile name to use
        /// </summary>
        public static string? SaslAwsProfile
        {
            get => _saslAwsProfile.Get();
            set => _saslAwsProfile.Set(value);
        }

        private static readonly __Value<string?> _saslAwsRegion = new __Value<string?>(() => __config.Get("saslAwsRegion"));
        /// <summary>
        /// AWS region where MSK is deployed.
        /// </summary>
        public static string? SaslAwsRegion
        {
            get => _saslAwsRegion.Get();
            set => _saslAwsRegion.Set(value);
        }

        private static readonly __Value<string?> _saslAwsRoleArn = new __Value<string?>(() => __config.Get("saslAwsRoleArn"));
        /// <summary>
        /// Arn of an AWS IAM role to assume
        /// </summary>
        public static string? SaslAwsRoleArn
        {
            get => _saslAwsRoleArn.Get();
            set => _saslAwsRoleArn.Set(value);
        }

        private static readonly __Value<string?> _saslMechanism = new __Value<string?>(() => __config.Get("saslMechanism") ?? Utilities.GetEnv("KAFKA_SASL_MECHANISM") ?? "plain");
        /// <summary>
        /// SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
        /// </summary>
        public static string? SaslMechanism
        {
            get => _saslMechanism.Get();
            set => _saslMechanism.Set(value);
        }

        private static readonly __Value<string?> _saslPassword = new __Value<string?>(() => __config.Get("saslPassword"));
        /// <summary>
        /// Password for SASL authentication.
        /// </summary>
        public static string? SaslPassword
        {
            get => _saslPassword.Get();
            set => _saslPassword.Set(value);
        }

        private static readonly __Value<string?> _saslTokenUrl = new __Value<string?>(() => __config.Get("saslTokenUrl"));
        /// <summary>
        /// The url to retrieve oauth2 tokens from, when using sasl mechanism oauthbearer
        /// </summary>
        public static string? SaslTokenUrl
        {
            get => _saslTokenUrl.Get();
            set => _saslTokenUrl.Set(value);
        }

        private static readonly __Value<string?> _saslUsername = new __Value<string?>(() => __config.Get("saslUsername"));
        /// <summary>
        /// Username for SASL authentication.
        /// </summary>
        public static string? SaslUsername
        {
            get => _saslUsername.Get();
            set => _saslUsername.Set(value);
        }

        private static readonly __Value<bool?> _skipTlsVerify = new __Value<bool?>(() => __config.GetBoolean("skipTlsVerify") ?? Utilities.GetEnvBoolean("KAFKA_SKIP_VERIFY") ?? false);
        /// <summary>
        /// Set this to true only if the target Kafka server is an insecure development instance.
        /// </summary>
        public static bool? SkipTlsVerify
        {
            get => _skipTlsVerify.Get();
            set => _skipTlsVerify.Set(value);
        }

        private static readonly __Value<int?> _timeout = new __Value<int?>(() => __config.GetInt32("timeout"));
        /// <summary>
        /// Timeout in seconds
        /// </summary>
        public static int? Timeout
        {
            get => _timeout.Get();
            set => _timeout.Set(value);
        }

        private static readonly __Value<bool?> _tlsEnabled = new __Value<bool?>(() => __config.GetBoolean("tlsEnabled") ?? Utilities.GetEnvBoolean("KAFKA_ENABLE_TLS") ?? true);
        /// <summary>
        /// Enable communication with the Kafka Cluster over TLS.
        /// </summary>
        public static bool? TlsEnabled
        {
            get => _tlsEnabled.Get();
            set => _tlsEnabled.Set(value);
        }

    }
}
