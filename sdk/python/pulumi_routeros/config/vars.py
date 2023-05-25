# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('routeros')


class _ExportableConfig(types.ModuleType):
    @property
    def ca_certificate(self) -> Optional[str]:
        """
        Path to MikroTik's certificate authority file.
        """
        return __config__.get('caCertificate') or _utilities.get_env('ROS_CA_CERTIFICATE')

    @property
    def hosturl(self) -> Optional[str]:
        """
        URL of the MikroTik router, default is TLS connection to REST. * API: api[s]://host[:port] * api://router.local *
        apis://router.local:8729 * REST: https://host * https://router.local * router.local * 127.0.0.1 export
        ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
        """
        return __config__.get('hosturl') or _utilities.get_env('ROS_HOSTURL')

    @property
    def insecure(self) -> Optional[bool]:
        """
        Whether to verify the SSL certificate or not.
        """
        return __config__.get_bool('insecure') or _utilities.get_env_bool('ROS_INSECURE')

    @property
    def password(self) -> Optional[str]:
        """
        Password for the MikroTik user.
        """
        return __config__.get('password') or _utilities.get_env('ROS_PASSWORD')

    @property
    def username(self) -> Optional[str]:
        """
        Username for the MikroTik WEB/Winbox. export ROS_USERNAME=admin or export MIKROTIK_USER=admin
        """
        return __config__.get('username') or _utilities.get_env('ROS_USERNAME')
