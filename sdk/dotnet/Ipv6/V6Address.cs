// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.Ipv6
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Routeros = Pulumi.Routeros;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ipv6Address = new Routeros.Ipv6.V6Address("ipv6Address", new()
    ///     {
    ///         Address = "fd55::1/64",
    ///         Disabled = false,
    ///         Interface = "ether1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/ipv6/address get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:Ipv6/v6Address:V6Address ipv6_address "*0"
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:Ipv6/v6Address:V6Address")]
    public partial class V6Address : global::Pulumi.CustomResource
    {
        /// <summary>
        /// &lt;em&gt;Resource ID type (.id / name). This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Output("___id_")]
        public Output<int?> ___id_ { get; private set; } = null!;

        /// <summary>
        /// &lt;em&gt;Resource path for CRUD operations. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Output("___path_")]
        public Output<string?> ___path_ { get; private set; } = null!;

        /// <summary>
        /// Name of the actual interface the logical one is bound to.
        /// </summary>
        [Output("actualInterface")]
        public Output<string> ActualInterface { get; private set; } = null!;

        /// <summary>
        /// IPv6 address. Using the eui*64 and from*pool options can transform the original address! [See docs](https://wiki.mikrotik.com/wiki/Manual:IPv6/Address#Properties)
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Whether to enable stateless address configuration. The prefix of that address is automatically advertised to hosts using ICMPv6 protocol. The option is set by default for addresses with prefix length 64.
        /// </summary>
        [Output("advertise")]
        public Output<bool?> Advertise { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
        /// </summary>
        [Output("dynamic")]
        public Output<bool> Dynamic { get; private set; } = null!;

        /// <summary>
        /// Whether to calculate EUI-64 address and use it as last 64 bits of the IPv6 address.
        /// </summary>
        [Output("eui64")]
        public Output<bool?> Eui64 { get; private set; } = null!;

        /// <summary>
        /// Name of the pool from which prefix will be taken to construct IPv6 address taking last part of the address from address property.
        /// </summary>
        [Output("fromPool")]
        public Output<string?> FromPool { get; private set; } = null!;

        /// <summary>
        /// Whether address is global.
        /// </summary>
        [Output("global")]
        public Output<bool> Global { get; private set; } = null!;

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        [Output("invalid")]
        public Output<bool> Invalid { get; private set; } = null!;

        /// <summary>
        /// Whether address is link local.
        /// </summary>
        [Output("linkLocal")]
        public Output<bool> LinkLocal { get; private set; } = null!;

        /// <summary>
        /// If set indicates that address is anycast address and Duplicate Address Detection should not be performed.
        /// </summary>
        [Output("noDad")]
        public Output<bool?> NoDad { get; private set; } = null!;


        /// <summary>
        /// Create a V6Address resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public V6Address(string name, V6AddressArgs args, CustomResourceOptions? options = null)
            : base("routeros:Ipv6/v6Address:V6Address", name, args ?? new V6AddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private V6Address(string name, Input<string> id, V6AddressState? state = null, CustomResourceOptions? options = null)
            : base("routeros:Ipv6/v6Address:V6Address", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing V6Address resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static V6Address Get(string name, Input<string> id, V6AddressState? state = null, CustomResourceOptions? options = null)
        {
            return new V6Address(name, id, state, options);
        }
    }

    public sealed class V6AddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;em&gt;Resource ID type (.id / name). This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___id_")]
        public Input<int>? ___id_ { get; set; }

        /// <summary>
        /// &lt;em&gt;Resource path for CRUD operations. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___path_")]
        public Input<string>? ___path_ { get; set; }

        /// <summary>
        /// IPv6 address. Using the eui*64 and from*pool options can transform the original address! [See docs](https://wiki.mikrotik.com/wiki/Manual:IPv6/Address#Properties)
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Whether to enable stateless address configuration. The prefix of that address is automatically advertised to hosts using ICMPv6 protocol. The option is set by default for addresses with prefix length 64.
        /// </summary>
        [Input("advertise")]
        public Input<bool>? Advertise { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Whether to calculate EUI-64 address and use it as last 64 bits of the IPv6 address.
        /// </summary>
        [Input("eui64")]
        public Input<bool>? Eui64 { get; set; }

        /// <summary>
        /// Name of the pool from which prefix will be taken to construct IPv6 address taking last part of the address from address property.
        /// </summary>
        [Input("fromPool")]
        public Input<string>? FromPool { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// If set indicates that address is anycast address and Duplicate Address Detection should not be performed.
        /// </summary>
        [Input("noDad")]
        public Input<bool>? NoDad { get; set; }

        public V6AddressArgs()
        {
        }
        public static new V6AddressArgs Empty => new V6AddressArgs();
    }

    public sealed class V6AddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;em&gt;Resource ID type (.id / name). This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___id_")]
        public Input<int>? ___id_ { get; set; }

        /// <summary>
        /// &lt;em&gt;Resource path for CRUD operations. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___path_")]
        public Input<string>? ___path_ { get; set; }

        /// <summary>
        /// Name of the actual interface the logical one is bound to.
        /// </summary>
        [Input("actualInterface")]
        public Input<string>? ActualInterface { get; set; }

        /// <summary>
        /// IPv6 address. Using the eui*64 and from*pool options can transform the original address! [See docs](https://wiki.mikrotik.com/wiki/Manual:IPv6/Address#Properties)
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Whether to enable stateless address configuration. The prefix of that address is automatically advertised to hosts using ICMPv6 protocol. The option is set by default for addresses with prefix length 64.
        /// </summary>
        [Input("advertise")]
        public Input<bool>? Advertise { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
        /// </summary>
        [Input("dynamic")]
        public Input<bool>? Dynamic { get; set; }

        /// <summary>
        /// Whether to calculate EUI-64 address and use it as last 64 bits of the IPv6 address.
        /// </summary>
        [Input("eui64")]
        public Input<bool>? Eui64 { get; set; }

        /// <summary>
        /// Name of the pool from which prefix will be taken to construct IPv6 address taking last part of the address from address property.
        /// </summary>
        [Input("fromPool")]
        public Input<string>? FromPool { get; set; }

        /// <summary>
        /// Whether address is global.
        /// </summary>
        [Input("global")]
        public Input<bool>? Global { get; set; }

        /// <summary>
        /// Name of the interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        [Input("invalid")]
        public Input<bool>? Invalid { get; set; }

        /// <summary>
        /// Whether address is link local.
        /// </summary>
        [Input("linkLocal")]
        public Input<bool>? LinkLocal { get; set; }

        /// <summary>
        /// If set indicates that address is anycast address and Duplicate Address Detection should not be performed.
        /// </summary>
        [Input("noDad")]
        public Input<bool>? NoDad { get; set; }

        public V6AddressState()
        {
        }
        public static new V6AddressState Empty => new V6AddressState();
    }
}