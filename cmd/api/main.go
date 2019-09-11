package main

import (
	"net/http"

	"github.com/carsonclarke570/lair-api/pkg/db"
	"github.com/carsonclarke570/lair-api/pkg/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	// j := []byte(`{
	// 	"id": 25,
	// 	"first_name": "Johnny",
	// 	"last_name": "Test",
	// }`)

	// user := models.User{
	// 	Base: models.Base{
	// 		ID: 1,
	// 	},
	// 	FirstName: "Johnny",
	// 	LastName:  "Test",
	// 	Email:     "jihn@test.com",
	// }

	// json.Unmarshal(j, &user)
	// fmt.Println(user)

	// Base      `json:",inline" db:",inline"`
	// FirstName string `json:"first_name" db:"first_name"`
	// LastName  string `json:"last_name" db:"last_name"`
	// Email     string `json:"email" db:"email"`
	// Hash      string `json:"-" db:"hash"`

	log.Info("Connecting to database..")
	sess, err := db.InitDB()
	if err != nil {
		log.WithError(err).Fatal("failed to connect to database")
	}
	defer sess.Close()

	router := router.CreateRouter(sess)
	log.Info("Starting HTTP server..")

	err = http.ListenAndServe(":8001", router)
	log.WithError(err).Error("error serving requests")
}
