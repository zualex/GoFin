package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDbConn() (*sql.DB, error) {
	db := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	log.Println("DB_HOST: " + os.Getenv("DB_HOST"))
	log.Println("DB_USER: " + os.Getenv("DB_USER"))
	log.Println("DB_PASSWORD: " + os.Getenv("DB_PASSWORD"))
	log.Println("DB_NAME: " + os.Getenv("DB_NAME"))

	dbConn, err := sql.Open("postgres", db)

	return dbConn, err
}
