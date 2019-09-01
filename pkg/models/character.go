package models

// Character represents a character sheet
type Character struct {
	Base       `json:",inline" db:",inline"`
	Class      string `json:"class" db:"class"`
	Race       string `json:"race" db:"race"`
	Alignment  string `json:"alignment" db:"alignment"`
	Level      uint   `json:"level" db:"level"`
	Experience uint   `json:"exp" db:"exp"`
	ArmorClass uint   `json:"ac" db:"ac"`
	// AbilitySet `json:"abilities" db:",inline"`
	// SavingSet  `json:"saving" db:",inline"`
}

// Attack represensts an attack
type Attack struct {
	Base `json:",inline" db:",inline"`
}
