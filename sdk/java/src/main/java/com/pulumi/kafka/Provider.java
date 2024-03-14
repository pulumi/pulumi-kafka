// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kafka.ProviderArgs;
import com.pulumi.kafka.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the kafka package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:kafka")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * CA certificate file to validate the server&#39;s certificate.
     * 
     */
    @Export(name="caCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> caCert;

    /**
     * @return CA certificate file to validate the server&#39;s certificate.
     * 
     */
    public Output<Optional<String>> caCert() {
        return Codegen.optional(this.caCert);
    }
    /**
     * Path to a CA certificate file to validate the server&#39;s certificate.
     * 
     * @deprecated
     * This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
     * 
     */
    @Deprecated /* This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead. */
    @Export(name="caCertFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> caCertFile;

    /**
     * @return Path to a CA certificate file to validate the server&#39;s certificate.
     * 
     */
    public Output<Optional<String>> caCertFile() {
        return Codegen.optional(this.caCertFile);
    }
    /**
     * The client certificate.
     * 
     */
    @Export(name="clientCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientCert;

    /**
     * @return The client certificate.
     * 
     */
    public Output<Optional<String>> clientCert() {
        return Codegen.optional(this.clientCert);
    }
    /**
     * Path to a file containing the client certificate.
     * 
     * @deprecated
     * This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
     * 
     */
    @Deprecated /* This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead. */
    @Export(name="clientCertFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientCertFile;

    /**
     * @return Path to a file containing the client certificate.
     * 
     */
    public Output<Optional<String>> clientCertFile() {
        return Codegen.optional(this.clientCertFile);
    }
    /**
     * The private key that the certificate was issued for.
     * 
     */
    @Export(name="clientKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientKey;

    /**
     * @return The private key that the certificate was issued for.
     * 
     */
    public Output<Optional<String>> clientKey() {
        return Codegen.optional(this.clientKey);
    }
    /**
     * Path to a file containing the private key that the certificate was issued for.
     * 
     * @deprecated
     * This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
     * 
     */
    @Deprecated /* This parameter is now deprecated and will be removed in a later release, please use `client_key` instead. */
    @Export(name="clientKeyFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientKeyFile;

    /**
     * @return Path to a file containing the private key that the certificate was issued for.
     * 
     */
    public Output<Optional<String>> clientKeyFile() {
        return Codegen.optional(this.clientKeyFile);
    }
    /**
     * The passphrase for the private key that the certificate was issued for.
     * 
     */
    @Export(name="clientKeyPassphrase", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientKeyPassphrase;

    /**
     * @return The passphrase for the private key that the certificate was issued for.
     * 
     */
    public Output<Optional<String>> clientKeyPassphrase() {
        return Codegen.optional(this.clientKeyPassphrase);
    }
    /**
     * AWS profile name to use
     * 
     */
    @Export(name="saslAwsProfile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saslAwsProfile;

    /**
     * @return AWS profile name to use
     * 
     */
    public Output<Optional<String>> saslAwsProfile() {
        return Codegen.optional(this.saslAwsProfile);
    }
    /**
     * AWS region where MSK is deployed.
     * 
     */
    @Export(name="saslAwsRegion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saslAwsRegion;

    /**
     * @return AWS region where MSK is deployed.
     * 
     */
    public Output<Optional<String>> saslAwsRegion() {
        return Codegen.optional(this.saslAwsRegion);
    }
    /**
     * Arn of an AWS IAM role to assume
     * 
     */
    @Export(name="saslAwsRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saslAwsRoleArn;

    /**
     * @return Arn of an AWS IAM role to assume
     * 
     */
    public Output<Optional<String>> saslAwsRoleArn() {
        return Codegen.optional(this.saslAwsRoleArn);
    }
    /**
     * SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
     * 
     */
    @Export(name="saslMechanism", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saslMechanism;

    /**
     * @return SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
     * 
     */
    public Output<Optional<String>> saslMechanism() {
        return Codegen.optional(this.saslMechanism);
    }
    /**
     * Password for SASL authentication.
     * 
     */
    @Export(name="saslPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saslPassword;

    /**
     * @return Password for SASL authentication.
     * 
     */
    public Output<Optional<String>> saslPassword() {
        return Codegen.optional(this.saslPassword);
    }
    /**
     * The url to retrieve oauth2 tokens from, when using sasl mechanism oauthbearer
     * 
     */
    @Export(name="saslTokenUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saslTokenUrl;

    /**
     * @return The url to retrieve oauth2 tokens from, when using sasl mechanism oauthbearer
     * 
     */
    public Output<Optional<String>> saslTokenUrl() {
        return Codegen.optional(this.saslTokenUrl);
    }
    /**
     * Username for SASL authentication.
     * 
     */
    @Export(name="saslUsername", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saslUsername;

    /**
     * @return Username for SASL authentication.
     * 
     */
    public Output<Optional<String>> saslUsername() {
        return Codegen.optional(this.saslUsername);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kafka", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
