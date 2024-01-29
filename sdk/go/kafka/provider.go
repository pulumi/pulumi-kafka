// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-kafka/sdk/v3/go/kafka/internal"
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
	// AWS region where MSK is deployed.
	SaslAwsRegion pulumi.StringPtrOutput `pulumi:"saslAwsRegion"`
	// SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
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
	if args.SaslMechanism == nil {
		if d := internal.GetEnvOrDefault("plain", nil, "KAFKA_SASL_MECHANISM"); d != nil {
			args.SaslMechanism = pulumi.StringPtr(d.(string))
		}
	}
	if args.SkipTlsVerify == nil {
		if d := internal.GetEnvOrDefault(false, internal.ParseEnvBool, "KAFKA_SKIP_VERIFY"); d != nil {
			args.SkipTlsVerify = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.TlsEnabled == nil {
		if d := internal.GetEnvOrDefault(true, internal.ParseEnvBool, "KAFKA_ENABLE_TLS"); d != nil {
			args.TlsEnabled = pulumi.BoolPtr(d.(bool))
		}
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// AWS region where MSK is deployed.
	SaslAwsRegion *string `pulumi:"saslAwsRegion"`
	// SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
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
	// AWS region where MSK is deployed.
	SaslAwsRegion pulumi.StringPtrInput
	// SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
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

// CA certificate file to validate the server's certificate.
func (o ProviderOutput) CaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CaCert }).(pulumi.StringPtrOutput)
}

// Path to a CA certificate file to validate the server's certificate.
//
// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
func (o ProviderOutput) CaCertFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CaCertFile }).(pulumi.StringPtrOutput)
}

// The client certificate.
func (o ProviderOutput) ClientCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientCert }).(pulumi.StringPtrOutput)
}

// Path to a file containing the client certificate.
//
// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
func (o ProviderOutput) ClientCertFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientCertFile }).(pulumi.StringPtrOutput)
}

// The private key that the certificate was issued for.
func (o ProviderOutput) ClientKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientKey }).(pulumi.StringPtrOutput)
}

// Path to a file containing the private key that the certificate was issued for.
//
// Deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
func (o ProviderOutput) ClientKeyFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientKeyFile }).(pulumi.StringPtrOutput)
}

// The passphrase for the private key that the certificate was issued for.
func (o ProviderOutput) ClientKeyPassphrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientKeyPassphrase }).(pulumi.StringPtrOutput)
}

// AWS region where MSK is deployed.
func (o ProviderOutput) SaslAwsRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SaslAwsRegion }).(pulumi.StringPtrOutput)
}

// SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
func (o ProviderOutput) SaslMechanism() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SaslMechanism }).(pulumi.StringPtrOutput)
}

// Password for SASL authentication.
func (o ProviderOutput) SaslPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SaslPassword }).(pulumi.StringPtrOutput)
}

// Username for SASL authentication.
func (o ProviderOutput) SaslUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SaslUsername }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
