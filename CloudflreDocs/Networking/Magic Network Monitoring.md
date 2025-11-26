Magic Network Monitoring
MagicNetworkMonitoring
Magic Network Monitoring
Configs
MagicNetworkMonitoring.Configs
Methods
client.MagicNetworkMonitoring.Configs.Get(ctx, query) (*
Configuration
, error)
get/accounts/{account_id}/mnm/config
Lists default sampling, router IPs and warp devices for account.
client.MagicNetworkMonitoring.Configs.New(ctx, params) (*
Configuration
, error)
post/accounts/{account_id}/mnm/config
Create a new network monitoring configuration.
client.MagicNetworkMonitoring.Configs.Update(ctx, params) (*
Configuration
, error)
put/accounts/{account_id}/mnm/config
Update an existing network monitoring configuration, requires the entire configuration to be updated at once.
client.MagicNetworkMonitoring.Configs.Edit(ctx, params) (*
Configuration
, error)
patch/accounts/{account_id}/mnm/config
Update fields in an existing network monitoring configuration.
client.MagicNetworkMonitoring.Configs.Delete(ctx, body) (*
Configuration
, error)
delete/accounts/{account_id}/mnm/config
Delete an existing network monitoring configuration.
Domain types
typeConfigurationstruct{…}
Magic Network Monitoring
Configs
Full
MagicNetworkMonitoring.Configs.Full
Methods
client.MagicNetworkMonitoring.Configs.Full.Get(ctx, query) (*
Configuration
, error)
get/accounts/{account_id}/mnm/config/full
Lists default sampling, router IPs, warp devices, and rules for account.
Magic Network Monitoring
Rules
MagicNetworkMonitoring.Rules
Methods
client.MagicNetworkMonitoring.Rules.List(ctx, query) (*SinglePage[
MagicNetworkMonitoringRule
], error)
get/accounts/{account_id}/mnm/rules
Lists network monitoring rules for account.
client.MagicNetworkMonitoring.Rules.Get(ctx, ruleID, query) (*
MagicNetworkMonitoringRule
, error)
get/accounts/{account_id}/mnm/rules/{rule_id}
List a single network monitoring rule for account.
client.MagicNetworkMonitoring.Rules.New(ctx, params) (*
MagicNetworkMonitoringRule
, error)
post/accounts/{account_id}/mnm/rules
Create network monitoring rules for account. Currently only supports creating a single rule per API request.
client.MagicNetworkMonitoring.Rules.Update(ctx, params) (*
MagicNetworkMonitoringRule
, error)
put/accounts/{account_id}/mnm/rules
Update network monitoring rules for account.
client.MagicNetworkMonitoring.Rules.Edit(ctx, ruleID, params) (*
MagicNetworkMonitoringRule
, error)
patch/accounts/{account_id}/mnm/rules/{rule_id}
Update a network monitoring rule for account.
client.MagicNetworkMonitoring.Rules.Delete(ctx, ruleID, body) (*
MagicNetworkMonitoringRule
, error)
delete/accounts/{account_id}/mnm/rules/{rule_id}
Delete a network monitoring rule for account.
Domain types
typeMagicNetworkMonitoringRulestruct{…}
Magic Network Monitoring
Rules
Advertisements
MagicNetworkMonitoring.Rules.Advertisements
Methods
client.MagicNetworkMonitoring.Rules.Advertisements.Edit(ctx, ruleID, params) (*
Advertisement
, error)
patch/accounts/{account_id}/mnm/rules/{rule_id}/advertisement
Update advertisement for rule.
Domain types
typeAdvertisementstruct{…}
Magic Network Monitoring
VPC Flows
MagicNetworkMonitoring.VPCFlows
Magic Network Monitoring
VPC Flows
Tokens
MagicNetworkMonitoring.VPCFlows.Tokens
Methods
client.MagicNetworkMonitoring.VPCFlows.Tokens.New(ctx, body) (*string, error)
post/accounts/{account_id}/mnm/vpc-flows/token
Generate authentication token for VPC flow logs export.