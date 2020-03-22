// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the kafka package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 *
 * > This content is derived from https://github.com/Mongey/terraform-provider-kafka/blob/master/website/docs/index.html.markdown.
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
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            if (!args || args.bootstrapServers === undefined) {
                throw new Error("Missing required property 'bootstrapServers'");
            }
            inputs["bootstrapServers"] = pulumi.output(args ? args.bootstrapServers : undefined).apply(JSON.stringify);
            inputs["caCert"] = (args ? args.caCert : undefined) || utilities.getEnv("KAFKA_CA_CERT");
            inputs["caCertFile"] = args ? args.caCertFile : undefined;
            inputs["clientCert"] = (args ? args.clientCert : undefined) || utilities.getEnv("KAFKA_CLIENT_CERT");
            inputs["clientCertFile"] = args ? args.clientCertFile : undefined;
            inputs["clientKey"] = (args ? args.clientKey : undefined) || utilities.getEnv("KAFKA_CLIENT_KEY");
            inputs["clientKeyFile"] = args ? args.clientKeyFile : undefined;
            inputs["saslMechanism"] = (args ? args.saslMechanism : undefined) || (utilities.getEnv("KAFKA_SASL_MECHANISM") || "plain");
            inputs["saslPassword"] = (args ? args.saslPassword : undefined) || utilities.getEnv("KAFKA_SASL_PASSWORD");
            inputs["saslUsername"] = (args ? args.saslUsername : undefined) || utilities.getEnv("KAFKA_SASL_USERNAME");
            inputs["skipTlsVerify"] = pulumi.output((args ? args.skipTlsVerify : undefined) || (utilities.getEnvBoolean("KAFKA_SKIP_VERIFY") || false)).apply(JSON.stringify);
            inputs["timeout"] = pulumi.output(args ? args.timeout : undefined).apply(JSON.stringify);
            inputs["tlsEnabled"] = pulumi.output((args ? args.tlsEnabled : undefined) || (utilities.getEnvBoolean("KAFKA_ENABLE_TLS") || true)).apply(JSON.stringify);
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * A list of kafka brokers
     */
    readonly bootstrapServers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * CA certificate file to validate the server's certificate.
     */
    readonly caCert?: pulumi.Input<string>;
    /**
     * Path to a CA certificate file to validate the server's certificate.
     * 
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.
     */
    readonly caCertFile?: pulumi.Input<string>;
    /**
     * The client certificate.
     */
    readonly clientCert?: pulumi.Input<string>;
    /**
     * Path to a file containing the client certificate.
     * 
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.
     */
    readonly clientCertFile?: pulumi.Input<string>;
    /**
     * The private key that the certificate was issued for.
     */
    readonly clientKey?: pulumi.Input<string>;
    /**
     * Path to a file containing the private key that the certificate was issued for.
     * 
     * @deprecated This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.
     */
    readonly clientKeyFile?: pulumi.Input<string>;
    /**
     * SASL mechanism, can be plain, scram-sha512, scram-sha256
     */
    readonly saslMechanism?: pulumi.Input<string>;
    /**
     * Password for SASL authentication.
     */
    readonly saslPassword?: pulumi.Input<string>;
    /**
     * Username for SASL authentication.
     */
    readonly saslUsername?: pulumi.Input<string>;
    /**
     * Set this to true only if the target Kafka server is an insecure development instance.
     */
    readonly skipTlsVerify?: pulumi.Input<boolean>;
    /**
     * Timeout in seconds
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * Enable communication with the Kafka Cluster over TLS.
     */
    readonly tlsEnabled?: pulumi.Input<boolean>;
}
