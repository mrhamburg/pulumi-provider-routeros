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
    /// ## # routeros.Iface.Vrrp (Resource)
    /// 
    /// ***
    /// 
    /// #### This is an alias for backwards compatibility between plugin versions.
    /// Please see documentation for routeros.Iface.InterfaceVrrp
    /// </summary>
    [RouterosResourceType("routeros:Iface/vrrp:Vrrp")]
    public partial class Vrrp : global::Pulumi.CustomResource
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

        /// <summary>
        /// Authentication method to use for VRRP advertisement packets.
        /// </summary>
        [Output("authentication")]
        public Output<string?> Authentication { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Allows combining multiple VRRP interfaces to maintain the same VRRP status within the group.
        /// </summary>
        [Output("groupMaster")]
        public Output<string> GroupMaster { get; private set; } = null!;

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// VRRP update interval in seconds. Defines how often master sends advertisement packets.
        /// </summary>
        [Output("interval")]
        public Output<string?> Interval { get; private set; } = null!;

        [Output("invalid")]
        public Output<bool> Invalid { get; private set; } = null!;

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

        /// <summary>
        /// Script to execute when the node is switched to the backup state.
        /// </summary>
        [Output("onBackup")]
        public Output<string?> OnBackup { get; private set; } = null!;

        /// <summary>
        /// Script to execute when the node fails.
        /// </summary>
        [Output("onFail")]
        public Output<string?> OnFail { get; private set; } = null!;

        /// <summary>
        /// Script to execute when the node is switched to master state.
        /// </summary>
        [Output("onMaster")]
        public Output<string?> OnMaster { get; private set; } = null!;

        /// <summary>
        /// Password required for authentication. Can be ignored if authentication is not used.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Whether the master node always has the priority. When set to 'no' the backup node will not be elected to be a master
        /// until the current master fails, even if the backup node has higher priority than the current master. This setting is
        /// ignored if the owner router becomes available
        /// </summary>
        [Output("preemptionMode")]
        public Output<bool?> PreemptionMode { get; private set; } = null!;

        /// <summary>
        /// Priority of VRRP node used in Master election algorithm. A higher number means higher priority. '255' is reserved for
        /// the router that owns VR IP and '0' is reserved for the Master router to indicate that it is releasing responsibility.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// Specifies the remote address of the other VRRP router for syncing connection tracking. If not set, the system
        /// autodetects the remote address via VRRP. The remote address is used only if sync-connection-tracking=yes.Sync connection
        /// tracking uses UDP port 8275.
        /// </summary>
        [Output("remoteAddress")]
        public Output<string?> RemoteAddress { get; private set; } = null!;

        [Output("running")]
        public Output<bool> Running { get; private set; } = null!;

        /// <summary>
        /// Synchronize connection tracking entries from Master to Backup device.
        /// </summary>
        [Output("syncConnectionTracking")]
        public Output<bool?> SyncConnectionTracking { get; private set; } = null!;

        /// <summary>
        /// A protocol that will be used by VRRPv3. Valid only if the version is 3.
        /// </summary>
        [Output("v3Protocol")]
        public Output<string?> V3Protocol { get; private set; } = null!;

        /// <summary>
        /// Which VRRP version to use.
        /// </summary>
        [Output("version")]
        public Output<int?> Version { get; private set; } = null!;

        /// <summary>
        /// Virtual Router identifier. Each Virtual router must have a unique id number.
        /// </summary>
        [Output("vrid")]
        public Output<int?> Vrid { get; private set; } = null!;


        /// <summary>
        /// Create a Vrrp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vrrp(string name, VrrpArgs args, CustomResourceOptions? options = null)
            : base("routeros:Iface/vrrp:Vrrp", name, args ?? new VrrpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vrrp(string name, Input<string> id, VrrpState? state = null, CustomResourceOptions? options = null)
            : base("routeros:Iface/vrrp:Vrrp", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Vrrp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vrrp Get(string name, Input<string> id, VrrpState? state = null, CustomResourceOptions? options = null)
        {
            return new Vrrp(name, id, state, options);
        }
    }

    public sealed class VrrpArgs : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Authentication method to use for VRRP advertisement packets.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Allows combining multiple VRRP interfaces to maintain the same VRRP status within the group.
        /// </summary>
        [Input("groupMaster")]
        public Input<string>? GroupMaster { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// VRRP update interval in seconds. Defines how often master sends advertisement packets.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

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

        /// <summary>
        /// Script to execute when the node is switched to the backup state.
        /// </summary>
        [Input("onBackup")]
        public Input<string>? OnBackup { get; set; }

        /// <summary>
        /// Script to execute when the node fails.
        /// </summary>
        [Input("onFail")]
        public Input<string>? OnFail { get; set; }

        /// <summary>
        /// Script to execute when the node is switched to master state.
        /// </summary>
        [Input("onMaster")]
        public Input<string>? OnMaster { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password required for authentication. Can be ignored if authentication is not used.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Whether the master node always has the priority. When set to 'no' the backup node will not be elected to be a master
        /// until the current master fails, even if the backup node has higher priority than the current master. This setting is
        /// ignored if the owner router becomes available
        /// </summary>
        [Input("preemptionMode")]
        public Input<bool>? PreemptionMode { get; set; }

        /// <summary>
        /// Priority of VRRP node used in Master election algorithm. A higher number means higher priority. '255' is reserved for
        /// the router that owns VR IP and '0' is reserved for the Master router to indicate that it is releasing responsibility.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Specifies the remote address of the other VRRP router for syncing connection tracking. If not set, the system
        /// autodetects the remote address via VRRP. The remote address is used only if sync-connection-tracking=yes.Sync connection
        /// tracking uses UDP port 8275.
        /// </summary>
        [Input("remoteAddress")]
        public Input<string>? RemoteAddress { get; set; }

        /// <summary>
        /// Synchronize connection tracking entries from Master to Backup device.
        /// </summary>
        [Input("syncConnectionTracking")]
        public Input<bool>? SyncConnectionTracking { get; set; }

        /// <summary>
        /// A protocol that will be used by VRRPv3. Valid only if the version is 3.
        /// </summary>
        [Input("v3Protocol")]
        public Input<string>? V3Protocol { get; set; }

        /// <summary>
        /// Which VRRP version to use.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        /// <summary>
        /// Virtual Router identifier. Each Virtual router must have a unique id number.
        /// </summary>
        [Input("vrid")]
        public Input<int>? Vrid { get; set; }

        public VrrpArgs()
        {
        }
        public static new VrrpArgs Empty => new VrrpArgs();
    }

    public sealed class VrrpState : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Authentication method to use for VRRP advertisement packets.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Allows combining multiple VRRP interfaces to maintain the same VRRP status within the group.
        /// </summary>
        [Input("groupMaster")]
        public Input<string>? GroupMaster { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// VRRP update interval in seconds. Defines how often master sends advertisement packets.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("invalid")]
        public Input<bool>? Invalid { get; set; }

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

        /// <summary>
        /// Script to execute when the node is switched to the backup state.
        /// </summary>
        [Input("onBackup")]
        public Input<string>? OnBackup { get; set; }

        /// <summary>
        /// Script to execute when the node fails.
        /// </summary>
        [Input("onFail")]
        public Input<string>? OnFail { get; set; }

        /// <summary>
        /// Script to execute when the node is switched to master state.
        /// </summary>
        [Input("onMaster")]
        public Input<string>? OnMaster { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password required for authentication. Can be ignored if authentication is not used.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Whether the master node always has the priority. When set to 'no' the backup node will not be elected to be a master
        /// until the current master fails, even if the backup node has higher priority than the current master. This setting is
        /// ignored if the owner router becomes available
        /// </summary>
        [Input("preemptionMode")]
        public Input<bool>? PreemptionMode { get; set; }

        /// <summary>
        /// Priority of VRRP node used in Master election algorithm. A higher number means higher priority. '255' is reserved for
        /// the router that owns VR IP and '0' is reserved for the Master router to indicate that it is releasing responsibility.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Specifies the remote address of the other VRRP router for syncing connection tracking. If not set, the system
        /// autodetects the remote address via VRRP. The remote address is used only if sync-connection-tracking=yes.Sync connection
        /// tracking uses UDP port 8275.
        /// </summary>
        [Input("remoteAddress")]
        public Input<string>? RemoteAddress { get; set; }

        [Input("running")]
        public Input<bool>? Running { get; set; }

        /// <summary>
        /// Synchronize connection tracking entries from Master to Backup device.
        /// </summary>
        [Input("syncConnectionTracking")]
        public Input<bool>? SyncConnectionTracking { get; set; }

        /// <summary>
        /// A protocol that will be used by VRRPv3. Valid only if the version is 3.
        /// </summary>
        [Input("v3Protocol")]
        public Input<string>? V3Protocol { get; set; }

        /// <summary>
        /// Which VRRP version to use.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        /// <summary>
        /// Virtual Router identifier. Each Virtual router must have a unique id number.
        /// </summary>
        [Input("vrid")]
        public Input<int>? Vrid { get; set; }

        public VrrpState()
        {
        }
        public static new VrrpState Empty => new VrrpState();
    }
}
