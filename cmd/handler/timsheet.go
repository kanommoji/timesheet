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

type RequestIncome struct {
	Year     int       `json:"year"`
	Month    int       `json:"month"`
	MemberID string    `json:"member_id"`
	Incomes  []Incomes `json:"incomes"`
}

type Incomes struct {
	Day                      int    `json:"day"`
	StartTimeAM              string `json:"start_time_am"`
	EndTimeAM                string `json:"end_time_am"`
	StartTimePM              string `json:"start_time_pm"`
	EndTimePM                string `json:"end_time_pm"`
	Overtime                 int    `json:"overtime"`
	TotalHours               int    `json:"total_hours"`
	CoachingCustomerCharging int    `json:"coaching_customer_charging"`
	CoachingPaymentRate      int    `json:"coaching_payment_rate"`
	TrainingWage             int    `json:"training_wage"`
	OtherWage                int    `json:"other_wage"`
	Company                  string `json:"company"`
	Description              string `json:"description"`
}

type TimesheetAPI struct {
	Timesheet timesheet.TimesheetGateways
	Income    timesheet.IncomeGateways
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	var date Date
	context.ShouldBindJSON(&date)

	// transactionTimesheet := api.Timesheet.GetSummary(date.Year, date.Month)

	context.JSON(http.StatusOK, []timesheet.TransactionTimesheet{
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
	},
	)
}

func (api TimesheetAPI) UpdateIncomeHandler(context *gin.Context) {
	var requestIncome RequestIncome
	context.ShouldBindJSON(&requestIncome)

	context.JSON(http.StatusOK, "")
}
