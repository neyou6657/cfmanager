Botnet Feed
BotnetFeed
Botnet Feed
ASN
BotnetFeed.ASN
Methods
client.BotnetFeed.ASN.DayReport(ctx, asnID, params) (*
ASNDayReportResponse
, error)
get/accounts/{account_id}/botnet_feed/asn/{asn_id}/day_report
Gets all the data the botnet tracking database has for a given ASN registered to user account for given date. If no date is given, it will return results for the previous day.
client.BotnetFeed.ASN.FullReport(ctx, asnID, query) (*
ASNFullReportResponse
, error)
get/accounts/{account_id}/botnet_feed/asn/{asn_id}/full_report
Gets all the data the botnet threat feed tracking database has for a given ASN registered to user account.
Botnet Feed
Configs
BotnetFeed.Configs
Botnet Feed
Configs
ASN
BotnetFeed.Configs.ASN
Methods
client.BotnetFeed.Configs.ASN.Get(ctx, query) (*
ConfigASNGetResponse
, error)
get/accounts/{account_id}/botnet_feed/configs/asn
Gets a list of all ASNs registered for a user for the DDoS Botnet Feed API.
client.BotnetFeed.Configs.ASN.Delete(ctx, asnID, body) (*
ConfigASNDeleteResponse
, error)
delete/accounts/{account_id}/botnet_feed/configs/asn/{asn_id}
Delete an ASN from botnet threat feed for a given user.