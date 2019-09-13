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

// AbilityScores represents a set of ability scores
type AbilityScores struct {
	STR int `json:"str" db:"str"`
	DEX int `json:"dex" db:"dex"`
	CON int `json:"con" db:"con"`
	INT int `json:"int" db:"int"`
	WIS int `json:"wis" db:"wis"`
	CHA int `json:"cha" db:"cha"`
}
