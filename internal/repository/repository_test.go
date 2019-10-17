package repository_test

import (
	"database/sql"
	"testing"
	"timesheet/internal/model"
	. "timesheet/internal/repository"

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
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetSummary(year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_UpdateIncomeByID_Input_Year_2018_Month_12_MemberID_001_Income_Should_Be_No_Error(t *testing.T) {
	year := 2018
	month := 12
	memberID := "001"
	incomes := model.Incomes{
		Day:                      28,
		StartTimeAMHours:         9,
		StartTimeAMMinutes:       0,
		StartTimeAMSeconds:       0,
		EndTimeAMHours:           12,
		EndTimeAMMinutes:         0,
		EndTimeAMSeconds:         0,
		StartTimePMHours:         13,
		StartTimePMMinutes:       0,
		StartTimePMSeconds:       0,
		EndTimePMHours:           18,
		EndTimePMMinutes:         0,
		EndTimePMSeconds:         0,
		Overtime:                 0,
		TotalHoursHours:          8,
		TotalHoursMinutes:        0,
		TotalHoursSeconds:        0,
		CoachingCustomerCharging: 15000.00,
		CoachingPaymentRate:      10000.00,
		TrainingWage:             0.00,
		OtherWage:                0.00,
		Company:                  "siam_chamnankit",
		Description:              "[KBTG] 2 Days Agile Project Management",
	}
	databaseConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()

	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateIncomeByID(year, month, memberID, incomes)

	assert.Equal(t, nil, err)
}

func Test_GetMemberByID_Input_MemberID_001_Should_Be_Member(t *testing.T) {
	expected := []model.Member{
		{
			ID:                    1,
			MemberID:              "001",
			Company:               "siam_chamnankit",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			Status:                "",
			TravelExpense:         0.00,
		},
		{
			ID:                    2,
			MemberID:              "001",
			Company:               "shuhari",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                0.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			Status:                "",
			TravelExpense:         0.00,
		},
	}
	memberID := "001"
	databaseConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetMemberByID(memberID)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetIncomes_Input_MemberID_006_Year_2018_Month_12_Should_Be_Incomes_Day_11_And_12(t *testing.T) {
	expected := []model.Incomes{
		{
			ID:                       58,
			MemberID:                 "006",
			Month:                    12,
			Year:                     2018,
			Day:                      11,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "work at TN",
		}, {
			ID:                       59,
			MemberID:                 "006",
			Month:                    12,
			Year:                     2018,
			Day:                      12,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "work at TN",
		},
	}
	memberID := "006"
	month := 12
	year := 2018
	databaseConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetIncomes(memberID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_UpdateTransactionTimsheet_Input_Transaction_MemberID_006_Should_Be_No_Error(t *testing.T) {
	transactionTimesheet := []model.TransactionTimesheet{
		{
			MemberID:               "006",
			MemberNameTH:           "ภาณุมาศ แสนโท",
			Month:                  12,
			Year:                   2018,
			Company:                "shuhari",
			Coaching:               0.00,
			Training:               0.00,
			Other:                  6500.00,
			TotalIncomes:           6500.00,
			Salary:                 25000.00,
			IncomeTax1:             0.00,
			SocialSecurity:         750.00,
			NetSalary:              24250.00,
			Wage:                   6500.00,
			IncomeTax53Percentage:  5,
			IncomeTax53:            325.00,
			NetWage:                6175.00,
			NetTransfer:            30425.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}
	databaseConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateTransactionTimsheet(transactionTimesheet)

	assert.Equal(t, nil, err)
}

func Test_UpdateTimsheetByID_Input_Timsheet_MemberID_006_Should_Be_No_Error(t *testing.T) {
	timesheet := []model.Payment{
		{
			MemberID:                      "006",
			Month:                         12,
			Year:                          2019,
			TotalHoursHours:               0,
			TotalHoursMinutes:             0,
			TotalHoursSeconds:             0,
			TotalCoachingCustomerCharging: 0.00,
			TotalCoachingPaymentRate:      0.00,
			TotalTrainigWage:              10000.00,
			TotalOtherWage:                0.00,
			PaymentWage:                   10000.00,
		},
	}

	databaseConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateTimsheetByID(timesheet)

	assert.Equal(t, nil, err)
}
