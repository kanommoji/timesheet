package timesheet_test

import (
	"testing"
	. "timesheet/internal/timesheet"

	"github.com/stretchr/testify/assert"
)

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
	incomeTax53Percentag := 10
	salary := 80000.00
	paymentWage := 155000.00

	actual := CalculateNetWage(incomeTax53Percentag, paymentWage, salary)

	assert.Equal(t, expected, actual)
}
