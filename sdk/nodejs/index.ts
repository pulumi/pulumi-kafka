// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { AclArgs, AclState } from "./acl";
export type Acl = import("./acl").Acl;
export const Acl: typeof import("./acl").Acl = null as any;
utilities.lazyLoad(exports, ["Acl"], () => require("./acl"));

export { GetTopicArgs, GetTopicResult, GetTopicOutputArgs } from "./getTopic";
export const getTopic: typeof import("./getTopic").getTopic = null as any;
export const getTopicOutput: typeof import("./getTopic").getTopicOutput = null as any;
utilities.lazyLoad(exports, ["getTopic","getTopicOutput"], () => require("./getTopic"));

export { GetTopicsResult } from "./getTopics";
export const getTopics: typeof import("./getTopics").getTopics = null as any;
export const getTopicsOutput: typeof import("./getTopics").getTopicsOutput = null as any;
utilities.lazyLoad(exports, ["getTopics","getTopicsOutput"], () => require("./getTopics"));

export * from "./provider";
import { Provider } from "./provider";

export { QuotaArgs, QuotaState } from "./quota";
export type Quota = import("./quota").Quota;
export const Quota: typeof import("./quota").Quota = null as any;
utilities.lazyLoad(exports, ["Quota"], () => require("./quota"));

export { TopicArgs, TopicState } from "./topic";
export type Topic = import("./topic").Topic;
export const Topic: typeof import("./topic").Topic = null as any;
utilities.lazyLoad(exports, ["Topic"], () => require("./topic"));

export { UserScramCredentialArgs, UserScramCredentialState } from "./userScramCredential";
export type UserScramCredential = import("./userScramCredential").UserScramCredential;
export const UserScramCredential: typeof import("./userScramCredential").UserScramCredential = null as any;
utilities.lazyLoad(exports, ["UserScramCredential"], () => require("./userScramCredential"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kafka:index/acl:Acl":
                return new Acl(name, <any>undefined, { urn })
            case "kafka:index/quota:Quota":
                return new Quota(name, <any>undefined, { urn })
            case "kafka:index/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            case "kafka:index/userScramCredential:UserScramCredential":
                return new UserScramCredential(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kafka", "index/acl", _module)
pulumi.runtime.registerResourceModule("kafka", "index/quota", _module)
pulumi.runtime.registerResourceModule("kafka", "index/topic", _module)
pulumi.runtime.registerResourceModule("kafka", "index/userScramCredential", _module)
pulumi.runtime.registerResourcePackage("kafka", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:kafka") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
