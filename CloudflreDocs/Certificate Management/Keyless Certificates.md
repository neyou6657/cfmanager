Keyless Certificates
KeylessCertificates
Methods
client.KeylessCertificates.List(ctx, query) (*SinglePage[
KeylessCertificate
], error)
get/zones/{zone_id}/keyless_certificates
List all Keyless SSL configurations for a given zone.
client.KeylessCertificates.Get(ctx, keylessCertificateID, query) (*
KeylessCertificate
, error)
get/zones/{zone_id}/keyless_certificates/{keyless_certificate_id}
Get details for one Keyless SSL configuration.
client.KeylessCertificates.New(ctx, params) (*
KeylessCertificate
, error)
post/zones/{zone_id}/keyless_certificates
Create Keyless SSL Configuration
client.KeylessCertificates.Edit(ctx, keylessCertificateID, params) (*
KeylessCertificate
, error)
patch/zones/{zone_id}/keyless_certificates/{keyless_certificate_id}
This will update attributes of a Keyless SSL. Consists of one or more of the following: host,name,port.
client.KeylessCertificates.Delete(ctx, keylessCertificateID, body) (*
KeylessCertificateDeleteResponse
, error)
delete/zones/{zone_id}/keyless_certificates/{keyless_certificate_id}
Delete Keyless SSL Configuration
Domain types
typeKeylessCertificatestruct{…}
typeTunnelstruct{…}
Configuration for using Keyless SSL through a Cloudflare Tunnel