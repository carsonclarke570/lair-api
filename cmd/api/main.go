package main

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := router.CreateRouter()
	log.Info("Starting HTTP server..")

	err := http.ListenAndServe(":80", router)
	log.WithError(err).Error("error serving requests")
}
