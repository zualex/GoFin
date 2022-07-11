package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/zualex/gofin/pkg/config"
)

func NewDbConn(env string) (*sql.DB, error) {
	if env == config.EnvTest {
		return NewTestDbConn()
	}

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

func OpenTestConn() (*sql.DB, error) {
	dbConn, err := NewTestDbConn()
	if err != nil {
		return nil, fmt.Errorf("could not connect to test database: %w", err)
	}

	return dbConn, dbConn.Ping() // Подумать как оптимизировать чтобы вызывать только 1 раз
}

func TruncateTables(db *sql.DB, tables []string) {
	for _, v := range tables {
		_, _ = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", v))
	}
}
