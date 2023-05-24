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
//			_, err := routeros.NewResourceInterfaceGre(ctx, "greHq", &routeros.ResourceInterfaceGreArgs{
//				Disabled:      pulumi.Bool(true),
//				RemoteAddress: pulumi.String("10.77.3.26"),
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
// # Import with the name of the gre interface in case of the example use gre-hq-1
//
// ```sh
//
//	$ pulumi import routeros:index/resourceInterfaceGre:ResourceInterfaceGre gre_hq gre-hq-1
//
// ```
type ResourceInterfaceGre struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  pulumi.StringPtrOutput `pulumi:"___path_"`
	ActualMtu pulumi.IntOutput       `pulumi:"actualMtu"`
	// Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
	AllowFastPath pulumi.BoolPtrOutput `pulumi:"allowFastPath"`
	// Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
	ClampTcpMss  pulumi.BoolPtrOutput   `pulumi:"clampTcpMss"`
	Comment      pulumi.StringPtrOutput `pulumi:"comment"`
	Disabled     pulumi.BoolPtrOutput   `pulumi:"disabled"`
	DontFragment pulumi.StringPtrOutput `pulumi:"dontFragment"`
	// Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
	Dscp pulumi.StringPtrOutput `pulumi:"dscp"`
	// When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc).
	IpsecSecret pulumi.StringPtrOutput `pulumi:"ipsecSecret"`
	// Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295
	Keepalive pulumi.StringPtrOutput `pulumi:"keepalive"`
	// Layer2 Maximum transmission unit.
	L2mtu        pulumi.IntOutput       `pulumi:"l2mtu"`
	LocalAddress pulumi.StringPtrOutput `pulumi:"localAddress"`
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu pulumi.StringOutput `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name          pulumi.StringOutput `pulumi:"name"`
	RemoteAddress pulumi.StringOutput `pulumi:"remoteAddress"`
	Running       pulumi.BoolOutput   `pulumi:"running"`
}

// NewResourceInterfaceGre registers a new resource with the given unique name, arguments, and options.
func NewResourceInterfaceGre(ctx *pulumi.Context,
	name string, args *ResourceInterfaceGreArgs, opts ...pulumi.ResourceOption) (*ResourceInterfaceGre, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RemoteAddress == nil {
		return nil, errors.New("invalid value for required argument 'RemoteAddress'")
	}
	if args.IpsecSecret != nil {
		args.IpsecSecret = pulumi.ToSecret(args.IpsecSecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"ipsecSecret",
	})
	opts = append(opts, secrets)
	var resource ResourceInterfaceGre
	err := ctx.RegisterResource("routeros:index/resourceInterfaceGre:ResourceInterfaceGre", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceInterfaceGre gets an existing ResourceInterfaceGre resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceInterfaceGre(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceInterfaceGreState, opts ...pulumi.ResourceOption) (*ResourceInterfaceGre, error) {
	var resource ResourceInterfaceGre
	err := ctx.ReadResource("routeros:index/resourceInterfaceGre:ResourceInterfaceGre", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceInterfaceGre resources.
type resourceInterfaceGreState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  *string `pulumi:"___path_"`
	ActualMtu *int    `pulumi:"actualMtu"`
	// Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
	AllowFastPath *bool `pulumi:"allowFastPath"`
	// Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
	ClampTcpMss  *bool   `pulumi:"clampTcpMss"`
	Comment      *string `pulumi:"comment"`
	Disabled     *bool   `pulumi:"disabled"`
	DontFragment *string `pulumi:"dontFragment"`
	// Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
	Dscp *string `pulumi:"dscp"`
	// When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc).
	IpsecSecret *string `pulumi:"ipsecSecret"`
	// Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295
	Keepalive *string `pulumi:"keepalive"`
	// Layer2 Maximum transmission unit.
	L2mtu        *int    `pulumi:"l2mtu"`
	LocalAddress *string `pulumi:"localAddress"`
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu *string `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name          *string `pulumi:"name"`
	RemoteAddress *string `pulumi:"remoteAddress"`
	Running       *bool   `pulumi:"running"`
}

type ResourceInterfaceGreState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_  pulumi.StringPtrInput
	ActualMtu pulumi.IntPtrInput
	// Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
	AllowFastPath pulumi.BoolPtrInput
	// Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
	ClampTcpMss  pulumi.BoolPtrInput
	Comment      pulumi.StringPtrInput
	Disabled     pulumi.BoolPtrInput
	DontFragment pulumi.StringPtrInput
	// Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
	Dscp pulumi.StringPtrInput
	// When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc).
	IpsecSecret pulumi.StringPtrInput
	// Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295
	Keepalive pulumi.StringPtrInput
	// Layer2 Maximum transmission unit.
	L2mtu        pulumi.IntPtrInput
	LocalAddress pulumi.StringPtrInput
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name          pulumi.StringPtrInput
	RemoteAddress pulumi.StringPtrInput
	Running       pulumi.BoolPtrInput
}

func (ResourceInterfaceGreState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInterfaceGreState)(nil)).Elem()
}

type resourceInterfaceGreArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
	AllowFastPath *bool `pulumi:"allowFastPath"`
	// Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
	ClampTcpMss  *bool   `pulumi:"clampTcpMss"`
	Comment      *string `pulumi:"comment"`
	Disabled     *bool   `pulumi:"disabled"`
	DontFragment *string `pulumi:"dontFragment"`
	// Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
	Dscp *string `pulumi:"dscp"`
	// When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc).
	IpsecSecret *string `pulumi:"ipsecSecret"`
	// Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295
	Keepalive    *string `pulumi:"keepalive"`
	LocalAddress *string `pulumi:"localAddress"`
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu *string `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name          *string `pulumi:"name"`
	RemoteAddress string  `pulumi:"remoteAddress"`
}

// The set of arguments for constructing a ResourceInterfaceGre resource.
type ResourceInterfaceGreArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
	AllowFastPath pulumi.BoolPtrInput
	// Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
	ClampTcpMss  pulumi.BoolPtrInput
	Comment      pulumi.StringPtrInput
	Disabled     pulumi.BoolPtrInput
	DontFragment pulumi.StringPtrInput
	// Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
	Dscp pulumi.StringPtrInput
	// When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc).
	IpsecSecret pulumi.StringPtrInput
	// Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295
	Keepalive    pulumi.StringPtrInput
	LocalAddress pulumi.StringPtrInput
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name          pulumi.StringPtrInput
	RemoteAddress pulumi.StringInput
}

func (ResourceInterfaceGreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInterfaceGreArgs)(nil)).Elem()
}

type ResourceInterfaceGreInput interface {
	pulumi.Input

	ToResourceInterfaceGreOutput() ResourceInterfaceGreOutput
	ToResourceInterfaceGreOutputWithContext(ctx context.Context) ResourceInterfaceGreOutput
}

func (*ResourceInterfaceGre) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInterfaceGre)(nil)).Elem()
}

func (i *ResourceInterfaceGre) ToResourceInterfaceGreOutput() ResourceInterfaceGreOutput {
	return i.ToResourceInterfaceGreOutputWithContext(context.Background())
}

func (i *ResourceInterfaceGre) ToResourceInterfaceGreOutputWithContext(ctx context.Context) ResourceInterfaceGreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceInterfaceGreOutput)
}

// ResourceInterfaceGreArrayInput is an input type that accepts ResourceInterfaceGreArray and ResourceInterfaceGreArrayOutput values.
// You can construct a concrete instance of `ResourceInterfaceGreArrayInput` via:
//
//	ResourceInterfaceGreArray{ ResourceInterfaceGreArgs{...} }
type ResourceInterfaceGreArrayInput interface {
	pulumi.Input

	ToResourceInterfaceGreArrayOutput() ResourceInterfaceGreArrayOutput
	ToResourceInterfaceGreArrayOutputWithContext(context.Context) ResourceInterfaceGreArrayOutput
}

type ResourceInterfaceGreArray []ResourceInterfaceGreInput

func (ResourceInterfaceGreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceInterfaceGre)(nil)).Elem()
}

func (i ResourceInterfaceGreArray) ToResourceInterfaceGreArrayOutput() ResourceInterfaceGreArrayOutput {
	return i.ToResourceInterfaceGreArrayOutputWithContext(context.Background())
}

func (i ResourceInterfaceGreArray) ToResourceInterfaceGreArrayOutputWithContext(ctx context.Context) ResourceInterfaceGreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceInterfaceGreArrayOutput)
}

// ResourceInterfaceGreMapInput is an input type that accepts ResourceInterfaceGreMap and ResourceInterfaceGreMapOutput values.
// You can construct a concrete instance of `ResourceInterfaceGreMapInput` via:
//
//	ResourceInterfaceGreMap{ "key": ResourceInterfaceGreArgs{...} }
type ResourceInterfaceGreMapInput interface {
	pulumi.Input

	ToResourceInterfaceGreMapOutput() ResourceInterfaceGreMapOutput
	ToResourceInterfaceGreMapOutputWithContext(context.Context) ResourceInterfaceGreMapOutput
}

type ResourceInterfaceGreMap map[string]ResourceInterfaceGreInput

func (ResourceInterfaceGreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceInterfaceGre)(nil)).Elem()
}

func (i ResourceInterfaceGreMap) ToResourceInterfaceGreMapOutput() ResourceInterfaceGreMapOutput {
	return i.ToResourceInterfaceGreMapOutputWithContext(context.Background())
}

func (i ResourceInterfaceGreMap) ToResourceInterfaceGreMapOutputWithContext(ctx context.Context) ResourceInterfaceGreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceInterfaceGreMapOutput)
}

type ResourceInterfaceGreOutput struct{ *pulumi.OutputState }

func (ResourceInterfaceGreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInterfaceGre)(nil)).Elem()
}

func (o ResourceInterfaceGreOutput) ToResourceInterfaceGreOutput() ResourceInterfaceGreOutput {
	return o
}

func (o ResourceInterfaceGreOutput) ToResourceInterfaceGreOutputWithContext(ctx context.Context) ResourceInterfaceGreOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o ResourceInterfaceGreOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o ResourceInterfaceGreOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

func (o ResourceInterfaceGreOutput) ActualMtu() pulumi.IntOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.IntOutput { return v.ActualMtu }).(pulumi.IntOutput)
}

// Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
func (o ResourceInterfaceGreOutput) AllowFastPath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.BoolPtrOutput { return v.AllowFastPath }).(pulumi.BoolPtrOutput)
}

// Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
func (o ResourceInterfaceGreOutput) ClampTcpMss() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.BoolPtrOutput { return v.ClampTcpMss }).(pulumi.BoolPtrOutput)
}

func (o ResourceInterfaceGreOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o ResourceInterfaceGreOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

func (o ResourceInterfaceGreOutput) DontFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringPtrOutput { return v.DontFragment }).(pulumi.StringPtrOutput)
}

// Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken from tunnelled traffic.
func (o ResourceInterfaceGreOutput) Dscp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringPtrOutput { return v.Dscp }).(pulumi.StringPtrOutput)
}

// When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc).
func (o ResourceInterfaceGreOutput) IpsecSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringPtrOutput { return v.IpsecSecret }).(pulumi.StringPtrOutput)
}

// Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295
func (o ResourceInterfaceGreOutput) Keepalive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringPtrOutput { return v.Keepalive }).(pulumi.StringPtrOutput)
}

// Layer2 Maximum transmission unit.
func (o ResourceInterfaceGreOutput) L2mtu() pulumi.IntOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.IntOutput { return v.L2mtu }).(pulumi.IntOutput)
}

func (o ResourceInterfaceGreOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringPtrOutput { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
func (o ResourceInterfaceGreOutput) Mtu() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringOutput { return v.Mtu }).(pulumi.StringOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o ResourceInterfaceGreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceInterfaceGreOutput) RemoteAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.StringOutput { return v.RemoteAddress }).(pulumi.StringOutput)
}

func (o ResourceInterfaceGreOutput) Running() pulumi.BoolOutput {
	return o.ApplyT(func(v *ResourceInterfaceGre) pulumi.BoolOutput { return v.Running }).(pulumi.BoolOutput)
}

type ResourceInterfaceGreArrayOutput struct{ *pulumi.OutputState }

func (ResourceInterfaceGreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceInterfaceGre)(nil)).Elem()
}

func (o ResourceInterfaceGreArrayOutput) ToResourceInterfaceGreArrayOutput() ResourceInterfaceGreArrayOutput {
	return o
}

func (o ResourceInterfaceGreArrayOutput) ToResourceInterfaceGreArrayOutputWithContext(ctx context.Context) ResourceInterfaceGreArrayOutput {
	return o
}

func (o ResourceInterfaceGreArrayOutput) Index(i pulumi.IntInput) ResourceInterfaceGreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceInterfaceGre {
		return vs[0].([]*ResourceInterfaceGre)[vs[1].(int)]
	}).(ResourceInterfaceGreOutput)
}

type ResourceInterfaceGreMapOutput struct{ *pulumi.OutputState }

func (ResourceInterfaceGreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceInterfaceGre)(nil)).Elem()
}

func (o ResourceInterfaceGreMapOutput) ToResourceInterfaceGreMapOutput() ResourceInterfaceGreMapOutput {
	return o
}

func (o ResourceInterfaceGreMapOutput) ToResourceInterfaceGreMapOutputWithContext(ctx context.Context) ResourceInterfaceGreMapOutput {
	return o
}

func (o ResourceInterfaceGreMapOutput) MapIndex(k pulumi.StringInput) ResourceInterfaceGreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceInterfaceGre {
		return vs[0].(map[string]*ResourceInterfaceGre)[vs[1].(string)]
	}).(ResourceInterfaceGreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInterfaceGreInput)(nil)).Elem(), &ResourceInterfaceGre{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInterfaceGreArrayInput)(nil)).Elem(), ResourceInterfaceGreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInterfaceGreMapInput)(nil)).Elem(), ResourceInterfaceGreMap{})
	pulumi.RegisterOutputType(ResourceInterfaceGreOutput{})
	pulumi.RegisterOutputType(ResourceInterfaceGreArrayOutput{})
	pulumi.RegisterOutputType(ResourceInterfaceGreMapOutput{})
}
