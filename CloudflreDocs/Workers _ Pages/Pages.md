# Pages

## Pages.Projects

### Methods

#### (resource) pages.projects > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "my-pages-app",
        "production_branch": "main"
      }'
```
`POST /accounts/{account_id}/pages/projects`

Create a new project.

#### (resource) pages.projects > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/pages/projects`

Fetch a list of all user projects.

#### (resource) pages.projects > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/pages/projects/{project_name}`

Delete a project by name.

#### (resource) pages.projects > (method) edit
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME \
  -X PATCH \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "my-pages-app",
        "production_branch": "main"
      }'
```
`PATCH /accounts/{account_id}/pages/projects/{project_name}`

Set new attributes for an existing project. Modify environment variables. To delete an environment variable, set the key to null.

#### (resource) pages.projects > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/pages/projects/{project_name}`

Fetch a project by name.

#### (resource) pages.projects > (method) purge_build_cache
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/purge_build_cache \
  -X POST \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`POST /accounts/{account_id}/pages/projects/{project_name}/purge_build_cache`

Purge all cached build artifacts for a Pages project

---

## Pages.Projects.Deployments

### Methods

#### (resource) pages.projects.deployments > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/deployments \
  -H 'Content-Type: multipart/form-data' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -F branch=staging \
  -F commit_dirty=false \
  -F commit_hash=a1b2c3d4e5f6 \
  -F commit_message='Update homepage' \
  -F manifest='{"index.html": "abc123", "style.css": "def456"}' \
  -F pages_build_output_dir=dist
```
`POST /accounts/{account_id}/pages/projects/{project_name}/deployments`

Start a new deployment from production. The repository and account must have already been authorized on the Cloudflare Pages dashboard.

#### (resource) pages.projects.deployments > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/deployments \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/pages/projects/{project_name}/deployments`

Fetch a list of project deployments.

#### (resource) pages.projects.deployments > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/deployments/$DEPLOYMENT_ID \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}`

Delete a deployment.

#### (resource) pages.projects.deployments > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/deployments/$DEPLOYMENT_ID \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}`

Fetch information about a deployment.

#### (resource) pages.projects.deployments > (method) retry
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/deployments/$DEPLOYMENT_ID/retry \
  -X POST \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry`

Retry a previous deployment.

#### (resource) pages.projects.deployments > (method) rollback
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/deployments/$DEPLOYMENT_ID/rollback \
  -X POST \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback`

Rollback the production deployment to a previous deployment. You can only rollback to succesful builds on production.

---

## Pages.Projects.Deployments.History.Logs

### Methods

#### (resource) pages.projects.deployments.history.logs > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/deployments/$DEPLOYMENT_ID/history/logs \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/history/logs`

Fetch deployment logs for a project.

---

## Pages.Projects.Domains

### Methods

#### (resource) pages.projects.domains > (method) create
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/domains \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
  -d '{
        "name": "this-is-my-domain-01.com"
      }'
```
`POST /accounts/{account_id}/pages/projects/{project_name}/domains`

Add a new domain for the Pages project.

#### (resource) pages.projects.domains > (method) list
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/domains \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/pages/projects/{project_name}/domains`

Fetch a list of all domains associated with a Pages project.

#### (resource) pages.projects.domains > (method) delete
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/domains/$DOMAIN_NAME \
  -X DELETE \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`DELETE /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}`

Delete a Pages project's domain.

#### (resource) pages.projects.domains > (method) edit
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/domains/$DOMAIN_NAME \
  -X PATCH \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`PATCH /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}`

Retry the validation status of a single domain.

#### (resource) pages.projects.domains > (method) get
```bash
curl https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/pages/projects/$PROJECT_NAME/domains/$DOMAIN_NAME \
  -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
```
`GET /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}`

Fetch a single domain.
