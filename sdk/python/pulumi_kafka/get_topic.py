# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = [
    'GetTopicResult',
    'AwaitableGetTopicResult',
    'get_topic',
    'get_topic_output',
]

@pulumi.output_type
class GetTopicResult:
    """
    A collection of values returned by getTopic.
    """
    def __init__(__self__, config=None, id=None, name=None, partitions=None, replication_factor=None):
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partitions and not isinstance(partitions, int):
            raise TypeError("Expected argument 'partitions' to be a int")
        pulumi.set(__self__, "partitions", partitions)
        if replication_factor and not isinstance(replication_factor, int):
            raise TypeError("Expected argument 'replication_factor' to be a int")
        pulumi.set(__self__, "replication_factor", replication_factor)

    @_builtins.property
    @pulumi.getter
    def config(self) -> Mapping[str, _builtins.str]:
        return pulumi.get(self, "config")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def name(self) -> _builtins.str:
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def partitions(self) -> _builtins.int:
        return pulumi.get(self, "partitions")

    @_builtins.property
    @pulumi.getter(name="replicationFactor")
    def replication_factor(self) -> _builtins.int:
        return pulumi.get(self, "replication_factor")


class AwaitableGetTopicResult(GetTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicResult(
            config=self.config,
            id=self.id,
            name=self.name,
            partitions=self.partitions,
            replication_factor=self.replication_factor)


def get_topic(name: Optional[_builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kafka:index:getTopic', __args__, opts=opts, typ=GetTopicResult).value

    return AwaitableGetTopicResult(
        config=pulumi.get(__ret__, 'config'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        partitions=pulumi.get(__ret__, 'partitions'),
        replication_factor=pulumi.get(__ret__, 'replication_factor'))
def get_topic_output(name: Optional[pulumi.Input[_builtins.str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTopicResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('kafka:index:getTopic', __args__, opts=opts, typ=GetTopicResult)
    return __ret__.apply(lambda __response__: GetTopicResult(
        config=pulumi.get(__response__, 'config'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        partitions=pulumi.get(__response__, 'partitions'),
        replication_factor=pulumi.get(__response__, 'replication_factor')))
