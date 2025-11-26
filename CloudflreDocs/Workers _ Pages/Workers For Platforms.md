Workers For Platforms
WorkersForPlatforms
Workers For Platforms
Dispatch
WorkersForPlatforms.Dispatch
Workers For Platforms
Dispatch
Namespaces
WorkersForPlatforms.Dispatch.Namespaces
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.List(ctx, query) (*SinglePage[
DispatchNamespaceListResponse
], error)
get/accounts/{account_id}/workers/dispatch/namespaces
Fetch a list of Workers for Platforms namespaces.
client.WorkersForPlatforms.Dispatch.Namespaces.Get(ctx, dispatchNamespace, query) (*
DispatchNamespaceGetResponse
, error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}
Get a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.New(ctx, params) (*
DispatchNamespaceNewResponse
, error)
post/accounts/{account_id}/workers/dispatch/namespaces
Create a new Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Delete(ctx, dispatchNamespace, body) (*
DispatchNamespaceDeleteResponse
, error)
delete/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}
Delete a Workers for Platforms namespace.
Workers For Platforms
Dispatch
Namespaces
Scripts
WorkersForPlatforms.Dispatch.Namespaces.Scripts
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Get(ctx, dispatchNamespace, scriptName, query) (*
Script
, error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}
Fetch information about a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Update(ctx, dispatchNamespace, scriptName, params) (*
DispatchNamespaceScriptUpdateResponse
, error)
put/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}
Upload a worker module to a Workers for Platforms namespace. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Delete(ctx, dispatchNamespace, scriptName, params) (*
DispatchNamespaceScriptDeleteResponse
, error)
delete/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}
Delete a worker from a Workers for Platforms namespace. This call has no response body on a successful delete.
Domain types
typeScriptstruct{…}
Details about a worker uploaded to a Workers for Platforms namespace.
Workers For Platforms
Dispatch
Namespaces
Scripts
Asset Upload
WorkersForPlatforms.Dispatch.Namespaces.Scripts.AssetUpload
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.AssetUpload.New(ctx, dispatchNamespace, scriptName, params) (*
DispatchNamespaceScriptAssetUploadNewResponse
, error)
post/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/assets-upload-session
Start uploading a collection of assets for use in a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.
Workers For Platforms
Dispatch
Namespaces
Scripts
Bindings
WorkersForPlatforms.Dispatch.Namespaces.Scripts.Bindings
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Bindings.Get(ctx, dispatchNamespace, scriptName, query) (*SinglePage[
DispatchNamespaceScriptBindingGetResponse
], error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/bindings
Fetch script bindings from a script uploaded to a Workers for Platforms namespace.
Workers For Platforms
Dispatch
Namespaces
Scripts
Content
WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content.Get(ctx, dispatchNamespace, scriptName, query) (*
Response
, error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content
Fetch script content from a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content.Update(ctx, dispatchNamespace, scriptName, params) (*
Script
, error)
put/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content
Put script content for a script uploaded to a Workers for Platforms namespace.
Workers For Platforms
Dispatch
Namespaces
Scripts
Secrets
WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.List(ctx, dispatchNamespace, scriptName, query) (*SinglePage[
DispatchNamespaceScriptSecretListResponse
], error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets
List secrets bound to a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Get(ctx, dispatchNamespace, scriptName, secretName, params) (*
DispatchNamespaceScriptSecretGetResponse
, error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets/{secret_name}
Get a given secret binding (value omitted) on a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Update(ctx, dispatchNamespace, scriptName, params) (*
DispatchNamespaceScriptSecretUpdateResponse
, error)
put/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets
Add a secret to a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.Delete(ctx, dispatchNamespace, scriptName, secretName, params) (*
DispatchNamespaceScriptSecretDeleteResponse
, error)
delete/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets/{secret_name}
Remove a secret from a script uploaded to a Workers for Platforms namespace.
Workers For Platforms
Dispatch
Namespaces
Scripts
Settings
WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.Get(ctx, dispatchNamespace, scriptName, query) (*
DispatchNamespaceScriptSettingGetResponse
, error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings
Get script settings from a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings.Edit(ctx, dispatchNamespace, scriptName, params) (*
DispatchNamespaceScriptSettingEditResponse
, error)
patch/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings
Patch script metadata, such as bindings.
Workers For Platforms
Dispatch
Namespaces
Scripts
Tags
WorkersForPlatforms.Dispatch.Namespaces.Scripts.Tags
Methods
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Tags.List(ctx, dispatchNamespace, scriptName, query) (*SinglePage[string], error)
get/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags
Fetch tags from a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Tags.Update(ctx, dispatchNamespace, scriptName, params) (*SinglePage[string], error)
put/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags
Put script tags for a script uploaded to a Workers for Platforms namespace.
client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Tags.Delete(ctx, dispatchNamespace, scriptName, tag, body) (*
DispatchNamespaceScriptTagDeleteResponse
, error)
delete/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags/{tag}
Delete script tag for a script uploaded to a Workers for Platforms namespace.