Argo
Argo
Argo
Smart Routing
Argo.SmartRouting
Methods
client.Argo.SmartRouting.Get(ctx, query) (*
SmartRoutingGetResponse
, error)
get/zones/{zone_id}/argo/smart_routing
Retrieves the value of Argo Smart Routing enablement setting.
client.Argo.SmartRouting.Edit(ctx, params) (*
SmartRoutingEditResponse
, error)
patch/zones/{zone_id}/argo/smart_routing
Configures the value of the Argo Smart Routing enablement setting.
Argo
Tiered Caching
Argo.TieredCaching
Methods
client.Argo.TieredCaching.Get(ctx, query) (*
TieredCachingGetResponse
, error)
get/zones/{zone_id}/argo/tiered_caching
Tiered Cache works by dividing Cloudflare's data centers into a hierarchy of lower-tiers and upper-tiers. If content is not cached in lower-tier data centers (generally the ones closest to a visitor), the lower-tier must ask an upper-tier to see if it has the content. If the upper-tier does not have the content, only the upper-tier can ask the origin for content. This practice improves bandwidth efficiency by limiting the number of data centers that can ask the origin for content, which reduces origin load and makes websites more cost-effective to operate. Additionally, Tiered Cache concentrates connections to origin servers so they come from a small number of data centers rather than the full set of network locations. This results in fewer open connections using server resources.
client.Argo.TieredCaching.Edit(ctx, params) (*
TieredCachingEditResponse
, error)
patch/zones/{zone_id}/argo/tiered_caching
Tiered Cache works by dividing Cloudflare's data centers into a hierarchy of lower-tiers and upper-tiers. If content is not cached in lower-tier data centers (generally the ones closest to a visitor), the lower-tier must ask an upper-tier to see if it has the content. If the upper-tier does not have the content, only the upper-tier can ask the origin for content. This practice improves bandwidth efficiency by limiting the number of data centers that can ask the origin for content, which reduces origin load and makes websites more cost-effective to operate. Additionally, Tiered Cache concentrates connections to origin servers so they come from a small number of data centers rather than the full set of network locations. This results in fewer open connections using server resources.