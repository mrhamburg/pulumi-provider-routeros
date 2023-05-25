// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CertificateArgs, CertificateState } from "./certificate";
export type Certificate = import("./certificate").Certificate;
export const Certificate: typeof import("./certificate").Certificate = null as any;
utilities.lazyLoad(exports, ["Certificate"], () => require("./certificate"));

export { IdentityArgs, IdentityState } from "./identity";
export type Identity = import("./identity").Identity;
export const Identity: typeof import("./identity").Identity = null as any;
utilities.lazyLoad(exports, ["Identity"], () => require("./identity"));

export { SchedulerArgs, SchedulerState } from "./scheduler";
export type Scheduler = import("./scheduler").Scheduler;
export const Scheduler: typeof import("./scheduler").Scheduler = null as any;
utilities.lazyLoad(exports, ["Scheduler"], () => require("./scheduler"));

export { SystemIdentityArgs, SystemIdentityState } from "./systemIdentity";
export type SystemIdentity = import("./systemIdentity").SystemIdentity;
export const SystemIdentity: typeof import("./systemIdentity").SystemIdentity = null as any;
utilities.lazyLoad(exports, ["SystemIdentity"], () => require("./systemIdentity"));

export { SystemSchedulerArgs, SystemSchedulerState } from "./systemScheduler";
export type SystemScheduler = import("./systemScheduler").SystemScheduler;
export const SystemScheduler: typeof import("./systemScheduler").SystemScheduler = null as any;
utilities.lazyLoad(exports, ["SystemScheduler"], () => require("./systemScheduler"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "routeros:System/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "routeros:System/identity:Identity":
                return new Identity(name, <any>undefined, { urn })
            case "routeros:System/scheduler:Scheduler":
                return new Scheduler(name, <any>undefined, { urn })
            case "routeros:System/systemIdentity:SystemIdentity":
                return new SystemIdentity(name, <any>undefined, { urn })
            case "routeros:System/systemScheduler:SystemScheduler":
                return new SystemScheduler(name, <any>undefined, { urn })
            case "routeros:System/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("routeros", "System/certificate", _module)
pulumi.runtime.registerResourceModule("routeros", "System/identity", _module)
pulumi.runtime.registerResourceModule("routeros", "System/scheduler", _module)
pulumi.runtime.registerResourceModule("routeros", "System/systemIdentity", _module)
pulumi.runtime.registerResourceModule("routeros", "System/systemScheduler", _module)
pulumi.runtime.registerResourceModule("routeros", "System/user", _module)