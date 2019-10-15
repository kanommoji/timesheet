package main

import (
	"database/sql"
	"log"
	"net/http"
	"timesheet/cmd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	if err != nil {
		log.Fatal("Cannot connect database", err.Error())
	}
	api := handler.TimesheetAPI{
		DBConnection: database,
	}
	router.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	router.POST("/addIncomeItem", api.UpdateIncomeHandler)
	router.StaticFS("/", http.Dir("ui"))
	router.Run()
}
