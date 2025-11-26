IAM
IAM
IAM
Permission Groups
IAM.PermissionGroups
Methods
client.IAM.PermissionGroups.List(ctx, params) (*V4PagePaginationArray[
PermissionGroupListResponse
], error)
get/accounts/{account_id}/iam/permission_groups
List all the permissions groups for an account.
client.IAM.PermissionGroups.Get(ctx, permissionGroupID, query) (*
PermissionGroupGetResponse
, error)
get/accounts/{account_id}/iam/permission_groups/{permission_group_id}
Get information about a specific permission group in an account.
IAM
Resource Groups
IAM.ResourceGroups
Methods
client.IAM.ResourceGroups.List(ctx, params) (*SinglePage[
ResourceGroupListResponse
], error)
get/accounts/{account_id}/iam/resource_groups
List all the resource groups for an account.
client.IAM.ResourceGroups.Get(ctx, resourceGroupID, query) (*
ResourceGroupGetResponse
, error)
get/accounts/{account_id}/iam/resource_groups/{resource_group_id}
Get information about a specific resource group in an account.
client.IAM.ResourceGroups.New(ctx, params) (*
ResourceGroupNewResponse
, error)
post/accounts/{account_id}/iam/resource_groups
Create a new Resource Group under the specified account.
client.IAM.ResourceGroups.Update(ctx, resourceGroupID, params) (*
ResourceGroupUpdateResponse
, error)
put/accounts/{account_id}/iam/resource_groups/{resource_group_id}
Modify an existing resource group.
client.IAM.ResourceGroups.Delete(ctx, resourceGroupID, body) (*
ResourceGroupDeleteResponse
, error)
delete/accounts/{account_id}/iam/resource_groups/{resource_group_id}
Remove a resource group from an account.
IAM
SSO
IAM.SSO
Methods
client.IAM.SSO.List(ctx, query) (*SinglePage[
SSOListResponse
], error)
get/accounts/{account_id}/sso_connectors
Get all SSO connectors
client.IAM.SSO.Get(ctx, ssoConnectorID, query) (*
SSOGetResponse
, error)
get/accounts/{account_id}/sso_connectors/{sso_connector_id}
Get single SSO connector
client.IAM.SSO.New(ctx, params) (*
SSONewResponse
, error)
post/accounts/{account_id}/sso_connectors
Initialize new SSO connector
client.IAM.SSO.Update(ctx, ssoConnectorID, params) (*
SSOUpdateResponse
, error)
patch/accounts/{account_id}/sso_connectors/{sso_connector_id}
Update SSO connector state
client.IAM.SSO.Delete(ctx, ssoConnectorID, body) (*
SSODeleteResponse
, error)
delete/accounts/{account_id}/sso_connectors/{sso_connector_id}
Delete SSO connector
client.IAM.SSO.BeginVerification(ctx, ssoConnectorID, body) (*
SSOBeginVerificationResponse
, error)
post/accounts/{account_id}/sso_connectors/{sso_connector_id}/begin_verification
Begin SSO connector verification
IAM
User Groups
IAM.UserGroups
Methods
client.IAM.UserGroups.List(ctx, params) (*V4PagePaginationArray[
UserGroupListResponse
], error)
get/accounts/{account_id}/iam/user_groups
List all the user groups for an account.
client.IAM.UserGroups.Get(ctx, userGroupID, query) (*
UserGroupGetResponse
, error)
get/accounts/{account_id}/iam/user_groups/{user_group_id}
Get information about a specific user group in an account.
client.IAM.UserGroups.New(ctx, params) (*
UserGroupNewResponse
, error)
post/accounts/{account_id}/iam/user_groups
Create a new user group under the specified account.
client.IAM.UserGroups.Update(ctx, userGroupID, params) (*
UserGroupUpdateResponse
, error)
put/accounts/{account_id}/iam/user_groups/{user_group_id}
Modify an existing user group.
client.IAM.UserGroups.Delete(ctx, userGroupID, body) (*
UserGroupDeleteResponse
, error)
delete/accounts/{account_id}/iam/user_groups/{user_group_id}
Remove a user group from an account.
IAM
User Groups
Members
IAM.UserGroups.Members
Methods
client.IAM.UserGroups.Members.List(ctx, userGroupID, params) (*V4PagePaginationArray[
UserGroupMemberListResponse
], error)
get/accounts/{account_id}/iam/user_groups/{user_group_id}/members
List all the members attached to a user group.
client.IAM.UserGroups.Members.New(ctx, userGroupID, params) (*
UserGroupMemberNewResponse
, error)
post/accounts/{account_id}/iam/user_groups/{user_group_id}/members
Add members to a User Group.
client.IAM.UserGroups.Members.Update(ctx, userGroupID, params) (*SinglePage[
UserGroupMemberUpdateResponse
], error)
put/accounts/{account_id}/iam/user_groups/{user_group_id}/members
Replace the set of members attached to a User Group.
client.IAM.UserGroups.Members.Delete(ctx, userGroupID, memberID, body) (*
UserGroupMemberDeleteResponse
, error)
delete/accounts/{account_id}/iam/user_groups/{user_group_id}/members/{member_id}
Remove a member from User Group