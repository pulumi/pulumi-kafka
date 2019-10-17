# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Acl(pulumi.CustomResource):
    acl_host: pulumi.Output[str]
    """
    Host from which principal listed in `acl_principal`
    will have access.
    """
    acl_operation: pulumi.Output[str]
    """
    Operation that is being allowed or denied. Valid
    values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
    `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
    """
    acl_permission_type: pulumi.Output[str]
    """
    Type of permission. Valid values are `Unknown`,
    `Any`, `Allow`, `Deny`.
    """
    acl_principal: pulumi.Output[str]
    """
    Principal that is being allowed or denied.
    """
    resource_name: pulumi.Output[str]
    """
    The name of the resource.
    """
    resource_pattern_type_filter: pulumi.Output[str]
    """
    The pattern filter. Valid values
    are `Prefixed`, `Any`, `Match`, `Literal`.
    """
    resource_type: pulumi.Output[str]
    """
    The type of resource. Valid values are `Unknown`,
    `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
    """
    def __init__(__self__, resource_name, opts=None, acl_host=None, acl_operation=None, acl_permission_type=None, acl_principal=None, resource_name=None, resource_pattern_type_filter=None, resource_type=None, __props__=None, __name__=None, __opts__=None):
        """
        A resource for managing Kafka ACLs.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl_host: Host from which principal listed in `acl_principal`
               will have access.
        :param pulumi.Input[str] acl_operation: Operation that is being allowed or denied. Valid
               values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
               `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
        :param pulumi.Input[str] acl_permission_type: Type of permission. Valid values are `Unknown`,
               `Any`, `Allow`, `Deny`.
        :param pulumi.Input[str] acl_principal: Principal that is being allowed or denied.
        :param pulumi.Input[str] resource_name: The name of the resource.
        :param pulumi.Input[str] resource_pattern_type_filter: The pattern filter. Valid values
               are `Prefixed`, `Any`, `Match`, `Literal`.
        :param pulumi.Input[str] resource_type: The type of resource. Valid values are `Unknown`,
               `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-kafka/blob/master/website/docs/r/acl.html.markdown.
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

            if acl_host is None:
                raise TypeError("Missing required property 'acl_host'")
            __props__['acl_host'] = acl_host
            if acl_operation is None:
                raise TypeError("Missing required property 'acl_operation'")
            __props__['acl_operation'] = acl_operation
            if acl_permission_type is None:
                raise TypeError("Missing required property 'acl_permission_type'")
            __props__['acl_permission_type'] = acl_permission_type
            if acl_principal is None:
                raise TypeError("Missing required property 'acl_principal'")
            __props__['acl_principal'] = acl_principal
            if resource_name is None:
                raise TypeError("Missing required property 'resource_name'")
            __props__['resource_name'] = resource_name
            __props__['resource_pattern_type_filter'] = resource_pattern_type_filter
            if resource_type is None:
                raise TypeError("Missing required property 'resource_type'")
            __props__['resource_type'] = resource_type
        super(Acl, __self__).__init__(
            'kafka:index/acl:Acl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, acl_host=None, acl_operation=None, acl_permission_type=None, acl_principal=None, resource_name=None, resource_pattern_type_filter=None, resource_type=None):
        """
        Get an existing Acl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl_host: Host from which principal listed in `acl_principal`
               will have access.
        :param pulumi.Input[str] acl_operation: Operation that is being allowed or denied. Valid
               values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
               `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
        :param pulumi.Input[str] acl_permission_type: Type of permission. Valid values are `Unknown`,
               `Any`, `Allow`, `Deny`.
        :param pulumi.Input[str] acl_principal: Principal that is being allowed or denied.
        :param pulumi.Input[str] resource_name: The name of the resource.
        :param pulumi.Input[str] resource_pattern_type_filter: The pattern filter. Valid values
               are `Prefixed`, `Any`, `Match`, `Literal`.
        :param pulumi.Input[str] resource_type: The type of resource. Valid values are `Unknown`,
               `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-kafka/blob/master/website/docs/r/acl.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["acl_host"] = acl_host
        __props__["acl_operation"] = acl_operation
        __props__["acl_permission_type"] = acl_permission_type
        __props__["acl_principal"] = acl_principal
        __props__["resource_name"] = resource_name
        __props__["resource_pattern_type_filter"] = resource_pattern_type_filter
        __props__["resource_type"] = resource_type
        return Acl(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

