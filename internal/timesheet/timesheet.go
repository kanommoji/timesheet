package timesheet

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type TimesheetGateways interface {
	UpdateIncomeByID(year, month int, memberID string) error
}

type Timesheet struct {
}

func (timesheet Timesheet) UpdateIncomeByID(year, month int, memberID string) error {
	return nil
}

func CalculateNetSalary(salary, incomeTax1, socialSecurity float64) float64 {
	return salary - incomeTax1 - socialSecurity
}

func CalculateNetWage(incomeTax53Percentag int, paymentWage, salary float64) float64 {

}
func Test_CalculateNetWage_Input_IncomeTax53Percentage_10_PaymentWage_155000_Salary_80000_Should_Be_67500(t *testing.T) {
	expected := 67500.00
	incomeTax53Percentag := 10
	salary := 80000.00
	paymentWage := 155000.00

	actual := CalculateNetWage(incomeTax53Percentag, paymentWage, salary)

	assert.Equal(t, expected, actual)
}
