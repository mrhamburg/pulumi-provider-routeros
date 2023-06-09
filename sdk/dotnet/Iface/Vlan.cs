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
    /// ## # routeros.Iface.Vlan (Resource)
    /// 
    /// ***
    /// 
    /// #### This is an alias for backwards compatibility between plugin versions.
    /// Please see documentation for routeros.Iface.InterfaceVlan
    /// </summary>
    [RouterosResourceType("routeros:Iface/vlan:Vlan")]
    public partial class Vlan : global::Pulumi.CustomResource
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
        /// ARP resolution protocol mode.
        /// </summary>
        [Output("arp")]
        public Output<string?> Arp { get; private set; } = null!;

        /// <summary>
        /// ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to
        /// the value of arp-timeout in IP/Settings, default is 30s. Can use postfix ms, s, M, h, d for milliseconds, seconds,
        /// minutes, hours or days. If no postfix is set then seconds (s) is used.
        /// </summary>
        [Output("arpTimeout")]
        public Output<string?> ArpTimeout { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Layer2 Maximum transmission unit.
        /// </summary>
        [Output("l2mtu")]
        public Output<int> L2mtu { get; private set; } = null!;

        [Output("loopProtect")]
        public Output<string?> LoopProtect { get; private set; } = null!;

        [Output("loopProtectDisableTime")]
        public Output<string?> LoopProtectDisableTime { get; private set; } = null!;

        [Output("loopProtectSendInterval")]
        public Output<string?> LoopProtectSendInterval { get; private set; } = null!;

        [Output("loopProtectStatus")]
        public Output<string> LoopProtectStatus { get; private set; } = null!;

        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        /// <summary>
        /// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
        /// </summary>
        [Output("mtu")]
        public Output<string> Mtu { get; private set; } = null!;

        /// <summary>
        /// Changing the name of this resource will force it to be recreated. &gt; The links of other configuration properties to this
        /// resource may be lost! &gt; Changing the name of the resource outside of a Terraform will result in a loss of control
        /// integrity for that resource!
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("running")]
        public Output<bool> Running { get; private set; } = null!;

        [Output("useServiceTag")]
        public Output<bool?> UseServiceTag { get; private set; } = null!;

        [Output("vlanId")]
        public Output<int> VlanId { get; private set; } = null!;


        /// <summary>
        /// Create a Vlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vlan(string name, VlanArgs args, CustomResourceOptions? options = null)
            : base("routeros:Iface/vlan:Vlan", name, args ?? new VlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vlan(string name, Input<string> id, VlanState? state = null, CustomResourceOptions? options = null)
            : base("routeros:Iface/vlan:Vlan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vlan Get(string name, Input<string> id, VlanState? state = null, CustomResourceOptions? options = null)
        {
            return new Vlan(name, id, state, options);
        }
    }

    public sealed class VlanArgs : global::Pulumi.ResourceArgs
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
        /// ARP resolution protocol mode.
        /// </summary>
        [Input("arp")]
        public Input<string>? Arp { get; set; }

        /// <summary>
        /// ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to
        /// the value of arp-timeout in IP/Settings, default is 30s. Can use postfix ms, s, M, h, d for milliseconds, seconds,
        /// minutes, hours or days. If no postfix is set then seconds (s) is used.
        /// </summary>
        [Input("arpTimeout")]
        public Input<string>? ArpTimeout { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        [Input("loopProtect")]
        public Input<string>? LoopProtect { get; set; }

        [Input("loopProtectDisableTime")]
        public Input<string>? LoopProtectDisableTime { get; set; }

        [Input("loopProtectSendInterval")]
        public Input<string>? LoopProtectSendInterval { get; set; }

        /// <summary>
        /// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
        /// </summary>
        [Input("mtu")]
        public Input<string>? Mtu { get; set; }

        /// <summary>
        /// Changing the name of this resource will force it to be recreated. &gt; The links of other configuration properties to this
        /// resource may be lost! &gt; Changing the name of the resource outside of a Terraform will result in a loss of control
        /// integrity for that resource!
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("useServiceTag")]
        public Input<bool>? UseServiceTag { get; set; }

        [Input("vlanId", required: true)]
        public Input<int> VlanId { get; set; } = null!;

        public VlanArgs()
        {
        }
        public static new VlanArgs Empty => new VlanArgs();
    }

    public sealed class VlanState : global::Pulumi.ResourceArgs
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
        /// ARP resolution protocol mode.
        /// </summary>
        [Input("arp")]
        public Input<string>? Arp { get; set; }

        /// <summary>
        /// ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to
        /// the value of arp-timeout in IP/Settings, default is 30s. Can use postfix ms, s, M, h, d for milliseconds, seconds,
        /// minutes, hours or days. If no postfix is set then seconds (s) is used.
        /// </summary>
        [Input("arpTimeout")]
        public Input<string>? ArpTimeout { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Layer2 Maximum transmission unit.
        /// </summary>
        [Input("l2mtu")]
        public Input<int>? L2mtu { get; set; }

        [Input("loopProtect")]
        public Input<string>? LoopProtect { get; set; }

        [Input("loopProtectDisableTime")]
        public Input<string>? LoopProtectDisableTime { get; set; }

        [Input("loopProtectSendInterval")]
        public Input<string>? LoopProtectSendInterval { get; set; }

        [Input("loopProtectStatus")]
        public Input<string>? LoopProtectStatus { get; set; }

        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
        /// </summary>
        [Input("mtu")]
        public Input<string>? Mtu { get; set; }

        /// <summary>
        /// Changing the name of this resource will force it to be recreated. &gt; The links of other configuration properties to this
        /// resource may be lost! &gt; Changing the name of the resource outside of a Terraform will result in a loss of control
        /// integrity for that resource!
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("running")]
        public Input<bool>? Running { get; set; }

        [Input("useServiceTag")]
        public Input<bool>? UseServiceTag { get; set; }

        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public VlanState()
        {
        }
        public static new VlanState Empty => new VlanState();
    }
}
