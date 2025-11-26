Page Rules
PageRules
Methods
client.PageRules.List(ctx, params) (*[]
PageRule
, error)
get/zones/{zone_id}/pagerules
Fetches Page Rules in a zone.
client.PageRules.Get(ctx, pageruleID, query) (*
PageRule
, error)
get/zones/{zone_id}/pagerules/{pagerule_id}
Fetches the details of a Page Rule.
client.PageRules.New(ctx, params) (*
PageRule
, error)
post/zones/{zone_id}/pagerules
Creates a new Page Rule.
client.PageRules.Update(ctx, pageruleID, params) (*
PageRule
, error)
put/zones/{zone_id}/pagerules/{pagerule_id}
Replaces the configuration of an existing Page Rule. The configuration of the updated Page Rule will exactly match the data passed in the API request.
client.PageRules.Edit(ctx, pageruleID, params) (*
PageRule
, error)
patch/zones/{zone_id}/pagerules/{pagerule_id}
Updates one or more fields of an existing Page Rule.
client.PageRules.Delete(ctx, pageruleID, body) (*
PageRuleDeleteResponse
, error)
delete/zones/{zone_id}/pagerules/{pagerule_id}
Deletes an existing Page Rule.
Domain types
typePageRulestruct{…}
typeTargetstruct{…}
URL target.