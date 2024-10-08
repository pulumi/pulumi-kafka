// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetTopicResult {
    private Map<String,String> config;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private Integer partitions;
    private Integer replicationFactor;

    private GetTopicResult() {}
    public Map<String,String> config() {
        return this.config;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public Integer partitions() {
        return this.partitions;
    }
    public Integer replicationFactor() {
        return this.replicationFactor;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTopicResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> config;
        private String id;
        private String name;
        private Integer partitions;
        private Integer replicationFactor;
        public Builder() {}
        public Builder(GetTopicResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.config = defaults.config;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.partitions = defaults.partitions;
    	      this.replicationFactor = defaults.replicationFactor;
        }

        @CustomType.Setter
        public Builder config(Map<String,String> config) {
            if (config == null) {
              throw new MissingRequiredPropertyException("GetTopicResult", "config");
            }
            this.config = config;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTopicResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetTopicResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder partitions(Integer partitions) {
            if (partitions == null) {
              throw new MissingRequiredPropertyException("GetTopicResult", "partitions");
            }
            this.partitions = partitions;
            return this;
        }
        @CustomType.Setter
        public Builder replicationFactor(Integer replicationFactor) {
            if (replicationFactor == null) {
              throw new MissingRequiredPropertyException("GetTopicResult", "replicationFactor");
            }
            this.replicationFactor = replicationFactor;
            return this;
        }
        public GetTopicResult build() {
            final var _resultValue = new GetTopicResult();
            _resultValue.config = config;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.partitions = partitions;
            _resultValue.replicationFactor = replicationFactor;
            return _resultValue;
        }
    }
}
