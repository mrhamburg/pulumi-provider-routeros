# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ResourceInterfaceListMemberArgs', 'ResourceInterfaceListMember']

@pulumi.input_type
class ResourceInterfaceListMemberArgs:
    def __init__(__self__, *,
                 interface: pulumi.Input[str],
                 list: pulumi.Input[str],
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ResourceInterfaceListMember resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        """
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "list", list)
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def list(self) -> pulumi.Input[str]:
        return pulumi.get(self, "list")

    @list.setter
    def list(self, value: pulumi.Input[str]):
        pulumi.set(self, "list", value)

    @property
    @pulumi.getter
    def ___id_(self) -> Optional[pulumi.Input[int]]:
        """
        <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        """
        return pulumi.get(self, "___id_")

    @___id_.setter
    def ___id_(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "___id_", value)

    @property
    @pulumi.getter
    def ___path_(self) -> Optional[pulumi.Input[str]]:
        """
        <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        """
        return pulumi.get(self, "___path_")

    @___path_.setter
    def ___path_(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "___path_", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)


@pulumi.input_type
class _ResourceInterfaceListMemberState:
    def __init__(__self__, *,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 dynamic: Optional[pulumi.Input[bool]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 list: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ResourceInterfaceListMember resources.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        """
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if dynamic is not None:
            pulumi.set(__self__, "dynamic", dynamic)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if list is not None:
            pulumi.set(__self__, "list", list)

    @property
    @pulumi.getter
    def ___id_(self) -> Optional[pulumi.Input[int]]:
        """
        <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        """
        return pulumi.get(self, "___id_")

    @___id_.setter
    def ___id_(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "___id_", value)

    @property
    @pulumi.getter
    def ___path_(self) -> Optional[pulumi.Input[str]]:
        """
        <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        """
        return pulumi.get(self, "___path_")

    @___path_.setter
    def ___path_(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "___path_", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def dynamic(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "dynamic")

    @dynamic.setter
    def dynamic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dynamic", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def list(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "list")

    @list.setter
    def list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "list", value)


class ResourceInterfaceListMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        list_member = routeros.ResourceInterfaceListMember("listMember",
            interface="ether1",
            list="my-list")
        ```

        ## Import

        #The ID can be found via API or the terminal #The command for the terminal is -> :put [/interface/list/member get [print show-ids]]

        ```sh
         $ pulumi import routeros:index/resourceInterfaceListMember:ResourceInterfaceListMember list_member "*0"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceInterfaceListMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        list_member = routeros.ResourceInterfaceListMember("listMember",
            interface="ether1",
            list="my-list")
        ```

        ## Import

        #The ID can be found via API or the terminal #The command for the terminal is -> :put [/interface/list/member get [print show-ids]]

        ```sh
         $ pulumi import routeros:index/resourceInterfaceListMember:ResourceInterfaceListMember list_member "*0"
        ```

        :param str resource_name: The name of the resource.
        :param ResourceInterfaceListMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceInterfaceListMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceInterfaceListMemberArgs.__new__(ResourceInterfaceListMemberArgs)

            __props__.__dict__["___id_"] = ___id_
            __props__.__dict__["___path_"] = ___path_
            __props__.__dict__["disabled"] = disabled
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if list is None and not opts.urn:
                raise TypeError("Missing required property 'list'")
            __props__.__dict__["list"] = list
            __props__.__dict__["dynamic"] = None
        super(ResourceInterfaceListMember, __self__).__init__(
            'routeros:index/resourceInterfaceListMember:ResourceInterfaceListMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ___id_: Optional[pulumi.Input[int]] = None,
            ___path_: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            dynamic: Optional[pulumi.Input[bool]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            list: Optional[pulumi.Input[str]] = None) -> 'ResourceInterfaceListMember':
        """
        Get an existing ResourceInterfaceListMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceInterfaceListMemberState.__new__(_ResourceInterfaceListMemberState)

        __props__.__dict__["___id_"] = ___id_
        __props__.__dict__["___path_"] = ___path_
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["dynamic"] = dynamic
        __props__.__dict__["interface"] = interface
        __props__.__dict__["list"] = list
        return ResourceInterfaceListMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ___id_(self) -> pulumi.Output[Optional[int]]:
        """
        <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        """
        return pulumi.get(self, "___id_")

    @property
    @pulumi.getter
    def ___path_(self) -> pulumi.Output[Optional[str]]:
        """
        <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        """
        return pulumi.get(self, "___path_")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def dynamic(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "dynamic")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def list(self) -> pulumi.Output[str]:
        return pulumi.get(self, "list")

