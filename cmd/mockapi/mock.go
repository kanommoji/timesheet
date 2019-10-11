package mockapi

import (
	"timesheet/internal/timesheet"

	"github.com/stretchr/testify/mock"
)

type MockTimesheet struct {
	mock.Mock
}

func (mock MockTimesheet) GetSummary(year, month int) []timesheet.TransactionTimesheet {
	argument := mock.Called(year, month)
	return argument.Get(0).([]timesheet.TransactionTimesheet)
}
