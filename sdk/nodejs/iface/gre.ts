// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # routeros.Iface.Gre (Resource)
 *
 * ***
 *
 * #### This is an alias for backwards compatibility between plugin versions.
 * Please see documentation for routeros.Iface.InterfaceGre
 */
export class Gre extends pulumi.CustomResource {
    /**
     * Get an existing Gre resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GreState, opts?: pulumi.CustomResourceOptions): Gre {
        return new Gre(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'routeros:Iface/gre:Gre';

    /**
     * Returns true if the given object is an instance of Gre.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gre {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gre.__pulumiType;
    }

    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___id_!: pulumi.Output<number | undefined>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___path_!: pulumi.Output<string | undefined>;
    public /*out*/ readonly actualMtu!: pulumi.Output<number>;
    /**
     * Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
     */
    public readonly allowFastPath!: pulumi.Output<boolean | undefined>;
    /**
     * Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for
     * received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP
     * overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is
     * changed.
     */
    public readonly clampTcpMss!: pulumi.Output<boolean | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    public readonly dontFragment!: pulumi.Output<string | undefined>;
    /**
     * Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
     */
    public readonly dscp!: pulumi.Output<string | undefined>;
    /**
     * When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default
     * phase2 uses sha1/aes128cbc).
     */
    public readonly ipsecSecret!: pulumi.Output<string | undefined>;
    /**
     * Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of
     * tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in
     * following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries -
     * number of retry attempts. KeepaliveInterval is integer 0..4294967295
     */
    public readonly keepalive!: pulumi.Output<string | undefined>;
    /**
     * Layer2 Maximum transmission unit.
     */
    public /*out*/ readonly l2mtu!: pulumi.Output<number>;
    public readonly localAddress!: pulumi.Output<string | undefined>;
    /**
     * Layer3 Maximum transmission unit ('auto', 0 .. 65535)
     */
    public readonly mtu!: pulumi.Output<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    public readonly name!: pulumi.Output<string>;
    public readonly remoteAddress!: pulumi.Output<string>;
    public /*out*/ readonly running!: pulumi.Output<boolean>;

    /**
     * Create a Gre resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GreArgs | GreState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GreState | undefined;
            resourceInputs["___id_"] = state ? state.___id_ : undefined;
            resourceInputs["___path_"] = state ? state.___path_ : undefined;
            resourceInputs["actualMtu"] = state ? state.actualMtu : undefined;
            resourceInputs["allowFastPath"] = state ? state.allowFastPath : undefined;
            resourceInputs["clampTcpMss"] = state ? state.clampTcpMss : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["dontFragment"] = state ? state.dontFragment : undefined;
            resourceInputs["dscp"] = state ? state.dscp : undefined;
            resourceInputs["ipsecSecret"] = state ? state.ipsecSecret : undefined;
            resourceInputs["keepalive"] = state ? state.keepalive : undefined;
            resourceInputs["l2mtu"] = state ? state.l2mtu : undefined;
            resourceInputs["localAddress"] = state ? state.localAddress : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["remoteAddress"] = state ? state.remoteAddress : undefined;
            resourceInputs["running"] = state ? state.running : undefined;
        } else {
            const args = argsOrState as GreArgs | undefined;
            if ((!args || args.remoteAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteAddress'");
            }
            resourceInputs["___id_"] = args ? args.___id_ : undefined;
            resourceInputs["___path_"] = args ? args.___path_ : undefined;
            resourceInputs["allowFastPath"] = args ? args.allowFastPath : undefined;
            resourceInputs["clampTcpMss"] = args ? args.clampTcpMss : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["dontFragment"] = args ? args.dontFragment : undefined;
            resourceInputs["dscp"] = args ? args.dscp : undefined;
            resourceInputs["ipsecSecret"] = args?.ipsecSecret ? pulumi.secret(args.ipsecSecret) : undefined;
            resourceInputs["keepalive"] = args ? args.keepalive : undefined;
            resourceInputs["localAddress"] = args ? args.localAddress : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["remoteAddress"] = args ? args.remoteAddress : undefined;
            resourceInputs["actualMtu"] = undefined /*out*/;
            resourceInputs["l2mtu"] = undefined /*out*/;
            resourceInputs["running"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["ipsecSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Gre.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Gre resources.
 */
export interface GreState {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    actualMtu?: pulumi.Input<number>;
    /**
     * Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
     */
    allowFastPath?: pulumi.Input<boolean>;
    /**
     * Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for
     * received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP
     * overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is
     * changed.
     */
    clampTcpMss?: pulumi.Input<boolean>;
    comment?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    dontFragment?: pulumi.Input<string>;
    /**
     * Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
     */
    dscp?: pulumi.Input<string>;
    /**
     * When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default
     * phase2 uses sha1/aes128cbc).
     */
    ipsecSecret?: pulumi.Input<string>;
    /**
     * Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of
     * tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in
     * following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries -
     * number of retry attempts. KeepaliveInterval is integer 0..4294967295
     */
    keepalive?: pulumi.Input<string>;
    /**
     * Layer2 Maximum transmission unit.
     */
    l2mtu?: pulumi.Input<number>;
    localAddress?: pulumi.Input<string>;
    /**
     * Layer3 Maximum transmission unit ('auto', 0 .. 65535)
     */
    mtu?: pulumi.Input<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    name?: pulumi.Input<string>;
    remoteAddress?: pulumi.Input<string>;
    running?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Gre resource.
 */
export interface GreArgs {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
     */
    allowFastPath?: pulumi.Input<boolean>;
    /**
     * Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for
     * received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP
     * overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is
     * changed.
     */
    clampTcpMss?: pulumi.Input<boolean>;
    comment?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    dontFragment?: pulumi.Input<string>;
    /**
     * Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
     */
    dscp?: pulumi.Input<string>;
    /**
     * When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default
     * phase2 uses sha1/aes128cbc).
     */
    ipsecSecret?: pulumi.Input<string>;
    /**
     * Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of
     * tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in
     * following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries -
     * number of retry attempts. KeepaliveInterval is integer 0..4294967295
     */
    keepalive?: pulumi.Input<string>;
    localAddress?: pulumi.Input<string>;
    /**
     * Layer3 Maximum transmission unit ('auto', 0 .. 65535)
     */
    mtu?: pulumi.Input<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    name?: pulumi.Input<string>;
    remoteAddress: pulumi.Input<string>;
}