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
    /// ## # routeros.ResourceDhcpServerLease (Resource)
    /// 
    /// Creates a DHCP lease on the mikrotik device.
    /// 
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
    ///     var dhcpLease = new Routeros.ResourceDhcpServerLease("dhcpLease", new()
    ///     {
    ///         Address = "10.0.0.2",
    ///         MacAddress = "AA:BB:CC:DD:11:22",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/ip/dhcp-server/lease get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:index/resourceDhcpServerLease:ResourceDhcpServerLease dhcp_lease "*0"
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:index/resourceDhcpServerLease:ResourceDhcpServerLease")]
    public partial class ResourceDhcpServerLease : global::Pulumi.CustomResource
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
        /// The IP address of the machine currently holding the DHCP lease.
        /// </summary>
        [Output("activeAddress")]
        public Output<string> ActiveAddress { get; private set; } = null!;

        /// <summary>
        /// Actual client-id of the client.
        /// </summary>
        [Output("activeClientId")]
        public Output<string> ActiveClientId { get; private set; } = null!;

        /// <summary>
        /// The hostname of the machine currently holding the DHCP lease.
        /// </summary>
        [Output("activeHostname")]
        public Output<string> ActiveHostname { get; private set; } = null!;

        /// <summary>
        /// The MAC address of of the machine currently holding the DHCP lease.
        /// </summary>
        [Output("activeMacAddress")]
        public Output<string> ActiveMacAddress { get; private set; } = null!;

        /// <summary>
        /// Actual dhcp server, which serves this client.
        /// </summary>
        [Output("activeServer")]
        public Output<string> ActiveServer { get; private set; } = null!;

        /// <summary>
        /// The IP address of the DHCP lease to be created.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Address list to which address will be added if lease is bound.
        /// </summary>
        [Output("addressLists")]
        public Output<string?> AddressLists { get; private set; } = null!;

        /// <summary>
        /// Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as hex dump.
        /// </summary>
        [Output("agentCircuitId")]
        public Output<string> AgentCircuitId { get; private set; } = null!;

        /// <summary>
        /// Remote ID, set by DHCP relay agent.
        /// </summary>
        [Output("agentRemoteId")]
        public Output<string> AgentRemoteId { get; private set; } = null!;

        /// <summary>
        /// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
        /// </summary>
        [Output("allowDualStackQueue")]
        public Output<bool?> AllowDualStackQueue { get; private set; } = null!;

        /// <summary>
        /// Send all replies as broadcasts.
        /// </summary>
        [Output("alwaysBroadcast")]
        public Output<bool?> AlwaysBroadcast { get; private set; } = null!;

        /// <summary>
        /// Whether to block access for this DHCP client (true|false).
        /// </summary>
        [Output("blockAccess")]
        public Output<bool?> BlockAccess { get; private set; } = null!;

        /// <summary>
        /// Whether the lease is blocked.
        /// </summary>
        [Output("blocked")]
        public Output<bool> Blocked { get; private set; } = null!;

        /// <summary>
        /// If specified, must match DHCP 'client identifier' option of the request.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Add additional DHCP options.
        /// </summary>
        [Output("dhcpOption")]
        public Output<string?> DhcpOption { get; private set; } = null!;

        /// <summary>
        /// Add additional set of DHCP options.
        /// </summary>
        [Output("dhcpOptionSet")]
        public Output<string?> DhcpOptionSet { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that specific device. Defaults to false.
        /// </summary>
        [Output("dynamic")]
        public Output<bool?> Dynamic { get; private set; } = null!;

        /// <summary>
        /// Time until lease expires.
        /// </summary>
        [Output("expiresAfter")]
        public Output<string> ExpiresAfter { get; private set; } = null!;

        /// <summary>
        /// The hostname of the device
        /// </summary>
        [Output("hostName")]
        public Output<string> HostName { get; private set; } = null!;

        /// <summary>
        /// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
        /// </summary>
        [Output("insertQueueBefore")]
        public Output<string?> InsertQueueBefore { get; private set; } = null!;

        [Output("lastSeen")]
        public Output<string> LastSeen { get; private set; } = null!;

        /// <summary>
        /// Time that the client may use the address. If set to 0s lease will never expire.
        /// </summary>
        [Output("leaseTime")]
        public Output<string?> LeaseTime { get; private set; } = null!;

        /// <summary>
        /// The MAC addreess of the DHCP lease to be created.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        /// <summary>
        /// Shows if this dynamic lease is authenticated by RADIUS or not.
        /// </summary>
        [Output("radius")]
        public Output<string> Radius { get; private set; } = null!;

        /// <summary>
        /// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
        /// </summary>
        [Output("rateLimit")]
        public Output<string?> RateLimit { get; private set; } = null!;

        /// <summary>
        /// Server name which serves this client.
        /// </summary>
        [Output("server")]
        public Output<string?> Server { get; private set; } = null!;

        /// <summary>
        /// Source MAC address.
        /// </summary>
        [Output("srcMacAddress")]
        public Output<string> SrcMacAddress { get; private set; } = null!;

        /// <summary>
        /// Lease status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// When this option is set server uses source MAC address instead of received CHADDR to assign address.
        /// </summary>
        [Output("useSrcMac")]
        public Output<bool?> UseSrcMac { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceDhcpServerLease resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceDhcpServerLease(string name, ResourceDhcpServerLeaseArgs args, CustomResourceOptions? options = null)
            : base("routeros:index/resourceDhcpServerLease:ResourceDhcpServerLease", name, args ?? new ResourceDhcpServerLeaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceDhcpServerLease(string name, Input<string> id, ResourceDhcpServerLeaseState? state = null, CustomResourceOptions? options = null)
            : base("routeros:index/resourceDhcpServerLease:ResourceDhcpServerLease", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceDhcpServerLease resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceDhcpServerLease Get(string name, Input<string> id, ResourceDhcpServerLeaseState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceDhcpServerLease(name, id, state, options);
        }
    }

    public sealed class ResourceDhcpServerLeaseArgs : global::Pulumi.ResourceArgs
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
        /// The IP address of the DHCP lease to be created.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Address list to which address will be added if lease is bound.
        /// </summary>
        [Input("addressLists")]
        public Input<string>? AddressLists { get; set; }

        /// <summary>
        /// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
        /// </summary>
        [Input("allowDualStackQueue")]
        public Input<bool>? AllowDualStackQueue { get; set; }

        /// <summary>
        /// Send all replies as broadcasts.
        /// </summary>
        [Input("alwaysBroadcast")]
        public Input<bool>? AlwaysBroadcast { get; set; }

        /// <summary>
        /// Whether to block access for this DHCP client (true|false).
        /// </summary>
        [Input("blockAccess")]
        public Input<bool>? BlockAccess { get; set; }

        /// <summary>
        /// If specified, must match DHCP 'client identifier' option of the request.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Add additional DHCP options.
        /// </summary>
        [Input("dhcpOption")]
        public Input<string>? DhcpOption { get; set; }

        /// <summary>
        /// Add additional set of DHCP options.
        /// </summary>
        [Input("dhcpOptionSet")]
        public Input<string>? DhcpOptionSet { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that specific device. Defaults to false.
        /// </summary>
        [Input("dynamic")]
        public Input<bool>? Dynamic { get; set; }

        /// <summary>
        /// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
        /// </summary>
        [Input("insertQueueBefore")]
        public Input<string>? InsertQueueBefore { get; set; }

        /// <summary>
        /// Time that the client may use the address. If set to 0s lease will never expire.
        /// </summary>
        [Input("leaseTime")]
        public Input<string>? LeaseTime { get; set; }

        /// <summary>
        /// The MAC addreess of the DHCP lease to be created.
        /// </summary>
        [Input("macAddress", required: true)]
        public Input<string> MacAddress { get; set; } = null!;

        /// <summary>
        /// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
        /// </summary>
        [Input("rateLimit")]
        public Input<string>? RateLimit { get; set; }

        /// <summary>
        /// Server name which serves this client.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// When this option is set server uses source MAC address instead of received CHADDR to assign address.
        /// </summary>
        [Input("useSrcMac")]
        public Input<bool>? UseSrcMac { get; set; }

        public ResourceDhcpServerLeaseArgs()
        {
        }
        public static new ResourceDhcpServerLeaseArgs Empty => new ResourceDhcpServerLeaseArgs();
    }

    public sealed class ResourceDhcpServerLeaseState : global::Pulumi.ResourceArgs
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
        /// The IP address of the machine currently holding the DHCP lease.
        /// </summary>
        [Input("activeAddress")]
        public Input<string>? ActiveAddress { get; set; }

        /// <summary>
        /// Actual client-id of the client.
        /// </summary>
        [Input("activeClientId")]
        public Input<string>? ActiveClientId { get; set; }

        /// <summary>
        /// The hostname of the machine currently holding the DHCP lease.
        /// </summary>
        [Input("activeHostname")]
        public Input<string>? ActiveHostname { get; set; }

        /// <summary>
        /// The MAC address of of the machine currently holding the DHCP lease.
        /// </summary>
        [Input("activeMacAddress")]
        public Input<string>? ActiveMacAddress { get; set; }

        /// <summary>
        /// Actual dhcp server, which serves this client.
        /// </summary>
        [Input("activeServer")]
        public Input<string>? ActiveServer { get; set; }

        /// <summary>
        /// The IP address of the DHCP lease to be created.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Address list to which address will be added if lease is bound.
        /// </summary>
        [Input("addressLists")]
        public Input<string>? AddressLists { get; set; }

        /// <summary>
        /// Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as hex dump.
        /// </summary>
        [Input("agentCircuitId")]
        public Input<string>? AgentCircuitId { get; set; }

        /// <summary>
        /// Remote ID, set by DHCP relay agent.
        /// </summary>
        [Input("agentRemoteId")]
        public Input<string>? AgentRemoteId { get; set; }

        /// <summary>
        /// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
        /// </summary>
        [Input("allowDualStackQueue")]
        public Input<bool>? AllowDualStackQueue { get; set; }

        /// <summary>
        /// Send all replies as broadcasts.
        /// </summary>
        [Input("alwaysBroadcast")]
        public Input<bool>? AlwaysBroadcast { get; set; }

        /// <summary>
        /// Whether to block access for this DHCP client (true|false).
        /// </summary>
        [Input("blockAccess")]
        public Input<bool>? BlockAccess { get; set; }

        /// <summary>
        /// Whether the lease is blocked.
        /// </summary>
        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// If specified, must match DHCP 'client identifier' option of the request.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Add additional DHCP options.
        /// </summary>
        [Input("dhcpOption")]
        public Input<string>? DhcpOption { get; set; }

        /// <summary>
        /// Add additional set of DHCP options.
        /// </summary>
        [Input("dhcpOptionSet")]
        public Input<string>? DhcpOptionSet { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that specific device. Defaults to false.
        /// </summary>
        [Input("dynamic")]
        public Input<bool>? Dynamic { get; set; }

        /// <summary>
        /// Time until lease expires.
        /// </summary>
        [Input("expiresAfter")]
        public Input<string>? ExpiresAfter { get; set; }

        /// <summary>
        /// The hostname of the device
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
        /// </summary>
        [Input("insertQueueBefore")]
        public Input<string>? InsertQueueBefore { get; set; }

        [Input("lastSeen")]
        public Input<string>? LastSeen { get; set; }

        /// <summary>
        /// Time that the client may use the address. If set to 0s lease will never expire.
        /// </summary>
        [Input("leaseTime")]
        public Input<string>? LeaseTime { get; set; }

        /// <summary>
        /// The MAC addreess of the DHCP lease to be created.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// Shows if this dynamic lease is authenticated by RADIUS or not.
        /// </summary>
        [Input("radius")]
        public Input<string>? Radius { get; set; }

        /// <summary>
        /// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
        /// </summary>
        [Input("rateLimit")]
        public Input<string>? RateLimit { get; set; }

        /// <summary>
        /// Server name which serves this client.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Source MAC address.
        /// </summary>
        [Input("srcMacAddress")]
        public Input<string>? SrcMacAddress { get; set; }

        /// <summary>
        /// Lease status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// When this option is set server uses source MAC address instead of received CHADDR to assign address.
        /// </summary>
        [Input("useSrcMac")]
        public Input<bool>? UseSrcMac { get; set; }

        public ResourceDhcpServerLeaseState()
        {
        }
        public static new ResourceDhcpServerLeaseState Empty => new ResourceDhcpServerLeaseState();
    }
}
