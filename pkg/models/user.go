package models

// User represents a user of the application
type User struct {
	Base  `json:",inline" db:",inline"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

// GetBase from model.Model
func (u *User) GetBase() *Base {
	return &u.Base
}

// SetID from model.Model
func (u *User) SetID(id uint) {
	u.ID = id
}

// TableName from model.Model
func (*User) TableName() string {
	return "users"
}
