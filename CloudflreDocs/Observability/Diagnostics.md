Diagnostics
Diagnostics
Diagnostics
Endpoint Healthchecks
Diagnostics.EndpointHealthchecks
Methods
client.Diagnostics.EndpointHealthchecks.List(ctx, query) (*
EndpointHealthcheckListResponse
, error)
get/accounts/{account_id}/diagnostics/endpoint-healthchecks
List Endpoint Health Checks.
client.Diagnostics.EndpointHealthchecks.New(ctx, params) (*
EndpointHealthcheckNewResponse
, error)
post/accounts/{account_id}/diagnostics/endpoint-healthchecks
Create Endpoint Health Check.
client.Diagnostics.EndpointHealthchecks.Get(ctx, id, query) (*
EndpointHealthcheckGetResponse
, error)
get/accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}
Get a single Endpoint Health Check.
client.Diagnostics.EndpointHealthchecks.Delete(ctx, id, body) (*
EndpointHealthcheckDeleteResponse
, error)
delete/accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}
Delete Endpoint Health Check.
client.Diagnostics.EndpointHealthchecks.Update(ctx, id, params) (*
EndpointHealthcheckUpdateResponse
, error)
put/accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}
Update a Endpoint Health Check.
Domain types
typeEndpointHealthcheckstruct{…}
Diagnostics
Traceroutes
Diagnostics.Traceroutes
Methods
client.Diagnostics.Traceroutes.New(ctx, params) (*SinglePage[
Traceroute
], error)
post/accounts/{account_id}/diagnostics/traceroute
Run traceroutes from Cloudflare colos.
Domain types
typeTraceroutestruct{…}