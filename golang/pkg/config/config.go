package config

import (
	"database/sql"
)

type Config struct {
	Env        string
	Db         *sql.DB
	Routes     map[string]Route
	Currencies []string
}

const EnvTest = "test"

func New(env string, dbConn *sql.DB) *Config {
	return &Config{env, dbConn, routes, currencies}
}
