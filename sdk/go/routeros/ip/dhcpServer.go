// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Ip.DhcpServer (Resource)
//
// ***
//
// #### This is an alias for backwards compatibility between plugin versions.
// Please see documentation for Ip.DhcpIpServer
type DhcpServer struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// Whether to add dynamic ARP entry.
	AddArp pulumi.BoolPtrOutput `pulumi:"addArp"`
	// IP pool, from which to take IP addresses for the clients. If set to static-only, then only the clients that have a
	// static lease (added in lease submenu) will be allowed.
	AddressPool pulumi.StringPtrOutput `pulumi:"addressPool"`
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	// Requires IPv6 DHCP Server to have this option enabled as well to work properly.
	AllowDualStackQueue pulumi.BoolPtrOutput `pulumi:"allowDualStackQueue"`
	// Always send replies as broadcasts even if destination IP is known.
	AlwaysBroadcast pulumi.BoolPtrOutput `pulumi:"alwaysBroadcast"`
	// Option changes the way how a server responds to DHCP requests.
	Authoritative pulumi.StringPtrOutput `pulumi:"authoritative"`
	// Accepts two predefined options or time value: * forever - lease never expires * lease-time - use time from lease-time
	// parameter
	BootpLeaseTime pulumi.StringPtrOutput `pulumi:"bootpLeaseTime"`
	// Support for BOOTP clients.
	BootpSupport pulumi.StringPtrOutput `pulumi:"bootpSupport"`
	// Specifies whether to limit specific number of clients per single MAC address.
	ClientMacLimit pulumi.IntPtrOutput    `pulumi:"clientMacLimit"`
	Comment        pulumi.StringPtrOutput `pulumi:"comment"`
	// Allows to disable/enable conflict detection. If option is enabled, then whenever server tries to assign a lease it will
	// send ICMP and ARP messages to detect whether such address in the network already exist. If any of above get reply
	// address is considered already used. Conflict detection must be disabled when any kind of DHCP client limitation per port
	// or per mac is used.
	ConflictDetection pulumi.BoolPtrOutput `pulumi:"conflictDetection"`
	// If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. If set to none - there is no
	// threshold (all DHCP packets are processed).
	DelayThreshold pulumi.StringPtrOutput `pulumi:"delayThreshold"`
	// Use custom set of DHCP options defined in option sets menu.
	DhcpOptionSet pulumi.StringPtrOutput `pulumi:"dhcpOptionSet"`
	Disabled      pulumi.BoolPtrOutput   `pulumi:"disabled"`
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic pulumi.BoolOutput `pulumi:"dynamic"`
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore pulumi.StringPtrOutput `pulumi:"insertQueueBefore"`
	// Name of the interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	Invalid   pulumi.BoolOutput   `pulumi:"invalid"`
	// A script that will be executed after a lease is assigned or de-assigned.
	LeaseScript pulumi.StringPtrOutput `pulumi:"leaseScript"`
	// The time that a client may use the assigned address. The client will try to renew this address after half of this time
	// and will request a new address after the time limit expires.
	LeaseTime pulumi.StringPtrOutput `pulumi:"leaseTime"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name        pulumi.StringOutput    `pulumi:"name"`
	ParentQueue pulumi.StringPtrOutput `pulumi:"parentQueue"`
	// The IP address of the relay this DHCP server.
	Relay pulumi.StringPtrOutput `pulumi:"relay"`
	// The address which the DHCP client must send requests to in order to renew an IP address lease.
	SrcAddress pulumi.StringPtrOutput `pulumi:"srcAddress"`
	// Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client.
	UseFramedAsClassless pulumi.BoolPtrOutput `pulumi:"useFramedAsClassless"`
	// Whether to use RADIUS server.
	UseRadius pulumi.StringPtrOutput `pulumi:"useRadius"`
}

// NewDhcpServer registers a new resource with the given unique name, arguments, and options.
func NewDhcpServer(ctx *pulumi.Context,
	name string, args *DhcpServerArgs, opts ...pulumi.ResourceOption) (*DhcpServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	var resource DhcpServer
	err := ctx.RegisterResource("routeros:Ip/dhcpServer:DhcpServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpServer gets an existing DhcpServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpServerState, opts ...pulumi.ResourceOption) (*DhcpServer, error) {
	var resource DhcpServer
	err := ctx.ReadResource("routeros:Ip/dhcpServer:DhcpServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpServer resources.
type dhcpServerState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Whether to add dynamic ARP entry.
	AddArp *bool `pulumi:"addArp"`
	// IP pool, from which to take IP addresses for the clients. If set to static-only, then only the clients that have a
	// static lease (added in lease submenu) will be allowed.
	AddressPool *string `pulumi:"addressPool"`
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	// Requires IPv6 DHCP Server to have this option enabled as well to work properly.
	AllowDualStackQueue *bool `pulumi:"allowDualStackQueue"`
	// Always send replies as broadcasts even if destination IP is known.
	AlwaysBroadcast *bool `pulumi:"alwaysBroadcast"`
	// Option changes the way how a server responds to DHCP requests.
	Authoritative *string `pulumi:"authoritative"`
	// Accepts two predefined options or time value: * forever - lease never expires * lease-time - use time from lease-time
	// parameter
	BootpLeaseTime *string `pulumi:"bootpLeaseTime"`
	// Support for BOOTP clients.
	BootpSupport *string `pulumi:"bootpSupport"`
	// Specifies whether to limit specific number of clients per single MAC address.
	ClientMacLimit *int    `pulumi:"clientMacLimit"`
	Comment        *string `pulumi:"comment"`
	// Allows to disable/enable conflict detection. If option is enabled, then whenever server tries to assign a lease it will
	// send ICMP and ARP messages to detect whether such address in the network already exist. If any of above get reply
	// address is considered already used. Conflict detection must be disabled when any kind of DHCP client limitation per port
	// or per mac is used.
	ConflictDetection *bool `pulumi:"conflictDetection"`
	// If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. If set to none - there is no
	// threshold (all DHCP packets are processed).
	DelayThreshold *string `pulumi:"delayThreshold"`
	// Use custom set of DHCP options defined in option sets menu.
	DhcpOptionSet *string `pulumi:"dhcpOptionSet"`
	Disabled      *bool   `pulumi:"disabled"`
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic *bool `pulumi:"dynamic"`
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore *string `pulumi:"insertQueueBefore"`
	// Name of the interface.
	Interface *string `pulumi:"interface"`
	Invalid   *bool   `pulumi:"invalid"`
	// A script that will be executed after a lease is assigned or de-assigned.
	LeaseScript *string `pulumi:"leaseScript"`
	// The time that a client may use the assigned address. The client will try to renew this address after half of this time
	// and will request a new address after the time limit expires.
	LeaseTime *string `pulumi:"leaseTime"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name        *string `pulumi:"name"`
	ParentQueue *string `pulumi:"parentQueue"`
	// The IP address of the relay this DHCP server.
	Relay *string `pulumi:"relay"`
	// The address which the DHCP client must send requests to in order to renew an IP address lease.
	SrcAddress *string `pulumi:"srcAddress"`
	// Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client.
	UseFramedAsClassless *bool `pulumi:"useFramedAsClassless"`
	// Whether to use RADIUS server.
	UseRadius *string `pulumi:"useRadius"`
}

type DhcpServerState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Whether to add dynamic ARP entry.
	AddArp pulumi.BoolPtrInput
	// IP pool, from which to take IP addresses for the clients. If set to static-only, then only the clients that have a
	// static lease (added in lease submenu) will be allowed.
	AddressPool pulumi.StringPtrInput
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	// Requires IPv6 DHCP Server to have this option enabled as well to work properly.
	AllowDualStackQueue pulumi.BoolPtrInput
	// Always send replies as broadcasts even if destination IP is known.
	AlwaysBroadcast pulumi.BoolPtrInput
	// Option changes the way how a server responds to DHCP requests.
	Authoritative pulumi.StringPtrInput
	// Accepts two predefined options or time value: * forever - lease never expires * lease-time - use time from lease-time
	// parameter
	BootpLeaseTime pulumi.StringPtrInput
	// Support for BOOTP clients.
	BootpSupport pulumi.StringPtrInput
	// Specifies whether to limit specific number of clients per single MAC address.
	ClientMacLimit pulumi.IntPtrInput
	Comment        pulumi.StringPtrInput
	// Allows to disable/enable conflict detection. If option is enabled, then whenever server tries to assign a lease it will
	// send ICMP and ARP messages to detect whether such address in the network already exist. If any of above get reply
	// address is considered already used. Conflict detection must be disabled when any kind of DHCP client limitation per port
	// or per mac is used.
	ConflictDetection pulumi.BoolPtrInput
	// If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. If set to none - there is no
	// threshold (all DHCP packets are processed).
	DelayThreshold pulumi.StringPtrInput
	// Use custom set of DHCP options defined in option sets menu.
	DhcpOptionSet pulumi.StringPtrInput
	Disabled      pulumi.BoolPtrInput
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic pulumi.BoolPtrInput
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore pulumi.StringPtrInput
	// Name of the interface.
	Interface pulumi.StringPtrInput
	Invalid   pulumi.BoolPtrInput
	// A script that will be executed after a lease is assigned or de-assigned.
	LeaseScript pulumi.StringPtrInput
	// The time that a client may use the assigned address. The client will try to renew this address after half of this time
	// and will request a new address after the time limit expires.
	LeaseTime pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name        pulumi.StringPtrInput
	ParentQueue pulumi.StringPtrInput
	// The IP address of the relay this DHCP server.
	Relay pulumi.StringPtrInput
	// The address which the DHCP client must send requests to in order to renew an IP address lease.
	SrcAddress pulumi.StringPtrInput
	// Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client.
	UseFramedAsClassless pulumi.BoolPtrInput
	// Whether to use RADIUS server.
	UseRadius pulumi.StringPtrInput
}

func (DhcpServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpServerState)(nil)).Elem()
}

type dhcpServerArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Whether to add dynamic ARP entry.
	AddArp *bool `pulumi:"addArp"`
	// IP pool, from which to take IP addresses for the clients. If set to static-only, then only the clients that have a
	// static lease (added in lease submenu) will be allowed.
	AddressPool *string `pulumi:"addressPool"`
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	// Requires IPv6 DHCP Server to have this option enabled as well to work properly.
	AllowDualStackQueue *bool `pulumi:"allowDualStackQueue"`
	// Always send replies as broadcasts even if destination IP is known.
	AlwaysBroadcast *bool `pulumi:"alwaysBroadcast"`
	// Option changes the way how a server responds to DHCP requests.
	Authoritative *string `pulumi:"authoritative"`
	// Accepts two predefined options or time value: * forever - lease never expires * lease-time - use time from lease-time
	// parameter
	BootpLeaseTime *string `pulumi:"bootpLeaseTime"`
	// Support for BOOTP clients.
	BootpSupport *string `pulumi:"bootpSupport"`
	// Specifies whether to limit specific number of clients per single MAC address.
	ClientMacLimit *int    `pulumi:"clientMacLimit"`
	Comment        *string `pulumi:"comment"`
	// Allows to disable/enable conflict detection. If option is enabled, then whenever server tries to assign a lease it will
	// send ICMP and ARP messages to detect whether such address in the network already exist. If any of above get reply
	// address is considered already used. Conflict detection must be disabled when any kind of DHCP client limitation per port
	// or per mac is used.
	ConflictDetection *bool `pulumi:"conflictDetection"`
	// If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. If set to none - there is no
	// threshold (all DHCP packets are processed).
	DelayThreshold *string `pulumi:"delayThreshold"`
	// Use custom set of DHCP options defined in option sets menu.
	DhcpOptionSet *string `pulumi:"dhcpOptionSet"`
	Disabled      *bool   `pulumi:"disabled"`
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore *string `pulumi:"insertQueueBefore"`
	// Name of the interface.
	Interface string `pulumi:"interface"`
	// A script that will be executed after a lease is assigned or de-assigned.
	LeaseScript *string `pulumi:"leaseScript"`
	// The time that a client may use the assigned address. The client will try to renew this address after half of this time
	// and will request a new address after the time limit expires.
	LeaseTime *string `pulumi:"leaseTime"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name        *string `pulumi:"name"`
	ParentQueue *string `pulumi:"parentQueue"`
	// The IP address of the relay this DHCP server.
	Relay *string `pulumi:"relay"`
	// The address which the DHCP client must send requests to in order to renew an IP address lease.
	SrcAddress *string `pulumi:"srcAddress"`
	// Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client.
	UseFramedAsClassless *bool `pulumi:"useFramedAsClassless"`
	// Whether to use RADIUS server.
	UseRadius *string `pulumi:"useRadius"`
}

// The set of arguments for constructing a DhcpServer resource.
type DhcpServerArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Whether to add dynamic ARP entry.
	AddArp pulumi.BoolPtrInput
	// IP pool, from which to take IP addresses for the clients. If set to static-only, then only the clients that have a
	// static lease (added in lease submenu) will be allowed.
	AddressPool pulumi.StringPtrInput
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	// Requires IPv6 DHCP Server to have this option enabled as well to work properly.
	AllowDualStackQueue pulumi.BoolPtrInput
	// Always send replies as broadcasts even if destination IP is known.
	AlwaysBroadcast pulumi.BoolPtrInput
	// Option changes the way how a server responds to DHCP requests.
	Authoritative pulumi.StringPtrInput
	// Accepts two predefined options or time value: * forever - lease never expires * lease-time - use time from lease-time
	// parameter
	BootpLeaseTime pulumi.StringPtrInput
	// Support for BOOTP clients.
	BootpSupport pulumi.StringPtrInput
	// Specifies whether to limit specific number of clients per single MAC address.
	ClientMacLimit pulumi.IntPtrInput
	Comment        pulumi.StringPtrInput
	// Allows to disable/enable conflict detection. If option is enabled, then whenever server tries to assign a lease it will
	// send ICMP and ARP messages to detect whether such address in the network already exist. If any of above get reply
	// address is considered already used. Conflict detection must be disabled when any kind of DHCP client limitation per port
	// or per mac is used.
	ConflictDetection pulumi.BoolPtrInput
	// If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. If set to none - there is no
	// threshold (all DHCP packets are processed).
	DelayThreshold pulumi.StringPtrInput
	// Use custom set of DHCP options defined in option sets menu.
	DhcpOptionSet pulumi.StringPtrInput
	Disabled      pulumi.BoolPtrInput
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore pulumi.StringPtrInput
	// Name of the interface.
	Interface pulumi.StringInput
	// A script that will be executed after a lease is assigned or de-assigned.
	LeaseScript pulumi.StringPtrInput
	// The time that a client may use the assigned address. The client will try to renew this address after half of this time
	// and will request a new address after the time limit expires.
	LeaseTime pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name        pulumi.StringPtrInput
	ParentQueue pulumi.StringPtrInput
	// The IP address of the relay this DHCP server.
	Relay pulumi.StringPtrInput
	// The address which the DHCP client must send requests to in order to renew an IP address lease.
	SrcAddress pulumi.StringPtrInput
	// Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client.
	UseFramedAsClassless pulumi.BoolPtrInput
	// Whether to use RADIUS server.
	UseRadius pulumi.StringPtrInput
}

func (DhcpServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpServerArgs)(nil)).Elem()
}

type DhcpServerInput interface {
	pulumi.Input

	ToDhcpServerOutput() DhcpServerOutput
	ToDhcpServerOutputWithContext(ctx context.Context) DhcpServerOutput
}

func (*DhcpServer) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpServer)(nil)).Elem()
}

func (i *DhcpServer) ToDhcpServerOutput() DhcpServerOutput {
	return i.ToDhcpServerOutputWithContext(context.Background())
}

func (i *DhcpServer) ToDhcpServerOutputWithContext(ctx context.Context) DhcpServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpServerOutput)
}

// DhcpServerArrayInput is an input type that accepts DhcpServerArray and DhcpServerArrayOutput values.
// You can construct a concrete instance of `DhcpServerArrayInput` via:
//
//	DhcpServerArray{ DhcpServerArgs{...} }
type DhcpServerArrayInput interface {
	pulumi.Input

	ToDhcpServerArrayOutput() DhcpServerArrayOutput
	ToDhcpServerArrayOutputWithContext(context.Context) DhcpServerArrayOutput
}

type DhcpServerArray []DhcpServerInput

func (DhcpServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpServer)(nil)).Elem()
}

func (i DhcpServerArray) ToDhcpServerArrayOutput() DhcpServerArrayOutput {
	return i.ToDhcpServerArrayOutputWithContext(context.Background())
}

func (i DhcpServerArray) ToDhcpServerArrayOutputWithContext(ctx context.Context) DhcpServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpServerArrayOutput)
}

// DhcpServerMapInput is an input type that accepts DhcpServerMap and DhcpServerMapOutput values.
// You can construct a concrete instance of `DhcpServerMapInput` via:
//
//	DhcpServerMap{ "key": DhcpServerArgs{...} }
type DhcpServerMapInput interface {
	pulumi.Input

	ToDhcpServerMapOutput() DhcpServerMapOutput
	ToDhcpServerMapOutputWithContext(context.Context) DhcpServerMapOutput
}

type DhcpServerMap map[string]DhcpServerInput

func (DhcpServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpServer)(nil)).Elem()
}

func (i DhcpServerMap) ToDhcpServerMapOutput() DhcpServerMapOutput {
	return i.ToDhcpServerMapOutputWithContext(context.Background())
}

func (i DhcpServerMap) ToDhcpServerMapOutputWithContext(ctx context.Context) DhcpServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpServerMapOutput)
}

type DhcpServerOutput struct{ *pulumi.OutputState }

func (DhcpServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpServer)(nil)).Elem()
}

func (o DhcpServerOutput) ToDhcpServerOutput() DhcpServerOutput {
	return o
}

func (o DhcpServerOutput) ToDhcpServerOutputWithContext(ctx context.Context) DhcpServerOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o DhcpServerOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o DhcpServerOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// Whether to add dynamic ARP entry.
func (o DhcpServerOutput) AddArp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolPtrOutput { return v.AddArp }).(pulumi.BoolPtrOutput)
}

// IP pool, from which to take IP addresses for the clients. If set to static-only, then only the clients that have a
// static lease (added in lease submenu) will be allowed.
func (o DhcpServerOutput) AddressPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.AddressPool }).(pulumi.StringPtrOutput)
}

// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
// Requires IPv6 DHCP Server to have this option enabled as well to work properly.
func (o DhcpServerOutput) AllowDualStackQueue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolPtrOutput { return v.AllowDualStackQueue }).(pulumi.BoolPtrOutput)
}

// Always send replies as broadcasts even if destination IP is known.
func (o DhcpServerOutput) AlwaysBroadcast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolPtrOutput { return v.AlwaysBroadcast }).(pulumi.BoolPtrOutput)
}

// Option changes the way how a server responds to DHCP requests.
func (o DhcpServerOutput) Authoritative() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.Authoritative }).(pulumi.StringPtrOutput)
}

// Accepts two predefined options or time value: * forever - lease never expires * lease-time - use time from lease-time
// parameter
func (o DhcpServerOutput) BootpLeaseTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.BootpLeaseTime }).(pulumi.StringPtrOutput)
}

// Support for BOOTP clients.
func (o DhcpServerOutput) BootpSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.BootpSupport }).(pulumi.StringPtrOutput)
}

// Specifies whether to limit specific number of clients per single MAC address.
func (o DhcpServerOutput) ClientMacLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.IntPtrOutput { return v.ClientMacLimit }).(pulumi.IntPtrOutput)
}

func (o DhcpServerOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Allows to disable/enable conflict detection. If option is enabled, then whenever server tries to assign a lease it will
// send ICMP and ARP messages to detect whether such address in the network already exist. If any of above get reply
// address is considered already used. Conflict detection must be disabled when any kind of DHCP client limitation per port
// or per mac is used.
func (o DhcpServerOutput) ConflictDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolPtrOutput { return v.ConflictDetection }).(pulumi.BoolPtrOutput)
}

// If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. If set to none - there is no
// threshold (all DHCP packets are processed).
func (o DhcpServerOutput) DelayThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.DelayThreshold }).(pulumi.StringPtrOutput)
}

// Use custom set of DHCP options defined in option sets menu.
func (o DhcpServerOutput) DhcpOptionSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.DhcpOptionSet }).(pulumi.StringPtrOutput)
}

func (o DhcpServerOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
// modified.
func (o DhcpServerOutput) Dynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolOutput { return v.Dynamic }).(pulumi.BoolOutput)
}

// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
func (o DhcpServerOutput) InsertQueueBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.InsertQueueBefore }).(pulumi.StringPtrOutput)
}

// Name of the interface.
func (o DhcpServerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o DhcpServerOutput) Invalid() pulumi.BoolOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolOutput { return v.Invalid }).(pulumi.BoolOutput)
}

// A script that will be executed after a lease is assigned or de-assigned.
func (o DhcpServerOutput) LeaseScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.LeaseScript }).(pulumi.StringPtrOutput)
}

// The time that a client may use the assigned address. The client will try to renew this address after half of this time
// and will request a new address after the time limit expires.
func (o DhcpServerOutput) LeaseTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.LeaseTime }).(pulumi.StringPtrOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o DhcpServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DhcpServerOutput) ParentQueue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.ParentQueue }).(pulumi.StringPtrOutput)
}

// The IP address of the relay this DHCP server.
func (o DhcpServerOutput) Relay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.Relay }).(pulumi.StringPtrOutput)
}

// The address which the DHCP client must send requests to in order to renew an IP address lease.
func (o DhcpServerOutput) SrcAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.SrcAddress }).(pulumi.StringPtrOutput)
}

// Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client.
func (o DhcpServerOutput) UseFramedAsClassless() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.BoolPtrOutput { return v.UseFramedAsClassless }).(pulumi.BoolPtrOutput)
}

// Whether to use RADIUS server.
func (o DhcpServerOutput) UseRadius() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServer) pulumi.StringPtrOutput { return v.UseRadius }).(pulumi.StringPtrOutput)
}

type DhcpServerArrayOutput struct{ *pulumi.OutputState }

func (DhcpServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpServer)(nil)).Elem()
}

func (o DhcpServerArrayOutput) ToDhcpServerArrayOutput() DhcpServerArrayOutput {
	return o
}

func (o DhcpServerArrayOutput) ToDhcpServerArrayOutputWithContext(ctx context.Context) DhcpServerArrayOutput {
	return o
}

func (o DhcpServerArrayOutput) Index(i pulumi.IntInput) DhcpServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpServer {
		return vs[0].([]*DhcpServer)[vs[1].(int)]
	}).(DhcpServerOutput)
}

type DhcpServerMapOutput struct{ *pulumi.OutputState }

func (DhcpServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpServer)(nil)).Elem()
}

func (o DhcpServerMapOutput) ToDhcpServerMapOutput() DhcpServerMapOutput {
	return o
}

func (o DhcpServerMapOutput) ToDhcpServerMapOutputWithContext(ctx context.Context) DhcpServerMapOutput {
	return o
}

func (o DhcpServerMapOutput) MapIndex(k pulumi.StringInput) DhcpServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpServer {
		return vs[0].(map[string]*DhcpServer)[vs[1].(string)]
	}).(DhcpServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpServerInput)(nil)).Elem(), &DhcpServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpServerArrayInput)(nil)).Elem(), DhcpServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpServerMapInput)(nil)).Elem(), DhcpServerMap{})
	pulumi.RegisterOutputType(DhcpServerOutput{})
	pulumi.RegisterOutputType(DhcpServerArrayOutput{})
	pulumi.RegisterOutputType(DhcpServerMapOutput{})
}
