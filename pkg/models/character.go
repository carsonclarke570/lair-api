package models

// Character represents a DnD 5e character sheet
type Character struct {
	Base      `json:",inline" db:",inline"`
	Name      string `json:"name" db:"name"`
	Image     string `json:"img" db:"img"`
	Race      string `json:"race" db:"race"`
	Level     uint   `json:"level" db:"level"`
	Size      string `json:"size" db:"size"`
	Alignment string `json:"alignment" db:"alignment"`

	ArmorClass uint   `json:"ac" db:"ac"`
	Armor      string `json:"armor" db:"armor"`
	HitPoints  uint   `json:"hit_points" db:"hit_points"`
	HitDie     string `json:"hit_die" db:"hit_die"`
	Speed      uint   `json:"speed" db:"speed"`
	Initiative uint   `json:"initiative" db:"initiative"`

	// TO-DO: Make multi-classing
	// Class     string `json:"class" db:"class"`

	AbilityScores `json:",inline" db:",inline"`

	// TO-DO: Saving throws
	// TO-DO: Skills

	// TO-DO: Vulnerabilities, Resistances, Immunities, Senses, Passive Perception, Languages
	PlayerID uint `json:"player_id" db:"player_id"`
}

// GetBase from model.Model
func (c *Character) GetBase() *Base {
	return &c.Base
}

// SetID from model.Model
func (c *Character) SetID(id uint) {
	c.ID = id
}

// TableName from model.Model
func (*Character) TableName() string {
	return "characters"
}
