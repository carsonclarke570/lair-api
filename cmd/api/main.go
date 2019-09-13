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
	log.Info("Starting HTTP server..")

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "80"
	}
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	log.WithError(err).Error("error serving requests")
}
