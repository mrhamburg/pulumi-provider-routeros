// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.Inputs
{

    public sealed class DatasourceFirewallNatArgs : global::Pulumi.InvokeArgs
    {
        [Input("action", required: true)]
        public string Action { get; set; } = null!;

        [Input("addressList", required: true)]
        public string AddressList { get; set; } = null!;

        [Input("addressListTimeout", required: true)]
        public string AddressListTimeout { get; set; } = null!;

        [Input("bytes", required: true)]
        public int Bytes { get; set; }

        [Input("chain", required: true)]
        public string Chain { get; set; } = null!;

        [Input("comment", required: true)]
        public string Comment { get; set; } = null!;

        [Input("connectionBytes", required: true)]
        public string ConnectionBytes { get; set; } = null!;

        [Input("connectionLimit", required: true)]
        public string ConnectionLimit { get; set; } = null!;

        [Input("connectionMark", required: true)]
        public string ConnectionMark { get; set; } = null!;

        [Input("connectionRate", required: true)]
        public string ConnectionRate { get; set; } = null!;

        [Input("connectionType", required: true)]
        public string ConnectionType { get; set; } = null!;

        [Input("content", required: true)]
        public string Content { get; set; } = null!;

        [Input("disabled", required: true)]
        public bool Disabled { get; set; }

        [Input("dscp", required: true)]
        public int Dscp { get; set; }

        [Input("dstAddress", required: true)]
        public string DstAddress { get; set; } = null!;

        [Input("dstAddressList", required: true)]
        public string DstAddressList { get; set; } = null!;

        [Input("dstAddressType", required: true)]
        public string DstAddressType { get; set; } = null!;

        [Input("dstLimit", required: true)]
        public string DstLimit { get; set; } = null!;

        [Input("dstPort", required: true)]
        public string DstPort { get; set; } = null!;

        [Input("dynamic", required: true)]
        public bool Dynamic { get; set; }

        [Input("filter")]
        private Dictionary<string, object>? _filter;

        /// <summary>
        /// Additional request filtering options.
        /// </summary>
        public Dictionary<string, object> Filter
        {
            get => _filter ?? (_filter = new Dictionary<string, object>());
            set => _filter = value;
        }

        [Input("fragment", required: true)]
        public bool Fragment { get; set; }

        [Input("hotspot", required: true)]
        public string Hotspot { get; set; } = null!;

        [Input("icmpOptions", required: true)]
        public string IcmpOptions { get; set; } = null!;

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("inBridgePort", required: true)]
        public string InBridgePort { get; set; } = null!;

        [Input("inBridgePortList", required: true)]
        public string InBridgePortList { get; set; } = null!;

        [Input("inInterface", required: true)]
        public string InInterface { get; set; } = null!;

        [Input("inInterfaceList", required: true)]
        public string InInterfaceList { get; set; } = null!;

        [Input("ingressPriority", required: true)]
        public int IngressPriority { get; set; }

        [Input("invalid", required: true)]
        public bool Invalid { get; set; }

        [Input("ipsecPolicy", required: true)]
        public string IpsecPolicy { get; set; } = null!;

        [Input("ipv4Options", required: true)]
        public string Ipv4Options { get; set; } = null!;

        [Input("jumpTarget", required: true)]
        public string JumpTarget { get; set; } = null!;

        [Input("layer7Protocol", required: true)]
        public string Layer7Protocol { get; set; } = null!;

        [Input("limit", required: true)]
        public string Limit { get; set; } = null!;

        [Input("log", required: true)]
        public bool Log { get; set; }

        [Input("logPrefix", required: true)]
        public string LogPrefix { get; set; } = null!;

        [Input("nth", required: true)]
        public string Nth { get; set; } = null!;

        [Input("outBridgePort", required: true)]
        public string OutBridgePort { get; set; } = null!;

        [Input("outBridgePortList", required: true)]
        public string OutBridgePortList { get; set; } = null!;

        [Input("outInterface", required: true)]
        public string OutInterface { get; set; } = null!;

        [Input("outInterfaceList", required: true)]
        public string OutInterfaceList { get; set; } = null!;

        [Input("packetMark", required: true)]
        public string PacketMark { get; set; } = null!;

        [Input("packetSize", required: true)]
        public string PacketSize { get; set; } = null!;

        [Input("packets", required: true)]
        public int Packets { get; set; }

        [Input("perConnectionClassifier", required: true)]
        public string PerConnectionClassifier { get; set; } = null!;

        [Input("port", required: true)]
        public string Port { get; set; } = null!;

        [Input("priority", required: true)]
        public int Priority { get; set; }

        [Input("protocol", required: true)]
        public string Protocol { get; set; } = null!;

        [Input("psd", required: true)]
        public string Psd { get; set; } = null!;

        [Input("random", required: true)]
        public int Random { get; set; }

        [Input("routingMark", required: true)]
        public string RoutingMark { get; set; } = null!;

        [Input("sameNotByDst", required: true)]
        public bool SameNotByDst { get; set; }

        [Input("srcAddress", required: true)]
        public string SrcAddress { get; set; } = null!;

        [Input("srcAddressList", required: true)]
        public string SrcAddressList { get; set; } = null!;

        [Input("srcAddressType", required: true)]
        public string SrcAddressType { get; set; } = null!;

        [Input("srcMacAddress", required: true)]
        public string SrcMacAddress { get; set; } = null!;

        [Input("srcPort", required: true)]
        public string SrcPort { get; set; } = null!;

        [Input("tcpMss", required: true)]
        public string TcpMss { get; set; } = null!;

        [Input("time", required: true)]
        public string Time { get; set; } = null!;

        [Input("toAddresses", required: true)]
        public string ToAddresses { get; set; } = null!;

        [Input("toPorts", required: true)]
        public string ToPorts { get; set; } = null!;

        [Input("ttl", required: true)]
        public string Ttl { get; set; } = null!;

        public DatasourceFirewallNatArgs()
        {
        }
        public static new DatasourceFirewallNatArgs Empty => new DatasourceFirewallNatArgs();
    }
}
