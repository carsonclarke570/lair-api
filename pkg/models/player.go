package models

// Player represents a player in a campaign
type Player struct {
	Base       `json:",inline" db:",inline"`
	UserID     uint `json:"user_id" db:"user_id"`
	CampaignID uint `json:"campaign_id" db:"campaign_id"`
}

// GetBase from model.Model
func (u *Player) GetBase() *Base {
	return &u.Base
}

// SetID from model.Model
func (u *Player) SetID(id uint) {
	u.ID = id
}

// TableName from model.Model
func (*Player) TableName() string {
	return "players"
}
