// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros
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
    ///     var user1 = new Routeros.ResourceInterfaceOpenVPNServer("user1", new()
    ///     {
    ///         User = "user1",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             routeros_ovpn_server.Server,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/interface/ovpn-server get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:index/resourceInterfaceOpenVPNServer:ResourceInterfaceOpenVPNServer user1 *29
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:index/resourceInterfaceOpenVPNServer:ResourceInterfaceOpenVPNServer")]
    public partial class ResourceInterfaceOpenVPNServer : global::Pulumi.CustomResource
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
        /// The address of the remote side.
        /// </summary>
        [Output("clientAddress")]
        public Output<string> ClientAddress { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Encryption characteristics.
        /// </summary>
        [Output("encoding")]
        public Output<string> Encoding { get; private set; } = null!;

        /// <summary>
        /// Layer2 Maximum transmission unit.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// Interface name (Example: ovpn-in1).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("running")]
        public Output<bool> Running { get; private set; } = null!;

        /// <summary>
        /// Connection uptime.
        /// </summary>
        [Output("uptime")]
        public Output<string> Uptime { get; private set; } = null!;

        /// <summary>
        /// User name used for authentication.
        /// </summary>
        [Output("user")]
        public Output<string?> User { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceInterfaceOpenVPNServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceInterfaceOpenVPNServer(string name, ResourceInterfaceOpenVPNServerArgs? args = null, CustomResourceOptions? options = null)
            : base("routeros:index/resourceInterfaceOpenVPNServer:ResourceInterfaceOpenVPNServer", name, args ?? new ResourceInterfaceOpenVPNServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceInterfaceOpenVPNServer(string name, Input<string> id, ResourceInterfaceOpenVPNServerState? state = null, CustomResourceOptions? options = null)
            : base("routeros:index/resourceInterfaceOpenVPNServer:ResourceInterfaceOpenVPNServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceInterfaceOpenVPNServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceInterfaceOpenVPNServer Get(string name, Input<string> id, ResourceInterfaceOpenVPNServerState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceInterfaceOpenVPNServer(name, id, state, options);
        }
    }

    public sealed class ResourceInterfaceOpenVPNServerArgs : global::Pulumi.ResourceArgs
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

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Interface name (Example: ovpn-in1).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User name used for authentication.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public ResourceInterfaceOpenVPNServerArgs()
        {
        }
        public static new ResourceInterfaceOpenVPNServerArgs Empty => new ResourceInterfaceOpenVPNServerArgs();
    }

    public sealed class ResourceInterfaceOpenVPNServerState : global::Pulumi.ResourceArgs
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
        /// The address of the remote side.
        /// </summary>
        [Input("clientAddress")]
        public Input<string>? ClientAddress { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Encryption characteristics.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Layer2 Maximum transmission unit.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Interface name (Example: ovpn-in1).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("running")]
        public Input<bool>? Running { get; set; }

        /// <summary>
        /// Connection uptime.
        /// </summary>
        [Input("uptime")]
        public Input<string>? Uptime { get; set; }

        /// <summary>
        /// User name used for authentication.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public ResourceInterfaceOpenVPNServerState()
        {
        }
        public static new ResourceInterfaceOpenVPNServerState Empty => new ResourceInterfaceOpenVPNServerState();
    }
}
