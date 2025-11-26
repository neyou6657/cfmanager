Abuse Reports
AbuseReports
Methods
client.AbuseReports.New(ctx, reportParam, params) (*string, error)
post/accounts/{account_id}/abuse-reports/{report_param}
Submit the Abuse Report of a particular type
client.AbuseReports.Get(ctx, reportParam, query) (*
AbuseReportGetResponse
, error)
get/accounts/{account_id}/abuse-reports/{report_param}
Retrieve the details of an abuse report.
client.AbuseReports.List(ctx, params) (*V4PagePagination[
AbuseReportListResponse
], error)
get/accounts/{account_id}/abuse-reports
List the abuse reports for a given account
Abuse Reports
Mitigations
AbuseReports.Mitigations
Methods
client.AbuseReports.Mitigations.List(ctx, reportID, params) (*V4PagePagination[
MitigationListResponse
], error)
get/accounts/{account_id}/abuse-reports/{report_id}/mitigations
List mitigations done to remediate the abuse report.
client.AbuseReports.Mitigations.Review(ctx, reportID, params) (*SinglePage[
MitigationReviewResponse
], error)
post/accounts/{account_id}/abuse-reports/{report_id}/mitigations/appeal
Request a review for mitigations on an account.