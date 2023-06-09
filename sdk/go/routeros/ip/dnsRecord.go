// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Ip.DnsRecord (Resource)
//
// ***
//
// #### This is an alias for backwards compatibility between plugin versions.
// Please see documentation for Ip.IpDnsRecord
type DnsRecord struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// The A record to be returend from the DNS hostname.
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
	AddressList pulumi.StringPtrOutput `pulumi:"addressList"`
	// Alias name for a domain name.
	Cname    pulumi.StringPtrOutput `pulumi:"cname"`
	Comment  pulumi.StringPtrOutput `pulumi:"comment"`
	Disabled pulumi.BoolPtrOutput   `pulumi:"disabled"`
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic pulumi.BoolOutput `pulumi:"dynamic"`
	// The IP address of a domain name server to which a particular DNS request must be forwarded.
	ForwardTo pulumi.StringPtrOutput `pulumi:"forwardTo"`
	// Whether the record will match requests for subdomains.
	MatchSubdomain pulumi.BoolPtrOutput `pulumi:"matchSubdomain"`
	// The domain name of the MX server.
	MxExchange pulumi.StringPtrOutput `pulumi:"mxExchange"`
	// Preference of the particular MX record.
	MxPreference pulumi.IntOutput `pulumi:"mxPreference"`
	// The name of the DNS hostname to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the authoritative domain name server for the particular record.
	Ns pulumi.StringPtrOutput `pulumi:"ns"`
	// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS
	// names to lowercase, you should write regex only with lowercase letters.
	Regexp pulumi.StringPtrOutput `pulumi:"regexp"`
	// The TCP or UDP port on which the service is to be found.
	SrvPort pulumi.IntOutput `pulumi:"srvPort"`
	// Priority of the particular SRV record.
	SrvPriority pulumi.IntOutput `pulumi:"srvPriority"`
	// The canonical hostname of the machine providing the service ends in a dot.
	SrvTarget pulumi.StringPtrOutput `pulumi:"srvTarget"`
	// Weight of the particular SRC record.
	SrvWeight pulumi.StringOutput `pulumi:"srvWeight"`
	// Textual information about the domain name.
	Text pulumi.StringPtrOutput `pulumi:"text"`
	// The ttl of the DNS record.
	Ttl pulumi.StringOutput `pulumi:"ttl"`
	// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDnsRecord registers a new resource with the given unique name, arguments, and options.
func NewDnsRecord(ctx *pulumi.Context,
	name string, args *DnsRecordArgs, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource DnsRecord
	err := ctx.RegisterResource("routeros:Ip/dnsRecord:DnsRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsRecord gets an existing DnsRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsRecordState, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	var resource DnsRecord
	err := ctx.ReadResource("routeros:Ip/dnsRecord:DnsRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsRecord resources.
type dnsRecordState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// The A record to be returend from the DNS hostname.
	Address *string `pulumi:"address"`
	// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
	AddressList *string `pulumi:"addressList"`
	// Alias name for a domain name.
	Cname    *string `pulumi:"cname"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic *bool `pulumi:"dynamic"`
	// The IP address of a domain name server to which a particular DNS request must be forwarded.
	ForwardTo *string `pulumi:"forwardTo"`
	// Whether the record will match requests for subdomains.
	MatchSubdomain *bool `pulumi:"matchSubdomain"`
	// The domain name of the MX server.
	MxExchange *string `pulumi:"mxExchange"`
	// Preference of the particular MX record.
	MxPreference *int `pulumi:"mxPreference"`
	// The name of the DNS hostname to be created.
	Name *string `pulumi:"name"`
	// Name of the authoritative domain name server for the particular record.
	Ns *string `pulumi:"ns"`
	// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS
	// names to lowercase, you should write regex only with lowercase letters.
	Regexp *string `pulumi:"regexp"`
	// The TCP or UDP port on which the service is to be found.
	SrvPort *int `pulumi:"srvPort"`
	// Priority of the particular SRV record.
	SrvPriority *int `pulumi:"srvPriority"`
	// The canonical hostname of the machine providing the service ends in a dot.
	SrvTarget *string `pulumi:"srvTarget"`
	// Weight of the particular SRC record.
	SrvWeight *string `pulumi:"srvWeight"`
	// Textual information about the domain name.
	Text *string `pulumi:"text"`
	// The ttl of the DNS record.
	Ttl *string `pulumi:"ttl"`
	// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
	Type *string `pulumi:"type"`
}

type DnsRecordState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// The A record to be returend from the DNS hostname.
	Address pulumi.StringPtrInput
	// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
	AddressList pulumi.StringPtrInput
	// Alias name for a domain name.
	Cname    pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
	// modified.
	Dynamic pulumi.BoolPtrInput
	// The IP address of a domain name server to which a particular DNS request must be forwarded.
	ForwardTo pulumi.StringPtrInput
	// Whether the record will match requests for subdomains.
	MatchSubdomain pulumi.BoolPtrInput
	// The domain name of the MX server.
	MxExchange pulumi.StringPtrInput
	// Preference of the particular MX record.
	MxPreference pulumi.IntPtrInput
	// The name of the DNS hostname to be created.
	Name pulumi.StringPtrInput
	// Name of the authoritative domain name server for the particular record.
	Ns pulumi.StringPtrInput
	// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS
	// names to lowercase, you should write regex only with lowercase letters.
	Regexp pulumi.StringPtrInput
	// The TCP or UDP port on which the service is to be found.
	SrvPort pulumi.IntPtrInput
	// Priority of the particular SRV record.
	SrvPriority pulumi.IntPtrInput
	// The canonical hostname of the machine providing the service ends in a dot.
	SrvTarget pulumi.StringPtrInput
	// Weight of the particular SRC record.
	SrvWeight pulumi.StringPtrInput
	// Textual information about the domain name.
	Text pulumi.StringPtrInput
	// The ttl of the DNS record.
	Ttl pulumi.StringPtrInput
	// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
	Type pulumi.StringPtrInput
}

func (DnsRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordState)(nil)).Elem()
}

type dnsRecordArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// The A record to be returend from the DNS hostname.
	Address *string `pulumi:"address"`
	// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
	AddressList *string `pulumi:"addressList"`
	// Alias name for a domain name.
	Cname    *string `pulumi:"cname"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// The IP address of a domain name server to which a particular DNS request must be forwarded.
	ForwardTo *string `pulumi:"forwardTo"`
	// Whether the record will match requests for subdomains.
	MatchSubdomain *bool `pulumi:"matchSubdomain"`
	// The domain name of the MX server.
	MxExchange *string `pulumi:"mxExchange"`
	// Preference of the particular MX record.
	MxPreference *int `pulumi:"mxPreference"`
	// The name of the DNS hostname to be created.
	Name *string `pulumi:"name"`
	// Name of the authoritative domain name server for the particular record.
	Ns *string `pulumi:"ns"`
	// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS
	// names to lowercase, you should write regex only with lowercase letters.
	Regexp *string `pulumi:"regexp"`
	// The TCP or UDP port on which the service is to be found.
	SrvPort *int `pulumi:"srvPort"`
	// Priority of the particular SRV record.
	SrvPriority *int `pulumi:"srvPriority"`
	// The canonical hostname of the machine providing the service ends in a dot.
	SrvTarget *string `pulumi:"srvTarget"`
	// Weight of the particular SRC record.
	SrvWeight *string `pulumi:"srvWeight"`
	// Textual information about the domain name.
	Text *string `pulumi:"text"`
	// The ttl of the DNS record.
	Ttl *string `pulumi:"ttl"`
	// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DnsRecord resource.
type DnsRecordArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// The A record to be returend from the DNS hostname.
	Address pulumi.StringPtrInput
	// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
	AddressList pulumi.StringPtrInput
	// Alias name for a domain name.
	Cname    pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// The IP address of a domain name server to which a particular DNS request must be forwarded.
	ForwardTo pulumi.StringPtrInput
	// Whether the record will match requests for subdomains.
	MatchSubdomain pulumi.BoolPtrInput
	// The domain name of the MX server.
	MxExchange pulumi.StringPtrInput
	// Preference of the particular MX record.
	MxPreference pulumi.IntPtrInput
	// The name of the DNS hostname to be created.
	Name pulumi.StringPtrInput
	// Name of the authoritative domain name server for the particular record.
	Ns pulumi.StringPtrInput
	// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS
	// names to lowercase, you should write regex only with lowercase letters.
	Regexp pulumi.StringPtrInput
	// The TCP or UDP port on which the service is to be found.
	SrvPort pulumi.IntPtrInput
	// Priority of the particular SRV record.
	SrvPriority pulumi.IntPtrInput
	// The canonical hostname of the machine providing the service ends in a dot.
	SrvTarget pulumi.StringPtrInput
	// Weight of the particular SRC record.
	SrvWeight pulumi.StringPtrInput
	// Textual information about the domain name.
	Text pulumi.StringPtrInput
	// The ttl of the DNS record.
	Ttl pulumi.StringPtrInput
	// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
	Type pulumi.StringInput
}

func (DnsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordArgs)(nil)).Elem()
}

type DnsRecordInput interface {
	pulumi.Input

	ToDnsRecordOutput() DnsRecordOutput
	ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput
}

func (*DnsRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (i *DnsRecord) ToDnsRecordOutput() DnsRecordOutput {
	return i.ToDnsRecordOutputWithContext(context.Background())
}

func (i *DnsRecord) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordOutput)
}

// DnsRecordArrayInput is an input type that accepts DnsRecordArray and DnsRecordArrayOutput values.
// You can construct a concrete instance of `DnsRecordArrayInput` via:
//
//	DnsRecordArray{ DnsRecordArgs{...} }
type DnsRecordArrayInput interface {
	pulumi.Input

	ToDnsRecordArrayOutput() DnsRecordArrayOutput
	ToDnsRecordArrayOutputWithContext(context.Context) DnsRecordArrayOutput
}

type DnsRecordArray []DnsRecordInput

func (DnsRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordArray) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return i.ToDnsRecordArrayOutputWithContext(context.Background())
}

func (i DnsRecordArray) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordArrayOutput)
}

// DnsRecordMapInput is an input type that accepts DnsRecordMap and DnsRecordMapOutput values.
// You can construct a concrete instance of `DnsRecordMapInput` via:
//
//	DnsRecordMap{ "key": DnsRecordArgs{...} }
type DnsRecordMapInput interface {
	pulumi.Input

	ToDnsRecordMapOutput() DnsRecordMapOutput
	ToDnsRecordMapOutputWithContext(context.Context) DnsRecordMapOutput
}

type DnsRecordMap map[string]DnsRecordInput

func (DnsRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordMap) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return i.ToDnsRecordMapOutputWithContext(context.Background())
}

func (i DnsRecordMap) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordMapOutput)
}

type DnsRecordOutput struct{ *pulumi.OutputState }

func (DnsRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (o DnsRecordOutput) ToDnsRecordOutput() DnsRecordOutput {
	return o
}

func (o DnsRecordOutput) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o DnsRecordOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o DnsRecordOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// The A record to be returend from the DNS hostname.
func (o DnsRecordOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
func (o DnsRecordOutput) AddressList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.AddressList }).(pulumi.StringPtrOutput)
}

// Alias name for a domain name.
func (o DnsRecordOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.Cname }).(pulumi.StringPtrOutput)
}

func (o DnsRecordOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o DnsRecordOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Configuration item created by software, not by management interface. It is not exported, and cannot be directly
// modified.
func (o DnsRecordOutput) Dynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.BoolOutput { return v.Dynamic }).(pulumi.BoolOutput)
}

// The IP address of a domain name server to which a particular DNS request must be forwarded.
func (o DnsRecordOutput) ForwardTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.ForwardTo }).(pulumi.StringPtrOutput)
}

// Whether the record will match requests for subdomains.
func (o DnsRecordOutput) MatchSubdomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.BoolPtrOutput { return v.MatchSubdomain }).(pulumi.BoolPtrOutput)
}

// The domain name of the MX server.
func (o DnsRecordOutput) MxExchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.MxExchange }).(pulumi.StringPtrOutput)
}

// Preference of the particular MX record.
func (o DnsRecordOutput) MxPreference() pulumi.IntOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntOutput { return v.MxPreference }).(pulumi.IntOutput)
}

// The name of the DNS hostname to be created.
func (o DnsRecordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of the authoritative domain name server for the particular record.
func (o DnsRecordOutput) Ns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.Ns }).(pulumi.StringPtrOutput)
}

// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS
// names to lowercase, you should write regex only with lowercase letters.
func (o DnsRecordOutput) Regexp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.Regexp }).(pulumi.StringPtrOutput)
}

// The TCP or UDP port on which the service is to be found.
func (o DnsRecordOutput) SrvPort() pulumi.IntOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntOutput { return v.SrvPort }).(pulumi.IntOutput)
}

// Priority of the particular SRV record.
func (o DnsRecordOutput) SrvPriority() pulumi.IntOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntOutput { return v.SrvPriority }).(pulumi.IntOutput)
}

// The canonical hostname of the machine providing the service ends in a dot.
func (o DnsRecordOutput) SrvTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.SrvTarget }).(pulumi.StringPtrOutput)
}

// Weight of the particular SRC record.
func (o DnsRecordOutput) SrvWeight() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.SrvWeight }).(pulumi.StringOutput)
}

// Textual information about the domain name.
func (o DnsRecordOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.Text }).(pulumi.StringPtrOutput)
}

// The ttl of the DNS record.
func (o DnsRecordOutput) Ttl() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Ttl }).(pulumi.StringOutput)
}

// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
func (o DnsRecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DnsRecordArrayOutput struct{ *pulumi.OutputState }

func (DnsRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) Index(i pulumi.IntInput) DnsRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].([]*DnsRecord)[vs[1].(int)]
	}).(DnsRecordOutput)
}

type DnsRecordMapOutput struct{ *pulumi.OutputState }

func (DnsRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) MapIndex(k pulumi.StringInput) DnsRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].(map[string]*DnsRecord)[vs[1].(string)]
	}).(DnsRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordInput)(nil)).Elem(), &DnsRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordArrayInput)(nil)).Elem(), DnsRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordMapInput)(nil)).Elem(), DnsRecordMap{})
	pulumi.RegisterOutputType(DnsRecordOutput{})
	pulumi.RegisterOutputType(DnsRecordArrayOutput{})
	pulumi.RegisterOutputType(DnsRecordMapOutput{})
}
