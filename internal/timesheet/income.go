package timesheet

type IncomeGateways interface {
	UpdateIncomeByID(year, month int, memberID string) error
}

type Income struct {
}

func (income Income) UpdateIncomeByID(year, month int, memberID string) error {
	return nil
}
