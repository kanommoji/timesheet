package timesheet

type TimesheetGateways interface {
	GetSummary(year, month int) []TransactionTimesheet
	UpdateIncomeByID(year, month int, memberID string) error
}

type Timesheet struct {
}

func (timesheet Timesheet) GetSummary(year, month int) []TransactionTimesheet {
	return []TransactionTimesheet{}
}

func (timesheet Timesheet) UpdateIncomeByID(year, month int, memberID string) error {
	return nil
}
