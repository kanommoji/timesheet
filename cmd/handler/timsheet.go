package handler

import (
	"net/http"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type TimesheetAPI struct {
	Timesheet timesheet.TimesheetGateways
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	var date Date
	context.ShouldBindJSON(&date)

	// api.Timesheet.GetSummary(date.Year, date.Month)

	context.JSON(http.StatusOK, []timesheet.TransactionTimesheet{
		{
			MemberID:               "001",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2018,
			Company:                "Siam Chamnankit",
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
			Company:                "Shuhari",
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
	})
}
