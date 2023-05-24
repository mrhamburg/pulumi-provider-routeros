// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as routeros from "@pulumi/routeros";
 *
 * const test3a = new routeros.ResourceCapsManAaa("test3a", {
 *     calledFormat: "ssid",
 *     macMode: "as-username-and-password",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import routeros:index/resourceCapsManAaa:ResourceCapsManAaa test_3a .
 * ```
 */
export class ResourceCapsManAaa extends pulumi.CustomResource {
    /**
     * Get an existing ResourceCapsManAaa resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceCapsManAaaState, opts?: pulumi.CustomResourceOptions): ResourceCapsManAaa {
        return new ResourceCapsManAaa(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'routeros:index/resourceCapsManAaa:ResourceCapsManAaa';

    /**
     * Returns true if the given object is an instance of ResourceCapsManAaa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceCapsManAaa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceCapsManAaa.__pulumiType;
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
     * Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
     */
    public readonly calledFormat!: pulumi.Output<string | undefined>;
    /**
     * When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
     */
    public readonly interimUpdate!: pulumi.Output<string>;
    /**
     * If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
     */
    public readonly macCaching!: pulumi.Output<string>;
    /**
     * Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
     */
    public readonly macFormat!: pulumi.Output<string | undefined>;
    /**
     * By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
     */
    public readonly macMode!: pulumi.Output<string | undefined>;

    /**
     * Create a ResourceCapsManAaa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourceCapsManAaaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceCapsManAaaArgs | ResourceCapsManAaaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceCapsManAaaState | undefined;
            resourceInputs["___id_"] = state ? state.___id_ : undefined;
            resourceInputs["___path_"] = state ? state.___path_ : undefined;
            resourceInputs["calledFormat"] = state ? state.calledFormat : undefined;
            resourceInputs["interimUpdate"] = state ? state.interimUpdate : undefined;
            resourceInputs["macCaching"] = state ? state.macCaching : undefined;
            resourceInputs["macFormat"] = state ? state.macFormat : undefined;
            resourceInputs["macMode"] = state ? state.macMode : undefined;
        } else {
            const args = argsOrState as ResourceCapsManAaaArgs | undefined;
            resourceInputs["___id_"] = args ? args.___id_ : undefined;
            resourceInputs["___path_"] = args ? args.___path_ : undefined;
            resourceInputs["calledFormat"] = args ? args.calledFormat : undefined;
            resourceInputs["interimUpdate"] = args ? args.interimUpdate : undefined;
            resourceInputs["macCaching"] = args ? args.macCaching : undefined;
            resourceInputs["macFormat"] = args ? args.macFormat : undefined;
            resourceInputs["macMode"] = args ? args.macMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceCapsManAaa.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceCapsManAaa resources.
 */
export interface ResourceCapsManAaaState {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
     */
    calledFormat?: pulumi.Input<string>;
    /**
     * When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
     */
    interimUpdate?: pulumi.Input<string>;
    /**
     * If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
     */
    macCaching?: pulumi.Input<string>;
    /**
     * Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
     */
    macFormat?: pulumi.Input<string>;
    /**
     * By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
     */
    macMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceCapsManAaa resource.
 */
export interface ResourceCapsManAaaArgs {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    /**
     * Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
     */
    calledFormat?: pulumi.Input<string>;
    /**
     * When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
     */
    interimUpdate?: pulumi.Input<string>;
    /**
     * If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
     */
    macCaching?: pulumi.Input<string>;
    /**
     * Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
     */
    macFormat?: pulumi.Input<string>;
    /**
     * By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.
     */
    macMode?: pulumi.Input<string>;
}
