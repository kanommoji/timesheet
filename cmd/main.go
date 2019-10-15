package main

import (
	"database/sql"
	"log"
	"net/http"
	"timesheet/cmd/handler"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	if err != nil {
		log.Fatal("Cannot connect database", err.Error())
	}
	defer database.Close()

	api := handler.TimesheetAPI{
		TimesheetRepository: timesheet.TimesheetRepository{
			DBConnection: database,
		},
	}
	router.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	router.POST("/addIncomeItem", api.UpdateIncomeHandler)
	router.StaticFS("/", http.Dir("ui"))
	router.Run()
}
