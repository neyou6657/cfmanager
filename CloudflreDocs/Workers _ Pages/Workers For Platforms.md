# Workers For Platforms

## WorkersForPlatforms.Dispatch.Namespaces

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "my-dispatch-namespace"
      }'
```
`POST /accounts/{account_id}/workers/dispatch/namespaces`

Create a new Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces`

Fetch a list of Workers for Platforms namespaces.

#### (resource) workers_for_platforms.dispatch.namespaces > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}`

Delete a Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}`

Get a Workers for Platforms namespace.

---

## WorkersForPlatforms.Dispatch.Namespaces.Scripts

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces.scripts > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME \
  -X PUT \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -F metadata='{}'
```
`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}`

Upload a worker module to a Workers for Platforms namespace. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}`

Delete a worker from a Workers for Platforms namespace. This call has no response body on a successful delete.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}`

Fetch information about a script uploaded to a Workers for Platforms namespace.

---

## WorkersForPlatforms.Dispatch.Namespaces.Scripts.AssetUpload

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.asset_upload > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/assets-upload-session \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "manifest": {
          "foo": {
            "hash": "hash",
            "size": 0
          }
        }
      }'
```
`POST /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/assets-upload-session`

Start uploading a collection of assets for use in a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.

---

## WorkersForPlatforms.Dispatch.Namespaces.Scripts.Content

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.content > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/content \
  -X PUT \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -F metadata='{}'
```
`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content`

Put script content for a script uploaded to a Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.content > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/content \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/content`

Fetch script content from a script uploaded to a Workers for Platforms namespace.

---

## WorkersForPlatforms.Dispatch.Namespaces.Scripts.Settings

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.settings > (method) edit
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/settings \
  -X PATCH \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`PATCH /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings`

Patch script metadata, such as bindings.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.settings > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/settings \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings`

Get script settings from a script uploaded to a Workers for Platforms namespace.

---

## WorkersForPlatforms.Dispatch.Namespaces.Scripts.Bindings

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.bindings > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/bindings \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/bindings`

Fetch script bindings from a script uploaded to a Workers for Platforms namespace.

---

## WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.secrets > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/secrets \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "myBinding",
        "text": "My secret.",
        "type": "secret_text"
      }'
```
`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets`

Add a secret to a script uploaded to a Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.secrets > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/secrets \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets`

List secrets bound to a script uploaded to a Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.secrets > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/secrets/$SECRET_NAME \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets/{secret_name}`

Remove a secret from a script uploaded to a Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.secrets > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/secrets/$SECRET_NAME \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets/{secret_name}`

Get a given secret binding (value omitted) on a script uploaded to a Workers for Platforms namespace.

---

## WorkersForPlatforms.Dispatch.Namespaces.Scripts.Tags

### Methods

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.tags > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/tags \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '[
        "my-team",
        "my-public-api"
      ]'
```
`PUT /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags`

Put script tags for a script uploaded to a Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.tags > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/tags \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags`

Fetch tags from a script uploaded to a Workers for Platforms namespace.

#### (resource) workers_for_platforms.dispatch.namespaces.scripts.tags > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/dispatch/namespaces/$DISPATCH_NAMESPACE/scripts/$SCRIPT_NAME/tags/$TAG \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/tags/{tag}`

Delete script tag for a script uploaded to a Workers for Platforms namespace.
