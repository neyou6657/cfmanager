Hyperdrive
Hyperdrive
Domain types
typeConfigurationinterface{…}
typeHyperdrivestruct{…}
Hyperdrive
Configs
Hyperdrive.Configs
Methods
client.Hyperdrive.Configs.List(ctx, query) (*SinglePage[
Hyperdrive
], error)
get/accounts/{account_id}/hyperdrive/configs
Returns a list of Hyperdrives.
client.Hyperdrive.Configs.Get(ctx, hyperdriveID, query) (*
Hyperdrive
, error)
get/accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
Returns the specified Hyperdrive configuration.
client.Hyperdrive.Configs.New(ctx, params) (*
Hyperdrive
, error)
post/accounts/{account_id}/hyperdrive/configs
Creates and returns a new Hyperdrive configuration.
client.Hyperdrive.Configs.Update(ctx, hyperdriveID, params) (*
Hyperdrive
, error)
put/accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
Updates and returns the specified Hyperdrive configuration.
client.Hyperdrive.Configs.Edit(ctx, hyperdriveID, params) (*
Hyperdrive
, error)
patch/accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
Patches and returns the specified Hyperdrive configuration. Custom caching settings are not kept if caching is disabled.
client.Hyperdrive.Configs.Delete(ctx, hyperdriveID, body) (*
ConfigDeleteResponse
, error)
delete/accounts/{account_id}/hyperdrive/configs/{hyperdrive_id}
Deletes the specified Hyperdrive.