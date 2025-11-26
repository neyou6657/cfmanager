Vectorize
Vectorize
Vectorize
Indexes
Vectorize.Indexes
Methods
client.Vectorize.Indexes.List(ctx, query) (*SinglePage[
CreateIndex
], error)
get/accounts/{account_id}/vectorize/v2/indexes
Returns a list of Vectorize Indexes
client.Vectorize.Indexes.Get(ctx, indexName, query) (*
CreateIndex
, error)
get/accounts/{account_id}/vectorize/v2/indexes/{index_name}
Returns the specified Vectorize Index.
client.Vectorize.Indexes.New(ctx, params) (*
CreateIndex
, error)
post/accounts/{account_id}/vectorize/v2/indexes
Creates and returns a new Vectorize Index.
client.Vectorize.Indexes.Delete(ctx, indexName, body) (*unknown, error)
delete/accounts/{account_id}/vectorize/v2/indexes/{index_name}
Deletes the specified Vectorize Index.
client.Vectorize.Indexes.Insert(ctx, indexName, params) (*
IndexInsertResponse
, error)
post/accounts/{account_id}/vectorize/v2/indexes/{index_name}/insert
Inserts vectors into the specified index and returns a mutation id corresponding to the vectors enqueued for insertion.
client.Vectorize.Indexes.Query(ctx, indexName, params) (*
IndexQueryResponse
, error)
post/accounts/{account_id}/vectorize/v2/indexes/{index_name}/query
Finds vectors closest to a given vector in an index.
client.Vectorize.Indexes.Upsert(ctx, indexName, params) (*
IndexUpsertResponse
, error)
post/accounts/{account_id}/vectorize/v2/indexes/{index_name}/upsert
Upserts vectors into the specified index, creating them if they do not exist and returns a mutation id corresponding to the vectors enqueued for upsertion.
client.Vectorize.Indexes.DeleteByIDs(ctx, indexName, params) (*
IndexDeleteByIDsResponse
, error)
post/accounts/{account_id}/vectorize/v2/indexes/{index_name}/delete_by_ids
Delete a set of vectors from an index by their vector identifiers.
client.Vectorize.Indexes.GetByIDs(ctx, indexName, params) (*
IndexGetByIDsResponse
, error)
post/accounts/{account_id}/vectorize/v2/indexes/{index_name}/get_by_ids
Get a set of vectors from an index by their vector identifiers.
client.Vectorize.Indexes.Info(ctx, indexName, query) (*
IndexInfoResponse
, error)
get/accounts/{account_id}/vectorize/v2/indexes/{index_name}/info
Get information about a vectorize index.
client.Vectorize.Indexes.ListVectors(ctx, indexName, params) (*
IndexListVectorsResponse
, error)
get/accounts/{account_id}/vectorize/v2/indexes/{index_name}/list
Returns a paginated list of vector identifiers from the specified index.
Domain types
typeCreateIndexstruct{…}
typeIndexDeleteVectorsByIDstruct{…}
typeIndexDimensionConfigurationstruct{…}
typeIndexInsertstruct{…}
typeIndexQuerystruct{…}
typeIndexUpsertstruct{…}
Vectorize
Indexes
Metadata Index
Vectorize.Indexes.MetadataIndex
Methods
client.Vectorize.Indexes.MetadataIndex.List(ctx, indexName, query) (*
IndexMetadataIndexListResponse
, error)
get/accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/list
List Metadata Indexes for the specified Vectorize Index.
client.Vectorize.Indexes.MetadataIndex.New(ctx, indexName, params) (*
IndexMetadataIndexNewResponse
, error)
post/accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/create
Enable metadata filtering based on metadata property. Limited to 10 properties.
client.Vectorize.Indexes.MetadataIndex.Delete(ctx, indexName, params) (*
IndexMetadataIndexDeleteResponse
, error)
post/accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/delete
Allow Vectorize to delete the specified metadata index.