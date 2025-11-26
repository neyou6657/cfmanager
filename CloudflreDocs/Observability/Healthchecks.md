Healthchecks
Healthchecks
Methods
client.Healthchecks.List(ctx, params) (*V4PagePaginationArray[
Healthcheck
], error)
get/zones/{zone_id}/healthchecks
List configured health checks.
client.Healthchecks.Get(ctx, healthcheckID, query) (*
Healthcheck
, error)
get/zones/{zone_id}/healthchecks/{healthcheck_id}
Fetch a single configured health check.
client.Healthchecks.New(ctx, params) (*
Healthcheck
, error)
post/zones/{zone_id}/healthchecks
Create a new health check.
client.Healthchecks.Update(ctx, healthcheckID, params) (*
Healthcheck
, error)
put/zones/{zone_id}/healthchecks/{healthcheck_id}
Update a configured health check.
client.Healthchecks.Edit(ctx, healthcheckID, params) (*
Healthcheck
, error)
patch/zones/{zone_id}/healthchecks/{healthcheck_id}
Patch a configured health check.
client.Healthchecks.Delete(ctx, healthcheckID, body) (*
HealthcheckDeleteResponse
, error)
delete/zones/{zone_id}/healthchecks/{healthcheck_id}
Delete a health check.
Domain types
typeCheckRegionstring
WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS and ENTERPRISE customers only).
typeHealthcheckstruct{…}
typeHTTPConfigurationstruct{…}
Parameters specific to an HTTP or HTTPS health check.
typeQueryHealthcheckstruct{…}
typeTCPConfigurationstruct{…}
Parameters specific to TCP health check.
Healthchecks
Previews
Healthchecks.Previews
Methods
client.Healthchecks.Previews.Get(ctx, healthcheckID, query) (*
Healthcheck
, error)
get/zones/{zone_id}/healthchecks/preview/{healthcheck_id}
Fetch a single configured health check preview.
client.Healthchecks.Previews.New(ctx, params) (*
Healthcheck
, error)
post/zones/{zone_id}/healthchecks/preview
Create a new preview health check.
client.Healthchecks.Previews.Delete(ctx, healthcheckID, body) (*
PreviewDeleteResponse
, error)
delete/zones/{zone_id}/healthchecks/preview/{healthcheck_id}
Delete a health check.