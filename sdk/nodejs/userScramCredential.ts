// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class UserScramCredential extends pulumi.CustomResource {
    /**
     * Get an existing UserScramCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserScramCredentialState, opts?: pulumi.CustomResourceOptions): UserScramCredential {
        return new UserScramCredential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kafka:index/userScramCredential:UserScramCredential';

    /**
     * Returns true if the given object is an instance of UserScramCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserScramCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserScramCredential.__pulumiType;
    }

    /**
     * The password of the credential
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The number of SCRAM iterations used when generating the credential
     */
    public readonly scramIterations!: pulumi.Output<number | undefined>;
    /**
     * The SCRAM mechanism used to generate the credential (SCRAM-SHA-256, SCRAM-SHA-512)
     */
    public readonly scramMechanism!: pulumi.Output<string>;
    /**
     * The name of the credential
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a UserScramCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserScramCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserScramCredentialArgs | UserScramCredentialState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserScramCredentialState | undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["scramIterations"] = state ? state.scramIterations : undefined;
            resourceInputs["scramMechanism"] = state ? state.scramMechanism : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserScramCredentialArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.scramMechanism === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scramMechanism'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["scramIterations"] = args ? args.scramIterations : undefined;
            resourceInputs["scramMechanism"] = args ? args.scramMechanism : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(UserScramCredential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserScramCredential resources.
 */
export interface UserScramCredentialState {
    /**
     * The password of the credential
     */
    password?: pulumi.Input<string>;
    /**
     * The number of SCRAM iterations used when generating the credential
     */
    scramIterations?: pulumi.Input<number>;
    /**
     * The SCRAM mechanism used to generate the credential (SCRAM-SHA-256, SCRAM-SHA-512)
     */
    scramMechanism?: pulumi.Input<string>;
    /**
     * The name of the credential
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserScramCredential resource.
 */
export interface UserScramCredentialArgs {
    /**
     * The password of the credential
     */
    password: pulumi.Input<string>;
    /**
     * The number of SCRAM iterations used when generating the credential
     */
    scramIterations?: pulumi.Input<number>;
    /**
     * The SCRAM mechanism used to generate the credential (SCRAM-SHA-256, SCRAM-SHA-512)
     */
    scramMechanism: pulumi.Input<string>;
    /**
     * The name of the credential
     */
    username: pulumi.Input<string>;
}