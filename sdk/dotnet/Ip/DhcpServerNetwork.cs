// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.Ip
{
    /// <summary>
    /// ## # routeros.Ip.DhcpServerNetwork (Resource)
    /// 
    /// ***
    /// 
    /// #### This is an alias for backwards compatibility between plugin versions.
    /// Please see documentation for routeros.Ip.DhcpIpServerNetwork
    /// </summary>
    [RouterosResourceType("routeros:Ip/dhcpServerNetwork:DhcpServerNetwork")]
    public partial class DhcpServerNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// &lt;em&gt;Resource ID type (.id / name). This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Output("___id_")]
        public Output<int?> ___id_ { get; private set; } = null!;

        /// <summary>
        /// &lt;em&gt;Resource path for CRUD operations. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Output("___path_")]
        public Output<string?> ___path_ { get; private set; } = null!;

        /// <summary>
        /// The network DHCP server(s) will lease addresses from.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Boot filename.
        /// </summary>
        [Output("bootFileName")]
        public Output<string?> BootFileName { get; private set; } = null!;

        /// <summary>
        /// A comma-separated list of IP addresses for one or more CAPsMAN system managers. DHCP Option 138 (capwap) will be used.
        /// </summary>
        [Output("capsManager")]
        public Output<string?> CapsManager { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Add additional DHCP options from the option list.
        /// </summary>
        [Output("dhcpOption")]
        public Output<string?> DhcpOption { get; private set; } = null!;

        /// <summary>
        /// Add an additional set of DHCP options.
        /// </summary>
        [Output("dhcpOptionSet")]
        public Output<string?> DhcpOptionSet { get; private set; } = null!;

        /// <summary>
        /// If set, then DHCP Server will not pass dynamic DNS servers configured on the router to the DHCP clients if no DNS Server
        /// in DNS-server is set.
        /// </summary>
        [Output("dnsNone")]
        public Output<bool?> DnsNone { get; private set; } = null!;

        /// <summary>
        /// the DHCP client will use these as the default DNS servers. Two comma-separated DNS servers can be specified to be used
        /// by the DHCP client as primary and secondary DNS servers.
        /// </summary>
        [Output("dnsServer")]
        public Output<string?> DnsServer { get; private set; } = null!;

        /// <summary>
        /// The DHCP client will use this as the 'DNS domain' setting for the network adapter.
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
        /// modified.
        /// </summary>
        [Output("dynamic")]
        public Output<bool> Dynamic { get; private set; } = null!;

        /// <summary>
        /// The default gateway to be used by DHCP Client.
        /// </summary>
        [Output("gateway")]
        public Output<string?> Gateway { get; private set; } = null!;

        /// <summary>
        /// The actual network mask is to be used by the DHCP client. If set to '0' - netmask from network address will be used.
        /// </summary>
        [Output("netmask")]
        public Output<int?> Netmask { get; private set; } = null!;

        /// <summary>
        /// The IP address of the next server to use in bootstrap.
        /// </summary>
        [Output("nextServer")]
        public Output<string?> NextServer { get; private set; } = null!;

        /// <summary>
        /// The DHCP client will use these as the default NTP servers. Two comma-separated NTP servers can be specified to be used
        /// by the DHCP client as primary and secondary NTP servers
        /// </summary>
        [Output("ntpServer")]
        public Output<string?> NtpServer { get; private set; } = null!;

        /// <summary>
        /// The Windows DHCP client will use these as the default WINS servers. Two comma-separated WINS servers can be specified to
        /// be used by the DHCP client as primary and secondary WINS servers
        /// </summary>
        [Output("winsServer")]
        public Output<string?> WinsServer { get; private set; } = null!;


        /// <summary>
        /// Create a DhcpServerNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DhcpServerNetwork(string name, DhcpServerNetworkArgs args, CustomResourceOptions? options = null)
            : base("routeros:Ip/dhcpServerNetwork:DhcpServerNetwork", name, args ?? new DhcpServerNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DhcpServerNetwork(string name, Input<string> id, DhcpServerNetworkState? state = null, CustomResourceOptions? options = null)
            : base("routeros:Ip/dhcpServerNetwork:DhcpServerNetwork", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DhcpServerNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DhcpServerNetwork Get(string name, Input<string> id, DhcpServerNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new DhcpServerNetwork(name, id, state, options);
        }
    }

    public sealed class DhcpServerNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;em&gt;Resource ID type (.id / name). This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___id_")]
        public Input<int>? ___id_ { get; set; }

        /// <summary>
        /// &lt;em&gt;Resource path for CRUD operations. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___path_")]
        public Input<string>? ___path_ { get; set; }

        /// <summary>
        /// The network DHCP server(s) will lease addresses from.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Boot filename.
        /// </summary>
        [Input("bootFileName")]
        public Input<string>? BootFileName { get; set; }

        /// <summary>
        /// A comma-separated list of IP addresses for one or more CAPsMAN system managers. DHCP Option 138 (capwap) will be used.
        /// </summary>
        [Input("capsManager")]
        public Input<string>? CapsManager { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Add additional DHCP options from the option list.
        /// </summary>
        [Input("dhcpOption")]
        public Input<string>? DhcpOption { get; set; }

        /// <summary>
        /// Add an additional set of DHCP options.
        /// </summary>
        [Input("dhcpOptionSet")]
        public Input<string>? DhcpOptionSet { get; set; }

        /// <summary>
        /// If set, then DHCP Server will not pass dynamic DNS servers configured on the router to the DHCP clients if no DNS Server
        /// in DNS-server is set.
        /// </summary>
        [Input("dnsNone")]
        public Input<bool>? DnsNone { get; set; }

        /// <summary>
        /// the DHCP client will use these as the default DNS servers. Two comma-separated DNS servers can be specified to be used
        /// by the DHCP client as primary and secondary DNS servers.
        /// </summary>
        [Input("dnsServer")]
        public Input<string>? DnsServer { get; set; }

        /// <summary>
        /// The DHCP client will use this as the 'DNS domain' setting for the network adapter.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The default gateway to be used by DHCP Client.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// The actual network mask is to be used by the DHCP client. If set to '0' - netmask from network address will be used.
        /// </summary>
        [Input("netmask")]
        public Input<int>? Netmask { get; set; }

        /// <summary>
        /// The IP address of the next server to use in bootstrap.
        /// </summary>
        [Input("nextServer")]
        public Input<string>? NextServer { get; set; }

        /// <summary>
        /// The DHCP client will use these as the default NTP servers. Two comma-separated NTP servers can be specified to be used
        /// by the DHCP client as primary and secondary NTP servers
        /// </summary>
        [Input("ntpServer")]
        public Input<string>? NtpServer { get; set; }

        /// <summary>
        /// The Windows DHCP client will use these as the default WINS servers. Two comma-separated WINS servers can be specified to
        /// be used by the DHCP client as primary and secondary WINS servers
        /// </summary>
        [Input("winsServer")]
        public Input<string>? WinsServer { get; set; }

        public DhcpServerNetworkArgs()
        {
        }
        public static new DhcpServerNetworkArgs Empty => new DhcpServerNetworkArgs();
    }

    public sealed class DhcpServerNetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;em&gt;Resource ID type (.id / name). This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___id_")]
        public Input<int>? ___id_ { get; set; }

        /// <summary>
        /// &lt;em&gt;Resource path for CRUD operations. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___path_")]
        public Input<string>? ___path_ { get; set; }

        /// <summary>
        /// The network DHCP server(s) will lease addresses from.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Boot filename.
        /// </summary>
        [Input("bootFileName")]
        public Input<string>? BootFileName { get; set; }

        /// <summary>
        /// A comma-separated list of IP addresses for one or more CAPsMAN system managers. DHCP Option 138 (capwap) will be used.
        /// </summary>
        [Input("capsManager")]
        public Input<string>? CapsManager { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Add additional DHCP options from the option list.
        /// </summary>
        [Input("dhcpOption")]
        public Input<string>? DhcpOption { get; set; }

        /// <summary>
        /// Add an additional set of DHCP options.
        /// </summary>
        [Input("dhcpOptionSet")]
        public Input<string>? DhcpOptionSet { get; set; }

        /// <summary>
        /// If set, then DHCP Server will not pass dynamic DNS servers configured on the router to the DHCP clients if no DNS Server
        /// in DNS-server is set.
        /// </summary>
        [Input("dnsNone")]
        public Input<bool>? DnsNone { get; set; }

        /// <summary>
        /// the DHCP client will use these as the default DNS servers. Two comma-separated DNS servers can be specified to be used
        /// by the DHCP client as primary and secondary DNS servers.
        /// </summary>
        [Input("dnsServer")]
        public Input<string>? DnsServer { get; set; }

        /// <summary>
        /// The DHCP client will use this as the 'DNS domain' setting for the network adapter.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
        /// modified.
        /// </summary>
        [Input("dynamic")]
        public Input<bool>? Dynamic { get; set; }

        /// <summary>
        /// The default gateway to be used by DHCP Client.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// The actual network mask is to be used by the DHCP client. If set to '0' - netmask from network address will be used.
        /// </summary>
        [Input("netmask")]
        public Input<int>? Netmask { get; set; }

        /// <summary>
        /// The IP address of the next server to use in bootstrap.
        /// </summary>
        [Input("nextServer")]
        public Input<string>? NextServer { get; set; }

        /// <summary>
        /// The DHCP client will use these as the default NTP servers. Two comma-separated NTP servers can be specified to be used
        /// by the DHCP client as primary and secondary NTP servers
        /// </summary>
        [Input("ntpServer")]
        public Input<string>? NtpServer { get; set; }

        /// <summary>
        /// The Windows DHCP client will use these as the default WINS servers. Two comma-separated WINS servers can be specified to
        /// be used by the DHCP client as primary and secondary WINS servers
        /// </summary>
        [Input("winsServer")]
        public Input<string>? WinsServer { get; set; }

        public DhcpServerNetworkState()
        {
        }
        public static new DhcpServerNetworkState Empty => new DhcpServerNetworkState();
    }
}
