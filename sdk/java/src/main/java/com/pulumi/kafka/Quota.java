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

/**
 * A resource for managing Kafka quotas.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kafka.Quota;
 * import com.pulumi.kafka.QuotaArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var quota = new Quota("quota", QuotaArgs.builder()
 *             .entityName("app_consumer")
 *             .entityType("client-id")
 *             .config(Map.ofEntries(
 *                 Map.entry("consumer_byte_rate", "5000000"),
 *                 Map.entry("producer_byte_rate", "2500000")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="kafka:index/quota:Quota")
public class Quota extends com.pulumi.resources.CustomResource {
    /**
     * A map of string k/v attributes.
     * 
     */
    @Export(name="config", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> config;

    /**
     * @return A map of string k/v attributes.
     * 
     */
    public Output<Optional<Map<String,Object>>> config() {
        return Codegen.optional(this.config);
    }
    /**
     * The name of the entity to target.
     * 
     */
    @Export(name="entityName", refs={String.class}, tree="[0]")
    private Output<String> entityName;

    /**
     * @return The name of the entity to target.
     * 
     */
    public Output<String> entityName() {
        return this.entityName;
    }
    /**
     * The type of entity. Valid values are `client-id`, `user`, `ip`.
     * 
     */
    @Export(name="entityType", refs={String.class}, tree="[0]")
    private Output<String> entityType;

    /**
     * @return The type of entity. Valid values are `client-id`, `user`, `ip`.
     * 
     */
    public Output<String> entityType() {
        return this.entityType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Quota(java.lang.String name) {
        this(name, QuotaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Quota(java.lang.String name, QuotaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Quota(java.lang.String name, QuotaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kafka:index/quota:Quota", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Quota(java.lang.String name, Output<java.lang.String> id, @Nullable QuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kafka:index/quota:Quota", name, state, makeResourceOptions(options, id), false);
    }

    private static QuotaArgs makeArgs(QuotaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? QuotaArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Quota get(java.lang.String name, Output<java.lang.String> id, @Nullable QuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Quota(name, id, state, options);
    }
}
