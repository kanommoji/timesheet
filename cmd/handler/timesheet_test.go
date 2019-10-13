package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	. "timesheet/cmd/handler"
	"timesheet/cmd/mockapi"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetSummaryHandler_Input_Year_2018_Month_12_Should_Be_Timesheet(t *testing.T) {
	expected := `[{"member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"siam_chamnankit","coaching":85000,"training":30000,"other":40000,"total_incomes":155000,"salary":80000,"income_tax_1":5000,"social_security":0,"net_salary":75000,"wage":75000,"income_tax_53_percentage":10,"income_tax_53":7500,"net_wage":67500,"net_transfer":142500,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":"","comment":""},{"member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"shuhari","coaching":0,"training":40000,"other":0,"total_incomes":40000,"salary":0,"income_tax_1":0,"social_security":0,"net_salary":0,"wage":40000,"income_tax_53_percentage":10,"income_tax_53":4000,"net_wage":36000,"net_transfer":36000,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":"","comment":""}]`
	date := Date{
		Year:  2018,
		Month: 12,
	}
	jsonRequest, _ := json.Marshal(date)
	request := httptest.NewRequest("POST", "/showSummaryTimesheet", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockTimesheet := new(mockapi.MockTimesheet)
	mockTimesheet.On("GetSummary", 2018, 12).Return([]timesheet.TransactionTimesheet{
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
	})

	api := TimesheetAPI{
		Timesheet: mockTimesheet,
	}

	testRoute := gin.Default()
	testRoute.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}
