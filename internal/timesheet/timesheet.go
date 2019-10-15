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
