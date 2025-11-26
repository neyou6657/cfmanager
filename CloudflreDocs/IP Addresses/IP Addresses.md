Zones
Zones
Methods
client.Zones.List(ctx, query) (*V4PagePaginationArray[
Zone
], error)
get/zones
Lists, searches, sorts, and filters your zones. Listing zones across more than 500 accounts is currently not allowed.
client.Zones.Get(ctx, query) (*
Zone
, error)
get/zones/{zone_id}
Zone Details
client.Zones.New(ctx, body) (*
Zone
, error)
post/zones
Create Zone
client.Zones.Edit(ctx, params) (*
Zone
, error)
patch/zones/{zone_id}
Edits a zone. Only one zone property can be changed at a time.
client.Zones.Delete(ctx, body) (*
ZoneDeleteResponse
, error)
delete/zones/{zone_id}
Deletes an existing zone.
Domain types
typeTypestring
A full zone implies that DNS is hosted with Cloudflare. A partial zone is typically a partner-hosted zone or a CNAME setup.
typeZonestruct{…}
Zones
Activation Check
Zones.ActivationCheck
Methods
client.Zones.ActivationCheck.Trigger(ctx, body) (*
ActivationCheckTriggerResponse
, error)
put/zones/{zone_id}/activation_check
Triggeres a new activation check for a PENDING Zone. This can be triggered every 5 min for paygo/ent customers, every hour for FREE Zones.
Zones
Custom Nameservers
Zones.CustomNameservers
Methods
client.Zones.CustomNameservers.Get(ctx, query) (*
CustomNameserverGetResponse
, error)
Deprecated
get/zones/{zone_id}/custom_ns
Deprecated
Use DNS settings API instead.
Get metadata for account-level custom nameservers on a zone.
Deprecated in favor of Show DNS Settings.
client.Zones.CustomNameservers.Update(ctx, params) (*SinglePage[string], error)
Deprecated
put/zones/{zone_id}/custom_ns
Deprecated
Use DNS settings API instead.
Set metadata for account-level custom nameservers on a zone.
If you would like new zones in the account to use account custom nameservers by default, use PUT /accounts/:identifier to set the account setting use_account_custom_ns_by_default to true.
Deprecated in favor of Update DNS Settings.
Zones
Holds
Zones.Holds
Methods
client.Zones.Holds.Get(ctx, query) (*
ZoneHold
, error)
get/zones/{zone_id}/hold
Retrieve whether the zone is subject to a zone hold, and metadata about the hold.
client.Zones.Holds.New(ctx, params) (*
ZoneHold
, error)
post/zones/{zone_id}/hold
Enforce a zone hold on the zone, blocking the creation and activation of zones with this zone's hostname.
client.Zones.Holds.Edit(ctx, params) (*
ZoneHold
, error)
patch/zones/{zone_id}/hold
Update the hold_after and/or include_subdomains values on an existing zone hold. The hold is enabled if the hold_after date-time value is in the past.
client.Zones.Holds.Delete(ctx, params) (*
ZoneHold
, error)
delete/zones/{zone_id}/hold
Stop enforcement of a zone hold on the zone, permanently or temporarily, allowing the creation and activation of zones with this zone's hostname.
Domain types
typeZoneHoldstruct{…}
Zones
Plans
Zones.Plans
Methods
client.Zones.Plans.List(ctx, query) (*SinglePage[
AvailableRatePlan
], error)
get/zones/{zone_id}/available_plans
Lists available plans the zone can subscribe to.
client.Zones.Plans.Get(ctx, planIdentifier, query) (*
AvailableRatePlan
, error)
get/zones/{zone_id}/available_plans/{plan_identifier}
Details of the available plan that the zone can subscribe to.
Domain types
typeAvailableRatePlanstruct{…}
Zones
Rate Plans
Zones.RatePlans
Methods
client.Zones.RatePlans.Get(ctx, query) (*SinglePage[
RatePlanGetResponse
], error)
get/zones/{zone_id}/available_rate_plans
Lists all rate plans the zone can subscribe to.
Zones
Settings
Zones.Settings
Methods
client.Zones.Settings.Get(ctx, settingID, query) (*
SettingGetResponse
, error)
get/zones/{zone_id}/settings/{setting_id}
Fetch a single zone setting by name
client.Zones.Settings.Edit(ctx, settingID, params) (*
SettingEditResponse
, error)
patch/zones/{zone_id}/settings/{setting_id}
Updates a single zone setting by the identifier
Domain types
typeAdvancedDDoSstruct{…}
Advanced protection from Distributed Denial of Service (DDoS) attacks on your website. This is an uneditable value that is 'on' in the case of Business and Enterprise zones.
typeAegisstruct{…}
Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your account so that you can increase your origin security by only allowing traffic from a small list of IP addresses.
typeAlwaysOnlinestruct{…}
When enabled, Cloudflare serves limited copies of web pages available from the Internet Archive's Wayback Machine if your server is offline. Refer to Always Online for more information.
typeAlwaysUseHTTPSstruct{…}
typeAutomaticHTTPSRewritesstruct{…}
typeAutomaticPlatformOptimizationstruct{…}
typeBrotlistruct{…}
When the client requesting an asset supports the Brotli compression algorithm, Cloudflare will serve a Brotli compressed version of the asset.
typeBrowserCacheTTLstruct{…}
typeBrowserCheckstruct{…}
typeCacheLevelstruct{…}
typeChallengeTTLstruct{…}
Specify how long a visitor is allowed access to your site after successfully completing a challenge (such as a CAPTCHA). After the TTL has expired the visitor will have to complete a new challenge. We recommend a 15 - 45 minute setting and will attempt to honor any setting above 45 minutes. (https://support.cloudflare.com/hc/en-us/articles/200170136).
typeCiphersstruct{…}
An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.
typeDevelopmentModestruct{…}
Development Mode temporarily allows you to enter development mode for your websites if you need to make changes to your site. This will bypass Cloudflare's accelerated cache and slow down your site, but is useful if you are making changes to cacheable content (like images, css, or JavaScript) and would like to see those changes right away. Once entered, development mode will last for 3 hours and then automatically toggle off.
typeEarlyHintsstruct{…}
When enabled, Cloudflare will attempt to speed up overall page loads by serving 103 responses with Link headers from the final response. Refer to Early Hints for more information.
typeEmailObfuscationstruct{…}
typeFontSettingsstruct{…}
Enhance your website's font delivery with Cloudflare Fonts. Deliver Google Hosted fonts from your own domain, boost performance, and enhance user privacy. Refer to the Cloudflare Fonts documentation for more information.
typeH2Prioritizationstruct{…}
HTTP/2 Edge Prioritization optimises the delivery of resources served through HTTP/2 to improve page load performance. It also supports fine control of content delivery when used in conjunction with Workers.
typeHotlinkProtectionstruct{…}
When enabled, the Hotlink Protection option ensures that other sites cannot suck up your bandwidth by building pages that use images hosted on your site. Anytime a request for an image on your site hits Cloudflare, we check to ensure that it's not another site requesting them. People will still be able to download and view images from your page, but other sites won't be able to steal them for use on their own pages. (https://support.cloudflare.com/hc/en-us/articles/200170026).
typeHTTP2struct{…}
HTTP2 enabled for this zone.
typeHTTP3struct{…}
HTTP3 enabled for this zone.
typeImageResizingstruct{…}
Image Transformations provides on-demand resizing, conversion and optimization for images served through Cloudflare's network. Refer to the Image Transformations documentation for more information.
typeIPGeolocationstruct{…}
typeIPV6struct{…}
Enable IPv6 on all subdomains that are Cloudflare enabled. (https://support.cloudflare.com/hc/en-us/articles/200168586).
typeMinTLSVersionstruct{…}
Only accepts HTTPS requests that use at least the TLS protocol version specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
typeMiragestruct{…}
typeNELstruct{…}
Enable Network Error Logging reporting on your zone. (Beta) 
typeOpportunisticEncryptionstruct{…}
typeOpportunisticOnionstruct{…}
Add an Alt-Svc header to all legitimate requests from Tor, allowing the connection to use our onion services instead of exit nodes.
typeOrangeToOrangestruct{…}
Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also on Cloudflare.
typeOriginErrorPagePassThrustruct{…}
typeOriginMaxHTTPVersionstruct{…}
typePolishstruct{…}
typePrefetchPreloadstruct{…}
Cloudflare will prefetch any URLs that are included in the response headers. This is limited to Enterprise Zones.
typeProxyReadTimeoutstruct{…}
Maximum time between two read operations from origin.
typePseudoIPV4struct{…}
The value set for the Pseudo IPv4 setting.
typeResponseBufferingstruct{…}
typeRocketLoaderstruct{…}
typeSecurityHeadersstruct{…}
Cloudflare security header for a zone.
typeSecurityLevelstruct{…}
typeServerSideExcludesstruct{…}
If there is sensitive content on your website that you want visible to real visitors, but that you want to hide from suspicious visitors, all you have to do is wrap the content with Cloudflare SSE tags. Wrap any content that you want to be excluded from suspicious visitors in the following SSE tags: . For example: Bad visitors won't see my phone number, 555-555-5555 . Note: SSE only will work with HTML. If you have HTML minification enabled, you won't see the SSE tags in your HTML source when it's served through Cloudflare. SSE will still function in this case, as Cloudflare's HTML minification and SSE functionality occur on-the-fly as the resource moves through our network to the visitor's computer. (https://support.cloudflare.com/hc/en-us/articles/200170036).
typeSortQueryStringForCachestruct{…}
typeSSLstruct{…}
typeSSLRecommenderstruct{…}
Enrollment in the SSL/TLS Recommender service which tries to detect and recommend (by sending periodic emails) the most secure SSL/TLS setting your origin servers support.
typeTLS1_3struct{…}
Enables Crypto TLS 1.3 feature for a zone.
typeTLSClientAuthstruct{…}
TLS Client Auth requires Cloudflare to connect to your origin server using a client certificate (Enterprise Only).
typeTrueClientIPHeaderstruct{…}
typeWAFstruct{…}
typeWebPstruct{…}
When the client requesting the image supports the WebP image codec, and WebP offers a performance advantage over the original image format, Cloudflare will serve a WebP version of the original image.
typeWebsocketstruct{…}
WebSockets are open connections sustained between the client and the origin server. Inside a WebSockets connection, the client and the origin can pass data back and forth without having to reestablish sessions. This makes exchanging data within a WebSockets connection fast. WebSockets are often used for real-time applications such as live chat and gaming. For more information refer to Can I use Cloudflare with Websockets.
typeZeroRTTstruct{…}
0-RTT session resumption enabled for this zone.
Zones
Subscriptions
Zones.Subscriptions
Methods
client.Zones.Subscriptions.Get(ctx, query) (*
Subscription
, error)
get/zones/{zone_id}/subscription
Lists zone subscription details.
client.Zones.Subscriptions.New(ctx, params) (*
Subscription
, error)
post/zones/{zone_id}/subscription
Create a zone subscription, either plan or add-ons.
client.Zones.Subscriptions.Update(ctx, params) (*
Subscription
, error)
put/zones/{zone_id}/subscription
Updates zone subscriptions, either plan or add-ons.