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
    ///     var rootCa = new Routeros.System.Certificate("rootCa", new()
    ///     {
    ///         CommonName = "RootCA",
    ///         KeyUsages = new[]
    ///         {
    ///             "key-cert-sign",
    ///             "crl-sign",
    ///         },
    ///         Trusted = true,
    ///         Signs = new[]
    ///         {
    ///             null,
    ///         },
    ///     });
    /// 
    ///     // digitalSignature: Used for entity and data origin authentication with integrity.
    ///     // keyEncipherment:  Used to encrypt symmetric key, which is then transferred to target.
    ///     // keyAgreement:     Enables use of key agreement to establish symmetric key with target. 
    ///     var serverCrt = new Routeros.System.Certificate("serverCrt", new()
    ///     {
    ///         CommonName = "server.crt",
    ///         KeyUsages = new[]
    ///         {
    ///             "digital-signature",
    ///             "key-encipherment",
    ///             "tls-server",
    ///         },
    ///         Signs = new[]
    ///         {
    ///             new Routeros.System.Inputs.CertificateSignArgs
    ///             {
    ///                 Ca = rootCa.Name,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var clientCrt = new Routeros.System.Certificate("clientCrt", new()
    ///     {
    ///         CommonName = "client.crt",
    ///         KeySize = "prime256v1",
    ///         KeyUsages = new[]
    ///         {
    ///             "digital-signature",
    ///             "key-agreement",
    ///             "tls-client",
    ///         },
    ///         Signs = new[]
    ///         {
    ///             new Routeros.System.Inputs.CertificateSignArgs
    ///             {
    ///                 Ca = rootCa.Name,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var unsignedCrt = new Routeros.System.Certificate("unsignedCrt", new()
    ///     {
    ///         CommonName = "unsigned.crt",
    ///         KeySize = "1024",
    ///         SubjectAltName = "DNS:router.lan,DNS:myrouter.lan,IP:192.168.88.1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/certificate get [print show-ids]] #If you plan to manipulate the certificate requiring signing, you need to correctly fill in the sign{} section. #Changes in the sign{} section will not cause changes in the certificate. It's not a bug, it's a feature!
    /// 
    /// ```sh
    ///  $ pulumi import routeros:System/certificate:Certificate client *9D
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:System/certificate:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
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
        /// &lt;em&gt;A set of transformations for field names. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Output("___skip_")]
        public Output<string?> ___skip_ { get; private set; } = null!;

        /// <summary>
        /// Authority Key Identifier.
        /// </summary>
        [Output("akid")]
        public Output<string> Akid { get; private set; } = null!;

        [Output("authority")]
        public Output<string> Authority { get; private set; } = null!;

        [Output("ca")]
        public Output<string> Ca { get; private set; } = null!;

        [Output("caCrlHost")]
        public Output<string> CaCrlHost { get; private set; } = null!;

        [Output("caFingerprint")]
        public Output<string> CaFingerprint { get; private set; } = null!;

        /// <summary>
        /// Common Name (e.g. server FQDN or YOUR name).
        /// </summary>
        [Output("commonName")]
        public Output<string> CommonName { get; private set; } = null!;

        [Output("copyFrom")]
        public Output<string?> CopyFrom { get; private set; } = null!;

        /// <summary>
        /// Country Name (2 letter code).
        /// </summary>
        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        [Output("crl")]
        public Output<string> Crl { get; private set; } = null!;

        /// <summary>
        /// Certificate lifetime.
        /// </summary>
        [Output("daysValid")]
        public Output<int> DaysValid { get; private set; } = null!;

        [Output("digestAlgorithm")]
        public Output<bool> DigestAlgorithm { get; private set; } = null!;

        [Output("dsa")]
        public Output<bool> Dsa { get; private set; } = null!;

        /// <summary>
        /// Set to true if certificate is expired.
        /// </summary>
        [Output("expired")]
        public Output<bool> Expired { get; private set; } = null!;

        [Output("expiresAfter")]
        public Output<string> ExpiresAfter { get; private set; } = null!;

        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The date after which certificate wil be invalid.
        /// </summary>
        [Output("invalidAfter")]
        public Output<string> InvalidAfter { get; private set; } = null!;

        /// <summary>
        /// The date before which certificate is invalid.
        /// </summary>
        [Output("invalidBefore")]
        public Output<string> InvalidBefore { get; private set; } = null!;

        [Output("issued")]
        public Output<string> Issued { get; private set; } = null!;

        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        [Output("keySize")]
        public Output<string> KeySize { get; private set; } = null!;

        [Output("keyType")]
        public Output<string> KeyType { get; private set; } = null!;

        /// <summary>
        /// Detailed key usage descriptions can be found in RFC 5280.
        /// </summary>
        [Output("keyUsages")]
        public Output<ImmutableArray<string>> KeyUsages { get; private set; } = null!;

        /// <summary>
        /// Locality Name (eg, city).
        /// </summary>
        [Output("locality")]
        public Output<string?> Locality { get; private set; } = null!;

        /// <summary>
        /// Name of the certificate. Name can be edited.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Organizational Unit Name (eg, section)
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        [Output("privateKey")]
        public Output<bool> PrivateKey { get; private set; } = null!;

        [Output("reqFingerprint")]
        public Output<string> ReqFingerprint { get; private set; } = null!;

        [Output("revoked")]
        public Output<string> Revoked { get; private set; } = null!;

        [Output("scepUrl")]
        public Output<string> ScepUrl { get; private set; } = null!;

        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        [Output("signs")]
        public Output<ImmutableArray<Outputs.CertificateSign>> Signs { get; private set; } = null!;

        /// <summary>
        /// Subject Key Identifier.
        /// </summary>
        [Output("skid")]
        public Output<string> Skid { get; private set; } = null!;

        [Output("smartCardKey")]
        public Output<string> SmartCardKey { get; private set; } = null!;

        /// <summary>
        /// State or Province Name (full name).
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Shows current status of scep client.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// SANs (subject alternative names).
        /// </summary>
        [Output("subjectAltName")]
        public Output<string?> SubjectAltName { get; private set; } = null!;

        /// <summary>
        /// If set to yes certificate is included 'in trusted certificate chain'.
        /// </summary>
        [Output("trusted")]
        public Output<bool?> Trusted { get; private set; } = null!;

        /// <summary>
        /// Organizational Unit Name (eg, section).
        /// </summary>
        [Output("unit")]
        public Output<string?> Unit { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("routeros:System/certificate:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("routeros:System/certificate:Certificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
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
        /// &lt;em&gt;A set of transformations for field names. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___skip_")]
        public Input<string>? ___skip_ { get; set; }

        /// <summary>
        /// Common Name (e.g. server FQDN or YOUR name).
        /// </summary>
        [Input("commonName", required: true)]
        public Input<string> CommonName { get; set; } = null!;

        [Input("copyFrom")]
        public Input<string>? CopyFrom { get; set; }

        /// <summary>
        /// Country Name (2 letter code).
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Certificate lifetime.
        /// </summary>
        [Input("daysValid")]
        public Input<int>? DaysValid { get; set; }

        [Input("keySize")]
        public Input<string>? KeySize { get; set; }

        [Input("keyUsages")]
        private InputList<string>? _keyUsages;

        /// <summary>
        /// Detailed key usage descriptions can be found in RFC 5280.
        /// </summary>
        public InputList<string> KeyUsages
        {
            get => _keyUsages ?? (_keyUsages = new InputList<string>());
            set => _keyUsages = value;
        }

        /// <summary>
        /// Locality Name (eg, city).
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// Name of the certificate. Name can be edited.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organizational Unit Name (eg, section)
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("signs")]
        private InputList<Inputs.CertificateSignArgs>? _signs;
        public InputList<Inputs.CertificateSignArgs> Signs
        {
            get => _signs ?? (_signs = new InputList<Inputs.CertificateSignArgs>());
            set => _signs = value;
        }

        /// <summary>
        /// State or Province Name (full name).
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// SANs (subject alternative names).
        /// </summary>
        [Input("subjectAltName")]
        public Input<string>? SubjectAltName { get; set; }

        /// <summary>
        /// If set to yes certificate is included 'in trusted certificate chain'.
        /// </summary>
        [Input("trusted")]
        public Input<bool>? Trusted { get; set; }

        /// <summary>
        /// Organizational Unit Name (eg, section).
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }

    public sealed class CertificateState : global::Pulumi.ResourceArgs
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
        /// &lt;em&gt;A set of transformations for field names. This is an internal service field, setting a value is not required.&lt;/em&gt;
        /// </summary>
        [Input("___skip_")]
        public Input<string>? ___skip_ { get; set; }

        /// <summary>
        /// Authority Key Identifier.
        /// </summary>
        [Input("akid")]
        public Input<string>? Akid { get; set; }

        [Input("authority")]
        public Input<string>? Authority { get; set; }

        [Input("ca")]
        public Input<string>? Ca { get; set; }

        [Input("caCrlHost")]
        public Input<string>? CaCrlHost { get; set; }

        [Input("caFingerprint")]
        public Input<string>? CaFingerprint { get; set; }

        /// <summary>
        /// Common Name (e.g. server FQDN or YOUR name).
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        [Input("copyFrom")]
        public Input<string>? CopyFrom { get; set; }

        /// <summary>
        /// Country Name (2 letter code).
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("crl")]
        public Input<string>? Crl { get; set; }

        /// <summary>
        /// Certificate lifetime.
        /// </summary>
        [Input("daysValid")]
        public Input<int>? DaysValid { get; set; }

        [Input("digestAlgorithm")]
        public Input<bool>? DigestAlgorithm { get; set; }

        [Input("dsa")]
        public Input<bool>? Dsa { get; set; }

        /// <summary>
        /// Set to true if certificate is expired.
        /// </summary>
        [Input("expired")]
        public Input<bool>? Expired { get; set; }

        [Input("expiresAfter")]
        public Input<string>? ExpiresAfter { get; set; }

        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The date after which certificate wil be invalid.
        /// </summary>
        [Input("invalidAfter")]
        public Input<string>? InvalidAfter { get; set; }

        /// <summary>
        /// The date before which certificate is invalid.
        /// </summary>
        [Input("invalidBefore")]
        public Input<string>? InvalidBefore { get; set; }

        [Input("issued")]
        public Input<string>? Issued { get; set; }

        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        [Input("keySize")]
        public Input<string>? KeySize { get; set; }

        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        [Input("keyUsages")]
        private InputList<string>? _keyUsages;

        /// <summary>
        /// Detailed key usage descriptions can be found in RFC 5280.
        /// </summary>
        public InputList<string> KeyUsages
        {
            get => _keyUsages ?? (_keyUsages = new InputList<string>());
            set => _keyUsages = value;
        }

        /// <summary>
        /// Locality Name (eg, city).
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// Name of the certificate. Name can be edited.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organizational Unit Name (eg, section)
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("privateKey")]
        public Input<bool>? PrivateKey { get; set; }

        [Input("reqFingerprint")]
        public Input<string>? ReqFingerprint { get; set; }

        [Input("revoked")]
        public Input<string>? Revoked { get; set; }

        [Input("scepUrl")]
        public Input<string>? ScepUrl { get; set; }

        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        [Input("signs")]
        private InputList<Inputs.CertificateSignGetArgs>? _signs;
        public InputList<Inputs.CertificateSignGetArgs> Signs
        {
            get => _signs ?? (_signs = new InputList<Inputs.CertificateSignGetArgs>());
            set => _signs = value;
        }

        /// <summary>
        /// Subject Key Identifier.
        /// </summary>
        [Input("skid")]
        public Input<string>? Skid { get; set; }

        [Input("smartCardKey")]
        public Input<string>? SmartCardKey { get; set; }

        /// <summary>
        /// State or Province Name (full name).
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Shows current status of scep client.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// SANs (subject alternative names).
        /// </summary>
        [Input("subjectAltName")]
        public Input<string>? SubjectAltName { get; set; }

        /// <summary>
        /// If set to yes certificate is included 'in trusted certificate chain'.
        /// </summary>
        [Input("trusted")]
        public Input<bool>? Trusted { get; set; }

        /// <summary>
        /// Organizational Unit Name (eg, section).
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public CertificateState()
        {
        }
        public static new CertificateState Empty => new CertificateState();
    }
}
