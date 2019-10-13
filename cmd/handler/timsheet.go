package handler

import (
	"net/http"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type TimesheetAPI struct {
	Timesheet timesheet.TimesheetGateways
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	var date Date
	context.ShouldBindJSON(&date)

	transactionTimesheet := api.Timesheet.GetSummary(date.Year, date.Month)

	context.JSON(http.StatusOK, transactionTimesheet)
}
