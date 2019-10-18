package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"timesheet/cmd/handler"
	"timesheet/config"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var databaseConfigPath string
	flag.StringVar(&databaseConfigPath, "config", "../timesheet/config/database.yml", "Database config path")
	flag.Parse()

	databaseConfig, err := config.SetupDatabaseConfig(databaseConfigPath)
	if err != nil {
		log.Fatal("Cannot read config", err.Error())
	}

	databaseConnection, err := sql.Open("mysql", databaseConfig.GetURI())
	if err != nil {
		log.Fatal("Cannot connect database", err.Error())
	}
	defer databaseConnection.Close()

	api := handler.TimesheetAPI{
		Timesheet: timesheet.Timesheet{},
		TimesheetRepository: repository.TimesheetRepository{
			DatabaseConnection: databaseConnection,
		},
	}

	router := gin.Default()

	router.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	router.POST("/addIncomeItem", api.UpdateIncomeHandler)
	router.POST("/calculatePayment", api.CalculatePaymentHandler)
	router.StaticFS("/", http.Dir("ui"))
	log.Fatal(router.Run())
}
