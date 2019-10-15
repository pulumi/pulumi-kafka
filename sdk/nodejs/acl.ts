// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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

    public readonly aclHost!: pulumi.Output<string>;
    public readonly aclOperation!: pulumi.Output<string>;
    public readonly aclPermissionType!: pulumi.Output<string>;
    public readonly aclPrincipal!: pulumi.Output<string>;
    /**
     * The name of the resouce
     */
    public readonly resourceName!: pulumi.Output<string>;
    public readonly resourcePatternTypeFilter!: pulumi.Output<string | undefined>;
    public readonly resourceType!: pulumi.Output<string>;

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
            inputs["resourceName"] = state ? state.resourceName : undefined;
            inputs["resourcePatternTypeFilter"] = state ? state.resourcePatternTypeFilter : undefined;
            inputs["resourceType"] = state ? state.resourceType : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if (!args || args.aclHost === undefined) {
                throw new Error("Missing required property 'aclHost'");
            }
            if (!args || args.aclOperation === undefined) {
                throw new Error("Missing required property 'aclOperation'");
            }
            if (!args || args.aclPermissionType === undefined) {
                throw new Error("Missing required property 'aclPermissionType'");
            }
            if (!args || args.aclPrincipal === undefined) {
                throw new Error("Missing required property 'aclPrincipal'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            if (!args || args.resourceType === undefined) {
                throw new Error("Missing required property 'resourceType'");
            }
            inputs["aclHost"] = args ? args.aclHost : undefined;
            inputs["aclOperation"] = args ? args.aclOperation : undefined;
            inputs["aclPermissionType"] = args ? args.aclPermissionType : undefined;
            inputs["aclPrincipal"] = args ? args.aclPrincipal : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["resourcePatternTypeFilter"] = args ? args.resourcePatternTypeFilter : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
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
    readonly aclHost?: pulumi.Input<string>;
    readonly aclOperation?: pulumi.Input<string>;
    readonly aclPermissionType?: pulumi.Input<string>;
    readonly aclPrincipal?: pulumi.Input<string>;
    /**
     * The name of the resouce
     */
    readonly resourceName?: pulumi.Input<string>;
    readonly resourcePatternTypeFilter?: pulumi.Input<string>;
    readonly resourceType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    readonly aclHost: pulumi.Input<string>;
    readonly aclOperation: pulumi.Input<string>;
    readonly aclPermissionType: pulumi.Input<string>;
    readonly aclPrincipal: pulumi.Input<string>;
    /**
     * The name of the resouce
     */
    readonly resourceName: pulumi.Input<string>;
    readonly resourcePatternTypeFilter?: pulumi.Input<string>;
    readonly resourceType: pulumi.Input<string>;
}
