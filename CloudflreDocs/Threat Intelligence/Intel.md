Intel
Intel
Intel
ASN
Intel.ASN
Methods
client.Intel.ASN.Get(ctx, asn, query) (*
ASN
, error)
get/accounts/{account_id}/intel/asn/{asn}
Gets an overview of the Autonomous System Number (ASN) and a list of subnets for it.
Intel
ASN
Subnets
Intel.ASN.Subnets
Methods
client.Intel.ASN.Subnets.Get(ctx, asn, query) (*
ASNSubnetGetResponse
, error)
get/accounts/{account_id}/intel/asn/{asn}/subnets
Get ASN Subnets.
Intel
Attack Surface Report
Intel.AttackSurfaceReport
Intel
Attack Surface Report
Issue Types
Intel.AttackSurfaceReport.IssueTypes
Methods
client.Intel.AttackSurfaceReport.IssueTypes.Get(ctx, query) (*SinglePage[string], error)
get/accounts/{account_id}/intel/attack-surface-report/issue-types
Get Security Center Issues Types
Intel
Attack Surface Report
Issues
Intel.AttackSurfaceReport.Issues
Methods
client.Intel.AttackSurfaceReport.Issues.List(ctx, params) (*V4PagePagination[
AttackSurfaceReportIssueListResponse
], error)
Deprecated
get/accounts/{account_id}/intel/attack-surface-report/issues
Get Security Center Issues
client.Intel.AttackSurfaceReport.Issues.Class(ctx, params) (*[]
AttackSurfaceReportIssueClassResponse
, error)
Deprecated
get/accounts/{account_id}/intel/attack-surface-report/issues/class
Get Security Center Issue Counts by Class
client.Intel.AttackSurfaceReport.Issues.Severity(ctx, params) (*[]
AttackSurfaceReportIssueSeverityResponse
, error)
Deprecated
get/accounts/{account_id}/intel/attack-surface-report/issues/severity
Get Security Center Issue Counts by Severity
client.Intel.AttackSurfaceReport.Issues.Type(ctx, params) (*[]
AttackSurfaceReportIssueTypeResponse
, error)
Deprecated
get/accounts/{account_id}/intel/attack-surface-report/issues/type
Get Security Center Issue Counts by Type
client.Intel.AttackSurfaceReport.Issues.Dismiss(ctx, issueID, params) (*
AttackSurfaceReportIssueDismissResponse
, error)
Deprecated
put/accounts/{account_id}/intel/attack-surface-report/{issue_id}/dismiss
Archive Security Center Insight
Domain types
typeIssueTypestring
typeSeverityQueryParamstring
Intel
DNS
Intel.DNS
Methods
client.Intel.DNS.List(ctx, params) (*V4PagePagination[
DNS
], error)
get/accounts/{account_id}/intel/dns
Gets a list of all the domains that have resolved to a specific IP address.
Domain types
typeDNSstruct{…}
Intel
Domain History
Intel.DomainHistory
Methods
client.Intel.DomainHistory.Get(ctx, params) (*[]
DomainHistory
, error)
get/accounts/{account_id}/intel/domain-history
Gets historical security threat and content categories currently and previously assigned to a domain.
Domain types
typeDomainHistorystruct{…}
Intel
Domains
Intel.Domains
Methods
client.Intel.Domains.Get(ctx, params) (*
Domain
, error)
get/accounts/{account_id}/intel/domain
Gets security details and statistics about a domain.
Domain types
typeDomainstruct{…}
Intel
Domains
Bulks
Intel.Domains.Bulks
Methods
client.Intel.Domains.Bulks.Get(ctx, params) (*[]
DomainBulkGetResponse
, error)
get/accounts/{account_id}/intel/domain/bulk
Same as summary.
Intel
Indicator Feeds
Intel.IndicatorFeeds
Methods
client.Intel.IndicatorFeeds.List(ctx, query) (*SinglePage[
IndicatorFeedListResponse
], error)
get/accounts/{account_id}/intel/indicator-feeds
Get indicator feeds owned by this account
client.Intel.IndicatorFeeds.Get(ctx, feedID, query) (*
IndicatorFeedGetResponse
, error)
get/accounts/{account_id}/intel/indicator-feeds/{feed_id}
Get indicator feed metadata
client.Intel.IndicatorFeeds.New(ctx, params) (*
IndicatorFeedNewResponse
, error)
post/accounts/{account_id}/intel/indicator-feeds
Create new indicator feed
client.Intel.IndicatorFeeds.Update(ctx, feedID, params) (*
IndicatorFeedUpdateResponse
, error)
put/accounts/{account_id}/intel/indicator-feeds/{feed_id}
Update indicator feed metadata
client.Intel.IndicatorFeeds.Data(ctx, feedID, query) (*string, error)
get/accounts/{account_id}/intel/indicator-feeds/{feed_id}/data
Get indicator feed data
Intel
Indicator Feeds
Downloads
Intel.IndicatorFeeds.Downloads
Intel
Indicator Feeds
Permissions
Intel.IndicatorFeeds.Permissions
Methods
client.Intel.IndicatorFeeds.Permissions.List(ctx, query) (*[]
IndicatorFeedPermissionListResponse
, error)
get/accounts/{account_id}/intel/indicator-feeds/permissions/view
List indicator feed permissions
client.Intel.IndicatorFeeds.Permissions.New(ctx, params) (*
IndicatorFeedPermissionNewResponse
, error)
put/accounts/{account_id}/intel/indicator-feeds/permissions/add
Grant permission to indicator feed
client.Intel.IndicatorFeeds.Permissions.Delete(ctx, params) (*
IndicatorFeedPermissionDeleteResponse
, error)
put/accounts/{account_id}/intel/indicator-feeds/permissions/remove
Revoke permission to indicator feed
Intel
Indicator Feeds
Snapshots
Intel.IndicatorFeeds.Snapshots
Methods
client.Intel.IndicatorFeeds.Snapshots.Update(ctx, feedID, params) (*
IndicatorFeedSnapshotUpdateResponse
, error)
put/accounts/{account_id}/intel/indicator-feeds/{feed_id}/snapshot
Update indicator feed data
Intel
IP Lists
Intel.IPLists
Methods
client.Intel.IPLists.Get(ctx, query) (*SinglePage[
IPList
], error)
get/accounts/{account_id}/intel/ip-list
Get IP Lists.
Domain types
typeIPListstruct{…}
Intel
IPs
Intel.IPs
Methods
client.Intel.IPs.Get(ctx, params) (*[]
IP
, error)
get/accounts/{account_id}/intel/ip
Gets the geolocation, ASN, infrastructure type of the ASN, and any security threat categories of an IP address. Must provide ip query parameters. For example, /intel/ip?ipv4=1.1.1.1 or /intel/ip?ipv6=2001:db8::1.
Domain types
typeIPstruct{…}
Intel
Miscategorizations
Intel.Miscategorizations
Methods
client.Intel.Miscategorizations.New(ctx, params) (*
MiscategorizationNewResponse
, error)
post/accounts/{account_id}/intel/miscategorization
Allows you to submit requests to change a domain’s category.
Intel
Sinkholes
Intel.Sinkholes
Methods
client.Intel.Sinkholes.List(ctx, query) (*SinglePage[
Sinkhole
], error)
get/accounts/{account_id}/intel/sinkholes
List sinkholes owned by this account
Domain types
typeSinkholestruct{…}
Intel
Whois
Intel.Whois
Methods
client.Intel.Whois.Get(ctx, params) (*
WhoisGetResponse
, error)
get/accounts/{account_id}/intel/whois
Get WHOIS Record
Domain types
typeWhoisstruct{…}