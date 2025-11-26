Filters
Filters
Methods
client.Filters.List(ctx, params) (*V4PagePaginationArray[
FirewallFilter
], error)
Deprecated
get/zones/{zone_id}/filters
Deprecated
The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Fetches filters in a zone. You can filter the results using several optional parameters.
client.Filters.Get(ctx, filterID, query) (*
FirewallFilter
, error)
Deprecated
get/zones/{zone_id}/filters/{filter_id}
Deprecated
The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Fetches the details of a filter.
client.Filters.New(ctx, params) (*SinglePage[
FirewallFilter
], error)
Deprecated
post/zones/{zone_id}/filters
Deprecated
The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Creates one or more filters.
client.Filters.Update(ctx, filterID, params) (*
FirewallFilter
, error)
Deprecated
put/zones/{zone_id}/filters/{filter_id}
Deprecated
The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Updates an existing filter.
client.Filters.Delete(ctx, filterID, body) (*
FilterDeleteResponse
, error)
Deprecated
delete/zones/{zone_id}/filters/{filter_id}
Deprecated
The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Deletes an existing filter.
client.Filters.BulkUpdate(ctx, params) (*SinglePage[
FirewallFilter
], error)
Deprecated
put/zones/{zone_id}/filters
Deprecated
The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Updates one or more existing filters.
client.Filters.BulkDelete(ctx, params) (*[]
FilterBulkDeleteResponse
, error)
Deprecated
delete/zones/{zone_id}/filters
Deprecated
The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Deletes one or more existing filters.
Domain types
typeFirewallFilterstruct{…}