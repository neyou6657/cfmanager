Hostnames
Hostnames
Hostnames
Settings
Hostnames.Settings
Hostnames
Settings
TLS
Hostnames.Settings.TLS
Methods
client.Hostnames.Settings.TLS.Get(ctx, settingID, query) (*SinglePage[
SettingTLSGetResponse
], error)
get/zones/{zone_id}/hostnames/settings/{setting_id}
List the requested TLS setting for the hostnames under this zone.
client.Hostnames.Settings.TLS.Update(ctx, settingID, hostname, params) (*
Setting
, error)
put/zones/{zone_id}/hostnames/settings/{setting_id}/{hostname}
Update the tls setting value for the hostname.
client.Hostnames.Settings.TLS.Delete(ctx, settingID, hostname, body) (*
SettingTLSDeleteResponse
, error)
delete/zones/{zone_id}/hostnames/settings/{setting_id}/{hostname}
Delete the tls setting value for the hostname.
Domain types
typeSettingstruct{…}
typeSettingValueUnioninterface{…}
The tls setting value.