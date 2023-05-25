// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CertificateSign struct {
	// Which CA to use if signing issued certificates.
	Ca *string `pulumi:"ca"`
	// CRL host if issuing CA certificate.
	CaCrlHost *string `pulumi:"caCrlHost"`
}

// CertificateSignInput is an input type that accepts CertificateSignArgs and CertificateSignOutput values.
// You can construct a concrete instance of `CertificateSignInput` via:
//
//	CertificateSignArgs{...}
type CertificateSignInput interface {
	pulumi.Input

	ToCertificateSignOutput() CertificateSignOutput
	ToCertificateSignOutputWithContext(context.Context) CertificateSignOutput
}

type CertificateSignArgs struct {
	// Which CA to use if signing issued certificates.
	Ca pulumi.StringPtrInput `pulumi:"ca"`
	// CRL host if issuing CA certificate.
	CaCrlHost pulumi.StringPtrInput `pulumi:"caCrlHost"`
}

func (CertificateSignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSign)(nil)).Elem()
}

func (i CertificateSignArgs) ToCertificateSignOutput() CertificateSignOutput {
	return i.ToCertificateSignOutputWithContext(context.Background())
}

func (i CertificateSignArgs) ToCertificateSignOutputWithContext(ctx context.Context) CertificateSignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSignOutput)
}

// CertificateSignArrayInput is an input type that accepts CertificateSignArray and CertificateSignArrayOutput values.
// You can construct a concrete instance of `CertificateSignArrayInput` via:
//
//	CertificateSignArray{ CertificateSignArgs{...} }
type CertificateSignArrayInput interface {
	pulumi.Input

	ToCertificateSignArrayOutput() CertificateSignArrayOutput
	ToCertificateSignArrayOutputWithContext(context.Context) CertificateSignArrayOutput
}

type CertificateSignArray []CertificateSignInput

func (CertificateSignArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateSign)(nil)).Elem()
}

func (i CertificateSignArray) ToCertificateSignArrayOutput() CertificateSignArrayOutput {
	return i.ToCertificateSignArrayOutputWithContext(context.Background())
}

func (i CertificateSignArray) ToCertificateSignArrayOutputWithContext(ctx context.Context) CertificateSignArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSignArrayOutput)
}

type CertificateSignOutput struct{ *pulumi.OutputState }

func (CertificateSignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSign)(nil)).Elem()
}

func (o CertificateSignOutput) ToCertificateSignOutput() CertificateSignOutput {
	return o
}

func (o CertificateSignOutput) ToCertificateSignOutputWithContext(ctx context.Context) CertificateSignOutput {
	return o
}

// Which CA to use if signing issued certificates.
func (o CertificateSignOutput) Ca() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateSign) *string { return v.Ca }).(pulumi.StringPtrOutput)
}

// CRL host if issuing CA certificate.
func (o CertificateSignOutput) CaCrlHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateSign) *string { return v.CaCrlHost }).(pulumi.StringPtrOutput)
}

type CertificateSignArrayOutput struct{ *pulumi.OutputState }

func (CertificateSignArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateSign)(nil)).Elem()
}

func (o CertificateSignArrayOutput) ToCertificateSignArrayOutput() CertificateSignArrayOutput {
	return o
}

func (o CertificateSignArrayOutput) ToCertificateSignArrayOutputWithContext(ctx context.Context) CertificateSignArrayOutput {
	return o
}

func (o CertificateSignArrayOutput) Index(i pulumi.IntInput) CertificateSignOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateSign {
		return vs[0].([]CertificateSign)[vs[1].(int)]
	}).(CertificateSignOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSignInput)(nil)).Elem(), CertificateSignArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSignArrayInput)(nil)).Elem(), CertificateSignArray{})
	pulumi.RegisterOutputType(CertificateSignOutput{})
	pulumi.RegisterOutputType(CertificateSignArrayOutput{})
}
