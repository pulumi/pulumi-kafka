# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['QuotaArgs', 'Quota']

@pulumi.input_type
class QuotaArgs:
    def __init__(__self__, *,
                 entity_name: pulumi.Input[builtins.str],
                 entity_type: pulumi.Input[builtins.str],
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Quota resource.
        :param pulumi.Input[builtins.str] entity_name: The name of the entity to target.
        :param pulumi.Input[builtins.str] entity_type: The type of entity. Valid values are `client-id`, `user`, `ip`.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] config: A map of string k/v attributes.
        """
        pulumi.set(__self__, "entity_name", entity_name)
        pulumi.set(__self__, "entity_type", entity_type)
        if config is not None:
            pulumi.set(__self__, "config", config)

    @property
    @pulumi.getter(name="entityName")
    def entity_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the entity to target.
        """
        return pulumi.get(self, "entity_name")

    @entity_name.setter
    def entity_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "entity_name", value)

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> pulumi.Input[builtins.str]:
        """
        The type of entity. Valid values are `client-id`, `user`, `ip`.
        """
        return pulumi.get(self, "entity_type")

    @entity_type.setter
    def entity_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "entity_type", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map of string k/v attributes.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "config", value)


@pulumi.input_type
class _QuotaState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 entity_name: Optional[pulumi.Input[builtins.str]] = None,
                 entity_type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Quota resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] config: A map of string k/v attributes.
        :param pulumi.Input[builtins.str] entity_name: The name of the entity to target.
        :param pulumi.Input[builtins.str] entity_type: The type of entity. Valid values are `client-id`, `user`, `ip`.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if entity_name is not None:
            pulumi.set(__self__, "entity_name", entity_name)
        if entity_type is not None:
            pulumi.set(__self__, "entity_type", entity_type)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map of string k/v attributes.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="entityName")
    def entity_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the entity to target.
        """
        return pulumi.get(self, "entity_name")

    @entity_name.setter
    def entity_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "entity_name", value)

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of entity. Valid values are `client-id`, `user`, `ip`.
        """
        return pulumi.get(self, "entity_type")

    @entity_type.setter
    def entity_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "entity_type", value)


class Quota(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 entity_name: Optional[pulumi.Input[builtins.str]] = None,
                 entity_type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A resource for managing Kafka quotas.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_kafka as kafka

        quota = kafka.Quota("quota",
            entity_name="app_consumer",
            entity_type="client-id",
            config={
                "consumer_byte_rate": "5000000",
                "producer_byte_rate": "2500000",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] config: A map of string k/v attributes.
        :param pulumi.Input[builtins.str] entity_name: The name of the entity to target.
        :param pulumi.Input[builtins.str] entity_type: The type of entity. Valid values are `client-id`, `user`, `ip`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QuotaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A resource for managing Kafka quotas.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_kafka as kafka

        quota = kafka.Quota("quota",
            entity_name="app_consumer",
            entity_type="client-id",
            config={
                "consumer_byte_rate": "5000000",
                "producer_byte_rate": "2500000",
            })
        ```

        :param str resource_name: The name of the resource.
        :param QuotaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QuotaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 entity_name: Optional[pulumi.Input[builtins.str]] = None,
                 entity_type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QuotaArgs.__new__(QuotaArgs)

            __props__.__dict__["config"] = config
            if entity_name is None and not opts.urn:
                raise TypeError("Missing required property 'entity_name'")
            __props__.__dict__["entity_name"] = entity_name
            if entity_type is None and not opts.urn:
                raise TypeError("Missing required property 'entity_type'")
            __props__.__dict__["entity_type"] = entity_type
        super(Quota, __self__).__init__(
            'kafka:index/quota:Quota',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            entity_name: Optional[pulumi.Input[builtins.str]] = None,
            entity_type: Optional[pulumi.Input[builtins.str]] = None) -> 'Quota':
        """
        Get an existing Quota resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] config: A map of string k/v attributes.
        :param pulumi.Input[builtins.str] entity_name: The name of the entity to target.
        :param pulumi.Input[builtins.str] entity_type: The type of entity. Valid values are `client-id`, `user`, `ip`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QuotaState.__new__(_QuotaState)

        __props__.__dict__["config"] = config
        __props__.__dict__["entity_name"] = entity_name
        __props__.__dict__["entity_type"] = entity_type
        return Quota(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        A map of string k/v attributes.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="entityName")
    def entity_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the entity to target.
        """
        return pulumi.get(self, "entity_name")

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> pulumi.Output[builtins.str]:
        """
        The type of entity. Valid values are `client-id`, `user`, `ip`.
        """
        return pulumi.get(self, "entity_type")

