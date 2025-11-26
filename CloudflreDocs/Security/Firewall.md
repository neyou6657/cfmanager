Firewall
Firewall
Firewall
Access Rules
Firewall.AccessRules
Methods
client.Firewall.AccessRules.List(ctx, params) (*V4PagePaginationArray[
AccessRuleListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/firewall/access_rules/rules
Fetches IP Access rules of an account or zone. These rules apply to all the zones in the account or zone. You can filter the results using several optional parameters.
client.Firewall.AccessRules.Get(ctx, ruleID, query) (*
AccessRuleGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/firewall/access_rules/rules/{rule_id}
Fetches the details of an IP Access rule defined.
client.Firewall.AccessRules.New(ctx, params) (*
AccessRuleNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/firewall/access_rules/rules
Creates a new IP Access rule for an account or zone. The rule will apply to all zones in the account or zone.
Note: To create an IP Access rule that applies to a single zone, refer to the IP Access rules for a zone endpoints.
client.Firewall.AccessRules.Edit(ctx, ruleID, params) (*
AccessRuleEditResponse
, error)
patch/{accounts_or_zones}/{account_or_zone_id}/firewall/access_rules/rules/{rule_id}
Updates an IP Access rule defined.
Note: This operation will affect all zones in the account or zone.
client.Firewall.AccessRules.Delete(ctx, ruleID, body) (*
AccessRuleDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/firewall/access_rules/rules/{rule_id}
Deletes an existing IP Access rule defined.
Note: This operation will affect all zones in the account or zone.
Domain types
typeAccessRuleCIDRConfigurationstruct{…}
typeAccessRuleIPConfigurationstruct{…}
typeASNConfigurationstruct{…}
typeCountryConfigurationstruct{…}
typeIPV6Configurationstruct{…}
Firewall
Lockdowns
Firewall.Lockdowns
Methods
client.Firewall.Lockdowns.List(ctx, params) (*V4PagePaginationArray[
Lockdown
], error)
get/zones/{zone_id}/firewall/lockdowns
Fetches Zone Lockdown rules. You can filter the results using several optional parameters.
client.Firewall.Lockdowns.Get(ctx, lockDownsID, query) (*
Lockdown
, error)
get/zones/{zone_id}/firewall/lockdowns/{lock_downs_id}
Fetches the details of a Zone Lockdown rule.
client.Firewall.Lockdowns.New(ctx, params) (*
Lockdown
, error)
post/zones/{zone_id}/firewall/lockdowns
Creates a new Zone Lockdown rule.
client.Firewall.Lockdowns.Update(ctx, lockDownsID, params) (*
Lockdown
, error)
put/zones/{zone_id}/firewall/lockdowns/{lock_downs_id}
Updates an existing Zone Lockdown rule.
client.Firewall.Lockdowns.Delete(ctx, lockDownsID, body) (*
LockdownDeleteResponse
, error)
delete/zones/{zone_id}/firewall/lockdowns/{lock_downs_id}
Deletes an existing Zone Lockdown rule.
Domain types
typeConfiguration[]
ConfigurationItem
A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule. You can include any number of ip or ip_range configurations.
typeLockdownstruct{…}
typeLockdownCIDRConfigurationstruct{…}
typeLockdownIPConfigurationstruct{…}
typeLockdownURLstring
Firewall
Rules
Firewall.Rules
Methods
client.Firewall.Rules.List(ctx, params) (*V4PagePaginationArray[
FirewallRule
], error)
Deprecated
get/zones/{zone_id}/firewall/rules
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Fetches firewall rules in a zone. You can filter the results using several optional parameters.
client.Firewall.Rules.Get(ctx, ruleID, query) (*
FirewallRule
, error)
Deprecated
get/zones/{zone_id}/firewall/rules/{rule_id}
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Fetches the details of a firewall rule.
client.Firewall.Rules.New(ctx, params) (*SinglePage[
FirewallRule
], error)
Deprecated
post/zones/{zone_id}/firewall/rules
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Create one or more firewall rules.
client.Firewall.Rules.Update(ctx, ruleID, params) (*
FirewallRule
, error)
Deprecated
put/zones/{zone_id}/firewall/rules/{rule_id}
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Updates an existing firewall rule.
client.Firewall.Rules.Edit(ctx, ruleID, body) (*SinglePage[
FirewallRule
], error)
Deprecated
patch/zones/{zone_id}/firewall/rules/{rule_id}
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Updates the priority of an existing firewall rule.
client.Firewall.Rules.Delete(ctx, ruleID, body) (*
FirewallRule
, error)
Deprecated
delete/zones/{zone_id}/firewall/rules/{rule_id}
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Deletes an existing firewall rule.
client.Firewall.Rules.BulkUpdate(ctx, params) (*SinglePage[
FirewallRule
], error)
Deprecated
put/zones/{zone_id}/firewall/rules
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Updates one or more existing firewall rules.
client.Firewall.Rules.BulkEdit(ctx, params) (*SinglePage[
FirewallRule
], error)
Deprecated
patch/zones/{zone_id}/firewall/rules
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Updates the priority of existing firewall rules.
client.Firewall.Rules.BulkDelete(ctx, body) (*SinglePage[
FirewallRule
], error)
Deprecated
delete/zones/{zone_id}/firewall/rules
Deprecated
The Firewall Rules API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.
Deletes existing firewall rules.
Domain types
typeDeletedFilterstruct{…}
typeFirewallRulestruct{…}
typeProductstring
A list of products to bypass for a request when using the bypass action.
Firewall
UA Rules
Firewall.UARules
Methods
client.Firewall.UARules.List(ctx, params) (*V4PagePaginationArray[
UARuleListResponse
], error)
get/zones/{zone_id}/firewall/ua_rules
Fetches User Agent Blocking rules in a zone. You can filter the results using several optional parameters.
client.Firewall.UARules.Get(ctx, uaRuleID, query) (*
UARuleGetResponse
, error)
get/zones/{zone_id}/firewall/ua_rules/{ua_rule_id}
Fetches the details of a User Agent Blocking rule.
client.Firewall.UARules.New(ctx, params) (*
UARuleNewResponse
, error)
post/zones/{zone_id}/firewall/ua_rules
Creates a new User Agent Blocking rule in a zone.
client.Firewall.UARules.Update(ctx, uaRuleID, params) (*
UARuleUpdateResponse
, error)
put/zones/{zone_id}/firewall/ua_rules/{ua_rule_id}
Updates an existing User Agent Blocking rule.
client.Firewall.UARules.Delete(ctx, uaRuleID, body) (*
UARuleDeleteResponse
, error)
delete/zones/{zone_id}/firewall/ua_rules/{ua_rule_id}
Deletes an existing User Agent Blocking rule.
Firewall
WAF
Firewall.WAF
Firewall
WAF
Overrides
Firewall.WAF.Overrides
Methods
client.Firewall.WAF.Overrides.List(ctx, params) (*V4PagePaginationArray[
Override
], error)
Deprecated
get/zones/{zone_id}/firewall/waf/overrides
Fetches the URI-based WAF overrides in a zone.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Overrides.Get(ctx, overridesID, query) (*
Override
, error)
Deprecated
get/zones/{zone_id}/firewall/waf/overrides/{overrides_id}
Fetches the details of a URI-based WAF override.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Overrides.New(ctx, params) (*
Override
, error)
Deprecated
post/zones/{zone_id}/firewall/waf/overrides
Creates a URI-based WAF override for a zone.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Overrides.Update(ctx, overridesID, params) (*
Override
, error)
Deprecated
put/zones/{zone_id}/firewall/waf/overrides/{overrides_id}
Updates an existing URI-based WAF override.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Overrides.Delete(ctx, overridesID, body) (*
WAFOverrideDeleteResponse
, error)
Deprecated
delete/zones/{zone_id}/firewall/waf/overrides/{overrides_id}
Deletes an existing URI-based WAF override.
Note: Applies only to the previous version of WAF managed rules.
Domain types
typeOverridestruct{…}
typeOverrideURLstring
typeRewriteActionstruct{…}
Specifies that, when a WAF rule matches, its configured action will be replaced by the action configured in this object.
typeWAFRulemap[string,
WAFRuleItem
]
An object that allows you to override the action of specific WAF rules. Each key of this object must be the ID of a WAF rule, and each value must be a valid WAF action. Unless you are disabling a rule, ensure that you also enable the rule group that this WAF rule belongs to. When creating a new URI-based WAF override, you must provide a groups object or a rules object.
Firewall
WAF
Packages
Firewall.WAF.Packages
Methods
client.Firewall.WAF.Packages.List(ctx, params) (*V4PagePaginationArray[
WAFPackageListResponse
], error)
Deprecated
get/zones/{zone_id}/firewall/waf/packages
Fetches WAF packages for a zone.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Packages.Get(ctx, packageID, query) (*
WAFPackageGetResponse
, error)
Deprecated
get/zones/{zone_id}/firewall/waf/packages/{package_id}
Fetches the details of a WAF package.
Note: Applies only to the previous version of WAF managed rules.
Firewall
WAF
Packages
Groups
Firewall.WAF.Packages.Groups
Methods
client.Firewall.WAF.Packages.Groups.List(ctx, packageID, params) (*V4PagePaginationArray[
Group
], error)
Deprecated
get/zones/{zone_id}/firewall/waf/packages/{package_id}/groups
Fetches the WAF rule groups in a WAF package.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Packages.Groups.Get(ctx, packageID, groupID, query) (*unknown, error)
Deprecated
get/zones/{zone_id}/firewall/waf/packages/{package_id}/groups/{group_id}
Fetches the details of a WAF rule group.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Packages.Groups.Edit(ctx, packageID, groupID, params) (*unknown, error)
Deprecated
patch/zones/{zone_id}/firewall/waf/packages/{package_id}/groups/{group_id}
Updates a WAF rule group. You can update the state (mode parameter) of a rule group.
Note: Applies only to the previous version of WAF managed rules.
Domain types
typeGroupstruct{…}
Firewall
WAF
Packages
Rules
Firewall.WAF.Packages.Rules
Methods
client.Firewall.WAF.Packages.Rules.List(ctx, packageID, params) (*V4PagePaginationArray[
WAFPackageRuleListResponse
], error)
Deprecated
get/zones/{zone_id}/firewall/waf/packages/{package_id}/rules
Fetches WAF rules in a WAF package.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Packages.Rules.Get(ctx, packageID, ruleID, query) (*unknown, error)
Deprecated
get/zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}
Fetches the details of a WAF rule in a WAF package.
Note: Applies only to the previous version of WAF managed rules.
client.Firewall.WAF.Packages.Rules.Edit(ctx, packageID, ruleID, params) (*
WAFPackageRuleEditResponse
, error)
Deprecated
patch/zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}
Updates a WAF rule. You can only update the mode/action of the rule.
Note: Applies only to the previous version of WAF managed rules.
Domain types
typeAllowedModesAnomalystring
Defines the mode anomaly. When set to on, the current WAF rule will be used when evaluating the request. Applies to anomaly detection WAF rules.
typeWAFRuleGroupstruct{…}
Defines the rule group to which the current WAF rule belongs.