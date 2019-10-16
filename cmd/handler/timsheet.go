package handler

import (
	"net/http"
	"timesheet/internal/model"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type CalculatePaymentRequest struct {
	MemberID string `json:"member_id"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
}

type IncomeRequest struct {
	Year     int           `json:"year"`
	Month    int           `json:"month"`
	MemberID string        `json:"member_id"`
	Incomes  model.Incomes `json:"incomes"`
}

type TimesheetAPI struct {
	Timesheet           timesheet.TimesheetGateways
	TimesheetRepository repository.TimesheetRepositoryGateways
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	var date Date
	err := context.ShouldBindJSON(&date)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	transactionTimesheet, _ := api.TimesheetRepository.GetSummary(date.Year, date.Month)
	context.JSON(http.StatusOK, transactionTimesheet)
}

func (api TimesheetAPI) UpdateIncomeHandler(context *gin.Context) {
	var requestIncome IncomeRequest
	err := context.ShouldBindJSON(&requestIncome)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, nil)
}

func (api TimesheetAPI) CalculatePaymentHandler(context *gin.Context) {

}
