Accounts
Accounts
Methods
client.Accounts.List(ctx, query) (*V4PagePaginationArray[
Account
], error)
get/accounts
List all accounts you have ownership or verified access to.
client.Accounts.Get(ctx, query) (*
Account
, error)
get/accounts/{account_id}
Get information about a specific account that you are a member of.
client.Accounts.New(ctx, body) (*
Account
, error)
post/accounts
Create an account (only available for tenant admins at this time)
client.Accounts.Update(ctx, params) (*
Account
, error)
put/accounts/{account_id}
Update an existing account.
client.Accounts.Delete(ctx, body) (*
AccountDeleteResponse
, error)
delete/accounts/{account_id}
Delete a specific account (only available for tenant admins at this time). This is a permanent operation that will delete any zones or other resources under the account
Domain types
typeAccountstruct{…}
Accounts
Account Organizations
Accounts.AccountOrganizations
Methods
client.Accounts.AccountOrganizations.New(ctx, params) (*
AccountOrganizationNewResponse
, error)
post/accounts/{account_id}/move
Move an account within an organization hierarchy or an account outside an organization. (Currently in Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
Accounts
Account Profile
Accounts.AccountProfile
Methods
client.Accounts.AccountProfile.Get(ctx, query) (*
AccountProfile
, error)
get/accounts/{account_id}/profile
Get account profile
client.Accounts.AccountProfile.Update(ctx, params) error
put/accounts/{account_id}/profile
Modify account profile
Domain types
typeAccountProfilestruct{…}
Accounts
Logs
Accounts.Logs
Accounts
Logs
Audit
Accounts.Logs.Audit
Methods
client.Accounts.Logs.Audit.List(ctx, params) (*CursorPaginationAfter[
LogAuditListResponse
], error)
get/accounts/{account_id}/logs/audit
Gets a list of audit logs for an account. 

 This is the beta release of Audit Logs Version 2. Since this is a beta version, there may be gaps or missing entries in the available audit logs. Be aware of the following limitations. 

Audit logs are available only for the past 30 days. 

Error handling is not yet implemented. 

Accounts
Members
Accounts.Members
Methods
client.Accounts.Members.List(ctx, params) (*V4PagePaginationArray[
Member
], error)
get/accounts/{account_id}/members
List all members of an account.
client.Accounts.Members.Get(ctx, memberID, query) (*
Member
, error)
get/accounts/{account_id}/members/{member_id}
Get information about a specific member of an account.
client.Accounts.Members.New(ctx, params) (*
Member
, error)
post/accounts/{account_id}/members
Add a user to the list of members for this account.
client.Accounts.Members.Update(ctx, memberID, params) (*
Member
, error)
put/accounts/{account_id}/members/{member_id}
Modify an account member.
client.Accounts.Members.Delete(ctx, memberID, body) (*
MemberDeleteResponse
, error)
delete/accounts/{account_id}/members/{member_id}
Remove a member from an account.
Domain types
typeStatusstring
Whether the user is a member of the organization or has an invitation pending.
Accounts
Roles
Accounts.Roles
Methods
client.Accounts.Roles.List(ctx, params) (*V4PagePaginationArray[
Role
], error)
get/accounts/{account_id}/roles
Get all available roles for an account.
client.Accounts.Roles.Get(ctx, roleID, query) (*
Role
, error)
get/accounts/{account_id}/roles/{role_id}
Get information about a specific role for an account.
Accounts
Subscriptions
Accounts.Subscriptions
Methods
client.Accounts.Subscriptions.Get(ctx, query) (*SinglePage[
Subscription
], error)
get/accounts/{account_id}/subscriptions
Lists all of an account's subscriptions.
client.Accounts.Subscriptions.New(ctx, params) (*
Subscription
, error)
post/accounts/{account_id}/subscriptions
Creates an account subscription.
client.Accounts.Subscriptions.Update(ctx, subscriptionIdentifier, params) (*
Subscription
, error)
put/accounts/{account_id}/subscriptions/{subscription_identifier}
Updates an account subscription.
client.Accounts.Subscriptions.Delete(ctx, subscriptionIdentifier, body) (*
SubscriptionDeleteResponse
, error)
delete/accounts/{account_id}/subscriptions/{subscription_identifier}
Deletes an account's subscription.
Accounts
Tokens
Accounts.Tokens
Methods
client.Accounts.Tokens.List(ctx, params) (*V4PagePaginationArray[
Token
], error)
get/accounts/{account_id}/tokens
List all Account Owned API tokens created for this account.
client.Accounts.Tokens.Get(ctx, tokenID, query) (*
Token
, error)
get/accounts/{account_id}/tokens/{token_id}
Get information about a specific Account Owned API token.
client.Accounts.Tokens.New(ctx, params) (*
TokenNewResponse
, error)
post/accounts/{account_id}/tokens
Create a new Account Owned API token.
client.Accounts.Tokens.Update(ctx, tokenID, params) (*
Token
, error)
put/accounts/{account_id}/tokens/{token_id}
Update an existing token.
client.Accounts.Tokens.Delete(ctx, tokenID, body) (*
TokenDeleteResponse
, error)
delete/accounts/{account_id}/tokens/{token_id}
Destroy an Account Owned API token.
client.Accounts.Tokens.Verify(ctx, query) (*
TokenVerifyResponse
, error)
get/accounts/{account_id}/tokens/verify
Test whether a token works.
Accounts
Tokens
Permission Groups
Accounts.Tokens.PermissionGroups
Methods
client.Accounts.Tokens.PermissionGroups.List(ctx, params) (*SinglePage[
TokenPermissionGroupListResponse
], error)
get/accounts/{account_id}/tokens/permission_groups
Find all available permission groups for Account Owned API Tokens
client.Accounts.Tokens.PermissionGroups.Get(ctx, params) (*[]
TokenPermissionGroupGetResponse
, error)
get/accounts/{account_id}/tokens/permission_groups
Find all available permission groups for Account Owned API Tokens
Accounts
Tokens
Value
Accounts.Tokens.Value
Methods
client.Accounts.Tokens.Value.Update(ctx, tokenID, params) (*
TokenValue
, error)
put/accounts/{account_id}/tokens/{token_id}/value
Roll the Account Owned API token secret.