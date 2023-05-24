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
//			_, err := routeros.NewResourceInterfaceListMember(ctx, "listMember", &routeros.ResourceInterfaceListMemberArgs{
//				Interface: pulumi.String("ether1"),
//				List:      pulumi.String("my-list"),
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
// #The ID can be found via API or the terminal #The command for the terminal is -> :put [/interface/list/member get [print show-ids]]
//
// ```sh
//
//	$ pulumi import routeros:index/resourceInterfaceListMember:ResourceInterfaceListMember list_member "*0"
//
// ```
type ResourceInterfaceListMember struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  pulumi.StringPtrOutput `pulumi:"___path_"`
	Disabled  pulumi.BoolPtrOutput   `pulumi:"disabled"`
	Dynamic   pulumi.BoolOutput      `pulumi:"dynamic"`
	Interface pulumi.StringOutput    `pulumi:"interface"`
	List      pulumi.StringOutput    `pulumi:"list"`
}

// NewResourceInterfaceListMember registers a new resource with the given unique name, arguments, and options.
func NewResourceInterfaceListMember(ctx *pulumi.Context,
	name string, args *ResourceInterfaceListMemberArgs, opts ...pulumi.ResourceOption) (*ResourceInterfaceListMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.List == nil {
		return nil, errors.New("invalid value for required argument 'List'")
	}
	var resource ResourceInterfaceListMember
	err := ctx.RegisterResource("routeros:index/resourceInterfaceListMember:ResourceInterfaceListMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceInterfaceListMember gets an existing ResourceInterfaceListMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceInterfaceListMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceInterfaceListMemberState, opts ...pulumi.ResourceOption) (*ResourceInterfaceListMember, error) {
	var resource ResourceInterfaceListMember
	err := ctx.ReadResource("routeros:index/resourceInterfaceListMember:ResourceInterfaceListMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceInterfaceListMember resources.
type resourceInterfaceListMemberState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  *string `pulumi:"___path_"`
	Disabled  *bool   `pulumi:"disabled"`
	Dynamic   *bool   `pulumi:"dynamic"`
	Interface *string `pulumi:"interface"`
	List      *string `pulumi:"list"`
}

type ResourceInterfaceListMemberState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  pulumi.StringPtrInput
	Disabled  pulumi.BoolPtrInput
	Dynamic   pulumi.BoolPtrInput
	Interface pulumi.StringPtrInput
	List      pulumi.StringPtrInput
}

func (ResourceInterfaceListMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInterfaceListMemberState)(nil)).Elem()
}

type resourceInterfaceListMemberArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  *string `pulumi:"___path_"`
	Disabled  *bool   `pulumi:"disabled"`
	Interface string  `pulumi:"interface"`
	List      string  `pulumi:"list"`
}

// The set of arguments for constructing a ResourceInterfaceListMember resource.
type ResourceInterfaceListMemberArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  pulumi.StringPtrInput
	Disabled  pulumi.BoolPtrInput
	Interface pulumi.StringInput
	List      pulumi.StringInput
}

func (ResourceInterfaceListMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInterfaceListMemberArgs)(nil)).Elem()
}

type ResourceInterfaceListMemberInput interface {
	pulumi.Input

	ToResourceInterfaceListMemberOutput() ResourceInterfaceListMemberOutput
	ToResourceInterfaceListMemberOutputWithContext(ctx context.Context) ResourceInterfaceListMemberOutput
}

func (*ResourceInterfaceListMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInterfaceListMember)(nil)).Elem()
}

func (i *ResourceInterfaceListMember) ToResourceInterfaceListMemberOutput() ResourceInterfaceListMemberOutput {
	return i.ToResourceInterfaceListMemberOutputWithContext(context.Background())
}

func (i *ResourceInterfaceListMember) ToResourceInterfaceListMemberOutputWithContext(ctx context.Context) ResourceInterfaceListMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceInterfaceListMemberOutput)
}

// ResourceInterfaceListMemberArrayInput is an input type that accepts ResourceInterfaceListMemberArray and ResourceInterfaceListMemberArrayOutput values.
// You can construct a concrete instance of `ResourceInterfaceListMemberArrayInput` via:
//
//	ResourceInterfaceListMemberArray{ ResourceInterfaceListMemberArgs{...} }
type ResourceInterfaceListMemberArrayInput interface {
	pulumi.Input

	ToResourceInterfaceListMemberArrayOutput() ResourceInterfaceListMemberArrayOutput
	ToResourceInterfaceListMemberArrayOutputWithContext(context.Context) ResourceInterfaceListMemberArrayOutput
}

type ResourceInterfaceListMemberArray []ResourceInterfaceListMemberInput

func (ResourceInterfaceListMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceInterfaceListMember)(nil)).Elem()
}

func (i ResourceInterfaceListMemberArray) ToResourceInterfaceListMemberArrayOutput() ResourceInterfaceListMemberArrayOutput {
	return i.ToResourceInterfaceListMemberArrayOutputWithContext(context.Background())
}

func (i ResourceInterfaceListMemberArray) ToResourceInterfaceListMemberArrayOutputWithContext(ctx context.Context) ResourceInterfaceListMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceInterfaceListMemberArrayOutput)
}

// ResourceInterfaceListMemberMapInput is an input type that accepts ResourceInterfaceListMemberMap and ResourceInterfaceListMemberMapOutput values.
// You can construct a concrete instance of `ResourceInterfaceListMemberMapInput` via:
//
//	ResourceInterfaceListMemberMap{ "key": ResourceInterfaceListMemberArgs{...} }
type ResourceInterfaceListMemberMapInput interface {
	pulumi.Input

	ToResourceInterfaceListMemberMapOutput() ResourceInterfaceListMemberMapOutput
	ToResourceInterfaceListMemberMapOutputWithContext(context.Context) ResourceInterfaceListMemberMapOutput
}

type ResourceInterfaceListMemberMap map[string]ResourceInterfaceListMemberInput

func (ResourceInterfaceListMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceInterfaceListMember)(nil)).Elem()
}

func (i ResourceInterfaceListMemberMap) ToResourceInterfaceListMemberMapOutput() ResourceInterfaceListMemberMapOutput {
	return i.ToResourceInterfaceListMemberMapOutputWithContext(context.Background())
}

func (i ResourceInterfaceListMemberMap) ToResourceInterfaceListMemberMapOutputWithContext(ctx context.Context) ResourceInterfaceListMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceInterfaceListMemberMapOutput)
}

type ResourceInterfaceListMemberOutput struct{ *pulumi.OutputState }

func (ResourceInterfaceListMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInterfaceListMember)(nil)).Elem()
}

func (o ResourceInterfaceListMemberOutput) ToResourceInterfaceListMemberOutput() ResourceInterfaceListMemberOutput {
	return o
}

func (o ResourceInterfaceListMemberOutput) ToResourceInterfaceListMemberOutputWithContext(ctx context.Context) ResourceInterfaceListMemberOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o ResourceInterfaceListMemberOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceListMember) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o ResourceInterfaceListMemberOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceListMember) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

func (o ResourceInterfaceListMemberOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceListMember) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

func (o ResourceInterfaceListMemberOutput) Dynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v *ResourceInterfaceListMember) pulumi.BoolOutput { return v.Dynamic }).(pulumi.BoolOutput)
}

func (o ResourceInterfaceListMemberOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceInterfaceListMember) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o ResourceInterfaceListMemberOutput) List() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceInterfaceListMember) pulumi.StringOutput { return v.List }).(pulumi.StringOutput)
}

type ResourceInterfaceListMemberArrayOutput struct{ *pulumi.OutputState }

func (ResourceInterfaceListMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceInterfaceListMember)(nil)).Elem()
}

func (o ResourceInterfaceListMemberArrayOutput) ToResourceInterfaceListMemberArrayOutput() ResourceInterfaceListMemberArrayOutput {
	return o
}

func (o ResourceInterfaceListMemberArrayOutput) ToResourceInterfaceListMemberArrayOutputWithContext(ctx context.Context) ResourceInterfaceListMemberArrayOutput {
	return o
}

func (o ResourceInterfaceListMemberArrayOutput) Index(i pulumi.IntInput) ResourceInterfaceListMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceInterfaceListMember {
		return vs[0].([]*ResourceInterfaceListMember)[vs[1].(int)]
	}).(ResourceInterfaceListMemberOutput)
}

type ResourceInterfaceListMemberMapOutput struct{ *pulumi.OutputState }

func (ResourceInterfaceListMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceInterfaceListMember)(nil)).Elem()
}

func (o ResourceInterfaceListMemberMapOutput) ToResourceInterfaceListMemberMapOutput() ResourceInterfaceListMemberMapOutput {
	return o
}

func (o ResourceInterfaceListMemberMapOutput) ToResourceInterfaceListMemberMapOutputWithContext(ctx context.Context) ResourceInterfaceListMemberMapOutput {
	return o
}

func (o ResourceInterfaceListMemberMapOutput) MapIndex(k pulumi.StringInput) ResourceInterfaceListMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceInterfaceListMember {
		return vs[0].(map[string]*ResourceInterfaceListMember)[vs[1].(string)]
	}).(ResourceInterfaceListMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInterfaceListMemberInput)(nil)).Elem(), &ResourceInterfaceListMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInterfaceListMemberArrayInput)(nil)).Elem(), ResourceInterfaceListMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInterfaceListMemberMapInput)(nil)).Elem(), ResourceInterfaceListMemberMap{})
	pulumi.RegisterOutputType(ResourceInterfaceListMemberOutput{})
	pulumi.RegisterOutputType(ResourceInterfaceListMemberArrayOutput{})
	pulumi.RegisterOutputType(ResourceInterfaceListMemberMapOutput{})
}
