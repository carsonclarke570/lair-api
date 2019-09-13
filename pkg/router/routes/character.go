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
	router.HandleFunc("/api/character_sheet", handlers.Create(sess, &models.CharacterSheet{})).Methods(http.MethodPost)
	router.HandleFunc("/api/character_sheet/{id}", handlers.Read(sess, &models.CharacterSheet{})).Methods(http.MethodGet)
	router.HandleFunc("/api/character_sheet/{id}", handlers.Update(sess, &models.CharacterSheet{})).Methods(http.MethodPut)
	router.HandleFunc("/api/character_sheet/{id}", handlers.Delete(sess, &models.CharacterSheet{})).Methods(http.MethodDelete)
	router.HandleFunc("/api/character_sheet", handlers.Filter(sess, &models.CharacterSheet{})).Methods(http.MethodGet)
}
