// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-kafka/sdk/v3/go/kafka/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// A list of kafka brokers
func GetBootstrapServers(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:bootstrapServers")
}

// CA certificate file to validate the server's certificate.
func GetCaCert(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:caCert")
}

// Path to a CA certificate file to validate the server's certificate.
//
// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
func GetCaCertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:caCertFile")
}

// The client certificate.
func GetClientCert(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:clientCert")
}

// Path to a file containing the client certificate.
//
// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
func GetClientCertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:clientCertFile")
}

// The private key that the certificate was issued for.
func GetClientKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:clientKey")
}

// Path to a file containing the private key that the certificate was issued for.
//
// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
func GetClientKeyFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:clientKeyFile")
}

// The passphrase for the private key that the certificate was issued for.
func GetClientKeyPassphrase(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:clientKeyPassphrase")
}

// SASL mechanism, can be plain, scram-sha512, scram-sha256
func GetSaslMechanism(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "kafka:saslMechanism")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault("plain", nil, "KAFKA_SASL_MECHANISM"); d != nil {
		value = d.(string)
	}
	return value
}

// Password for SASL authentication.
func GetSaslPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:saslPassword")
}

// Username for SASL authentication.
func GetSaslUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "kafka:saslUsername")
}

// Set this to true only if the target Kafka server is an insecure development instance.
func GetSkipTlsVerify(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "kafka:skipTlsVerify")
	if err == nil {
		return v
	}
	var value bool
	if d := internal.GetEnvOrDefault(false, internal.ParseEnvBool, "KAFKA_SKIP_VERIFY"); d != nil {
		value = d.(bool)
	}
	return value
}

// Timeout in seconds
func GetTimeout(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "kafka:timeout")
}

// Enable communication with the Kafka Cluster over TLS.
func GetTlsEnabled(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "kafka:tlsEnabled")
	if err == nil {
		return v
	}
	var value bool
	if d := internal.GetEnvOrDefault(true, internal.ParseEnvBool, "KAFKA_ENABLE_TLS"); d != nil {
		value = d.(bool)
	}
	return value
}
