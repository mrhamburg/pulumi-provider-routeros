// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package routeros

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/mrhamburg/pulumi-provider-routeros/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/terraform-routeros/terraform-provider-routeros/routeros"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "routeros"
	// modules:
	mainMod      = "index" // the routeros module
	rosInterface = "Iface"
	rosCapsMan   = "CapsMan"
	rosIp        = "Ip"
	rosIpv6      = "Ipv6"
	rosOpenVpn   = "OpenVpn"
	rosPpp       = "Ppp"
	rosRouting   = "Routing"
	rosSystem    = "System"
)

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// routerOsMember manufactures a type token for the RouterOs package and the given module and type.
func routerOsMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// routerOsType manufactures a type token for the RouterOs package and the given module and type.
func routerOsType(mod string, typ string) tokens.Type {
	return tokens.Type(routerOsMember(mod, typ))
}

// routerOsDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the RouterOs package and names the file by simply lower casing the data
// source's first character.
func routerOsDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return routerOsMember(mod+"/"+fn, res)
}

// routerOsResource manufactures a standard resource token given a module and resource name.
// It automatically uses the routeros package and names the file by simply lower casing the resource's
// first character.
func routerOsResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return routerOsType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(routeros.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "routeros",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumi",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing routeros cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "routeros", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/mrhamburg/pulumi-provider-routeros",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "terraform-routeros",
		Config: map[string]*tfbridge.SchemaInfo{
			"hosturl": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ROS_HOSTURL"},
				},
			},
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ROS_USERNAME"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ROS_PASSWORD"},
				},
				Secret: boolRef(true),
			},
			"ca_certificate": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ROS_CA_CERTIFICATE"},
				},
				Secret: boolRef(true),
			},
			"insecure": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ROS_INSECURE"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"routeros_capsman_aaa":               {Tok: routerOsResource(rosCapsMan, "Aaa")},
			"routeros_capsman_channel":           {Tok: routerOsResource(rosCapsMan, "Channel")},
			"routeros_capsman_configuration":     {Tok: routerOsResource(rosCapsMan, "Configuration")},
			"routeros_capsman_datapath":          {Tok: routerOsResource(rosCapsMan, "Datapath")},
			"routeros_capsman_manager":           {Tok: routerOsResource(rosCapsMan, "Manager")},
			"routeros_capsman_manager_interface": {Tok: routerOsResource(rosCapsMan, "ManagerInterface")},
			"routeros_capsman_provisioning":      {Tok: routerOsResource(rosCapsMan, "Provisioning")},
			"routeros_capsman_rates":             {Tok: routerOsResource(rosCapsMan, "Rates")},
			"routeros_capsman_security":          {Tok: routerOsResource(rosCapsMan, "Security")},
			"routeros_bridge":                    {Tok: routerOsResource(rosInterface, "Bridge")},
			"routeros_bridge_port":               {Tok: routerOsResource(rosInterface, "BridgePort")},
			"routeros_bridge_vlan":               {Tok: routerOsResource(rosInterface, "BridgeVlan")},
			"routeros_interface_bridge":          {Tok: routerOsResource(rosInterface, "InterfaceBridge")},
			"routeros_interface_bridge_port":     {Tok: routerOsResource(rosInterface, "InterfaceBridgePort")},
			"routeros_interface_bridge_vlan":     {Tok: routerOsResource(rosInterface, "InterfaceBridgeVlan")},
			"routeros_interface_gre":             {Tok: routerOsResource(rosInterface, "InterfaceGre")},
			"routeros_gre":                       {Tok: routerOsResource(rosInterface, "Gre")},
			"routeros_interface_list":            {Tok: routerOsResource(rosInterface, "List")},
			"routeros_interface_list_member":     {Tok: routerOsResource(rosInterface, "ListMember")},
			"routeros_interface_ovpn_server":     {Tok: routerOsResource(rosInterface, "OpenVpnServer")},
			"routeros_vlan":                      {Tok: routerOsResource(rosInterface, "Vlan")},
			"routeros_interface_vlan":            {Tok: routerOsResource(rosInterface, "InterfaceVlan")},
			"routeros_vrrp":                      {Tok: routerOsResource(rosInterface, "Vrrp")},
			"routeros_interface_vrrp":            {Tok: routerOsResource(rosInterface, "InterfaceVrrp")},
			"routeros_wireguard":                 {Tok: routerOsResource(rosInterface, "Wireguard")},
			"routeros_interface_wireguard":       {Tok: routerOsResource(rosInterface, "InterfaceWireguard")},
			"routeros_wireguard_peer":            {Tok: routerOsResource(rosInterface, "WireguardPeer")},
			"routeros_interface_wireguard_peer":  {Tok: routerOsResource(rosInterface, "InterfaceWireguardPeer")},
			"routeros_ip_address":                {Tok: routerOsResource(rosIp, "V4Address")},
			"routeros_dhcp_client":               {Tok: routerOsResource(rosIp, "DhcpClient")},
			"routeros_dhcp_server":               {Tok: routerOsResource(rosIp, "DhcpServer")},
			"routeros_dhcp_server_lease":         {Tok: routerOsResource(rosIp, "DhcpServerLease")},
			"routeros_dhcp_server_network":       {Tok: routerOsResource(rosIp, "DhcpServerNetwork")},
			"routeros_ip_dhcp_client":            {Tok: routerOsResource(rosIp, "DhcpIpClient")},
			"routeros_ip_dhcp_server":            {Tok: routerOsResource(rosIp, "DhcpIpServer")},
			"routeros_ip_dhcp_server_lease":      {Tok: routerOsResource(rosIp, "DhcpIpServerLease")},
			"routeros_ip_dhcp_server_network":    {Tok: routerOsResource(rosIp, "DhcpIpServerNetwork")},
			"routeros_dns":                       {Tok: routerOsResource(rosIp, "Dns")},
			"routeros_dns_record":                {Tok: routerOsResource(rosIp, "DnsRecord")},
			"routeros_firewall_addr_list":        {Tok: routerOsResource(rosIp, "FirewallAddrList")},
			"routeros_firewall_filter":           {Tok: routerOsResource(rosIp, "FirewallFilter")},
			"routeros_firewall_mangle":           {Tok: routerOsResource(rosIp, "FirewallMangle")},
			"routeros_firewall_nat":              {Tok: routerOsResource(rosIp, "FirewallNat")},
			"routeros_ip_dns":                    {Tok: routerOsResource(rosIp, "IpDns")},
			"routeros_ip_dns_record":             {Tok: routerOsResource(rosIp, "IpDnsRecord")},
			"routeros_ip_firewall_addr_list":     {Tok: routerOsResource(rosIp, "IpFirewallAddrList")},
			"routeros_ip_firewall_filter":        {Tok: routerOsResource(rosIp, "IpFirewallFilter")},
			"routeros_ip_firewall_mangle":        {Tok: routerOsResource(rosIp, "IpFirewallMangle")},
			"routeros_ip_firewall_nat":           {Tok: routerOsResource(rosIp, "IpFirewallNat")},
			"routeros_ip_pool":                   {Tok: routerOsResource(rosIp, "Pool")},
			"routeros_ip_route":                  {Tok: routerOsResource(rosIp, "Route")},
			"routeros_ip_service":                {Tok: routerOsResource(rosIp, "Service")},
			"routeros_ipv6_address":              {Tok: routerOsResource(rosIpv6, "V6Address")},
			"routeros_ipv6_firewall_filter":      {Tok: routerOsResource(rosIpv6, "FirewallFilter")},
			"routeros_ipv6_route":                {Tok: routerOsResource(rosIpv6, "Route")},
			"routeros_ovpn_server":               {Tok: routerOsResource(rosOpenVpn, "Server")},
			"routeros_ppp_profile":               {Tok: routerOsResource(rosPpp, "Profile")},
			"routeros_ppp_secret":                {Tok: routerOsResource(rosPpp, "Secret")},
			"routeros_routing_table":             {Tok: routerOsResource(rosRouting, "Routing")},
			"routeros_system_certificate":        {Tok: routerOsResource(rosSystem, "Certificate")},
			"routeros_identity":                  {Tok: routerOsResource(rosSystem, "Identity")},
			"routeros_system_identity":           {Tok: routerOsResource(rosSystem, "SystemIdentity")},
			"routeros_system_scheduler":          {Tok: routerOsResource(rosSystem, "SystemScheduler")},
			"routeros_scheduler":                 {Tok: routerOsResource(rosSystem, "Scheduler")},
			"routeros_system_user":               {Tok: routerOsResource(rosSystem, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"routeros_firewall": {
				Tok: routerOsDataSource(rosIp, "getFirewall"),
			},
			"routeros_interfaces": {
				Tok: routerOsDataSource(rosInterface, "getInterfaces"),
			},
			"routeros_ip_addresses": {
				Tok: routerOsDataSource(rosIp, "getAddresses"),
			},
			"routeros_ip_routes": {
				Tok: routerOsDataSource(rosIp, "getRoutes"),
			},
			"routeros_ipv6_addresses": {
				Tok: routerOsDataSource(rosIpv6, "getIpv6Addresses"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
