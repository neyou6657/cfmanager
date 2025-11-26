Snippets
Snippets
Methods
client.Snippets.List(ctx, params) (*V4PagePaginationArray[
SnippetListResponse
], error)
get/zones/{zone_id}/snippets
Fetches all snippets belonging to the zone.
client.Snippets.Get(ctx, snippetName, query) (*
SnippetGetResponse
, error)
get/zones/{zone_id}/snippets/{snippet_name}
Fetches a snippet belonging to the zone.
client.Snippets.Update(ctx, snippetName, params) (*
SnippetUpdateResponse
, error)
put/zones/{zone_id}/snippets/{snippet_name}
Creates or updates a snippet belonging to the zone.
client.Snippets.Delete(ctx, snippetName, body) (*string, error)
delete/zones/{zone_id}/snippets/{snippet_name}
Deletes a snippet belonging to the zone.
Snippets
Content
Snippets.Content
Methods
client.Snippets.Content.Get(ctx, snippetName, query) (*
Response
, error)
get/zones/{zone_id}/snippets/{snippet_name}/content
Fetches the content of a snippet belonging to the zone.
Snippets
Rules
Snippets.Rules
Methods
client.Snippets.Rules.List(ctx, query) (*SinglePage[
RuleListResponse
], error)
get/zones/{zone_id}/snippets/snippet_rules
Fetches all snippet rules belonging to the zone.
client.Snippets.Rules.Update(ctx, params) (*SinglePage[
RuleUpdateResponse
], error)
put/zones/{zone_id}/snippets/snippet_rules
Updates all snippet rules belonging to the zone.
client.Snippets.Rules.Delete(ctx, body) (*SinglePage[
RuleDeleteResponse
], error)
delete/zones/{zone_id}/snippets/snippet_rules
Deletes all snippet rules belonging to the zone.