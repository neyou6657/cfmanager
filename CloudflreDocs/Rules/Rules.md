Web3
Web3
Web3
Hostnames
Web3.Hostnames
Methods
client.Web3.Hostnames.List(ctx, query) (*SinglePage[
Hostname
], error)
get/zones/{zone_id}/web3/hostnames
List Web3 Hostnames
client.Web3.Hostnames.Get(ctx, identifier, query) (*
Hostname
, error)
get/zones/{zone_id}/web3/hostnames/{identifier}
Web3 Hostname Details
client.Web3.Hostnames.New(ctx, params) (*
Hostname
, error)
post/zones/{zone_id}/web3/hostnames
Create Web3 Hostname
client.Web3.Hostnames.Edit(ctx, identifier, params) (*
Hostname
, error)
patch/zones/{zone_id}/web3/hostnames/{identifier}
Edit Web3 Hostname
client.Web3.Hostnames.Delete(ctx, identifier, body) (*
HostnameDeleteResponse
, error)
delete/zones/{zone_id}/web3/hostnames/{identifier}
Delete Web3 Hostname
Domain types
typeHostnamestruct{…}
Web3
Hostnames
IPFS Universal Paths
Web3.Hostnames.IPFSUniversalPaths
Web3
Hostnames
IPFS Universal Paths
Content Lists
Web3.Hostnames.IPFSUniversalPaths.ContentLists
Methods
client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Get(ctx, identifier, query) (*
ContentList
, error)
get/zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list
IPFS Universal Path Gateway Content List Details
client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Update(ctx, identifier, params) (*
ContentList
, error)
put/zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list
Update IPFS Universal Path Gateway Content List
Domain types
typeContentListstruct{…}
Web3
Hostnames
IPFS Universal Paths
Content Lists
Entries
Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries
Methods
client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.List(ctx, identifier, query) (*
HostnameIPFSUniversalPathContentListEntryListResponse
, error)
get/zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries
List IPFS Universal Path Gateway Content List Entries
client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.Get(ctx, identifier, contentListEntryIdentifier, query) (*
HostnameIPFSUniversalPathContentListEntryGetResponse
, error)
get/zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries/{content_list_entry_identifier}
IPFS Universal Path Gateway Content List Entry Details
client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.New(ctx, identifier, params) (*
HostnameIPFSUniversalPathContentListEntryNewResponse
, error)
post/zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries
Create IPFS Universal Path Gateway Content List Entry
client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.Update(ctx, identifier, contentListEntryIdentifier, params) (*
HostnameIPFSUniversalPathContentListEntryUpdateResponse
, error)
put/zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries/{content_list_entry_identifier}
Edit IPFS Universal Path Gateway Content List Entry
client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.Delete(ctx, identifier, contentListEntryIdentifier, body) (*
HostnameIPFSUniversalPathContentListEntryDeleteResponse
, error)
delete/zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries/{content_list_entry_identifier}
Delete IPFS Universal Path Gateway Content List Entry