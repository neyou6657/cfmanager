Token Validation
TokenValidation
Token Validation
Configuration
TokenValidation.Configuration
Methods
client.TokenValidation.Configuration.List(ctx, params) (*V4PagePaginationArray[
TokenConfig
], error)
get/zones/{zone_id}/token_validation/config
Lists all token validation configurations for this zone
client.TokenValidation.Configuration.Get(ctx, configID, query) (*
TokenConfig
, error)
get/zones/{zone_id}/token_validation/config/{config_id}
Get a single Token Configuration
client.TokenValidation.Configuration.New(ctx, params) (*
TokenConfig
, error)
post/zones/{zone_id}/token_validation/config
Create a new Token Validation configuration
client.TokenValidation.Configuration.Edit(ctx, configID, params) (*
ConfigurationEditResponse
, error)
patch/zones/{zone_id}/token_validation/config/{config_id}
Edit fields of an existing Token Configuration
client.TokenValidation.Configuration.Delete(ctx, configID, body) (*
ConfigurationDeleteResponse
, error)
delete/zones/{zone_id}/token_validation/config/{config_id}
Delete Token Configuration
Domain types
typeTokenConfigstruct{…}
Token Validation
Configuration
Credentials
TokenValidation.Configuration.Credentials
Methods
client.TokenValidation.Configuration.Credentials.Update(ctx, configID, params) (*
ConfigurationCredentialUpdateResponse
, error)
put/zones/{zone_id}/token_validation/config/{config_id}/credentials
Update Token Configuration credentials
Token Validation
Rules
TokenValidation.Rules
Methods
client.TokenValidation.Rules.List(ctx, params) (*V4PagePaginationArray[
TokenValidationRule
], error)
get/zones/{zone_id}/token_validation/rules
List token validation rules
client.TokenValidation.Rules.New(ctx, params) (*
TokenValidationRule
, error)
post/zones/{zone_id}/token_validation/rules
Create a token validation rule.
client.TokenValidation.Rules.BulkNew(ctx, params) (*SinglePage[
TokenValidationRule
], error)
post/zones/{zone_id}/token_validation/rules/bulk
Create zone token validation rules.
A request can create multiple Token Validation Rules.
client.TokenValidation.Rules.BulkEdit(ctx, params) (*SinglePage[
TokenValidationRule
], error)
patch/zones/{zone_id}/token_validation/rules/bulk
Edit token validation rules.
A request can update multiple Token Validation Rules.
Rules can be re-ordered using the position field.
Returns all updated rules.
client.TokenValidation.Rules.Get(ctx, ruleID, query) (*
TokenValidationRule
, error)
get/zones/{zone_id}/token_validation/rules/{rule_id}
Get a zone token validation rule.
client.TokenValidation.Rules.Delete(ctx, ruleID, body) (*
RuleDeleteResponse
, error)
delete/zones/{zone_id}/token_validation/rules/{rule_id}
Delete a zone token validation rule.
client.TokenValidation.Rules.Edit(ctx, ruleID, params) (*
TokenValidationRule
, error)
patch/zones/{zone_id}/token_validation/rules/{rule_id}
Edit a zone token validation rule.
Domain types
typeTokenValidationRulestruct{…}
A Token Validation rule that can enforce security policies using JWT Tokens.