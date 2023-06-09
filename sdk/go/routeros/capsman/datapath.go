// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package capsman

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-routeros/sdk/go/routeros/CapsMan"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := CapsMan.NewDatapath(ctx, "testDatapath", &CapsMan.DatapathArgs{
//				Arp:                      pulumi.String("local-proxy-arp"),
//				Bridge:                   pulumi.String("bridge"),
//				BridgeCost:               pulumi.Int(100),
//				BridgeHorizon:            pulumi.Int(200),
//				ClientToClientForwarding: pulumi.Bool(true),
//				Comment:                  pulumi.String("test_datapath"),
//				InterfaceList:            pulumi.String("static"),
//				L2mtu:                    pulumi.Int(1450),
//				LocalForwarding:          pulumi.Bool(true),
//				Mtu:                      pulumi.Int(1500),
//				VlanId:                   pulumi.Int(101),
//				VlanMode:                 pulumi.String("no-tag"),
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
// # Import with the name of the CAPsMAN datapath configuration in case of the example use test-datapath-config
//
// ```sh
//
//	$ pulumi import routeros:CapsMan/datapath:Datapath test_datapath test-datapath-config
//
// ```
type Datapath struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// ARP mode. See [docs](https://wiki.mikrotik.com/wiki/Manual:IP/ARP#ARP_Modes) for info.
	Arp pulumi.StringPtrOutput `pulumi:"arp"`
	// Bridge to which particular interface should be automatically added as port. Required only when local-forwarding is not used.
	Bridge pulumi.StringPtrOutput `pulumi:"bridge"`
	// Bridge port cost to use when adding as bridge port.
	BridgeCost pulumi.IntPtrOutput `pulumi:"bridgeCost"`
	// Bridge horizon to use when adding as bridge port.
	BridgeHorizon pulumi.IntPtrOutput `pulumi:"bridgeHorizon"`
	// Controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN.
	ClientToClientForwarding pulumi.BoolPtrOutput   `pulumi:"clientToClientForwarding"`
	Comment                  pulumi.StringPtrOutput `pulumi:"comment"`
	// Interface list name.
	InterfaceList pulumi.StringPtrOutput `pulumi:"interfaceList"`
	// Layer2 MTU size.
	L2mtu pulumi.IntPtrOutput `pulumi:"l2mtu"`
	// Controls forwarding mode. If disabled, all L2 and L3 data will be forwarded to CAPsMAN, and further forwarding decisions will be made only then. See [docs](https://wiki.mikrotik.com/wiki/Manual:CAPsMAN#Local_Forwarding_Mode) for info.
	LocalForwarding pulumi.BoolPtrOutput `pulumi:"localForwarding"`
	// MTU size.
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringOutput `pulumi:"name"`
	// OpenFlow switch to add interface to, as port when enabled.
	OpenflowSwitch pulumi.StringPtrOutput `pulumi:"openflowSwitch"`
	// VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging.
	VlanId pulumi.IntPtrOutput `pulumi:"vlanId"`
	// VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)
	VlanMode pulumi.StringPtrOutput `pulumi:"vlanMode"`
}

// NewDatapath registers a new resource with the given unique name, arguments, and options.
func NewDatapath(ctx *pulumi.Context,
	name string, args *DatapathArgs, opts ...pulumi.ResourceOption) (*Datapath, error) {
	if args == nil {
		args = &DatapathArgs{}
	}

	var resource Datapath
	err := ctx.RegisterResource("routeros:CapsMan/datapath:Datapath", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatapath gets an existing Datapath resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatapath(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatapathState, opts ...pulumi.ResourceOption) (*Datapath, error) {
	var resource Datapath
	err := ctx.ReadResource("routeros:CapsMan/datapath:Datapath", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datapath resources.
type datapathState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// ARP mode. See [docs](https://wiki.mikrotik.com/wiki/Manual:IP/ARP#ARP_Modes) for info.
	Arp *string `pulumi:"arp"`
	// Bridge to which particular interface should be automatically added as port. Required only when local-forwarding is not used.
	Bridge *string `pulumi:"bridge"`
	// Bridge port cost to use when adding as bridge port.
	BridgeCost *int `pulumi:"bridgeCost"`
	// Bridge horizon to use when adding as bridge port.
	BridgeHorizon *int `pulumi:"bridgeHorizon"`
	// Controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN.
	ClientToClientForwarding *bool   `pulumi:"clientToClientForwarding"`
	Comment                  *string `pulumi:"comment"`
	// Interface list name.
	InterfaceList *string `pulumi:"interfaceList"`
	// Layer2 MTU size.
	L2mtu *int `pulumi:"l2mtu"`
	// Controls forwarding mode. If disabled, all L2 and L3 data will be forwarded to CAPsMAN, and further forwarding decisions will be made only then. See [docs](https://wiki.mikrotik.com/wiki/Manual:CAPsMAN#Local_Forwarding_Mode) for info.
	LocalForwarding *bool `pulumi:"localForwarding"`
	// MTU size.
	Mtu *int `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// OpenFlow switch to add interface to, as port when enabled.
	OpenflowSwitch *string `pulumi:"openflowSwitch"`
	// VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging.
	VlanId *int `pulumi:"vlanId"`
	// VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)
	VlanMode *string `pulumi:"vlanMode"`
}

type DatapathState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// ARP mode. See [docs](https://wiki.mikrotik.com/wiki/Manual:IP/ARP#ARP_Modes) for info.
	Arp pulumi.StringPtrInput
	// Bridge to which particular interface should be automatically added as port. Required only when local-forwarding is not used.
	Bridge pulumi.StringPtrInput
	// Bridge port cost to use when adding as bridge port.
	BridgeCost pulumi.IntPtrInput
	// Bridge horizon to use when adding as bridge port.
	BridgeHorizon pulumi.IntPtrInput
	// Controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN.
	ClientToClientForwarding pulumi.BoolPtrInput
	Comment                  pulumi.StringPtrInput
	// Interface list name.
	InterfaceList pulumi.StringPtrInput
	// Layer2 MTU size.
	L2mtu pulumi.IntPtrInput
	// Controls forwarding mode. If disabled, all L2 and L3 data will be forwarded to CAPsMAN, and further forwarding decisions will be made only then. See [docs](https://wiki.mikrotik.com/wiki/Manual:CAPsMAN#Local_Forwarding_Mode) for info.
	LocalForwarding pulumi.BoolPtrInput
	// MTU size.
	Mtu pulumi.IntPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// OpenFlow switch to add interface to, as port when enabled.
	OpenflowSwitch pulumi.StringPtrInput
	// VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging.
	VlanId pulumi.IntPtrInput
	// VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)
	VlanMode pulumi.StringPtrInput
}

func (DatapathState) ElementType() reflect.Type {
	return reflect.TypeOf((*datapathState)(nil)).Elem()
}

type datapathArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// ARP mode. See [docs](https://wiki.mikrotik.com/wiki/Manual:IP/ARP#ARP_Modes) for info.
	Arp *string `pulumi:"arp"`
	// Bridge to which particular interface should be automatically added as port. Required only when local-forwarding is not used.
	Bridge *string `pulumi:"bridge"`
	// Bridge port cost to use when adding as bridge port.
	BridgeCost *int `pulumi:"bridgeCost"`
	// Bridge horizon to use when adding as bridge port.
	BridgeHorizon *int `pulumi:"bridgeHorizon"`
	// Controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN.
	ClientToClientForwarding *bool   `pulumi:"clientToClientForwarding"`
	Comment                  *string `pulumi:"comment"`
	// Interface list name.
	InterfaceList *string `pulumi:"interfaceList"`
	// Layer2 MTU size.
	L2mtu *int `pulumi:"l2mtu"`
	// Controls forwarding mode. If disabled, all L2 and L3 data will be forwarded to CAPsMAN, and further forwarding decisions will be made only then. See [docs](https://wiki.mikrotik.com/wiki/Manual:CAPsMAN#Local_Forwarding_Mode) for info.
	LocalForwarding *bool `pulumi:"localForwarding"`
	// MTU size.
	Mtu *int `pulumi:"mtu"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// OpenFlow switch to add interface to, as port when enabled.
	OpenflowSwitch *string `pulumi:"openflowSwitch"`
	// VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging.
	VlanId *int `pulumi:"vlanId"`
	// VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)
	VlanMode *string `pulumi:"vlanMode"`
}

// The set of arguments for constructing a Datapath resource.
type DatapathArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// ARP mode. See [docs](https://wiki.mikrotik.com/wiki/Manual:IP/ARP#ARP_Modes) for info.
	Arp pulumi.StringPtrInput
	// Bridge to which particular interface should be automatically added as port. Required only when local-forwarding is not used.
	Bridge pulumi.StringPtrInput
	// Bridge port cost to use when adding as bridge port.
	BridgeCost pulumi.IntPtrInput
	// Bridge horizon to use when adding as bridge port.
	BridgeHorizon pulumi.IntPtrInput
	// Controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN.
	ClientToClientForwarding pulumi.BoolPtrInput
	Comment                  pulumi.StringPtrInput
	// Interface list name.
	InterfaceList pulumi.StringPtrInput
	// Layer2 MTU size.
	L2mtu pulumi.IntPtrInput
	// Controls forwarding mode. If disabled, all L2 and L3 data will be forwarded to CAPsMAN, and further forwarding decisions will be made only then. See [docs](https://wiki.mikrotik.com/wiki/Manual:CAPsMAN#Local_Forwarding_Mode) for info.
	LocalForwarding pulumi.BoolPtrInput
	// MTU size.
	Mtu pulumi.IntPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// OpenFlow switch to add interface to, as port when enabled.
	OpenflowSwitch pulumi.StringPtrInput
	// VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging.
	VlanId pulumi.IntPtrInput
	// VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)
	VlanMode pulumi.StringPtrInput
}

func (DatapathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datapathArgs)(nil)).Elem()
}

type DatapathInput interface {
	pulumi.Input

	ToDatapathOutput() DatapathOutput
	ToDatapathOutputWithContext(ctx context.Context) DatapathOutput
}

func (*Datapath) ElementType() reflect.Type {
	return reflect.TypeOf((**Datapath)(nil)).Elem()
}

func (i *Datapath) ToDatapathOutput() DatapathOutput {
	return i.ToDatapathOutputWithContext(context.Background())
}

func (i *Datapath) ToDatapathOutputWithContext(ctx context.Context) DatapathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatapathOutput)
}

// DatapathArrayInput is an input type that accepts DatapathArray and DatapathArrayOutput values.
// You can construct a concrete instance of `DatapathArrayInput` via:
//
//	DatapathArray{ DatapathArgs{...} }
type DatapathArrayInput interface {
	pulumi.Input

	ToDatapathArrayOutput() DatapathArrayOutput
	ToDatapathArrayOutputWithContext(context.Context) DatapathArrayOutput
}

type DatapathArray []DatapathInput

func (DatapathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datapath)(nil)).Elem()
}

func (i DatapathArray) ToDatapathArrayOutput() DatapathArrayOutput {
	return i.ToDatapathArrayOutputWithContext(context.Background())
}

func (i DatapathArray) ToDatapathArrayOutputWithContext(ctx context.Context) DatapathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatapathArrayOutput)
}

// DatapathMapInput is an input type that accepts DatapathMap and DatapathMapOutput values.
// You can construct a concrete instance of `DatapathMapInput` via:
//
//	DatapathMap{ "key": DatapathArgs{...} }
type DatapathMapInput interface {
	pulumi.Input

	ToDatapathMapOutput() DatapathMapOutput
	ToDatapathMapOutputWithContext(context.Context) DatapathMapOutput
}

type DatapathMap map[string]DatapathInput

func (DatapathMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datapath)(nil)).Elem()
}

func (i DatapathMap) ToDatapathMapOutput() DatapathMapOutput {
	return i.ToDatapathMapOutputWithContext(context.Background())
}

func (i DatapathMap) ToDatapathMapOutputWithContext(ctx context.Context) DatapathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatapathMapOutput)
}

type DatapathOutput struct{ *pulumi.OutputState }

func (DatapathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datapath)(nil)).Elem()
}

func (o DatapathOutput) ToDatapathOutput() DatapathOutput {
	return o
}

func (o DatapathOutput) ToDatapathOutputWithContext(ctx context.Context) DatapathOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o DatapathOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o DatapathOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// ARP mode. See [docs](https://wiki.mikrotik.com/wiki/Manual:IP/ARP#ARP_Modes) for info.
func (o DatapathOutput) Arp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringPtrOutput { return v.Arp }).(pulumi.StringPtrOutput)
}

// Bridge to which particular interface should be automatically added as port. Required only when local-forwarding is not used.
func (o DatapathOutput) Bridge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringPtrOutput { return v.Bridge }).(pulumi.StringPtrOutput)
}

// Bridge port cost to use when adding as bridge port.
func (o DatapathOutput) BridgeCost() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.IntPtrOutput { return v.BridgeCost }).(pulumi.IntPtrOutput)
}

// Bridge horizon to use when adding as bridge port.
func (o DatapathOutput) BridgeHorizon() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.IntPtrOutput { return v.BridgeHorizon }).(pulumi.IntPtrOutput)
}

// Controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN.
func (o DatapathOutput) ClientToClientForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.BoolPtrOutput { return v.ClientToClientForwarding }).(pulumi.BoolPtrOutput)
}

func (o DatapathOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Interface list name.
func (o DatapathOutput) InterfaceList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringPtrOutput { return v.InterfaceList }).(pulumi.StringPtrOutput)
}

// Layer2 MTU size.
func (o DatapathOutput) L2mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.IntPtrOutput { return v.L2mtu }).(pulumi.IntPtrOutput)
}

// Controls forwarding mode. If disabled, all L2 and L3 data will be forwarded to CAPsMAN, and further forwarding decisions will be made only then. See [docs](https://wiki.mikrotik.com/wiki/Manual:CAPsMAN#Local_Forwarding_Mode) for info.
func (o DatapathOutput) LocalForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.BoolPtrOutput { return v.LocalForwarding }).(pulumi.BoolPtrOutput)
}

// MTU size.
func (o DatapathOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.IntPtrOutput { return v.Mtu }).(pulumi.IntPtrOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o DatapathOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OpenFlow switch to add interface to, as port when enabled.
func (o DatapathOutput) OpenflowSwitch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringPtrOutput { return v.OpenflowSwitch }).(pulumi.StringPtrOutput)
}

// VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging.
func (o DatapathOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.IntPtrOutput { return v.VlanId }).(pulumi.IntPtrOutput)
}

// VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)
func (o DatapathOutput) VlanMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datapath) pulumi.StringPtrOutput { return v.VlanMode }).(pulumi.StringPtrOutput)
}

type DatapathArrayOutput struct{ *pulumi.OutputState }

func (DatapathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datapath)(nil)).Elem()
}

func (o DatapathArrayOutput) ToDatapathArrayOutput() DatapathArrayOutput {
	return o
}

func (o DatapathArrayOutput) ToDatapathArrayOutputWithContext(ctx context.Context) DatapathArrayOutput {
	return o
}

func (o DatapathArrayOutput) Index(i pulumi.IntInput) DatapathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Datapath {
		return vs[0].([]*Datapath)[vs[1].(int)]
	}).(DatapathOutput)
}

type DatapathMapOutput struct{ *pulumi.OutputState }

func (DatapathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datapath)(nil)).Elem()
}

func (o DatapathMapOutput) ToDatapathMapOutput() DatapathMapOutput {
	return o
}

func (o DatapathMapOutput) ToDatapathMapOutputWithContext(ctx context.Context) DatapathMapOutput {
	return o
}

func (o DatapathMapOutput) MapIndex(k pulumi.StringInput) DatapathOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Datapath {
		return vs[0].(map[string]*Datapath)[vs[1].(string)]
	}).(DatapathOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatapathInput)(nil)).Elem(), &Datapath{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatapathArrayInput)(nil)).Elem(), DatapathArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatapathMapInput)(nil)).Elem(), DatapathMap{})
	pulumi.RegisterOutputType(DatapathOutput{})
	pulumi.RegisterOutputType(DatapathArrayOutput{})
	pulumi.RegisterOutputType(DatapathMapOutput{})
}
