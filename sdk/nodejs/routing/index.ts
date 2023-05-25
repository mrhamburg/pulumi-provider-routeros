// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { RoutingArgs, RoutingState } from "./routing";
export type Routing = import("./routing").Routing;
export const Routing: typeof import("./routing").Routing = null as any;
utilities.lazyLoad(exports, ["Routing"], () => require("./routing"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "routeros:Routing/routing:Routing":
                return new Routing(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("routeros", "Routing/routing", _module)