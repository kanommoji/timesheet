package timesheet

type TimesheetGateways interface {
	GetSummary(year, month int) []TransactionTimesheet
}
