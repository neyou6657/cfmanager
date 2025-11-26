Turnstile
Turnstile
Turnstile
Widgets
Turnstile.Widgets
Methods
client.Turnstile.Widgets.List(ctx, params) (*V4PagePaginationArray[
WidgetListResponse
], error)
get/accounts/{account_id}/challenges/widgets
Lists all turnstile widgets of an account.
client.Turnstile.Widgets.Get(ctx, sitekey, query) (*
Widget
, error)
get/accounts/{account_id}/challenges/widgets/{sitekey}
Show a single challenge widget configuration.
client.Turnstile.Widgets.New(ctx, params) (*
Widget
, error)
post/accounts/{account_id}/challenges/widgets
Lists challenge widgets.
client.Turnstile.Widgets.Update(ctx, sitekey, params) (*
Widget
, error)
put/accounts/{account_id}/challenges/widgets/{sitekey}
Update the configuration of a widget.
client.Turnstile.Widgets.Delete(ctx, sitekey, body) (*
Widget
, error)
delete/accounts/{account_id}/challenges/widgets/{sitekey}
Destroy a Turnstile Widget.
client.Turnstile.Widgets.RotateSecret(ctx, sitekey, params) (*
Widget
, error)
post/accounts/{account_id}/challenges/widgets/{sitekey}/rotate_secret
Generate a new secret key for this widget. If invalidate_immediately is set to false, the previous secret remains valid for 2 hours.
Note that secrets cannot be rotated again during the grace period.
Domain types
typeWidgetstruct{…}
A Turnstile widget's detailed configuration
typeWidgetDomainstring
Hosts as a hostname or IPv4/IPv6 address represented by strings. The widget will only work on these domains, and their subdomains.