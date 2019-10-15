package timesheet

import "timesheet/internal/database"

type TimesheetRepositoryMariaDB struct {
	DataMapper database.DataMapperMariaDB
}
