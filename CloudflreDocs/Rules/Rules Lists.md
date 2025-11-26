Rules Lists
Rules
Rules Lists
Lists
Rules.Lists
Methods
client.Rules.Lists.List(ctx, query) (*SinglePage[
ListsList
], error)
get/accounts/{account_id}/rules/lists
Fetches all lists in the account.
client.Rules.Lists.Get(ctx, listID, query) (*
ListGetResponse
, error)
get/accounts/{account_id}/rules/lists/{list_id}
Fetches the details of a list.
client.Rules.Lists.New(ctx, params) (*
ListNewResponse
, error)
post/accounts/{account_id}/rules/lists
Creates a new list of the specified kind.
client.Rules.Lists.Update(ctx, listID, params) (*
ListUpdateResponse
, error)
put/accounts/{account_id}/rules/lists/{list_id}
Updates the description of a list.
client.Rules.Lists.Delete(ctx, listID, body) (*
ListDeleteResponse
, error)
delete/accounts/{account_id}/rules/lists/{list_id}
Deletes a specific list and all its items.
Domain types
typeHostnamestruct{…}
Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).
typeListsListstruct{…}
typeRedirectstruct{…}
The definition of the redirect.
Rules Lists
Lists
Bulk Operations
Rules.Lists.BulkOperations
Methods
client.Rules.Lists.BulkOperations.Get(ctx, operationID, query) (*
ListBulkOperationGetResponse
, error)
get/accounts/{account_id}/rules/lists/bulk_operations/{operation_id}
Gets the current status of an asynchronous operation on a list.
The status property can have one of the following values: pending, running, completed, or failed. If the status is failed, the error property will contain a message describing the error.
Rules Lists
Lists
Items
Rules.Lists.Items
Methods
client.Rules.Lists.Items.List(ctx, listID, params) (*CursorPaginationAfter[
ListItemListResponse
], error)
get/accounts/{account_id}/rules/lists/{list_id}/items
Fetches all the items in the list.
client.Rules.Lists.Items.Get(ctx, listID, itemID, query) (*
ListItemGetResponse
, error)
get/accounts/{account_id}/rules/lists/{list_id}/items/{item_id}
Fetches a list item in the list.
client.Rules.Lists.Items.New(ctx, listID, params) (*
ListItemNewResponse
, error)
post/accounts/{account_id}/rules/lists/{list_id}/items
Appends new items to the list.
This operation is asynchronous. To get current the operation status, invoke the Get bulk operation status endpoint with the returned operation_id.
There is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.
client.Rules.Lists.Items.Update(ctx, listID, params) (*
ListItemUpdateResponse
, error)
put/accounts/{account_id}/rules/lists/{list_id}/items
Removes all existing items from the list and adds the provided items to the list.
This operation is asynchronous. To get current the operation status, invoke the Get bulk operation status endpoint with the returned operation_id.
There is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.
client.Rules.Lists.Items.Delete(ctx, listID, params) (*
ListItemDeleteResponse
, error)
delete/accounts/{account_id}/rules/lists/{list_id}/items
Removes one or more items from a list.
This operation is asynchronous. To get current the operation status, invoke the Get bulk operation status endpoint with the returned operation_id.
There is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.
Domain types
typeListCursorstruct{…}
typeListItemstruct{…}