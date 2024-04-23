// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-kafka/sdk/v3/go/kafka/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A resource for managing Kafka topics. Increases partition count without destroying the topic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kafka/sdk/v3/go/kafka"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kafka.NewTopic(ctx, "logs", &kafka.TopicArgs{
//				Name:              pulumi.String("systemd_logs"),
//				ReplicationFactor: pulumi.Int(2),
//				Partitions:        pulumi.Int(100),
//				Config: pulumi.Map{
//					"segment.ms":     pulumi.Any("20000"),
//					"cleanup.policy": pulumi.Any("compact"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Topics can be imported using their ARN, e.g.
//
// ```sh
// $ pulumi import kafka:index/topic:Topic logs systemd_logs
// ```
type Topic struct {
	pulumi.CustomResourceState

	// A map of string k/v attributes.
	Config pulumi.MapOutput `pulumi:"config"`
	// The name of the topic.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of partitions the topic should have.
	Partitions pulumi.IntOutput `pulumi:"partitions"`
	// The number of replicas the topic should have.
	ReplicationFactor pulumi.IntOutput `pulumi:"replicationFactor"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Partitions == nil {
		return nil, errors.New("invalid value for required argument 'Partitions'")
	}
	if args.ReplicationFactor == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationFactor'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("kafka:index/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("kafka:index/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// A map of string k/v attributes.
	Config map[string]interface{} `pulumi:"config"`
	// The name of the topic.
	Name *string `pulumi:"name"`
	// The number of partitions the topic should have.
	Partitions *int `pulumi:"partitions"`
	// The number of replicas the topic should have.
	ReplicationFactor *int `pulumi:"replicationFactor"`
}

type TopicState struct {
	// A map of string k/v attributes.
	Config pulumi.MapInput
	// The name of the topic.
	Name pulumi.StringPtrInput
	// The number of partitions the topic should have.
	Partitions pulumi.IntPtrInput
	// The number of replicas the topic should have.
	ReplicationFactor pulumi.IntPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// A map of string k/v attributes.
	Config map[string]interface{} `pulumi:"config"`
	// The name of the topic.
	Name *string `pulumi:"name"`
	// The number of partitions the topic should have.
	Partitions int `pulumi:"partitions"`
	// The number of replicas the topic should have.
	ReplicationFactor int `pulumi:"replicationFactor"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// A map of string k/v attributes.
	Config pulumi.MapInput
	// The name of the topic.
	Name pulumi.StringPtrInput
	// The number of partitions the topic should have.
	Partitions pulumi.IntInput
	// The number of replicas the topic should have.
	ReplicationFactor pulumi.IntInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//	TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//	TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

// A map of string k/v attributes.
func (o TopicOutput) Config() pulumi.MapOutput {
	return o.ApplyT(func(v *Topic) pulumi.MapOutput { return v.Config }).(pulumi.MapOutput)
}

// The name of the topic.
func (o TopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of partitions the topic should have.
func (o TopicOutput) Partitions() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.Partitions }).(pulumi.IntOutput)
}

// The number of replicas the topic should have.
func (o TopicOutput) ReplicationFactor() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.ReplicationFactor }).(pulumi.IntOutput)
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].([]*Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].(map[string]*Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicArrayInput)(nil)).Elem(), TopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicMapInput)(nil)).Elem(), TopicMap{})
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
