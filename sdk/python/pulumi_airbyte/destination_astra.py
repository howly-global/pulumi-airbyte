# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DestinationAstraArgs', 'DestinationAstra']

@pulumi.input_type
class DestinationAstraArgs:
    def __init__(__self__, *,
                 configuration: pulumi.Input['DestinationAstraConfigurationArgs'],
                 workspace_id: pulumi.Input[str],
                 definition_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DestinationAstra resource.
        :param pulumi.Input['DestinationAstraConfigurationArgs'] configuration: The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination
               configuration, as well as to provide type safety for the configuration passed to the destination. The configuration
               model is composed of four parts: * Processing configuration * Embedding configuration * Indexing configuration *
               Advanced configuration Processing, embedding and advanced configuration are provided by this base class, while the
               indexing configuration is provided by the destination connector in the sub class.
        :param pulumi.Input[str] definition_id: The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
               replacement if changed.
        :param pulumi.Input[str] name: Name of the destination e.g. dev-mysql-instance.
        """
        pulumi.set(__self__, "configuration", configuration)
        pulumi.set(__self__, "workspace_id", workspace_id)
        if definition_id is not None:
            pulumi.set(__self__, "definition_id", definition_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Input['DestinationAstraConfigurationArgs']:
        """
        The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination
        configuration, as well as to provide type safety for the configuration passed to the destination. The configuration
        model is composed of four parts: * Processing configuration * Embedding configuration * Indexing configuration *
        Advanced configuration Processing, embedding and advanced configuration are provided by this base class, while the
        indexing configuration is provided by the destination connector in the sub class.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: pulumi.Input['DestinationAstraConfigurationArgs']):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workspace_id", value)

    @property
    @pulumi.getter(name="definitionId")
    def definition_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
        replacement if changed.
        """
        return pulumi.get(self, "definition_id")

    @definition_id.setter
    def definition_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "definition_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the destination e.g. dev-mysql-instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DestinationAstraState:
    def __init__(__self__, *,
                 configuration: Optional[pulumi.Input['DestinationAstraConfigurationArgs']] = None,
                 definition_id: Optional[pulumi.Input[str]] = None,
                 destination_id: Optional[pulumi.Input[str]] = None,
                 destination_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DestinationAstra resources.
        :param pulumi.Input['DestinationAstraConfigurationArgs'] configuration: The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination
               configuration, as well as to provide type safety for the configuration passed to the destination. The configuration
               model is composed of four parts: * Processing configuration * Embedding configuration * Indexing configuration *
               Advanced configuration Processing, embedding and advanced configuration are provided by this base class, while the
               indexing configuration is provided by the destination connector in the sub class.
        :param pulumi.Input[str] definition_id: The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
               replacement if changed.
        :param pulumi.Input[str] name: Name of the destination e.g. dev-mysql-instance.
        """
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if definition_id is not None:
            pulumi.set(__self__, "definition_id", definition_id)
        if destination_id is not None:
            pulumi.set(__self__, "destination_id", destination_id)
        if destination_type is not None:
            pulumi.set(__self__, "destination_type", destination_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if workspace_id is not None:
            pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['DestinationAstraConfigurationArgs']]:
        """
        The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination
        configuration, as well as to provide type safety for the configuration passed to the destination. The configuration
        model is composed of four parts: * Processing configuration * Embedding configuration * Indexing configuration *
        Advanced configuration Processing, embedding and advanced configuration are provided by this base class, while the
        indexing configuration is provided by the destination connector in the sub class.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['DestinationAstraConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="definitionId")
    def definition_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
        replacement if changed.
        """
        return pulumi.get(self, "definition_id")

    @definition_id.setter
    def definition_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "definition_id", value)

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_id")

    @destination_id.setter
    def destination_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_id", value)

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_type")

    @destination_type.setter
    def destination_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the destination e.g. dev-mysql-instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_id", value)


class DestinationAstra(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['DestinationAstraConfigurationArgs']]] = None,
                 definition_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DestinationAstra resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DestinationAstraConfigurationArgs']] configuration: The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination
               configuration, as well as to provide type safety for the configuration passed to the destination. The configuration
               model is composed of four parts: * Processing configuration * Embedding configuration * Indexing configuration *
               Advanced configuration Processing, embedding and advanced configuration are provided by this base class, while the
               indexing configuration is provided by the destination connector in the sub class.
        :param pulumi.Input[str] definition_id: The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
               replacement if changed.
        :param pulumi.Input[str] name: Name of the destination e.g. dev-mysql-instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DestinationAstraArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DestinationAstra resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DestinationAstraArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DestinationAstraArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['DestinationAstraConfigurationArgs']]] = None,
                 definition_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DestinationAstraArgs.__new__(DestinationAstraArgs)

            if configuration is None and not opts.urn:
                raise TypeError("Missing required property 'configuration'")
            __props__.__dict__["configuration"] = configuration
            __props__.__dict__["definition_id"] = definition_id
            __props__.__dict__["name"] = name
            if workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_id'")
            __props__.__dict__["workspace_id"] = workspace_id
            __props__.__dict__["destination_id"] = None
            __props__.__dict__["destination_type"] = None
        super(DestinationAstra, __self__).__init__(
            'airbyte:index/destinationAstra:DestinationAstra',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configuration: Optional[pulumi.Input[pulumi.InputType['DestinationAstraConfigurationArgs']]] = None,
            definition_id: Optional[pulumi.Input[str]] = None,
            destination_id: Optional[pulumi.Input[str]] = None,
            destination_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'DestinationAstra':
        """
        Get an existing DestinationAstra resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DestinationAstraConfigurationArgs']] configuration: The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination
               configuration, as well as to provide type safety for the configuration passed to the destination. The configuration
               model is composed of four parts: * Processing configuration * Embedding configuration * Indexing configuration *
               Advanced configuration Processing, embedding and advanced configuration are provided by this base class, while the
               indexing configuration is provided by the destination connector in the sub class.
        :param pulumi.Input[str] definition_id: The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
               replacement if changed.
        :param pulumi.Input[str] name: Name of the destination e.g. dev-mysql-instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DestinationAstraState.__new__(_DestinationAstraState)

        __props__.__dict__["configuration"] = configuration
        __props__.__dict__["definition_id"] = definition_id
        __props__.__dict__["destination_id"] = destination_id
        __props__.__dict__["destination_type"] = destination_type
        __props__.__dict__["name"] = name
        __props__.__dict__["workspace_id"] = workspace_id
        return DestinationAstra(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output['outputs.DestinationAstraConfiguration']:
        """
        The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination
        configuration, as well as to provide type safety for the configuration passed to the destination. The configuration
        model is composed of four parts: * Processing configuration * Embedding configuration * Indexing configuration *
        Advanced configuration Processing, embedding and advanced configuration are provided by this base class, while the
        indexing configuration is provided by the destination connector in the sub class.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="definitionId")
    def definition_id(self) -> pulumi.Output[Optional[str]]:
        """
        The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
        replacement if changed.
        """
        return pulumi.get(self, "definition_id")

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "destination_id")

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the destination e.g. dev-mysql-instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workspace_id")
