Secrets Store
SecretsStore
Secrets Store
Quota
SecretsStore.Quota
Methods
client.SecretsStore.Quota.Get(ctx, query) (*
QuotaGetResponse
, error)
get/accounts/{account_id}/secrets_store/quota
Lists the number of secrets used in the account.
Secrets Store
Stores
SecretsStore.Stores
Methods
client.SecretsStore.Stores.List(ctx, params) (*V4PagePaginationArray[
StoreListResponse
], error)
get/accounts/{account_id}/secrets_store/stores
Lists all the stores in an account
client.SecretsStore.Stores.New(ctx, params) (*SinglePage[
StoreNewResponse
], error)
post/accounts/{account_id}/secrets_store/stores
Creates a store in the account
client.SecretsStore.Stores.Delete(ctx, storeID, body) (*
StoreDeleteResponse
, error)
delete/accounts/{account_id}/secrets_store/stores/{store_id}
Deletes a single store
Secrets Store
Stores
Secrets
SecretsStore.Stores.Secrets
Methods
client.SecretsStore.Stores.Secrets.List(ctx, storeID, params) (*V4PagePaginationArray[
StoreSecretListResponse
], error)
get/accounts/{account_id}/secrets_store/stores/{store_id}/secrets
Lists all store secrets
client.SecretsStore.Stores.Secrets.Get(ctx, storeID, secretID, query) (*
StoreSecretGetResponse
, error)
get/accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}
Returns details of a single secret
client.SecretsStore.Stores.Secrets.New(ctx, storeID, params) (*SinglePage[
StoreSecretNewResponse
], error)
post/accounts/{account_id}/secrets_store/stores/{store_id}/secrets
Creates a secret in the account
client.SecretsStore.Stores.Secrets.Edit(ctx, storeID, secretID, params) (*
StoreSecretEditResponse
, error)
patch/accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}
Updates a single secret
client.SecretsStore.Stores.Secrets.Delete(ctx, storeID, secretID, body) (*
StoreSecretDeleteResponse
, error)
delete/accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}
Deletes a single secret
client.SecretsStore.Stores.Secrets.BulkDelete(ctx, storeID, body) (*SinglePage[
StoreSecretBulkDeleteResponse
], error)
delete/accounts/{account_id}/secrets_store/stores/{store_id}/secrets
Deletes one or more secrets
client.SecretsStore.Stores.Secrets.Duplicate(ctx, storeID, secretID, params) (*
StoreSecretDuplicateResponse
, error)
post/accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}/duplicate
Duplicates the secret, keeping the value