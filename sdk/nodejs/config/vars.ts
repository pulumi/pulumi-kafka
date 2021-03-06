// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("kafka");

/**
 * A list of kafka brokers
 */
export let bootstrapServers: string[] | undefined = __config.getObject<string[]>("bootstrapServers");
/**
 * CA certificate file to validate the server's certificate.
 */
export let caCert: string | undefined = __config.get("caCert") || utilities.getEnv("KAFKA_CA_CERT");
/**
 * Path to a CA certificate file to validate the server's certificate.
 */
export let caCertFile: string | undefined = __config.get("caCertFile");
/**
 * The client certificate.
 */
export let clientCert: string | undefined = __config.get("clientCert") || utilities.getEnv("KAFKA_CLIENT_CERT");
/**
 * Path to a file containing the client certificate.
 */
export let clientCertFile: string | undefined = __config.get("clientCertFile");
/**
 * The private key that the certificate was issued for.
 */
export let clientKey: string | undefined = __config.get("clientKey") || utilities.getEnv("KAFKA_CLIENT_KEY");
/**
 * Path to a file containing the private key that the certificate was issued for.
 */
export let clientKeyFile: string | undefined = __config.get("clientKeyFile");
/**
 * The passphrase for the private key that the certificate was issued for.
 */
export let clientKeyPassphrase: string | undefined = __config.get("clientKeyPassphrase");
/**
 * SASL mechanism, can be plain, scram-sha512, scram-sha256
 */
export let saslMechanism: string | undefined = __config.get("saslMechanism") || (utilities.getEnv("KAFKA_SASL_MECHANISM") || "plain");
/**
 * Password for SASL authentication.
 */
export let saslPassword: string | undefined = __config.get("saslPassword") || utilities.getEnv("KAFKA_SASL_PASSWORD");
/**
 * Username for SASL authentication.
 */
export let saslUsername: string | undefined = __config.get("saslUsername") || utilities.getEnv("KAFKA_SASL_USERNAME");
/**
 * Set this to true only if the target Kafka server is an insecure development instance.
 */
export let skipTlsVerify: boolean | undefined = __config.getObject<boolean>("skipTlsVerify") || (<any>utilities.getEnvBoolean("KAFKA_SKIP_VERIFY") || false);
/**
 * Timeout in seconds
 */
export let timeout: number | undefined = __config.getObject<number>("timeout");
/**
 * Enable communication with the Kafka Cluster over TLS.
 */
export let tlsEnabled: boolean | undefined = __config.getObject<boolean>("tlsEnabled") || (<any>utilities.getEnvBoolean("KAFKA_ENABLE_TLS") || true);
