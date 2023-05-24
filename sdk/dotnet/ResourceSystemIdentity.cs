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
    ///     var identity = new Routeros.ResourceSystemIdentity("identity");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import routeros:index/resourceSystemIdentity:ResourceSystemIdentity identity .
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:index/resourceSystemIdentity:ResourceSystemIdentity")]
    public partial class ResourceSystemIdentity : global::Pulumi.CustomResource
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
        /// Device name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceSystemIdentity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceSystemIdentity(string name, ResourceSystemIdentityArgs? args = null, CustomResourceOptions? options = null)
            : base("routeros:index/resourceSystemIdentity:ResourceSystemIdentity", name, args ?? new ResourceSystemIdentityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceSystemIdentity(string name, Input<string> id, ResourceSystemIdentityState? state = null, CustomResourceOptions? options = null)
            : base("routeros:index/resourceSystemIdentity:ResourceSystemIdentity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceSystemIdentity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceSystemIdentity Get(string name, Input<string> id, ResourceSystemIdentityState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceSystemIdentity(name, id, state, options);
        }
    }

    public sealed class ResourceSystemIdentityArgs : global::Pulumi.ResourceArgs
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
        /// Device name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ResourceSystemIdentityArgs()
        {
        }
        public static new ResourceSystemIdentityArgs Empty => new ResourceSystemIdentityArgs();
    }

    public sealed class ResourceSystemIdentityState : global::Pulumi.ResourceArgs
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
        /// Device name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ResourceSystemIdentityState()
        {
        }
        public static new ResourceSystemIdentityState Empty => new ResourceSystemIdentityState();
    }
}
