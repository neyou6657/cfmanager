Audit Logs
AuditLogs
Methods
client.AuditLogs.List(ctx, params) (*V4PagePaginationArray[
AuditLog
], error)
get/accounts/{account_id}/audit_logs
Gets a list of audit logs for an account. Can be filtered by who made the change, on which zone, and the timeframe of the change.