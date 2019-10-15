package main

import (
	"net/http"
	"timesheet/cmd/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	var api handler.TimesheetAPI
	router.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	router.POST("/addIncomeItem", api.UpdateIncomeHandler)
	router.StaticFS("/", http.Dir("ui"))
	router.Run()
}
