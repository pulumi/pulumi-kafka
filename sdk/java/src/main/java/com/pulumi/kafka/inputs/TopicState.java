// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TopicState extends com.pulumi.resources.ResourceArgs {

    public static final TopicState Empty = new TopicState();

    /**
     * A map of string k/v attributes.
     * 
     */
    @Import(name="config")
    private @Nullable Output<Map<String,Object>> config;

    /**
     * @return A map of string k/v attributes.
     * 
     */
    public Optional<Output<Map<String,Object>>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * The name of the topic.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the topic.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The number of partitions the topic should have.
     * 
     */
    @Import(name="partitions")
    private @Nullable Output<Integer> partitions;

    /**
     * @return The number of partitions the topic should have.
     * 
     */
    public Optional<Output<Integer>> partitions() {
        return Optional.ofNullable(this.partitions);
    }

    /**
     * The number of replicas the topic should have.
     * 
     */
    @Import(name="replicationFactor")
    private @Nullable Output<Integer> replicationFactor;

    /**
     * @return The number of replicas the topic should have.
     * 
     */
    public Optional<Output<Integer>> replicationFactor() {
        return Optional.ofNullable(this.replicationFactor);
    }

    private TopicState() {}

    private TopicState(TopicState $) {
        this.config = $.config;
        this.name = $.name;
        this.partitions = $.partitions;
        this.replicationFactor = $.replicationFactor;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TopicState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TopicState $;

        public Builder() {
            $ = new TopicState();
        }

        public Builder(TopicState defaults) {
            $ = new TopicState(Objects.requireNonNull(defaults));
        }

        /**
         * @param config A map of string k/v attributes.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<Map<String,Object>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config A map of string k/v attributes.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,Object> config) {
            return config(Output.of(config));
        }

        /**
         * @param name The name of the topic.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the topic.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partitions The number of partitions the topic should have.
         * 
         * @return builder
         * 
         */
        public Builder partitions(@Nullable Output<Integer> partitions) {
            $.partitions = partitions;
            return this;
        }

        /**
         * @param partitions The number of partitions the topic should have.
         * 
         * @return builder
         * 
         */
        public Builder partitions(Integer partitions) {
            return partitions(Output.of(partitions));
        }

        /**
         * @param replicationFactor The number of replicas the topic should have.
         * 
         * @return builder
         * 
         */
        public Builder replicationFactor(@Nullable Output<Integer> replicationFactor) {
            $.replicationFactor = replicationFactor;
            return this;
        }

        /**
         * @param replicationFactor The number of replicas the topic should have.
         * 
         * @return builder
         * 
         */
        public Builder replicationFactor(Integer replicationFactor) {
            return replicationFactor(Output.of(replicationFactor));
        }

        public TopicState build() {
            return $;
        }
    }

}
