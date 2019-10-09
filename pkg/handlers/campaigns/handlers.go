package handlers

import (
	"net/http"

	"upper.io/db.v3/lib/sqlbuilder"
)

// Players returns the list of players in the campaign
func Players(sess sqlbuilder.Database) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {

	}
}
