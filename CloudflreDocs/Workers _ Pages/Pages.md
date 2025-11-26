Pages
Pages
Pages
Projects
Pages.Projects
Methods
client.Pages.Projects.List(ctx, params) (*V4PagePaginationArray[
ProjectListResponse
], error)
get/accounts/{account_id}/pages/projects
Fetch a list of all user projects.
client.Pages.Projects.Get(ctx, projectName, query) (*
ProjectGetResponse
, error)
get/accounts/{account_id}/pages/projects/{project_name}
Fetch a project by name.
client.Pages.Projects.New(ctx, params) (*
ProjectNewResponse
, error)
post/accounts/{account_id}/pages/projects
Create a new project.
client.Pages.Projects.Edit(ctx, projectName, params) (*
ProjectEditResponse
, error)
patch/accounts/{account_id}/pages/projects/{project_name}
Set new attributes for an existing project. Modify environment variables. To delete an environment variable, set the key to null.
client.Pages.Projects.Delete(ctx, projectName, body) (*
ProjectDeleteResponse
, error)
delete/accounts/{account_id}/pages/projects/{project_name}
Delete a project by name.
client.Pages.Projects.PurgeBuildCache(ctx, projectName, body) (*
ProjectPurgeBuildCacheResponse
, error)
post/accounts/{account_id}/pages/projects/{project_name}/purge_build_cache
Purge all cached build artifacts for a Pages project
Domain types
typeDeploymentstruct{…}
typeProjectstruct{…}
typeStagestruct{…}
The status of the deployment.
Pages
Projects
Deployments
Pages.Projects.Deployments
Methods
client.Pages.Projects.Deployments.List(ctx, projectName, params) (*V4PagePaginationArray[
ProjectDeploymentListResponse
], error)
get/accounts/{account_id}/pages/projects/{project_name}/deployments
Fetch a list of project deployments.
client.Pages.Projects.Deployments.Get(ctx, projectName, deploymentID, query) (*
ProjectDeploymentGetResponse
, error)
get/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}
Fetch information about a deployment.
client.Pages.Projects.Deployments.New(ctx, projectName, params) (*
ProjectDeploymentNewResponse
, error)
post/accounts/{account_id}/pages/projects/{project_name}/deployments
Start a new deployment from production. The repository and account must have already been authorized on the Cloudflare Pages dashboard.
client.Pages.Projects.Deployments.Delete(ctx, projectName, deploymentID, body) (*
ProjectDeploymentDeleteResponse
, error)
delete/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}
Delete a deployment.
client.Pages.Projects.Deployments.Retry(ctx, projectName, deploymentID, body) (*
ProjectDeploymentRetryResponse
, error)
post/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry
Retry a previous deployment.
client.Pages.Projects.Deployments.Rollback(ctx, projectName, deploymentID, body) (*
ProjectDeploymentRollbackResponse
, error)
post/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback
Rollback the production deployment to a previous deployment. You can only rollback to succesful builds on production.
Pages
Projects
Deployments
History
Pages.Projects.Deployments.History
Pages
Projects
Deployments
History
Logs
Pages.Projects.Deployments.History.Logs
Methods
client.Pages.Projects.Deployments.History.Logs.Get(ctx, projectName, deploymentID, query) (*
ProjectDeploymentHistoryLogGetResponse
, error)
get/accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/history/logs
Fetch deployment logs for a project.
Pages
Projects
Domains
Pages.Projects.Domains
Methods
client.Pages.Projects.Domains.List(ctx, projectName, query) (*SinglePage[
ProjectDomainListResponse
], error)
get/accounts/{account_id}/pages/projects/{project_name}/domains
Fetch a list of all domains associated with a Pages project.
client.Pages.Projects.Domains.Get(ctx, projectName, domainName, query) (*
ProjectDomainGetResponse
, error)
get/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}
Fetch a single domain.
client.Pages.Projects.Domains.New(ctx, projectName, params) (*
ProjectDomainNewResponse
, error)
post/accounts/{account_id}/pages/projects/{project_name}/domains
Add a new domain for the Pages project.
client.Pages.Projects.Domains.Edit(ctx, projectName, domainName, body) (*
ProjectDomainEditResponse
, error)
patch/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}
Retry the validation status of a single domain.
client.Pages.Projects.Domains.Delete(ctx, projectName, domainName, body) (*
ProjectDomainDeleteResponse
, error)
delete/accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}
Delete a Pages project's domain.