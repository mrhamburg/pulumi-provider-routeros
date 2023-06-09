// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.CapsMan
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
    ///     var testConfiguration = new Routeros.CapsMan.Configuration("testConfiguration");
    /// 
    ///     var testProvisioning = new Routeros.CapsMan.Provisioning("testProvisioning", new()
    ///     {
    ///         MasterConfiguration = "cfg1",
    ///         Action = "create-disabled",
    ///         NamePrefix = "cap-",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             testConfiguration,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt;
    /// 
    /// :put [/caps-man/provisioning get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:CapsMan/provisioning:Provisioning test_provisioning "*B"
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:CapsMan/provisioning:Provisioning")]
    public partial class Provisioning : global::Pulumi.CustomResource
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
        /// Provisioning action.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Regular expression to match radios by common name. Each CAP's common name identifier can be found under "/caps-man radio" as value "REMOTE-CAP-NAME"
        /// </summary>
        [Output("commonNameRegexp")]
        public Output<string?> CommonNameRegexp { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Match radios by supported wireless modes.
        /// </summary>
        [Output("hwSupportedModes")]
        public Output<string?> HwSupportedModes { get; private set; } = null!;

        /// <summary>
        /// Regular expression to match radios by router identity.
        /// </summary>
        [Output("identityRegexp")]
        public Output<string?> IdentityRegexp { get; private set; } = null!;

        /// <summary>
        /// Match CAPs with IPs within configured address range.
        /// </summary>
        [Output("ipAddressRanges")]
        public Output<string?> IpAddressRanges { get; private set; } = null!;

        /// <summary>
        /// If action specifies to create interfaces, then a new master interface with its configuration set to this configuration profile will be created
        /// </summary>
        [Output("masterConfiguration")]
        public Output<string> MasterConfiguration { get; private set; } = null!;

        /// <summary>
        /// Specify the syntax of the CAP interface name creation.
        /// </summary>
        [Output("nameFormat")]
        public Output<string?> NameFormat { get; private set; } = null!;

        /// <summary>
        /// Name prefix which can be used in the name-format for creating the CAP interface names.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// MAC address of radio to be matched, empty MAC (00:00:00:00:00:00) means match all MAC addresses.
        /// </summary>
        [Output("radioMac")]
        public Output<string?> RadioMac { get; private set; } = null!;

        /// <summary>
        /// If action specifies to create interfaces, then a new slave interface for each configuration profile in this list is created.
        /// </summary>
        [Output("slaveConfigurations")]
        public Output<string?> SlaveConfigurations { get; private set; } = null!;


        /// <summary>
        /// Create a Provisioning resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provisioning(string name, ProvisioningArgs args, CustomResourceOptions? options = null)
            : base("routeros:CapsMan/provisioning:Provisioning", name, args ?? new ProvisioningArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Provisioning(string name, Input<string> id, ProvisioningState? state = null, CustomResourceOptions? options = null)
            : base("routeros:CapsMan/provisioning:Provisioning", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Provisioning resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Provisioning Get(string name, Input<string> id, ProvisioningState? state = null, CustomResourceOptions? options = null)
        {
            return new Provisioning(name, id, state, options);
        }
    }

    public sealed class ProvisioningArgs : global::Pulumi.ResourceArgs
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
        /// Provisioning action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Regular expression to match radios by common name. Each CAP's common name identifier can be found under "/caps-man radio" as value "REMOTE-CAP-NAME"
        /// </summary>
        [Input("commonNameRegexp")]
        public Input<string>? CommonNameRegexp { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Match radios by supported wireless modes.
        /// </summary>
        [Input("hwSupportedModes")]
        public Input<string>? HwSupportedModes { get; set; }

        /// <summary>
        /// Regular expression to match radios by router identity.
        /// </summary>
        [Input("identityRegexp")]
        public Input<string>? IdentityRegexp { get; set; }

        /// <summary>
        /// Match CAPs with IPs within configured address range.
        /// </summary>
        [Input("ipAddressRanges")]
        public Input<string>? IpAddressRanges { get; set; }

        /// <summary>
        /// If action specifies to create interfaces, then a new master interface with its configuration set to this configuration profile will be created
        /// </summary>
        [Input("masterConfiguration", required: true)]
        public Input<string> MasterConfiguration { get; set; } = null!;

        /// <summary>
        /// Specify the syntax of the CAP interface name creation.
        /// </summary>
        [Input("nameFormat")]
        public Input<string>? NameFormat { get; set; }

        /// <summary>
        /// Name prefix which can be used in the name-format for creating the CAP interface names.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// MAC address of radio to be matched, empty MAC (00:00:00:00:00:00) means match all MAC addresses.
        /// </summary>
        [Input("radioMac")]
        public Input<string>? RadioMac { get; set; }

        /// <summary>
        /// If action specifies to create interfaces, then a new slave interface for each configuration profile in this list is created.
        /// </summary>
        [Input("slaveConfigurations")]
        public Input<string>? SlaveConfigurations { get; set; }

        public ProvisioningArgs()
        {
        }
        public static new ProvisioningArgs Empty => new ProvisioningArgs();
    }

    public sealed class ProvisioningState : global::Pulumi.ResourceArgs
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
        /// Provisioning action.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Regular expression to match radios by common name. Each CAP's common name identifier can be found under "/caps-man radio" as value "REMOTE-CAP-NAME"
        /// </summary>
        [Input("commonNameRegexp")]
        public Input<string>? CommonNameRegexp { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Match radios by supported wireless modes.
        /// </summary>
        [Input("hwSupportedModes")]
        public Input<string>? HwSupportedModes { get; set; }

        /// <summary>
        /// Regular expression to match radios by router identity.
        /// </summary>
        [Input("identityRegexp")]
        public Input<string>? IdentityRegexp { get; set; }

        /// <summary>
        /// Match CAPs with IPs within configured address range.
        /// </summary>
        [Input("ipAddressRanges")]
        public Input<string>? IpAddressRanges { get; set; }

        /// <summary>
        /// If action specifies to create interfaces, then a new master interface with its configuration set to this configuration profile will be created
        /// </summary>
        [Input("masterConfiguration")]
        public Input<string>? MasterConfiguration { get; set; }

        /// <summary>
        /// Specify the syntax of the CAP interface name creation.
        /// </summary>
        [Input("nameFormat")]
        public Input<string>? NameFormat { get; set; }

        /// <summary>
        /// Name prefix which can be used in the name-format for creating the CAP interface names.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// MAC address of radio to be matched, empty MAC (00:00:00:00:00:00) means match all MAC addresses.
        /// </summary>
        [Input("radioMac")]
        public Input<string>? RadioMac { get; set; }

        /// <summary>
        /// If action specifies to create interfaces, then a new slave interface for each configuration profile in this list is created.
        /// </summary>
        [Input("slaveConfigurations")]
        public Input<string>? SlaveConfigurations { get; set; }

        public ProvisioningState()
        {
        }
        public static new ProvisioningState Empty => new ProvisioningState();
    }
}
