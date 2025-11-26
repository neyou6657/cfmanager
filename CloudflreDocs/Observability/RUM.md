RUM
RUM
RUM
Rules
RUM.Rules
Methods
client.RUM.Rules.List(ctx, rulesetID, query) (*
RuleListResponse
, error)
get/accounts/{account_id}/rum/v2/{ruleset_id}/rules
Lists all the rules in a Web Analytics ruleset.
client.RUM.Rules.New(ctx, rulesetID, params) (*
RUMRule
, error)
post/accounts/{account_id}/rum/v2/{ruleset_id}/rule
Creates a new rule in a Web Analytics ruleset.
client.RUM.Rules.Update(ctx, rulesetID, ruleID, params) (*
RUMRule
, error)
put/accounts/{account_id}/rum/v2/{ruleset_id}/rule/{rule_id}
Updates a rule in a Web Analytics ruleset.
client.RUM.Rules.Delete(ctx, rulesetID, ruleID, body) (*
RuleDeleteResponse
, error)
delete/accounts/{account_id}/rum/v2/{ruleset_id}/rule/{rule_id}
Deletes an existing rule from a Web Analytics ruleset.
client.RUM.Rules.BulkNew(ctx, rulesetID, params) (*
RuleBulkNewResponse
, error)
post/accounts/{account_id}/rum/v2/{ruleset_id}/rules
Modifies one or more rules in a Web Analytics ruleset with a single request.
Domain types
typeRUMRulestruct{…}
RUM
Site Info
RUM.SiteInfo
Methods
client.RUM.SiteInfo.List(ctx, params) (*V4PagePaginationArray[
Site
], error)
get/accounts/{account_id}/rum/site_info/list
Lists all Web Analytics sites of an account.
client.RUM.SiteInfo.Get(ctx, siteID, query) (*
Site
, error)
get/accounts/{account_id}/rum/site_info/{site_id}
Retrieves a Web Analytics site.
client.RUM.SiteInfo.New(ctx, params) (*
Site
, error)
post/accounts/{account_id}/rum/site_info
Creates a new Web Analytics site.
client.RUM.SiteInfo.Update(ctx, siteID, params) (*
Site
, error)
put/accounts/{account_id}/rum/site_info/{site_id}
Updates an existing Web Analytics site.
client.RUM.SiteInfo.Delete(ctx, siteID, body) (*
SiteInfoDeleteResponse
, error)
delete/accounts/{account_id}/rum/site_info/{site_id}
Deletes an existing Web Analytics site.
Domain types
typeSitestruct{…}