// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # System.Scheduler (Resource)
//
// ***
//
// #### This is an alias for backwards compatibility between plugin versions.
// Please see documentation for System.SystemScheduler
type Scheduler struct {
	pulumi.CustomResourceState

	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrOutput `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrOutput `pulumi:"___path_"`
	Comment  pulumi.StringPtrOutput `pulumi:"comment"`
	Disabled pulumi.BoolPtrOutput   `pulumi:"disabled"`
	// Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
	// otherwise it is executed repeatedly at the time interval is specified.
	Interval pulumi.StringOutput `pulumi:"interval"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name    pulumi.StringOutput `pulumi:"name"`
	NextRun pulumi.StringOutput `pulumi:"nextRun"`
	// Name of the script to execute. It must be presented at /system script.
	OnEvent pulumi.StringOutput `pulumi:"onEvent"`
	Owner   pulumi.StringOutput `pulumi:"owner"`
	// List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
	// full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
	// used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
	// that grants user management rights. Should be used together with the write policy. Allows also to see global variables
	// created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
	// configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
	// Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
	// Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
	// displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
	// ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
	// access to the router's configuration, except for user management. This policy does not allow to read the configuration,
	// so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// This counter is incremented each time the script is executed.
	RunCount pulumi.StringOutput `pulumi:"runCount"`
	// Date of the first script execution.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
	// start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
	// and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
	// will not run at startup.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewScheduler registers a new resource with the given unique name, arguments, and options.
func NewScheduler(ctx *pulumi.Context,
	name string, args *SchedulerArgs, opts ...pulumi.ResourceOption) (*Scheduler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OnEvent == nil {
		return nil, errors.New("invalid value for required argument 'OnEvent'")
	}
	var resource Scheduler
	err := ctx.RegisterResource("routeros:System/scheduler:Scheduler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduler gets an existing Scheduler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchedulerState, opts ...pulumi.ResourceOption) (*Scheduler, error) {
	var resource Scheduler
	err := ctx.ReadResource("routeros:System/scheduler:Scheduler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Scheduler resources.
type schedulerState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
	// otherwise it is executed repeatedly at the time interval is specified.
	Interval *string `pulumi:"interval"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name    *string `pulumi:"name"`
	NextRun *string `pulumi:"nextRun"`
	// Name of the script to execute. It must be presented at /system script.
	OnEvent *string `pulumi:"onEvent"`
	Owner   *string `pulumi:"owner"`
	// List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
	// full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
	// used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
	// that grants user management rights. Should be used together with the write policy. Allows also to see global variables
	// created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
	// configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
	// Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
	// Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
	// displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
	// ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
	// access to the router's configuration, except for user management. This policy does not allow to read the configuration,
	// so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
	Policies []string `pulumi:"policies"`
	// This counter is incremented each time the script is executed.
	RunCount *string `pulumi:"runCount"`
	// Date of the first script execution.
	StartDate *string `pulumi:"startDate"`
	// Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
	// start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
	// and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
	// will not run at startup.
	StartTime *string `pulumi:"startTime"`
}

type SchedulerState struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
	// otherwise it is executed repeatedly at the time interval is specified.
	Interval pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name    pulumi.StringPtrInput
	NextRun pulumi.StringPtrInput
	// Name of the script to execute. It must be presented at /system script.
	OnEvent pulumi.StringPtrInput
	Owner   pulumi.StringPtrInput
	// List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
	// full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
	// used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
	// that grants user management rights. Should be used together with the write policy. Allows also to see global variables
	// created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
	// configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
	// Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
	// Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
	// displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
	// ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
	// access to the router's configuration, except for user management. This policy does not allow to read the configuration,
	// so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
	Policies pulumi.StringArrayInput
	// This counter is incremented each time the script is executed.
	RunCount pulumi.StringPtrInput
	// Date of the first script execution.
	StartDate pulumi.StringPtrInput
	// Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
	// start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
	// and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
	// will not run at startup.
	StartTime pulumi.StringPtrInput
}

func (SchedulerState) ElementType() reflect.Type {
	return reflect.TypeOf((*schedulerState)(nil)).Elem()
}

type schedulerArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ *int `pulumi:"___id_"`
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ *string `pulumi:"___path_"`
	Comment  *string `pulumi:"comment"`
	Disabled *bool   `pulumi:"disabled"`
	// Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
	// otherwise it is executed repeatedly at the time interval is specified.
	Interval *string `pulumi:"interval"`
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name *string `pulumi:"name"`
	// Name of the script to execute. It must be presented at /system script.
	OnEvent string `pulumi:"onEvent"`
	// List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
	// full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
	// used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
	// that grants user management rights. Should be used together with the write policy. Allows also to see global variables
	// created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
	// configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
	// Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
	// Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
	// displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
	// ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
	// access to the router's configuration, except for user management. This policy does not allow to read the configuration,
	// so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
	Policies []string `pulumi:"policies"`
	// Date of the first script execution.
	StartDate *string `pulumi:"startDate"`
	// Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
	// start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
	// and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
	// will not run at startup.
	StartTime *string `pulumi:"startTime"`
}

// The set of arguments for constructing a Scheduler resource.
type SchedulerArgs struct {
	// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
	___id_ pulumi.IntPtrInput
	// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
	___path_ pulumi.StringPtrInput
	Comment  pulumi.StringPtrInput
	Disabled pulumi.BoolPtrInput
	// Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
	// otherwise it is executed repeatedly at the time interval is specified.
	Interval pulumi.StringPtrInput
	// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
	// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
	// integrity for that resource!
	Name pulumi.StringPtrInput
	// Name of the script to execute. It must be presented at /system script.
	OnEvent pulumi.StringInput
	// List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
	// full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
	// used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
	// that grants user management rights. Should be used together with the write policy. Allows also to see global variables
	// created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
	// configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
	// Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
	// Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
	// displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
	// ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
	// access to the router's configuration, except for user management. This policy does not allow to read the configuration,
	// so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
	Policies pulumi.StringArrayInput
	// Date of the first script execution.
	StartDate pulumi.StringPtrInput
	// Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
	// start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
	// and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
	// will not run at startup.
	StartTime pulumi.StringPtrInput
}

func (SchedulerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schedulerArgs)(nil)).Elem()
}

type SchedulerInput interface {
	pulumi.Input

	ToSchedulerOutput() SchedulerOutput
	ToSchedulerOutputWithContext(ctx context.Context) SchedulerOutput
}

func (*Scheduler) ElementType() reflect.Type {
	return reflect.TypeOf((**Scheduler)(nil)).Elem()
}

func (i *Scheduler) ToSchedulerOutput() SchedulerOutput {
	return i.ToSchedulerOutputWithContext(context.Background())
}

func (i *Scheduler) ToSchedulerOutputWithContext(ctx context.Context) SchedulerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulerOutput)
}

// SchedulerArrayInput is an input type that accepts SchedulerArray and SchedulerArrayOutput values.
// You can construct a concrete instance of `SchedulerArrayInput` via:
//
//	SchedulerArray{ SchedulerArgs{...} }
type SchedulerArrayInput interface {
	pulumi.Input

	ToSchedulerArrayOutput() SchedulerArrayOutput
	ToSchedulerArrayOutputWithContext(context.Context) SchedulerArrayOutput
}

type SchedulerArray []SchedulerInput

func (SchedulerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Scheduler)(nil)).Elem()
}

func (i SchedulerArray) ToSchedulerArrayOutput() SchedulerArrayOutput {
	return i.ToSchedulerArrayOutputWithContext(context.Background())
}

func (i SchedulerArray) ToSchedulerArrayOutputWithContext(ctx context.Context) SchedulerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulerArrayOutput)
}

// SchedulerMapInput is an input type that accepts SchedulerMap and SchedulerMapOutput values.
// You can construct a concrete instance of `SchedulerMapInput` via:
//
//	SchedulerMap{ "key": SchedulerArgs{...} }
type SchedulerMapInput interface {
	pulumi.Input

	ToSchedulerMapOutput() SchedulerMapOutput
	ToSchedulerMapOutputWithContext(context.Context) SchedulerMapOutput
}

type SchedulerMap map[string]SchedulerInput

func (SchedulerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Scheduler)(nil)).Elem()
}

func (i SchedulerMap) ToSchedulerMapOutput() SchedulerMapOutput {
	return i.ToSchedulerMapOutputWithContext(context.Background())
}

func (i SchedulerMap) ToSchedulerMapOutputWithContext(ctx context.Context) SchedulerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulerMapOutput)
}

type SchedulerOutput struct{ *pulumi.OutputState }

func (SchedulerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scheduler)(nil)).Elem()
}

func (o SchedulerOutput) ToSchedulerOutput() SchedulerOutput {
	return o
}

func (o SchedulerOutput) ToSchedulerOutputWithContext(ctx context.Context) SchedulerOutput {
	return o
}

// <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
func (o SchedulerOutput) ___id_() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.IntPtrOutput { return v.___id_ }).(pulumi.IntPtrOutput)
}

// <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
func (o SchedulerOutput) ___path_() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringPtrOutput { return v.___path_ }).(pulumi.StringPtrOutput)
}

func (o SchedulerOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o SchedulerOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Interval between two script executions, if time interval is set to zero, the script is only executed at its start time,
// otherwise it is executed repeatedly at the time interval is specified.
func (o SchedulerOutput) Interval() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.Interval }).(pulumi.StringOutput)
}

// Changing the name of this resource will force it to be recreated. > The links of other configuration properties to this
// resource may be lost! > Changing the name of the resource outside of a Terraform will result in a loss of control
// integrity for that resource!
func (o SchedulerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SchedulerOutput) NextRun() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.NextRun }).(pulumi.StringOutput)
}

// Name of the script to execute. It must be presented at /system script.
func (o SchedulerOutput) OnEvent() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.OnEvent }).(pulumi.StringOutput)
}

func (o SchedulerOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// List of applicable policies: * dude - Policy that grants rights to log in to dude server. * ftp - Policy that grants
// full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be
// used together with read/write policies. * password - Policy that grants rights to change the password. * policy - Policy
// that grants user management rights. Should be used together with the write policy. Allows also to see global variables
// created by other users (requires also 'test' policy). * read - Policy that grants read access to the router's
// configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP. * reboot -
// Policy that allows rebooting the router. * romon - Policy that grants rights to connect to RoMon server. * sensitive -
// Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not
// displayed. * sniff - Policy that grants rights to use packet sniffer tool. * test - Policy that grants rights to run
// ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands. * write - Policy that grants write
// access to the router's configuration, except for user management. This policy does not allow to read the configuration,
// so make sure to enable read policy as well. policy = ["ftp", "read", "write"]
func (o SchedulerOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// This counter is incremented each time the script is executed.
func (o SchedulerOutput) RunCount() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.RunCount }).(pulumi.StringOutput)
}

// Date of the first script execution.
func (o SchedulerOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and
// start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup
// and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler
// will not run at startup.
func (o SchedulerOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheduler) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

type SchedulerArrayOutput struct{ *pulumi.OutputState }

func (SchedulerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Scheduler)(nil)).Elem()
}

func (o SchedulerArrayOutput) ToSchedulerArrayOutput() SchedulerArrayOutput {
	return o
}

func (o SchedulerArrayOutput) ToSchedulerArrayOutputWithContext(ctx context.Context) SchedulerArrayOutput {
	return o
}

func (o SchedulerArrayOutput) Index(i pulumi.IntInput) SchedulerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Scheduler {
		return vs[0].([]*Scheduler)[vs[1].(int)]
	}).(SchedulerOutput)
}

type SchedulerMapOutput struct{ *pulumi.OutputState }

func (SchedulerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Scheduler)(nil)).Elem()
}

func (o SchedulerMapOutput) ToSchedulerMapOutput() SchedulerMapOutput {
	return o
}

func (o SchedulerMapOutput) ToSchedulerMapOutputWithContext(ctx context.Context) SchedulerMapOutput {
	return o
}

func (o SchedulerMapOutput) MapIndex(k pulumi.StringInput) SchedulerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Scheduler {
		return vs[0].(map[string]*Scheduler)[vs[1].(string)]
	}).(SchedulerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulerInput)(nil)).Elem(), &Scheduler{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulerArrayInput)(nil)).Elem(), SchedulerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulerMapInput)(nil)).Elem(), SchedulerMap{})
	pulumi.RegisterOutputType(SchedulerOutput{})
	pulumi.RegisterOutputType(SchedulerArrayOutput{})
	pulumi.RegisterOutputType(SchedulerMapOutput{})
}