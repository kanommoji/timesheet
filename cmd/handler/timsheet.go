package handler

import (
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

}
