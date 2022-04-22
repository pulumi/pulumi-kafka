// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Quota struct {
	pulumi.CustomResourceState

	// A map of string k/v properties.
	Config pulumi.MapOutput `pulumi:"config"`
	// The name of the entity
	EntityName pulumi.StringOutput `pulumi:"entityName"`
	// The type of the entity (client-id, user, ip)
	EntityType pulumi.StringOutput `pulumi:"entityType"`
}

// NewQuota registers a new resource with the given unique name, arguments, and options.
func NewQuota(ctx *pulumi.Context,
	name string, args *QuotaArgs, opts ...pulumi.ResourceOption) (*Quota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntityName == nil {
		return nil, errors.New("invalid value for required argument 'EntityName'")
	}
	if args.EntityType == nil {
		return nil, errors.New("invalid value for required argument 'EntityType'")
	}
	var resource Quota
	err := ctx.RegisterResource("kafka:index/quota:Quota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuota gets an existing Quota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaState, opts ...pulumi.ResourceOption) (*Quota, error) {
	var resource Quota
	err := ctx.ReadResource("kafka:index/quota:Quota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Quota resources.
type quotaState struct {
	// A map of string k/v properties.
	Config map[string]interface{} `pulumi:"config"`
	// The name of the entity
	EntityName *string `pulumi:"entityName"`
	// The type of the entity (client-id, user, ip)
	EntityType *string `pulumi:"entityType"`
}

type QuotaState struct {
	// A map of string k/v properties.
	Config pulumi.MapInput
	// The name of the entity
	EntityName pulumi.StringPtrInput
	// The type of the entity (client-id, user, ip)
	EntityType pulumi.StringPtrInput
}

func (QuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaState)(nil)).Elem()
}

type quotaArgs struct {
	// A map of string k/v properties.
	Config map[string]interface{} `pulumi:"config"`
	// The name of the entity
	EntityName string `pulumi:"entityName"`
	// The type of the entity (client-id, user, ip)
	EntityType string `pulumi:"entityType"`
}

// The set of arguments for constructing a Quota resource.
type QuotaArgs struct {
	// A map of string k/v properties.
	Config pulumi.MapInput
	// The name of the entity
	EntityName pulumi.StringInput
	// The type of the entity (client-id, user, ip)
	EntityType pulumi.StringInput
}

func (QuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaArgs)(nil)).Elem()
}

type QuotaInput interface {
	pulumi.Input

	ToQuotaOutput() QuotaOutput
	ToQuotaOutputWithContext(ctx context.Context) QuotaOutput
}

func (*Quota) ElementType() reflect.Type {
	return reflect.TypeOf((**Quota)(nil)).Elem()
}

func (i *Quota) ToQuotaOutput() QuotaOutput {
	return i.ToQuotaOutputWithContext(context.Background())
}

func (i *Quota) ToQuotaOutputWithContext(ctx context.Context) QuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaOutput)
}

// QuotaArrayInput is an input type that accepts QuotaArray and QuotaArrayOutput values.
// You can construct a concrete instance of `QuotaArrayInput` via:
//
//          QuotaArray{ QuotaArgs{...} }
type QuotaArrayInput interface {
	pulumi.Input

	ToQuotaArrayOutput() QuotaArrayOutput
	ToQuotaArrayOutputWithContext(context.Context) QuotaArrayOutput
}

type QuotaArray []QuotaInput

func (QuotaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Quota)(nil)).Elem()
}

func (i QuotaArray) ToQuotaArrayOutput() QuotaArrayOutput {
	return i.ToQuotaArrayOutputWithContext(context.Background())
}

func (i QuotaArray) ToQuotaArrayOutputWithContext(ctx context.Context) QuotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaArrayOutput)
}

// QuotaMapInput is an input type that accepts QuotaMap and QuotaMapOutput values.
// You can construct a concrete instance of `QuotaMapInput` via:
//
//          QuotaMap{ "key": QuotaArgs{...} }
type QuotaMapInput interface {
	pulumi.Input

	ToQuotaMapOutput() QuotaMapOutput
	ToQuotaMapOutputWithContext(context.Context) QuotaMapOutput
}

type QuotaMap map[string]QuotaInput

func (QuotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Quota)(nil)).Elem()
}

func (i QuotaMap) ToQuotaMapOutput() QuotaMapOutput {
	return i.ToQuotaMapOutputWithContext(context.Background())
}

func (i QuotaMap) ToQuotaMapOutputWithContext(ctx context.Context) QuotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaMapOutput)
}

type QuotaOutput struct{ *pulumi.OutputState }

func (QuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Quota)(nil)).Elem()
}

func (o QuotaOutput) ToQuotaOutput() QuotaOutput {
	return o
}

func (o QuotaOutput) ToQuotaOutputWithContext(ctx context.Context) QuotaOutput {
	return o
}

type QuotaArrayOutput struct{ *pulumi.OutputState }

func (QuotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Quota)(nil)).Elem()
}

func (o QuotaArrayOutput) ToQuotaArrayOutput() QuotaArrayOutput {
	return o
}

func (o QuotaArrayOutput) ToQuotaArrayOutputWithContext(ctx context.Context) QuotaArrayOutput {
	return o
}

func (o QuotaArrayOutput) Index(i pulumi.IntInput) QuotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Quota {
		return vs[0].([]*Quota)[vs[1].(int)]
	}).(QuotaOutput)
}

type QuotaMapOutput struct{ *pulumi.OutputState }

func (QuotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Quota)(nil)).Elem()
}

func (o QuotaMapOutput) ToQuotaMapOutput() QuotaMapOutput {
	return o
}

func (o QuotaMapOutput) ToQuotaMapOutputWithContext(ctx context.Context) QuotaMapOutput {
	return o
}

func (o QuotaMapOutput) MapIndex(k pulumi.StringInput) QuotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Quota {
		return vs[0].(map[string]*Quota)[vs[1].(string)]
	}).(QuotaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaInput)(nil)).Elem(), &Quota{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaArrayInput)(nil)).Elem(), QuotaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaMapInput)(nil)).Elem(), QuotaMap{})
	pulumi.RegisterOutputType(QuotaOutput{})
	pulumi.RegisterOutputType(QuotaArrayOutput{})
	pulumi.RegisterOutputType(QuotaMapOutput{})
}
