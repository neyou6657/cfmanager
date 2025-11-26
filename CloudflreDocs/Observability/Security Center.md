Security Center
SecurityCenter
Security Center
Insights
SecurityCenter.Insights
Methods
client.SecurityCenter.Insights.List(ctx, params) (*V4PagePagination[
InsightListResponse
], error)
get/{accounts_or_zones}/{account_or_zone_id}/security-center/insights
Get Security Center Insights
client.SecurityCenter.Insights.Dismiss(ctx, issueID, params) (*
InsightDismissResponse
, error)
put/{accounts_or_zones}/{account_or_zone_id}/security-center/insights/{issue_id}/dismiss
Archive Security Center Insight
Security Center
Insights
Class
SecurityCenter.Insights.Class
Methods
client.SecurityCenter.Insights.Class.Get(ctx, params) (*[]
InsightClassGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/security-center/insights/class
Get Security Center Insight Counts by Class
Security Center
Insights
Severity
SecurityCenter.Insights.Severity
Methods
client.SecurityCenter.Insights.Severity.Get(ctx, params) (*[]
InsightSeverityGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/security-center/insights/severity
Get Security Center Insight Counts by Severity
Security Center
Insights
Type
SecurityCenter.Insights.Type
Methods
client.SecurityCenter.Insights.Type.Get(ctx, params) (*[]
InsightTypeGetResponse
, error)
get/{accounts_or_zones}/{account_or_zone_id}/security-center/insights/type
Get Security Center Insight Counts by Type