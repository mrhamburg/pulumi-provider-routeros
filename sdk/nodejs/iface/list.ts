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
 * const list = new routeros.iface.List("list", {});
 * ```
 *
 * ## Import
 *
 * Import with the name of the interface list in case of the example use my-list
 *
 * ```sh
 *  $ pulumi import routeros:Iface/list:List list my-list
 * ```
 */
export class List extends pulumi.CustomResource {
    /**
     * Get an existing List resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListState, opts?: pulumi.CustomResourceOptions): List {
        return new List(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'routeros:Iface/list:List';

    /**
     * Returns true if the given object is an instance of List.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is List {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === List.__pulumiType;
    }

    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___id_!: pulumi.Output<number | undefined>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___path_!: pulumi.Output<string | undefined>;
    public /*out*/ readonly builtin!: pulumi.Output<boolean>;
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
     */
    public /*out*/ readonly dynamic!: pulumi.Output<boolean>;
    public readonly exclude!: pulumi.Output<string | undefined>;
    public readonly include!: pulumi.Output<string | undefined>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a List resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListArgs | ListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListState | undefined;
            resourceInputs["___id_"] = state ? state.___id_ : undefined;
            resourceInputs["___path_"] = state ? state.___path_ : undefined;
            resourceInputs["builtin"] = state ? state.builtin : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamic"] = state ? state.dynamic : undefined;
            resourceInputs["exclude"] = state ? state.exclude : undefined;
            resourceInputs["include"] = state ? state.include : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ListArgs | undefined;
            resourceInputs["___id_"] = args ? args.___id_ : undefined;
            resourceInputs["___path_"] = args ? args.___path_ : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["exclude"] = args ? args.exclude : undefined;
            resourceInputs["include"] = args ? args.include : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["builtin"] = undefined /*out*/;
            resourceInputs["dynamic"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(List.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering List resources.
 */
export interface ListState {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    builtin?: pulumi.Input<boolean>;
    comment?: pulumi.Input<string>;
    /**
     * Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
     */
    dynamic?: pulumi.Input<boolean>;
    exclude?: pulumi.Input<string>;
    include?: pulumi.Input<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a List resource.
 */
export interface ListArgs {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    exclude?: pulumi.Input<string>;
    include?: pulumi.Input<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    name?: pulumi.Input<string>;
}
