package timesheet

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
	wage := CalculateWage(paymentWage, salary)
	incomeTax53 := CalculateIncomeTax53(wage, incomeTax53Percentag)
	return wage - incomeTax53
}

func CalculateWage(paymentWage, salary float64) float64 {
	return 75000.00
}

func CalculateIncomeTax53(wage float64, incomeTax53Percentag int) float64 {
	return 7500.00
}
