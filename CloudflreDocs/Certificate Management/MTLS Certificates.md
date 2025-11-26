MTLS Certificates
MTLSCertificates
Methods
client.MTLSCertificates.List(ctx, query) (*SinglePage[
MTLSCertificate
], error)
get/accounts/{account_id}/mtls_certificates
Lists all mTLS certificates.
client.MTLSCertificates.Get(ctx, mtlsCertificateID, query) (*
MTLSCertificate
, error)
get/accounts/{account_id}/mtls_certificates/{mtls_certificate_id}
Fetches a single mTLS certificate.
client.MTLSCertificates.New(ctx, params) (*
MTLSCertificateNewResponse
, error)
post/accounts/{account_id}/mtls_certificates
Upload a certificate that you want to use with mTLS-enabled Cloudflare services.
client.MTLSCertificates.Delete(ctx, mtlsCertificateID, body) (*
MTLSCertificate
, error)
delete/accounts/{account_id}/mtls_certificates/{mtls_certificate_id}
Deletes the mTLS certificate unless the certificate is in use by one or more Cloudflare services.
Domain types
typeMTLSCertificatestruct{…}
MTLS Certificates
Associations
MTLSCertificates.Associations
Methods
client.MTLSCertificates.Associations.Get(ctx, mtlsCertificateID, query) (*SinglePage[
CertificateAsssociation
], error)
get/accounts/{account_id}/mtls_certificates/{mtls_certificate_id}/associations
Lists all active associations between the certificate and Cloudflare services.
Domain types
typeCertificateAsssociationstruct{…}