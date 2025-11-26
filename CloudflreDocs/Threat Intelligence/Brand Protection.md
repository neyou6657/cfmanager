Brand Protection
BrandProtection
Methods
client.BrandProtection.Submit(ctx, body) (*
BrandProtectionSubmitResponse
, error)
post/accounts/{account_id}/brand-protection/submit
Return new URL submissions
client.BrandProtection.URLInfo(ctx, query) (*SinglePage[
BrandProtectionURLInfoResponse
], error)
get/accounts/{account_id}/brand-protection/url-info
Return submitted URLs based on ID
Domain types
typeInfostruct{…}
typeSubmitstruct{…}
Brand Protection
Logo Matches
BrandProtection.LogoMatches
Methods
client.BrandProtection.LogoMatches.Get(ctx, params) (*
LogoMatchGetResponse
, error)
get/accounts/{account_id}/brand-protection/logo-matches
Return matches for logo queries based on ID
client.BrandProtection.LogoMatches.Download(ctx, params) (*
LogoMatchDownloadResponse
, error)
get/accounts/{account_id}/brand-protection/logo-matches/download
Return matches as CSV for logo queries based on ID
Brand Protection
Logos
BrandProtection.Logos
Methods
client.BrandProtection.Logos.New(ctx, params) (*
LogoNewResponse
, error)
post/accounts/{account_id}/brand-protection/logos
Return new saved logo queries created from image files
client.BrandProtection.Logos.Delete(ctx, logoID, body) error
delete/accounts/{account_id}/brand-protection/logos/{logo_id}
Return a success message after deleting saved logo queries by ID
Brand Protection
Matches
BrandProtection.Matches
Methods
client.BrandProtection.Matches.Get(ctx, params) (*
MatchGetResponse
, error)
get/accounts/{account_id}/brand-protection/matches
Return matches for string queries based on ID
client.BrandProtection.Matches.Download(ctx, params) (*
MatchDownloadResponse
, error)
get/accounts/{account_id}/brand-protection/matches/download
Return matches as CSV for string queries based on ID
Brand Protection
Queries
BrandProtection.Queries
Methods
client.BrandProtection.Queries.New(ctx, params) error
post/accounts/{account_id}/brand-protection/queries
Return a success message after creating new saved string queries
client.BrandProtection.Queries.Delete(ctx, params) error
delete/accounts/{account_id}/brand-protection/queries
Return a success message after deleting saved string queries by ID
client.BrandProtection.Queries.Bulk(ctx, params) error
post/accounts/{account_id}/brand-protection/queries/bulk
Return a success message after creating new saved string queries in bulk