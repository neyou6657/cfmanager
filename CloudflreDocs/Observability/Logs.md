Logs
Logs
Logs
Control
Logs.Control
Logs
Control
Cmb
Logs.Control.Cmb
Logs
Control
Cmb
Config
Logs.Control.Cmb.Config
Methods
client.Logs.Control.Cmb.Config.Get(ctx, query) (*
CmbConfig
, error)
get/accounts/{account_id}/logs/control/cmb/config
Gets CMB config.
client.Logs.Control.Cmb.Config.New(ctx, params) (*
CmbConfig
, error)
post/accounts/{account_id}/logs/control/cmb/config
Updates CMB config.
client.Logs.Control.Cmb.Config.Delete(ctx, body) (*
ControlCmbConfigDeleteResponse
, error)
delete/accounts/{account_id}/logs/control/cmb/config
Deletes CMB config.
Domain types
typeCmbConfigstruct{…}
Logs
Control
Retention
Logs.Control.Retention
Methods
client.Logs.Control.Retention.Get(ctx, query) (*
ControlRetentionGetResponse
, error)
get/zones/{zone_id}/logs/control/retention/flag
Gets log retention flag for Logpull API.
client.Logs.Control.Retention.New(ctx, params) (*
ControlRetentionNewResponse
, error)
post/zones/{zone_id}/logs/control/retention/flag
Updates log retention flag for Logpull API.
Logs
RayID
Logs.RayID
Methods
client.Logs.RayID.Get(ctx, RayID, params) (*unknown, error)
get/zones/{zone_id}/logs/rayids/{ray_id}
The /rayids api route allows lookups by specific rayid. The rayids route will return zero, one, or more records (ray ids are not unique).
Logs
Received
Logs.Received
Methods
client.Logs.Received.Get(ctx, params) (*unknown, error)
get/zones/{zone_id}/logs/received
The /received api route allows customers to retrieve their edge HTTP logs. The basic access pattern is "give me all the logs for zone Z for minute M", where the minute M refers to the time records were received at Cloudflare's central data center. start is inclusive, and end is exclusive. Because of that, to get all data, at minutely cadence, starting at 10AM, the proper values are: start=2018-05-20T10:00:00Z&end=2018-05-20T10:01:00Z, then start=2018-05-20T10:01:00Z&end=2018-05-20T10:02:00Z and so on; the overlap will be handled properly.
Logs
Received
Fields
Logs.Received.Fields
Methods
client.Logs.Received.Fields.Get(ctx, query) (*
ReceivedFieldGetResponse
, error)
get/zones/{zone_id}/logs/received/fields
Lists all fields available. The response is json object with key-value pairs, where keys are field names, and values are descriptions.