package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/carsonclarke570/lair-api/pkg/models"
	"github.com/gorilla/mux"
	"upper.io/db.v3/lib/sqlbuilder"
)

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
	_, ok := model.(models.Embedded)
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
			resp.WriteHeader(http.StatusNoContent)
			return
		}

		if ok {
			err = m.Elem().Interface().(models.Embedded).ReadChildren(sess)
			if err != nil {
				http.Error(resp, err.Error(), 400)
				return
			}
		}

		resp.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(m.Elem().Interface().(models.Model))
	}
}

// Update updates all fields given in the body
func Update(sess sqlbuilder.Database, model models.Model) func(http.ResponseWriter, *http.Request) {
	m := reflect.New(reflect.TypeOf(model))
	return func(resp http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}

		col := sess.Collection(model.TableName())
		res := col.Find(id)
		err = res.One(m.Interface())
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}

		m2 := m.Elem().Interface()
		m2.(models.Model).GetBase().Modified = time.Now()
		json.NewDecoder(req.Body).Decode(m.Interface())
		err = res.Update(m2)
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}
	}
}

// Delete removes an entry in the database
func Delete(sess sqlbuilder.Database, model models.Model) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}

		col := sess.Collection(model.TableName())
		err = col.Find(id).Delete()
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}
	}
}

// Filter filters results
func Filter(sess sqlbuilder.Database, model models.Model) func(http.ResponseWriter, *http.Request) {
	m := reflect.New(reflect.SliceOf(reflect.TypeOf(model)))

	return func(resp http.ResponseWriter, req *http.Request) {

		pg := NewPagination()
		clauses := []string{}
		for key, values := range req.URL.Query() {
			ok, err := pg.parseQuery(key, values[0])
			if err != nil {
				http.Error(resp, err.Error(), 400)
				return
			}
			if !ok {
				clauses = append(clauses, fmt.Sprintf("%s = \"%s\"", key, values[0]))
			}
		}

		where := ""
		if len(clauses) > 0 {
			where = fmt.Sprintf("%s %s", "WHERE", strings.Join(clauses, " AND "))
		}

		query := fmt.Sprintf("SELECT * FROM %s %s ORDER BY %s %s LIMIT %d OFFSET %d;",
			model.TableName(), where, pg.OrderBy, pg.Order, pg.N, (pg.Page-1)*pg.N,
		)

		rows, err := sess.Query(query)
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}

		err = sqlbuilder.NewIterator(rows).All(m.Interface())
		if err != nil {
			http.Error(resp, err.Error(), 400)
			return
		}

		resp.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(m.Elem().Interface())
	}
}
