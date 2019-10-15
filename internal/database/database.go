package database

import "database/sql"

type DataMapperMariaDB struct {
	DBConnection *sql.DB
}
