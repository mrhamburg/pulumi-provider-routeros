// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Ip.IpDns (Resource)
//
// A MikroTik router with DNS feature enabled can be set as a DNS server for any DNS-compliant client.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-routeros/sdk/go/routeros/Ip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ip.NewDns(ctx, "dns-server", &Ip.DnsArgs{
//				AllowRemoteRequests: pulumi.Bool(true),
//				Servers:             pulumi.String("2606:4700:4700::1111,1.1.1.1,2606:4700:4700::1001,1.0.0.1"),
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
// #The DNS Settings can not be imported.
//
// #Terraform will ignore the current settings and will overwrite the current settings with the settings defined in Terraform.
type IpDns struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// Specifies whether to allow network requests.
	AllowRemoteRequests pulumi.BoolPtrOutput `pulumi:"allowRemoteRequests"`
	// Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
	// time. Shorter TTL received from DNS servers are respected. *Default: 1w*
	CacheMaxTtl pulumi.StringOutput `pulumi:"cacheMaxTtl"`
	// Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
	CacheSize pulumi.IntOutput `pulumi:"cacheSize"`
	// Shows the currently used cache size in KiB.
	CacheUsed pulumi.IntOutput `pulumi:"cacheUsed"`
	// Specifies how many DoH concurrent queries are allowed.
	DohMaxConcurrentQueries pulumi.IntOutput `pulumi:"dohMaxConcurrentQueries"`
	// Specifies how many concurrent connections to the DoH server are allowed.
	DohMaxServerConnections pulumi.IntOutput `pulumi:"dohMaxServerConnections"`
	// Specifies how long to wait for query response from the DoH server.
	DohTimeout pulumi.StringOutput `pulumi:"dohTimeout"`
	// List of dynamically added DNS server from different services, for example, DHCP.
	DynamicServers pulumi.StringOutput `pulumi:"dynamicServers"`
	// Specifies how much concurrent queries are allowed. *Default: 100*
	MaxConcurrentQueries pulumi.IntOutput `pulumi:"maxConcurrentQueries"`
	// Specifies how much concurrent TCP sessions are allowed. *Default: 20*
	MaxConcurrentTcpSessions pulumi.IntOutput `pulumi:"maxConcurrentTcpSessions"`
	// Maximum size of allowed UDP packet. *Default: 4096*
	MaxUdpPacketSize pulumi.IntOutput `pulumi:"maxUdpPacketSize"`
	// Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
	QueryServerTimeout pulumi.StringOutput `pulumi:"queryServerTimeout"`
	// Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
	// query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
	QueryTotalTimeout pulumi.StringOutput `pulumi:"queryTotalTimeout"`
	// List of DNS server IPv4/IPv6 addresses.
	Servers pulumi.StringPtrOutput `pulumi:"servers"`
	// DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
	// fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
	// on the device.
	UseDohServer pulumi.StringPtrOutput `pulumi:"useDohServer"`
	// DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
	VerifyDohCert pulumi.BoolPtrOutput `pulumi:"verifyDohCert"`
}

// NewIpDns registers a new resource with the given unique name, arguments, and options.
func NewIpDns(ctx *pulumi.Context,
	name string, args *IpDnsArgs, opts ...pulumi.ResourceOption) (*IpDns, error) {
	if args == nil {
		args = &IpDnsArgs{}
	}

	var resource IpDns
	err := ctx.RegisterResource("routeros:Ip/ipDns:IpDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpDns gets an existing IpDns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpDnsState, opts ...pulumi.ResourceOption) (*IpDns, error) {
	var resource IpDns
	err := ctx.ReadResource("routeros:Ip/ipDns:IpDns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpDns resources.
type ipDnsState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Specifies whether to allow network requests.
	AllowRemoteRequests *bool `pulumi:"allowRemoteRequests"`
	// Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
	// time. Shorter TTL received from DNS servers are respected. *Default: 1w*
	CacheMaxTtl *string `pulumi:"cacheMaxTtl"`
	// Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
	CacheSize *int `pulumi:"cacheSize"`
	// Shows the currently used cache size in KiB.
	CacheUsed *int `pulumi:"cacheUsed"`
	// Specifies how many DoH concurrent queries are allowed.
	DohMaxConcurrentQueries *int `pulumi:"dohMaxConcurrentQueries"`
	// Specifies how many concurrent connections to the DoH server are allowed.
	DohMaxServerConnections *int `pulumi:"dohMaxServerConnections"`
	// Specifies how long to wait for query response from the DoH server.
	DohTimeout *string `pulumi:"dohTimeout"`
	// List of dynamically added DNS server from different services, for example, DHCP.
	DynamicServers *string `pulumi:"dynamicServers"`
	// Specifies how much concurrent queries are allowed. *Default: 100*
	MaxConcurrentQueries *int `pulumi:"maxConcurrentQueries"`
	// Specifies how much concurrent TCP sessions are allowed. *Default: 20*
	MaxConcurrentTcpSessions *int `pulumi:"maxConcurrentTcpSessions"`
	// Maximum size of allowed UDP packet. *Default: 4096*
	MaxUdpPacketSize *int `pulumi:"maxUdpPacketSize"`
	// Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
	QueryServerTimeout *string `pulumi:"queryServerTimeout"`
	// Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
	// query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
	QueryTotalTimeout *string `pulumi:"queryTotalTimeout"`
	// List of DNS server IPv4/IPv6 addresses.
	Servers *string `pulumi:"servers"`
	// DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
	// fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
	// on the device.
	UseDohServer *string `pulumi:"useDohServer"`
	// DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
	VerifyDohCert *bool `pulumi:"verifyDohCert"`
}

type IpDnsState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Specifies whether to allow network requests.
	AllowRemoteRequests pulumi.BoolPtrInput
	// Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
	// time. Shorter TTL received from DNS servers are respected. *Default: 1w*
	CacheMaxTtl pulumi.StringPtrInput
	// Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
	CacheSize pulumi.IntPtrInput
	// Shows the currently used cache size in KiB.
	CacheUsed pulumi.IntPtrInput
	// Specifies how many DoH concurrent queries are allowed.
	DohMaxConcurrentQueries pulumi.IntPtrInput
	// Specifies how many concurrent connections to the DoH server are allowed.
	DohMaxServerConnections pulumi.IntPtrInput
	// Specifies how long to wait for query response from the DoH server.
	DohTimeout pulumi.StringPtrInput
	// List of dynamically added DNS server from different services, for example, DHCP.
	DynamicServers pulumi.StringPtrInput
	// Specifies how much concurrent queries are allowed. *Default: 100*
	MaxConcurrentQueries pulumi.IntPtrInput
	// Specifies how much concurrent TCP sessions are allowed. *Default: 20*
	MaxConcurrentTcpSessions pulumi.IntPtrInput
	// Maximum size of allowed UDP packet. *Default: 4096*
	MaxUdpPacketSize pulumi.IntPtrInput
	// Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
	QueryServerTimeout pulumi.StringPtrInput
	// Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
	// query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
	QueryTotalTimeout pulumi.StringPtrInput
	// List of DNS server IPv4/IPv6 addresses.
	Servers pulumi.StringPtrInput
	// DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
	// fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
	// on the device.
	UseDohServer pulumi.StringPtrInput
	// DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
	VerifyDohCert pulumi.BoolPtrInput
}

func (IpDnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipDnsState)(nil)).Elem()
}

type ipDnsArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// Specifies whether to allow network requests.
	AllowRemoteRequests *bool `pulumi:"allowRemoteRequests"`
	// Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
	// time. Shorter TTL received from DNS servers are respected. *Default: 1w*
	CacheMaxTtl *string `pulumi:"cacheMaxTtl"`
	// Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
	CacheSize *int `pulumi:"cacheSize"`
	// Specifies how many DoH concurrent queries are allowed.
	DohMaxConcurrentQueries *int `pulumi:"dohMaxConcurrentQueries"`
	// Specifies how many concurrent connections to the DoH server are allowed.
	DohMaxServerConnections *int `pulumi:"dohMaxServerConnections"`
	// Specifies how long to wait for query response from the DoH server.
	DohTimeout *string `pulumi:"dohTimeout"`
	// Specifies how much concurrent queries are allowed. *Default: 100*
	MaxConcurrentQueries *int `pulumi:"maxConcurrentQueries"`
	// Specifies how much concurrent TCP sessions are allowed. *Default: 20*
	MaxConcurrentTcpSessions *int `pulumi:"maxConcurrentTcpSessions"`
	// Maximum size of allowed UDP packet. *Default: 4096*
	MaxUdpPacketSize *int `pulumi:"maxUdpPacketSize"`
	// Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
	QueryServerTimeout *string `pulumi:"queryServerTimeout"`
	// Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
	// query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
	QueryTotalTimeout *string `pulumi:"queryTotalTimeout"`
	// List of DNS server IPv4/IPv6 addresses.
	Servers *string `pulumi:"servers"`
	// DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
	// fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
	// on the device.
	UseDohServer *string `pulumi:"useDohServer"`
	// DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
	VerifyDohCert *bool `pulumi:"verifyDohCert"`
}

// The set of arguments for constructing a IpDns resource.
type IpDnsArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// Specifies whether to allow network requests.
	AllowRemoteRequests pulumi.BoolPtrInput
	// Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
	// time. Shorter TTL received from DNS servers are respected. *Default: 1w*
	CacheMaxTtl pulumi.StringPtrInput
	// Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
	CacheSize pulumi.IntPtrInput
	// Specifies how many DoH concurrent queries are allowed.
	DohMaxConcurrentQueries pulumi.IntPtrInput
	// Specifies how many concurrent connections to the DoH server are allowed.
	DohMaxServerConnections pulumi.IntPtrInput
	// Specifies how long to wait for query response from the DoH server.
	DohTimeout pulumi.StringPtrInput
	// Specifies how much concurrent queries are allowed. *Default: 100*
	MaxConcurrentQueries pulumi.IntPtrInput
	// Specifies how much concurrent TCP sessions are allowed. *Default: 20*
	MaxConcurrentTcpSessions pulumi.IntPtrInput
	// Maximum size of allowed UDP packet. *Default: 4096*
	MaxUdpPacketSize pulumi.IntPtrInput
	// Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
	QueryServerTimeout pulumi.StringPtrInput
	// Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
	// query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
	QueryTotalTimeout pulumi.StringPtrInput
	// List of DNS server IPv4/IPv6 addresses.
	Servers pulumi.StringPtrInput
	// DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
	// fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
	// on the device.
	UseDohServer pulumi.StringPtrInput
	// DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
	VerifyDohCert pulumi.BoolPtrInput
}

func (IpDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipDnsArgs)(nil)).Elem()
}

type IpDnsInput interface {
	pulumi.Input

	ToIpDnsOutput() IpDnsOutput
	ToIpDnsOutputWithContext(ctx context.Context) IpDnsOutput
}

func (*IpDns) ElementType() reflect.Type {
	return reflect.TypeOf((**IpDns)(nil)).Elem()
}

func (i *IpDns) ToIpDnsOutput() IpDnsOutput {
	return i.ToIpDnsOutputWithContext(context.Background())
}

func (i *IpDns) ToIpDnsOutputWithContext(ctx context.Context) IpDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpDnsOutput)
}

// IpDnsArrayInput is an input type that accepts IpDnsArray and IpDnsArrayOutput values.
// You can construct a concrete instance of `IpDnsArrayInput` via:
//
//	IpDnsArray{ IpDnsArgs{...} }
type IpDnsArrayInput interface {
	pulumi.Input

	ToIpDnsArrayOutput() IpDnsArrayOutput
	ToIpDnsArrayOutputWithContext(context.Context) IpDnsArrayOutput
}

type IpDnsArray []IpDnsInput

func (IpDnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpDns)(nil)).Elem()
}

func (i IpDnsArray) ToIpDnsArrayOutput() IpDnsArrayOutput {
	return i.ToIpDnsArrayOutputWithContext(context.Background())
}

func (i IpDnsArray) ToIpDnsArrayOutputWithContext(ctx context.Context) IpDnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpDnsArrayOutput)
}

// IpDnsMapInput is an input type that accepts IpDnsMap and IpDnsMapOutput values.
// You can construct a concrete instance of `IpDnsMapInput` via:
//
//	IpDnsMap{ "key": IpDnsArgs{...} }
type IpDnsMapInput interface {
	pulumi.Input

	ToIpDnsMapOutput() IpDnsMapOutput
	ToIpDnsMapOutputWithContext(context.Context) IpDnsMapOutput
}

type IpDnsMap map[string]IpDnsInput

func (IpDnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpDns)(nil)).Elem()
}

func (i IpDnsMap) ToIpDnsMapOutput() IpDnsMapOutput {
	return i.ToIpDnsMapOutputWithContext(context.Background())
}

func (i IpDnsMap) ToIpDnsMapOutputWithContext(ctx context.Context) IpDnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpDnsMapOutput)
}

type IpDnsOutput struct{ *pulumi.OutputState }

func (IpDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpDns)(nil)).Elem()
}

func (o IpDnsOutput) ToIpDnsOutput() IpDnsOutput {
	return o
}

func (o IpDnsOutput) ToIpDnsOutputWithContext(ctx context.Context) IpDnsOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o IpDnsOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o IpDnsOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// Specifies whether to allow network requests.
func (o IpDnsOutput) AllowRemoteRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpDns) pulumi.BoolPtrOutput { return v.AllowRemoteRequests }).(pulumi.BoolPtrOutput)
}

// Maximum time-to-live for cache records. In other words, cache records will expire unconditionally after cache-max-ttl
// time. Shorter TTL received from DNS servers are respected. *Default: 1w*
func (o IpDnsOutput) CacheMaxTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringOutput { return v.CacheMaxTtl }).(pulumi.StringOutput)
}

// Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*
func (o IpDnsOutput) CacheSize() pulumi.IntOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntOutput { return v.CacheSize }).(pulumi.IntOutput)
}

// Shows the currently used cache size in KiB.
func (o IpDnsOutput) CacheUsed() pulumi.IntOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntOutput { return v.CacheUsed }).(pulumi.IntOutput)
}

// Specifies how many DoH concurrent queries are allowed.
func (o IpDnsOutput) DohMaxConcurrentQueries() pulumi.IntOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntOutput { return v.DohMaxConcurrentQueries }).(pulumi.IntOutput)
}

// Specifies how many concurrent connections to the DoH server are allowed.
func (o IpDnsOutput) DohMaxServerConnections() pulumi.IntOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntOutput { return v.DohMaxServerConnections }).(pulumi.IntOutput)
}

// Specifies how long to wait for query response from the DoH server.
func (o IpDnsOutput) DohTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringOutput { return v.DohTimeout }).(pulumi.StringOutput)
}

// List of dynamically added DNS server from different services, for example, DHCP.
func (o IpDnsOutput) DynamicServers() pulumi.StringOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringOutput { return v.DynamicServers }).(pulumi.StringOutput)
}

// Specifies how much concurrent queries are allowed. *Default: 100*
func (o IpDnsOutput) MaxConcurrentQueries() pulumi.IntOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntOutput { return v.MaxConcurrentQueries }).(pulumi.IntOutput)
}

// Specifies how much concurrent TCP sessions are allowed. *Default: 20*
func (o IpDnsOutput) MaxConcurrentTcpSessions() pulumi.IntOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntOutput { return v.MaxConcurrentTcpSessions }).(pulumi.IntOutput)
}

// Maximum size of allowed UDP packet. *Default: 4096*
func (o IpDnsOutput) MaxUdpPacketSize() pulumi.IntOutput {
	return o.ApplyT(func(v *IpDns) pulumi.IntOutput { return v.MaxUdpPacketSize }).(pulumi.IntOutput)
}

// Specifies how long to wait for query response from one server. Time can be specified in milliseconds. *Default: 2s*
func (o IpDnsOutput) QueryServerTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringOutput { return v.QueryServerTimeout }).(pulumi.StringOutput)
}

// Specifies how long to wait for query response in total. Note that this setting must be configured taking into account
// query_server_timeout and number of used DNS server. Time can be specified in milliseconds. *Default: 10s*
func (o IpDnsOutput) QueryTotalTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringOutput { return v.QueryTotalTimeout }).(pulumi.StringOutput)
}

// List of DNS server IPv4/IPv6 addresses.
func (o IpDnsOutput) Servers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringPtrOutput { return v.Servers }).(pulumi.StringPtrOutput)
}

// DNS over HTTPS (DoH) server URL. > Mikrotik strongly suggest not use third-party download links for certificate
// fetching. Use the Certificate Authority's own website. > RouterOS prioritize DoH over DNS server if both are configured
// on the device.
func (o IpDnsOutput) UseDohServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpDns) pulumi.StringPtrOutput { return v.UseDohServer }).(pulumi.StringPtrOutput)
}

// DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).
func (o IpDnsOutput) VerifyDohCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpDns) pulumi.BoolPtrOutput { return v.VerifyDohCert }).(pulumi.BoolPtrOutput)
}

type IpDnsArrayOutput struct{ *pulumi.OutputState }

func (IpDnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpDns)(nil)).Elem()
}

func (o IpDnsArrayOutput) ToIpDnsArrayOutput() IpDnsArrayOutput {
	return o
}

func (o IpDnsArrayOutput) ToIpDnsArrayOutputWithContext(ctx context.Context) IpDnsArrayOutput {
	return o
}

func (o IpDnsArrayOutput) Index(i pulumi.IntInput) IpDnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpDns {
		return vs[0].([]*IpDns)[vs[1].(int)]
	}).(IpDnsOutput)
}

type IpDnsMapOutput struct{ *pulumi.OutputState }

func (IpDnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpDns)(nil)).Elem()
}

func (o IpDnsMapOutput) ToIpDnsMapOutput() IpDnsMapOutput {
	return o
}

func (o IpDnsMapOutput) ToIpDnsMapOutputWithContext(ctx context.Context) IpDnsMapOutput {
	return o
}

func (o IpDnsMapOutput) MapIndex(k pulumi.StringInput) IpDnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpDns {
		return vs[0].(map[string]*IpDns)[vs[1].(string)]
	}).(IpDnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpDnsInput)(nil)).Elem(), &IpDns{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpDnsArrayInput)(nil)).Elem(), IpDnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpDnsMapInput)(nil)).Elem(), IpDnsMap{})
	pulumi.RegisterOutputType(IpDnsOutput{})
	pulumi.RegisterOutputType(IpDnsArrayOutput{})
	pulumi.RegisterOutputType(IpDnsMapOutput{})
}
