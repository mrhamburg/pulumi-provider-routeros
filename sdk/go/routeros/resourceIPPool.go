// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package routeros

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-routeros/sdk/go/routeros"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := routeros.NewResourceIPPool(ctx, "pool", &routeros.ResourceIPPoolArgs{
//				Ranges: pulumi.StringArray{
//					pulumi.String("10.0.0.100-10.0.0.200"),
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
// Import with the name of the ip pool in case of the example use my_ip_pool
//
// ```sh
//
//	$ pulumi import routeros:index/resourceIPPool:ResourceIPPool pool my_ip_pool
//
// ```
type ResourceIPPool struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	Comment  pulumi.StringPtrOutput `pulumi:"comment"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringOutput `pulumi:"name"`
	// When address is acquired from pool that has no free addresses, and next-pool property is set to another pool, then next IP address will be acquired from next-pool.
	NextPool pulumi.StringPtrOutput `pulumi:"nextPool"`
	// IP address list of non-overlapping IP address ranges in form of: ["from1-to1", "from2-to2", ..., "fromN-toN"]. For example, ["10.0.0.1-10.0.0.27", "10.0.0.32-10.0.0.47"]
	Ranges pulumi.StringArrayOutput `pulumi:"ranges"`
}

// NewResourceIPPool registers a new resource with the given unique name, arguments, and options.
func NewResourceIPPool(ctx *pulumi.Context,
	name string, args *ResourceIPPoolArgs, opts ...pulumi.ResourceOption) (*ResourceIPPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ranges == nil {
		return nil, errors.New("invalid value for required argument 'Ranges'")
	}
	var resource ResourceIPPool
	err := ctx.RegisterResource("routeros:index/resourceIPPool:ResourceIPPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceIPPool gets an existing ResourceIPPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceIPPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceIPPoolState, opts ...pulumi.ResourceOption) (*ResourceIPPool, error) {
	var resource ResourceIPPool
	err := ctx.ReadResource("routeros:index/resourceIPPool:ResourceIPPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceIPPool resources.
type resourceIPPoolState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	Comment  *string `pulumi:"comment"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// When address is acquired from pool that has no free addresses, and next-pool property is set to another pool, then next IP address will be acquired from next-pool.
	NextPool *string `pulumi:"nextPool"`
	// IP address list of non-overlapping IP address ranges in form of: ["from1-to1", "from2-to2", ..., "fromN-toN"]. For example, ["10.0.0.1-10.0.0.27", "10.0.0.32-10.0.0.47"]
	Ranges []string `pulumi:"ranges"`
}

type ResourceIPPoolState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// When address is acquired from pool that has no free addresses, and next-pool property is set to another pool, then next IP address will be acquired from next-pool.
	NextPool pulumi.StringPtrInput
	// IP address list of non-overlapping IP address ranges in form of: ["from1-to1", "from2-to2", ..., "fromN-toN"]. For example, ["10.0.0.1-10.0.0.27", "10.0.0.32-10.0.0.47"]
	Ranges pulumi.StringArrayInput
}

func (ResourceIPPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceIPPoolState)(nil)).Elem()
}

type resourceIPPoolArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	Comment  *string `pulumi:"comment"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// When address is acquired from pool that has no free addresses, and next-pool property is set to another pool, then next IP address will be acquired from next-pool.
	NextPool *string `pulumi:"nextPool"`
	// IP address list of non-overlapping IP address ranges in form of: ["from1-to1", "from2-to2", ..., "fromN-toN"]. For example, ["10.0.0.1-10.0.0.27", "10.0.0.32-10.0.0.47"]
	Ranges []string `pulumi:"ranges"`
}

// The set of arguments for constructing a ResourceIPPool resource.
type ResourceIPPoolArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// When address is acquired from pool that has no free addresses, and next-pool property is set to another pool, then next IP address will be acquired from next-pool.
	NextPool pulumi.StringPtrInput
	// IP address list of non-overlapping IP address ranges in form of: ["from1-to1", "from2-to2", ..., "fromN-toN"]. For example, ["10.0.0.1-10.0.0.27", "10.0.0.32-10.0.0.47"]
	Ranges pulumi.StringArrayInput
}

func (ResourceIPPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceIPPoolArgs)(nil)).Elem()
}

type ResourceIPPoolInput interface {
	pulumi.Input

	ToResourceIPPoolOutput() ResourceIPPoolOutput
	ToResourceIPPoolOutputWithContext(ctx context.Context) ResourceIPPoolOutput
}

func (*ResourceIPPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIPPool)(nil)).Elem()
}

func (i *ResourceIPPool) ToResourceIPPoolOutput() ResourceIPPoolOutput {
	return i.ToResourceIPPoolOutputWithContext(context.Background())
}

func (i *ResourceIPPool) ToResourceIPPoolOutputWithContext(ctx context.Context) ResourceIPPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIPPoolOutput)
}

// ResourceIPPoolArrayInput is an input type that accepts ResourceIPPoolArray and ResourceIPPoolArrayOutput values.
// You can construct a concrete instance of `ResourceIPPoolArrayInput` via:
//
//	ResourceIPPoolArray{ ResourceIPPoolArgs{...} }
type ResourceIPPoolArrayInput interface {
	pulumi.Input

	ToResourceIPPoolArrayOutput() ResourceIPPoolArrayOutput
	ToResourceIPPoolArrayOutputWithContext(context.Context) ResourceIPPoolArrayOutput
}

type ResourceIPPoolArray []ResourceIPPoolInput

func (ResourceIPPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceIPPool)(nil)).Elem()
}

func (i ResourceIPPoolArray) ToResourceIPPoolArrayOutput() ResourceIPPoolArrayOutput {
	return i.ToResourceIPPoolArrayOutputWithContext(context.Background())
}

func (i ResourceIPPoolArray) ToResourceIPPoolArrayOutputWithContext(ctx context.Context) ResourceIPPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIPPoolArrayOutput)
}

// ResourceIPPoolMapInput is an input type that accepts ResourceIPPoolMap and ResourceIPPoolMapOutput values.
// You can construct a concrete instance of `ResourceIPPoolMapInput` via:
//
//	ResourceIPPoolMap{ "key": ResourceIPPoolArgs{...} }
type ResourceIPPoolMapInput interface {
	pulumi.Input

	ToResourceIPPoolMapOutput() ResourceIPPoolMapOutput
	ToResourceIPPoolMapOutputWithContext(context.Context) ResourceIPPoolMapOutput
}

type ResourceIPPoolMap map[string]ResourceIPPoolInput

func (ResourceIPPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceIPPool)(nil)).Elem()
}

func (i ResourceIPPoolMap) ToResourceIPPoolMapOutput() ResourceIPPoolMapOutput {
	return i.ToResourceIPPoolMapOutputWithContext(context.Background())
}

func (i ResourceIPPoolMap) ToResourceIPPoolMapOutputWithContext(ctx context.Context) ResourceIPPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIPPoolMapOutput)
}

type ResourceIPPoolOutput struct{ *pulumi.OutputState }

func (ResourceIPPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIPPool)(nil)).Elem()
}

func (o ResourceIPPoolOutput) ToResourceIPPoolOutput() ResourceIPPoolOutput {
	return o
}

func (o ResourceIPPoolOutput) ToResourceIPPoolOutputWithContext(ctx context.Context) ResourceIPPoolOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o ResourceIPPoolOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceIPPool) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o ResourceIPPoolOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIPPool) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

func (o ResourceIPPoolOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIPPool) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o ResourceIPPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceIPPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// When address is acquired from pool that has no free addresses, and next-pool property is set to another pool, then next IP address will be acquired from next-pool.
func (o ResourceIPPoolOutput) NextPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIPPool) pulumi.StringPtrOutput { return v.NextPool }).(pulumi.StringPtrOutput)
}

// IP address list of non-overlapping IP address ranges in form of: ["from1-to1", "from2-to2", ..., "fromN-toN"]. For example, ["10.0.0.1-10.0.0.27", "10.0.0.32-10.0.0.47"]
func (o ResourceIPPoolOutput) Ranges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceIPPool) pulumi.StringArrayOutput { return v.Ranges }).(pulumi.StringArrayOutput)
}

type ResourceIPPoolArrayOutput struct{ *pulumi.OutputState }

func (ResourceIPPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceIPPool)(nil)).Elem()
}

func (o ResourceIPPoolArrayOutput) ToResourceIPPoolArrayOutput() ResourceIPPoolArrayOutput {
	return o
}

func (o ResourceIPPoolArrayOutput) ToResourceIPPoolArrayOutputWithContext(ctx context.Context) ResourceIPPoolArrayOutput {
	return o
}

func (o ResourceIPPoolArrayOutput) Index(i pulumi.IntInput) ResourceIPPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceIPPool {
		return vs[0].([]*ResourceIPPool)[vs[1].(int)]
	}).(ResourceIPPoolOutput)
}

type ResourceIPPoolMapOutput struct{ *pulumi.OutputState }

func (ResourceIPPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceIPPool)(nil)).Elem()
}

func (o ResourceIPPoolMapOutput) ToResourceIPPoolMapOutput() ResourceIPPoolMapOutput {
	return o
}

func (o ResourceIPPoolMapOutput) ToResourceIPPoolMapOutputWithContext(ctx context.Context) ResourceIPPoolMapOutput {
	return o
}

func (o ResourceIPPoolMapOutput) MapIndex(k pulumi.StringInput) ResourceIPPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceIPPool {
		return vs[0].(map[string]*ResourceIPPool)[vs[1].(string)]
	}).(ResourceIPPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIPPoolInput)(nil)).Elem(), &ResourceIPPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIPPoolArrayInput)(nil)).Elem(), ResourceIPPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIPPoolMapInput)(nil)).Elem(), ResourceIPPoolMap{})
	pulumi.RegisterOutputType(ResourceIPPoolOutput{})
	pulumi.RegisterOutputType(ResourceIPPoolArrayOutput{})
	pulumi.RegisterOutputType(ResourceIPPoolMapOutput{})
}
