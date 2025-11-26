Origin CA Certificates
OriginCACertificates
Methods
client.OriginCACertificates.List(ctx, query) (*V4PagePaginationArray[
OriginCACertificate
], error)
get/certificates
List all existing Origin CA certificates for a given zone. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint (see above).
client.OriginCACertificates.Get(ctx, certificateID) (*
OriginCACertificate
, error)
get/certificates/{certificate_id}
Get an existing Origin CA certificate by its serial number. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint (see above).
client.OriginCACertificates.New(ctx, body) (*
OriginCACertificate
, error)
post/certificates
Create an Origin CA certificate. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint (see above).
client.OriginCACertificates.Delete(ctx, certificateID) (*
OriginCACertificateDeleteResponse
, error)
delete/certificates/{certificate_id}
Revoke an existing Origin CA certificate by its serial number. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint (see above).
Domain types
typeOriginCACertificatestruct{…}