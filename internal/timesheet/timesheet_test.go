package timesheet_test

import (
	"testing"
	"timesheet/internal/model"
	. "timesheet/internal/timesheet"

	"github.com/stretchr/testify/assert"
)

func Test_CalculatePaymentSummary_Input_Member_MemberID_001_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			Company:               "siam_chamnankit",
			Coaching:              0.00,
			Training:              155000.00,
			Other:                 0.00,
			TotalIncomes:          155000.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			NetSalary:             75000.00,
			Wage:                  75000.00,
			IncomeTax53Percentage: 10,
			IncomeTax53:           7500.00,
			NetWage:               67500.00,
			NetTransfer:           142500.00,
		},
		{
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			Company:               "shuhari",
			Coaching:              0.00,
			Training:              20000.00,
			Other:                 0.00,
			TotalIncomes:          20000.00,
			Salary:                0.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			NetSalary:             0.00,
			Wage:                  20000.00,
			IncomeTax53Percentage: 10,
			IncomeTax53:           2000.00,
			NetWage:               18000.00,
			NetTransfer:           18000.00,
		},
	}
	member := []model.Member{
		{
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
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      3,
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
			TrainingWage:             20000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(member, incomes, year, month)

	assert.Equal(t, expected, actual)
}

func Test_CalculatePaymentSummary_Input_Member_MemberID_001_Should_Be_Append_One_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			Company:               "siam_chamnankit",
			Coaching:              0.00,
			Training:              155000.00,
			Other:                 0.00,
			TotalIncomes:          155000.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			NetSalary:             75000.00,
			Wage:                  75000.00,
			IncomeTax53Percentage: 10,
			IncomeTax53:           7500.00,
			NetWage:               67500.00,
			NetTransfer:           142500.00,
		},
	}
	member := []model.Member{
		{
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
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(member, incomes, year, month)

	assert.Equal(t, expected, actual)
}
func Test_CalculatePayment_Input_Income_CoachingCustomerCharging_15000_CoachingPaymentRate_10000_TrainigWage_20000_Should_Be_Payment(t *testing.T) {
	expected := model.Payment{
		TotalHoursHours:               16,
		TotalHoursMinutes:             0,
		TotalHoursSeconds:             0,
		TotalCoachingCustomerCharging: 15000.00,
		TotalCoachingPaymentRate:      10000.00,
		TotalTrainigWage:              20000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   30000.00,
	}
	incomes := []model.Incomes{
		{
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
		},
		{
			Day:                      3,
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
			TrainingWage:             20000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	timesheet := Timesheet{}

	actual := timesheet.CalculatePayment(incomes)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalHour_Input_Incomes_Should_Be_Time_18_0_0(t *testing.T) {
	expected := model.Time{
		Hours:   18,
		Minutes: 0,
		Seconds: 0,
	}
	incomes := []model.Incomes{
		{
			StartTimeAMHours:   9,
			StartTimeAMMinutes: 0,
			StartTimeAMSeconds: 0,
			EndTimeAMHours:     12,
			EndTimeAMMinutes:   0,
			EndTimeAMSeconds:   0,
			StartTimePMHours:   13,
			StartTimePMMinutes: 0,
			StartTimePMSeconds: 0,
			EndTimePMHours:     18,
			EndTimePMMinutes:   30,
			EndTimePMSeconds:   0,
			Overtime:           0,
		},
		{
			StartTimeAMHours:   9,
			StartTimeAMMinutes: 0,
			StartTimeAMSeconds: 0,
			EndTimeAMHours:     12,
			EndTimeAMMinutes:   0,
			EndTimeAMSeconds:   0,
			StartTimePMHours:   13,
			StartTimePMMinutes: 0,
			StartTimePMSeconds: 0,
			EndTimePMHours:     18,
			EndTimePMMinutes:   30,
			EndTimePMSeconds:   0,
			Overtime:           1,
		},
	}

	actual := CalculateTotalHour(incomes)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetSalary_Input_Salary_80000_IncomeTax1_5000_SocialSecurity_0_Should_Be_75000(t *testing.T) {
	expected := 75000.00
	salary := 80000.00
	incomeTax1 := 5000.00
	socialSecurity := 0.00

	actual := CalculateNetSalary(salary, incomeTax1, socialSecurity)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetWage_Input_IncomeTax53Percentage_10_PaymentWage_155000_Salary_80000_Should_Be_67500(t *testing.T) {
	expected := 67500.00
	incomeTax53Percentage := 10
	salary := 80000.00
	paymentWage := 155000.00
	status := ""

	actual := CalculateNetWage(incomeTax53Percentage, paymentWage, salary, status)

	assert.Equal(t, expected, actual)
}

func Test_CalculateIncomeTax53_Input_IncomeTax53Percentage_10_Wage_40000_Should_Be_4000(t *testing.T) {
	expected := 4000.00
	incomeTax53Percentage := 10
	wage := 40000.00

	actual := CalculateIncomeTax53(wage, incomeTax53Percentage)

	assert.Equal(t, expected, actual)
}

func Test_CalculateWage_Input_PaymentWage_155000_Salary_80000_Should_Be_750000(t *testing.T) {
	expected := 75000.00
	paymentWage := 155000.00
	salary := 80000.00
	status := ""

	actual := CalculateWage(paymentWage, salary, status)

	assert.Equal(t, expected, actual)
}

func Test_CalculateWage_Input_PaymentWage_155000_Salary_80000_Status_Salary_Should_Be_155000(t *testing.T) {
	expected := 155000.00
	paymentWage := 155000.00
	salary := 80000.00
	status := "salary"

	actual := CalculateWage(paymentWage, salary, status)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetTransfer_Input_NetSalary_75000_NetWage_67500_Should_Be_142500(t *testing.T) {
	expected := 142500.00
	netSalary := 75000.00
	netWage := 67500.00

	actual := CalculateNetTransfer(netSalary, netWage)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalPaymentWage_Input_CoachingPaymentRate_85000_TrainingWage_70000_OtherWage_40000_Should_Be_195000(t *testing.T) {
	expected := 195000.00
	coachingPaymentRate := 85000.00
	trainingWage := 70000.00
	otherWage := 40000.00

	actual := CalculateTotalPaymentWage(coachingPaymentRate, trainingWage, otherWage)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingCustomerCharging_Input_Incomes_Should_Be_60000(t *testing.T) {
	expected := 60000.00
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
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
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
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
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := CalculateTotalCoachingCustomerCharging(incomes)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalOtherWage_Input_Incomes_Company_Shuhari_TravelExpense_0_Should_Be_45000(t *testing.T) {
	expected := 45000.00
	company := "shuhari"
	travelExpense := 0.00
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
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
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                25000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
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
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                20000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := CalculateTotalOtherWage(incomes, company, travelExpense)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalOtherWage_Input_Incomes_TravelExpense_0_Should_Be_55000(t *testing.T) {
	expected := 55000.00
	company := ""
	travelExpense := 0.00
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                10000.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
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
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                25000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
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
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                20000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := CalculateTotalOtherWage(incomes, company, travelExpense)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingPaymentRate_Input_Incomes_Company_SiamChamnankit_Should_Be_10000(t *testing.T) {
	expected := 10000.00
	company := "siam_chamnankit"
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
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
		},
		{
			Day:                      30,
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
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := CalculateTotalCoachingPaymentRate(incomes, company)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingPaymentRate_Input_Incomes_Should_Be_20000(t *testing.T) {
	expected := 20000.00
	company := ""
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
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
		},
		{
			Day:                      30,
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
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := CalculateTotalCoachingPaymentRate(incomes, company)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalTrainingWage_Input_Incomes_Company_Shuhari_Should_Be_20000(t *testing.T) {
	expected := 20000.00
	company := "shuhari"
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
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
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
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
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := CalculateTotalTrainingWage(incomes, company)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalTrainingWage_Input_Incomes_Should_Be_30000(t *testing.T) {
	expected := 30000.00
	company := ""
	incomes := []model.Incomes{
		{
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
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
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
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
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
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := CalculateTotalTrainingWage(incomes, company)

	assert.Equal(t, expected, actual)
}
