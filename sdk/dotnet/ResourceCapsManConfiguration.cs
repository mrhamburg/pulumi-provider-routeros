// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros
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
    ///     var testConfiguration = new Routeros.ResourceCapsManConfiguration("testConfiguration", new()
    ///     {
    ///         Comment = "Comment",
    ///         Country = "no_country_set",
    ///         DisconnectTimeout = "1s150ms",
    ///         Distance = "indoors",
    ///         FrameLifetime = "0.12",
    ///         GuardInterval = "long",
    ///         HideSsid = true,
    ///         HwProtectionMode = "rts-cts",
    ///         HwRetries = 1,
    ///         Installation = "indoor",
    ///         KeepaliveFrames = "enabled",
    ///         LoadBalancingGroup = "",
    ///         MaxStaCount = 1,
    ///         Mode = "ap",
    ///         MulticastHelper = "full",
    ///         RxChains = new[]
    ///         {
    ///             1,
    ///             3,
    ///         },
    ///         Ssid = "SSID",
    ///         TxChains = new[]
    ///         {
    ///             0,
    ///             2,
    ///         },
    ///     });
    /// 
    ///     var testChannel = new Routeros.ResourceCapsManChannel("testChannel");
    /// 
    ///     var testDatapath = new Routeros.ResourceCapsManDatapath("testDatapath");
    /// 
    ///     var testRates = new Routeros.ResourceCapsManRates("testRates");
    /// 
    ///     var testSecurity = new Routeros.ResourceCapsManSecurity("testSecurity");
    /// 
    ///     var testConfiguration2 = new Routeros.ResourceCapsManConfiguration("testConfiguration2", new()
    ///     {
    ///         Channel = 
    ///         {
    ///             { "config", testChannel.Name },
    ///             { "band", "2ghz-b/g/n" },
    ///             { "control_channel_width", "10mhz" },
    ///             { "extension_channel", "eCee" },
    ///             { "frequency", "2412" },
    ///             { "reselect_interval", "1h" },
    ///             { "save_selected", "true" },
    ///             { "secondary_frequency", "disabled" },
    ///             { "skip_dfs_channels", "true" },
    ///             { "tx_power", "20" },
    ///         },
    ///         Datapath = 
    ///         {
    ///             { "config", testDatapath.Name },
    ///             { "arp", "local-proxy-arp" },
    ///             { "bridge", "bridge" },
    ///             { "bridge_cost", "100" },
    ///             { "bridge_horizon", "200" },
    ///             { "client_to_client_forwarding", "true" },
    ///             { "interface_list", "static" },
    ///             { "l2mtu", "1450" },
    ///             { "local_forwarding", "true" },
    ///             { "mtu", "1500" },
    ///             { "vlan_id", "101" },
    ///             { "vlan_mode", "no-tag" },
    ///         },
    ///         Rates = 
    ///         {
    ///             { "config", testRates.Name },
    ///             { "basic", "1Mbps,5.5Mbps,6Mbps,18Mbps,36Mbps,54Mbps" },
    ///             { "ht_basic_mcs", "mcs-0,mcs-7,mcs-11,mcs-14,mcs-16,mcs-21" },
    ///             { "ht_supported_mcs", "mcs-3,mcs-8,mcs-10,mcs-13,mcs-17,mcs-18" },
    ///             { "supported", "2Mbps,11Mbps,9Mbps,12Mbps,24Mbps,48Mbps" },
    ///             { "vht_basic_mcs", "none" },
    ///             { "vht_supported_mcs", "mcs0-9,mcs0-7" },
    ///         },
    ///         Security = 
    ///         {
    ///             { "config", testSecurity.Name },
    ///             { "authentication_types", "wpa-psk,wpa-eap" },
    ///             { "disable_pmkid", "true" },
    ///             { "eap_methods", "eap-tls,passthrough" },
    ///             { "eap_radius_accounting", "true" },
    ///             { "encryption", "aes-ccm,tkip" },
    ///             { "group_encryption", "aes-ccm" },
    ///             { "group_key_update", "1h" },
    ///             { "passphrase", "AAAAAAAAA" },
    ///             { "tls_certificate", "none" },
    ///             { "tls_mode", "verify-certificate" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             testChannel,
    ///             testDatapath,
    ///             testRates,
    ///             testSecurity,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import with the name of the CAPsMAN configuration in case of the example use test_configuration_name
    /// 
    /// ```sh
    ///  $ pulumi import routeros:index/resourceCapsManConfiguration:ResourceCapsManConfiguration test_configuration_2 test_configuration_name
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:index/resourceCapsManConfiguration:ResourceCapsManConfiguration")]
    public partial class ResourceCapsManConfiguration : global::Pulumi.CustomResource
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
        /// &lt;em&gt;A set of transformations for field names. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Output("___ts_")]
        public Output<string?> ___ts_ { get; private set; } = null!;

        /// <summary>
        /// Channel inline settings.
        /// </summary>
        [Output("channel")]
        public Output<ImmutableDictionary<string, string>?> Channel { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies default value of scan-list. Value no*country*set is an FCC compliant set of channels.
        /// </summary>
        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        /// <summary>
        /// Datapath inline settings.
        /// </summary>
        [Output("datapath")]
        public Output<ImmutableDictionary<string, string>?> Datapath { get; private set; } = null!;

        /// <summary>
        /// This interval is measured from third sending failure on the lowest data rate. At this point 3 * (hw-retries + 1) frame transmits on the lowest data rate had failed. During disconnect-timeout packet transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully during disconnect-timeout, the connection is closed, and this event is logged as "extensive data loss". Successful frame transmission resets this timer.
        /// </summary>
        [Output("disconnectTimeout")]
        public Output<string?> DisconnectTimeout { get; private set; } = null!;

        /// <summary>
        /// How long to wait for confirmation of unicast frames (ACKs) before considering transmission unsuccessful, or in short ACK-Timeout.
        /// </summary>
        [Output("distance")]
        public Output<string?> Distance { get; private set; } = null!;

        /// <summary>
        /// Discard frames that have been queued for sending longer than frame-lifetime. By default, when value of this property is 0, frames are discarded only after connection is closed (format: 0.00 sec).
        /// </summary>
        [Output("frameLifetime")]
        public Output<string?> FrameLifetime { get; private set; } = null!;

        /// <summary>
        /// Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how this may affect throughput). "any" will use either short or long, depending on data rate, "long" will use long.
        /// </summary>
        [Output("guardInterval")]
        public Output<string?> GuardInterval { get; private set; } = null!;

        /// <summary>
        /// This property has effect only in AP mode. Setting it to yes can remove this network from the list of wireless networks that are shown by some client software. Changing this setting does not improve the security of the wireless network, because SSID is included in other frames sent by the AP.
        /// </summary>
        [Output("hideSsid")]
        public Output<bool> HideSsid { get; private set; } = null!;

        /// <summary>
        /// Frame protection support property. [See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#Frame_protection_support_(RTS/CTS)).
        /// </summary>
        [Output("hwProtectionMode")]
        public Output<string?> HwProtectionMode { get; private set; } = null!;

        /// <summary>
        /// Number of times sending frame is retried without considering it a transmission failure. [See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless)
        /// </summary>
        [Output("hwRetries")]
        public Output<int?> HwRetries { get; private set; } = null!;

        /// <summary>
        /// Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.
        /// </summary>
        [Output("installation")]
        public Output<string?> Installation { get; private set; } = null!;

        /// <summary>
        /// If a client has not communicated for around 20 seconds, AP sends a "keepalive-frame".
        /// </summary>
        [Output("keepaliveFrames")]
        public Output<string?> KeepaliveFrames { get; private set; } = null!;

        /// <summary>
        /// Tags the interface to the load balancing group. For a client to connect to interface in this group, the interface should have the same number of already connected clients as all other interfaces in the group or smaller. Useful in setups where ranges of CAPs mostly overlap.
        /// </summary>
        [Output("loadBalancingGroup")]
        public Output<string?> LoadBalancingGroup { get; private set; } = null!;

        /// <summary>
        /// Maximum number of associated clients.
        /// </summary>
        [Output("maxStaCount")]
        public Output<int?> MaxStaCount { get; private set; } = null!;

        /// <summary>
        /// Set operational mode. Only **ap** currently supported.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// When set to full multicast packets will be sent with unicast destination MAC address, resolving multicast problem on a wireless link. This option should be enabled only on the access point, clients should be configured in station-bridge mode.
        /// </summary>
        [Output("multicastHelper")]
        public Output<string?> MulticastHelper { get; private set; } = null!;

        /// <summary>
        /// Changing the name of this resource will force it to be recreated. &gt; The links of other configuration properties to this
        /// resource may be lost! &gt; Changing the name of the resource outside of a Terraform will result in a loss of control
        /// integrity for that resource!
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Rates inline settings.
        /// </summary>
        [Output("rates")]
        public Output<ImmutableDictionary<string, string>?> Rates { get; private set; } = null!;

        /// <summary>
        /// Which antennas to use for receive.
        /// </summary>
        [Output("rxChains")]
        public Output<ImmutableArray<int>> RxChains { get; private set; } = null!;

        /// <summary>
        /// Security inline settings.
        /// </summary>
        [Output("security")]
        public Output<ImmutableDictionary<string, string>?> Security { get; private set; } = null!;

        /// <summary>
        /// SSID (service set identifier) is a name broadcast in the beacons that identifies wireless network.
        /// </summary>
        [Output("ssid")]
        public Output<string?> Ssid { get; private set; } = null!;

        /// <summary>
        /// Which antennas to use for transmit.
        /// </summary>
        [Output("txChains")]
        public Output<ImmutableArray<int>> TxChains { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceCapsManConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceCapsManConfiguration(string name, ResourceCapsManConfigurationArgs? args = null, CustomResourceOptions? options = null)
            : base("routeros:index/resourceCapsManConfiguration:ResourceCapsManConfiguration", name, args ?? new ResourceCapsManConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceCapsManConfiguration(string name, Input<string> id, ResourceCapsManConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("routeros:index/resourceCapsManConfiguration:ResourceCapsManConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceCapsManConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceCapsManConfiguration Get(string name, Input<string> id, ResourceCapsManConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceCapsManConfiguration(name, id, state, options);
        }
    }

    public sealed class ResourceCapsManConfigurationArgs : global::Pulumi.ResourceArgs
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
        /// &lt;em&gt;A set of transformations for field names. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___ts_")]
        public Input<string>? ___ts_ { get; set; }

        [Input("channel")]
        private InputMap<string>? _channel;

        /// <summary>
        /// Channel inline settings.
        /// </summary>
        public InputMap<string> Channel
        {
            get => _channel ?? (_channel = new InputMap<string>());
            set => _channel = value;
        }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies default value of scan-list. Value no*country*set is an FCC compliant set of channels.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("datapath")]
        private InputMap<string>? _datapath;

        /// <summary>
        /// Datapath inline settings.
        /// </summary>
        public InputMap<string> Datapath
        {
            get => _datapath ?? (_datapath = new InputMap<string>());
            set => _datapath = value;
        }

        /// <summary>
        /// This interval is measured from third sending failure on the lowest data rate. At this point 3 * (hw-retries + 1) frame transmits on the lowest data rate had failed. During disconnect-timeout packet transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully during disconnect-timeout, the connection is closed, and this event is logged as "extensive data loss". Successful frame transmission resets this timer.
        /// </summary>
        [Input("disconnectTimeout")]
        public Input<string>? DisconnectTimeout { get; set; }

        /// <summary>
        /// How long to wait for confirmation of unicast frames (ACKs) before considering transmission unsuccessful, or in short ACK-Timeout.
        /// </summary>
        [Input("distance")]
        public Input<string>? Distance { get; set; }

        /// <summary>
        /// Discard frames that have been queued for sending longer than frame-lifetime. By default, when value of this property is 0, frames are discarded only after connection is closed (format: 0.00 sec).
        /// </summary>
        [Input("frameLifetime")]
        public Input<string>? FrameLifetime { get; set; }

        /// <summary>
        /// Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how this may affect throughput). "any" will use either short or long, depending on data rate, "long" will use long.
        /// </summary>
        [Input("guardInterval")]
        public Input<string>? GuardInterval { get; set; }

        /// <summary>
        /// This property has effect only in AP mode. Setting it to yes can remove this network from the list of wireless networks that are shown by some client software. Changing this setting does not improve the security of the wireless network, because SSID is included in other frames sent by the AP.
        /// </summary>
        [Input("hideSsid")]
        public Input<bool>? HideSsid { get; set; }

        /// <summary>
        /// Frame protection support property. [See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#Frame_protection_support_(RTS/CTS)).
        /// </summary>
        [Input("hwProtectionMode")]
        public Input<string>? HwProtectionMode { get; set; }

        /// <summary>
        /// Number of times sending frame is retried without considering it a transmission failure. [See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless)
        /// </summary>
        [Input("hwRetries")]
        public Input<int>? HwRetries { get; set; }

        /// <summary>
        /// Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.
        /// </summary>
        [Input("installation")]
        public Input<string>? Installation { get; set; }

        /// <summary>
        /// If a client has not communicated for around 20 seconds, AP sends a "keepalive-frame".
        /// </summary>
        [Input("keepaliveFrames")]
        public Input<string>? KeepaliveFrames { get; set; }

        /// <summary>
        /// Tags the interface to the load balancing group. For a client to connect to interface in this group, the interface should have the same number of already connected clients as all other interfaces in the group or smaller. Useful in setups where ranges of CAPs mostly overlap.
        /// </summary>
        [Input("loadBalancingGroup")]
        public Input<string>? LoadBalancingGroup { get; set; }

        /// <summary>
        /// Maximum number of associated clients.
        /// </summary>
        [Input("maxStaCount")]
        public Input<int>? MaxStaCount { get; set; }

        /// <summary>
        /// Set operational mode. Only **ap** currently supported.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// When set to full multicast packets will be sent with unicast destination MAC address, resolving multicast problem on a wireless link. This option should be enabled only on the access point, clients should be configured in station-bridge mode.
        /// </summary>
        [Input("multicastHelper")]
        public Input<string>? MulticastHelper { get; set; }

        /// <summary>
        /// Changing the name of this resource will force it to be recreated. &gt; The links of other configuration properties to this
        /// resource may be lost! &gt; Changing the name of the resource outside of a Terraform will result in a loss of control
        /// integrity for that resource!
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rates")]
        private InputMap<string>? _rates;

        /// <summary>
        /// Rates inline settings.
        /// </summary>
        public InputMap<string> Rates
        {
            get => _rates ?? (_rates = new InputMap<string>());
            set => _rates = value;
        }

        [Input("rxChains")]
        private InputList<int>? _rxChains;

        /// <summary>
        /// Which antennas to use for receive.
        /// </summary>
        public InputList<int> RxChains
        {
            get => _rxChains ?? (_rxChains = new InputList<int>());
            set => _rxChains = value;
        }

        [Input("security")]
        private InputMap<string>? _security;

        /// <summary>
        /// Security inline settings.
        /// </summary>
        public InputMap<string> Security
        {
            get => _security ?? (_security = new InputMap<string>());
            set => _security = value;
        }

        /// <summary>
        /// SSID (service set identifier) is a name broadcast in the beacons that identifies wireless network.
        /// </summary>
        [Input("ssid")]
        public Input<string>? Ssid { get; set; }

        [Input("txChains")]
        private InputList<int>? _txChains;

        /// <summary>
        /// Which antennas to use for transmit.
        /// </summary>
        public InputList<int> TxChains
        {
            get => _txChains ?? (_txChains = new InputList<int>());
            set => _txChains = value;
        }

        public ResourceCapsManConfigurationArgs()
        {
        }
        public static new ResourceCapsManConfigurationArgs Empty => new ResourceCapsManConfigurationArgs();
    }

    public sealed class ResourceCapsManConfigurationState : global::Pulumi.ResourceArgs
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
        /// &lt;em&gt;A set of transformations for field names. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___ts_")]
        public Input<string>? ___ts_ { get; set; }

        [Input("channel")]
        private InputMap<string>? _channel;

        /// <summary>
        /// Channel inline settings.
        /// </summary>
        public InputMap<string> Channel
        {
            get => _channel ?? (_channel = new InputMap<string>());
            set => _channel = value;
        }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies default value of scan-list. Value no*country*set is an FCC compliant set of channels.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("datapath")]
        private InputMap<string>? _datapath;

        /// <summary>
        /// Datapath inline settings.
        /// </summary>
        public InputMap<string> Datapath
        {
            get => _datapath ?? (_datapath = new InputMap<string>());
            set => _datapath = value;
        }

        /// <summary>
        /// This interval is measured from third sending failure on the lowest data rate. At this point 3 * (hw-retries + 1) frame transmits on the lowest data rate had failed. During disconnect-timeout packet transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully during disconnect-timeout, the connection is closed, and this event is logged as "extensive data loss". Successful frame transmission resets this timer.
        /// </summary>
        [Input("disconnectTimeout")]
        public Input<string>? DisconnectTimeout { get; set; }

        /// <summary>
        /// How long to wait for confirmation of unicast frames (ACKs) before considering transmission unsuccessful, or in short ACK-Timeout.
        /// </summary>
        [Input("distance")]
        public Input<string>? Distance { get; set; }

        /// <summary>
        /// Discard frames that have been queued for sending longer than frame-lifetime. By default, when value of this property is 0, frames are discarded only after connection is closed (format: 0.00 sec).
        /// </summary>
        [Input("frameLifetime")]
        public Input<string>? FrameLifetime { get; set; }

        /// <summary>
        /// Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how this may affect throughput). "any" will use either short or long, depending on data rate, "long" will use long.
        /// </summary>
        [Input("guardInterval")]
        public Input<string>? GuardInterval { get; set; }

        /// <summary>
        /// This property has effect only in AP mode. Setting it to yes can remove this network from the list of wireless networks that are shown by some client software. Changing this setting does not improve the security of the wireless network, because SSID is included in other frames sent by the AP.
        /// </summary>
        [Input("hideSsid")]
        public Input<bool>? HideSsid { get; set; }

        /// <summary>
        /// Frame protection support property. [See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#Frame_protection_support_(RTS/CTS)).
        /// </summary>
        [Input("hwProtectionMode")]
        public Input<string>? HwProtectionMode { get; set; }

        /// <summary>
        /// Number of times sending frame is retried without considering it a transmission failure. [See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless)
        /// </summary>
        [Input("hwRetries")]
        public Input<int>? HwRetries { get; set; }

        /// <summary>
        /// Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.
        /// </summary>
        [Input("installation")]
        public Input<string>? Installation { get; set; }

        /// <summary>
        /// If a client has not communicated for around 20 seconds, AP sends a "keepalive-frame".
        /// </summary>
        [Input("keepaliveFrames")]
        public Input<string>? KeepaliveFrames { get; set; }

        /// <summary>
        /// Tags the interface to the load balancing group. For a client to connect to interface in this group, the interface should have the same number of already connected clients as all other interfaces in the group or smaller. Useful in setups where ranges of CAPs mostly overlap.
        /// </summary>
        [Input("loadBalancingGroup")]
        public Input<string>? LoadBalancingGroup { get; set; }

        /// <summary>
        /// Maximum number of associated clients.
        /// </summary>
        [Input("maxStaCount")]
        public Input<int>? MaxStaCount { get; set; }

        /// <summary>
        /// Set operational mode. Only **ap** currently supported.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// When set to full multicast packets will be sent with unicast destination MAC address, resolving multicast problem on a wireless link. This option should be enabled only on the access point, clients should be configured in station-bridge mode.
        /// </summary>
        [Input("multicastHelper")]
        public Input<string>? MulticastHelper { get; set; }

        /// <summary>
        /// Changing the name of this resource will force it to be recreated. &gt; The links of other configuration properties to this
        /// resource may be lost! &gt; Changing the name of the resource outside of a Terraform will result in a loss of control
        /// integrity for that resource!
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rates")]
        private InputMap<string>? _rates;

        /// <summary>
        /// Rates inline settings.
        /// </summary>
        public InputMap<string> Rates
        {
            get => _rates ?? (_rates = new InputMap<string>());
            set => _rates = value;
        }

        [Input("rxChains")]
        private InputList<int>? _rxChains;

        /// <summary>
        /// Which antennas to use for receive.
        /// </summary>
        public InputList<int> RxChains
        {
            get => _rxChains ?? (_rxChains = new InputList<int>());
            set => _rxChains = value;
        }

        [Input("security")]
        private InputMap<string>? _security;

        /// <summary>
        /// Security inline settings.
        /// </summary>
        public InputMap<string> Security
        {
            get => _security ?? (_security = new InputMap<string>());
            set => _security = value;
        }

        /// <summary>
        /// SSID (service set identifier) is a name broadcast in the beacons that identifies wireless network.
        /// </summary>
        [Input("ssid")]
        public Input<string>? Ssid { get; set; }

        [Input("txChains")]
        private InputList<int>? _txChains;

        /// <summary>
        /// Which antennas to use for transmit.
        /// </summary>
        public InputList<int> TxChains
        {
            get => _txChains ?? (_txChains = new InputList<int>());
            set => _txChains = value;
        }

        public ResourceCapsManConfigurationState()
        {
        }
        public static new ResourceCapsManConfigurationState Empty => new ResourceCapsManConfigurationState();
    }
}
