Page Shield
PageShield
Methods
client.PageShield.Get(ctx, query) (*
Setting
, error)
get/zones/{zone_id}/page_shield
Fetches the Page Shield settings.
client.PageShield.Update(ctx, params) (*
PageShieldUpdateResponse
, error)
put/zones/{zone_id}/page_shield
Updates Page Shield settings.
Domain types
typeSettingstruct{…}
Page Shield
Connections
PageShield.Connections
Methods
client.PageShield.Connections.List(ctx, params) (*SinglePage[
Connection
], error)
get/zones/{zone_id}/page_shield/connections
Lists all connections detected by Page Shield.
client.PageShield.Connections.Get(ctx, connectionID, query) (*
Connection
, error)
get/zones/{zone_id}/page_shield/connections/{connection_id}
Fetches a connection detected by Page Shield by connection ID.
Domain types
typeConnectionstruct{…}
Page Shield
Cookies
PageShield.Cookies
Methods
client.PageShield.Cookies.List(ctx, params) (*SinglePage[
CookieListResponse
], error)
get/zones/{zone_id}/page_shield/cookies
Lists all cookies collected by Page Shield.
client.PageShield.Cookies.Get(ctx, cookieID, query) (*
CookieGetResponse
, error)
get/zones/{zone_id}/page_shield/cookies/{cookie_id}
Fetches a cookie collected by Page Shield by cookie ID.
Page Shield
Policies
PageShield.Policies
Methods
client.PageShield.Policies.List(ctx, query) (*SinglePage[
PolicyListResponse
], error)
get/zones/{zone_id}/page_shield/policies
Lists all Page Shield policies.
client.PageShield.Policies.Get(ctx, policyID, query) (*
PolicyGetResponse
, error)
get/zones/{zone_id}/page_shield/policies/{policy_id}
Fetches a Page Shield policy by ID.
client.PageShield.Policies.New(ctx, params) (*
PolicyNewResponse
, error)
post/zones/{zone_id}/page_shield/policies
Create a Page Shield policy.
client.PageShield.Policies.Update(ctx, policyID, params) (*
PolicyUpdateResponse
, error)
put/zones/{zone_id}/page_shield/policies/{policy_id}
Update a Page Shield policy by ID.
client.PageShield.Policies.Delete(ctx, policyID, body) error
delete/zones/{zone_id}/page_shield/policies/{policy_id}
Delete a Page Shield policy by ID.
Domain types
typePolicystruct{…}
Page Shield
Scripts
PageShield.Scripts
Methods
client.PageShield.Scripts.List(ctx, params) (*SinglePage[
Script
], error)
get/zones/{zone_id}/page_shield/scripts
Lists all scripts detected by Page Shield.
client.PageShield.Scripts.Get(ctx, scriptID, query) (*
ScriptGetResponse
, error)
get/zones/{zone_id}/page_shield/scripts/{script_id}
Fetches a script detected by Page Shield by script ID.
Domain types
typeScriptstruct{…}