package timesheet

type TimesheetGateways interface {
	GetSummary(year, month int) []TransactionTimesheet
	UpdateIncomeByID(year, month int, memberID string) error
}

func (timesheet Timesheet) UpdateIncomeByID(year, month int, memberID string) error {
	return nil
}
