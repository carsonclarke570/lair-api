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
)

var (
	user = "root"
	pass = "lairpw"
	host = "db"
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

	settings := mysql.ConnectionURL{
		User:     user,
		Password: pass,
		Host:     host,
		Database: "lair",
	}

	sess, err := mysql.Open(settings)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
