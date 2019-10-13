package timesheet

type TimesheetGateways interface {
	GetSummary(year, month int) []TransactionTimesheet
}

type Timesheet struct {
}

func (timesheet Timesheet) GetSummary(year, month int) []TransactionTimesheet {
	return []TransactionTimesheet{}
}
