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
//			_, err := CapsMan.NewRates(ctx, "testRates", &CapsMan.RatesArgs{
//				Basics: pulumi.StringArray{
//					pulumi.String("1Mbps"),
//					pulumi.String("5.5Mbps"),
//					pulumi.String("6Mbps"),
//					pulumi.String("18Mbps"),
//					pulumi.String("36Mbps"),
//					pulumi.String("54Mbps"),
//				},
//				Comment: pulumi.String("test_rates"),
//				HtBasicMcs: pulumi.StringArray{
//					pulumi.String("mcs-0"),
//					pulumi.String("mcs-7"),
//					pulumi.String("mcs-11"),
//					pulumi.String("mcs-14"),
//					pulumi.String("mcs-16"),
//					pulumi.String("mcs-21"),
//				},
//				HtSupportedMcs: pulumi.StringArray{
//					pulumi.String("mcs-3"),
//					pulumi.String("mcs-8"),
//					pulumi.String("mcs-10"),
//					pulumi.String("mcs-13"),
//					pulumi.String("mcs-17"),
//					pulumi.String("mcs-18"),
//				},
//				Supporteds: pulumi.StringArray{
//					pulumi.String("2Mbps"),
//					pulumi.String("11Mbps"),
//					pulumi.String("9Mbps"),
//					pulumi.String("12Mbps"),
//					pulumi.String("24Mbps"),
//					pulumi.String("48Mbps"),
//				},
//				VhtBasicMcs:     pulumi.String("none"),
//				VhtSupportedMcs: pulumi.String("mcs0-9,mcs0-7"),
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
// # Import with the name of the CAPsMAN rates configuration in case of the example use test-rates-config
//
// ```sh
//
//	$ pulumi import routeros:CapsMan/rates:Rates test_rates test-rates-config
//
// ```
type Rates struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// List of basic rates. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP.
	Basics  pulumi.StringArrayOutput `pulumi:"basics"`
	Comment pulumi.StringPtrOutput   `pulumi:"comment"`
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
	HtBasicMcs pulumi.StringArrayOutput `pulumi:"htBasicMcs"`
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
	HtSupportedMcs pulumi.StringArrayOutput `pulumi:"htSupportedMcs"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringOutput `pulumi:"name"`
	// List of supported rates. Two devices will communicate only using rates that are supported by both devices.
	Supporteds pulumi.StringArrayOutput `pulumi:"supporteds"`
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9
	VhtBasicMcs pulumi.StringOutput `pulumi:"vhtBasicMcs"`
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9
	VhtSupportedMcs pulumi.StringOutput `pulumi:"vhtSupportedMcs"`
}

// NewRates registers a new resource with the given unique name, arguments, and options.
func NewRates(ctx *pulumi.Context,
	name string, args *RatesArgs, opts ...pulumi.ResourceOption) (*Rates, error) {
	if args == nil {
		args = &RatesArgs{}
	}

	var resource Rates
	err := ctx.RegisterResource("routeros:CapsMan/rates:Rates", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRates gets an existing Rates resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRates(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RatesState, opts ...pulumi.ResourceOption) (*Rates, error) {
	var resource Rates
	err := ctx.ReadResource("routeros:CapsMan/rates:Rates", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rates resources.
type ratesState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// List of basic rates. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP.
	Basics  []string `pulumi:"basics"`
	Comment *string  `pulumi:"comment"`
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
	HtBasicMcs []string `pulumi:"htBasicMcs"`
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
	HtSupportedMcs []string `pulumi:"htSupportedMcs"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// List of supported rates. Two devices will communicate only using rates that are supported by both devices.
	Supporteds []string `pulumi:"supporteds"`
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9
	VhtBasicMcs *string `pulumi:"vhtBasicMcs"`
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9
	VhtSupportedMcs *string `pulumi:"vhtSupportedMcs"`
}

type RatesState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// List of basic rates. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP.
	Basics  pulumi.StringArrayInput
	Comment pulumi.StringPtrInput
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
	HtBasicMcs pulumi.StringArrayInput
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
	HtSupportedMcs pulumi.StringArrayInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// List of supported rates. Two devices will communicate only using rates that are supported by both devices.
	Supporteds pulumi.StringArrayInput
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9
	VhtBasicMcs pulumi.StringPtrInput
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9
	VhtSupportedMcs pulumi.StringPtrInput
}

func (RatesState) ElementType() reflect.Type {
	return reflect.TypeOf((*ratesState)(nil)).Elem()
}

type ratesArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// List of basic rates. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP.
	Basics  []string `pulumi:"basics"`
	Comment *string  `pulumi:"comment"`
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
	HtBasicMcs []string `pulumi:"htBasicMcs"`
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
	HtSupportedMcs []string `pulumi:"htSupportedMcs"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// List of supported rates. Two devices will communicate only using rates that are supported by both devices.
	Supporteds []string `pulumi:"supporteds"`
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9
	VhtBasicMcs *string `pulumi:"vhtBasicMcs"`
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9
	VhtSupportedMcs *string `pulumi:"vhtSupportedMcs"`
}

// The set of arguments for constructing a Rates resource.
type RatesArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// List of basic rates. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP.
	Basics  pulumi.StringArrayInput
	Comment pulumi.StringPtrInput
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
	HtBasicMcs pulumi.StringArrayInput
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
	HtSupportedMcs pulumi.StringArrayInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// List of supported rates. Two devices will communicate only using rates that are supported by both devices.
	Supporteds pulumi.StringArrayInput
	// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9
	VhtBasicMcs pulumi.StringPtrInput
	// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9
	VhtSupportedMcs pulumi.StringPtrInput
}

func (RatesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ratesArgs)(nil)).Elem()
}

type RatesInput interface {
	pulumi.Input

	ToRatesOutput() RatesOutput
	ToRatesOutputWithContext(ctx context.Context) RatesOutput
}

func (*Rates) ElementType() reflect.Type {
	return reflect.TypeOf((**Rates)(nil)).Elem()
}

func (i *Rates) ToRatesOutput() RatesOutput {
	return i.ToRatesOutputWithContext(context.Background())
}

func (i *Rates) ToRatesOutputWithContext(ctx context.Context) RatesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RatesOutput)
}

// RatesArrayInput is an input type that accepts RatesArray and RatesArrayOutput values.
// You can construct a concrete instance of `RatesArrayInput` via:
//
//	RatesArray{ RatesArgs{...} }
type RatesArrayInput interface {
	pulumi.Input

	ToRatesArrayOutput() RatesArrayOutput
	ToRatesArrayOutputWithContext(context.Context) RatesArrayOutput
}

type RatesArray []RatesInput

func (RatesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rates)(nil)).Elem()
}

func (i RatesArray) ToRatesArrayOutput() RatesArrayOutput {
	return i.ToRatesArrayOutputWithContext(context.Background())
}

func (i RatesArray) ToRatesArrayOutputWithContext(ctx context.Context) RatesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RatesArrayOutput)
}

// RatesMapInput is an input type that accepts RatesMap and RatesMapOutput values.
// You can construct a concrete instance of `RatesMapInput` via:
//
//	RatesMap{ "key": RatesArgs{...} }
type RatesMapInput interface {
	pulumi.Input

	ToRatesMapOutput() RatesMapOutput
	ToRatesMapOutputWithContext(context.Context) RatesMapOutput
}

type RatesMap map[string]RatesInput

func (RatesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rates)(nil)).Elem()
}

func (i RatesMap) ToRatesMapOutput() RatesMapOutput {
	return i.ToRatesMapOutputWithContext(context.Background())
}

func (i RatesMap) ToRatesMapOutputWithContext(ctx context.Context) RatesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RatesMapOutput)
}

type RatesOutput struct{ *pulumi.OutputState }

func (RatesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rates)(nil)).Elem()
}

func (o RatesOutput) ToRatesOutput() RatesOutput {
	return o
}

func (o RatesOutput) ToRatesOutputWithContext(ctx context.Context) RatesOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o RatesOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rates) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o RatesOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// List of basic rates. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP.
func (o RatesOutput) Basics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringArrayOutput { return v.Basics }).(pulumi.StringArrayOutput)
}

func (o RatesOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
func (o RatesOutput) HtBasicMcs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringArrayOutput { return v.HtBasicMcs }).(pulumi.StringArrayOutput)
}

// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
func (o RatesOutput) HtSupportedMcs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringArrayOutput { return v.HtSupportedMcs }).(pulumi.StringArrayOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o RatesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of supported rates. Two devices will communicate only using rates that are supported by both devices.
func (o RatesOutput) Supporteds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringArrayOutput { return v.Supporteds }).(pulumi.StringArrayOutput)
}

// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9
func (o RatesOutput) VhtBasicMcs() pulumi.StringOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringOutput { return v.VhtBasicMcs }).(pulumi.StringOutput)
}

// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9
func (o RatesOutput) VhtSupportedMcs() pulumi.StringOutput {
	return o.ApplyT(func(v *Rates) pulumi.StringOutput { return v.VhtSupportedMcs }).(pulumi.StringOutput)
}

type RatesArrayOutput struct{ *pulumi.OutputState }

func (RatesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rates)(nil)).Elem()
}

func (o RatesArrayOutput) ToRatesArrayOutput() RatesArrayOutput {
	return o
}

func (o RatesArrayOutput) ToRatesArrayOutputWithContext(ctx context.Context) RatesArrayOutput {
	return o
}

func (o RatesArrayOutput) Index(i pulumi.IntInput) RatesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rates {
		return vs[0].([]*Rates)[vs[1].(int)]
	}).(RatesOutput)
}

type RatesMapOutput struct{ *pulumi.OutputState }

func (RatesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rates)(nil)).Elem()
}

func (o RatesMapOutput) ToRatesMapOutput() RatesMapOutput {
	return o
}

func (o RatesMapOutput) ToRatesMapOutputWithContext(ctx context.Context) RatesMapOutput {
	return o
}

func (o RatesMapOutput) MapIndex(k pulumi.StringInput) RatesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rates {
		return vs[0].(map[string]*Rates)[vs[1].(string)]
	}).(RatesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RatesInput)(nil)).Elem(), &Rates{})
	pulumi.RegisterInputType(reflect.TypeOf((*RatesArrayInput)(nil)).Elem(), RatesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RatesMapInput)(nil)).Elem(), RatesMap{})
	pulumi.RegisterOutputType(RatesOutput{})
	pulumi.RegisterOutputType(RatesArrayOutput{})
	pulumi.RegisterOutputType(RatesMapOutput{})
}