Pipelines
Pipelines
Methods
client.Pipelines.List(ctx, params) (*
PipelineListResponse
, error)
Deprecated
get/accounts/{account_id}/pipelines
Deprecated
Use list_v1 instead. This endpoint will be removed in the future.
[DEPRECATED] List, filter, and paginate pipelines in an account. Use the new /pipelines/v1/pipelines endpoint instead.
client.Pipelines.Get(ctx, pipelineName, query) (*
PipelineGetResponse
, error)
Deprecated
get/accounts/{account_id}/pipelines/{pipeline_name}
Deprecated
Use get_v1 instead. This endpoint will be removed in the future.
[DEPRECATED] Get configuration of a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.
client.Pipelines.New(ctx, params) (*
PipelineNewResponse
, error)
Deprecated
post/accounts/{account_id}/pipelines
Deprecated
Use create_v1 instead. This endpoint will be removed in the future.
[DEPRECATED] Create a new pipeline. Use the new /pipelines/v1/pipelines endpoint instead.
client.Pipelines.Update(ctx, pipelineName, params) (*
PipelineUpdateResponse
, error)
Deprecated
put/accounts/{account_id}/pipelines/{pipeline_name}
Deprecated
The v1 API does not support updates. This endpoint will be removed in the future.
[DEPRECATED] Update an existing pipeline. Use the new /pipelines/v1/pipelines endpoint instead.
client.Pipelines.Delete(ctx, pipelineName, body) error
Deprecated
delete/accounts/{account_id}/pipelines/{pipeline_name}
Deprecated
Use delete_v1 instead. This endpoint will be removed in the future.
[DEPRECATED] Delete a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.
client.Pipelines.ListV1(ctx, params) (*V4PagePaginationArray[
PipelineListV1Response
], error)
get/accounts/{account_id}/pipelines/v1/pipelines
List/Filter Pipelines in Account.
client.Pipelines.GetV1(ctx, pipelineID, query) (*
PipelineGetV1Response
, error)
get/accounts/{account_id}/pipelines/v1/pipelines/{pipeline_id}
Get Pipelines Details.
client.Pipelines.NewV1(ctx, params) (*
PipelineNewV1Response
, error)
post/accounts/{account_id}/pipelines/v1/pipelines
Create a new Pipeline.
client.Pipelines.DeleteV1(ctx, pipelineID, body) error
delete/accounts/{account_id}/pipelines/v1/pipelines/{pipeline_id}
Delete Pipeline in Account.
client.Pipelines.ValidateSql(ctx, params) (*
PipelineValidateSqlResponse
, error)
post/accounts/{account_id}/pipelines/v1/validate_sql
Validate Arroyo SQL.
Pipelines
Sinks
Pipelines.Sinks
Methods
client.Pipelines.Sinks.List(ctx, params) (*V4PagePaginationArray[
SinkListResponse
], error)
get/accounts/{account_id}/pipelines/v1/sinks
List/Filter Sinks in Account.
client.Pipelines.Sinks.Get(ctx, sinkID, query) (*
SinkGetResponse
, error)
get/accounts/{account_id}/pipelines/v1/sinks/{sink_id}
Get Sink Details.
client.Pipelines.Sinks.New(ctx, params) (*
SinkNewResponse
, error)
post/accounts/{account_id}/pipelines/v1/sinks
Create a new Sink.
client.Pipelines.Sinks.Delete(ctx, sinkID, params) error
delete/accounts/{account_id}/pipelines/v1/sinks/{sink_id}
Delete Pipeline in Account.
Pipelines
Streams
Pipelines.Streams
Methods
client.Pipelines.Streams.List(ctx, params) (*V4PagePaginationArray[
StreamListResponse
], error)
get/accounts/{account_id}/pipelines/v1/streams
List/Filter Streams in Account.
client.Pipelines.Streams.Get(ctx, streamID, query) (*
StreamGetResponse
, error)
get/accounts/{account_id}/pipelines/v1/streams/{stream_id}
Get Stream Details.
client.Pipelines.Streams.New(ctx, params) (*
StreamNewResponse
, error)
post/accounts/{account_id}/pipelines/v1/streams
Create a new Stream.
client.Pipelines.Streams.Update(ctx, streamID, params) (*
StreamUpdateResponse
, error)
patch/accounts/{account_id}/pipelines/v1/streams/{stream_id}
Update a Stream.
client.Pipelines.Streams.Delete(ctx, streamID, params) error
delete/accounts/{account_id}/pipelines/v1/streams/{stream_id}
Delete Stream in Account.