// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # routeros.Ip.DhcpServerLease (Resource)
 *
 * ***
 *
 * #### This is an alias for backwards compatibility between plugin versions.
 * Please see documentation for routeros.Ip.DhcpIpServerLease
 */
export class DhcpServerLease extends pulumi.CustomResource {
    /**
     * Get an existing DhcpServerLease resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DhcpServerLeaseState, opts?: pulumi.CustomResourceOptions): DhcpServerLease {
        return new DhcpServerLease(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'routeros:Ip/dhcpServerLease:DhcpServerLease';

    /**
     * Returns true if the given object is an instance of DhcpServerLease.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DhcpServerLease {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DhcpServerLease.__pulumiType;
    }

    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___id_!: pulumi.Output<number | undefined>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___path_!: pulumi.Output<string | undefined>;
    /**
     * The IP address of the machine currently holding the DHCP lease.
     */
    public /*out*/ readonly activeAddress!: pulumi.Output<string>;
    /**
     * Actual client-id of the client.
     */
    public /*out*/ readonly activeClientId!: pulumi.Output<string>;
    /**
     * The hostname of the machine currently holding the DHCP lease.
     */
    public /*out*/ readonly activeHostname!: pulumi.Output<string>;
    /**
     * The MAC address of of the machine currently holding the DHCP lease.
     */
    public /*out*/ readonly activeMacAddress!: pulumi.Output<string>;
    /**
     * Actual dhcp server, which serves this client.
     */
    public /*out*/ readonly activeServer!: pulumi.Output<string>;
    /**
     * The IP address of the DHCP lease to be created.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * Address list to which address will be added if lease is bound.
     */
    public readonly addressLists!: pulumi.Output<string | undefined>;
    /**
     * Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as
     * hex dump.
     */
    public /*out*/ readonly agentCircuitId!: pulumi.Output<string>;
    /**
     * Remote ID, set by DHCP relay agent.
     */
    public /*out*/ readonly agentRemoteId!: pulumi.Output<string>;
    /**
     * Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
     */
    public readonly allowDualStackQueue!: pulumi.Output<boolean | undefined>;
    /**
     * Send all replies as broadcasts.
     */
    public readonly alwaysBroadcast!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to block access for this DHCP client (true|false).
     */
    public readonly blockAccess!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the lease is blocked.
     */
    public /*out*/ readonly blocked!: pulumi.Output<boolean>;
    /**
     * If specified, must match DHCP 'client identifier' option of the request.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Add additional DHCP options.
     */
    public readonly dhcpOption!: pulumi.Output<string | undefined>;
    /**
     * Add additional set of DHCP options.
     */
    public readonly dhcpOptionSet!: pulumi.Output<string | undefined>;
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
     * specific device. Defaults to false.
     */
    public readonly dynamic!: pulumi.Output<boolean | undefined>;
    /**
     * Time until lease expires.
     */
    public /*out*/ readonly expiresAfter!: pulumi.Output<string>;
    /**
     * The hostname of the device
     */
    public /*out*/ readonly hostName!: pulumi.Output<string>;
    /**
     * Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
     */
    public readonly insertQueueBefore!: pulumi.Output<string | undefined>;
    public /*out*/ readonly lastSeen!: pulumi.Output<string>;
    /**
     * Time that the client may use the address. If set to 0s lease will never expire.
     */
    public readonly leaseTime!: pulumi.Output<string | undefined>;
    /**
     * The MAC addreess of the DHCP lease to be created.
     */
    public readonly macAddress!: pulumi.Output<string>;
    /**
     * Shows if this dynamic lease is authenticated by RADIUS or not.
     */
    public /*out*/ readonly radius!: pulumi.Output<string>;
    /**
     * Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
     */
    public readonly rateLimit!: pulumi.Output<string | undefined>;
    /**
     * Server name which serves this client.
     */
    public readonly server!: pulumi.Output<string | undefined>;
    /**
     * Source MAC address.
     */
    public /*out*/ readonly srcMacAddress!: pulumi.Output<string>;
    /**
     * Lease status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * When this option is set server uses source MAC address instead of received CHADDR to assign address.
     */
    public readonly useSrcMac!: pulumi.Output<boolean | undefined>;

    /**
     * Create a DhcpServerLease resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DhcpServerLeaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DhcpServerLeaseArgs | DhcpServerLeaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DhcpServerLeaseState | undefined;
            resourceInputs["___id_"] = state ? state.___id_ : undefined;
            resourceInputs["___path_"] = state ? state.___path_ : undefined;
            resourceInputs["activeAddress"] = state ? state.activeAddress : undefined;
            resourceInputs["activeClientId"] = state ? state.activeClientId : undefined;
            resourceInputs["activeHostname"] = state ? state.activeHostname : undefined;
            resourceInputs["activeMacAddress"] = state ? state.activeMacAddress : undefined;
            resourceInputs["activeServer"] = state ? state.activeServer : undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["addressLists"] = state ? state.addressLists : undefined;
            resourceInputs["agentCircuitId"] = state ? state.agentCircuitId : undefined;
            resourceInputs["agentRemoteId"] = state ? state.agentRemoteId : undefined;
            resourceInputs["allowDualStackQueue"] = state ? state.allowDualStackQueue : undefined;
            resourceInputs["alwaysBroadcast"] = state ? state.alwaysBroadcast : undefined;
            resourceInputs["blockAccess"] = state ? state.blockAccess : undefined;
            resourceInputs["blocked"] = state ? state.blocked : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dhcpOption"] = state ? state.dhcpOption : undefined;
            resourceInputs["dhcpOptionSet"] = state ? state.dhcpOptionSet : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["dynamic"] = state ? state.dynamic : undefined;
            resourceInputs["expiresAfter"] = state ? state.expiresAfter : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["insertQueueBefore"] = state ? state.insertQueueBefore : undefined;
            resourceInputs["lastSeen"] = state ? state.lastSeen : undefined;
            resourceInputs["leaseTime"] = state ? state.leaseTime : undefined;
            resourceInputs["macAddress"] = state ? state.macAddress : undefined;
            resourceInputs["radius"] = state ? state.radius : undefined;
            resourceInputs["rateLimit"] = state ? state.rateLimit : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["srcMacAddress"] = state ? state.srcMacAddress : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["useSrcMac"] = state ? state.useSrcMac : undefined;
        } else {
            const args = argsOrState as DhcpServerLeaseArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.macAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'macAddress'");
            }
            resourceInputs["___id_"] = args ? args.___id_ : undefined;
            resourceInputs["___path_"] = args ? args.___path_ : undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["addressLists"] = args ? args.addressLists : undefined;
            resourceInputs["allowDualStackQueue"] = args ? args.allowDualStackQueue : undefined;
            resourceInputs["alwaysBroadcast"] = args ? args.alwaysBroadcast : undefined;
            resourceInputs["blockAccess"] = args ? args.blockAccess : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dhcpOption"] = args ? args.dhcpOption : undefined;
            resourceInputs["dhcpOptionSet"] = args ? args.dhcpOptionSet : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["dynamic"] = args ? args.dynamic : undefined;
            resourceInputs["insertQueueBefore"] = args ? args.insertQueueBefore : undefined;
            resourceInputs["leaseTime"] = args ? args.leaseTime : undefined;
            resourceInputs["macAddress"] = args ? args.macAddress : undefined;
            resourceInputs["rateLimit"] = args ? args.rateLimit : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["useSrcMac"] = args ? args.useSrcMac : undefined;
            resourceInputs["activeAddress"] = undefined /*out*/;
            resourceInputs["activeClientId"] = undefined /*out*/;
            resourceInputs["activeHostname"] = undefined /*out*/;
            resourceInputs["activeMacAddress"] = undefined /*out*/;
            resourceInputs["activeServer"] = undefined /*out*/;
            resourceInputs["agentCircuitId"] = undefined /*out*/;
            resourceInputs["agentRemoteId"] = undefined /*out*/;
            resourceInputs["blocked"] = undefined /*out*/;
            resourceInputs["expiresAfter"] = undefined /*out*/;
            resourceInputs["hostName"] = undefined /*out*/;
            resourceInputs["lastSeen"] = undefined /*out*/;
            resourceInputs["radius"] = undefined /*out*/;
            resourceInputs["srcMacAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DhcpServerLease.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DhcpServerLease resources.
 */
export interface DhcpServerLeaseState {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * The IP address of the machine currently holding the DHCP lease.
     */
    activeAddress?: pulumi.Input<string>;
    /**
     * Actual client-id of the client.
     */
    activeClientId?: pulumi.Input<string>;
    /**
     * The hostname of the machine currently holding the DHCP lease.
     */
    activeHostname?: pulumi.Input<string>;
    /**
     * The MAC address of of the machine currently holding the DHCP lease.
     */
    activeMacAddress?: pulumi.Input<string>;
    /**
     * Actual dhcp server, which serves this client.
     */
    activeServer?: pulumi.Input<string>;
    /**
     * The IP address of the DHCP lease to be created.
     */
    address?: pulumi.Input<string>;
    /**
     * Address list to which address will be added if lease is bound.
     */
    addressLists?: pulumi.Input<string>;
    /**
     * Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as
     * hex dump.
     */
    agentCircuitId?: pulumi.Input<string>;
    /**
     * Remote ID, set by DHCP relay agent.
     */
    agentRemoteId?: pulumi.Input<string>;
    /**
     * Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
     */
    allowDualStackQueue?: pulumi.Input<boolean>;
    /**
     * Send all replies as broadcasts.
     */
    alwaysBroadcast?: pulumi.Input<boolean>;
    /**
     * Whether to block access for this DHCP client (true|false).
     */
    blockAccess?: pulumi.Input<boolean>;
    /**
     * Whether the lease is blocked.
     */
    blocked?: pulumi.Input<boolean>;
    /**
     * If specified, must match DHCP 'client identifier' option of the request.
     */
    clientId?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    /**
     * Add additional DHCP options.
     */
    dhcpOption?: pulumi.Input<string>;
    /**
     * Add additional set of DHCP options.
     */
    dhcpOptionSet?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    /**
     * Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
     * specific device. Defaults to false.
     */
    dynamic?: pulumi.Input<boolean>;
    /**
     * Time until lease expires.
     */
    expiresAfter?: pulumi.Input<string>;
    /**
     * The hostname of the device
     */
    hostName?: pulumi.Input<string>;
    /**
     * Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
     */
    insertQueueBefore?: pulumi.Input<string>;
    lastSeen?: pulumi.Input<string>;
    /**
     * Time that the client may use the address. If set to 0s lease will never expire.
     */
    leaseTime?: pulumi.Input<string>;
    /**
     * The MAC addreess of the DHCP lease to be created.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * Shows if this dynamic lease is authenticated by RADIUS or not.
     */
    radius?: pulumi.Input<string>;
    /**
     * Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
     */
    rateLimit?: pulumi.Input<string>;
    /**
     * Server name which serves this client.
     */
    server?: pulumi.Input<string>;
    /**
     * Source MAC address.
     */
    srcMacAddress?: pulumi.Input<string>;
    /**
     * Lease status.
     */
    status?: pulumi.Input<string>;
    /**
     * When this option is set server uses source MAC address instead of received CHADDR to assign address.
     */
    useSrcMac?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DhcpServerLease resource.
 */
export interface DhcpServerLeaseArgs {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * The IP address of the DHCP lease to be created.
     */
    address: pulumi.Input<string>;
    /**
     * Address list to which address will be added if lease is bound.
     */
    addressLists?: pulumi.Input<string>;
    /**
     * Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
     */
    allowDualStackQueue?: pulumi.Input<boolean>;
    /**
     * Send all replies as broadcasts.
     */
    alwaysBroadcast?: pulumi.Input<boolean>;
    /**
     * Whether to block access for this DHCP client (true|false).
     */
    blockAccess?: pulumi.Input<boolean>;
    /**
     * If specified, must match DHCP 'client identifier' option of the request.
     */
    clientId?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    /**
     * Add additional DHCP options.
     */
    dhcpOption?: pulumi.Input<string>;
    /**
     * Add additional set of DHCP options.
     */
    dhcpOptionSet?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    /**
     * Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
     * specific device. Defaults to false.
     */
    dynamic?: pulumi.Input<boolean>;
    /**
     * Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
     */
    insertQueueBefore?: pulumi.Input<string>;
    /**
     * Time that the client may use the address. If set to 0s lease will never expire.
     */
    leaseTime?: pulumi.Input<string>;
    /**
     * The MAC addreess of the DHCP lease to be created.
     */
    macAddress: pulumi.Input<string>;
    /**
     * Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
     */
    rateLimit?: pulumi.Input<string>;
    /**
     * Server name which serves this client.
     */
    server?: pulumi.Input<string>;
    /**
     * When this option is set server uses source MAC address instead of received CHADDR to assign address.
     */
    useSrcMac?: pulumi.Input<boolean>;
}