Origin TLS Client Auth
OriginTLSClientAuth
Methods
client.OriginTLSClientAuth.List(ctx, query) (*SinglePage[
OriginTLSClientAuthListResponse
], error)
get/zones/{zone_id}/origin_tls_client_auth
List Certificates
client.OriginTLSClientAuth.Get(ctx, certificateID, query) (*
OriginTLSClientAuthGetResponse
, error)
get/zones/{zone_id}/origin_tls_client_auth/{certificate_id}
Get Certificate Details
client.OriginTLSClientAuth.New(ctx, params) (*
OriginTLSClientAuthNewResponse
, error)
post/zones/{zone_id}/origin_tls_client_auth
Upload your own certificate you want Cloudflare to use for edge-to-origin communication to override the shared certificate. Please note that it is important to keep only one certificate active. Also, make sure to enable zone-level authenticated origin pulls by making a PUT call to settings endpoint to see the uploaded certificate in use.
client.OriginTLSClientAuth.Delete(ctx, certificateID, body) (*
OriginTLSClientAuthDeleteResponse
, error)
delete/zones/{zone_id}/origin_tls_client_auth/{certificate_id}
Delete Certificate
Domain types
typeZoneAuthenticatedOriginPullstruct{…}
Origin TLS Client Auth
Hostnames
OriginTLSClientAuth.Hostnames
Methods
client.OriginTLSClientAuth.Hostnames.Get(ctx, hostname, query) (*
AuthenticatedOriginPull
, error)
get/zones/{zone_id}/origin_tls_client_auth/hostnames/{hostname}
Get the Hostname Status for Client Authentication
client.OriginTLSClientAuth.Hostnames.Update(ctx, params) (*SinglePage[
HostnameUpdateResponse
], error)
put/zones/{zone_id}/origin_tls_client_auth/hostnames
Associate a hostname to a certificate and enable, disable or invalidate the association. If disabled, client certificate will not be sent to the hostname even if activated at the zone level. 100 maximum associations on a single certificate are allowed. Note: Use a null value for parameter enabled to invalidate the association.
Domain types
typeAuthenticatedOriginPullstruct{…}
Origin TLS Client Auth
Hostnames
Certificates
OriginTLSClientAuth.Hostnames.Certificates
Methods
client.OriginTLSClientAuth.Hostnames.Certificates.List(ctx, query) (*SinglePage[
HostnameCertificateListResponse
], error)
get/zones/{zone_id}/origin_tls_client_auth/hostnames/certificates
List Certificates
client.OriginTLSClientAuth.Hostnames.Certificates.Get(ctx, certificateID, query) (*
HostnameCertificateGetResponse
, error)
get/zones/{zone_id}/origin_tls_client_auth/hostnames/certificates/{certificate_id}
Get the certificate by ID to be used for client authentication on a hostname.
client.OriginTLSClientAuth.Hostnames.Certificates.New(ctx, params) (*
HostnameCertificateNewResponse
, error)
post/zones/{zone_id}/origin_tls_client_auth/hostnames/certificates
Upload a certificate to be used for client authentication on a hostname. 10 hostname certificates per zone are allowed.
client.OriginTLSClientAuth.Hostnames.Certificates.Delete(ctx, certificateID, body) (*
HostnameCertificateDeleteResponse
, error)
delete/zones/{zone_id}/origin_tls_client_auth/hostnames/certificates/{certificate_id}
Delete Hostname Client Certificate
Domain types
typeCertificatestruct{…}
Origin TLS Client Auth
Settings
OriginTLSClientAuth.Settings
Methods
client.OriginTLSClientAuth.Settings.Get(ctx, query) (*
SettingGetResponse
, error)
get/zones/{zone_id}/origin_tls_client_auth/settings
Get whether zone-level authenticated origin pulls is enabled or not. It is false by default.
client.OriginTLSClientAuth.Settings.Update(ctx, params) (*
SettingUpdateResponse
, error)
put/zones/{zone_id}/origin_tls_client_auth/settings
Enable or disable zone-level authenticated origin pulls. 'enabled' should be set true either before/after the certificate is uploaded to see the certificate in use.