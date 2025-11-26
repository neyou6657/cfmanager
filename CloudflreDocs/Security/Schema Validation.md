Schema Validation
SchemaValidation
Schema Validation
Schemas
SchemaValidation.Schemas
Methods
client.SchemaValidation.Schemas.List(ctx, params) (*V4PagePaginationArray[
PublicSchema
], error)
get/zones/{zone_id}/schema_validation/schemas
List all uploaded schemas
client.SchemaValidation.Schemas.Get(ctx, schemaID, params) (*
PublicSchema
, error)
get/zones/{zone_id}/schema_validation/schemas/{schema_id}
Get details of a schema
client.SchemaValidation.Schemas.New(ctx, params) (*
PublicSchema
, error)
post/zones/{zone_id}/schema_validation/schemas
Upload a schema
client.SchemaValidation.Schemas.Edit(ctx, schemaID, params) (*
PublicSchema
, error)
patch/zones/{zone_id}/schema_validation/schemas/{schema_id}
Edit details of a schema to enable validation
client.SchemaValidation.Schemas.Delete(ctx, schemaID, body) (*
SchemaDeleteResponse
, error)
delete/zones/{zone_id}/schema_validation/schemas/{schema_id}
Delete a schema
Domain types
typePublicSchemastruct{…}
A schema used in schema validation
Schema Validation
Settings
SchemaValidation.Settings
Methods
client.SchemaValidation.Settings.Get(ctx, query) (*
SettingGetResponse
, error)
get/zones/{zone_id}/schema_validation/settings
Get global schema validation settings
client.SchemaValidation.Settings.Update(ctx, params) (*
SettingUpdateResponse
, error)
put/zones/{zone_id}/schema_validation/settings
Update global schema validation settings
client.SchemaValidation.Settings.Edit(ctx, params) (*
SettingEditResponse
, error)
patch/zones/{zone_id}/schema_validation/settings
Edit global schema validation settings
Schema Validation
Settings
Operations
SchemaValidation.Settings.Operations
Methods
client.SchemaValidation.Settings.Operations.List(ctx, params) (*V4PagePaginationArray[
SettingOperationListResponse
], error)
get/zones/{zone_id}/schema_validation/settings/operations
List per-operation schema validation settings
client.SchemaValidation.Settings.Operations.Get(ctx, operationID, query) (*
SettingOperationGetResponse
, error)
get/zones/{zone_id}/schema_validation/settings/operations/{operation_id}
Get per-operation schema validation setting
client.SchemaValidation.Settings.Operations.Update(ctx, operationID, params) (*
SettingOperationUpdateResponse
, error)
put/zones/{zone_id}/schema_validation/settings/operations/{operation_id}
Update per-operation schema validation setting
client.SchemaValidation.Settings.Operations.BulkEdit(ctx, params) (*
SettingOperationBulkEditResponse
, error)
patch/zones/{zone_id}/schema_validation/settings/operations
Bulk edit per-operation schema validation settings
client.SchemaValidation.Settings.Operations.Delete(ctx, operationID, body) (*
SettingOperationDeleteResponse
, error)
delete/zones/{zone_id}/schema_validation/settings/operations/{operation_id}
Delete per-operation schema validation setting