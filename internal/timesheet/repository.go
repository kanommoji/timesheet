package timesheet

import "database/sql"

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) []TransactionTimesheet
}

type TimesheetRepository struct {
	DBConnection *sql.DB
}

func (repository TimesheetRepository) GetSummary(year, month int) []TransactionTimesheet {
	return []TransactionTimesheet{}
}
