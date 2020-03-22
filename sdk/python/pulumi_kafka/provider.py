# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, bootstrap_servers=None, ca_cert=None, ca_cert_file=None, client_cert=None, client_cert_file=None, client_key=None, client_key_file=None, sasl_mechanism=None, sasl_password=None, sasl_username=None, skip_tls_verify=None, timeout=None, tls_enabled=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the kafka package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        > This content is derived from https://github.com/Mongey/terraform-provider-kafka/blob/master/website/docs/index.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] bootstrap_servers: A list of kafka brokers
        :param pulumi.Input[str] ca_cert: CA certificate file to validate the server's certificate.
        :param pulumi.Input[str] ca_cert_file: Path to a CA certificate file to validate the server's certificate.
        :param pulumi.Input[str] client_cert: The client certificate.
        :param pulumi.Input[str] client_cert_file: Path to a file containing the client certificate.
        :param pulumi.Input[str] client_key: The private key that the certificate was issued for.
        :param pulumi.Input[str] client_key_file: Path to a file containing the private key that the certificate was issued for.
        :param pulumi.Input[str] sasl_mechanism: SASL mechanism, can be plain, scram-sha512, scram-sha256
        :param pulumi.Input[str] sasl_password: Password for SASL authentication.
        :param pulumi.Input[str] sasl_username: Username for SASL authentication.
        :param pulumi.Input[bool] skip_tls_verify: Set this to true only if the target Kafka server is an insecure development instance.
        :param pulumi.Input[float] timeout: Timeout in seconds
        :param pulumi.Input[bool] tls_enabled: Enable communication with the Kafka Cluster over TLS.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if bootstrap_servers is None:
                raise TypeError("Missing required property 'bootstrap_servers'")
            __props__['bootstrap_servers'] = pulumi.Output.from_input(bootstrap_servers).apply(json.dumps) if bootstrap_servers is not None else None
            if ca_cert is None:
                ca_cert = utilities.get_env('KAFKA_CA_CERT')
            __props__['ca_cert'] = ca_cert
            __props__['ca_cert_file'] = ca_cert_file
            if client_cert is None:
                client_cert = utilities.get_env('KAFKA_CLIENT_CERT')
            __props__['client_cert'] = client_cert
            __props__['client_cert_file'] = client_cert_file
            if client_key is None:
                client_key = utilities.get_env('KAFKA_CLIENT_KEY')
            __props__['client_key'] = client_key
            __props__['client_key_file'] = client_key_file
            if sasl_mechanism is None:
                sasl_mechanism = (utilities.get_env('KAFKA_SASL_MECHANISM') or 'plain')
            __props__['sasl_mechanism'] = sasl_mechanism
            if sasl_password is None:
                sasl_password = utilities.get_env('KAFKA_SASL_PASSWORD')
            __props__['sasl_password'] = sasl_password
            if sasl_username is None:
                sasl_username = utilities.get_env('KAFKA_SASL_USERNAME')
            __props__['sasl_username'] = sasl_username
            if skip_tls_verify is None:
                skip_tls_verify = (utilities.get_env_bool('KAFKA_SKIP_VERIFY') or False)
            __props__['skip_tls_verify'] = pulumi.Output.from_input(skip_tls_verify).apply(json.dumps) if skip_tls_verify is not None else None
            __props__['timeout'] = pulumi.Output.from_input(timeout).apply(json.dumps) if timeout is not None else None
            if tls_enabled is None:
                tls_enabled = (utilities.get_env_bool('KAFKA_ENABLE_TLS') or True)
            __props__['tls_enabled'] = pulumi.Output.from_input(tls_enabled).apply(json.dumps) if tls_enabled is not None else None
        super(Provider, __self__).__init__(
            'kafka',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

