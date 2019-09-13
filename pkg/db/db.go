package db

import (
	"os"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

const (
	envUser = "DB_USER"
	envPass = "DB_PASS"
	envHost = "DB_HOST"
	envData = "DB_DATA"
)

var (
	user = "root"
	pass = "lairpw"
	host = "localhost"
	db   = "lair"
)

// InitDB initializes the database connection
func InitDB() (sqlbuilder.Database, error) {
	if u := os.Getenv(envUser); u != "" {
		user = u
	}
	if p := os.Getenv(envPass); p != "" {
		pass = p
	}
	if h := os.Getenv(envHost); h != "" {
		host = h
	}
	if d := os.Getenv(envData); d != "" {
		db = d
	}

	settings := mysql.ConnectionURL{
		User:     user,
		Password: pass,
		Host:     host,
		Database: db,
	}

	sess, err := mysql.Open(settings)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
