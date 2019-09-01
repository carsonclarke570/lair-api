package models

// Base represnts fields common to all models
type Base struct {
	ID   uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// AbilitySet represents a set of ability scores
type AbilitySet struct {
	Strength     uint `json:"str" db:"str"`
	Dexterity    uint `json:"dex" db:"dex"`
	Constitution uint `json:"con" db:"con"`
	Intelligence uint `json:"int" db:"int"`
	Wisdom       uint `json:"wis" db:"wis"`
	Charisma     uint `json:"cha" db:"cha"`
}

// SavingSet represents a set of saving throws
type SavingSet struct {
	Strength     uint `json:"str" db:"str_save"`
	Dexterity    uint `json:"dex" db:"dex_save"`
	Constitution uint `json:"con" db:"con_save"`
	Intelligence uint `json:"int" db:"int_save"`
	Wisdom       uint `json:"wis" db:"wis_save"`
	Charisma     uint `json:"cha" db:"cha_save"`
}
