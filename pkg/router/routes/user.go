package routes

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/handlers"
	"github.com/carsonclarke570/lair-api/pkg/models"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

// UserRoutes routes users
func UserRoutes(sess sqlbuilder.Database, router *mux.Router) {
	router.HandleFunc("/users", handlers.Create(sess, &models.User{})).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlers.Read(sess, &models.User{})).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.Update(sess, &models.User{})).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.Delete(sess, &models.User{})).Methods(http.MethodDelete)
	router.HandleFunc("/users", handlers.Filter(sess, &models.User{})).Methods(http.MethodGet)
}
