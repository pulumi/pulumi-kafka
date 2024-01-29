// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the kafka package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'kafka';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * CA certificate file to validate the server's certificate.
     */
    public readonly caCert!: pulumi.Output<string | undefined>;
    /**
     * Path to a CA certificate file to validate the server's certificate.
     *
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
     */
    public readonly caCertFile!: pulumi.Output<string | undefined>;
    /**
     * The client certificate.
     */
    public readonly clientCert!: pulumi.Output<string | undefined>;
    /**
     * Path to a file containing the client certificate.
     *
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
     */
    public readonly clientCertFile!: pulumi.Output<string | undefined>;
    /**
     * The private key that the certificate was issued for.
     */
    public readonly clientKey!: pulumi.Output<string | undefined>;
    /**
     * Path to a file containing the private key that the certificate was issued for.
     *
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
     */
    public readonly clientKeyFile!: pulumi.Output<string | undefined>;
    /**
     * The passphrase for the private key that the certificate was issued for.
     */
    public readonly clientKeyPassphrase!: pulumi.Output<string | undefined>;
    /**
     * AWS region where MSK is deployed.
     */
    public readonly saslAwsRegion!: pulumi.Output<string | undefined>;
    /**
     * SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
     */
    public readonly saslMechanism!: pulumi.Output<string | undefined>;
    /**
     * Password for SASL authentication.
     */
    public readonly saslPassword!: pulumi.Output<string | undefined>;
    /**
     * Username for SASL authentication.
     */
    public readonly saslUsername!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.bootstrapServers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bootstrapServers'");
            }
            resourceInputs["bootstrapServers"] = pulumi.output(args ? args.bootstrapServers : undefined).apply(JSON.stringify);
            resourceInputs["caCert"] = args ? args.caCert : undefined;
            resourceInputs["caCertFile"] = args ? args.caCertFile : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["clientCertFile"] = args ? args.clientCertFile : undefined;
            resourceInputs["clientKey"] = args ? args.clientKey : undefined;
            resourceInputs["clientKeyFile"] = args ? args.clientKeyFile : undefined;
            resourceInputs["clientKeyPassphrase"] = args ? args.clientKeyPassphrase : undefined;
            resourceInputs["saslAwsRegion"] = args ? args.saslAwsRegion : undefined;
            resourceInputs["saslMechanism"] = (args ? args.saslMechanism : undefined) ?? (utilities.getEnv("KAFKA_SASL_MECHANISM") || "plain");
            resourceInputs["saslPassword"] = args ? args.saslPassword : undefined;
            resourceInputs["saslUsername"] = args ? args.saslUsername : undefined;
            resourceInputs["skipTlsVerify"] = pulumi.output((args ? args.skipTlsVerify : undefined) ?? (utilities.getEnvBoolean("KAFKA_SKIP_VERIFY") || false)).apply(JSON.stringify);
            resourceInputs["timeout"] = pulumi.output(args ? args.timeout : undefined).apply(JSON.stringify);
            resourceInputs["tlsEnabled"] = pulumi.output((args ? args.tlsEnabled : undefined) ?? (utilities.getEnvBoolean("KAFKA_ENABLE_TLS") || true)).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * A list of kafka brokers
     */
    bootstrapServers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * CA certificate file to validate the server's certificate.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Path to a CA certificate file to validate the server's certificate.
     *
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
     */
    caCertFile?: pulumi.Input<string>;
    /**
     * The client certificate.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Path to a file containing the client certificate.
     *
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
     */
    clientCertFile?: pulumi.Input<string>;
    /**
     * The private key that the certificate was issued for.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * Path to a file containing the private key that the certificate was issued for.
     *
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
     */
    clientKeyFile?: pulumi.Input<string>;
    /**
     * The passphrase for the private key that the certificate was issued for.
     */
    clientKeyPassphrase?: pulumi.Input<string>;
    /**
     * AWS region where MSK is deployed.
     */
    saslAwsRegion?: pulumi.Input<string>;
    /**
     * SASL mechanism, can be plain, scram-sha512, scram-sha256, aws-iam
     */
    saslMechanism?: pulumi.Input<string>;
    /**
     * Password for SASL authentication.
     */
    saslPassword?: pulumi.Input<string>;
    /**
     * Username for SASL authentication.
     */
    saslUsername?: pulumi.Input<string>;
    /**
     * Set this to true only if the target Kafka server is an insecure development instance.
     */
    skipTlsVerify?: pulumi.Input<boolean>;
    /**
     * Timeout in seconds
     */
    timeout?: pulumi.Input<number>;
    /**
     * Enable communication with the Kafka Cluster over TLS.
     */
    tlsEnabled?: pulumi.Input<boolean>;
}
