package router

import "github.com/gorilla/mux"

// CreateRouter creates a new router to handle request
func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}
