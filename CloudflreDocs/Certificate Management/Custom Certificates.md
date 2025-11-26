Custom Certificates
CustomCertificates
Methods
client.CustomCertificates.List(ctx, params) (*V4PagePaginationArray[
CustomCertificate
], error)
get/zones/{zone_id}/custom_certificates
List, search, and filter all of your custom SSL certificates. The higher priority will break ties across overlapping 'legacy_custom' certificates, but 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
client.CustomCertificates.Get(ctx, customCertificateID, query) (*
CustomCertificate
, error)
get/zones/{zone_id}/custom_certificates/{custom_certificate_id}
SSL Configuration Details
client.CustomCertificates.New(ctx, params) (*
CustomCertificate
, error)
post/zones/{zone_id}/custom_certificates
Upload a new SSL certificate for a zone.
client.CustomCertificates.Edit(ctx, customCertificateID, params) (*
CustomCertificate
, error)
patch/zones/{zone_id}/custom_certificates/{custom_certificate_id}
Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing a configuration for sni_custom certificates will result in a new resource id being returned, and the previous one being deleted.
client.CustomCertificates.Delete(ctx, customCertificateID, body) (*
CustomCertificateDeleteResponse
, error)
delete/zones/{zone_id}/custom_certificates/{custom_certificate_id}
Remove a SSL certificate from a zone.
Domain types
typeCustomCertificatestruct{…}
typeGeoRestrictionsstruct{…}
Specify the region where your private key can be held locally for optimal TLS performance. HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Options allow distribution to only to U.S. data centers, only to E.U. data centers, or only to highest security data centers. Default distribution is to all Cloudflare datacenters, for optimal performance.
typeStatusstring
Client Certificates may be active or revoked, and the pending_reactivation or pending_revocation represent in-progress asynchronous transitions
Custom Certificates
Prioritize
CustomCertificates.Prioritize
Methods
client.CustomCertificates.Prioritize.Update(ctx, params) (*SinglePage[
CustomCertificate
], error)
put/zones/{zone_id}/custom_certificates/prioritize
If a zone has multiple SSL certificates, you can set the order in which they should be used during a request. The higher priority will break ties across overlapping 'legacy_custom' certificates.