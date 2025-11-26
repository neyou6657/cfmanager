API Gateway
APIGateway
API Gateway
Configurations
APIGateway.Configurations
Methods
client.APIGateway.Configurations.Get(ctx, params) (*
Configuration
, error)
get/zones/{zone_id}/api_gateway/configuration
Retrieve information about specific configuration properties
client.APIGateway.Configurations.Update(ctx, params) (*
Configuration
, error)
put/zones/{zone_id}/api_gateway/configuration
Update configuration properties
Domain types
typeConfigurationstruct{…}
API Gateway
Discovery
APIGateway.Discovery
Methods
client.APIGateway.Discovery.Get(ctx, query) (*
DiscoveryGetResponse
, error)
get/zones/{zone_id}/api_gateway/discovery
Retrieve the most up to date view of discovered operations, rendered as OpenAPI schemas
Domain types
typeDiscoveryOperationstruct{…}
API Gateway
Discovery
Operations
APIGateway.Discovery.Operations
Methods
client.APIGateway.Discovery.Operations.List(ctx, params) (*V4PagePaginationArray[
DiscoveryOperation
], error)
get/zones/{zone_id}/api_gateway/discovery/operations
Retrieve the most up to date view of discovered operations
client.APIGateway.Discovery.Operations.Edit(ctx, operationID, params) (*
DiscoveryOperationEditResponse
, error)
patch/zones/{zone_id}/api_gateway/discovery/operations/{operation_id}
Update the state on a discovered operation
client.APIGateway.Discovery.Operations.BulkEdit(ctx, params) (*
DiscoveryOperationBulkEditResponse
, error)
patch/zones/{zone_id}/api_gateway/discovery/operations
Update the state on one or more discovered operations
API Gateway
Expression Template
APIGateway.ExpressionTemplate
API Gateway
Expression Template
Fallthrough
APIGateway.ExpressionTemplate.Fallthrough
Methods
client.APIGateway.ExpressionTemplate.Fallthrough.New(ctx, params) (*
ExpressionTemplateFallthroughNewResponse
, error)
post/zones/{zone_id}/api_gateway/expression-template/fallthrough
Generate fallthrough WAF expression template from a set of API hosts
API Gateway
Operations
APIGateway.Operations
Methods
client.APIGateway.Operations.List(ctx, params) (*V4PagePaginationArray[
OperationListResponse
], error)
get/zones/{zone_id}/api_gateway/operations
Retrieve information about all operations on a zone
client.APIGateway.Operations.Get(ctx, operationID, params) (*
OperationGetResponse
, error)
get/zones/{zone_id}/api_gateway/operations/{operation_id}
Retrieve information about an operation
client.APIGateway.Operations.New(ctx, params) (*
OperationNewResponse
, error)
post/zones/{zone_id}/api_gateway/operations/item
Add one operation to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.
client.APIGateway.Operations.Delete(ctx, operationID, body) (*
OperationDeleteResponse
, error)
delete/zones/{zone_id}/api_gateway/operations/{operation_id}
Delete an operation
client.APIGateway.Operations.BulkNew(ctx, params) (*SinglePage[
OperationBulkNewResponse
], error)
post/zones/{zone_id}/api_gateway/operations
Add one or more operations to a zone. Endpoints can contain path variables. Host, method, endpoint will be normalized to a canoncial form when creating an operation and must be unique on the zone. Inserting an operation that matches an existing one will return the record of the already existing operation and update its last_updated date.
client.APIGateway.Operations.BulkDelete(ctx, body) (*
OperationBulkDeleteResponse
, error)
delete/zones/{zone_id}/api_gateway/operations
Delete multiple operations
Domain types
typeAPIShieldstruct{…}
API Gateway
Operations
Schema Validation
APIGateway.Operations.SchemaValidation
Methods
client.APIGateway.Operations.SchemaValidation.Get(ctx, operationID, query) (*
OperationSchemaValidationGetResponse
, error)
Deprecated
get/zones/{zone_id}/api_gateway/operations/{operation_id}/schema_validation
Deprecated
Use Schema Validation API instead.
Retrieves operation-level schema validation settings on the zone
client.APIGateway.Operations.SchemaValidation.Update(ctx, operationID, params) (*
OperationSchemaValidationUpdateResponse
, error)
Deprecated
put/zones/{zone_id}/api_gateway/operations/{operation_id}/schema_validation
Deprecated
Use Schema Validation API instead.
Updates operation-level schema validation settings on the zone
client.APIGateway.Operations.SchemaValidation.Edit(ctx, params) (*
OperationSchemaValidationEditResponse
, error)
Deprecated
patch/zones/{zone_id}/api_gateway/operations/schema_validation
Deprecated
Use Schema Validation API instead.
Updates multiple operation-level schema validation settings on the zone
API Gateway
Schemas
APIGateway.Schemas
Methods
client.APIGateway.Schemas.List(ctx, params) (*
SchemaListResponse
, error)
get/zones/{zone_id}/api_gateway/schemas
Retrieve operations and features as OpenAPI schemas
API Gateway
Settings
APIGateway.Settings
API Gateway
Settings
Schema Validation
APIGateway.Settings.SchemaValidation
Methods
client.APIGateway.Settings.SchemaValidation.Get(ctx, query) (*
SettingSchemaValidationGetResponse
, error)
Deprecated
get/zones/{zone_id}/api_gateway/settings/schema_validation
Deprecated
Use Schema Validation API instead.
Retrieves zone level schema validation settings currently set on the zone
client.APIGateway.Settings.SchemaValidation.Update(ctx, params) (*
SettingSchemaValidationUpdateResponse
, error)
Deprecated
put/zones/{zone_id}/api_gateway/settings/schema_validation
Deprecated
Use Schema Validation API instead.
Updates zone level schema validation settings on the zone
client.APIGateway.Settings.SchemaValidation.Edit(ctx, params) (*
SettingSchemaValidationEditResponse
, error)
Deprecated
patch/zones/{zone_id}/api_gateway/settings/schema_validation
Deprecated
Use Schema Validation API instead.
Updates zone level schema validation settings on the zone
API Gateway
User Schemas
APIGateway.UserSchemas
Methods
client.APIGateway.UserSchemas.List(ctx, params) (*V4PagePaginationArray[
OldPublicSchema
], error)
Deprecated
get/zones/{zone_id}/api_gateway/user_schemas
Deprecated
Use Schema Validation API instead.
Retrieve information about all schemas on a zone
client.APIGateway.UserSchemas.Get(ctx, schemaID, params) (*
OldPublicSchema
, error)
Deprecated
get/zones/{zone_id}/api_gateway/user_schemas/{schema_id}
Deprecated
Use Schema Validation API instead.
Retrieve information about a specific schema on a zone
client.APIGateway.UserSchemas.New(ctx, params) (*
UserSchemaNewResponse
, error)
Deprecated
post/zones/{zone_id}/api_gateway/user_schemas
Deprecated
Use Schema Validation API instead.
Upload a schema to a zone
client.APIGateway.UserSchemas.Edit(ctx, schemaID, params) (*
OldPublicSchema
, error)
Deprecated
patch/zones/{zone_id}/api_gateway/user_schemas/{schema_id}
Deprecated
Use Schema Validation API instead.
Enable validation for a schema
client.APIGateway.UserSchemas.Delete(ctx, schemaID, body) (*
UserSchemaDeleteResponse
, error)
Deprecated
delete/zones/{zone_id}/api_gateway/user_schemas/{schema_id}
Deprecated
Use Schema Validation API instead.
Delete a schema
Domain types
typeMessage[]
MessageItem
typeOldPublicSchemastruct{…}
API Gateway
User Schemas
Hosts
APIGateway.UserSchemas.Hosts
Methods
client.APIGateway.UserSchemas.Hosts.List(ctx, params) (*V4PagePaginationArray[
UserSchemaHostListResponse
], error)
Deprecated
get/zones/{zone_id}/api_gateway/user_schemas/hosts
Deprecated
Use Schema Validation API instead.
Retrieve schema hosts in a zone
API Gateway
User Schemas
Operations
APIGateway.UserSchemas.Operations
Methods
client.APIGateway.UserSchemas.Operations.List(ctx, schemaID, params) (*V4PagePaginationArray[
UserSchemaOperationListResponse
], error)
Deprecated
get/zones/{zone_id}/api_gateway/user_schemas/{schema_id}/operations
Deprecated
Use Schema Validation API instead.
Retrieves all operations from the schema. Operations that already exist in API Shield Endpoint Management will be returned as full operations.