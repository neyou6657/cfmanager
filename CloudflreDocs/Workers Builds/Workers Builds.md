# Workers Builds

## WorkersBuilds

### Methods

#### (resource) workers_builds > (method) get_account_limits
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/account/limits \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/account/limits`

Get account limits for Workers Builds.

#### (resource) workers_builds > (method) get_builds_by_version
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/builds \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/builds`

Get builds by version.

#### (resource) workers_builds > (method) get_latest_builds
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/builds/latest \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/builds/latest`

Get latest builds.

---

## WorkersBuilds.Triggers

### Methods

#### (resource) workers_builds.triggers > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "branch_excludes": [
          "string"
        ],
        "branch_includes": [
          "main"
        ],
        "build_command": "npm run build",
        "build_token_uuid": "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
        "deploy_command": "npx wrangler deploy",
        "external_script_id": "my-worker",
        "path_excludes": [
          "*.md"
        ],
        "path_includes": [
          "*"
        ],
        "repo_connection_uuid": "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
        "root_directory": "/",
        "trigger_name": "Production Deploy"
      }'
```
`POST /accounts/{account_id}/builds/triggers`

Create a build trigger.

#### (resource) workers_builds.triggers > (method) update
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers/$TRIGGER_UUID \
  -X PATCH \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "branch_excludes": [
          "string"
        ],
        "branch_includes": [
          "main"
        ],
        "build_command": "npm run build",
        "deploy_command": "npm run deploy",
        "path_excludes": [
          "*.md"
        ],
        "path_includes": [
          "src/**"
        ],
        "root_directory": "/",
        "trigger_name": "Production Deploy"
      }'
```
`PATCH /accounts/{account_id}/builds/triggers/{trigger_uuid}`

Update a build trigger.

#### (resource) workers_builds.triggers > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/workers/$EXTERNAL_SCRIPT_ID/triggers \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/workers/{external_script_id}/triggers`

List build triggers for a worker.

#### (resource) workers_builds.triggers > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers/$TRIGGER_UUID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/builds/triggers/{trigger_uuid}`

Delete a build trigger.

#### (resource) workers_builds.triggers > (method) create_build
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers/$TRIGGER_UUID/builds \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "branch": "main",
        "commit_hash": "abc123def456"
      }'
```
`POST /accounts/{account_id}/builds/triggers/{trigger_uuid}/builds`

Create a build for a trigger.

#### (resource) workers_builds.triggers > (method) purge_cache
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers/$TRIGGER_UUID/purge_build_cache \
  -X POST \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`POST /accounts/{account_id}/builds/triggers/{trigger_uuid}/purge_build_cache`

Purge build cache for a trigger.

---

## WorkersBuilds.Triggers.EnvironmentVariables

### Methods

#### (resource) workers_builds.triggers.environment_variables > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers/$TRIGGER_UUID/environment_variables \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/triggers/{trigger_uuid}/environment_variables`

List environment variables for a build trigger.

#### (resource) workers_builds.triggers.environment_variables > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers/$TRIGGER_UUID/environment_variables/$ENVIRONMENT_VARIABLE_KEY \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/builds/triggers/{trigger_uuid}/environment_variables/{environment_variable_key}`

Delete an environment variable for a build trigger.

#### (resource) workers_builds.triggers.environment_variables > (method) upsert
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/triggers/$TRIGGER_UUID/environment_variables \
  -X PATCH \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "API_KEY": {
          "is_secret": true,
          "value": "secret-key"
        },
        "NODE_ENV": {
          "is_secret": false,
          "value": "production"
        }
      }'
```
`PATCH /accounts/{account_id}/builds/triggers/{trigger_uuid}/environment_variables`

Upsert environment variables for a build trigger.

---

## WorkersBuilds.Tokens

### Methods

#### (resource) workers_builds.tokens > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/tokens \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "build_token_name": "My Build Token",
        "build_token_secret": "super-secret-token",
        "cloudflare_token_id": "cf-token-123"
      }'
```
`POST /accounts/{account_id}/builds/tokens`

Create a build token.

#### (resource) workers_builds.tokens > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/tokens \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/tokens`

List build tokens.

#### (resource) workers_builds.tokens > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/tokens/$BUILD_TOKEN_UUID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/builds/tokens/{build_token_uuid}`

Delete a build token.

---

## WorkersBuilds.Repos.Connections

### Methods

#### (resource) workers_builds.repos.connections > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/repos/connections/$REPO_CONNECTION_UUID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/builds/repos/connections/{repo_connection_uuid}`

Delete a repository connection.

#### (resource) workers_builds.repos.connections > (method) upsert
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/repos/connections \
  -X PUT \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "provider_account_id": "cloudflare",
        "provider_account_name": "Cloudflare",
        "provider_type": "github",
        "repo_id": "workers-sdk",
        "repo_name": "workers-sdk"
      }'
```
`PUT /accounts/{account_id}/builds/repos/connections`

Upsert a repository connection.

---

## WorkersBuilds.Repos.ConfigAutofill

### Methods

#### (resource) workers_builds.repos.config_autofill > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/repos/$PROVIDER_TYPE/$PROVIDER_ACCOUNT_ID/$REPO_ID/config_autofill \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/repos/{provider_type}/{provider_account_id}/{repo_id}/config_autofill`

Get configuration autofill for a repository.

---

## WorkersBuilds.Builds

### Methods

#### (resource) workers_builds.builds > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/workers/$EXTERNAL_SCRIPT_ID/builds \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/workers/{external_script_id}/builds`

List builds for a worker.

#### (resource) workers_builds.builds > (method) cancel
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/builds/$BUILD_UUID/cancel \
  -X PUT \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`PUT /accounts/{account_id}/builds/builds/{build_uuid}/cancel`

Cancel a build.

#### (resource) workers_builds.builds > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/builds/$BUILD_UUID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/builds/{build_uuid}`

Get a build.

---

## WorkersBuilds.Builds.Logs

### Methods

#### (resource) workers_builds.builds.logs > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/builds/builds/$BUILD_UUID/logs \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/builds/builds/{build_uuid}/logs`

Get build logs.
