Registrar
Registrar
Registrar
Domains
Registrar.Domains
Methods
client.Registrar.Domains.List(ctx, query) (*SinglePage[
Domain
], error)
get/accounts/{account_id}/registrar/domains
List domains handled by Registrar.
client.Registrar.Domains.Get(ctx, domainName, query) (*
DomainGetResponse
, error)
get/accounts/{account_id}/registrar/domains/{domain_name}
Show individual domain.
client.Registrar.Domains.Update(ctx, domainName, params) (*
DomainUpdateResponse
, error)
put/accounts/{account_id}/registrar/domains/{domain_name}
Update individual domain.
Domain types
typeDomainstruct{…}