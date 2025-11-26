Addressing
Addressing
Addressing
Address Maps
Addressing.AddressMaps
Methods
client.Addressing.AddressMaps.List(ctx, query) (*SinglePage[
AddressMap
], error)
get/accounts/{account_id}/addressing/address_maps
List all address maps owned by the account.
client.Addressing.AddressMaps.Get(ctx, addressMapID, query) (*
AddressMapGetResponse
, error)
get/accounts/{account_id}/addressing/address_maps/{address_map_id}
Show a particular address map owned by the account.
client.Addressing.AddressMaps.New(ctx, params) (*
AddressMapNewResponse
, error)
post/accounts/{account_id}/addressing/address_maps
Create a new address map under the account.
client.Addressing.AddressMaps.Edit(ctx, addressMapID, params) (*
AddressMap
, error)
patch/accounts/{account_id}/addressing/address_maps/{address_map_id}
Modify properties of an address map owned by the account.
client.Addressing.AddressMaps.Delete(ctx, addressMapID, body) (*
AddressMapDeleteResponse
, error)
delete/accounts/{account_id}/addressing/address_maps/{address_map_id}
Delete a particular address map owned by the account. An Address Map must be disabled before it can be deleted.
Domain types
typeAddressMapstruct{…}
typeKindstring
The type of the membership.
Addressing
Address Maps
Accounts
Addressing.AddressMaps.Accounts
Methods
client.Addressing.AddressMaps.Accounts.Update(ctx, addressMapID, params) (*
AddressMapAccountUpdateResponse
, error)
put/accounts/{account_id}/addressing/address_maps/{address_map_id}/accounts/{account_id}
Add an account as a member of a particular address map.
client.Addressing.AddressMaps.Accounts.Delete(ctx, addressMapID, body) (*
AddressMapAccountDeleteResponse
, error)
delete/accounts/{account_id}/addressing/address_maps/{address_map_id}/accounts/{account_id}
Remove an account as a member of a particular address map.
Addressing
Address Maps
IPs
Addressing.AddressMaps.IPs
Methods
client.Addressing.AddressMaps.IPs.Update(ctx, addressMapID, ipAddress, params) (*
AddressMapIPUpdateResponse
, error)
put/accounts/{account_id}/addressing/address_maps/{address_map_id}/ips/{ip_address}
Add an IP from a prefix owned by the account to a particular address map.
client.Addressing.AddressMaps.IPs.Delete(ctx, addressMapID, ipAddress, body) (*
AddressMapIPDeleteResponse
, error)
delete/accounts/{account_id}/addressing/address_maps/{address_map_id}/ips/{ip_address}
Remove an IP from a particular address map.
Addressing
Address Maps
Zones
Addressing.AddressMaps.Zones
Methods
client.Addressing.AddressMaps.Zones.Update(ctx, addressMapID, params) (*
AddressMapZoneUpdateResponse
, error)
put/accounts/{account_id}/addressing/address_maps/{address_map_id}/zones/{zone_id}
Add a zone as a member of a particular address map.
client.Addressing.AddressMaps.Zones.Delete(ctx, addressMapID, body) (*
AddressMapZoneDeleteResponse
, error)
delete/accounts/{account_id}/addressing/address_maps/{address_map_id}/zones/{zone_id}
Remove a zone as a member of a particular address map.
Addressing
LOA Documents
Addressing.LOADocuments
Methods
client.Addressing.LOADocuments.Get(ctx, loaDocumentID, query) (*
Response
, error)
get/accounts/{account_id}/addressing/loa_documents/{loa_document_id}/download
Download specified LOA document under the account.
Addressing
Prefixes
Addressing.Prefixes
Methods
client.Addressing.Prefixes.List(ctx, query) (*SinglePage[
Prefix
], error)
get/accounts/{account_id}/addressing/prefixes
List all prefixes owned by the account.
client.Addressing.Prefixes.Get(ctx, prefixID, query) (*
Prefix
, error)
get/accounts/{account_id}/addressing/prefixes/{prefix_id}
List a particular prefix owned by the account.
client.Addressing.Prefixes.New(ctx, params) (*
Prefix
, error)
post/accounts/{account_id}/addressing/prefixes
Add a new prefix under the account.
client.Addressing.Prefixes.Edit(ctx, prefixID, params) (*
Prefix
, error)
patch/accounts/{account_id}/addressing/prefixes/{prefix_id}
Modify the description for a prefix owned by the account.
client.Addressing.Prefixes.Delete(ctx, prefixID, body) (*
PrefixDeleteResponse
, error)
delete/accounts/{account_id}/addressing/prefixes/{prefix_id}
Delete an unapproved prefix owned by the account.
Domain types
typePrefixstruct{…}
Addressing
Prefixes
Advertisement Status
Addressing.Prefixes.AdvertisementStatus
Methods
client.Addressing.Prefixes.AdvertisementStatus.Get(ctx, prefixID, query) (*
PrefixAdvertisementStatusGetResponse
, error)
Deprecated
get/accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/status
View the current advertisement state for a prefix.
Deprecated: Prefer the BGP Prefixes endpoints, which additionally allow for advertising and withdrawing subnets of an IP prefix.
client.Addressing.Prefixes.AdvertisementStatus.Edit(ctx, prefixID, params) (*
PrefixAdvertisementStatusEditResponse
, error)
Deprecated
patch/accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/status
Advertise or withdraw the BGP route for a prefix.
Deprecated: Prefer the BGP Prefixes endpoints, which additionally allow for advertising and withdrawing subnets of an IP prefix.
Addressing
Prefixes
BGP Prefixes
Addressing.Prefixes.BGPPrefixes
Methods
client.Addressing.Prefixes.BGPPrefixes.List(ctx, prefixID, query) (*SinglePage[
BGPPrefix
], error)
get/accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes
List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to control which specific subnets are advertised to the Internet. It is possible to advertise subnets more specific than an IP Prefix by creating more specific BGP Prefixes.
client.Addressing.Prefixes.BGPPrefixes.Get(ctx, prefixID, bgpPrefixID, query) (*
BGPPrefix
, error)
get/accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}
Retrieve a single BGP Prefix according to its identifier
client.Addressing.Prefixes.BGPPrefixes.New(ctx, prefixID, params) (*
BGPPrefix
, error)
post/accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes
Create a BGP prefix, controlling the BGP advertisement status of a specific subnet. When created, BGP prefixes are initially withdrawn, and can be advertised with the Update BGP Prefix API.
client.Addressing.Prefixes.BGPPrefixes.Edit(ctx, prefixID, bgpPrefixID, params) (*
BGPPrefix
, error)
patch/accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}
Update the properties of a BGP Prefix, such as the on demand advertisement status (advertised or withdrawn).
Domain types
typeBGPPrefixstruct{…}
Addressing
Prefixes
Delegations
Addressing.Prefixes.Delegations
Methods
client.Addressing.Prefixes.Delegations.List(ctx, prefixID, query) (*SinglePage[
Delegations
], error)
get/accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations
List all delegations for a given account IP prefix.
client.Addressing.Prefixes.Delegations.New(ctx, prefixID, params) (*
Delegations
, error)
post/accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations
Create a new account delegation for a given IP prefix.
client.Addressing.Prefixes.Delegations.Delete(ctx, prefixID, delegationID, body) (*
PrefixDelegationDeleteResponse
, error)
delete/accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations/{delegation_id}
Delete an account delegation for a given IP prefix.
Domain types
typeDelegationsstruct{…}
Addressing
Prefixes
Service Bindings
Addressing.Prefixes.ServiceBindings
Methods
client.Addressing.Prefixes.ServiceBindings.List(ctx, prefixID, query) (*SinglePage[
ServiceBinding
], error)
get/accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings
List the Cloudflare services this prefix is currently bound to. Traffic sent to an address within an IP prefix will be routed to the Cloudflare service of the most-specific Service Binding matching the address. Example: binding 192.0.2.0/24 to Cloudflare Magic Transit and 192.0.2.1/32 to the Cloudflare CDN would route traffic for 192.0.2.1 to the CDN, and traffic for all other IPs in the prefix to Cloudflare Magic Transit.
client.Addressing.Prefixes.ServiceBindings.Get(ctx, prefixID, bindingID, query) (*
ServiceBinding
, error)
get/accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings/{binding_id}
Fetch a single Service Binding
client.Addressing.Prefixes.ServiceBindings.New(ctx, prefixID, params) (*
ServiceBinding
, error)
post/accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings
Creates a new Service Binding, routing traffic to IPs within the given CIDR to a service running on Cloudflare's network. NOTE: The first Service Binding created for an IP Prefix must exactly match the IP Prefix's CIDR. Subsequent Service Bindings may be created with a more-specific CIDR. Refer to the Service Bindings Documentation for compatibility details.
client.Addressing.Prefixes.ServiceBindings.Delete(ctx, prefixID, bindingID, body) (*
PrefixServiceBindingDeleteResponse
, error)
delete/accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings/{binding_id}
Delete a Service Binding
Domain types
typeServiceBindingstruct{…}
Addressing
Regional Hostnames
Addressing.RegionalHostnames
Methods
client.Addressing.RegionalHostnames.List(ctx, query) (*SinglePage[
RegionalHostnameListResponse
], error)
get/zones/{zone_id}/addressing/regional_hostnames
List all Regional Hostnames within a zone.
client.Addressing.RegionalHostnames.Get(ctx, hostname, query) (*
RegionalHostnameGetResponse
, error)
get/zones/{zone_id}/addressing/regional_hostnames/{hostname}
Fetch the configuration for a specific Regional Hostname, within a zone.
client.Addressing.RegionalHostnames.New(ctx, params) (*
RegionalHostnameNewResponse
, error)
post/zones/{zone_id}/addressing/regional_hostnames
Create a new Regional Hostname entry. Cloudflare will only use data centers that are physically located within the chosen region to decrypt and service HTTPS traffic. Learn more about Regional Services.
client.Addressing.RegionalHostnames.Edit(ctx, hostname, params) (*
RegionalHostnameEditResponse
, error)
patch/zones/{zone_id}/addressing/regional_hostnames/{hostname}
Update the configuration for a specific Regional Hostname. Only the region_key of a hostname is mutable.
client.Addressing.RegionalHostnames.Delete(ctx, hostname, body) (*
RegionalHostnameDeleteResponse
, error)
delete/zones/{zone_id}/addressing/regional_hostnames/{hostname}
Delete the region configuration for a specific Regional Hostname.
Addressing
Regional Hostnames
Regions
Addressing.RegionalHostnames.Regions
Methods
client.Addressing.RegionalHostnames.Regions.List(ctx, query) (*SinglePage[
RegionalHostnameRegionListResponse
], error)
get/accounts/{account_id}/addressing/regional_hostnames/regions
List all Regional Services regions available for use by this account.
Addressing
Services
Addressing.Services
Methods
client.Addressing.Services.List(ctx, query) (*SinglePage[
ServiceListResponse
], error)
get/accounts/{account_id}/addressing/services
Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a service running on the Cloudflare network to enable a Cloudflare product on the IP addresses. This endpoint can be used as a reference of available services on the Cloudflare network, and their service IDs.