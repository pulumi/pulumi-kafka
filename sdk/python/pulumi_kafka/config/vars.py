# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('kafka')


class _ExportableConfig(types.ModuleType):
    @property
    def bootstrap_servers(self) -> Optional[str]:
        """
        A list of kafka brokers
        """
        return __config__.get('bootstrapServers')

    @property
    def ca_cert(self) -> Optional[str]:
        """
        CA certificate file to validate the server's certificate.
        """
        return __config__.get('caCert')

    @property
    def ca_cert_file(self) -> Optional[str]:
        """
        Path to a CA certificate file to validate the server's certificate.
        """
        return __config__.get('caCertFile')

    @property
    def client_cert(self) -> Optional[str]:
        """
        The client certificate.
        """
        return __config__.get('clientCert')

    @property
    def client_cert_file(self) -> Optional[str]:
        """
        Path to a file containing the client certificate.
        """
        return __config__.get('clientCertFile')

    @property
    def client_key(self) -> Optional[str]:
        """
        The private key that the certificate was issued for.
        """
        return __config__.get('clientKey')

    @property
    def client_key_file(self) -> Optional[str]:
        """
        Path to a file containing the private key that the certificate was issued for.
        """
        return __config__.get('clientKeyFile')

    @property
    def client_key_passphrase(self) -> Optional[str]:
        """
        The passphrase for the private key that the certificate was issued for.
        """
        return __config__.get('clientKeyPassphrase')

    @property
    def sasl_mechanism(self) -> str:
        """
        SASL mechanism, can be plain, scram-sha512, scram-sha256
        """
        return __config__.get('saslMechanism') or (_utilities.get_env('KAFKA_SASL_MECHANISM') or 'plain')

    @property
    def sasl_password(self) -> Optional[str]:
        """
        Password for SASL authentication.
        """
        return __config__.get('saslPassword')

    @property
    def sasl_username(self) -> Optional[str]:
        """
        Username for SASL authentication.
        """
        return __config__.get('saslUsername')

    @property
    def skip_tls_verify(self) -> bool:
        """
        Set this to true only if the target Kafka server is an insecure development instance.
        """
        return __config__.get_bool('skipTlsVerify') or (_utilities.get_env_bool('KAFKA_SKIP_VERIFY') or False)

    @property
    def timeout(self) -> Optional[int]:
        """
        Timeout in seconds
        """
        return __config__.get_int('timeout')

    @property
    def tls_enabled(self) -> bool:
        """
        Enable communication with the Kafka Cluster over TLS.
        """
        return __config__.get_bool('tlsEnabled') or (_utilities.get_env_bool('KAFKA_ENABLE_TLS') or True)

