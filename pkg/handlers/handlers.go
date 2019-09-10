package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/carsonclarke570/lair-api/pkg/models"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

var ()

func Create(sess sqlbuilder.Database, model models.Model) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {

	}
}

// Read reads a model from the database
func Read(sess sqlbuilder.Database, model models.Model) func(http.ResponseWriter, *http.Request) {
	m := reflect.New(reflect.TypeOf(model))
	return func(resp http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}

		col := sess.Collection(model.TableName())
		err = col.Find(id).One(m.Interface())
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}

		resp.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(m.Elem().Interface().(models.Model))
	}
}
