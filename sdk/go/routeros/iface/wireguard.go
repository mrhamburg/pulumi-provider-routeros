// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iface

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Iface.Wireguard (Resource)
//
// ***
//
// #### This is an alias for backwards compatibility between plugin versions.
// Please see documentation for Iface.InterfaceWireguard
type Wireguard struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	Comment  pulumi.StringPtrOutput `pulumi:"comment"`
	Disabled pulumi.BoolPtrOutput   `pulumi:"disabled"`
	// Port for WireGuard service to listen on for incoming sessions.
	ListenPort pulumi.IntOutput `pulumi:"listenPort"`
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu pulumi.StringOutput `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringOutput `pulumi:"name"`
	// A base64 private key. If not specified, it will be automatically generated upon interface creation.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// A base64 public key is calculated from the private key.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	Running   pulumi.BoolOutput   `pulumi:"running"`
}

// NewWireguard registers a new resource with the given unique name, arguments, and options.
func NewWireguard(ctx *pulumi.Context,
	name string, args *WireguardArgs, opts ...pulumi.ResourceOption) (*Wireguard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ListenPort == nil {
		return nil, errors.New("invalid value for required argument 'ListenPort'")
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	var resource Wireguard
	err := ctx.RegisterResource("routeros:Iface/wireguard:Wireguard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWireguard gets an existing Wireguard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWireguard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WireguardState, opts ...pulumi.ResourceOption) (*Wireguard, error) {
	var resource Wireguard
	err := ctx.ReadResource("routeros:Iface/wireguard:Wireguard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wireguard resources.
type wireguardState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// Port for WireGuard service to listen on for incoming sessions.
	ListenPort *int `pulumi:"listenPort"`
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu *string `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// A base64 private key. If not specified, it will be automatically generated upon interface creation.
	PrivateKey *string `pulumi:"privateKey"`
	// A base64 public key is calculated from the private key.
	PublicKey *string `pulumi:"publicKey"`
	Running   *bool   `pulumi:"running"`
}

type WireguardState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// Port for WireGuard service to listen on for incoming sessions.
	ListenPort pulumi.IntPtrInput
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// A base64 private key. If not specified, it will be automatically generated upon interface creation.
	PrivateKey pulumi.StringPtrInput
	// A base64 public key is calculated from the private key.
	PublicKey pulumi.StringPtrInput
	Running   pulumi.BoolPtrInput
}

func (WireguardState) ElementType() reflect.Type {
	return reflect.TypeOf((*wireguardState)(nil)).Elem()
}

type wireguardArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// Port for WireGuard service to listen on for incoming sessions.
	ListenPort int `pulumi:"listenPort"`
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu *string `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// A base64 private key. If not specified, it will be automatically generated upon interface creation.
	PrivateKey *string `pulumi:"privateKey"`
}

// The set of arguments for constructing a Wireguard resource.
type WireguardArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// Port for WireGuard service to listen on for incoming sessions.
	ListenPort pulumi.IntInput
	// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
	Mtu pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// A base64 private key. If not specified, it will be automatically generated upon interface creation.
	PrivateKey pulumi.StringPtrInput
}

func (WireguardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wireguardArgs)(nil)).Elem()
}

type WireguardInput interface {
	pulumi.Input

	ToWireguardOutput() WireguardOutput
	ToWireguardOutputWithContext(ctx context.Context) WireguardOutput
}

func (*Wireguard) ElementType() reflect.Type {
	return reflect.TypeOf((**Wireguard)(nil)).Elem()
}

func (i *Wireguard) ToWireguardOutput() WireguardOutput {
	return i.ToWireguardOutputWithContext(context.Background())
}

func (i *Wireguard) ToWireguardOutputWithContext(ctx context.Context) WireguardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WireguardOutput)
}

// WireguardArrayInput is an input type that accepts WireguardArray and WireguardArrayOutput values.
// You can construct a concrete instance of `WireguardArrayInput` via:
//
//	WireguardArray{ WireguardArgs{...} }
type WireguardArrayInput interface {
	pulumi.Input

	ToWireguardArrayOutput() WireguardArrayOutput
	ToWireguardArrayOutputWithContext(context.Context) WireguardArrayOutput
}

type WireguardArray []WireguardInput

func (WireguardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wireguard)(nil)).Elem()
}

func (i WireguardArray) ToWireguardArrayOutput() WireguardArrayOutput {
	return i.ToWireguardArrayOutputWithContext(context.Background())
}

func (i WireguardArray) ToWireguardArrayOutputWithContext(ctx context.Context) WireguardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WireguardArrayOutput)
}

// WireguardMapInput is an input type that accepts WireguardMap and WireguardMapOutput values.
// You can construct a concrete instance of `WireguardMapInput` via:
//
//	WireguardMap{ "key": WireguardArgs{...} }
type WireguardMapInput interface {
	pulumi.Input

	ToWireguardMapOutput() WireguardMapOutput
	ToWireguardMapOutputWithContext(context.Context) WireguardMapOutput
}

type WireguardMap map[string]WireguardInput

func (WireguardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wireguard)(nil)).Elem()
}

func (i WireguardMap) ToWireguardMapOutput() WireguardMapOutput {
	return i.ToWireguardMapOutputWithContext(context.Background())
}

func (i WireguardMap) ToWireguardMapOutputWithContext(ctx context.Context) WireguardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WireguardMapOutput)
}

type WireguardOutput struct{ *pulumi.OutputState }

func (WireguardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wireguard)(nil)).Elem()
}

func (o WireguardOutput) ToWireguardOutput() WireguardOutput {
	return o
}

func (o WireguardOutput) ToWireguardOutputWithContext(ctx context.Context) WireguardOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o WireguardOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o WireguardOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

func (o WireguardOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o WireguardOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Port for WireGuard service to listen on for incoming sessions.
func (o WireguardOutput) ListenPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.IntOutput { return v.ListenPort }).(pulumi.IntOutput)
}

// Layer3 Maximum transmission unit ('auto', 0 .. 65535)
func (o WireguardOutput) Mtu() pulumi.StringOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.StringOutput { return v.Mtu }).(pulumi.StringOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o WireguardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A base64 private key. If not specified, it will be automatically generated upon interface creation.
func (o WireguardOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// A base64 public key is calculated from the private key.
func (o WireguardOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

func (o WireguardOutput) Running() pulumi.BoolOutput {
	return o.ApplyT(func(v *Wireguard) pulumi.BoolOutput { return v.Running }).(pulumi.BoolOutput)
}

type WireguardArrayOutput struct{ *pulumi.OutputState }

func (WireguardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wireguard)(nil)).Elem()
}

func (o WireguardArrayOutput) ToWireguardArrayOutput() WireguardArrayOutput {
	return o
}

func (o WireguardArrayOutput) ToWireguardArrayOutputWithContext(ctx context.Context) WireguardArrayOutput {
	return o
}

func (o WireguardArrayOutput) Index(i pulumi.IntInput) WireguardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wireguard {
		return vs[0].([]*Wireguard)[vs[1].(int)]
	}).(WireguardOutput)
}

type WireguardMapOutput struct{ *pulumi.OutputState }

func (WireguardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wireguard)(nil)).Elem()
}

func (o WireguardMapOutput) ToWireguardMapOutput() WireguardMapOutput {
	return o
}

func (o WireguardMapOutput) ToWireguardMapOutputWithContext(ctx context.Context) WireguardMapOutput {
	return o
}

func (o WireguardMapOutput) MapIndex(k pulumi.StringInput) WireguardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wireguard {
		return vs[0].(map[string]*Wireguard)[vs[1].(string)]
	}).(WireguardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WireguardInput)(nil)).Elem(), &Wireguard{})
	pulumi.RegisterInputType(reflect.TypeOf((*WireguardArrayInput)(nil)).Elem(), WireguardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WireguardMapInput)(nil)).Elem(), WireguardMap{})
	pulumi.RegisterOutputType(WireguardOutput{})
	pulumi.RegisterOutputType(WireguardArrayOutput{})
	pulumi.RegisterOutputType(WireguardMapOutput{})
}
