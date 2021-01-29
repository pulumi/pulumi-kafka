// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A resource for managing Kafka ACLs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-kafka/sdk/v2/go/kafka/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kafka.NewAcl(ctx, "test", &kafka.AclArgs{
// 			AclResourceName:   pulumi.String("syslog"),
// 			AclResourceType:   pulumi.String("Topic"),
// 			AclPrincipal:      pulumi.String("User:Alice"),
// 			AclHost:           pulumi.String("*"),
// 			AclOperation:      pulumi.String("Write"),
// 			AclPermissionType: pulumi.String("Deny"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Acl struct {
	pulumi.CustomResourceState

	// Host from which principal listed in `aclPrincipal`
	// will have access.
	AclHost pulumi.StringOutput `pulumi:"aclHost"`
	// Operation that is being allowed or denied. Valid
	// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
	// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
	AclOperation pulumi.StringOutput `pulumi:"aclOperation"`
	// Type of permission. Valid values are `Unknown`,
	// `Any`, `Allow`, `Deny`.
	AclPermissionType pulumi.StringOutput `pulumi:"aclPermissionType"`
	// Principal that is being allowed or denied.
	AclPrincipal pulumi.StringOutput `pulumi:"aclPrincipal"`
	// The name of the resource.
	AclResourceName pulumi.StringOutput `pulumi:"aclResourceName"`
	// The type of resource. Valid values are `Unknown`,
	// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
	AclResourceType pulumi.StringOutput `pulumi:"aclResourceType"`
	// The pattern filter. Valid values
	// are `Prefixed`, `Any`, `Match`, `Literal`.
	ResourcePatternTypeFilter pulumi.StringPtrOutput `pulumi:"resourcePatternTypeFilter"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclHost == nil {
		return nil, errors.New("invalid value for required argument 'AclHost'")
	}
	if args.AclOperation == nil {
		return nil, errors.New("invalid value for required argument 'AclOperation'")
	}
	if args.AclPermissionType == nil {
		return nil, errors.New("invalid value for required argument 'AclPermissionType'")
	}
	if args.AclPrincipal == nil {
		return nil, errors.New("invalid value for required argument 'AclPrincipal'")
	}
	if args.AclResourceName == nil {
		return nil, errors.New("invalid value for required argument 'AclResourceName'")
	}
	if args.AclResourceType == nil {
		return nil, errors.New("invalid value for required argument 'AclResourceType'")
	}
	var resource Acl
	err := ctx.RegisterResource("kafka:index/acl:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("kafka:index/acl:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
	// Host from which principal listed in `aclPrincipal`
	// will have access.
	AclHost *string `pulumi:"aclHost"`
	// Operation that is being allowed or denied. Valid
	// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
	// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
	AclOperation *string `pulumi:"aclOperation"`
	// Type of permission. Valid values are `Unknown`,
	// `Any`, `Allow`, `Deny`.
	AclPermissionType *string `pulumi:"aclPermissionType"`
	// Principal that is being allowed or denied.
	AclPrincipal *string `pulumi:"aclPrincipal"`
	// The name of the resource.
	AclResourceName *string `pulumi:"aclResourceName"`
	// The type of resource. Valid values are `Unknown`,
	// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
	AclResourceType *string `pulumi:"aclResourceType"`
	// The pattern filter. Valid values
	// are `Prefixed`, `Any`, `Match`, `Literal`.
	ResourcePatternTypeFilter *string `pulumi:"resourcePatternTypeFilter"`
}

type AclState struct {
	// Host from which principal listed in `aclPrincipal`
	// will have access.
	AclHost pulumi.StringPtrInput
	// Operation that is being allowed or denied. Valid
	// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
	// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
	AclOperation pulumi.StringPtrInput
	// Type of permission. Valid values are `Unknown`,
	// `Any`, `Allow`, `Deny`.
	AclPermissionType pulumi.StringPtrInput
	// Principal that is being allowed or denied.
	AclPrincipal pulumi.StringPtrInput
	// The name of the resource.
	AclResourceName pulumi.StringPtrInput
	// The type of resource. Valid values are `Unknown`,
	// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
	AclResourceType pulumi.StringPtrInput
	// The pattern filter. Valid values
	// are `Prefixed`, `Any`, `Match`, `Literal`.
	ResourcePatternTypeFilter pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// Host from which principal listed in `aclPrincipal`
	// will have access.
	AclHost string `pulumi:"aclHost"`
	// Operation that is being allowed or denied. Valid
	// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
	// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
	AclOperation string `pulumi:"aclOperation"`
	// Type of permission. Valid values are `Unknown`,
	// `Any`, `Allow`, `Deny`.
	AclPermissionType string `pulumi:"aclPermissionType"`
	// Principal that is being allowed or denied.
	AclPrincipal string `pulumi:"aclPrincipal"`
	// The name of the resource.
	AclResourceName string `pulumi:"aclResourceName"`
	// The type of resource. Valid values are `Unknown`,
	// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
	AclResourceType string `pulumi:"aclResourceType"`
	// The pattern filter. Valid values
	// are `Prefixed`, `Any`, `Match`, `Literal`.
	ResourcePatternTypeFilter *string `pulumi:"resourcePatternTypeFilter"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	// Host from which principal listed in `aclPrincipal`
	// will have access.
	AclHost pulumi.StringInput
	// Operation that is being allowed or denied. Valid
	// values are `Unknown`, `Any`, `All`, `Read`, `Write`, `Create`, `Delete`, `Alter`,
	// `Describe`, `ClusterAction`, `DescribeConfigs`, `AlterConfigs`, `IdempotentWrite`.
	AclOperation pulumi.StringInput
	// Type of permission. Valid values are `Unknown`,
	// `Any`, `Allow`, `Deny`.
	AclPermissionType pulumi.StringInput
	// Principal that is being allowed or denied.
	AclPrincipal pulumi.StringInput
	// The name of the resource.
	AclResourceName pulumi.StringInput
	// The type of resource. Valid values are `Unknown`,
	// `Any`, `Topic`, `Group`, `Cluster`, `TransactionalID`.
	AclResourceType pulumi.StringInput
	// The pattern filter. Valid values
	// are `Prefixed`, `Any`, `Match`, `Literal`.
	ResourcePatternTypeFilter pulumi.StringPtrInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((*Acl)(nil))
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

type AclOutput struct {
	*pulumi.OutputState
}

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Acl)(nil))
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AclOutput{})
}
