// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kafka

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Acl struct {
	pulumi.CustomResourceState

	AclHost           pulumi.StringOutput `pulumi:"aclHost"`
	AclOperation      pulumi.StringOutput `pulumi:"aclOperation"`
	AclPermissionType pulumi.StringOutput `pulumi:"aclPermissionType"`
	AclPrincipal      pulumi.StringOutput `pulumi:"aclPrincipal"`
	// The name of the resouce
	AclResourceName           pulumi.StringOutput    `pulumi:"aclResourceName"`
	AclResourceType           pulumi.StringOutput    `pulumi:"aclResourceType"`
	ResourcePatternTypeFilter pulumi.StringPtrOutput `pulumi:"resourcePatternTypeFilter"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil || args.AclHost == nil {
		return nil, errors.New("missing required argument 'AclHost'")
	}
	if args == nil || args.AclOperation == nil {
		return nil, errors.New("missing required argument 'AclOperation'")
	}
	if args == nil || args.AclPermissionType == nil {
		return nil, errors.New("missing required argument 'AclPermissionType'")
	}
	if args == nil || args.AclPrincipal == nil {
		return nil, errors.New("missing required argument 'AclPrincipal'")
	}
	if args == nil || args.AclResourceName == nil {
		return nil, errors.New("missing required argument 'AclResourceName'")
	}
	if args == nil || args.AclResourceType == nil {
		return nil, errors.New("missing required argument 'AclResourceType'")
	}
	if args == nil {
		args = &AclArgs{}
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
	AclHost           *string `pulumi:"aclHost"`
	AclOperation      *string `pulumi:"aclOperation"`
	AclPermissionType *string `pulumi:"aclPermissionType"`
	AclPrincipal      *string `pulumi:"aclPrincipal"`
	// The name of the resouce
	AclResourceName           *string `pulumi:"aclResourceName"`
	AclResourceType           *string `pulumi:"aclResourceType"`
	ResourcePatternTypeFilter *string `pulumi:"resourcePatternTypeFilter"`
}

type AclState struct {
	AclHost           pulumi.StringPtrInput
	AclOperation      pulumi.StringPtrInput
	AclPermissionType pulumi.StringPtrInput
	AclPrincipal      pulumi.StringPtrInput
	// The name of the resouce
	AclResourceName           pulumi.StringPtrInput
	AclResourceType           pulumi.StringPtrInput
	ResourcePatternTypeFilter pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	AclHost           string `pulumi:"aclHost"`
	AclOperation      string `pulumi:"aclOperation"`
	AclPermissionType string `pulumi:"aclPermissionType"`
	AclPrincipal      string `pulumi:"aclPrincipal"`
	// The name of the resouce
	AclResourceName           string  `pulumi:"aclResourceName"`
	AclResourceType           string  `pulumi:"aclResourceType"`
	ResourcePatternTypeFilter *string `pulumi:"resourcePatternTypeFilter"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	AclHost           pulumi.StringInput
	AclOperation      pulumi.StringInput
	AclPermissionType pulumi.StringInput
	AclPrincipal      pulumi.StringInput
	// The name of the resouce
	AclResourceName           pulumi.StringInput
	AclResourceType           pulumi.StringInput
	ResourcePatternTypeFilter pulumi.StringPtrInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}
