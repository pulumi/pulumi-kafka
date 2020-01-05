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
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/Mongey/terraform-provider-kafka/blob/master/website/docs/index.html.markdown.
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
            __props__['ca_cert'] = ca_cert
            __props__['ca_cert_file'] = ca_cert_file
            __props__['client_cert'] = client_cert
            __props__['client_cert_file'] = client_cert_file
            __props__['client_key'] = client_key
            __props__['client_key_file'] = client_key_file
            __props__['sasl_mechanism'] = sasl_mechanism
            __props__['sasl_password'] = sasl_password
            __props__['sasl_username'] = sasl_username
            __props__['skip_tls_verify'] = pulumi.Output.from_input(skip_tls_verify).apply(json.dumps) if skip_tls_verify is not None else None
            __props__['timeout'] = pulumi.Output.from_input(timeout).apply(json.dumps) if timeout is not None else None
            __props__['tls_enabled'] = pulumi.Output.from_input(tls_enabled).apply(json.dumps) if tls_enabled is not None else None
        super(Provider, __self__).__init__(
            'kafka',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Provider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/Mongey/terraform-provider-kafka/blob/master/website/docs/index.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        return Provider(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

