Client Certificates
ClientCertificates
Methods
client.ClientCertificates.List(ctx, params) (*V4PagePaginationArray[
ClientCertificate
], error)
get/zones/{zone_id}/client_certificates
List all of your Zone's API Shield mTLS Client Certificates by Status and/or using Pagination
client.ClientCertificates.Get(ctx, clientCertificateID, query) (*
ClientCertificate
, error)
get/zones/{zone_id}/client_certificates/{client_certificate_id}
Get Details for a single mTLS API Shield Client Certificate
client.ClientCertificates.New(ctx, params) (*
ClientCertificate
, error)
post/zones/{zone_id}/client_certificates
Create a new API Shield mTLS Client Certificate
client.ClientCertificates.Edit(ctx, clientCertificateID, params) (*
ClientCertificate
, error)
patch/zones/{zone_id}/client_certificates/{client_certificate_id}
If a API Shield mTLS Client Certificate is in a pending_revocation state, you may reactivate it with this endpoint.
client.ClientCertificates.Delete(ctx, clientCertificateID, body) (*
ClientCertificate
, error)
delete/zones/{zone_id}/client_certificates/{client_certificate_id}
Set a API Shield mTLS Client Certificate to pending_revocation status for processing to revoked status.
Domain types
typeClientCertificatestruct{…}