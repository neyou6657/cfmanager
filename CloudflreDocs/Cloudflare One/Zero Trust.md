Zero Trust
ZeroTrust
Zero Trust
Access
ZeroTrust.Access
Zero Trust
Access
AI Controls
ZeroTrust.Access.AIControls
Zero Trust
Access
AI Controls
Mcp
ZeroTrust.Access.AIControls.Mcp
Zero Trust
Access
AI Controls
Mcp
Portals
ZeroTrust.Access.AIControls.Mcp.Portals
Methods
client.ZeroTrust.Access.AIControls.Mcp.Portals.List(ctx, params) (*V4PagePaginationArray[
AccessAIControlMcpPortalListResponse
], error)
get/accounts/{account_id}/access/ai-controls/mcp/portals
List MCP Portals
client.ZeroTrust.Access.AIControls.Mcp.Portals.New(ctx, params) (*
AccessAIControlMcpPortalNewResponse
, error)
post/accounts/{account_id}/access/ai-controls/mcp/portals
Create a new MCP Portal
client.ZeroTrust.Access.AIControls.Mcp.Portals.Read(ctx, id, query) (*
AccessAIControlMcpPortalReadResponse
, error)
get/accounts/{account_id}/access/ai-controls/mcp/portals/{id}
Read details of an MCP Portal
client.ZeroTrust.Access.AIControls.Mcp.Portals.Update(ctx, id, params) (*
AccessAIControlMcpPortalUpdateResponse
, error)
put/accounts/{account_id}/access/ai-controls/mcp/portals/{id}
Update a MCP Portal
client.ZeroTrust.Access.AIControls.Mcp.Portals.Delete(ctx, id, body) (*
AccessAIControlMcpPortalDeleteResponse
, error)
delete/accounts/{account_id}/access/ai-controls/mcp/portals/{id}
Delete a MCP Portal
Zero Trust
Access
AI Controls
Mcp
Servers
ZeroTrust.Access.AIControls.Mcp.Servers
Methods
client.ZeroTrust.Access.AIControls.Mcp.Servers.List(ctx, params) (*V4PagePaginationArray[
AccessAIControlMcpServerListResponse
], error)
get/accounts/{account_id}/access/ai-controls/mcp/servers
List MCP Servers
client.ZeroTrust.Access.AIControls.Mcp.Servers.New(ctx, params) (*
AccessAIControlMcpServerNewResponse
, error)
post/accounts/{account_id}/access/ai-controls/mcp/servers
Create a new MCP Server
client.ZeroTrust.Access.AIControls.Mcp.Servers.Read(ctx, id, query) (*
AccessAIControlMcpServerReadResponse
, error)
get/accounts/{account_id}/access/ai-controls/mcp/servers/{id}
Read the details of a MCP Server
client.ZeroTrust.Access.AIControls.Mcp.Servers.Update(ctx, id, params) (*
AccessAIControlMcpServerUpdateResponse
, error)
put/accounts/{account_id}/access/ai-controls/mcp/servers/{id}
Update a MCP Server
client.ZeroTrust.Access.AIControls.Mcp.Servers.Delete(ctx, id, body) (*
AccessAIControlMcpServerDeleteResponse
, error)
delete/accounts/{account_id}/access/ai-controls/mcp/servers/{id}
Delete a MCP Server
client.ZeroTrust.Access.AIControls.Mcp.Servers.Sync(ctx, id, body) (*
AccessAIControlMcpServerSyncResponse
, error)
post/accounts/{account_id}/access/ai-controls/mcp/servers/{id}/sync
Sync MCP Server Capabilities
Zero Trust
Access
Applications
ZeroTrust.Access.Applications
Methods
client.ZeroTrust.Access.Applications.List(ctx, params) (*V4PagePaginationArray[
AccessApplicationListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/apps
Lists all Access applications in an account or zone.
client.ZeroTrust.Access.Applications.Get(ctx, appID, query) (*
AccessApplicationGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}
Fetches information about an Access application.
client.ZeroTrust.Access.Applications.New(ctx, params) (*
AccessApplicationNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/apps
Adds a new application to Access.
client.ZeroTrust.Access.Applications.Update(ctx, appID, params) (*
AccessApplicationUpdateResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}
Updates an Access application.
client.ZeroTrust.Access.Applications.Delete(ctx, appID, body) (*
AccessApplicationDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}
Deletes an application from Access.
client.ZeroTrust.Access.Applications.RevokeTokens(ctx, appID, body) (*
AccessApplicationRevokeTokensResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/revoke_tokens
Revokes all tokens issued for an application.
Domain types
typeAllowedHeadersstring
typeAllowedIdPsstring
The identity providers selected for application.
typeAllowedMethodsstring
typeAllowedOriginsstring
typeAppIDstring
Identifier.
typeApplicationinterface{…}
typeApplicationPolicystruct{…}
typeApplicationSCIMConfigstruct{…}
Configuration for provisioning to this application via SCIM. This is currently in closed beta.
typeApplicationTypestring
The application type.
typeCORSHeadersstruct{…}
typeDecisionstring
The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.
typeOIDCSaaSAppstruct{…}
typeSaaSAppNameIDFormatstring
The format of the name identifier sent to the SaaS application.
typeSAMLSaaSAppstruct{…}
typeSCIMConfigAuthenticationHTTPBasicstruct{…}
Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.
typeSCIMConfigAuthenticationOauth2struct{…}
Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning to an application.
typeSCIMConfigAuthenticationOAuthBearerTokenstruct{…}
Attributes for configuring OAuth Bearer Token authentication scheme for SCIM provisioning to an application.
typeSCIMConfigMappingstruct{…}
Transformations and filters applied to resources before they are provisioned in the remote SCIM service.
typeSelfHostedDomainsstring
A domain that Access will secure.
Zero Trust
Access
Applications
CAs
ZeroTrust.Access.Applications.CAs
Methods
client.ZeroTrust.Access.Applications.CAs.List(ctx, params) (*V4PagePaginationArray[
CA
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/apps/ca
Lists short-lived certificate CAs and their public keys.
client.ZeroTrust.Access.Applications.CAs.Get(ctx, appID, query) (*
CA
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/ca
Fetches a short-lived certificate CA and its public key.
client.ZeroTrust.Access.Applications.CAs.New(ctx, appID, body) (*
CA
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/ca
Generates a new short-lived certificate CA and public key.
client.ZeroTrust.Access.Applications.CAs.Delete(ctx, appID, body) (*
AccessApplicationCADeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/ca
Deletes a short-lived certificate CA.
Domain types
typeCAstruct{…}
Zero Trust
Access
Applications
Policies
ZeroTrust.Access.Applications.Policies
Methods
client.ZeroTrust.Access.Applications.Policies.List(ctx, appID, params) (*V4PagePaginationArray[
AccessApplicationPolicyListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/policies
Lists Access policies configured for an application. Returns both exclusively scoped and reusable policies used by the application.
client.ZeroTrust.Access.Applications.Policies.Get(ctx, appID, policyID, query) (*
AccessApplicationPolicyGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/policies/{policy_id}
Fetches a single Access policy configured for an application. Returns both exclusively owned and reusable policies used by the application.
client.ZeroTrust.Access.Applications.Policies.New(ctx, appID, params) (*
AccessApplicationPolicyNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/policies
Creates a policy applying exclusive to a single application that defines the users or groups who can reach it. We recommend creating a reusable policy instead and subsequently referencing its ID in the application's 'policies' array.
client.ZeroTrust.Access.Applications.Policies.Update(ctx, appID, policyID, params) (*
AccessApplicationPolicyUpdateResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/policies/{policy_id}
Updates an Access policy specific to an application. To update a reusable policy, use the /account or zones/{account or zone_id}/policies/{uid} endpoint.
client.ZeroTrust.Access.Applications.Policies.Delete(ctx, appID, policyID, body) (*
AccessApplicationPolicyDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/policies/{policy_id}
Deletes an Access policy specific to an application. To delete a reusable policy, use the /account or zones/{account or zone_id}/policies/{uid} endpoint.
Domain types
typeAccessDevicePostureRulestruct{…}
Enforces a device posture rule has run successfully
typeAccessRuleinterface{…}
Matches an Access group.
typeAnyValidServiceTokenRulestruct{…}
Matches any valid Access Service Token
typeAuthenticationMethodRulestruct{…}
Enforce different MFA options
typeAzureGroupRulestruct{…}
Matches an Azure group. Requires an Azure identity provider.
typeCertificateRulestruct{…}
Matches any valid client certificate.
typeCountryRulestruct{…}
Matches a specific country
typeDomainRulestruct{…}
Match an entire email domain.
typeEmailListRulestruct{…}
Matches an email address from a list.
typeEmailRulestruct{…}
Matches a specific email.
typeEveryoneRulestruct{…}
Matches everyone.
typeExternalEvaluationRulestruct{…}
Create Allow or Block policies which evaluate the user based on custom criteria.
typeGitHubOrganizationRulestruct{…}
Matches a Github organization. Requires a Github identity provider.
typeGroupRulestruct{…}
Matches an Access group.
typeGSuiteGroupRulestruct{…}
Matches a group in Google Workspace. Requires a Google Workspace identity provider.
typeIPListRulestruct{…}
Matches an IP address from a list.
typeIPRulestruct{…}
Matches an IP address block.
typeOktaGroupRulestruct{…}
Matches an Okta group. Requires an Okta identity provider.
typeSAMLGroupRulestruct{…}
Matches a SAML group. Requires a SAML identity provider.
typeServiceTokenRulestruct{…}
Matches a specific Access Service Token
Zero Trust
Access
Applications
Policy Tests
ZeroTrust.Access.Applications.PolicyTests
Methods
client.ZeroTrust.Access.Applications.PolicyTests.Get(ctx, policyTestID, query) (*
AccessApplicationPolicyTestGetResponse
, error)
get/accounts/{account_id}/access/policy-tests/{policy_test_id}
Fetches the current status of a given Access policy test.
client.ZeroTrust.Access.Applications.PolicyTests.New(ctx, params) (*
AccessApplicationPolicyTestNewResponse
, error)
post/accounts/{account_id}/access/policy-tests
Starts an Access policy test.
Zero Trust
Access
Applications
Policy Tests
Users
ZeroTrust.Access.Applications.PolicyTests.Users
Methods
client.ZeroTrust.Access.Applications.PolicyTests.Users.List(ctx, policyTestID, params) (*V4PagePaginationArray[
AccessApplicationPolicyTestUserListResponse
], error)
get/accounts/{account_id}/access/policy-tests/{policy_test_id}/users
Fetches a single page of user results from an Access policy test.
Zero Trust
Access
Applications
Settings
ZeroTrust.Access.Applications.Settings
Methods
client.ZeroTrust.Access.Applications.Settings.Update(ctx, appID, params) (*
AccessApplicationSettingUpdateResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/settings
Updates Access application settings.
client.ZeroTrust.Access.Applications.Settings.Edit(ctx, appID, params) (*
AccessApplicationSettingEditResponse
, error)
patch/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/settings
Updates Access application settings.
Zero Trust
Access
Applications
User Policy Checks
ZeroTrust.Access.Applications.UserPolicyChecks
Methods
client.ZeroTrust.Access.Applications.UserPolicyChecks.List(ctx, appID, query) (*
AccessApplicationUserPolicyCheckListResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/apps/{app_id}/user_policy_checks
Tests if a specific user has permission to access an application.
Domain types
typeUserPolicyCheckGeostruct{…}
Zero Trust
Access
Bookmarks
ZeroTrust.Access.Bookmarks
Methods
client.ZeroTrust.Access.Bookmarks.List(ctx, query) (*SinglePage[
Bookmark
], error)
Deprecated
get/accounts/{account_id}/access/bookmarks
Lists Bookmark applications.
client.ZeroTrust.Access.Bookmarks.Get(ctx, bookmarkID, query) (*
Bookmark
, error)
Deprecated
get/accounts/{account_id}/access/bookmarks/{bookmark_id}
Fetches a single Bookmark application.
client.ZeroTrust.Access.Bookmarks.New(ctx, bookmarkID, params) (*
Bookmark
, error)
Deprecated
post/accounts/{account_id}/access/bookmarks/{bookmark_id}
Create a new Bookmark application.
client.ZeroTrust.Access.Bookmarks.Update(ctx, bookmarkID, params) (*
Bookmark
, error)
Deprecated
put/accounts/{account_id}/access/bookmarks/{bookmark_id}
Updates a configured Bookmark application.
client.ZeroTrust.Access.Bookmarks.Delete(ctx, bookmarkID, body) (*
AccessBookmarkDeleteResponse
, error)
Deprecated
delete/accounts/{account_id}/access/bookmarks/{bookmark_id}
Deletes a Bookmark application.
Domain types
typeBookmarkstruct{…}
Zero Trust
Access
Certificates
ZeroTrust.Access.Certificates
Methods
client.ZeroTrust.Access.Certificates.List(ctx, params) (*V4PagePaginationArray[
Certificate
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/certificates
Lists all mTLS root certificates.
client.ZeroTrust.Access.Certificates.Get(ctx, certificateID, query) (*
Certificate
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/certificates/{certificate_id}
Fetches a single mTLS certificate.
client.ZeroTrust.Access.Certificates.New(ctx, params) (*
Certificate
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/certificates
Adds a new mTLS root certificate to Access.
client.ZeroTrust.Access.Certificates.Update(ctx, certificateID, params) (*
Certificate
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/certificates/{certificate_id}
Updates a configured mTLS certificate.
client.ZeroTrust.Access.Certificates.Delete(ctx, certificateID, body) (*
AccessCertificateDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/access/certificates/{certificate_id}
Deletes an mTLS certificate.
Domain types
typeAssociatedHostnamesstring
A fully-qualified domain name (FQDN).
typeCertificatestruct{…}
Zero Trust
Access
Certificates
Settings
ZeroTrust.Access.Certificates.Settings
Methods
client.ZeroTrust.Access.Certificates.Settings.Get(ctx, query) (*SinglePage[
CertificateSettings
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/certificates/settings
List all mTLS hostname settings for this account or zone.
client.ZeroTrust.Access.Certificates.Settings.Update(ctx, params) (*SinglePage[
CertificateSettings
], error)
put/{accounts_or_zones}/{account_or_zone_id}/access/certificates/settings
Updates an mTLS certificate's hostname settings.
Domain types
typeCertificateSettingsstruct{…}
Zero Trust
Access
Custom Pages
ZeroTrust.Access.CustomPages
Methods
client.ZeroTrust.Access.CustomPages.List(ctx, params) (*V4PagePaginationArray[
CustomPageWithoutHTML
], error)
get/accounts/{account_id}/access/custom_pages
List custom pages
client.ZeroTrust.Access.CustomPages.Get(ctx, customPageID, query) (*
CustomPage
, error)
get/accounts/{account_id}/access/custom_pages/{custom_page_id}
Fetches a custom page and also returns its HTML.
client.ZeroTrust.Access.CustomPages.New(ctx, params) (*
CustomPageWithoutHTML
, error)
post/accounts/{account_id}/access/custom_pages
Create a custom page
client.ZeroTrust.Access.CustomPages.Update(ctx, customPageID, params) (*
CustomPageWithoutHTML
, error)
put/accounts/{account_id}/access/custom_pages/{custom_page_id}
Update a custom page
client.ZeroTrust.Access.CustomPages.Delete(ctx, customPageID, body) (*
AccessCustomPageDeleteResponse
, error)
delete/accounts/{account_id}/access/custom_pages/{custom_page_id}
Delete a custom page
Domain types
typeCustomPagestruct{…}
typeCustomPageWithoutHTMLstruct{…}
Zero Trust
Access
Gateway CA
ZeroTrust.Access.GatewayCA
Methods
client.ZeroTrust.Access.GatewayCA.List(ctx, query) (*SinglePage[
AccessGatewayCAListResponse
], error)
get/accounts/{account_id}/access/gateway_ca
Lists SSH Certificate Authorities (CA).
client.ZeroTrust.Access.GatewayCA.New(ctx, body) (*
AccessGatewayCANewResponse
, error)
post/accounts/{account_id}/access/gateway_ca
Adds a new SSH Certificate Authority (CA).
client.ZeroTrust.Access.GatewayCA.Delete(ctx, certificateID, body) (*
AccessGatewayCADeleteResponse
, error)
delete/accounts/{account_id}/access/gateway_ca/{certificate_id}
Deletes an SSH Certificate Authority.
Zero Trust
Access
Groups
ZeroTrust.Access.Groups
Methods
client.ZeroTrust.Access.Groups.List(ctx, params) (*V4PagePaginationArray[
AccessGroupListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/groups
Lists all Access groups.
client.ZeroTrust.Access.Groups.Get(ctx, groupID, query) (*
AccessGroupGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/groups/{group_id}
Fetches a single Access group.
client.ZeroTrust.Access.Groups.New(ctx, params) (*
AccessGroupNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/groups
Creates a new Access group.
client.ZeroTrust.Access.Groups.Update(ctx, groupID, params) (*
AccessGroupUpdateResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/groups/{group_id}
Updates a configured Access group.
client.ZeroTrust.Access.Groups.Delete(ctx, groupID, body) (*
AccessGroupDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/access/groups/{group_id}
Deletes an Access group.
Domain types
typeZeroTrustGroupstruct{…}
Zero Trust
Access
Infrastructure
ZeroTrust.Access.Infrastructure
Zero Trust
Access
Infrastructure
Targets
ZeroTrust.Access.Infrastructure.Targets
Methods
client.ZeroTrust.Access.Infrastructure.Targets.List(ctx, params) (*V4PagePaginationArray[
AccessInfrastructureTargetListResponse
], error)
get/accounts/{account_id}/infrastructure/targets
Lists and sorts an account’s targets. Filters are optional and are ANDed together.
client.ZeroTrust.Access.Infrastructure.Targets.Get(ctx, targetID, query) (*
AccessInfrastructureTargetGetResponse
, error)
get/accounts/{account_id}/infrastructure/targets/{target_id}
Get target
client.ZeroTrust.Access.Infrastructure.Targets.New(ctx, params) (*
AccessInfrastructureTargetNewResponse
, error)
post/accounts/{account_id}/infrastructure/targets
Create new target
client.ZeroTrust.Access.Infrastructure.Targets.Update(ctx, targetID, params) (*
AccessInfrastructureTargetUpdateResponse
, error)
put/accounts/{account_id}/infrastructure/targets/{target_id}
Update target
client.ZeroTrust.Access.Infrastructure.Targets.Delete(ctx, targetID, body) error
delete/accounts/{account_id}/infrastructure/targets/{target_id}
Delete target
client.ZeroTrust.Access.Infrastructure.Targets.BulkUpdate(ctx, params) (*SinglePage[
AccessInfrastructureTargetBulkUpdateResponse
], error)
put/accounts/{account_id}/infrastructure/targets/batch
Adds one or more targets.
client.ZeroTrust.Access.Infrastructure.Targets.BulkDelete(ctx, body) error
Deprecated
delete/accounts/{account_id}/infrastructure/targets/batch
Removes one or more targets.
client.ZeroTrust.Access.Infrastructure.Targets.BulkDeleteV2(ctx, params) error
post/accounts/{account_id}/infrastructure/targets/batch_delete
Removes one or more targets.
Zero Trust
Access
Keys
ZeroTrust.Access.Keys
Methods
client.ZeroTrust.Access.Keys.Get(ctx, query) (*
AccessKeyGetResponse
, error)
get/accounts/{account_id}/access/keys
Gets the Access key rotation settings for an account.
client.ZeroTrust.Access.Keys.Update(ctx, params) (*
AccessKeyUpdateResponse
, error)
put/accounts/{account_id}/access/keys
Updates the Access key rotation settings for an account.
client.ZeroTrust.Access.Keys.Rotate(ctx, body) (*
AccessKeyRotateResponse
, error)
post/accounts/{account_id}/access/keys/rotate
Perfoms a key rotation for an account.
Zero Trust
Access
Logs
ZeroTrust.Access.Logs
Zero Trust
Access
Logs
Access Requests
ZeroTrust.Access.Logs.AccessRequests
Methods
client.ZeroTrust.Access.Logs.AccessRequests.List(ctx, params) (*[]
AccessRequest
, error)
get/accounts/{account_id}/access/logs/access_requests
Gets a list of Access authentication audit logs for an account.
Zero Trust
Access
Logs
SCIM
ZeroTrust.Access.Logs.SCIM
Domain types
typeAccessRequeststruct{…}
Zero Trust
Access
Logs
SCIM
Updates
ZeroTrust.Access.Logs.SCIM.Updates
Methods
client.ZeroTrust.Access.Logs.SCIM.Updates.List(ctx, params) (*V4PagePaginationArray[
AccessLogSCIMUpdateListResponse
], error)
get/accounts/{account_id}/access/logs/scim/updates
Lists Access SCIM update logs that maintain a record of updates made to User and Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).
Zero Trust
Access
Policies
ZeroTrust.Access.Policies
Methods
client.ZeroTrust.Access.Policies.List(ctx, params) (*V4PagePaginationArray[
AccessPolicyListResponse
], error)
get/accounts/{account_id}/access/policies
Lists Access reusable policies.
client.ZeroTrust.Access.Policies.Get(ctx, policyID, query) (*
AccessPolicyGetResponse
, error)
get/accounts/{account_id}/access/policies/{policy_id}
Fetches a single Access reusable policy.
client.ZeroTrust.Access.Policies.New(ctx, params) (*
AccessPolicyNewResponse
, error)
post/accounts/{account_id}/access/policies
Creates a new Access reusable policy.
client.ZeroTrust.Access.Policies.Update(ctx, policyID, params) (*
AccessPolicyUpdateResponse
, error)
put/accounts/{account_id}/access/policies/{policy_id}
Updates a Access reusable policy.
client.ZeroTrust.Access.Policies.Delete(ctx, policyID, body) (*
AccessPolicyDeleteResponse
, error)
delete/accounts/{account_id}/access/policies/{policy_id}
Deletes an Access reusable policy.
Domain types
typeApprovalGroupstruct{…}
A group of email addresses that can approve a temporary authentication request.
typePolicystruct{…}
Zero Trust
Access
Service Tokens
ZeroTrust.Access.ServiceTokens
Methods
client.ZeroTrust.Access.ServiceTokens.List(ctx, params) (*V4PagePaginationArray[
ServiceToken
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/service_tokens
Lists all service tokens.
client.ZeroTrust.Access.ServiceTokens.Get(ctx, serviceTokenID, query) (*
ServiceToken
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/service_tokens/{service_token_id}
Fetches a single service token.
client.ZeroTrust.Access.ServiceTokens.New(ctx, params) (*
AccessServiceTokenNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/service_tokens
Generates a new service token. Note: This is the only time you can get the Client Secret. If you lose the Client Secret, you will have to rotate the Client Secret or create a new service token.
client.ZeroTrust.Access.ServiceTokens.Update(ctx, serviceTokenID, params) (*
ServiceToken
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/service_tokens/{service_token_id}
Updates a configured service token.
client.ZeroTrust.Access.ServiceTokens.Delete(ctx, serviceTokenID, body) (*
ServiceToken
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/access/service_tokens/{service_token_id}
Deletes a service token.
client.ZeroTrust.Access.ServiceTokens.Refresh(ctx, serviceTokenID, body) (*
ServiceToken
, error)
post/accounts/{account_id}/access/service_tokens/{service_token_id}/refresh
Refreshes the expiration of a service token.
client.ZeroTrust.Access.ServiceTokens.Rotate(ctx, serviceTokenID, params) (*
AccessServiceTokenRotateResponse
, error)
post/accounts/{account_id}/access/service_tokens/{service_token_id}/rotate
Generates a new Client Secret for a service token and revokes the old one.
Domain types
typeServiceTokenstruct{…}
Zero Trust
Access
Tags
ZeroTrust.Access.Tags
Methods
client.ZeroTrust.Access.Tags.List(ctx, params) (*V4PagePaginationArray[
Tag
], error)
get/accounts/{account_id}/access/tags
List tags
client.ZeroTrust.Access.Tags.Get(ctx, tagName, query) (*
Tag
, error)
get/accounts/{account_id}/access/tags/{tag_name}
Get a tag
client.ZeroTrust.Access.Tags.New(ctx, params) (*
Tag
, error)
post/accounts/{account_id}/access/tags
Create a tag
client.ZeroTrust.Access.Tags.Update(ctx, tagName, params) (*
Tag
, error)
put/accounts/{account_id}/access/tags/{tag_name}
Update a tag
client.ZeroTrust.Access.Tags.Delete(ctx, tagName, body) (*
AccessTagDeleteResponse
, error)
delete/accounts/{account_id}/access/tags/{tag_name}
Delete a tag
Domain types
typeTagstruct{…}
A tag
Zero Trust
Access
Users
ZeroTrust.Access.Users
Methods
client.ZeroTrust.Access.Users.List(ctx, params) (*V4PagePaginationArray[
AccessUserListResponse
], error)
get/accounts/{account_id}/access/users
Gets a list of users for an account.
Domain types
typeAccessUserstruct{…}
Zero Trust
Access
Users
Active Sessions
ZeroTrust.Access.Users.ActiveSessions
Methods
client.ZeroTrust.Access.Users.ActiveSessions.List(ctx, userID, query) (*SinglePage[
AccessUserActiveSessionListResponse
], error)
get/accounts/{account_id}/access/users/{user_id}/active_sessions
Get active sessions for a single user.
client.ZeroTrust.Access.Users.ActiveSessions.Get(ctx, userID, nonce, query) (*
AccessUserActiveSessionGetResponse
, error)
get/accounts/{account_id}/access/users/{user_id}/active_sessions/{nonce}
Get an active session for a single user.
Zero Trust
Access
Users
Failed Logins
ZeroTrust.Access.Users.FailedLogins
Methods
client.ZeroTrust.Access.Users.FailedLogins.List(ctx, userID, query) (*SinglePage[
AccessUserFailedLoginListResponse
], error)
get/accounts/{account_id}/access/users/{user_id}/failed_logins
Get all failed login attempts for a single user.
Zero Trust
Access
Users
Last Seen Identity
ZeroTrust.Access.Users.LastSeenIdentity
Methods
client.ZeroTrust.Access.Users.LastSeenIdentity.Get(ctx, userID, query) (*
Identity
, error)
get/accounts/{account_id}/access/users/{user_id}/last_seen_identity
Get last seen identity for a single user.
Domain types
typeIdentitystruct{…}
Zero Trust
Connectivity Settings
ZeroTrust.ConnectivitySettings
Methods
client.ZeroTrust.ConnectivitySettings.Get(ctx, query) (*
ConnectivitySettingGetResponse
, error)
get/accounts/{account_id}/zerotrust/connectivity_settings
Gets the Zero Trust Connectivity Settings for the given account.
client.ZeroTrust.ConnectivitySettings.Edit(ctx, params) (*
ConnectivitySettingEditResponse
, error)
patch/accounts/{account_id}/zerotrust/connectivity_settings
Updates the Zero Trust Connectivity Settings for the given account.
Zero Trust
Devices
ZeroTrust.Devices
Methods
client.ZeroTrust.Devices.List(ctx, query) (*SinglePage[
Device
], error)
Deprecated
get/accounts/{account_id}/devices
List WARP devices. Not supported when multi-user mode is enabled for the account.
Deprecated: please use one of the following endpoints instead:
GET /accounts/{account_id}/devices/physical-devices
GET /accounts/{account_id}/devices/registrations
client.ZeroTrust.Devices.Get(ctx, deviceID, query) (*
DeviceGetResponse
, error)
Deprecated
get/accounts/{account_id}/devices/{device_id}
Fetches a single WARP device. Not supported when multi-user mode is enabled for the account.
Deprecated: please use one of the following endpoints instead:
GET /accounts/{account_id}/devices/physical-devices/{device_id}
GET /accounts/{account_id}/devices/registrations/{registration_id}
Domain types
typeDevicestruct{…}
Zero Trust
Devices
Devices
ZeroTrust.Devices.Devices
Methods
client.ZeroTrust.Devices.Devices.List(ctx, params) (*CursorPagination[
DeviceDeviceListResponse
], error)
get/accounts/{account_id}/devices/physical-devices
Lists WARP devices.
client.ZeroTrust.Devices.Devices.Get(ctx, deviceID, params) (*
DeviceDeviceGetResponse
, error)
get/accounts/{account_id}/devices/physical-devices/{device_id}
Fetches a single WARP device.
client.ZeroTrust.Devices.Devices.Delete(ctx, deviceID, body) (*
DeviceDeviceDeleteResponse
, error)
delete/accounts/{account_id}/devices/physical-devices/{device_id}
Deletes a WARP device.
client.ZeroTrust.Devices.Devices.Revoke(ctx, deviceID, body) (*
DeviceDeviceRevokeResponse
, error)
post/accounts/{account_id}/devices/physical-devices/{device_id}/revoke
Revokes all WARP registrations associated with the specified device.
Zero Trust
Devices
DEX Tests
ZeroTrust.Devices.DEXTests
Methods
client.ZeroTrust.Devices.DEXTests.List(ctx, query) (*SinglePage[
DeviceDEXTestListResponse
], error)
get/accounts/{account_id}/dex/devices/dex_tests
Fetch all DEX tests
client.ZeroTrust.Devices.DEXTests.Get(ctx, dexTestID, query) (*
DeviceDEXTestGetResponse
, error)
get/accounts/{account_id}/dex/devices/dex_tests/{dex_test_id}
Fetch a single DEX test.
client.ZeroTrust.Devices.DEXTests.New(ctx, params) (*
DeviceDEXTestNewResponse
, error)
post/accounts/{account_id}/dex/devices/dex_tests
Create a DEX test.
client.ZeroTrust.Devices.DEXTests.Update(ctx, dexTestID, params) (*
DeviceDEXTestUpdateResponse
, error)
put/accounts/{account_id}/dex/devices/dex_tests/{dex_test_id}
Update a DEX test.
client.ZeroTrust.Devices.DEXTests.Delete(ctx, dexTestID, body) (*
DeviceDEXTestDeleteResponse
, error)
delete/accounts/{account_id}/dex/devices/dex_tests/{dex_test_id}
Delete a Device DEX test. Returns the remaining device dex tests for the account.
Domain types
typeSchemaDatastruct{…}
The configuration object which contains the details for the WARP client to conduct the test.
typeSchemaHTTPstruct{…}
Zero Trust
Devices
Fleet Status
ZeroTrust.Devices.FleetStatus
Methods
client.ZeroTrust.Devices.FleetStatus.Get(ctx, deviceID, params) (*
DeviceFleetStatusGetResponse
, error)
get/accounts/{account_id}/dex/devices/{device_id}/fleet-status/live
Get the live status of a latest device given device_id from the device_state table
Zero Trust
Devices
Networks
ZeroTrust.Devices.Networks
Methods
client.ZeroTrust.Devices.Networks.List(ctx, query) (*SinglePage[
DeviceNetwork
], error)
get/accounts/{account_id}/devices/networks
Fetches a list of managed networks for an account.
client.ZeroTrust.Devices.Networks.Get(ctx, networkID, query) (*
DeviceNetwork
, error)
get/accounts/{account_id}/devices/networks/{network_id}
Fetches details for a single managed network.
client.ZeroTrust.Devices.Networks.New(ctx, params) (*
DeviceNetwork
, error)
post/accounts/{account_id}/devices/networks
Creates a new device managed network.
client.ZeroTrust.Devices.Networks.Update(ctx, networkID, params) (*
DeviceNetwork
, error)
put/accounts/{account_id}/devices/networks/{network_id}
Updates a configured device managed network.
client.ZeroTrust.Devices.Networks.Delete(ctx, networkID, body) (*SinglePage[
DeviceNetwork
], error)
delete/accounts/{account_id}/devices/networks/{network_id}
Deletes a device managed network and fetches a list of the remaining device managed networks for an account.
Domain types
typeDeviceNetworkstruct{…}
Zero Trust
Devices
Override Codes
ZeroTrust.Devices.OverrideCodes
Methods
client.ZeroTrust.Devices.OverrideCodes.List(ctx, deviceID, query) (*SinglePage[
DeviceOverrideCodeListResponse
], error)
Deprecated
get/accounts/{account_id}/devices/{device_id}/override_codes
Fetches a one-time use admin override code for a device. This relies on the Admin Override setting being enabled in your device configuration. Not supported when multi-user mode is enabled for the account. Deprecated: please use GET /accounts/{account_id}/devices/registrations/{registration_id}/override_codes instead.
client.ZeroTrust.Devices.OverrideCodes.Get(ctx, registrationID, query) (*
DeviceOverrideCodeGetResponse
, error)
get/accounts/{account_id}/devices/registrations/{registration_id}/override_codes
Fetches one-time use admin override codes for a registration. This relies on the Admin Override setting being enabled in your device configuration.
Zero Trust
Devices
Policies
ZeroTrust.Devices.Policies
Domain types
typeDevicePolicyCertificatesstruct{…}
typeFallbackDomainstruct{…}
typeFallbackDomainPolicy[]
FallbackDomain
typeSettingsPolicystruct{…}
typeSplitTunnelExcludeinterface{…}
typeSplitTunnelIncludeinterface{…}
Zero Trust
Devices
Policies
Custom
ZeroTrust.Devices.Policies.Custom
Methods
client.ZeroTrust.Devices.Policies.Custom.List(ctx, query) (*SinglePage[
SettingsPolicy
], error)
get/accounts/{account_id}/devices/policies
Fetches a list of the device settings profiles for an account.
client.ZeroTrust.Devices.Policies.Custom.Get(ctx, policyID, query) (*
SettingsPolicy
, error)
get/accounts/{account_id}/devices/policy/{policy_id}
Fetches a device settings profile by ID.
client.ZeroTrust.Devices.Policies.Custom.New(ctx, params) (*
SettingsPolicy
, error)
post/accounts/{account_id}/devices/policy
Creates a device settings profile to be applied to certain devices matching the criteria.
client.ZeroTrust.Devices.Policies.Custom.Edit(ctx, policyID, params) (*
SettingsPolicy
, error)
patch/accounts/{account_id}/devices/policy/{policy_id}
Updates a configured device settings profile.
client.ZeroTrust.Devices.Policies.Custom.Delete(ctx, policyID, body) (*SinglePage[
SettingsPolicy
], error)
delete/accounts/{account_id}/devices/policy/{policy_id}
Deletes a device settings profile and fetches a list of the remaining profiles for an account.
Zero Trust
Devices
Policies
Custom
Excludes
ZeroTrust.Devices.Policies.Custom.Excludes
Methods
client.ZeroTrust.Devices.Policies.Custom.Excludes.Get(ctx, policyID, query) (*SinglePage[
SplitTunnelExclude
], error)
get/accounts/{account_id}/devices/policy/{policy_id}/exclude
Fetches the list of routes excluded from the WARP client's tunnel for a specific device settings profile.
client.ZeroTrust.Devices.Policies.Custom.Excludes.Update(ctx, policyID, params) (*SinglePage[
SplitTunnelExclude
], error)
put/accounts/{account_id}/devices/policy/{policy_id}/exclude
Sets the list of routes excluded from the WARP client's tunnel for a specific device settings profile.
Zero Trust
Devices
Policies
Custom
Fallback Domains
ZeroTrust.Devices.Policies.Custom.FallbackDomains
Methods
client.ZeroTrust.Devices.Policies.Custom.FallbackDomains.Get(ctx, policyID, query) (*SinglePage[
FallbackDomain
], error)
get/accounts/{account_id}/devices/policy/{policy_id}/fallback_domains
Fetches the list of domains to bypass Gateway DNS resolution from a specified device settings profile. These domains will use the specified local DNS resolver instead.
client.ZeroTrust.Devices.Policies.Custom.FallbackDomains.Update(ctx, policyID, params) (*SinglePage[
FallbackDomain
], error)
put/accounts/{account_id}/devices/policy/{policy_id}/fallback_domains
Sets the list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead. This will only apply to the specified device settings profile.
Zero Trust
Devices
Policies
Custom
Includes
ZeroTrust.Devices.Policies.Custom.Includes
Methods
client.ZeroTrust.Devices.Policies.Custom.Includes.Get(ctx, policyID, query) (*SinglePage[
SplitTunnelInclude
], error)
get/accounts/{account_id}/devices/policy/{policy_id}/include
Fetches the list of routes included in the WARP client's tunnel for a specific device settings profile.
client.ZeroTrust.Devices.Policies.Custom.Includes.Update(ctx, policyID, params) (*SinglePage[
SplitTunnelInclude
], error)
put/accounts/{account_id}/devices/policy/{policy_id}/include
Sets the list of routes included in the WARP client's tunnel for a specific device settings profile.
Zero Trust
Devices
Policies
Default
ZeroTrust.Devices.Policies.Default
Methods
client.ZeroTrust.Devices.Policies.Default.Get(ctx, query) (*
DevicePolicyDefaultGetResponse
, error)
get/accounts/{account_id}/devices/policy
Fetches the default device settings profile for an account.
client.ZeroTrust.Devices.Policies.Default.Edit(ctx, params) (*
DevicePolicyDefaultEditResponse
, error)
patch/accounts/{account_id}/devices/policy
Updates the default device settings profile for an account.
Zero Trust
Devices
Policies
Default
Certificates
ZeroTrust.Devices.Policies.Default.Certificates
Methods
client.ZeroTrust.Devices.Policies.Default.Certificates.Get(ctx, query) (*
DevicePolicyCertificates
, error)
get/zones/{zone_id}/devices/policy/certificates
Fetches device certificate provisioning.
client.ZeroTrust.Devices.Policies.Default.Certificates.Edit(ctx, params) (*
DevicePolicyCertificates
, error)
patch/zones/{zone_id}/devices/policy/certificates
Enable Zero Trust Clients to provision a certificate, containing a x509 subject, and referenced by Access device posture policies when the client visits MTLS protected domains. This facilitates device posture without a WARP session.
Zero Trust
Devices
Policies
Default
Excludes
ZeroTrust.Devices.Policies.Default.Excludes
Methods
client.ZeroTrust.Devices.Policies.Default.Excludes.Get(ctx, query) (*SinglePage[
SplitTunnelExclude
], error)
get/accounts/{account_id}/devices/policy/exclude
Fetches the list of routes excluded from the WARP client's tunnel.
client.ZeroTrust.Devices.Policies.Default.Excludes.Update(ctx, params) (*SinglePage[
SplitTunnelExclude
], error)
put/accounts/{account_id}/devices/policy/exclude
Sets the list of routes excluded from the WARP client's tunnel.
Zero Trust
Devices
Policies
Default
Fallback Domains
ZeroTrust.Devices.Policies.Default.FallbackDomains
Methods
client.ZeroTrust.Devices.Policies.Default.FallbackDomains.Get(ctx, query) (*SinglePage[
FallbackDomain
], error)
get/accounts/{account_id}/devices/policy/fallback_domains
Fetches a list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead.
client.ZeroTrust.Devices.Policies.Default.FallbackDomains.Update(ctx, params) (*SinglePage[
FallbackDomain
], error)
put/accounts/{account_id}/devices/policy/fallback_domains
Sets the list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead.
Zero Trust
Devices
Policies
Default
Includes
ZeroTrust.Devices.Policies.Default.Includes
Methods
client.ZeroTrust.Devices.Policies.Default.Includes.Get(ctx, query) (*SinglePage[
SplitTunnelInclude
], error)
get/accounts/{account_id}/devices/policy/include
Fetches the list of routes included in the WARP client's tunnel.
client.ZeroTrust.Devices.Policies.Default.Includes.Update(ctx, params) (*SinglePage[
SplitTunnelInclude
], error)
put/accounts/{account_id}/devices/policy/include
Sets the list of routes included in the WARP client's tunnel.
Zero Trust
Devices
Posture
ZeroTrust.Devices.Posture
Methods
client.ZeroTrust.Devices.Posture.List(ctx, query) (*SinglePage[
DevicePostureRule
], error)
get/accounts/{account_id}/devices/posture
Fetches device posture rules for a Zero Trust account.
client.ZeroTrust.Devices.Posture.Get(ctx, ruleID, query) (*
DevicePostureRule
, error)
get/accounts/{account_id}/devices/posture/{rule_id}
Fetches a single device posture rule.
client.ZeroTrust.Devices.Posture.New(ctx, params) (*
DevicePostureRule
, error)
post/accounts/{account_id}/devices/posture
Creates a new device posture rule.
client.ZeroTrust.Devices.Posture.Update(ctx, ruleID, params) (*
DevicePostureRule
, error)
put/accounts/{account_id}/devices/posture/{rule_id}
Updates a device posture rule.
client.ZeroTrust.Devices.Posture.Delete(ctx, ruleID, body) (*
DevicePostureDeleteResponse
, error)
delete/accounts/{account_id}/devices/posture/{rule_id}
Deletes a device posture rule.
Domain types
typeCarbonblackInputstring
typeClientCertificateInputstruct{…}
typeCrowdstrikeInputstruct{…}
typeDeviceInputinterface{…}
The value to be checked against.
typeDeviceMatchstruct{…}
typeDevicePostureRulestruct{…}
typeDiskEncryptionInputstruct{…}
typeDomainJoinedInputstruct{…}
typeFileInputstruct{…}
typeFirewallInputstruct{…}
typeIntuneInputstruct{…}
typeKolideInputstruct{…}
typeOSVersionInputstruct{…}
typeSentineloneInputstruct{…}
typeSentineloneS2sInputstruct{…}
typeTaniumInputstruct{…}
typeUniqueClientIDInputstruct{…}
typeWorkspaceOneInputstruct{…}
Zero Trust
Devices
Posture
Integrations
ZeroTrust.Devices.Posture.Integrations
Methods
client.ZeroTrust.Devices.Posture.Integrations.List(ctx, query) (*SinglePage[
Integration
], error)
get/accounts/{account_id}/devices/posture/integration
Fetches the list of device posture integrations for an account.
client.ZeroTrust.Devices.Posture.Integrations.Get(ctx, integrationID, query) (*
Integration
, error)
get/accounts/{account_id}/devices/posture/integration/{integration_id}
Fetches details for a single device posture integration.
client.ZeroTrust.Devices.Posture.Integrations.New(ctx, params) (*
Integration
, error)
post/accounts/{account_id}/devices/posture/integration
Create a new device posture integration.
client.ZeroTrust.Devices.Posture.Integrations.Edit(ctx, integrationID, params) (*
Integration
, error)
patch/accounts/{account_id}/devices/posture/integration/{integration_id}
Updates a configured device posture integration.
client.ZeroTrust.Devices.Posture.Integrations.Delete(ctx, integrationID, body) (*unknown, error)
delete/accounts/{account_id}/devices/posture/integration/{integration_id}
Delete a configured device posture integration.
Domain types
typeIntegrationstruct{…}
Zero Trust
Devices
Registrations
ZeroTrust.Devices.Registrations
Methods
client.ZeroTrust.Devices.Registrations.List(ctx, params) (*CursorPagination[
DeviceRegistrationListResponse
], error)
get/accounts/{account_id}/devices/registrations
Lists WARP registrations.
client.ZeroTrust.Devices.Registrations.Get(ctx, registrationID, params) (*
DeviceRegistrationGetResponse
, error)
get/accounts/{account_id}/devices/registrations/{registration_id}
Fetches a single WARP registration.
client.ZeroTrust.Devices.Registrations.Delete(ctx, registrationID, body) (*
DeviceRegistrationDeleteResponse
, error)
delete/accounts/{account_id}/devices/registrations/{registration_id}
Deletes a WARP registration.
client.ZeroTrust.Devices.Registrations.BulkDelete(ctx, params) (*
DeviceRegistrationBulkDeleteResponse
, error)
delete/accounts/{account_id}/devices/registrations
Deletes a list of WARP registrations.
client.ZeroTrust.Devices.Registrations.Revoke(ctx, params) (*
DeviceRegistrationRevokeResponse
, error)
post/accounts/{account_id}/devices/registrations/revoke
Revokes a list of WARP registrations.
client.ZeroTrust.Devices.Registrations.Unrevoke(ctx, params) (*
DeviceRegistrationUnrevokeResponse
, error)
post/accounts/{account_id}/devices/registrations/unrevoke
Unrevokes a list of WARP registrations.
Zero Trust
Devices
Resilience
ZeroTrust.Devices.Resilience
Zero Trust
Devices
Resilience
Global WARP Override
ZeroTrust.Devices.Resilience.GlobalWARPOverride
Methods
client.ZeroTrust.Devices.Resilience.GlobalWARPOverride.Get(ctx, query) (*
DeviceResilienceGlobalWARPOverrideGetResponse
, error)
get/accounts/{account_id}/devices/resilience/disconnect
Fetch the Global WARP override state.
client.ZeroTrust.Devices.Resilience.GlobalWARPOverride.New(ctx, params) (*
DeviceResilienceGlobalWARPOverrideNewResponse
, error)
post/accounts/{account_id}/devices/resilience/disconnect
Sets the Global WARP override state.
Zero Trust
Devices
Revoke
ZeroTrust.Devices.Revoke
Methods
client.ZeroTrust.Devices.Revoke.New(ctx, params) (*unknown, error)
Deprecated
post/accounts/{account_id}/devices/revoke
Revokes a list of devices. Not supported when multi-user mode is enabled.
Deprecated: please use POST /accounts/{account_id}/devices/registrations/revoke instead.
Zero Trust
Devices
Settings
ZeroTrust.Devices.Settings
Methods
client.ZeroTrust.Devices.Settings.Get(ctx, query) (*
DeviceSettings
, error)
get/accounts/{account_id}/devices/settings
Describes the current device settings for a Zero Trust account.
client.ZeroTrust.Devices.Settings.Update(ctx, params) (*
DeviceSettings
, error)
put/accounts/{account_id}/devices/settings
Updates the current device settings for a Zero Trust account.
client.ZeroTrust.Devices.Settings.Edit(ctx, params) (*
DeviceSettings
, error)
patch/accounts/{account_id}/devices/settings
Patches the current device settings for a Zero Trust account.
client.ZeroTrust.Devices.Settings.Delete(ctx, body) (*
DeviceSettings
, error)
delete/accounts/{account_id}/devices/settings
Resets the current device settings for a Zero Trust account.
Domain types
typeDeviceSettingsstruct{…}
Zero Trust
Devices
Unrevoke
ZeroTrust.Devices.Unrevoke
Methods
client.ZeroTrust.Devices.Unrevoke.New(ctx, params) (*unknown, error)
Deprecated
post/accounts/{account_id}/devices/unrevoke
Unrevokes a list of devices. Not supported when multi-user mode is enabled.
Deprecated: please use POST /accounts/{account_id}/devices/registrations/unrevoke instead.
Zero Trust
DEX
ZeroTrust.DEX
Domain types
typeDigitalExperienceMonitorstruct{…}
typeNetworkPathstruct{…}
typeNetworkPathResponsestruct{…}
typePercentilesstruct{…}
Zero Trust
DEX
Colos
ZeroTrust.DEX.Colos
Methods
client.ZeroTrust.DEX.Colos.List(ctx, params) (*SinglePage[
DEXColoListResponse
], error)
get/accounts/{account_id}/dex/colos
List Cloudflare colos that account's devices were connected to during a time period, sorted by usage starting from the most used colo. Colos without traffic are also returned and sorted alphabetically.
Zero Trust
DEX
Commands
ZeroTrust.DEX.Commands
Methods
client.ZeroTrust.DEX.Commands.List(ctx, params) (*V4PagePagination[
DEXCommandListResponse
], error)
get/accounts/{account_id}/dex/commands
Retrieves a paginated list of commands issued to devices under the specified account, optionally filtered by time range, device, or other parameters
client.ZeroTrust.DEX.Commands.New(ctx, params) (*
DEXCommandNewResponse
, error)
post/accounts/{account_id}/dex/commands
Initiate commands for up to 10 devices per account
Zero Trust
DEX
Commands
Devices
ZeroTrust.DEX.Commands.Devices
Methods
client.ZeroTrust.DEX.Commands.Devices.List(ctx, params) (*V4PagePagination[
DEXCommandDeviceListResponse
], error)
get/accounts/{account_id}/dex/commands/devices
List devices with WARP client support for remote captures which have been connected in the last 1 hour.
Zero Trust
DEX
Commands
Downloads
ZeroTrust.DEX.Commands.Downloads
Methods
client.ZeroTrust.DEX.Commands.Downloads.Get(ctx, commandID, filename, query) (*
Response
, error)
get/accounts/{account_id}/dex/commands/{command_id}/downloads/{filename}
Downloads artifacts for an executed command. Bulk downloads are not supported
Zero Trust
DEX
Commands
Quota
ZeroTrust.DEX.Commands.Quota
Methods
client.ZeroTrust.DEX.Commands.Quota.Get(ctx, query) (*
DEXCommandQuotaGetResponse
, error)
get/accounts/{account_id}/dex/commands/quota
Retrieves the current quota usage and limits for device commands within a specific account, including the time when the quota will reset
Zero Trust
DEX
Fleet Status
ZeroTrust.DEX.FleetStatus
Methods
client.ZeroTrust.DEX.FleetStatus.Live(ctx, params) (*
DEXFleetStatusLiveResponse
, error)
get/accounts/{account_id}/dex/fleet-status/live
List details for live (up to 60 minutes) devices using WARP
client.ZeroTrust.DEX.FleetStatus.OverTime(ctx, params) error
get/accounts/{account_id}/dex/fleet-status/over-time
List details for devices using WARP, up to 7 days
Domain types
typeLiveStatstruct{…}
Zero Trust
DEX
Fleet Status
Devices
ZeroTrust.DEX.FleetStatus.Devices
Methods
client.ZeroTrust.DEX.FleetStatus.Devices.List(ctx, params) (*V4PagePaginationArray[
DEXFleetStatusDeviceListResponse
], error)
get/accounts/{account_id}/dex/fleet-status/devices
List details for devices using WARP
Zero Trust
DEX
HTTP Tests
ZeroTrust.DEX.HTTPTests
Methods
client.ZeroTrust.DEX.HTTPTests.Get(ctx, testID, params) (*
HTTPDetails
, error)
get/accounts/{account_id}/dex/http-tests/{test_id}
Get test details and aggregate performance metrics for an http test for a given time period between 1 hour and 7 days.
Domain types
typeHTTPDetailsstruct{…}
Zero Trust
DEX
HTTP Tests
Percentiles
ZeroTrust.DEX.HTTPTests.Percentiles
Methods
client.ZeroTrust.DEX.HTTPTests.Percentiles.Get(ctx, testID, params) (*
HTTPDetailsPercentiles
, error)
get/accounts/{account_id}/dex/http-tests/{test_id}/percentiles
Get percentiles for an http test for a given time period between 1 hour and 7 days.
Domain types
typeHTTPDetailsPercentilesstruct{…}
typeTestStatOverTimestruct{…}
Zero Trust
DEX
Tests
ZeroTrust.DEX.Tests
Methods
client.ZeroTrust.DEX.Tests.List(ctx, params) (*V4PagePagination[
Tests
], error)
get/accounts/{account_id}/dex/tests/overview
List DEX tests with overview metrics
Domain types
typeAggregateTimePeriodstruct{…}
typeTestsstruct{…}
Zero Trust
DEX
Tests
Unique Devices
ZeroTrust.DEX.Tests.UniqueDevices
Methods
client.ZeroTrust.DEX.Tests.UniqueDevices.List(ctx, params) (*
UniqueDevices
, error)
get/accounts/{account_id}/dex/tests/unique-devices
Returns unique count of devices that have run synthetic application monitoring tests in the past 7 days.
Domain types
typeUniqueDevicesstruct{…}
Zero Trust
DEX
Traceroute Test Results
ZeroTrust.DEX.TracerouteTestResults
Zero Trust
DEX
Traceroute Test Results
Network Path
ZeroTrust.DEX.TracerouteTestResults.NetworkPath
Methods
client.ZeroTrust.DEX.TracerouteTestResults.NetworkPath.Get(ctx, testResultID, query) (*
DEXTracerouteTestResultNetworkPathGetResponse
, error)
get/accounts/{account_id}/dex/traceroute-test-results/{test_result_id}/network-path
Get a breakdown of hops and performance metrics for a specific traceroute test run
Zero Trust
DEX
Traceroute Tests
ZeroTrust.DEX.TracerouteTests
Methods
client.ZeroTrust.DEX.TracerouteTests.Get(ctx, testID, params) (*
Traceroute
, error)
get/accounts/{account_id}/dex/traceroute-tests/{test_id}
Get test details and aggregate performance metrics for an traceroute test for a given time period between 1 hour and 7 days.
client.ZeroTrust.DEX.TracerouteTests.Percentiles(ctx, testID, params) (*
DEXTracerouteTestPercentilesResponse
, error)
get/accounts/{account_id}/dex/traceroute-tests/{test_id}/percentiles
Get percentiles for a traceroute test for a given time period between 1 hour and 7 days.
client.ZeroTrust.DEX.TracerouteTests.NetworkPath(ctx, testID, params) (*
NetworkPathResponse
, error)
get/accounts/{account_id}/dex/traceroute-tests/{test_id}/network-path
Get a breakdown of metrics by hop for individual traceroute test runs
Domain types
typeTraceroutestruct{…}
Zero Trust
DEX
WARP Change Events
ZeroTrust.DEX.WARPChangeEvents
Methods
client.ZeroTrust.DEX.WARPChangeEvents.Get(ctx, params) (*[]
DexwarpChangeEventGetResponse
, error)
get/accounts/{account_id}/dex/warp-change-events
List WARP configuration and enablement toggle change events by device.
Zero Trust
DLP
ZeroTrust.DLP
Zero Trust
DLP
Datasets
ZeroTrust.DLP.Datasets
Methods
client.ZeroTrust.DLP.Datasets.List(ctx, query) (*SinglePage[
Dataset
], error)
get/accounts/{account_id}/dlp/datasets
Fetch all datasets
client.ZeroTrust.DLP.Datasets.Get(ctx, datasetID, query) (*
DLPDatasetGetResponse
, error)
get/accounts/{account_id}/dlp/datasets/{dataset_id}
Fetch a specific dataset
client.ZeroTrust.DLP.Datasets.New(ctx, params) (*
DLPDatasetNewResponse
, error)
post/accounts/{account_id}/dlp/datasets
Create a new dataset
client.ZeroTrust.DLP.Datasets.Update(ctx, datasetID, params) (*
DLPDatasetUpdateResponse
, error)
put/accounts/{account_id}/dlp/datasets/{dataset_id}
Update details about a dataset
client.ZeroTrust.DLP.Datasets.Delete(ctx, datasetID, body) error
delete/accounts/{account_id}/dlp/datasets/{dataset_id}
This deletes all versions of the dataset.
Domain types
typeDatasetstruct{…}
typeDatasetArray[]
Dataset
typeDatasetCreationstruct{…}
Zero Trust
DLP
Datasets
Upload
ZeroTrust.DLP.Datasets.Upload
Methods
client.ZeroTrust.DLP.Datasets.Upload.New(ctx, datasetID, body) (*
DLPDatasetUploadNewResponse
, error)
post/accounts/{account_id}/dlp/datasets/{dataset_id}/upload
Prepare to upload a new version of a dataset
client.ZeroTrust.DLP.Datasets.Upload.Edit(ctx, datasetID, version, dataset, body) (*
DLPDatasetUploadEditResponse
, error)
post/accounts/{account_id}/dlp/datasets/{dataset_id}/upload/{version}
This is used for single-column EDMv1 and Custom Word Lists. The EDM format can only be created in the Cloudflare dashboard. For other clients, this operation can only be used for non-secret Custom Word Lists. The body must be a UTF-8 encoded, newline (NL or CRNL) separated list of words to be matched.
Domain types
typeNewVersionstruct{…}
Zero Trust
DLP
Datasets
Versions
ZeroTrust.DLP.Datasets.Versions
Methods
client.ZeroTrust.DLP.Datasets.Versions.New(ctx, datasetID, version, params) (*SinglePage[
DLPDatasetVersionNewResponse
], error)
post/accounts/{account_id}/dlp/datasets/{dataset_id}/versions/{version}
This is used for multi-column EDMv2 datasets. The EDMv2 format can only be created in the Cloudflare dashboard. The columns in the response appear in the same order as in the request.
Zero Trust
DLP
Datasets
Versions
Entries
ZeroTrust.DLP.Datasets.Versions.Entries
Methods
client.ZeroTrust.DLP.Datasets.Versions.Entries.New(ctx, datasetID, version, entryID, datasetVersionEntry, body) (*
DLPDatasetVersionEntryNewResponse
, error)
post/accounts/{account_id}/dlp/datasets/{dataset_id}/versions/{version}/entries/{entry_id}
This is used for multi-column EDMv2 datasets. The EDMv2 format can only be created in the Cloudflare dashboard.
Zero Trust
DLP
Email
ZeroTrust.DLP.Email
Zero Trust
DLP
Email
Account Mapping
ZeroTrust.DLP.Email.AccountMapping
Methods
client.ZeroTrust.DLP.Email.AccountMapping.Get(ctx, query) (*
DLPEmailAccountMappingGetResponse
, error)
get/accounts/{account_id}/dlp/email/account_mapping
Get mapping
client.ZeroTrust.DLP.Email.AccountMapping.New(ctx, params) (*
DLPEmailAccountMappingNewResponse
, error)
post/accounts/{account_id}/dlp/email/account_mapping
Create mapping
Zero Trust
DLP
Email
Rules
ZeroTrust.DLP.Email.Rules
Methods
client.ZeroTrust.DLP.Email.Rules.List(ctx, query) (*SinglePage[
DLPEmailRuleListResponse
], error)
get/accounts/{account_id}/dlp/email/rules
Lists all email scanner rules for an account.
client.ZeroTrust.DLP.Email.Rules.Get(ctx, ruleID, query) (*
DLPEmailRuleGetResponse
, error)
get/accounts/{account_id}/dlp/email/rules/{rule_id}
Get an email scanner rule
client.ZeroTrust.DLP.Email.Rules.New(ctx, params) (*
DLPEmailRuleNewResponse
, error)
post/accounts/{account_id}/dlp/email/rules
Create email scanner rule
client.ZeroTrust.DLP.Email.Rules.Update(ctx, ruleID, params) (*
DLPEmailRuleUpdateResponse
, error)
put/accounts/{account_id}/dlp/email/rules/{rule_id}
Update email scanner rule
client.ZeroTrust.DLP.Email.Rules.Delete(ctx, ruleID, body) (*
DLPEmailRuleDeleteResponse
, error)
delete/accounts/{account_id}/dlp/email/rules/{rule_id}
Delete email scanner rule
client.ZeroTrust.DLP.Email.Rules.BulkEdit(ctx, params) (*
DLPEmailRuleBulkEditResponse
, error)
patch/accounts/{account_id}/dlp/email/rules
Update email scanner rule priorities
Zero Trust
DLP
Entries
ZeroTrust.DLP.Entries
Methods
client.ZeroTrust.DLP.Entries.List(ctx, query) (*SinglePage[
DLPEntryListResponse
], error)
get/accounts/{account_id}/dlp/entries
Lists all DLP entries in an account.
client.ZeroTrust.DLP.Entries.Get(ctx, entryID, query) (*
DLPEntryGetResponse
, error)
get/accounts/{account_id}/dlp/entries/{entry_id}
Fetches a DLP entry by ID.
client.ZeroTrust.DLP.Entries.New(ctx, params) (*
DLPEntryNewResponse
, error)
post/accounts/{account_id}/dlp/entries
Creates a DLP custom entry.
client.ZeroTrust.DLP.Entries.Update(ctx, entryID, params) (*
DLPEntryUpdateResponse
, error)
put/accounts/{account_id}/dlp/entries/{entry_id}
Updates a DLP entry.
client.ZeroTrust.DLP.Entries.Delete(ctx, entryID, body) (*
DLPEntryDeleteResponse
, error)
delete/accounts/{account_id}/dlp/entries/{entry_id}
Deletes a DLP custom entry.
Zero Trust
DLP
Entries
Custom
ZeroTrust.DLP.Entries.Custom
Methods
client.ZeroTrust.DLP.Entries.Custom.New(ctx, params) (*
DLPEntryCustomNewResponse
, error)
post/accounts/{account_id}/dlp/entries
Creates a DLP custom entry.
client.ZeroTrust.DLP.Entries.Custom.Update(ctx, entryID, params) (*
DLPEntryCustomUpdateResponse
, error)
put/accounts/{account_id}/dlp/entries/custom/{entry_id}
Updates a DLP custom entry.
client.ZeroTrust.DLP.Entries.Custom.Delete(ctx, entryID, body) (*
DLPEntryCustomDeleteResponse
, error)
delete/accounts/{account_id}/dlp/entries/{entry_id}
Deletes a DLP custom entry.
client.ZeroTrust.DLP.Entries.Custom.Get(ctx, entryID, query) (*
DLPEntryCustomGetResponse
, error)
get/accounts/{account_id}/dlp/entries/{entry_id}
Fetches a DLP entry by ID.
client.ZeroTrust.DLP.Entries.Custom.List(ctx, query) (*SinglePage[
DLPEntryCustomListResponse
], error)
get/accounts/{account_id}/dlp/entries
Lists all DLP entries in an account.
Zero Trust
DLP
Entries
Integration
ZeroTrust.DLP.Entries.Integration
Methods
client.ZeroTrust.DLP.Entries.Integration.New(ctx, params) (*
DLPEntryIntegrationNewResponse
, error)
post/accounts/{account_id}/dlp/entries/integration
Integration entries can't be created, this will update an existing integration entry. This is needed for our generated terraform API.
client.ZeroTrust.DLP.Entries.Integration.Update(ctx, entryID, params) (*
DLPEntryIntegrationUpdateResponse
, error)
put/accounts/{account_id}/dlp/entries/integration/{entry_id}
Updates a DLP entry.
client.ZeroTrust.DLP.Entries.Integration.Delete(ctx, entryID, body) (*
DLPEntryIntegrationDeleteResponse
, error)
delete/accounts/{account_id}/dlp/entries/integration/{entry_id}
This is a no-op as integration entires can't be deleted but is needed for our generated terraform API.
client.ZeroTrust.DLP.Entries.Integration.Get(ctx, entryID, query) (*
DLPEntryIntegrationGetResponse
, error)
get/accounts/{account_id}/dlp/entries/{entry_id}
Fetches a DLP entry by ID.
client.ZeroTrust.DLP.Entries.Integration.List(ctx, query) (*SinglePage[
DLPEntryIntegrationListResponse
], error)
get/accounts/{account_id}/dlp/entries
Lists all DLP entries in an account.
Zero Trust
DLP
Entries
Predefined
ZeroTrust.DLP.Entries.Predefined
Methods
client.ZeroTrust.DLP.Entries.Predefined.New(ctx, params) (*
DLPEntryPredefinedNewResponse
, error)
post/accounts/{account_id}/dlp/entries/predefined
Predefined entries can't be created, this will update an existing predefined entry. This is needed for our generated terraform API.
client.ZeroTrust.DLP.Entries.Predefined.Update(ctx, entryID, params) (*
DLPEntryPredefinedUpdateResponse
, error)
put/accounts/{account_id}/dlp/entries/predefined/{entry_id}
Updates a DLP entry.
client.ZeroTrust.DLP.Entries.Predefined.Delete(ctx, entryID, body) (*
DLPEntryPredefinedDeleteResponse
, error)
delete/accounts/{account_id}/dlp/entries/predefined/{entry_id}
This is a no-op as predefined entires can't be deleted but is needed for our generated terraform API.
client.ZeroTrust.DLP.Entries.Predefined.Get(ctx, entryID, query) (*
DLPEntryPredefinedGetResponse
, error)
get/accounts/{account_id}/dlp/entries/{entry_id}
Fetches a DLP entry by ID.
client.ZeroTrust.DLP.Entries.Predefined.List(ctx, query) (*SinglePage[
DLPEntryPredefinedListResponse
], error)
get/accounts/{account_id}/dlp/entries
Lists all DLP entries in an account.
Zero Trust
DLP
Limits
ZeroTrust.DLP.Limits
Methods
client.ZeroTrust.DLP.Limits.List(ctx, query) (*
DLPLimitListResponse
, error)
get/accounts/{account_id}/dlp/limits
Fetch limits associated with DLP for account
Zero Trust
DLP
Patterns
ZeroTrust.DLP.Patterns
Methods
client.ZeroTrust.DLP.Patterns.Validate(ctx, params) (*
DLPPatternValidateResponse
, error)
post/accounts/{account_id}/dlp/patterns/validate
Validates whether this pattern is a valid regular expression. Rejects it if the regular expression is too complex or can match an unbounded-length string. The regex will be rejected if it uses * or +. Bound the maximum number of characters that can be matched using a range, e.g. {1,100}.
Zero Trust
DLP
Payload Logs
ZeroTrust.DLP.PayloadLogs
Methods
client.ZeroTrust.DLP.PayloadLogs.Get(ctx, query) (*
DLPPayloadLogGetResponse
, error)
get/accounts/{account_id}/dlp/payload_log
Get payload log settings
client.ZeroTrust.DLP.PayloadLogs.Update(ctx, params) (*
DLPPayloadLogUpdateResponse
, error)
put/accounts/{account_id}/dlp/payload_log
Set payload log settings
Zero Trust
DLP
Profiles
ZeroTrust.DLP.Profiles
Methods
client.ZeroTrust.DLP.Profiles.List(ctx, params) (*SinglePage[
Profile
], error)
get/accounts/{account_id}/dlp/profiles
Lists all DLP profiles in an account.
client.ZeroTrust.DLP.Profiles.Get(ctx, profileID, query) (*
DLPProfileGetResponse
, error)
get/accounts/{account_id}/dlp/profiles/{profile_id}
Fetches a DLP profile by ID.
Domain types
typeContextAwarenessstruct{…}
Scan the context of predefined entries to only return matches surrounded by keywords.
typeProfileinterface{…}
typeSkipConfigurationstruct{…}
Content types to exclude from context analysis and return all matches.
Zero Trust
DLP
Profiles
Custom
ZeroTrust.DLP.Profiles.Custom
Methods
client.ZeroTrust.DLP.Profiles.Custom.Get(ctx, profileID, query) (*
DLPProfileCustomGetResponse
, error)
get/accounts/{account_id}/dlp/profiles/custom/{profile_id}
Fetches a custom DLP profile by id.
client.ZeroTrust.DLP.Profiles.Custom.New(ctx, params) (*
DLPProfileCustomNewResponse
, error)
post/accounts/{account_id}/dlp/profiles/custom
Creates a DLP custom profile.
client.ZeroTrust.DLP.Profiles.Custom.Update(ctx, profileID, params) (*
DLPProfileCustomUpdateResponse
, error)
put/accounts/{account_id}/dlp/profiles/custom/{profile_id}
Updates a DLP custom profile.
client.ZeroTrust.DLP.Profiles.Custom.Delete(ctx, profileID, body) (*
DLPProfileCustomDeleteResponse
, error)
delete/accounts/{account_id}/dlp/profiles/custom/{profile_id}
Deletes a DLP custom profile.
Domain types
typeCustomProfilestruct{…}
typePatternstruct{…}
Zero Trust
DLP
Profiles
Predefined
ZeroTrust.DLP.Profiles.Predefined
Methods
client.ZeroTrust.DLP.Profiles.Predefined.Get(ctx, profileID, query) (*
DLPProfilePredefinedGetResponse
, error)
get/accounts/{account_id}/dlp/profiles/predefined/{profile_id}/config
This is similar to get_predefined but only returns entries that are enabled. This is needed for our terraform API Fetches a predefined DLP profile by id.
client.ZeroTrust.DLP.Profiles.Predefined.Update(ctx, profileID, params) (*
DLPProfilePredefinedUpdateResponse
, error)
put/accounts/{account_id}/dlp/profiles/predefined/{profile_id}/config
This is similar to update_predefined but only returns entries that are enabled. This is needed for our terraform API Updates a DLP predefined profile. Only supports enabling/disabling entries.
client.ZeroTrust.DLP.Profiles.Predefined.Delete(ctx, profileID, body) (*
DLPProfilePredefinedDeleteResponse
, error)
delete/accounts/{account_id}/dlp/profiles/predefined/{profile_id}
This is a no-op as predefined profiles can't be deleted but is needed for our generated terraform API.
Domain types
typePredefinedProfilestruct{…}
Zero Trust
Gateway
ZeroTrust.Gateway
Methods
client.ZeroTrust.Gateway.List(ctx, query) (*
GatewayListResponse
, error)
get/accounts/{account_id}/gateway
Retrieve information about the current Zero Trust account.
client.ZeroTrust.Gateway.New(ctx, body) (*
GatewayNewResponse
, error)
post/accounts/{account_id}/gateway
Create a Zero Trust account for an existing Cloudflare account.
Zero Trust
Gateway
App Types
ZeroTrust.Gateway.AppTypes
Methods
client.ZeroTrust.Gateway.AppTypes.List(ctx, query) (*SinglePage[
AppType
], error)
get/accounts/{account_id}/gateway/app_types
List all application and application type mappings.
Domain types
typeAppTypeinterface{…}
Zero Trust
Gateway
Audit SSH Settings
ZeroTrust.Gateway.AuditSSHSettings
Methods
client.ZeroTrust.Gateway.AuditSSHSettings.Get(ctx, query) (*
GatewaySettings
, error)
get/accounts/{account_id}/gateway/audit_ssh_settings
Retrieve all Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.
client.ZeroTrust.Gateway.AuditSSHSettings.Update(ctx, params) (*
GatewaySettings
, error)
put/accounts/{account_id}/gateway/audit_ssh_settings
Update Zero Trust Audit SSH and SSH with Access for Infrastructure settings for an account.
client.ZeroTrust.Gateway.AuditSSHSettings.RotateSeed(ctx, body) (*
GatewaySettings
, error)
post/accounts/{account_id}/gateway/audit_ssh_settings/rotate_seed
Rotate the SSH account seed that generates the host key identity when connecting through the Cloudflare SSH Proxy.
Domain types
typeGatewaySettingsstruct{…}
Zero Trust
Gateway
Categories
ZeroTrust.Gateway.Categories
Methods
client.ZeroTrust.Gateway.Categories.List(ctx, query) (*SinglePage[
Category
], error)
get/accounts/{account_id}/gateway/categories
List all categories.
Domain types
typeCategorystruct{…}
Zero Trust
Gateway
Certificates
ZeroTrust.Gateway.Certificates
Methods
client.ZeroTrust.Gateway.Certificates.List(ctx, query) (*SinglePage[
GatewayCertificateListResponse
], error)
get/accounts/{account_id}/gateway/certificates
List all Zero Trust certificates for an account.
client.ZeroTrust.Gateway.Certificates.Get(ctx, certificateID, query) (*
GatewayCertificateGetResponse
, error)
get/accounts/{account_id}/gateway/certificates/{certificate_id}
Get a single Zero Trust certificate.
client.ZeroTrust.Gateway.Certificates.New(ctx, params) (*
GatewayCertificateNewResponse
, error)
post/accounts/{account_id}/gateway/certificates
Create a new Zero Trust certificate.
client.ZeroTrust.Gateway.Certificates.Delete(ctx, certificateID, body) (*
GatewayCertificateDeleteResponse
, error)
delete/accounts/{account_id}/gateway/certificates/{certificate_id}
Delete a gateway-managed Zero Trust certificate. You must deactivate the certificate from the edge (inactive) before deleting it.
client.ZeroTrust.Gateway.Certificates.Activate(ctx, certificateID, params) (*
GatewayCertificateActivateResponse
, error)
post/accounts/{account_id}/gateway/certificates/{certificate_id}/activate
Bind a single Zero Trust certificate to the edge.
client.ZeroTrust.Gateway.Certificates.Deactivate(ctx, certificateID, params) (*
GatewayCertificateDeactivateResponse
, error)
post/accounts/{account_id}/gateway/certificates/{certificate_id}/deactivate
Unbind a single Zero Trust certificate from the edge.
Zero Trust
Gateway
Configurations
ZeroTrust.Gateway.Configurations
Methods
client.ZeroTrust.Gateway.Configurations.Get(ctx, query) (*
GatewayConfigurationGetResponse
, error)
get/accounts/{account_id}/gateway/configuration
Retrieve the current Zero Trust account configuration.
client.ZeroTrust.Gateway.Configurations.Update(ctx, params) (*
GatewayConfigurationUpdateResponse
, error)
put/accounts/{account_id}/gateway/configuration
Update the current Zero Trust account configuration.
client.ZeroTrust.Gateway.Configurations.Edit(ctx, params) (*
GatewayConfigurationEditResponse
, error)
patch/accounts/{account_id}/gateway/configuration
Update (PATCH) a single subcollection of settings such as antivirus, tls_decrypt, activity_log, block_page, browser_isolation, fips, body_scanning, or certificate without updating the entire configuration object. This endpoint returns an error if any settings collection lacks proper configuration.
Domain types
typeActivityLogSettingsstruct{…}
Specify activity log settings.
typeAntiVirusSettingsstruct{…}
Specify anti-virus settings.
typeBlockPageSettingsstruct{…}
Specify block page layout settings.
typeBodyScanningSettingsstruct{…}
Specify the DLP inspection mode.
typeBrowserIsolationSettingsstruct{…}
Specify Clientless Browser Isolation settings.
typeCustomCertificateSettingsstruct{…}
Specify custom certificate settings for BYO-PKI. This field is deprecated; use certificate instead.
typeExtendedEmailMatchingstruct{…}
Configures user email settings for firewall policies. When you enable this, the system standardizes email addresses in the identity portion of the rule to match extended email variants in firewall policies. When you disable this setting, the system matches email addresses exactly as you provide them. Enable this setting if your email uses . or + modifiers.
typeFipsSettingsstruct{…}
Specify FIPS settings.
typeGatewayConfigurationSettingsstruct{…}
Specify account settings.
typeNotificationSettingsstruct{…}
Configure the message the user's device shows during an antivirus scan.
typeProtocolDetectionstruct{…}
Specify whether to detect protocols from the initial bytes of client traffic.
typeTLSSettingsstruct{…}
Specify whether to inspect encrypted HTTP traffic.
Zero Trust
Gateway
Configurations
Custom Certificate
ZeroTrust.Gateway.Configurations.CustomCertificate
Methods
client.ZeroTrust.Gateway.Configurations.CustomCertificate.Get(ctx, query) (*
CustomCertificateSettings
, error)
Deprecated
get/accounts/{account_id}/gateway/configuration/custom_certificate
Retrieve the current Zero Trust certificate configuration.
Zero Trust
Gateway
Lists
ZeroTrust.Gateway.Lists
Methods
client.ZeroTrust.Gateway.Lists.List(ctx, params) (*SinglePage[
GatewayList
], error)
get/accounts/{account_id}/gateway/lists
Fetch all Zero Trust lists for an account.
client.ZeroTrust.Gateway.Lists.Get(ctx, listID, query) (*
GatewayList
, error)
get/accounts/{account_id}/gateway/lists/{list_id}
Fetch a single Zero Trust list.
client.ZeroTrust.Gateway.Lists.New(ctx, params) (*
GatewayListNewResponse
, error)
post/accounts/{account_id}/gateway/lists
Creates a new Zero Trust list.
client.ZeroTrust.Gateway.Lists.Update(ctx, listID, params) (*
GatewayList
, error)
put/accounts/{account_id}/gateway/lists/{list_id}
Updates a configured Zero Trust list. Skips updating list items if not included in the payload. A non empty list items will overwrite the existing list.
client.ZeroTrust.Gateway.Lists.Edit(ctx, listID, params) (*
GatewayList
, error)
patch/accounts/{account_id}/gateway/lists/{list_id}
Appends or removes an item from a configured Zero Trust list.
client.ZeroTrust.Gateway.Lists.Delete(ctx, listID, body) (*
GatewayListDeleteResponse
, error)
delete/accounts/{account_id}/gateway/lists/{list_id}
Deletes a Zero Trust list.
Domain types
typeGatewayItemstruct{…}
typeGatewayListstruct{…}
Zero Trust
Gateway
Lists
Items
ZeroTrust.Gateway.Lists.Items
Methods
client.ZeroTrust.Gateway.Lists.Items.List(ctx, listID, query) (*SinglePage[[]
GatewayItem
], error)
get/accounts/{account_id}/gateway/lists/{list_id}/items
Fetch all items in a single Zero Trust list.
Zero Trust
Gateway
Locations
ZeroTrust.Gateway.Locations
Methods
client.ZeroTrust.Gateway.Locations.List(ctx, query) (*SinglePage[
Location
], error)
get/accounts/{account_id}/gateway/locations
List Zero Trust Gateway locations for an account.
client.ZeroTrust.Gateway.Locations.Get(ctx, locationID, query) (*
Location
, error)
get/accounts/{account_id}/gateway/locations/{location_id}
Get a single Zero Trust Gateway location.
client.ZeroTrust.Gateway.Locations.New(ctx, params) (*
Location
, error)
post/accounts/{account_id}/gateway/locations
Create a new Zero Trust Gateway location.
client.ZeroTrust.Gateway.Locations.Update(ctx, locationID, params) (*
Location
, error)
put/accounts/{account_id}/gateway/locations/{location_id}
Update a configured Zero Trust Gateway location.
client.ZeroTrust.Gateway.Locations.Delete(ctx, locationID, body) (*
GatewayLocationDeleteResponse
, error)
delete/accounts/{account_id}/gateway/locations/{location_id}
Delete a configured Zero Trust Gateway location.
Domain types
typeDOHEndpointstruct{…}
typeDOTEndpointstruct{…}
typeEndpointstruct{…}
Configure the destination endpoints for this location.
typeIPNetworkstruct{…}
typeIPV4Endpointstruct{…}
typeIPV6Endpointstruct{…}
typeIPV6Networkstruct{…}
typeLocationstruct{…}
Zero Trust
Gateway
Logging
ZeroTrust.Gateway.Logging
Methods
client.ZeroTrust.Gateway.Logging.Get(ctx, query) (*
LoggingSetting
, error)
get/accounts/{account_id}/gateway/logging
Retrieve the current logging settings for the Zero Trust account.
client.ZeroTrust.Gateway.Logging.Update(ctx, params) (*
LoggingSetting
, error)
put/accounts/{account_id}/gateway/logging
Update logging settings for the current Zero Trust account.
Domain types
typeLoggingSettingstruct{…}
Zero Trust
Gateway
Proxy Endpoints
ZeroTrust.Gateway.ProxyEndpoints
Methods
client.ZeroTrust.Gateway.ProxyEndpoints.List(ctx, query) (*SinglePage[
ProxyEndpoint
], error)
get/accounts/{account_id}/gateway/proxy_endpoints
List all Zero Trust Gateway proxy endpoints for an account.
client.ZeroTrust.Gateway.ProxyEndpoints.Get(ctx, proxyEndpointID, query) (*
ProxyEndpoint
, error)
get/accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}
Get a single Zero Trust Gateway proxy endpoint.
client.ZeroTrust.Gateway.ProxyEndpoints.New(ctx, params) (*
ProxyEndpoint
, error)
post/accounts/{account_id}/gateway/proxy_endpoints
Create a new Zero Trust Gateway proxy endpoint.
client.ZeroTrust.Gateway.ProxyEndpoints.Edit(ctx, proxyEndpointID, params) (*
ProxyEndpoint
, error)
patch/accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}
Update a configured Zero Trust Gateway proxy endpoint.
client.ZeroTrust.Gateway.ProxyEndpoints.Delete(ctx, proxyEndpointID, body) (*
GatewayProxyEndpointDeleteResponse
, error)
delete/accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}
Delete a configured Zero Trust Gateway proxy endpoint.
Domain types
typeGatewayIPsstring
Specify an IPv4 or IPv6 CIDR. Limit IPv6 to a maximum of /109 and IPv4 to a maximum of /25.
typeProxyEndpointinterface{…}
Zero Trust
Gateway
Rules
ZeroTrust.Gateway.Rules
Methods
client.ZeroTrust.Gateway.Rules.List(ctx, query) (*SinglePage[
GatewayRule
], error)
get/accounts/{account_id}/gateway/rules
List Zero Trust Gateway rules for an account.
client.ZeroTrust.Gateway.Rules.Get(ctx, ruleID, query) (*
GatewayRule
, error)
get/accounts/{account_id}/gateway/rules/{rule_id}
Get a single Zero Trust Gateway rule.
client.ZeroTrust.Gateway.Rules.New(ctx, params) (*
GatewayRule
, error)
post/accounts/{account_id}/gateway/rules
Create a new Zero Trust Gateway rule.
client.ZeroTrust.Gateway.Rules.Update(ctx, ruleID, params) (*
GatewayRule
, error)
put/accounts/{account_id}/gateway/rules/{rule_id}
Update a configured Zero Trust Gateway rule.
client.ZeroTrust.Gateway.Rules.Delete(ctx, ruleID, body) (*
GatewayRuleDeleteResponse
, error)
delete/accounts/{account_id}/gateway/rules/{rule_id}
Delete a Zero Trust Gateway rule.
client.ZeroTrust.Gateway.Rules.ListTenant(ctx, query) (*SinglePage[
GatewayRule
], error)
get/accounts/{account_id}/gateway/rules/tenant
List Zero Trust Gateway rules for the parent account of an account in the MSP configuration.
client.ZeroTrust.Gateway.Rules.ResetExpiration(ctx, ruleID, body) (*
GatewayRule
, error)
post/accounts/{account_id}/gateway/rules/{rule_id}/reset_expiration
Resets the expiration of a Zero Trust Gateway Rule if its duration elapsed and it has a default duration. The Zero Trust Gateway Rule must have values for both expiration.expires_at and expiration.duration.
Domain types
typeDNSResolverSettingsV4struct{…}
typeDNSResolverSettingsV6struct{…}
typeGatewayFilterstring
Specify the protocol or layer to use.
typeGatewayRulestruct{…}
typeRuleSettingstruct{…}
Defines settings for this rule. Settings apply only to specific rule types and must use compatible selectors. If Terraform detects drift, confirm the setting supports your rule type and check whether the API modifies the value. Use API-returned values in your configuration to prevent drift.
typeSchedulestruct{…}
Defines the schedule for activating DNS policies. Settable only for dns and dns_resolver rules.
Zero Trust
Identity Providers
ZeroTrust.IdentityProviders
Methods
client.ZeroTrust.IdentityProviders.List(ctx, params) (*V4PagePaginationArray[
IdentityProviderListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/access/identity_providers
Lists all configured identity providers.
client.ZeroTrust.IdentityProviders.Get(ctx, identityProviderID, query) (*
IdentityProvider
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/identity_providers/{identity_provider_id}
Fetches a configured identity provider.
client.ZeroTrust.IdentityProviders.New(ctx, params) (*
IdentityProvider
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/identity_providers
Adds a new identity provider to Access.
client.ZeroTrust.IdentityProviders.Update(ctx, identityProviderID, params) (*
IdentityProvider
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/identity_providers/{identity_provider_id}
Updates a configured identity provider.
client.ZeroTrust.IdentityProviders.Delete(ctx, identityProviderID, body) (*
IdentityProviderDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/access/identity_providers/{identity_provider_id}
Deletes an identity provider from Access.
Domain types
typeAzureADstruct{…}
typeGenericOAuthConfigstruct{…}
typeIdentityProviderinterface{…}
typeIdentityProviderSCIMConfigstruct{…}
The configuration settings for enabling a System for Cross-Domain Identity Management (SCIM) with the identity provider.
typeIdentityProviderTypestring
The type of identity provider. To determine the value for a specific provider, refer to our developer documentation.
Zero Trust
Identity Providers
SCIM
ZeroTrust.IdentityProviders.SCIM
Zero Trust
Identity Providers
SCIM
Groups
ZeroTrust.IdentityProviders.SCIM.Groups
Methods
client.ZeroTrust.IdentityProviders.SCIM.Groups.List(ctx, identityProviderID, params) (*V4PagePaginationArray[
ZeroTrustGroup
], error)
get/accounts/{account_id}/access/identity_providers/{identity_provider_id}/scim/groups
Lists SCIM Group resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).
Zero Trust
Identity Providers
SCIM
Users
ZeroTrust.IdentityProviders.SCIM.Users
Methods
client.ZeroTrust.IdentityProviders.SCIM.Users.List(ctx, identityProviderID, params) (*V4PagePaginationArray[
AccessUser
], error)
get/accounts/{account_id}/access/identity_providers/{identity_provider_id}/scim/users
Lists SCIM User resources synced to Cloudflare via the System for Cross-domain Identity Management (SCIM).
Zero Trust
Networks
ZeroTrust.Networks
Zero Trust
Networks
Hostname Routes
ZeroTrust.Networks.HostnameRoutes
Methods
client.ZeroTrust.Networks.HostnameRoutes.List(ctx, params) (*V4PagePaginationArray[
HostnameRoute
], error)
get/accounts/{account_id}/zerotrust/routes/hostname
Lists and filters hostname routes in an account.
client.ZeroTrust.Networks.HostnameRoutes.Get(ctx, hostnameRouteID, query) (*
HostnameRoute
, error)
get/accounts/{account_id}/zerotrust/routes/hostname/{hostname_route_id}
Get a hostname route.
client.ZeroTrust.Networks.HostnameRoutes.New(ctx, params) (*
HostnameRoute
, error)
post/accounts/{account_id}/zerotrust/routes/hostname
Create a hostname route.
client.ZeroTrust.Networks.HostnameRoutes.Edit(ctx, hostnameRouteID, params) (*
HostnameRoute
, error)
patch/accounts/{account_id}/zerotrust/routes/hostname/{hostname_route_id}
Updates a hostname route.
client.ZeroTrust.Networks.HostnameRoutes.Delete(ctx, hostnameRouteID, body) (*
HostnameRoute
, error)
delete/accounts/{account_id}/zerotrust/routes/hostname/{hostname_route_id}
Delete a hostname route.
Domain types
typeHostnameRoutestruct{…}
Zero Trust
Networks
Routes
ZeroTrust.Networks.Routes
Methods
client.ZeroTrust.Networks.Routes.List(ctx, params) (*V4PagePaginationArray[
Teamnet
], error)
get/accounts/{account_id}/teamnet/routes
Lists and filters private network routes in an account.
client.ZeroTrust.Networks.Routes.Get(ctx, routeID, query) (*
Route
, error)
get/accounts/{account_id}/teamnet/routes/{route_id}
Get a private network route in an account.
client.ZeroTrust.Networks.Routes.New(ctx, params) (*
Route
, error)
post/accounts/{account_id}/teamnet/routes
Routes a private network through a Cloudflare Tunnel.
client.ZeroTrust.Networks.Routes.Edit(ctx, routeID, params) (*
Route
, error)
patch/accounts/{account_id}/teamnet/routes/{route_id}
Updates an existing private network route in an account. The fields that are meant to be updated should be provided in the body of the request.
client.ZeroTrust.Networks.Routes.Delete(ctx, routeID, body) (*
Route
, error)
delete/accounts/{account_id}/teamnet/routes/{route_id}
Deletes a private network route from an account.
Domain types
typeNetworkRoutestruct{…}
typeRoutestruct{…}
typeTeamnetstruct{…}
Zero Trust
Networks
Routes
IPs
ZeroTrust.Networks.Routes.IPs
Methods
client.ZeroTrust.Networks.Routes.IPs.Get(ctx, ip, params) (*
Teamnet
, error)
get/accounts/{account_id}/teamnet/routes/ip/{ip}
Fetches routes that contain the given IP address.
Zero Trust
Networks
Routes
Networks
ZeroTrust.Networks.Routes.Networks
Methods
client.ZeroTrust.Networks.Routes.Networks.New(ctx, ipNetworkEncoded, params) (*
Route
, error)
Deprecated
post/accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}
Deprecated
This endpoint and its related APIs are deprecated in favor of the equivalent Tunnel Route (without CIDR) APIs.
Routes a private network through a Cloudflare Tunnel. The CIDR in ip_network_encoded must be written in URL-encoded format.
client.ZeroTrust.Networks.Routes.Networks.Edit(ctx, ipNetworkEncoded, body) (*
Route
, error)
Deprecated
patch/accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}
Deprecated
This endpoint and its related APIs are deprecated in favor of the equivalent Tunnel Route (without CIDR) APIs.
Updates an existing private network route in an account. The CIDR in ip_network_encoded must be written in URL-encoded format.
client.ZeroTrust.Networks.Routes.Networks.Delete(ctx, ipNetworkEncoded, params) (*
Route
, error)
Deprecated
delete/accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}
Deprecated
This endpoint and its related APIs are deprecated in favor of the equivalent Tunnel Route (without CIDR) APIs.
Deletes a private network route from an account. The CIDR in ip_network_encoded must be written in URL-encoded format. If no virtual_network_id is provided it will delete the route from the default vnet. If no tun_type is provided it will fetch the type from the tunnel_id or if that is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided it will delete the route from that tunnel, otherwise it will delete the route based on the vnet and tun_type.
Zero Trust
Networks
Subnets
ZeroTrust.Networks.Subnets
Methods
client.ZeroTrust.Networks.Subnets.List(ctx, params) (*V4PagePaginationArray[
NetworkSubnetListResponse
], error)
get/accounts/{account_id}/zerotrust/subnets
Lists and filters subnets in an account.
Zero Trust
Networks
Subnets
Cloudflare Source
ZeroTrust.Networks.Subnets.CloudflareSource
Methods
client.ZeroTrust.Networks.Subnets.CloudflareSource.Update(ctx, addressFamily, params) (*
NetworkSubnetCloudflareSourceUpdateResponse
, error)
patch/accounts/{account_id}/zerotrust/subnets/cloudflare_source/{address_family}
Updates the Cloudflare Source subnet of the given address family
Zero Trust
Networks
Virtual Networks
ZeroTrust.Networks.VirtualNetworks
Methods
client.ZeroTrust.Networks.VirtualNetworks.List(ctx, params) (*SinglePage[
VirtualNetwork
], error)
get/accounts/{account_id}/teamnet/virtual_networks
Lists and filters virtual networks in an account.
client.ZeroTrust.Networks.VirtualNetworks.Get(ctx, virtualNetworkID, query) (*
VirtualNetwork
, error)
get/accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}
Get a virtual network.
client.ZeroTrust.Networks.VirtualNetworks.New(ctx, params) (*
VirtualNetwork
, error)
post/accounts/{account_id}/teamnet/virtual_networks
Adds a new virtual network to an account.
client.ZeroTrust.Networks.VirtualNetworks.Edit(ctx, virtualNetworkID, params) (*
VirtualNetwork
, error)
patch/accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}
Updates an existing virtual network.
client.ZeroTrust.Networks.VirtualNetworks.Delete(ctx, virtualNetworkID, body) (*
VirtualNetwork
, error)
delete/accounts/{account_id}/teamnet/virtual_networks/{virtual_network_id}
Deletes an existing virtual network.
Domain types
typeVirtualNetworkstruct{…}
Zero Trust
Organizations
ZeroTrust.Organizations
Methods
client.ZeroTrust.Organizations.List(ctx, query) (*
Organization
, error)
get/{accounts_or_zones}/{account_or_zone_id}/access/organizations
Returns the configuration for your Zero Trust organization.
client.ZeroTrust.Organizations.New(ctx, params) (*
Organization
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/organizations
Sets up a Zero Trust organization for your account or zone.
client.ZeroTrust.Organizations.Update(ctx, params) (*
Organization
, error)
put/{accounts_or_zones}/{account_or_zone_id}/access/organizations
Updates the configuration for your Zero Trust organization.
client.ZeroTrust.Organizations.RevokeUsers(ctx, params) (*
OrganizationRevokeUsersResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/access/organizations/revoke_user
Revokes a user's access across all applications.
Domain types
typeLoginDesignstruct{…}
typeOrganizationstruct{…}
Zero Trust
Organizations
DOH
ZeroTrust.Organizations.DOH
Methods
client.ZeroTrust.Organizations.DOH.Get(ctx, query) (*
OrganizationDOHGetResponse
, error)
get/accounts/{account_id}/access/organizations/doh
Returns the DoH settings for your Zero Trust organization.
client.ZeroTrust.Organizations.DOH.Update(ctx, params) (*
OrganizationDOHUpdateResponse
, error)
put/accounts/{account_id}/access/organizations/doh
Updates the DoH settings for your Zero Trust organization.
Zero Trust
Risk Scoring
ZeroTrust.RiskScoring
Methods
client.ZeroTrust.RiskScoring.Get(ctx, userID, query) (*
RiskScoringGetResponse
, error)
get/accounts/{account_id}/zt_risk_scoring/{user_id}
Get risk event/score information for a specific user
client.ZeroTrust.RiskScoring.Reset(ctx, userID, body) (*
RiskScoringResetResponse
, error)
post/accounts/{account_id}/zt_risk_scoring/{user_id}/reset
Clear the risk score for a particular user
Zero Trust
Risk Scoring
Behaviours
ZeroTrust.RiskScoring.Behaviours
Methods
client.ZeroTrust.RiskScoring.Behaviours.Get(ctx, query) (*
RiskScoringBehaviourGetResponse
, error)
get/accounts/{account_id}/zt_risk_scoring/behaviors
Get all behaviors and associated configuration
client.ZeroTrust.RiskScoring.Behaviours.Update(ctx, params) (*
RiskScoringBehaviourUpdateResponse
, error)
put/accounts/{account_id}/zt_risk_scoring/behaviors
Update configuration for risk behaviors
Zero Trust
Risk Scoring
Integrations
ZeroTrust.RiskScoring.Integrations
Methods
client.ZeroTrust.RiskScoring.Integrations.List(ctx, query) (*SinglePage[
RiskScoringIntegrationListResponse
], error)
get/accounts/{account_id}/zt_risk_scoring/integrations
List all risk score integrations for the account.
client.ZeroTrust.RiskScoring.Integrations.Get(ctx, integrationID, query) (*
RiskScoringIntegrationGetResponse
, error)
get/accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}
Get risk score integration by id.
client.ZeroTrust.RiskScoring.Integrations.New(ctx, params) (*
RiskScoringIntegrationNewResponse
, error)
post/accounts/{account_id}/zt_risk_scoring/integrations
Create new risk score integration.
client.ZeroTrust.RiskScoring.Integrations.Update(ctx, integrationID, params) (*
RiskScoringIntegrationUpdateResponse
, error)
put/accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}
Overwrite the reference_id, tenant_url, and active values with the ones provided.
client.ZeroTrust.RiskScoring.Integrations.Delete(ctx, integrationID, body) (*
RiskScoringIntegrationDeleteResponse
, error)
delete/accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}
Delete a risk score integration.
Zero Trust
Risk Scoring
Integrations
References
ZeroTrust.RiskScoring.Integrations.References
Methods
client.ZeroTrust.RiskScoring.Integrations.References.Get(ctx, referenceID, query) (*
RiskScoringIntegrationReferenceGetResponse
, error)
get/accounts/{account_id}/zt_risk_scoring/integrations/reference_id/{reference_id}
Get risk score integration by reference id.
Zero Trust
Risk Scoring
Summary
ZeroTrust.RiskScoring.Summary
Methods
client.ZeroTrust.RiskScoring.Summary.Get(ctx, query) (*
RiskScoringSummaryGetResponse
, error)
get/accounts/{account_id}/zt_risk_scoring/summary
Get risk score info for all users in the account
Zero Trust
Seats
ZeroTrust.Seats
Methods
client.ZeroTrust.Seats.Edit(ctx, params) (*SinglePage[
Seat
], error)
patch/accounts/{account_id}/access/seats
Removes a user from a Zero Trust seat when both access_seat and gateway_seat are set to false.
Domain types
typeSeatstruct{…}
Zero Trust
Tunnels
ZeroTrust.Tunnels
Methods
client.ZeroTrust.Tunnels.List(ctx, params) (*V4PagePaginationArray[
TunnelListResponse
], error)
get/accounts/{account_id}/tunnels
Lists and filters all types of Tunnels in an account.
Domain types
typeConnectionstruct{…}
Zero Trust
Tunnels
Cloudflared
ZeroTrust.Tunnels.Cloudflared
Methods
client.ZeroTrust.Tunnels.Cloudflared.List(ctx, params) (*V4PagePaginationArray[
CloudflareTunnel
], error)
get/accounts/{account_id}/cfd_tunnel
Lists and filters Cloudflare Tunnels in an account.
client.ZeroTrust.Tunnels.Cloudflared.Get(ctx, tunnelID, query) (*
CloudflareTunnel
, error)
get/accounts/{account_id}/cfd_tunnel/{tunnel_id}
Fetches a single Cloudflare Tunnel.
client.ZeroTrust.Tunnels.Cloudflared.New(ctx, params) (*
CloudflareTunnel
, error)
post/accounts/{account_id}/cfd_tunnel
Creates a new Cloudflare Tunnel in an account.
client.ZeroTrust.Tunnels.Cloudflared.Edit(ctx, tunnelID, params) (*
CloudflareTunnel
, error)
patch/accounts/{account_id}/cfd_tunnel/{tunnel_id}
Updates an existing Cloudflare Tunnel.
client.ZeroTrust.Tunnels.Cloudflared.Delete(ctx, tunnelID, body) (*
CloudflareTunnel
, error)
delete/accounts/{account_id}/cfd_tunnel/{tunnel_id}
Deletes a Cloudflare Tunnel from an account.
Zero Trust
Tunnels
Cloudflared
Configurations
ZeroTrust.Tunnels.Cloudflared.Configurations
Methods
client.ZeroTrust.Tunnels.Cloudflared.Configurations.Get(ctx, tunnelID, query) (*
TunnelCloudflaredConfigurationGetResponse
, error)
get/accounts/{account_id}/cfd_tunnel/{tunnel_id}/configurations
Gets the configuration for a remotely-managed tunnel
client.ZeroTrust.Tunnels.Cloudflared.Configurations.Update(ctx, tunnelID, params) (*
TunnelCloudflaredConfigurationUpdateResponse
, error)
put/accounts/{account_id}/cfd_tunnel/{tunnel_id}/configurations
Adds or updates the configuration for a remotely-managed tunnel.
Zero Trust
Tunnels
Cloudflared
Connections
ZeroTrust.Tunnels.Cloudflared.Connections
Methods
client.ZeroTrust.Tunnels.Cloudflared.Connections.Get(ctx, tunnelID, query) (*SinglePage[
Client
], error)
get/accounts/{account_id}/cfd_tunnel/{tunnel_id}/connections
Fetches connection details for a Cloudflare Tunnel.
client.ZeroTrust.Tunnels.Cloudflared.Connections.Delete(ctx, tunnelID, params) (*
TunnelCloudflaredConnectionDeleteResponse
, error)
delete/accounts/{account_id}/cfd_tunnel/{tunnel_id}/connections
Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel independently of its current state. If no connector id (client_id) is provided all connectors will be removed. We recommend running this command after rotating tokens.
Domain types
typeClientstruct{…}
A client (typically cloudflared) that maintains connections to a Cloudflare data center.
Zero Trust
Tunnels
Cloudflared
Connectors
ZeroTrust.Tunnels.Cloudflared.Connectors
Methods
client.ZeroTrust.Tunnels.Cloudflared.Connectors.Get(ctx, tunnelID, connectorID, query) (*
Client
, error)
get/accounts/{account_id}/cfd_tunnel/{tunnel_id}/connectors/{connector_id}
Fetches connector and connection details for a Cloudflare Tunnel.
Zero Trust
Tunnels
Cloudflared
Management
ZeroTrust.Tunnels.Cloudflared.Management
Methods
client.ZeroTrust.Tunnels.Cloudflared.Management.New(ctx, tunnelID, params) (*string, error)
post/accounts/{account_id}/cfd_tunnel/{tunnel_id}/management
Gets a management token used to access the management resources (i.e. Streaming Logs) of a tunnel.
Zero Trust
Tunnels
Cloudflared
Token
ZeroTrust.Tunnels.Cloudflared.Token
Methods
client.ZeroTrust.Tunnels.Cloudflared.Token.Get(ctx, tunnelID, query) (*string, error)
get/accounts/{account_id}/cfd_tunnel/{tunnel_id}/token
Gets the token used to associate cloudflared with a specific tunnel.
Zero Trust
Tunnels
WARP Connector
ZeroTrust.Tunnels.WARPConnector
Methods
client.ZeroTrust.Tunnels.WARPConnector.List(ctx, params) (*V4PagePaginationArray[
TunnelWARPConnectorListResponse
], error)
get/accounts/{account_id}/warp_connector
Lists and filters Warp Connector Tunnels in an account.
client.ZeroTrust.Tunnels.WARPConnector.Get(ctx, tunnelID, query) (*
TunnelWARPConnectorGetResponse
, error)
get/accounts/{account_id}/warp_connector/{tunnel_id}
Fetches a single Warp Connector Tunnel.
client.ZeroTrust.Tunnels.WARPConnector.New(ctx, params) (*
TunnelWARPConnectorNewResponse
, error)
post/accounts/{account_id}/warp_connector
Creates a new Warp Connector Tunnel in an account.
client.ZeroTrust.Tunnels.WARPConnector.Edit(ctx, tunnelID, params) (*
TunnelWARPConnectorEditResponse
, error)
patch/accounts/{account_id}/warp_connector/{tunnel_id}
Updates an existing Warp Connector Tunnel.
client.ZeroTrust.Tunnels.WARPConnector.Delete(ctx, tunnelID, body) (*
TunnelWARPConnectorDeleteResponse
, error)
delete/accounts/{account_id}/warp_connector/{tunnel_id}
Deletes a Warp Connector Tunnel from an account.
Zero Trust
Tunnels
WARP Connector
Token
ZeroTrust.Tunnels.WARPConnector.Token
Methods
client.ZeroTrust.Tunnels.WARPConnector.Token.Get(ctx, tunnelID, query) (*string, error)
get/accounts/{account_id}/warp_connector/{tunnel_id}/token
Gets the token used to associate warp device with a specific Warp Connector tunnel.