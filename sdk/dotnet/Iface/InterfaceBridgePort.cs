// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.Iface
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Routeros = Pulumi.Routeros;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bridgePort = new Routeros.Iface.InterfaceBridgePort("bridgePort", new()
    ///     {
    ///         Bridge = "bridge",
    ///         Interface = "ether5",
    ///         Pvid = 50,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/interface/bridge/port get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:Iface/interfaceBridgePort:InterfaceBridgePort bridge_port "*0"
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:Iface/interfaceBridgePort:InterfaceBridgePort")]
    public partial class InterfaceBridgePort : global::Pulumi.CustomResource
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
        /// When enabled, prevents a port moving from discarding into forwarding state if no BPDUs are received from the neighboring bridge. The port will change into a forwarding state only when a BPDU is received. This property only has an effect when protocol-mode is set to rstp or mstp and edge is set to no.
        /// </summary>
        [Output("autoIsolate")]
        public Output<bool?> AutoIsolate { get; private set; } = null!;

        /// <summary>
        /// This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Output("bpduGuard")]
        public Output<bool?> BpduGuard { get; private set; } = null!;

        [Output("bridge")]
        public Output<string> Bridge { get; private set; } = null!;

        /// <summary>
        /// When enabled, bridge floods broadcast traffic to all bridge egress ports. When disabled, drops broadcast traffic on egress ports.
        /// </summary>
        [Output("broadcastFlood")]
        public Output<bool?> BroadcastFlood { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("debugInfo")]
        public Output<string> DebugInfo { get; private set; } = null!;

        /// <summary>
        /// Root bridge ID (bridge priority and the bridge MAC address).
        /// </summary>
        [Output("designatedBridge")]
        public Output<string> DesignatedBridge { get; private set; } = null!;

        /// <summary>
        /// Designated cost.
        /// </summary>
        [Output("designatedCost")]
        public Output<string> DesignatedCost { get; private set; } = null!;

        /// <summary>
        /// Designated port number.
        /// </summary>
        [Output("designatedPortNumber")]
        public Output<int> DesignatedPortNumber { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
        /// </summary>
        [Output("dynamic")]
        public Output<bool> Dynamic { get; private set; } = null!;

        /// <summary>
        /// Set port as edge port or non-edge port, or enable edge discovery. Edge ports are connected to a LAN that has no other bridges attached.
        /// </summary>
        [Output("edge")]
        public Output<string?> Edge { get; private set; } = null!;

        /// <summary>
        /// Whether port is an edge port or not.
        /// </summary>
        [Output("edgePort")]
        public Output<bool> EdgePort { get; private set; } = null!;

        /// <summary>
        /// Whether port is set to automatically detect edge ports.
        /// </summary>
        [Output("edgePortDiscovery")]
        public Output<bool> EdgePortDiscovery { get; private set; } = null!;

        /// <summary>
        /// Whether registration table is used instead of forwarding data base.
        /// </summary>
        [Output("externalFdbStatus")]
        public Output<bool> ExternalFdbStatus { get; private set; } = null!;

        /// <summary>
        /// Enables IGMP Fast leave feature on the port.
        /// </summary>
        [Output("fastLeave")]
        public Output<bool?> FastLeave { get; private set; } = null!;

        /// <summary>
        /// Shows if the port is not blocked by (R/M)STP.
        /// </summary>
        [Output("forwarding")]
        public Output<bool> Forwarding { get; private set; } = null!;

        /// <summary>
        /// Specifies allowed ingress frame types on a bridge port. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Output("frameTypes")]
        public Output<string?> FrameTypes { get; private set; } = null!;

        /// <summary>
        /// Use split horizon bridging to prevent bridging loops. Set the same value for group of ports, to prevent them from sending data to ports with the same horizon value. Split horizon is a software feature that disables hardware offloading. This value is integer '0'..'429496729' or 'none'.
        /// </summary>
        [Output("horizon")]
        public Output<string?> Horizon { get; private set; } = null!;

        /// <summary>
        /// Enable or disable Hardware Offloading of the interface.
        /// </summary>
        [Output("hw")]
        public Output<bool> Hw { get; private set; } = null!;

        /// <summary>
        /// Hardware offloading state.
        /// </summary>
        [Output("hwOffload")]
        public Output<bool> HwOffload { get; private set; } = null!;

        /// <summary>
        /// Switch chip used by the port.
        /// </summary>
        [Output("hwOffloadGroup")]
        public Output<string> HwOffloadGroup { get; private set; } = null!;

        [Output("inactive")]
        public Output<bool> Inactive { get; private set; } = null!;

        /// <summary>
        /// Enables or disables VLAN ingress filtering, which checks if the ingress port is a member of the received VLAN ID in the bridge VLAN table. Should be used with frame-types to specify if the ingress traffic should be tagged or untagged. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Output("ingressFiltering")]
        public Output<bool> IngressFiltering { get; private set; } = null!;

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Path cost to the interface for MSTI0 inside a region. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Output("internalPathCost")]
        public Output<int?> InternalPathCost { get; private set; } = null!;

        /// <summary>
        /// Changes MAC learning behaviour on a bridge port
        /// </summary>
        [Output("learn")]
        public Output<string?> Learn { get; private set; } = null!;

        /// <summary>
        /// Shows whether the port is capable of learning MAC addresses.
        /// </summary>
        [Output("learning")]
        public Output<bool> Learning { get; private set; } = null!;

        /// <summary>
        /// Changes the state of a bridge port whether IGMP membership reports are going to be forwarded to this port.
        /// </summary>
        [Output("multicastRouter")]
        public Output<string?> MulticastRouter { get; private set; } = null!;

        [Output("nextid")]
        public Output<string> Nextid { get; private set; } = null!;

        /// <summary>
        /// Path cost to the interface, used by STP to determine the "best" path, used by MSTP todetermine "best" path between regions. This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Output("pathCost")]
        public Output<string?> PathCost { get; private set; } = null!;

        /// <summary>
        /// Specifies if a bridge port is connected to a bridge using a point-to-point link for faster convergence in case of failure. This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Output("pointToPoint")]
        public Output<string?> PointToPoint { get; private set; } = null!;

        /// <summary>
        /// Whether the port is connected to a bridge port using full-duplex (true) or half-duplex (false).
        /// </summary>
        [Output("pointToPointPort")]
        public Output<bool> PointToPointPort { get; private set; } = null!;

        /// <summary>
        /// Port number will be assigned in the order that ports got added to the bridge, but this is only true until reboot. After reboot internal numbering will be used.
        /// </summary>
        [Output("portNumber")]
        public Output<int> PortNumber { get; private set; } = null!;

        /// <summary>
        /// The priority of the interface, used by STP to determine the root port, used by MSTP to determine root port between regions.
        /// </summary>
        [Output("priority")]
        public Output<string?> Priority { get; private set; } = null!;

        /// <summary>
        /// ort VLAN ID (pvid) specifies which VLAN the untagged ingress traffic is assigned to. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Output("pvid")]
        public Output<int> Pvid { get; private set; } = null!;

        /// <summary>
        /// Enable the restricted role on a port, used by STP to forbid a port becoming a root port. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Output("restrictedRole")]
        public Output<bool?> RestrictedRole { get; private set; } = null!;

        /// <summary>
        /// Disable topology change notification (TCN) sending on a port, used by STP to forbid network topology changes to propagate. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Output("restrictedTcn")]
        public Output<bool?> RestrictedTcn { get; private set; } = null!;

        /// <summary>
        /// (R/M)STP algorithm assigned role of the port
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The total cost of the path to the root-bridge.
        /// </summary>
        [Output("rootPathCost")]
        public Output<int> RootPathCost { get; private set; } = null!;

        /// <summary>
        /// Whether the port is sending RSTP or MSTP BPDU types. A port will transit to STP type when RSTP/MSTP enabled port receives a STP BPDU
        /// </summary>
        [Output("sendingRstp")]
        public Output<string> SendingRstp { get; private set; } = null!;

        /// <summary>
        /// Port status ('in-bridge' - port is enabled).
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Forces all packets to be treated as untagged packets. Packets on ingress port will be tagged with another VLAN tag regardless if a VLAN tag already exists, packets will be tagged with a VLAN ID that matches the pvid value and will use EtherType that is specified in ether-type. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Output("tagStacking")]
        public Output<bool?> TagStacking { get; private set; } = null!;

        /// <summary>
        /// When enabled, it allows to forward DHCP packets towards DHCP server through this port. Mainly used to limit unauthorized servers to provide malicious information for users. This property only has effect when dhcp-snooping is set to yes.
        /// </summary>
        [Output("trusted")]
        public Output<bool?> Trusted { get; private set; } = null!;

        /// <summary>
        /// When enabled, bridge floods unknown multicast traffic to all bridge egress ports.
        /// </summary>
        [Output("unknownMulticastFlood")]
        public Output<bool?> UnknownMulticastFlood { get; private set; } = null!;

        /// <summary>
        /// When enabled, bridge floods unknown unicast traffic to all bridge egress ports.
        /// </summary>
        [Output("unknownUnicastFlood")]
        public Output<bool?> UnknownUnicastFlood { get; private set; } = null!;


        /// <summary>
        /// Create a InterfaceBridgePort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterfaceBridgePort(string name, InterfaceBridgePortArgs args, CustomResourceOptions? options = null)
            : base("routeros:Iface/interfaceBridgePort:InterfaceBridgePort", name, args ?? new InterfaceBridgePortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InterfaceBridgePort(string name, Input<string> id, InterfaceBridgePortState? state = null, CustomResourceOptions? options = null)
            : base("routeros:Iface/interfaceBridgePort:InterfaceBridgePort", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InterfaceBridgePort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterfaceBridgePort Get(string name, Input<string> id, InterfaceBridgePortState? state = null, CustomResourceOptions? options = null)
        {
            return new InterfaceBridgePort(name, id, state, options);
        }
    }

    public sealed class InterfaceBridgePortArgs : global::Pulumi.ResourceArgs
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
        /// When enabled, prevents a port moving from discarding into forwarding state if no BPDUs are received from the neighboring bridge. The port will change into a forwarding state only when a BPDU is received. This property only has an effect when protocol-mode is set to rstp or mstp and edge is set to no.
        /// </summary>
        [Input("autoIsolate")]
        public Input<bool>? AutoIsolate { get; set; }

        /// <summary>
        /// This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Input("bpduGuard")]
        public Input<bool>? BpduGuard { get; set; }

        [Input("bridge", required: true)]
        public Input<string> Bridge { get; set; } = null!;

        /// <summary>
        /// When enabled, bridge floods broadcast traffic to all bridge egress ports. When disabled, drops broadcast traffic on egress ports.
        /// </summary>
        [Input("broadcastFlood")]
        public Input<bool>? BroadcastFlood { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Set port as edge port or non-edge port, or enable edge discovery. Edge ports are connected to a LAN that has no other bridges attached.
        /// </summary>
        [Input("edge")]
        public Input<string>? Edge { get; set; }

        /// <summary>
        /// Enables IGMP Fast leave feature on the port.
        /// </summary>
        [Input("fastLeave")]
        public Input<bool>? FastLeave { get; set; }

        /// <summary>
        /// Specifies allowed ingress frame types on a bridge port. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("frameTypes")]
        public Input<string>? FrameTypes { get; set; }

        /// <summary>
        /// Use split horizon bridging to prevent bridging loops. Set the same value for group of ports, to prevent them from sending data to ports with the same horizon value. Split horizon is a software feature that disables hardware offloading. This value is integer '0'..'429496729' or 'none'.
        /// </summary>
        [Input("horizon")]
        public Input<string>? Horizon { get; set; }

        /// <summary>
        /// Enable or disable Hardware Offloading of the interface.
        /// </summary>
        [Input("hw")]
        public Input<bool>? Hw { get; set; }

        /// <summary>
        /// Enables or disables VLAN ingress filtering, which checks if the ingress port is a member of the received VLAN ID in the bridge VLAN table. Should be used with frame-types to specify if the ingress traffic should be tagged or untagged. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("ingressFiltering")]
        public Input<bool>? IngressFiltering { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// Path cost to the interface for MSTI0 inside a region. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Input("internalPathCost")]
        public Input<int>? InternalPathCost { get; set; }

        /// <summary>
        /// Changes MAC learning behaviour on a bridge port
        /// </summary>
        [Input("learn")]
        public Input<string>? Learn { get; set; }

        /// <summary>
        /// Changes the state of a bridge port whether IGMP membership reports are going to be forwarded to this port.
        /// </summary>
        [Input("multicastRouter")]
        public Input<string>? MulticastRouter { get; set; }

        /// <summary>
        /// Path cost to the interface, used by STP to determine the "best" path, used by MSTP todetermine "best" path between regions. This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Input("pathCost")]
        public Input<string>? PathCost { get; set; }

        /// <summary>
        /// Specifies if a bridge port is connected to a bridge using a point-to-point link for faster convergence in case of failure. This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Input("pointToPoint")]
        public Input<string>? PointToPoint { get; set; }

        /// <summary>
        /// The priority of the interface, used by STP to determine the root port, used by MSTP to determine root port between regions.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// ort VLAN ID (pvid) specifies which VLAN the untagged ingress traffic is assigned to. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("pvid", required: true)]
        public Input<int> Pvid { get; set; } = null!;

        /// <summary>
        /// Enable the restricted role on a port, used by STP to forbid a port becoming a root port. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Input("restrictedRole")]
        public Input<bool>? RestrictedRole { get; set; }

        /// <summary>
        /// Disable topology change notification (TCN) sending on a port, used by STP to forbid network topology changes to propagate. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Input("restrictedTcn")]
        public Input<bool>? RestrictedTcn { get; set; }

        /// <summary>
        /// Forces all packets to be treated as untagged packets. Packets on ingress port will be tagged with another VLAN tag regardless if a VLAN tag already exists, packets will be tagged with a VLAN ID that matches the pvid value and will use EtherType that is specified in ether-type. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("tagStacking")]
        public Input<bool>? TagStacking { get; set; }

        /// <summary>
        /// When enabled, it allows to forward DHCP packets towards DHCP server through this port. Mainly used to limit unauthorized servers to provide malicious information for users. This property only has effect when dhcp-snooping is set to yes.
        /// </summary>
        [Input("trusted")]
        public Input<bool>? Trusted { get; set; }

        /// <summary>
        /// When enabled, bridge floods unknown multicast traffic to all bridge egress ports.
        /// </summary>
        [Input("unknownMulticastFlood")]
        public Input<bool>? UnknownMulticastFlood { get; set; }

        /// <summary>
        /// When enabled, bridge floods unknown unicast traffic to all bridge egress ports.
        /// </summary>
        [Input("unknownUnicastFlood")]
        public Input<bool>? UnknownUnicastFlood { get; set; }

        public InterfaceBridgePortArgs()
        {
        }
        public static new InterfaceBridgePortArgs Empty => new InterfaceBridgePortArgs();
    }

    public sealed class InterfaceBridgePortState : global::Pulumi.ResourceArgs
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
        /// When enabled, prevents a port moving from discarding into forwarding state if no BPDUs are received from the neighboring bridge. The port will change into a forwarding state only when a BPDU is received. This property only has an effect when protocol-mode is set to rstp or mstp and edge is set to no.
        /// </summary>
        [Input("autoIsolate")]
        public Input<bool>? AutoIsolate { get; set; }

        /// <summary>
        /// This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Input("bpduGuard")]
        public Input<bool>? BpduGuard { get; set; }

        [Input("bridge")]
        public Input<string>? Bridge { get; set; }

        /// <summary>
        /// When enabled, bridge floods broadcast traffic to all bridge egress ports. When disabled, drops broadcast traffic on egress ports.
        /// </summary>
        [Input("broadcastFlood")]
        public Input<bool>? BroadcastFlood { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("debugInfo")]
        public Input<string>? DebugInfo { get; set; }

        /// <summary>
        /// Root bridge ID (bridge priority and the bridge MAC address).
        /// </summary>
        [Input("designatedBridge")]
        public Input<string>? DesignatedBridge { get; set; }

        /// <summary>
        /// Designated cost.
        /// </summary>
        [Input("designatedCost")]
        public Input<string>? DesignatedCost { get; set; }

        /// <summary>
        /// Designated port number.
        /// </summary>
        [Input("designatedPortNumber")]
        public Input<int>? DesignatedPortNumber { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
        /// </summary>
        [Input("dynamic")]
        public Input<bool>? Dynamic { get; set; }

        /// <summary>
        /// Set port as edge port or non-edge port, or enable edge discovery. Edge ports are connected to a LAN that has no other bridges attached.
        /// </summary>
        [Input("edge")]
        public Input<string>? Edge { get; set; }

        /// <summary>
        /// Whether port is an edge port or not.
        /// </summary>
        [Input("edgePort")]
        public Input<bool>? EdgePort { get; set; }

        /// <summary>
        /// Whether port is set to automatically detect edge ports.
        /// </summary>
        [Input("edgePortDiscovery")]
        public Input<bool>? EdgePortDiscovery { get; set; }

        /// <summary>
        /// Whether registration table is used instead of forwarding data base.
        /// </summary>
        [Input("externalFdbStatus")]
        public Input<bool>? ExternalFdbStatus { get; set; }

        /// <summary>
        /// Enables IGMP Fast leave feature on the port.
        /// </summary>
        [Input("fastLeave")]
        public Input<bool>? FastLeave { get; set; }

        /// <summary>
        /// Shows if the port is not blocked by (R/M)STP.
        /// </summary>
        [Input("forwarding")]
        public Input<bool>? Forwarding { get; set; }

        /// <summary>
        /// Specifies allowed ingress frame types on a bridge port. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("frameTypes")]
        public Input<string>? FrameTypes { get; set; }

        /// <summary>
        /// Use split horizon bridging to prevent bridging loops. Set the same value for group of ports, to prevent them from sending data to ports with the same horizon value. Split horizon is a software feature that disables hardware offloading. This value is integer '0'..'429496729' or 'none'.
        /// </summary>
        [Input("horizon")]
        public Input<string>? Horizon { get; set; }

        /// <summary>
        /// Enable or disable Hardware Offloading of the interface.
        /// </summary>
        [Input("hw")]
        public Input<bool>? Hw { get; set; }

        /// <summary>
        /// Hardware offloading state.
        /// </summary>
        [Input("hwOffload")]
        public Input<bool>? HwOffload { get; set; }

        /// <summary>
        /// Switch chip used by the port.
        /// </summary>
        [Input("hwOffloadGroup")]
        public Input<string>? HwOffloadGroup { get; set; }

        [Input("inactive")]
        public Input<bool>? Inactive { get; set; }

        /// <summary>
        /// Enables or disables VLAN ingress filtering, which checks if the ingress port is a member of the received VLAN ID in the bridge VLAN table. Should be used with frame-types to specify if the ingress traffic should be tagged or untagged. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("ingressFiltering")]
        public Input<bool>? IngressFiltering { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Path cost to the interface for MSTI0 inside a region. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Input("internalPathCost")]
        public Input<int>? InternalPathCost { get; set; }

        /// <summary>
        /// Changes MAC learning behaviour on a bridge port
        /// </summary>
        [Input("learn")]
        public Input<string>? Learn { get; set; }

        /// <summary>
        /// Shows whether the port is capable of learning MAC addresses.
        /// </summary>
        [Input("learning")]
        public Input<bool>? Learning { get; set; }

        /// <summary>
        /// Changes the state of a bridge port whether IGMP membership reports are going to be forwarded to this port.
        /// </summary>
        [Input("multicastRouter")]
        public Input<string>? MulticastRouter { get; set; }

        [Input("nextid")]
        public Input<string>? Nextid { get; set; }

        /// <summary>
        /// Path cost to the interface, used by STP to determine the "best" path, used by MSTP todetermine "best" path between regions. This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Input("pathCost")]
        public Input<string>? PathCost { get; set; }

        /// <summary>
        /// Specifies if a bridge port is connected to a bridge using a point-to-point link for faster convergence in case of failure. This property has no effect when protocol-mode is set to none.
        /// </summary>
        [Input("pointToPoint")]
        public Input<string>? PointToPoint { get; set; }

        /// <summary>
        /// Whether the port is connected to a bridge port using full-duplex (true) or half-duplex (false).
        /// </summary>
        [Input("pointToPointPort")]
        public Input<bool>? PointToPointPort { get; set; }

        /// <summary>
        /// Port number will be assigned in the order that ports got added to the bridge, but this is only true until reboot. After reboot internal numbering will be used.
        /// </summary>
        [Input("portNumber")]
        public Input<int>? PortNumber { get; set; }

        /// <summary>
        /// The priority of the interface, used by STP to determine the root port, used by MSTP to determine root port between regions.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// ort VLAN ID (pvid) specifies which VLAN the untagged ingress traffic is assigned to. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("pvid")]
        public Input<int>? Pvid { get; set; }

        /// <summary>
        /// Enable the restricted role on a port, used by STP to forbid a port becoming a root port. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Input("restrictedRole")]
        public Input<bool>? RestrictedRole { get; set; }

        /// <summary>
        /// Disable topology change notification (TCN) sending on a port, used by STP to forbid network topology changes to propagate. This property only has effect when protocol-mode is set to mstp.
        /// </summary>
        [Input("restrictedTcn")]
        public Input<bool>? RestrictedTcn { get; set; }

        /// <summary>
        /// (R/M)STP algorithm assigned role of the port
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The total cost of the path to the root-bridge.
        /// </summary>
        [Input("rootPathCost")]
        public Input<int>? RootPathCost { get; set; }

        /// <summary>
        /// Whether the port is sending RSTP or MSTP BPDU types. A port will transit to STP type when RSTP/MSTP enabled port receives a STP BPDU
        /// </summary>
        [Input("sendingRstp")]
        public Input<string>? SendingRstp { get; set; }

        /// <summary>
        /// Port status ('in-bridge' - port is enabled).
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Forces all packets to be treated as untagged packets. Packets on ingress port will be tagged with another VLAN tag regardless if a VLAN tag already exists, packets will be tagged with a VLAN ID that matches the pvid value and will use EtherType that is specified in ether-type. This property only has effect when vlan-filtering is set to yes.
        /// </summary>
        [Input("tagStacking")]
        public Input<bool>? TagStacking { get; set; }

        /// <summary>
        /// When enabled, it allows to forward DHCP packets towards DHCP server through this port. Mainly used to limit unauthorized servers to provide malicious information for users. This property only has effect when dhcp-snooping is set to yes.
        /// </summary>
        [Input("trusted")]
        public Input<bool>? Trusted { get; set; }

        /// <summary>
        /// When enabled, bridge floods unknown multicast traffic to all bridge egress ports.
        /// </summary>
        [Input("unknownMulticastFlood")]
        public Input<bool>? UnknownMulticastFlood { get; set; }

        /// <summary>
        /// When enabled, bridge floods unknown unicast traffic to all bridge egress ports.
        /// </summary>
        [Input("unknownUnicastFlood")]
        public Input<bool>? UnknownUnicastFlood { get; set; }

        public InterfaceBridgePortState()
        {
        }
        public static new InterfaceBridgePortState Empty => new InterfaceBridgePortState();
    }
}
