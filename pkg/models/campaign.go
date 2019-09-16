package models

// Campaign represents a collection of players and a DM
type Campaign struct {
	Base         `json:",inline" db:",inline"`
	Name         string `json:"name" db:"name"`
	GameMasterID uint   `json:"dm_id" db:"dm_id"`
}

// GetBase from model.Model
func (u *Campaign) GetBase() *Base {
	return &u.Base
}

// SetID from model.Model
func (u *Campaign) SetID(id uint) {
	u.ID = id
}

// TableName from model.Model
func (*Campaign) TableName() string {
	return "campaigns"
}
