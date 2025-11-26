URL Normalization
URLNormalization
Methods
client.URLNormalization.Get(ctx, query) (*
URLNormalizationGetResponse
, error)
get/zones/{zone_id}/url_normalization
Fetches the current URL Normalization settings.
client.URLNormalization.Update(ctx, params) (*
URLNormalizationUpdateResponse
, error)
put/zones/{zone_id}/url_normalization
Updates the URL Normalization settings.
client.URLNormalization.Delete(ctx, body) error
delete/zones/{zone_id}/url_normalization
Deletes the URL Normalization settings.