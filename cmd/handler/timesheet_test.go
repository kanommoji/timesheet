package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	. "timesheet/cmd/handler"
	"timesheet/cmd/mockapi"
	"timesheet/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetSummaryHandler_Input_Year_2018_Month_12_Should_Be_Timesheet(t *testing.T) {
	expected := `[{"id":1,"member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"siam_chamnankit","coaching":85000,"training":30000,"other":40000,"total_incomes":155000,"salary":80000,"income_tax_1":5000,"social_security":0,"net_salary":75000,"wage":75000,"income_tax_53_percentage":10,"income_tax_53":7500,"net_wage":67500,"net_transfer":142500,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":"","comment":""},{"id":3,"member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"shuhari","coaching":0,"training":40000,"other":0,"total_incomes":40000,"salary":0,"income_tax_1":0,"social_security":0,"net_salary":0,"wage":40000,"income_tax_53_percentage":10,"income_tax_53":4000,"net_wage":36000,"net_transfer":36000,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":"","comment":""}]`
	date := Date{
		Year:  2018,
		Month: 12,
	}
	jsonRequest, _ := json.Marshal(date)
	request := httptest.NewRequest("POST", "/showSummaryTimesheet", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockTimesheet := new(mockapi.MockTimesheet)
	mockTimesheet.On("GetSummary", 2018, 12).Return([]model.TransactionTimesheet{
		{
			ID:                     1,
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
			ID:                     3,
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
	}, nil)

	api := TimesheetAPI{
		TimesheetRepository: mockTimesheet,
	}

	testRoute := gin.Default()
	testRoute.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}

func Test_UpdateIncomeHandler_Input_Year_2018_Month_12_MemberID_001_Income_Should_Be_Status_200(t *testing.T) {
	expectedStatus := http.StatusOK
	expected := `{"year":2018,"month":12,"member_id":"001","incomes":[{"day":28,"start_time_am":"09:00:00","end_time_am":"12:00:00","start_time_pm":"13:00:00","end_time_pm":"18:00:00","overtime":0,"total_hours":8,"coaching_customer_charging":15000,"coaching_payment_rate":10000,"training_wage":0,"other_wage":0,"company":"siam_chamnankit","description":"[KBTG] 2 Days Agile Project Management"}]}`
	requestIncome := RequestIncome{
		Year:     2018,
		Month:    12,
		MemberID: "001",
		Incomes: []model.Incomes{
			{
				Day:                      28,
				StartTimeAM:              "09:00:00",
				EndTimeAM:                "12:00:00",
				StartTimePM:              "13:00:00",
				EndTimePM:                "18:00:00",
				Overtime:                 0,
				TotalHours:               8,
				CoachingCustomerCharging: 15000,
				CoachingPaymentRate:      10000,
				TrainingWage:             0,
				OtherWage:                0,
				Company:                  "siam_chamnankit",
				Description:              "[KBTG] 2 Days Agile Project Management",
			},
		},
	}
	jsonRequest, _ := json.Marshal(requestIncome)
	request := httptest.NewRequest("POST", "/addIncomeItem", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockTimesheet := new(mockapi.MockTimesheet)
	mockTimesheet.On("UpdateIncomeByID", 2018, 12, "001").Return(nil)

	api := TimesheetAPI{
		Timesheet: mockTimesheet,
	}

	testRoute := gin.Default()
	testRoute.POST("/addIncomeItem", api.UpdateIncomeHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()

	assert.Equal(t, expectedStatus, response.StatusCode)
	assert.Equal(t, expected, string(jsonRequest))
}
