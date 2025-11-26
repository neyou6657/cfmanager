Network Interconnects
NetworkInterconnects
Network Interconnects
CNIs
NetworkInterconnects.CNIs
Methods
client.NetworkInterconnects.CNIs.List(ctx, params) (*
CNIListResponse
, error)
get/accounts/{account_id}/cni/cnis
List existing CNI objects
client.NetworkInterconnects.CNIs.Get(ctx, cni, query) (*
CNIGetResponse
, error)
get/accounts/{account_id}/cni/cnis/{cni}
Get information about a CNI object
client.NetworkInterconnects.CNIs.New(ctx, params) (*
CNINewResponse
, error)
post/accounts/{account_id}/cni/cnis
Create a new CNI object
client.NetworkInterconnects.CNIs.Update(ctx, cni, params) (*
CNIUpdateResponse
, error)
put/accounts/{account_id}/cni/cnis/{cni}
Modify stored information about a CNI object
client.NetworkInterconnects.CNIs.Delete(ctx, cni, body) error
delete/accounts/{account_id}/cni/cnis/{cni}
Delete a specified CNI object
Network Interconnects
Interconnects
NetworkInterconnects.Interconnects
Methods
client.NetworkInterconnects.Interconnects.List(ctx, params) (*
InterconnectListResponse
, error)
get/accounts/{account_id}/cni/interconnects
List existing interconnects
client.NetworkInterconnects.Interconnects.Get(ctx, icon, query) (*
InterconnectGetResponse
, error)
get/accounts/{account_id}/cni/interconnects/{icon}
Get information about an interconnect object
client.NetworkInterconnects.Interconnects.New(ctx, params) (*
InterconnectNewResponse
, error)
post/accounts/{account_id}/cni/interconnects
Create a new interconnect
client.NetworkInterconnects.Interconnects.Delete(ctx, icon, body) error
delete/accounts/{account_id}/cni/interconnects/{icon}
Delete an interconnect object
client.NetworkInterconnects.Interconnects.LOA(ctx, icon, query) error
get/accounts/{account_id}/cni/interconnects/{icon}/loa
Generate the Letter of Authorization (LOA) for a given interconnect
client.NetworkInterconnects.Interconnects.Status(ctx, icon, query) (*
InterconnectStatusResponse
, error)
get/accounts/{account_id}/cni/interconnects/{icon}/status
Get the current status of an interconnect object
Network Interconnects
Settings
NetworkInterconnects.Settings
Methods
client.NetworkInterconnects.Settings.Get(ctx, query) (*
SettingGetResponse
, error)
get/accounts/{account_id}/cni/settings
Get the current settings for the active account
client.NetworkInterconnects.Settings.Update(ctx, params) (*
SettingUpdateResponse
, error)
put/accounts/{account_id}/cni/settings
Update the current settings for the active account
Network Interconnects
Slots
NetworkInterconnects.Slots
Methods
client.NetworkInterconnects.Slots.List(ctx, params) (*
SlotListResponse
, error)
get/accounts/{account_id}/cni/slots
Retrieve a list of all slots matching the specified parameters
client.NetworkInterconnects.Slots.Get(ctx, slot, query) (*
SlotGetResponse
, error)
get/accounts/{account_id}/cni/slots/{slot}
Get information about the specified slot