package models

// User represents a user of the application
type User struct {
	Base      `json:",inline" db:",inline"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Hash      string `json:"-" db:"hash"`
}
