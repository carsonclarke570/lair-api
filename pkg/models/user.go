package models

// User represents a user of the application
type User struct {
	Base      `json:",inline" db:",inline"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Hash      string `json:"-" db:"hash"`
}

// GetBase from model.Model
func (u *User) GetBase() *Base {
	return &u.Base
}

func (u *User) SetID(id uint) {
	u.ID = id
}

// TableName from model.Model
func (*User) TableName() string {
	return "Users"
}

// FilterParams from model.Model
func (u *User) FilterParams() map[string]string {
	return map[string]string{}
}
