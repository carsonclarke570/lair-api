package router

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/router/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"upper.io/db.v3/lib/sqlbuilder"
)

var rs = [](func(sqlbuilder.Database, *mux.Router)){
	routes.UserRoutes,
	routes.CharacterRoutes,
	routes.CampaignRoutes,
	routes.PlayerRoutes,
}

// CreateRouter creates a new router to handle request
func CreateRouter(sess sqlbuilder.Database) http.Handler {
	router := mux.NewRouter()
	for _, r := range rs {
		r(sess, router)
	}

	// Add middleware for CORS policy
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://lair-ui.herokuapp.com", "http://localhost:3000"},
		Debug:          true,
	})

	return middleware.Handler(router)
}
