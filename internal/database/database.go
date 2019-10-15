package database

import "database/sql"

func DBConnect(url string) (*sql.DB, error) {
	database, err := sql.Open("mysql", url)
	if err != nil {
		return database, err
	}
	return database, nil
}
