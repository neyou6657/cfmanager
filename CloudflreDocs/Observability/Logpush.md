Logpush
Logpush
Logpush
Datasets
Logpush.Datasets
Logpush
Datasets
Fields
Logpush.Datasets.Fields
Methods
client.Logpush.Datasets.Fields.Get(ctx, datasetID, query) (*
DatasetFieldGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/logpush/datasets/{dataset_id}/fields
Lists all fields available for a dataset. The response result is. an object with key-value pairs, where keys are field names, and values are descriptions.
Logpush
Datasets
Jobs
Logpush.Datasets.Jobs
Methods
client.Logpush.Datasets.Jobs.Get(ctx, datasetID, query) (*SinglePage[
LogpushJob
], error)
get/{accounts_or_zones}/{account_or_zone_id}/logpush/datasets/{dataset_id}/jobs
Lists Logpush jobs for an account or zone for a dataset.
Logpush
Edge
Logpush.Edge
Methods
client.Logpush.Edge.Get(ctx, query) (*SinglePage[
InstantLogpushJob
], error)
get/zones/{zone_id}/logpush/edge/jobs
Lists Instant Logs jobs for a zone.
client.Logpush.Edge.New(ctx, params) (*
InstantLogpushJob
, error)
post/zones/{zone_id}/logpush/edge/jobs
Creates a new Instant Logs job for a zone.
Domain types
typeInstantLogpushJobstruct{…}
Logpush
Jobs
Logpush.Jobs
Methods
client.Logpush.Jobs.List(ctx, query) (*SinglePage[
LogpushJob
], error)
get/{accounts_or_zones}/{account_or_zone_id}/logpush/jobs
Lists Logpush jobs for an account or zone.
client.Logpush.Jobs.Get(ctx, jobID, query) (*
LogpushJob
, error)
get/{accounts_or_zones}/{account_or_zone_id}/logpush/jobs/{job_id}
Gets the details of a Logpush job.
client.Logpush.Jobs.New(ctx, params) (*
LogpushJob
, error)
post/{accounts_or_zones}/{account_or_zone_id}/logpush/jobs
Creates a new Logpush job for an account or zone.
client.Logpush.Jobs.Update(ctx, jobID, params) (*
LogpushJob
, error)
put/{accounts_or_zones}/{account_or_zone_id}/logpush/jobs/{job_id}
Updates a Logpush job.
client.Logpush.Jobs.Delete(ctx, jobID, body) (*
JobDeleteResponse
, error)
delete/{accounts_or_zones}/{account_or_zone_id}/logpush/jobs/{job_id}
Deletes a Logpush job.
Domain types
typeLogpushJobstruct{…}
typeOutputOptionsstruct{…}
The structured replacement for logpull_options. When including this field, the logpull_option field will be ignored.
Logpush
Ownership
Logpush.Ownership
Methods
client.Logpush.Ownership.New(ctx, params) (*
OwnershipNewResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/logpush/ownership
Gets a new ownership challenge sent to your destination.
client.Logpush.Ownership.Validate(ctx, params) (*
OwnershipValidation
, error)
post/{accounts_or_zones}/{account_or_zone_id}/logpush/ownership/validate
Validates ownership challenge of the destination.
Domain types
typeOwnershipValidationstruct{…}
Logpush
Validate
Logpush.Validate
Methods
client.Logpush.Validate.Destination(ctx, params) (*
ValidateDestinationResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/logpush/validate/destination
Validates destination.
client.Logpush.Validate.DestinationExists(ctx, params) (*
ValidateDestinationExistsResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/logpush/validate/destination/exists
Checks if there is an existing job with a destination.
client.Logpush.Validate.Origin(ctx, params) (*
ValidateOriginResponse
, error)
post/{accounts_or_zones}/{account_or_zone_id}/logpush/validate/origin
Validates logpull origin with logpull_options.