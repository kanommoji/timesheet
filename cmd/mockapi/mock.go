package mockapi

import (
	"timesheet/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockTimesheet struct {
	mock.Mock
}

func (mock MockTimesheet) GetSummary(year, month int) ([]model.TransactionTimesheet, error) {
	argument := mock.Called(year, month)
	return argument.Get(0).([]model.TransactionTimesheet), argument.Error(1)
}

func (mock MockTimesheet) UpdateIncomeByID(year, month int, memberID string) error {
	argument := mock.Called(year, month, memberID)
	return argument.Error(0)
}
