package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	db, err = NewTestDbConn()
	if err != nil {
		return -1, fmt.Errorf("could not connect to test database: %w", err)
	}

	// truncates all test data after the tests are run
	defer func() {
		for _, t := range []string{"wallets"} {
			_, _ = db.Exec(fmt.Sprintf("DELETE FROM %s", t))
		}

		db.Close()
	}()

	return m.Run(), nil
}
