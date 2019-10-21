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
	var request Date
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	transactionTimesheet, _ := api.TimesheetRepository.GetSummary(request.Year, request.Month)
	context.JSON(http.StatusOK, transactionTimesheet)
}

func (api TimesheetAPI) CreateIncomeHandler(context *gin.Context) {
	var request IncomeRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = api.TimesheetRepository.CreateIncome(request.Year, request.Month, request.MemberID, request.Incomes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (api TimesheetAPI) CalculatePaymentHandler(context *gin.Context) {
	var request CalculatePaymentRequest
	err := context.BindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	incomes, err := api.TimesheetRepository.GetIncomes(request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	payments := api.Timesheet.CalculatePayment([]model.Incomes{})

	err = api.TimesheetRepository.VerifyTimesheet(payments, request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	members, err := api.TimesheetRepository.GetMemberByID(request.MemberID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	transactionTimesheet := api.Timesheet.CalculatePaymentSummary(members, incomes, request.Year, request.Month)

	err = api.TimesheetRepository.VerifyTransactionTimsheet(transactionTimesheet)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
