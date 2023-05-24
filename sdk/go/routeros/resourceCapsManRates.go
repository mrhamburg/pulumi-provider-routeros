// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package routeros

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
//	"github.com/pulumi/pulumi-routeros/sdk/go/routeros"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := routeros.NewResourceCapsManRates(ctx, "testRates", &routeros.ResourceCapsManRatesArgs{
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
//	$ pulumi import routeros:index/resourceCapsManRates:ResourceCapsManRates test_rates test-rates-config
//
// ```
type ResourceCapsManRates struct {
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

// NewResourceCapsManRates registers a new resource with the given unique name, arguments, and options.
func NewResourceCapsManRates(ctx *pulumi.Context,
	name string, args *ResourceCapsManRatesArgs, opts ...pulumi.ResourceOption) (*ResourceCapsManRates, error) {
	if args == nil {
		args = &ResourceCapsManRatesArgs{}
	}

	var resource ResourceCapsManRates
	err := ctx.RegisterResource("routeros:index/resourceCapsManRates:ResourceCapsManRates", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceCapsManRates gets an existing ResourceCapsManRates resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceCapsManRates(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceCapsManRatesState, opts ...pulumi.ResourceOption) (*ResourceCapsManRates, error) {
	var resource ResourceCapsManRates
	err := ctx.ReadResource("routeros:index/resourceCapsManRates:ResourceCapsManRates", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceCapsManRates resources.
type resourceCapsManRatesState struct {
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

type ResourceCapsManRatesState struct {
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

func (ResourceCapsManRatesState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceCapsManRatesState)(nil)).Elem()
}

type resourceCapsManRatesArgs struct {
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

// The set of arguments for constructing a ResourceCapsManRates resource.
type ResourceCapsManRatesArgs struct {
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

func (ResourceCapsManRatesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceCapsManRatesArgs)(nil)).Elem()
}

type ResourceCapsManRatesInput interface {
	pulumi.Input

	ToResourceCapsManRatesOutput() ResourceCapsManRatesOutput
	ToResourceCapsManRatesOutputWithContext(ctx context.Context) ResourceCapsManRatesOutput
}

func (*ResourceCapsManRates) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceCapsManRates)(nil)).Elem()
}

func (i *ResourceCapsManRates) ToResourceCapsManRatesOutput() ResourceCapsManRatesOutput {
	return i.ToResourceCapsManRatesOutputWithContext(context.Background())
}

func (i *ResourceCapsManRates) ToResourceCapsManRatesOutputWithContext(ctx context.Context) ResourceCapsManRatesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCapsManRatesOutput)
}

// ResourceCapsManRatesArrayInput is an input type that accepts ResourceCapsManRatesArray and ResourceCapsManRatesArrayOutput values.
// You can construct a concrete instance of `ResourceCapsManRatesArrayInput` via:
//
//	ResourceCapsManRatesArray{ ResourceCapsManRatesArgs{...} }
type ResourceCapsManRatesArrayInput interface {
	pulumi.Input

	ToResourceCapsManRatesArrayOutput() ResourceCapsManRatesArrayOutput
	ToResourceCapsManRatesArrayOutputWithContext(context.Context) ResourceCapsManRatesArrayOutput
}

type ResourceCapsManRatesArray []ResourceCapsManRatesInput

func (ResourceCapsManRatesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceCapsManRates)(nil)).Elem()
}

func (i ResourceCapsManRatesArray) ToResourceCapsManRatesArrayOutput() ResourceCapsManRatesArrayOutput {
	return i.ToResourceCapsManRatesArrayOutputWithContext(context.Background())
}

func (i ResourceCapsManRatesArray) ToResourceCapsManRatesArrayOutputWithContext(ctx context.Context) ResourceCapsManRatesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCapsManRatesArrayOutput)
}

// ResourceCapsManRatesMapInput is an input type that accepts ResourceCapsManRatesMap and ResourceCapsManRatesMapOutput values.
// You can construct a concrete instance of `ResourceCapsManRatesMapInput` via:
//
//	ResourceCapsManRatesMap{ "key": ResourceCapsManRatesArgs{...} }
type ResourceCapsManRatesMapInput interface {
	pulumi.Input

	ToResourceCapsManRatesMapOutput() ResourceCapsManRatesMapOutput
	ToResourceCapsManRatesMapOutputWithContext(context.Context) ResourceCapsManRatesMapOutput
}

type ResourceCapsManRatesMap map[string]ResourceCapsManRatesInput

func (ResourceCapsManRatesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceCapsManRates)(nil)).Elem()
}

func (i ResourceCapsManRatesMap) ToResourceCapsManRatesMapOutput() ResourceCapsManRatesMapOutput {
	return i.ToResourceCapsManRatesMapOutputWithContext(context.Background())
}

func (i ResourceCapsManRatesMap) ToResourceCapsManRatesMapOutputWithContext(ctx context.Context) ResourceCapsManRatesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCapsManRatesMapOutput)
}

type ResourceCapsManRatesOutput struct{ *pulumi.OutputState }

func (ResourceCapsManRatesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceCapsManRates)(nil)).Elem()
}

func (o ResourceCapsManRatesOutput) ToResourceCapsManRatesOutput() ResourceCapsManRatesOutput {
	return o
}

func (o ResourceCapsManRatesOutput) ToResourceCapsManRatesOutputWithContext(ctx context.Context) ResourceCapsManRatesOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o ResourceCapsManRatesOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o ResourceCapsManRatesOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// List of basic rates. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP.
func (o ResourceCapsManRatesOutput) Basics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringArrayOutput { return v.Basics }).(pulumi.StringArrayOutput)
}

func (o ResourceCapsManRatesOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
func (o ResourceCapsManRatesOutput) HtBasicMcs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringArrayOutput { return v.HtBasicMcs }).(pulumi.StringArrayOutput)
}

// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
func (o ResourceCapsManRatesOutput) HtSupportedMcs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringArrayOutput { return v.HtSupportedMcs }).(pulumi.StringArrayOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o ResourceCapsManRatesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of supported rates. Two devices will communicate only using rates that are supported by both devices.
func (o ResourceCapsManRatesOutput) Supporteds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringArrayOutput { return v.Supporteds }).(pulumi.StringArrayOutput)
}

// Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9
func (o ResourceCapsManRatesOutput) VhtBasicMcs() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringOutput { return v.VhtBasicMcs }).(pulumi.StringOutput)
}

// Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9
func (o ResourceCapsManRatesOutput) VhtSupportedMcs() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceCapsManRates) pulumi.StringOutput { return v.VhtSupportedMcs }).(pulumi.StringOutput)
}

type ResourceCapsManRatesArrayOutput struct{ *pulumi.OutputState }

func (ResourceCapsManRatesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceCapsManRates)(nil)).Elem()
}

func (o ResourceCapsManRatesArrayOutput) ToResourceCapsManRatesArrayOutput() ResourceCapsManRatesArrayOutput {
	return o
}

func (o ResourceCapsManRatesArrayOutput) ToResourceCapsManRatesArrayOutputWithContext(ctx context.Context) ResourceCapsManRatesArrayOutput {
	return o
}

func (o ResourceCapsManRatesArrayOutput) Index(i pulumi.IntInput) ResourceCapsManRatesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceCapsManRates {
		return vs[0].([]*ResourceCapsManRates)[vs[1].(int)]
	}).(ResourceCapsManRatesOutput)
}

type ResourceCapsManRatesMapOutput struct{ *pulumi.OutputState }

func (ResourceCapsManRatesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceCapsManRates)(nil)).Elem()
}

func (o ResourceCapsManRatesMapOutput) ToResourceCapsManRatesMapOutput() ResourceCapsManRatesMapOutput {
	return o
}

func (o ResourceCapsManRatesMapOutput) ToResourceCapsManRatesMapOutputWithContext(ctx context.Context) ResourceCapsManRatesMapOutput {
	return o
}

func (o ResourceCapsManRatesMapOutput) MapIndex(k pulumi.StringInput) ResourceCapsManRatesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceCapsManRates {
		return vs[0].(map[string]*ResourceCapsManRates)[vs[1].(string)]
	}).(ResourceCapsManRatesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCapsManRatesInput)(nil)).Elem(), &ResourceCapsManRates{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCapsManRatesArrayInput)(nil)).Elem(), ResourceCapsManRatesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCapsManRatesMapInput)(nil)).Elem(), ResourceCapsManRatesMap{})
	pulumi.RegisterOutputType(ResourceCapsManRatesOutput{})
	pulumi.RegisterOutputType(ResourceCapsManRatesArrayOutput{})
	pulumi.RegisterOutputType(ResourceCapsManRatesMapOutput{})
}
