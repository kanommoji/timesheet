package timesheet

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) []TransactionTimesheet
}

func (timesheet Timesheet) GetSummary(year, month int) []TransactionTimesheet {
	return []TransactionTimesheet{}
}
