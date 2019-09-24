package models

import (
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Campaign represents a collection of players and a DM
type Campaign struct {
	Base         `json:",inline" db:",inline"`
	Name         string   `json:"name" db:"name"`
	GameMasterID uint     `json:"dm_id" db:"dm_id"`
	Image        string   `json:"img,omitempty" db:"img,omitempty"`
	Players      []Player `json:"players" db:"-"`
}

// GetBase from model.Model
func (c *Campaign) GetBase() *Base {
	return &c.Base
}

// SetID from model.Model
func (c *Campaign) SetID(id uint) {
	c.ID = id
}

// TableName from model.Model
func (*Campaign) TableName() string {
	return "campaigns"
}

// ReadChildren from model.Embedded
func (c *Campaign) ReadChildren(sess sqlbuilder.Database) error {
	children := make([]Player, 0)
	collection := sess.Collection("players")

	err := collection.Find(db.Cond{"campaign_id": c.ID}).All(&children)
	if err != nil {
		return err
	}
	c.Players = children

	return nil
}

// DeleteChildren from model.Embedded
func (c *Campaign) DeleteChildren(sess sqlbuilder.Database) error {
	collection := sess.Collection(c.TableName())

	return collection.Find(db.Cond{"campaign_id": c.ID}).Delete()
}
