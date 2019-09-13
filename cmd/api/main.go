package main

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/db"
	"github.com/carsonclarke570/lair-api/pkg/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Connecting to database..")
	sess, err := db.InitDB()
	if err != nil {
		log.WithError(err).Fatal("failed to connect to database")
	}
	defer sess.Close()

	router := router.CreateRouter(sess)
	log.Info("Starting HTTP server..")

	err = http.ListenAndServe(":80", router)
	log.WithError(err).Error("error serving requests")
}
