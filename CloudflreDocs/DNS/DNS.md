DNS
DNS
DNS
Analytics
DNS.Analytics
DNS
Analytics
Reports
DNS.Analytics.Reports
Methods
client.DNS.Analytics.Reports.Get(ctx, params) (*
Report
, error)
get/zones/{zone_id}/dns_analytics/report
Retrieves a list of summarised aggregate metrics over a given time period.
See Analytics API properties for detailed information about the available query parameters.
Domain types
typeReportstruct{…}
DNS
Analytics
Reports
Bytimes
DNS.Analytics.Reports.Bytimes
Methods
client.DNS.Analytics.Reports.Bytimes.Get(ctx, params) (*
ByTime
, error)
get/zones/{zone_id}/dns_analytics/report/bytime
Retrieves a list of aggregate metrics grouped by time interval.
See Analytics API properties for detailed information about the available query parameters.
Domain types
typeByTimestruct{…}
DNS
DNSSEC
DNS.DNSSEC
Methods
client.DNS.DNSSEC.Get(ctx, query) (*
DNSSEC
, error)
get/zones/{zone_id}/dnssec
Details about DNSSEC status and configuration.
client.DNS.DNSSEC.Edit(ctx, params) (*
DNSSEC
, error)
patch/zones/{zone_id}/dnssec
Enable or disable DNSSEC.
client.DNS.DNSSEC.Delete(ctx, body) (*string, error)
delete/zones/{zone_id}/dnssec
Delete DNSSEC.
Domain types
typeDNSSECstruct{…}
DNS
Records
DNS.Records
Methods
client.DNS.Records.List(ctx, params) (*V4PagePaginationArray[
RecordResponse
], error)
get/zones/{zone_id}/dns_records
List, search, sort, and filter a zones' DNS records.
client.DNS.Records.Get(ctx, dnsRecordID, query) (*
RecordResponse
, error)
get/zones/{zone_id}/dns_records/{dns_record_id}
DNS Record Details
client.DNS.Records.New(ctx, params) (*
RecordResponse
, error)
post/zones/{zone_id}/dns_records
Create a new DNS record for a zone.
Notes:
A/AAAA records cannot exist on the same name as CNAME records.
NS records cannot exist on the same name as any other record type.
Domain names are always represented in Punycode, even if Unicode characters were used when creating the record.
client.DNS.Records.Update(ctx, dnsRecordID, params) (*
RecordResponse
, error)
put/zones/{zone_id}/dns_records/{dns_record_id}
Overwrite an existing DNS record.
Notes:
A/AAAA records cannot exist on the same name as CNAME records.
NS records cannot exist on the same name as any other record type.
Domain names are always represented in Punycode, even if Unicode characters were used when creating the record.
client.DNS.Records.Edit(ctx, dnsRecordID, params) (*
RecordResponse
, error)
patch/zones/{zone_id}/dns_records/{dns_record_id}
Update an existing DNS record.
Notes:
A/AAAA records cannot exist on the same name as CNAME records.
NS records cannot exist on the same name as any other record type.
Domain names are always represented in Punycode, even if Unicode characters were used when creating the record.
client.DNS.Records.Delete(ctx, dnsRecordID, body) (*
RecordDeleteResponse
, error)
delete/zones/{zone_id}/dns_records/{dns_record_id}
Delete DNS Record
client.DNS.Records.Export(ctx, query) (*string, error)
get/zones/{zone_id}/dns_records/export
You can export your BIND config through this endpoint.
See the documentation for more information.
client.DNS.Records.Import(ctx, params) (*
RecordImportResponse
, error)
post/zones/{zone_id}/dns_records/import
You can upload your BIND config through this endpoint. It assumes that cURL is called from a location with bind_config.txt (valid BIND config) present.
See the documentation for more information.
client.DNS.Records.Scan(ctx, params) (*
RecordScanResponse
, error)
Deprecated
post/zones/{zone_id}/dns_records/scan
Deprecated
This endpoint is deprecated in favor of a new asynchronous version. Please use the /scan/trigger and /scan/review endpoints instead.
Scan for common DNS records on your domain and automatically add them to your zone. Useful if you haven't updated your nameservers yet.
client.DNS.Records.ScanTrigger(ctx, body) (*
RecordScanTriggerResponse
, error)
post/zones/{zone_id}/dns_records/scan/trigger
Initiates an asynchronous scan for common DNS records on your domain. Note that this does not automatically add records to your zone. The scan runs in the background, and results can be reviewed later using the /scan/review endpoints. Useful if you haven't updated your nameservers yet.
client.DNS.Records.ScanReview(ctx, params) (*
RecordScanReviewResponse
, error)
post/zones/{zone_id}/dns_records/scan/review
Accept or reject DNS records found by the DNS records scan. Accepted records will be permanently added to the zone, while rejected records will be permanently deleted.
client.DNS.Records.ScanList(ctx, query) (*SinglePage[
RecordResponse
], error)
get/zones/{zone_id}/dns_records/scan/review
Retrieves the list of DNS records discovered up to this point by the asynchronous scan. These records are temporary until explicitly accepted or rejected via POST /scan/review. Additional records may be discovered by the scan later.
client.DNS.Records.Batch(ctx, params) (*
RecordBatchResponse
, error)
post/zones/{zone_id}/dns_records/batch
Send a Batch of DNS Record API calls to be executed together.
Notes:
Although Cloudflare will execute the batched operations in a single database transaction, Cloudflare's distributed KV store must treat each record change as a single key-value pair. This means that the propagation of changes is not atomic. See the documentation for more information.
The operations you specify within the /batch request body are always executed in the following order:
Deletes
Patches
Puts
Posts
Domain types
typeARecordstruct{…}
typeAAAARecordstruct{…}
typeBatchPatchinterface{…}
typeBatchPutinterface{…}
typeCAARecordstruct{…}
typeCERTRecordstruct{…}
typeCNAMERecordstruct{…}
typeDNSKEYRecordstruct{…}
typeDSRecordstruct{…}
typeHTTPSRecordstruct{…}
typeLOCRecordstruct{…}
typeMXRecordstruct{…}
typeNAPTRRecordstruct{…}
typeNSRecordstruct{…}
typePTRRecordstruct{…}
typeRecordinterface{…}
typeRecordResponseinterface{…}
typeRecordTagsstring
Individual tag of the form name:value (the name must consist of only letters, numbers, underscores and hyphens)
typeSMIMEARecordstruct{…}
typeSRVRecordstruct{…}
typeSSHFPRecordstruct{…}
typeSVCBRecordstruct{…}
typeTLSARecordstruct{…}
typeTTLinterface{…}
Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'. Value must be between 60 and 86400, with the minimum reduced to 30 for Enterprise zones.
typeTXTRecordstruct{…}
typeURIRecordstruct{…}
DNS
Settings
DNS.Settings
DNS
Settings
Account
DNS.Settings.Account
Methods
client.DNS.Settings.Account.Get(ctx, query) (*
SettingAccountGetResponse
, error)
get/accounts/{account_id}/dns_settings
Show DNS settings for an account
client.DNS.Settings.Account.Edit(ctx, params) (*
SettingAccountEditResponse
, error)
patch/accounts/{account_id}/dns_settings
Update DNS settings for an account
DNS
Settings
Account
Views
DNS.Settings.Account.Views
Methods
client.DNS.Settings.Account.Views.List(ctx, params) (*V4PagePaginationArray[
SettingAccountViewListResponse
], error)
get/accounts/{account_id}/dns_settings/views
List DNS Internal Views for an Account
client.DNS.Settings.Account.Views.Get(ctx, viewID, query) (*
SettingAccountViewGetResponse
, error)
get/accounts/{account_id}/dns_settings/views/{view_id}
Get DNS Internal View
client.DNS.Settings.Account.Views.New(ctx, params) (*
SettingAccountViewNewResponse
, error)
post/accounts/{account_id}/dns_settings/views
Create Internal DNS View for an account
client.DNS.Settings.Account.Views.Edit(ctx, viewID, params) (*
SettingAccountViewEditResponse
, error)
patch/accounts/{account_id}/dns_settings/views/{view_id}
Update an existing Internal DNS View
client.DNS.Settings.Account.Views.Delete(ctx, viewID, body) (*
SettingAccountViewDeleteResponse
, error)
delete/accounts/{account_id}/dns_settings/views/{view_id}
Delete an existing Internal DNS View
DNS
Settings
Zone
DNS.Settings.Zone
Methods
client.DNS.Settings.Zone.Get(ctx, query) (*
SettingZoneGetResponse
, error)
get/zones/{zone_id}/dns_settings
Show DNS settings for a zone
client.DNS.Settings.Zone.Edit(ctx, params) (*
SettingZoneEditResponse
, error)
patch/zones/{zone_id}/dns_settings
Update DNS settings for a zone
DNS
Zone Transfers
DNS.ZoneTransfers
DNS
Zone Transfers
ACLs
DNS.ZoneTransfers.ACLs
Methods
client.DNS.ZoneTransfers.ACLs.List(ctx, query) (*SinglePage[
ACL
], error)
get/accounts/{account_id}/secondary_dns/acls
List ACLs.
client.DNS.ZoneTransfers.ACLs.Get(ctx, aclID, query) (*
ACL
, error)
get/accounts/{account_id}/secondary_dns/acls/{acl_id}
Get ACL.
client.DNS.ZoneTransfers.ACLs.New(ctx, params) (*
ACL
, error)
post/accounts/{account_id}/secondary_dns/acls
Create ACL.
client.DNS.ZoneTransfers.ACLs.Update(ctx, aclID, params) (*
ACL
, error)
put/accounts/{account_id}/secondary_dns/acls/{acl_id}
Modify ACL.
client.DNS.ZoneTransfers.ACLs.Delete(ctx, aclID, body) (*
ZoneTransferACLDeleteResponse
, error)
delete/accounts/{account_id}/secondary_dns/acls/{acl_id}
Delete ACL.
Domain types
typeACLstruct{…}
DNS
Zone Transfers
Force AXFR
DNS.ZoneTransfers.ForceAXFR
Methods
client.DNS.ZoneTransfers.ForceAXFR.New(ctx, params) (*
ForceAXFR
, error)
post/zones/{zone_id}/secondary_dns/force_axfr
Sends AXFR zone transfer request to primary nameserver(s).
Domain types
typeForceAXFRstring
When force_axfr query parameter is set to true, the response is a simple string.
DNS
Zone Transfers
Incoming
DNS.ZoneTransfers.Incoming
Methods
client.DNS.ZoneTransfers.Incoming.Get(ctx, query) (*
ZoneTransferIncomingGetResponse
, error)
get/zones/{zone_id}/secondary_dns/incoming
Get secondary zone configuration for incoming zone transfers.
client.DNS.ZoneTransfers.Incoming.New(ctx, params) (*
ZoneTransferIncomingNewResponse
, error)
post/zones/{zone_id}/secondary_dns/incoming
Create secondary zone configuration for incoming zone transfers.
client.DNS.ZoneTransfers.Incoming.Update(ctx, params) (*
ZoneTransferIncomingUpdateResponse
, error)
put/zones/{zone_id}/secondary_dns/incoming
Update secondary zone configuration for incoming zone transfers.
client.DNS.ZoneTransfers.Incoming.Delete(ctx, body) (*
ZoneTransferIncomingDeleteResponse
, error)
delete/zones/{zone_id}/secondary_dns/incoming
Delete secondary zone configuration for incoming zone transfers.
Domain types
typeIncomingstruct{…}
DNS
Zone Transfers
Outgoing
DNS.ZoneTransfers.Outgoing
Methods
client.DNS.ZoneTransfers.Outgoing.Get(ctx, query) (*
ZoneTransferOutgoingGetResponse
, error)
get/zones/{zone_id}/secondary_dns/outgoing
Get primary zone configuration for outgoing zone transfers.
client.DNS.ZoneTransfers.Outgoing.New(ctx, params) (*
ZoneTransferOutgoingNewResponse
, error)
post/zones/{zone_id}/secondary_dns/outgoing
Create primary zone configuration for outgoing zone transfers.
client.DNS.ZoneTransfers.Outgoing.Update(ctx, params) (*
ZoneTransferOutgoingUpdateResponse
, error)
put/zones/{zone_id}/secondary_dns/outgoing
Update primary zone configuration for outgoing zone transfers.
client.DNS.ZoneTransfers.Outgoing.Delete(ctx, body) (*
ZoneTransferOutgoingDeleteResponse
, error)
delete/zones/{zone_id}/secondary_dns/outgoing
Delete primary zone configuration for outgoing zone transfers.
client.DNS.ZoneTransfers.Outgoing.Disable(ctx, params) (*
DisableTransfer
, error)
post/zones/{zone_id}/secondary_dns/outgoing/disable
Disable outgoing zone transfers for primary zone and clears IXFR backlog of primary zone.
client.DNS.ZoneTransfers.Outgoing.Enable(ctx, params) (*
EnableTransfer
, error)
post/zones/{zone_id}/secondary_dns/outgoing/enable
Enable outgoing zone transfers for primary zone.
client.DNS.ZoneTransfers.Outgoing.ForceNotify(ctx, params) (*string, error)
post/zones/{zone_id}/secondary_dns/outgoing/force_notify
Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
Domain types
typeDisableTransferstring
The zone transfer status of a primary zone.
typeEnableTransferstring
The zone transfer status of a primary zone.
typeOutgoingstruct{…}
typeOutgoingStatusstring
The zone transfer status of a primary zone.
DNS
Zone Transfers
Outgoing
Status
DNS.ZoneTransfers.Outgoing.Status
Methods
client.DNS.ZoneTransfers.Outgoing.Status.Get(ctx, query) (*
EnableTransfer
, error)
get/zones/{zone_id}/secondary_dns/outgoing/status
Get primary zone transfer status.
DNS
Zone Transfers
Peers
DNS.ZoneTransfers.Peers
Methods
client.DNS.ZoneTransfers.Peers.List(ctx, query) (*SinglePage[
Peer
], error)
get/accounts/{account_id}/secondary_dns/peers
List Peers.
client.DNS.ZoneTransfers.Peers.Get(ctx, peerID, query) (*
Peer
, error)
get/accounts/{account_id}/secondary_dns/peers/{peer_id}
Get Peer.
client.DNS.ZoneTransfers.Peers.New(ctx, params) (*
Peer
, error)
post/accounts/{account_id}/secondary_dns/peers
Create Peer.
client.DNS.ZoneTransfers.Peers.Update(ctx, peerID, params) (*
Peer
, error)
put/accounts/{account_id}/secondary_dns/peers/{peer_id}
Modify Peer.
client.DNS.ZoneTransfers.Peers.Delete(ctx, peerID, body) (*
ZoneTransferPeerDeleteResponse
, error)
delete/accounts/{account_id}/secondary_dns/peers/{peer_id}
Delete Peer.
Domain types
typePeerstruct{…}
DNS
Zone Transfers
TSIGs
DNS.ZoneTransfers.TSIGs
Methods
client.DNS.ZoneTransfers.TSIGs.List(ctx, query) (*SinglePage[
TSIG
], error)
get/accounts/{account_id}/secondary_dns/tsigs
List TSIGs.
client.DNS.ZoneTransfers.TSIGs.Get(ctx, tsigID, query) (*
TSIG
, error)
get/accounts/{account_id}/secondary_dns/tsigs/{tsig_id}
Get TSIG.
client.DNS.ZoneTransfers.TSIGs.New(ctx, params) (*
TSIG
, error)
post/accounts/{account_id}/secondary_dns/tsigs
Create TSIG.
client.DNS.ZoneTransfers.TSIGs.Update(ctx, tsigID, params) (*
TSIG
, error)
put/accounts/{account_id}/secondary_dns/tsigs/{tsig_id}
Modify TSIG.
client.DNS.ZoneTransfers.TSIGs.Delete(ctx, tsigID, body) (*
ZoneTransferTSIGDeleteResponse
, error)
delete/accounts/{account_id}/secondary_dns/tsigs/{tsig_id}
Delete TSIG.
Domain types
typeTSIGstruct{…}