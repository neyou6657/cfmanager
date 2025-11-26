Calls
Calls
Calls
SFU
Calls.SFU
Methods
client.Calls.SFU.List(ctx, query) (*SinglePage[
SFUListResponse
], error)
get/accounts/{account_id}/calls/apps
Lists all apps in the Cloudflare account
client.Calls.SFU.Get(ctx, appID, query) (*
SFUGetResponse
, error)
get/accounts/{account_id}/calls/apps/{app_id}
Fetches details for a single Calls app.
client.Calls.SFU.New(ctx, params) (*
SFUNewResponse
, error)
post/accounts/{account_id}/calls/apps
Creates a new Cloudflare calls app. An app is an unique enviroment where each Session can access all Tracks within the app.
client.Calls.SFU.Update(ctx, appID, params) (*
SFUUpdateResponse
, error)
put/accounts/{account_id}/calls/apps/{app_id}
Edit details for a single app.
client.Calls.SFU.Delete(ctx, appID, body) (*
SFUDeleteResponse
, error)
delete/accounts/{account_id}/calls/apps/{app_id}
Deletes an app from Cloudflare Calls
Calls
TURN
Calls.TURN
Methods
client.Calls.TURN.List(ctx, query) (*SinglePage[
TURNListResponse
], error)
get/accounts/{account_id}/calls/turn_keys
Lists all TURN keys in the Cloudflare account
client.Calls.TURN.Get(ctx, keyID, query) (*
TURNGetResponse
, error)
get/accounts/{account_id}/calls/turn_keys/{key_id}
Fetches details for a single TURN key.
client.Calls.TURN.New(ctx, params) (*
TURNNewResponse
, error)
post/accounts/{account_id}/calls/turn_keys
Creates a new Cloudflare Calls TURN key.
client.Calls.TURN.Update(ctx, keyID, params) (*
TURNUpdateResponse
, error)
put/accounts/{account_id}/calls/turn_keys/{key_id}
Edit details for a single TURN key.
client.Calls.TURN.Delete(ctx, keyID, body) (*
TURNDeleteResponse
, error)
delete/accounts/{account_id}/calls/turn_keys/{key_id}
Deletes a TURN key from Cloudflare Calls