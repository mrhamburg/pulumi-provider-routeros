# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 hosturl: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] ca_certificate: Path to MikroTik's certificate authority file.
        :param pulumi.Input[str] hosturl: URL of the MikroTik router, default is TLS connection to REST. * API: api[s]://host[:port] * api://router.local *
               apis://router.local:8729 * REST: https://host * https://router.local * router.local * 127.0.0.1 export
               ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
        :param pulumi.Input[bool] insecure: Whether to verify the SSL certificate or not.
        :param pulumi.Input[str] password: Password for the MikroTik user.
        :param pulumi.Input[str] username: Username for the MikroTik WEB/Winbox. export ROS_USERNAME=admin or export MIKROTIK_USER=admin
        """
        if ca_certificate is None:
            ca_certificate = _utilities.get_env('ROS_CA_CERTIFICATE')
        if ca_certificate is not None:
            pulumi.set(__self__, "ca_certificate", ca_certificate)
        if hosturl is None:
            hosturl = _utilities.get_env('ROS_HOSTURL')
        if hosturl is not None:
            pulumi.set(__self__, "hosturl", hosturl)
        if insecure is None:
            insecure = _utilities.get_env_bool('ROS_INSECURE')
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if password is None:
            password = _utilities.get_env('ROS_PASSWORD')
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is None:
            username = _utilities.get_env('ROS_USERNAME')
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Path to MikroTik's certificate authority file.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter
    def hosturl(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the MikroTik router, default is TLS connection to REST. * API: api[s]://host[:port] * api://router.local *
        apis://router.local:8729 * REST: https://host * https://router.local * router.local * 127.0.0.1 export
        ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
        """
        return pulumi.get(self, "hosturl")

    @hosturl.setter
    def hosturl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hosturl", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to verify the SSL certificate or not.
        """
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for the MikroTik user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Username for the MikroTik WEB/Winbox. export ROS_USERNAME=admin or export MIKROTIK_USER=admin
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 hosturl: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the routeros package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_certificate: Path to MikroTik's certificate authority file.
        :param pulumi.Input[str] hosturl: URL of the MikroTik router, default is TLS connection to REST. * API: api[s]://host[:port] * api://router.local *
               apis://router.local:8729 * REST: https://host * https://router.local * router.local * 127.0.0.1 export
               ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
        :param pulumi.Input[bool] insecure: Whether to verify the SSL certificate or not.
        :param pulumi.Input[str] password: Password for the MikroTik user.
        :param pulumi.Input[str] username: Username for the MikroTik WEB/Winbox. export ROS_USERNAME=admin or export MIKROTIK_USER=admin
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the routeros package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 hosturl: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if ca_certificate is None:
                ca_certificate = _utilities.get_env('ROS_CA_CERTIFICATE')
            __props__.__dict__["ca_certificate"] = None if ca_certificate is None else pulumi.Output.secret(ca_certificate)
            if hosturl is None:
                hosturl = _utilities.get_env('ROS_HOSTURL')
            __props__.__dict__["hosturl"] = hosturl
            if insecure is None:
                insecure = _utilities.get_env_bool('ROS_INSECURE')
            __props__.__dict__["insecure"] = pulumi.Output.from_input(insecure).apply(pulumi.runtime.to_json) if insecure is not None else None
            if password is None:
                password = _utilities.get_env('ROS_PASSWORD')
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if username is None:
                username = _utilities.get_env('ROS_USERNAME')
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["caCertificate", "password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'routeros',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> pulumi.Output[Optional[str]]:
        """
        Path to MikroTik's certificate authority file.
        """
        return pulumi.get(self, "ca_certificate")

    @property
    @pulumi.getter
    def hosturl(self) -> pulumi.Output[Optional[str]]:
        """
        URL of the MikroTik router, default is TLS connection to REST. * API: api[s]://host[:port] * api://router.local *
        apis://router.local:8729 * REST: https://host * https://router.local * router.local * 127.0.0.1 export
        ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
        """
        return pulumi.get(self, "hosturl")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password for the MikroTik user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        Username for the MikroTik WEB/Winbox. export ROS_USERNAME=admin or export MIKROTIK_USER=admin
        """
        return pulumi.get(self, "username")

