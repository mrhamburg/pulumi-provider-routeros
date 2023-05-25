// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.Ppp
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
    ///     var test = new Routeros.Ppp.Secret("test", new()
    ///     {
    ///         Password = "123",
    ///         Profile = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/ppp/secret get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:Ppp/secret:Secret test *6
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:Ppp/secret:Secret")]
    public partial class Secret : global::Pulumi.CustomResource
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
        /// For PPTP and L2TP it is the IP address a client must connect from. For PPPoE it is the MAC address (written in CAPITAL letters) a client must  connect from. For ISDN it is the caller's number (that may or may not be  provided by the operator) the client may dial-in from.
        /// </summary>
        [Output("callerId")]
        public Output<string?> CallerId { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// IPv6 routes.
        /// </summary>
        [Output("ipv6Routes")]
        public Output<ImmutableArray<string>> Ipv6Routes { get; private set; } = null!;

        [Output("lastCallerId")]
        public Output<string> LastCallerId { get; private set; } = null!;

        [Output("lastDisconnectReason")]
        public Output<string> LastDisconnectReason { get; private set; } = null!;

        [Output("lastLoggedOut")]
        public Output<string> LastLoggedOut { get; private set; } = null!;

        /// <summary>
        /// Maximal amount of bytes for a session that client can upload.
        /// </summary>
        [Output("limitBytesIn")]
        public Output<int?> LimitBytesIn { get; private set; } = null!;

        /// <summary>
        /// Maximal amount of bytes for a session that client can download.
        /// </summary>
        [Output("limitBytesOut")]
        public Output<int?> LimitBytesOut { get; private set; } = null!;

        /// <summary>
        /// IP address that will be set locally on ppp interface.
        /// </summary>
        [Output("localAddress")]
        public Output<string?> LocalAddress { get; private set; } = null!;

        /// <summary>
        /// Name used for authentication.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password used for authentication.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Which user profile to use.
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;

        /// <summary>
        /// IP address that will be assigned to remote ppp interface.
        /// </summary>
        [Output("remoteAddress")]
        public Output<string?> RemoteAddress { get; private set; } = null!;

        /// <summary>
        /// IPv6 prefix assigned to ppp client. Prefix is added to ND prefix list enabling stateless address auto-configuration on ppp interface.Available starting from v5.0.
        /// </summary>
        [Output("remoteIpv6Prefix")]
        public Output<string?> RemoteIpv6Prefix { get; private set; } = null!;

        /// <summary>
        /// Routes  that appear on the server when the client is connected. The route  format is: dst-address gateway metric (for example, 10.1.0.0/ 24  10.0.0.1 1). Other syntax is not acceptable since it can be represented  in incorrect way. Several routes may be specified separated with commas.  This parameter will be ignored for OpenVPN.
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<string>> Routes { get; private set; } = null!;

        /// <summary>
        /// Specifies the services that particular user will be able to use.
        /// </summary>
        [Output("service")]
        public Output<string?> Service { get; private set; } = null!;


        /// <summary>
        /// Create a Secret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Secret(string name, SecretArgs? args = null, CustomResourceOptions? options = null)
            : base("routeros:Ppp/secret:Secret", name, args ?? new SecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Secret(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
            : base("routeros:Ppp/secret:Secret", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Secret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Secret Get(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
        {
            return new Secret(name, id, state, options);
        }
    }

    public sealed class SecretArgs : global::Pulumi.ResourceArgs
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
        /// For PPTP and L2TP it is the IP address a client must connect from. For PPPoE it is the MAC address (written in CAPITAL letters) a client must  connect from. For ISDN it is the caller's number (that may or may not be  provided by the operator) the client may dial-in from.
        /// </summary>
        [Input("callerId")]
        public Input<string>? CallerId { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("ipv6Routes")]
        private InputList<string>? _ipv6Routes;

        /// <summary>
        /// IPv6 routes.
        /// </summary>
        public InputList<string> Ipv6Routes
        {
            get => _ipv6Routes ?? (_ipv6Routes = new InputList<string>());
            set => _ipv6Routes = value;
        }

        /// <summary>
        /// Maximal amount of bytes for a session that client can upload.
        /// </summary>
        [Input("limitBytesIn")]
        public Input<int>? LimitBytesIn { get; set; }

        /// <summary>
        /// Maximal amount of bytes for a session that client can download.
        /// </summary>
        [Input("limitBytesOut")]
        public Input<int>? LimitBytesOut { get; set; }

        /// <summary>
        /// IP address that will be set locally on ppp interface.
        /// </summary>
        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        /// <summary>
        /// Name used for authentication.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password used for authentication.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Which user profile to use.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// IP address that will be assigned to remote ppp interface.
        /// </summary>
        [Input("remoteAddress")]
        public Input<string>? RemoteAddress { get; set; }

        /// <summary>
        /// IPv6 prefix assigned to ppp client. Prefix is added to ND prefix list enabling stateless address auto-configuration on ppp interface.Available starting from v5.0.
        /// </summary>
        [Input("remoteIpv6Prefix")]
        public Input<string>? RemoteIpv6Prefix { get; set; }

        [Input("routes")]
        private InputList<string>? _routes;

        /// <summary>
        /// Routes  that appear on the server when the client is connected. The route  format is: dst-address gateway metric (for example, 10.1.0.0/ 24  10.0.0.1 1). Other syntax is not acceptable since it can be represented  in incorrect way. Several routes may be specified separated with commas.  This parameter will be ignored for OpenVPN.
        /// </summary>
        public InputList<string> Routes
        {
            get => _routes ?? (_routes = new InputList<string>());
            set => _routes = value;
        }

        /// <summary>
        /// Specifies the services that particular user will be able to use.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public SecretArgs()
        {
        }
        public static new SecretArgs Empty => new SecretArgs();
    }

    public sealed class SecretState : global::Pulumi.ResourceArgs
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
        /// For PPTP and L2TP it is the IP address a client must connect from. For PPPoE it is the MAC address (written in CAPITAL letters) a client must  connect from. For ISDN it is the caller's number (that may or may not be  provided by the operator) the client may dial-in from.
        /// </summary>
        [Input("callerId")]
        public Input<string>? CallerId { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("ipv6Routes")]
        private InputList<string>? _ipv6Routes;

        /// <summary>
        /// IPv6 routes.
        /// </summary>
        public InputList<string> Ipv6Routes
        {
            get => _ipv6Routes ?? (_ipv6Routes = new InputList<string>());
            set => _ipv6Routes = value;
        }

        [Input("lastCallerId")]
        public Input<string>? LastCallerId { get; set; }

        [Input("lastDisconnectReason")]
        public Input<string>? LastDisconnectReason { get; set; }

        [Input("lastLoggedOut")]
        public Input<string>? LastLoggedOut { get; set; }

        /// <summary>
        /// Maximal amount of bytes for a session that client can upload.
        /// </summary>
        [Input("limitBytesIn")]
        public Input<int>? LimitBytesIn { get; set; }

        /// <summary>
        /// Maximal amount of bytes for a session that client can download.
        /// </summary>
        [Input("limitBytesOut")]
        public Input<int>? LimitBytesOut { get; set; }

        /// <summary>
        /// IP address that will be set locally on ppp interface.
        /// </summary>
        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        /// <summary>
        /// Name used for authentication.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password used for authentication.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Which user profile to use.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// IP address that will be assigned to remote ppp interface.
        /// </summary>
        [Input("remoteAddress")]
        public Input<string>? RemoteAddress { get; set; }

        /// <summary>
        /// IPv6 prefix assigned to ppp client. Prefix is added to ND prefix list enabling stateless address auto-configuration on ppp interface.Available starting from v5.0.
        /// </summary>
        [Input("remoteIpv6Prefix")]
        public Input<string>? RemoteIpv6Prefix { get; set; }

        [Input("routes")]
        private InputList<string>? _routes;

        /// <summary>
        /// Routes  that appear on the server when the client is connected. The route  format is: dst-address gateway metric (for example, 10.1.0.0/ 24  10.0.0.1 1). Other syntax is not acceptable since it can be represented  in incorrect way. Several routes may be specified separated with commas.  This parameter will be ignored for OpenVPN.
        /// </summary>
        public InputList<string> Routes
        {
            get => _routes ?? (_routes = new InputList<string>());
            set => _routes = value;
        }

        /// <summary>
        /// Specifies the services that particular user will be able to use.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public SecretState()
        {
        }
        public static new SecretState Empty => new SecretState();
    }
}