Browser Rendering
BrowserRendering
Browser Rendering
Content
BrowserRendering.Content
Methods
client.BrowserRendering.Content.New(ctx, params) (*string, error)
post/accounts/{account_id}/browser-rendering/content
Fetches rendered HTML content from provided URL or HTML. Check available options like gotoOptions and waitFor* to control page load behaviour.
Browser Rendering
Json
BrowserRendering.Json
Methods
client.BrowserRendering.Json.New(ctx, params) (*
JsonNewResponse
, error)
post/accounts/{account_id}/browser-rendering/json
Gets json from a webpage from a provided URL or HTML. Pass prompt or schema in the body. Control page loading with gotoOptions and waitFor* options.
Browser Rendering
Links
BrowserRendering.Links
Methods
client.BrowserRendering.Links.New(ctx, params) (*[]string, error)
post/accounts/{account_id}/browser-rendering/links
Get links from a web page.
Browser Rendering
Markdown
BrowserRendering.Markdown
Methods
client.BrowserRendering.Markdown.New(ctx, params) (*string, error)
post/accounts/{account_id}/browser-rendering/markdown
Gets markdown of a webpage from provided URL or HTML. Control page loading with gotoOptions and waitFor* options.
Browser Rendering
PDF
BrowserRendering.PDF
Methods
client.BrowserRendering.PDF.New(ctx, params) (*
Response
, error)
post/accounts/{account_id}/browser-rendering/pdf
Fetches rendered PDF from provided URL or HTML. Check available options like gotoOptions and waitFor* to control page load behaviour.
Browser Rendering
Scrape
BrowserRendering.Scrape
Methods
client.BrowserRendering.Scrape.New(ctx, params) (*[]
ScrapeNewResponse
, error)
post/accounts/{account_id}/browser-rendering/scrape
Get meta attributes like height, width, text and others of selected elements.
Browser Rendering
Screenshot
BrowserRendering.Screenshot
Methods
client.BrowserRendering.Screenshot.New(ctx, params) (*
ScreenshotNewResponse
, error)
post/accounts/{account_id}/browser-rendering/screenshot
Takes a screenshot of a webpage from provided URL or HTML. Control page loading with gotoOptions and waitFor* options. Customize screenshots with viewport, fullPage, clip and others.
Browser Rendering
Snapshot
BrowserRendering.Snapshot
Methods
client.BrowserRendering.Snapshot.New(ctx, params) (*
SnapshotNewResponse
, error)
post/accounts/{account_id}/browser-rendering/snapshot
Returns the page's HTML content and screenshot. Control page loading with gotoOptions and waitFor* options. Customize screenshots with viewport, fullPage, clip and others.