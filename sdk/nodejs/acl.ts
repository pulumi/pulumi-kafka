// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * A resource for managing Kafka ACLs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as kafka from "@pulumi/kafka";
 *
 * const test = new kafka.Acl("test", {
 *     aclResourceName: "syslog",
 *     aclResourceType: "Topic",
 *     aclPrincipal: "User:Alice",
 *     aclHost: "*",
 *     aclOperation: "Write",
 *     aclPermissionType: "Deny",
 * });
 * ```
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kafka:index/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * Host from which principal listed in `aclPrincipal`
     * will have access.
     */
    public readonly aclHost!: pulumi.Output<string>;
    /**
     * Operation that is being allowed or denied. Valid
     * values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
     * `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
     */
    public readonly aclOperation!: pulumi.Output<string>;
    /**
     * Type of permission. Valid values are `Unknown`,
     * `Any`, `Allow`, `Deny`.
     */
    public readonly aclPermissionType!: pulumi.Output<string>;
    /**
     * Principal that is being allowed or denied.
     */
    public readonly aclPrincipal!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public readonly aclResourceName!: pulumi.Output<string>;
    /**
     * The type of resource. Valid values are `Unknown`,
     * `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
     */
    public readonly aclResourceType!: pulumi.Output<string>;
    /**
     * The pattern filter. Valid values
     * are `Prefixed`, `Any`, `Match`, `Literal`.
     */
    public readonly resourcePatternTypeFilter!: pulumi.Output<string | undefined>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AclState | undefined;
            inputs["aclHost"] = state ? state.aclHost : undefined;
            inputs["aclOperation"] = state ? state.aclOperation : undefined;
            inputs["aclPermissionType"] = state ? state.aclPermissionType : undefined;
            inputs["aclPrincipal"] = state ? state.aclPrincipal : undefined;
            inputs["aclResourceName"] = state ? state.aclResourceName : undefined;
            inputs["aclResourceType"] = state ? state.aclResourceType : undefined;
            inputs["resourcePatternTypeFilter"] = state ? state.resourcePatternTypeFilter : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if ((!args || args.aclHost === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclHost'");
            }
            if ((!args || args.aclOperation === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclOperation'");
            }
            if ((!args || args.aclPermissionType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclPermissionType'");
            }
            if ((!args || args.aclPrincipal === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclPrincipal'");
            }
            if ((!args || args.aclResourceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclResourceName'");
            }
            if ((!args || args.aclResourceType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclResourceType'");
            }
            inputs["aclHost"] = args ? args.aclHost : undefined;
            inputs["aclOperation"] = args ? args.aclOperation : undefined;
            inputs["aclPermissionType"] = args ? args.aclPermissionType : undefined;
            inputs["aclPrincipal"] = args ? args.aclPrincipal : undefined;
            inputs["aclResourceName"] = args ? args.aclResourceName : undefined;
            inputs["aclResourceType"] = args ? args.aclResourceType : undefined;
            inputs["resourcePatternTypeFilter"] = args ? args.resourcePatternTypeFilter : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Acl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * Host from which principal listed in `aclPrincipal`
     * will have access.
     */
    readonly aclHost?: pulumi.Input<string>;
    /**
     * Operation that is being allowed or denied. Valid
     * values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
     * `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
     */
    readonly aclOperation?: pulumi.Input<string>;
    /**
     * Type of permission. Valid values are `Unknown`,
     * `Any`, `Allow`, `Deny`.
     */
    readonly aclPermissionType?: pulumi.Input<string>;
    /**
     * Principal that is being allowed or denied.
     */
    readonly aclPrincipal?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    readonly aclResourceName?: pulumi.Input<string>;
    /**
     * The type of resource. Valid values are `Unknown`,
     * `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
     */
    readonly aclResourceType?: pulumi.Input<string>;
    /**
     * The pattern filter. Valid values
     * are `Prefixed`, `Any`, `Match`, `Literal`.
     */
    readonly resourcePatternTypeFilter?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * Host from which principal listed in `aclPrincipal`
     * will have access.
     */
    readonly aclHost: pulumi.Input<string>;
    /**
     * Operation that is being allowed or denied. Valid
     * values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
     * `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
     */
    readonly aclOperation: pulumi.Input<string>;
    /**
     * Type of permission. Valid values are `Unknown`,
     * `Any`, `Allow`, `Deny`.
     */
    readonly aclPermissionType: pulumi.Input<string>;
    /**
     * Principal that is being allowed or denied.
     */
    readonly aclPrincipal: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    readonly aclResourceName: pulumi.Input<string>;
    /**
     * The type of resource. Valid values are `Unknown`,
     * `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
     */
    readonly aclResourceType: pulumi.Input<string>;
    /**
     * The pattern filter. Valid values
     * are `Prefixed`, `Any`, `Match`, `Literal`.
     */
    readonly resourcePatternTypeFilter?: pulumi.Input<string>;
}
