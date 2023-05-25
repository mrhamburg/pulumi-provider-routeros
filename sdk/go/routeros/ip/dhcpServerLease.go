// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Ip.DhcpServerLease (Resource)
//
// ***
//
// #### This is an alias for backwards compatibility between plugin versions.
// Please see documentation for Ip.DhcpIpServerLease
type DhcpServerLease struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// The IP address of the machine currently holding the DHCP lease.
	ActiveAddress pulumi.StringOutput `pulumi:"activeAddress"`
	// Actual client-id of the client.
	ActiveClientId pulumi.StringOutput `pulumi:"activeClientId"`
	// The hostname of the machine currently holding the DHCP lease.
	ActiveHostname pulumi.StringOutput `pulumi:"activeHostname"`
	// The MAC address of of the machine currently holding the DHCP lease.
	ActiveMacAddress pulumi.StringOutput `pulumi:"activeMacAddress"`
	// Actual dhcp server, which serves this client.
	ActiveServer pulumi.StringOutput `pulumi:"activeServer"`
	// The IP address of the DHCP lease to be created.
	Address pulumi.StringOutput `pulumi:"address"`
	// Address list to which address will be added if lease is bound.
	AddressLists pulumi.StringPtrOutput `pulumi:"addressLists"`
	// Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as
	// hex dump.
	AgentCircuitId pulumi.StringOutput `pulumi:"agentCircuitId"`
	// Remote ID, set by DHCP relay agent.
	AgentRemoteId pulumi.StringOutput `pulumi:"agentRemoteId"`
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	AllowDualStackQueue pulumi.BoolPtrOutput `pulumi:"allowDualStackQueue"`
	// Send all replies as broadcasts.
	AlwaysBroadcast pulumi.BoolPtrOutput `pulumi:"alwaysBroadcast"`
	// Whether to block access for this DHCP client (true|false).
	BlockAccess pulumi.BoolPtrOutput `pulumi:"blockAccess"`
	// Whether the lease is blocked.
	Blocked pulumi.BoolOutput `pulumi:"blocked"`
	// If specified, must match DHCP 'client identifier' option of the request.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	Comment  pulumi.StringPtrOutput `pulumi:"comment"`
	// Add additional DHCP options.
	DhcpOption pulumi.StringPtrOutput `pulumi:"dhcpOption"`
	// Add additional set of DHCP options.
	DhcpOptionSet pulumi.StringPtrOutput `pulumi:"dhcpOptionSet"`
	Disabled      pulumi.BoolPtrOutput   `pulumi:"disabled"`
	// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
	// specific device. Defaults to false.
	Dynamic pulumi.BoolPtrOutput `pulumi:"dynamic"`
	// Time until lease expires.
	ExpiresAfter pulumi.StringOutput `pulumi:"expiresAfter"`
	// The hostname of the device
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore pulumi.StringPtrOutput `pulumi:"insertQueueBefore"`
	LastSeen          pulumi.StringOutput    `pulumi:"lastSeen"`
	// Time that the client may use the address. If set to 0s lease will never expire.
	LeaseTime pulumi.StringPtrOutput `pulumi:"leaseTime"`
	// The MAC addreess of the DHCP lease to be created.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// Shows if this dynamic lease is authenticated by RADIUS or not.
	Radius pulumi.StringOutput `pulumi:"radius"`
	// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
	RateLimit pulumi.StringPtrOutput `pulumi:"rateLimit"`
	// Server name which serves this client.
	Server pulumi.StringPtrOutput `pulumi:"server"`
	// Source MAC address.
	SrcMacAddress pulumi.StringOutput `pulumi:"srcMacAddress"`
	// Lease status.
	Status pulumi.StringOutput `pulumi:"status"`
	// When this option is set server uses source MAC address instead of received CHADDR to assign address.
	UseSrcMac pulumi.BoolPtrOutput `pulumi:"useSrcMac"`
}

// NewDhcpServerLease registers a new resource with the given unique name, arguments, and options.
func NewDhcpServerLease(ctx *pulumi.Context,
	name string, args *DhcpServerLeaseArgs, opts ...pulumi.ResourceOption) (*DhcpServerLease, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.MacAddress == nil {
		return nil, errors.New("invalid value for required argument 'MacAddress'")
	}
	var resource DhcpServerLease
	err := ctx.RegisterResource("routeros:Ip/dhcpServerLease:DhcpServerLease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpServerLease gets an existing DhcpServerLease resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpServerLease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpServerLeaseState, opts ...pulumi.ResourceOption) (*DhcpServerLease, error) {
	var resource DhcpServerLease
	err := ctx.ReadResource("routeros:Ip/dhcpServerLease:DhcpServerLease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpServerLease resources.
type dhcpServerLeaseState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// The IP address of the machine currently holding the DHCP lease.
	ActiveAddress *string `pulumi:"activeAddress"`
	// Actual client-id of the client.
	ActiveClientId *string `pulumi:"activeClientId"`
	// The hostname of the machine currently holding the DHCP lease.
	ActiveHostname *string `pulumi:"activeHostname"`
	// The MAC address of of the machine currently holding the DHCP lease.
	ActiveMacAddress *string `pulumi:"activeMacAddress"`
	// Actual dhcp server, which serves this client.
	ActiveServer *string `pulumi:"activeServer"`
	// The IP address of the DHCP lease to be created.
	Address *string `pulumi:"address"`
	// Address list to which address will be added if lease is bound.
	AddressLists *string `pulumi:"addressLists"`
	// Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as
	// hex dump.
	AgentCircuitId *string `pulumi:"agentCircuitId"`
	// Remote ID, set by DHCP relay agent.
	AgentRemoteId *string `pulumi:"agentRemoteId"`
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	AllowDualStackQueue *bool `pulumi:"allowDualStackQueue"`
	// Send all replies as broadcasts.
	AlwaysBroadcast *bool `pulumi:"alwaysBroadcast"`
	// Whether to block access for this DHCP client (true|false).
	BlockAccess *bool `pulumi:"blockAccess"`
	// Whether the lease is blocked.
	Blocked *bool `pulumi:"blocked"`
	// If specified, must match DHCP 'client identifier' option of the request.
	ClientId *string `pulumi:"clientId"`
	Comment  *string `pulumi:"comment"`
	// Add additional DHCP options.
	DhcpOption *string `pulumi:"dhcpOption"`
	// Add additional set of DHCP options.
	DhcpOptionSet *string `pulumi:"dhcpOptionSet"`
	Disabled      *bool   `pulumi:"disabled"`
	// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
	// specific device. Defaults to false.
	Dynamic *bool `pulumi:"dynamic"`
	// Time until lease expires.
	ExpiresAfter *string `pulumi:"expiresAfter"`
	// The hostname of the device
	HostName *string `pulumi:"hostName"`
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore *string `pulumi:"insertQueueBefore"`
	LastSeen          *string `pulumi:"lastSeen"`
	// Time that the client may use the address. If set to 0s lease will never expire.
	LeaseTime *string `pulumi:"leaseTime"`
	// The MAC addreess of the DHCP lease to be created.
	MacAddress *string `pulumi:"macAddress"`
	// Shows if this dynamic lease is authenticated by RADIUS or not.
	Radius *string `pulumi:"radius"`
	// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
	RateLimit *string `pulumi:"rateLimit"`
	// Server name which serves this client.
	Server *string `pulumi:"server"`
	// Source MAC address.
	SrcMacAddress *string `pulumi:"srcMacAddress"`
	// Lease status.
	Status *string `pulumi:"status"`
	// When this option is set server uses source MAC address instead of received CHADDR to assign address.
	UseSrcMac *bool `pulumi:"useSrcMac"`
}

type DhcpServerLeaseState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// The IP address of the machine currently holding the DHCP lease.
	ActiveAddress pulumi.StringPtrInput
	// Actual client-id of the client.
	ActiveClientId pulumi.StringPtrInput
	// The hostname of the machine currently holding the DHCP lease.
	ActiveHostname pulumi.StringPtrInput
	// The MAC address of of the machine currently holding the DHCP lease.
	ActiveMacAddress pulumi.StringPtrInput
	// Actual dhcp server, which serves this client.
	ActiveServer pulumi.StringPtrInput
	// The IP address of the DHCP lease to be created.
	Address pulumi.StringPtrInput
	// Address list to which address will be added if lease is bound.
	AddressLists pulumi.StringPtrInput
	// Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as
	// hex dump.
	AgentCircuitId pulumi.StringPtrInput
	// Remote ID, set by DHCP relay agent.
	AgentRemoteId pulumi.StringPtrInput
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	AllowDualStackQueue pulumi.BoolPtrInput
	// Send all replies as broadcasts.
	AlwaysBroadcast pulumi.BoolPtrInput
	// Whether to block access for this DHCP client (true|false).
	BlockAccess pulumi.BoolPtrInput
	// Whether the lease is blocked.
	Blocked pulumi.BoolPtrInput
	// If specified, must match DHCP 'client identifier' option of the request.
	ClientId pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	// Add additional DHCP options.
	DhcpOption pulumi.StringPtrInput
	// Add additional set of DHCP options.
	DhcpOptionSet pulumi.StringPtrInput
	Disabled      pulumi.BoolPtrInput
	// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
	// specific device. Defaults to false.
	Dynamic pulumi.BoolPtrInput
	// Time until lease expires.
	ExpiresAfter pulumi.StringPtrInput
	// The hostname of the device
	HostName pulumi.StringPtrInput
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore pulumi.StringPtrInput
	LastSeen          pulumi.StringPtrInput
	// Time that the client may use the address. If set to 0s lease will never expire.
	LeaseTime pulumi.StringPtrInput
	// The MAC addreess of the DHCP lease to be created.
	MacAddress pulumi.StringPtrInput
	// Shows if this dynamic lease is authenticated by RADIUS or not.
	Radius pulumi.StringPtrInput
	// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
	RateLimit pulumi.StringPtrInput
	// Server name which serves this client.
	Server pulumi.StringPtrInput
	// Source MAC address.
	SrcMacAddress pulumi.StringPtrInput
	// Lease status.
	Status pulumi.StringPtrInput
	// When this option is set server uses source MAC address instead of received CHADDR to assign address.
	UseSrcMac pulumi.BoolPtrInput
}

func (DhcpServerLeaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpServerLeaseState)(nil)).Elem()
}

type dhcpServerLeaseArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// The IP address of the DHCP lease to be created.
	Address string `pulumi:"address"`
	// Address list to which address will be added if lease is bound.
	AddressLists *string `pulumi:"addressLists"`
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	AllowDualStackQueue *bool `pulumi:"allowDualStackQueue"`
	// Send all replies as broadcasts.
	AlwaysBroadcast *bool `pulumi:"alwaysBroadcast"`
	// Whether to block access for this DHCP client (true|false).
	BlockAccess *bool `pulumi:"blockAccess"`
	// If specified, must match DHCP 'client identifier' option of the request.
	ClientId *string `pulumi:"clientId"`
	Comment  *string `pulumi:"comment"`
	// Add additional DHCP options.
	DhcpOption *string `pulumi:"dhcpOption"`
	// Add additional set of DHCP options.
	DhcpOptionSet *string `pulumi:"dhcpOptionSet"`
	Disabled      *bool   `pulumi:"disabled"`
	// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
	// specific device. Defaults to false.
	Dynamic *bool `pulumi:"dynamic"`
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore *string `pulumi:"insertQueueBefore"`
	// Time that the client may use the address. If set to 0s lease will never expire.
	LeaseTime *string `pulumi:"leaseTime"`
	// The MAC addreess of the DHCP lease to be created.
	MacAddress string `pulumi:"macAddress"`
	// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
	RateLimit *string `pulumi:"rateLimit"`
	// Server name which serves this client.
	Server *string `pulumi:"server"`
	// When this option is set server uses source MAC address instead of received CHADDR to assign address.
	UseSrcMac *bool `pulumi:"useSrcMac"`
}

// The set of arguments for constructing a DhcpServerLease resource.
type DhcpServerLeaseArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// The IP address of the DHCP lease to be created.
	Address pulumi.StringInput
	// Address list to which address will be added if lease is bound.
	AddressLists pulumi.StringPtrInput
	// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
	AllowDualStackQueue pulumi.BoolPtrInput
	// Send all replies as broadcasts.
	AlwaysBroadcast pulumi.BoolPtrInput
	// Whether to block access for this DHCP client (true|false).
	BlockAccess pulumi.BoolPtrInput
	// If specified, must match DHCP 'client identifier' option of the request.
	ClientId pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	// Add additional DHCP options.
	DhcpOption pulumi.StringPtrInput
	// Add additional set of DHCP options.
	DhcpOptionSet pulumi.StringPtrInput
	Disabled      pulumi.BoolPtrInput
	// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
	// specific device. Defaults to false.
	Dynamic pulumi.BoolPtrInput
	// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
	InsertQueueBefore pulumi.StringPtrInput
	// Time that the client may use the address. If set to 0s lease will never expire.
	LeaseTime pulumi.StringPtrInput
	// The MAC addreess of the DHCP lease to be created.
	MacAddress pulumi.StringInput
	// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
	RateLimit pulumi.StringPtrInput
	// Server name which serves this client.
	Server pulumi.StringPtrInput
	// When this option is set server uses source MAC address instead of received CHADDR to assign address.
	UseSrcMac pulumi.BoolPtrInput
}

func (DhcpServerLeaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpServerLeaseArgs)(nil)).Elem()
}

type DhcpServerLeaseInput interface {
	pulumi.Input

	ToDhcpServerLeaseOutput() DhcpServerLeaseOutput
	ToDhcpServerLeaseOutputWithContext(ctx context.Context) DhcpServerLeaseOutput
}

func (*DhcpServerLease) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpServerLease)(nil)).Elem()
}

func (i *DhcpServerLease) ToDhcpServerLeaseOutput() DhcpServerLeaseOutput {
	return i.ToDhcpServerLeaseOutputWithContext(context.Background())
}

func (i *DhcpServerLease) ToDhcpServerLeaseOutputWithContext(ctx context.Context) DhcpServerLeaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpServerLeaseOutput)
}

// DhcpServerLeaseArrayInput is an input type that accepts DhcpServerLeaseArray and DhcpServerLeaseArrayOutput values.
// You can construct a concrete instance of `DhcpServerLeaseArrayInput` via:
//
//	DhcpServerLeaseArray{ DhcpServerLeaseArgs{...} }
type DhcpServerLeaseArrayInput interface {
	pulumi.Input

	ToDhcpServerLeaseArrayOutput() DhcpServerLeaseArrayOutput
	ToDhcpServerLeaseArrayOutputWithContext(context.Context) DhcpServerLeaseArrayOutput
}

type DhcpServerLeaseArray []DhcpServerLeaseInput

func (DhcpServerLeaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpServerLease)(nil)).Elem()
}

func (i DhcpServerLeaseArray) ToDhcpServerLeaseArrayOutput() DhcpServerLeaseArrayOutput {
	return i.ToDhcpServerLeaseArrayOutputWithContext(context.Background())
}

func (i DhcpServerLeaseArray) ToDhcpServerLeaseArrayOutputWithContext(ctx context.Context) DhcpServerLeaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpServerLeaseArrayOutput)
}

// DhcpServerLeaseMapInput is an input type that accepts DhcpServerLeaseMap and DhcpServerLeaseMapOutput values.
// You can construct a concrete instance of `DhcpServerLeaseMapInput` via:
//
//	DhcpServerLeaseMap{ "key": DhcpServerLeaseArgs{...} }
type DhcpServerLeaseMapInput interface {
	pulumi.Input

	ToDhcpServerLeaseMapOutput() DhcpServerLeaseMapOutput
	ToDhcpServerLeaseMapOutputWithContext(context.Context) DhcpServerLeaseMapOutput
}

type DhcpServerLeaseMap map[string]DhcpServerLeaseInput

func (DhcpServerLeaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpServerLease)(nil)).Elem()
}

func (i DhcpServerLeaseMap) ToDhcpServerLeaseMapOutput() DhcpServerLeaseMapOutput {
	return i.ToDhcpServerLeaseMapOutputWithContext(context.Background())
}

func (i DhcpServerLeaseMap) ToDhcpServerLeaseMapOutputWithContext(ctx context.Context) DhcpServerLeaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpServerLeaseMapOutput)
}

type DhcpServerLeaseOutput struct{ *pulumi.OutputState }

func (DhcpServerLeaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpServerLease)(nil)).Elem()
}

func (o DhcpServerLeaseOutput) ToDhcpServerLeaseOutput() DhcpServerLeaseOutput {
	return o
}

func (o DhcpServerLeaseOutput) ToDhcpServerLeaseOutputWithContext(ctx context.Context) DhcpServerLeaseOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o DhcpServerLeaseOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o DhcpServerLeaseOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// The IP address of the machine currently holding the DHCP lease.
func (o DhcpServerLeaseOutput) ActiveAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.ActiveAddress }).(pulumi.StringOutput)
}

// Actual client-id of the client.
func (o DhcpServerLeaseOutput) ActiveClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.ActiveClientId }).(pulumi.StringOutput)
}

// The hostname of the machine currently holding the DHCP lease.
func (o DhcpServerLeaseOutput) ActiveHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.ActiveHostname }).(pulumi.StringOutput)
}

// The MAC address of of the machine currently holding the DHCP lease.
func (o DhcpServerLeaseOutput) ActiveMacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.ActiveMacAddress }).(pulumi.StringOutput)
}

// Actual dhcp server, which serves this client.
func (o DhcpServerLeaseOutput) ActiveServer() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.ActiveServer }).(pulumi.StringOutput)
}

// The IP address of the DHCP lease to be created.
func (o DhcpServerLeaseOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Address list to which address will be added if lease is bound.
func (o DhcpServerLeaseOutput) AddressLists() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.AddressLists }).(pulumi.StringPtrOutput)
}

// Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else this value is displayed as
// hex dump.
func (o DhcpServerLeaseOutput) AgentCircuitId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.AgentCircuitId }).(pulumi.StringOutput)
}

// Remote ID, set by DHCP relay agent.
func (o DhcpServerLeaseOutput) AgentRemoteId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.AgentRemoteId }).(pulumi.StringOutput)
}

// Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification.
func (o DhcpServerLeaseOutput) AllowDualStackQueue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.BoolPtrOutput { return v.AllowDualStackQueue }).(pulumi.BoolPtrOutput)
}

// Send all replies as broadcasts.
func (o DhcpServerLeaseOutput) AlwaysBroadcast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.BoolPtrOutput { return v.AlwaysBroadcast }).(pulumi.BoolPtrOutput)
}

// Whether to block access for this DHCP client (true|false).
func (o DhcpServerLeaseOutput) BlockAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.BoolPtrOutput { return v.BlockAccess }).(pulumi.BoolPtrOutput)
}

// Whether the lease is blocked.
func (o DhcpServerLeaseOutput) Blocked() pulumi.BoolOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.BoolOutput { return v.Blocked }).(pulumi.BoolOutput)
}

// If specified, must match DHCP 'client identifier' option of the request.
func (o DhcpServerLeaseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o DhcpServerLeaseOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Add additional DHCP options.
func (o DhcpServerLeaseOutput) DhcpOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.DhcpOption }).(pulumi.StringPtrOutput)
}

// Add additional set of DHCP options.
func (o DhcpServerLeaseOutput) DhcpOptionSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.DhcpOptionSet }).(pulumi.StringPtrOutput)
}

func (o DhcpServerLeaseOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that
// specific device. Defaults to false.
func (o DhcpServerLeaseOutput) Dynamic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.BoolPtrOutput { return v.Dynamic }).(pulumi.BoolPtrOutput)
}

// Time until lease expires.
func (o DhcpServerLeaseOutput) ExpiresAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.ExpiresAfter }).(pulumi.StringOutput)
}

// The hostname of the device
func (o DhcpServerLeaseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.
func (o DhcpServerLeaseOutput) InsertQueueBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.InsertQueueBefore }).(pulumi.StringPtrOutput)
}

func (o DhcpServerLeaseOutput) LastSeen() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.LastSeen }).(pulumi.StringOutput)
}

// Time that the client may use the address. If set to 0s lease will never expire.
func (o DhcpServerLeaseOutput) LeaseTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.LeaseTime }).(pulumi.StringPtrOutput)
}

// The MAC addreess of the DHCP lease to be created.
func (o DhcpServerLeaseOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// Shows if this dynamic lease is authenticated by RADIUS or not.
func (o DhcpServerLeaseOutput) Radius() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.Radius }).(pulumi.StringOutput)
}

// Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static.
func (o DhcpServerLeaseOutput) RateLimit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.RateLimit }).(pulumi.StringPtrOutput)
}

// Server name which serves this client.
func (o DhcpServerLeaseOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringPtrOutput { return v.Server }).(pulumi.StringPtrOutput)
}

// Source MAC address.
func (o DhcpServerLeaseOutput) SrcMacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.SrcMacAddress }).(pulumi.StringOutput)
}

// Lease status.
func (o DhcpServerLeaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// When this option is set server uses source MAC address instead of received CHADDR to assign address.
func (o DhcpServerLeaseOutput) UseSrcMac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpServerLease) pulumi.BoolPtrOutput { return v.UseSrcMac }).(pulumi.BoolPtrOutput)
}

type DhcpServerLeaseArrayOutput struct{ *pulumi.OutputState }

func (DhcpServerLeaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpServerLease)(nil)).Elem()
}

func (o DhcpServerLeaseArrayOutput) ToDhcpServerLeaseArrayOutput() DhcpServerLeaseArrayOutput {
	return o
}

func (o DhcpServerLeaseArrayOutput) ToDhcpServerLeaseArrayOutputWithContext(ctx context.Context) DhcpServerLeaseArrayOutput {
	return o
}

func (o DhcpServerLeaseArrayOutput) Index(i pulumi.IntInput) DhcpServerLeaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpServerLease {
		return vs[0].([]*DhcpServerLease)[vs[1].(int)]
	}).(DhcpServerLeaseOutput)
}

type DhcpServerLeaseMapOutput struct{ *pulumi.OutputState }

func (DhcpServerLeaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpServerLease)(nil)).Elem()
}

func (o DhcpServerLeaseMapOutput) ToDhcpServerLeaseMapOutput() DhcpServerLeaseMapOutput {
	return o
}

func (o DhcpServerLeaseMapOutput) ToDhcpServerLeaseMapOutputWithContext(ctx context.Context) DhcpServerLeaseMapOutput {
	return o
}

func (o DhcpServerLeaseMapOutput) MapIndex(k pulumi.StringInput) DhcpServerLeaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpServerLease {
		return vs[0].(map[string]*DhcpServerLease)[vs[1].(string)]
	}).(DhcpServerLeaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpServerLeaseInput)(nil)).Elem(), &DhcpServerLease{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpServerLeaseArrayInput)(nil)).Elem(), DhcpServerLeaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpServerLeaseMapInput)(nil)).Elem(), DhcpServerLeaseMap{})
	pulumi.RegisterOutputType(DhcpServerLeaseOutput{})
	pulumi.RegisterOutputType(DhcpServerLeaseArrayOutput{})
	pulumi.RegisterOutputType(DhcpServerLeaseMapOutput{})
}