package db

import (
	"reflect"
	"time"

	"github.com/carsonclarke570/lair-api/pkg/models"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Create adds a new Model to the database. Returns the id of the entry.
func Create(sess sqlbuilder.Database, model models.Model) (uint, error) {
	base := model.GetBase()
	base.ID = 0
	base.Created = time.Now()
	base.Modified = time.Now()

	col := sess.Collection(model.TableName())
	id, err := col.Insert(model)
	if err != nil {
		return 0, err
	}
	return uint(id.(int64)), nil
}

// Read reads a model from the database.
func Read(sess sqlbuilder.Database, model models.Model) (models.Model, error) {
	m := reflect.New(reflect.TypeOf(model))

	col := sess.Collection(model.TableName())
	res := col.Find(model.GetBase().ID)
	err := res.One(m.Interface())
	if err != nil {
		return nil, err
	}

	return m.Elem().Interface().(models.Model), nil
}

// Update updates an existing model in the database.
func Update(sess sqlbuilder.Database, model models.Model) error {
	return nil
}

// Delete deletes an entry in the database
func Delete(sess sqlbuilder.Database, model models.Model) error {
	col := sess.Collection(model.TableName())
	res := col.Find(model.GetBase().ID)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Filter filters results on the table
func Filter(sess sqlbuilder.Database, model models.Model) ([]models.Model, error) {
	return []models.Model{}, nil
}
