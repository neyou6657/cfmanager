# Workers

## Workers.Beta.Workers

### Methods

#### (resource) workers.beta.workers > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "my-worker",
        "tags": [
          "my-team",
          "my-public-api"
        ]
      }'
```
`POST /accounts/{account_id}/workers/workers`

Create a new Worker.

#### (resource) workers.beta.workers > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "my-worker",
        "tags": [
          "my-team",
          "my-public-api"
        ]
      }'
```
`PUT /accounts/{account_id}/workers/workers/{worker_id}`

Perform a complete replacement of a Worker, where omitted properties are set to their default values. This is the exact same as the Create Worker endpoint, but operates on an existing Worker. To perform a partial update instead, use the Edit Worker endpoint.

#### (resource) workers.beta.workers > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/workers`

List all Workers for an account.

#### (resource) workers.beta.workers > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/workers/{worker_id}`

Delete a Worker and all its associated resources (versions, deployments, etc.).

#### (resource) workers.beta.workers > (method) edit
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID \
  -X PATCH \
  -H 'Content-Type: application/merge-patch+json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "logpush": true,
        "name": "my-worker",
        "observability": {},
        "subdomain": {},
        "tags": [
          "my-team",
          "my-public-api"
        ],
        "tail_consumers": [
          {
            "name": "my-tail-consumer"
          }
        ]
      }'
```
`PATCH /accounts/{account_id}/workers/workers/{worker_id}`

Perform a partial update on a Worker, where omitted properties are left unchanged from their current values.

#### (resource) workers.beta.workers > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/workers/{worker_id}`

Get details about a specific Worker.

---

## Workers.Beta.Workers.Versions

### Methods

#### (resource) workers.beta.workers.versions > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID/versions \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "bindings": [
          {
            "name": "MY_ENV_VAR",
            "text": "my_data",
            "type": "plain_text"
          }
        ],
        "compatibility_date": "2021-01-01",
        "compatibility_flags": [
          "nodejs_compat"
        ],
        "main_module": "index.js",
        "usage_model": "standard"
      }'
```
`POST /accounts/{account_id}/workers/workers/{worker_id}/versions`

Create a new version.

#### (resource) workers.beta.workers.versions > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID/versions \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/workers/{worker_id}/versions`

List all versions for a Worker.

#### (resource) workers.beta.workers.versions > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID/versions/$VERSION_ID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}`

Delete a version.

#### (resource) workers.beta.workers.versions > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/workers/$WORKER_ID/versions/$VERSION_ID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}`

Get details about a specific version.

---

## Workers.Routes

### Methods

#### (resource) workers.routes > (method) create
```bash
curl https://api.cloudflare.com/client/v4/zones/$ZONE_ID/workers/routes \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "pattern": "example.com/*",
        "script": "my-workers-script"
      }'
```
`POST /zones/{zone_id}/workers/routes`

Creates a route that maps a URL pattern to a Worker.

#### (resource) workers.routes > (method) update
```bash
curl https://api.cloudflare.com/client/v4/zones/$ZONE_ID/workers/routes/$ROUTE_ID \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "pattern": "example.com/*",
        "script": "my-workers-script"
      }'
```
`PUT /zones/{zone_id}/workers/routes/{route_id}`

Updates the URL pattern or Worker associated with a route.

#### (resource) workers.routes > (method) list
```bash
curl https://api.cloudflare.com/client/v4/zones/$ZONE_ID/workers/routes \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /zones/{zone_id}/workers/routes`

Returns routes for a zone.

#### (resource) workers.routes > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/zones/$ZONE_ID/workers/routes/$ROUTE_ID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /zones/{zone_id}/workers/routes/{route_id}`

Deletes a route.

#### (resource) workers.routes > (method) get
```bash
curl https://api.cloudflare.com/client/v4/zones/$ZONE_ID/workers/routes/$ROUTE_ID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /zones/{zone_id}/workers/routes/{route_id}`

Returns information about a route, including URL pattern and Worker.

---

## Workers.Assets.Upload

### Methods

#### (resource) workers.assets.upload > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/assets/upload \
  -H 'Content-Type: multipart/form-data' \
  -F body='{"foo":"string"}'
```
`POST /accounts/{account_id}/workers/assets/upload`

Upload assets ahead of creating a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.

---

## Workers.Scripts

### Methods

#### (resource) workers.scripts > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME \
  -X PUT \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -F metadata='{}'
```
`PUT /accounts/{account_id}/workers/scripts/{script_name}`

Upload a worker module. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.

#### (resource) workers.scripts > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts`

Fetch a list of uploaded workers.

#### (resource) workers.scripts > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/scripts/{script_name}`

Delete your worker. This call has no response body on a successful delete.

#### (resource) workers.scripts > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}`

Fetch raw script content for your worker. Note this is the original script content, not JSON encoded.

#### (resource) workers.scripts > (method) search
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts-search \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts-search`

Search for Workers in an account.

---

## Workers.Scripts.Assets.Upload

### Methods

#### (resource) workers.scripts.assets.upload > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/assets-upload-session \
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
`POST /accounts/{account_id}/workers/scripts/{script_name}/assets-upload-session`

Start uploading a collection of assets for use in a Worker version. To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.

---

## Workers.Scripts.Subdomain

### Methods

#### (resource) workers.scripts.subdomain > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/subdomain \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "enabled": true
      }'
```
`POST /accounts/{account_id}/workers/scripts/{script_name}/subdomain`

Enable or disable the Worker on the workers.dev subdomain.

#### (resource) workers.scripts.subdomain > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/subdomain \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/scripts/{script_name}/subdomain`

Disable all workers.dev subdomains for a Worker.

#### (resource) workers.scripts.subdomain > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/subdomain \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/subdomain`

Get if the Worker is available on the workers.dev subdomain.

---

## Workers.Scripts.Schedules

### Methods

#### (resource) workers.scripts.schedules > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/schedules \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '[
        {
          "cron": "*/30 * * * *"
        }
      ]'
```
`PUT /accounts/{account_id}/workers/scripts/{script_name}/schedules`

Updates Cron Triggers for a Worker.

#### (resource) workers.scripts.schedules > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/schedules \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/schedules`

Fetches Cron Triggers for a Worker.

---

## Workers.Scripts.Tail

### Methods

#### (resource) workers.scripts.tail > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/tails \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{}'
```
`POST /accounts/{account_id}/workers/scripts/{script_name}/tails`

Starts a tail that receives logs and exception from a Worker.

#### (resource) workers.scripts.tail > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/tails/$ID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/scripts/{script_name}/tails/{id}`

Deletes a tail from a Worker.

#### (resource) workers.scripts.tail > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/tails \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/tails`

Get list of tails currently deployed on a Worker.

---

## Workers.Scripts.Content

### Methods

#### (resource) workers.scripts.content > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/content \
  -X PUT \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -F metadata='{}'
```
`PUT /accounts/{account_id}/workers/scripts/{script_name}/content`

Put script content without touching config or metadata.

#### (resource) workers.scripts.content > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/content/v2 \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/content/v2`

Fetch script content only.

---

## Workers.Scripts.Settings

### Methods

#### (resource) workers.scripts.settings > (method) edit
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/script-settings \
  -X PATCH \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "tags": [
          "my-team",
          "my-public-api"
        ]
      }'
```
`PATCH /accounts/{account_id}/workers/scripts/{script_name}/script-settings`

Patch script-level settings when using Worker Versions. Including but not limited to Logpush and Tail Consumers.

#### (resource) workers.scripts.settings > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/script-settings \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/script-settings`

Get script-level settings when using Worker Versions. Includes Logpush and Tail Consumers.

---

## Workers.Scripts.Deployments

### Methods

#### (resource) workers.scripts.deployments > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/deployments \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "strategy": "percentage",
        "versions": [
          {
            "percentage": 100,
            "version_id": "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
          }
        ]
      }'
```
`POST /accounts/{account_id}/workers/scripts/{script_name}/deployments`

Deployments configure how Worker Versions are deployed to traffic. A deployment can consist of one or two versions of a Worker.

#### (resource) workers.scripts.deployments > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/deployments \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/deployments`

List of Worker Deployments. The first deployment in the list is the latest deployment actively serving traffic.

#### (resource) workers.scripts.deployments > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/deployments/$DEPLOYMENT_ID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}`

Delete a Worker Deployment. The latest deployment, which is actively serving traffic, cannot be deleted. All other deployments can be deleted.

#### (resource) workers.scripts.deployments > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/deployments/$DEPLOYMENT_ID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}`

Get information about a Worker Deployment.

---

## Workers.Scripts.Versions

### Methods

#### (resource) workers.scripts.versions > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/versions \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -F metadata='{"main_module":"worker.js"}'
```
`POST /accounts/{account_id}/workers/scripts/{script_name}/versions`

Upload a Worker Version without deploying to Cloudflare's network. You can find more about the multipart metadata on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.

#### (resource) workers.scripts.versions > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/versions \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/versions`

List of Worker Versions. The first version in the list is the latest version.

#### (resource) workers.scripts.versions > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/versions/$VERSION_ID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/versions/{version_id}`

Get Version Detail

---

## Workers.Scripts.Secrets

### Methods

#### (resource) workers.scripts.secrets > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/secrets \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "myBinding",
        "text": "My secret.",
        "type": "secret_text"
      }'
```
`PUT /accounts/{account_id}/workers/scripts/{script_name}/secrets`

Add a secret to a script.

#### (resource) workers.scripts.secrets > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/secrets \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/secrets`

List secrets bound to a script.

#### (resource) workers.scripts.secrets > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/secrets/$SECRET_NAME \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}`

Remove a secret from a script.

#### (resource) workers.scripts.secrets > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/secrets/$SECRET_NAME \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/secrets/{secret_name}`

Get a given secret binding (value omitted) on a script.

---

## Workers.Scripts.ScriptAndVersionSettings

### Methods

#### (resource) workers.scripts.script_and_version_settings > (method) edit
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/settings \
  -X PATCH \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`PATCH /accounts/{account_id}/workers/scripts/{script_name}/settings`

Patch metadata or config, such as bindings or usage model.

#### (resource) workers.scripts.script_and_version_settings > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/scripts/$SCRIPT_NAME/settings \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/scripts/{script_name}/settings`

Get metadata and config, such as bindings or usage model.

---

## Workers.AccountSettings

### Methods

#### (resource) workers.account_settings > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/account-settings \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{}'
```
`PUT /accounts/{account_id}/workers/account-settings`

Creates Worker account settings for an account.

#### (resource) workers.account_settings > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/account-settings \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/account-settings`

Fetches Worker account settings for an account.

---

## Workers.Domains

### Methods

#### (resource) workers.domains > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/domains \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "hostname": "foo.example.com",
        "service": "foo",
        "zone_id": "593c9c94de529bbbfaac7c53ced0447d",
        "environment": "production"
      }'
```
`PUT /accounts/{account_id}/workers/domains`

Attaches a Worker to a zone and hostname.

#### (resource) workers.domains > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/domains \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/domains`

Lists all Worker Domains for an account.

#### (resource) workers.domains > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/domains/$DOMAIN_ID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/domains/{domain_id}`

Detaches a Worker from a zone and hostname.

#### (resource) workers.domains > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/domains/$DOMAIN_ID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/domains/{domain_id}`

Gets a Worker domain.

---

## Workers.Subdomains

### Methods

#### (resource) workers.subdomains > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/subdomain \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "subdomain": "my-subdomain"
      }'
```
`PUT /accounts/{account_id}/workers/subdomain`

Creates a Workers subdomain for an account.

#### (resource) workers.subdomains > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/subdomain \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/workers/subdomain`

Deletes a Workers subdomain for an account.

#### (resource) workers.subdomains > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/subdomain \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/workers/subdomain`

Returns a Workers subdomain for an account.

---

## Workers.Observability.Telemetry

### Methods

#### (resource) workers.observability.telemetry > (method) keys
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/observability/telemetry/keys \
  -H 'Content-Type: application/json' \
  -H "X-Auth-Email: $CLOUDFLARE_EMAIL" \
  -H "X-Auth-Key: $CLOUDFLARE_API_KEY" \
  -d '{}'
```
`POST /accounts/{account_id}/workers/observability/telemetry/keys`

List all the keys in your telemetry events.

#### (resource) workers.observability.telemetry > (method) query
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/observability/telemetry/query \
  -H 'Content-Type: application/json' \
  -H "X-Auth-Email: $CLOUDFLARE_EMAIL" \
  -H "X-Auth-Key: $CLOUDFLARE_API_KEY" \
  -d '{
        "queryId": "queryId",
        "timeframe": {
          "from": 0,
          "to": 0
        }
      }'
```
`POST /accounts/{account_id}/workers/observability/telemetry/query`

Runs a temporary or saved query

#### (resource) workers.observability.telemetry > (method) values
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/workers/observability/telemetry/values \
  -H 'Content-Type: application/json' \
  -H "X-Auth-Email: $CLOUDFLARE_EMAIL" \
  -H "X-Auth-Key: $CLOUDFLARE_API_KEY" \
  -d '{
        "datasets": [
          "string"
        ],
        "key": "key",
        "timeframe": {
          "from": 0,
          "to": 0
        },
        "type": "string"
      }'
```
`POST /accounts/{account_id}/workers/observability/telemetry/values`

List unique values found in your events
