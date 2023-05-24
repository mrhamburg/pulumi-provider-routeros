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
//			_, err := routeros.NewResourceUser(ctx, "test", &routeros.ResourceUserArgs{
//				Address:  pulumi.String("0.0.0.0/0"),
//				Comment:  pulumi.String("Test User"),
//				Group:    pulumi.String("read"),
//				Password: pulumi.String("secret"),
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
// #The ID can be found via API or the terminal #The command for the terminal is -> :put [/user get [print show-ids]]
//
// ```sh
//
//	$ pulumi import routeros:index/resourceUser:ResourceUser test *1
//
// ```
type ResourceUser struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// Host or network address from which the user is allowed to log in.
	Address  pulumi.StringPtrOutput `pulumi:"address"`
	Comment  pulumi.StringPtrOutput `pulumi:"comment"`
	Disabled pulumi.BoolPtrOutput   `pulumi:"disabled"`
	// Password expired.
	Expired pulumi.BoolOutput `pulumi:"expired"`
	// Name of the group the user belongs to.
	Group pulumi.StringOutput `pulumi:"group"`
	// Read-only field. Last time and date when a user logged in.
	LastLoggedIn pulumi.StringOutput `pulumi:"lastLoggedIn"`
	// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
	Name pulumi.StringOutput `pulumi:"name"`
	// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
	Password pulumi.StringPtrOutput `pulumi:"password"`
}

// NewResourceUser registers a new resource with the given unique name, arguments, and options.
func NewResourceUser(ctx *pulumi.Context,
	name string, args *ResourceUserArgs, opts ...pulumi.ResourceOption) (*ResourceUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	var resource ResourceUser
	err := ctx.RegisterResource("routeros:index/resourceUser:ResourceUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceUser gets an existing ResourceUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceUserState, opts ...pulumi.ResourceOption) (*ResourceUser, error) {
	var resource ResourceUser
	err := ctx.ReadResource("routeros:index/resourceUser:ResourceUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceUser resources.
type resourceUserState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Host or network address from which the user is allowed to log in.
	Address  *string `pulumi:"address"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// Password expired.
	Expired *bool `pulumi:"expired"`
	// Name of the group the user belongs to.
	Group *string `pulumi:"group"`
	// Read-only field. Last time and date when a user logged in.
	LastLoggedIn *string `pulumi:"lastLoggedIn"`
	// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
	Name *string `pulumi:"name"`
	// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
	Password *string `pulumi:"password"`
}

type ResourceUserState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Host or network address from which the user is allowed to log in.
	Address  pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// Password expired.
	Expired pulumi.BoolPtrInput
	// Name of the group the user belongs to.
	Group pulumi.StringPtrInput
	// Read-only field. Last time and date when a user logged in.
	LastLoggedIn pulumi.StringPtrInput
	// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
	Name pulumi.StringPtrInput
	// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
	Password pulumi.StringPtrInput
}

func (ResourceUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceUserState)(nil)).Elem()
}

type resourceUserArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Host or network address from which the user is allowed to log in.
	Address  *string `pulumi:"address"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// Name of the group the user belongs to.
	Group string `pulumi:"group"`
	// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
	Name *string `pulumi:"name"`
	// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
	Password *string `pulumi:"password"`
}

// The set of arguments for constructing a ResourceUser resource.
type ResourceUserArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Host or network address from which the user is allowed to log in.
	Address  pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// Name of the group the user belongs to.
	Group pulumi.StringInput
	// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
	Name pulumi.StringPtrInput
	// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
	Password pulumi.StringPtrInput
}

func (ResourceUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceUserArgs)(nil)).Elem()
}

type ResourceUserInput interface {
	pulumi.Input

	ToResourceUserOutput() ResourceUserOutput
	ToResourceUserOutputWithContext(ctx context.Context) ResourceUserOutput
}

func (*ResourceUser) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceUser)(nil)).Elem()
}

func (i *ResourceUser) ToResourceUserOutput() ResourceUserOutput {
	return i.ToResourceUserOutputWithContext(context.Background())
}

func (i *ResourceUser) ToResourceUserOutputWithContext(ctx context.Context) ResourceUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceUserOutput)
}

// ResourceUserArrayInput is an input type that accepts ResourceUserArray and ResourceUserArrayOutput values.
// You can construct a concrete instance of `ResourceUserArrayInput` via:
//
//	ResourceUserArray{ ResourceUserArgs{...} }
type ResourceUserArrayInput interface {
	pulumi.Input

	ToResourceUserArrayOutput() ResourceUserArrayOutput
	ToResourceUserArrayOutputWithContext(context.Context) ResourceUserArrayOutput
}

type ResourceUserArray []ResourceUserInput

func (ResourceUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceUser)(nil)).Elem()
}

func (i ResourceUserArray) ToResourceUserArrayOutput() ResourceUserArrayOutput {
	return i.ToResourceUserArrayOutputWithContext(context.Background())
}

func (i ResourceUserArray) ToResourceUserArrayOutputWithContext(ctx context.Context) ResourceUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceUserArrayOutput)
}

// ResourceUserMapInput is an input type that accepts ResourceUserMap and ResourceUserMapOutput values.
// You can construct a concrete instance of `ResourceUserMapInput` via:
//
//	ResourceUserMap{ "key": ResourceUserArgs{...} }
type ResourceUserMapInput interface {
	pulumi.Input

	ToResourceUserMapOutput() ResourceUserMapOutput
	ToResourceUserMapOutputWithContext(context.Context) ResourceUserMapOutput
}

type ResourceUserMap map[string]ResourceUserInput

func (ResourceUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceUser)(nil)).Elem()
}

func (i ResourceUserMap) ToResourceUserMapOutput() ResourceUserMapOutput {
	return i.ToResourceUserMapOutputWithContext(context.Background())
}

func (i ResourceUserMap) ToResourceUserMapOutputWithContext(ctx context.Context) ResourceUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceUserMapOutput)
}

type ResourceUserOutput struct{ *pulumi.OutputState }

func (ResourceUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceUser)(nil)).Elem()
}

func (o ResourceUserOutput) ToResourceUserOutput() ResourceUserOutput {
	return o
}

func (o ResourceUserOutput) ToResourceUserOutputWithContext(ctx context.Context) ResourceUserOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o ResourceUserOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o ResourceUserOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// Host or network address from which the user is allowed to log in.
func (o ResourceUserOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ResourceUserOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o ResourceUserOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Password expired.
func (o ResourceUserOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.BoolOutput { return v.Expired }).(pulumi.BoolOutput)
}

// Name of the group the user belongs to.
func (o ResourceUserOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// Read-only field. Last time and date when a user logged in.
func (o ResourceUserOutput) LastLoggedIn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.StringOutput { return v.LastLoggedIn }).(pulumi.StringOutput)
}

// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
func (o ResourceUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
func (o ResourceUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

type ResourceUserArrayOutput struct{ *pulumi.OutputState }

func (ResourceUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceUser)(nil)).Elem()
}

func (o ResourceUserArrayOutput) ToResourceUserArrayOutput() ResourceUserArrayOutput {
	return o
}

func (o ResourceUserArrayOutput) ToResourceUserArrayOutputWithContext(ctx context.Context) ResourceUserArrayOutput {
	return o
}

func (o ResourceUserArrayOutput) Index(i pulumi.IntInput) ResourceUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceUser {
		return vs[0].([]*ResourceUser)[vs[1].(int)]
	}).(ResourceUserOutput)
}

type ResourceUserMapOutput struct{ *pulumi.OutputState }

func (ResourceUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceUser)(nil)).Elem()
}

func (o ResourceUserMapOutput) ToResourceUserMapOutput() ResourceUserMapOutput {
	return o
}

func (o ResourceUserMapOutput) ToResourceUserMapOutputWithContext(ctx context.Context) ResourceUserMapOutput {
	return o
}

func (o ResourceUserMapOutput) MapIndex(k pulumi.StringInput) ResourceUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceUser {
		return vs[0].(map[string]*ResourceUser)[vs[1].(string)]
	}).(ResourceUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceUserInput)(nil)).Elem(), &ResourceUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceUserArrayInput)(nil)).Elem(), ResourceUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceUserMapInput)(nil)).Elem(), ResourceUserMap{})
	pulumi.RegisterOutputType(ResourceUserOutput{})
	pulumi.RegisterOutputType(ResourceUserArrayOutput{})
	pulumi.RegisterOutputType(ResourceUserMapOutput{})
}
