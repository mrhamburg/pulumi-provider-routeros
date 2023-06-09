# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AaaArgs', 'Aaa']

@pulumi.input_type
class AaaArgs:
    def __init__(__self__, *,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 called_format: Optional[pulumi.Input[str]] = None,
                 interim_update: Optional[pulumi.Input[str]] = None,
                 mac_caching: Optional[pulumi.Input[str]] = None,
                 mac_format: Optional[pulumi.Input[str]] = None,
                 mac_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Aaa resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] called_format: Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
        :param pulumi.Input[str] interim_update: When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
        :param pulumi.Input[str] mac_caching: If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
        :param pulumi.Input[str] mac_format: Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
        :param pulumi.Input[str] mac_mode: By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
        """
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if called_format is not None:
            pulumi.set(__self__, "called_format", called_format)
        if interim_update is not None:
            pulumi.set(__self__, "interim_update", interim_update)
        if mac_caching is not None:
            pulumi.set(__self__, "mac_caching", mac_caching)
        if mac_format is not None:
            pulumi.set(__self__, "mac_format", mac_format)
        if mac_mode is not None:
            pulumi.set(__self__, "mac_mode", mac_mode)

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
    @pulumi.getter(name="calledFormat")
    def called_format(self) -> Optional[pulumi.Input[str]]:
        """
        Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
        """
        return pulumi.get(self, "called_format")

    @called_format.setter
    def called_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "called_format", value)

    @property
    @pulumi.getter(name="interimUpdate")
    def interim_update(self) -> Optional[pulumi.Input[str]]:
        """
        When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
        """
        return pulumi.get(self, "interim_update")

    @interim_update.setter
    def interim_update(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interim_update", value)

    @property
    @pulumi.getter(name="macCaching")
    def mac_caching(self) -> Optional[pulumi.Input[str]]:
        """
        If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
        """
        return pulumi.get(self, "mac_caching")

    @mac_caching.setter
    def mac_caching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_caching", value)

    @property
    @pulumi.getter(name="macFormat")
    def mac_format(self) -> Optional[pulumi.Input[str]]:
        """
        Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
        """
        return pulumi.get(self, "mac_format")

    @mac_format.setter
    def mac_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_format", value)

    @property
    @pulumi.getter(name="macMode")
    def mac_mode(self) -> Optional[pulumi.Input[str]]:
        """
        By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
        """
        return pulumi.get(self, "mac_mode")

    @mac_mode.setter
    def mac_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_mode", value)


@pulumi.input_type
class _AaaState:
    def __init__(__self__, *,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 called_format: Optional[pulumi.Input[str]] = None,
                 interim_update: Optional[pulumi.Input[str]] = None,
                 mac_caching: Optional[pulumi.Input[str]] = None,
                 mac_format: Optional[pulumi.Input[str]] = None,
                 mac_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Aaa resources.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] called_format: Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
        :param pulumi.Input[str] interim_update: When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
        :param pulumi.Input[str] mac_caching: If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
        :param pulumi.Input[str] mac_format: Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
        :param pulumi.Input[str] mac_mode: By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
        """
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if called_format is not None:
            pulumi.set(__self__, "called_format", called_format)
        if interim_update is not None:
            pulumi.set(__self__, "interim_update", interim_update)
        if mac_caching is not None:
            pulumi.set(__self__, "mac_caching", mac_caching)
        if mac_format is not None:
            pulumi.set(__self__, "mac_format", mac_format)
        if mac_mode is not None:
            pulumi.set(__self__, "mac_mode", mac_mode)

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
    @pulumi.getter(name="calledFormat")
    def called_format(self) -> Optional[pulumi.Input[str]]:
        """
        Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
        """
        return pulumi.get(self, "called_format")

    @called_format.setter
    def called_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "called_format", value)

    @property
    @pulumi.getter(name="interimUpdate")
    def interim_update(self) -> Optional[pulumi.Input[str]]:
        """
        When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
        """
        return pulumi.get(self, "interim_update")

    @interim_update.setter
    def interim_update(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interim_update", value)

    @property
    @pulumi.getter(name="macCaching")
    def mac_caching(self) -> Optional[pulumi.Input[str]]:
        """
        If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
        """
        return pulumi.get(self, "mac_caching")

    @mac_caching.setter
    def mac_caching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_caching", value)

    @property
    @pulumi.getter(name="macFormat")
    def mac_format(self) -> Optional[pulumi.Input[str]]:
        """
        Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
        """
        return pulumi.get(self, "mac_format")

    @mac_format.setter
    def mac_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_format", value)

    @property
    @pulumi.getter(name="macMode")
    def mac_mode(self) -> Optional[pulumi.Input[str]]:
        """
        By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
        """
        return pulumi.get(self, "mac_mode")

    @mac_mode.setter
    def mac_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_mode", value)


class Aaa(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 called_format: Optional[pulumi.Input[str]] = None,
                 interim_update: Optional[pulumi.Input[str]] = None,
                 mac_caching: Optional[pulumi.Input[str]] = None,
                 mac_format: Optional[pulumi.Input[str]] = None,
                 mac_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        test3a = routeros.caps_man.Aaa("test3a",
            called_format="ssid",
            mac_mode="as-username-and-password")
        ```

        ## Import

        ```sh
         $ pulumi import routeros:CapsMan/aaa:Aaa test_3a .
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] called_format: Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
        :param pulumi.Input[str] interim_update: When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
        :param pulumi.Input[str] mac_caching: If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
        :param pulumi.Input[str] mac_format: Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
        :param pulumi.Input[str] mac_mode: By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AaaArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        test3a = routeros.caps_man.Aaa("test3a",
            called_format="ssid",
            mac_mode="as-username-and-password")
        ```

        ## Import

        ```sh
         $ pulumi import routeros:CapsMan/aaa:Aaa test_3a .
        ```

        :param str resource_name: The name of the resource.
        :param AaaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AaaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 called_format: Optional[pulumi.Input[str]] = None,
                 interim_update: Optional[pulumi.Input[str]] = None,
                 mac_caching: Optional[pulumi.Input[str]] = None,
                 mac_format: Optional[pulumi.Input[str]] = None,
                 mac_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AaaArgs.__new__(AaaArgs)

            __props__.__dict__["___id_"] = ___id_
            __props__.__dict__["___path_"] = ___path_
            __props__.__dict__["called_format"] = called_format
            __props__.__dict__["interim_update"] = interim_update
            __props__.__dict__["mac_caching"] = mac_caching
            __props__.__dict__["mac_format"] = mac_format
            __props__.__dict__["mac_mode"] = mac_mode
        super(Aaa, __self__).__init__(
            'routeros:CapsMan/aaa:Aaa',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ___id_: Optional[pulumi.Input[int]] = None,
            ___path_: Optional[pulumi.Input[str]] = None,
            called_format: Optional[pulumi.Input[str]] = None,
            interim_update: Optional[pulumi.Input[str]] = None,
            mac_caching: Optional[pulumi.Input[str]] = None,
            mac_format: Optional[pulumi.Input[str]] = None,
            mac_mode: Optional[pulumi.Input[str]] = None) -> 'Aaa':
        """
        Get an existing Aaa resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] called_format: Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
        :param pulumi.Input[str] interim_update: When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
        :param pulumi.Input[str] mac_caching: If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
        :param pulumi.Input[str] mac_format: Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
        :param pulumi.Input[str] mac_mode: By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AaaState.__new__(_AaaState)

        __props__.__dict__["___id_"] = ___id_
        __props__.__dict__["___path_"] = ___path_
        __props__.__dict__["called_format"] = called_format
        __props__.__dict__["interim_update"] = interim_update
        __props__.__dict__["mac_caching"] = mac_caching
        __props__.__dict__["mac_format"] = mac_format
        __props__.__dict__["mac_mode"] = mac_mode
        return Aaa(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="calledFormat")
    def called_format(self) -> pulumi.Output[Optional[str]]:
        """
        Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
        """
        return pulumi.get(self, "called_format")

    @property
    @pulumi.getter(name="interimUpdate")
    def interim_update(self) -> pulumi.Output[str]:
        """
        When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
        """
        return pulumi.get(self, "interim_update")

    @property
    @pulumi.getter(name="macCaching")
    def mac_caching(self) -> pulumi.Output[str]:
        """
        If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
        """
        return pulumi.get(self, "mac_caching")

    @property
    @pulumi.getter(name="macFormat")
    def mac_format(self) -> pulumi.Output[Optional[str]]:
        """
        Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
        """
        return pulumi.get(self, "mac_format")

    @property
    @pulumi.getter(name="macMode")
    def mac_mode(self) -> pulumi.Output[Optional[str]]:
        """
        By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
        """
        return pulumi.get(self, "mac_mode")

