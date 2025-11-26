Connectivity
Connectivity
Connectivity
Directory
Connectivity.Directory
Connectivity
Directory
Services
Connectivity.Directory.Services
Methods
client.Connectivity.Directory.Services.List(ctx, params) (*V4PagePaginationArray[
DirectoryServiceListResponse
], error)
get/accounts/{account_id}/connectivity/directory/services
List connectivity services
client.Connectivity.Directory.Services.New(ctx, params) (*
DirectoryServiceNewResponse
, error)
post/accounts/{account_id}/connectivity/directory/services
Create connectivity service
client.Connectivity.Directory.Services.Get(ctx, serviceID, query) (*
DirectoryServiceGetResponse
, error)
get/accounts/{account_id}/connectivity/directory/services/{service_id}
Get connectivity service
client.Connectivity.Directory.Services.Update(ctx, serviceID, params) (*
DirectoryServiceUpdateResponse
, error)
put/accounts/{account_id}/connectivity/directory/services/{service_id}
Update connectivity service
client.Connectivity.Directory.Services.Delete(ctx, serviceID, body) error
delete/accounts/{account_id}/connectivity/directory/services/{service_id}
Delete connectivity service