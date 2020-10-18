# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bootstrap_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 ca_cert_file: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_cert_file: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 client_key_file: Optional[pulumi.Input[str]] = None,
                 client_key_passphrase: Optional[pulumi.Input[str]] = None,
                 sasl_mechanism: Optional[pulumi.Input[str]] = None,
                 sasl_password: Optional[pulumi.Input[str]] = None,
                 sasl_username: Optional[pulumi.Input[str]] = None,
                 skip_tls_verify: Optional[pulumi.Input[bool]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the kafka package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bootstrap_servers: A list of kafka brokers
        :param pulumi.Input[str] ca_cert: CA certificate file to validate the server's certificate.
        :param pulumi.Input[str] ca_cert_file: Path to a CA certificate file to validate the server's certificate.
        :param pulumi.Input[str] client_cert: The client certificate.
        :param pulumi.Input[str] client_cert_file: Path to a file containing the client certificate.
        :param pulumi.Input[str] client_key: The private key that the certificate was issued for.
        :param pulumi.Input[str] client_key_file: Path to a file containing the private key that the certificate was issued for.
        :param pulumi.Input[str] client_key_passphrase: The passphrase for the private key that the certificate was issued for.
        :param pulumi.Input[str] sasl_mechanism: SASL mechanism, can be plain, scram-sha512, scram-sha256
        :param pulumi.Input[str] sasl_password: Password for SASL authentication.
        :param pulumi.Input[str] sasl_username: Username for SASL authentication.
        :param pulumi.Input[bool] skip_tls_verify: Set this to true only if the target Kafka server is an insecure development instance.
        :param pulumi.Input[int] timeout: Timeout in seconds
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if bootstrap_servers is None:
                raise TypeError("Missing required property 'bootstrap_servers'")
            __props__['bootstrap_servers'] = pulumi.Output.from_input(bootstrap_servers).apply(pulumi.runtime.to_json) if bootstrap_servers is not None else None
            if ca_cert is None:
                ca_cert = _utilities.get_env('KAFKA_CA_CERT')
            __props__['ca_cert'] = ca_cert
            if ca_cert_file is not None:
                warnings.warn("This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.", DeprecationWarning)
                pulumi.log.warn("ca_cert_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.")
            __props__['ca_cert_file'] = ca_cert_file
            if client_cert is None:
                client_cert = _utilities.get_env('KAFKA_CLIENT_CERT')
            __props__['client_cert'] = client_cert
            if client_cert_file is not None:
                warnings.warn("This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.", DeprecationWarning)
                pulumi.log.warn("client_cert_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.")
            __props__['client_cert_file'] = client_cert_file
            if client_key is None:
                client_key = _utilities.get_env('KAFKA_CLIENT_KEY')
            __props__['client_key'] = client_key
            if client_key_file is not None:
                warnings.warn("This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.", DeprecationWarning)
                pulumi.log.warn("client_key_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.")
            __props__['client_key_file'] = client_key_file
            __props__['client_key_passphrase'] = client_key_passphrase
            if sasl_mechanism is None:
                sasl_mechanism = (_utilities.get_env('KAFKA_SASL_MECHANISM') or 'plain')
            __props__['sasl_mechanism'] = sasl_mechanism
            if sasl_password is None:
                sasl_password = _utilities.get_env('KAFKA_SASL_PASSWORD')
            __props__['sasl_password'] = sasl_password
            if sasl_username is None:
                sasl_username = _utilities.get_env('KAFKA_SASL_USERNAME')
            __props__['sasl_username'] = sasl_username
            if skip_tls_verify is None:
                skip_tls_verify = (_utilities.get_env_bool('KAFKA_SKIP_VERIFY') or False)
            __props__['skip_tls_verify'] = pulumi.Output.from_input(skip_tls_verify).apply(pulumi.runtime.to_json) if skip_tls_verify is not None else None
            __props__['timeout'] = pulumi.Output.from_input(timeout).apply(pulumi.runtime.to_json) if timeout is not None else None
            if tls_enabled is None:
                tls_enabled = (_utilities.get_env_bool('KAFKA_ENABLE_TLS') or True)
            __props__['tls_enabled'] = pulumi.Output.from_input(tls_enabled).apply(pulumi.runtime.to_json) if tls_enabled is not None else None
        super(Provider, __self__).__init__(
            'kafka',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

