package repository_test

import (
	"database/sql"
	"testing"
	"timesheet/internal/model"
	"timesheet/internal/repository"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func Test_GetSummary_Input_Year_2018_Month_17_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			ID:                     1,
			MemberID:               "001",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			Company:                "siam_chamnankit",
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}
	databaseConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	year := 2017
	month := 12
	repository := repository.TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetSummary(year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
