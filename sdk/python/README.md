[![Build Status](https://travis-ci.com/pulumi/pulumi-kafka.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-kafka)

# Kafka Resource Provider

The Kafka resource provider for Pulumi lets you manage Kafka resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/kafka

or `yarn`:

    $ yarn add @pulumi/kafka

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_kafka

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-kafka/sdk/go/...

## Configuration

The following configuration points are available:

* kafka:bootstrapServers - (Required) A list of host:port addresses that will be used to discover the full set of alive brokers.
* kafka:caCert - (Optional) The CA certificate or path to a CA certificate file to validate the server's certificate.
* kafka:clientCert - (Optional) The client certificate or path to a file containing the client certificate -- Use for Client authentication to Kafka.
* kafka:clientKey - (Optional) The private key or path to a file containing the private key that the client certificate was issued for.
* kafka:skipTlsVerify - (Optional) Skip TLS verification. Default `false`.
* kafka:tlsEnabled - (Optional) Enable communication with the Kafka Cluster over TLS. Default `false`.
* kafka:saslUsername - (Optional) Username for SASL authentication.
* kafka:saslPassword - (Optional) Password for SASL authentication.
* kafka:saslMechanism - (Optional) Mechanism for SASL authentication. Allowed values are `plain`, `scram-sha512` and `scram-sha256`. Default `plain`.
* kafka:timeout - (Optional) Timeout in seconds. Default `120`.

## Reference

For detailed reference documentation, please visit [the API docs](https://pulumi.io/reference/pkg/nodejs/@pulumi/kafka/index.html).
