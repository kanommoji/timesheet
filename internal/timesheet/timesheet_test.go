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
