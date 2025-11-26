Workflows
Workflows
Methods
client.Workflows.List(ctx, params) (*V4PagePaginationArray[
WorkflowListResponse
], error)
get/accounts/{account_id}/workflows
List all Workflows
client.Workflows.Get(ctx, workflowName, query) (*
WorkflowGetResponse
, error)
get/accounts/{account_id}/workflows/{workflow_name}
Get Workflow details
client.Workflows.Update(ctx, workflowName, params) (*
WorkflowUpdateResponse
, error)
put/accounts/{account_id}/workflows/{workflow_name}
Create/modify Workflow
client.Workflows.Delete(ctx, workflowName, body) (*
WorkflowDeleteResponse
, error)
delete/accounts/{account_id}/workflows/{workflow_name}
Deletes a Workflow. This only deletes the Workflow and does not delete or modify any Worker associated to this Workflow or bounded to it.
Workflows
Instances
Workflows.Instances
Methods
client.Workflows.Instances.List(ctx, workflowName, params) (*V4PagePaginationArray[
InstanceListResponse
], error)
get/accounts/{account_id}/workflows/{workflow_name}/instances
List of workflow instances
client.Workflows.Instances.Get(ctx, workflowName, instanceID, query) (*
InstanceGetResponse
, error)
get/accounts/{account_id}/workflows/{workflow_name}/instances/{instance_id}
Get logs and status from instance
client.Workflows.Instances.New(ctx, workflowName, params) (*
InstanceNewResponse
, error)
post/accounts/{account_id}/workflows/{workflow_name}/instances
Create a new workflow instance
client.Workflows.Instances.Bulk(ctx, workflowName, params) (*SinglePage[
InstanceBulkResponse
], error)
post/accounts/{account_id}/workflows/{workflow_name}/instances/batch
Batch create new Workflow instances
Workflows
Instances
Events
Workflows.Instances.Events
Methods
client.Workflows.Instances.Events.New(ctx, workflowName, instanceID, eventType, params) (*
InstanceEventNewResponse
, error)
post/accounts/{account_id}/workflows/{workflow_name}/instances/{instance_id}/events/{event_type}
Send event to instance
Workflows
Instances
Status
Workflows.Instances.Status
Methods
client.Workflows.Instances.Status.Edit(ctx, workflowName, instanceID, params) (*
InstanceStatusEditResponse
, error)
patch/accounts/{account_id}/workflows/{workflow_name}/instances/{instance_id}/status
Change status of instance
Workflows
Versions
Workflows.Versions
Methods
client.Workflows.Versions.List(ctx, workflowName, params) (*V4PagePaginationArray[
VersionListResponse
], error)
get/accounts/{account_id}/workflows/{workflow_name}/versions
List deployed Workflow versions
client.Workflows.Versions.Get(ctx, workflowName, versionID, query) (*
VersionGetResponse
, error)
get/accounts/{account_id}/workflows/{workflow_name}/versions/{version_id}
Get Workflow version details