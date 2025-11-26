Rulesets
Rulesets
Methods
client.Rulesets.List(ctx, params) (*CursorPagination[
RulesetListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/rulesets
Fetches all rulesets.
client.Rulesets.Get(ctx, rulesetID, query) (*
RulesetGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}
Fetches the latest version of an account or zone ruleset.
client.Rulesets.New(ctx, params) (*
RulesetNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/rulesets
Creates a ruleset.
client.Rulesets.Update(ctx, rulesetID, params) (*
RulesetUpdateResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}
Updates an account or zone ruleset, creating a new version.
client.Rulesets.Delete(ctx, rulesetID, body) error
delete/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}
Deletes all versions of an existing account or zone ruleset.
Domain types
typeKindstring
The kind of the ruleset.
typePhasestring
The phase of the ruleset.
typeRulesetstruct{…}
A ruleset object.
Rulesets
Phases
Rulesets.Phases
Methods
client.Rulesets.Phases.Get(ctx, rulesetPhase, query) (*
PhaseGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/rulesets/phases/{ruleset_phase}/entrypoint
Fetches the latest version of the account or zone entry point ruleset for a given phase.
client.Rulesets.Phases.Update(ctx, rulesetPhase, params) (*
PhaseUpdateResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/rulesets/phases/{ruleset_phase}/entrypoint
Updates an account or zone entry point ruleset, creating a new version.
Rulesets
Phases
Versions
Rulesets.Phases.Versions
Methods
client.Rulesets.Phases.Versions.List(ctx, rulesetPhase, query) (*SinglePage[
PhaseVersionListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/rulesets/phases/{ruleset_phase}/entrypoint/versions
Fetches the versions of an account or zone entry point ruleset.
client.Rulesets.Phases.Versions.Get(ctx, rulesetPhase, rulesetVersion, query) (*
PhaseVersionGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/rulesets/phases/{ruleset_phase}/entrypoint/versions/{ruleset_version}
Fetches a specific version of an account or zone entry point ruleset.
Rulesets
Rules
Rulesets.Rules
Methods
client.Rulesets.Rules.New(ctx, rulesetID, params) (*
RuleNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/rules
Adds a new rule to an account or zone ruleset. The rule will be added to the end of the existing list of rules in the ruleset by default.
client.Rulesets.Rules.Edit(ctx, rulesetID, ruleID, params) (*
RuleEditResponse
, error)
patch/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/rules/{rule_id}
Updates an existing rule in an account or zone ruleset.
client.Rulesets.Rules.Delete(ctx, rulesetID, ruleID, body) (*
RuleDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/rules/{rule_id}
Deletes an existing rule from an account or zone ruleset.
Domain types
typeBlockRulestruct{…}
typeCompressResponseRulestruct{…}
typeDDoSDynamicRulestruct{…}
typeExecuteRulestruct{…}
typeForceConnectionCloseRulestruct{…}
typeLogCustomFieldRulestruct{…}
typeLogRulestruct{…}
typeLoggingstruct{…}
An object configuring the rule's logging behavior.
typeManagedChallengeRulestruct{…}
typeRedirectRulestruct{…}
typeRewriteRulestruct{…}
typeRouteRulestruct{…}
typeRulesetRulestruct{…}
typeScoreRulestruct{…}
typeServeErrorRulestruct{…}
typeSetCacheSettingsRulestruct{…}
typeSetConfigRulestruct{…}
typeSkipRulestruct{…}
Rulesets
Versions
Rulesets.Versions
Methods
client.Rulesets.Versions.List(ctx, rulesetID, query) (*SinglePage[
VersionListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/versions
Fetches the versions of an account or zone ruleset.
client.Rulesets.Versions.Get(ctx, rulesetID, rulesetVersion, query) (*
VersionGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/versions/{ruleset_version}
Fetches a specific version of an account or zone ruleset.
client.Rulesets.Versions.Delete(ctx, rulesetID, rulesetVersion, body) error
delete/{accounts_or_zones}/{account_or_zone_id}/rulesets/{ruleset_id}/versions/{ruleset_version}
Deletes an existing version of an account or zone ruleset.