// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as routeros from "@pulumi/routeros";
 *
 * const address = new routeros.ip.V4Address("address", {
 *     address: "10.0.0.1",
 *     "interface": "bridge",
 *     network: "10.0.0.0/24",
 * });
 * ```
 *
 * ## Import
 *
 * #The ID can be found via API or the terminal #The command for the terminal is -> :put [/ip/address get [print show-ids]]
 *
 * ```sh
 *  $ pulumi import routeros:Ip/v4Address:V4Address address "*0"
 * ```
 */
export class V4Address extends pulumi.CustomResource {
    /**
     * Get an existing V4Address resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: V4AddressState, opts?: pulumi.CustomResourceOptions): V4Address {
        return new V4Address(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'routeros:Ip/v4Address:V4Address';

    /**
     * Returns true if the given object is an instance of V4Address.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is V4Address {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === V4Address.__pulumiType;
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
     * Name of the actual interface the logical one is bound to.
     */
    public /*out*/ readonly actualInterface!: pulumi.Output<string>;
    /**
     * IP address.
     */
    public readonly address!: pulumi.Output<string>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
     */
    public /*out*/ readonly dynamic!: pulumi.Output<boolean>;
    /**
     * Name of the interface.
     */
    public readonly interface!: pulumi.Output<string>;
    public /*out*/ readonly invalid!: pulumi.Output<boolean>;
    /**
     * IP address for the network. For point-to-point links it should be the address of the remote end. Starting from v5RC6 this parameter is configurable only for addresses with /32 netmask (point to point links)
     */
    public readonly network!: pulumi.Output<string>;

    /**
     * Create a V4Address resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: V4AddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: V4AddressArgs | V4AddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as V4AddressState | undefined;
            resourceInputs["___id_"] = state ? state.___id_ : undefined;
            resourceInputs["___path_"] = state ? state.___path_ : undefined;
            resourceInputs["actualInterface"] = state ? state.actualInterface : undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["dynamic"] = state ? state.dynamic : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["invalid"] = state ? state.invalid : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
        } else {
            const args = argsOrState as V4AddressArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            resourceInputs["___id_"] = args ? args.___id_ : undefined;
            resourceInputs["___path_"] = args ? args.___path_ : undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["actualInterface"] = undefined /*out*/;
            resourceInputs["dynamic"] = undefined /*out*/;
            resourceInputs["invalid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(V4Address.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering V4Address resources.
 */
export interface V4AddressState {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * Name of the actual interface the logical one is bound to.
     */
    actualInterface?: pulumi.Input<string>;
    /**
     * IP address.
     */
    address?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    /**
     * Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
     */
    dynamic?: pulumi.Input<boolean>;
    /**
     * Name of the interface.
     */
    interface?: pulumi.Input<string>;
    invalid?: pulumi.Input<boolean>;
    /**
     * IP address for the network. For point-to-point links it should be the address of the remote end. Starting from v5RC6 this parameter is configurable only for addresses with /32 netmask (point to point links)
     */
    network?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a V4Address resource.
 */
export interface V4AddressArgs {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * IP address.
     */
    address: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    /**
     * Name of the interface.
     */
    interface: pulumi.Input<string>;
    /**
     * IP address for the network. For point-to-point links it should be the address of the remote end. Starting from v5RC6 this parameter is configurable only for addresses with /32 netmask (point to point links)
     */
    network?: pulumi.Input<string>;
}