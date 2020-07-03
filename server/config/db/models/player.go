package models

import (
	"github.com/jinzhu/gorm"
	"github.com/moonrailgun/PUBGo/server/config/pubg/model"
	"time"
)

type ModelPlayer struct {
	gorm.Model
	AccountId        string `gorm:"not null;unique_index:index_account"`
	Name             string `gorm:"not null;unique_index:index_account"`
	ShardID          string
	AccountCreatedAt time.Time
	AccountUpdatedAt time.Time
}

func (m *ModelPlayer) ParseFromPUBG(data model.Player) {
	m.AccountId = data.ID
	m.Name = data.Name
	m.ShardID = data.ShardID
	m.AccountCreatedAt = data.CreatedAt
	m.AccountUpdatedAt = data.UpdatedAt
}
