package mockapi

import (
	"github.com/stretchr/testify/mock"
)

type MockIncomeRepository struct {
	mock.Mock
}

func (mockIncomeRepository *MockIncomeRepository) UpdateIncomeByID(year, month int, memberID string) error {
	argument := mockIncomeRepository.Called(year, month, memberID)
	return argument.Error(0)
}
