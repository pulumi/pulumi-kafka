// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./acl";
export * from "./provider";
export * from "./topic";

// Export sub-modules:
import * as config from "./config";

export {
    config,
};

// Import resources to register:
import { Acl } from "./acl";
import { Topic } from "./topic";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kafka:index/acl:Acl":
                return new Acl(name, <any>undefined, { urn })
            case "kafka:index/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kafka", "index/acl", _module)
pulumi.runtime.registerResourceModule("kafka", "index/topic", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("kafka", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:kafka") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
