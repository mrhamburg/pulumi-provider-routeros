// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package routeros

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # DatasourceFirewall (Data Source)
//
// This datasource contains all supported firewall resources:
// - addressList
// - nat
// - mangle
// - rules (aka filter)
func DatasourceFirewall(ctx *pulumi.Context, args *DatasourceFirewallArgs, opts ...pulumi.InvokeOption) (*DatasourceFirewallResult, error) {
	var rv DatasourceFirewallResult
	err := ctx.Invoke("routeros:index/datasourceFirewall:DatasourceFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking DatasourceFirewall.
type DatasourceFirewallArgs struct {
	AddressLists []DatasourceFirewallAddressList `pulumi:"addressLists"`
	Mangles      []DatasourceFirewallMangle      `pulumi:"mangles"`
	Nats         []DatasourceFirewallNat         `pulumi:"nats"`
	Rules        []DatasourceFirewallRule        `pulumi:"rules"`
}

// A collection of values returned by DatasourceFirewall.
type DatasourceFirewallResult struct {
	AddressLists []DatasourceFirewallAddressList `pulumi:"addressLists"`
	// The provider-assigned unique ID for this managed resource.
	Id      string                     `pulumi:"id"`
	Mangles []DatasourceFirewallMangle `pulumi:"mangles"`
	Nats    []DatasourceFirewallNat    `pulumi:"nats"`
	Rules   []DatasourceFirewallRule   `pulumi:"rules"`
}

func DatasourceFirewallOutput(ctx *pulumi.Context, args DatasourceFirewallOutputArgs, opts ...pulumi.InvokeOption) DatasourceFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (DatasourceFirewallResult, error) {
			args := v.(DatasourceFirewallArgs)
			r, err := DatasourceFirewall(ctx, &args, opts...)
			var s DatasourceFirewallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(DatasourceFirewallResultOutput)
}

// A collection of arguments for invoking DatasourceFirewall.
type DatasourceFirewallOutputArgs struct {
	AddressLists DatasourceFirewallAddressListArrayInput `pulumi:"addressLists"`
	Mangles      DatasourceFirewallMangleArrayInput      `pulumi:"mangles"`
	Nats         DatasourceFirewallNatArrayInput         `pulumi:"nats"`
	Rules        DatasourceFirewallRuleArrayInput        `pulumi:"rules"`
}

func (DatasourceFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceFirewallArgs)(nil)).Elem()
}

// A collection of values returned by DatasourceFirewall.
type DatasourceFirewallResultOutput struct{ *pulumi.OutputState }

func (DatasourceFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceFirewallResult)(nil)).Elem()
}

func (o DatasourceFirewallResultOutput) ToDatasourceFirewallResultOutput() DatasourceFirewallResultOutput {
	return o
}

func (o DatasourceFirewallResultOutput) ToDatasourceFirewallResultOutputWithContext(ctx context.Context) DatasourceFirewallResultOutput {
	return o
}

func (o DatasourceFirewallResultOutput) AddressLists() DatasourceFirewallAddressListArrayOutput {
	return o.ApplyT(func(v DatasourceFirewallResult) []DatasourceFirewallAddressList { return v.AddressLists }).(DatasourceFirewallAddressListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o DatasourceFirewallResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DatasourceFirewallResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o DatasourceFirewallResultOutput) Mangles() DatasourceFirewallMangleArrayOutput {
	return o.ApplyT(func(v DatasourceFirewallResult) []DatasourceFirewallMangle { return v.Mangles }).(DatasourceFirewallMangleArrayOutput)
}

func (o DatasourceFirewallResultOutput) Nats() DatasourceFirewallNatArrayOutput {
	return o.ApplyT(func(v DatasourceFirewallResult) []DatasourceFirewallNat { return v.Nats }).(DatasourceFirewallNatArrayOutput)
}

func (o DatasourceFirewallResultOutput) Rules() DatasourceFirewallRuleArrayOutput {
	return o.ApplyT(func(v DatasourceFirewallResult) []DatasourceFirewallRule { return v.Rules }).(DatasourceFirewallRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasourceFirewallResultOutput{})
}
