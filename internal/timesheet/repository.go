package timesheet

import (
	"database/sql"
)

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) ([]TransactionTimesheet, error)
}

type TimesheetRepository struct {
	DBConnection *sql.DB
}

func (repository TimesheetRepository) GetSummary(year, month int) ([]TransactionTimesheet, error) {
	var transactionTimesheetList []TransactionTimesheet
	var transactionTimesheet TransactionTimesheet
	statement, err := repository.DBConnection.Prepare(`SELECT * FROM timesheet.transactions WHERE year = ? AND month = ?`)
	if err != nil {
		return transactionTimesheetList, err
	}
	row, err := statement.Query(year, month)
	if err != nil {
		return transactionTimesheetList, err
	}
	for row.Next() {
		err = row.Scan(
			&transactionTimesheet.ID,
			&transactionTimesheet.MemberID,
			&transactionTimesheet.Month,
			&transactionTimesheet.Year,
			&transactionTimesheet.Company,
			&transactionTimesheet.MemberNameTH,
			&transactionTimesheet.Coaching,
			&transactionTimesheet.Training,
			&transactionTimesheet.Other,
			&transactionTimesheet.TotalIncomes,
			&transactionTimesheet.Salary,
			&transactionTimesheet.IncomeTax1,
			&transactionTimesheet.SocialSecurity,
			&transactionTimesheet.NetSalary,
			&transactionTimesheet.Wage,
			&transactionTimesheet.IncomeTax53Percentage,
			&transactionTimesheet.IncomeTax53,
			&transactionTimesheet.NetWage,
			&transactionTimesheet.NetTransfer,
			&transactionTimesheet.StatusCheckingTransfer,
			&transactionTimesheet.DateTransfer,
			&transactionTimesheet.Comment,
		)
		transactionTimesheetList = append(transactionTimesheetList, transactionTimesheet)
	}

	return transactionTimesheetList, nil
}
