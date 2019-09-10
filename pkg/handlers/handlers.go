package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/carsonclarke570/lair-api/pkg/models"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

var ()

// Create creates a new entry in a database tabel
func Create(sess sqlbuilder.Database, model models.Model) func(http.ResponseWriter, *http.Request) {
	m := reflect.New(reflect.TypeOf(model))
	return func(resp http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(m.Interface())

		m2 := m.Elem().Interface().(models.Model)
		base := m2.GetBase()
		base.ID = 0
		base.Created = time.Now()
		base.Modified = time.Now()

		col := sess.Collection(model.TableName())
		id, err := col.Insert(m2)
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}
		resp.Header().Set("Location", fmt.Sprintf("%d", id))
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
