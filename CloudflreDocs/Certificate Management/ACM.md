ACM
ACM
ACM
Total TLS
ACM.TotalTLS
Methods
client.ACM.TotalTLS.Get(ctx, query) (*
TotalTLSGetResponse
, error)
get/zones/{zone_id}/acm/total_tls
Get Total TLS Settings for a Zone.
client.ACM.TotalTLS.New(ctx, params) (*
TotalTLSNewResponse
, error)
post/zones/{zone_id}/acm/total_tls
Set Total TLS Settings or disable the feature for a Zone.
Domain types
typeCertificateAuthoritystring
The Certificate Authority that Total TLS certificates will be issued through.