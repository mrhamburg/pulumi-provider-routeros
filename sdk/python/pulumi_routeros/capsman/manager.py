# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ManagerArgs', 'Manager']

@pulumi.input_type
class ManagerArgs:
    def __init__(__self__, *,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 require_peer_certificate: Optional[pulumi.Input[bool]] = None,
                 upgrade_policy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Manager resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ca_certificate: Device CA certificate.
        :param pulumi.Input[str] certificate: Device certificate.
        :param pulumi.Input[bool] enabled: Disable or enable CAPsMAN functionality.
        :param pulumi.Input[str] package_path: Folder location for the RouterOS packages. For example, use '/upgrade' to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.
        :param pulumi.Input[bool] require_peer_certificate: Require all connecting CAPs to have a valid certificate.
        :param pulumi.Input[str] upgrade_policy: Upgrade policy options.
        """
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if ca_certificate is not None:
            pulumi.set(__self__, "ca_certificate", ca_certificate)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if package_path is not None:
            pulumi.set(__self__, "package_path", package_path)
        if require_peer_certificate is not None:
            pulumi.set(__self__, "require_peer_certificate", require_peer_certificate)
        if upgrade_policy is not None:
            pulumi.set(__self__, "upgrade_policy", upgrade_policy)

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
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Device CA certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Device certificate.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable or enable CAPsMAN functionality.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> Optional[pulumi.Input[str]]:
        """
        Folder location for the RouterOS packages. For example, use '/upgrade' to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.
        """
        return pulumi.get(self, "package_path")

    @package_path.setter
    def package_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_path", value)

    @property
    @pulumi.getter(name="requirePeerCertificate")
    def require_peer_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Require all connecting CAPs to have a valid certificate.
        """
        return pulumi.get(self, "require_peer_certificate")

    @require_peer_certificate.setter
    def require_peer_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_peer_certificate", value)

    @property
    @pulumi.getter(name="upgradePolicy")
    def upgrade_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Upgrade policy options.
        """
        return pulumi.get(self, "upgrade_policy")

    @upgrade_policy.setter
    def upgrade_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upgrade_policy", value)


@pulumi.input_type
class _ManagerState:
    def __init__(__self__, *,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 require_peer_certificate: Optional[pulumi.Input[bool]] = None,
                 upgrade_policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Manager resources.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ca_certificate: Device CA certificate.
        :param pulumi.Input[str] certificate: Device certificate.
        :param pulumi.Input[bool] enabled: Disable or enable CAPsMAN functionality.
        :param pulumi.Input[str] package_path: Folder location for the RouterOS packages. For example, use '/upgrade' to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.
        :param pulumi.Input[bool] require_peer_certificate: Require all connecting CAPs to have a valid certificate.
        :param pulumi.Input[str] upgrade_policy: Upgrade policy options.
        """
        if ___id_ is not None:
            pulumi.set(__self__, "___id_", ___id_)
        if ___path_ is not None:
            pulumi.set(__self__, "___path_", ___path_)
        if ca_certificate is not None:
            pulumi.set(__self__, "ca_certificate", ca_certificate)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if package_path is not None:
            pulumi.set(__self__, "package_path", package_path)
        if require_peer_certificate is not None:
            pulumi.set(__self__, "require_peer_certificate", require_peer_certificate)
        if upgrade_policy is not None:
            pulumi.set(__self__, "upgrade_policy", upgrade_policy)

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
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Device CA certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Device certificate.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable or enable CAPsMAN functionality.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> Optional[pulumi.Input[str]]:
        """
        Folder location for the RouterOS packages. For example, use '/upgrade' to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.
        """
        return pulumi.get(self, "package_path")

    @package_path.setter
    def package_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_path", value)

    @property
    @pulumi.getter(name="requirePeerCertificate")
    def require_peer_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Require all connecting CAPs to have a valid certificate.
        """
        return pulumi.get(self, "require_peer_certificate")

    @require_peer_certificate.setter
    def require_peer_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_peer_certificate", value)

    @property
    @pulumi.getter(name="upgradePolicy")
    def upgrade_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Upgrade policy options.
        """
        return pulumi.get(self, "upgrade_policy")

    @upgrade_policy.setter
    def upgrade_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upgrade_policy", value)


class Manager(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 require_peer_certificate: Optional[pulumi.Input[bool]] = None,
                 upgrade_policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        test_manager = routeros.caps_man.Manager("testManager",
            enabled=True,
            upgrade_policy="require-same-version")
        ```

        ## Import

        ```sh
         $ pulumi import routeros:CapsMan/manager:Manager test_manager .
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ca_certificate: Device CA certificate.
        :param pulumi.Input[str] certificate: Device certificate.
        :param pulumi.Input[bool] enabled: Disable or enable CAPsMAN functionality.
        :param pulumi.Input[str] package_path: Folder location for the RouterOS packages. For example, use '/upgrade' to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.
        :param pulumi.Input[bool] require_peer_certificate: Require all connecting CAPs to have a valid certificate.
        :param pulumi.Input[str] upgrade_policy: Upgrade policy options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ManagerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_routeros as routeros

        test_manager = routeros.caps_man.Manager("testManager",
            enabled=True,
            upgrade_policy="require-same-version")
        ```

        ## Import

        ```sh
         $ pulumi import routeros:CapsMan/manager:Manager test_manager .
        ```

        :param str resource_name: The name of the resource.
        :param ManagerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ___id_: Optional[pulumi.Input[int]] = None,
                 ___path_: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 require_peer_certificate: Optional[pulumi.Input[bool]] = None,
                 upgrade_policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ManagerArgs.__new__(ManagerArgs)

            __props__.__dict__["___id_"] = ___id_
            __props__.__dict__["___path_"] = ___path_
            __props__.__dict__["ca_certificate"] = ca_certificate
            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["package_path"] = package_path
            __props__.__dict__["require_peer_certificate"] = require_peer_certificate
            __props__.__dict__["upgrade_policy"] = upgrade_policy
        super(Manager, __self__).__init__(
            'routeros:CapsMan/manager:Manager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ___id_: Optional[pulumi.Input[int]] = None,
            ___path_: Optional[pulumi.Input[str]] = None,
            ca_certificate: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            package_path: Optional[pulumi.Input[str]] = None,
            require_peer_certificate: Optional[pulumi.Input[bool]] = None,
            upgrade_policy: Optional[pulumi.Input[str]] = None) -> 'Manager':
        """
        Get an existing Manager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ___id_: <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ___path_: <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
        :param pulumi.Input[str] ca_certificate: Device CA certificate.
        :param pulumi.Input[str] certificate: Device certificate.
        :param pulumi.Input[bool] enabled: Disable or enable CAPsMAN functionality.
        :param pulumi.Input[str] package_path: Folder location for the RouterOS packages. For example, use '/upgrade' to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.
        :param pulumi.Input[bool] require_peer_certificate: Require all connecting CAPs to have a valid certificate.
        :param pulumi.Input[str] upgrade_policy: Upgrade policy options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ManagerState.__new__(_ManagerState)

        __props__.__dict__["___id_"] = ___id_
        __props__.__dict__["___path_"] = ___path_
        __props__.__dict__["ca_certificate"] = ca_certificate
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["package_path"] = package_path
        __props__.__dict__["require_peer_certificate"] = require_peer_certificate
        __props__.__dict__["upgrade_policy"] = upgrade_policy
        return Manager(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> pulumi.Output[Optional[str]]:
        """
        Device CA certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional[str]]:
        """
        Device certificate.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable or enable CAPsMAN functionality.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> pulumi.Output[Optional[str]]:
        """
        Folder location for the RouterOS packages. For example, use '/upgrade' to specify the upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.
        """
        return pulumi.get(self, "package_path")

    @property
    @pulumi.getter(name="requirePeerCertificate")
    def require_peer_certificate(self) -> pulumi.Output[Optional[bool]]:
        """
        Require all connecting CAPs to have a valid certificate.
        """
        return pulumi.get(self, "require_peer_certificate")

    @property
    @pulumi.getter(name="upgradePolicy")
    def upgrade_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Upgrade policy options.
        """
        return pulumi.get(self, "upgrade_policy")

