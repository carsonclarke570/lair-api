package routes

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/handlers"
	"github.com/carsonclarke570/lair-api/pkg/models"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

// CampaignRoutes routes campaigns
func CampaignRoutes(sess sqlbuilder.Database, router *mux.Router) {
	router.HandleFunc("/api/campaign", handlers.Create(sess, &models.Campaign{})).Methods(http.MethodPost)
	router.HandleFunc("/api/campaign/{id}", handlers.Read(sess, &models.Campaign{})).Methods(http.MethodGet)
	router.HandleFunc("/api/campaign/{id}", handlers.Update(sess, &models.Campaign{})).Methods(http.MethodPut)
	router.HandleFunc("/api/campaign/{id}", handlers.Delete(sess, &models.Campaign{})).Methods(http.MethodDelete)
	router.HandleFunc("/api/campaign", handlers.Filter(sess, &models.Campaign{})).Methods(http.MethodGet)
}
