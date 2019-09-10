package models

//Model represents a common interface for models
type Model interface {

	// GetBase returns the Base of the model
	GetBase() *Base

	// SetID sets the ID of the model
	SetID(id uint)

	// TableName gives the name of the database table for the model
	TableName() string

	// // FilterParams gives a map of all params
	// FilterParams() map[string]string
}
