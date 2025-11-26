DNS Firewall
DNSFirewall
Methods
client.DNSFirewall.List(ctx, params) (*V4PagePaginationArray[
DNSFirewallListResponse
], error)
get/accounts/{account_id}/dns_firewall
List DNS Firewall clusters for an account
client.DNSFirewall.Get(ctx, dnsFirewallID, query) (*
DNSFirewallGetResponse
, error)
get/accounts/{account_id}/dns_firewall/{dns_firewall_id}
Show a single DNS Firewall cluster for an account
client.DNSFirewall.New(ctx, params) (*
DNSFirewallNewResponse
, error)
post/accounts/{account_id}/dns_firewall
Create a DNS Firewall cluster
client.DNSFirewall.Edit(ctx, dnsFirewallID, params) (*
DNSFirewallEditResponse
, error)
patch/accounts/{account_id}/dns_firewall/{dns_firewall_id}
Modify the configuration of a DNS Firewall cluster
client.DNSFirewall.Delete(ctx, dnsFirewallID, body) (*
DNSFirewallDeleteResponse
, error)
delete/accounts/{account_id}/dns_firewall/{dns_firewall_id}
Delete a DNS Firewall cluster
Domain types
typeAttackMitigationstruct{…}
Attack mitigation settings
typeFirewallIPsstring
Cloudflare-assigned DNS IPv4 address
typeUpstreamIPsstring
Upstream DNS Server IPv4 address
DNS Firewall
Analytics
DNSFirewall.Analytics
DNS Firewall
Analytics
Reports
DNSFirewall.Analytics.Reports
Methods
client.DNSFirewall.Analytics.Reports.Get(ctx, dnsFirewallID, params) (*
Report
, error)
get/accounts/{account_id}/dns_firewall/{dns_firewall_id}/dns_analytics/report
Retrieves a list of summarised aggregate metrics over a given time period.
See Analytics API properties for detailed information about the available query parameters.
DNS Firewall
Analytics
Reports
Bytimes
DNSFirewall.Analytics.Reports.Bytimes
Methods
client.DNSFirewall.Analytics.Reports.Bytimes.Get(ctx, dnsFirewallID, params) (*
ByTime
, error)
get/accounts/{account_id}/dns_firewall/{dns_firewall_id}/dns_analytics/report/bytime
Retrieves a list of aggregate metrics grouped by time interval.
See Analytics API properties for detailed information about the available query parameters.
DNS Firewall
Reverse DNS
DNSFirewall.ReverseDNS
Methods
client.DNSFirewall.ReverseDNS.Get(ctx, dnsFirewallID, query) (*
ReverseDNSGetResponse
, error)
get/accounts/{account_id}/dns_firewall/{dns_firewall_id}/reverse_dns
Show reverse DNS configuration (PTR records) for a DNS Firewall cluster
client.DNSFirewall.ReverseDNS.Edit(ctx, dnsFirewallID, params) (*
ReverseDNSEditResponse
, error)
patch/accounts/{account_id}/dns_firewall/{dns_firewall_id}/reverse_dns
Update reverse DNS configuration (PTR records) for a DNS Firewall cluster