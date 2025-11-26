Managed Transforms
ManagedTransforms
Methods
client.ManagedTransforms.List(ctx, query) (*
ManagedTransformListResponse
, error)
get/zones/{zone_id}/managed_headers
Fetches a list of all Managed Transforms.
client.ManagedTransforms.Edit(ctx, params) (*
ManagedTransformEditResponse
, error)
patch/zones/{zone_id}/managed_headers
Updates the status of one or more Managed Transforms.
client.ManagedTransforms.Delete(ctx, body) error
delete/zones/{zone_id}/managed_headers
Disables all Managed Transforms.