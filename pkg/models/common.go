package models

import (
	"time"
)

// Base represnts fields common to all models
type Base struct {
	ID       uint      `json:"id" db:"id,omitempty"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}
