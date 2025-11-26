Speed
Speed
Domain types
typeLabeledRegionstruct{…}
A test region with a label.
typeLighthouseReportstruct{…}
The Lighthouse report.
typeTrendstruct{…}
Speed
Availabilities
Speed.Availabilities
Methods
client.Speed.Availabilities.List(ctx, query) (*
Availability
, error)
get/zones/{zone_id}/speed_api/availabilities
Retrieves quota for all plans, as well as the current zone quota.
Domain types
typeAvailabilitystruct{…}
Speed
Pages
Speed.Pages
Methods
client.Speed.Pages.List(ctx, query) (*SinglePage[
PageListResponse
], error)
get/zones/{zone_id}/speed_api/pages
Lists all webpages which have been tested.
client.Speed.Pages.Trend(ctx, url, params) (*
Trend
, error)
get/zones/{zone_id}/speed_api/pages/{url}/trend
Lists the core web vital metrics trend over time for a specific page.
Speed
Pages
Tests
Speed.Pages.Tests
Methods
client.Speed.Pages.Tests.List(ctx, url, params) (*V4PagePaginationArray[
Test
], error)
get/zones/{zone_id}/speed_api/pages/{url}/tests
Test history (list of tests) for a specific webpage.
client.Speed.Pages.Tests.Get(ctx, url, testID, query) (*
Test
, error)
get/zones/{zone_id}/speed_api/pages/{url}/tests/{test_id}
Retrieves the result of a specific test.
client.Speed.Pages.Tests.New(ctx, url, params) (*
Test
, error)
post/zones/{zone_id}/speed_api/pages/{url}/tests
Starts a test for a specific webpage, in a specific region.
client.Speed.Pages.Tests.Delete(ctx, url, params) (*
PageTestDeleteResponse
, error)
delete/zones/{zone_id}/speed_api/pages/{url}/tests
Deletes all tests for a specific webpage from a specific region. Deleted tests are still counted as part of the quota.
Domain types
typeTeststruct{…}
Speed
Schedule
Speed.Schedule
Methods
client.Speed.Schedule.Get(ctx, url, params) (*
Schedule
, error)
get/zones/{zone_id}/speed_api/schedule/{url}
Retrieves the test schedule for a page in a specific region.
client.Speed.Schedule.New(ctx, url, params) (*
ScheduleNewResponse
, error)
post/zones/{zone_id}/speed_api/schedule/{url}
Creates a scheduled test for a page.
client.Speed.Schedule.Delete(ctx, url, params) (*
ScheduleDeleteResponse
, error)
delete/zones/{zone_id}/speed_api/schedule/{url}
Deletes a scheduled test for a page.
Domain types
typeSchedulestruct{…}
The test schedule.