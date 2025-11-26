Alerting
Alerting
Alerting
Available Alerts
Alerting.AvailableAlerts
Methods
client.Alerting.AvailableAlerts.List(ctx, query) (*
AvailableAlertListResponse
, error)
get/accounts/{account_id}/alerting/v3/available_alerts
Gets a list of all alert types for which an account is eligible.
Alerting
Destinations
Alerting.Destinations
Alerting
Destinations
Eligible
Alerting.Destinations.Eligible
Methods
client.Alerting.Destinations.Eligible.Get(ctx, query) (*
DestinationEligibleGetResponse
, error)
get/accounts/{account_id}/alerting/v3/destinations/eligible
Get a list of all delivery mechanism types for which an account is eligible.
Alerting
Destinations
Pagerduty
Alerting.Destinations.Pagerduty
Methods
client.Alerting.Destinations.Pagerduty.Get(ctx, query) (*SinglePage[
Pagerduty
], error)
get/accounts/{account_id}/alerting/v3/destinations/pagerduty
Get a list of all configured PagerDuty services.
client.Alerting.Destinations.Pagerduty.New(ctx, body) (*
DestinationPagerdutyNewResponse
, error)
post/accounts/{account_id}/alerting/v3/destinations/pagerduty/connect
Creates a new token for integrating with PagerDuty.
client.Alerting.Destinations.Pagerduty.Delete(ctx, body) (*
DestinationPagerdutyDeleteResponse
, error)
delete/accounts/{account_id}/alerting/v3/destinations/pagerduty
Deletes all the PagerDuty Services connected to the account.
client.Alerting.Destinations.Pagerduty.Link(ctx, tokenID, query) (*
DestinationPagerdutyLinkResponse
, error)
get/accounts/{account_id}/alerting/v3/destinations/pagerduty/connect/{token_id}
Links PagerDuty with the account using the integration token.
Domain types
typePagerdutystruct{…}
Alerting
Destinations
Webhooks
Alerting.Destinations.Webhooks
Methods
client.Alerting.Destinations.Webhooks.List(ctx, query) (*SinglePage[
Webhooks
], error)
get/accounts/{account_id}/alerting/v3/destinations/webhooks
Gets a list of all configured webhook destinations.
client.Alerting.Destinations.Webhooks.Get(ctx, webhookID, query) (*
Webhooks
, error)
get/accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}
Get details for a single webhooks destination.
client.Alerting.Destinations.Webhooks.New(ctx, params) (*
DestinationWebhookNewResponse
, error)
post/accounts/{account_id}/alerting/v3/destinations/webhooks
Creates a new webhook destination.
client.Alerting.Destinations.Webhooks.Update(ctx, webhookID, params) (*
DestinationWebhookUpdateResponse
, error)
put/accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}
Update a webhook destination.
client.Alerting.Destinations.Webhooks.Delete(ctx, webhookID, body) (*
DestinationWebhookDeleteResponse
, error)
delete/accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}
Delete a configured webhook destination.
Domain types
typeWebhooksstruct{…}
Alerting
History
Alerting.History
Methods
client.Alerting.History.List(ctx, params) (*V4PagePaginationArray[
History
], error)
get/accounts/{account_id}/alerting/v3/history
Gets a list of history records for notifications sent to an account. The records are displayed for last x number of days based on the zone plan (free = 30, pro = 30, biz = 30, ent = 90).
Domain types
typeHistorystruct{…}
Alerting
Policies
Alerting.Policies
Methods
client.Alerting.Policies.List(ctx, query) (*SinglePage[
Policy
], error)
get/accounts/{account_id}/alerting/v3/policies
Get a list of all Notification policies.
client.Alerting.Policies.Get(ctx, policyID, query) (*
Policy
, error)
get/accounts/{account_id}/alerting/v3/policies/{policy_id}
Get details for a single policy.
client.Alerting.Policies.New(ctx, params) (*
PolicyNewResponse
, error)
post/accounts/{account_id}/alerting/v3/policies
Creates a new Notification policy.
client.Alerting.Policies.Update(ctx, policyID, params) (*
PolicyUpdateResponse
, error)
put/accounts/{account_id}/alerting/v3/policies/{policy_id}
Update a Notification policy.
client.Alerting.Policies.Delete(ctx, policyID, body) (*
PolicyDeleteResponse
, error)
delete/accounts/{account_id}/alerting/v3/policies/{policy_id}
Delete a Notification policy.
Domain types
typeMechanismstruct{…}
List of IDs that will be used when dispatching a notification. IDs for email type will be the email address.
typePolicystruct{…}
typePolicyFilterstruct{…}
Optional filters that allow you to be alerted only on a subset of events for that alert type based on some criteria. This is only available for select alert types. See alert type documentation for more details.
Alerting
Silences
Alerting.Silences
Methods
client.Alerting.Silences.List(ctx, query) (*SinglePage[
SilenceListResponse
], error)
get/accounts/{account_id}/alerting/v3/silences
Gets a list of silences for an account.
client.Alerting.Silences.Get(ctx, silenceID, query) (*
SilenceGetResponse
, error)
get/accounts/{account_id}/alerting/v3/silences/{silence_id}
Gets a specific silence for an account.
client.Alerting.Silences.New(ctx, params) (*
SilenceNewResponse
, error)
post/accounts/{account_id}/alerting/v3/silences
Creates a new silence for an account.
client.Alerting.Silences.Update(ctx, params) (*SinglePage[
SilenceUpdateResponse
], error)
put/accounts/{account_id}/alerting/v3/silences
Updates existing silences for an account.
client.Alerting.Silences.Delete(ctx, silenceID, body) (*
SilenceDeleteResponse
, error)
delete/accounts/{account_id}/alerting/v3/silences/{silence_id}
Deletes an existing silence for an account.