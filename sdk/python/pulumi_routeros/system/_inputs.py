# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CertificateSignArgs',
]

@pulumi.input_type
class CertificateSignArgs:
    def __init__(__self__, *,
                 ca: Optional[pulumi.Input[str]] = None,
                 ca_crl_host: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] ca: Which CA to use if signing issued certificates.
        :param pulumi.Input[str] ca_crl_host: CRL host if issuing CA certificate.
        """
        if ca is not None:
            pulumi.set(__self__, "ca", ca)
        if ca_crl_host is not None:
            pulumi.set(__self__, "ca_crl_host", ca_crl_host)

    @property
    @pulumi.getter
    def ca(self) -> Optional[pulumi.Input[str]]:
        """
        Which CA to use if signing issued certificates.
        """
        return pulumi.get(self, "ca")

    @ca.setter
    def ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca", value)

    @property
    @pulumi.getter(name="caCrlHost")
    def ca_crl_host(self) -> Optional[pulumi.Input[str]]:
        """
        CRL host if issuing CA certificate.
        """
        return pulumi.get(self, "ca_crl_host")

    @ca_crl_host.setter
    def ca_crl_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_crl_host", value)

