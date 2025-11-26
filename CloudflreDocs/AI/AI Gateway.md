AI Gateway
AIGateway
Methods
client.AIGateway.List(ctx, params) (*V4PagePaginationArray[
AIGatewayListResponse
], error)
get/accounts/{account_id}/ai-gateway/gateways
List Gateways
client.AIGateway.Get(ctx, id, query) (*
AIGatewayGetResponse
, error)
get/accounts/{account_id}/ai-gateway/gateways/{id}
Fetch a Gateway
client.AIGateway.New(ctx, params) (*
AIGatewayNewResponse
, error)
post/accounts/{account_id}/ai-gateway/gateways
Create a new Gateway
client.AIGateway.Update(ctx, id, params) (*
AIGatewayUpdateResponse
, error)
put/accounts/{account_id}/ai-gateway/gateways/{id}
Update a Gateway
client.AIGateway.Delete(ctx, id, body) (*
AIGatewayDeleteResponse
, error)
delete/accounts/{account_id}/ai-gateway/gateways/{id}
Delete a Gateway
AI Gateway
Datasets
AIGateway.Datasets
Methods
client.AIGateway.Datasets.List(ctx, gatewayID, params) (*V4PagePaginationArray[
DatasetListResponse
], error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/datasets
List Datasets
client.AIGateway.Datasets.Get(ctx, gatewayID, id, query) (*
DatasetGetResponse
, error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/datasets/{id}
Fetch a Dataset
client.AIGateway.Datasets.New(ctx, gatewayID, params) (*
DatasetNewResponse
, error)
post/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/datasets
Create a new Dataset
client.AIGateway.Datasets.Update(ctx, gatewayID, id, params) (*
DatasetUpdateResponse
, error)
put/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/datasets/{id}
Update a Dataset
client.AIGateway.Datasets.Delete(ctx, gatewayID, id, body) (*
DatasetDeleteResponse
, error)
delete/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/datasets/{id}
Delete a Dataset
AI Gateway
Evaluation Types
AIGateway.EvaluationTypes
Methods
client.AIGateway.EvaluationTypes.List(ctx, params) (*V4PagePaginationArray[
EvaluationTypeListResponse
], error)
get/accounts/{account_id}/ai-gateway/evaluation-types
List Evaluators
AI Gateway
Evaluations
AIGateway.Evaluations
Methods
client.AIGateway.Evaluations.List(ctx, gatewayID, params) (*V4PagePaginationArray[
EvaluationListResponse
], error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/evaluations
List Evaluations
client.AIGateway.Evaluations.Get(ctx, gatewayID, id, query) (*
EvaluationGetResponse
, error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/evaluations/{id}
Fetch a Evaluation
client.AIGateway.Evaluations.New(ctx, gatewayID, params) (*
EvaluationNewResponse
, error)
post/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/evaluations
Create a new Evaluation
client.AIGateway.Evaluations.Delete(ctx, gatewayID, id, body) (*
EvaluationDeleteResponse
, error)
delete/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/evaluations/{id}
Delete a Evaluation
AI Gateway
Logs
AIGateway.Logs
Methods
client.AIGateway.Logs.List(ctx, gatewayID, params) (*V4PagePaginationArray[
LogListResponse
], error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs
List Gateway Logs
client.AIGateway.Logs.Get(ctx, gatewayID, id, query) (*
LogGetResponse
, error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}
Get Gateway Log Detail
client.AIGateway.Logs.Edit(ctx, gatewayID, id, params) (*
LogEditResponse
, error)
patch/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}
Patch Gateway Log
client.AIGateway.Logs.Delete(ctx, gatewayID, params) (*
LogDeleteResponse
, error)
delete/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs
Delete Gateway Logs
client.AIGateway.Logs.Request(ctx, gatewayID, id, query) (*
LogRequestResponse
, error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}/request
Get Gateway Log Request
client.AIGateway.Logs.Response(ctx, gatewayID, id, query) (*
LogResponseResponse
, error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}/response
Get Gateway Log Response
AI Gateway
URLs
AIGateway.URLs
Methods
client.AIGateway.URLs.Get(ctx, gatewayID, provider, query) (*string, error)
get/accounts/{account_id}/ai-gateway/gateways/{gateway_id}/url/{provider}
Get Gateway URL