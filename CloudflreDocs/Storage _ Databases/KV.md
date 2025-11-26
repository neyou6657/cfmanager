KV
KV
KV
Namespaces
KV.Namespaces
Methods
client.KV.Namespaces.List(ctx, params) (*V4PagePaginationArray[
Namespace
], error)
get/accounts/{account_id}/storage/kv/namespaces
Returns the namespaces owned by an account.
client.KV.Namespaces.Get(ctx, namespaceID, query) (*
Namespace
, error)
get/accounts/{account_id}/storage/kv/namespaces/{namespace_id}
Get the namespace corresponding to the given ID.
client.KV.Namespaces.New(ctx, params) (*
Namespace
, error)
post/accounts/{account_id}/storage/kv/namespaces
Creates a namespace under the given title. A 400 is returned if the account already owns a namespace with this title. A namespace must be explicitly deleted to be replaced.
client.KV.Namespaces.Update(ctx, namespaceID, params) (*
Namespace
, error)
put/accounts/{account_id}/storage/kv/namespaces/{namespace_id}
Modifies a namespace's title.
client.KV.Namespaces.Delete(ctx, namespaceID, body) (*
NamespaceDeleteResponse
, error)
delete/accounts/{account_id}/storage/kv/namespaces/{namespace_id}
Deletes the namespace corresponding to the given ID.
client.KV.Namespaces.BulkUpdate(ctx, namespaceID, params) (*
NamespaceBulkUpdateResponse
, error)
put/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk
Write multiple keys and values at once. Body should be an array of up to 10,000 key-value pairs to be stored, along with optional expiration information. Existing values and expirations will be overwritten. If neither expiration nor expiration_ttl is specified, the key-value pair will never expire. If both are set, expiration_ttl is used and expiration is ignored. The entire request size must be 100 megabytes or less.
client.KV.Namespaces.BulkDelete(ctx, namespaceID, params) (*
NamespaceBulkDeleteResponse
, error)
post/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk/delete
Remove multiple KV pairs from the namespace. Body should be an array of up to 10,000 keys to be removed.
client.KV.Namespaces.BulkGet(ctx, namespaceID, params) (*
NamespaceBulkGetResponse
, error)
post/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk/get
Retrieve up to 100 KV pairs from the namespace. Keys must contain text-based values. JSON values can optionally be parsed instead of being returned as a string value. Metadata can be included if withMetadata is true.
Domain types
typeNamespacestruct{…}
KV
Namespaces
Keys
KV.Namespaces.Keys
Methods
client.KV.Namespaces.Keys.List(ctx, namespaceID, params) (*CursorPaginationAfter[
Key
], error)
get/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/keys
Lists a namespace's keys.
client.KV.Namespaces.Keys.BulkUpdate(ctx, namespaceID, params) (*
NamespaceKeyBulkUpdateResponse
, error)
Deprecated
put/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk
Deprecated
Please use kv.namespaces.bulk_update instead
Write multiple keys and values at once. Body should be an array of up to 10,000 key-value pairs to be stored, along with optional expiration information. Existing values and expirations will be overwritten. If neither expiration nor expiration_ttl is specified, the key-value pair will never expire. If both are set, expiration_ttl is used and expiration is ignored. The entire request size must be 100 megabytes or less.
client.KV.Namespaces.Keys.BulkDelete(ctx, namespaceID, params) (*
NamespaceKeyBulkDeleteResponse
, error)
Deprecated
post/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk/delete
Deprecated
Please use kv.namespaces.bulk_delete instead
Remove multiple KV pairs from the namespace. Body should be an array of up to 10,000 keys to be removed.
client.KV.Namespaces.Keys.BulkGet(ctx, namespaceID, params) (*
NamespaceKeyBulkGetResponse
, error)
Deprecated
post/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/bulk/get
Deprecated
Please use kv.namespaces.bulk_get instead
Retrieve up to 100 KV pairs from the namespace. Keys must contain text-based values. JSON values can optionally be parsed instead of being returned as a string value. Metadata can be included if withMetadata is true.
Domain types
typeKeystruct{…}
A name for a value. A value stored under a given key may be retrieved via the same key.
KV
Namespaces
Metadata
KV.Namespaces.Metadata
Methods
client.KV.Namespaces.Metadata.Get(ctx, namespaceID, keyName, query) (*
NamespaceMetadataGetResponse
, error)
get/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/metadata/{key_name}
Returns the metadata associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, :, !, %) in the key name.
KV
Namespaces
Values
KV.Namespaces.Values
Methods
client.KV.Namespaces.Values.Get(ctx, namespaceID, keyName, query) (*
Response
, error)
get/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}
Returns the value associated with the given key in the given namespace. Use URL-encoding to use special characters (for example, :, !, %) in the key name. If the KV-pair is set to expire at some point, the expiration time as measured in seconds since the UNIX epoch will be returned in the expiration response header.
client.KV.Namespaces.Values.Update(ctx, namespaceID, keyName, params) (*
NamespaceValueUpdateResponse
, error)
put/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}
Write a value identified by a key. Use URL-encoding to use special characters (for example, :, !, %) in the key name. Body should be the value to be stored. If JSON metadata to be associated with the key/value pair is needed, use multipart/form-data content type for your PUT request (see dropdown below in REQUEST BODY SCHEMA). Existing values, expirations, and metadata will be overwritten. If neither expiration nor expiration_ttl is specified, the key-value pair will never expire. If both are set, expiration_ttl is used and expiration is ignored.
client.KV.Namespaces.Values.Delete(ctx, namespaceID, keyName, body) (*
NamespaceValueDeleteResponse
, error)
delete/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}
Remove a KV pair from the namespace. Use URL-encoding to use special characters (for example, :, !, %) in the key name.