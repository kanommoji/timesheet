package main

import (
	"database/sql"
	"log"
	"net/http"
	"timesheet/cmd/handler"
	"timesheet/internal/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	dBConnection, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	if err != nil {
		log.Fatal("Cannot connect database", err.Error())
	}
	dataMapper := database.DataMapperMariaDB{
		DBConnection: dBConnection,
	}

	api := handler.TimesheetAPI{
		DataMapperMariaDB: dataMapper,
	}

	router.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	router.POST("/addIncomeItem", api.UpdateIncomeHandler)
	router.StaticFS("/", http.Dir("ui"))
	router.Run()
}
