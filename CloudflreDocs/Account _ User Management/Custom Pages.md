Custom Pages
CustomPages
Methods
client.CustomPages.List(ctx, query) (*SinglePage[
CustomPageListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/custom_pages
Fetches all the custom pages.
client.CustomPages.Get(ctx, identifier, query) (*
CustomPageGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/custom_pages/{identifier}
Fetches the details of a custom page.
client.CustomPages.Update(ctx, identifier, params) (*
CustomPageUpdateResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/custom_pages/{identifier}
Updates the configuration of an existing custom page.