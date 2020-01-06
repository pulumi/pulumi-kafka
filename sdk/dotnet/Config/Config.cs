// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Kafka
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("kafka");

        /// <summary>
        /// A list of kafka brokers
        /// </summary>
        public static ImmutableArray<string> BootstrapServers { get; set; } = __config.GetObject<ImmutableArray<string>>("bootstrapServers");

        /// <summary>
        /// CA certificate file to validate the server's certificate.
        /// </summary>
        public static string? CaCert { get; set; } = __config.Get("caCert") ?? Utilities.GetEnv("KAFKA_CA_CERT");

        /// <summary>
        /// Path to a CA certificate file to validate the server's certificate.
        /// </summary>
        public static string? CaCertFile { get; set; } = __config.Get("caCertFile");

        /// <summary>
        /// The client certificate.
        /// </summary>
        public static string? ClientCert { get; set; } = __config.Get("clientCert") ?? Utilities.GetEnv("KAFKA_CLIENT_CERT");

        /// <summary>
        /// Path to a file containing the client certificate.
        /// </summary>
        public static string? ClientCertFile { get; set; } = __config.Get("clientCertFile");

        /// <summary>
        /// The private key that the certificate was issued for.
        /// </summary>
        public static string? ClientKey { get; set; } = __config.Get("clientKey") ?? Utilities.GetEnv("KAFKA_CLIENT_KEY");

        /// <summary>
        /// Path to a file containing the private key that the certificate was issued for.
        /// </summary>
        public static string? ClientKeyFile { get; set; } = __config.Get("clientKeyFile");

        /// <summary>
        /// SASL mechanism, can be plain, scram-sha512, scram-sha256
        /// </summary>
        public static string? SaslMechanism { get; set; } = __config.Get("saslMechanism") ?? Utilities.GetEnv("KAFKA_SASL_MECHANISM") ?? "plain";

        /// <summary>
        /// Password for SASL authentication.
        /// </summary>
        public static string? SaslPassword { get; set; } = __config.Get("saslPassword") ?? Utilities.GetEnv("KAFKA_SASL_PASSWORD");

        /// <summary>
        /// Username for SASL authentication.
        /// </summary>
        public static string? SaslUsername { get; set; } = __config.Get("saslUsername") ?? Utilities.GetEnv("KAFKA_SASL_USERNAME");

        /// <summary>
        /// Set this to true only if the target Kafka server is an insecure development instance.
        /// </summary>
        public static bool? SkipTlsVerify { get; set; } = __config.GetBoolean("skipTlsVerify") ?? Utilities.GetEnvBoolean("KAFKA_SKIP_VERIFY") ?? false;

        /// <summary>
        /// Timeout in seconds
        /// </summary>
        public static int? Timeout { get; set; } = __config.GetInt32("timeout");

        /// <summary>
        /// Enable communication with the Kafka Cluster over TLS.
        /// </summary>
        public static bool? TlsEnabled { get; set; } = __config.GetBoolean("tlsEnabled") ?? Utilities.GetEnvBoolean("KAFKA_ENABLE_TLS") ?? true;

    }
    namespace ConfigTypes
    {
    }
}
