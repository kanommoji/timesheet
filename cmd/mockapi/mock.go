package mockapi

import (
	"timesheet/internal/timesheet"

	"github.com/stretchr/testify/mock"
)

type MockTimesheet struct {
	mock.Mock
}

func (mock MockTimesheet) GetSummary(year, month int) ([]timesheet.TransactionTimesheet, error) {
	argument := mock.Called(year, month)
	return argument.Get(0).([]timesheet.TransactionTimesheet), argument.Error(1)
}

func (mock MockTimesheet) UpdateIncomeByID(year, month int, memberID string, income []timesheet.Incomes) bool {
	argument := mock.Called(year, month, memberID)
	return argument.Bool(0)
}
