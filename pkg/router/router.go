package router

import (
	"github.com/carsonclarke570/lair-api/pkg/router/routes"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

var rs = [](func(sqlbuilder.Database, *mux.Router)){
	routes.UserRoutes,
	routes.CharacterRoutes,
	routes.CampaignRoutes,
	routes.PlayerRoutes,
}

// CreateRouter creates a new router to handle request
func CreateRouter(sess sqlbuilder.Database) *mux.Router {
	router := mux.NewRouter()
	for _, r := range rs {
		r(sess, router)
	}
	return router
}
