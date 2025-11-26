Load Balancers
LoadBalancers
Methods
client.LoadBalancers.List(ctx, query) (*SinglePage[
LoadBalancer
], error)
get/zones/{zone_id}/load_balancers
List configured load balancers.
client.LoadBalancers.Get(ctx, loadBalancerID, query) (*
LoadBalancer
, error)
get/zones/{zone_id}/load_balancers/{load_balancer_id}
Fetch a single configured load balancer.
client.LoadBalancers.New(ctx, params) (*
LoadBalancer
, error)
post/zones/{zone_id}/load_balancers
Create a new load balancer.
client.LoadBalancers.Update(ctx, loadBalancerID, params) (*
LoadBalancer
, error)
put/zones/{zone_id}/load_balancers/{load_balancer_id}
Update a configured load balancer.
client.LoadBalancers.Edit(ctx, loadBalancerID, params) (*
LoadBalancer
, error)
patch/zones/{zone_id}/load_balancers/{load_balancer_id}
Apply changes to an existing load balancer, overwriting the supplied properties.
client.LoadBalancers.Delete(ctx, loadBalancerID, body) (*
LoadBalancerDeleteResponse
, error)
delete/zones/{zone_id}/load_balancers/{load_balancer_id}
Delete a configured load balancer.
Domain types
typeAdaptiveRoutingstruct{…}
Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests. For example, zero-downtime failover occurs immediately when an origin becomes unavailable due to HTTP 521, 522, or 523 response codes. If there is another healthy origin in the same pool, the request is retried once against this alternate origin.
typeCheckRegionstring
WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS: Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (ENTERPRISE customers only).
typeDefaultPoolsstring
A pool ID.
typeFilterOptionsstruct{…}
Filter options for a particular resource type (pool or origin). Use null to reset.
typeHeaderstruct{…}
The request header is used to pass additional information with an HTTP request. Currently supported header is 'Host'.
typeHoststring
typeLoadBalancerstruct{…}
typeLoadSheddingstruct{…}
Configures load shedding policies and percentages for the pool.
typeLocationStrategystruct{…}
Controls location-based steering for non-proxied requests. See steering_policy to learn how steering is affected.
typeNotificationFilterstruct{…}
Filter pool and origin health notifications by resource type or health status. Use null to reset.
typeOriginstruct{…}
typeOriginSteeringstruct{…}
Configures origin steering for the pool. Controls how origins are selected for new sessions and traffic without session affinity.
typeRandomSteeringstruct{…}
Configures pool weights.
steering_policy="random": A random pool is selected with probability proportional to pool weights.
steering_policy="least_outstanding_requests": Use pool weights to scale each pool's outstanding requests.
steering_policy="least_connections": Use pool weights to scale each pool's open connections.
typeRulesstruct{…}
A rule object containing conditions and overrides for this load balancer to evaluate.
typeSessionAffinitystring
Specifies the type of session affinity the load balancer should use unless specified as "none". The supported types are: - "cookie": On the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy, then a new origin server is calculated and used. - "ip_cookie": Behaves the same as "cookie" except the initial origin selection is stable and based on the client's ip address. - "header": On the first request to a proxied load balancer, a session key based on the configured HTTP headers (see session_affinity_attributes.headers) is generated, encoding the request headers used for storing in the load balancer session state which origin the request will be forwarded to. Subsequent requests to the load balancer with the same headers will be sent to the same origin server, for the duration of the session and as long as the origin server remains healthy. If the session has been idle for the duration of session_affinity_ttl seconds or the origin server is unhealthy, then a new origin server is calculated and used. See headers in session_affinity_attributes for additional required configuration.
typeSessionAffinityAttributesstruct{…}
Configures attributes for session affinity.
typeSteeringPolicystring
Steering Policy for this load balancer.
"off": Use default_pools.
"geo": Use region_pools/country_pools/pop_pools. For non-proxied requests, the country for country_pools is determined by location_strategy.
"random": Select a pool randomly.
"dynamic_latency": Use round trip time to select the closest pool in default_pools (requires pool health checks).
"proximity": Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by location_strategy for non-proxied requests.
"least_outstanding_requests": Select a pool by taking into consideration random_steering weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
"least_connections": Select a pool by taking into consideration random_steering weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
"": Will map to "geo" if you use region_pools/country_pools/pop_pools otherwise "off".
Load Balancers
Monitor Groups
LoadBalancers.MonitorGroups
Methods
client.LoadBalancers.MonitorGroups.List(ctx, query) (*SinglePage[
MonitorGroup
], error)
get/accounts/{account_id}/load_balancers/monitor_groups
List configured monitor groups.
client.LoadBalancers.MonitorGroups.Get(ctx, monitorGroupID, query) (*
MonitorGroup
, error)
get/accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
Fetch a single configured monitor group.
client.LoadBalancers.MonitorGroups.New(ctx, params) (*
MonitorGroup
, error)
post/accounts/{account_id}/load_balancers/monitor_groups
Create a new monitor group.
client.LoadBalancers.MonitorGroups.Update(ctx, monitorGroupID, params) (*
MonitorGroup
, error)
put/accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
Modify a configured monitor group.
client.LoadBalancers.MonitorGroups.Edit(ctx, monitorGroupID, params) (*
MonitorGroup
, error)
patch/accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
Apply changes to an existing monitor group, overwriting the supplied properties.
client.LoadBalancers.MonitorGroups.Delete(ctx, monitorGroupID, body) (*
MonitorGroup
, error)
delete/accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
Delete a configured monitor group.
Domain types
typeMonitorGroupstruct{…}
Load Balancers
Monitors
LoadBalancers.Monitors
Methods
client.LoadBalancers.Monitors.List(ctx, query) (*SinglePage[
Monitor
], error)
get/accounts/{account_id}/load_balancers/monitors
List configured monitors for an account.
client.LoadBalancers.Monitors.Get(ctx, monitorID, query) (*
Monitor
, error)
get/accounts/{account_id}/load_balancers/monitors/{monitor_id}
List a single configured monitor for an account.
client.LoadBalancers.Monitors.New(ctx, params) (*
Monitor
, error)
post/accounts/{account_id}/load_balancers/monitors
Create a configured monitor.
client.LoadBalancers.Monitors.Update(ctx, monitorID, params) (*
Monitor
, error)
put/accounts/{account_id}/load_balancers/monitors/{monitor_id}
Modify a configured monitor.
client.LoadBalancers.Monitors.Edit(ctx, monitorID, params) (*
Monitor
, error)
patch/accounts/{account_id}/load_balancers/monitors/{monitor_id}
Apply changes to an existing monitor, overwriting the supplied properties.
client.LoadBalancers.Monitors.Delete(ctx, monitorID, body) (*
MonitorDeleteResponse
, error)
delete/accounts/{account_id}/load_balancers/monitors/{monitor_id}
Delete a configured monitor.
Domain types
typeMonitorstruct{…}
Load Balancers
Monitors
Previews
LoadBalancers.Monitors.Previews
Methods
client.LoadBalancers.Monitors.Previews.New(ctx, monitorID, params) (*
MonitorPreviewNewResponse
, error)
post/accounts/{account_id}/load_balancers/monitors/{monitor_id}/preview
Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.
Load Balancers
Monitors
References
LoadBalancers.Monitors.References
Methods
client.LoadBalancers.Monitors.References.Get(ctx, monitorID, query) (*SinglePage[
MonitorReferenceGetResponse
], error)
get/accounts/{account_id}/load_balancers/monitors/{monitor_id}/references
Get the list of resources that reference the provided monitor.
Load Balancers
Pools
LoadBalancers.Pools
Methods
client.LoadBalancers.Pools.List(ctx, params) (*SinglePage[
Pool
], error)
get/accounts/{account_id}/load_balancers/pools
List configured pools.
client.LoadBalancers.Pools.Get(ctx, poolID, query) (*
Pool
, error)
get/accounts/{account_id}/load_balancers/pools/{pool_id}
Fetch a single configured pool.
client.LoadBalancers.Pools.New(ctx, params) (*
Pool
, error)
post/accounts/{account_id}/load_balancers/pools
Create a new pool.
client.LoadBalancers.Pools.Update(ctx, poolID, params) (*
Pool
, error)
put/accounts/{account_id}/load_balancers/pools/{pool_id}
Modify a configured pool.
client.LoadBalancers.Pools.Edit(ctx, poolID, params) (*
Pool
, error)
patch/accounts/{account_id}/load_balancers/pools/{pool_id}
Apply changes to an existing pool, overwriting the supplied properties.
client.LoadBalancers.Pools.Delete(ctx, poolID, body) (*
PoolDeleteResponse
, error)
delete/accounts/{account_id}/load_balancers/pools/{pool_id}
Delete a configured pool.
client.LoadBalancers.Pools.BulkEdit(ctx, params) (*SinglePage[
Pool
], error)
patch/accounts/{account_id}/load_balancers/pools
Apply changes to a number of existing pools, overwriting the supplied properties. Pools are ordered by ascending name. Returns the list of affected pools. Supports the standard pagination query parameters, either limit/offset or per_page/page.
Domain types
typePoolstruct{…}
Load Balancers
Pools
Health
LoadBalancers.Pools.Health
Methods
client.LoadBalancers.Pools.Health.Get(ctx, poolID, query) (*
PoolHealthGetResponse
, error)
get/accounts/{account_id}/load_balancers/pools/{pool_id}/health
Fetch the latest pool health status for a single pool.
client.LoadBalancers.Pools.Health.New(ctx, poolID, params) (*
PoolHealthNewResponse
, error)
post/accounts/{account_id}/load_balancers/pools/{pool_id}/preview
Preview pool health using provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.
Load Balancers
Pools
References
LoadBalancers.Pools.References
Methods
client.LoadBalancers.Pools.References.Get(ctx, poolID, query) (*SinglePage[
PoolReferenceGetResponse
], error)
get/accounts/{account_id}/load_balancers/pools/{pool_id}/references
Get the list of resources that reference the provided pool.
Load Balancers
Previews
LoadBalancers.Previews
Methods
client.LoadBalancers.Previews.Get(ctx, previewID, query) (*
PreviewGetResponse
, error)
get/accounts/{account_id}/load_balancers/preview/{preview_id}
Get the result of a previous preview operation using the provided preview_id.
Load Balancers
Regions
LoadBalancers.Regions
Methods
client.LoadBalancers.Regions.List(ctx, params) (*unknown, error)
get/accounts/{account_id}/load_balancers/regions
List all region mappings.
client.LoadBalancers.Regions.Get(ctx, regionID, query) (*unknown, error)
get/accounts/{account_id}/load_balancers/regions/{region_id}
Get a single region mapping.
Load Balancers
Searches
LoadBalancers.Searches
Methods
client.LoadBalancers.Searches.List(ctx, params) (*V4PagePagination[
SearchListResponse
], error)
get/accounts/{account_id}/load_balancers/search
Search for Load Balancing resources.