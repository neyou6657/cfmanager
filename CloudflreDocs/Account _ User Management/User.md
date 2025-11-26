User
User
Methods
client.User.Get(ctx) (*
UserGetResponse
, error)
get/user
User Details
client.User.Edit(ctx, body) (*
UserEditResponse
, error)
patch/user
Edit part of your user details.
User
Audit Logs
User.AuditLogs
Methods
client.User.AuditLogs.List(ctx, query) (*V4PagePaginationArray[
AuditLog
], error)
get/user/audit_logs
Gets a list of audit logs for a user account. Can be filtered by who made the change, on which zone, and the timeframe of the change.
User
Billing
User.Billing
User
Billing
History
User.Billing.History
Methods
client.User.Billing.History.List(ctx, query) (*V4PagePaginationArray[
BillingHistory
], error)
Deprecated
get/user/billing/history
Accesses your billing history object.
Domain types
typeBillingHistorystruct{…}
User
Billing
Profile
User.Billing.Profile
Methods
client.User.Billing.Profile.Get(ctx) (*
BillingProfileGetResponse
, error)
Deprecated
get/user/billing/profile
Accesses your billing profile object.
User
Invites
User.Invites
Methods
client.User.Invites.List(ctx) (*SinglePage[
Invite
], error)
get/user/invites
Lists all invitations associated with my user.
client.User.Invites.Get(ctx, inviteID) (*
Invite
, error)
get/user/invites/{invite_id}
Gets the details of an invitation.
client.User.Invites.Edit(ctx, inviteID, body) (*
Invite
, error)
patch/user/invites/{invite_id}
Responds to an invitation.
Domain types
typeInvitestruct{…}
User
Organizations
User.Organizations
Methods
client.User.Organizations.List(ctx, query) (*V4PagePaginationArray[
Organization
], error)
Deprecated
get/user/organizations
Lists organizations the user is associated with.
client.User.Organizations.Get(ctx, organizationID) (*
OrganizationGetResponse
, error)
Deprecated
get/user/organizations/{organization_id}
Gets a specific organization the user is associated with.
client.User.Organizations.Delete(ctx, organizationID) (*
OrganizationDeleteResponse
, error)
Deprecated
delete/user/organizations/{organization_id}
Removes association to an organization.
Domain types
typeOrganizationstruct{…}
User
Subscriptions
User.Subscriptions
Methods
client.User.Subscriptions.Get(ctx) (*SinglePage[
Subscription
], error)
get/user/subscriptions
Lists all of a user's subscriptions.
client.User.Subscriptions.Update(ctx, identifier, body) (*unknown, error)
put/user/subscriptions/{identifier}
Updates a user's subscriptions.
client.User.Subscriptions.Delete(ctx, identifier) (*
SubscriptionDeleteResponse
, error)
delete/user/subscriptions/{identifier}
Deletes a user's subscription.
User
Tokens
User.Tokens
Methods
client.User.Tokens.List(ctx, query) (*V4PagePaginationArray[
Token
], error)
get/user/tokens
List all access tokens you created.
client.User.Tokens.Get(ctx, tokenID) (*
Token
, error)
get/user/tokens/{token_id}
Get information about a specific token.
client.User.Tokens.New(ctx, body) (*
TokenNewResponse
, error)
post/user/tokens
Create a new access token.
client.User.Tokens.Update(ctx, tokenID, body) (*
Token
, error)
put/user/tokens/{token_id}
Update an existing token.
client.User.Tokens.Delete(ctx, tokenID) (*
TokenDeleteResponse
, error)
delete/user/tokens/{token_id}
Destroy a token.
client.User.Tokens.Verify(ctx) (*
TokenVerifyResponse
, error)
get/user/tokens/verify
Test whether a token works.
User
Tokens
Permission Groups
User.Tokens.PermissionGroups
Methods
client.User.Tokens.PermissionGroups.List(ctx, query) (*SinglePage[
TokenPermissionGroupListResponse
], error)
get/user/tokens/permission_groups
Find all available permission groups for API Tokens
User
Tokens
Value
User.Tokens.Value
Methods
client.User.Tokens.Value.Update(ctx, tokenID, body) (*
TokenValue
, error)
put/user/tokens/{token_id}/value
Roll the token secret.