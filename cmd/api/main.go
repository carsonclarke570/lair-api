package main

import (
	"fmt"
	"net/http"
	"os"

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
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "80"
	}
	log.Info(fmt.Sprintf("Starting HTTP server on port %s..", port))
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	log.WithError(err).Error("error serving requests")
}
