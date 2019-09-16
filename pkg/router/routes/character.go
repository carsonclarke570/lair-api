package routes

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/handlers"
	"github.com/carsonclarke570/lair-api/pkg/models"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

// CharacterRoutes routes characters
func CharacterRoutes(sess sqlbuilder.Database, router *mux.Router) {
	router.HandleFunc("/api/character", handlers.Create(sess, &models.Character{})).Methods(http.MethodPost)
	router.HandleFunc("/api/character/{id}", handlers.Read(sess, &models.Character{})).Methods(http.MethodGet)
	router.HandleFunc("/api/character/{id}", handlers.Update(sess, &models.Character{})).Methods(http.MethodPut)
	router.HandleFunc("/api/character/{id}", handlers.Delete(sess, &models.Character{})).Methods(http.MethodDelete)
	router.HandleFunc("/api/character", handlers.Filter(sess, &models.Character{})).Methods(http.MethodGet)
}
