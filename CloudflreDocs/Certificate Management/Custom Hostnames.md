Custom Hostnames
CustomHostnames
Methods
client.CustomHostnames.List(ctx, params) (*V4PagePaginationArray[
CustomHostnameListResponse
], error)
get/zones/{zone_id}/custom_hostnames
List, search, sort, and filter all of your custom hostnames.
client.CustomHostnames.Get(ctx, customHostnameID, query) (*
CustomHostnameGetResponse
, error)
get/zones/{zone_id}/custom_hostnames/{custom_hostname_id}
Custom Hostname Details
client.CustomHostnames.New(ctx, params) (*
CustomHostnameNewResponse
, error)
post/zones/{zone_id}/custom_hostnames
Add a new custom hostname and request that an SSL certificate be issued for it. One of three validation methods—http, txt, email—should be used, with 'http' recommended if the CNAME is already in place (or will be soon). Specifying 'email' will send an email to the WHOIS contacts on file for the base domain plus hostmaster, postmaster, webmaster, admin, administrator. If http is used and the domain is not already pointing to the Managed CNAME host, the PATCH method must be used once it is (to complete validation). Enable bundling of certificates using the custom_cert_bundle field. The bundling process requires the following condition One certificate in the bundle must use an RSA, and the other must use an ECDSA.
client.CustomHostnames.Edit(ctx, customHostnameID, params) (*
CustomHostnameEditResponse
, error)
patch/zones/{zone_id}/custom_hostnames/{custom_hostname_id}
Modify SSL configuration for a custom hostname. When sent with SSL config that matches existing config, used to indicate that hostname should pass domain control validation (DCV). Can also be used to change validation type, e.g., from 'http' to 'email'. Bundle an existing certificate with another certificate by using the "custom_cert_bundle" field. The bundling process supports combining certificates as long as the following condition is met. One certificate must use the RSA algorithm, and the other must use the ECDSA algorithm.
client.CustomHostnames.Delete(ctx, customHostnameID, body) (*
CustomHostnameDeleteResponse
, error)
delete/zones/{zone_id}/custom_hostnames/{custom_hostname_id}
Delete Custom Hostname (and any issued SSL certificates)
Domain types
typeBundleMethodstring
A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.
typeCustomHostnamestruct{…}
typeDCVMethodstring
Domain control validation (DCV) method used for this hostname.
typeDomainValidationTypestring
Level of validation to be used for this hostname. Domain validation (dv) must be used.
Custom Hostnames
Certificate Pack
CustomHostnames.CertificatePack
Custom Hostnames
Certificate Pack
Certificates
CustomHostnames.CertificatePack.Certificates
Methods
client.CustomHostnames.CertificatePack.Certificates.Update(ctx, customHostnameID, certificatePackID, certificateID, params) (*
CertificatePackCertificateUpdateResponse
, error)
put/zones/{zone_id}/custom_hostnames/{custom_hostname_id}/certificate_pack/{certificate_pack_id}/certificates/{certificate_id}
Replace a single custom certificate within a certificate pack that contains two bundled certificates. The replacement must adhere to the following constraints. You can only replace an RSA certificate with another RSA certificate or an ECDSA certificate with another ECDSA certificate.
client.CustomHostnames.CertificatePack.Certificates.Delete(ctx, customHostnameID, certificatePackID, certificateID, body) (*
CertificatePackCertificateDeleteResponse
, error)
delete/zones/{zone_id}/custom_hostnames/{custom_hostname_id}/certificate_pack/{certificate_pack_id}/certificates/{certificate_id}
Delete a single custom certificate from a certificate pack that contains two bundled certificates. Deletion is subject to the following constraints. You cannot delete a certificate if it is the only remaining certificate in the pack. At least one certificate must remain in the pack.
Custom Hostnames
Fallback Origin
CustomHostnames.FallbackOrigin
Methods
client.CustomHostnames.FallbackOrigin.Get(ctx, query) (*
FallbackOriginGetResponse
, error)
get/zones/{zone_id}/custom_hostnames/fallback_origin
Get Fallback Origin for Custom Hostnames
client.CustomHostnames.FallbackOrigin.Update(ctx, params) (*
FallbackOriginUpdateResponse
, error)
put/zones/{zone_id}/custom_hostnames/fallback_origin
Update Fallback Origin for Custom Hostnames
client.CustomHostnames.FallbackOrigin.Delete(ctx, body) (*
FallbackOriginDeleteResponse
, error)
delete/zones/{zone_id}/custom_hostnames/fallback_origin
Delete Fallback Origin for Custom Hostnames