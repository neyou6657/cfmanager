Workers
Workers
Domain types
typeMigrationStepstruct{…}
typeSingleStepMigrationstruct{…}
A single set of migrations to apply.
typeWorkerMetadatastruct{…}
JSON-encoded metadata about the uploaded parts and Worker configuration.
Workers
Account Settings
Workers.AccountSettings
Methods
client.Workers.AccountSettings.Get(ctx, query) (*
AccountSettingGetResponse
, error)
get/accounts/{account_id}/workers/account-settings
Fetches Worker account settings for an account.
client.Workers.AccountSettings.Update(ctx, params) (*
AccountSettingUpdateResponse
, error)
put/accounts/{account_id}/workers/account-settings
Creates Worker account settings for an account.
Workers
Assets
Workers.Assets
Workers
Assets
Upload
Workers.Assets.Upload
Methods
client.Workers.Assets.Upload.New(ctx, params) (*
AssetUploadNewResponse
, error)
post/accounts/{account_id}/workers/assets/upload
Upload assets ahead of creating a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.
Workers
Beta
Workers.Beta
Workers
Beta
Workers
Workers.Beta.Workers
Methods
client.Workers.Beta.Workers.List(ctx, params) (*V4PagePaginationArray[
Worker
], error)
get/accounts/{account_id}/workers/workers
List all Workers for an account.
client.Workers.Beta.Workers.Get(ctx, workerID, query) (*
Worker
, error)
get/accounts/{account_id}/workers/workers/{worker_id}
Get details about a specific Worker.
client.Workers.Beta.Workers.New(ctx, params) (*
Worker
, error)
post/accounts/{account_id}/workers/workers
Create a new Worker.
client.Workers.Beta.Workers.Update(ctx, workerID, params) (*
Worker
, error)
put/accounts/{account_id}/workers/workers/{worker_id}
Perform a complete replacement of a Worker, where omitted properties are set to their default values. This is the exact same as the Create Worker endpoint, but operates on an existing Worker. To perform a partial update instead, use the Edit Worker endpoint.
client.Workers.Beta.Workers.Edit(ctx, workerID, params) (*
Worker
, error)
patch/accounts/{account_id}/workers/workers/{worker_id}
Perform a partial update on a Worker, where omitted properties are left unchanged from their current values.
client.Workers.Beta.Workers.Delete(ctx, workerID, body) (*
BetaWorkerDeleteResponse
, error)
delete/accounts/{account_id}/workers/workers/{worker_id}
Delete a Worker and all its associated resources (versions, deployments, etc.).
Domain types
typeWorkerstruct{…}
Workers
Beta
Workers
Versions
Workers.Beta.Workers.Versions
Methods
client.Workers.Beta.Workers.Versions.List(ctx, workerID, params) (*V4PagePaginationArray[
Version
], error)
get/accounts/{account_id}/workers/workers/{worker_id}/versions
List all versions for a Worker.
client.Workers.Beta.Workers.Versions.Get(ctx, workerID, versionID, params) (*
Version
, error)
get/accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}
Get details about a specific version.
client.Workers.Beta.Workers.Versions.New(ctx, workerID, params) (*
Version
, error)
post/accounts/{account_id}/workers/workers/{worker_id}/versions
Create a new version.
client.Workers.Beta.Workers.Versions.Delete(ctx, workerID, versionID, body) (*
BetaWorkerVersionDeleteResponse
, error)
delete/accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}
Delete a version.
Domain types
typeVersionstruct{…}
Workers
Domains
Workers.Domains
Methods
client.Workers.Domains.List(ctx, params) (*SinglePage[
Domain
], error)
get/accounts/{account_id}/workers/domains
Lists all Worker Domains for an account.
client.Workers.Domains.Get(ctx, domainID, query) (*
Domain
, error)
get/accounts/{account_id}/workers/domains/{domain_id}
Gets a Worker domain.
client.Workers.Domains.Update(ctx, params) (*
Domain
, error)
put/accounts/{account_id}/workers/domains
Attaches a Worker to a zone and hostname.
client.Workers.Domains.Delete(ctx, domainID, body) error
delete/accounts/{account_id}/workers/domains/{domain_id}
Detaches a Worker from a zone and hostname.
Domain types
typeDomainstruct{…}
Workers
Observability
Workers.Observability
Workers
Observability
Telemetry
Workers.Observability.Telemetry
Methods
client.Workers.Observability.Telemetry.Keys(ctx, params) (*SinglePage[
ObservabilityTelemetryKeysResponse
], error)
post/accounts/{account_id}/workers/observability/telemetry/keys
List all the keys in your telemetry events.
client.Workers.Observability.Telemetry.Query(ctx, params) (*
ObservabilityTelemetryQueryResponse
, error)
post/accounts/{account_id}/workers/observability/telemetry/query
Runs a temporary or saved query
client.Workers.Observability.Telemetry.Values(ctx, params) (*SinglePage[
ObservabilityTelemetryValuesResponse
], error)
post/accounts/{account_id}/workers/observability/telemetry/values
List unique values found in your events
Workers
Routes
Workers.Routes
Methods
client.Workers.Routes.List(ctx, query) (*SinglePage[
RouteListResponse
], error)
get/zones/{zone_id}/workers/routes
Returns routes for a zone.
client.Workers.Routes.Get(ctx, routeID, query) (*
RouteGetResponse
, error)
get/zones/{zone_id}/workers/routes/{route_id}
Returns information about a route, including URL pattern and Worker.
client.Workers.Routes.New(ctx, params) (*
RouteNewResponse
, error)
post/zones/{zone_id}/workers/routes
Creates a route that maps a URL pattern to a Worker.
client.Workers.Routes.Update(ctx, routeID, params) (*
RouteUpdateResponse
, error)
put/zones/{zone_id}/workers/routes/{route_id}
Updates the URL pattern or Worker associated with a route.
client.Workers.Routes.Delete(ctx, routeID, body) (*
RouteDeleteResponse
, error)
delete/zones/{zone_id}/workers/routes/{route_id}
Deletes a route.
Workers
Scripts
Workers.Scripts
Methods
client.Workers.Scripts.List(ctx, params) (*SinglePage[
ScriptListResponse
], error)
get/accounts/{account_id}/workers/scripts
Fetch a list of uploaded workers.
client.Workers.Scripts.Search(ctx, params) (*[]
ScriptSearchResponse
, error)
get/accounts/{account_id}/workers/scripts-search
Search for Workers in an account.
client.Workers.Scripts.Get(ctx, scriptName, query) (*string, error)
get/accounts/{account_id}/workers/scripts/{script_name}
Fetch raw script content for your worker. Note this is the original script content, not JSON encoded.
client.Workers.Scripts.Update(ctx, scriptName, params) (*
ScriptUpdateResponse
, error)
put/accounts/{account_id}/workers/scripts/{script_name}
Upload a worker module. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
client.Workers.Scripts.Delete(ctx, scriptName, params) (*
ScriptDeleteResponse
, error)
delete/accounts/{account_id}/workers/scripts/{script_name}
Delete your worker. This call has no response body on a successful delete.
Domain types
typeScriptstruct{…}
typeScriptSettingstruct{…}
Workers
Scripts
Assets
Workers.Scripts.Assets
Workers
Scripts
Assets
Upload
Workers.Scripts.Assets.Upload
Methods
client.Workers.Scripts.Assets.Upload.New(ctx, scriptName, params) (*
ScriptAssetUploadNewResponse
, error)
post/accounts/{account_id}/workers/scripts/{script_name}/assets-upload-session
Start uploading a collection of assets for use in a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.
Workers
Scripts
Content
Workers.Scripts.Content
Methods
client.Workers.Scripts.Content.Get(ctx, scriptName, query) (*
Response
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/content/v2
Fetch script content only.
client.Workers.Scripts.Content.Update(ctx, scriptName, params) (*
Script
, error)
put/accounts/{account_id}/workers/scripts/{script_name}/content
Put script content without touching config or metadata.
Workers
Scripts
Deployments
Workers.Scripts.Deployments
Methods
client.Workers.Scripts.Deployments.List(ctx, scriptName, query) (*
ScriptDeploymentListResponse
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/deployments
List of Worker Deployments. The first deployment in the list is the latest deployment actively serving traffic.
client.Workers.Scripts.Deployments.New(ctx, scriptName, params) (*
Deployment
, error)
post/accounts/{account_id}/workers/scripts/{script_name}/deployments
Deployments configure how Worker Versions are deployed to traffic. A deployment can consist of one or two versions of a Worker.
client.Workers.Scripts.Deployments.Get(ctx, scriptName, deploymentID, query) (*
Deployment
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}
Get information about a Worker Deployment.
client.Workers.Scripts.Deployments.Delete(ctx, scriptName, deploymentID, body) (*
ScriptDeploymentDeleteResponse
, error)
delete/accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}
Delete a Worker Deployment. The latest deployment, which is actively serving traffic, cannot be deleted. All other deployments can be deleted.
Domain types
typeDeploymentstruct{…}
Workers
Scripts
Schedules
Workers.Scripts.Schedules
Methods
client.Workers.Scripts.Schedules.Get(ctx, scriptName, query) (*
ScriptScheduleGetResponse
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/schedules
Fetches Cron Triggers for a Worker.
client.Workers.Scripts.Schedules.Update(ctx, scriptName, params) (*
ScriptScheduleUpdateResponse
, error)
put/accounts/{account_id}/workers/scripts/{script_name}/schedules
Updates Cron Triggers for a Worker.
Workers
Scripts
Script And Version Settings
Workers.Scripts.ScriptAndVersionSettings
Methods
client.Workers.Scripts.ScriptAndVersionSettings.Get(ctx, scriptName, query) (*
ScriptScriptAndVersionSettingGetResponse
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/settings
Get metadata and config, such as bindings or usage model.
client.Workers.Scripts.ScriptAndVersionSettings.Edit(ctx, scriptName, params) (*
ScriptScriptAndVersionSettingEditResponse
, error)
patch/accounts/{account_id}/workers/scripts/{script_name}/settings
Patch metadata or config, such as bindings or usage model.
Workers
Scripts
Secrets
Workers.Scripts.Secrets
Methods
client.Workers.Scripts.Secrets.List(ctx, scriptName, query) (*SinglePage[
ScriptSecretListResponse
], error)
get/accounts/{account_id}/workers/scripts/{script_name}/secrets
List secrets bound to a script.
client.Workers.Scripts.Secrets.Get(ctx, scriptName, secretName, params) (*
ScriptSecretGetResponse
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}
Get a given secret binding (value omitted) on a script.
client.Workers.Scripts.Secrets.Update(ctx, scriptName, params) (*
ScriptSecretUpdateResponse
, error)
put/accounts/{account_id}/workers/scripts/{script_name}/secrets
Add a secret to a script.
client.Workers.Scripts.Secrets.Delete(ctx, scriptName, secretName, params) (*
ScriptSecretDeleteResponse
, error)
delete/accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}
Remove a secret from a script.
Workers
Scripts
Settings
Workers.Scripts.Settings
Methods
client.Workers.Scripts.Settings.Get(ctx, scriptName, query) (*
ScriptSetting
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/script-settings
Get script-level settings when using Worker Versions. Includes Logpush and Tail Consumers.
client.Workers.Scripts.Settings.Edit(ctx, scriptName, params) (*
ScriptSetting
, error)
patch/accounts/{account_id}/workers/scripts/{script_name}/script-settings
Patch script-level settings when using Worker Versions. Including but not limited to Logpush and Tail Consumers.
Workers
Scripts
Subdomain
Workers.Scripts.Subdomain
Methods
client.Workers.Scripts.Subdomain.Get(ctx, scriptName, query) (*
ScriptSubdomainGetResponse
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/subdomain
Get if the Worker is available on the workers.dev subdomain.
client.Workers.Scripts.Subdomain.New(ctx, scriptName, params) (*
ScriptSubdomainNewResponse
, error)
post/accounts/{account_id}/workers/scripts/{script_name}/subdomain
Enable or disable the Worker on the workers.dev subdomain.
client.Workers.Scripts.Subdomain.Delete(ctx, scriptName, body) (*
ScriptSubdomainDeleteResponse
, error)
delete/accounts/{account_id}/workers/scripts/{script_name}/subdomain
Disable all workers.dev subdomains for a Worker.
Workers
Scripts
Tail
Workers.Scripts.Tail
Methods
client.Workers.Scripts.Tail.Get(ctx, scriptName, query) (*
ScriptTailGetResponse
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/tails
Get list of tails currently deployed on a Worker.
client.Workers.Scripts.Tail.New(ctx, scriptName, params) (*
ScriptTailNewResponse
, error)
post/accounts/{account_id}/workers/scripts/{script_name}/tails
Starts a tail that receives logs and exception from a Worker.
client.Workers.Scripts.Tail.Delete(ctx, scriptName, id, body) (*
ScriptTailDeleteResponse
, error)
delete/accounts/{account_id}/workers/scripts/{script_name}/tails/{id}
Deletes a tail from a Worker.
Domain types
typeConsumerScriptstruct{…}
A reference to a script that will consume logs from the attached Worker.
Workers
Scripts
Versions
Workers.Scripts.Versions
Methods
client.Workers.Scripts.Versions.List(ctx, scriptName, params) (*V4PagePagination[
ScriptVersionListResponse
], error)
get/accounts/{account_id}/workers/scripts/{script_name}/versions
List of Worker Versions. The first version in the list is the latest version.
client.Workers.Scripts.Versions.Get(ctx, scriptName, versionID, query) (*
ScriptVersionGetResponse
, error)
get/accounts/{account_id}/workers/scripts/{script_name}/versions/{version_id}
Get Version Detail
client.Workers.Scripts.Versions.New(ctx, scriptName, params) (*
ScriptVersionNewResponse
, error)
post/accounts/{account_id}/workers/scripts/{script_name}/versions
Upload a Worker Version without deploying to Cloudflare's network. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
Workers
Subdomains
Workers.Subdomains
Methods
client.Workers.Subdomains.Get(ctx, query) (*
SubdomainGetResponse
, error)
get/accounts/{account_id}/workers/subdomain
Returns a Workers subdomain for an account.
client.Workers.Subdomains.Update(ctx, params) (*
SubdomainUpdateResponse
, error)
put/accounts/{account_id}/workers/subdomain
Creates a Workers subdomain for an account.
client.Workers.Subdomains.Delete(ctx, body) error
delete/accounts/{account_id}/workers/subdomain
Deletes a Workers subdomain for an account.