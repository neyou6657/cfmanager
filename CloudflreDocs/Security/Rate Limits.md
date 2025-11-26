Rate Limits
RateLimits
Methods
client.RateLimits.List(ctx, params) (*V4PagePaginationArray[
RateLimit
], error)
Deprecated
get/zones/{zone_id}/rate_limits
Deprecated
Rate limiting API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version for full details.
Fetches the rate limits for a zone.
client.RateLimits.Get(ctx, rateLimitID, query) (*
RateLimit
, error)
Deprecated
get/zones/{zone_id}/rate_limits/{rate_limit_id}
Deprecated
Rate limiting API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version for full details.
Fetches the details of a rate limit.
client.RateLimits.New(ctx, params) (*
RateLimit
, error)
Deprecated
post/zones/{zone_id}/rate_limits
Deprecated
Rate limiting API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version for full details.
Creates a new rate limit for a zone. Refer to the object definition for a list of required attributes.
client.RateLimits.Edit(ctx, rateLimitID, params) (*
RateLimit
, error)
Deprecated
put/zones/{zone_id}/rate_limits/{rate_limit_id}
Deprecated
Rate limiting API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version for full details.
Updates an existing rate limit.
client.RateLimits.Delete(ctx, rateLimitID, body) (*
RateLimitDeleteResponse
, error)
Deprecated
delete/zones/{zone_id}/rate_limits/{rate_limit_id}
Deprecated
Rate limiting API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version for full details.
Deletes an existing rate limit.
Domain types
typeActionstring
The action to apply to a matched request. The log action is only available on an Enterprise plan.
typeRateLimitstruct{…}