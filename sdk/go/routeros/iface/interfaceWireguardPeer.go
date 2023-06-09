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
//			testWgInterface, err := Iface.NewInterfaceWireguard(ctx, "testWgInterface", &Iface.InterfaceWireguardArgs{
//				ListenPort: pulumi.Int(13231),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Iface.NewInterfaceWireguardPeer(ctx, "wgPeer", &Iface.InterfaceWireguardPeerArgs{
//				Interface: testWgInterface.Name,
//				PublicKey: pulumi.String("MY_BASE_64_PUBLIC_KEY"),
//				AllowedAddresses: pulumi.StringArray{
//					pulumi.String("192.168.0.0/16"),
//					pulumi.String("172.16.0.0/12"),
//					pulumi.String("10.0.0.0/8"),
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
// #The ID can be found via API or the terminal #The command for the terminal is -> :put [/interface/wireguard/peers get [print show-ids]]
//
// ```sh
//
//	$ pulumi import routeros:Iface/interfaceWireguardPeer:InterfaceWireguardPeer wg_peer "*0"
//
// ```
type InterfaceWireguardPeer struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
	AllowedAddresses pulumi.StringArrayOutput `pulumi:"allowedAddresses"`
	Comment          pulumi.StringPtrOutput   `pulumi:"comment"`
	// The most recent source IP address of correctly authenticated packets from the peer.
	CurrentEndpointAddress pulumi.StringOutput `pulumi:"currentEndpointAddress"`
	// The most recent source IP port of correctly authenticated packets from the peer.
	CurrentEndpointPort pulumi.IntOutput     `pulumi:"currentEndpointPort"`
	Disabled            pulumi.BoolPtrOutput `pulumi:"disabled"`
	// An endpoint IP or hostname can be left blank to allow remote connection from any address.
	EndpointAddress pulumi.StringOutput `pulumi:"endpointAddress"`
	// An endpoint port can be left blank to allow remote connection from any port.
	EndpointPort pulumi.StringOutput `pulumi:"endpointPort"`
	// Name of the interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Time in seconds after the last successful handshake.
	LastHandshake pulumi.StringOutput `pulumi:"lastHandshake"`
	// A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds.
	PersistentKeepalive pulumi.StringPtrOutput `pulumi:"persistentKeepalive"`
	// A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
	PresharedKey pulumi.StringPtrOutput `pulumi:"presharedKey"`
	// The remote peer's calculated public key.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The total amount of bytes received from the peer.
	Rx pulumi.StringOutput `pulumi:"rx"`
	// The total amount of bytes transmitted to the peer.
	Tx pulumi.StringOutput `pulumi:"tx"`
}

// NewInterfaceWireguardPeer registers a new resource with the given unique name, arguments, and options.
func NewInterfaceWireguardPeer(ctx *pulumi.Context,
	name string, args *InterfaceWireguardPeerArgs, opts ...pulumi.ResourceOption) (*InterfaceWireguardPeer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	if args.PresharedKey != nil {
		args.PresharedKey = pulumi.ToSecret(args.PresharedKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"presharedKey",
	})
	opts = append(opts, secrets)
	var resource InterfaceWireguardPeer
	err := ctx.RegisterResource("routeros:Iface/interfaceWireguardPeer:InterfaceWireguardPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterfaceWireguardPeer gets an existing InterfaceWireguardPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterfaceWireguardPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfaceWireguardPeerState, opts ...pulumi.ResourceOption) (*InterfaceWireguardPeer, error) {
	var resource InterfaceWireguardPeer
	err := ctx.ReadResource("routeros:Iface/interfaceWireguardPeer:InterfaceWireguardPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterfaceWireguardPeer resources.
type interfaceWireguardPeerState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
	AllowedAddresses []string `pulumi:"allowedAddresses"`
	Comment          *string  `pulumi:"comment"`
	// The most recent source IP address of correctly authenticated packets from the peer.
	CurrentEndpointAddress *string `pulumi:"currentEndpointAddress"`
	// The most recent source IP port of correctly authenticated packets from the peer.
	CurrentEndpointPort *int  `pulumi:"currentEndpointPort"`
	Disabled            *bool `pulumi:"disabled"`
	// An endpoint IP or hostname can be left blank to allow remote connection from any address.
	EndpointAddress *string `pulumi:"endpointAddress"`
	// An endpoint port can be left blank to allow remote connection from any port.
	EndpointPort *string `pulumi:"endpointPort"`
	// Name of the interface.
	Interface *string `pulumi:"interface"`
	// Time in seconds after the last successful handshake.
	LastHandshake *string `pulumi:"lastHandshake"`
	// A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds.
	PersistentKeepalive *string `pulumi:"persistentKeepalive"`
	// A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
	PresharedKey *string `pulumi:"presharedKey"`
	// The remote peer's calculated public key.
	PublicKey *string `pulumi:"publicKey"`
	// The total amount of bytes received from the peer.
	Rx *string `pulumi:"rx"`
	// The total amount of bytes transmitted to the peer.
	Tx *string `pulumi:"tx"`
}

type InterfaceWireguardPeerState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
	AllowedAddresses pulumi.StringArrayInput
	Comment          pulumi.StringPtrInput
	// The most recent source IP address of correctly authenticated packets from the peer.
	CurrentEndpointAddress pulumi.StringPtrInput
	// The most recent source IP port of correctly authenticated packets from the peer.
	CurrentEndpointPort pulumi.IntPtrInput
	Disabled            pulumi.BoolPtrInput
	// An endpoint IP or hostname can be left blank to allow remote connection from any address.
	EndpointAddress pulumi.StringPtrInput
	// An endpoint port can be left blank to allow remote connection from any port.
	EndpointPort pulumi.StringPtrInput
	// Name of the interface.
	Interface pulumi.StringPtrInput
	// Time in seconds after the last successful handshake.
	LastHandshake pulumi.StringPtrInput
	// A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds.
	PersistentKeepalive pulumi.StringPtrInput
	// A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
	PresharedKey pulumi.StringPtrInput
	// The remote peer's calculated public key.
	PublicKey pulumi.StringPtrInput
	// The total amount of bytes received from the peer.
	Rx pulumi.StringPtrInput
	// The total amount of bytes transmitted to the peer.
	Tx pulumi.StringPtrInput
}

func (InterfaceWireguardPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceWireguardPeerState)(nil)).Elem()
}

type interfaceWireguardPeerArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
	AllowedAddresses []string `pulumi:"allowedAddresses"`
	Comment          *string  `pulumi:"comment"`
	Disabled         *bool    `pulumi:"disabled"`
	// An endpoint IP or hostname can be left blank to allow remote connection from any address.
	EndpointAddress *string `pulumi:"endpointAddress"`
	// An endpoint port can be left blank to allow remote connection from any port.
	EndpointPort *string `pulumi:"endpointPort"`
	// Name of the interface.
	Interface string `pulumi:"interface"`
	// A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds.
	PersistentKeepalive *string `pulumi:"persistentKeepalive"`
	// A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
	PresharedKey *string `pulumi:"presharedKey"`
	// The remote peer's calculated public key.
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a InterfaceWireguardPeer resource.
type InterfaceWireguardPeerArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
	AllowedAddresses pulumi.StringArrayInput
	Comment          pulumi.StringPtrInput
	Disabled         pulumi.BoolPtrInput
	// An endpoint IP or hostname can be left blank to allow remote connection from any address.
	EndpointAddress pulumi.StringPtrInput
	// An endpoint port can be left blank to allow remote connection from any port.
	EndpointPort pulumi.StringPtrInput
	// Name of the interface.
	Interface pulumi.StringInput
	// A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds.
	PersistentKeepalive pulumi.StringPtrInput
	// A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
	PresharedKey pulumi.StringPtrInput
	// The remote peer's calculated public key.
	PublicKey pulumi.StringInput
}

func (InterfaceWireguardPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceWireguardPeerArgs)(nil)).Elem()
}

type InterfaceWireguardPeerInput interface {
	pulumi.Input

	ToInterfaceWireguardPeerOutput() InterfaceWireguardPeerOutput
	ToInterfaceWireguardPeerOutputWithContext(ctx context.Context) InterfaceWireguardPeerOutput
}

func (*InterfaceWireguardPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceWireguardPeer)(nil)).Elem()
}

func (i *InterfaceWireguardPeer) ToInterfaceWireguardPeerOutput() InterfaceWireguardPeerOutput {
	return i.ToInterfaceWireguardPeerOutputWithContext(context.Background())
}

func (i *InterfaceWireguardPeer) ToInterfaceWireguardPeerOutputWithContext(ctx context.Context) InterfaceWireguardPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceWireguardPeerOutput)
}

// InterfaceWireguardPeerArrayInput is an input type that accepts InterfaceWireguardPeerArray and InterfaceWireguardPeerArrayOutput values.
// You can construct a concrete instance of `InterfaceWireguardPeerArrayInput` via:
//
//	InterfaceWireguardPeerArray{ InterfaceWireguardPeerArgs{...} }
type InterfaceWireguardPeerArrayInput interface {
	pulumi.Input

	ToInterfaceWireguardPeerArrayOutput() InterfaceWireguardPeerArrayOutput
	ToInterfaceWireguardPeerArrayOutputWithContext(context.Context) InterfaceWireguardPeerArrayOutput
}

type InterfaceWireguardPeerArray []InterfaceWireguardPeerInput

func (InterfaceWireguardPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterfaceWireguardPeer)(nil)).Elem()
}

func (i InterfaceWireguardPeerArray) ToInterfaceWireguardPeerArrayOutput() InterfaceWireguardPeerArrayOutput {
	return i.ToInterfaceWireguardPeerArrayOutputWithContext(context.Background())
}

func (i InterfaceWireguardPeerArray) ToInterfaceWireguardPeerArrayOutputWithContext(ctx context.Context) InterfaceWireguardPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceWireguardPeerArrayOutput)
}

// InterfaceWireguardPeerMapInput is an input type that accepts InterfaceWireguardPeerMap and InterfaceWireguardPeerMapOutput values.
// You can construct a concrete instance of `InterfaceWireguardPeerMapInput` via:
//
//	InterfaceWireguardPeerMap{ "key": InterfaceWireguardPeerArgs{...} }
type InterfaceWireguardPeerMapInput interface {
	pulumi.Input

	ToInterfaceWireguardPeerMapOutput() InterfaceWireguardPeerMapOutput
	ToInterfaceWireguardPeerMapOutputWithContext(context.Context) InterfaceWireguardPeerMapOutput
}

type InterfaceWireguardPeerMap map[string]InterfaceWireguardPeerInput

func (InterfaceWireguardPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterfaceWireguardPeer)(nil)).Elem()
}

func (i InterfaceWireguardPeerMap) ToInterfaceWireguardPeerMapOutput() InterfaceWireguardPeerMapOutput {
	return i.ToInterfaceWireguardPeerMapOutputWithContext(context.Background())
}

func (i InterfaceWireguardPeerMap) ToInterfaceWireguardPeerMapOutputWithContext(ctx context.Context) InterfaceWireguardPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceWireguardPeerMapOutput)
}

type InterfaceWireguardPeerOutput struct{ *pulumi.OutputState }

func (InterfaceWireguardPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceWireguardPeer)(nil)).Elem()
}

func (o InterfaceWireguardPeerOutput) ToInterfaceWireguardPeerOutput() InterfaceWireguardPeerOutput {
	return o
}

func (o InterfaceWireguardPeerOutput) ToInterfaceWireguardPeerOutputWithContext(ctx context.Context) InterfaceWireguardPeerOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o InterfaceWireguardPeerOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o InterfaceWireguardPeerOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
func (o InterfaceWireguardPeerOutput) AllowedAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringArrayOutput { return v.AllowedAddresses }).(pulumi.StringArrayOutput)
}

func (o InterfaceWireguardPeerOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The most recent source IP address of correctly authenticated packets from the peer.
func (o InterfaceWireguardPeerOutput) CurrentEndpointAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.CurrentEndpointAddress }).(pulumi.StringOutput)
}

// The most recent source IP port of correctly authenticated packets from the peer.
func (o InterfaceWireguardPeerOutput) CurrentEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.IntOutput { return v.CurrentEndpointPort }).(pulumi.IntOutput)
}

func (o InterfaceWireguardPeerOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// An endpoint IP or hostname can be left blank to allow remote connection from any address.
func (o InterfaceWireguardPeerOutput) EndpointAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.EndpointAddress }).(pulumi.StringOutput)
}

// An endpoint port can be left blank to allow remote connection from any port.
func (o InterfaceWireguardPeerOutput) EndpointPort() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.EndpointPort }).(pulumi.StringOutput)
}

// Name of the interface.
func (o InterfaceWireguardPeerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Time in seconds after the last successful handshake.
func (o InterfaceWireguardPeerOutput) LastHandshake() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.LastHandshake }).(pulumi.StringOutput)
}

// A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds.
func (o InterfaceWireguardPeerOutput) PersistentKeepalive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringPtrOutput { return v.PersistentKeepalive }).(pulumi.StringPtrOutput)
}

// A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
func (o InterfaceWireguardPeerOutput) PresharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringPtrOutput { return v.PresharedKey }).(pulumi.StringPtrOutput)
}

// The remote peer's calculated public key.
func (o InterfaceWireguardPeerOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// The total amount of bytes received from the peer.
func (o InterfaceWireguardPeerOutput) Rx() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.Rx }).(pulumi.StringOutput)
}

// The total amount of bytes transmitted to the peer.
func (o InterfaceWireguardPeerOutput) Tx() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceWireguardPeer) pulumi.StringOutput { return v.Tx }).(pulumi.StringOutput)
}

type InterfaceWireguardPeerArrayOutput struct{ *pulumi.OutputState }

func (InterfaceWireguardPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterfaceWireguardPeer)(nil)).Elem()
}

func (o InterfaceWireguardPeerArrayOutput) ToInterfaceWireguardPeerArrayOutput() InterfaceWireguardPeerArrayOutput {
	return o
}

func (o InterfaceWireguardPeerArrayOutput) ToInterfaceWireguardPeerArrayOutputWithContext(ctx context.Context) InterfaceWireguardPeerArrayOutput {
	return o
}

func (o InterfaceWireguardPeerArrayOutput) Index(i pulumi.IntInput) InterfaceWireguardPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InterfaceWireguardPeer {
		return vs[0].([]*InterfaceWireguardPeer)[vs[1].(int)]
	}).(InterfaceWireguardPeerOutput)
}

type InterfaceWireguardPeerMapOutput struct{ *pulumi.OutputState }

func (InterfaceWireguardPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterfaceWireguardPeer)(nil)).Elem()
}

func (o InterfaceWireguardPeerMapOutput) ToInterfaceWireguardPeerMapOutput() InterfaceWireguardPeerMapOutput {
	return o
}

func (o InterfaceWireguardPeerMapOutput) ToInterfaceWireguardPeerMapOutputWithContext(ctx context.Context) InterfaceWireguardPeerMapOutput {
	return o
}

func (o InterfaceWireguardPeerMapOutput) MapIndex(k pulumi.StringInput) InterfaceWireguardPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InterfaceWireguardPeer {
		return vs[0].(map[string]*InterfaceWireguardPeer)[vs[1].(string)]
	}).(InterfaceWireguardPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceWireguardPeerInput)(nil)).Elem(), &InterfaceWireguardPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceWireguardPeerArrayInput)(nil)).Elem(), InterfaceWireguardPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceWireguardPeerMapInput)(nil)).Elem(), InterfaceWireguardPeerMap{})
	pulumi.RegisterOutputType(InterfaceWireguardPeerOutput{})
	pulumi.RegisterOutputType(InterfaceWireguardPeerArrayOutput{})
	pulumi.RegisterOutputType(InterfaceWireguardPeerMapOutput{})
}
