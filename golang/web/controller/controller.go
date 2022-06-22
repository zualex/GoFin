package controller

import "database/sql"

type Controller struct {
	Database *sql.DB
}
