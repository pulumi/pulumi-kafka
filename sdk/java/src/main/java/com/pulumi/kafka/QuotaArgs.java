// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaArgs extends com.pulumi.resources.ResourceArgs {

    public static final QuotaArgs Empty = new QuotaArgs();

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
     * The name of the entity to target.
     * 
     */
    @Import(name="entityName", required=true)
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
    @Import(name="entityType", required=true)
    private Output<String> entityType;

    /**
     * @return The type of entity. Valid values are `client-id`, `user`, `ip`.
     * 
     */
    public Output<String> entityType() {
        return this.entityType;
    }

    private QuotaArgs() {}

    private QuotaArgs(QuotaArgs $) {
        this.config = $.config;
        this.entityName = $.entityName;
        this.entityType = $.entityType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaArgs $;

        public Builder() {
            $ = new QuotaArgs();
        }

        public Builder(QuotaArgs defaults) {
            $ = new QuotaArgs(Objects.requireNonNull(defaults));
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
         * @param entityName The name of the entity to target.
         * 
         * @return builder
         * 
         */
        public Builder entityName(Output<String> entityName) {
            $.entityName = entityName;
            return this;
        }

        /**
         * @param entityName The name of the entity to target.
         * 
         * @return builder
         * 
         */
        public Builder entityName(String entityName) {
            return entityName(Output.of(entityName));
        }

        /**
         * @param entityType The type of entity. Valid values are `client-id`, `user`, `ip`.
         * 
         * @return builder
         * 
         */
        public Builder entityType(Output<String> entityType) {
            $.entityType = entityType;
            return this;
        }

        /**
         * @param entityType The type of entity. Valid values are `client-id`, `user`, `ip`.
         * 
         * @return builder
         * 
         */
        public Builder entityType(String entityType) {
            return entityType(Output.of(entityType));
        }

        public QuotaArgs build() {
            $.entityName = Objects.requireNonNull($.entityName, "expected parameter 'entityName' to be non-null");
            $.entityType = Objects.requireNonNull($.entityType, "expected parameter 'entityType' to be non-null");
            return $;
        }
    }

}