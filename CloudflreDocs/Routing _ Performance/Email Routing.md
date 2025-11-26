Email Routing
EmailRouting
Methods
client.EmailRouting.Get(ctx, query) (*
Settings
, error)
get/zones/{zone_id}/email/routing
Get information about the settings for your Email Routing zone.
client.EmailRouting.Disable(ctx, params) (*
Settings
, error)
Deprecated
post/zones/{zone_id}/email/routing/disable
Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.
client.EmailRouting.Enable(ctx, params) (*
Settings
, error)
Deprecated
post/zones/{zone_id}/email/routing/enable
Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
Domain types
typeSettingsstruct{…}
Email Routing
Addresses
EmailRouting.Addresses
Methods
client.EmailRouting.Addresses.List(ctx, params) (*V4PagePaginationArray[
Address
], error)
get/accounts/{account_id}/email/routing/addresses
Lists existing destination addresses.
client.EmailRouting.Addresses.Get(ctx, destinationAddressIdentifier, query) (*
Address
, error)
get/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}
Gets information for a specific destination email already created.
client.EmailRouting.Addresses.New(ctx, params) (*
Address
, error)
post/accounts/{account_id}/email/routing/addresses
Create a destination address to forward your emails to. Destination addresses need to be verified before they can be used.
client.EmailRouting.Addresses.Delete(ctx, destinationAddressIdentifier, body) (*
Address
, error)
delete/accounts/{account_id}/email/routing/addresses/{destination_address_identifier}
Deletes a specific destination address.
Domain types
typeAddressstruct{…}
Email Routing
DNS
EmailRouting.DNS
Methods
client.EmailRouting.DNS.Get(ctx, params) (*
DNSGetResponse
, error)
get/zones/{zone_id}/email/routing/dns
Show the DNS records needed to configure your Email Routing zone.
client.EmailRouting.DNS.New(ctx, params) (*
Settings
, error)
post/zones/{zone_id}/email/routing/dns
Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
client.EmailRouting.DNS.Edit(ctx, params) (*
Settings
, error)
patch/zones/{zone_id}/email/routing/dns
Unlock MX Records previously locked by Email Routing.
client.EmailRouting.DNS.Delete(ctx, body) (*SinglePage[
DNSRecord
], error)
delete/zones/{zone_id}/email/routing/dns
Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.
Domain types
typeDNSRecordstruct{…}
List of records needed to enable an Email Routing zone.
Email Routing
Rules
EmailRouting.Rules
Methods
client.EmailRouting.Rules.List(ctx, params) (*V4PagePaginationArray[
EmailRoutingRule
], error)
get/zones/{zone_id}/email/routing/rules
Lists existing routing rules.
client.EmailRouting.Rules.Get(ctx, ruleIdentifier, query) (*
EmailRoutingRule
, error)
get/zones/{zone_id}/email/routing/rules/{rule_identifier}
Get information for a specific routing rule already created.
client.EmailRouting.Rules.New(ctx, params) (*
EmailRoutingRule
, error)
post/zones/{zone_id}/email/routing/rules
Rules consist of a set of criteria for matching emails (such as an email being sent to a specific custom email address) plus a set of actions to take on the email (like forwarding it to a specific destination address).
client.EmailRouting.Rules.Update(ctx, ruleIdentifier, params) (*
EmailRoutingRule
, error)
put/zones/{zone_id}/email/routing/rules/{rule_identifier}
Update actions and matches, or enable/disable specific routing rules.
client.EmailRouting.Rules.Delete(ctx, ruleIdentifier, body) (*
EmailRoutingRule
, error)
delete/zones/{zone_id}/email/routing/rules/{rule_identifier}
Delete a specific routing rule.
Domain types
typeActionstruct{…}
Actions pattern.
typeEmailRoutingRulestruct{…}
typeMatcherstruct{…}
Matching pattern to forward your actions.
Email Routing
Rules
Catch Alls
EmailRouting.Rules.CatchAlls
Methods
client.EmailRouting.Rules.CatchAlls.Get(ctx, query) (*
RuleCatchAllGetResponse
, error)
get/zones/{zone_id}/email/routing/rules/catch_all
Get information on the default catch-all routing rule.
client.EmailRouting.Rules.CatchAlls.Update(ctx, params) (*
RuleCatchAllUpdateResponse
, error)
put/zones/{zone_id}/email/routing/rules/catch_all
Enable or disable catch-all routing rule, or change action to forward to specific destination address.
Domain types
typeCatchAllActionstruct{…}
Action for the catch-all routing rule.
typeCatchAllMatcherstruct{…}
Matcher for catch-all routing rule.