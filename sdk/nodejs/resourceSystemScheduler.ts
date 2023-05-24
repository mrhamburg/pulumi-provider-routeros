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
 * const schedule1 = new routeros.ResourceSystemScheduler("schedule1", {onEvent: "script name"});
 * ```
 *
 * ## Import
 *
 * #The ID can be found via API or the terminal #The command for the terminal is -> :put [/system/scheduler get [print show-ids]]
 *
 * ```sh
 *  $ pulumi import routeros:index/resourceSystemScheduler:ResourceSystemScheduler schedule1 "*0"
 * ```
 */
export class ResourceSystemScheduler extends pulumi.CustomResource {
    /**
     * Get an existing ResourceSystemScheduler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceSystemSchedulerState, opts?: pulumi.CustomResourceOptions): ResourceSystemScheduler {
        return new ResourceSystemScheduler(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'routeros:index/resourceSystemScheduler:ResourceSystemScheduler';

    /**
     * Returns true if the given object is an instance of ResourceSystemScheduler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceSystemScheduler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceSystemScheduler.__pulumiType;
    }

    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___id_!: pulumi.Output<number | undefined>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    public readonly ___path_!: pulumi.Output<string | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
     * otherwise it is executed repeatedly at the time interval is specified.
     */
    public readonly interval!: pulumi.Output<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly nextRun!: pulumi.Output<string>;
    /**
     * Name of the script to execute. It must be presented at /system script.
     */
    public readonly onEvent!: pulumi.Output<string>;
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
     * full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
     * used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
     * that grants user management rights. Should be used together with the write policy. Allows also to see global variables
     * created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
     * configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
     * Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
     * Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
     * displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
     * ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
     * access to the router's configuration, except for user management. This policy does not allow to read the configuration,
     * so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
     */
    public readonly policies!: pulumi.Output<string[]>;
    /**
     * This counter is incremented each time the script is executed.
     */
    public /*out*/ readonly runCount!: pulumi.Output<string>;
    /**
     * Date of the first script execution.
     */
    public readonly startDate!: pulumi.Output<string>;
    /**
     * Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
     * start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
     * and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
     * will not run at startup.
     */
    public readonly startTime!: pulumi.Output<string>;

    /**
     * Create a ResourceSystemScheduler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceSystemSchedulerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceSystemSchedulerArgs | ResourceSystemSchedulerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceSystemSchedulerState | undefined;
            resourceInputs["___id_"] = state ? state.___id_ : undefined;
            resourceInputs["___path_"] = state ? state.___path_ : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nextRun"] = state ? state.nextRun : undefined;
            resourceInputs["onEvent"] = state ? state.onEvent : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["runCount"] = state ? state.runCount : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as ResourceSystemSchedulerArgs | undefined;
            if ((!args || args.onEvent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'onEvent'");
            }
            resourceInputs["___id_"] = args ? args.___id_ : undefined;
            resourceInputs["___path_"] = args ? args.___path_ : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["onEvent"] = args ? args.onEvent : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["startDate"] = args ? args.startDate : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["nextRun"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["runCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceSystemScheduler.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceSystemScheduler resources.
 */
export interface ResourceSystemSchedulerState {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    /**
     * Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
     * otherwise it is executed repeatedly at the time interval is specified.
     */
    interval?: pulumi.Input<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    name?: pulumi.Input<string>;
    nextRun?: pulumi.Input<string>;
    /**
     * Name of the script to execute. It must be presented at /system script.
     */
    onEvent?: pulumi.Input<string>;
    owner?: pulumi.Input<string>;
    /**
     * List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
     * full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
     * used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
     * that grants user management rights. Should be used together with the write policy. Allows also to see global variables
     * created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
     * configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
     * Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
     * Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
     * displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
     * ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
     * access to the router's configuration, except for user management. This policy does not allow to read the configuration,
     * so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This counter is incremented each time the script is executed.
     */
    runCount?: pulumi.Input<string>;
    /**
     * Date of the first script execution.
     */
    startDate?: pulumi.Input<string>;
    /**
     * Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
     * start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
     * and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
     * will not run at startup.
     */
    startTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceSystemScheduler resource.
 */
export interface ResourceSystemSchedulerArgs {
    /**
     * <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
     */
    ___id_?: pulumi.Input<number>;
    /**
     * <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
     */
    ___path_?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    disabled?: pulumi.Input<boolean>;
    /**
     * Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
     * otherwise it is executed repeatedly at the time interval is specified.
     */
    interval?: pulumi.Input<string>;
    /**
     * Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
     * resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
     * integrity for that resource!
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the script to execute. It must be presented at /system script.
     */
    onEvent: pulumi.Input<string>;
    /**
     * List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
     * full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
     * used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
     * that grants user management rights. Should be used together with the write policy. Allows also to see global variables
     * created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
     * configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
     * Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
     * Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
     * displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
     * ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
     * access to the router's configuration, except for user management. This policy does not allow to read the configuration,
     * so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date of the first script execution.
     */
    startDate?: pulumi.Input<string>;
    /**
     * Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
     * start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
     * and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
     * will not run at startup.
     */
    startTime?: pulumi.Input<string>;
}
