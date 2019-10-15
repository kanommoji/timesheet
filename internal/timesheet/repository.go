package timesheet

import (
	"database/sql"
)

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) ([]TransactionTimesheet, error)
	UpdateIncomeByID(year, month int, memberID string, income []Incomes) bool
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

func (repository TimesheetRepository) UpdateIncomeByID(year, month int, memberID string, income []Incomes) bool {
	statement, err := repository.DBConnection.Prepare(`INSERT INTO incomes (member_id, month, year, day, start_time_am, end_time_am, start_time_pm, end_time_pm, overtime, total_hours, coaching_customer_charging, coaching_payment_rate, training_wage, other_wage, company, description) VALUES ( ? , ? , ?, ? , ? , ?, ? , ? , ?, ? , ? , ?, ? , ? , ?, ? )`)
	if err != nil {
		return false
	}
	_, err = statement.Exec(memberID, month, year, income[0].Day, income[0].StartTimeAM, income[0].EndTimeAM, income[0].StartTimePM, income[0].EndTimePM, income[0].Overtime, income[0].TotalHours, income[0].CoachingCustomerCharging, income[0].CoachingPaymentRate, income[0].TrainingWage, income[0].OtherWage, income[0].Company, income[0].Description)
	if err != nil {
		return false
	}
	return true
}
