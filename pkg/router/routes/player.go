package routes

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/handlers"
	"github.com/carsonclarke570/lair-api/pkg/models"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

// PlayerRoutes routes players
func PlayerRoutes(sess sqlbuilder.Database, router *mux.Router) {
	router.HandleFunc("/api/player", handlers.Create(sess, &models.Player{})).Methods(http.MethodPost)
	router.HandleFunc("/api/player/{id}", handlers.Read(sess, &models.Player{})).Methods(http.MethodGet)
	router.HandleFunc("/api/player/{id}", handlers.Update(sess, &models.Player{})).Methods(http.MethodPut)
	router.HandleFunc("/api/player/{id}", handlers.Delete(sess, &models.Player{})).Methods(http.MethodDelete)
	router.HandleFunc("/api/player", handlers.Filter(sess, &models.Player{})).Methods(http.MethodGet)
}
