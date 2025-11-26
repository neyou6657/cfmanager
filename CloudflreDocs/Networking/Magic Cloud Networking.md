Magic Cloud Networking
MagicCloudNetworking
Magic Cloud Networking
Catalog Syncs
MagicCloudNetworking.CatalogSyncs
Methods
client.MagicCloudNetworking.CatalogSyncs.List(ctx, query) (*SinglePage[
CatalogSyncListResponse
], error)
get/accounts/{account_id}/magic/cloud/catalog-syncs
List Catalog Syncs (Closed Beta).
client.MagicCloudNetworking.CatalogSyncs.Get(ctx, syncID, query) (*
CatalogSyncGetResponse
, error)
get/accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}
Read a Catalog Sync (Closed Beta).
client.MagicCloudNetworking.CatalogSyncs.New(ctx, params) (*
CatalogSyncNewResponse
, error)
post/accounts/{account_id}/magic/cloud/catalog-syncs
Create a new Catalog Sync (Closed Beta).
client.MagicCloudNetworking.CatalogSyncs.Update(ctx, syncID, params) (*
CatalogSyncUpdateResponse
, error)
put/accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}
Update a Catalog Sync (Closed Beta).
client.MagicCloudNetworking.CatalogSyncs.Edit(ctx, syncID, params) (*
CatalogSyncEditResponse
, error)
patch/accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}
Update a Catalog Sync (Closed Beta).
client.MagicCloudNetworking.CatalogSyncs.Delete(ctx, syncID, params) (*
CatalogSyncDeleteResponse
, error)
delete/accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}
Delete a Catalog Sync (Closed Beta).
client.MagicCloudNetworking.CatalogSyncs.Refresh(ctx, syncID, body) (*string, error)
post/accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}/refresh
Refresh a Catalog Sync's destination by running the sync policy against latest resource catalog (Closed Beta).
Magic Cloud Networking
Catalog Syncs
Prebuilt Policies
MagicCloudNetworking.CatalogSyncs.PrebuiltPolicies
Methods
client.MagicCloudNetworking.CatalogSyncs.PrebuiltPolicies.List(ctx, params) (*SinglePage[
CatalogSyncPrebuiltPolicyListResponse
], error)
get/accounts/{account_id}/magic/cloud/catalog-syncs/prebuilt-policies
List prebuilt catalog sync policies (Closed Beta).
Magic Cloud Networking
Cloud Integrations
MagicCloudNetworking.CloudIntegrations
Methods
client.MagicCloudNetworking.CloudIntegrations.List(ctx, params) (*SinglePage[
CloudIntegrationListResponse
], error)
get/accounts/{account_id}/magic/cloud/providers
List Cloud Integrations (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.Get(ctx, providerID, params) (*
CloudIntegrationGetResponse
, error)
get/accounts/{account_id}/magic/cloud/providers/{provider_id}
Read a Cloud Integration (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.New(ctx, params) (*
CloudIntegrationNewResponse
, error)
post/accounts/{account_id}/magic/cloud/providers
Create a new Cloud Integration (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.Update(ctx, providerID, params) (*
CloudIntegrationUpdateResponse
, error)
put/accounts/{account_id}/magic/cloud/providers/{provider_id}
Update a Cloud Integration (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.Edit(ctx, providerID, params) (*
CloudIntegrationEditResponse
, error)
patch/accounts/{account_id}/magic/cloud/providers/{provider_id}
Update a Cloud Integration (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.Delete(ctx, providerID, body) (*
CloudIntegrationDeleteResponse
, error)
delete/accounts/{account_id}/magic/cloud/providers/{provider_id}
Delete a Cloud Integration (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.DiscoverAll(ctx, body) (*
CloudIntegrationDiscoverAllResponse
, error)
post/accounts/{account_id}/magic/cloud/providers/discover
Run discovery for all Cloud Integrations in an account (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.Discover(ctx, providerID, params) (*
CloudIntegrationDiscoverResponse
, error)
post/accounts/{account_id}/magic/cloud/providers/{provider_id}/discover
Run discovery for a Cloud Integration (Closed Beta).
client.MagicCloudNetworking.CloudIntegrations.InitialSetup(ctx, providerID, query) (*
CloudIntegrationInitialSetupResponse
, error)
get/accounts/{account_id}/magic/cloud/providers/{provider_id}/initial_setup
Get initial configuration to complete Cloud Integration setup (Closed Beta).
Magic Cloud Networking
On Ramps
MagicCloudNetworking.OnRamps
Methods
client.MagicCloudNetworking.OnRamps.List(ctx, params) (*SinglePage[
OnRampListResponse
], error)
get/accounts/{account_id}/magic/cloud/onramps
List On-ramps (Closed Beta).
client.MagicCloudNetworking.OnRamps.Get(ctx, onrampID, params) (*
OnRampGetResponse
, error)
get/accounts/{account_id}/magic/cloud/onramps/{onramp_id}
Read an On-ramp (Closed Beta).
client.MagicCloudNetworking.OnRamps.New(ctx, params) (*
OnRampNewResponse
, error)
post/accounts/{account_id}/magic/cloud/onramps
Create a new On-ramp (Closed Beta).
client.MagicCloudNetworking.OnRamps.Update(ctx, onrampID, params) (*
OnRampUpdateResponse
, error)
put/accounts/{account_id}/magic/cloud/onramps/{onramp_id}
Update an On-ramp (Closed Beta).
client.MagicCloudNetworking.OnRamps.Edit(ctx, onrampID, params) (*
OnRampEditResponse
, error)
patch/accounts/{account_id}/magic/cloud/onramps/{onramp_id}
Update an On-ramp (Closed Beta).
client.MagicCloudNetworking.OnRamps.Delete(ctx, onrampID, params) (*
OnRampDeleteResponse
, error)
delete/accounts/{account_id}/magic/cloud/onramps/{onramp_id}
Delete an On-ramp (Closed Beta).
client.MagicCloudNetworking.OnRamps.Apply(ctx, onrampID, body) (*
OnRampApplyResponse
, error)
post/accounts/{account_id}/magic/cloud/onramps/{onramp_id}/apply
Apply an On-ramp (Closed Beta).
client.MagicCloudNetworking.OnRamps.Export(ctx, onrampID, body) (*
Response
, error)
post/accounts/{account_id}/magic/cloud/onramps/{onramp_id}/export
Export an On-ramp to terraform ready file(s) (Closed Beta).
client.MagicCloudNetworking.OnRamps.Plan(ctx, onrampID, body) (*
OnRampPlanResponse
, error)
post/accounts/{account_id}/magic/cloud/onramps/{onramp_id}/plan
Plan an On-ramp (Closed Beta).
Magic Cloud Networking
On Ramps
Address Spaces
MagicCloudNetworking.OnRamps.AddressSpaces
Methods
client.MagicCloudNetworking.OnRamps.AddressSpaces.List(ctx, query) (*
OnRampAddressSpaceListResponse
, error)
get/accounts/{account_id}/magic/cloud/onramps/magic_wan_address_space
Read the Magic WAN Address Space (Closed Beta).
client.MagicCloudNetworking.OnRamps.AddressSpaces.Update(ctx, params) (*
OnRampAddressSpaceUpdateResponse
, error)
put/accounts/{account_id}/magic/cloud/onramps/magic_wan_address_space
Update the Magic WAN Address Space (Closed Beta).
client.MagicCloudNetworking.OnRamps.AddressSpaces.Edit(ctx, params) (*
OnRampAddressSpaceEditResponse
, error)
patch/accounts/{account_id}/magic/cloud/onramps/magic_wan_address_space
Update the Magic WAN Address Space (Closed Beta).
Magic Cloud Networking
Resources
MagicCloudNetworking.Resources
Methods
client.MagicCloudNetworking.Resources.List(ctx, params) (*V4PagePaginationArray[
ResourceListResponse
], error)
get/accounts/{account_id}/magic/cloud/resources
List resources in the Resource Catalog (Closed Beta).
client.MagicCloudNetworking.Resources.Get(ctx, resourceID, params) (*
ResourceGetResponse
, error)
get/accounts/{account_id}/magic/cloud/resources/{resource_id}
Read an resource from the Resource Catalog (Closed Beta).
client.MagicCloudNetworking.Resources.Export(ctx, params) (*
Response
, error)
get/accounts/{account_id}/magic/cloud/resources/export
Export resources in the Resource Catalog as a JSON file (Closed Beta).
client.MagicCloudNetworking.Resources.PolicyPreview(ctx, params) (*string, error)
post/accounts/{account_id}/magic/cloud/resources/policy-preview
Preview Rego query result against the latest resource catalog (Closed Beta).