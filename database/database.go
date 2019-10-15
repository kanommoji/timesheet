package database

import "database/sql"

func DBConnect(url string) (*sql.DB, error) {
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	if err != nil {
		return database, err
	}
	return database, nil
}
