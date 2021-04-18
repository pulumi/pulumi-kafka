# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 bootstrap_servers: pulumi.Input[Sequence[pulumi.Input[str]]],
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
                 tls_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
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
        pulumi.set(__self__, "bootstrap_servers", bootstrap_servers)
        if ca_cert is not None:
            pulumi.set(__self__, "ca_cert", ca_cert)
        if ca_cert_file is not None:
            warnings.warn("""This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.""", DeprecationWarning)
            pulumi.log.warn("""ca_cert_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.""")
        if ca_cert_file is not None:
            pulumi.set(__self__, "ca_cert_file", ca_cert_file)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if client_cert_file is not None:
            warnings.warn("""This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.""", DeprecationWarning)
            pulumi.log.warn("""client_cert_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.""")
        if client_cert_file is not None:
            pulumi.set(__self__, "client_cert_file", client_cert_file)
        if client_key is not None:
            pulumi.set(__self__, "client_key", client_key)
        if client_key_file is not None:
            warnings.warn("""This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.""", DeprecationWarning)
            pulumi.log.warn("""client_key_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.""")
        if client_key_file is not None:
            pulumi.set(__self__, "client_key_file", client_key_file)
        if client_key_passphrase is not None:
            pulumi.set(__self__, "client_key_passphrase", client_key_passphrase)
        if sasl_mechanism is None:
            sasl_mechanism = (_utilities.get_env('KAFKA_SASL_MECHANISM') or 'plain')
        if sasl_mechanism is not None:
            pulumi.set(__self__, "sasl_mechanism", sasl_mechanism)
        if sasl_password is not None:
            pulumi.set(__self__, "sasl_password", sasl_password)
        if sasl_username is not None:
            pulumi.set(__self__, "sasl_username", sasl_username)
        if skip_tls_verify is None:
            skip_tls_verify = (_utilities.get_env_bool('KAFKA_SKIP_VERIFY') or False)
        if skip_tls_verify is not None:
            pulumi.set(__self__, "skip_tls_verify", skip_tls_verify)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if tls_enabled is None:
            tls_enabled = (_utilities.get_env_bool('KAFKA_ENABLE_TLS') or True)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter(name="bootstrapServers")
    def bootstrap_servers(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of kafka brokers
        """
        return pulumi.get(self, "bootstrap_servers")

    @bootstrap_servers.setter
    def bootstrap_servers(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "bootstrap_servers", value)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate file to validate the server's certificate.
        """
        return pulumi.get(self, "ca_cert")

    @ca_cert.setter
    def ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert", value)

    @property
    @pulumi.getter(name="caCertFile")
    def ca_cert_file(self) -> Optional[pulumi.Input[str]]:
        """
        Path to a CA certificate file to validate the server's certificate.
        """
        return pulumi.get(self, "ca_cert_file")

    @ca_cert_file.setter
    def ca_cert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert_file", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        The client certificate.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter(name="clientCertFile")
    def client_cert_file(self) -> Optional[pulumi.Input[str]]:
        """
        Path to a file containing the client certificate.
        """
        return pulumi.get(self, "client_cert_file")

    @client_cert_file.setter
    def client_cert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert_file", value)

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> Optional[pulumi.Input[str]]:
        """
        The private key that the certificate was issued for.
        """
        return pulumi.get(self, "client_key")

    @client_key.setter
    def client_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key", value)

    @property
    @pulumi.getter(name="clientKeyFile")
    def client_key_file(self) -> Optional[pulumi.Input[str]]:
        """
        Path to a file containing the private key that the certificate was issued for.
        """
        return pulumi.get(self, "client_key_file")

    @client_key_file.setter
    def client_key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key_file", value)

    @property
    @pulumi.getter(name="clientKeyPassphrase")
    def client_key_passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        The passphrase for the private key that the certificate was issued for.
        """
        return pulumi.get(self, "client_key_passphrase")

    @client_key_passphrase.setter
    def client_key_passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key_passphrase", value)

    @property
    @pulumi.getter(name="saslMechanism")
    def sasl_mechanism(self) -> Optional[pulumi.Input[str]]:
        """
        SASL mechanism, can be plain, scram-sha512, scram-sha256
        """
        return pulumi.get(self, "sasl_mechanism")

    @sasl_mechanism.setter
    def sasl_mechanism(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sasl_mechanism", value)

    @property
    @pulumi.getter(name="saslPassword")
    def sasl_password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for SASL authentication.
        """
        return pulumi.get(self, "sasl_password")

    @sasl_password.setter
    def sasl_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sasl_password", value)

    @property
    @pulumi.getter(name="saslUsername")
    def sasl_username(self) -> Optional[pulumi.Input[str]]:
        """
        Username for SASL authentication.
        """
        return pulumi.get(self, "sasl_username")

    @sasl_username.setter
    def sasl_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sasl_username", value)

    @property
    @pulumi.getter(name="skipTlsVerify")
    def skip_tls_verify(self) -> Optional[pulumi.Input[bool]]:
        """
        Set this to true only if the target Kafka server is an insecure development instance.
        """
        return pulumi.get(self, "skip_tls_verify")

    @skip_tls_verify.setter
    def skip_tls_verify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_tls_verify", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout in seconds
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable communication with the Kafka Cluster over TLS.
        """
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)


class Provider(pulumi.ProviderResource):
    @overload
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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the kafka package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
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
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if bootstrap_servers is None and not opts.urn:
                raise TypeError("Missing required property 'bootstrap_servers'")
            __props__.__dict__["bootstrap_servers"] = pulumi.Output.from_input(bootstrap_servers).apply(pulumi.runtime.to_json) if bootstrap_servers is not None else None
            __props__.__dict__["ca_cert"] = ca_cert
            if ca_cert_file is not None and not opts.urn:
                warnings.warn("""This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.""", DeprecationWarning)
                pulumi.log.warn("""ca_cert_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `ca_cert` instead.""")
            __props__.__dict__["ca_cert_file"] = ca_cert_file
            __props__.__dict__["client_cert"] = client_cert
            if client_cert_file is not None and not opts.urn:
                warnings.warn("""This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.""", DeprecationWarning)
                pulumi.log.warn("""client_cert_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_cert` instead.""")
            __props__.__dict__["client_cert_file"] = client_cert_file
            __props__.__dict__["client_key"] = client_key
            if client_key_file is not None and not opts.urn:
                warnings.warn("""This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.""", DeprecationWarning)
                pulumi.log.warn("""client_key_file is deprecated: This parameter is now deprecated and will be removed in a later release, please use `client_key` instead.""")
            __props__.__dict__["client_key_file"] = client_key_file
            __props__.__dict__["client_key_passphrase"] = client_key_passphrase
            if sasl_mechanism is None:
                sasl_mechanism = (_utilities.get_env('KAFKA_SASL_MECHANISM') or 'plain')
            __props__.__dict__["sasl_mechanism"] = sasl_mechanism
            __props__.__dict__["sasl_password"] = sasl_password
            __props__.__dict__["sasl_username"] = sasl_username
            if skip_tls_verify is None:
                skip_tls_verify = (_utilities.get_env_bool('KAFKA_SKIP_VERIFY') or False)
            __props__.__dict__["skip_tls_verify"] = pulumi.Output.from_input(skip_tls_verify).apply(pulumi.runtime.to_json) if skip_tls_verify is not None else None
            __props__.__dict__["timeout"] = pulumi.Output.from_input(timeout).apply(pulumi.runtime.to_json) if timeout is not None else None
            if tls_enabled is None:
                tls_enabled = (_utilities.get_env_bool('KAFKA_ENABLE_TLS') or True)
            __props__.__dict__["tls_enabled"] = pulumi.Output.from_input(tls_enabled).apply(pulumi.runtime.to_json) if tls_enabled is not None else None
        super(Provider, __self__).__init__(
            'kafka',
            resource_name,
            __props__,
            opts)

