// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("kafka:index:getTopic", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopic.
 */
export interface GetTopicArgs {
    name: string;
}

/**
 * A collection of values returned by getTopic.
 */
export interface GetTopicResult {
    readonly config: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly partitions: number;
    readonly replicationFactor: number;
}
export function getTopicOutput(args: GetTopicOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicResult> {
    return pulumi.output(args).apply((a: any) => getTopic(a, opts))
}

/**
 * A collection of arguments for invoking getTopic.
 */
export interface GetTopicOutputArgs {
    name: pulumi.Input<string>;
}
