# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSourceLeverHiringResult',
    'AwaitableGetSourceLeverHiringResult',
    'get_source_lever_hiring',
    'get_source_lever_hiring_output',
]

@pulumi.output_type
class GetSourceLeverHiringResult:
    """
    A collection of values returned by getSourceLeverHiring.
    """
    def __init__(__self__, configuration=None, id=None, name=None, source_id=None, source_type=None, workspace_id=None):
        if configuration and not isinstance(configuration, str):
            raise TypeError("Expected argument 'configuration' to be a str")
        pulumi.set(__self__, "configuration", configuration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if source_id and not isinstance(source_id, str):
            raise TypeError("Expected argument 'source_id' to be a str")
        pulumi.set(__self__, "source_id", source_id)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def configuration(self) -> str:
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> str:
        return pulumi.get(self, "source_id")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        return pulumi.get(self, "workspace_id")


class AwaitableGetSourceLeverHiringResult(GetSourceLeverHiringResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSourceLeverHiringResult(
            configuration=self.configuration,
            id=self.id,
            name=self.name,
            source_id=self.source_id,
            source_type=self.source_type,
            workspace_id=self.workspace_id)


def get_source_lever_hiring(source_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSourceLeverHiringResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['sourceId'] = source_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('airbyte:index/getSourceLeverHiring:getSourceLeverHiring', __args__, opts=opts, typ=GetSourceLeverHiringResult).value

    return AwaitableGetSourceLeverHiringResult(
        configuration=pulumi.get(__ret__, 'configuration'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        source_id=pulumi.get(__ret__, 'source_id'),
        source_type=pulumi.get(__ret__, 'source_type'),
        workspace_id=pulumi.get(__ret__, 'workspace_id'))


@_utilities.lift_output_func(get_source_lever_hiring)
def get_source_lever_hiring_output(source_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSourceLeverHiringResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
