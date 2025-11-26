URL Scanner
URLScanner
Domain types
typeURLScannerDomainstruct{…}
typeURLScannerTaskstruct{…}
URL Scanner
Responses
URLScanner.Responses
Methods
client.URLScanner.Responses.Get(ctx, responseID, query) (*string, error)
get/accounts/{account_id}/urlscanner/v2/responses/{response_id}
Returns the raw response of the network request. Find the response_id in the data.requests.response.hash.
URL Scanner
Scans
URLScanner.Scans
Methods
client.URLScanner.Scans.List(ctx, params) (*
ScanListResponse
, error)
get/accounts/{account_id}/urlscanner/v2/search
Use a subset of ElasticSearch Query syntax to filter scans. Some example queries:

- 'path:"/bundles/jquery.js"': Searches for scans who requested resources with the given path.
- 'page.asn:AS24940 AND hash:xxx': Websites hosted in AS24940 where a resource with the given hash was downloaded.
- 'page.domain:microsoft* AND verdicts.malicious:true AND NOT page.domain:microsoft.com': malicious scans whose hostname starts with "microsoft".
- 'apikey:me AND date:[2025-01 TO 2025-02]': my scans from 2025 January to 2025 February.
client.URLScanner.Scans.Get(ctx, scanID, query) (*
ScanGetResponse
, error)
get/accounts/{account_id}/urlscanner/v2/result/{scan_id}
Get URL scan by uuid
client.URLScanner.Scans.New(ctx, params) (*
ScanNewResponse
, error)
post/accounts/{account_id}/urlscanner/v2/scan
Submit a URL to scan. Check limits at https://developers.cloudflare.com/security-center/investigate/scan-limits/.
client.URLScanner.Scans.BulkNew(ctx, params) (*[]
ScanBulkNewResponse
, error)
post/accounts/{account_id}/urlscanner/v2/bulk
Submit URLs to scan. Check limits at https://developers.cloudflare.com/security-center/investigate/scan-limits/ and take into account scans submitted in bulk have lower priority and may take longer to finish.
client.URLScanner.Scans.HAR(ctx, scanID, query) (*
ScanHARResponse
, error)
get/accounts/{account_id}/urlscanner/v2/har/{scan_id}
Get a URL scan's HAR file. See HAR spec at http://www.softwareishard.com/blog/har-12-spec/.
client.URLScanner.Scans.Screenshot(ctx, scanID, params) (*
Response
, error)
get/accounts/{account_id}/urlscanner/v2/screenshots/{scan_id}.png
Get scan's screenshot by resolution (desktop/mobile/tablet).
client.URLScanner.Scans.DOM(ctx, scanID, query) (*string, error)
get/accounts/{account_id}/urlscanner/v2/dom/{scan_id}
Returns a plain text response, with the scan's DOM content as rendered by Chrome.