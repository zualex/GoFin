package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewDbConn() (*sql.DB, error) {
	db := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	dbConn, err := sql.Open("postgres", db)

	return dbConn, err
}

func NewTestDbConn() (*sql.DB, error) {
	db := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_TEST_NAME"))

	dbConn, err := sql.Open("postgres", db)

	return dbConn, err
}
