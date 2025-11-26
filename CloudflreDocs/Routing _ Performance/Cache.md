Cache
Cache
Methods
client.Cache.Purge(ctx, params) (*
CachePurgeResponse
, error)
post/zones/{zone_id}/purge_cache
Purge All Cached Content
Removes ALL files from Cloudflare's cache. All tiers can purge everything.
{"purge_everything": true}

Purge Cached Content by URL
Granularly removes one or more files from Cloudflare's cache by specifying URLs. All tiers can purge by URL.
To purge files with custom cache keys, include the headers used to compute the cache key as in the example. If you have a device type or geo in your cache key, you will need to include the CF-Device-Type or CF-IPCountry headers. If you have lang in your cache key, you will need to include the Accept-Language header.
NB: When including the Origin header, be sure to include the scheme and hostname. The port number can be omitted if it is the default port (80 for http, 443 for https), but must be included otherwise.
Single file purge example with files:
{"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}

Single file purge example with url and header pairs:
{"files": [{url: "http://www.example.com/cat_picture.jpg", headers: { "CF-IPCountry": "US", "CF-Device-Type": "desktop", "Accept-Language": "zh-CN" }}, {url: "http://www.example.com/dog_picture.jpg", headers: { "CF-IPCountry": "EU", "CF-Device-Type": "mobile", "Accept-Language": "en-US" }}]}

Purge Cached Content by Tag, Host or Prefix
Granularly removes one or more files from Cloudflare's cache either by specifying the host, the associated Cache-Tag, or a Prefix.
Flex purge with tags:
{"tags": ["a-cache-tag", "another-cache-tag"]}

Flex purge with hosts:
{"hosts": ["www.example.com", "images.example.com"]}

Flex purge with prefixes:
{"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}

Availability and limits
please refer to purge cache availability and limits documentation page.
Cache
Cache Reserve
Cache.CacheReserve
Methods
client.Cache.CacheReserve.Get(ctx, query) (*
CacheReserveGetResponse
, error)
get/zones/{zone_id}/cache/cache_reserve
Increase cache lifetimes by automatically storing all cacheable files into Cloudflare's persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the developer docs for more information.
client.Cache.CacheReserve.Edit(ctx, params) (*
CacheReserveEditResponse
, error)
patch/zones/{zone_id}/cache/cache_reserve
Increase cache lifetimes by automatically storing all cacheable files into Cloudflare's persistent object storage buckets. Requires Cache Reserve subscription. Note: using Tiered Cache with Cache Reserve is highly recommended to reduce Reserve operations costs. See the developer docs for more information.
client.Cache.CacheReserve.Status(ctx, query) (*
CacheReserveStatusResponse
, error)
get/zones/{zone_id}/cache/cache_reserve_clear
You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.
client.Cache.CacheReserve.Clear(ctx, params) (*
CacheReserveClearResponse
, error)
post/zones/{zone_id}/cache/cache_reserve_clear
You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.
Domain types
typeCacheReservestring
The identifier of the caching setting.
typeCacheReserveClearstring
ID of the zone setting.
typeStatestring
The current state of the Cache Reserve Clear operation.
Cache
Regional Tiered Cache
Cache.RegionalTieredCache
Methods
client.Cache.RegionalTieredCache.Get(ctx, query) (*
RegionalTieredCacheGetResponse
, error)
get/zones/{zone_id}/cache/regional_tiered_cache
Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.
client.Cache.RegionalTieredCache.Edit(ctx, params) (*
RegionalTieredCacheEditResponse
, error)
patch/zones/{zone_id}/cache/regional_tiered_cache
Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.
Domain types
typeRegionalTieredCachestring
The identifier of the caching setting.
Cache
Smart Tiered Cache
Cache.SmartTieredCache
Methods
client.Cache.SmartTieredCache.Get(ctx, query) (*
SmartTieredCacheGetResponse
, error)
get/zones/{zone_id}/cache/tiered_cache_smart_topology_enable
Smart Tiered Cache dynamically selects the single closest upper tier for each of your website’s origins with no configuration required, using our in-house performance and routing data. Cloudflare collects latency data for each request to an origin, and uses the latency data to determine how well any upper-tier data center is connected with an origin. As a result, Cloudflare can select the data center with the lowest latency to be the upper-tier for an origin.
client.Cache.SmartTieredCache.Edit(ctx, params) (*
SmartTieredCacheEditResponse
, error)
patch/zones/{zone_id}/cache/tiered_cache_smart_topology_enable
Smart Tiered Cache dynamically selects the single closest upper tier for each of your website’s origins with no configuration required, using our in-house performance and routing data. Cloudflare collects latency data for each request to an origin, and uses the latency data to determine how well any upper-tier data center is connected with an origin. As a result, Cloudflare can select the data center with the lowest latency to be the upper-tier for an origin.
client.Cache.SmartTieredCache.Delete(ctx, body) (*
SmartTieredCacheDeleteResponse
, error)
delete/zones/{zone_id}/cache/tiered_cache_smart_topology_enable
Smart Tiered Cache dynamically selects the single closest upper tier for each of your website’s origins with no configuration required, using our in-house performance and routing data. Cloudflare collects latency data for each request to an origin, and uses the latency data to determine how well any upper-tier data center is connected with an origin. As a result, Cloudflare can select the data center with the lowest latency to be the upper-tier for an origin.
Cache
Variants
Cache.Variants
Methods
client.Cache.Variants.Get(ctx, query) (*
VariantGetResponse
, error)
get/zones/{zone_id}/cache/variants
Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.
client.Cache.Variants.Edit(ctx, params) (*
VariantEditResponse
, error)
patch/zones/{zone_id}/cache/variants
Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.
client.Cache.Variants.Delete(ctx, body) (*
VariantDeleteResponse
, error)
delete/zones/{zone_id}/cache/variants
Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.
Domain types
typeCacheVariantstruct{…}
Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.