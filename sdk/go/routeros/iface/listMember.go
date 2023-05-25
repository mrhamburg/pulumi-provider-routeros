// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iface

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
//	"github.com/pulumi/pulumi-routeros/sdk/go/routeros/Iface"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Iface.NewListMember(ctx, "listMember", &Iface.ListMemberArgs{
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
//	$ pulumi import routeros:Iface/listMember:ListMember list_member "*0"
//
// ```
type ListMember struct {
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

// NewListMember registers a new resource with the given unique name, arguments, and options.
func NewListMember(ctx *pulumi.Context,
	name string, args *ListMemberArgs, opts ...pulumi.ResourceOption) (*ListMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.List == nil {
		return nil, errors.New("invalid value for required argument 'List'")
	}
	var resource ListMember
	err := ctx.RegisterResource("routeros:Iface/listMember:ListMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListMember gets an existing ListMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListMemberState, opts ...pulumi.ResourceOption) (*ListMember, error) {
	var resource ListMember
	err := ctx.ReadResource("routeros:Iface/listMember:ListMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListMember resources.
type listMemberState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  *string `pulumi:"___path_"`
	Disabled  *bool   `pulumi:"disabled"`
	Dynamic   *bool   `pulumi:"dynamic"`
	Interface *string `pulumi:"interface"`
	List      *string `pulumi:"list"`
}

type ListMemberState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  pulumi.StringPtrInput
	Disabled  pulumi.BoolPtrInput
	Dynamic   pulumi.BoolPtrInput
	Interface pulumi.StringPtrInput
	List      pulumi.StringPtrInput
}

func (ListMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*listMemberState)(nil)).Elem()
}

type listMemberArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  *string `pulumi:"___path_"`
	Disabled  *bool   `pulumi:"disabled"`
	Interface string  `pulumi:"interface"`
	List      string  `pulumi:"list"`
}

// The set of arguments for constructing a ListMember resource.
type ListMemberArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  pulumi.StringPtrInput
	Disabled  pulumi.BoolPtrInput
	Interface pulumi.StringInput
	List      pulumi.StringInput
}

func (ListMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listMemberArgs)(nil)).Elem()
}

type ListMemberInput interface {
	pulumi.Input

	ToListMemberOutput() ListMemberOutput
	ToListMemberOutputWithContext(ctx context.Context) ListMemberOutput
}

func (*ListMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ListMember)(nil)).Elem()
}

func (i *ListMember) ToListMemberOutput() ListMemberOutput {
	return i.ToListMemberOutputWithContext(context.Background())
}

func (i *ListMember) ToListMemberOutputWithContext(ctx context.Context) ListMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListMemberOutput)
}

// ListMemberArrayInput is an input type that accepts ListMemberArray and ListMemberArrayOutput values.
// You can construct a concrete instance of `ListMemberArrayInput` via:
//
//	ListMemberArray{ ListMemberArgs{...} }
type ListMemberArrayInput interface {
	pulumi.Input

	ToListMemberArrayOutput() ListMemberArrayOutput
	ToListMemberArrayOutputWithContext(context.Context) ListMemberArrayOutput
}

type ListMemberArray []ListMemberInput

func (ListMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ListMember)(nil)).Elem()
}

func (i ListMemberArray) ToListMemberArrayOutput() ListMemberArrayOutput {
	return i.ToListMemberArrayOutputWithContext(context.Background())
}

func (i ListMemberArray) ToListMemberArrayOutputWithContext(ctx context.Context) ListMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListMemberArrayOutput)
}

// ListMemberMapInput is an input type that accepts ListMemberMap and ListMemberMapOutput values.
// You can construct a concrete instance of `ListMemberMapInput` via:
//
//	ListMemberMap{ "key": ListMemberArgs{...} }
type ListMemberMapInput interface {
	pulumi.Input

	ToListMemberMapOutput() ListMemberMapOutput
	ToListMemberMapOutputWithContext(context.Context) ListMemberMapOutput
}

type ListMemberMap map[string]ListMemberInput

func (ListMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ListMember)(nil)).Elem()
}

func (i ListMemberMap) ToListMemberMapOutput() ListMemberMapOutput {
	return i.ToListMemberMapOutputWithContext(context.Background())
}

func (i ListMemberMap) ToListMemberMapOutputWithContext(ctx context.Context) ListMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListMemberMapOutput)
}

type ListMemberOutput struct{ *pulumi.OutputState }

func (ListMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListMember)(nil)).Elem()
}

func (o ListMemberOutput) ToListMemberOutput() ListMemberOutput {
	return o
}

func (o ListMemberOutput) ToListMemberOutputWithContext(ctx context.Context) ListMemberOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o ListMemberOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ListMember) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o ListMemberOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ListMember) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

func (o ListMemberOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ListMember) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

func (o ListMemberOutput) Dynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v *ListMember) pulumi.BoolOutput { return v.Dynamic }).(pulumi.BoolOutput)
}

func (o ListMemberOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *ListMember) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o ListMemberOutput) List() pulumi.StringOutput {
	return o.ApplyT(func(v *ListMember) pulumi.StringOutput { return v.List }).(pulumi.StringOutput)
}

type ListMemberArrayOutput struct{ *pulumi.OutputState }

func (ListMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ListMember)(nil)).Elem()
}

func (o ListMemberArrayOutput) ToListMemberArrayOutput() ListMemberArrayOutput {
	return o
}

func (o ListMemberArrayOutput) ToListMemberArrayOutputWithContext(ctx context.Context) ListMemberArrayOutput {
	return o
}

func (o ListMemberArrayOutput) Index(i pulumi.IntInput) ListMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ListMember {
		return vs[0].([]*ListMember)[vs[1].(int)]
	}).(ListMemberOutput)
}

type ListMemberMapOutput struct{ *pulumi.OutputState }

func (ListMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ListMember)(nil)).Elem()
}

func (o ListMemberMapOutput) ToListMemberMapOutput() ListMemberMapOutput {
	return o
}

func (o ListMemberMapOutput) ToListMemberMapOutputWithContext(ctx context.Context) ListMemberMapOutput {
	return o
}

func (o ListMemberMapOutput) MapIndex(k pulumi.StringInput) ListMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ListMember {
		return vs[0].(map[string]*ListMember)[vs[1].(string)]
	}).(ListMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListMemberInput)(nil)).Elem(), &ListMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListMemberArrayInput)(nil)).Elem(), ListMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListMemberMapInput)(nil)).Elem(), ListMemberMap{})
	pulumi.RegisterOutputType(ListMemberOutput{})
	pulumi.RegisterOutputType(ListMemberArrayOutput{})
	pulumi.RegisterOutputType(ListMemberMapOutput{})
}