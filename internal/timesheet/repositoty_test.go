package timesheet_test

import (
	"database/sql"
	"testing"
	"timesheet/internal/timesheet"

	"github.com/stretchr/testify/assert"
)

func Test_TimesheetRepository(t *testing.T) {
	dBConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")

	t.Run("Test_GetSummary_Input_Year_2018_Month_12_Should_Be_TransactionTimesheet", func(t *testing.T) {
		expected := []timesheet.TransactionTimesheet{
			{
				MemberID:               "001",
				MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
				Month:                  12,
				Year:                   2018,
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
			}, {
				MemberID:               "001",
				MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
				Month:                  12,
				Year:                   2018,
				Company:                "shuhari",
				Coaching:               0.00,
				Training:               40000.00,
				Other:                  0.00,
				TotalIncomes:           40000.00,
				Salary:                 0.00,
				IncomeTax1:             0.00,
				SocialSecurity:         0.00,
				NetSalary:              0.00,
				Wage:                   40000.00,
				IncomeTax53Percentage:  10,
				IncomeTax53:            4000.00,
				NetWage:                36000.00,
				NetTransfer:            36000.00,
				StatusCheckingTransfer: "รอการตรวจสอบ",
				DateTransfer:           "",
				Comment:                "",
			},
		}
		year := 2018
		month := 12
		repository := timesheet.TimesheetRepository{
			DBConnection: dBConnection,
		}
		actual := repository.GetSummary(year, month)

		assert.Equal(t, expected, actual)
	})

}
