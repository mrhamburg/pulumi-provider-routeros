// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.System
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
    ///     var test = new Routeros.System.User("test", new()
    ///     {
    ///         Address = "0.0.0.0/0",
    ///         Comment = "Test User",
    ///         Group = "read",
    ///         Password = "secret",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/user get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:System/user:User test *1
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:System/user:User")]
    public partial class User : global::Pulumi.CustomResource
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
        /// Host or network address from which the user is allowed to log in.
        /// </summary>
        [Output("address")]
        public Output<string?> Address { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Password expired.
        /// </summary>
        [Output("expired")]
        public Output<bool> Expired { get; private set; } = null!;

        /// <summary>
        /// Name of the group the user belongs to.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// Read-only field. Last time and date when a user logged in.
        /// </summary>
        [Output("lastLoggedIn")]
        public Output<string> LastLoggedIn { get; private set; } = null!;

        /// <summary>
        /// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("routeros:System/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("routeros:System/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
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
        /// Host or network address from which the user is allowed to log in.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Name of the group the user belongs to.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
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

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
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
        /// Host or network address from which the user is allowed to log in.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Password expired.
        /// </summary>
        [Input("expired")]
        public Input<bool>? Expired { get; set; }

        /// <summary>
        /// Name of the group the user belongs to.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Read-only field. Last time and date when a user logged in.
        /// </summary>
        [Input("lastLoggedIn")]
        public Input<string>? LastLoggedIn { get; set; }

        /// <summary>
        /// User name. Although it must start with an alphanumeric character, it may contain '*', '_', '.' and '@' symbols.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// User  password. If not specified, it is left blank (hit [Enter] when logging  in). It conforms to standard Unix characteristics of passwords and may  contain letters, digits, '*' and '_' symbols.
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

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}