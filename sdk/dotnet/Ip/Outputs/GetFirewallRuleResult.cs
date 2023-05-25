// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.Ip.Outputs
{

    [OutputType]
    public sealed class GetFirewallRuleResult
    {
        public readonly string Action;
        public readonly string AddressListTimeout;
        public readonly int Bytes;
        public readonly string Chain;
        public readonly string Comment;
        public readonly string ConnectionBytes;
        public readonly string ConnectionLimit;
        public readonly string ConnectionMark;
        public readonly string ConnectionNatState;
        public readonly string ConnectionRate;
        public readonly string ConnectionState;
        public readonly string ConnectionType;
        public readonly string Content;
        public readonly bool Disabled;
        public readonly int Dscp;
        public readonly string DstAddress;
        public readonly string DstAddressList;
        public readonly string DstAddressType;
        public readonly string DstLimit;
        public readonly string DstPort;
        public readonly bool Dynamic;
        /// <summary>
        /// Additional request filtering options.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Filter;
        public readonly bool Fragment;
        public readonly string Hotspot;
        public readonly bool HwOffload;
        public readonly string IcmpOptions;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        public readonly string InBridgePort;
        public readonly string InBridgePortList;
        public readonly string InInterface;
        public readonly string InInterfaceList;
        public readonly int IngressPriority;
        public readonly bool Invalid;
        public readonly string IpsecPolicy;
        public readonly string Ipv4Options;
        public readonly string JumpTarget;
        public readonly string Layer7Protocol;
        public readonly string Limit;
        public readonly bool Log;
        public readonly string LogPrefix;
        public readonly string Nth;
        public readonly string OutBridgePort;
        public readonly string OutBridgePortList;
        public readonly string OutInterface;
        public readonly string OutInterfaceList;
        public readonly string PacketMark;
        public readonly string PacketSize;
        public readonly int Packets;
        public readonly string PerConnectionClassifier;
        public readonly string Port;
        public readonly int Priority;
        public readonly string Protocol;
        public readonly string Psd;
        public readonly int Random;
        public readonly string RejectWith;
        public readonly string RoutingMark;
        public readonly string RoutingTable;
        public readonly string SrcAddress;
        public readonly string SrcAddressList;
        public readonly string SrcAddressType;
        public readonly string SrcMacAddress;
        public readonly string SrcPort;
        public readonly string TcpFlags;
        public readonly string TcpMss;
        public readonly string Time;
        public readonly string TlsHost;
        public readonly string Ttl;

        [OutputConstructor]
        private GetFirewallRuleResult(
            string action,

            string addressListTimeout,

            int bytes,

            string chain,

            string comment,

            string connectionBytes,

            string connectionLimit,

            string connectionMark,

            string connectionNatState,

            string connectionRate,

            string connectionState,

            string connectionType,

            string content,

            bool disabled,

            int dscp,

            string dstAddress,

            string dstAddressList,

            string dstAddressType,

            string dstLimit,

            string dstPort,

            bool dynamic,

            ImmutableDictionary<string, object>? filter,

            bool fragment,

            string hotspot,

            bool hwOffload,

            string icmpOptions,

            string id,

            string inBridgePort,

            string inBridgePortList,

            string inInterface,

            string inInterfaceList,

            int ingressPriority,

            bool invalid,

            string ipsecPolicy,

            string ipv4Options,

            string jumpTarget,

            string layer7Protocol,

            string limit,

            bool log,

            string logPrefix,

            string nth,

            string outBridgePort,

            string outBridgePortList,

            string outInterface,

            string outInterfaceList,

            string packetMark,

            string packetSize,

            int packets,

            string perConnectionClassifier,

            string port,

            int priority,

            string protocol,

            string psd,

            int random,

            string rejectWith,

            string routingMark,

            string routingTable,

            string srcAddress,

            string srcAddressList,

            string srcAddressType,

            string srcMacAddress,

            string srcPort,

            string tcpFlags,

            string tcpMss,

            string time,

            string tlsHost,

            string ttl)
        {
            Action = action;
            AddressListTimeout = addressListTimeout;
            Bytes = bytes;
            Chain = chain;
            Comment = comment;
            ConnectionBytes = connectionBytes;
            ConnectionLimit = connectionLimit;
            ConnectionMark = connectionMark;
            ConnectionNatState = connectionNatState;
            ConnectionRate = connectionRate;
            ConnectionState = connectionState;
            ConnectionType = connectionType;
            Content = content;
            Disabled = disabled;
            Dscp = dscp;
            DstAddress = dstAddress;
            DstAddressList = dstAddressList;
            DstAddressType = dstAddressType;
            DstLimit = dstLimit;
            DstPort = dstPort;
            Dynamic = dynamic;
            Filter = filter;
            Fragment = fragment;
            Hotspot = hotspot;
            HwOffload = hwOffload;
            IcmpOptions = icmpOptions;
            Id = id;
            InBridgePort = inBridgePort;
            InBridgePortList = inBridgePortList;
            InInterface = inInterface;
            InInterfaceList = inInterfaceList;
            IngressPriority = ingressPriority;
            Invalid = invalid;
            IpsecPolicy = ipsecPolicy;
            Ipv4Options = ipv4Options;
            JumpTarget = jumpTarget;
            Layer7Protocol = layer7Protocol;
            Limit = limit;
            Log = log;
            LogPrefix = logPrefix;
            Nth = nth;
            OutBridgePort = outBridgePort;
            OutBridgePortList = outBridgePortList;
            OutInterface = outInterface;
            OutInterfaceList = outInterfaceList;
            PacketMark = packetMark;
            PacketSize = packetSize;
            Packets = packets;
            PerConnectionClassifier = perConnectionClassifier;
            Port = port;
            Priority = priority;
            Protocol = protocol;
            Psd = psd;
            Random = random;
            RejectWith = rejectWith;
            RoutingMark = routingMark;
            RoutingTable = routingTable;
            SrcAddress = srcAddress;
            SrcAddressList = srcAddressList;
            SrcAddressType = srcAddressType;
            SrcMacAddress = srcMacAddress;
            SrcPort = srcPort;
            TcpFlags = tcpFlags;
            TcpMss = tcpMss;
            Time = time;
            TlsHost = tlsHost;
            Ttl = ttl;
        }
    }
}