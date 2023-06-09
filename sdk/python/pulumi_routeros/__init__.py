# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_routeros.capsman as __capsman
    capsman = __capsman
    import pulumi_routeros.config as __config
    config = __config
    import pulumi_routeros.iface as __iface
    iface = __iface
    import pulumi_routeros.ip as __ip
    ip = __ip
    import pulumi_routeros.ipv6 as __ipv6
    ipv6 = __ipv6
    import pulumi_routeros.openvpn as __openvpn
    openvpn = __openvpn
    import pulumi_routeros.ppp as __ppp
    ppp = __ppp
    import pulumi_routeros.routing as __routing
    routing = __routing
    import pulumi_routeros.system as __system
    system = __system
else:
    capsman = _utilities.lazy_import('pulumi_routeros.capsman')
    config = _utilities.lazy_import('pulumi_routeros.config')
    iface = _utilities.lazy_import('pulumi_routeros.iface')
    ip = _utilities.lazy_import('pulumi_routeros.ip')
    ipv6 = _utilities.lazy_import('pulumi_routeros.ipv6')
    openvpn = _utilities.lazy_import('pulumi_routeros.openvpn')
    ppp = _utilities.lazy_import('pulumi_routeros.ppp')
    routing = _utilities.lazy_import('pulumi_routeros.routing')
    system = _utilities.lazy_import('pulumi_routeros.system')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "routeros",
  "mod": "CapsMan/aaa",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/aaa:Aaa": "Aaa"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/channel",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/channel:Channel": "Channel"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/configuration",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/configuration:Configuration": "Configuration"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/datapath",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/datapath:Datapath": "Datapath"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/manager",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/manager:Manager": "Manager"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/managerInterface",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/managerInterface:ManagerInterface": "ManagerInterface"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/provisioning",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/provisioning:Provisioning": "Provisioning"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/rates",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/rates:Rates": "Rates"
  }
 },
 {
  "pkg": "routeros",
  "mod": "CapsMan/security",
  "fqn": "pulumi_routeros.capsman",
  "classes": {
   "routeros:CapsMan/security:Security": "Security"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/bridge",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/bridge:Bridge": "Bridge"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/bridgePort",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/bridgePort:BridgePort": "BridgePort"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/bridgeVlan",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/bridgeVlan:BridgeVlan": "BridgeVlan"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/gre",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/gre:Gre": "Gre"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceBridge",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceBridge:InterfaceBridge": "InterfaceBridge"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceBridgePort",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceBridgePort:InterfaceBridgePort": "InterfaceBridgePort"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceBridgeVlan",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceBridgeVlan:InterfaceBridgeVlan": "InterfaceBridgeVlan"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceGre",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceGre:InterfaceGre": "InterfaceGre"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceVlan",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceVlan:InterfaceVlan": "InterfaceVlan"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceVrrp",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceVrrp:InterfaceVrrp": "InterfaceVrrp"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceWireguard",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceWireguard:InterfaceWireguard": "InterfaceWireguard"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/interfaceWireguardPeer",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/interfaceWireguardPeer:InterfaceWireguardPeer": "InterfaceWireguardPeer"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/list",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/list:List": "List"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/listMember",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/listMember:ListMember": "ListMember"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/openVpnServer",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/openVpnServer:OpenVpnServer": "OpenVpnServer"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/vlan",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/vlan:Vlan": "Vlan"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/vrrp",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/vrrp:Vrrp": "Vrrp"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/wireguard",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/wireguard:Wireguard": "Wireguard"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Iface/wireguardPeer",
  "fqn": "pulumi_routeros.iface",
  "classes": {
   "routeros:Iface/wireguardPeer:WireguardPeer": "WireguardPeer"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpClient",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpClient:DhcpClient": "DhcpClient"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpIpClient",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpIpClient:DhcpIpClient": "DhcpIpClient"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpIpServer",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpIpServer:DhcpIpServer": "DhcpIpServer"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpIpServerLease",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpIpServerLease:DhcpIpServerLease": "DhcpIpServerLease"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpIpServerNetwork",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpIpServerNetwork:DhcpIpServerNetwork": "DhcpIpServerNetwork"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpServer",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpServer:DhcpServer": "DhcpServer"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpServerLease",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpServerLease:DhcpServerLease": "DhcpServerLease"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dhcpServerNetwork",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dhcpServerNetwork:DhcpServerNetwork": "DhcpServerNetwork"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dns",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dns:Dns": "Dns"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/dnsRecord",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/dnsRecord:DnsRecord": "DnsRecord"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/firewallAddrList",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/firewallAddrList:FirewallAddrList": "FirewallAddrList"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/firewallFilter",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/firewallFilter:FirewallFilter": "FirewallFilter"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/firewallMangle",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/firewallMangle:FirewallMangle": "FirewallMangle"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/firewallNat",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/firewallNat:FirewallNat": "FirewallNat"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/ipDns",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/ipDns:IpDns": "IpDns"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/ipDnsRecord",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/ipDnsRecord:IpDnsRecord": "IpDnsRecord"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/ipFirewallAddrList",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/ipFirewallAddrList:IpFirewallAddrList": "IpFirewallAddrList"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/ipFirewallFilter",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/ipFirewallFilter:IpFirewallFilter": "IpFirewallFilter"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/ipFirewallMangle",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/ipFirewallMangle:IpFirewallMangle": "IpFirewallMangle"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/ipFirewallNat",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/ipFirewallNat:IpFirewallNat": "IpFirewallNat"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/pool",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/pool:Pool": "Pool"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/route",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/route:Route": "Route"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/service",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/service:Service": "Service"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ip/v4Address",
  "fqn": "pulumi_routeros.ip",
  "classes": {
   "routeros:Ip/v4Address:V4Address": "V4Address"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ipv6/firewallFilter",
  "fqn": "pulumi_routeros.ipv6",
  "classes": {
   "routeros:Ipv6/firewallFilter:FirewallFilter": "FirewallFilter"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ipv6/route",
  "fqn": "pulumi_routeros.ipv6",
  "classes": {
   "routeros:Ipv6/route:Route": "Route"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ipv6/v6Address",
  "fqn": "pulumi_routeros.ipv6",
  "classes": {
   "routeros:Ipv6/v6Address:V6Address": "V6Address"
  }
 },
 {
  "pkg": "routeros",
  "mod": "OpenVpn/server",
  "fqn": "pulumi_routeros.openvpn",
  "classes": {
   "routeros:OpenVpn/server:Server": "Server"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ppp/profile",
  "fqn": "pulumi_routeros.ppp",
  "classes": {
   "routeros:Ppp/profile:Profile": "Profile"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Ppp/secret",
  "fqn": "pulumi_routeros.ppp",
  "classes": {
   "routeros:Ppp/secret:Secret": "Secret"
  }
 },
 {
  "pkg": "routeros",
  "mod": "Routing/routing",
  "fqn": "pulumi_routeros.routing",
  "classes": {
   "routeros:Routing/routing:Routing": "Routing"
  }
 },
 {
  "pkg": "routeros",
  "mod": "System/certificate",
  "fqn": "pulumi_routeros.system",
  "classes": {
   "routeros:System/certificate:Certificate": "Certificate"
  }
 },
 {
  "pkg": "routeros",
  "mod": "System/identity",
  "fqn": "pulumi_routeros.system",
  "classes": {
   "routeros:System/identity:Identity": "Identity"
  }
 },
 {
  "pkg": "routeros",
  "mod": "System/scheduler",
  "fqn": "pulumi_routeros.system",
  "classes": {
   "routeros:System/scheduler:Scheduler": "Scheduler"
  }
 },
 {
  "pkg": "routeros",
  "mod": "System/systemIdentity",
  "fqn": "pulumi_routeros.system",
  "classes": {
   "routeros:System/systemIdentity:SystemIdentity": "SystemIdentity"
  }
 },
 {
  "pkg": "routeros",
  "mod": "System/systemScheduler",
  "fqn": "pulumi_routeros.system",
  "classes": {
   "routeros:System/systemScheduler:SystemScheduler": "SystemScheduler"
  }
 },
 {
  "pkg": "routeros",
  "mod": "System/user",
  "fqn": "pulumi_routeros.system",
  "classes": {
   "routeros:System/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "routeros",
  "token": "pulumi:providers:routeros",
  "fqn": "pulumi_routeros",
  "class": "Provider"
 }
]
"""
)
