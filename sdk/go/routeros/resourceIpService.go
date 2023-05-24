// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package routeros

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
//	"github.com/pulumi/pulumi-routeros/sdk/go/routeros"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tlsService := map[string]interface{}{
//				"api-ssl": 8729,
//				"www-ssl": 443,
//			}
//			disableService := map[string]interface{}{
//				"api":    8728,
//				"ftp":    21,
//				"telnet": 23,
//				"www":    80,
//			}
//			enableService := map[string]interface{}{
//				"ssh":    22,
//				"winbox": 8291,
//			}
//			tlsCert, err := routeros.NewResourceSystemCertificate(ctx, "tlsCert", &routeros.ResourceSystemCertificateArgs{
//				CommonName: pulumi.String("Mikrotik Router"),
//				DaysValid:  pulumi.Int(3650),
//				KeyUsages: pulumi.StringArray{
//					pulumi.String("key-cert-sign"),
//					pulumi.String("crl-sign"),
//					pulumi.String("digital-signature"),
//					pulumi.String("key-agreement"),
//					pulumi.String("tls-server"),
//				},
//				KeySize: pulumi.String("prime256v1"),
//				Signs: routeros.ResourceSystemCertificateSignArray{
//					nil,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			var tls []*routeros.ResourceIpService
//			for key0, val0 := range tlsService {
//				__res, err := routeros.NewResourceIpService(ctx, fmt.Sprintf("tls-%v", key0), &routeros.ResourceIpServiceArgs{
//					Numbers:     pulumi.String(key0),
//					Port:        pulumi.Float64(val0),
//					Certificate: tlsCert.Name,
//					TlsVersion:  pulumi.String("only-1.2"),
//					Disabled:    pulumi.Bool(false),
//				})
//				if err != nil {
//					return err
//				}
//				tls = append(tls, __res)
//			}
//			var disabled []*routeros.ResourceIpService
//			for key0, val0 := range disableService {
//				__res, err := routeros.NewResourceIpService(ctx, fmt.Sprintf("disabled-%v", key0), &routeros.ResourceIpServiceArgs{
//					Numbers:  pulumi.String(key0),
//					Port:     pulumi.Float64(val0),
//					Disabled: pulumi.Bool(true),
//				})
//				if err != nil {
//					return err
//				}
//				disabled = append(disabled, __res)
//			}
//			var enabled []*routeros.ResourceIpService
//			for key0, val0 := range enableService {
//				__res, err := routeros.NewResourceIpService(ctx, fmt.Sprintf("enabled-%v", key0), &routeros.ResourceIpServiceArgs{
//					Numbers:  pulumi.String(key0),
//					Port:     pulumi.Float64(val0),
//					Disabled: pulumi.Bool(false),
//				})
//				if err != nil {
//					return err
//				}
//				enabled = append(enabled, __res)
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Import with the name of the ip service in case of the example use www-ssl
//
// ```sh
//
//	$ pulumi import routeros:index/resourceIpService:ResourceIpService www_ssl www-ssl
//
// ```
type ResourceIpService struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	// List of IP/IPv6 prefixes from which the service is accessible.
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// The name of the certificate used by a particular service. Applicable only for services that depend on certificates ( www-ssl, api-ssl ).
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	Disabled    pulumi.BoolPtrOutput   `pulumi:"disabled"`
	Invalid     pulumi.BoolOutput      `pulumi:"invalid"`
	// Service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the service whose settings will be changed ( api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl ).
	Numbers pulumi.StringOutput `pulumi:"numbers"`
	// The port particular service listens on.
	Port pulumi.IntOutput `pulumi:"port"`
	// Specifies which TLS versions to allow by a particular service.
	TlsVersion pulumi.StringPtrOutput `pulumi:"tlsVersion"`
	// Specify which VRF instance to use by a particular service.
	Vrf pulumi.StringPtrOutput `pulumi:"vrf"`
}

// NewResourceIpService registers a new resource with the given unique name, arguments, and options.
func NewResourceIpService(ctx *pulumi.Context,
	name string, args *ResourceIpServiceArgs, opts ...pulumi.ResourceOption) (*ResourceIpService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Numbers == nil {
		return nil, errors.New("invalid value for required argument 'Numbers'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	var resource ResourceIpService
	err := ctx.RegisterResource("routeros:index/resourceIpService:ResourceIpService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceIpService gets an existing ResourceIpService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceIpService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceIpServiceState, opts ...pulumi.ResourceOption) (*ResourceIpService, error) {
	var resource ResourceIpService
	err := ctx.ReadResource("routeros:index/resourceIpService:ResourceIpService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceIpService resources.
type resourceIpServiceState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// List of IP/IPv6 prefixes from which the service is accessible.
	Address *string `pulumi:"address"`
	// The name of the certificate used by a particular service. Applicable only for services that depend on certificates ( www-ssl, api-ssl ).
	Certificate *string `pulumi:"certificate"`
	Disabled    *bool   `pulumi:"disabled"`
	Invalid     *bool   `pulumi:"invalid"`
	// Service name.
	Name *string `pulumi:"name"`
	// The name of the service whose settings will be changed ( api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl ).
	Numbers *string `pulumi:"numbers"`
	// The port particular service listens on.
	Port *int `pulumi:"port"`
	// Specifies which TLS versions to allow by a particular service.
	TlsVersion *string `pulumi:"tlsVersion"`
	// Specify which VRF instance to use by a particular service.
	Vrf *string `pulumi:"vrf"`
}

type ResourceIpServiceState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// List of IP/IPv6 prefixes from which the service is accessible.
	Address pulumi.StringPtrInput
	// The name of the certificate used by a particular service. Applicable only for services that depend on certificates ( www-ssl, api-ssl ).
	Certificate pulumi.StringPtrInput
	Disabled    pulumi.BoolPtrInput
	Invalid     pulumi.BoolPtrInput
	// Service name.
	Name pulumi.StringPtrInput
	// The name of the service whose settings will be changed ( api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl ).
	Numbers pulumi.StringPtrInput
	// The port particular service listens on.
	Port pulumi.IntPtrInput
	// Specifies which TLS versions to allow by a particular service.
	TlsVersion pulumi.StringPtrInput
	// Specify which VRF instance to use by a particular service.
	Vrf pulumi.StringPtrInput
}

func (ResourceIpServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceIpServiceState)(nil)).Elem()
}

type resourceIpServiceArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	// List of IP/IPv6 prefixes from which the service is accessible.
	Address *string `pulumi:"address"`
	// The name of the certificate used by a particular service. Applicable only for services that depend on certificates ( www-ssl, api-ssl ).
	Certificate *string `pulumi:"certificate"`
	Disabled    *bool   `pulumi:"disabled"`
	// The name of the service whose settings will be changed ( api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl ).
	Numbers string `pulumi:"numbers"`
	// The port particular service listens on.
	Port int `pulumi:"port"`
	// Specifies which TLS versions to allow by a particular service.
	TlsVersion *string `pulumi:"tlsVersion"`
	// Specify which VRF instance to use by a particular service.
	Vrf *string `pulumi:"vrf"`
}

// The set of arguments for constructing a ResourceIpService resource.
type ResourceIpServiceArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	// List of IP/IPv6 prefixes from which the service is accessible.
	Address pulumi.StringPtrInput
	// The name of the certificate used by a particular service. Applicable only for services that depend on certificates ( www-ssl, api-ssl ).
	Certificate pulumi.StringPtrInput
	Disabled    pulumi.BoolPtrInput
	// The name of the service whose settings will be changed ( api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl ).
	Numbers pulumi.StringInput
	// The port particular service listens on.
	Port pulumi.IntInput
	// Specifies which TLS versions to allow by a particular service.
	TlsVersion pulumi.StringPtrInput
	// Specify which VRF instance to use by a particular service.
	Vrf pulumi.StringPtrInput
}

func (ResourceIpServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceIpServiceArgs)(nil)).Elem()
}

type ResourceIpServiceInput interface {
	pulumi.Input

	ToResourceIpServiceOutput() ResourceIpServiceOutput
	ToResourceIpServiceOutputWithContext(ctx context.Context) ResourceIpServiceOutput
}

func (*ResourceIpService) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIpService)(nil)).Elem()
}

func (i *ResourceIpService) ToResourceIpServiceOutput() ResourceIpServiceOutput {
	return i.ToResourceIpServiceOutputWithContext(context.Background())
}

func (i *ResourceIpService) ToResourceIpServiceOutputWithContext(ctx context.Context) ResourceIpServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIpServiceOutput)
}

// ResourceIpServiceArrayInput is an input type that accepts ResourceIpServiceArray and ResourceIpServiceArrayOutput values.
// You can construct a concrete instance of `ResourceIpServiceArrayInput` via:
//
//	ResourceIpServiceArray{ ResourceIpServiceArgs{...} }
type ResourceIpServiceArrayInput interface {
	pulumi.Input

	ToResourceIpServiceArrayOutput() ResourceIpServiceArrayOutput
	ToResourceIpServiceArrayOutputWithContext(context.Context) ResourceIpServiceArrayOutput
}

type ResourceIpServiceArray []ResourceIpServiceInput

func (ResourceIpServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceIpService)(nil)).Elem()
}

func (i ResourceIpServiceArray) ToResourceIpServiceArrayOutput() ResourceIpServiceArrayOutput {
	return i.ToResourceIpServiceArrayOutputWithContext(context.Background())
}

func (i ResourceIpServiceArray) ToResourceIpServiceArrayOutputWithContext(ctx context.Context) ResourceIpServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIpServiceArrayOutput)
}

// ResourceIpServiceMapInput is an input type that accepts ResourceIpServiceMap and ResourceIpServiceMapOutput values.
// You can construct a concrete instance of `ResourceIpServiceMapInput` via:
//
//	ResourceIpServiceMap{ "key": ResourceIpServiceArgs{...} }
type ResourceIpServiceMapInput interface {
	pulumi.Input

	ToResourceIpServiceMapOutput() ResourceIpServiceMapOutput
	ToResourceIpServiceMapOutputWithContext(context.Context) ResourceIpServiceMapOutput
}

type ResourceIpServiceMap map[string]ResourceIpServiceInput

func (ResourceIpServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceIpService)(nil)).Elem()
}

func (i ResourceIpServiceMap) ToResourceIpServiceMapOutput() ResourceIpServiceMapOutput {
	return i.ToResourceIpServiceMapOutputWithContext(context.Background())
}

func (i ResourceIpServiceMap) ToResourceIpServiceMapOutputWithContext(ctx context.Context) ResourceIpServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIpServiceMapOutput)
}

type ResourceIpServiceOutput struct{ *pulumi.OutputState }

func (ResourceIpServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIpService)(nil)).Elem()
}

func (o ResourceIpServiceOutput) ToResourceIpServiceOutput() ResourceIpServiceOutput {
	return o
}

func (o ResourceIpServiceOutput) ToResourceIpServiceOutputWithContext(ctx context.Context) ResourceIpServiceOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o ResourceIpServiceOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o ResourceIpServiceOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

// List of IP/IPv6 prefixes from which the service is accessible.
func (o ResourceIpServiceOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

// The name of the certificate used by a particular service. Applicable only for services that depend on certificates ( www-ssl, api-ssl ).
func (o ResourceIpServiceOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o ResourceIpServiceOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

func (o ResourceIpServiceOutput) Invalid() pulumi.BoolOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.BoolOutput { return v.Invalid }).(pulumi.BoolOutput)
}

// Service name.
func (o ResourceIpServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the service whose settings will be changed ( api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl ).
func (o ResourceIpServiceOutput) Numbers() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.StringOutput { return v.Numbers }).(pulumi.StringOutput)
}

// The port particular service listens on.
func (o ResourceIpServiceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Specifies which TLS versions to allow by a particular service.
func (o ResourceIpServiceOutput) TlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.StringPtrOutput { return v.TlsVersion }).(pulumi.StringPtrOutput)
}

// Specify which VRF instance to use by a particular service.
func (o ResourceIpServiceOutput) Vrf() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIpService) pulumi.StringPtrOutput { return v.Vrf }).(pulumi.StringPtrOutput)
}

type ResourceIpServiceArrayOutput struct{ *pulumi.OutputState }

func (ResourceIpServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceIpService)(nil)).Elem()
}

func (o ResourceIpServiceArrayOutput) ToResourceIpServiceArrayOutput() ResourceIpServiceArrayOutput {
	return o
}

func (o ResourceIpServiceArrayOutput) ToResourceIpServiceArrayOutputWithContext(ctx context.Context) ResourceIpServiceArrayOutput {
	return o
}

func (o ResourceIpServiceArrayOutput) Index(i pulumi.IntInput) ResourceIpServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceIpService {
		return vs[0].([]*ResourceIpService)[vs[1].(int)]
	}).(ResourceIpServiceOutput)
}

type ResourceIpServiceMapOutput struct{ *pulumi.OutputState }

func (ResourceIpServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceIpService)(nil)).Elem()
}

func (o ResourceIpServiceMapOutput) ToResourceIpServiceMapOutput() ResourceIpServiceMapOutput {
	return o
}

func (o ResourceIpServiceMapOutput) ToResourceIpServiceMapOutputWithContext(ctx context.Context) ResourceIpServiceMapOutput {
	return o
}

func (o ResourceIpServiceMapOutput) MapIndex(k pulumi.StringInput) ResourceIpServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceIpService {
		return vs[0].(map[string]*ResourceIpService)[vs[1].(string)]
	}).(ResourceIpServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIpServiceInput)(nil)).Elem(), &ResourceIpService{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIpServiceArrayInput)(nil)).Elem(), ResourceIpServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIpServiceMapInput)(nil)).Elem(), ResourceIpServiceMap{})
	pulumi.RegisterOutputType(ResourceIpServiceOutput{})
	pulumi.RegisterOutputType(ResourceIpServiceArrayOutput{})
	pulumi.RegisterOutputType(ResourceIpServiceMapOutput{})
}
