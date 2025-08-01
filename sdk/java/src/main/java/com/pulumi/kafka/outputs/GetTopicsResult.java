// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.kafka.outputs.GetTopicsList;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetTopicsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list containing all the topics.
     * 
     */
    private List<GetTopicsList> lists;

    private GetTopicsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list containing all the topics.
     * 
     */
    public List<GetTopicsList> lists() {
        return this.lists;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTopicsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<GetTopicsList> lists;
        public Builder() {}
        public Builder(GetTopicsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.lists = defaults.lists;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTopicsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lists(List<GetTopicsList> lists) {
            if (lists == null) {
              throw new MissingRequiredPropertyException("GetTopicsResult", "lists");
            }
            this.lists = lists;
            return this;
        }
        public Builder lists(GetTopicsList... lists) {
            return lists(List.of(lists));
        }
        public GetTopicsResult build() {
            final var _resultValue = new GetTopicsResult();
            _resultValue.id = id;
            _resultValue.lists = lists;
            return _resultValue;
        }
    }
}
