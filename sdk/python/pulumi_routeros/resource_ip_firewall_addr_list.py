# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ResourceIPFirewallAddrListArgs', 'ResourceIPFirewallAddrList']

@pulumi.input_type
class ResourceIPFirewallAddrListArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 list: pulumi.Input[str],
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 timeout: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ResourceIPFirewallAddrList resource.
        :param pulumi.Input[str] address: A single IP address or range of IPs to add to address list or DNS name. You can input for example,
               '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving.
        :param pulumi.Input[str] list: Name for the address list of the added IP address.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] timeout: Time after address will be removed from address list. If timeout is not specified, the address will be stored into the
               address list permanently. > Please plan your work logic based on the fact that after the timeout > the resource has been
               destroyed outside of Terraform.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "list", list)
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        A single IP address or range of IPs to add to address list or DNS name. You can input for example,
        '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def list(self) -> pulumi.Input[str]:
        """
        Name for the address list of the added IP address.
        """
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
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        Time after address will be removed from address list. If timeout is not specified, the address will be stored into the
        address list permanently. > Please plan your work logic based on the fact that after the timeout > the resource has been
        destroyed outside of Terraform.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _ResourceIPFirewallAddrListState:
    def __init__(__self__, *,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 dynamic: Optional[pulumi.Input[bool]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ResourceIPFirewallAddrList resources.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] address: A single IP address or range of IPs to add to address list or DNS name. You can input for example,
               '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving.
        :param pulumi.Input[str] creation_time: Rule creation time
        :param pulumi.Input[bool] dynamic: Configuration item created by software, not by management interface. It is not exported, and cannot be directly
               modified.
        :param pulumi.Input[str] list: Name for the address list of the added IP address.
        :param pulumi.Input[str] timeout: Time after address will be removed from address list. If timeout is not specified, the address will be stored into the
               address list permanently. > Please plan your work logic based on the fact that after the timeout > the resource has been
               destroyed outside of Terraform.
        """
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if dynamic is not None:
            pulumi.set(__self__, "dynamic", dynamic)
        if list is not None:
            pulumi.set(__self__, "list", list)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

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
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        A single IP address or range of IPs to add to address list or DNS name. You can input for example,
        '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        Rule creation time
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

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
        """
        Configuration item created by software, not by management interface. It is not exported, and cannot be directly
        modified.
        """
        return pulumi.get(self, "dynamic")

    @dynamic.setter
    def dynamic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dynamic", value)

    @property
    @pulumi.getter
    def list(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the address list of the added IP address.
        """
        return pulumi.get(self, "list")

    @list.setter
    def list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "list", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        Time after address will be removed from address list. If timeout is not specified, the address will be stored into the
        address list permanently. > Please plan your work logic based on the fact that after the timeout > the resource has been
        destroyed outside of Terraform.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)


class ResourceIPFirewallAddrList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        example_list = routeros.ResourceIPFirewallAddrList("exampleList",
            address="1.1.1.1",
            list="Example List")
        ```

        ## Import

        #The ID can be found via API or the terminal #The command for the terminal is -> :put [/ip/firewall/address-list get [print show-ids]]

        ```sh
         $ pulumi import routeros:index/resourceIPFirewallAddrList:ResourceIPFirewallAddrList example_list "*0"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] address: A single IP address or range of IPs to add to address list or DNS name. You can input for example,
               '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving.
        :param pulumi.Input[str] list: Name for the address list of the added IP address.
        :param pulumi.Input[str] timeout: Time after address will be removed from address list. If timeout is not specified, the address will be stored into the
               address list permanently. > Please plan your work logic based on the fact that after the timeout > the resource has been
               destroyed outside of Terraform.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceIPFirewallAddrListArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        example_list = routeros.ResourceIPFirewallAddrList("exampleList",
            address="1.1.1.1",
            list="Example List")
        ```

        ## Import

        #The ID can be found via API or the terminal #The command for the terminal is -> :put [/ip/firewall/address-list get [print show-ids]]

        ```sh
         $ pulumi import routeros:index/resourceIPFirewallAddrList:ResourceIPFirewallAddrList example_list "*0"
        ```

        :param str resource_name: The name of the resource.
        :param ResourceIPFirewallAddrListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceIPFirewallAddrListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceIPFirewallAddrListArgs.__new__(ResourceIPFirewallAddrListArgs)

            __props__.__dict__["___id_"] = ___id_
            __props__.__dict__["___path_"] = ___path_
            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["comment"] = comment
            __props__.__dict__["disabled"] = disabled
            if list is None and not opts.urn:
                raise TypeError("Missing required property 'list'")
            __props__.__dict__["list"] = list
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["dynamic"] = None
        super(ResourceIPFirewallAddrList, __self__).__init__(
            'routeros:index/resourceIPFirewallAddrList:ResourceIPFirewallAddrList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ___id_: Optional[pulumi.Input[int]] = None,
            ___path_: Optional[pulumi.Input[str]] = None,
            address: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            dynamic: Optional[pulumi.Input[bool]] = None,
            list: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[str]] = None) -> 'ResourceIPFirewallAddrList':
        """
        Get an existing ResourceIPFirewallAddrList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] address: A single IP address or range of IPs to add to address list or DNS name. You can input for example,
               '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving.
        :param pulumi.Input[str] creation_time: Rule creation time
        :param pulumi.Input[bool] dynamic: Configuration item created by software, not by management interface. It is not exported, and cannot be directly
               modified.
        :param pulumi.Input[str] list: Name for the address list of the added IP address.
        :param pulumi.Input[str] timeout: Time after address will be removed from address list. If timeout is not specified, the address will be stored into the
               address list permanently. > Please plan your work logic based on the fact that after the timeout > the resource has been
               destroyed outside of Terraform.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceIPFirewallAddrListState.__new__(_ResourceIPFirewallAddrListState)

        __props__.__dict__["___id_"] = ___id_
        __props__.__dict__["___path_"] = ___path_
        __props__.__dict__["address"] = address
        __props__.__dict__["comment"] = comment
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["dynamic"] = dynamic
        __props__.__dict__["list"] = list
        __props__.__dict__["timeout"] = timeout
        return ResourceIPFirewallAddrList(resource_name, opts=opts, __props__=__props__)

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
    def address(self) -> pulumi.Output[str]:
        """
        A single IP address or range of IPs to add to address list or DNS name. You can input for example,
        '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        Rule creation time
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def dynamic(self) -> pulumi.Output[bool]:
        """
        Configuration item created by software, not by management interface. It is not exported, and cannot be directly
        modified.
        """
        return pulumi.get(self, "dynamic")

    @property
    @pulumi.getter
    def list(self) -> pulumi.Output[str]:
        """
        Name for the address list of the added IP address.
        """
        return pulumi.get(self, "list")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[str]]:
        """
        Time after address will be removed from address list. If timeout is not specified, the address will be stored into the
        address list permanently. > Please plan your work logic based on the fact that after the timeout > the resource has been
        destroyed outside of Terraform.
        """
        return pulumi.get(self, "timeout")

