Resource Sharing
ResourceSharing
Methods
client.ResourceSharing.List(ctx, params) (*V4PagePaginationArray[
ResourceSharingListResponse
], error)
get/accounts/{account_id}/shares
Lists all account shares.
client.ResourceSharing.Get(ctx, shareID, params) (*
ResourceSharingGetResponse
, error)
get/accounts/{account_id}/shares/{share_id}
Fetches share by ID.
client.ResourceSharing.New(ctx, params) (*
ResourceSharingNewResponse
, error)
post/accounts/{account_id}/shares
Create a new share
client.ResourceSharing.Update(ctx, shareID, params) (*
ResourceSharingUpdateResponse
, error)
put/accounts/{account_id}/shares/{share_id}
Updating is not immediate, an updated share object with a new status will be returned.
client.ResourceSharing.Delete(ctx, shareID, body) (*
ResourceSharingDeleteResponse
, error)
delete/accounts/{account_id}/shares/{share_id}
Deletion is not immediate, an updated share object with a new status will be returned.
Resource Sharing
Recipients
ResourceSharing.Recipients
Methods
client.ResourceSharing.Recipients.List(ctx, shareID, params) (*V4PagePaginationArray[
RecipientListResponse
], error)
get/accounts/{account_id}/shares/{share_id}/recipients
List share recipients by share ID.
client.ResourceSharing.Recipients.Get(ctx, shareID, recipientID, params) (*
RecipientGetResponse
, error)
get/accounts/{account_id}/shares/{share_id}/recipients/{recipient_id}
Get share recipient by ID.
client.ResourceSharing.Recipients.New(ctx, shareID, params) (*
RecipientNewResponse
, error)
post/accounts/{account_id}/shares/{share_id}/recipients
Create a new share recipient
client.ResourceSharing.Recipients.Delete(ctx, shareID, recipientID, body) (*
RecipientDeleteResponse
, error)
delete/accounts/{account_id}/shares/{share_id}/recipients/{recipient_id}
Deletion is not immediate, an updated share recipient object with a new status will be returned.
Resource Sharing
Resources
ResourceSharing.Resources
Methods
client.ResourceSharing.Resources.List(ctx, shareID, params) (*V4PagePaginationArray[
ResourceListResponse
], error)
get/accounts/{account_id}/shares/{share_id}/resources
List share resources by share ID.
client.ResourceSharing.Resources.Get(ctx, shareID, resourceID, query) (*
ResourceGetResponse
, error)
get/accounts/{account_id}/shares/{share_id}/resources/{resource_id}
Get share resource by ID.
client.ResourceSharing.Resources.New(ctx, shareID, params) (*
ResourceNewResponse
, error)
post/accounts/{account_id}/shares/{share_id}/resources
Create a new share resource
client.ResourceSharing.Resources.Update(ctx, shareID, resourceID, params) (*
ResourceUpdateResponse
, error)
put/accounts/{account_id}/shares/{share_id}/resources/{resource_id}
Update is not immediate, an updated share resource object with a new status will be returned.
client.ResourceSharing.Resources.Delete(ctx, shareID, resourceID, body) (*
ResourceDeleteResponse
, error)
delete/accounts/{account_id}/shares/{share_id}/resources/{resource_id}
Deletion is not immediate, an updated share resource object with a new status will be returned.