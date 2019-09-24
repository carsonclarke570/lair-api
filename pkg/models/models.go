package models

import "upper.io/db.v3/lib/sqlbuilder"

//Model represents a common interface for models
type Model interface {

	// GetBase returns the Base of the model
	GetBase() *Base

	// SetID sets the ID of the model
	SetID(id uint)

	// TableName gives the name of the database table for the model
	TableName() string
}

// Embedded represents a model that has children models
type Embedded interface {
	ReadChildren(sess sqlbuilder.Database) error
	DeleteChildren(sess sqlbuilder.Database) error
}
