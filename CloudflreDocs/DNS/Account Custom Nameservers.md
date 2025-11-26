Account Custom Nameservers
CustomNameservers
Methods
client.CustomNameservers.Get(ctx, query) (*SinglePage[
CustomNameserver
], error)
get/accounts/{account_id}/custom_ns
List an account's custom nameservers.
client.CustomNameservers.New(ctx, params) (*
CustomNameserver
, error)
post/accounts/{account_id}/custom_ns
Add Account Custom Nameserver
client.CustomNameservers.Delete(ctx, customNSID, body) (*SinglePage[string], error)
delete/accounts/{account_id}/custom_ns/{custom_ns_id}
Delete Account Custom Nameserver
Domain types
typeCustomNameserverstruct{…}
A single account custom nameserver.