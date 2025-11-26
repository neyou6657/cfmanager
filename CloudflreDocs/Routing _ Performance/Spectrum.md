Spectrum
Spectrum
Domain types
typeDNSstruct{…}
The name and type of DNS record for the Spectrum application.
typeEdgeIPsinterface{…}
The anycast edge IP configuration for the hostname of this application.
typeOriginDNSstruct{…}
The name and type of DNS record for the Spectrum application.
typeOriginPortUnioninterface{…}
The destination port at the origin. Only specified in conjunction with origin_dns. May use an integer to specify a single origin port, for example 1000, or a string to specify a range of origin ports, for example "1000-2000". Notes: If specifying a port range, the number of ports in the range must match the number of ports specified in the "protocol" field.
Spectrum
Analytics
Spectrum.Analytics
Spectrum
Analytics
Aggregates
Spectrum.Analytics.Aggregates
Spectrum
Analytics
Aggregates
Currents
Spectrum.Analytics.Aggregates.Currents
Methods
client.Spectrum.Analytics.Aggregates.Currents.Get(ctx, params) (*[]
AnalyticsAggregateCurrentGetResponse
, error)
get/zones/{zone_id}/spectrum/analytics/aggregate/current
Retrieves analytics aggregated from the last minute of usage on Spectrum applications underneath a given zone.
Spectrum
Analytics
Events
Spectrum.Analytics.Events
Domain types
typeDimensionstring
Spectrum
Analytics
Events
Bytimes
Spectrum.Analytics.Events.Bytimes
Methods
client.Spectrum.Analytics.Events.Bytimes.Get(ctx, params) (*
AnalyticsEventBytimeGetResponse
, error)
get/zones/{zone_id}/spectrum/analytics/events/bytime
Retrieves a list of aggregate metrics grouped by time interval.
Spectrum
Analytics
Events
Summaries
Spectrum.Analytics.Events.Summaries
Methods
client.Spectrum.Analytics.Events.Summaries.Get(ctx, params) (*
AnalyticsEventSummaryGetResponse
, error)
get/zones/{zone_id}/spectrum/analytics/events/summary
Retrieves a list of summarised aggregate metrics over a given time period.
Spectrum
Apps
Spectrum.Apps
Methods
client.Spectrum.Apps.List(ctx, params) (*V4PagePaginationArray[
AppListResponse
], error)
get/zones/{zone_id}/spectrum/apps
Retrieves a list of currently existing Spectrum applications inside a zone.
client.Spectrum.Apps.Get(ctx, appID, query) (*
AppGetResponse
, error)
get/zones/{zone_id}/spectrum/apps/{app_id}
Gets the application configuration of a specific application inside a zone.
client.Spectrum.Apps.New(ctx, params) (*
AppNewResponse
, error)
post/zones/{zone_id}/spectrum/apps
Creates a new Spectrum application from a configuration using a name for the origin.
client.Spectrum.Apps.Update(ctx, appID, params) (*
AppUpdateResponse
, error)
put/zones/{zone_id}/spectrum/apps/{app_id}
Updates a previously existing application's configuration that uses a name for the origin.
client.Spectrum.Apps.Delete(ctx, appID, body) (*
AppDeleteResponse
, error)
delete/zones/{zone_id}/spectrum/apps/{app_id}
Deletes a previously existing application.