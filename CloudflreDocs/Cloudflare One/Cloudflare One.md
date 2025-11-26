SSL
SSL
SSL
Analyze
SSL.Analyze
Methods
client.SSL.Analyze.New(ctx, params) (*
AnalyzeNewResponse
, error)
post/zones/{zone_id}/ssl/analyze
Returns the set of hostnames, the signature algorithm, and the expiration date of the certificate.
SSL
Certificate Packs
SSL.CertificatePacks
Methods
client.SSL.CertificatePacks.List(ctx, params) (*SinglePage[
CertificatePackListResponse
], error)
get/zones/{zone_id}/ssl/certificate_packs
For a given zone, list all active certificate packs.
client.SSL.CertificatePacks.Get(ctx, certificatePackID, query) (*
CertificatePackGetResponse
, error)
get/zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}
For a given zone, get a certificate pack.
client.SSL.CertificatePacks.New(ctx, params) (*
CertificatePackNewResponse
, error)
post/zones/{zone_id}/ssl/certificate_packs/order
For a given zone, order an advanced certificate pack.
client.SSL.CertificatePacks.Edit(ctx, certificatePackID, params) (*
CertificatePackEditResponse
, error)
patch/zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}
For a given zone, restart validation or add cloudflare branding for an advanced certificate pack. The former is only a validation operation for a Certificate Pack in a validation_timed_out status.
client.SSL.CertificatePacks.Delete(ctx, certificatePackID, body) (*
CertificatePackDeleteResponse
, error)
delete/zones/{zone_id}/ssl/certificate_packs/{certificate_pack_id}
For a given zone, delete an advanced certificate pack.
Domain types
typeHoststring
typeRequestValidityfloat64
The number of days for which the certificate should be valid.
typeStatusstring
Status of certificate pack.
typeValidationMethodstring
Validation method in use for a certificate pack order.
SSL
Certificate Packs
Quota
SSL.CertificatePacks.Quota
Methods
client.SSL.CertificatePacks.Quota.Get(ctx, query) (*
CertificatePackQuotaGetResponse
, error)
get/zones/{zone_id}/ssl/certificate_packs/quota
For a given zone, list certificate pack quotas.
SSL
Recommendations
SSL.Recommendations
Methods
client.SSL.Recommendations.Get(ctx, query) (*
RecommendationGetResponse
, error)
Deprecated
get/zones/{zone_id}/ssl/recommendation
Deprecated
SSL/TLS Recommender has been decommissioned in favor of Automatic SSL/TLS
Retrieve the SSL/TLS Recommender's recommendation for a zone.
SSL
Universal
SSL.Universal
SSL
Universal
Settings
SSL.Universal.Settings
Methods
client.SSL.Universal.Settings.Get(ctx, query) (*
UniversalSSLSettings
, error)
get/zones/{zone_id}/ssl/universal/settings
Get Universal SSL Settings for a Zone.
client.SSL.Universal.Settings.Edit(ctx, params) (*
UniversalSSLSettings
, error)
patch/zones/{zone_id}/ssl/universal/settings
Patch Universal SSL Settings for a Zone.
Domain types
typeUniversalSSLSettingsstruct{…}
SSL
Verification
SSL.Verification
Methods
client.SSL.Verification.Get(ctx, params) (*[]
Verification
, error)
get/zones/{zone_id}/ssl/verification
Get SSL Verification Info for a Zone.
client.SSL.Verification.Edit(ctx, certificatePackID, params) (*
VerificationEditResponse
, error)
patch/zones/{zone_id}/ssl/verification/{certificate_pack_id}
Edit SSL validation method for a certificate pack. A PATCH request will request an immediate validation check on any certificate, and return the updated status. If a validation method is provided, the validation will be immediately attempted using that method.
Domain types
typeVerificationstruct{…}