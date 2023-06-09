// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-routeros/sdk/go/routeros"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "routeros:Ip/dhcpClient:DhcpClient":
		r = &DhcpClient{}
	case "routeros:Ip/dhcpIpClient:DhcpIpClient":
		r = &DhcpIpClient{}
	case "routeros:Ip/dhcpIpServer:DhcpIpServer":
		r = &DhcpIpServer{}
	case "routeros:Ip/dhcpIpServerLease:DhcpIpServerLease":
		r = &DhcpIpServerLease{}
	case "routeros:Ip/dhcpIpServerNetwork:DhcpIpServerNetwork":
		r = &DhcpIpServerNetwork{}
	case "routeros:Ip/dhcpServer:DhcpServer":
		r = &DhcpServer{}
	case "routeros:Ip/dhcpServerLease:DhcpServerLease":
		r = &DhcpServerLease{}
	case "routeros:Ip/dhcpServerNetwork:DhcpServerNetwork":
		r = &DhcpServerNetwork{}
	case "routeros:Ip/dns:Dns":
		r = &Dns{}
	case "routeros:Ip/dnsRecord:DnsRecord":
		r = &DnsRecord{}
	case "routeros:Ip/firewallAddrList:FirewallAddrList":
		r = &FirewallAddrList{}
	case "routeros:Ip/firewallFilter:FirewallFilter":
		r = &FirewallFilter{}
	case "routeros:Ip/firewallMangle:FirewallMangle":
		r = &FirewallMangle{}
	case "routeros:Ip/firewallNat:FirewallNat":
		r = &FirewallNat{}
	case "routeros:Ip/ipDns:IpDns":
		r = &IpDns{}
	case "routeros:Ip/ipDnsRecord:IpDnsRecord":
		r = &IpDnsRecord{}
	case "routeros:Ip/ipFirewallAddrList:IpFirewallAddrList":
		r = &IpFirewallAddrList{}
	case "routeros:Ip/ipFirewallFilter:IpFirewallFilter":
		r = &IpFirewallFilter{}
	case "routeros:Ip/ipFirewallMangle:IpFirewallMangle":
		r = &IpFirewallMangle{}
	case "routeros:Ip/ipFirewallNat:IpFirewallNat":
		r = &IpFirewallNat{}
	case "routeros:Ip/pool:Pool":
		r = &Pool{}
	case "routeros:Ip/route:Route":
		r = &Route{}
	case "routeros:Ip/service:Service":
		r = &Service{}
	case "routeros:Ip/v4Address:V4Address":
		r = &V4Address{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := routeros.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpIpClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpIpServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpIpServerLease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpIpServerNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpServerLease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dhcpServerNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/dnsRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/firewallAddrList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/firewallFilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/firewallMangle",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/firewallNat",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/ipDns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/ipDnsRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/ipFirewallAddrList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/ipFirewallFilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/ipFirewallMangle",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/ipFirewallNat",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/pool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"routeros",
		"Ip/v4Address",
		&module{version},
	)
}
