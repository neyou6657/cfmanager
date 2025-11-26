Leaked Credential Checks
LeakedCredentialChecks
Methods
client.LeakedCredentialChecks.Get(ctx, query) (*
LeakedCredentialCheckGetResponse
, error)
get/zones/{zone_id}/leaked-credential-checks
Retrieves the current status of Leaked Credential Checks.
client.LeakedCredentialChecks.New(ctx, params) (*
LeakedCredentialCheckNewResponse
, error)
post/zones/{zone_id}/leaked-credential-checks
Updates the current status of Leaked Credential Checks.
Leaked Credential Checks
Detections
LeakedCredentialChecks.Detections
Methods
client.LeakedCredentialChecks.Detections.List(ctx, query) (*SinglePage[
DetectionListResponse
], error)
get/zones/{zone_id}/leaked-credential-checks/detections
List user-defined detection patterns for Leaked Credential Checks.
client.LeakedCredentialChecks.Detections.New(ctx, params) (*
DetectionNewResponse
, error)
post/zones/{zone_id}/leaked-credential-checks/detections
Create user-defined detection pattern for Leaked Credential Checks.
client.LeakedCredentialChecks.Detections.Update(ctx, detectionID, params) (*
DetectionUpdateResponse
, error)
put/zones/{zone_id}/leaked-credential-checks/detections/{detection_id}
Update user-defined detection pattern for Leaked Credential Checks.
client.LeakedCredentialChecks.Detections.Delete(ctx, detectionID, body) (*
DetectionDeleteResponse
, error)
delete/zones/{zone_id}/leaked-credential-checks/detections/{detection_id}
Remove user-defined detection pattern for Leaked Credential Checks.