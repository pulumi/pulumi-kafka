// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kafka.QuotaArgs;
import com.pulumi.kafka.Utilities;
import com.pulumi.kafka.inputs.QuotaState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="kafka:index/quota:Quota")
public class Quota extends com.pulumi.resources.CustomResource {
    /**
     * A map of string k/v properties.
     * 
     */
    @Export(name="config", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> config;

    /**
     * @return A map of string k/v properties.
     * 
     */
    public Output<Optional<Map<String,Object>>> config() {
        return Codegen.optional(this.config);
    }
    /**
     * The name of the entity
     * 
     */
    @Export(name="entityName", type=String.class, parameters={})
    private Output<String> entityName;

    /**
     * @return The name of the entity
     * 
     */
    public Output<String> entityName() {
        return this.entityName;
    }
    /**
     * The type of the entity (client-id, user, ip)
     * 
     */
    @Export(name="entityType", type=String.class, parameters={})
    private Output<String> entityType;

    /**
     * @return The type of the entity (client-id, user, ip)
     * 
     */
    public Output<String> entityType() {
        return this.entityType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Quota(String name) {
        this(name, QuotaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Quota(String name, QuotaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Quota(String name, QuotaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kafka:index/quota:Quota", name, args == null ? QuotaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Quota(String name, Output<String> id, @Nullable QuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kafka:index/quota:Quota", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Quota get(String name, Output<String> id, @Nullable QuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Quota(name, id, state, options);
    }
}
