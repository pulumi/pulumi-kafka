// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the kafka package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// CA certificate file to validate the server's certificate.
	CaCert pulumi.StringPtrOutput `pulumi:"caCert"`
	// Path to a CA certificate file to validate the server's certificate.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
	CaCertFile pulumi.StringPtrOutput `pulumi:"caCertFile"`
	// The client certificate.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// Path to a file containing the client certificate.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
	ClientCertFile pulumi.StringPtrOutput `pulumi:"clientCertFile"`
	// The private key that the certificate was issued for.
	ClientKey pulumi.StringPtrOutput `pulumi:"clientKey"`
	// Path to a file containing the private key that the certificate was issued for.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
	ClientKeyFile pulumi.StringPtrOutput `pulumi:"clientKeyFile"`
	// The passphrase for the private key that the certificate was issued for.
	ClientKeyPassphrase pulumi.StringPtrOutput `pulumi:"clientKeyPassphrase"`
	// SASL mechanism, can be plain, scram-sha512, scram-sha256
	SaslMechanism pulumi.StringPtrOutput `pulumi:"saslMechanism"`
	// Password for SASL authentication.
	SaslPassword pulumi.StringPtrOutput `pulumi:"saslPassword"`
	// Username for SASL authentication.
	SaslUsername pulumi.StringPtrOutput `pulumi:"saslUsername"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BootstrapServers == nil {
		return nil, errors.New("invalid value for required argument 'BootstrapServers'")
	}
	if isZero(args.SaslMechanism) {
		args.SaslMechanism = pulumi.StringPtr(getEnvOrDefault("plain", nil, "KAFKA_SASL_MECHANISM").(string))
	}
	if isZero(args.SkipTlsVerify) {
		args.SkipTlsVerify = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "KAFKA_SKIP_VERIFY").(bool))
	}
	if isZero(args.TlsEnabled) {
		args.TlsEnabled = pulumi.BoolPtr(getEnvOrDefault(true, parseEnvBool, "KAFKA_ENABLE_TLS").(bool))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:kafka", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// A list of kafka brokers
	BootstrapServers []string `pulumi:"bootstrapServers"`
	// CA certificate file to validate the server's certificate.
	CaCert *string `pulumi:"caCert"`
	// Path to a CA certificate file to validate the server's certificate.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
	CaCertFile *string `pulumi:"caCertFile"`
	// The client certificate.
	ClientCert *string `pulumi:"clientCert"`
	// Path to a file containing the client certificate.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
	ClientCertFile *string `pulumi:"clientCertFile"`
	// The private key that the certificate was issued for.
	ClientKey *string `pulumi:"clientKey"`
	// Path to a file containing the private key that the certificate was issued for.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
	ClientKeyFile *string `pulumi:"clientKeyFile"`
	// The passphrase for the private key that the certificate was issued for.
	ClientKeyPassphrase *string `pulumi:"clientKeyPassphrase"`
	// SASL mechanism, can be plain, scram-sha512, scram-sha256
	SaslMechanism *string `pulumi:"saslMechanism"`
	// Password for SASL authentication.
	SaslPassword *string `pulumi:"saslPassword"`
	// Username for SASL authentication.
	SaslUsername *string `pulumi:"saslUsername"`
	// Set this to true only if the target Kafka server is an insecure development instance.
	SkipTlsVerify *bool `pulumi:"skipTlsVerify"`
	// Timeout in seconds
	Timeout *int `pulumi:"timeout"`
	// Enable communication with the Kafka Cluster over TLS.
	TlsEnabled *bool `pulumi:"tlsEnabled"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// A list of kafka brokers
	BootstrapServers pulumi.StringArrayInput
	// CA certificate file to validate the server's certificate.
	CaCert pulumi.StringPtrInput
	// Path to a CA certificate file to validate the server's certificate.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
	CaCertFile pulumi.StringPtrInput
	// The client certificate.
	ClientCert pulumi.StringPtrInput
	// Path to a file containing the client certificate.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
	ClientCertFile pulumi.StringPtrInput
	// The private key that the certificate was issued for.
	ClientKey pulumi.StringPtrInput
	// Path to a file containing the private key that the certificate was issued for.
	//
	// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
	ClientKeyFile pulumi.StringPtrInput
	// The passphrase for the private key that the certificate was issued for.
	ClientKeyPassphrase pulumi.StringPtrInput
	// SASL mechanism, can be plain, scram-sha512, scram-sha256
	SaslMechanism pulumi.StringPtrInput
	// Password for SASL authentication.
	SaslPassword pulumi.StringPtrInput
	// Username for SASL authentication.
	SaslUsername pulumi.StringPtrInput
	// Set this to true only if the target Kafka server is an insecure development instance.
	SkipTlsVerify pulumi.BoolPtrInput
	// Timeout in seconds
	Timeout pulumi.IntPtrInput
	// Enable communication with the Kafka Cluster over TLS.
	TlsEnabled pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
