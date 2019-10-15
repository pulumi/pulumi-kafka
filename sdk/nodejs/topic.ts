// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kafka:index/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * A map of string k/v attributes.
     */
    public readonly config!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the topic.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of partitions.
     */
    public readonly partitions!: pulumi.Output<number>;
    /**
     * Number of replicas.
     */
    public readonly replicationFactor!: pulumi.Output<number>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TopicState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partitions"] = state ? state.partitions : undefined;
            inputs["replicationFactor"] = state ? state.replicationFactor : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            if (!args || args.partitions === undefined) {
                throw new Error("Missing required property 'partitions'");
            }
            if (!args || args.replicationFactor === undefined) {
                throw new Error("Missing required property 'replicationFactor'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partitions"] = args ? args.partitions : undefined;
            inputs["replicationFactor"] = args ? args.replicationFactor : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * A map of string k/v attributes.
     */
    readonly config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the topic.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Number of partitions.
     */
    readonly partitions?: pulumi.Input<number>;
    /**
     * Number of replicas.
     */
    readonly replicationFactor?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * A map of string k/v attributes.
     */
    readonly config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the topic.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Number of partitions.
     */
    readonly partitions: pulumi.Input<number>;
    /**
     * Number of replicas.
     */
    readonly replicationFactor: pulumi.Input<number>;
}
