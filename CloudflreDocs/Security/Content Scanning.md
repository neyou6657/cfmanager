Content Scanning
ContentScanning
Methods
client.ContentScanning.Enable(ctx, body) (*
ContentScanningEnableResponse
, error)
post/zones/{zone_id}/content-upload-scan/enable
Enable Content Scanning.
client.ContentScanning.Disable(ctx, body) (*
ContentScanningDisableResponse
, error)
post/zones/{zone_id}/content-upload-scan/disable
Disable Content Scanning.
client.ContentScanning.New(ctx, params) (*
ContentScanningNewResponse
, error)
put/zones/{zone_id}/content-upload-scan/settings
Update the Content Scanning status.
client.ContentScanning.Update(ctx, params) (*
ContentScanningUpdateResponse
, error)
put/zones/{zone_id}/content-upload-scan/settings
Update the Content Scanning status.
client.ContentScanning.Get(ctx, query) (*
ContentScanningGetResponse
, error)
get/zones/{zone_id}/content-upload-scan/settings
Retrieve the current status of Content Scanning.
Content Scanning
Payloads
ContentScanning.Payloads
Methods
client.ContentScanning.Payloads.List(ctx, query) (*SinglePage[
PayloadListResponse
], error)
get/zones/{zone_id}/content-upload-scan/payloads
Get a list of existing custom scan expressions for Content Scanning.
client.ContentScanning.Payloads.New(ctx, params) (*SinglePage[
PayloadNewResponse
], error)
post/zones/{zone_id}/content-upload-scan/payloads
Add custom scan expressions for Content Scanning.
client.ContentScanning.Payloads.Delete(ctx, expressionID, body) (*SinglePage[
PayloadDeleteResponse
], error)
delete/zones/{zone_id}/content-upload-scan/payloads/{expression_id}
Delete a Content Scan Custom Expression.
Content Scanning
Settings
ContentScanning.Settings
Methods
client.ContentScanning.Settings.Get(ctx, query) (*
SettingGetResponse
, error)
get/zones/{zone_id}/content-upload-scan/settings
Retrieve the current status of Content Scanning.