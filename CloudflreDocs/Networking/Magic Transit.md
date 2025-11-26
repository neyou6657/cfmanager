Magic Transit
MagicTransit
Domain types
typeHealthCheckstruct{…}
typeHealthCheckRatestring
How frequent the health check is run. The default value is mid.
typeHealthCheckTypestring
The type of healthcheck to run, reply or request. The default value is reply.
Magic Transit
Apps
MagicTransit.Apps
Methods
client.MagicTransit.Apps.List(ctx, query) (*SinglePage[
AppListResponse
], error)
get/accounts/{account_id}/magic/apps
Lists Apps associated with an account.
client.MagicTransit.Apps.New(ctx, params) (*
AppNewResponse
, error)
post/accounts/{account_id}/magic/apps
Creates a new App for an account
client.MagicTransit.Apps.Update(ctx, accountAppID, params) (*
AppUpdateResponse
, error)
put/accounts/{account_id}/magic/apps/{account_app_id}
Updates an Account App
client.MagicTransit.Apps.Edit(ctx, accountAppID, params) (*
AppEditResponse
, error)
patch/accounts/{account_id}/magic/apps/{account_app_id}
Updates an Account App
client.MagicTransit.Apps.Delete(ctx, accountAppID, body) (*
AppDeleteResponse
, error)
delete/accounts/{account_id}/magic/apps/{account_app_id}
Deletes specific Account App.
Magic Transit
Cf Interconnects
MagicTransit.CfInterconnects
Methods
client.MagicTransit.CfInterconnects.List(ctx, params) (*
CfInterconnectListResponse
, error)
get/accounts/{account_id}/magic/cf_interconnects
Lists interconnects associated with an account.
client.MagicTransit.CfInterconnects.Get(ctx, cfInterconnectID, params) (*
CfInterconnectGetResponse
, error)
get/accounts/{account_id}/magic/cf_interconnects/{cf_interconnect_id}
Lists details for a specific interconnect.
client.MagicTransit.CfInterconnects.Update(ctx, cfInterconnectID, params) (*
CfInterconnectUpdateResponse
, error)
put/accounts/{account_id}/magic/cf_interconnects/{cf_interconnect_id}
Updates a specific interconnect associated with an account. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.CfInterconnects.BulkUpdate(ctx, params) (*
CfInterconnectBulkUpdateResponse
, error)
put/accounts/{account_id}/magic/cf_interconnects
Updates multiple interconnects associated with an account. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
Magic Transit
Connectors
MagicTransit.Connectors
Methods
client.MagicTransit.Connectors.List(ctx, query) (*SinglePage[
ConnectorListResponse
], error)
get/accounts/{account_id}/magic/connectors
List Connectors
client.MagicTransit.Connectors.Get(ctx, connectorID, query) (*
ConnectorGetResponse
, error)
get/accounts/{account_id}/magic/connectors/{connector_id}
Fetch Connector
client.MagicTransit.Connectors.New(ctx, params) (*
ConnectorNewResponse
, error)
post/accounts/{account_id}/magic/connectors
Add a connector to your account
client.MagicTransit.Connectors.Update(ctx, connectorID, params) (*
ConnectorUpdateResponse
, error)
put/accounts/{account_id}/magic/connectors/{connector_id}
Replace Connector or Re-provision License Key
client.MagicTransit.Connectors.Edit(ctx, connectorID, params) (*
ConnectorEditResponse
, error)
patch/accounts/{account_id}/magic/connectors/{connector_id}
Edit Connector to update specific properties or Re-provision License Key
client.MagicTransit.Connectors.Delete(ctx, connectorID, body) (*
ConnectorDeleteResponse
, error)
delete/accounts/{account_id}/magic/connectors/{connector_id}
Remove a connector from your account
Magic Transit
Connectors
Events
MagicTransit.Connectors.Events
Methods
client.MagicTransit.Connectors.Events.List(ctx, connectorID, params) (*
ConnectorEventListResponse
, error)
get/accounts/{account_id}/magic/connectors/{connector_id}/telemetry/events
List Events
client.MagicTransit.Connectors.Events.Get(ctx, connectorID, eventT, eventN, query) (*
ConnectorEventGetResponse
, error)
get/accounts/{account_id}/magic/connectors/{connector_id}/telemetry/events/{event_t}.{event_n}
Get Event
Magic Transit
Connectors
Events
Latest
MagicTransit.Connectors.Events.Latest
Methods
client.MagicTransit.Connectors.Events.Latest.List(ctx, connectorID, query) (*
ConnectorEventLatestListResponse
, error)
get/accounts/{account_id}/magic/connectors/{connector_id}/telemetry/events/latest
Get latest Events
Magic Transit
Connectors
Snapshots
MagicTransit.Connectors.Snapshots
Methods
client.MagicTransit.Connectors.Snapshots.List(ctx, connectorID, params) (*
ConnectorSnapshotListResponse
, error)
get/accounts/{account_id}/magic/connectors/{connector_id}/telemetry/snapshots
List Snapshots
client.MagicTransit.Connectors.Snapshots.Get(ctx, connectorID, snapshotT, query) (*
ConnectorSnapshotGetResponse
, error)
get/accounts/{account_id}/magic/connectors/{connector_id}/telemetry/snapshots/{snapshot_t}
Get Snapshot
Magic Transit
Connectors
Snapshots
Latest
MagicTransit.Connectors.Snapshots.Latest
Methods
client.MagicTransit.Connectors.Snapshots.Latest.List(ctx, connectorID, query) (*
ConnectorSnapshotLatestListResponse
, error)
get/accounts/{account_id}/magic/connectors/{connector_id}/telemetry/snapshots/latest
Get latest Snapshots
Magic Transit
GRE Tunnels
MagicTransit.GRETunnels
Methods
client.MagicTransit.GRETunnels.List(ctx, params) (*
GRETunnelListResponse
, error)
get/accounts/{account_id}/magic/gre_tunnels
Lists GRE tunnels associated with an account.
client.MagicTransit.GRETunnels.Get(ctx, greTunnelID, params) (*
GRETunnelGetResponse
, error)
get/accounts/{account_id}/magic/gre_tunnels/{gre_tunnel_id}
Lists informtion for a specific GRE tunnel.
client.MagicTransit.GRETunnels.New(ctx, params) (*
GRETunnelNewResponse
, error)
post/accounts/{account_id}/magic/gre_tunnels
Creates a new GRE tunnel. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.GRETunnels.Update(ctx, greTunnelID, params) (*
GRETunnelUpdateResponse
, error)
put/accounts/{account_id}/magic/gre_tunnels/{gre_tunnel_id}
Updates a specific GRE tunnel. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.GRETunnels.Delete(ctx, greTunnelID, params) (*
GRETunnelDeleteResponse
, error)
delete/accounts/{account_id}/magic/gre_tunnels/{gre_tunnel_id}
Disables and removes a specific static GRE tunnel. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.GRETunnels.BulkUpdate(ctx, params) (*
GRETunnelBulkUpdateResponse
, error)
put/accounts/{account_id}/magic/gre_tunnels
Updates multiple GRE tunnels. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
Magic Transit
IPSEC Tunnels
MagicTransit.IPSECTunnels
Methods
client.MagicTransit.IPSECTunnels.List(ctx, params) (*
IPSECTunnelListResponse
, error)
get/accounts/{account_id}/magic/ipsec_tunnels
Lists IPsec tunnels associated with an account.
client.MagicTransit.IPSECTunnels.Get(ctx, ipsecTunnelID, params) (*
IPSECTunnelGetResponse
, error)
get/accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}
Lists details for a specific IPsec tunnel.
client.MagicTransit.IPSECTunnels.New(ctx, params) (*
IPSECTunnelNewResponse
, error)
post/accounts/{account_id}/magic/ipsec_tunnels
Creates a new IPsec tunnel associated with an account. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.IPSECTunnels.Update(ctx, ipsecTunnelID, params) (*
IPSECTunnelUpdateResponse
, error)
put/accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}
Updates a specific IPsec tunnel associated with an account. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.IPSECTunnels.Delete(ctx, ipsecTunnelID, params) (*
IPSECTunnelDeleteResponse
, error)
delete/accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}
Disables and removes a specific static IPsec Tunnel associated with an account. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.IPSECTunnels.BulkUpdate(ctx, params) (*
IPSECTunnelBulkUpdateResponse
, error)
put/accounts/{account_id}/magic/ipsec_tunnels
Update multiple IPsec tunnels associated with an account. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes.
client.MagicTransit.IPSECTunnels.PSKGenerate(ctx, ipsecTunnelID, params) (*
IPSECTunnelPSKGenerateResponse
, error)
post/accounts/{account_id}/magic/ipsec_tunnels/{ipsec_tunnel_id}/psk_generate
Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session. Use ?validate_only=true as an optional query parameter to only run validation without persisting changes. After a PSK is generated, the PSK is immediately persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a safe place.
Domain types
typePSKMetadatastruct{…}
The PSK metadata that includes when the PSK was generated.
Magic Transit
PCAPs
MagicTransit.PCAPs
Methods
client.MagicTransit.PCAPs.List(ctx, query) (*SinglePage[
PCAPListResponse
], error)
get/accounts/{account_id}/pcaps
Lists all packet capture requests for an account.
client.MagicTransit.PCAPs.Get(ctx, pcapID, query) (*
PCAPGetResponse
, error)
get/accounts/{account_id}/pcaps/{pcap_id}
Get information for a PCAP request by id.
client.MagicTransit.PCAPs.New(ctx, params) (*
PCAPNewResponse
, error)
post/accounts/{account_id}/pcaps
Create new PCAP request for account.
client.MagicTransit.PCAPs.Stop(ctx, pcapID, body) error
put/accounts/{account_id}/pcaps/{pcap_id}/stop
Stop full PCAP.
Domain types
typePCAPstruct{…}
typePCAPFilterstruct{…}
The packet capture filter. When this field is empty, all packets are captured.
Magic Transit
PCAPs
Download
MagicTransit.PCAPs.Download
Methods
client.MagicTransit.PCAPs.Download.Get(ctx, pcapID, query) (*
Response
, error)
get/accounts/{account_id}/pcaps/{pcap_id}/download
Download PCAP information into a file. Response is a binary PCAP file.
Magic Transit
PCAPs
Ownership
MagicTransit.PCAPs.Ownership
Methods
client.MagicTransit.PCAPs.Ownership.Get(ctx, query) (*SinglePage[
Ownership
], error)
get/accounts/{account_id}/pcaps/ownership
List all buckets configured for use with PCAPs API.
client.MagicTransit.PCAPs.Ownership.New(ctx, params) (*
Ownership
, error)
post/accounts/{account_id}/pcaps/ownership
Adds an AWS or GCP bucket to use with full packet captures.
client.MagicTransit.PCAPs.Ownership.Delete(ctx, ownershipID, body) error
delete/accounts/{account_id}/pcaps/ownership/{ownership_id}
Deletes buckets added to the packet captures API.
client.MagicTransit.PCAPs.Ownership.Validate(ctx, params) (*
Ownership
, error)
post/accounts/{account_id}/pcaps/ownership/validate
Validates buckets added to the packet captures API.
Domain types
typeOwnershipstruct{…}
Magic Transit
Routes
MagicTransit.Routes
Methods
client.MagicTransit.Routes.List(ctx, query) (*
RouteListResponse
, error)
get/accounts/{account_id}/magic/routes
List all Magic static routes.
client.MagicTransit.Routes.Get(ctx, routeID, query) (*
RouteGetResponse
, error)
get/accounts/{account_id}/magic/routes/{route_id}
Get a specific Magic static route.
client.MagicTransit.Routes.New(ctx, params) (*
RouteNewResponse
, error)
post/accounts/{account_id}/magic/routes
Creates a new Magic static route. Use ?validate_only=true as an optional query parameter to run validation only without persisting changes.
client.MagicTransit.Routes.Update(ctx, routeID, params) (*
RouteUpdateResponse
, error)
put/accounts/{account_id}/magic/routes/{route_id}
Update a specific Magic static route. Use ?validate_only=true as an optional query parameter to run validation only without persisting changes.
client.MagicTransit.Routes.Delete(ctx, routeID, body) (*
RouteDeleteResponse
, error)
delete/accounts/{account_id}/magic/routes/{route_id}
Disable and remove a specific Magic static route.
client.MagicTransit.Routes.BulkUpdate(ctx, params) (*
RouteBulkUpdateResponse
, error)
put/accounts/{account_id}/magic/routes
Update multiple Magic static routes. Use ?validate_only=true as an optional query parameter to run validation only without persisting changes. Only fields for a route that need to be changed need be provided.
client.MagicTransit.Routes.Empty(ctx, body) (*
RouteEmptyResponse
, error)
delete/accounts/{account_id}/magic/routes
Delete multiple Magic static routes.
Domain types
typeScopestruct{…}
Used only for ECMP routes.
Magic Transit
Sites
MagicTransit.Sites
Methods
client.MagicTransit.Sites.List(ctx, params) (*SinglePage[
Site
], error)
get/accounts/{account_id}/magic/sites
Lists Sites associated with an account. Use connectorid query param to return sites where connectorid matches either site.ConnectorID or site.SecondaryConnectorID.
client.MagicTransit.Sites.Get(ctx, siteID, params) (*
Site
, error)
get/accounts/{account_id}/magic/sites/{site_id}
Get a specific Site.
client.MagicTransit.Sites.New(ctx, params) (*
Site
, error)
post/accounts/{account_id}/magic/sites
Creates a new Site
client.MagicTransit.Sites.Update(ctx, siteID, params) (*
Site
, error)
put/accounts/{account_id}/magic/sites/{site_id}
Update a specific Site.
client.MagicTransit.Sites.Edit(ctx, siteID, params) (*
Site
, error)
patch/accounts/{account_id}/magic/sites/{site_id}
Patch a specific Site.
client.MagicTransit.Sites.Delete(ctx, siteID, body) (*
Site
, error)
delete/accounts/{account_id}/magic/sites/{site_id}
Remove a specific Site.
Domain types
typeSitestruct{…}
typeSiteLocationstruct{…}
Location of site in latitude and longitude.
Magic Transit
Sites
ACLs
MagicTransit.Sites.ACLs
Methods
client.MagicTransit.Sites.ACLs.List(ctx, siteID, query) (*SinglePage[
ACL
], error)
get/accounts/{account_id}/magic/sites/{site_id}/acls
Lists Site ACLs associated with an account.
client.MagicTransit.Sites.ACLs.Get(ctx, siteID, aclID, query) (*
ACL
, error)
get/accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}
Get a specific Site ACL.
client.MagicTransit.Sites.ACLs.New(ctx, siteID, params) (*
ACL
, error)
post/accounts/{account_id}/magic/sites/{site_id}/acls
Creates a new Site ACL.
client.MagicTransit.Sites.ACLs.Update(ctx, siteID, aclID, params) (*
ACL
, error)
put/accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}
Update a specific Site ACL.
client.MagicTransit.Sites.ACLs.Edit(ctx, siteID, aclID, params) (*
ACL
, error)
patch/accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}
Patch a specific Site ACL.
client.MagicTransit.Sites.ACLs.Delete(ctx, siteID, aclID, body) (*
ACL
, error)
delete/accounts/{account_id}/magic/sites/{site_id}/acls/{acl_id}
Remove a specific Site ACL.
Domain types
typeACLstruct{…}
Bidirectional ACL policy for network traffic within a site.
typeACLConfigurationstruct{…}
typeAllowedProtocolstring
Array of allowed communication protocols between configured LANs. If no protocols are provided, all protocols are allowed.
typeSubnetstring
A valid IPv4 address.
Magic Transit
Sites
LANs
MagicTransit.Sites.LANs
Methods
client.MagicTransit.Sites.LANs.List(ctx, siteID, query) (*SinglePage[
LAN
], error)
get/accounts/{account_id}/magic/sites/{site_id}/lans
Lists Site LANs associated with an account.
client.MagicTransit.Sites.LANs.Get(ctx, siteID, lanID, query) (*
LAN
, error)
get/accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}
Get a specific Site LAN.
client.MagicTransit.Sites.LANs.New(ctx, siteID, params) (*SinglePage[
LAN
], error)
post/accounts/{account_id}/magic/sites/{site_id}/lans
Creates a new Site LAN. If the site is in high availability mode, static_addressing is required along with secondary and virtual address.
client.MagicTransit.Sites.LANs.Update(ctx, siteID, lanID, params) (*
LAN
, error)
put/accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}
Update a specific Site LAN.
client.MagicTransit.Sites.LANs.Edit(ctx, siteID, lanID, params) (*
LAN
, error)
patch/accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}
Patch a specific Site LAN.
client.MagicTransit.Sites.LANs.Delete(ctx, siteID, lanID, body) (*
LAN
, error)
delete/accounts/{account_id}/magic/sites/{site_id}/lans/{lan_id}
Remove a specific Site LAN.
Domain types
typeDHCPRelaystruct{…}
typeDHCPServerstruct{…}
typeLANstruct{…}
typeLANStaticAddressingstruct{…}
If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP). However, if in high availability mode, static_address is required along with secondary and virtual address.
typeNatstruct{…}
typeRoutedSubnetstruct{…}
Magic Transit
Sites
WANs
MagicTransit.Sites.WANs
Methods
client.MagicTransit.Sites.WANs.List(ctx, siteID, query) (*SinglePage[
WAN
], error)
get/accounts/{account_id}/magic/sites/{site_id}/wans
Lists Site WANs associated with an account.
client.MagicTransit.Sites.WANs.Get(ctx, siteID, wanID, query) (*
WAN
, error)
get/accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}
Get a specific Site WAN.
client.MagicTransit.Sites.WANs.New(ctx, siteID, params) (*SinglePage[
WAN
], error)
post/accounts/{account_id}/magic/sites/{site_id}/wans
Creates a new Site WAN.
client.MagicTransit.Sites.WANs.Update(ctx, siteID, wanID, params) (*
WAN
, error)
put/accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}
Update a specific Site WAN.
client.MagicTransit.Sites.WANs.Edit(ctx, siteID, wanID, params) (*
WAN
, error)
patch/accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}
Patch a specific Site WAN.
client.MagicTransit.Sites.WANs.Delete(ctx, siteID, wanID, body) (*
WAN
, error)
delete/accounts/{account_id}/magic/sites/{site_id}/wans/{wan_id}
Remove a specific Site WAN.
Domain types
typeWANstruct{…}
typeWANStaticAddressingstruct{…}
(optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.