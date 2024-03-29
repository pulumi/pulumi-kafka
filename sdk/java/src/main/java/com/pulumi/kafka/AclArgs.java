// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kafka;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AclArgs extends com.pulumi.resources.ResourceArgs {

    public static final AclArgs Empty = new AclArgs();

    /**
     * Host from which principal listed in `acl_principal`
     * will have access.
     * 
     */
    @Import(name="aclHost", required=true)
    private Output<String> aclHost;

    /**
     * @return Host from which principal listed in `acl_principal`
     * will have access.
     * 
     */
    public Output<String> aclHost() {
        return this.aclHost;
    }

    /**
     * Operation that is being allowed or denied. Valid
     * values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
     * `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
     * 
     */
    @Import(name="aclOperation", required=true)
    private Output<String> aclOperation;

    /**
     * @return Operation that is being allowed or denied. Valid
     * values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
     * `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
     * 
     */
    public Output<String> aclOperation() {
        return this.aclOperation;
    }

    /**
     * Type of permission. Valid values are `Unknown`,
     * `Any`, `Allow`, `Deny`.
     * 
     */
    @Import(name="aclPermissionType", required=true)
    private Output<String> aclPermissionType;

    /**
     * @return Type of permission. Valid values are `Unknown`,
     * `Any`, `Allow`, `Deny`.
     * 
     */
    public Output<String> aclPermissionType() {
        return this.aclPermissionType;
    }

    /**
     * Principal that is being allowed or denied.
     * 
     */
    @Import(name="aclPrincipal", required=true)
    private Output<String> aclPrincipal;

    /**
     * @return Principal that is being allowed or denied.
     * 
     */
    public Output<String> aclPrincipal() {
        return this.aclPrincipal;
    }

    /**
     * The name of the resource.
     * 
     */
    @Import(name="aclResourceName", required=true)
    private Output<String> aclResourceName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> aclResourceName() {
        return this.aclResourceName;
    }

    /**
     * The type of resource. Valid values are `Unknown`,
     * `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
     * 
     */
    @Import(name="aclResourceType", required=true)
    private Output<String> aclResourceType;

    /**
     * @return The type of resource. Valid values are `Unknown`,
     * `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
     * 
     */
    public Output<String> aclResourceType() {
        return this.aclResourceType;
    }

    /**
     * The pattern filter. Valid values
     * are `Prefixed`, `Any`, `Match`, `Literal`. Default `Literal`.
     * 
     */
    @Import(name="resourcePatternTypeFilter")
    private @Nullable Output<String> resourcePatternTypeFilter;

    /**
     * @return The pattern filter. Valid values
     * are `Prefixed`, `Any`, `Match`, `Literal`. Default `Literal`.
     * 
     */
    public Optional<Output<String>> resourcePatternTypeFilter() {
        return Optional.ofNullable(this.resourcePatternTypeFilter);
    }

    private AclArgs() {}

    private AclArgs(AclArgs $) {
        this.aclHost = $.aclHost;
        this.aclOperation = $.aclOperation;
        this.aclPermissionType = $.aclPermissionType;
        this.aclPrincipal = $.aclPrincipal;
        this.aclResourceName = $.aclResourceName;
        this.aclResourceType = $.aclResourceType;
        this.resourcePatternTypeFilter = $.resourcePatternTypeFilter;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AclArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AclArgs $;

        public Builder() {
            $ = new AclArgs();
        }

        public Builder(AclArgs defaults) {
            $ = new AclArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclHost Host from which principal listed in `acl_principal`
         * will have access.
         * 
         * @return builder
         * 
         */
        public Builder aclHost(Output<String> aclHost) {
            $.aclHost = aclHost;
            return this;
        }

        /**
         * @param aclHost Host from which principal listed in `acl_principal`
         * will have access.
         * 
         * @return builder
         * 
         */
        public Builder aclHost(String aclHost) {
            return aclHost(Output.of(aclHost));
        }

        /**
         * @param aclOperation Operation that is being allowed or denied. Valid
         * values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
         * `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
         * 
         * @return builder
         * 
         */
        public Builder aclOperation(Output<String> aclOperation) {
            $.aclOperation = aclOperation;
            return this;
        }

        /**
         * @param aclOperation Operation that is being allowed or denied. Valid
         * values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
         * `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
         * 
         * @return builder
         * 
         */
        public Builder aclOperation(String aclOperation) {
            return aclOperation(Output.of(aclOperation));
        }

        /**
         * @param aclPermissionType Type of permission. Valid values are `Unknown`,
         * `Any`, `Allow`, `Deny`.
         * 
         * @return builder
         * 
         */
        public Builder aclPermissionType(Output<String> aclPermissionType) {
            $.aclPermissionType = aclPermissionType;
            return this;
        }

        /**
         * @param aclPermissionType Type of permission. Valid values are `Unknown`,
         * `Any`, `Allow`, `Deny`.
         * 
         * @return builder
         * 
         */
        public Builder aclPermissionType(String aclPermissionType) {
            return aclPermissionType(Output.of(aclPermissionType));
        }

        /**
         * @param aclPrincipal Principal that is being allowed or denied.
         * 
         * @return builder
         * 
         */
        public Builder aclPrincipal(Output<String> aclPrincipal) {
            $.aclPrincipal = aclPrincipal;
            return this;
        }

        /**
         * @param aclPrincipal Principal that is being allowed or denied.
         * 
         * @return builder
         * 
         */
        public Builder aclPrincipal(String aclPrincipal) {
            return aclPrincipal(Output.of(aclPrincipal));
        }

        /**
         * @param aclResourceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder aclResourceName(Output<String> aclResourceName) {
            $.aclResourceName = aclResourceName;
            return this;
        }

        /**
         * @param aclResourceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder aclResourceName(String aclResourceName) {
            return aclResourceName(Output.of(aclResourceName));
        }

        /**
         * @param aclResourceType The type of resource. Valid values are `Unknown`,
         * `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
         * 
         * @return builder
         * 
         */
        public Builder aclResourceType(Output<String> aclResourceType) {
            $.aclResourceType = aclResourceType;
            return this;
        }

        /**
         * @param aclResourceType The type of resource. Valid values are `Unknown`,
         * `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
         * 
         * @return builder
         * 
         */
        public Builder aclResourceType(String aclResourceType) {
            return aclResourceType(Output.of(aclResourceType));
        }

        /**
         * @param resourcePatternTypeFilter The pattern filter. Valid values
         * are `Prefixed`, `Any`, `Match`, `Literal`. Default `Literal`.
         * 
         * @return builder
         * 
         */
        public Builder resourcePatternTypeFilter(@Nullable Output<String> resourcePatternTypeFilter) {
            $.resourcePatternTypeFilter = resourcePatternTypeFilter;
            return this;
        }

        /**
         * @param resourcePatternTypeFilter The pattern filter. Valid values
         * are `Prefixed`, `Any`, `Match`, `Literal`. Default `Literal`.
         * 
         * @return builder
         * 
         */
        public Builder resourcePatternTypeFilter(String resourcePatternTypeFilter) {
            return resourcePatternTypeFilter(Output.of(resourcePatternTypeFilter));
        }

        public AclArgs build() {
            if ($.aclHost == null) {
                throw new MissingRequiredPropertyException("AclArgs", "aclHost");
            }
            if ($.aclOperation == null) {
                throw new MissingRequiredPropertyException("AclArgs", "aclOperation");
            }
            if ($.aclPermissionType == null) {
                throw new MissingRequiredPropertyException("AclArgs", "aclPermissionType");
            }
            if ($.aclPrincipal == null) {
                throw new MissingRequiredPropertyException("AclArgs", "aclPrincipal");
            }
            if ($.aclResourceName == null) {
                throw new MissingRequiredPropertyException("AclArgs", "aclResourceName");
            }
            if ($.aclResourceType == null) {
                throw new MissingRequiredPropertyException("AclArgs", "aclResourceType");
            }
            return $;
        }
    }

}
