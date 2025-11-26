Durable Objects
DurableObjects
Durable Objects
Namespaces
DurableObjects.Namespaces
Methods
client.DurableObjects.Namespaces.List(ctx, params) (*V4PagePaginationArray[
Namespace
], error)
get/accounts/{account_id}/workers/durable_objects/namespaces
Returns the Durable Object namespaces owned by an account.
Domain types
typeNamespacestruct{…}
Durable Objects
Namespaces
Objects
DurableObjects.Namespaces.Objects
Methods
client.DurableObjects.Namespaces.Objects.List(ctx, id, params) (*CursorPaginationAfter[
DurableObject
], error)
get/accounts/{account_id}/workers/durable_objects/namespaces/{id}/objects
Returns the Durable Objects in a given namespace.
Domain types
typeDurableObjectstruct{…}