// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kafka.TopicArgs;
import com.pulumi.kafka.Utilities;
import com.pulumi.kafka.inputs.TopicState;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A resource for managing Kafka topics. Increases partition count without destroying the topic.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kafka.Topic;
 * import com.pulumi.kafka.TopicArgs;
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
 *         var logs = new Topic(&#34;logs&#34;, TopicArgs.builder()        
 *             .name(&#34;systemd_logs&#34;)
 *             .replicationFactor(2)
 *             .partitions(100)
 *             .config(Map.ofEntries(
 *                 Map.entry(&#34;segment.ms&#34;, &#34;20000&#34;),
 *                 Map.entry(&#34;cleanup.policy&#34;, &#34;compact&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Topics can be imported using their ARN, e.g.
 * 
 * ```sh
 * $ pulumi import kafka:index/topic:Topic logs systemd_logs
 * ```
 * 
 */
@ResourceType(type="kafka:index/topic:Topic")
public class Topic extends com.pulumi.resources.CustomResource {
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
     * The name of the topic.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the topic.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The number of partitions the topic should have.
     * 
     */
    @Export(name="partitions", refs={Integer.class}, tree="[0]")
    private Output<Integer> partitions;

    /**
     * @return The number of partitions the topic should have.
     * 
     */
    public Output<Integer> partitions() {
        return this.partitions;
    }
    /**
     * The number of replicas the topic should have.
     * 
     */
    @Export(name="replicationFactor", refs={Integer.class}, tree="[0]")
    private Output<Integer> replicationFactor;

    /**
     * @return The number of replicas the topic should have.
     * 
     */
    public Output<Integer> replicationFactor() {
        return this.replicationFactor;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Topic(String name) {
        this(name, TopicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Topic(String name, TopicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Topic(String name, TopicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kafka:index/topic:Topic", name, args == null ? TopicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Topic(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kafka:index/topic:Topic", name, state, makeResourceOptions(options, id));
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
    public static Topic get(String name, Output<String> id, @Nullable TopicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Topic(name, id, state, options);
    }
}
