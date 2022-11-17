// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetTopicPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTopicPlainArgs Empty = new GetTopicPlainArgs();

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    private GetTopicPlainArgs() {}

    private GetTopicPlainArgs(GetTopicPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTopicPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTopicPlainArgs $;

        public Builder() {
            $ = new GetTopicPlainArgs();
        }

        public Builder(GetTopicPlainArgs defaults) {
            $ = new GetTopicPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetTopicPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
