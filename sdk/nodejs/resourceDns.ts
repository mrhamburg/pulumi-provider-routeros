// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # routeros.ResourceDns (Resource)
 *
 * A MikroTik router with DNS feature enabled can be set as a DNS server for any DNS-compliant client.
 *
 * ## Import
 *
 * #The DNS Settings can not be imported.
 *
 * #Terraform will ignore the current settings and will overwrite the current settings with the settings defined in Terraform.
 */
export class ResourceDns extends pulumi.CustomResource {
    /**
     * Get an existing ResourceDns resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceDnsState, opts?: pulumi.CustomResourceOptions): ResourceDns {
        return new ResourceDns(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'routeros:index/resourceDns:ResourceDns';

    /**
     * Returns true if the given object is an instance of ResourceDns.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceDns {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceDns.__pulumiType;
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
     * Specifies whether to allow network requests.
     */
    public readonly allowRemoteRequests!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
     * time. Shorter TTL received from DNS servers are respected. *Default: 1w*
     */
    public readonly cacheMaxTtl!: pulumi.Output<string>;
    /**
     * Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
     */
    public readonly cacheSize!: pulumi.Output<number>;
    /**
     * Shows the currently used cache size in KiB.
     */
    public /*out*/ readonly cacheUsed!: pulumi.Output<number>;
    /**
     * Specifies how many DoH concurrent queries are allowed.
     */
    public readonly dohMaxConcurrentQueries!: pulumi.Output<number>;
    /**
     * Specifies how many concurrent connections to the DoH server are allowed.
     */
    public readonly dohMaxServerConnections!: pulumi.Output<number>;
    /**
     * Specifies how long to wait for query response from the DoH server.
     */
    public readonly dohTimeout!: pulumi.Output<string>;
    /**
     * List of dynamically added DNS server from different services, for example, DHCP.
     */
    public /*out*/ readonly dynamicServers!: pulumi.Output<string>;
    /**
     * Specifies how much concurrent queries are allowed. *Default: 100*
     */
    public readonly maxConcurrentQueries!: pulumi.Output<number>;
    /**
     * Specifies how much concurrent TCP sessions are allowed. *Default: 20*
     */
    public readonly maxConcurrentTcpSessions!: pulumi.Output<number>;
    /**
     * Maximum size of allowed UDP packet. *Default: 4096*
     */
    public readonly maxUdpPacketSize!: pulumi.Output<number>;
    /**
     * Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
     */
    public readonly queryServerTimeout!: pulumi.Output<string>;
    /**
     * Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
     * query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
     */
    public readonly queryTotalTimeout!: pulumi.Output<string>;
    /**
     * List of DNS server IPv4/IPv6 addresses.
     */
    public readonly servers!: pulumi.Output<string | undefined>;
    /**
     * DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
     * fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
     * on the device.
     */
    public readonly useDohServer!: pulumi.Output<string | undefined>;
    /**
     * DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
     */
    public readonly verifyDohCert!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ResourceDns resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourceDnsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceDnsArgs | ResourceDnsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceDnsState | undefined;
            resourceInputs["___id_"] = state ? state.___id_ : undefined;
            resourceInputs["___path_"] = state ? state.___path_ : undefined;
            resourceInputs["allowRemoteRequests"] = state ? state.allowRemoteRequests : undefined;
            resourceInputs["cacheMaxTtl"] = state ? state.cacheMaxTtl : undefined;
            resourceInputs["cacheSize"] = state ? state.cacheSize : undefined;
            resourceInputs["cacheUsed"] = state ? state.cacheUsed : undefined;
            resourceInputs["dohMaxConcurrentQueries"] = state ? state.dohMaxConcurrentQueries : undefined;
            resourceInputs["dohMaxServerConnections"] = state ? state.dohMaxServerConnections : undefined;
            resourceInputs["dohTimeout"] = state ? state.dohTimeout : undefined;
            resourceInputs["dynamicServers"] = state ? state.dynamicServers : undefined;
            resourceInputs["maxConcurrentQueries"] = state ? state.maxConcurrentQueries : undefined;
            resourceInputs["maxConcurrentTcpSessions"] = state ? state.maxConcurrentTcpSessions : undefined;
            resourceInputs["maxUdpPacketSize"] = state ? state.maxUdpPacketSize : undefined;
            resourceInputs["queryServerTimeout"] = state ? state.queryServerTimeout : undefined;
            resourceInputs["queryTotalTimeout"] = state ? state.queryTotalTimeout : undefined;
            resourceInputs["servers"] = state ? state.servers : undefined;
            resourceInputs["useDohServer"] = state ? state.useDohServer : undefined;
            resourceInputs["verifyDohCert"] = state ? state.verifyDohCert : undefined;
        } else {
            const args = argsOrState as ResourceDnsArgs | undefined;
            resourceInputs["___id_"] = args ? args.___id_ : undefined;
            resourceInputs["___path_"] = args ? args.___path_ : undefined;
            resourceInputs["allowRemoteRequests"] = args ? args.allowRemoteRequests : undefined;
            resourceInputs["cacheMaxTtl"] = args ? args.cacheMaxTtl : undefined;
            resourceInputs["cacheSize"] = args ? args.cacheSize : undefined;
            resourceInputs["dohMaxConcurrentQueries"] = args ? args.dohMaxConcurrentQueries : undefined;
            resourceInputs["dohMaxServerConnections"] = args ? args.dohMaxServerConnections : undefined;
            resourceInputs["dohTimeout"] = args ? args.dohTimeout : undefined;
            resourceInputs["maxConcurrentQueries"] = args ? args.maxConcurrentQueries : undefined;
            resourceInputs["maxConcurrentTcpSessions"] = args ? args.maxConcurrentTcpSessions : undefined;
            resourceInputs["maxUdpPacketSize"] = args ? args.maxUdpPacketSize : undefined;
            resourceInputs["queryServerTimeout"] = args ? args.queryServerTimeout : undefined;
            resourceInputs["queryTotalTimeout"] = args ? args.queryTotalTimeout : undefined;
            resourceInputs["servers"] = args ? args.servers : undefined;
            resourceInputs["useDohServer"] = args ? args.useDohServer : undefined;
            resourceInputs["verifyDohCert"] = args ? args.verifyDohCert : undefined;
            resourceInputs["cacheUsed"] = undefined /*out*/;
            resourceInputs["dynamicServers"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceDns.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceDns resources.
 */
export interface ResourceDnsState {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * Specifies whether to allow network requests.
     */
    allowRemoteRequests?: pulumi.Input<boolean>;
    /**
     * Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
     * time. Shorter TTL received from DNS servers are respected. *Default: 1w*
     */
    cacheMaxTtl?: pulumi.Input<string>;
    /**
     * Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
     */
    cacheSize?: pulumi.Input<number>;
    /**
     * Shows the currently used cache size in KiB.
     */
    cacheUsed?: pulumi.Input<number>;
    /**
     * Specifies how many DoH concurrent queries are allowed.
     */
    dohMaxConcurrentQueries?: pulumi.Input<number>;
    /**
     * Specifies how many concurrent connections to the DoH server are allowed.
     */
    dohMaxServerConnections?: pulumi.Input<number>;
    /**
     * Specifies how long to wait for query response from the DoH server.
     */
    dohTimeout?: pulumi.Input<string>;
    /**
     * List of dynamically added DNS server from different services, for example, DHCP.
     */
    dynamicServers?: pulumi.Input<string>;
    /**
     * Specifies how much concurrent queries are allowed. *Default: 100*
     */
    maxConcurrentQueries?: pulumi.Input<number>;
    /**
     * Specifies how much concurrent TCP sessions are allowed. *Default: 20*
     */
    maxConcurrentTcpSessions?: pulumi.Input<number>;
    /**
     * Maximum size of allowed UDP packet. *Default: 4096*
     */
    maxUdpPacketSize?: pulumi.Input<number>;
    /**
     * Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
     */
    queryServerTimeout?: pulumi.Input<string>;
    /**
     * Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
     * query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
     */
    queryTotalTimeout?: pulumi.Input<string>;
    /**
     * List of DNS server IPv4/IPv6 addresses.
     */
    servers?: pulumi.Input<string>;
    /**
     * DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
     * fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
     * on the device.
     */
    useDohServer?: pulumi.Input<string>;
    /**
     * DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
     */
    verifyDohCert?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ResourceDns resource.
 */
export interface ResourceDnsArgs {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * Specifies whether to allow network requests.
     */
    allowRemoteRequests?: pulumi.Input<boolean>;
    /**
     * Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
     * time. Shorter TTL received from DNS servers are respected. *Default: 1w*
     */
    cacheMaxTtl?: pulumi.Input<string>;
    /**
     * Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
     */
    cacheSize?: pulumi.Input<number>;
    /**
     * Specifies how many DoH concurrent queries are allowed.
     */
    dohMaxConcurrentQueries?: pulumi.Input<number>;
    /**
     * Specifies how many concurrent connections to the DoH server are allowed.
     */
    dohMaxServerConnections?: pulumi.Input<number>;
    /**
     * Specifies how long to wait for query response from the DoH server.
     */
    dohTimeout?: pulumi.Input<string>;
    /**
     * Specifies how much concurrent queries are allowed. *Default: 100*
     */
    maxConcurrentQueries?: pulumi.Input<number>;
    /**
     * Specifies how much concurrent TCP sessions are allowed. *Default: 20*
     */
    maxConcurrentTcpSessions?: pulumi.Input<number>;
    /**
     * Maximum size of allowed UDP packet. *Default: 4096*
     */
    maxUdpPacketSize?: pulumi.Input<number>;
    /**
     * Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
     */
    queryServerTimeout?: pulumi.Input<string>;
    /**
     * Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
     * query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
     */
    queryTotalTimeout?: pulumi.Input<string>;
    /**
     * List of DNS server IPv4/IPv6 addresses.
     */
    servers?: pulumi.Input<string>;
    /**
     * DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
     * fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
     * on the device.
     */
    useDohServer?: pulumi.Input<string>;
    /**
     * DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
     */
    verifyDohCert?: pulumi.Input<boolean>;
}
