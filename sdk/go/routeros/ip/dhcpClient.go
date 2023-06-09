// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Ip.DhcpClient (Resource)
//
// ***
//
// #### This is an alias for backwards compatibility between plugin versions.
// Please see documentation for Ip.DhcpIpClient
type DhcpClient struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// Whether to install default route in routing table received from DHCP server.
	AddDefaultRoute pulumi.StringOutput `pulumi:"addDefaultRoute"`
	// IP address and netmask, which is assigned to DHCP Client from the Server.
	Address pulumi.StringOutput    `pulumi:"address"`
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Distance of default route. Applicable if add-default-route is set to yes.
	DefaultRouteDistance pulumi.IntPtrOutput `pulumi:"defaultRouteDistance"`
	// Options that are sent to the DHCP server.
	DhcpOptions pulumi.StringPtrOutput `pulumi:"dhcpOptions"`
	// The IP address of the DHCP server.
	DhcpServer pulumi.StringOutput  `pulumi:"dhcpServer"`
	Disabled   pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic pulumi.BoolOutput `pulumi:"dynamic"`
	// A time when the lease expires (specified by the DHCP server).
	ExpiresAfter pulumi.StringOutput `pulumi:"expiresAfter"`
	// The IP address of the gateway which is assigned by DHCP server.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Name of the interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	Invalid   pulumi.BoolOutput   `pulumi:"invalid"`
	// The IP address of the first DNS resolver, that was assigned by the DHCP server.
	PrimaryDns pulumi.StringOutput `pulumi:"primaryDns"`
	// The IP address of the primary NTP server, assigned by the DHCP server.
	PrimaryNtp pulumi.StringOutput `pulumi:"primaryNtp"`
	// The IP address of the second DNS resolver, assigned by the DHCP server.
	SecondaryDns pulumi.StringOutput `pulumi:"secondaryDns"`
	// The IP address of the secondary NTP server, assigned by the DHCP server.
	SecondaryNtp pulumi.StringOutput `pulumi:"secondaryNtp"`
	Status       pulumi.StringOutput `pulumi:"status"`
	// Whether to accept the DNS settings advertised by DHCP Server (will override the settings put in the /ip dns submenu).
	UsePeerDns pulumi.BoolPtrOutput `pulumi:"usePeerDns"`
	// Whether to accept the NTP settings advertised by DHCP Server (will override the settings put in the /system ntp client
	// submenu).
	UsePeerNtp pulumi.BoolPtrOutput `pulumi:"usePeerNtp"`
}

// NewDhcpClient registers a new resource with the given unique name, arguments, and options.
func NewDhcpClient(ctx *pulumi.Context,
	name string, args *DhcpClientArgs, opts ...pulumi.ResourceOption) (*DhcpClient, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	var resource DhcpClient
	err := ctx.RegisterResource("routeros:Ip/dhcpClient:DhcpClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpClient gets an existing DhcpClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpClientState, opts ...pulumi.ResourceOption) (*DhcpClient, error) {
	var resource DhcpClient
	err := ctx.ReadResource("routeros:Ip/dhcpClient:DhcpClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpClient resources.
type dhcpClientState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Whether to install default route in routing table received from DHCP server.
	AddDefaultRoute *string `pulumi:"addDefaultRoute"`
	// IP address and netmask, which is assigned to DHCP Client from the Server.
	Address *string `pulumi:"address"`
	Comment *string `pulumi:"comment"`
	// Distance of default route. Applicable if add-default-route is set to yes.
	DefaultRouteDistance *int `pulumi:"defaultRouteDistance"`
	// Options that are sent to the DHCP server.
	DhcpOptions *string `pulumi:"dhcpOptions"`
	// The IP address of the DHCP server.
	DhcpServer *string `pulumi:"dhcpServer"`
	Disabled   *bool   `pulumi:"disabled"`
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic *bool `pulumi:"dynamic"`
	// A time when the lease expires (specified by the DHCP server).
	ExpiresAfter *string `pulumi:"expiresAfter"`
	// The IP address of the gateway which is assigned by DHCP server.
	Gateway *string `pulumi:"gateway"`
	// Name of the interface.
	Interface *string `pulumi:"interface"`
	Invalid   *bool   `pulumi:"invalid"`
	// The IP address of the first DNS resolver, that was assigned by the DHCP server.
	PrimaryDns *string `pulumi:"primaryDns"`
	// The IP address of the primary NTP server, assigned by the DHCP server.
	PrimaryNtp *string `pulumi:"primaryNtp"`
	// The IP address of the second DNS resolver, assigned by the DHCP server.
	SecondaryDns *string `pulumi:"secondaryDns"`
	// The IP address of the secondary NTP server, assigned by the DHCP server.
	SecondaryNtp *string `pulumi:"secondaryNtp"`
	Status       *string `pulumi:"status"`
	// Whether to accept the DNS settings advertised by DHCP Server (will override the settings put in the /ip dns submenu).
	UsePeerDns *bool `pulumi:"usePeerDns"`
	// Whether to accept the NTP settings advertised by DHCP Server (will override the settings put in the /system ntp client
	// submenu).
	UsePeerNtp *bool `pulumi:"usePeerNtp"`
}

type DhcpClientState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Whether to install default route in routing table received from DHCP server.
	AddDefaultRoute pulumi.StringPtrInput
	// IP address and netmask, which is assigned to DHCP Client from the Server.
	Address pulumi.StringPtrInput
	Comment pulumi.StringPtrInput
	// Distance of default route. Applicable if add-default-route is set to yes.
	DefaultRouteDistance pulumi.IntPtrInput
	// Options that are sent to the DHCP server.
	DhcpOptions pulumi.StringPtrInput
	// The IP address of the DHCP server.
	DhcpServer pulumi.StringPtrInput
	Disabled   pulumi.BoolPtrInput
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic pulumi.BoolPtrInput
	// A time when the lease expires (specified by the DHCP server).
	ExpiresAfter pulumi.StringPtrInput
	// The IP address of the gateway which is assigned by DHCP server.
	Gateway pulumi.StringPtrInput
	// Name of the interface.
	Interface pulumi.StringPtrInput
	Invalid   pulumi.BoolPtrInput
	// The IP address of the first DNS resolver, that was assigned by the DHCP server.
	PrimaryDns pulumi.StringPtrInput
	// The IP address of the primary NTP server, assigned by the DHCP server.
	PrimaryNtp pulumi.StringPtrInput
	// The IP address of the second DNS resolver, assigned by the DHCP server.
	SecondaryDns pulumi.StringPtrInput
	// The IP address of the secondary NTP server, assigned by the DHCP server.
	SecondaryNtp pulumi.StringPtrInput
	Status       pulumi.StringPtrInput
	// Whether to accept the DNS settings advertised by DHCP Server (will override the settings put in the /ip dns submenu).
	UsePeerDns pulumi.BoolPtrInput
	// Whether to accept the NTP settings advertised by DHCP Server (will override the settings put in the /system ntp client
	// submenu).
	UsePeerNtp pulumi.BoolPtrInput
}

func (DhcpClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpClientState)(nil)).Elem()
}

type dhcpClientArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Whether to install default route in routing table received from DHCP server.
	AddDefaultRoute *string `pulumi:"addDefaultRoute"`
	Comment         *string `pulumi:"comment"`
	// Distance of default route. Applicable if add-default-route is set to yes.
	DefaultRouteDistance *int `pulumi:"defaultRouteDistance"`
	// Options that are sent to the DHCP server.
	DhcpOptions *string `pulumi:"dhcpOptions"`
	Disabled    *bool   `pulumi:"disabled"`
	// Name of the interface.
	Interface string `pulumi:"interface"`
	// Whether to accept the DNS settings advertised by DHCP Server (will override the settings put in the /ip dns submenu).
	UsePeerDns *bool `pulumi:"usePeerDns"`
	// Whether to accept the NTP settings advertised by DHCP Server (will override the settings put in the /system ntp client
	// submenu).
	UsePeerNtp *bool `pulumi:"usePeerNtp"`
}

// The set of arguments for constructing a DhcpClient resource.
type DhcpClientArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Whether to install default route in routing table received from DHCP server.
	AddDefaultRoute pulumi.StringPtrInput
	Comment         pulumi.StringPtrInput
	// Distance of default route. Applicable if add-default-route is set to yes.
	DefaultRouteDistance pulumi.IntPtrInput
	// Options that are sent to the DHCP server.
	DhcpOptions pulumi.StringPtrInput
	Disabled    pulumi.BoolPtrInput
	// Name of the interface.
	Interface pulumi.StringInput
	// Whether to accept the DNS settings advertised by DHCP Server (will override the settings put in the /ip dns submenu).
	UsePeerDns pulumi.BoolPtrInput
	// Whether to accept the NTP settings advertised by DHCP Server (will override the settings put in the /system ntp client
	// submenu).
	UsePeerNtp pulumi.BoolPtrInput
}

func (DhcpClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpClientArgs)(nil)).Elem()
}

type DhcpClientInput interface {
	pulumi.Input

	ToDhcpClientOutput() DhcpClientOutput
	ToDhcpClientOutputWithContext(ctx context.Context) DhcpClientOutput
}

func (*DhcpClient) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpClient)(nil)).Elem()
}

func (i *DhcpClient) ToDhcpClientOutput() DhcpClientOutput {
	return i.ToDhcpClientOutputWithContext(context.Background())
}

func (i *DhcpClient) ToDhcpClientOutputWithContext(ctx context.Context) DhcpClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpClientOutput)
}

// DhcpClientArrayInput is an input type that accepts DhcpClientArray and DhcpClientArrayOutput values.
// You can construct a concrete instance of `DhcpClientArrayInput` via:
//
//	DhcpClientArray{ DhcpClientArgs{...} }
type DhcpClientArrayInput interface {
	pulumi.Input

	ToDhcpClientArrayOutput() DhcpClientArrayOutput
	ToDhcpClientArrayOutputWithContext(context.Context) DhcpClientArrayOutput
}

type DhcpClientArray []DhcpClientInput

func (DhcpClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpClient)(nil)).Elem()
}

func (i DhcpClientArray) ToDhcpClientArrayOutput() DhcpClientArrayOutput {
	return i.ToDhcpClientArrayOutputWithContext(context.Background())
}

func (i DhcpClientArray) ToDhcpClientArrayOutputWithContext(ctx context.Context) DhcpClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpClientArrayOutput)
}

// DhcpClientMapInput is an input type that accepts DhcpClientMap and DhcpClientMapOutput values.
// You can construct a concrete instance of `DhcpClientMapInput` via:
//
//	DhcpClientMap{ "key": DhcpClientArgs{...} }
type DhcpClientMapInput interface {
	pulumi.Input

	ToDhcpClientMapOutput() DhcpClientMapOutput
	ToDhcpClientMapOutputWithContext(context.Context) DhcpClientMapOutput
}

type DhcpClientMap map[string]DhcpClientInput

func (DhcpClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpClient)(nil)).Elem()
}

func (i DhcpClientMap) ToDhcpClientMapOutput() DhcpClientMapOutput {
	return i.ToDhcpClientMapOutputWithContext(context.Background())
}

func (i DhcpClientMap) ToDhcpClientMapOutputWithContext(ctx context.Context) DhcpClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpClientMapOutput)
}

type DhcpClientOutput struct{ *pulumi.OutputState }

func (DhcpClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpClient)(nil)).Elem()
}

func (o DhcpClientOutput) ToDhcpClientOutput() DhcpClientOutput {
	return o
}

func (o DhcpClientOutput) ToDhcpClientOutputWithContext(ctx context.Context) DhcpClientOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o DhcpClientOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o DhcpClientOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// Whether to install default route in routing table received from DHCP server.
func (o DhcpClientOutput) AddDefaultRoute() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.AddDefaultRoute }).(pulumi.StringOutput)
}

// IP address and netmask, which is assigned to DHCP Client from the Server.
func (o DhcpClientOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

func (o DhcpClientOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Distance of default route. Applicable if add-default-route is set to yes.
func (o DhcpClientOutput) DefaultRouteDistance() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.IntPtrOutput { return v.DefaultRouteDistance }).(pulumi.IntPtrOutput)
}

// Options that are sent to the DHCP server.
func (o DhcpClientOutput) DhcpOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringPtrOutput { return v.DhcpOptions }).(pulumi.StringPtrOutput)
}

// The IP address of the DHCP server.
func (o DhcpClientOutput) DhcpServer() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.DhcpServer }).(pulumi.StringOutput)
}

func (o DhcpClientOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
// modified.
func (o DhcpClientOutput) Dynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.BoolOutput { return v.Dynamic }).(pulumi.BoolOutput)
}

// A time when the lease expires (specified by the DHCP server).
func (o DhcpClientOutput) ExpiresAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.ExpiresAfter }).(pulumi.StringOutput)
}

// The IP address of the gateway which is assigned by DHCP server.
func (o DhcpClientOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Name of the interface.
func (o DhcpClientOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o DhcpClientOutput) Invalid() pulumi.BoolOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.BoolOutput { return v.Invalid }).(pulumi.BoolOutput)
}

// The IP address of the first DNS resolver, that was assigned by the DHCP server.
func (o DhcpClientOutput) PrimaryDns() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.PrimaryDns }).(pulumi.StringOutput)
}

// The IP address of the primary NTP server, assigned by the DHCP server.
func (o DhcpClientOutput) PrimaryNtp() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.PrimaryNtp }).(pulumi.StringOutput)
}

// The IP address of the second DNS resolver, assigned by the DHCP server.
func (o DhcpClientOutput) SecondaryDns() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.SecondaryDns }).(pulumi.StringOutput)
}

// The IP address of the secondary NTP server, assigned by the DHCP server.
func (o DhcpClientOutput) SecondaryNtp() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.SecondaryNtp }).(pulumi.StringOutput)
}

func (o DhcpClientOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Whether to accept the DNS settings advertised by DHCP Server (will override the settings put in the /ip dns submenu).
func (o DhcpClientOutput) UsePeerDns() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.BoolPtrOutput { return v.UsePeerDns }).(pulumi.BoolPtrOutput)
}

// Whether to accept the NTP settings advertised by DHCP Server (will override the settings put in the /system ntp client
// submenu).
func (o DhcpClientOutput) UsePeerNtp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DhcpClient) pulumi.BoolPtrOutput { return v.UsePeerNtp }).(pulumi.BoolPtrOutput)
}

type DhcpClientArrayOutput struct{ *pulumi.OutputState }

func (DhcpClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpClient)(nil)).Elem()
}

func (o DhcpClientArrayOutput) ToDhcpClientArrayOutput() DhcpClientArrayOutput {
	return o
}

func (o DhcpClientArrayOutput) ToDhcpClientArrayOutputWithContext(ctx context.Context) DhcpClientArrayOutput {
	return o
}

func (o DhcpClientArrayOutput) Index(i pulumi.IntInput) DhcpClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpClient {
		return vs[0].([]*DhcpClient)[vs[1].(int)]
	}).(DhcpClientOutput)
}

type DhcpClientMapOutput struct{ *pulumi.OutputState }

func (DhcpClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpClient)(nil)).Elem()
}

func (o DhcpClientMapOutput) ToDhcpClientMapOutput() DhcpClientMapOutput {
	return o
}

func (o DhcpClientMapOutput) ToDhcpClientMapOutputWithContext(ctx context.Context) DhcpClientMapOutput {
	return o
}

func (o DhcpClientMapOutput) MapIndex(k pulumi.StringInput) DhcpClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpClient {
		return vs[0].(map[string]*DhcpClient)[vs[1].(string)]
	}).(DhcpClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpClientInput)(nil)).Elem(), &DhcpClient{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpClientArrayInput)(nil)).Elem(), DhcpClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpClientMapInput)(nil)).Elem(), DhcpClientMap{})
	pulumi.RegisterOutputType(DhcpClientOutput{})
	pulumi.RegisterOutputType(DhcpClientArrayOutput{})
	pulumi.RegisterOutputType(DhcpClientMapOutput{})
}
