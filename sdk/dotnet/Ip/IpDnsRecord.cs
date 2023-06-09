// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Routeros.Ip
{
    /// <summary>
    /// ## # routeros.Ip.IpDnsRecord (Resource)
    /// 
    /// Creates a DNS record on the MikroTik device.
    /// 
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
    ///     var nameRecord = new Routeros.Ip.IpDnsRecord("nameRecord", new()
    ///     {
    ///         Address = "192.168.88.1",
    ///         Type = "A",
    ///     });
    /// 
    ///     var regexpRecord = new Routeros.Ip.IpDnsRecord("regexpRecord", new()
    ///     {
    ///         Address = "192.168.88.1",
    ///         Regexp = ".*pool.ntp.org",
    ///         Type = "A",
    ///     });
    /// 
    ///     var aaaaRecord = new Routeros.Ip.DnsRecord("aaaaRecord", new()
    ///     {
    ///         Address = "ff00::1",
    ///         Type = "AAAA",
    ///     });
    /// 
    ///     var cnameRecord = new Routeros.Ip.DnsRecord("cnameRecord", new()
    ///     {
    ///         Cname = "ipv4.lan",
    ///         Type = "CNAME",
    ///     });
    /// 
    ///     var fwdRecord = new Routeros.Ip.DnsRecord("fwdRecord", new()
    ///     {
    ///         ForwardTo = "127.0.0.1",
    ///         Type = "FWD",
    ///     });
    /// 
    ///     var mxRecord = new Routeros.Ip.DnsRecord("mxRecord", new()
    ///     {
    ///         MxExchange = "127.0.0.1",
    ///         MxPreference = 10,
    ///         Type = "MX",
    ///     });
    /// 
    ///     var nsRecord = new Routeros.Ip.DnsRecord("nsRecord", new()
    ///     {
    ///         Ns = "127.0.0.1",
    ///         Type = "NS",
    ///     });
    /// 
    ///     var nxdomainRecord = new Routeros.Ip.DnsRecord("nxdomainRecord", new()
    ///     {
    ///         Type = "NXDOMAIN",
    ///     });
    /// 
    ///     var srvRecord = new Routeros.Ip.DnsRecord("srvRecord", new()
    ///     {
    ///         SrvPort = 8080,
    ///         SrvPriority = 10,
    ///         SrvTarget = "127.0.0.1",
    ///         SrvWeight = "100",
    ///         Type = "SRV",
    ///     });
    /// 
    ///     var txtRecord = new Routeros.Ip.DnsRecord("txtRecord", new()
    ///     {
    ///         Text = "dW6MrI3nBy3eJgYWH3QAg1Cwk_TvjFESOuKo+mp6nm1",
    ///         Type = "TXT",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #The ID can be found via API or the terminal #The command for the terminal is -&gt; :put [/ip/dns/static get [print show-ids]]
    /// 
    /// ```sh
    ///  $ pulumi import routeros:Ip/ipDnsRecord:IpDnsRecord name_record "*0"
    /// ```
    /// </summary>
    [RouterosResourceType("routeros:Ip/ipDnsRecord:IpDnsRecord")]
    public partial class IpDnsRecord : global::Pulumi.CustomResource
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
        /// The A record to be returend from the DNS hostname.
        /// </summary>
        [Output("address")]
        public Output<string?> Address { get; private set; } = null!;

        /// <summary>
        /// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
        /// </summary>
        [Output("addressList")]
        public Output<string?> AddressList { get; private set; } = null!;

        /// <summary>
        /// Alias name for a domain name.
        /// </summary>
        [Output("cname")]
        public Output<string?> Cname { get; private set; } = null!;

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
        /// The IP address of a domain name server to which a particular DNS request must be forwarded.
        /// </summary>
        [Output("forwardTo")]
        public Output<string?> ForwardTo { get; private set; } = null!;

        /// <summary>
        /// Whether the record will match requests for subdomains.
        /// </summary>
        [Output("matchSubdomain")]
        public Output<bool?> MatchSubdomain { get; private set; } = null!;

        /// <summary>
        /// The domain name of the MX server.
        /// </summary>
        [Output("mxExchange")]
        public Output<string?> MxExchange { get; private set; } = null!;

        /// <summary>
        /// Preference of the particular MX record.
        /// </summary>
        [Output("mxPreference")]
        public Output<int> MxPreference { get; private set; } = null!;

        /// <summary>
        /// The name of the DNS hostname to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of the authoritative domain name server for the particular record.
        /// </summary>
        [Output("ns")]
        public Output<string?> Ns { get; private set; } = null!;

        /// <summary>
        /// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS names to lowercase, you should write regex only with lowercase letters.
        /// </summary>
        [Output("regexp")]
        public Output<string?> Regexp { get; private set; } = null!;

        /// <summary>
        /// The TCP or UDP port on which the service is to be found.
        /// </summary>
        [Output("srvPort")]
        public Output<int> SrvPort { get; private set; } = null!;

        /// <summary>
        /// Priority of the particular SRV record.
        /// </summary>
        [Output("srvPriority")]
        public Output<int> SrvPriority { get; private set; } = null!;

        /// <summary>
        /// The canonical hostname of the machine providing the service ends in a dot.
        /// </summary>
        [Output("srvTarget")]
        public Output<string?> SrvTarget { get; private set; } = null!;

        /// <summary>
        /// Weight of the particular SRC record.
        /// </summary>
        [Output("srvWeight")]
        public Output<string> SrvWeight { get; private set; } = null!;

        /// <summary>
        /// Textual information about the domain name.
        /// </summary>
        [Output("text")]
        public Output<string?> Text { get; private set; } = null!;

        /// <summary>
        /// The ttl of the DNS record.
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IpDnsRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpDnsRecord(string name, IpDnsRecordArgs args, CustomResourceOptions? options = null)
            : base("routeros:Ip/ipDnsRecord:IpDnsRecord", name, args ?? new IpDnsRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpDnsRecord(string name, Input<string> id, IpDnsRecordState? state = null, CustomResourceOptions? options = null)
            : base("routeros:Ip/ipDnsRecord:IpDnsRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpDnsRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpDnsRecord Get(string name, Input<string> id, IpDnsRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new IpDnsRecord(name, id, state, options);
        }
    }

    public sealed class IpDnsRecordArgs : global::Pulumi.ResourceArgs
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
        /// The A record to be returend from the DNS hostname.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
        /// </summary>
        [Input("addressList")]
        public Input<string>? AddressList { get; set; }

        /// <summary>
        /// Alias name for a domain name.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The IP address of a domain name server to which a particular DNS request must be forwarded.
        /// </summary>
        [Input("forwardTo")]
        public Input<string>? ForwardTo { get; set; }

        /// <summary>
        /// Whether the record will match requests for subdomains.
        /// </summary>
        [Input("matchSubdomain")]
        public Input<bool>? MatchSubdomain { get; set; }

        /// <summary>
        /// The domain name of the MX server.
        /// </summary>
        [Input("mxExchange")]
        public Input<string>? MxExchange { get; set; }

        /// <summary>
        /// Preference of the particular MX record.
        /// </summary>
        [Input("mxPreference")]
        public Input<int>? MxPreference { get; set; }

        /// <summary>
        /// The name of the DNS hostname to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the authoritative domain name server for the particular record.
        /// </summary>
        [Input("ns")]
        public Input<string>? Ns { get; set; }

        /// <summary>
        /// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS names to lowercase, you should write regex only with lowercase letters.
        /// </summary>
        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        /// <summary>
        /// The TCP or UDP port on which the service is to be found.
        /// </summary>
        [Input("srvPort")]
        public Input<int>? SrvPort { get; set; }

        /// <summary>
        /// Priority of the particular SRV record.
        /// </summary>
        [Input("srvPriority")]
        public Input<int>? SrvPriority { get; set; }

        /// <summary>
        /// The canonical hostname of the machine providing the service ends in a dot.
        /// </summary>
        [Input("srvTarget")]
        public Input<string>? SrvTarget { get; set; }

        /// <summary>
        /// Weight of the particular SRC record.
        /// </summary>
        [Input("srvWeight")]
        public Input<string>? SrvWeight { get; set; }

        /// <summary>
        /// Textual information about the domain name.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        /// <summary>
        /// The ttl of the DNS record.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IpDnsRecordArgs()
        {
        }
        public static new IpDnsRecordArgs Empty => new IpDnsRecordArgs();
    }

    public sealed class IpDnsRecordState : global::Pulumi.ResourceArgs
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
        /// The A record to be returend from the DNS hostname.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
        /// </summary>
        [Input("addressList")]
        public Input<string>? AddressList { get; set; }

        /// <summary>
        /// Alias name for a domain name.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

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
        /// The IP address of a domain name server to which a particular DNS request must be forwarded.
        /// </summary>
        [Input("forwardTo")]
        public Input<string>? ForwardTo { get; set; }

        /// <summary>
        /// Whether the record will match requests for subdomains.
        /// </summary>
        [Input("matchSubdomain")]
        public Input<bool>? MatchSubdomain { get; set; }

        /// <summary>
        /// The domain name of the MX server.
        /// </summary>
        [Input("mxExchange")]
        public Input<string>? MxExchange { get; set; }

        /// <summary>
        /// Preference of the particular MX record.
        /// </summary>
        [Input("mxPreference")]
        public Input<int>? MxPreference { get; set; }

        /// <summary>
        /// The name of the DNS hostname to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the authoritative domain name server for the particular record.
        /// </summary>
        [Input("ns")]
        public Input<string>? Ns { get; set; }

        /// <summary>
        /// DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS names to lowercase, you should write regex only with lowercase letters.
        /// </summary>
        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        /// <summary>
        /// The TCP or UDP port on which the service is to be found.
        /// </summary>
        [Input("srvPort")]
        public Input<int>? SrvPort { get; set; }

        /// <summary>
        /// Priority of the particular SRV record.
        /// </summary>
        [Input("srvPriority")]
        public Input<int>? SrvPriority { get; set; }

        /// <summary>
        /// The canonical hostname of the machine providing the service ends in a dot.
        /// </summary>
        [Input("srvTarget")]
        public Input<string>? SrvTarget { get; set; }

        /// <summary>
        /// Weight of the particular SRC record.
        /// </summary>
        [Input("srvWeight")]
        public Input<string>? SrvWeight { get; set; }

        /// <summary>
        /// Textual information about the domain name.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        /// <summary>
        /// The ttl of the DNS record.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IpDnsRecordState()
        {
        }
        public static new IpDnsRecordState Empty => new IpDnsRecordState();
    }
}
